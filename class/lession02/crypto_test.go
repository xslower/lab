package main

import (
	"testing"
)

func TestEccPk(t *testing.T) {
	pri, pub := NewEccKey()
	priStr := pri.GetPri()
	pubStr := pub.GetPub()
	echo(priStr)
	echo(pubStr)
	pri = ToEccPriKey(priStr)
	pub = ToEccPubKey(pubStr)
	msg := []byte(`hello world`)
	cipher, _ := pub.Encode(msg)
	plain, _ := pri.Decode(cipher)
	echo(plain)
	signed := pri.Sign(msg)
	veri := pub.Verify(signed, msg)
	echo(veri)
}

func TestRsaPk(t *testing.T) {
	pri, pub := NewRsaKey()
	priStr := pri.GetPri()
	pubStr := pub.GetPub()
	echo(priStr)
	echo(pubStr)
	pri = ToRsaPriKey(priStr)
	pub = ToRsaPubKey(pubStr)
	msg := []byte(`rsa hahaha`)
	cipher, _ := pub.Encode(msg)
	plain, _ := pri.Decode(cipher)
	echo(plain)
	signed := pri.Sign(msg)
	veri := pub.Verify(signed, msg)
	echo(veri)
}
