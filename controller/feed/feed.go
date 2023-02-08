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
// 注意：如果用户已登录，推送的视频流中不应该包含本人发布的视频
func List(c *gin.Context) {

	c.JSON(http.StatusOK, ListResponse{
		Response:  common.Response{StatusCode: 0},
		VideoList: common.DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
