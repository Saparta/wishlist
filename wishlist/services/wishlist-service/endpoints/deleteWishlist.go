package endpoints

import (
	"context"

	pb "github.com/Saparta/wishlist/wishlist/services/wishlist-service/proto"
	"github.com/Saparta/wishlist/wishlist/services/wishlist-service/shared"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (w *WishlistService) DeleteWishlist(ctx context.Context, request *pb.DeleteWishlistRequest) (*pb.DeleteWishlistResponse, error) {
	dbPool, err := shared.ConnectToDatabase(ctx)
	if err != nil {
		return nil, err
	}

	rows, err := dbPool.Query(ctx, `
	WITH authorized_user AS (
    SELECT w.id
    FROM shared s
    JOIN wishlists w ON w.id = s.wishlist_id
    WHERE w.id = $1 AND s.user_id = $2 AND s.is_owner = TRUE
	)
	DELETE FROM wishlists
	WHERE id IN (SELECT id FROM authorized_user);`, request.WishlistId, request.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer rows.Close()
	return &pb.DeleteWishlistResponse{}, nil
}
