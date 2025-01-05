package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	pb "github.com/Saparta/wishlist/wishlist/services/wishlist-service/proto"
	"google.golang.org/grpc"
)

type wishlistService struct {
	pb.UnimplementedWishlistServiceServer
}

type contextKey string

const (
	DBSession contextKey = "dbSession"
)

func (w *wishlistService) CreateWishlist(ctx context.Context, request *pb.CreateWishlistRequest) (*pb.CreateWishlistResponse, error) {
	return &pb.CreateWishlistResponse{}, nil
}

func createTables(dbpool *pgxpool.Pool) error {
	query := `
	CREATE TABLE IF NOT EXISTS wishlists(
		id UUID PRIMARY KEY,
		FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
		title VARCHAR(255),
		description VARCHAR(255),
		is_public boolean,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		last_modified TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		last_opened TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS items(
		id UUID PRIMARY KEY,
		FOREIGN KEY(wishlist_id) REFERENCES wishlists(id) ON DELETE CASCADE,
		name VARCHAR(255),
		url TEXT,
		price REAL,
		is_gifted boolean,
		FOREIGN KEY(gifted_by) REFERENCES users(id) ON DELETE SET NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS shared(
		id UUID PRIMARY KEY,
		FOREIGN KEY(wishlist_id) REFERENCES wishlists(id) ON DELETE CASCADE,
		FOREIGN KEY(shared_with) REFERENCES users(id) ON DELETE CASCADE,
		can_edit boolean,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := dbpool.Exec(context.Background(), query)
	if err != nil {
		return err
	}

	return nil
}

func setUpDb(channel chan *pgxpool.Pool) {
	var dsn string = os.Getenv("DSN")

	dbPool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	err = createTables(dbPool)
	if err != nil {
		log.Fatalf("Failed to create users table: %v\n", err)
	}

	channel <- dbPool
}

func DBUnaryServerInterceptor(session *pgxpool.Pool) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(context.WithValue(ctx, DBSession, session), req)
	}
}

func main() {
	var dbChannel chan *pgxpool.Pool = make(chan *pgxpool.Pool)
	godotenv.Load()
	go setUpDb(dbChannel)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8081))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var dbPool *pgxpool.Pool = <-dbChannel
	defer dbPool.Close()
	opts := []grpc.ServerOption{grpc.ChainUnaryInterceptor(DBUnaryServerInterceptor(dbPool))}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterWishlistServiceServer(grpcServer, &wishlistService{})
	grpcServer.Serve(lis)

}
