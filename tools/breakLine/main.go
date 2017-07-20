package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

var (
	file  = `F:\Pictures\Temp\1.txt`
	input = ``
)

func main() {
	str := `[\~\!\@\#\$\%\^\&\*\(\)\-\+\[\]\{\}\|\/\\\.\,\'\"\?\:\;\s\w]{6,}`

	reg := regexp.MustCompile(str)
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		echo(err)
		return
	}
	ret := reg.ReplaceAll(bytes, []byte(``))
	// ret := reg.ReplaceAllString(str, ``)
	// ret := reg.FindAllString(input, -1)
	ioutil.WriteFile(file, ret, 0666)
	// echo(string(ret))
}

func echo(msg ...interface{}) {
	fmt.Println(msg...)
}

func breakLine() {
	//todo 为什么会有两段长句连起来的情况
	//把所有中括号内部的内容删除
	reg := regexp.MustCompile(`(?U)\[.*\]`)
	input = reg.ReplaceAllString(input, ``)
	//删除两个斜杠//后面的内容
	reg = regexp.MustCompile(`(?U)\/\/.*\n`)
	input := reg.ReplaceAllString(input, ``)

	// 把与句号等同的标点都替换为句号，用于断句
	symbol := []string{"\n", "！", "？", "；"}
	for _, sym := range symbol {
		input = strings.Replace(input, sym, "。", -1)
	}
	symbol = []string{"、", "：", "“", "”", `"`}
	for _, sym := range symbol {
		input = strings.Replace(input, sym, "，", -1)
	}
	out := []string{}
	// 以自然句为单位，以免出现句子A结尾+句子B开头的组合
	sentence := strings.Split(input, "。")
	for _, sen := range sentence {
		sen = strings.TrimSpace(sen)
		if len(sen) == 0 {
			continue
		}
		parts := strings.Split(sen, `，`)
		var tmp = ``
		for _, p := range parts {
			p = strings.TrimSpace(p)
			if len(p) == 0 {
				continue
			}
			if len(tmp)+len(p) < 46 || len(tmp) < 10 {
				if len(tmp) == 0 {
					tmp = p
				} else {
					tmp += ` ` + p
				}
			} else {
				out = append(out, tmp)
				tmp = p
			}

		}
		if len(tmp) < 13 {
			tmp = out[len(out)-1] + ` ` + tmp
			out[len(out)-1] = tmp
		} else {
			out = append(out, tmp)
		}

	}

	for i := 0; i < len(out); i++ {
		echo(out[i])
		// if i < len(out)-1 && len(out[i]) < 16 && len(out[i+1]) < 45 {
		// 	echo(out[i], out[i+1])
		// 	i++
		// } else {
		// }
	}
}
