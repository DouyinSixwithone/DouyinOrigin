package repository

import "gorm.io/gorm"

// Video 存储Video的基本信息，以gorm.Model.ID作为视频的id；
// 表中未存储点赞数、评论数和用户的具体信息，可通过调用其他文件中的接口获取；
// eg.在repository/comment文件中定义GetCommentCountById函数；调用service/user中的GetUserInfo函数得到用户信息。
type Video struct {
	gorm.Model
	AuthorId uint
	PlayUrl  string
	CoverUrl string
	title    string
}
