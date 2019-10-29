package loadgen

import (
	"fmt"
	"go_concurrent/ch4/loadgen/lib"
	"time"
	"context"
	"math"
	"sync/atomic"
	"errors"
)

func Show() {
	fmt.Printf("show \n")
}

type myGenerator struct {
	caller      lib.Caller
	timeoutNS   time.Duration
	lps         uint32
	durationNS  time.Duration
	concurrency uint32
	tickets     lib.GoTickets
	ctx         context.Context    // 上下文,用户通过这个实现取消,暂停
	cancelFunc  context.CancelFunc //  取消回调函数
	callCount   int64
	status      uint32
	resultCh    chan *lib.CallResult
}

// new myGenerator

func NewGenerator(set ParamSet) (lib.Generator, error) {
	fmt.Printf("new a load generator...")
	if err := set.Check(); err != nil {
		return nil, err
	}

	gen := &myGenerator{
		caller:     set.Caller,
		timeoutNS:  set.TimeoutNS,
		lps:        set.LPS,
		durationNS: set.DurationNS,
		resultCh:   set.ResultCh,
		status:     lib.STATUS_ORIGINAL,
	}

	if err := gen.init(); err != nil {
		return nil, err
	}

	return gen, nil
}

// 初始化函数
func (gen *myGenerator) init() error {
	fmt.Printf("initializing the load generator...")
	// 并发量 = 总的超时时间 / 发生一个载荷所需要的总时间
	var total = int64(gen.timeoutNS)/int64(1e9/gen.lps) + 1
	if total > math.MaxInt32 {
		total = math.MaxInt32
	}

	gen.concurrency = uint32(total)
	tickets, err := lib.NewGoTickets(gen.concurrency)
	if err != nil {
		return err
	}

	gen.tickets = tickets
	fmt.Printf("initializing done. concurrency= %d\n", gen.concurrency)
	return nil
}

// start
func (gen *myGenerator) Start() bool {
	fmt.Printf("starting load generator...")
	// 可以将初始化完成的载荷器启动,或者将已经停止的载荷器启动
	if !atomic.CompareAndSwapUint32(&gen.status, lib.STATUS_ORIGINAL, lib.STATUS_STARTING) {
		if !atomic.CompareAndSwapUint32(&gen.status, lib.STATUS_STOPPED, lib.STATUS_STARTING) {
			return false
		}
	}

	var throttle <-chan time.Time
	if gen.lps > 0 {
		// 计算1s内每次的发送间隔
		interval := time.Duration(1e9 / gen.lps)
		fmt.Printf("setting throttle(%v)...", interval)
		throttle = time.Tick(interval)
	}

	gen.ctx, gen.cancelFunc = context.WithTimeout(context.Background(), gen.durationNS)

	gen.callCount = 0

	atomic.StoreUint32(&gen.status, lib.STATUS_STARTED)

	go func() {
		fmt.Printf("generator loads...")
		// load
		gen.genLoad(throttle)
		fmt.Printf("stopped.(call count:%d)\n", gen.callCount)
	}()
	// 异步直接返回
	return true
}

func (gen *myGenerator) Stop() bool {
	if !atomic.CompareAndSwapUint32(&gen.status, lib.STATUS_STARTED, lib.STATUS_STARTING) {
		return false
	}

	gen.cancelFunc()
	for {
		if atomic.LoadUint32(&gen.status) == lib.STATUS_STOPPED {
			break
		}
		time.Sleep(time.Microsecond)
	}
	return true
}

func (gen *myGenerator) Status() uint32 {
	return atomic.LoadUint32(&gen.status)
}

func (gen *myGenerator) CallCount() int64 {
	return atomic.LoadInt64(&gen.callCount)
}

func (gen *myGenerator) genLoad(throttle <-chan time.Time) {
	for {
		// 先校验一下是否有关闭请求
		select {
		case <-gen.ctx.Done():
			gen.prepareToStop(gen.ctx.Err())
			return
		default:

		}

		// gen Async call
		gen.asyncCall()

		if gen.lps > 0 {
			select {
			// 如果下面两个通道都准备就绪了,那么select随机一个,如果选到了地一个case的话
			// 那么就需要在下一次进入for循环的时候重新判断一下是否执行了停止
			// 也就是for循环中的地一个select存在的意义
			case <-throttle: // 等待下个周期来临
			case <-gen.ctx.Done():
				gen.prepareToStop(gen.ctx.Err())
				return
			}
		}
	}
}

