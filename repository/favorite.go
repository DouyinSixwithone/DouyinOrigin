package repository

import "gorm.io/gorm"

// Favorite 存储点赞信息，表示user点赞了video。
type Favorite struct {
	gorm.Model
	UserId  uint
	VideoId uint
}

func GetFavoriteCountById(id uint) uint {
	return 0
}

// IsBFavoriteA 用户B是否给视频A点赞
func IsBFavoriteA(idA uint, idB uint) bool {
	if idB == 0 {
		return false
	}
	// 未完善
	return false
}
