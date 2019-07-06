package chapter3

import (
	"fmt"
	"unicode/utf8"
	"strings"
	"strconv"
)

func ShowNumber() {
	var u uint8 = 255
	//因为这里是8位的无符号int，所以u+1：11111111 + 1 --> 00000000
	fmt.Println(u, u+1, u*u)

	fmt.Printf("7 ^ 9 = %d\n", 7^9)

	//z=x&^y:y的某一位为1,z对应的那一位就是0,否则就取x对应的那一位的值
	fmt.Printf("7 &^ 9 = %d\n", 7&^9)

	x := 1<<1 | 1<<5
	y := 1<<1 | 1<<2
	fmt.Printf("x集合：%08b\n", x)
	fmt.Printf("y集合：%08b\n", y)
	fmt.Printf("x,y交集：%08b\n", x&y)
	fmt.Printf("x,y并集：%08b\n", x|y)
	fmt.Printf("x,y对称差：%08b\n", x^y)
	fmt.Printf("x,y 差集：%08b\n", x&^y)

	fmt.Printf("8进制数：0777=%d\n", 0777)
	num := 0666
	fmt.Printf("%d的八进制[%o]/十六进制[%x]/二进制[%b]\n", num, num, num, num)
	//与上面一行的输出效果一样,[1]表示Printf重复使用第一个操作数,#的话是输出进制的前缀
	fmt.Printf("%d的八进制[%#[1]o]/十六进制[%#[1]x]/二进制[%[1]b]\n", num)

	chinese := "你好啊"
	fmt.Printf("中文的字节数:%d\n", len(chinese))

	full := "hello world"
	fmt.Printf("full[1:4]-->%v\n", full[1:4])
	end := "world"
	fmt.Printf("compare two string:%v\n", "hello "+end == full)

	mix := "hello, 世界"
	fmt.Printf("byte number:%d\n", len(mix))
	fmt.Printf("char number:%d\n", utf8.RuneCountInString(mix))
	for i := 0; i < len(mix); {
		r, size := utf8.DecodeRuneInString(mix)
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
}

func BaseName(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func BaseNameWithLib(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	slash = strings.LastIndex(s, ".")
	s = s[:slash]
	return s
}

func FormatNumber(s string) string {
	if len(s) < 3 {
		return s
	}

	var tips = ""
	for i := len(s) - 1; i >= 0; i-- {
		if (len(s)-i)%3 == 0 {
			tips = "," + s[i:i+3] + tips
		}

		if i == 0 {
			if len(s)%3 != 0 {
				tips = s[0:len(s)%3] + tips
			}else {
				tips = tips[1:]
			}
		}
	}

	return tips
}

func ConvertToString(number int32) {
	x := number
	y := fmt.Sprintf("%d",x)
	fmt.Printf("int --> string:%v\n",y)
	fmt.Printf("10 --> 2:%v\n",strconv.FormatInt(int64(x),2))
}

func ShowConstant(){
	fmt.Printf("Monday:%d\n",Monday)
}

type WeekDay int

const (
	Sunday WeekDay = iota
	Monday
	Tuesday
)