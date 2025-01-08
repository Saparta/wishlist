package endpoints

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Saparta/wishlist/wishlist/services/wishlist-service/models"
	pb "github.com/Saparta/wishlist/wishlist/services/wishlist-service/proto"
	"github.com/Saparta/wishlist/wishlist/services/wishlist-service/shared"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (w *WishlistService) CreateWishlist(ctx context.Context, request *pb.CreateWishlistRequest) (*pb.CreateWishlistResponse, error) {
	dbPool, ok := ctx.Value(shared.DBSession).(*pgxpool.Pool)
	if !ok {
		return nil, status.Error(codes.Internal, "Failed to retrieve database connection from context")
	}

	rows, err := dbPool.Query(ctx,
		`INSERT INTO wishlists (id, user_id, title, description, is_public) VALUES (gen_random_uuid(), $1, $2, $3, $4) RETURNING id, user_id, title, description, is_public, created_at, last_modified, last_opened`, request.UserId, request.Title, request.Description, request.IsPublic)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer rows.Close()

	var wish models.Wishlist
	rows.Next()
	err = rows.Scan(&wish.ID, &wish.UserID, &wish.Title, &wish.Description, &wish.IsPublic, &wish.CreatedAt, &wish.LastModified, &wish.LastOpened)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreateWishlistResponse{
		Id:           &wish.ID,
		UserId:       &wish.UserID,
		Description:  &wish.Description,
		IsPublic:     &wish.IsPublic,
		CreatedAt:    timestamppb.New(wish.CreatedAt),
		LastModified: timestamppb.New(wish.LastModified),
		LastOpened:   timestamppb.New(wish.LastOpened),
	}, nil
}
