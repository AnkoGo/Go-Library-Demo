package main

import (
	"fmt"
	"io"
	)
func main3267788() {
	byte_slice:=[]byte{0,1,2,3,4,5,6}

	var R =io.SectionReader{}
	src,err:=R.Read(byte_slice)

	fmt.Println(src,err)
}

