package main

import (
	"fmt"
	"time"
)

var max = 9999999

func getNextStep(cur int) (ret int) {
	if cur < 100 {
		return 100
	} else if cur < 1000 {
		return cur
	}
	ret = cur * 30 / 100
	return
}

func copy_() {
	var start = time.Now().UnixNano()
	var a = make([]int8, max)
	for i := 0; i < max; i++ {
		a[i] = 2
	}
	var b = make([]int8, max)
	for i := 0; i < max; i++ {
		b[i] = 3
	}
	// var c = make([]int, max)
	var end = time.Now().UnixNano()
	echo(end - start)
	start = time.Now().UnixNano()
	copy(b, a)
	end = time.Now().UnixNano()
	echo(end - start)
	start = time.Now().UnixNano()
	for i := 0; i < max; i++ {
		b[i] = a[i]
	}
	end = time.Now().UnixNano()
	echo(end - start)
}

func speed() {
	// var re = make(map[int]int, 100000000)
	var start = time.Now().UnixNano()
	var re = make([]int, 5)
	for i := 0; i < max; i++ {
		if len(re) <= i {
			tmp := re
			re = make([]int, len(re)+getNextStep(len(re)))
			// for k, _ := range tmp {
			// 	re[k] = tmp[k]
			// }
			copy(re, tmp)
		}
		re[i] = i
	}
	var end = time.Now().UnixNano()
	echo(end - start)
	start = time.Now().UnixNano()
	var co = make([]int, 5)
	for i := 0; i < max; i++ {
		co = append(co, i)
	}
	end = time.Now().UnixNano()
	echo(end - start)
	// time.Sleep(time.Hour)
}

func _map_() {
	var m = make(map[int32]int32, max)
	for i := int32(0); i < int32(max); i++ {
		m[i] = i
	}
	j := int32(max)
	m[j] = j
	m[j+1] = j + 1
	time.Sleep(time.Hour)
}

func main() {
	// copy_()
	// speed()
	_map_()
}
func echo(i ...interface{}) {
	fmt.Println(i...)
}
