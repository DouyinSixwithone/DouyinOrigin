package user

import (
	"Douyin/common"
	"Douyin/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InfoResponse struct {
	common.Response
	User user.Info `json:"user"`
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if u, exist := user.UsersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, InfoResponse{
			Response: common.Response{StatusCode: 0},
			User:     u,
		})
	} else {
		c.JSON(http.StatusOK, InfoResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
