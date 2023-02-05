package user

import (
	"Douyin/middleware/jwt"
	"Douyin/repository"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type LoginInfo struct {
	UserId uint   `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

func GetLoginInfo(username string, password string) (LoginInfo, error) {

	//1.合法性检验
	err := checkLogin(username, password)
	if err != nil {
		return LoginInfo{}, err
	}

	//2.获得id
	id, err := repository.GetIdByName(username)
	if err != nil {
		return LoginInfo{}, err
	}

	//3.获得token
	token, err := jwt.CreateToken(id)
	if err != nil {
		return LoginInfo{}, err
	}

	return LoginInfo{
		UserId: id,
		Token:  token,
	}, nil
}

func checkLogin(name string, pass string) error {
	if name == "" {
		return errors.New("user name is empty")
	}
	if len(name) > MaxUsernameLength {
		return errors.New("user name length exceeds the limit")
	}
	if len(pass) < MinPasswordLength {
		return errors.New("password is too short")
	}
	if len(pass) > MaxPasswordLength {
		return errors.New("password length exceeds the limit")
	}
	// 验证密码是否正确
	hashedPass, err := repository.GetPassByName(name)
	if err != nil {
		return errors.New("user name doesn't exist")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass)); err != nil {
		return errors.New("incorrect password")
	}
	return nil
}
