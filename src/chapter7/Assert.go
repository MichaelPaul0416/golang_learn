package chapter7

import (
	"io"
	"os"
	"fmt"
	"bytes"
)

func AssertTip(c bool) {
	var w io.Writer
	w = os.Stdout
	//w.Close()//此时w声明的类型为io.Writer,不具备Close方法

	f, ok := w.(*os.File) //断言w为*os.File类型-->true
	fmt.Printf("assert io.Writer --> *os.File:%t\n", ok)
	if c {
		f.Close() //返回的f是断言的类型，也就是*os.File,具有Close方法
	}
}

func ShowErrorType() {
	_, err := os.Open("/hello/world")
	fmt.Printf("%v\n", err)  //输出错误信息
	fmt.Printf("%#v\n", err) //输出原始信息
}

//使用断言来返回错误
func AssertWithError() {
	_, err := os.Open("/no/such/file")

	if pe, ok := err.(*os.PathError); ok {
		//如果类型断言成功，那么err被转换为pe的实际类型，也就是断言类型PathError
		fmt.Printf("assert err is *os.PathError:%t\n", ok)
		err = pe.Err
	}
	fmt.Printf("%s\n", err)
}

//通过接口类型断言来查询特性--> java中的instanceOf 关键字
type AbstractInvoker interface {
	Invoke(s string) (i int, err error)
}

type HttpInvoker interface {
	DoInvokerByHttp(code int, req string) (i int, err error)
}

type RpcInvoker interface {
	DoInvokerByRpc(host, port, inter string) (i int, err error)
}

type Protocol struct {
	Proto string
	Code  int
}

type Restful Protocol

type Dubbo Protocol

func (r Restful) String() string {
	var buf bytes.Buffer
	buf.WriteString("Restful:{")
	buf.WriteString("Protocol:" + r.Proto + ",")
	buf.WriteString("Code:")
	fmt.Fprintf(&buf, "%d}", r.Code)
	return buf.String()
}

func (r Restful) Invoke(s string) (i int, err error) {
	fmt.Printf("do invoke[%s]-->%s\n", s,r)
	return 0, nil
}

func (r Restful) DoInvokerByHttp(code int, req string) (i int, err error) {
	if r.Code >> 1 == 1{
		fmt.Printf("restful protocol[code:%d,req:%s] --> %s\n",code,req,r)
		return 0,nil
	}

	r.Invoke(req)
	return 0,nil
}

func (d Dubbo) String() string {
	var buf bytes.Buffer
	buf.WriteString("Dubbo:{")
	buf.WriteString("Protocol:" + d.Proto + ",")
	buf.WriteString("Code:")
	fmt.Fprintf(&buf, "%d}", d.Code)
	return buf.String()
}

func (d Dubbo) Invoke(s string) (i int,err error){
	fmt.Printf("do dubbo invoke[%s] --> %s\n",s,d)
	return 0,nil
}

func (d Dubbo) DoInvokerByRpc(host,port,inter string)(i int,err error){
	if d.Code >> 2 == 1{
		fmt.Printf("dubbo protocol[to %s:%s--%s] --> %s\n",host,port,inter,d);
		return 0,nil
	}
	d.Invoke("rpc service")
	return 0,nil
}

func DynamicTypeByAssert(invoker AbstractInvoker){
	if http,ok := invoker.(HttpInvoker);ok{
		http.DoInvokerByHttp(200,"http request message")
		return
	}

	if rpc,ok := invoker.(RpcInvoker);ok{
		rpc.DoInvokerByRpc("192.168.56.1","6379","rpc request message")
		return
	}

	invoker.Invoke("abstract invoke request")
}

func RealType(x interface{})string{
	if x == nil{
		return "NULL"
	}else if _,ok := x.(int);ok{
		return "int"
	}else if _,ok := x.(uint);ok{
		return "uint"
	}else if b,ok := x.(bool);ok{
		if b{
			return "true"
		}else {
			return "false"
		}
	}else {
		return "..."
	}
}