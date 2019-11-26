package main

import (
	"fmt"
	"unicode/utf8"
)

func main2356() {
	str:="\u8bf7\u6c42\u7684\u6570\u636e"//unicode码值的16进制表示
	fmt.Println(str)
	fmt.Println(utf8.ValidString(str))
}
