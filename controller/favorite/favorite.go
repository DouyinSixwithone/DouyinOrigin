package favorite

import (
	"Douyin/controller"
	"Douyin/controller/publish"
	"Douyin/controller/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := user.UsersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, controller.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, controller.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, publish.VideoListResponse{
		Response: controller.Response{
			StatusCode: 0,
		},
		VideoList: controller.DemoVideos,
	})
}
