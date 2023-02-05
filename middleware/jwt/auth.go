package jwt

import (
	"Douyin/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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
		tokenStr = strings.Fields(tokenStr)[1]

		//解析token并验证token是否超时
		token, err := ParseToken(tokenStr)
		if err != nil || time.Now().Unix() > token.ExpiresAt {
			c.JSON(http.StatusUnauthorized, invalidResp)
			c.Abort()
			return
		}
		c.Set("user_id", token.UserId)
		c.Next()
	}
}

// AuthWithoutLogin 鉴权中间件，若携带token则解析出user_id，若未携带则放入默认值0
func AuthWithoutLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userId uint
		tokenStr := c.Query("token")
		if tokenStr == "" {
			tokenStr = c.PostForm("token")
		}
		//没有获取到token
		if tokenStr == "" {
			userId = 0
		} else {
			tokenStr = strings.Fields(tokenStr)[1]
			token, err := ParseToken(tokenStr)
			if err != nil || time.Now().Unix() > token.ExpiresAt {
				userId = 0
			} else {
				userId = token.UserId
			}
		}
		c.Set("user_Id", userId)
		c.Next()
	}
}
