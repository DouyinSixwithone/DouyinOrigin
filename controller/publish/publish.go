package publish

import (
	"Douyin/controller"
	"Douyin/controller/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

type VideoListResponse struct {
	controller.Response
	VideoList []controller.Video `json:"video_list"`
}

// Publish check token then save upload file to data directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	if _, exist := user.UsersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, controller.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, controller.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	user := user.UsersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./data/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, controller.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, controller.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: controller.Response{
			StatusCode: 0,
		},
		VideoList: controller.DemoVideos,
	})
}
