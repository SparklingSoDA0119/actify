package rest

type Result int

type RestResult struct {
	Code    Result
	Message string
}

type RestResultJson struct {
	Code    Result      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

const (
	Success Result = iota
	Error          = 0x90000000

	// Common
	InvalidJson = 0x90000001 + iota
	DBInsertError
	DBGetError
	DBUpdateError
	DBDeleteError
	QueryParamError

	//User
	DuplecateEmail = 0x90001001 + iota
	DuplecatePhoneNum
	InvalidDateFormat
	WrongPasswordEncrypt
	WrongPasswordDecrypt
)

func NewRestJsonResp(code Result) *RestResultJson {
	return &RestResultJson{
		Code:    code,
		Message: MessageFromCode(code),
	}
}

func MessageFromCode(code Result) string {
	switch code {
	case Success:
		return "Success"
	case InvalidJson:
		return "Invalid Json"
	case DBInsertError:
		return "Database insert error."
	case DBGetError:
		return "Database get error."
	case DBUpdateError:
		return "Database update error."
	case DBDeleteError:
		return "Database delete error."
	case DuplecateEmail:
		return "Not allow Email"
	case DuplecatePhoneNum:
		return "Not allow Phone number"
	case InvalidDateFormat:
		return "Invalid birthdate format, expected YYYY-MM-DD"
	case WrongPasswordEncrypt:
		return "Wrong password(Encrypt)"
	case WrongPasswordDecrypt:
		return "Wrong password(Decrypt)"
	default:
		return "Unknown"
	}
}