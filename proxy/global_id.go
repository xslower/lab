package main

import (
	`strconv`
	`sync`
)

var (
	gGid = &GlobalId{ids: make(map[string]int)}
)

type GlobalId struct {
	ids map[string]int
	mux sync.Mutex //`这里如果并发大的话,可以给每个id分配一个锁mux map[string]sync.Mutex`
}

func (this *GlobalId) getId(key string) int {
	if _, ok := this.ids[key]; !ok {
		return 0
	}
	this.mux.Lock()
	id := this.ids[key]
	id += icr_step
	this.ids[key] = id
	this.mux.Unlock()
	return id
}

func (this *GlobalId) setId(key string, id int) {
	this.mux.Lock()
	this.ids[key] = id
	this.mux.Unlock()
}

func (this *GlobalId) clear() {
	this.mux.Lock()
	for key, _ := range this.ids {
		delete(this.ids, key)
	}
	this.mux.Unlock()
	// this.ids = make(map[string]int)
}

func initGlobalId(hkey, db, table, pk string) int {
	var tn = "`" + table + "`"
	if db != `` {
		tn = "`" + db + "`." + tn
	}
	sql := "select max(`" + pk + "`) from " + tn
	conn := gConnManager.getLocal(hkey)
	_, dset := conn.Query(sql)
	col := dset[0][0]
	id, _ := strconv.Atoi(string(col))
	start := icr_start
	step := icr_step
	if start == step {
		start = 0
	}
	remainder := id % step
	id = id - remainder + step + start
	gGid.setId(table, id)
	return id
}
