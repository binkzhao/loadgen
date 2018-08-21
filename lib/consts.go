package lib

// 声明代表载荷发生器状态的常量
const (
	// 代表原始
	STATUS_ORIGINAL uint32 = 0
	// 代表正在启动
	STATUS_STARING uint32 = 1
	// 代表已启动
	STATUS_STARTED uint32 = 2
	// 代表正在停止
	STATUS_STOPPING uint32 = 3
	// 代表已停止
	STATUS_STOPPED = 4
)

// 保留 1 ~ 1000 给载荷承受方使用。
const (
	RET_CODE_SUCCESS              RetCode = 0    // 成功。
	RET_CODE_WARNING_CALL_TIMEOUT         = 1001 // 调用超时警告。
	RET_CODE_ERROR_CALL                   = 2001 // 调用错误。
	RET_CODE_ERROR_RESPONSE               = 2002 // 响应内容错误。
	RET_CODE_ERROR_CALEE                  = 2003 // 被调用方（被测软件）的内部错误。
	RET_CODE_FATAL_CALL                   = 3001 // 调用过程中发生了致命错误！
)
