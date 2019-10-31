package main

import (
	"os"
	"sync"
	"errors"
	"io"
)

const (
	FIXED_LENGTH = 16
)

type DataFile interface {
	Read() (rsn int64, d Data, err error)

	Write(d Data) (wsn int64, err error)

	// 读下标
	RSN() int64

	// 写下标
	WSN() int64

	DataLen() int32

	Close() error
}

type Data []byte

type myDataFile struct {
	f       *os.File
	fmutex  sync.RWMutex
	woffset int64
	roffset int64
	// to protect woffset for a thread safe
	wmutex sync.Mutex
	// protect roffset for thread safe
	rmutex  sync.Mutex
	datalen uint32
}

func NewDataFile(path string, dataLen uint32) (DataFile, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	if dataLen == 0 {
		return nil, errors.New("invalid data length!")
	}

	df := &myDataFile{
		f:       f,
		datalen: dataLen,
	}
	return df, nil
}

func (md *myDataFile) Read() (rsn int64, d Data, err error) {
	var offset int64
	md.rmutex.Lock()
	offset = md.roffset
	md.roffset += int64(md.datalen)
	md.rmutex.Unlock()

	rsn = offset / int64(md.datalen)
	bytes := make([]byte, md.datalen)
	for {
		md.fmutex.RLock()
		// 从偏移量offset处开始读取，读取bytes长度的字节
		_, err = md.f.ReadAt(bytes, offset)
		if err != nil {
			// 如果读取到的是文件末尾，那么就开始下一次的读取
			if err == io.EOF {
				md.fmutex.RUnlock()
				// next time
				continue
			}
			// 否则的话，直接返回
			md.fmutex.RUnlock()
			return
		}
		d = bytes
		md.fmutex.RUnlock()
		return
	}
}

func (md *myDataFile) Write(d Data) (wsn int64, err error) {
	var offset int64
	md.wmutex.Lock()
	offset = md.woffset
	md.woffset += int64(md.datalen)
	md.wmutex.Unlock()

	wsn = offset / int64(md.datalen)
	var bytes []byte
	if len(d) > int(md.datalen) {
		bytes = d[0:md.datalen]
	} else {
		bytes = d
	}
	md.fmutex.Lock()
	defer md.fmutex.Unlock()
	_,err = md.f.Write(bytes)

	return
}


func (md *myDataFile) RSN() int64{
	md.rmutex.Lock()
	defer md.rmutex.Unlock()
	return md.roffset / int64(md.datalen)
}

func (md *myDataFile) WSN() int64{
	md.wmutex.Lock()
	defer md.wmutex.Unlock()
	return md.woffset / int64(md.datalen)
}

func (md *myDataFile) DataLen() int32{
	md.wmutex.Lock()
	defer  md.wmutex.Unlock()
	return int32(md.woffset)
}

func (md *myDataFile) Close() error{
	return md.f.Close()
}