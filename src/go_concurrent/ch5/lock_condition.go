package main

import (
	"sync"
	"errors"
)

// 使用独占锁+条件变量 实现生产者消费者模型
type Producer struct {
	*container
}

type Consumer struct {
	*container
}

func (p *Producer) produce(data int) error {
	p.put(data)
	return nil
}

func (c *Consumer) consume() (int, error) {
	return c.get()
}

type container struct {
	// container
	cap int
	c        []int
	lock     sync.Mutex
	notFull  *sync.Cond
	notEmpty *sync.Cond
}

func NewContainer(c int) (*container, error) {
	if c < 0 {
		return nil, errors.New("container cap must > 0")
	}

	con := &container{
		cap:c,
		c: make([]int,0),
	}

	con.notFull = sync.NewCond(&con.lock)
	con.notEmpty = sync.NewCond(&con.lock)

	return con, nil

}
func (c *container) put(data int) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.isFull() {
		c.notFull.Wait()
	}

	c.c = append(c.c, data)
	c.notEmpty.Broadcast()
}

func (c *container) get() (int, error) {
	// 类比java中Object#notify方法
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.isEmpty() {
		c.notEmpty.Wait()
	}

	d := c.c[0]
	if d == 0 {
		return 0, errors.New("empty data")
	}
	c.c = c.c[1:]
	c.notFull.Broadcast()
	return d, nil
}

func (c *container) isFull() bool {
	return len(c.c) == c.cap
}

func (c *container) isEmpty() bool {
	return len(c.c) == 0
}

func (c *container) size() int {
	return len(c.c)
}
