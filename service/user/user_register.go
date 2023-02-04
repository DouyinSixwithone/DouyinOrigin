package user

import (
	"Douyin/common"
	"Douyin/middleware/jwt"
	"Douyin/repository"
	"errors"
)

const (
	MaxUsernameLength = 32 //用户名最大长度
	MaxPasswordLength = 32 //密码最大长度
	MinPasswordLength = 6  //密码最小长度
)

type RegisterInfo struct {
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

func GetRegisterInfo(username string, password string) (RegisterInfo, error) {

	//1.合法性检验
	err := checkParam(username, password)
	if err != nil {
		return RegisterInfo{}, err
	}

	//2.新建用户
	u, err := createUser(username, password)
	if err != nil {
		return RegisterInfo{}, err
	}

	//3.获得token
	token, err := jwt.CreateToken(u)
	if err != nil {
		return RegisterInfo{}, err
	}

	return RegisterInfo{
		UserId: u.Id,
		Token:  token,
	}, nil
}

func checkParam(username string, password string) error {
	u := repository.GetUserbyName(username)
	if username == u.Name {
		return errors.New("user name already exist")
	}
	if username == "" {
		return errors.New("user name is empty")
	}
	if len(username) > MaxUsernameLength {
		return errors.New("user name length exceeds the limit")
	}
	if len(password) < MinPasswordLength {
		return errors.New("password is too short")
	}
	if len(password) > MaxPasswordLength {
		return errors.New("password length exceeds the limit")
	}
	return nil
}

func createUser(username string, password string) (common.User, error) {
	return common.User{}, nil
}
