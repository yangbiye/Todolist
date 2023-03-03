package service

import (
	"errors"

	"github.com/yby/todo_list/models"
	util "github.com/yby/todo_list/pkg/util/jwt"
)

func LoginUser(userID string, pwd string) (string, error) {
	user, res := models.LoginUser(userID)

	if res.Error != nil {
		return "", res.Error
	}

	if pwd != user.Password {
		return "", errors.New("密码不正确")
	}

	token, err := util.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}
	return token, nil
}

func RegisterUser(pwd string, email string) (uint, string, error) {
	if pwd == "" || email == "" {
		return 0, "", errors.New("请完整填写邮箱和密码")
	}
	user, result := models.RegisterUser(pwd, email)
	if result.Error != nil {
		return 0, "", result.Error
	}

	token, err := util.GenerateToken(user.ID)

	if err != nil {
		return 0, "", err
	}

	return user.ID, token, nil
}
