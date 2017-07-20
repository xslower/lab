package main

import (
	"errors"
	"fmt"
	"unsafe"
)

const (
	// Maximum number of key/value pairs a bucket can hold.
	bucketCntBits = 3
	bucketCnt     = 1 << bucketCntBits
)

type String struct {
	ptr uintptr
	ln  int
}

type Slice struct {
	arr uintptr
	ln  int
	cp  int
}

type Map struct {
	// Note: the format of the Hmap is encoded in ../../cmd/internal/gc/reflect.go and
	// ../reflect/type.go. Don't change this structure without also changing that code!
	count     int // # live cells == size of map.  Must be first (used by len() builtin)
	flags     uint8
	B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
	noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
	hash0     uint32 // hash seed

	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

	// If both key and value do not contain pointers and are inline, then we mark bucket
	// type as containing no pointers. This avoids scanning such maps.
	// However, bmap.overflow is a pointer. In order to keep overflow buckets
	// alive, we store pointers to all overflow buckets in hmap.overflow.
	// Overflow is used only if key and value do not contain pointers.
	// overflow[0] contains overflow buckets for hmap.buckets.
	// overflow[1] contains overflow buckets for hmap.oldbuckets.
	// The first indirection allows us to reduce static size of hmap.
	// The second indirection allows to store a pointer to the slice in hiter.
	overflow *[2]*[]*bmap
}

// A bucket for a Go map.
type bmap struct {
	// tophash generally contains the top byte of the hash value
	// for each key in this bucket. If tophash[0] < minTopHash,
	// tophash[0] is a bucket evacuation state instead.
	tophash [bucketCnt]uint8
	// Followed by bucketCnt keys and then bucketCnt values.
	// NOTE: packing all the keys together and then all the values together makes the
	// code a bit more complicated than alternating key/value/key/value/... but it allows
	// us to eliminate padding which would be needed for, e.g., map[int64]int8.
	// Followed by an overflow pointer.
}

func forceChangeSlice() {
	b := make([]byte, 0, 100)
	c := b
	c = append(c, 'a', 'b')
	echo(b, c)
	ptr := (*Slice)(unsafe.Pointer(&b))
	ptr.ln += 2
	b = *(*[]byte)(unsafe.Pointer(ptr))
	echo(b, c)
}
func testMap() {
	var m = make(map[string]int32, 10)
	m[`hello`] = 123
	//	m[`aa`] = 19
	ptr := (*Map)(unsafe.Pointer(&m))
	echo(m, *ptr)
}
func getString() string {
	return `abcdef`
}
func testSlice() []int {
	t := make([]int, 0, 5)
	t = append(t, 1, 2, 4)
	return t
}

//s 如果=`abc`，则为常量，无法修改，只能=[]byte才能修改。
func testString(s string) {
	//	b := make([]byte, 0, 100)
	//	b = append(b, 'a', 'b', 'c')
	//	s := string(b)
	t := s
	ptr := (*String)(unsafe.Pointer(&s))
	*(*byte)(unsafe.Pointer(ptr.ptr)) = 'y'
	echo(s, t)
}
func testDeepCopy() {
	//	s := string([]byte{'a', 'b', 'c'})
	s := `abc` + `def`
	b := []byte(s)
	b[1] = 'q'
	echo(string(b), s)
}

var ERR = errors.New(`aaaa`)

func copyIfc(i interface{}) {
	m := i.(map[int]int)
	m[2] = 888
}
func main() {
	m := map[int]int{1: 2, 3: 4}
	copyIfc(m)
	echo(m, ERR)
}
func echo(i ...interface{}) {
	fmt.Println(i...)
}
