package main

import (
	"fmt"
	"time"
)

type node struct {
	next  *node
	pre   *node
	key   string
	value interface{}
}

type LruCache struct {
	capacity int
	count    int
	//head,tail都不存储数据
	head *node
	tail *node
	//注册表，判断是否有这个key
	register map[string]*node
}

const DefaultCacheSize = 16

func (cache *LruCache) Put(key string, value interface{}) {

	cache.initIfNecessary()

	//先判断是否存在这个key
	if _, ok := cache.register[key]; !ok {
		//不存在这个key

		//继续判断是否满了
		if cache.fullCache() {
			//删除最后一个节点，然后将新的节点加入到head.next
			removeLastNode(cache)

			//将新的key-value封装为一个节点，加入到head.next
			h := newNode(key, value)
			addToHead(cache, h)

		} else {
			//无论cache是否是空的，新加入的都是加在head
			h := newNode(key, value)
			addToHead(cache, h)
		}
	} else {
		//存在这个key，那么旧的节点删除，然后将新的值重新封装为一个node，加入到head.next
		cache.DeleteSpecificNode(key)

		addToHead(cache, newNode(key, value))
	}
}

func (cache *LruCache) Get(key string) interface{} {
	if _, ok := cache.register[key]; !ok {
		return nil
	}

	f := cache.register[key].value
	cache.DeleteSpecificNode(key)
	return f
}

func (cache *LruCache) DeleteSpecificNode(key string) {
	if _, ok := cache.register[key]; !ok {
		return
	}

	v := cache.register[key] //这里获取的其实是一个node对象，而不是指针
	p := v.pre               //返回的本身就是一个指针
	n := v.next
	p.next = n
	n.pre = p

	v.next = nil
	v.pre = nil
	cache.count --
}

func (cache *LruCache) EmptyCache() bool {
	return cache.head.next == cache.tail && cache.count == 0
}

func newNode(key string, v interface{}) *node {
	n := new(node) //new函数返回的其实是一个指针变量
	n.key = key
	n.value = v
	return n
}

func addToHead(cache *LruCache, h *node) {
	n := cache.head.next
	h.next = n
	n.pre = h
	cache.head.next = h
	h.pre = cache.head

	//将h写入注册表
	cache.register[h.key] = h
	cache.count ++
}

func removeLastNode(cache *LruCache) {
	pre := cache.tail.pre
	pre.pre.next = cache.tail
	cache.tail.pre = pre.pre
	pre.next = nil
	pre.pre = nil

	cache.count --
}

func (cache *LruCache) fullCache() bool {
	if cache == nil {
		panic("empty cache and please init it")
	}

	return cache.count >= cache.capacity
}

//先假设是线程安全的，不考虑那么多
func (cache *LruCache) initIfNecessary() {
	if cache.capacity > 0 {
		return
	}

	cache.capacity = DefaultCacheSize
	cache.register = make(map[string]*node)
	h := newNode("HEAD",nil)
	h.pre = nil
	t := newNode("TAIL",nil)
	t.next = nil
	h.next = t
	t.pre = h

	cache.head = h
	cache.tail = t
}

func PrintCache(cache LruCache) {
	if cache.EmptyCache() {
		fmt.Printf("empty cache\n")
		return
	}

	for s := cache.head; s != nil; s = s.next {
		fmt.Printf("node:%s\t",s.key)
	}
	fmt.Println()
}

func main() {
	lru := new(LruCache)
	lru.Put("hello", func(t time.Time) {
		fmt.Printf("this is hello:%v\n", t)
	})

	lru.Put("world", func(m int) {
		fmt.Printf("this is world:%d\n", m)
	})

	lru.Put("hello", func(t time.Time) {
		fmt.Printf("secondary show hello:%v\n",t)
	})

	PrintCache(*lru)

	h := lru.Get("hello")
	//将h断言为一个匿名函数，入参只有一个，就是time.Time
	h1 := h.(func(t time.Time))
	h1(time.Now())

	w := lru.Get("world")
	w1 := w.(func(m int))
	w1(1)

	fmt.Printf("empty cache:%t\n", lru.EmptyCache())
}
