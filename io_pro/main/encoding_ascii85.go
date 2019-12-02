package main

import (
	"bytes"
	"encoding/ascii85"
	"fmt"
)

func main() {

	fmt.Println("------------编码---------------")
	dst_ls:=make([]byte,15)
	src_ls:=[]byte{97,98,99,100,101,102}

	//将src编码成最多MaxEncodedLen(len(src))数据写入dst，返回实际写入的字节数。编码每4字节一段进行一次，
	//最后一个片段采用特殊的处理方式，因此不应将本函数用于处理大数据流的某一独立数据块。
	//一般来说ascii85编码数据会被'<~'和'~>'包括起来，函数并未添加上它们。
	n:=ascii85.Encode(dst_ls,src_ls)//每次取4个字节为一段来进行转码，不够4个就取光来转码，不过无论去多少个，最后一定生成5个字节输出
	fmt.Println(n)
	fmt.Println("源切片",src_ls,"字符串显示为：",string(src_ls))
	fmt.Println("转换后的目的切片",dst_ls,"字符串显示为：",string(dst_ls))

	//输出：
	//	8,这个8=(len(src_ls)/4)*5+(len(src_ls)%4)+1,具体表示什么真的不好说
	//	源切片 [97 98 99 100 101 102] 字符串显示为： abcdef
	//	转换后的目的切片 [64 58 69 95 87 65 83 40 111 66 0 0 0 0 0] 字符串显示为： @:E_WAS(oB，编码后的
	//	字节序列的长度为(len(src_ls) + 3) / 4 * 5,当然可以通过ascii85.MaxEncodedLen（）获取，转码后的字节序列的长度一定会是5的倍数

	fmt.Println("-------------解码--------------")

	//将src解码后写入dst，返回写入dst的字节数、从src解码的字节数。如果src含有非法数据，函数将返回成功执行的数据（两个数字）
	//和CorruptInputError。如果flush为真，则函数会认为src代表输入流的结尾，完全处理src，而不会等待另一个32字节的数据块。
	//
	//函数会忽略src中的空格和控制字符，一般来说ascii85编码数据会被'<~'和'~>'包括起来，但是调用者应自行去掉它们。
	dst_ls111:=make([]byte,15)
	src_ls111:=[]byte{64 ,58, 69, 95, 87, 65, 83, 40}//这个数据就是上面转码后的数据，这种ascii85.Decode的第三个参数必须设置为true才可以正确的解码
	//src_ls111:=[]byte{64 ,58, 69, 95, 87, 65, 83, 40, 111, 66}//这个数据就是上面转码后的数据
	ndst, nsrc, err := ascii85.Decode(dst_ls111, src_ls111, true)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(ndst)
	fmt.Println("解码的源序列字节数为：",nsrc,"要解码数据为：",src_ls111)
	fmt.Println("解码后的序列的字节数为：",ndst,"解码后的数据为：",dst_ls111)
	//输出：
	//	解码的源序列字节数为： 10 要解码数据为： [64 58 69 95 87 65 83 40 111 66]
	//	解码后的序列的字节数为： 8 解码后的数据为： [97 98 99 100 101 102 0 0 0 0 0 0 0 0 0]

	fmt.Println("--------------ascii85.Encode()封装函数ascii85.NewEncoder()-----------------")

	// MaxEncodedLen返回n个源字节的编码的最大长度。

	// 编码后的字节序列的长度为(len(src_ls) + 3) / 4 * 5,当然可以通过ascii85.MaxEncodedLen（）获取，转码后的字节序列的长度一定会是5的倍数
	fmt.Println(ascii85.MaxEncodedLen(8))//10
	fmt.Println(ascii85.MaxEncodedLen(9))//15

	fmt.Println("-------------------------------")
	// NewEncoder返回一个新的ascii85流编码器。 写入返回的写入器的数据将被编码，然后写入w。
	// Ascii85编码以32位块运行； 完成写入后，调用者必须关闭返回的编码器以刷新任何尾随的部分块。
	// 底层也是通过ascii85.Encode（）实现的，不过放在2个函数里面实现，一个是ascii85.NewEncoder，另外一个是writeCloser.Close()
	ls_bf:=make([]byte,0,15)
	buffer := bytes.NewBuffer(ls_bf)
	writeCloser := ascii85.NewEncoder(buffer)

	ls_src222:=[]byte{97,98,99,100,101,102}
	i, err111 := writeCloser.Write(ls_src222)//写入什么
	if err111 !=nil{
		fmt.Println(err111)
	}
	// Close刷新编码器的所有未决输出。必须调用
	//调用Close之后调用Write会出错。
	err222:=writeCloser.Close()//不够4个字节的尾字节都会在这里编码，也就是说，如果不写这句话的话，
	// 那么编码将只有5个字节而不会有后面的3个字节，因为后面的3个字节是由最尾的2个源切片的字节编码后生成的。
	if err222 !=nil{
		fmt.Println(err222)
	}

	fmt.Println(i)
	fmt.Println("源切片：",ls_src222,"字符串显示为：",string(ls_src222),"源切片长度：",len(ls_src222))
	fmt.Println("转换后的目的切片：",buffer.Bytes(),"字符串显示为：",buffer.String(),"转换后的目的切片长度：",buffer.Len())
	//输出：
	//	6
	//	源切片： [97 98 99 100 101 102] 字符串显示为： abcdef 源切片长度： 6
	//	转换后的目的切片： [64 58 69 95 87 65 83 40] 字符串显示为： @:E_WAS( 转换后的目的切片长度： 8

	//从输出可以看出比ascii85.Encode（）的编码少了2个字节，至于为什么不大清楚，不过可以肯定的是，这种方式返回的写入字节数会跟源字节数相等
	//从ascii85.Encode（）的文档中可以知道Encode（）不大适合处理大数据流，而眼前的这个函数的话肯定是适合处理大数据流的。
	//对于这个编码生成的序列我们不能采用上面的ascii85.Decode()函数来进行解码，否则解码后的字节序列会丢掉后面的2个字节，还原不了源字节序列。
	//而应该采用下面讲到的这个函数来进行解码

	fmt.Println("--------------函数ascii85.NewEncoder()-----------------")

	//src_ls222:=[]byte{64 ,58, 69, 95, 87 ,65 ,83, 40,}//这个数据是无法通过ascii85.NewDecoder（）进行解码的
	src_ls222:=[]byte{64 ,58, 69, 95, 87 ,65 ,83, 40,111 ,66}//只能解码由ascii85.Encode()编码后的数据，不能解码由ascii85.newEncode()编码后的数据，至于为什么不大清楚
	//reader := bytes.NewReader(src_ls222)//这种方式创建io.Reader也是可以的
	reader := bytes.NewBuffer(src_ls222)
	newDecoder_reader := ascii85.NewDecoder(reader)
	dst_ls222:=make([]byte,15)
	n2, err333 := newDecoder_reader.Read(dst_ls222)
	if err333 !=nil{
		fmt.Println(err333)
	}
	fmt.Println(n2)

	fmt.Println("解码的源序列字节数为：",len(src_ls222),"要解码数据为：",string(src_ls222))
	fmt.Println("解码后的序列的字节数为：", n2,"解码后的数据为：",dst_ls222)
	//输出：
	//	8
	//	解码的源序列字节数为： 10 要解码数据为： @:E_WAS(oB
	//	解码后的序列的字节数为： 8 解码后的数据为： [97 98 99 100 101 102 0 0 0 0 0 0 0 0 0]

	//总结，解码最好通过ascii85.Decoder（）来进行解码，但是编码的话我们可以通过2中方式来进行编码，分别是ascii85.NewEncoder（）和ascii85.Encoder（）

}





















