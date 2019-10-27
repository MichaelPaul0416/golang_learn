package lib

import "time"

// 封装了一系列的bean

// 原生的请求
type RawReq struct {
	ID  int64
	Req []byte
}

// 原生的响应结构
type RawResp struct {
	ID     int64
	Resp   []byte
	Err    error
	Elapse time.Duration
}

// 结果码
type RetCode int

const (
	RET_CODE_SUCCESS              RetCode = 0
	RET_CODE_WAITING_CALL_TIMEOUT         = 1001
	RET_CODE_ERROR_CALL                   = 2001
	RET_CODE_ERROR_RESPONSE               = 2002
	RET_CODE_ERROR_CALEE                  = 2003 // 服务器内部错误
	RET_CODE_FATAL_CALL                   = 3001 // 载荷器内部错误
)

//调用结果,返回个给调用方的结果
type CallResult struct {
	ID     int64
	Req    RawReq
	Resp   RawResp
	Code   RetCode
	Msg    string
	Elapse time.Duration
}

func GetRetCodePlain(code RetCode) string {
	var plain string
	switch code {
	case RET_CODE_SUCCESS:
		plain = "Success"
	case RET_CODE_WAITING_CALL_TIMEOUT:
		plain = "caller timeout waiting"
	case RET_CODE_ERROR_CALL:
		plain = "caller error"
	case RET_CODE_ERROR_RESPONSE:
		plain = "response to caller error"
	case RET_CODE_ERROR_CALEE:
		plain = "server inner error"
	case RET_CODE_FATAL_CALL:
		plain = "soft inner error"
	default:
		plain = "unknown error"
	}
	return plain
}

// 载荷发生器的状态量
const (
	STATUS_ORIGINAL uint32 = 0
	STATUS_STARTING uint32 = 1
	STATUS_STARTED  uint32 = 2
	STATUS_STOPPING uint32 = 3
	STATUS_STOPPED  uint32 = 4
)

type Generator interface {
	Start() bool

	Stop() bool

	Status() uint32

	CallCount() int64
}
