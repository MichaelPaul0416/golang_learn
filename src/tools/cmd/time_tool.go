package main

import (
	"time"
	"fmt"
	"errors"
	"strings"
	"flag"
)

type interval []time.Duration

// 绑定了下面两个方法，相当于就是实现了Value接口
func (i *interval) String() string{
	return fmt.Sprintf("%v",*i)
}

func (i *interval) Set(value string) error{
	if len(*i) > 0{
		return errors.New("interval flag already set")
	}

	for _,dt := range strings.Split(value,","){
		duration,err := time.ParseDuration(dt)
		if err != nil{
			return err
		}
		// 加到slice中
		*i = append(*i,duration)
	}
	return nil
}

var intervalFlag interval

func init(){
	flag.Var(&intervalFlag,"deltaT","comma-separated list of input")
}

func main(){
	flag.Parse()
	fmt.Println(intervalFlag)
}