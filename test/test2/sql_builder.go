package main

import (
	`fmt`
	`strconv`
	`strings`
)

/**
* @desc 大体的sql builder思路：
 */
const (
	ZArrayTypeInt = iota
	ZArrayTypeString
)

type ZArray struct {
	Type   int8
	intArr []int
	strArr []string
}

func NewZArray(elements ...interface{}) (z *ZArray) {
	z = &ZArray{}
	if len(elements) == 0 {
		return
	}
	switch v := elements[0].(type) {
	case int:
		z.Type = ZArrayTypeInt
		z.intArr = make([]int, len(elements))
	case string:
		z.Type = ZArrayTypeString
		z.strArr = make([]string, len(elements))
	default:
		fmt.Print(v)
	}

	return
}
func NewIntZArray(elements ...int) (z *ZArray) {
	z = &ZArray{Type: ZArrayTypeInt}
	if len(elements) == 0 {
		return
	}
	z.intArr = make([]int, len(elements))
	_ = copy(z.intArr, elements)
	return
}
func NewStrZArray(elements ...string) (z *ZArray) {
	z = &ZArray{Type: ZArrayTypeString}
	if len(elements) == 0 {
		return
	}
	z.strArr = make([]string, len(elements))
	_ = copy(z.strArr, elements)
	return
}
func (this *ZArray) IntArr() []int {
	return this.intArr
}
func (this *ZArray) StrArr() []string {
	return this.strArr
}

type Expr interface {
	GetRawSql() string
	//GetPrepareSql() string
}

type ExprEqual struct {
	Field string
	Value interface{}
}

func (this *ExprEqual) GetRawSql() string {
	expr := this.Field + ` = `
	switch val := this.Value.(type) {
	case string:
		expr += `'` + val + `'`
	case int:
		expr += strconv.Itoa(val)
	default:
		panic(`value type error`)
	}
	return expr
}

type ExprIn struct {
	Field  string
	Values interface{}
}

func (this *ExprIn) GetRawSql() string {
	expr := this.Field + ` IN (`
	switch val := this.Values.(type) {
	case []int:
		expr += strconv.Itoa(val[0])
		for _, v := range val[1:] {
			expr += `, ` + strconv.Itoa(v)
		}
	case []string:
		expr += `'` + strings.Join(val, `', '`) + `'`
	default:
		fmt.Print(val)
	}
	expr += `)`
	return expr
}

//默认不支持字符串的大小比较
type ExprBetween struct {
	Field string
	Start int
	End   int
}

func (this *ExprBetween) GetRawSql() string {
	return `BETWEEN`
}

type ExprLike struct {
	Field string
	Value string
}

func (this *ExprLike) GetRawSql() string {
	return `LIKE`
}

var ValidOp = map[string]bool{`>`: true, `>=`: true, `<`: true, `<=`: true}

type ExprOp struct {
	Field string
	Op    string
	Value int
}

const (
	AND = `AND`
	OR  = `OR`
)

type Cond struct {
	rel      string
	exprs    []Expr
	subConds []*Cond
}

func (this *Cond) Or() *Cond {
	c := &Cond{rel: OR}
	this.exprs = append(this.exprs, c)
	return c
}

func (this *Cond) LoadMap(condList map[string]interface{}) *Cond {
	for field, value := range condList {
		this.exprs = append(this.exprs, &ExprEqual{field, value})
	}
	return this
}

func (this *Cond) Equal(field string, value interface{}) *Cond {
	this.exprs = append(this.exprs, &ExprEqual{field, value})
	return this
}

func (this *Cond) Between(field string, start, end int) *Cond {
	this.exprs = append(this.exprs, &ExprBetween{field, start, end})
	return this
}
func (this *Cond) In(field string, elements ...interface{}) *Cond {
	return this
}

func NewCond() *Cond {
	c := &Cond{rel: AND}
	return c
}

func (this *Cond) GetRawSql() string {
	statement := ``
	for i, expr := range this.exprs {
		if i > 1 {
			statement += this.rel
		}
		statement += `(` + expr.GetRawSql() + `)`
	}
	return statement
}

func sample() {

}
