package user

import (
	"Douyin/common"
	"Douyin/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type InfoResponse struct {
	common.Response
	User common.User `json:"user"`
}

func Info(c *gin.Context) {

	//方法1：从token中解析出id，解析的步骤已经在中间件中写好，直接调用get方法即可
	//idToken, ok := c.Get("user_id")
	//id, err := idToken.(uint)

	//方法2：直接调用query方法得到id，代码如下
	idStr := c.Query("user_id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, InfoResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	// 因为要返回is_follow的信息，所以接口传入了两个id
	userInfo, err := user.GetUserInfo(uint(id), uint(id))
	if err != nil {
		c.JSON(http.StatusOK, InfoResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	c.JSON(http.StatusOK, InfoResponse{
		Response: common.Response{StatusCode: 0},
		User:     userInfo,
	})
	return
}
