package main

import (
	"fmt"
	"os"
)

func main346289() {
	file, e := os.Open("test/config.xml")
	check_EncDec_err(e)
	info, e := file.Stat()
	check_EncDec_err(e)

	dst_byte := make([]byte, info.Size())
	n, e := file.Read(dst_byte)
	check_EncDec_err(e)

	fmt.Println("读取到的文件的字节数是：", n)
	fmt.Println("读取到的文件字节是：", dst_byte)
	fmt.Printf("读取到的文件字符串是：\n%v\n", string(dst_byte))

	ret, e := file.Seek(0, 0)
	check_EncDec_err(e)

	fmt.Println("将文件读写指针移动到：", ret)

	dst_byte1 := make([]byte, info.Size())
	n1, e := file.Read(dst_byte1)
	check_EncDec_err(e)

	fmt.Println("读取到的文件的字节数是：", n1)
	fmt.Println("读取到的文件字节是：", dst_byte)
	fmt.Printf("读取到的文件字符串是：\n%v\n", string(dst_byte))
}
