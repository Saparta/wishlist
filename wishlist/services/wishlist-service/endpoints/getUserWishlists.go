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

	// Execute the query
	rows, err := dbPool.Query(ctx, `
        SELECT w.id AS wishlist_id, 
            w.title AS wishlist_title, 
            w.description AS wishlist_description,
            w.is_public AS is_public,
            i.id AS item_id,
            i.name AS item_name,
            i.url AS item_url,
            i.price AS item_price,
            i.is_gifted AS item_is_gifted,
            i.gifted_by AS item_gifted_by
        FROM wishlists w
        LEFT JOIN items i ON w.id = i.wishlist_id
        WHERE w.user_id = $1
    `, userID)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to query wishlists: "+err.Error())
	}
	defer rows.Close()

	var wishlists []*pb.Wishlist
	var currentWishlist *pb.Wishlist

	// Iterate through the rows
	for rows.Next() {
		var wishlistID, wishlistTitle, wishlistDescription string
		var isPublic bool
		var itemID, itemName, itemUrl, itemGiftedBy string
		var itemPrice float64
		var itemIsGifted bool

		// Scan the row into variables
		err := rows.Scan(
			&wishlistID,
			&wishlistTitle,
			&wishlistDescription,
			&isPublic,
			&itemID,
			&itemName,
			&itemUrl,
			&itemPrice,
			&itemIsGifted,
			&itemGiftedBy,
		)
		if err != nil {
			return nil, status.Error(codes.Internal, "Failed to scan row: "+err.Error())
		}

		// If we don't have a current wishlist or it's a new wishlist, create a new one
		if currentWishlist == nil || currentWishlist.Id != wishlistID {
			// If we already have a current wishlist, append it to the result set
			if currentWishlist != nil {
				wishlists = append(wishlists, currentWishlist)
			}

			// Create a new wishlist entry
			currentWishlist = &pb.Wishlist{
				Id:          wishlistID,
				Title:       wishlistTitle,
				Description: wishlistDescription,
				IsPublic:    isPublic,
				Items:       []*pb.WishlistItem{}, // Initialize with an empty slice
			}
		}

		// Add item to the current wishlist (if itemID is not empty)
		if itemID != "" {
			item := &pb.WishlistItem{
				Id:       itemID,
				Name:     itemName,
				Url:      itemUrl,
				Price:    itemPrice,
				IsGifted: itemIsGifted,
				GiftedBy: &itemGiftedBy, // This could be nil if the value is NULL
			}
			currentWishlist.Items = append(currentWishlist.Items, item)
		}
	}

	// Don't forget to append the last wishlist if it's not nil
	if currentWishlist != nil {
		wishlists = append(wishlists, currentWishlist)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Internal, "Error occurred while reading rows: "+err.Error())
	}

	// Return the final result
	return &pb.GetUserWishlistsResponse{
		Wishlists: wishlists,
	}, nil
}
