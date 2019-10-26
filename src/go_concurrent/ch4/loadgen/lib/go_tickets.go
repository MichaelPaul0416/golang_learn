package lib

import (
	"fmt"
	"errors"
)

// goroutine的票据池.类似java的线程池

// 定义票据池的接口
type GoTickets interface {
	Take()

	Return()

	Active() bool

	Remainder() uint32

	Total() uint32
}

// 接口实现类,封装在内部
type myGoTickets struct {
	total uint32
	// 票据容器,其实就是一个带缓冲区的通道
	ticketCh chan struct{}
	active   bool
}

func NewGoTickets(total uint32) (GoTickets, error) {
	gt := myGoTickets{}
	if !gt.init(total){
		errMsg := fmt.Sprintf("The goroutine ticket pool can not be initialized and total(%d)\n",total)
		return nil,errors.New(errMsg)
	}
	return &gt,nil
}


func (gt *myGoTickets) init(total uint32)bool{
	if gt.active{
		return false
	}

	if total ==0 {
		return false
	}

	ch := make(chan struct{},total)
	n := int(total)

	for i:=0;i<n;i++{
		ch <- struct{}{}
	}

	gt.ticketCh = ch
	gt.total = total
	gt.active = true
	return true
}

func (gt *myGoTickets) Take(){
	<- gt.ticketCh
}

func (gt *myGoTickets) Return(){
	gt.ticketCh <- struct{}{}
}

func (gt * myGoTickets) Active() bool{
	return gt.active
}

func (gt *myGoTickets) Total() uint32{
	return gt.total
}

func (gt *myGoTickets) Remainder() uint32 {
	return uint32(len(gt.ticketCh))
}