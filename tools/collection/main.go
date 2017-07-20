package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	var s = time.Now()
	var dir = `E:/Work/Js/project/puzzle/res`
	var fis, _ = ioutil.ReadDir(dir)
	for _, fi := range fis {
		if fi.IsDir() || fi.Name() == `_index_` {
			continue
		}
		fmt.Println(fi.Name())
	}
	var e = time.Now()
	e.Add(time.Duration(1000000))
	fmt.Println(e.Sub(s))
}
