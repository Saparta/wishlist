package endpoints

import (
	"context"
	"fmt"
	"strings"

	pb "github.com/Saparta/wishlist/wishlist/services/wishlist-service/proto"
	"github.com/Saparta/wishlist/wishlist/services/wishlist-service/shared"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (W *WishlistService) ShareWishlist(ctx context.Context, request *pb.ShareWishlistRequest) (*pb.ShareWishlistResponse, error) {
	dbPool, err := shared.ConnectToDatabase(ctx)
	if err != nil {
		return nil, err
	}

	placeholders := []string{}
	args := []any{*request.WishlistId, *request.UserId, *request.CanEdit}

	for i, userId := range request.ShareWith {
		// +4 because $0, $1, $2, $3 are taken or invalid
		placeholders = append(placeholders, fmt.Sprintf("($1, $%d, $3)", i+4))
		args = append(args, userId)
	}

	query := fmt.Sprintf(`
    WITH authorized_users AS (
        SELECT w.id AS wishlist_id, s.shared_with
        FROM wishlists w
        LEFT JOIN shared s ON w.id = s.wishlist_id
        WHERE w.id = $1
          AND (w.user_id = $2 OR (s.shared_with = $2 AND s.can_edit = TRUE))
    )
    INSERT INTO shared (wishlist_id, shared_with, can_edit)
    VALUES %s
    ON CONFLICT DO NOTHING
		RETURNING shared_with;
`, strings.Join(placeholders, ", "))

	rows, err := dbPool.Query(ctx, query, args...)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to query database: %s", err.Error())
	}
	var sharedUserId string
	sharedWith, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (string, error) {
		err := row.Scan(&sharedUserId)
		if err != nil {
			return "", status.Errorf(codes.Internal, "Failed to parse query results: %s", err.Error())
		}
		return sharedUserId, nil
	})
	if err != nil {
		return nil, err
	}

	return &pb.ShareWishlistResponse{SharedWith: sharedWith}, nil
}
