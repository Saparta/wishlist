package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var sessionSecret = []byte(os.Getenv("SESSION_SECRET"))
var cookieName = os.Getenv("SESSION_COOKIE_NAME")

type SessionClaims struct {
	UserID string `json:"uid"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func SetSessionCookie(w http.ResponseWriter, userID, email string) error {
	claims := SessionClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(sessionSecret)
	if err != nil { return err }

	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    signed,
		Path:     "/",
		HttpOnly: true,
		Secure:   os.Getenv("SESSION_COOKIE_SECURE") == "true",
		SameSite: http.SameSiteLaxMode, // configurable
		Expires:  time.Now().Add(30 * 24 * time.Hour),
		Domain:   os.Getenv("SESSION_COOKIE_DOMAIN"),
	}
	http.SetCookie(w, cookie)
	return nil
}

func ClearSessionCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0,0),
		MaxAge:   -1,
		HttpOnly: true,
	})
}