package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func setUpDb(channel chan *pgxpool.Pool) {
	godotenv.Load()
	var dsn string = os.Getenv("DSN")

	// Create a connection pool
	dbPool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	channel <- dbPool
}

type User struct {
	ID        string    `json:"id"`
	OAuthID   string    `json:"oauth_id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func getUsers(ctx *gin.Context, dbpool *pgxpool.Pool) {
	rows, err := dbpool.Query(context.Background(), "SELECT id, oauth_id, email, name, created_at FROM users")
	if err != nil {
		log.Printf("Query error: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	defer rows.Close()

	// Collect the users into a slice
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.OAuthID, &user.Email, &user.Name, &user.CreatedAt)
		if err != nil {
			log.Printf("Row scan error: %v\n", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user data"})
			return
		}
		users = append(users, user)
	}

	// Return the users as JSON
	ctx.JSON(http.StatusOK, users)
}

func main() {
	var dbChannel chan *pgxpool.Pool = make(chan *pgxpool.Pool)
	go setUpDb(dbChannel)

	var r *gin.Engine = gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	var dbPool *pgxpool.Pool = <-dbChannel
	defer dbPool.Close()

	r.GET("/users", func(ctx *gin.Context) { getUsers(ctx, dbPool) })

	r.Run()
}
