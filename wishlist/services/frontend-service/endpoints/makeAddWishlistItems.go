// call add wishlist items that is in wishlists service
// similar to makeCreateWishlist
// register the endpoint with gin(library) in main

package endpoints

import (
	"net/http"

	"github.com/Saparta/wishlist/wishlist/services/frontend-service/models"
	pb "github.com/Saparta/wishlist/wishlist/services/frontend-service/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

func MakeAddWishlistItem(ctx *gin.Context, client pb.WishlistServiceClient) {
	var newItem models.Item

	// Bind the incoming JSON request to the Item struct
	if err := ctx.BindJSON(&newItem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	// Validate that WishlistId is provided in the request
	if newItem.WishlistID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Wishlist ID is required"})
		return
	}

	// Prepare the AddWishlistItem request for gRPC
	resp, err := client.AddWishlistItem(ctx.Request.Context(),
		&pb.AddWishlistItemRequest{
			WishlistId: &newItem.WishlistID,
			Name:       &newItem.Name,
			Url:        &newItem.Url,
			Price:      &newItem.Price,
			IsGifted:   &newItem.IsGifted,
			GiftedBy:   &newItem.GiftedBy,
		})

	// Handle errors from the gRPC call
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": st.Message(), "details": st.Code()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error", "details": err.Error()})
		}
		return
	}

	// Return a successful response with the added item
	ctx.JSON(http.StatusOK, gin.H{
		"message":    "Item added to wishlist successfully",
		"added_item": resp,
	})
}
