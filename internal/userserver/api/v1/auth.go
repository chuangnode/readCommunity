package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"readCommunity/internal/pkg/app"
	"readCommunity/internal/pkg/errcode"
	"readCommunity/internal/pkg/utils/cryptutil"
	"readCommunity/internal/pkg/utils/jwt"
	"readCommunity/internal/userserver/service"
)

type Auth struct {
	UserName string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}

// @Summary  生成token
// @Produce  json
// @Param    username  body      string  true  "username"
// @Param    password  body      string  true  "password"
// @Success  200       {string}  json    "{"code": 200,"data":{},"msg":"ok"}"
//@Router    /auth [post]
func GetAuth(c *gin.Context) {
	var auth Auth
	if err := c.ShouldBind(&auth); err != nil {
		fmt.Printf("c.shouldbindjson failed,err:", err)
	}
	code := errcode.Success
	data := make(map[string]interface{})
	password := cryptutil.EncryptUtil(auth.Password, auth.UserName)
	//auth := Auth{UserName: username, Password: password}
	/*if err := validation.ValidateStruct(auth); err != nil {
		code = errcode.INVALID_PARAMS
	} else {*/
	isExist, err := service.CheckAuth(auth.UserName, password)
	if err != nil {
		code = errcode.ErrPasswordIncorrect
	}
	if isExist {
		token, err := jwt.GenerateToken(auth.UserName, password)
		if err != nil {
			code = errcode.ErrEncodingFailed
		} else {
			data["token"] = token
		}
	} else {
		code = errcode.ErrUserNameNotExist
	}
	//}
	app.ResponseInfo(c, code, data)
}
