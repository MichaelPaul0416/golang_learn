package chapter9

import (
	"sync"
	"fmt"
)

/**
并发非阻塞的缓存实现
 */
type Memo struct {
	f     Func
	cache map[string]result

	//测试证明,使用两把锁,进行锁粒度的细化,在性能上要高于使用一把锁锁住整个Get方法

	//下列方法Get不是并发安全的,存在竞态的问题,一种解决办法就是加上锁
	lock sync.Mutex //加上互斥锁

	//上述的lock对象锁住的是整个get方法,下面可以使用两个不同的锁,分别锁住get中的两次过程[query,update]
	//但是呢,仅仅使用两把锁,锁住两个不同的过程,又会存在并发的问题,假设对于同一个key,3个goroutine在竞争,然后A抢到了query锁,在判断为空之后,释放query锁,然后去获取了update锁
	//在获取到update锁的之后,同时B抢到了query锁,由于此时A还没有更新或者还没更新结束,B判断的也是为空,所以B同样的也要去竞争update锁,这样的话就存在了竞态的问题
	queryLock  sync.Mutex
	updateLock sync.Mutex
}

type result struct {
	value interface{}
	err   error
}

type Func func(key string) (interface{}, error)

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

/**
原先的get过程被拆分为两个步骤
1.用于查询
2.在查询没有返回结果的时候进行更新
使用两把锁,做到锁粒度的细化
 */
func (memo *Memo) GetByTwoStep(key string) (interface{}, error) {
	fmt.Printf("key:%s\n", key)
	memo.queryLock.Lock()
	res, ok := memo.cache[key]
	memo.queryLock.Unlock()

	if !ok {
		res.value, res.err = memo.f(key)
		memo.updateLock.Lock()
		memo.cache[key] = res
		memo.updateLock.Unlock()
	}
	return res.value, res.err
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.lock.Lock()
	defer memo.lock.Unlock()
	//这里不是并发安全的,如果多个同时get的话,这里可能if都成立
	res, ok := memo.cache[key]
	if !ok {
		//如果没有缓存对应的key的值的话,那么就去调用缓存的函数,让其执行然后缓存结果
		res.value, res.err = memo.f(key)
		//fmt.Printf("call http url:%s\n",key)
		memo.cache[key] = res
	}
	return res.value, res.err
}

//下面采用通道的机制,实现并发的非阻塞缓存
type entry struct {
	res   result
	ready chan struct{}
}

type Memo2 struct {
	f     Func
	lock  sync.Mutex
	cache map[string]*entry
}

func New2(f Func) *Memo2 {
	return &Memo2{f: f, cache: make(map[string]*entry)}
}

func (memo2 *Memo2) GetFromChannel(key string)(interface{},error){
	memo2.lock.Lock()//无论是哪个goroutine进来,都先获取lock
	//这里不能用ok判断,因为cache中维护的只是一个entry的指针,而不是实际的数据,要判断指针是否为空
	//e,ok := memo2.cache[key]
	e := memo2.cache[key]
	if e == nil{
		//生成一个指针,然后赋值给cache数组,再释放锁,最后调用f函数去更新指针对应的结构体的值
		e = &entry{ready:make(chan struct{})}
		memo2.cache[key] = e
		memo2.lock.Unlock()

		e.res.value,e.res.err = memo2.f(key)
		close(e.ready)//发出通道关闭的信号,让其他在else分支中等待的goroutine继续往下执行
	}else{
		memo2.lock.Unlock()
		//第一个goroutine进入if分支,其它相同key的goroutine进入else分支,并且在if中执行close(e.ready)之前,一直阻塞住
		//换句话说,当不阻塞了,也就是执行if分支的那个goroutine执行完了memo2.f方法,将数据塞入到entry.res中了,此时可以直接获取了
		<- e.ready
	}

	//不管是走if的goroutine还是走else的goroutine,此时的数据都已经准备好了,可以直接获取
	return e.res.value,e.res.err
}

//封装一个请求消息和一个接受结果的通道
type request struct {
	key string
	//response必须是双向通道,计算的goroutine将计算结果发送出去,调用Get方法的goroutine将结果从中接受
	response chan result
}

//Memo3 --> request --> result(chan) --> [interface{},error]

type Memo3 struct {
	reqs chan request
}

func New3(f Func)*Memo3{
	memo3 := &Memo3{reqs: make(chan request)}
	go memo3.server(f)//启动基于这个memo3对象的server监听
	return memo3
}

func (memo3 *Memo3) server(f Func){
	cache := make(map[string]*entry)
	for req := range memo3.reqs{
		e := cache[req.key]
		if e == nil{
			e = &entry{ready:make(chan struct{})}
			cache[req.key] = e
			go e.call(f,req.key)//异步执行
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func,key string){
	e.res.value,e.res.err = f(key)

	//数据接受完毕,发送close事件
	close(e.ready)
}

func (e *entry) deliver(response chan<- result){
	//执行完call方法之后,原本阻塞在<- e.ready的goroutine就会继续往下走
	<- e.ready
	//将entry中的result通过之前注册的chan发送回去
	response <- e.res
}

func (memo3 *Memo3) GetFromMonitorGoroutine(key string)(interface{},error){
	response := make(chan result)
	memo3.reqs <- request{key:key,response:response}
	res := <- response
	return res.value,res.err
}

func (memo3 *Memo3) Close(){
	close(memo3.reqs)
}
