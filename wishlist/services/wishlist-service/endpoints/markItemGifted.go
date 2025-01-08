package endpoints

import (
	"context"
	"errors"

	pb "github.com/Saparta/wishlist/wishlist/services/wishlist-service/proto"
	"github.com/Saparta/wishlist/wishlist/services/wishlist-service/shared"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (w *WishlistService) MarkItemGifted(ctx context.Context, request *pb.MarkItemGiftedRequest) (*pb.MarkItemGiftedResponse, error) {
	dbPool, ok := ctx.Value(shared.DBSession).(*pgxpool.Pool)
	if !ok {
		return nil, status.Error(codes.Internal, "Failed to retrieve database connection from context")
	}

	userID := request.UserId
	if userID == nil || *userID == "" {
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	}
	rows, err := dbPool.Query(ctx, `
	UPDATE items SET 
		gifted_by = $1,
		is_gifted = True
	WHERE 
		id = $2 AND is_gifted = False
	RETURNING is_gifted;`, userID, request.ItemId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failure to query items: %v", err)
	}
	
	var success bool
	_, err = pgx.CollectExactlyOneRow(rows, func(row pgx.CollectableRow) (any, error) {
		return nil, nil
	})

	if errors.Is(err, pgx.ErrNoRows) {
		success = false
		return &pb.MarkItemGiftedResponse{
			Success: &success,
		}, nil
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, "Failure to scan query result: %v", err)
	}

	success = true
	return &pb.MarkItemGiftedResponse{
		Success: &success,
	}, nil
}
