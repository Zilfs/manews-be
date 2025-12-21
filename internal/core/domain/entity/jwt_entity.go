package entity

import "github.com/golang-jwt/jwt/v4"

type JwtData struct {
	UserID float64 `json:"user_id"`
	jwt.RegisteredClaims
}
