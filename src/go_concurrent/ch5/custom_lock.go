package main

import (
	"errors"
)

type Lock interface {
	Lock() error

	UnLock() error
}

type unReentrantLock struct {
	exclusive bool
	ch        chan struct{}
}

func NewUnReentrantLock() (Lock, error) {
	ch := make(chan struct{}, 1)
	lk := unReentrantLock{
		exclusive: false,
		ch:        ch,
	}
	return &lk, nil
}

func (ur *unReentrantLock) Lock() error {
	if ur.ch == nil {
		return errors.New("channel is nil or empty")
	}
	select {
	case ur.ch <- struct{}{}:
		return nil
	}
}

func TryCloseableChannel(ch <-chan struct{}) (interface{}, bool) {
	if p, ok := <-ch; !ok {
		return nil, true
	} else {
		return p, false
	}
}

func (ur *unReentrantLock) UnLock() error {

	if ur.ch == nil {
		return errors.New("channel is nil or empty")
	}
	select {
	case _, ok := <-ur.ch:
		if !ok{
			return errors.New("close channel")
		}
		return nil
	}
}
