package main

import "fmt"

func f() {
	f()
}

func main87890() {
	defer func() {
		fmt.Println("出现异常===========================")
		recover() // 对于栈溢出，无法防止程序崩溃，也无法捕获异常或者恢复它
	}()
	f()

}
