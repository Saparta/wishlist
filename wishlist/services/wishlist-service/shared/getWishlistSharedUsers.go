package shared

import (
	"context"

	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetWishlistSharedUsers(ctx context.Context, wishlistId string, doneChannel chan []string, errChannel chan error) {
	dbPool, err := ConnectToDatabase(ctx)
	if err != nil {
		errChannel <- status.Errorf(codes.Internal, "Failed to connect to database: %s", err.Error())
		return
	}
	rows, err := dbPool.Query(ctx, "SELECT shared_with FROM shared WHERE wishlist_id = $1", wishlistId)
	if err != nil {
		errChannel <- status.Errorf(codes.Internal, "Failed to query database: %s", err.Error())
		return
	}

	allSharedWithUsers, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (string, error) {
		var userId string
		err = row.Scan(&userId)
		if err != nil {
			return "", status.Errorf(codes.Internal, "Failed to parse query results: %s", err.Error())
		}
		return userId, nil
	})
	if err != nil {
		errChannel <- err
		return
	}
	doneChannel <- allSharedWithUsers
}

func BatchGetWishlistsSharedUsers(ctx context.Context, wishlistIds []string, doneChannel chan map[string][]string, errChannel chan error) {
	dbPool, err := ConnectToDatabase(ctx)
	if err != nil {
		errChannel <- status.Errorf(codes.Internal, "Failed to connect to database: %s", err.Error())
		return
	}

	rows, err := dbPool.Query(ctx,
		"SELECT wishlist_id, shared_with FROM shared WHERE wishlist_id = ANY($1)", wishlistIds)
	if err != nil {
		errChannel <- status.Errorf(codes.Internal, "Failed to query database: %s", err.Error())
		return
	}

	sharedUsers := make(map[string][]string)
	var wishId, sharedWith string
	_, err = pgx.ForEachRow(rows, []any{&wishId, &sharedWith}, func() error {
		usersShared, found := sharedUsers[wishId]
		if found {
			sharedUsers[wishId] = append(usersShared, sharedWith)
			return nil
		}
		sharedUsers[wishId] = []string{sharedWith}
		return nil
	})

	if err != nil {
		errChannel <- status.Errorf(codes.Internal, "Failed to parse query results: %s", err.Error())
		return
	}

	doneChannel <- sharedUsers
}
