package endpoints

import (
	"net/http"

	"github.com/Saparta/wishlist/wishlist/services/frontend-service/models"
	pb "github.com/Saparta/wishlist/wishlist/services/frontend-service/proto"
	"github.com/gin-gonic/gin"
)

func MakeModifyWishlistItem(ctx *gin.Context, client pb.WishlistServiceClient) {
	// Parse the wishlist ID and item ID from the URL parameters
	wishlistID := ctx.Param("wishlistId")
	itemID := ctx.Param("itemId")

	// Check if wishlistID and itemID are provided
	if wishlistID == "" || itemID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Wishlist ID and Item ID are required"})
		return
	}

	// Parse the request body for the modified wishlist item data
	var updatedItem models.Item
	if err := ctx.BindJSON(&updatedItem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	// Call the gRPC service to modify the wishlist item
	resp, err := client.ModifyWishlistItem(ctx.Request.Context(),
		&pb.ModifyWishlistItemRequest{
			Id:         &itemID,
			UserId:     &updatedItem.UserID,
			Name:       &updatedItem.Name,
			Url:        &updatedItem.Url,
			Price:      &updatedItem.Price,
			GiftedStatus:   &updatedItem.IsGifted,
		})
	if err != nil {
		// Handle error from gRPC call
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the response from the gRPC service
	ctx.JSON(http.StatusOK, resp)
}
