package chapter2

import "fmt"

//相当于java里面声明了两个实例对象，但是class都是同一个，即使都是使用了相同的底层类型float64,但是他们的类型也是不同的
//不同类型的不能用算是表达式进行比较和合并
type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)
func init(){
	fmt.Printf("%s\n","convert function init...")
}
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string  {
	return fmt.Sprintf("%g 摄氏度",c)
}
