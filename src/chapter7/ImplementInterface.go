package chapter7

import (
	"fmt"
	"bytes"
	"flag"
)

//定义一个空接口
type Empty interface{}

type Celsius struct {
	value float64
}

//摄氏度转为华氏度
func (c *Celsius) CToF(t float64) {
	f := 5*t/9 + 32
	c.value = f
}

func (c *Celsius) FToC(t float64) {
	f := (t - 32) * 9 / 5
	c.value = f
}

func (c *Celsius) Init(f float64) {
	c.value = f
}

type CelsiusFlag struct {
	Celsius
}

func (cf *CelsiusFlag) String() string{
	var buf bytes.Buffer
	buf.WriteString("Celsius:{")
	fmt.Fprintf(&buf, "%.2f", cf.value)
	buf.WriteByte('}')
	return buf.String()
}

//flag.Value方法需要一个Set和String接口方法
func (cf *CelsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "℃":
		cf.Celsius = Celsius{value: value}
		return nil
	case "F", "℉":
		cf.Celsius = Celsius{value: value}
		cf.CToF(cf.value) //将摄氏度转为华氏度
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func ChangeTemperature(name string, value Celsius, usage string) *Celsius {
	f := CelsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
