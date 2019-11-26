package main

import (
	"fmt"
	"reflect"
)

//func main() {
//	Myperson(map[string]interface{}{"name":"anko","age":18})
//}
//
//func Myperson(b map[string]interface{})  {
//	fmt.Println(b)
//}

//这个结构体作为 关键字传参封装函数
type keyArgs struct{
	name interface{}
	age interface{}
}
//对结构体进行默认值的检查
func keyArgsDefaultValue(defaultValue ...interface{}) *keyArgs{
	k,ok:= defaultValue[0].(*keyArgs)
	//defaultValue2 := defaultValue
	//断言失败，默认值
	if !ok{
		k1:=&keyArgs{defaultValue[0],defaultValue[1]}
		return k1
	}else {
		//设置值
		for key, value0 := range defaultValue {
			switch reflect.ValueOf(*k).Field(key).Kind(){

			case reflect.String:
				k.name=value0.(string)
			case reflect.Int:
				k.age=value0.(int)

			}
		}
		return k
	}



}

func defaultmy(defaultFun func(defaultValue ...interface{}) *keyArgs,k1 ...*keyArgs,) (kArg *keyArgs) {
	if len(k1)==0{
		//默认值
		kArg =defaultFun("anko1",18)
	}else {
		//设置值
		kArg=defaultFun(k1[0])
	}
	return
}

func Myperson(defaultFun func(defaultValue ...interface{}) *keyArgs,k1 ...*keyArgs) {
	var kArg *keyArgs=defaultmy(defaultFun,k1...)
	fmt.Printf("%T---%#v\n",kArg,kArg)
}


func main() {

	//Myperson(keyArgsDefaultValue)
	Myperson(keyArgsDefaultValue,&keyArgs{"anko",20})



}



