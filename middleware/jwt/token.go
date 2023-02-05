package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	UserId   uint   `json:"user_id"`
	UserName string `json:"username"`
	jwt.StandardClaims
}

var jwtKey = []byte("Douyin_Sixwithone")

func CreateToken(id uint, name string) (string, error) {
	expireTime := time.Now().Add(24 * time.Hour) //过期时间
	nowTime := time.Now()                        //当前时间
	claims := MyClaims{
		UserId:   id,
		UserName: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间戳
			IssuedAt:  nowTime.Unix(),    //当前时间戳
			Issuer:    "douyin",          //颁发者签名
			Subject:   "userToken",       //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
