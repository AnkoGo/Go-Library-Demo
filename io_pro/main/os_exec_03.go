package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)



func main() {

	fmt.Println("-------------------")
	cmd := exec.Command(`ping`,"www.baidu.com" )

	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	//simplifiedchinese.GB18030
	fmt.Println(out.String())

	fmt.Println("======***********========")



	cmd.Wait()//必须在所有管道读取之后才wait,他主要是关闭管道,释放相关的资源的作用

}







