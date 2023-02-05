package user

import (
	"Douyin/common"
	"Douyin/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginResponse struct {
	common.Response
	UserId uint   `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if u, exist := user.UsersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, LoginResponse{
			Response: common.Response{StatusCode: 0},
			UserId:   u.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, LoginResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
