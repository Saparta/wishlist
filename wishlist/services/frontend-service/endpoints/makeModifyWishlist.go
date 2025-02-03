package endpoints

import (
	"net/http"

	"github.com/Saparta/wishlist/wishlist/services/frontend-service/models"
	pb "github.com/Saparta/wishlist/wishlist/services/frontend-service/proto"
	"github.com/gin-gonic/gin"
)

func MakeModifyWishlist(ctx *gin.Context, client pb.WishlistServiceClient) {

	// Parse the wishlist ID from the URL parameter
	wishlistID := ctx.Param("id")

	if wishlistID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Wishlist ID is required"})
		return
	}

	// Parse the request body for the modified wishlist data
	var updatedWishlist models.Wishlist
	if err := ctx.BindJSON(&updatedWishlist); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	// Call the gRPC service to modify the wishlist
	resp, err := client.ModifyWishlist(ctx.Request.Context(),
		&pb.ModifyWishlistRequest{
			Id:          &wishlistID,
			UserId:      &updatedWishlist.UserID,
			Title:       &updatedWishlist.Title,
			Description: &updatedWishlist.Description,
			IsPublic:    &updatedWishlist.IsPublic,
		})
	if err != nil {
		// Handle error from gRPC call
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the response
	ctx.JSON(http.StatusOK, resp)
}
