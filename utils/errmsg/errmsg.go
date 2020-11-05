package errmsg

const (
	SUCCESS = 200
	FAILED  = 500

	// 注册相关
	PASSWORD_INCONSISTENCY    = 1001
	USER_IS_EXISTED           = 1002
	NAME_IS_EXISTED           = 1003
	PASSWORD_ENCRYPTION_ERROR = 1004
	USER_REGISTER_FAILED      = 1005

	// 登录相关
	PARAM_INCORRECT = 2001

	// 用户相关
	USER_NON_LOGIN = 3001

	// 视频相关
	VIDEO_SAVE_FAILED   = 4001
	VIDEO_NON_EXISTED   = 4002
	LIST_DISPLAY_FALIED = 4003
	VIDEO_DELETE_FAILED = 4004
)

var codeMsg = map[int]string{
	SUCCESS: "SUCCESS",
	FAILED:  "FAILED",

	PASSWORD_INCONSISTENCY:    "两次输入密码不一致",
	USER_IS_EXISTED:           "用户已存在",
	NAME_IS_EXISTED:           "名字已被占用",
	PASSWORD_ENCRYPTION_ERROR: "密码加密失败",
	USER_REGISTER_FAILED:      "注册失败",

	PARAM_INCORRECT: "用户名或密码错误",

	USER_NON_LOGIN: "用户未登录",

	VIDEO_SAVE_FAILED:   "视频保存失败",
	VIDEO_NON_EXISTED:   "视频不存在",
	LIST_DISPLAY_FALIED: "视频列表展示失败",
	VIDEO_DELETE_FAILED: "视频删除失败",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
