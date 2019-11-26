package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)
type MynopCloser struct {
	io.Writer
}

func (w MynopCloser) Close() error {
	return nil
}

func main8606() {
	// Compressor返回一个新的压缩编写器，写入w。
	//必须使用WriteCloser的Close方法将待处理数据刷新到w。
	// Compressor本身必须可以安全地同时从多个goroutine调用，但是每个返回的writer一次只能由一个goroutine使用。


	//首先创建一个文件
	_, e := os.OpenFile("main/zip/zip1.zip", os.O_CREATE, 777)//打开新建的压缩zip文件准备读写
	checkErr_zip(e)


	file_dst, e := os.OpenFile("main/zip/zip1.zip", os.O_RDWR, 777)//打开新建的压缩zip文件准备读写
	checkErr_zip(e)
	var zipCom zip.Compressor//zipCom是一个接收装载器的压缩写入器函数
	zipCom=func(w io.Writer) (io.WriteCloser, error) { return &MynopCloser{w}, nil }
	wtCloser, e := zipCom(file_dst)
	checkErr_zip(e)

	checkErr_zip(e)
	//打开要压缩的文件
	file_src, e := os.OpenFile("main/tar_data/text1.txt", os.O_RDWR, 777)//打开文件准备读写
	checkErr_zip(e)
	info_src, e := file_src.Stat()
	checkErr_zip(e)
	byte_src:=make([]byte,info_src.Size())
	n, e := file_src.Read(byte_src)
	checkErr_zip(e)
	fmt.Println("要压缩的文件的字符串是：",string(byte_src))
	fmt.Println("要压缩的文件的字节数是：",n)
	//开始将text1中的字节写入压缩文件中去
	i, e := wtCloser.Write(byte_src)
	checkErr_zip(e)

	fmt.Println("成功写入压缩文件的字节数是（只有这个值等于上面的要压缩的文件的字节数的话才是成功写入了）：",i)

	defer func() {
		e2 := wtCloser.Close()
		checkErr_zip(e2)
		e2 = file_src.Close()
		checkErr_zip(e2)
		e2 = file_dst.Close()
		checkErr_zip(e2)
	}()

}
func checkErr_zip(err error)  {
	if err != nil{
		fmt.Println(err)
	}
}