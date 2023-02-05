package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// CreateToken 根据id生成Token
func CreateToken(id uint) (string, error) {
	expireTime := time.Now().Add(24 * time.Hour) //过期时间
	nowTime := time.Now()                        //当前时间
	claims := MyClaims{
		UserId: id,
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

// ParseToken 解析token
func ParseToken(tokenStr string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err == nil && token != nil {
		if claim, ok := token.Claims.(*MyClaims); ok && token.Valid {
			return claim, nil
		}
	}
	return nil, errors.New("invalid token")
}
