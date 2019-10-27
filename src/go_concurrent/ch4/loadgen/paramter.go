package loadgen

import (
	"go_concurrent/ch4/loadgen/lib"
	"time"
	"bytes"
	"strings"
	"fmt"
	"errors"
)

// 载荷器的暴露入参
type ParamSet struct {
	Caller     lib.Caller
	TimeoutNS  time.Duration
	LPS        uint32               // 每秒载荷量
	DurationNS time.Duration        // 载荷器持续测试时间
	ResultCh   chan *lib.CallResult // 传输server返回的调用结果的通道
}

func (params *ParamSet) Check() error {
	var errorMsgs []string

	if params.Caller == nil{
		errorMsgs = append(errorMsgs,"empty caller")
	}

	if params.TimeoutNS == 0{
		errorMsgs = append(errorMsgs,"invalid timeout")
	}

	if params.LPS == 0{
		errorMsgs = append(errorMsgs,"invalid lps(load per second)")
	}

	if params.DurationNS == 0{
		errorMsgs = append(errorMsgs,"invalid response timeout")
	}

	if params.ResultCh == nil{
		errorMsgs = append(errorMsgs,"nil channel for receive results")
	}

	var buf bytes.Buffer
	buf.WriteString("errors by checking paramSet:")
	if errorMsgs != nil{
		errMsg := strings.Join(errorMsgs,",")
		buf.WriteString(fmt.Sprintf("not passed (%s)",errMsg))
		fmt.Printf("check param: %s\n",buf.String())
		return errors.New(errMsg)
	}

	buf.WriteString(
		fmt.Sprintf("check passed. (timeoutNs=%s,lps=%d,durationNs=%s)",
			params.TimeoutNS,params.LPS,params.DurationNS))

	fmt.Printf("check param: %s\n",buf.String())
	return nil
}
