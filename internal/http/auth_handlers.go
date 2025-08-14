package httpapi

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"wishlist/services/frontend-service/internal/auth"
	"wishlist/services/frontend-service/internal/store"
)

type MeResponse struct {
	User *store.User `json:"user"`
}

func RegisterAuthRoutes(r *gin.Engine, db *store.Store) {
	// CORS for frontend
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONTEND_ORIGIN"))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PATCH,DELETE,OPTIONS")
		if c.Request.Method == http.MethodOptions {
			c.Status(http.StatusNoContent); c.Abort(); return
		}
		c.Next()
	})

	g := r.Group("/api/auth")
	g.GET("/google/start", func(c *gin.Context) {
		state, err := auth.RandomState()
		if err != nil { c.String(500, "state err"); return }
		auth.SetStateCookie(c.Writer, state)
		url := auth.GoogleAuthConfig().AuthCodeURL(state, oauth2.AccessTypeOnline)
		c.Redirect(http.StatusTemporaryRedirect, url)
	})

	g.GET("/google/callback", func(c *gin.Context) {
		state := c.Query("state")
		code := c.Query("code")
		if state == "" || code == "" { c.String(400, "missing state/code"); return }

		// Validate state cookie
		sc, _ := c.Request.Cookie("oauth_state")
		if sc == nil || sc.Value != state { c.String(400, "bad state"); return }

		tok, err := auth.GoogleAuthConfig().Exchange(context.Background(), code)
		if err != nil { c.String(400, "token exchange failed"); return }

		// Get userinfo via https://www.googleapis.com/oauth2/v3/userinfo
		// Note: You can also decode id_token; userinfo is straightforward.
		client := auth.GoogleAuthConfig().Client(context.Background(), tok)
		resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
		if err != nil { c.String(400, "userinfo failed"); return }
		defer resp.Body.Close()

		var ui struct{
			Sub string `json:"sub"`
			Email string `json:"email"`
			Name string `json:"name"`
			Picture string `json:"picture"`
			EmailVerified bool `json:"email_verified"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&ui); err != nil { c.String(400, "decode failed"); return }

		// Upsert user in DB
		user, err := db.FindOrCreateUserByEmail(c, ui.Email, ui.Name, ui.Picture)
		if err != nil { c.String(500, "db err"); return }

		// Set signed session cookie
		if err := auth.SetSessionCookie(c.Writer, user.ID, user.Email); err != nil {
			c.String(500, "session err"); return
		}

		// Redirect back to frontend home
		c.Redirect(http.StatusTemporaryRedirect, os.Getenv("FRONTEND_ORIGIN")+"/")
	})

	g.GET("/me", func(c *gin.Context) {
		claims, ok := auth.ExtractClaimsFromRequest(c.Request)
		if !ok { c.JSON(200, MeResponse{User: nil}); return }

		user, err := db.GetUserByID(c, claims.UserID)
		if err != nil { c.JSON(200, MeResponse{User: nil}); return }

		c.JSON(200, MeResponse{User: user})
	})

	g.POST("/logout", func(c *gin.Context) {
		auth.ClearSessionCookie(c.Writer)
		c.Status(204)
	})
}