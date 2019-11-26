package main

import "fmt"

func main67942() {
	ls := []int{1, 2, 3, 4, 5}
	ch := make(chan int, 1)

	for _, i := range ls {
	loop:
		select {
		//取值
		case x := <-ch:
			fmt.Println(x)

			if x !=5{
				goto loop
			}
		//存值
		case ch <- i:
			fmt.Println("存值：", i)

			if i ==5 {
				goto loop
			}

		}
	}

}
