package endpoints

import (
	"context"
	"errors"
	"time"

	pb "github.com/Saparta/wishlist/wishlist/services/wishlist-service/proto"
	"github.com/Saparta/wishlist/wishlist/services/wishlist-service/shared"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (W *WishlistService) ModifyWishlist(ctx context.Context, request *pb.ModifyWishlistRequest) (*pb.ModifyWishlistResponse, error) {
	dbPool, err := shared.ConnectToDatabase(ctx)
	if err != nil {
		return nil, err
	}
	
	rows, err := dbPool.Query(ctx, `
	WITH authorized_users AS (
    SELECT w.id AS wishlist_id
    FROM wishlists w
    LEFT JOIN shared s ON w.id = s.wishlist_id
    WHERE w.id = $1
      AND (w.user_id = $2 OR (s.shared_with = $2 AND s.can_edit = TRUE))
	)
	UPDATE wishlists SET
		title = $3,
		description = $4,
		is_public = $5
	WHERE 
		id in (SELECT wishlist_id FROM authorized_users)
	RETURNING id, user_id, title, description, is_public, created_at, last_modified, last_opened;`, request.Id, request.UserId, request.Title, request.Description, request.IsPublic)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to query database: %s", err.Error())
	}
	result, err := pgx.CollectOneRow(rows, func(row pgx.CollectableRow) (*pb.ModifyWishlistResponse, error) {
		var result pb.ModifyWishlistResponse
		var createdAt, lastModified, lastOpened time.Time
		err = row.Scan(&result.Id, &result.UserId, &result.Title, &result.Description, &result.IsPublic, &createdAt, &lastModified, &lastOpened)
		if err != nil {
			return nil, err
		}
		result.CanEdit = func() *bool { b := true; return &b }()
		result.CreatedAt = timestamppb.New(createdAt)
		result.LastModified = timestamppb.New(lastModified)
		result.LastOpened = timestamppb.New(lastOpened)
		return &result, nil
	})

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Errorf(codes.InvalidArgument, "User isn't authorized")
		}
		return nil, status.Errorf(codes.Internal, "Parsing results failed: %s", err.Error())
	}

	return result, nil
}
