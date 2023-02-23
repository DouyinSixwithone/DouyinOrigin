package comment

import (
	"Douyin/common"
	"Douyin/service/comment"
	"Douyin/service/user"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusOK, ListResponse{
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
		c.JSON(http.StatusOK, ListResponse{
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
		c.JSON(http.StatusOK, ListResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	if actionType == 1 {
		content := c.Query("content_text")
		var sendComment comment.Comment
		sendComment.UserId = userId
		sendComment.VideoId = videoId
		sendComment.Content = content
		sendComment.CreateDate = time.Now().Format("2006-01-02 15:04:05")
		_, err := comment.Send(sendComment)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, common.Response{
				StatusCode: 0,
				StatusMsg:  "success!",
			})
			return
		}
	} else {
		commentId, err := strconv.ParseInt(c.Query("comment_id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, ListResponse{
				Response: common.Response{
					StatusCode: 1,
					StatusMsg:  err.Error(),
				},
			})
			return
		}
		err = comment.Delete(commentId)
		if err != nil {
			c.JSON(http.StatusOK, ListResponse{
				Response: common.Response{
					StatusCode: 1,
					StatusMsg:  err.Error(),
				},
			})
			return
		} else {
			c.JSON(http.StatusOK, ListResponse{
				Response: common.Response{
					StatusCode: 0,
					StatusMsg:  "delete successfully",
				},
			})
		}
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
	commentList, err := comment.GetList(videoId)
	if err != nil {
		c.JSON(http.StatusOK, ListResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	var idx = 0
	var List []common.Comment
	for _, comment := range commentList {
		List[idx].Id = comment.ID
		List[idx].Content = comment.Content
		List[idx].CreateDate = comment.CreateDate
		List[idx].User, _ = user.GetUserInfo(comment.UserId, comment.UserId)
		idx = idx + 1
	}
	c.JSON(http.StatusOK, ListResponse{
		Response: common.Response{
			StatusCode: 0,
			StatusMsg:  "查表成功！",
		},
		CommentList: List,
	})
	return
}
