package endpoints

import (
	"context"

	pb "github.com/Saparta/wishlist/wishlist/services/wishlist-service/proto"
	"github.com/Saparta/wishlist/wishlist/services/wishlist-service/shared"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (w *WishlistService) UnmarkItemGifted(ctx context.Context, request *pb.UnmarkItemGiftedRequest) (*pb.UnmarkItemGiftedResponse, error) {
	dbPool, ok := ctx.Value(shared.DBSession).(*pgxpool.Pool)
	if !ok {
		return nil, status.Error(codes.Internal, "Failed to retrieve database connection from context")
	}
	if request.UserId == nil || *request.UserId == "" {
		return nil, status.Error(codes.InvalidArgument, "UserID is required")
	}

	rows, err := dbPool.Query(ctx, `
	UPDATE items SET
		gifted_by = NULL,
		is_gifted = FALSE
	WHERE
		gifted_by = $1 AND id = $2 AND is_gifted = True
	RETURNING is_gifted;`, request.UserId, request.ItemId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to query DB: %v", err.Error())
	}
	defer rows.Close()

	return &pb.UnmarkItemGiftedResponse{}, nil
}
