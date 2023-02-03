package relation

import (
	"Douyin/controller"
	"Douyin/controller/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserListResponse struct {
	controller.Response
	UserList []controller.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := user.UsersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, controller.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, controller.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: controller.Response{
			StatusCode: 0,
		},
		UserList: []controller.User{controller.DemoUser},
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: controller.Response{
			StatusCode: 0,
		},
		UserList: []controller.User{controller.DemoUser},
	})
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: controller.Response{
			StatusCode: 0,
		},
		UserList: []controller.User{controller.DemoUser},
	})
}
