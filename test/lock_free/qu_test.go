package main

import (
	"runtime"
	"time"
	"fmt"
	"testing"
)

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
	//co_num = 10
	chp    = make(chan bool, co_num)
	chc    = make(chan bool, co_num)
)

func producer(q QueueIfc) {
	for i := 0; i < co_num; i++ {
		go func(i int) {
			k := i % len(val)
			for j := 0; j < 10000; j++ {
				var r = q.Set(val[k])
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
func concumer(q QueueIfc) {
	for i := 0; i < co_num; i++ {
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

func estQueue(t *testing.T) {
	// maxProcs := runtime.NumCPU() // 获取cpu个数
	runtime.GOMAXPROCS(4) //限制同时运行的goroutines数量
	var start = time.Now()
	var q = NewQueue()
	producer(q)
	// time.Sleep(100000)
	concumer(q)
	for i := 0; i < co_num; i++ {
		<-chp
		<-chc
	}
	var end = time.Now()
	fmt.Println(end.Sub(start))
	fmt.Println(q)

	//mq := NewMtxQueue()
	mq := NewChQueue()
	start = time.Now()
	producer(mq)
	concumer(mq)
	for i := 0; i < co_num; i++ {
		<-chp
		<-chc
	}
	end = time.Now()
	fmt.Println(end.Sub(start))
	fmt.Println(mq)
}

func echo(d interface{}) {
	fmt.Println(d)
}
