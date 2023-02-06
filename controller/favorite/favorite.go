package favorite

import (
	"Douyin/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListResponse struct {
	common.Response
	List []common.Video `json:"video_list"`
}

// Action no practical effect
func Action(c *gin.Context) {
	c.JSON(http.StatusOK, common.Response{StatusCode: 0})
}

// List all users have same favorite video list
func List(c *gin.Context) {
	c.JSON(http.StatusOK, ListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		List: common.DemoVideos,
	})
}
