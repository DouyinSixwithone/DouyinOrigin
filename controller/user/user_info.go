package user

import (
	"Douyin/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserResponse struct {
	common.Response
	User common.User `json:"user"`
}

// UsersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var UsersLoginInfo = map[string]common.User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if user, exist := UsersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: common.Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
