/* ASN.1
 */

package main

import (
	"encoding/asn1"
	"fmt"
	"os"
)
type person struct {
	Name string
	Age int
	Isfat bool
}


func main23845() {
	P:=person{
		Name:  "anko",
		Age:   18,
		Isfat: false,
	}
	mdata, err := asn1.Marshal(P)
	checkError(err)
	fmt.Println(mdata,"字符串为：",string(mdata))

	var n person
	ls, err1 := asn1.Unmarshal(mdata, &n)
	checkError(err1)
	fmt.Println(ls)
	fmt.Println("After marshal/unmarshal: ", n)

	fmt.Println("-------------------------------------")

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}