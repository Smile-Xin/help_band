package errmsg

const (
	SUCCESS = 200 //访问成功

	//User错误代码
	INEXISTENCE_USER = 1001 //未找到用户
	EXIST_USER       = 1002 // 用户重复

	// Task错误代码
	TASK_NOT_EXIST = 2001 // 任务不存在

	// TaskComment错误代码
	TASK_COMMENT_NOT_EXIST = 3001 // 任务评论不存在

	// DATABASE_WRITE_FAIL 操作数据库错误
	DATABASE_WRITE_FAIL = 4396
)

var codemsg = map[uint]string{
	SUCCESS: "OK", //访问成功 200

	//User
	INEXISTENCE_USER: "未找到用户", // 1001
	EXIST_USER:       "用户重复",  // 1002

	// Task
	TASK_NOT_EXIST: "任务不存在", // 2001

	// TaskComment
	TASK_COMMENT_NOT_EXIST: "任务评论不存在", // 3001

	DATABASE_WRITE_FAIL: "数据库操作错误", // 4396
}

func GetErrMsg(code uint) string {
	return codemsg[code]
}
