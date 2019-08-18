package chapter9

import (
	"fmt"
	"time"
	"sync"
	"bytes"
)

var x, y int
/**
下面这个函数可能会有多个结果
x=0,y=0
x=1,y=0
x=1,y=2
x=0,y=2
 */
func Memory_1() {
	go func() {
		x = 1
		fmt.Printf("y:%d\n", y)
	}()

	go func() {
		y = 2
		fmt.Printf("x:%d\n", y)
	}()

	time.Sleep(1 * time.Second)
}

type person struct {
	name  string
	age   int
	email string
}

func (p person) String() string {
	var buf bytes.Buffer
	buf.WriteString("person{name:")
	buf.WriteString(p.name)
	buf.WriteString(",age:")
	fmt.Fprintf(&buf, "%d,email:%s}", p.age, p.email)
	return buf.String()
}

//共享的全局变量
var students map[string]person

/**
使用读写锁确保初始化一次
 */
var rwlock sync.RWMutex

func PersonInfo(id string) {
	rwlock.RLock()
	if students != nil {
		p := students[id]
		rwlock.RUnlock()
		fmt.Printf("get person:%s\n", p)
		return
	}

	rwlock.RUnlock()
	rwlock.Lock()
	if students == nil {
		initPerson()

	}
	p := students[id]
	fmt.Printf("get person after init:%s\n", p)
	rwlock.Unlock()
}

/**
使用sync.Once保证初始化安全
 */
var loadInitOnce sync.Once
func InitAndQueryBySyncOnce(id string) {
 	loadInitOnce.Do(initPerson)
 	fmt.Printf("get person:%s\n",students[id])
}

func initPerson() {
	students = map[string]person{
		"001": loadPerson("1", "1", 1),
		"002": loadPerson("2", "2", 2),
	}
	fmt.Printf("init done\n")
}

func loadPerson(n, e string, a int) person {
	p := person{n, a, e}
	return p
}
