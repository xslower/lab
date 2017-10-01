package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"encoding/json"
	"github.com/json-iterator/go"
	// mjson "github.com/resure-tech/lib/encoding/json"
)

type Aa struct {
	A int    `json:"aaa" auto:"1"`
	B string `json:"bbb"`
	C int16  `json:"c"`
	D []byte `json:"d"`
	E uint8
	f uint8
}
type Bb struct {
	Aa
}

func (a *Aa) Whaha() {
	echo(`aa`)
}

type ff func()

func prof() {
	cnt := 5000
	tt := &Aa{A: 8, B: `pwofnsodpfijw`, C: 9977, D: []byte{'a', 'b', 'c', 'z'}, E: 244}
	a := []*Aa{tt, tt, tt}
	start := time.Now()
	// b := &Aa{}
	var stm []byte
	for i := 0; i < cnt; i++ {
		// stm, _ = mjson.Marshal(a)
		// _ = ByteToStt(b, stm)

	}
	end := time.Now()
	echo(string(stm))
	echo(end.Sub(start))
	start2 := time.Now()
	for i := 0; i < cnt; i++ {
		stm, _ = json.Marshal(a)
		// json.Unmarshal(stm, b)
	}
	end2 := time.Now()
	echo(string(stm))
	echo(end2.Sub(start2))

	start3 := time.Now()
	for i := 0; i < cnt; i++ {
		stm, _ = jsoniter.Marshal(a)
		// jsoniter.Unmarshal(stm, b)
	}
	end3 := time.Now()
	echo(end3.Sub(start3))
	// a := &Aa{}
	// start := time.Now()
	// cnt := 1000000
	// for i := 0; i < cnt; i++ {
	// 	r := reflect.ValueOf(a).Elem().Field(0)
	// 	if r.Kind() == reflect.Int {
	// 		r.SetInt(999)
	// 	}

	// }
	// end := time.Now()
	// echo(end.Sub(start))

	// echo(a)
	// start2 := time.Now()
	// for i := 0; i < cnt; i++ {
	// 	t := reflect.TypeOf(a).Elem().Field(0)
	// 	if t.Type.Kind() == reflect.Int {
	// 		ptr := uintptr(unsafe.Pointer(a)) + t.Offset
	// 		*((*int)(unsafe.Pointer(ptr))) = 888
	// 	}
	// }
	// end2 := time.Now()
	// echo(end2.Sub(start2))
	// echo(a)
}

func main() {
	a := `aa`
	ref := reflect.ValueOf(a)
	echo(ref.String())
}

func assign() {
	//struct
	a := &Aa{}
	vo := reflect.ValueOf(a).Elem()
	to := vo.Type()
	for i := 0; i < to.NumField(); i++ {
		ft := to.Field(i)
		fv := vo.Field(i)
		echo(ft.Name, ft.Type)
		if ft.Type.Name() == `int` {
			*((*int)(unsafe.Pointer(fv.UnsafeAddr()))) = 50
			//			fv.Set(reflect.ValueOf(20))
		}
		echo(fv.Type(), fv.Interface())
	}
	//direct assign to interface
	var p *Aa = &Aa{}
	a = &Aa{A: 123}
	ref := reflect.ValueOf(p).Elem()
	ref.Set(reflect.ValueOf(a))
	echo(p)
}

func usafe() {
	a := Aa{}
	ref_to := reflect.TypeOf(&a).Elem()
	//	ref_vo := reflect.ValueOf(&a).Elem()
	//	ref_vo.FieldByName()
	field, _ := ref_to.FieldByName(`A`)
	fptr := uintptr(unsafe.Pointer(&a)) + field.Offset
	*((*int)(unsafe.Pointer(fptr))) = 50
	echo(a)
}

func sliceAppend() {
	slc := make([]Aa, 0)
	ptr := &slc
	ref_type := reflect.TypeOf(slc).Elem()
	ref_ptr := reflect.ValueOf(ptr)
	ref_val := ref_ptr.Elem()
	one := reflect.New(ref_type).Elem()
	one.Field(0).SetInt(88)
	refv2 := reflect.Append(ref_val, one)
	if ref_val == refv2 { //两边不相等
		echo(`rv = rv2`)
	}
	//	rv.Index(0).Field(0).SetInt(99)
	ref_val.Set(refv2) //此步是关键，不然无法影响到原指针ptr
	echo(ref_val, refv2, slc)

}

func sliceCreate() {
	slc := make([]Aa, 0)
	//	ref_type := reflect.TypeOf(slc)
	var ifc interface{} = &slc
	ref_val := reflect.ValueOf(ifc).Elem()
	new_val := reflect.MakeSlice(ref_val.Type(), 3, 3)
	fld := new_val.Index(0).Field(0)
	fld2 := fld
	if fld.Kind() == reflect.Int {

		fld2.SetInt(99)
	}
	ref_val.Set(new_val)
	echo(slc, new_val)

}

func tag() {
	var a = Aa{}
	t := reflect.TypeOf(a)
	tag := t.Field(0).Tag
	x, y := tag.Lookup(`auto`)
	echo(tag.Get(`col`), x, y)
}

func getAddrNSet() {
	a := &Aa{}
	v := reflect.ValueOf(a).Elem()
	f := v.Field(0).Addr().Interface()
	switch ff := f.(type) {
	case *int:
		*ff = 100
	}
	echo(a)
}
func testZero() {
	a := ``
	v := reflect.ValueOf(a)
	echo(v.IsValid())
	zero := reflect.Zero(v.Type())
	if zero.Interface() == v.Interface() {
		echo(`haha`)
	}
	echo(zero, a)
}
func testIfcCopy() {
	//	i := []int{1, 23, 4}
	i := map[int]int{1: 2, 2: 3}
	//	var d interface{} = i
	v := reflect.ValueOf(i)
	e := v.MapIndex(reflect.ValueOf(1))
	e.SetInt(77)
	echo(i, v, e)
}
func getFieldValue(fld reflect.Value) (val []byte) {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, fld.Interface())
	val = buf.Bytes()
	return
}

func setFieldValue(fldPrr reflect.Value, val []byte) {
	fld := fldPrr.Elem()
	buf := bytes.NewBuffer(val)
	binary.Read(buf, binary.BigEndian, fld.Addr().Interface())
	return
}
func testConvert() {

	var a int64 = 199
	ref := reflect.ValueOf(a)
	by := getFieldValue(ref)
	by = append(by, by...)
	echo(by)
	var b int64
	br := reflect.ValueOf(&b)
	setFieldValue(br, by)
	echo(b)

}

type Slice struct {
	arr uintptr
	ln  int
	cp  int
}

func reverseBytes(b []byte) {
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}
}
func final() {
	if exception := recover(); exception != nil {
		log.Println(exception)
	}
}
func throw(err error) {
	if err != nil {
		panic(err.Error())
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
	_ = reflect.Array
	_ = errors.New(``)
}
