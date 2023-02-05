package repository

import "gorm.io/gorm"

// Relation 存储关注信息，表示Follower关注User
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
	return false, nil
}
