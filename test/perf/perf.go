package main

import (
	"fmt"
	"time"
	"unsafe"
)

type Bytes []byte

func (b Bytes) Append(d ...byte) {
	b = append([]byte(b), d...)
	echo(b)
}
func (b *Bytes) AppendString(str ...string) {
	for _, s := range str {
		*b = append([]byte(*b), s...)
	}

}

type Slice struct {
	arr uintptr
	ln  int
	cp  int
}

func perf() {
	b := make([]byte, 800)
	ptr := &b
	for i := 1; i < 10; i++ {
		*ptr = append(*ptr, `hahaha `...)
	}
}

func perf2() {
	b := make(Bytes, 800)
	ptr := &b
	for i := 1; i < 10; i++ {
		ptr.AppendString(`hahaha `)
	}
}

func perform() {
	cnt := 100000
	start := time.Now()
	for i := 0; i < cnt; i++ {
		perf()
	}
	end := time.Now()
	echo(end.Sub(start))

	start = time.Now()
	for i := 0; i < cnt; i++ {
		perf2()
	}
	end = time.Now()
	echo(end.Sub(start))
}

func main() {
	b := make([]byte, 0, 100)
	c := b
	c = append(c, 'a', 'b')
	echo(b, c)
	ptr := (*Slice)(unsafe.Pointer(&b))
	ptr.ln += 2
	b = *(*Bytes)(unsafe.Pointer(ptr))
	echo(b, c)
}
func echo(i ...interface{}) {
	fmt.Println(i...)
}
