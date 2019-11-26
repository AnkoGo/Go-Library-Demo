package main

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
)

func update( aa11 *interface{}){
	*aa11=5
}

func main34356() {
	var aa11 interface{}=4
	fmt.Printf("aa11更改之前:%T---%v\n",aa11,aa11)
	update(&aa11)
	fmt.Printf("aa11更改之后:%T---%v\n",aa11,aa11)

	var r io.Reader
	r=bytes.NewBuffer([]byte{'a','b','c'})
	fmt.Printf("%#v\n",reflect.TypeOf(r).Elem())
	//输出：
	//	aa11更改之前:int---4
	//	aa11更改之后:int---5
	//	&reflect.rtype{size:0x28, ptrdata:0x8, hash:0x64cbaa3a, tflag:0x7, align:0x8, fieldAlign:0x8, kind:0x19, alg:(*reflect.typeAlg)(0x620620), gcdata:(*uint8)(0x5499b0), str:33168, ptrToThis:253920}
	a:=3
	b:=3
	fmt.Printf("%#v\n",reflect.DeepEqual(&a,&b))
}
