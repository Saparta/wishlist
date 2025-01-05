package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/Saparta/wishlist/wishlist/services/wishlist-service/proto"
	"google.golang.org/grpc"
)

type wishlistService struct {
	pb.UnimplementedWishlistServiceServer
}

func (w *wishlistService) CreateWishlist(ctx context.Context, request *pb.CreateWishlistRequest) (*pb.CreateWishlistResponse, error) {
	return &pb.CreateWishlistResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8081))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterWishlistServiceServer(grpcServer, &wishlistService{})
	grpcServer.Serve(lis)

}
