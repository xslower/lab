package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

var (
	testDir = ``
	dir     = ``
)

func servFileList(w http.ResponseWriter, r *http.Request) {
	var res = dir + `\res`
	var fis, _ = ioutil.ReadDir(res)
	for _, fi := range fis {
		if fi.IsDir() {
			continue
		}
		fmt.Fprint(w, fi.Name(), "\r\n")
	}
}

func main() {
	var start = time.Now()
	if testDir != `` {
		dir = testDir
	} else {
		dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	}

	fmt.Println(dir)
	http.HandleFunc(`/res_index`, servFileList)
	http.Handle(`/`, http.FileServer(http.Dir(dir)))
	go http.ListenAndServe(`:80`, nil)
	//监听中断
	var c = make(chan os.Signal)
	signal.Notify(c)
	var sig = <-c
	handleSignal(sig)
	//计时
	var end = time.Now()
	fmt.Println(end.Sub(start))
	time.Sleep(1000)
}

func handleSignal(sig os.Signal) {
	var log = ``
	if sig == syscall.SIGINT {
		log = `系统中止`
	}
	fmt.Println(sig, log)
}
