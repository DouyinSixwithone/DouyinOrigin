package user

import (
	"Douyin/common"
	"Douyin/repository"
	"errors"
	"fmt"
)

// GetUserInfo 传入两个id，表示查询user的信息以及follower是否关注了user
func GetUserInfo(userId uint, followerId uint) (common.User, error) {
	// 1. 参数合法性检验
	err := checkUserInfo(userId, followerId)
	if err != nil {
		return common.User{}, err
	}
	// 2. 获取信息并返回
	return common.User{
		Id:            userId,
		Name:          repository.GetNameById(userId),
		FollowCount:   repository.GetFollowCountById(userId),
		FollowerCount: repository.GetFollowerCountById(userId),
		IsFollow:      repository.IsBFollowA(userId, followerId),
	}, nil
}

func checkUserInfo(userId uint, followerId uint) error {
	if !repository.IsUserExistById(userId) {
		return errors.New(fmt.Sprintf("user(id=%v) doesn't exist", userId))
	}
	if !repository.IsUserExistById(followerId) && followerId != 0 {
		return errors.New(fmt.Sprintf("user(id=%v) doesn't exist", followerId))
	}
	return nil
}
