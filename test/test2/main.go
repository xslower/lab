package main

import (
	//	`encoding/json`

	"fmt"
	// "github.com/resure-tech/lib/encoding/json"
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
func strSlc2Map(slc []string) (m map[string]bool) {
	for i := 0; i < len(slc); i++ {
		m[slc[i]] = true
	}
	return
}

func setMap(m *map[string]bool) {
	(*m)[`aa`] = true
}

func main() {
	// str := []string{`a`, `b`, `c`, `d`}
	// m := strSlc2Map(str)
	var a float64
	if a == 0.0 {
		echo(true)
	}
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
