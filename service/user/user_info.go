package user

import (
	"Douyin/common"
	"Douyin/repository"
	"errors"
	"fmt"
)

// GetUserInfo 传入两个id，表示查询user的信息以及follower是否关注了user
func GetUserInfo(userId uint, followerId uint) (common.User, error) {
	//1.参数合法性检验
	err := checkUserInfo(userId, followerId)
	if err != nil {
		return common.User{}, err
	}

	//2.得到需要的信息
	name, err := repository.GetNameById(userId)
	if err != nil {
		return common.User{}, err
	}
	followCount, err := repository.GetFollowCountById(userId)
	if err != nil {
		return common.User{}, err
	}
	followerCount, err := repository.GetFollowerCountById(userId)
	if err != nil {
		return common.User{}, err
	}
	isFollow, err := repository.IsBFollowA(userId, followerId)
	if err != nil {
		return common.User{}, err
	}

	return common.User{
		Id:            userId,
		Name:          name,
		FollowCount:   followCount,
		FollowerCount: followerCount,
		IsFollow:      isFollow,
	}, nil
}

func checkUserInfo(userId uint, followerId uint) error {
	if !repository.IsUserExistById(userId) {
		return errors.New(fmt.Sprintf("user(id=%v) doesn't exist", userId))
	}
	if !repository.IsUserExistById(followerId) {
		return errors.New(fmt.Sprintf("user(id=%v) doesn't exist", followerId))
	}
	return nil
}
