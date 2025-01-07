package endpoints

import (
	"context"
	"time"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	pb "github.com/Saparta/wishlist/wishlist/services/wishlist-service/proto"
	"github.com/Saparta/wishlist/wishlist/services/wishlist-service/shared"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (w *WishlistService) GetUserWishlists(ctx context.Context, request *pb.GetUserWishlistsRequest) (*pb.GetUserWishlistsResponse, error) {
	dbPool, ok := ctx.Value(shared.DBSession).(*pgxpool.Pool)
	if !ok {
		return nil, status.Error(codes.Internal, "Failed to retrieve database connection from context")
	}

	userID := request.GetUserId()
	if userID == "" {
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
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
			i.gifted_by AS item_gifted_by
			i.created_at AS item_created_at
    FROM wishlists w
		LEFT JOIN items i
		ON w.id = i.wishlist_id
    WHERE w.user_id = $1
    `, userID)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to query wishlists: "+err.Error())
	}

	wishMap := make(map[string]*pb.Wishlist)
	var wishlistId, title, description, name, url, giftedBy, itemId *string
	var price *float32
	var isPublic, isGifted *bool
	var lastOpened, lastModified, createdAt *time.Time
	_, err = pgx.ForEachRow(rows, []any{&wishlistId, &title, &description, &isPublic, &lastOpened, &lastModified, &itemId, &name, &url, &price, &isGifted, &giftedBy, &createdAt}, func() error {
		val, found := wishMap[*wishlistId]
		var allItems []*pb.WishlistItem

		if itemId != nil {
			allItems = append(allItems,
				&pb.WishlistItem{
					Id:        *itemId,
					Name:      *name,
					Url:       *url,
					Price:     *price,
					IsGifted:  *isGifted,
					GiftedBy:  *giftedBy,
					CreatedAt: timestamppb.New(*createdAt),
				})
		}

		if !found {
			wishMap[*wishlistId] = &pb.Wishlist{
				Id:           *wishlistId,
				UserId:       userID,
				Title:        *title,
				Description:  *description,
				IsPublic:     *isPublic,
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
	return &pb.GetUserWishlistsResponse{
		Wishlists: wishlists,
	}, nil
}
