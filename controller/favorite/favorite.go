package favorite

import (
	"Douyin/common"
	"Douyin/controller/publish"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Action no practical effect, just check if token is valid
func Action(c *gin.Context) {
	token := c.Query("token")

	if _, exist := common.UsersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, common.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// List all users have same favorite video list
func List(c *gin.Context) {
	c.JSON(http.StatusOK, publish.VideoListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		VideoList: common.DemoVideos,
	})
}
