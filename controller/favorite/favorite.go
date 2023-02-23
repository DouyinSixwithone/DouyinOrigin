package favorite

import (
	"Douyin/common"
	"Douyin/service/favorite"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ListResponse struct {
	common.Response
	List []common.Video `json:"video_list"`
}

// Action no practical effect
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

	var strVideoId = c.Query("video_id")
	VideoId, err := strconv.ParseUint(strVideoId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, ListResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	err = favorite.FavoriteAction(userId, uint(VideoId), uint(actionType))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 0,
			StatusMsg:  "success!",
		})
	}

}

// List all users have same favorite video list
func List(c *gin.Context) {

	//从token中解析出id，得到的是当前登录用户的id
	idToken, _ := c.Get("user_id")
	userHostId, ok := idToken.(uint)
	if !ok {
		c.JSON(http.StatusOK, ListResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  "invalid user id",
			},
		})
		return
	}

	var strUserId = c.Query("user_id")
	userId, err := strconv.ParseUint(strUserId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, ListResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	userNewId := uint(userId)
	if userNewId == 0 {
		userNewId = userHostId
	}

	favoriteList := favorite.GetFavoriteList(userNewId, userHostId)

	if err != nil {
		c.JSON(http.StatusBadRequest, ListResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  "查表失败！",
			},
			List: nil,
		})
	} else {
		c.JSON(http.StatusOK, ListResponse{
			Response: common.Response{
				StatusCode: 0,
				StatusMsg:  "查表成功！",
			},
			List: favoriteList,
		})
	}
}
