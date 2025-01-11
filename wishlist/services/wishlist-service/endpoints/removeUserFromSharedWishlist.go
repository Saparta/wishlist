package endpoints

import (
	"context"

	pb "github.com/Saparta/wishlist/wishlist/services/wishlist-service/proto"
	"github.com/Saparta/wishlist/wishlist/services/wishlist-service/shared"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (w *WishlistService) RemoveUserFromSharedWishlist(ctx context.Context, request *pb.RemoveUserFromSharedWishlistRequest) (*pb.RemoveUserFromSharedWishlistResponse, error) {
	dbPool, err := shared.ConnectToDatabase(ctx)
	if err != nil {
		return nil, err
	}

	rows, err := dbPool.Query(ctx, `
	WITH authorized_users AS (
  SELECT s.wishlist_id
  FROM shared s
  WHERE s.wishlist_id = $1
    AND s.user_id = $2
    AND (s.can_edit = TRUE OR s.is_owner = TRUE)
	)
	DELETE FROM shared
	WHERE user_id = $3 AND
	is_owner = FALSE AND
	wishlist_id in (SELECT wishlist_id from authorized_users);`,
		request.WishlistId, request.UserId, request.UserToRemove)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer rows.Close()
	return &pb.RemoveUserFromSharedWishlistResponse{}, nil
}
