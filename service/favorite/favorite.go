package favorite

import (
	"Douyin/common"
	"Douyin/repository"
	"Douyin/service/user"
)

func FavoriteAction(userId uint, VideoId uint, actionType uint) (err error) {

	isFavorite := repository.IsBFavoriteA(VideoId, userId)

	// 点赞操作
	if actionType == 1 {
		if !isFavorite {
			err := repository.AddFavorite(VideoId, userId)
			if err != nil {
				return err
			}
		}
		// 取消点赞操作
	} else {
		if isFavorite {
			err := repository.DeleteFavorite(VideoId, userId)
			if err != nil {
				return err
			}
		}
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

		tmp.FavoriteCount = repository.GetFavoriteCountById(tmp.Id)
		tmp.CommentCount = repository.GetCommentCountById(tmp.Id)
		tmp.IsFavorite = repository.IsBFavoriteA(tmp.Id, userHostId)

		favoriteVideoList = append(favoriteVideoList, tmp)
	}
	return favoriteVideoList

}

func FavoriteGet(userId uint) []repository.Video {

	var favoriteList = repository.GetFavoriteById(userId)

	videoList := make([]repository.Video, 0)

	for _, v := range favoriteList {
		var video = repository.GetVideoById(v.VideoId)
		videoList = append(videoList, video)
	}

	return videoList
}
