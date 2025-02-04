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
    SELECT s.wishlist_id
    FROM shared s
    WHERE s.wishlist_id = $1
      AND s.user_id = $2
      AND (s.can_edit = TRUE OR s.is_owner = TRUE)
	),
	updated_wishlist AS (
    UPDATE wishlists
    SET
      title = COALESCE($3, title),
      description = COALESCE($4, description),
      is_public = COALESCE($5, is_public)
    WHERE id IN (SELECT wishlist_id FROM authorized_users)
    RETURNING id, title, description, is_public
	),
	update_shared AS (
    UPDATE shared
    SET last_modified = CURRENT_TIMESTAMP
    WHERE wishlist_id IN (SELECT wishlist_id FROM authorized_users)
      AND user_id = $2
    RETURNING user_id, created_at, last_modified, last_opened, wishlist_id
	)
	SELECT 
    uw.id AS wishlist_id,
    us.user_id,
    uw.title AS wishlist_title,
    uw.description AS wishlist_description,
    uw.is_public AS is_public,
    us.created_at,
    us.last_modified,
    us.last_opened
	FROM updated_wishlist uw
	JOIN update_shared us ON uw.id = us.wishlist_id;`, request.Id, request.UserId, request.Title, request.Description, request.IsPublic)
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
