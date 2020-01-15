package main

import (
	"github.com/jessevdk/go-flags"
	"fmt"
)

type option struct {
	// 设置长短两种配置
	Verbose []bool `short:"v" long:"verbose" desc:"show verbose debug message"`
}

type complex struct {
	intFlag        int            `short:"i" long:"int" desc:"int value"`
	intSlice       []int          `long:"intSlice" desc:"int slice flag value"`
	boolFlag       bool           `short:"b" long:"bool" desc:"bool value"`
	BoolSlice      []bool         `long:"boolslice" description:"bool slice flag value"`
	//FloatFlag      float64        `long:"float", description:"float64 flag value"`
	FloatSlice     []float64      `long:"floatslice" description:"float64 slice flag value"`
	StringFlag     string         `short:"s" long:"string" description:"string flag value"`
	StringSlice    []string       `long:"strslice" description:"string slice flag value"`
	PtrStringSlice []*string      `long:"pstrslice" description:"slice of pointer of string flag value"`
	Call           func(string)   `long:"call" description:"callback"`
	IntMap         map[string]int `long:"intmap" description:"A map from string to int"`
}

func main() {
	//simpleExample()
	var cp complex
	s,err := flags.Parse(&cp)
	if err != nil{
		fmt.Printf("error:%v\n",err)
		return
	}
	fmt.Println(s)

}

func simpleExample() {
	var opt option
	flags.Parse(&opt)
	fmt.Println(opt)
}
