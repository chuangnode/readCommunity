package validation

import (
	"fmt"
	"github.com/dlclark/regexp2"
	"github.com/go-playground/validator/v10"
	"regexp"
)

var validate *validator.Validate

// 初始化中文翻译
func init() {
	validate = validator.New()
	_ = validate.RegisterValidation("validatepwd", ValidatePwd)
	_ = validate.RegisterValidation("validatePhone", validatePhone)
}

// ValidateStruct: 验证结构体
func ValidateStruct(dataStruct interface{}) error {
	//zh_ch := zh.New()
	//uni := ut.New(zh_ch)
	//trans, _ := uni.GetTranslator("zh")
	err := validate.Struct(dataStruct)
	return err
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
func validatePhone(fl validator.FieldLevel) bool {
	// 匹配规则 1.第一位为1； 2.第2位[3-9]; 3.[1-9]{9}
	regPhone := "^1[3-9]{1}\\d{9}$"
	reg := regexp.MustCompile(regPhone)
	return reg.MatchString(fl.Field().String())
}
