package main

import (
	// "crypto"
	// "crypto/ecdsa"
	// "crypto/elliptic"
	// "crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"github.com/btcsuite/btcd/btcec"
)

func NewEccKey() (epri *EccPriKey, epub *EccPubKey) {
	pri, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		echo(err)
		return
	}
	pub := (*btcec.PublicKey)(&pri.PublicKey)
	epri = &EccPriKey{pri}
	epub = &EccPubKey{pub}
	return
}

type EccPriKey struct {
	*btcec.PrivateKey
}
type EccPubKey struct {
	*btcec.PublicKey
}

func ToEccPriKey(priStr string) (epi *EccPriKey) {
	priBytes, err := hex.DecodeString(priStr)
	if err != nil {
		echo(err)
		return
	}
	priKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), priBytes)
	epi = &EccPriKey{priKey}
	return
}

func ToEccPubKey(pubStr string) (epu *EccPubKey) {
	pubBytes, err := hex.DecodeString(pubStr)
	if err != nil {
		echo(err)
		return
	}
	pubKey, err := btcec.ParsePubKey(pubBytes, btcec.S256())
	if err != nil {
		echo(err)
		return
	}
	epu = &EccPubKey{pubKey}
	return
}

func (this *EccPriKey) GetPri() (pri string) {
	var stream = this.Serialize()
	pri = hex.EncodeToString(stream)
	return
}

func (this *EccPubKey) GetPub() (pub string) {
	var stream = this.SerializeCompressed()
	pub = hex.EncodeToString(stream)
	return
}

func (this *EccPriKey) Sign(src []byte) (signed []byte) {
	var h = sha256.New()
	h.Write(src)
	var hash = h.Sum(nil)
	signature, err := this.PrivateKey.Sign(hash)
	if err != nil {
		echo(err)
		return
	}
	signed = signature.Serialize()
	return
}

func (this *EccPriKey) Decode(cipher []byte) (plain []byte, err error) {
	plain, err = btcec.Decrypt(this.PrivateKey, cipher)
	return
}

func (this *EccPubKey) Verify(signed, src []byte) bool {
	var h = sha256.New()
	h.Write(src)
	var hash = h.Sum(nil)
	signature, err := btcec.ParseSignature(signed, btcec.S256())
	if err != nil {
		return false
	}
	return signature.Verify(hash, this.PublicKey)
}

func (this *EccPubKey) Encode(plain []byte) (cipher []byte, err error) {
	cipher, err = btcec.Encrypt(this.PublicKey, plain)
	return
}
