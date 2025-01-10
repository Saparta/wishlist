package endpoints

import (
	"context"

	pb "github.com/Saparta/wishlist/wishlist/services/wishlist-service/proto"
	"github.com/Saparta/wishlist/wishlist/services/wishlist-service/shared"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (w *WishlistService) DeleteWishlistItem(ctx context.Context, request *pb.DeleteWishlistItemRequest) (*pb.DeleteWishlistItemResponse, error) {
	dbPool, err := shared.ConnectToDatabase(ctx)
	if err != nil {
		return nil, err
	}
	_, err = dbPool.Query(ctx,
		`
		WITH authorized_users AS (
    SELECT i.id AS item_id
    FROM items i
    JOIN wishlists w ON i.wishlist_id = w.id
    LEFT JOIN shared s ON w.id = s.wishlist_id
    WHERE i.id = $1
      AND (w.user_id = $2 OR (s.shared_with = $2 AND s.can_edit = TRUE))
		)
		DELETE FROM items 
		WHERE id IN (SELECT item_id FROM authorized_users);
		`, request.ItemId, request.UserId)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error executing query: %s", err.Error())
	}
	return &pb.DeleteWishlistItemResponse{}, nil
}
