package main

import "testing"

/**
基准测试,所有的测试函数都要以BenchMark开头,并且入参的类型是*testing.B
go test -bench 正则表达式
正则表达式表示要进行基准测试的文件,如果是目录下所有文件,则用.
 */
func BenchmarkDecode(b *testing.B){

	//测试ch7_rest_service中的decodeXml
	for i:=0;i<b.N;i++{
		DecoderXml()
	}
}
