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
		SELECT w.id AS wishlist_id, w.user_id, s.shared_with FROM
    wishlists w
    LEFT JOIN shared s ON w.id = s.wishlist_id
    WHERE w.id = $1
      AND (w.user_id = $2 OR (s.shared_with = $2 AND s.can_edit = TRUE))
	)
	DELETE FROM shared
	WHERE shared_with = $3 AND
	wishlist_id in (SELECT wishlist_id from authorized_users)`,
		request.WishlistId, request.UserId, request.UserToRemove)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer rows.Close()
	return &pb.RemoveUserFromSharedWishlistResponse{}, nil
}
