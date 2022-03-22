package controllers

type ResCode int

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParma
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy

	CodeNeedLogin
	CodeInvalidAuth
	CodeInvalidAuthFormat
)

var codeMsg = map[ResCode]string{
	CodeSuccess:           "success",
	CodeInvalidParma:      "请求参数错误",
	CodeUserExist:         "用户已存在",
	CodeUserNotExist:      "用户不存在",
	CodeInvalidPassword:   "密码错误",
	CodeServerBusy:        "服务繁忙",
	CodeNeedLogin:         "请先登录",
	CodeInvalidAuth:       "无效的token",
	CodeInvalidAuthFormat: "认证格式有误",
}

// Msg 根据错误码获得信息
func (c ResCode) Msg() string {
	msg, ok := codeMsg[c]
	if !ok {
		return codeMsg[CodeServerBusy]
	}
	return msg
}
