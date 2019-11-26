package main

import (
	"fmt"
	"reflect"
)

func main457882() {

	type ppp struct {
		name string
		age int
	}
	var p ppp
	//=ppp{"anko",88}
	fmt.Printf("%#v\n",p)
	if reflect.ValueOf(p).IsValid(){
		fmt.Println("有效值，非zero Value")
	}else {
		fmt.Println("zero Value,无效值")
	}
	//只要是reflect.ValueOf（）返回的值都不会是reflect.Value{}这个零值
	fmt.Printf("%#v\n",reflect.ValueOf(reflect.ValueOf(&p).Elem()).IsZero())
	fmt.Printf("%#v\n",reflect.ValueOf(reflect.ValueOf(p)).IsZero())//这个有没有疑问的？
	fmt.Printf("%#v\n",reflect.ValueOf(reflect.Value{}).IsZero())//这个有没有疑问的？
	fmt.Printf("%#v\n",reflect.ValueOf(p).IsZero())
	fmt.Printf("%#v\n",reflect.ValueOf(&p).IsZero())

	fmt.Println()
	var p1 *ppp
	fmt.Printf("%#v\n",reflect.ValueOf(reflect.ValueOf(&p1).Elem()).IsZero())
	fmt.Printf("%#v\n",reflect.ValueOf(reflect.ValueOf(p1).Elem()).IsZero())
	fmt.Printf("%#v\n",reflect.ValueOf(reflect.ValueOf(p1)).IsZero())
	fmt.Printf("%#v\n",reflect.ValueOf(p1).IsZero())
	fmt.Printf("%#v\n",reflect.ValueOf(&p1).IsZero())


	fmt.Println()
	var p2 *ppp
	//panic: reflect: call of reflect.Value.IsNil on struct Value(reflect.Value结构体----字段指针指向*ppp)
	//fmt.Printf("%#v\n",reflect.ValueOf(reflect.ValueOf(&p2).Elem()).IsNil())
	//panic: reflect: call of reflect.Value.IsNil on struct Value(reflect.Value结构体----字段指针指向ppp)
	//fmt.Printf("%#v\n",reflect.ValueOf(reflect.ValueOf(p2).Elem()).IsNil())
	//panic: reflect: call of reflect.Value.IsNil on struct Value(reflect.Value结构体----字段指针指向*ppp)
	//fmt.Printf("%#v\n",reflect.ValueOf(reflect.ValueOf(p2)).IsNil())
	fmt.Printf("%#v\n",reflect.ValueOf(p2).IsNil())
	fmt.Printf("%#v\n",reflect.ValueOf(&p2).IsNil())

	//输出：
	//	main.ppp{name:"", age:0}
	//	有效值，非zero Value
	//	false
	//	false
	//	true
	//	false
	//
	//	false
	//	true
	//	false
	//	true
	//	false
	//
	//	true
	//	false
}
