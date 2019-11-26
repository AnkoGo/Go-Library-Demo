//例子1-------------------------------------------------------------
package main

import "fmt"

func demo(recoverShallowerPanicAtFirst bool) {
	fmt.Println("====================")
	defer func() {
		if !recoverShallowerPanicAtFirst{
			// 恢复恐慌1
			defer fmt.Println("恐慌", recover(), "被恢复了")
		}
		defer func() {
			//  恢复恐慌2
			fmt.Println("恐慌", recover(), "被恢复了")
		}()
		if recoverShallowerPanicAtFirst {
			//  恢复恐慌1
			defer fmt.Println("恐慌", recover(), "被恢复了")
		}
		defer fmt.Println("现在有两个恐慌共存")
		panic(2)
	}()
	panic(1)
}

func main34876() {
	demo(true)
	demo(false)
}



//例子2-------------------------------------------------------------
//package main
//
//import "fmt"
//
//func demo() {
//	defer func() {
//		// recover panic 1
//		defer fmt.Println(" (done).")
//		defer recover()
//		defer fmt.Println("To recover panic 1 ...")
//
//		defer func() {
//			// recover panic 2
//			fmt.Println("panic", recover(), "is recovered")
//		}()
//
//		defer fmt.Println("now, two active panics coexist")
//		panic(2)
//	}()
//	panic(1)
//}
//
//func main() {
//	demo()
//}




















