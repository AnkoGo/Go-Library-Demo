package main

import (
	"fmt"
	"reflect"
	"time"
)

func main346787() {
	chan_int3 := make(chan int, 3)
	go func() {
		n:=0
		for n<20000000{
			chan_int3<-n
			n++
		}
	}()
	//time.Sleep(2e9)//确保上面的g程存值
loop:
	//for {
		T_After:=time.After(3e9)
		v,ok:=reflect.ValueOf(T_After).TryRecv()
		fmt.Println(v,ok)

		if v.Kind()!=reflect.Invalid{
			fmt.Println("3s时间到，准备执行退出循环",v)
			//break loop
			return
		}
		//select {
		//case i:=<-chan_int3:
		//	fmt.Println("取到通道的值为：",i)
			for m:=0;m<200000000 ;m++  {
				m++
			}
			goto loop
		//}
	//}
	fmt.Println("for循环已经退出。。。准备结束整个程序")
}
