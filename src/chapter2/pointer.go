package chapter2

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n",false,"输出内容不换行")
var sep = flag.String("s"," ","各个输入参数将使用该字符拼接")

func main(){
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(),*sep))
	if !*n{
		fmt.Println()
	}
}
