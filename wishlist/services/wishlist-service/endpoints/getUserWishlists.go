package endpoints

import (
	"context"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	pb "github.com/Saparta/wishlist/wishlist/services/wishlist-service/proto"
	"github.com/Saparta/wishlist/wishlist/services/wishlist-service/shared"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Get every wishlist that a specific user has and all the items within them.
func (w *WishlistService) GetUserWishlists(ctx context.Context, request *pb.GetUserWishlistsRequest) (*pb.GetUserWishlistsResponse, error) {
	dbPool, ok := ctx.Value(shared.DBSession).(*pgxpool.Pool)
	if !ok {
		return nil, status.Error(codes.Internal, "Failed to retrieve database connection from context")
	}

	// get user id
	userID := request.GetUserId()
	if userID == "" {
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	}

	// query the wishlists for user id
	rows, err := dbPool.Query(ctx, `
        SELECT w.id AS wishlist_id, 
			w.title AS wishlist_title, 
			w.description AS wishlist_description,
			w.is_public AS is_public,
			i.id AS item_id,
			i.name AS item_name,
			i.url AS item_url,
			i.price AS item_price ,
			i.is_gifted AS item_is_gifted,
			i.gifted_by AS item_gifted_by
        FROM wishlists w
		LEFT JOIN items i
		ON w.id = i.wishlist_id
        WHERE w.user_id = $1
    `, userID)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to query wishlists: "+err.Error())
	}
	defer rows.Close()

	// Return the wishlists response
	return &pb.GetUserWishlistsResponse{}, nil
}
