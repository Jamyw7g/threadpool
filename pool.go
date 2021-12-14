package pool

import (
	"runtime"
	"sync"
)

type Pool struct {
	wg *sync.WaitGroup
	ch chan struct{}
}

func NewPool(num int) *Pool {
	if num <= 0 {
		num = runtime.NumCPU()
	}

	return &Pool{
		wg: &sync.WaitGroup{},
		ch: make(chan struct{}, num),
	}
}

func (p *Pool) Add() {
	p.ch <- struct{}{}
	p.wg.Add(1)
}

func (p *Pool) Done() {
	<-p.ch
	p.wg.Done()
}

func (p *Pool) Wait() {
	p.wg.Wait()
}
