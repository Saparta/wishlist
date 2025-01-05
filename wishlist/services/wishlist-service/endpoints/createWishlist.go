package endpoints

import (
	"context"
	"log"
	"time"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	pb "github.com/Saparta/wishlist/wishlist/services/wishlist-service/proto"
	"github.com/Saparta/wishlist/wishlist/services/wishlist-service/shared"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type wishlist struct {
	ID           string
	UserID       string
	Title        string
	Description  string
	IsPublic     bool
	CreatedAt    time.Time
	LastModified time.Time
	LastOpened   time.Time
}

func (w *WishlistService) CreateWishlist(ctx context.Context, request *pb.CreateWishlistRequest) (*pb.CreateWishlistResponse, error) {
	dbPool, ok := ctx.Value(shared.DBSession).(*pgxpool.Pool)
	if !ok {
		return nil, status.Error(codes.Internal, "Failed to retrieve database connection from context")
	}

	rows, err := dbPool.Query(context.Background(),
		`INSERT INTO wishlists (id, user_id, title, description, is_public) VALUES ($1, $2, $3, $4, $5) RETURNING id, user_id, title, description, is_public, created_at, last_modified, last_opened`, uuid.New(), request.UserId, request.Title, request.Description, request.IsPublic)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer rows.Close()

	var wish wishlist
	rows.Next()
	err = rows.Scan(&wish.ID, &wish.UserID, &wish.Title, &wish.Description, &wish.IsPublic, &wish.CreatedAt, &wish.LastModified, &wish.LastOpened)
	if err != nil {
		log.Printf("Row scan error: %v\n", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreateWishlistResponse{
		Id:           wish.ID,
		UserId:       wish.UserID,
		Description:  wish.Description,
		IsPublic:     wish.IsPublic,
		CreatedAt:    timestamppb.New(wish.CreatedAt),
		LastModified: timestamppb.New(wish.LastModified),
		LatOpened:    timestamppb.New(wish.LastOpened),
	}, nil
}
