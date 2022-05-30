package errcode

//go:generate codegen -type=int
//go:generate codegen -type=int -doc -output ./error_code_generated.md

//通用：基本错误
const (
	// Success - 200: OK.
	Success int = iota + 100001

	// ErrUnknown - 500: Internal server error.
	ErrUnknown

	// ErrBind - 400: Error occurred while binding the request body to the struct.
	ErrBind

	// ErrValidation - 400: Validation failed.
	ErrValidation

	// ErrTokenInvalid - 400: Token invalid.
	ErrTokenInvalid

	// ErrIdInvalid - 400: id invalid
	ErrIdInvalid

	/*SUCCESS = 200
	ERROR = 500
	INVALID_PARAMS = 400

	ERROR_AUTH_CHECK_TOKEN_FAIL = 100101   //1-2: 服务；3-4: 模块; 5-6:错误码序列号
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 100102
	ERROR_AUTH_TOKEN = 100103
	ERROR_AUTH = 100104*/
)

// 通用: 数据库类错误
const (
	// ErrDataBase - 500: Database error.
	ErrDataBase int = iota + 100101
)

// 通用: 认证授权类错误
const (
	// ErrEncrypt - 401: Error occurred while encrypting the user password.
	ErrEncrypt int = iota + 100201

	// ErrTokenExpired - 401: Token expired.
	ErrTokenExpired

	// ErrPasswordIncorrect - 401: Password was incorrect.
	ErrPasswordIncorrect

	// ErrPermissionDenied - 403: Permission denied.
	ErrPermissionDenied
)

//通用: 编解码类错误
const (
	// ErrEncodingFailed - 500: Encoding failed due to an error with the data.
	ErrEncodingFailed int = iota + 100301

	//ErrDecodingFailed - 500: Decoding failed due to an error with the data.
	ErrDecodingFailed

	// ErrInvalidJson - 500: Data is not valid JSON.
	ErrInvalidJson

	// ErrEncodingJSON - 500: JSON data could not be encoded.
	ErrEncodingJSON

	// ErrDecodingJSON - 500: JSON data could not be decoded.
	ErrDecodingJSON

	// ErrInvalidYaml - 500: Data is not valid Yaml.
	ErrInvalidYaml

	// ErrEncodingYaml - 500: Yaml data could not be encoded.
	ErrEncodingYaml

	// ErrDecodingYaml - 500: Yaml data could not be decoded.
	ErrDecodingYaml
)
