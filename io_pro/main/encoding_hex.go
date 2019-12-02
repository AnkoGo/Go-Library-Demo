package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

func main() {
	// Package hex实现十六进制编码和解码。

	fmt.Println("-------------hex.Encode()和hex.EncodeToString()编码数据-----------------")



	src_ls:=[]byte{'a','b','c','d'}
	// EncodedLen返回n个源字节的编码长度。
	//具体来说，它返回n * 2。
	enc_len := hex.EncodedLen(len(src_ls))
	dst_ls:=make([]byte,enc_len)

	//将src编码为dst的EncodedLen（len（src））字节。 为方便起见，它返回写入dst的字节数，但此值始终为EncodedLen（len（src））。
	//编码实现十六进制编码。
	dst_n := hex.Encode(dst_ls, src_ls)//enc_len==dst_n


	fmt.Println("编码前的数据字节为：",src_ls)
	fmt.Println("编码后的字节长度为：",dst_n)
	fmt.Println("编码后的数据字节为：",dst_ls)

	//输出：
	//	编码前的数据字节为： [97 98 99 100]
	//	编码后的字节长度为： 8
	//	编码后的数据字节为： [54 49 54 50 54 51 54 52]

	//底层采用的是encode()

	fmt.Println()
	dst_str := hex.EncodeToString(src_ls)
	fmt.Println("编码前的字节为：",src_ls)
	fmt.Println("编码后的字节长度为：",len(dst_str))
	fmt.Println("编码后的数据字符串为：",dst_str)
	fmt.Println("编码后的数据字节为：",[]byte(dst_str))

	//输出：
	//	编码前的字节为： [97 98 99 100]
	//	编码后的字节长度为： 8
	//	编码后的数据字符串为： 61626364
	//	编码后的数据字节为： [54 49 54 50 54 51 54 52]

	fmt.Println("-------------hex.NewEncoder新建编码写入器，.Write()编码数据-----------------")

	buffer := new(bytes.Buffer)
	encoder_wt := hex.NewEncoder(buffer)

	// Writer是包装基本Write方法的接口。
	// Write将p的len（p）个字节写入基础数据流。
	//返回从p（0 <= n <= len（p））写入的字节数，以及遇到的任何导致写入提前停止的错误。
	//如果写入返回n <len（p），则必须返回一个非nil错误。
	//写操作不得修改切片数据，即使是临时的也不行。
	//实现不得保留p。
	n, enc_err := encoder_wt.Write(src_ls)
	checkErr_hex(enc_err)

	fmt.Println("编码前的字节数据为：",src_ls)
	fmt.Println("编码写入缓存的原始字节数：",n)//但是不是编码写入的字节数，即是要编码的字节数，而不是编码后的字节数
	fmt.Println("编码后的数据为：",buffer)
	fmt.Println("编码后的数据字节为：",buffer.Bytes())//其实这个并不算是读取

	//输出：
	//	编码前的字节数据为： [97 98 99 100]
	//	编码写入缓存的原始字节数： 4
	//	编码后的数据为： 61626364
	//	编码后的数据字节为： [54 49 54 50 54 51 54 52]



	fmt.Println("-------------------hex.Decode和hex.DecodeString解码数据--------------------------")

	decodedLen := hex.DecodedLen(len(dst_ls)) //我们采用上面编码过的数据进行解码
	dec_dst_ls:=make([]byte,decodedLen)//创造装载解码后的数据的容器

	//解码将src解码为DecodedLen（len（src））字节，并返回写入dst的实际字节数。
	//解码期望src仅包含十六进制字符，并且src具有偶数长度。
	//如果输入格式错误，则Decode返回错误发生之前解码的字节数。
	i, dec_err := hex.Decode(dec_dst_ls, dst_ls)
	checkErr_hex(dec_err)

	fmt.Println("解码前的字节数据：",dst_ls)
	fmt.Println("解码后的字节个数：",i)
	fmt.Println("解码后的字节数据：",dec_dst_ls)

	//输出：
	//	解码前的字节数据： [54 49 54 50 54 51 54 52]
	//	解码后的字节个数： 4
	//	解码后的字节数据： [97 98 99 100]

	fmt.Println()
	dec_dst_byte, dec_err111 := hex.DecodeString(dst_str)
	checkErr_hex(dec_err111)

	fmt.Println("解码前的字节数据：",[]byte(dst_str))
	fmt.Println("解码前的字符串数据：",dst_str)
	fmt.Println("解码后的字节数据：",dec_dst_byte)
	fmt.Println("解码后的字符串数据：",string(dec_dst_byte))
	//输出：
	//	解码前的字节数据： [54 49 54 50 54 51 54 52]
	//	解码前的字符串数据： 61626364
	//	解码后的字节数据： [97 98 99 100]
	//	解码后的字符串数据： abcd


	fmt.Println("------------hex.NewDecoder(yy)依据要解码的内容yy来新建解码读取器和.Read（xx）读取解码数据到xx-------------------")

	fmt.Println(buffer.Bytes())//这并不移动字节指针off,底层是return b.buf[b.off:]，返回的是一个buffer中的字段buf底层切片的副本而已
								//所以可以多次获取都不会影响到buffer
	fmt.Println(buffer.Bytes())
	fmt.Println(buffer.Bytes())
	bf_ls:=buffer.Bytes()//备份buffer,以便下面打印
	// NewDecoder返回io.Reader，该io.Reader解码r中的十六进制字符。
	// NewDecoder期望r仅包含偶数个十六进制字符。
	dec_Rd := hex.NewDecoder(buffer)

	dec_dst_ls1111:=make([]byte,hex.DecodedLen(buffer.Len()))//解码后的数据的转载器
	n2, dec_err222 := dec_Rd.Read(dec_dst_ls1111)//读取到哪里去
	checkErr_hex(dec_err222)

	fmt.Println("解码前的字节数据：",buffer.Bytes())//从buffer里面读取光数据了，所以这里没了
	fmt.Println("解码前的缓存数据：",bf_ls)//
	fmt.Println("解码后的字节个数：",n2)
	fmt.Println("解码后的字节数据：",dec_dst_ls1111)
	fmt.Println("解码后的字符串数据：",string(dec_dst_ls1111))
	//输出：
	//	[54 49 54 50 54 51 54 52]
	//	[54 49 54 50 54 51 54 52]
	//	[54 49 54 50 54 51 54 52]
	//	解码前的字节数据： []
	//	解码前的缓存数据： [54 49 54 50 54 51 54 52]
	//	解码后的字节个数： 4
	//	解码后的字节数据： [97 98 99 100]
	//	解码后的字符串数据： abcd

	fmt.Println("-------------------hex.Dump()堆内存方式显示字节切片-------------------------")

	//转储返回一个字符串，这个字符串包含给定数据的十六进制转储。 十六进制转储的格式与命令行上“ hexdump -C”的输出匹配。
	//底层是采用的hex.dumper这个结构体实现的

	//dump_str := hex.Dump(src_ls)
	dump_str := hex.Dump(dst_ls)
	fmt.Println(dump_str)
	//src_ls输出：
	//00000000  61 62 63 64                                       |abcd|
	//dst_ls输出：
	//00000000  36 31 36 32 36 33 36 34                           |61626364|
	//由上可知，事实上他只要求是一个合法的切片，而并不管是不是hex进制的字节切片与否

	fmt.Println("-------------------hex.Dump()堆内存方式显示字节切片-------------------------")

	newBuffer := new(bytes.Buffer)
	dumper_wtclo := hex.Dumper(newBuffer)

	src_wtclo:=src_ls//同理，这里根本不管写入的是什么进制的字节切片，但是我们一般是写入的是16进制的字节切片，
	// 因为这2个函数都是想要显示的是16进制的堆内存形式的数据
	//src_wtclo:=dst_ls


	// Writer是包装基本Write方法的接口。
	// Write将p的len（p）个字节写入基础数据流。
	//返回从p（0 <= n <= len（p））写入的字节数，以及遇到的任何导致写入提前停止的错误。
	//如果写入返回n <len（p），则必须返回一个非nil错误。
	//写操作不得修改切片数据，即使是临时的也不行。
	//实现不得保留p。
	n3, enc_err333 := dumper_wtclo.Write(src_wtclo)
	checkErr_hex(enc_err333)
	dumper_wtclo.Close()//必须写上，不写会导致数据不完整

	fmt.Println("写入newBuffer的原始字节个数为：",n3)//不是指接收器的字节个数，是指原始的字节个数
	fmt.Println("newBuffer的字节为：",newBuffer.Bytes())
	fmt.Println("newBuffer的字符串为：",newBuffer.String())
	//src_ls输出：
	//	写入newBuffer的原始字节个数为： 4
	//	newBuffer的字节为： [48 48 48 48 48 48 48 48 32 32 54 49 32 54 50 32 54 51 32 54 52 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 124 97 98 99 100 124 10]
	//	newBuffer的字符串为： 00000000  61 62 63 64                                       |abcd|

	//dst_ls输出：
	//	写入newBuffer的原始字节个数为： 8
	//	newBuffer的字节为： [48 48 48 48 48 48 48 48 32 32 51 54 32 51 49 32 51 54 32 51 50 32 51 54 32 51 51 32 51 54 32 51 52 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 32 124 54 49 54 50 54 51 54 52 124 10]
	//	newBuffer的字符串为： 00000000  36 31 36 32 36 33 36 34                           |61626364|

	//课外：
	// InvalidByteError值描述了十六进制字符串中无效字节导致的错误。
	//hex.InvalidByteError()



}

func checkErr_hex(err error)  {
	if err != nil{
		fmt.Print(err)
	}
}













