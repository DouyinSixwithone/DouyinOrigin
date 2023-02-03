package feed

import (
	"Douyin/controller"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	controller.Response
	VideoList []controller.Video `json:"video_list,omitempty"`
	NextTime  int64              `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	c.JSON(http.StatusOK, FeedResponse{
		Response:  controller.Response{StatusCode: 0},
		VideoList: controller.DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
