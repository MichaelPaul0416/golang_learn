package main

import (
	"testing"
	"fmt"
	"time"
)

func TestContainer(t *testing.T) {
	c, err := NewContainer(2)
	if err != nil {
		t.Log("get container error")
		return
	}
	producer := Producer{
		c,
	}
	// 3个生产者，5个消费者
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Printf("ready put number:%d\n",1)
			producer.put(1)
			fmt.Printf("put number:%d and len:%d\n", 1,producer.size())
		}()
	}

	consumer := Consumer{
		c,
	}

	time.Sleep(1 * time.Second)
	for i := 0; i < 1; i++ {
		go func() {
			fmt.Printf("ready get number\n")
			d,err := consumer.get()
			if err != nil{
				return
			}
			fmt.Printf("get number:%d\n",d)
		}()
	}

	// 直接睡眠模拟，计数器太麻烦
	time.Sleep(1000 * time.Second)
}
