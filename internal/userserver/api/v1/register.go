package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"readCommunity/internal/pkg/app"
	"readCommunity/internal/pkg/errcode"
	"readCommunity/internal/pkg/validation"
	"readCommunity/internal/userserver/service"
)

type ValidateError struct {
	Key   string
	Error string
}

func Register(c *gin.Context) {
	var registerParams service.RegisterRequest
	if err := c.ShouldBind(&registerParams); err != nil {
		fmt.Printf("register params bind failed, err: ", err)
		app.ResponseInfo(c, errcode.ErrBind, nil)
		return
	}

	// 验证参数
	if err := validation.ValidateStruct(registerParams); err != nil {
		//TODO 处理validator返回的错误信息
		validationErrors := err.(validator.ValidationErrors)
		var code int
		for _, validationError := range validationErrors {
			value := validationError.Value().(string)
		//field:字段名，tag: 限制属性
			switch validationError.Field() {
				case "UserName":
					if value == "" {
						code = errcode.ErrUserNameNotNull
					} else if len(value) < 3 || len(value) >30 {
						code = errcode.ErrUserNameLenNotMatch
					}
			case "Password":
				if value == "" {
					code = errcode.ErrPasswordNotNull
				} else if len(value) < 6 || len(value) > 30 {
					code = errcode.ErrPasswordLenNotMatch
				} else {
					code = errcode.ErrPasswordInvalid
				}
			case "Email":
				code = errcode.ErrEmailInvalid
			case "Phone":
				code = errcode.ErrPhoneInvalid
			}
		}
		fmt.Println("v1/register.go Register validate failed, err: ", err)
		app.ResponseInfo(c, code, nil)
		return
	}
	err := service.RegisterService(registerParams)
	if err != nil {
		fmt.Printf("v1/register.go Register Service failed, err: %s\n", err)
		app.ResponseInfo(c, errcode.ErrUserRegister, nil)
	} else {
		app.ResponseInfo(c, errcode.Success, nil)
	}
}
