package main

import (
	"fmt"
	"github.com/json-iterator/go"
	"github.com/resure-tech/lib/encoding/json"
	"github.com/vmihailenco/msgpack"
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
	Host string `json:"host"`
	User string
	Pass string
}

func main() {
	in := Info{`abcd`, `aaa`, `bbb`}
	b, _ := msgpack.Marshal(in)
	on := Info{}
	_ = msgpack.Unmarshal(b, &on)
	echo(b, string(b), on)
}
func echo(i ...interface{}) {
	fmt.Println(i...)
}
func placeHolder() {
	_ = json.Decoder{}
	_ = jsoniter.Nil
	var _ msgpack.Marshaler
}
