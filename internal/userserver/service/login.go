package service

import (
	"fmt"
	"readCommunity/internal/pkg/utils/cryptutil"
	"readCommunity/internal/userserver/model"
)

type LoginParams struct {
	UserName string `form:"user_name" json:"user_name" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}

func LoginService(params LoginParams) (bool, error) {
	var user model.User
	pwd := cryptutil.EncryptUtil(params.Password, params.UserName)
	fmt.Printf("password: %v\n", pwd)
	istrue, err := user.CheckPwd(params.UserName, pwd)
	if err != nil {
		return false, err
	}
	if istrue {
		return true, nil
	}
	return false, nil
}
