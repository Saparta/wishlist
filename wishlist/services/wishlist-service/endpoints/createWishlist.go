package endpoints

import (
	"context"
	"fmt"

	pb "github.com/Saparta/wishlist/wishlist/services/wishlist-service/proto"
	"github.com/Saparta/wishlist/wishlist/services/wishlist-service/shared"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (w *WishlistService) CreateWishlist(ctx context.Context, request *pb.CreateWishlistRequest) (*pb.CreateWishlistResponse, error) {
	dbPool, ok := ctx.Value(shared.DBSession).(*pgxpool.Pool)
	if !ok {
		return nil, fmt.Errorf("failed to retrieve database connection from context")
	}

	wishlistID := uuid.New()

	_, err := dbPool.Query(context.Background(), `INSERT INTO wishlists (id, user_id, title, description, is_public) VALUES ($1, $2, $3, $4, $5)`, wishlistID, request.UserId, request.Title, request.Description, request.IsPublic)

	if err != nil {
		return nil, fmt.Errorf("failed to insert wishlist into database") // Handle
	}

	return &pb.CreateWishlistResponse{}, nil
}
