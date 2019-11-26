package main

import (
	"fmt"
)

func main0897(){
	x := "hello"
	y:=&x

	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%p---%v\n",&x,x)
	fmt.Printf("%p---%v\n",y,*y)

	*y="HELLO"
	fmt.Printf("%p---%v\n",&x,x)
	fmt.Printf("%p---%v\n",y,*y)
	fmt.Println(*y)


	z:=x[1:2]
	fmt.Printf("%p---%v\n",&z,z)
	q:=&z
	fmt.Printf("%p---%v\n",q,*q)
	*q="e"
	fmt.Printf("%p---%v\n",&z,z)
	fmt.Printf("%p---%v\n",q,*q)

	fmt.Printf("%p---%v\n",&x,x)
	fmt.Printf("%p---%v\n",y,*y)

	//输出如下：
	//	hello
	//	0xc0000301f0
	//	0xc0000301f0---hello
	//	0xc0000301f0---hello
	//	0xc0000301f0---HELLO
	//	0xc0000301f0---HELLO
	//	HELLO
	//	0xc000030260---E
	//	0xc000030260---E
	//	0xc000030260---e
	//	0xc000030260---e
	//	0xc0000301f0---HELLO
	//	0xc0000301f0---HELLO

	//从上面可以知道确实不可以更改字符串的单个元素
}