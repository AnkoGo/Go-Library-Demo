package main

import (
	"fmt"
	"unsafe"
)

func changeValue(a unsafe.Pointer){

	fmt.Printf("changeValue_a:\t%T--%p--%v\n",a,a,a)
	i:=(*int)(a)
	fmt.Println("changeValue_i:\t",i)
	*i=5
}

func main346773() {
	var cls int=4
	var a  =unsafe.Pointer(&cls)

	fmt.Printf("main_a:\t%T=====%p======%v\n",a,a,a)
	fmt.Printf("changeValue前cls:\t%T_____%p_____%v\n",cls,&cls,cls)
	changeValue(a)
	fmt.Printf("changeValue后cls:\t%T_____%p_____%v\n",cls,&cls,cls)

}
