package main

import (
	"fmt"
	"github.com/json-iterator/go"
)

var str = []byte(`{
	"01":{
		"host":"rm-uf6tz1g9l077nkd0do.mysql.rds.aliyuncs.com",
		"user":"rdsroot",
		"pass":"yinnuo123!@#"
	},
	"02":{
		"host":"rm-uf6tz1g9l077nkd0do.mysql.rds.aliyuncs.com",
		"user":"rdsroot",
		"pass":"yinnuo123!@#"
	}
	"partition":{
		"default":"01"
	}
}`)

type Info struct {
	Host string
	User string
	Pass string
}

func main() {
	info := map[string]*Info{}
	err := jsoniter.Unmarshal(str, &info)
	for _, in := range info {
		echo(in)
	}
	echo(err)
}
func echo(i ...interface{}) {
	fmt.Println(i...)
}
