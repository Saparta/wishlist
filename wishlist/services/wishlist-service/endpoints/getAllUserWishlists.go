package endpoints

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/Saparta/wishlist/wishlist/services/wishlist-service/proto"
	"github.com/Saparta/wishlist/wishlist/services/wishlist-service/shared"
	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (w *WishlistService) GetAllUserWishlists(ctx context.Context, request *pb.GetAllUserWishlistsRequest) (*pb.GetAllUserWishlistsResponse, error) {
	dbPool, err := shared.ConnectToDatabase(ctx)
	if err != nil {
		return nil, err
	}

	rows, err := dbPool.Query(ctx, `
    SELECT
			w.id AS wishlist_id, 
			w.title AS wishlist_title, 
			w.description AS wishlist_description,
			w.is_public AS is_public,
			w.last_opened AS wishlist_last_opened,
			w.last_modified AS wishlist_last_modified,
			i.id AS item_id,
			i.name AS item_name,
			i.url AS item_url,
			i.price AS item_price ,
			i.is_gifted AS item_is_gifted,
			i.gifted_by AS item_gifted_by,
			i.created_at AS item_created_at
    FROM wishlists w
		LEFT JOIN items i
		ON w.id = i.wishlist_id
    WHERE w.user_id = $1;
    `, request.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to query wishlists: %v", err.Error())
	}

	var wishlistId, title, description, name, url, giftedBy, itemId *string
	var price *float32
	var isPublic, isGifted *bool
	var lastOpened, lastModified, createdAt *time.Time

	wishMap := make(map[string]*pb.Wishlist)
	_, err = pgx.ForEachRow(rows, []any{&wishlistId, &title, &description, &isPublic, &lastOpened, &lastModified, &itemId, &name, &url, &price, &isGifted, &giftedBy, &createdAt}, func() error {
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
				CanEdit:      func() *bool { b := true; return &b }(), // This is just a pointer to true
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

	var wishlists []*pb.Wishlist
	for _, value := range wishMap {
		wishlists = append(wishlists, value)
	}

	// Return the wishlists response
	return &pb.GetAllUserWishlistsResponse{
		Wishlists: wishlists,
	}, nil
}
