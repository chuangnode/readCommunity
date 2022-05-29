package service

import (
	"readCommunity/internal/pkg/utils/cryptutil"
	"readCommunity/internal/userserver/model"
)

// 注册需要参数
type RegisterRequest struct {
	UserName string `form:"username" json:"user_name" validate:"required,min=3,max=30"`
	Password string `form:"password" json:"password" validate:"required,validatepwd"`
	NickName string `form:"nickname" json:"nick_name" validate:"-"`
	Email    string `form:"email" json:"email" validate:"email"`
	Phone    string `form:"phone" json:"phone" validate:"validatePhone"`
}

func RegisterService(request RegisterRequest) error {
	pwd := cryptutil.EncryptUtil(request.Password, request.UserName)
	user := model.User{
		UserName: request.UserName,
		NickName: request.NickName,
		Password: pwd,
		Email:    request.Email,
		Phone:    request.Phone,
	}
	err := user.AddUser(user)
	return err
}
