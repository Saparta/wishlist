package shared

import (
	"context"

	"github.com/Saparta/wishlist/wishlist/services/wishlist-service/proto"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func MarkItem(ctx context.Context, request *proto.ItemMarkingRequest, markingOption *bool, succChannel chan *proto.ItemMarkingResponse, errChannel chan error) {
	if markingOption == nil {
		succChannel <- &proto.ItemMarkingResponse{}
		return
	}
	dbPool, ok := ctx.Value(DBSession).(*pgxpool.Pool)
	if !ok {
		errChannel <- status.Error(codes.Internal, "Failed to retrieve database connection from context")
	}

	var userId *string
	if *markingOption {
		userId = request.UserId
	}

	row := dbPool.QueryRow(ctx, `
	UPDATE items SET 
		gifted_by = $1,
		is_gifted = $2
		last_modified = CURRENT_TIMESTAMP
	WHERE 
		id = $3 AND is_gifted = $4
	RETURNING id;`, userId, *markingOption, request.ItemId, !*markingOption)
	var id string
	err := row.Scan(&id)
	if err != nil {
		if err == pgx.ErrNoRows {
			succChannel <- &proto.ItemMarkingResponse{Success: false}
			return
		} else {
			errChannel <- status.Errorf(codes.Internal, "Failure to query items: %v", err)
			return
		}
	}
	succChannel <- &proto.ItemMarkingResponse{Success: true}
}
