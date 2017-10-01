package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type User struct {
	Name string
	Age  int
}

type Out struct {
	Age  int
	Name string
}

func New() *User {
	return &User{Name: "viney", Age: 32}
}

type Config struct {
	LogPath   string `yaml:"log_path"`
	LogLevel  uint8  `yaml:"log_level"`
	Port      uint16
	LocalPow  uint8  `yaml:"local_pow"`
	CacheDrv  string `yaml:"cache_drv"`
	CacheArgs string `yaml:"cache_args"`
	//LocalPow uint8  `json:"local_pow"`
	//GodisArg string `json:"godis_arg"`
}

func main() {
	bytes, _ := ioutil.ReadFile(`folder/cfg.yml`)
	c := &Config{}
	_ = yaml.Unmarshal(bytes, &c)
	echo(c.CacheArgs)
	return
	// 实例化User
	u := New()

	// 对User编码
	b, err := yaml.Marshal(u)
	if err != nil {
		fmt.Println("encode fail: " + err.Error())
	}
	echo(string(b))
	// 对User解码
	var out Out
	if err := yaml.Unmarshal(b, &out); err != nil {
		fmt.Println("decode fail: " + err.Error())
	}

	fmt.Println(out)
}
func echo(d ...interface{}) {
	fmt.Println(d...)
}

func place() {
	_ = 1
}
