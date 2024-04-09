package errmsg

const (
	SUCCESS = 200 //访问成功

	//User错误代码
	INEXISTENCE_USER  = 1001 //未找到用户
	EXIST_USER        = 1002 // 用户重复
	ERROR_PASSWORD    = 1003 // 密码错误
	TOKEN_FAIL        = 1004 // token错误
	INEXISTENCE_TOKEN = 1005 // token不存在
	TOKEN_TYPE_ERROR  = 1006 // token格式错误
	TOKEN_PARSE_ERROR = 1007 // token解析错误
	INSUFFICIENT_ROLE = 1008 // 权限不足

	// Task错误代码
	TASK_NOT_EXIST = 2001 // 任务不存在

	// TaskComment错误代码
	TASK_COMMENT_NOT_EXIST = 3001 // 任务评论不存在

	// article 业务代码
	EXIST_ARTICLE       = 6001 // EXIST_ARTICLE 已存在
	INEXISTENCE_ARTICLE = 6002 // INEXISTENCE_ARTICLE 不存在

	// category 业务代码
	EXIST_CATEGORY       = 7001 // EXIST_CATEGORY 已存在分类
	INEXISTENCE_CATEGORY = 7002 // INEXISTENCE_CATEGORY 不存在分类

	// letter 业务代码
	LETTER_EXIST     = 8001 // 信件已存在
	LETTER_NOT_EXIST = 8002 // 信件不存在
	LETTER_USER_SAME = 8003 // 信件用户相同

	// Message 业务代码
	MESSAGE_EXIST = 9001 // 留言已存在

	// DATABASE_WRITE_FAIL 操作数据库错误
	DATABASE_WRITE_FAIL = 4396

	// TRANSPORT_ERR 网络传输错误
	TRANSPORT_ERR = 4004

	// 七牛云上传错误
	QN_UPLOAD_ERROR = 5001
)

var codemsg = map[uint]string{
	SUCCESS: "OK", //访问成功 200

	//User
	INEXISTENCE_USER:  "未找到用户",     // 1001
	EXIST_USER:        "用户重复",      // 1002
	ERROR_PASSWORD:    "密码错误",      // 1003
	TOKEN_FAIL:        "token错误",   // 1004
	INEXISTENCE_TOKEN: "token不存在",  // 1005
	TOKEN_TYPE_ERROR:  "token格式错误", // 1006
	TOKEN_PARSE_ERROR: "token解析错误", // 1007
	INSUFFICIENT_ROLE: "权限不足",      // 1008

	// Task
	TASK_NOT_EXIST: "任务不存在", // 2001

	// TaskComment
	TASK_COMMENT_NOT_EXIST: "任务评论不存在", // 3001

	// Article
	EXIST_ARTICLE:       "文章已存在", // 6001
	INEXISTENCE_ARTICLE: "文章不存在", // 6002

	// Category
	EXIST_CATEGORY:       "分类已存在", // 7001
	INEXISTENCE_CATEGORY: "分类不存在", // 7002

	// Letter
	LETTER_EXIST:     "信件已存在",  // 8001
	LETTER_NOT_EXIST: "信件不存在",  // 8002
	LETTER_USER_SAME: "信件用户相同", // 8003
	// Message
	MESSAGE_EXIST: "留言已存在", // 9001

	DATABASE_WRITE_FAIL: "数据库操作错误", // 4396

	TRANSPORT_ERR: "网络传输错误", // 4004

	QN_UPLOAD_ERROR: "七牛云上传错误", // 5001
}

func GetErrMsg(code uint) string {
	return codemsg[code]
}
