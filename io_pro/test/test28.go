package main

import (
	"fmt"
)

type T_ls interface {
	Name() *P
}

type P struct {
}

func (*P) Name() *P {
	fmt.Println("name....")
	return nil
}

var p P

func ret_ls() T_ls {
	fmt.Println("ret_ls....")
	return p.Name()
}

func main3472336() {
	ls := ret_ls()

	fmt.Printf("类型：%T,值：%v\n", ls, ls)

	
	if ls != (*P)(nil) {
		fmt.Println("不是nil")
	}
	if ls == (*P)(nil) {
		fmt.Println("是nil")
	}

	// var str []byte
	// str=nil

	// if str != nil {
	// 	fmt.Println("str不是nil 000")
	// }

	// if str == nil {
	// 	fmt.Println("str是nil 00011")
	// }

	// if reflect.TypeOf(str) != nil {
	// 	fmt.Println("str不是nil 111")
	// }
	// if reflect.ValueOf(str).IsNil(){
	// 	fmt.Println("str是nil 222")
	// }

	// if reflect.TypeOf(nil) == nil {
	// 	fmt.Println("是nil")
	// }


}
