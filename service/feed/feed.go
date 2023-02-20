package feed

import (
	"Douyin/common"
	"Douyin/repository"
	"Douyin/service/user"
	"fmt"
	"time"
)

func GetFeedList(guestId uint, lastTime int64) ([]common.Video, int64) {
	var feedVideoList []common.Video
	feedVideoList = make([]common.Video, 0)

	videoList := FeedGet(lastTime)
	var newTime int64 = 0
	for _, x := range videoList {
		var tmp common.Video
		tmp.Id = x.ID
		tmp.PlayUrl = x.PlayUrl
		tmp.CoverUrl = x.CoverUrl
		tmp.Title = x.Title

		feedUser, _ := user.GetUserInfo(x.AuthorId, guestId)
		tmp.Author = feedUser

		tmp.FavoriteCount = repository.GetFavoriteCountById(tmp.Id)
		tmp.CommentCount = repository.GetCommentCountById(tmp.Id)
		tmp.IsFavorite = repository.IsBFavoriteA(tmp.Id, guestId)

		feedVideoList = append(feedVideoList, tmp)
		newTime = x.CreatedAt.Unix()
	}
	return feedVideoList, newTime
}

func FeedGet(lastTime int64) []repository.Video {
	var strTime time.Time
	if lastTime == 1676733878 {
		strTime = time.Now()
	} else {
		strTime = time.Unix(lastTime/1000, 0)
	}

	// strTime := fmt.Sprint(time.Unix(lastTime, 0).Format("2006-01-02 15:04:05"))
	fmt.Println("查询到的时间", strTime)

	var videoList = repository.GetVideoListByTime(strTime)
	return videoList
}
