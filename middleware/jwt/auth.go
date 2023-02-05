package jwt

import (
	"Douyin/common"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var invalidResp = common.Response{
	StatusCode: 1,
	StatusMsg:  "Invalid token",
}

// AuthWithLogin 鉴权中间件，为登录用户鉴权并设置user_id
func AuthWithLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Query("token")
		if tokenStr == "" {
			tokenStr = c.PostForm("token")
		}
		//没有获取到token
		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, invalidResp)
			c.Abort()
			return
		}

		// 解析token
		token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, invalidResp)
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*MyClaims)
		if !ok || !token.Valid || time.Now().Unix() > claims.ExpiresAt {
			c.JSON(http.StatusUnauthorized, invalidResp)
			c.Abort()
			return
		}
		c.Set("user_id", claims.UserId)
		c.Next()
	}
}
