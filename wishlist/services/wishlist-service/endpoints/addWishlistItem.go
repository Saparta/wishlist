package endpoints

import (
	"context"
	"log"

	"github.com/Saparta/wishlist/wishlist/services/wishlist-service/models"
	pb "github.com/Saparta/wishlist/wishlist/services/wishlist-service/proto"
	"github.com/Saparta/wishlist/wishlist/services/wishlist-service/shared"
	"github.com/jackc/pgx/v5/pgxpool"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/google/uuid"
)

func (w *WishlistService) AddWishlistItem(ctx context.Context, request *pb.AddWishlistItemRequest) (*pb.AddWishlistItemResponse, error) {
	dbPool, ok := ctx.Value(shared.DBSession).(*pgxpool.Pool)
	if !ok {
		return nil, status.Error(codes.Internal, "Failed to retrieve database connection from context")
	}

	query := `INSERT INTO items (id, wishlist_id, name, url, price, is_gifted, gifted_by) 
	SELECT $1, $2, $3, $4, $5, $6, $7
	FROM wishlists
	WHERE id = $2 AND user_id = $8
	RETURNING *;`

	rows, err := dbPool.Query(ctx, query,
		uuid.New(), request.WishlistId, request.Name, request.Url, request.Price,
		request.IsGifted, request.GiftedBy, request.UserId)
	if err != nil {
		log.Print("query failure")
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer rows.Close()
	
	if !rows.Next() {
		return nil, status.Error(codes.NotFound, "No matching wishlist found or ownership check failed")
	}
	var item models.Item
	err = rows.Scan(&item.ID, &item.WishlistID, &item.Name, &item.Url, &item.Price, &item.IsGifted, &item.GiftedBy, &item.CreatedAt)
	if err != nil {
		log.Print("Row scan failure")
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.AddWishlistItemResponse{
		Id:         item.ID,
		WishlistId: item.WishlistID,
		Name:       item.Name,
		Url:        item.Url,
		Price:      item.Price,
		IsGifted:   item.IsGifted,
		GiftedBy:   item.GiftedBy,
		CreatedAt:  timestamppb.New(item.CreatedAt),
	}, nil
}
