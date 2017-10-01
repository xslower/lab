package main

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"io/ioutil"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	_ = b
	c, _ := ioutil.ReadAll(r.Body)
	echo(string(c), string(b))
	fmt.Fprint(w, string(c))
}

func main() {

	log := zerolog.New(os.Stdout)
	log.Info().Msg(msg)
	log.Info().Msg(`haha`)
	http.HandleFunc(`/`, hello)
	// http.ListenAndServe(`:81`, nil)
}

func echo(i ...interface{}) {
	fmt.Println(i...)
}
func placeHolder() {
	_ = json.Decoder{}
	_ = os.Stdout
	// _ = jsoniter.Nil
}
