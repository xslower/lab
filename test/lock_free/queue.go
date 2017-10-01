package main

import (
	"runtime"
	"sync/atomic"
	"sync"
)

/**
 * pi与ci之间空一个位置，用以标示空或满，这样不用额外的容量竞争变量
 */
type Queue struct {
	mod int64    //这里mod只能=2^n，取余即可使用“&”操作
	pi  int64    //生产者指针
	ci  int64    //消费者指针
	dt  []string //数据堆
	bit []bool   //标示可读还是可写的状态
}

func NewQueue() (r *Queue) {
	var ln int64 = 16
	return &Queue{
		dt:  make([]string, ln),
		bit: make([]bool, ln),
		mod: ln - 1}
}

func (this *Queue) Set(v string) (ret bool) {
	var pi, ci, mod int64
	mod = this.mod
	for {
		pi = atomic.LoadInt64(&this.pi)
		ci = atomic.LoadInt64(&this.ci)
		if pi-ci >= mod { //full
			return false
		}
		var npi = pi + 1
		if atomic.CompareAndSwapInt64(&this.pi, pi, npi) {
			break
		} else {
			// echo(`Set swap failed`)
			runtime.Gosched()
		}

	}
	var d = &this.dt[pi&mod]
	var b = &this.bit[pi&mod]
	var cnt = 0
	for {
		if *b == false {
			*b = true
			*d = v
			return true
		}
		// echo(`Set bit blocked`)
		cnt++
		if cnt > 10 {
			return false
		}
		// time.Sleep(time.Microsecond)
		runtime.Gosched()
	}

}

func (this *Queue) Get() (v string) {
	var pi, ci, mod int64
	mod = this.mod
	for {
		pi = atomic.LoadInt64(&this.pi)
		ci = atomic.LoadInt64(&this.ci)
		if pi == ci { //empty
			return
		}
		var nci = ci + 1
		if atomic.CompareAndSwapInt64(&this.ci, ci, nci) {
			break
		} else {
			// echo(`~Get swap failed~`)
			runtime.Gosched()
		}
	}
	var d = &this.dt[ci&mod]
	var b = &this.bit[ci&mod]
	var cnt = 0
	for {
		if *b {
			*b = false
			v = *d
			return
		}
		// echo(`~Get bit blocked!`)
		cnt++
		if cnt > 10 {
			return
		}
		// time.Sleep(time.Microsecond*10)
		runtime.Gosched()

	}
}

type MtxQueue struct {
	mod int64 //这里mod只能=2^n，取余即可使用“&”操作
	pi  int64
	ci  int64
	dt  []string
	bit []bool
	// mtx sync.Mutex
	mtx sync.RWMutex
}

func NewMtxQueue() (r *MtxQueue) {
	var ln int64 = 16
	return &MtxQueue{
		dt:  make([]string, ln),
		bit: make([]bool, ln),
		mod: ln - 1}
}

func (this *MtxQueue) Set(v string) (ret bool) {
	defer this.mtx.Unlock()
	this.mtx.Lock()
	var pi = this.pi
	var ci = this.ci
	var mod = this.mod
	if pi-ci >= mod { //满了
		ret = false
	} else {
		this.pi = pi + 1
		this.dt[pi&mod] = v
		ret = true
	}
	return

}

func (this *MtxQueue) Get() (v string) {
	defer this.mtx.RUnlock()
	this.mtx.RLock()
	var pi, ci, mod int64
	mod = this.mod
	pi = this.pi
	ci = this.ci
	if pi == ci { //empty
		return
	}
	this.ci = ci + 1
	v = this.dt[ci&mod]
	return
}

type QueueIfc interface {
	Set(v string) bool
	Get() string
}

func NewChQueue() (q *ChQueue) {
	q = &ChQueue{make(chan string, 16)}
	return
}

type ChQueue struct {
	ch chan string
}

func (this *ChQueue) Set(v string) bool {
	this.ch <- v
	return true
}
func (this *ChQueue) Get() (v string) {
	v = <- this.ch
	return
}
