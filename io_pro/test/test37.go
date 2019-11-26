package main

import "fmt"

type add struct {
	x int
}

func (a add) add1(y ...int) int {
	return a.x + y[0]
}

func main92738() {
	var ad = add{}
	sum := ad.add1(5, 6)
	fmt.Println(sum)
}
