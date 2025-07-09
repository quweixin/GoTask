package forms

import "github.com/golang-jwt/jwt/v5"

var JWTSecret = []byte("your_secret_key")

type CustomClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
