package token_jwt

import "github.com/dgrijalva/jwt-go"

// jwtCustomClaims are custom claims extending default ones.
type Claims struct {
	ID       int32  `json:"id"`
	Identity string `json:"identity"`
	jwt.StandardClaims
}

const Key = "7d6543d7862a07edf7902086f39b4b9a"
