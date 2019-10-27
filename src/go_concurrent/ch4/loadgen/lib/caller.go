package lib

import "time"

// 暴露给外部用的API接口
type Caller interface {
	// 调用方自己构建请求
	BuildReq() RawReq

	Call(req []byte,timeoutNS time.Duration)([]byte,error)

	CheckResp(req RawReq,resq RawResp) *CallResult
}
