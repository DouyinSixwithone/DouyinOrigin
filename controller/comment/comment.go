package comment

import (
	"Douyin/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ActionResponse struct {
	common.Response
	Comment common.Comment `json:"comment,omitempty"`
}

type ListResponse struct {
	common.Response
	List []common.Comment `json:"comment_list,omitempty"`
}

// Action no practical effect
func Action(c *gin.Context) {
	actionType := c.Query("action_type")
	if actionType == "1" {
		text := c.Query("comment_text")
		c.JSON(http.StatusOK, ActionResponse{Response: common.Response{StatusCode: 0},
			Comment: common.Comment{
				Id:         1,
				User:       common.DemoUser,
				Content:    text,
				CreateDate: "05-01",
			}})
		return
	}
	c.JSON(http.StatusOK, common.Response{StatusCode: 0})
}

// List all videos have same demo comment list
func List(c *gin.Context) {
	c.JSON(http.StatusOK, ListResponse{
		Response: common.Response{StatusCode: 0},
		List:     common.DemoComments,
	})
}
