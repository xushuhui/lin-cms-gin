package errcode

const (
	SUCCESS = iota
	InvalidParams
	DBError
	ServerError
	AuthCheckTokenFail
	AuthCheckTokenTimeout
	ErrorAuthToken
	TimeoutAuthToken
	ErrorPassWord
	UserFound
	UserNotFound
	BookNotFound
	BookTitleRepetition
	GroupNotFound
	GroupFound
	RootGroupNotAllowDelete
	GuestGroupNotAllowDelete

	PermissionNotFound
)

var MsgFlags = map[int]string{
	SUCCESS:                  "ok",
	InvalidParams:            "请求参数错误",
	DBError:                  "数据库错误",
	ServerError:              "系统异常，请联系管理员！",
	AuthCheckTokenFail:       "Token鉴权失败",
	AuthCheckTokenTimeout:    "Token已超时",
	ErrorAuthToken:           "Token错误",
	ErrorPassWord:            "密码错误",
	UserFound:                "用户已存在",
	UserNotFound:             "用户不存在",
	BookNotFound:             "书籍不存在",
	BookTitleRepetition:      "书籍标题重复",
	GroupNotFound:            "分组不存在",
	RootGroupNotAllowDelete:  "root分组不允许删除",
	GuestGroupNotAllowDelete: "guest分组不允许删除",
	GroupFound:               "分组已存在",
	PermissionNotFound:       "权限不存在",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[InvalidParams]
}
