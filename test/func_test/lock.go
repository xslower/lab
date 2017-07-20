package main

import (
	"fmt"
	"runtime"
	"sync"
	// "sync/atomic"
	"time"
)

/**
 * pi与ci之间空一个位置，用以标示空或满，这样不用额外的容量竞争变量
 */
type Queue struct {
	mod int64 //这里mod只能=2^n，取余即可使用“&”操作
	pi  int64
	ci  int64
	dt  []string
	bit []bool
	// mtx sync.Mutex
	mtx sync.RWMutex
}

func NewQueue() (r *Queue) {
	var ln int64 = 16
	return &Queue{
		dt:  make([]string, ln),
		bit: make([]bool, ln),
		mod: ln - 1}
}

func (this *Queue) Set(v string) (ret bool) {
	// defer this.mtx.Unlock()
	// this.mtx.Lock()
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

func (this *Queue) Get() (v string) {
	// defer this.mtx.RUnlock()
	// this.mtx.RLock()
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

var (
	val = []string{
		`aaaaaaaaaaaa`,
		`bbbbbbbbbbbb`,
		`ccccccccccccc`,
		`ddddddddddd`,
		`eeeeeeeee`,
		`ffffffffff`,
		`gggggggggg`,
		`hhhhhhhhh`,
		`iiiiiiiii`,
		`lllllllllll`,
	}
	cnt = 10
	chp = make(chan bool, cnt)
	chc = make(chan bool, cnt)
)

func producer(q *Queue) {
	for i := 0; i < cnt; i++ {
		go func(i int) {
			for j := 0; j < 10000; j++ {
				var r = q.Set(val[i])
				if !r {
					// time.Sleep(time.Microsecond)
					runtime.Gosched()
				}
			}
			chp <- true
		}(i)
	}
}
func concumer(q *Queue) {
	for i := 0; i < cnt; i++ {
		go func() {
			for j := 0; j < 10000; j++ {
				var r = q.Get()
				if r == `` {
					// time.Sleep(time.Microsecond)
					runtime.Gosched()
				}
				// fmt.Println(r)
			}
			chc <- true
		}()
	}
}

func main() {
	// maxProcs := runtime.NumCPU() // 获取cpu个数
	runtime.GOMAXPROCS(4) //限制同时运行的goroutines数量
	var start = time.Now()
	var q = NewQueue()
	producer(q)
	// time.Sleep(100000)
	concumer(q)
	for i := 0; i < cnt; i++ {
		<-chp
		<-chc
	}
	var end = time.Now()
	fmt.Println(*q)
	fmt.Println(end.Sub(start))
}
