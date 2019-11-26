package main

import (
	"fmt"
	"time"
)

func main34792() {
	chan_int3 := make(chan int, 3)
	chan_int4 := make(chan int, 4)
	go func() {
		n:=0
		for n<20000000{
			chan_int3<-n
			n++
		}
	}()
	time.Sleep(2e9)//确保上面的g程存值

	go func() {
		time.Sleep(2e9)
		n:=0
		for n<200{
			chan_int4<-n
			n++
		}
	}()

	for {

		select {
		case i:=<-chan_int4:
			fmt.Println("取到chan_int4通道的值为==============：",i)
		case i:=<-chan_int3:
			fmt.Println("取到chan_int3通道的值为：",i)
			for m:=0;m<200000 ;m++  {
				m++
			}
		}
	}
	fmt.Println("for循环已经退出。。。准备结束整个程序")
}
