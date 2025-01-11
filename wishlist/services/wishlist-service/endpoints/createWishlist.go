package endpoints

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/Saparta/wishlist/wishlist/services/wishlist-service/proto"
	"github.com/Saparta/wishlist/wishlist/services/wishlist-service/shared"
	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (w *WishlistService) CreateWishlist(ctx context.Context, request *pb.CreateWishlistRequest) (*pb.CreateWishlistResponse, error) {
	dbPool, err := shared.ConnectToDatabase(ctx)
	if err != nil {
		return nil, err
	}

	rows, err := dbPool.Query(ctx, `
		WITH updated_wishlists AS (
			INSERT INTO wishlists (title, description, is_public) 
			VALUES ($2, $3, $4) 
			RETURNING id, title, description, is_public
		),
		updated_shared AS (
			INSERT INTO shared (wishlist_id, user_id, is_owner, can_edit)
			SELECT id, $1, TRUE, TRUE
			FROM updated_wishlists
			RETURNING *
		)
		SELECT 
			w.id, s.user_id, w.title, w.description,
			w.is_public, s.created_at, s.last_modified, s.last_opened
			FROM updated_wishlists w JOIN updated_shared s ON w.id = s.wishlist_id;`,
		request.UserId, request.Title, request.Description, request.IsPublic)
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