func (gen *myGenerator) prepareToStop(ctxError error) {
	fmt.Printf("prepare to stop load generator (cause: %s)...", ctxError)
	atomic.CompareAndSwapUint32(&gen.status, lib.STATUS_STARTED, lib.STATUS_STOPPING)
	fmt.Printf("close result channel...\n")
	close(gen.resultCh)
	atomic.StoreUint32(&gen.status, lib.STATUS_STOPPED)
}

// 载荷器的一次周期到了之后,执行该方法进行发送以及结果的处理
func (gen *myGenerator) asyncCall() {
	// 从goroutine的池中获取一个
	gen.tickets.Take()
	go func() {
		defer func() {
			// 转换panic

			if p := recover(); p != nil {
				// 类型断言,将p断言为error
				err, ok := interface{}(p).(error)
				var errMsg string
				if ok {
					errMsg = fmt.Sprintf("async call panic! (error: %s)", err)
				} else {
					errMsg = fmt.Sprintf("async call panic! (cause: %v)", err)
				}
				fmt.Printf("error: %s\n", errMsg)
				result := &lib.CallResult{
					ID:   -1,
					Code: lib.RET_CODE_FATAL_CALL,
					Msg:  errMsg,
				}
				gen.sendResult(result)
			}
			// 将goroutine资源归还给池
			gen.tickets.Return()
		}()
		// 回调调用方实现的构建请求
		rawReq := gen.caller.BuildReq()
		// 调用状态: 0-未调用或者调用中;1-调用完成,2-调用超时
		var callStatus uint32

		// 设置超时后的回调
		timer := time.AfterFunc(gen.timeoutNS, func() {
			// 可能cas失败或者说,旧值不是0,而是1或者2
			if !atomic.CompareAndSwapUint32(&callStatus, 0, 2) {
				return
			}
			result := &lib.CallResult{
				ID:     rawReq.ID,
				Req:    rawReq,
				Code:   lib.RET_CODE_WAITING_CALL_TIMEOUT,
				Msg:    fmt.Sprintf("timeout! (expected: < %v)", gen.timeoutNS),
				Elapse: gen.timeoutNS,
			}
			gen.sendResult(result)
		})

		rawResp := gen.callOne(&rawReq)
		if !atomic.CompareAndSwapUint32(&callStatus, 0, 1) {
			return
		}
		// 停止,重置
		timer.Stop()
		var result *lib.CallResult
		if rawResp.Err != nil {
			result = &lib.CallResult{
				ID:     rawResp.ID,
				Req:    rawReq,
				Code:   lib.RET_CODE_ERROR_CALL,
				Msg:    rawResp.Err.Error(),
				Elapse: rawResp.Elapse,
			}
		} else {
			// callback
			result = gen.caller.CheckResp(rawReq, *rawResp)
			result.Elapse = rawResp.Elapse
		}

		gen.sendResult(result)
	}()

}

func (gen *myGenerator) sendResult(result *lib.CallResult) bool {
	if atomic.LoadUint32(&gen.status) != lib.STATUS_STARTED {
		// 打印被忽略的结果
		gen.printIgnoreResult(result, "stopped load generator")
		return false
	}

	select {
	case gen.resultCh <- result:
		return true
	default:
		gen.printIgnoreResult(result, "full result channel")
		return false
	}
}

func (gen *myGenerator) printIgnoreResult(result *lib.CallResult, cause string) {
	resultMsg := fmt.Sprintf(
		"ID=%d,Code=%d,Msg=%s,Elapse=%v", result.ID, result.Code, result.Msg, result.Elapse)

	fmt.Printf("ignored result:%s,(cause:%s)\n", resultMsg, cause)
}

func (gen *myGenerator) callOne(req *lib.RawReq) *lib.RawResp {
	atomic.AddInt64(&gen.callCount, 1)
	if req == nil {
		return &lib.RawResp{
			ID:  -1,
			Err: errors.New("Invalid raw request."),
		}
	}

	start := time.Now().UnixNano()
	// 回调接口实现类的call方法,交给接口实现类自己去实现
	resp, err := gen.caller.Call(req.Req, gen.timeoutNS)
	end := time.Now().UnixNano()
	elapseTime := time.Duration(end - start)
	var rawResp lib.RawResp
	if err != nil {
		errMsg := fmt.Sprintf("sysc call error:%s.", err)
		rawResp = lib.RawResp{
			ID:     req.ID,
			Err:    errors.New(errMsg),
			Elapse: elapseTime,
		}
	} else {
		rawResp = lib.RawResp{
			ID:     req.ID,
			Resp:   resp,
			Elapse: elapseTime,
		}
	}

	return &rawResp
}
