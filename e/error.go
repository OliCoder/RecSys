package e

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_AUTH                     = 10001
	ERROR_AUTH_TOKEN               = 10002
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 10003
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 10004

	ERROR_NOT_EXIST_MOVIRE = 20001

	PAGE_NOT_FOUND = 40004

	ERROR_RPC_CALL = 50001
)

var ErrorCodeMsg = map[int]string{
	SUCCESS:                        "OK",
	ERROR:                          "Fail",
	INVALID_PARAMS:                 "Invalid Params",
	ERROR_NOT_EXIST_MOVIRE:         "Movie No Exist",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token Check Failed",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token Check Timeout",
	ERROR_AUTH_TOKEN:               "Generate Token Failed",
	ERROR_AUTH:                     "Token Error",
	PAGE_NOT_FOUND:                 "Page Not Found",
	ERROR_RPC_CALL:                 "Rpc Call Failed",
}

func GetErrorCodeMsg(code int) string {
	msg, ok := ErrorCodeMsg[code]
	if ok {
		return msg
	}
	return ErrorCodeMsg[ERROR]
}
