package main

import (
	`biz/tools/tools`
	`encoding/json`
	`greentea/orm`
	`os/exec`
)

var (
	gConnManager *ConnManager
)

func NewConnManager(php, conv, db string) *ConnManager {
	if !tools.Exist(conv) {
		panic(conv + `is not exist!`)
	}
	var c = &ConnManager{php: php, conv: conv, db: db}
	c.initialize()
	return c
}

type ConnManager struct {
	connAll map[string][]*Conn
	connLcl map[string]*Conn
	php     string
	conv    string
	db      string
}

func (this *ConnManager) realloc() {
	if this.connAll != nil {
		for key, _ := range this.connAll {
			delete(this.connAll, key)
		}
	} else {
		this.connAll = make(map[string][]*Conn)
	}
	if this.connLcl != nil {
		for key, _ := range this.connLcl {
			delete(this.connLcl, key)
		}
	} else {
		this.connLcl = make(map[string]*Conn)
	}
}

//是不是应该增加mutex防止出现些乱七八糟的问题
func (this *ConnManager) initialize() {
	var options = []string{this.conv, this.db}
	ret, err := exec.Command(this.php, options...).Output()
	if err != nil {
		echo(`command error`, err.Error())
	}
	var nodes = map[string]map[string]map[string]string{}
	err = json.Unmarshal(ret, &nodes)
	if err != nil {
		echo(`json error`, err.Error())
	}
	this.realloc()
	for key, node := range nodes {
		for _, conf := range node {
			var cn = NewConn(conf)
			this.connAll[key] = append(this.connAll[key], cn)
			if conf[`role`] == `master` {
				this.connLcl[key] = cn
			}
			go cn.execSql()
		}
	}
}

func (this *ConnManager) getCluster(hkey string) []*Conn {
	return this.connAll[hkey]
}

func (this *ConnManager) getLocal(hkey string) *Conn {
	return this.connLcl[hkey]
}

func (this *ConnManager) exit() {
	for _, cnn := range this.connAll {
		for _, cn := range cnn {
			cn.exit <- true
		}
	}
}

//某个连接失败不影响其它连接
func NewConn(cnf map[string]string) *Conn {
	defer func() {
		if err := recover(); err != nil {
			logit(cnf[`host`] + `connect failed!`)
		}
	}()
	c := orm.NewMysqlConn(cnf)
	this := &Conn{host: cnf[`host`], sql: make(chan string, 10),
		exit: make(chan bool)}
	this.IConn = c
	return this
}

type Conn struct {
	orm.IConn
	host string
	sql  chan string
	exit chan bool
}

func (this *Conn) execSql() {
	defer func() {
		if err := recover(); err != nil {
			logit(err)
		}
		this.execSql()
	}()
	for {
		select {
		case sql := <-this.sql:
			id := this.Exec(sql)
			logit(sql, id)
		case <-this.exit:
			close(this.sql)
			close(this.exit)
			this.Close()
			logit(`Conntention:`, this.host, `Exit!`)
			return
		}
	}

}
