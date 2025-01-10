package endpoints

import (
	"context"
	"time"

	pb "github.com/Saparta/wishlist/wishlist/services/wishlist-service/proto"
	"github.com/Saparta/wishlist/wishlist/services/wishlist-service/shared"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (w *WishlistService) GetWishlistsSharedWithUser(ctx context.Context, request *pb.GetWishlistsSharedWithUserRequest) (*pb.GetWishlistsSharedWithUserResponse, error) {
	dbPool, err := shared.ConnectToDatabase(ctx)
	if err != nil {
		return nil, err
	}

	rows, err := dbPool.Query(ctx, `
	SELECT
			sw.id AS wishlist_id, 
			sw.title AS wishlist_title, 
			sw.description AS wishlist_description,
			sw.is_public AS is_public,
			sw.last_opened AS wishlist_last_opened,
			sw.last_modified AS wishlist_last_modified,
			i.id AS item_id,
			i.name AS item_name,
			i.url AS item_url,
			i.price AS item_price ,
			i.is_gifted AS item_is_gifted,
			i.gifted_by AS item_gifted_by,
			i.created_at AS item_created_at
	FROM
		(SELECT
				s.wishlist_id as id,
				s.shared_with as user_id,
				w.title,
				w.description,
				w.is_public,
				s.can_edit,
				s.last_opened,
				s.last_modified
		FROM
				shared s JOIN wishlists w ON s.wishlist_id = w.id
		WHERE s.shared_with = $1 AND (w.is_public = TRUE OR s.can_edit = TRUE)) sw 
		LEFT JOIN items i ON sw.id = i.wishlist_id
		ORDER BY sw.last_opened;
	;
	`, request.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failure to query wishlists: %v", err)
	}
	wishMap := make(map[string]*pb.Wishlist)
	var wishlistId, title, description, name, url, giftedBy, itemId *string
	var price *float32
	var isPublic, isGifted *bool
	var lastOpened, lastModified, createdAt *time.Time
	_, err = pgx.ForEachRow(rows, []any{&wishlistId, &title, &description, &isPublic, &lastOpened, &lastModified, &itemId, &name, &url, &price, &isGifted, &giftedBy, &createdAt},
		func() error {
			val, found := wishMap[*wishlistId]
			var allItems []*pb.WishlistItem

			if itemId != nil {
				allItems = append(allItems,
					&pb.WishlistItem{
						Id:         itemId,
						Name:       name,
						Url:        url,
						Price:      price,
						IsGifted:   isGifted,
						GiftedBy:   giftedBy,
						CreatedAt:  timestamppb.New(*createdAt),
						WishlistId: wishlistId,
					})
			}

			if !found {
				wishMap[*wishlistId] = &pb.Wishlist{
					Id:           wishlistId,
					UserId:       request.UserId,
					Title:        title,
					Description:  description,
					IsPublic:     isPublic,
					CanEdit:      func() *bool { b := true; return &b }(), // This is just a pointer to True
					LastOpened:   timestamppb.New(*lastOpened),
					LastModified: timestamppb.New(*lastModified),
					Items:        allItems,
				}
			} else {
				val.Items = append(val.Items, allItems...)
			}
			return nil
		})

	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to process rows")
	}

	wishlistIds := make([]string, 0, len(wishMap))
	for key := range wishMap {
		wishlistIds = append(wishlistIds, key)
	}
	var getSharedWith = make(chan map[string][]string, 1)
	var errChannel chan error = make(chan error, 1)
	go shared.BatchGetWishlistsSharedUsers(ctx, wishlistIds, getSharedWith, errChannel)

	select {
	case sharedWithMap := <-getSharedWith:
		for id, sharedWith := range sharedWithMap {
			wishlist, found := wishMap[id]
			if found {
				wishlist.SharedWith = sharedWith
			}
		}
	case err := <-errChannel:
		return nil, err
	}

	var wishlists []*pb.Wishlist
	for _, value := range wishMap {
		wishlists = append(wishlists, value)
	}

	// Return the wishlists response
	return &pb.GetWishlistsSharedWithUserResponse{
		Wishlists: wishlists,
	}, nil
}
