package main

import "fmt"

//至少传递0个参数
func add1(ls_int ...int) int {
	num:=0
	for _, v := range ls_int {
		num+=v
	}
	return num
}

//至少传递一个参数
func add2(default1 int,ls_int ...int) int {
	num:=default1
	for _, v := range ls_int {
		num+=v
	}
	return num
}

func main457993() {
	fmt.Println(add1())
	fmt.Println(add1(3,5))
	fmt.Println(add2(3))
	fmt.Println(add2(3,7))
	fmt.Println(add2(3,7,5))
}
