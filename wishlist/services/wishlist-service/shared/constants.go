package shared

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ContextKey string

const (
	DBSession ContextKey = "dbSession"
)

func ConnectToDatabase(ctx context.Context) (*pgxpool.Pool, error) {
	dbPool, ok := ctx.Value(DBSession).(*pgxpool.Pool)
	if !ok {
		return nil, status.Error(codes.Internal, "Failed to retrieve database connection from context")
	}
	return dbPool, nil
}
