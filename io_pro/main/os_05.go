package main

import (
	"fmt"
	"os"
	"time"
)

func main56790() {
	attr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}, //其他变量如果不清楚可以不设定
	}
	//notepad.exe 打开main/tmp.txt文件
	p, err := os.StartProcess(`C:\Windows\System32\notepad.exe`, []string{`C:\Windows\System32\notepad.exe`,
		`C:\Users\Administrator\Desktop\go_pro\src\io_pro\main\test.txt`}, attr)//阻塞的，直到文件被关闭

	fmt.Println("1111111111")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)                  //&{3308 288 0 {{0 0} 0 0 0 0}}
	pro, _ := os.FindProcess(p.Pid) //查找进程
	fmt.Println(pro)                //&{3308 244 0 {{0 0} 0 0 0 0}}

	fmt.Println("------------------------")
	fmt.Println("333")
	time.Sleep(4)
	go func() {
		fmt.Println("444")
		time.Sleep(6e9)//nano
		//发送一个信号给进程p, 在windows中没有实现发送中断interrupt
		p.Signal(os.Kill) //kill process
	}()
	fmt.Println("222")
	pstat, err := p.Wait()
	if err != nil {
		fmt.Println("111",err)
	}

	fmt.Println(pstat) //exit status 1,如果是linux的话会显示signal: killed


	// StartProcess使用名称，argv和attr指定的程序，参数和属性启动新进程。 argv slice在新进程中将变为os.Args，因此它通常以程序名称开头。
	//如果调用的goroutine已使用runtime.LockOSThread锁定了操作系统线程，并修改了任何可继承的OS级线程状态（例如Linux或Plan 9命名空间），则新进程将继承调用者的线程状态。
	// StartProcess是一个低级接口。 os / exec软件包提供了更高级别的接口。
	//如果有错误，它将是* PathError类型。


}