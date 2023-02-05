package message

import (
	"Douyin/common"
	"Douyin/service/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"
)

var tempChat = map[string][]common.Message{}

var messageIdSequence = uint64(1)

type ChatResponse struct {
	common.Response
	MessageList []common.Message `json:"message_list"`
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	content := c.Query("content")

	if u, exist := user.UsersLoginInfo[token]; exist {
		userIdB, _ := strconv.Atoi(toUserId)
		chatKey := genChatKey(u.Id, uint(userIdB))

		atomic.AddUint64(&messageIdSequence, 1)
		curMessage := common.Message{
			Id:         uint(messageIdSequence),
			Content:    content,
			CreateTime: time.Now().Format(time.Kitchen),
		}

		if messages, exist := tempChat[chatKey]; exist {
			tempChat[chatKey] = append(messages, curMessage)
		} else {
			tempChat[chatKey] = []common.Message{curMessage}
		}
		c.JSON(http.StatusOK, common.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// MessageChat all users have same follow list
func MessageChat(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")

	if u, exist := user.UsersLoginInfo[token]; exist {
		userIdB, _ := strconv.Atoi(toUserId)
		chatKey := genChatKey(u.Id, uint(userIdB))

		c.JSON(http.StatusOK, ChatResponse{Response: common.Response{StatusCode: 0}, MessageList: tempChat[chatKey]})
	} else {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

func genChatKey(userIdA uint, userIdB uint) string {
	if userIdA > userIdB {
		return fmt.Sprintf("%d_%d", userIdB, userIdA)
	}
	return fmt.Sprintf("%d_%d", userIdA, userIdB)
}
