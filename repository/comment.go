package repository

import "gorm.io/gorm"

// Comment 存储评论信息，以gorm.Model.ID作为评论的id，表示user评论了video，内容为content；
// CreateDate可通过gorm.Model.CreatedAt获取。
type Comment struct {
	gorm.Model
	UserId  uint
	VideoId uint
	Content string
}

func GetCommentCountById(id uint) uint {
	return 0
}
