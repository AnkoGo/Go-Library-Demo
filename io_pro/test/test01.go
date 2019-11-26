package main

import (
	"fmt"
)

func main111() {
	x := make([]int, 1, 10)

	var a []int
	for i := 0; i < 10; i++ {

		a = append(x, i)
		fmt.Printf("%p---%p---%v----%v---%v   ===   %p---%p---%v----%v---%v\n", a, &a[0], len(a), cap(a), a, x, &x[0], len(x), cap(x), x)
	}

	fmt.Println("--------------------------------------------------")
	fmt.Printf("%p---%p---%v----%v---%v   ===   %p---%p---%v----%v---%v\n", a, &a[0], len(a), cap(a), a, x, &x[0], len(x), cap(x), x)
	fmt.Println("a:", a)
	fmt.Println("x:", x)
	fmt.Println("sdsdsd")
	
}
