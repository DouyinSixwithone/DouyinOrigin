package comment

import (
	"Douyin/common"
	"Douyin/repository"
	"Douyin/service/user"
)

func Send(userId uint, videoId uint, content string) (common.Comment, error) {
	newComment := repository.Comment{
		UserId:  userId,
		VideoId: videoId,
		Content: content,
	}
	err := repository.InsertComment(&newComment)
	if err != nil {
		return common.Comment{}, err
	}
	u, _ := user.GetUserInfo(userId, userId)
	retComment := common.Comment{
		Id:         newComment.ID,
		User:       u,
		Content:    content,
		CreateDate: newComment.CreatedAt.Format("01-02"),
	}
	return retComment, nil
}

func Delete(id uint) error {
	err := repository.DeleteComment(id)
	if err != nil {
		return err
	}
	return nil
}

func GetList(videoId uint) ([]common.Comment, error) {
	commentList, err := repository.GetCommentList(videoId)
	if err != nil {
		return nil, err
	}
	var retList []common.Comment
	for _, x := range commentList {
		u, err := user.GetUserInfo(x.UserId, x.UserId)
		if err != nil {
			return nil, err
		}
		comment := common.Comment{
			Id:         x.ID,
			User:       u,
			Content:    x.Content,
			CreateDate: x.CreatedAt.Format("01-02"),
		}
		retList = append(retList, comment)
	}
	return retList, nil
}
