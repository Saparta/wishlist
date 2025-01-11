package endpoints

import (
	"context"

	pb "github.com/Saparta/wishlist/wishlist/services/wishlist-service/proto"
	"github.com/Saparta/wishlist/wishlist/services/wishlist-service/shared"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (w *WishlistService) ClearWishlistItems(ctx context.Context, request *pb.ClearWishlistItemsRequest) (*pb.ClearWishlistItemsResponse, error) {
	dbPool, err := shared.ConnectToDatabase(ctx)
	if err != nil {
		return nil, err
	}

	query := `
	WITH deleted_items AS (
    DELETE FROM items
    WHERE wishlist_id = $1
      AND wishlist_id IN (
          SELECT wishlist_id 
					FROM shared 
					WHERE wishlist_id = $1 AND user_id = $2 AND is_owner = TRUE
      )
    RETURNING wishlist_id
	)
	UPDATE shared
	SET last_modified = CURRENT_TIMESTAMP
	WHERE wishlist_id in (SELECT * FROM deleted_items) AND user_id = $2;
`
	rows, err := dbPool.Query(ctx, query, request.WishlistId, request.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer rows.Close()
	return &pb.ClearWishlistItemsResponse{}, nil
}
