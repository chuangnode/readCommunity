package errcode

var MsgFlags = map[int] string {
	Success: "请求成功",
	ErrUnknown: "未知错误",
	ErrBind: "请求参数绑定错误",
	ErrValidation: "请求参数验证失败",
	ErrTokenInvalid: "token不可用",
	ErrIdInvalid: "Id不可用",

	ErrDataBase: "数据库错误",

	ErrEncrypt: "数据加密失败",
	ErrTokenExpired: "token已过期",
	ErrPasswordIncorrect: "密码不正确",
	ErrPermissionDenied: "无权限访问",

	ErrEncodingFailed: "编码失败",
	ErrDecodingFailed: "解码失败",
	ErrInvalidJson: "无效的Json格式",
	ErrEncodingJSON: "JSON编码失败",
	ErrDecodingJSON: "JSON解码失败",
	ErrInvalidYaml: "无效的yaml格式",
	ErrEncodingYaml: "yaml加码失败",
	ErrDecodingYaml: "yaml解析失败",

	ErrUserNameNotNull: "用户名不能为空",
	ErrUserNameNotExist: "用户名不存在",
	ErrUserNameLenNotMatch: "用户名长度3-30",
	ErrUserAlreadyExist: "用户已存在",
	ErrPasswordNotNull: "密码不能为空",
	ErrPasswordLenNotMatch: "密码必须为6到30个字符",
	ErrPasswordInvalid: "密码必须为大写字母、小写字母、特殊字符、数字其中至少2种",
	ErrPhoneInvalid: "手机号无效",
	ErrEmailInvalid: "邮箱格式不正确",
	ErrUserRegister: "用户注册失败",
	ErrUserLogin: "用户登陆失败",

	ErrBookGetFailed: "获取书籍列表失败",
	/*SUCCESS: "ok",
	ERROR: "fail",
	INVALID_PARAMS: "请求参数错误",
	ERROR_AUTH_CHECK_TOKEN_FAIL: "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已失效",
	ERROR_AUTH_TOKEN: "Token生成失败",
	ERROR_AUTH: "Token错误",*/
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ErrUnknown]
}
