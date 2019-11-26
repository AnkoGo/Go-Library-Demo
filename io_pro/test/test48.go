package main

import (
	"fmt"
	"reflect"
)

type Reader interface {
	MyRead(p []byte) (n int, err error)
}

type teach interface {
	//Reader	//teach接口和Reader接口相互独立
	MyWrite(p []byte) (n int, err error)
}

type teacher struct {
	name string
	age int
}
//实现了独立接口Reader
func (teacher)MyRead(p []byte) (n int, err error)  {
	fmt.Println("MyRead()方法执行中")
	return len(p), nil
}
//实现了独立接口teach
func (teacher)MyWrite(p []byte) (n int, err error)  {
	fmt.Println("MyWrite()方法执行中")
	return len(p), nil
}

func main088239() {
	T:=teacher{
		name: "anko",
		age:  20,
	}
	fmt.Println(T.name)
	fmt.Println(T.age)
	fmt.Println(T.MyRead([]byte{'a','b','c'}))
	fmt.Println(T.MyWrite([]byte{'x','y','z'}))

	fmt.Println("--------------")
	var iT teach
	if !reflect.ValueOf(iT).IsValid(){
		fmt.Println("iT此时为nil,不能调用任何方法，否则会报错，为此，我们将iT赋值给一个实现类实例对象：")
		iT=T
		//fmt.Println(iT.MyRead([]byte{'a','b','c'}))//不存在MyRead方法
		fmt.Println(iT.(Reader).MyRead([]byte{'a','b','c'}))
		fmt.Println(iT.MyWrite([]byte{'x','y','z'}))
	}else {
		fmt.Println("iT非nil")
		//fmt.Println(iT.MyRead([]byte{'a','b','c'}))//不存在MyRead方法
		fmt.Println(iT.(Reader).MyRead([]byte{'a','b','c'}))
		fmt.Println(iT.MyWrite([]byte{'x','y','z'}))
	}
}