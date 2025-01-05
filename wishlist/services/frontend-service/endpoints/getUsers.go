package endpoints

import (
	"context"
	"log"
	"net/http"

	"github.com/Saparta/wishlist/wishlist/services/frontend-service/models"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/gin-gonic/gin"
)

func GetUsers(ctx *gin.Context, dbpool *pgxpool.Pool) {
	rows, err := dbpool.Query(context.Background(), "SELECT id, oauth_id, email, name, created_at FROM users")
	if err != nil {
		log.Printf("Query error: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	defer rows.Close()

	// Collect the users into a slice
	var users []models.User
	for rows.Next() {
		var user models.User
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
