package chapter6

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	/**
	从左到右，slice代表的分别是64*0,64*1,64*2,64*3,64*n
	 */
	//位向量代表的数如下：64*index + slice[index]的二进制表示中任意位数的组合[如：slice[index]=1101001011,那么符合slice[index] & x != 0的数]
	//包含的数就是64 * index + x[x满足的条件slice[index] & x != 0]
	Words []uint64
}

type Count struct {
	num  int
	desc string
}

func (c *Count) Init(n int, d string) {
	c.num = n
	c.desc = d
}

func (c *Count) Add(){
	c.num ++
}

func (c Count) String() string {
	var buf bytes.Buffer
	buf.WriteString("Count{")
	buf.WriteString("num:")
	fmt.Fprintf(&buf,"%d\t",c.num)
	buf.WriteString("desc:" + c.desc + "}")
	return buf.String()
}

func (is *IntSet) AddItem(n uint64) {
	is.Words = append(is.Words, n)
}

//新增元素x到位向量slice中
func (is *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	//扩展数组，使得slice的len = int(x / 64) + 1
	for word >= len(is.Words) {
		is.Words = append(is.Words, 0)
	}
	is.Words[word] |= 1 << bit
}

//判断位向量slice中是否存在x
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	//s.Words[word] & (1 << bit) != 0 --> 代表的就是slice的那个位子上的数字包含了x%64的余数
	return word < len(s.Words) && s.Words[word]&(1<<bit) != 0
}

//合并Inset的位向量到当前位向量中
func (is *IntSet) UnionWith(a *IntSet) {
	for i, words := range a.Words {
		if i < len(is.Words) {
			is.Words[i] |= words
		} else {
			is.Words = append(is.Words, words)
		}
	}
}

//toString
func (is *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range is.Words {
		if word == 0 {
			continue
		}

		for j := 0; j < 64; j++ {
			//在每一位上的数字都是通过 1 << x%64 | slice[i]得到，所以只要word & (1 << uint(j)) != 0
			//就说明j存在，此时只需要把64^i加上就好
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				//当前位上包含的数存入
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}

	buf.WriteByte('}')
	return buf.String()
}
