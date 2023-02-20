package repository

import "gorm.io/gorm"

// Favorite 存储点赞信息，表示user点赞了video。
type Favorite struct {
	gorm.Model
	UserId  uint
	VideoId uint
}

func GetFavoriteCountById(id uint) uint {
	var cnt int64
	DB.Model(Favorite{}).Where("user_id = ?", id).Count(&cnt)
	return uint(cnt)
}

func GetFavoriteById(id uint) []Favorite {
	var favoriteList []Favorite
	DB.Model(Favorite{}).Where("user_id = ?", id).Find(&favoriteList)
	return favoriteList
}

// IsBFavoriteA 用户B是否给视频A点赞
func IsBFavoriteA(idA uint, idB uint) bool {
	if idB == 0 {
		return false
	} else {
		var isFavorite Favorite
		result := DB.Model(Favorite{}).Where("user_id = ? AND video_id = ?", idB, idA).First(&isFavorite)
		if result.Error != nil {
			return false
		} else {
			return true
		}
	}
	return false
}

func AddFavorite(idA uint, idB uint) error {

	newFavorite := Favorite{
		UserId:  idB,
		VideoId: idA,
	}
	// 只有登录用户才可以实现点赞操作
	if idB == 0 {
		return nil
	} else {
		if err := DB.Create(&newFavorite).Error; err != nil {
			return err
		}
	}
	return nil
}

func DeleteFavorite(idA uint, idB uint) error {

	oldFavorite := Favorite{}

	// 只有登录用户才可以实现取消点赞操作
	if idB == 0 {
		return nil
	} else if IsBFavoriteA(idA, idB) {
		DB.Where("user_id = ? AND video_id = ?", idB, idA).Take(&oldFavorite)
		DB.Delete(&oldFavorite)
	}
	return nil
}
