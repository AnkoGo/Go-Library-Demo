package main

import (
	"fmt"
	"time"
)

func main347834() {
	start()
	fmt.Println("start end")
	time.Sleep(8e9)
	fmt.Println("main end")
}

func start() {
	fmt.Println("here")
	go func() {
		time.Sleep(2e9)
		fmt.Println("goroutine running")
	}()

	//tick := time.Tick(1e8)
	//boom := time.After(5e8)
	//select {
	//case <-tick:
	//	fmt.Println("tick.")
	//case <-boom:
	//	fmt.Println("BOOM!")
	//	return
	//default:
	//	fmt.Println("." )
	//	time.Sleep(5e7)//2个5e7=1e8
	//}
}
