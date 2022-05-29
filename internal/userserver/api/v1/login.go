package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"readCommunity/initialize"
	"readCommunity/internal/pkg/app"
	"readCommunity/internal/pkg/errcode"
	"readCommunity/internal/userserver/service"
)

func Login(c *gin.Context) {
	var loginParams service.LoginParams
	if err := c.ShouldBind(&loginParams); err != nil {
		fmt.Printf("login bind params failed, err: %v\n", err)
		app.ResponseInfo(c, errcode.ErrBind, nil)
		return
	}
	if err := initialize.ValidateStruct(&loginParams); err != nil {
		fmt.Printf("validate failer,err: %v", err)
	}
	isTrue, err := service.LoginService(loginParams)
	if err != nil {
		app.ResponseInfo(c, errcode.ErrUserLogin, nil)
		return
	}
	if isTrue {
		app.ResponseInfo(c, errcode.Success, nil)
	} else {
		app.ResponseInfo(c, errcode.ErrUserLogin, nil)
	}
}
