package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/Saparta/wishlist/wishlist/services/user-service/endpoints"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// var googleAuthConfig = &oauth2.Config{}

func createUsersTable(dbpool *pgxpool.Pool) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY,
		oauth_id UUID NOT NULL,
		email TEXT NOT NULL,
		name TEXT NOT NULL,
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
	err = createUsersTable(dbPool)
	if err != nil {
		log.Fatalf("Failed to create users table: %v\n", err)
	}

	channel <- dbPool
}

// func GetUserInfo(accessToken string) (map[string]interface{}, error) {
// 	userInfoEndpoint := "https://www.googleapis.com/oauth2/v2/userinfo"
// 	resp, err := http.Get(fmt.Sprintf("%s?access_token=%s", userInfoEndpoint, accessToken))
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	var userInfo map[string]interface{}
// 	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
// 		return nil, err
// 	}

// 	return userInfo, nil
// }

// func SignJWT(userInfo map[string]interface{}) (string, error) {
// 	// Customize the claims as needed
// 	claims := jwt.MapClaims{
// 		"sub":   userInfo["id"],
// 		"name":  userInfo["name"],
// 		"email": userInfo["email"],
// 		"iss":   "oauth-app-golang",
// 		"exp":   time.Now().Add(time.Hour * 24 * 5).Unix(), // 5 days
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	signedToken, err := token.SignedString([]byte("your-secret-key")) // Replace with your actual secret key
// 	if err != nil {
// 		return "", err
// 	}

// 	return signedToken, nil
// }

// func googleCallbackHandler(ctx *gin.Context) {
// 	code := ctx.Query("code")
// 	token, err := googleAuthConfig.Exchange(context.Background(), code)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	userInfo, err := GetUserInfo(token.AccessToken)
// 	if err != nil {
// 		fmt.Println("Error getting user info: " + err.Error())
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	signedToken, err := SignJWT(userInfo)
// 	if err != nil {
// 		fmt.Println("Error signing token: " + err.Error())
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"token": signedToken})
// }

// func init() {
// 	googleAuthConfig = &oauth2.Config{
// 		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
// 		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
// 		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"), // Should be a frontend url
// 		Scopes:       []string{"profile", "email"},
// 		Endpoint:     google.Endpoint,
// 	}
// }

func main() {
	var dbChannel chan *pgxpool.Pool = make(chan *pgxpool.Pool)
	godotenv.Load()
	go setUpDb(dbChannel)

	var r *gin.Engine = gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	var dbPool *pgxpool.Pool = <-dbChannel
	defer dbPool.Close()

	r.Use(cors.Default())
	r.GET("/users", func(ctx *gin.Context) { endpoints.GetUsers(ctx, dbPool) })
	// r.GET("/auth/google/callback", googleCallbackHandler)

	r.Run()
}
