package repository

import "gorm.io/gorm"

// Relation 存储关注信息，表示follower关注了user。
type Relation struct {
	gorm.Model
	UserId     uint
	FollowerId uint
}

func GetFollowCountById(id uint) (uint, error) {
	return 0, nil
}

func GetFollowerCountById(id uint) (uint, error) {
	return 0, nil
}

func IsBFollowA(idA uint, idB uint) (bool, error) {
	if idA == idB {
		return false, nil
	}
	// 未完善
	return false, nil
}
