package main

import "fmt"

type person1 struct {
	name string
	age string
	info []string
}

func (p1 *person1)String()string  {
	var str string="==" + p1.name + p1.age
	return str
}
//不可以既实现指针类型的String()方法，又实现值类型的String()方法
//func (p1 person1)String()string  {
//	var str string="==" + p1.name + p1.age
//	return str
//}

type person2 struct {
	name string
	age string
	pp1 *person1
}


func main34677() {

	var p1 = &person1{name:"anko1--",age:"17",info:[]string{"info1","info2"}}
	var p2 = &person2{name:"anko2--",age:"18",pp1:p1}
	pp1 := p2.pp1
	fmt.Printf("%+v\n",*pp1)
	fmt.Printf("%+v\n",pp1)
	//{name:anko1-- age:17 info:[info1 info2]}
	//==anko1--17


}
