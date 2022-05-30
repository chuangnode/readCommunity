package errcode

// 用户模块相关错误
const(
	// ErrUserNameNotNull - 400: user name must not null.
	ErrUserNameNotNull int = iota + 110001

	// ErrUserNameNotExist - 400: user name not exist, please register.
	ErrUserNameNotExist

	// ErrUserNameLenNotMatch -400: user name length between 3 and 30.
	ErrUserNameLenNotMatch

	// ErrUserAlreadyExist - 500: user has algready exist, can't register again.
	ErrUserAlreadyExist

	// ErrPasswordNotNull - 400: password must not null.
	ErrPasswordNotNull

	// ErrPasswordLenNotMatch - 400: password length between 6 and 30.
	ErrPasswordLenNotMatch

	// ErrPasswordInvalid - 400: the password needs to meet at least 2 types of \
	//uppercase letters, lowercase letters, special characters and numbers.
	ErrPasswordInvalid

	// ErrPhoneInvalid - 400: phone is not valid.
	ErrPhoneInvalid

	// ErrEmailInvalid - 400: email is not valid
	ErrEmailInvalid

	// ErrUserRegister - 500: user register failed
	ErrUserRegister

	// ErrUserLogin - 500: user login failed
	ErrUserLogin
)
