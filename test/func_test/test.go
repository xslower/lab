package main

import (
	"encoding/json"
	"fmt"
	// `greentea/orm`
	"errors"
	// js "greentea/json"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"runtime/pprof"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	aaa int
)

type hehe int8
type HaHaHa struct {
	Aa string `imatag:"haha"`
	Bb int32
}

func (this HaHaHa) Hello(s string) {
	echo(`hello`, s)
}

type IA interface {
	aa(a int)
}
type IB interface {
	bb()
}

type AA struct {
	Aa int
	Bb int32
}

func (this *AA) aa(a int) {
	echo(a)
}

func (this *AA) bb() {
	echo(`bb`)
}

type Bb struct {
	IB
}

func getNextStep(cur int) (ret int) {
	if cur < 100 {
		return 100
	} else if cur < 1000 {
		return cur
	}
	ret = cur * 30 / 100
	return
}

func main() {

}

func prof() {
	memFile, err := os.Create(`memfile.log`)
	throw(err)
	pprof.WriteHeapProfile(memFile)
	defer memFile.Close()

}

func thr() {
	panic(`haha`)
}

func final() {
	if exception := recover(); exception != nil {
		var j, e = json.Marshal(exception)
		fmt.Println(j, e)
	}
}
func throw(err error) {
	if err != nil {
		panic(err)
	}
}
func check(err error) bool {
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}
func echo(i ...interface{}) {
	fmt.Println(i...)
}

func logit(data ...interface{}) {
	log.Println(data...)
}
func placeHolder() {
	_ = strings.Index(`abc`, `a`)
	_ = strconv.Itoa(1)
	_ = time.Now()
	_ = os.O_WRONLY
	_ = filepath.Ext(``)
	_, _ = ioutil.TempDir(`/tmp/`, `haha`)
	var a = []int{1, 2}
	sort.Ints(a)
	_ = reflect.TypeOf(a)
	_, _ = json.Marshal(`a`)
	_ = errors.New(``)
	_ = rand.New(rand.NewSource(0))
	_ = math.E
}
