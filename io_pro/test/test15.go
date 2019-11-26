package main

import (
	"fmt"
	"time"
)

func producer8(out chan int) {
	go generate(out)
}

func generate(out chan int) {
	fmt.Println("hello,i am producer")
	for i := 0; i < 10; i++ {
		fmt.Printf("producer one %v \n", i+1)
		out <- i

	}
	//close(out)
}

func consumer8(in <- chan int) {
Loop:
	for {
		select {
		case <-time.After(2e9):
			fmt.Printf("time out\n")
			break Loop
		case v1, ok := <-in:
			if ok {
				fmt.Printf("recive one %v %v \n", v1, ok)
				time.Sleep(3e9)
			}
			if !ok {
				fmt.Println("chan无知可取。。。。")

			}

		}
	}
}

func testCurrencyPatternsGenerators() {

	out := make(chan int)
	producer8(out)
	go consumer8(out)
	//for i:=0;i<100 ;i++  {
	//	fmt.Println(i)
	//}
	time.Sleep(30e9)
	fmt.Println("程序结束")
}
func main4592() {

	testCurrencyPatternsGenerators()
}
