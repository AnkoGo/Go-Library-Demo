package main

import (
	"fmt"
	"os"
)

func main19076() {
	mapping := func(key string) string {
		m := make(map[string]string)
		m = map[string]string{
			"world": "kitty",
			"hello": "hi",
		}
		if m[key] != "" {
			return m[key]
		}
		return key
	}
	s := "hello,world"            //  hello,world，由于hello world之前没有$符号，则无法利用map规则进行转换
	s1 := "$hello,$world $finish" //  hi,kitty finish，finish没有在map规则中，所以还是返回原来的值
	fmt.Println(os.Expand(s, mapping))
	fmt.Println(os.Expand(s1, mapping))

	//s2 := "hello $GOROOT11"
	s2 := "hello $GOROOT"
	fmt.Println(os.ExpandEnv(s2))//$GOROOT替换为环境变量的值，而h没有环境变量可以替换，返回空字符串



}