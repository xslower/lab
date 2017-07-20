package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
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
					// echo(`Full! wait`)
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
					// echo(`Empty! wait`)
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

func echo(d interface{}) {
	fmt.Println(d)
}
