package feed

import (
	"Douyin/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Response struct {
	common.Response
	VideoList []common.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

// Action same demo video list for every request
func Action(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Response:  common.Response{StatusCode: 0},
		VideoList: common.DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
