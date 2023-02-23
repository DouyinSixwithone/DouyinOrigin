package comment

import (
	"Douyin/repository"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserId     uint
	VideoId    uint64
	Content    string
	CreateDate string
}

func CountFromVideoId(videoId uint) uint {
	cnt := repository.GetCommentCountById(videoId)
	return cnt
}

func Send(comment Comment) (repository.Comment, error) {
	var commentInfo repository.Comment
	commentInfo.VideoId = comment.VideoId       //评论视频id传入
	commentInfo.UserId = comment.UserId         //评论用户id传入
	commentInfo.Content = comment.Content       //评论内容传入
	commentInfo.CreateDate = comment.CreateDate //评论时间
	commentRtn, err := repository.InsertComment(commentInfo)
	if err != nil {
		return repository.Comment{}, err
	}
	return commentRtn, nil
}

func Delete(id int64) error {
	err := repository.DeleteComment(id)
	if err != nil {
		return err
	}
	return nil
}

func GetList(videoId int64) ([]repository.Comment, error) {
	return repository.GetCommentList(videoId)
}
