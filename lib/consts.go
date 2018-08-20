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
