package endpoints

import (
	pb "github.com/Saparta/wishlist/wishlist/services/wishlist-service/proto"
)

type WishlistService struct {
	pb.UnimplementedWishlistServiceServer
}
