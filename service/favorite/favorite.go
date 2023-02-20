package favorite

import (
	"Douyin/common"
	"Douyin/repository"
	"Douyin/service/user"
	"errors"
)

func FavoriteAction(userId uint, VideoId uint, actionType uint) (err error) {
	if userId == 0 {
		return errors.New("user not logged in")
	}
	isFavorite := repository.IsBFavoriteA(VideoId, userId)
	if actionType == 1 && !isFavorite {
		// 点赞操作
		err := repository.AddFavorite(VideoId, userId)
		if err != nil {
			return err
		}
	} else if actionType == 0 && isFavorite {
		// 取消点赞操作
		err := repository.DeleteFavorite(VideoId, userId)
		if err != nil {
			return err
		}
	} else {
		return errors.New("invalid action")
	}
	return nil
}

func GetFavoriteList(userNewId uint, userHostId uint) []common.Video {

	var favoriteVideoList []common.Video
	favoriteVideoList = make([]common.Video, 0)

	videoList := FavoriteGet(userNewId)
	for _, x := range videoList {
		var tmp common.Video
		tmp.Id = x.ID
		tmp.PlayUrl = x.PlayUrl
		tmp.CoverUrl = x.CoverUrl
		tmp.Title = x.Title

		favoriteUser, _ := user.GetUserInfo(x.AuthorId, userHostId)
		tmp.Author = favoriteUser

		tmp.FavoriteCount = repository.GetFavoritedCountByVideoId(tmp.Id)
		tmp.CommentCount = repository.GetCommentCountById(tmp.Id)
		tmp.IsFavorite = repository.IsBFavoriteA(tmp.Id, userHostId)

		favoriteVideoList = append(favoriteVideoList, tmp)
	}
	return favoriteVideoList

}

func FavoriteGet(userId uint) []repository.Video {

	var favoriteList = repository.GetFavoriteByUserId(userId)

	videoList := make([]repository.Video, 0)

	for _, v := range favoriteList {
		var video = repository.GetVideoById(v.VideoId)
		videoList = append(videoList, video)
	}

	return videoList
}
