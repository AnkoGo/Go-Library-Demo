package main

import("fmt";"time")


func main45782() {
	for i:=0;i<3;i++{

		var a int
		fmt.Println("请输入一个值：")
		n,err:=fmt.Scanln(&a)
		if err!=nil{
			fmt.Println(err)
		}
		fmt.Println("输出：",a,"\t字符串个数为：",n)
		time.Sleep(2e9)

	}
	time.Sleep(5e9)
}