package endpoints

import (
	"context"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	pb "github.com/Saparta/wishlist/wishlist/services/wishlist-service/proto"
	"github.com/Saparta/wishlist/wishlist/services/wishlist-service/shared"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (w *WishlistService) CreateWishlist(ctx context.Context, request *pb.CreateWishlistRequest) (*pb.CreateWishlistResponse, error) {
	dbPool, ok := ctx.Value(shared.DBSession).(*pgxpool.Pool)
	if !ok {
		return nil, status.Error(codes.Internal, "Failed to retrieve database connection from context")
	}

	rows, err := dbPool.Query(context.Background(), `INSERT INTO wishlists (id, user_id, title, description, is_public) VALUES ($1, $2, $3, $4, $5)`, uuid.New(), request.UserId, request.Title, request.Description, request.IsPublic)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer rows.Close()

	return &pb.CreateWishlistResponse{}, nil
}
