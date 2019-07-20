package chapter7

import (
	"bytes"
	"fmt"
	"io"
)

type ByteCounter int

//封装了原本的输入流，并额外增加返回流中的字节数目
type WriteByteCounter struct {
	io.Writer
	num *int64
}

//自定义的接口类型，组合了三种接口，如果需要实现这个接口，那么需要同时实现下面的三种接口中的所有方法
type CustomerCloseableRW interface {
	io.Writer
	io.Reader
	io.Closer
}

type CloseableReaderWriter struct {
	w, r  int
	close bool
}

func (crw *CloseableReaderWriter) Write(p []byte) (int, error) {
	l := len(p)
	crw.w = l
	return l, nil
}

func (crw *CloseableReaderWriter) Read(p []byte) (int, error) {
	num := len(p)
	crw.r = num
	return num,nil
}

func (crw *CloseableReaderWriter) Close()(error){
	if crw.w > 0 && crw.r > 0{
		crw.close = true
		return nil
	}

	crw.close = false
	return fmt.Errorf("task has not done:[write:%d/read:%d]\n",crw.w,crw.r)
}

//实现Writer接口的Write方法
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) //len方法返回int，将其强转为ByteCounter类型
	return len(p), nil
}

//实现fmt的String接口方法
func (c *ByteCounter) String() string {
	var buf bytes.Buffer
	buf.WriteString("ByteCounter:")
	buf.WriteByte('{')
	fmt.Fprintf(&buf, "%d ", *c)
	buf.WriteByte('}')
	return buf.String()
}

func (wc *WriteByteCounter) CountingWriter(writer io.Writer) (io.Writer, *int64) {
	var p = int64(100)
	return writer, &p
}
