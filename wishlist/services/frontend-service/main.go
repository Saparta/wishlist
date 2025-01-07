package main

import (
	"log"
	"net/http"

	"github.com/Saparta/wishlist/wishlist/services/frontend-service/db"
	"github.com/Saparta/wishlist/wishlist/services/frontend-service/endpoints"
	pb "github.com/Saparta/wishlist/wishlist/services/frontend-service/proto"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	serverAddr = "localhost:8081"
)

func main() {
	var dbChannel chan *pgxpool.Pool = make(chan *pgxpool.Pool)
	godotenv.Load()
	go db.SetUpDb(dbChannel)

	var r *gin.Engine = gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	var dbPool *pgxpool.Pool = <-dbChannel
	defer dbPool.Close()

	r.Use(cors.Default())
	// Users database calls
	r.GET("/users", func(ctx *gin.Context) { endpoints.GetUsers(ctx, dbPool) })
	// r.GET("/auth/google/callback", googleCallbackHandler)

	// Setup for wishlist servce gRPC calls
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewWishlistServiceClient(conn)

	// Wishlist service requests
	r.POST("/wishlist", func(ctx *gin.Context) { endpoints.MakeCreateWishlist(ctx, client) })
	r.POST("/wishlist", func(ctx *gin.Context) { endpoints.MakeAddWishlistItem(ctx, client) })
	r.GET("/wishlist/:userId", func(ctx *gin.Context) { endpoints.MakeGetUserWishlists(ctx, client) })

	r.Run()
}

// var googleAuthConfig = &oauth2.Config{}

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
