package user

import (
	"Douyin/controller"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync/atomic"
)

var userIdSequence = int64(1)

type RegisterResponse struct {
	controller.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if _, exist := UsersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, LoginResponse{
			Response: controller.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		atomic.AddInt64(&userIdSequence, 1)
		newUser := controller.User{
			Id:   userIdSequence,
			Name: username,
		}
		UsersLoginInfo[token] = newUser
		c.JSON(http.StatusOK, LoginResponse{
			Response: controller.Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    username + password,
		})
	}
}
