package main

import (
	`fmt`
	"net/http"
	`os/exec`
	`strconv`
	`strings`
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

func serveCtrl(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	cmd := r.Form.Get(`cmd`)
	if cmd == `reload` {
		gConnManager.initialize()
		gGid.clear()
		logit(`db proxy reloaded!`)
	} else if cmd == `close` {
		gConnManager.exit()
	}

}

func serveProxy(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	hkey := r.Form.Get(`hkey`)
	if hkey == `` {
		fmt.Fprint(w, `host key is empty!`)
		return
	}
	conns := gConnManager.getCluster(hkey)
	if conns == nil {
		fmt.Fprint(w, `host key is error!`)
		return
	}
	sql := r.Form.Get(`sql`)
	if sql == `` {
		fmt.Fprint(w, `sql is empty`)
		return
	}
	need_pk := r.Form.Get(`need_pk`)
	var idstr = `ok`
	if need_pk != `` {
		table := r.Form.Get(`table`)
		if table == `` {
			fmt.Fprint(w, `table name is empty!`)
		}
		db := r.Form.Get(`db`)
		pk := r.Form.Get(`pk`)
		id := gGid.getId(table)
		if id == 0 {
			id = initGlobalId(hkey, db, table, pk)
		}
		idstr = strconv.Itoa(id)
		pksql := "SET `" + pk + "` = " + idstr + ", "
		search := `set `
		if strings.Index(sql, `SET `) > 0 {
			search = `SET `
		}
		sql = strings.Replace(sql, search, pksql, -1)

	}
	for _, cn := range conns {
		cn.sql <- sql
	}

	fmt.Fprint(w, idstr)

}
