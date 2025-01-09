package endpoints

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/Saparta/wishlist/wishlist/services/wishlist-service/proto"
	"github.com/Saparta/wishlist/wishlist/services/wishlist-service/shared"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (w *WishlistService) CreateWishlist(ctx context.Context, request *pb.CreateWishlistRequest) (*pb.CreateWishlistResponse, error) {
	dbPool, ok := ctx.Value(shared.DBSession).(*pgxpool.Pool)
	if !ok {
		return nil, status.Error(codes.Internal, "Failed to retrieve database connection from context")
	}

	rows, err := dbPool.Query(ctx,
		`INSERT INTO wishlists (id, user_id, title, description, is_public) VALUES (gen_random_uuid(), $1, $2, $3, $4) RETURNING id, user_id, title, description, is_public, created_at, last_modified, last_opened`, request.UserId, request.Title, request.Description, request.IsPublic)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer rows.Close()

	wish, err := pgx.CollectOneRow(rows, func(row pgx.CollectableRow) (*pb.CreateWishlistResponse, error) {
		var res pb.CreateWishlistResponse
		var createdAt, lastModified, lastOpened time.Time
		err := row.Scan(&res.Id, &res.UserId, &res.Title, &res.Description, &res.IsPublic, &createdAt, &lastModified, &lastOpened)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		res.CreatedAt = timestamppb.New(createdAt)
		res.LastModified = timestamppb.New(lastModified)
		res.LastOpened = timestamppb.New(lastOpened)
		return &res, nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return wish, nil
}
