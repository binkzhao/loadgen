package loadgen

import (
	"time"
	"loadgen/lib"
	"golang.org/x/net/context"
)

type myGenerator struct {
	caller      lib.Caller           // 调用器
	timeoutNS   time.Duration        // 相应超时时间
	lps         uint32               // 每秒载荷量
	durationNS  time.Duration        // 负载持续时间
	resultCh    chan *lib.CallResult // 调用结果通道
	concurrency uint32               // 载荷并发量
	tickets     lib.GoTickets        // goroutine 票池 ，数量由concurrency决定
	stopSign    chan struct{}        // 停止信号的传递通道
	ctx         context.Context      // 上下文
	cancelFunc  context.CancelFunc   // 取消函数
	callCount   int64                // 调用计数约等于 timeoutNS / (1e9 / lps) + 1
	status      uint32               // 状态
}

// 新建一个载荷发生器
func NewGenerator(
	caller lib.Caller,
	timeoutNS time.Duration,
	lps uint32,
	durationNS time.Duration,
	resultCh chan *lib.CallResult,
) (lib.Generator, error) {

	gen := &myGenerator{
		caller:     caller,
		timeoutNS:  timeoutNS,
		lps:        lps,
		durationNS: durationNS,
		resultCh:   resultCh,
	}

	return gen, nil
}
