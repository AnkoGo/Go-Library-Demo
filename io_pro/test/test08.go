package main

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"log"
)

func main34343() {
	//编码unicode成为GB18030码
	enc := simplifiedchinese.GB18030.NewEncoder() //申明一个编码器（相对于unicode来说确实是编码）
	encBuf := bytes.NewBuffer(make([]byte, 0))    //申明一个encBuf来存储转换后的原生字节
	// 序列（无任何编码的遵循unicode协议的字节序列，他是unicode值，但是unicode不是一种编码，而是一种规范和协议）

	// Writer包装另一个Writer对其UTF-8输出进行编码。
	//只要使用了返回的Writer，编码器就不能用于任何其他操作。
	writ := enc.Writer(encBuf) //申明一个对接encBuf缓存的写入器，这个写入器是由编码器生成的，
	// 因此它具有编码的功能， 在这儿是把GB18030转成原生字节的转码器
	writ.Write([]byte{34, 11, 8, 146, 147, 214, 236, 5, 16, 152, 148, 189, 4, 42, 0, 50, 10, 8, 12, 18, 6, 49, 50, 51, 52, 53, 54}) //把一个原生字节序列或者叫unicode字符串经过
	// 写入器（内含GB18030编码器）写入到encBuf缓存容器中去，此时缓存容器里面存的是GB18030编码后的
	//字节序列
	fmt.Println(encBuf)                    //这个是buf类型
	fmt.Println(encBuf.Bytes())            //转成字节类型才能打印
	fmt.Println("编码后的内容", encBuf.String()) //String()是对缓存容器里面的原生字节序列进行go内置的utf8解码
	//并且打印出utf8字符串,但是由于他是GB18030值，而go却采用的是utf8的解码器去解码，所以不会成功打印的！
	//--------------------------------------------------------------

	//解码GB18030成为unicode之后再编码成为utf8
	dec := simplifiedchinese.GB18030.NewDecoder() //声明一个解码器（相对于unicode来说确实是解码）
	read := dec.Reader(encBuf)                    //跟上面一样写入缓存需要一个写入器，读出缓存也需要一个读取器来读取缓存容器里面的内容
	decBuf, err := ioutil.ReadAll(read)           //上面说到编码器写入到缓存容器中，这时候解码器需要从
	//缓存容器中读取GB18030字节序列出来，并且返回新的字节切片
	if err != nil {
		log.Println(err)
	}
	fmt.Println(decBuf) //[232 191 153 230 174 181 229 134 133 229 174 185 230 152 175 232 166 129 232 162
	// 171 231 188 150 231 160 129 232 189 172 230 141 162]
	fmt.Println("解码后的内容", string(decBuf)) //跟上面string函数一样

	//缩写方式1如下：
	//reader := simplifiedchinese.GB18030.NewDecoder().Reader(resp.Body)
	//buf, err := ioutil.ReadAll(reader)

	//中级缩写方式2如下：
	////utf8转gbk
	//src:="编码转换内容内容"
	////src:=[]byte{78, 45, 86, 253, 78, 186}//unicode码值
	//fmt.Println(src)
	//fmt.Println([]byte(src))//[231 188 150 231 160 129 232 189 172 230 141
	//// 162 229 134 133 229 174 185 229 134 133 229 174 185]
	////每3个字节为一个中文字，这是典型的utf8编码后的字节序列
	//
	////将utf8的字节序列转化为gbk的字节序列
	//data, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GBK.NewEncoder()))
	////[177 224 194 235 215 170 187 187 196 218 200 221 196 218 200 221]
	////每2个字节为一个中文字
	//fmt.Println(data) //byte
	////因为他采用的是utf8来解码序列了，很明显这是gbk才能解码的序列
	//fmt.Println(string(data))  //打印为乱码:����ת����������
}
