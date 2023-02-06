package jwt

import "github.com/dgrijalva/jwt-go"

type MyClaims struct {
	UserId uint `json:"user_id"`
	jwt.StandardClaims
}

var jwtKey = []byte("Douyin_Sixwithone")
