package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)
//子协程
func consumer(stop <-chan bool) {
	for {
		select {
		//不停的检测来自系统的信号的管道的是否有值，一旦有值就停止子协程，stop是系统信号有无标志的管道
		case <-stop:
			fmt.Println("exit sub goroutine")
			return
		default:
			fmt.Println("running...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
func main() {
	stop := make(chan bool)
	var wg sync.WaitGroup
	//开启三个协程
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(stop <-chan bool) {
			defer wg.Done()
			consumer(stop)//协程主函数
		}(stop)
	}
	waitForSignal()//阻塞
	//一旦解开阻塞就说明有来自系统的信号了，接着关闭这个管道就相当于往这个管道上面发送2个false值
	//第一个false值是管道元素的零值，第二个值是管道的标志位值，当关闭管道的时候会将此标志设置为
	//false并且将这个值发送到管道里面去。
	close(stop)
	fmt.Println("stopping all jobs!")
	wg.Wait()//之所以 要上锁是因为多个进程不同步操作这段代码会导致管道里面的值来自多个进程！
}
func waitForSignal() {
	//创建信号管道，没指定缓存代表监听所有类型的信号，一般指定1即可
	sigs := make(chan os.Signal)
	signal.Notify(sigs, os.Interrupt)
	signal.Notify(sigs, syscall.SIGTERM)
	<-sigs//阻塞监听是否有来自系统的信号出现，一旦出现就解开阻塞
}
