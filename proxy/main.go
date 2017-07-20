/**
 * dbproxy.
 * php 访问db的代理，暂时只支持写代理。
 * 重要功能：
 * 1. 全局id
 * 2. 忽略mysql的复制机制，直接多次写入
 * 3. 队列
 * 4. 连接池
 */

package main

import (
	`flag`
	`github.com/astaxie/beego/config`
	"net/http"
	`strconv`
)

var (
	c         = flag.String(`c`, `config.ini`, `-c /path/to/config/dir`)
	icr_start = 1
	icr_step  = 1
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			logit(err)
		}
	}()
	flag.Parse()
	cnfPth := *c + `/config.ini`
	cnf, err := config.NewConfig(`ini`, cnfPth)
	throw(err)
	global, err := cnf.GetSection(`global`)
	throw(err)
	icr, err := cnf.GetSection(`icr_conf`)
	throw(err)
	start, err := strconv.Atoi(icr[`start`])
	if err == nil {
		icr_start = start
	}
	step, err := strconv.Atoi(icr[`step`])
	if err == nil {
		icr_step = step
	}
	var nodes = global[`nodes`]
	var php = global[`php`]
	var conv = *c + `/conv.php`
	gConnManager = NewConnManager(php, conv, nodes)
	http.HandleFunc(`/ctrl`, serveCtrl)
	http.HandleFunc(`/exec`, serveExec)
	http.HandleFunc(`/dbproxy`, serveProxy)
	err = http.ListenAndServe(global[`addr`], nil)
	if err != nil {
		println(err.Error())
	}
	gConnManager.exit()
}

// func handleSignal() {
// 	var c = make(chan os.Signal)
// 	signal.Notify(c, os.Interrupt, os.Kill, os.)
// }
