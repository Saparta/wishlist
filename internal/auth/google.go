package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOAuthConfig *oauth2.Config

func InitGoogleOAuth() {
	googleOAuthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes:       []string{"openid", "email", "profile"},
		Endpoint:     google.Endpoint,
	}
}

func RandomState() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

// Set a short-lived state cookie to mitigate CSRF on the OAuth dance
func SetStateCookie(w http.ResponseWriter, state string) {
	c := &http.Cookie{
		Name:     "oauth_state",
		Value:    state,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // true in prod behind HTTPS
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(10 * time.Minute),
	}
	http.SetCookie(w, c)
}

func GoogleAuthConfig() *oauth2.Config { return googleOAuthConfig }