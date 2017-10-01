package main

import (
	"fmt"
	"strings"
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
	b := `012345`
	var c string
	for i := 1; i < 1000; i++ {
		c += b
		_ = c
	}
}

func perf2() {
	b := `012345`
	c := make([]byte, 0, 100)
	for i := 1; i < 1000; i++ {
		c = append(c, b...)
		_ = c
	}
	var d string = string(c)
	_ = d
}

func perform() {
	cnt := 100
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
	perform()
}

func orgin() {
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
func placeHolder() {
	_ = strings.Compare(`a`, `b`)
}
