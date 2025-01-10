package endpoints

import (
	"context"
	"time"

	pb "github.com/Saparta/wishlist/wishlist/services/wishlist-service/proto"
	"github.com/Saparta/wishlist/wishlist/services/wishlist-service/shared"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (w *WishlistService) ModifyWishlistItem(ctx context.Context, request *pb.ModifyWishlistItemRequest) (*pb.ModifyWishlistItemResponse, error) {
	dbPool, err := shared.ConnectToDatabase(ctx)
	if err != nil {
		return nil, err
	}

	var done chan *pb.ItemMarkingResponse = make(chan *pb.ItemMarkingResponse)
	var errChan chan error = make(chan error)
	go shared.MarkItem(ctx, &pb.ItemMarkingRequest{UserId: request.UserId, ItemId: request.Id}, request.GiftedStatus, done, errChan)
	row, err := dbPool.Query(ctx, `
	WITH authorized_users AS (
    SELECT i.id AS item_id
    FROM items i
    JOIN wishlists w ON i.wishlist_id = w.id
    LEFT JOIN shared s ON w.id = s.wishlist_id
    WHERE i.id = $1
      AND (w.user_id = $2 OR (s.shared_with = $2 AND s.can_edit = TRUE))
	)
	UPDATE items
	SET
    name = $3,
    url = $4,
    price = $5
	WHERE id IN (SELECT item_id FROM authorized_users);`,
		request.Id, request.UserId, request.Name, request.Url, request.Price)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to query database: %s", err.Error())
	}
	defer row.Close()

	select {
	case <-done:
		row := dbPool.QueryRow(ctx, "SELECT id, wishlist_id, name, url, price, is_gifted, gifted_by, created_at FROM items WHERE id = $1;", request.Id)
		var response pb.ModifyWishlistItemResponse = pb.ModifyWishlistItemResponse{
			Item: &pb.WishlistItem{},
		}
		var createdAt time.Time
		err = row.Scan(&response.Item.Id, &response.Item.WishlistId, &response.Item.Name, &response.Item.Url, &response.Item.Price, &response.Item.IsGifted, &response.Item.GiftedBy, &createdAt)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to query database: %s", err.Error())
		}
		response.Item.CreatedAt = timestamppb.New(createdAt)
		return &response, nil
	case err := <-errChan:
		return nil, status.Errorf(codes.Internal, "Failed to query database: %s", err.Error())
	}
}
