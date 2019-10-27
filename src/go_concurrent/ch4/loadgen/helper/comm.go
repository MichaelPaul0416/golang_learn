package helper

import (
	"net"
	"bytes"
	"bufio"
	"go_concurrent/ch4/loadgen/lib"
	"time"
	"math/rand"
	"encoding/json"
	"fmt"
)

const (
	DELIM = '\n'
)

var operaors = []string{"+", "-", "*", "/"}

type TCPComm struct {
	addr string
}

func NewTCPComm(addr string) lib.Caller {
	return &TCPComm{addr: addr}
}

func (comm *TCPComm) BuildReq() lib.RawReq {
	id := time.Now().UnixNano()
	sreq := ServerReq{
		ID: id,
		Operands: []int{
			int(rand.Int31n(1000) + 1),
			int(rand.Int31n(1000) + 1),
		},
		Operator: func() string {
			return operaors[rand.Int31n(100)%4]
		}(),
	}

	bytes, err := json.Marshal(sreq)
	if err != nil {
		panic(err)
	}

	rawreq := lib.RawReq{ID: id, Req: bytes}
	return rawreq
}

func (comm *TCPComm) Call(req []byte, timeoutNS time.Duration) ([]byte, error) {
	conn, err := net.DialTimeout("tcp", comm.addr, timeoutNS)
	if err != nil {
		return nil, err
	}

	_, err = write(conn, req, DELIM)
	if err != nil {
		return nil, err
	}

	return read(conn,DELIM)
}

// 组装返回结果
func (comm *TCPComm) CheckResp(req lib.RawReq, resp lib.RawResp) *lib.CallResult{
	var commResult lib.CallResult
	commResult.ID = resp.ID
	commResult.Req = req
	commResult.Resp = resp

	var sreq ServerReq
	err := json.Unmarshal(req.Req,&sreq)
	if err != nil{
		commResult.Code = lib.RET_CODE_FATAL_CALL
		commResult.Msg = fmt.Sprintf("Incorrectly formatted req:%s\n",string(req.Req))
		return &commResult
	}

	var sresp ServerResp
	err = json.Unmarshal(resp.Resp,&sresp)
	if err != nil{
		commResult.Code = lib.RET_CODE_ERROR_RESPONSE
		commResult.Msg = fmt.Sprintf("Incorrectly formatted resp:%s\n",string(resp.Resp))
		return &commResult
	}

	if sresp.ID != sreq.ID{
		commResult.Code = lib.RET_CODE_ERROR_RESPONSE
		commResult.Msg = fmt.Sprintf("Inconsistent raw id!(%d != %d)\n",sreq.ID,sresp.ID)
		return &commResult
	}

	if sresp.Err != nil{
		commResult.Code = lib.RET_CODE_ERROR_CALEE
		commResult.Msg = fmt.Sprintf("abnormal server:%s\n",sresp.Err)
		return &commResult
	}

	if sresp.Result != op(sreq.Operands,sreq.Operator){
		commResult.Code = lib.RET_CODE_ERROR_RESPONSE
		commResult.Msg = fmt.Sprintf("incorrect result:%s\n",genFomula(sreq.Operands,sreq.Operator,sresp.Result,false))
		return &commResult
	}

	commResult.Code = lib.RET_CODE_SUCCESS
	commResult.Msg = fmt.Sprintf("success. (%s)",sresp.Formula)
	return &commResult

}

func read(conn net.Conn, delim byte) ([]byte, error) {
	readBytes := make([]byte, 1)
	var buffer bytes.Buffer
	for {
		_, err := conn.Read(readBytes)
		if err != nil {
			return nil, err
		}

		readByte := readBytes[0]
		if readByte == delim {
			break
		}
		buffer.WriteByte(readByte)
	}

	return buffer.Bytes(), nil
}

func write(conn net.Conn, content []byte, delim byte) (int, error) {
	writer := bufio.NewWriter(conn)
	n, err := writer.Write(content)
	if err == nil {
		writer.WriteByte(delim)
	}

	if err == nil {
		err = writer.Flush()
	}
	return n, err
}
