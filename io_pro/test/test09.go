package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main23468() {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	bw.WriteString("123")
	fmt.Println("--",b)
	c := bytes.NewBuffer(make([]byte, 0))
	bw.Reset(c)
	bw.WriteString("456")
	bw.Flush()
	fmt.Println("--",b)
	fmt.Println(c)
	//输出：
	//--
	//--
	//456
}