package lib

import "time"

// 用于表示原生请求的结构
type RawReq struct {
	ID  uint64
	Req []byte
}

// 用于表示原生相应结构
type RawResp struct {
	ID     uint64
	Resp   []byte
	err    error
	Elapse time.Duration
}

type CallResult struct {
	ID     int64         // ID
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
