package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

func main346() {
	testInput := "<P>Foo<P>\n\n<P>Bar</>\n"
	d := xml.NewDecoder(strings.NewReader(testInput))
	T1, err1111 := d.Token()
	check_EncDec_err(err1111)
	T2, err2222 := d.Token()
	check_EncDec_err(err2222)

	T3, err333 := d.Token()
	check_EncDec_err(err333)

	T4, err444 := d.Token()
	check_EncDec_err(err444)

	fmt.Println(T1)
	fmt.Println(T2)
	fmt.Println(T3)
	fmt.Println(T4)
}

func check_EncDec_err(err error) {
	if err != nil{
		fmt.Println("----",err)
	}
}