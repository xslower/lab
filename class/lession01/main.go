package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

func toMd5(in string) {
	h := md5.New()
	io.WriteString(h, in)
	ret := h.Sum(nil)
	fmt.Println(len(ret))
	m5 := hex.EncodeToString(ret)
	fmt.Println(m5, len(m5))
}

func toSha256(in string) {
	h := sha256.New()
	io.WriteString(h, in)
	ret := h.Sum(nil)
	fmt.Println(len(ret))
	s2 := hex.EncodeToString(ret)
	fmt.Println(s2, len(s2))
}

func main() {
	str := `hahaha`
	toMd5(str)
	toSha256(str)
}
