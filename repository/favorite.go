package repository

import "gorm.io/gorm"

// Favorite 存储点赞信息，表示user点赞了video。
type Favorite struct {
	gorm.Model
	UserId  uint
	VideoId uint
}
