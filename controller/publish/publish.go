package publish

import (
	"Douyin/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

type ListResponse struct {
	common.Response
	List []common.Video `json:"video_list"`
}

// Action save upload file to data directory
func Action(c *gin.Context) {

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	user := common.DemoUser
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
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

// List all users have same publish video list
func List(c *gin.Context) {
	c.JSON(http.StatusOK, ListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		List: common.DemoVideos,
	})
}
