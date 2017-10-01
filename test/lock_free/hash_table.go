/**
 * 使用原子操作代替锁的，hash table实现
 * 方便演示，key使用Int64
 */

package main

import (
	`sync/atomic`
	"runtime"
	"fmt"
	"sync"
)

func NewHashTable(pow uint8) (ht *HashTable) {
	num := int64(1 << pow)
	ht = &HashTable{
		mod:  num - 1,
		data: make([]interface{}, num),
		lock: make([]int32, num),
	}
	return
}

type HashTable struct {
	mod  int64 //mod=2^N-1
	data []interface{}
	lock []int32
}

func (this *HashTable) Put(key int64, val interface{}) {
	//取余，
	key = key & this.mod
	addr := &this.lock[key]
	var lck, nlck int32
	for {
		lck = atomic.LoadInt32(addr)
		if lck != 0 { //=0 才意味没有读，也没有写
			//fmt.Println(`[put] write lock occupied`)
			runtime.Gosched()
		}
		nlck = lck - 1
		if atomic.CompareAndSwapInt32(addr, lck, nlck) {
			break
		} else {
			//fmt.Println(`[put] write lock changed`)
			runtime.Gosched()
		}
	}
	this.data[key] = val
	atomic.StoreInt32(addr, 0)
	return
}

func (this *HashTable) Get(key int64) (val interface{}) {
	key = key & this.mod
	addr := &this.lock[key]
	var lck int32
	for {
		lck = atomic.LoadInt32(addr)
		if lck < 0 { //正在写入不能读，其它情况都能读
			//fmt.Println(`[get] read lock occupied`)
			runtime.Gosched()
		}
		atomic.AddInt32(addr, 1)
		break
	}
	val = this.data[key]
	atomic.AddInt32(addr, -1)
	return
}

//func NewMtxHT(pow uint8) (ht *MtxHashTable) {
//	num := int64(1 << pow)
//	ht = &MtxHashTable{
//		mod:  num - 1,
//		//data: make([]interface{}, num),
//		data: make(map[int64]interface{}),
//	}
//	return
//}
//
//type MtxHashTable struct {
//	mod  int64
//	data map[int64]interface{}
//	mtx  sync.RWMutex
//}
//
//func (this *MtxHashTable) Put(key int64, val interface{}) {
//	key = key & this.mod
//	this.mtx.Lock()
//	this.data[key] = val
//	this.mtx.Unlock()
//}
//
//func (this *MtxHashTable) Get(key int64) (val interface{}) {
//	key = key & this.mod
//	this.mtx.RLock()
//	val = this.data[key]
//	this.mtx.RUnlock()
//	return
//}

func NewMtxHT(pow uint8) (ht *MtxHashTable) {
	num := int64(1 << pow)
	ht = &MtxHashTable{
		mod:  num - 1,
		data: make([]interface{}, num),
		//data: make(map[int64]interface{}),
		mtx:make([]sync.RWMutex, num),
	}
	return
}

type MtxHashTable struct {
	mod  int64
	data []interface{}
	mtx  []sync.RWMutex
}

func (this *MtxHashTable) Put(key int64, val interface{}) {
	key = key & this.mod
	this.mtx[key].Lock()
	this.data[key] = val
	this.mtx[key].Unlock()
}

func (this *MtxHashTable) Get(key int64) (val interface{}) {
	key = key & this.mod
	this.mtx[key].RLock()
	val = this.data[key]
	this.mtx[key].RUnlock()
	return
}



type HashTableIfc interface {
	Put(key int64, val interface{})
	Get(key int64) interface{}
}

func placeHolder(){
	fmt.Print(``)
}