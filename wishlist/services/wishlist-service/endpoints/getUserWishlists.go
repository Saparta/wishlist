package endpoints

import (
	"context"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	"github.com/Saparta/wishlist/wishlist/services/wishlist-service/models"
	pb "github.com/Saparta/wishlist/wishlist/services/wishlist-service/proto"
	"github.com/Saparta/wishlist/wishlist/services/wishlist-service/shared"
	"github.com/jackc/pgx/v5/pgxpool"
)

type completeWishlist struct {
	models.Wishlist
	items []*models.Item
}

// Get every wishlist that a specific user has and all the items within them.
func (w *WishlistService) getUserWishlists(ctx context.Context, request *pb.getUserWishlistsRequest) (*pb.getUserWishlistsResponse, error) {
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
        SELECT w.id AS wishlist_id, 
			w.title AS wishlist_title, 
			w.description AS wishlist_description,
			w.is_public AS is_public,
			i.id AS item_id,
			i.name AS item_name,
			i.url AS item_url,
			i.price AS item_price ,
			i.is_gifted AS item_is_gifted,
			i.gifted_by AS item_gifted_by
        FROM wishlists w
		LEFT JOIN items i
		ON w.id = i.wishlist_id
        WHERE w.user_id = $1
    `, userID)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to query wishlists: "+err.Error())
	}
	defer rows.Close()

	// map from a string to the completeWishlist
	wishMap := make(map[string]completeWishlist)

	// loop over wishlists and fetch each corresponding item
	var wishes []completeWishlist
	for rows.Next() {
		var wish completeWishlist
		if err := rows.Scan(&wish.ID, &wish.Title, &wish.Description, &wish.IsPublic); err != nil {
			return nil, status.Error(codes.Internal, "Wishlist row scan error: "+err.Error())
		}

		// // query the items of the specific wishlist
		// itemRows, err := dbPool.Query(ctx, `
		// 	SELECT id, name, url, price, is_gifted, gifted_by, created_at
		// 	FROM items
		// 	WHERE wishlist_id = $1
		// `, wishlist.Id)

		// if err != nil {
		// 	return nil, status.Error(codes.Internal, "Failed to query wishlist items: "+err.Error())
		// }
		// defer itemRows.Close()

		var items []models.Item
		for itemRows.Next() {
			var item models.Item
			if err := rows.Scan(&item.ID, &item.Name, &item.Url, &item.Price, &item.IsGifted, &item.GiftedBy, &item.CreatedAt); err != nil {
				return nil, status.Error(codes.Internal, "Item row scan error: "+err.Error())
			}

			items = append(items, &item)

		}

		// append to wishlist
		wish.Items = items
		wishes = append(wishlists, &wishlist)
	}

	// Check for errors in iterating over rows
	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Internal, "Error occurred while reading rows: "+err.Error())
	}

	// Return the wishlists response
	return &pb.GetUserWishlistsResponse{
		Wishes: wishes,
	}, nil
}
