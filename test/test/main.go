package main

import (
	"encoding/json"
	"fmt"
	"github.com/resure-tech/lib/goutils/utils"
	"net/url"
	"os"
	"reflect"
	"strings"
)

type Slice struct {
	arr uintptr
	ln  int
	cp  int
}

type Stt struct {
	A string
	C map[string]string
}

func HashPassword(pwd string) (hashed string) {
	//todo salt改为随机码，并存在db里
	salt := `chrdw,hdhxt`
	hashed = utils.Sha256(salt + pwd)
	return
}
func de() (a string, b int) {
	defer func() {
		a = `opafijew`
		b = 999
	}()
	return
}
func main() {
	f, _ := os.OpenFile(`a.log`, os.O_CREATE, 0666)
	b := utils.NewBytes(100)
	b.Add('a')
	b.Add('\n')
	n, _ := f.Write(b.Bytes())
	n, _ = f.Write(b.Bytes())
	echo(n)
}
func testMail() {

	host := `smtp.163.com:25`
	user := `xslower@163.com`
	pass := `sh2345isAshitSB+`
	to := `xslower@163.com`
	subject := `this is a subject`
	body := `<html>this is a body</html>`
	err := utils.SendMail(host, user, pass, to, subject, body)
	echo(err)
}
func echo(d ...interface{}) {
	fmt.Println(d...)
}
func place() {
	_ = reflect.Map
	_ = utils.Rand(10)
	_ = json.Encoder{}
	_ = url.PathEscape(``)
	_ = strings.Index(``, ``)
	_ = os.Stdout
}
