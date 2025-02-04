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
  	SELECT i.id AS item_id, w.id as wishlist_id
  	FROM items i
  	JOIN wishlists w ON i.wishlist_id = w.id
  	LEFT JOIN shared s ON w.id = s.wishlist_id
  	WHERE i.id = $1
    	AND (s.is_owner = TRUE OR (s.user_id = $2 AND s.can_edit = TRUE))
	),
	updated_items AS (
		DELETE FROM items 
		WHERE id IN (SELECT item_id FROM authorized_users)
		RETURNING *
	),
	update_shared AS (
    UPDATE shared
    SET last_modified = CURRENT_TIMESTAMP
    WHERE wishlist_id IN (SELECT wishlist_id FROM authorized_users)
    AND user_id = $2
	)
	SELECT * FROM updated_items;
		`, request.ItemId, request.UserId)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error executing query: %s", err.Error())
	}
	return &pb.DeleteWishlistItemResponse{}, nil
}
