package main

import "fmt"

var f1 = func(i int) {
	fmt.Println("X")
}

func main234526() {
	//f1(0)
	//var f1 func(i int)
	f1:=func (i int)  {
		fmt.Println(i)
		if i>0{
			f1(i-1)
		}
	}
	f1(10)

}
