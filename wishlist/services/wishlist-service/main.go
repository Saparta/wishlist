package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	"github.com/Saparta/wishlist/wishlist/services/wishlist-service/db"
	"github.com/Saparta/wishlist/wishlist/services/wishlist-service/endpoints"
	pb "github.com/Saparta/wishlist/wishlist/services/wishlist-service/proto"
	"github.com/Saparta/wishlist/wishlist/services/wishlist-service/shared"
	"google.golang.org/grpc"
)

func DBUnaryServerInterceptor(session *pgxpool.Pool) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(context.WithValue(ctx, shared.DBSession, session), req)
	}
}

func main() {
	var dbChannel chan *pgxpool.Pool = make(chan *pgxpool.Pool)
	godotenv.Load()
	go db.SetUpDb(dbChannel)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8081))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var dbPool *pgxpool.Pool = <-dbChannel
	close(dbChannel)
	defer dbPool.Close()
	opts := []grpc.ServerOption{grpc.ChainUnaryInterceptor(DBUnaryServerInterceptor(dbPool))}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterWishlistServiceServer(grpcServer, &endpoints.WishlistService{})
	fmt.Println("Server is listening on port 8081. . .")
	grpcServer.Serve(lis)
}
