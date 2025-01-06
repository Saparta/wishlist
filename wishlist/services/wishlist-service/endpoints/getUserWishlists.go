package endpoints

import (
	"context"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	pb "github.com/Saparta/wishlist/wishlist/services/wishlist-service/proto"
	"github.com/Saparta/wishlist/wishlist/services/wishlist-service/shared"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Get every wishlist that a specific user has and all the items within them.
func (w *WishlistService) getUserWishlists(ctx context.Context, request *pb.getUserWishlistsRequest) (*pb.CreateWishlistResponse, error) {
	dbPool, ok := ctx.Value(shared.DBSession).(*pgxpool.Pool)
	if !ok {
		return nil, status.Error(codes.Internal, "Failed to retrieve database connection from context")
	}

	// get user id
	userID := request.GetUserId()
	if userID == "" {
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	}

	// query the wishlists for user id
	rows, err := dbPool.Query(ctx, `
        SELECT id, name, created_at
        FROM wishlists
        WHERE user_id = $1
    `, userID)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to query wishlists: "+err.Error())
	}
	defer rows.Close()

	// loop over wishlists and fetch each corresponding item
	var wish wishlist
	for rows.Next() {
		if err := rows.Scan(&wish.Id, &wish.Title, &wish.Description, &wish.IsPublic, &wish.createdAt, &wish.lastModified, &wish.lastOpened); err != nil {
			return nil, status.Error(codes.Internal, "Wishlist row scan error: "+err.Error())
		}

		// query the items of the specific wishlist
		itemRows, err := dbPool.Query(ctx, `
			SELECT id, name, url, price, is_gifted, gifted_by, created_at
			FROM items
			WHERE wishlist_id = $1
		`, wishlist.Id)
		if err != nil {
			return nil, status.Error(codes.Internal, "Failed to query wishlist items: "+err.Error())
		}
		defer itemRows.Close()

		var items []item
		for itemRows.Next() {
			var item item
			if err := rows.Scan(&items.Id, &items.Name, &item.URL, &item.Price, &item.IsGifted, &item.GiftedBy, &item.CreatedAt); err != nil {
				return nil, status.Error(codes.Internal, "Item row scan error: "+err.Error())
			}

			items = append(items, &item)

		}
	}
}
