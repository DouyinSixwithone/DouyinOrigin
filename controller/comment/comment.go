package comment

import (
	"Douyin/common"
	"Douyin/service/comment"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ActionResponse struct {
	common.Response
	Comment common.Comment `json:"comment,omitempty"`
}

type ListResponse struct {
	common.Response
	CommentList []common.Comment `json:"comment_list,omitempty"`
}

func Action(c *gin.Context) {

	//从token中解析出id，得到的是当前登录用户的id
	idToken, _ := c.Get("user_id")
	userId, ok := idToken.(uint)
	if !ok {
		c.JSON(http.StatusOK, ActionResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  "invalid user id",
			},
		})
		return
	}

	var strVideoId = c.Query("video_id")
	videoId, err := strconv.ParseUint(strVideoId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, ActionResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	var strActionType = c.Query("action_type")
	actionType, err := strconv.ParseUint(strActionType, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, ActionResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	if actionType == 1 {
		content := c.Query("comment_text")
		newComment, err := comment.Send(userId, uint(videoId), content)
		if err != nil {
			c.JSON(http.StatusOK, ActionResponse{
				Response: common.Response{
					StatusCode: 1,
					StatusMsg:  err.Error(),
				},
			})
			return
		}
		c.JSON(http.StatusOK, ActionResponse{
			Response: common.Response{
				StatusCode: 0,
				StatusMsg:  "add comment successfully",
			},
			Comment: newComment,
		})
	} else {
		commentId, err := strconv.ParseUint(c.Query("comment_id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, ActionResponse{
				Response: common.Response{
					StatusCode: 1,
					StatusMsg:  err.Error(),
				},
			})
			return
		}
		err = comment.Delete(uint(commentId))
		if err != nil {
			c.JSON(http.StatusOK, ActionResponse{
				Response: common.Response{
					StatusCode: 1,
					StatusMsg:  err.Error(),
				},
			})
			return
		}
		c.JSON(http.StatusOK, ActionResponse{
			Response: common.Response{
				StatusCode: 0,
				StatusMsg:  "delete successfully",
			},
		})
	}
}

func List(c *gin.Context) {
	videoId, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, ListResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	//调用service层评论函数
	commentList, err := comment.GetList(uint(videoId))
	if err != nil {
		c.JSON(http.StatusOK, ListResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	c.JSON(http.StatusOK, ListResponse{
		Response: common.Response{
			StatusCode: 0,
			StatusMsg:  "查表成功！",
		},
		CommentList: commentList,
	})
	return
}
