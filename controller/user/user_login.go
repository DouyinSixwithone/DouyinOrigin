package user

import (
	"Douyin/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginResponse struct {
	common.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if user, exist := UsersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, LoginResponse{
			Response: common.Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, LoginResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
