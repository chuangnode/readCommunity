package validation

import (
	"fmt"
	"github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"readCommunity/global"
	"regexp"
	"strings"
)

// 初始化中文翻译
/*func init() {
	validate = validator.New()
	_ = validate.RegisterValidation("validatepwd", ValidatePwd)
	_ = validate.RegisterValidation("validatePhone", validatePhone)
}
*/
func HandleValidatorError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": RemoveTopStruct(errs.Translate(global.Trans)),
	})
	return
}

func RemoveTopStruct(fields map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fields {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

// ValidatePwd: 验证密码
func ValidatePwd(fl validator.FieldLevel) bool {
	// 匹配规则
	// 1.长度>=6, <=30
	// 2.特殊字符、大写字母、小写字母、数字至少存在2种
	// 正向否定查找，四种字符不符合的只有单一存在的
	regPwdStr := "^(?![0-9]+$)(?![a-z]+$)(?![A-Z]+$)(?![!@#$%,.+-?]+$)[0-9a-zA-Z!@#$%,.+-?]{6,30}$"
	regPwd := regexp2.MustCompile(regPwdStr, 0)
	matchString, err := regPwd.MatchString(fl.Field().String())
	if err != nil {
		fmt.Printf("validate password failed, err: %s", err)
	}
	return matchString
}

// validatePhone: 验证手机号
func ValidatePhone(fl validator.FieldLevel) bool {
	// 匹配规则 1.第一位为1； 2.第2位[3-9]; 3.[1-9]{9}
	regPhone := "^1[3-9]{1}\\d{9}$"
	reg := regexp.MustCompile(regPhone)
	return reg.MatchString(fl.Field().String())
}
