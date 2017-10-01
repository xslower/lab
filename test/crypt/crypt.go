package main

import (
	"fmt"
	"github.com/resure-tech/lib/goutils/utils"
)

func main() {

	s := []byte(`abcdef`)
	k := []byte(`123456789012345678901234`)
	en, e := utils.TripleDesEncrypt(s, k)
	echo(string(en), e)
	org, e := utils.TripleDesDecrypt(en, k)
	echo(string(org), e)
}
func echo(d ...interface{}) {
	fmt.Println(d...)
}
func place() {

	_ = utils.Rand(10)

}
