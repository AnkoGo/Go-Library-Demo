package main

import (
	//"bytes"
	"fmt"
	"github.com/axgle/mahonia"
	"log"
	"os"
	"os/exec"
	//"io"
)

//src为要转换的字符串，srcCode为待转换的编码格式，targetCode为要转换的编码格式
func ConvertToByte(src string, srcCode string, targetCode string) []byte {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(targetCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	return cdata
}

func main667778() {
	for key, value := range os.Environ() {
		fmt.Println(key, ":",value)
	}
	// LookPath在由PATH环境变量命名的目录中搜索一个名为可执行文件的文件。
	//如果文件包含斜杠，则直接尝试使用该斜杠，并且不查询PATH。
	// LookPath还使用PATHEXT环境变量来匹配合适的候选者。
	//结果可能是绝对路径，也可能是相对于当前目录的路径。
	//path, err := exec.LookPath("go.exe")
	path, err := exec.LookPath(`C:\Go\bin\go.exe`)//也可以给一个完整的路径
	if err != nil {
		log.Fatal("在当前系统的环境变量中无法找到相应的文件")
	}
	fmt.Println(path)//C:\Go\bin\go.exe
	fmt.Printf("找到了，路径是：%s\n", path)

	fmt.Println("-------------------")
	cmd := exec.Command(`ping`,"www.baidu.com" )
	//cmd.Stdin = strings.NewReader("some input")//这里自己创建一个*reader,因为不是所有的命令都有输入的！


	fmt.Println(cmd.Stdin)
	//设置输入流
	writeCloser, err222 := cmd.StdinPipe()
	if err222 !=nil{
		fmt.Println(err222)
	}
	slice := make([]byte, 20)
	slice=append(slice,'b','a')
	n, err333 := writeCloser.Write(slice)//一定要在命令开始之前设置
	if err333 !=nil{
		fmt.Println(err333)
	}
	fmt.Println(n)
	by_slice:=make([]byte,22)
	fmt.Println(cmd.Stdin.Read(by_slice))
	fmt.Println(by_slice)

	fmt.Println("======***********========")

	//out, err777 := cmd.Output()//返回一个接受输出的[]byte结构，设置了这个的话就不用设置cmd.run()了，因为在这个方法里面写了
	out, err777 := cmd.CombinedOutput()//执行命令并返回标准输出和错误输出合并的切片。
	if err777 != nil {
		log.Fatal(err777)
	}
	//cmd.Start()
	//fmt.Println(cmd.SysProcAttr)//不是所有系统都会设置这个东西
	//fmt.Println(cmd.SysProcAttr.CmdLine)

	fmt.Println("cmd.ProcessState",cmd.ProcessState.String())//exit status 0
	fmt.Println("cmd.ProcessState",cmd.ProcessState.Pid())//32
	fmt.Println("cmd.ProcessState",cmd.ProcessState.Exited())//true
	fmt.Println("cmd.Process",cmd.Process)//包含一个已经存在的进程的信息，只有在调用Wait或Run后才可用,如&{8576 18446744073709551615 1 {{0 0} 0 0 0 0}}
	fmt.Println("cmd.Env",cmd.Path)//C:\windows\system32\ping.exe
	fmt.Println("cmd.Env",cmd.Env)//如为空切片，则是在当前进程的环境下执行。
	fmt.Println("cmd.Args",cmd.Args)//[ping www.baidu.com]
	fmt.Println("cmd.Dir",cmd.Dir)//空字符串表示调用者的进程当前目录下执行
	fmt.Printf("The date is %s\n", out)


	//设置错误的输出管道
	//stderrPipe, err555 := cmd.StderrPipe()//这个要在start之前设置
	//if err555 != nil {
	//	log.Fatal(err555)
	//}

	//设置输出流
	//var out bytes.Buffer
	//cmd.Stdout = &out
	//多种方式写入输出流
	//readCloser, err444 := cmd.StdoutPipe()
	//if err444 !=nil{
	//	fmt.Println(err444)
	//}
	//fmt.Println(readCloser)
	//
	//
	////err111 := cmd.Run()//这个阻塞然后接受数据最后关闭所有管道了
	//err111 := cmd.Start()//这个阻塞然后接受数据但是不关闭管道了
	//if err111 != nil {
	//	log.Fatal(err111)
	//}
	//
	//s2:=make([]byte,50)
	//stderrPipe.Read(s2)//这个只有读取错误才会输出错误
	//fmt.Println("stderrPipe:",s2)
	//
	//sl:=make([]byte,30)
	//fmt.Println(readCloser.Read(sl))






	cmd.Wait()//必须在所有管道读取之后才wait,他主要是关闭管道,释放相关的资源的作用
	//fmt.Println(sl)
	//fmt.Println(string(sl))


	////ls:=make([]byte,200)
	////utf8.EncodeRune(out.Bytes())
	////for len(out.Bytes()) > 0 {
	////	r, size := utf8.DecodeRune(out.Bytes())
	////	fmt.Printf("%c %v", r, size)
	////	b = b[size:]
	////}
	//////lrune:=bytes.Runes(ls)
	////fmt.Println(string(lrune))
	//
	////b := out.Bytes()
	////for len(out.Bytes()) > 0 {
	////	r, size := utf8.DecodeRune(out.Bytes())
	////	fmt.Printf("%c %v", r, size)
	////	b = b[size:]
	////}
	//fmt.Println("--------------------")
	//b := out.Bytes()
	//
	//
	//
	//fmt.Print()
	////fmt.Print(out.String())
	//fmt.Print(out)
	//fmt.Println("---2222-----------------")
	//for len(b) > 0 {
	//	r, size := utf8.DecodeRune(b)
	//	//fmt.Print(r,",")
	//	fmt.Printf("%v",string(r))
	//	b = b[size:]
	//}
	//
	////对string由gbk转码为utf8
	//fmt.Println("******************")
	//response := ConvertToByte(out.String(), "gbk", "utf8")
	//fmt.Println(response)
	//fmt.Println(string(response))
	//
	////写入到文件
	//f,err:=os.OpenFile("test4.txt",os.O_APPEND,0777)
	//if err !=nil{
	//	fmt.Println(err)
	//}
	//f.Write(response)
	//f.Close()
}







