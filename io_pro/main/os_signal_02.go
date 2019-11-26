package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main90124() {
	fmt.Println("111111111")
	//c := make(chan os.Signal, 3)
	//signal.Notify(c, os.Interrupt)
	// Block until a signal is received.


	ch := make(chan os.Signal, 1)
	//监控按键以及其他程序或者子协程的运行，用于监控或者重启程序以及程序配置的重新加载
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM,os.Interrupt)
	for {
		s := <-ch
		switch s {
		case syscall.SIGQUIT:
			log.Fatalln("SIGSTOP")
			return
		case syscall.SIGHUP:
			log.Fatalln("SIGHUP")
			return
		case syscall.SIGKILL:
			log.Fatalln("SIGKILL")
			return
		case os.Interrupt:
			log.Fatalln("os.Interrupt")
			return
		default:
			log.Fatalln("default")
			return
		}
	}


	//fmt.Println("222222222222")
	//s := <-c
	//fmt.Println("333333333333")
	//fmt.Println("Got signal:", s)
	//fmt.Println("344444444444444")

}