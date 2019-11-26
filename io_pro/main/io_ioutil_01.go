package main

import (
	"fmt"
	"io"
	"io/ioutil"
)
//创建一个io.Reader对象
type myReader struct {
	s string//字符串
	i int64//当前的字节索引
	prevRune int   // index of previous rune; or < 0

}

func newmyReader(s string) *myReader {
	return &myReader{
		s: s,//字符串，表示可以通过字符串创建
		i: 0,//读索引
		prevRune: -1,//Rune的索引
	}
}

func (r *myReader) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		r.i=0//置零指针
		return 0, io.EOF//返回io.EOF告诉调用方已经读到结尾了
	}
	r.prevRune = -1
	n = copy(b, r.s[r.i:])
	r.i += int64(n)

	return//读取部分，返回部分字节，但不是读到末尾了
}
func(r *myReader) resetIndex()  {
	ioutil.ReadAll(r)
}

func main345454() {
	mR:=newmyReader("abcdef124")
	ls:=make([]byte,13)
	n,err111:=mR.Read(ls)
	if err111 != nil {
		fmt.Println(err111)
	}
	fmt.Println("此次读取的ls:",ls,"长度为：",n)

	mR.resetIndex()
	bytes, e := ioutil.ReadAll(mR)
	if e != nil{
		fmt.Println("==================",e)
	}
	fmt.Println("接着上次的读取停止位置一次性读取剩余的全部bytes:",bytes)
	fmt.Println("len(bytes)",len(bytes),"cap(bytes)",cap(bytes))
}

