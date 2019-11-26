package main

import (
	"fmt"
)

func main1135(){
	x := "hello"
	for _, y := range x {
		x=fmt.Sprintf("%s %s",x,string(y))//连接多个字符串不会加配内存
	}
	fmt.Println(x)
}