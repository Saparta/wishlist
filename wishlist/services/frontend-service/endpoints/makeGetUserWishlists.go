package endpoints

import (
	"net/http"

	pb "github.com/Saparta/wishlist/wishlist/services/frontend-service/proto"
	"github.com/gin-gonic/gin"
)

func MakeGetUserWishlists(ctx *gin.Context, client pb.WishlistServiceClient) {
	id := ctx.Param("userId")
	resp, err := client.GetUserWishlists(ctx.Request.Context(),
		&pb.GetUserWishlistsRequest{UserId: id})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
