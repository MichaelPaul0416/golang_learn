package main

import (
	"flag"
	"fmt"
	"time"
	"errors"
	"strings"
)

var (
	intflag  int
	boolflag bool
	strflag  string
	timeI timeInterval
)

func init() {
	flag.IntVar(&intflag, "i", 0, "intput int value")
	flag.IntVar(&intflag,"input",0,"input int value for long")
	flag.BoolVar(&boolflag, "b", false, "input bool value")
	flag.StringVar(&strflag, "s", "hello", "input str value")
	flag.Var(&timeI,"delta","input delta time")
}

type timeInterval []time.Duration

func (i *timeInterval) String() string{
	return fmt.Sprint(*i)
}

func (i *timeInterval) Set(value string) error{
	if len(*i) > 0{
		return errors.New("时间间隔数组已经该被设置")
	}

	str := strings.Split(value,",")
	for _,s := range str{
		d,err := time.ParseDuration(s)
		if err != nil{
			return err
		}
		*i = append(*i,d)
	}
	return nil
}

func main() {
	flag.Parse()

	fmt.Printf("intflag:%v\n", intflag)
	fmt.Printf("boolflag:%v\n", boolflag)
	fmt.Printf("strflag:%v\n", strflag)

	fmt.Printf("-----打印未被解析的参数-----\n")
	fmt.Printf("-----未被解析的参数个数:%d-----\n", flag.NArg())
	for i := 0; i < flag.NArg(); i++ {
		// flag.Arg 返回未解析的参数
		fmt.Printf("arg(%d):%v\n",i,flag.Arg(i))
	}

	fmt.Println(timeI)

}