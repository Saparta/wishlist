package endpoints

import (
	"net/http"

	"github.com/Saparta/wishlist/wishlist/services/frontend-service/models"
	pb "github.com/Saparta/wishlist/wishlist/services/frontend-service/proto"
	"github.com/gin-gonic/gin"
)

func MakeCreateWishlist(ctx *gin.Context, client pb.WishlistServiceClient) {
	var newWishlist models.Wishlist

	if err := ctx.BindJSON(&newWishlist); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	resp, err := client.CreateWishlist(ctx.Request.Context(),
		&pb.CreateWishlistRequest{
			UserId:      newWishlist.UserID,
			Title:       newWishlist.Title,
			Description: newWishlist.Description,
			IsPublic:    newWishlist.IsPublic})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
