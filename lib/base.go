package lib

import "time"

// RetCode 表示结果代码的类型。
type RetCode int

// 用于表示原生请求的结构
type RawReq struct {
	ID  uint64
	Req []byte
}

// 用于表示原生相应结构
type RawResp struct {
	ID     uint64
	Resp   []byte
	Err    error
	Elapse time.Duration
}

// GetRetCodePlain 会依据结果代码返回相应的文字解释。
func GetRetCodePlain(code RetCode) string {
	var codePlain string
	switch code {
	case RET_CODE_SUCCESS:
		codePlain = "Success"
	case RET_CODE_WARNING_CALL_TIMEOUT:
		codePlain = "Call Timeout Warning"
	case RET_CODE_ERROR_CALL:
		codePlain = "Call Error"
	case RET_CODE_ERROR_RESPONSE:
		codePlain = "Response Error"
	case RET_CODE_ERROR_CALEE:
		codePlain = "Callee Error"
	case RET_CODE_FATAL_CALL:
		codePlain = "Call Fatal Error"
	default:
		codePlain = "Unknown result code"
	}
	return codePlain
}

type CallResult struct {
	ID     uint64         // ID
	Req    RawReq        // 原生请求
	Resp   RawResp       // 原生相应
	Code   RetCode       // 相应代码
	Msg    string        // 结果成因的简述
	Elapse time.Duration // 耗时
}

// 用于表示载荷发生器的接口
type Generator interface {
	// 启动载荷发生器
	Start() bool

	// 停止载荷发生器
	Stop() bool

	// 获取状态
	Status() uint32

	//获取调用计数，	每次启动会充值该计数
	CallCount() int64
}
