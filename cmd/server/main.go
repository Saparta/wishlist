package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"wishlist/services/frontend-service/internal/auth"
	"wishlist/services/frontend-service/internal/httpapi"
	"wishlist/services/frontend-service/internal/store"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	auth.InitGoogleOAuth()

	s := &store.Store{DB: db}
	r := gin.Default()
	httpapi.RegisterAuthRoutes(r, s)

	// health
	r.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	log.Println("listening on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}