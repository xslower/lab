package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var (
	a = flag.Int(`a`, 1, ``)
	b = flag.Int(`b`, 2, ``)
)

func handleHttp(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var x = r.Form.Get(`x`)
	var y = r.Form.Get(`y`)
	w.Write([]byte(`x=` + x + `y=` + y))
}

func echo(msg interface{}) {
	if b, ok := msg.([]byte); ok {
		fmt.Println(string(b))
	} else {
		fmt.Println(msg)
	}
}

type pack struct {
	Aa int
	Bb string
}

func main() {
	flag.Parse()
	echo(*a)
	echo(*b)

	p := &pack{1, `aaaaaaaaaaaaa`}
	str, _ := json.Marshal(p)
	fn := `db.txt`
	f, err := os.OpenFile(fn, os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		echo(err)
	}
	f.Write(str)
	f.Write([]byte{'\n'})
	f.Seek(0, 0)
	// f, _ = os.Open(fn)
	str2, _ := ioutil.ReadAll(f)
	str3 := strings.TrimSpace(string(str2))
	blks := strings.Split(str3, "\n")
	p2 := &pack{}
	echo(p2)
	_ = json.Unmarshal([]byte(blks[0]), p2)
	echo(p2)

	f.Close()
	//http.HandleFunc(`/`, handleHttp)
	//http.ListenAndServe(`:80`, nil)
}
