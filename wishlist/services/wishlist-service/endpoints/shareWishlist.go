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
  		SELECT s.wishlist_id
  		FROM shared s
  		WHERE s.wishlist_id = $1
    		AND s.user_id = $2
    		AND (s.can_edit = TRUE OR s.is_owner = TRUE)
		)
    INSERT INTO shared (wishlist_id, user_id, can_edit)
    VALUES %s
    ON CONFLICT DO NOTHING
		RETURNING user_id;
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
