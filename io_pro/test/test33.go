package main

import (
	"fmt"
	"strconv"
	"time"
)

func main568343() {
	ticker := time.NewTicker(2e9)
	defer ticker.Stop()
	ch1 := make(chan string, 3)
	ch2 := make(chan string, 3)

	go func() {
		for i := 0; i < 2000000; i++ {
			ch1 <- "ch1---" + strconv.Itoa(i)
		}
	}()

	go func() {
		for i := 0; i < 3000000; i++ {
			ch1 <- "ch2---" + strconv.Itoa(i)
		}
	}()
	for {
		select {
		case u := <-ch1:
			fmt.Println(u)
		case v := <-ch2:
			fmt.Println(v)
		case <-ticker.C:
			fmt.Println("2s时间到，进入 ticker case")

		default:
			fmt.Println("default case")
			break
		}
	}


}
