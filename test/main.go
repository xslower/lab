package main

import (
	"fmt"
)

func main() {
	s := `abcdef`
	if s == nil {
		echo(`nil`)
	}

}

func echo(d ...interface{}) {
	fmt.Print(d...)
}
