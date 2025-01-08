package util

import (
	"github.com/Saparta/wishlist/wishlist/services/frontend-service/models"
	pb "github.com/Saparta/wishlist/wishlist/services/frontend-service/proto"
)

func PrepareItemsForJSON(items []*pb.WishlistItem) []models.Item {
	var transformedItems []models.Item
	for _, item := range items {
		transformedItems = append(transformedItems, models.Item{
			ID:        *item.Id,
			Name:      *item.Name,
			Url:       *item.Url,
			Price:     *item.Price,
			IsGifted:  *item.IsGifted,
			GiftedBy:  *item.GiftedBy,
			CreatedAt: item.CreatedAt.AsTime(),
		})
	}
	return transformedItems
}

func PrepareWishlistsForJSON(respWishlists []*pb.Wishlist) []models.Wishlist {
	var wishlists []models.Wishlist
	for _, wishlist := range respWishlists {
		wishlists = append(wishlists, models.Wishlist{
			ID:           *wishlist.Id,
			UserID:       *wishlist.UserId,
			Title:        *wishlist.Title,
			Description:  *wishlist.Description,
			IsPublic:     *wishlist.IsPublic,
			CanEdit:      *wishlist.CanEdit,
			CreatedAt:    wishlist.CreatedAt.AsTime(),
			LastModified: wishlist.LastModified.AsTime(),
			LastOpened:   wishlist.LastModified.AsTime(),
			Items:        PrepareItemsForJSON(wishlist.Items),
		})
	}
	return wishlists
}
