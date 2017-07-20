package main

import (
	"fmt"
	// "github.com/btcsuite/btcd/btcec"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

//首先，不能设计总帐号，这样即使有流水日志，但是依然需要对账。信用卡就是例子
//因为没有总账，1)找零，需要有多个输出；2)合并花费，所以需要多个输入
//
//功能：
//1.创建帐号
//2.存储：序列化、反序列化
//3.测试用，发行货币
//
//1.查询剩余多少钱
//2.创建交易
type TxIn struct {
	Hash string
	Idx  int
	Sign string //保存的是针对
}

func (this *TxIn) Serialize() (s string) {
	s = this.Hash + `:` + strconv.Itoa(this.Idx) + `:` + this.Sign
	return
}
func (this *TxIn) UnSerialize(s string) {
	part := strings.Split(s, `:`)
	if len(part) != 3 {
		panic(`TxIn unserialize error: ` + s)
	}
	this.Hash = part[0]
	this.Idx, _ = strconv.Atoi(part[1])
	this.Sign = part[2]

}

type TxOut struct {
	Value    int
	PkScript string //保存的是公钥
}

func (this *TxOut) Serialize() (s string) {
	s = strconv.Itoa(this.Value) + `:` + this.PkScript
	return
}
func (this *TxOut) UnSerialize(s string) {
	part := strings.Split(s, `:`)
	if len(part) != 2 {
		panic(`TxOut unserialize error: ` + s)
	}
	this.Value, _ = strconv.Atoi(part[0])
	this.PkScript = part[1]
}

type Transaction struct {
	Hash   string
	TxIns  []*TxIn
	TxOuts []*TxOut
}

func (this *Transaction) MkHash() {
	h := sha256.New()
	s := this.serial()
	h.Write([]byte(s))
	hashed := h.Sum(nil)
	this.Hash = hex.EncodeToString(hashed)
}
func (this *Transaction) serial() (s string) {
	if len(this.TxIns) > 0 {
		s = this.TxIns[0].Serialize()
		for i := 1; i < len(this.TxIns); i++ {
			s += `|` + this.TxIns[i].Serialize()
		}
	}
	s += `!`
	if len(this.TxOuts) == 0 {
		panic(`trsanction error: no txOut!`)
	}
	s += this.TxOuts[0].Serialize()
	for i := 1; i < len(this.TxOuts); i++ {
		s += `|` + this.TxOuts[i].Serialize()
	}
	return
}
func (this *Transaction) Serialize() (s string) {
	if this.Hash == `` {
		this.MkHash()
	}
	s = this.Hash + `!` + this.serial()
	return
}
func (this *Transaction) UnSerialize(s string) {
	part := strings.Split(s, `!`)
	if len(part) != 3 {
		panic(`unserialize transaction error!` + s)
	}
	this.Hash = part[0]
	// this.TxIns = []*TxIn{}
	txIns := strings.Split(part[1], `|`)
	for _, tinStr := range txIns {
		tin := &TxIn{}
		tin.UnSerialize(tinStr)
		this.TxIns = append(this.TxIns, tin)
	}
	txOuts := strings.Split(part[2], `|`)
	for _, toutStr := range txOuts {
		tout := &TxOut{}
		tout.UnSerialize(toutStr)
		this.TxOuts = append(this.TxOuts, tout)
	}
}

type Db struct {
	txList []*Transaction
}

func (this *Db) Save() {
	f, err := os.OpenFile(db_file, os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		echo(err)
	}
	for _, tx := range this.txList {
		f.WriteString(tx.Serialize())
		f.WriteString("\n")
	}
}
func (this *Db) Load() {
	f, err := os.Open(db_file)
	if err != nil {
		echo(err)
	}
	str, _ := ioutil.ReadAll(f)

	part := strings.Split(string(str), "\n")
	for _, txStr := range part {
		tx := &Transaction{}
		tx.UnSerialize(txStr)
		this.txList = append(this.txList, tx)
	}
}
func (this *Db) AddTx(tx *Transaction) {
	this.txList = append(this.txList, tx)
}

var (
	c       = flag.String(`c`, `help`, `-c command`)
	p       = flag.String(`p`, ``, `-p parameter`)
	db_file = `tx_list`
	command func()
	mapping = map[string]func(p string){
		`create`: cmdCreateAccount,
		`help`:   cmdHelp,
		`money`:  cmdCreateMoney,
	}
	db = &Db{}
)

func init() {
	// mapping = map[string]func()
	db.Load()
}

//改为http吧，cmd不方便数据展示
func cmdHelp(p string) {
	echo(`-c:`)
	echo(`create: create account`)
}

func cmdCreateAccount(p string) {
	epri, epub := NewEccKey()
	pri := epri.GetPri()
	echo(`private:`, pri)
	pub := epub.GetPub()
	echo(`public:`, pub)
	return
}

func cmdCreateMoney(pub string) {
	txOut := &TxOut{50, pub}
	tx := &Transaction{}
	tx.TxOuts = append(tx.TxOuts, txOut)
	db.AddTx(tx)
	db.Save()
}

func cmdFindTx(pub string) {
	tos := []*TxOut{}
	for _, tx := range db.txList {
		for _, to := range tx.TxOuts {
			if to.PkScript == pub {
				tos = append(tos, to)
			}
		}
	}

}

func main() {
	flag.Parse()
	if f, ok := mapping[*c]; ok {
		f(*p)
	} else {
		cmdHelp(``)
	}

}

func echo(i ...interface{}) {
	fmt.Println(i...)
}
