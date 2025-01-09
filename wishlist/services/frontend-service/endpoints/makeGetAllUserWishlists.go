package endpoints

import (
	"net/http"

	pb "github.com/Saparta/wishlist/wishlist/services/frontend-service/proto"
	"github.com/Saparta/wishlist/wishlist/services/frontend-service/util"
	"github.com/gin-gonic/gin"
)

func MakeGetAllUserWishlists(ctx *gin.Context, client pb.WishlistServiceClient) {
	id := ctx.Param("userId")
	resp, err := client.GetAllUserWishlists(ctx.Request.Context(),
		&pb.GetAllUserWishlistsRequest{UserId: &id})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, util.PrepareWishlistsForJson(resp.Wishlists))
}
