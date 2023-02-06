package feed

import (
	"Douyin/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ListResponse struct {
	common.Response
	VideoList []common.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

// List same demo video list for every request
func List(c *gin.Context) {

	c.JSON(http.StatusOK, ListResponse{
		Response:  common.Response{StatusCode: 0},
		VideoList: common.DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
