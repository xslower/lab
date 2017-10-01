package main

import (
	"testing"
	"fmt"
	"runtime"
	"time"
)
var (
    _ht     = NewHashTable(4)
	strBase = []string{
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
	co_num  = 100
	chw     = make(chan bool, co_num)
	chr     = make(chan bool, co_num)
	run_num = 1000
)

func init() {

}
func TestBase(t *testing.T) {
	v := `pofiajwef`
	_ht.Put(123, v)
	nv := _ht.Get(123)
	fmt.Println(nv)

}

func writer(ht HashTableIfc) {
	for i := 0; i < co_num; i++ {
		go func(i int) {
			k := i % len(strBase)
			for j := 0; j < run_num; j++ {
				ht.Put(123, strBase[k])
			}
			chw <- true
		}(i)
	}
}
func reader(ht HashTableIfc) {
	for i := 0; i < co_num*2; i++ {
		go func() {
			for j := 0; j < run_num; j++ {
				var v = ht.Get(123)
				_ = v
			}
			chr <- true
		}()
	}
}
func TestProf(t *testing.T) {
	runtime.GOMAXPROCS(4)
	var start = time.Now()
	writer(_ht)
	reader(_ht)
	for i := 0; i < co_num; i++ {
		<-chw
		<-chr
	}
	var end = time.Now()
	fmt.Println(end.Sub(start))
	fmt.Println(_ht)

	ht := NewMtxHT(4)
	start = time.Now()
	writer(ht)
	reader(ht)
	for i := 0; i < co_num; i++ {
		<-chw
		<-chr
	}
	end = time.Now()
	fmt.Println(end.Sub(start))
	fmt.Println(ht)
}