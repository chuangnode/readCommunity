package initialize

import (
	"fmt"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"readCommunity/global"
	"readCommunity/internal/pkg/validation"
)

var v *validator.Validate

//InitTrans validator
func InitTrans(locale string) (err error) {
	v = validator.New()
	zhT := zh.New()
	enT := en.New()
	uni := ut.New(enT, zhT, enT)
	var ok bool
	global.Trans, ok = uni.GetTranslator(locale)
	if !ok {
		return fmt.Errorf("uni.GetTranslator(%s)", locale)
	}
	switch locale {
	case "en":
		_ = en_translations.RegisterDefaultTranslations(v, global.Trans)
	case "zh":
		err = zh_translations.RegisterDefaultTranslations(v, global.Trans)
	default:
		_ = en_translations.RegisterDefaultTranslations(v, global.Trans)
	}
	RegisterValidatorFunc(v, "validatepwd", "密码为6-30位字符,包括大写字母、小写字母、数字、特殊字符至少2种", validation.ValidatePwd)
	RegisterValidatorFunc(v, "validatePhone", "手机号需为11位有效联系方式", validation.ValidatePhone)
	return
}

type Func func(f1 validator.FieldLevel) bool

func RegisterValidatorFunc(v *validator.Validate, tag string, msgStr string, fn Func) {
	_ = v.RegisterValidation(tag, validator.Func(fn))
	_ = v.RegisterTranslation(tag, global.Trans, func(ut ut.Translator) error {
		return ut.Add(tag, "{0}"+msgStr, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(tag, fe.Field())
		return t
	})
}

func ValidateStruct(data interface{}) error {
	return v.Struct(data)
}
