package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/xslower/lab/test/thrift/gen-go/la"
)

type kickArt struct {
}

func (this *kickArt) Kick(id int32, name string) (r map[int32]string, err error) {
	echo(`client call `, id, name)
	r = make(map[int32]string)
	r[id] = name
	return
}
func (this *kickArt) Put(na *la.Art) (err error) {
	echo(na)
	return
}

func main() {
	//tf := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	tf := thrift.NewTTransportFactory()
	pf := thrift.NewTBinaryProtocolFactoryDefault()
	st, _ := thrift.NewTServerSocket(`:9090`)
	proc := la.NewKickArtProcessor(&kickArt{})
	server := thrift.NewTSimpleServer4(proc, st, tf, pf)
	server.Serve()
}

func echo(d ...interface{}) {
	fmt.Println(d...)
}
