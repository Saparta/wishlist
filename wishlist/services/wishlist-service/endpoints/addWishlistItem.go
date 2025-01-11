package endpoints

import (
	"context"
	"errors"
	"log"
	"time"

	pb "github.com/Saparta/wishlist/wishlist/services/wishlist-service/proto"
	"github.com/Saparta/wishlist/wishlist/services/wishlist-service/shared"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (w *WishlistService) AddWishlistItem(ctx context.Context, request *pb.AddWishlistItemRequest) (*pb.AddWishlistItemResponse, error) {
	dbPool, err := shared.ConnectToDatabase(ctx)
	if err != nil {
		return nil, err
	}

	rows, err := dbPool.Query(ctx, `
	WITH authorized_users AS (
		SELECT s.wishlist_id, s.user_id FROM
		shared s
		WHERE s.wishlist_id = $1 AND
			s.user_id = $5 AND (s.can_edit = TRUE OR s.is_owner = TRUE)
	),
	updated_items AS (
		INSERT INTO items (wishlist_id, name, url, price)
		SELECT $1, $2, $3, $4
		WHERE EXISTS (
    	SELECT 1 FROM authorized_users
		)
		RETURNING id, wishlist_id, name, url, price, is_gifted, gifted_by, created_at
	),
	update_shared AS (
    UPDATE shared
    SET last_modified = CURRENT_TIMESTAMP
    WHERE wishlist_id IN (SELECT wishlist_id FROM authorized_users)
    AND user_id = $5
	)
	SELECT * FROM updated_items;
	`, request.WishlistId, request.Name, request.Url, request.Price, request.UserId)
	if err != nil {
		log.Print("query failure")
		return nil, status.Error(codes.Internal, err.Error())
	}

	result, err := pgx.CollectOneRow(rows, func(row pgx.CollectableRow) (*pb.AddWishlistItemResponse, error) {
		var res pb.AddWishlistItemResponse
		var createdAt time.Time
		row.Scan(&res.Id, &res.WishlistId, &res.Name, &res.Url, &res.Price, &res.IsGifted, &res.GiftedBy, &createdAt)
		res.CreatedAt = timestamppb.New(createdAt)
		return &res, nil
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "Could not add item")
		} else {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return result, nil
}
