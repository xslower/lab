package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/xslower/lab/test/thrift/gen-go/la"
)

func main() {
	addr := `106.15.198.173:8123`
	sock, _ := thrift.NewTSocket(addr)
	// trans := transFty.GetTransport(sock)
	// transFty := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	// transFty := thrift.NewTTransportFactory()
	protoFty := thrift.NewTBinaryProtocolFactoryDefault()

	trans := thrift.NewTBufferedTransport(sock, 4096)
	cli := la.NewKickArtClientFactory(trans, protoFty)
	trans.Open()
	ret, err := cli.Kick(123, `haha`)
	trans.Close()
	fmt.Println(ret, err)
}
