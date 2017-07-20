package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	// `os`
)

func NewRsaKey() (rpri *RsaPriKey, rpub *RsaPubKey) {
	bits := 1024
	priKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		echo(err)
		return
	}
	pubKey := &priKey.PublicKey
	rpri = &RsaPriKey{priKey}
	rpub = &RsaPubKey{pubKey}
	return
}

type RsaPriKey struct {
	*rsa.PrivateKey
}

type RsaPubKey struct {
	*rsa.PublicKey
}

func ToRsaPriKey(priStr string) (rpr *RsaPriKey) {
	priBt, err := hex.DecodeString(priStr)
	if err != nil {
		echo(err)
		return
	}
	pri, err := x509.ParsePKCS1PrivateKey(priBt)
	if err != nil {
		echo(err)
		return
	}
	rpr = &RsaPriKey{pri}
	return
}

func ToRsaPubKey(pubStr string) (rpu *RsaPubKey) {
	pubBt, err := hex.DecodeString(pubStr)
	if err != nil {
		echo(err)
		return
	}
	pubIfc, err := x509.ParsePKIXPublicKey(pubBt)
	if err != nil {
		echo(err)
		return
	}
	pub := pubIfc.(*rsa.PublicKey)
	rpu = &RsaPubKey{pub}
	return
}

func (this *RsaPriKey) GetPri() (pri string) {
	stream := x509.MarshalPKCS1PrivateKey(this.PrivateKey)
	pri = hex.EncodeToString(stream)
	return
}

func (this *RsaPubKey) GetPub() (pub string) {
	stream, err := x509.MarshalPKIXPublicKey(this.PublicKey)
	if err != nil {
		echo(err)
		return
	}
	pub = hex.EncodeToString(stream)
	return
}

func (this *RsaPriKey) Decode(cipher []byte) (plain []byte, err error) {
	plain, err = rsa.DecryptPKCS1v15(rand.Reader, this.PrivateKey, cipher)
	return
}

func (this *RsaPubKey) Encode(plain []byte) (cipher []byte, err error) {
	cipher, err = rsa.EncryptPKCS1v15(rand.Reader, this.PublicKey, plain)
	return
}

func (this *RsaPriKey) Sign(src []byte) (signed []byte) {
	hash := crypto.SHA256
	h := hash.New()
	h.Write(src)
	hashed := h.Sum(nil)
	signed, err := rsa.SignPKCS1v15(rand.Reader, this.PrivateKey, hash, hashed)
	if err != nil {
		echo(err)
	}
	return
}

func (this *RsaPubKey) Verify(signed, src []byte) bool {
	hash := crypto.SHA256
	h := hash.New()
	h.Write(src)
	hashed := h.Sum(nil)
	err := rsa.VerifyPKCS1v15(this.PublicKey, hash, hashed, signed)
	if err != nil {
		return false
	}
	return true
}

func main() {

}

func echo(msg interface{}) {
	if b, ok := msg.([]byte); ok {
		fmt.Println(string(b))
	} else {
		fmt.Println(msg)
	}

}
