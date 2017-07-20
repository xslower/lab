package main

import (
	"unsafe"
	//	`encoding/json`
	"fmt"
	// "github.com/astaxie/beego/config"
	"reflect"
	"strconv"
	"strings"
	// `greentea/utils`
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

var (
	dsn = "rdsroot:yinnuo123!@#@tcp(rm-uf6tz1g9l077nkd0do.mysql.rds.aliyuncs.com:3306)/test?charset=utf8"
	//	dsn = "rdsroot:yinnuo123!@#@tcp(rm-uf6tz1g9l077nkd0d.mysql.rds.aliyuncs.com:3306)/test?charset=utf8"

	db, _ = sql.Open(`mysql`, dsn)
)

// Model Struct
type Mytest struct {
	Id         int    `orm:"auto"`
	Name       string `orm:"size(100)"`
	Category   string
	CreateTime string
}

func init() {
	// register model
	orm.RegisterModel(new(Mytest))

	// set default database
	orm.RegisterDataBase("default", "mysql", dsn, 30)

}
func testBeeOrm() {
	o := orm.NewOrm()

	//	user := Mytest{Name: "slene"}

	//	// insert
	//	id, _ := o.Insert(&user)
	//	echo(id)

	//	// update
	//	user.Name = "astaxie"
	//	num, _ := o.Update(&user)
	u := Mytest{Id: 1}
	_ = o.Read(&u)

	// delete
	//	num, _ = o.Delete(&u)
	//	echo(num)
}

func testNaked() {
	rows, _ := db.Query(`SELECT id,name,category FROM mytest WHERE id=1`)
	rows.Next()
	var mt = &Mytest{}
	rv := reflect.ValueOf(mt).Elem()
	fifc := make([]interface{}, rv.NumField())
	col, _ := rows.Columns()
	for i, cn := range col {

		v := rv.FieldByName(cn)
		switch rv.Kind() {
		case reflect.Int:
			fifc[i] = (*int)(unsafe.Pointer(v.UnsafeAddr()))
		case reflect.String:
			fifc[i] = (*string)(unsafe.Pointer(v.UnsafeAddr()))
		}
	}
	rows.Scan(fifc...)
	rows.Close()
	echo(mt)
	//	echo(col, mt)
}

func main() {
	testNaked()
	//	start := time.Now()
	//	for i := 0; i < 100; i++ {
	//		testNaked()
	//		//	testBeeOrm()
	//	}
	//	end := time.Now()
	//	echo(end.Sub(start))

	//	start = time.Now()
	//	for i := 0; i < 100; i++ {
	//		//		testNaked()
	//		testBeeOrm()
	//	}
	//	end = time.Now()
	//	echo(end.Sub(start))

}
func final() {
	if exception := recover(); exception != nil {
		log.Println(exception)
	}
}
func throw(err error) {
	if err != nil {
		panic(err.Error())
	}
}
func check(err error) bool {
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}
func echo(i ...interface{}) {
	fmt.Println(i...)
}

func logit(data ...interface{}) {
	log.Println(data...)
}
func placeHolder() {
	_ = strings.Index(`abc`, `a`)
	_ = strconv.Itoa(1)
	_ = time.Now()
	_ = os.O_WRONLY
	_ = filepath.Ext(``)
	_ = reflect.Array
}
