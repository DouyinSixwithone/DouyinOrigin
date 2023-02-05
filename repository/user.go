package repository

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string `json:"name"`
	Password      string `json:"password"`
	FollowCount   uint   `json:"follow_count"`
	FollowerCount uint   `json:"follower_count"`
}

func IsUserExistByName(name string) bool {
	var userExist = &User{}
	result := DB.Model(User{}).Where("name=?", name).First(&userExist)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// 如果返回记录不存在的错误，说明没有名为name的用户
		return false
	}
	return true
}

func InsertNewUser(u *User) error {
	if err := DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func GetIdByName(name string) (uint, error) {
	var userId = &User{}
	err := DB.Model(User{}).Where("name=?", name).First(&userId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}
	return userId.ID, nil
}

func GetPassByName(name string) (string, error) {
	var userId = &User{}
	err := DB.Model(User{}).Where("name=?", name).First(&userId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", err
	}
	return userId.Password, nil
}
