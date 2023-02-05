package publish

import (
	"Douyin/common"
	"Douyin/service/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

type VideoListResponse struct {
	common.Response
	VideoList []common.Video `json:"video_list"`
}

// Publish check token then save upload file to data directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	if _, exist := user.UsersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	u := user.UsersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", u.Id, filename)
	saveFile := filepath.Join("./data/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, common.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		VideoList: common.DemoVideos,
	})
}
