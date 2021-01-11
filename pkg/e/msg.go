package e

var MsgFlags = map[int]string{
	SUCCESS:                    "ok",
	ERROR:                      "fail",
	InvalidParams:              "请求参数错误",
	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token已超时",
	ErrorAuthToken:             "Token生成失败",
	ErrorAuth:                  "Token错误",
	ErrorAuthnameExist:         "用户名已存在",
	ErrorAccountIsDisabled:     "账号已禁用",
	ErrorAccountNotExist:       "账号不存在",
	ErrorTagExist:              "标签已存在",
	ErrorTagNotExist:           "标签不存在",
}

/**

 */
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[code]
}
