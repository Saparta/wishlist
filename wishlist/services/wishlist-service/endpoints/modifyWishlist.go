package endpoints

import (
	"context"
	"time"

	pb "github.com/Saparta/wishlist/wishlist/services/wishlist-service/proto"
	"github.com/Saparta/wishlist/wishlist/services/wishlist-service/shared"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (W *WishlistService) ModifyWishlist(ctx context.Context, request *pb.ModifyWishlistRequest) (*pb.ModifyWishlistResponse, error) {
	dbPool, ok := ctx.Value(shared.DBSession).(*pgxpool.Pool)
	if !ok {
		return nil, status.Error(codes.Internal, "Failed to retrieve database connection from context")
	}
	rows, err := dbPool.Query(ctx, `
	UPDATE wishlists SET
		title = $1,
		description = $2,
		is_public = $3
	WHERE 
		id = $4 AND user_id = $5
	RETURNING id, user_id, title, description, is_public, can_edit, created_at, last_modified, last_opened;`, request.Title, request.Description, request.Id, request.IsPublic, request.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to query database: %s", err.Error())
	}
	result, err := pgx.CollectOneRow(rows, func(row pgx.CollectableRow) (*pb.ModifyWishlistResponse, error) {
		var result pb.ModifyWishlistResponse
		var createdAt, lastModified, lastOpened time.Time
		err = row.Scan(&result.Id, &result.UserId, &result.Title, &result.Description, &result.IsPublic, &result.CanEdit, &createdAt, &lastModified, &lastOpened)
		if err != nil {
			return nil, err
		}
		result.CreatedAt = timestamppb.New(createdAt)
		result.LastModified = timestamppb.New(lastModified)
		result.LastOpened = timestamppb.New(lastOpened)
		return &result, nil
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Parsing results failed: %s", err.Error())
	}

	return result, nil
}
