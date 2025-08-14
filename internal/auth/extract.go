package auth

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
)

func GoogleAuthConfig() *oauth2.Config { return googleOAuthConfig }

func ExtractClaimsFromRequest(r *http.Request) (*SessionClaims, bool) {
	cookie, err := r.Cookie(os.Getenv("SESSION_COOKIE_NAME"))
	if err != nil {
		return nil, false
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &SessionClaims{}, func(token *jwt.Token) (interface{}, error) {
		return sessionSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, false
	}

	if claims, ok := token.Claims.(*SessionClaims); ok {
		return claims, true
	}
	return nil, false
}