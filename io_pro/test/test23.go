package main

import (
	"fmt"
	"os"
)

func main3465() {
	f := `D:\\con`//不能使用con，是bug
	file, err := os.Open(f)

	if err != nil {
		fmt.Println("打开文件发生了错误，错误信息为：", err)
	}
	fmt.Printf("file的类型是：%T,值是：%v\n", file, file)

	dst_byte := make([]byte, 5)
	n, err := file.Read(dst_byte)
	if err != nil {
		fmt.Println("读取发生了错误，错误信息为：", err)
	}
	fmt.Println("读取到的字节个数为：", n)
	fmt.Println("读取到的字节的装载切片为：", dst_byte)
}
