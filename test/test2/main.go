package main

import (
	//	`encoding/json`
	"fmt"
	"strconv"
	//	`strings`
	// "reflect"
	"net/http"
	"time"
)

var ()

type hehe int8
type HaHaHa struct {
	Aa string `imatag:"haha"`
	Bb int32
	Cc []rune
}

func (this HaHaHa) Hello(s string) {
	echo(`hello`, s)
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := a[1:3:5]
	echo(b)
	echo(cap(b))
}
func check(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
func echo(i ...interface{}) {
	fmt.Println(i...)
}
