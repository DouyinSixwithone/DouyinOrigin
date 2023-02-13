package publish

import (
	"Douyin/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListResponse struct {
	common.Response
	List []common.Video `json:"video_list"`
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
