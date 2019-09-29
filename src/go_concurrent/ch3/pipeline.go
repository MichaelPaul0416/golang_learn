package main

import (
	"os/exec"
	"fmt"
	"bufio"
)

func main() {
	cmd0 := exec.Command("echo", "-n", "my first command from golang")

	// 启动命令
	if err := cmd0.Start(); err != nil {
		fmt.Printf("the command can not be start...err:%v\n", err)
		return
	}

	// 创建一个能获取此命令的输出管道
	stdout0, err := cmd0.StdoutPipe()
	if err != nil {
		fmt.Printf("can not obtain pipeline:%v\n", err)
		return
	}

	// 读入的数据存入调用方传递给他的字节切片中
	//var outputBuf bytes.Buffer
	//// 带缓冲区的，循环接收
	//for{
	//	tmpOutput := make([]byte,5)
	//	n,err := stdout0.Read(tmpOutput)
	//	if err != nil{
	//		if err == io.EOF{
	//			break
	//		}else{
	//			fmt.Printf("can not read data from pipeline:%v\n",err)
	//			return
	//		}
	//	}
	//	if n > 0{
	//		outputBuf.Write(tmpOutput)
	//	}
	//}
	//
	//fmt.Printf("%s\n",outputBuf.String())


	// 或者直接使用带缓冲区的
	outputbuf0 := bufio.NewReader(stdout0)// 返回携带一个长度为4096的缓冲区
	// 第二个bool类型表明当前行是否还未读完，如果是false，那么利用for读出剩余的数据
	output0,_,err := outputbuf0.ReadLine()
	if err != nil{
		fmt.Printf("can not read data from pipeline:%v\n",err)
		return
	}
	fmt.Printf("%s\n",string(output0))
}
