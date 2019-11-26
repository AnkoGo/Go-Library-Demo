package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"
	"strings"
	"unicode/utf8"
)

func main2367() {

	fmt.Println("------------------Encode()编码数据---------------------")
	bf_ls := make([]byte, 0, 20)     //记得作为要写入的缓存我们只需要给一定大小的cap和给0 len
	buffer := bytes.NewBuffer(bf_ls) //申请一个缓存来装编码后的数据
	// NewEncoder返回一个新的编码器，它将在io.Writer上传输。
	gob_encoder := gob.NewEncoder(buffer)
	src_ls := []byte{'a', 'b', 'c', 'd', 'e'}
	//编码传输由空接口值表示的数据项，从而确保所有必需的类型信息都已首先传输。
	//将nil指针传递给Encoder会感到恐慌，因为它们无法通过gob传输。
	encode_err := gob_encoder.Encode(src_ls)
	checkErr_2(encode_err)

	fmt.Println("编码前的数据为：", src_ls)
	fmt.Println("编码后的数据为：", buffer)
	fmt.Println("编码后的数据字节为：", buffer.Bytes())
	fmt.Println("换行字节为：", '\n')
	//输出：
	//	编码前的数据为： [97 98 99 100 101]
	//	编码后的数据为：
	//	abcde
	//	编码后的数据字节为： [8 10 0 5 97 98 99 100 101]
	//	换行字节为： 10
	//解释：
	//	[8 10 0 5 97 98 99 100 101]，对于后面的5个肯定是原来的abcde的ascii值，对前面的4个：第一个字节8是除了第一个字节之外还有多少字节
	//	表示这个切片结构的编码数据，还有8个；第二个字节10代表字节切片类型的type id的倍数,第3个字节0不知道代表什么，第4个字节5代表真正的表示里面元素的个数

	fmt.Println("------------------Encode()编码数据111---------------------")

	bf_ls111 := make([]byte, 0, 20)
	buffer111 := bytes.NewBuffer(bf_ls111)

	gob_encoder111 := gob.NewEncoder(buffer111)
	src_ls111 := "abcd"

	encode_err111 := gob_encoder111.Encode(src_ls111)
	checkErr_2(encode_err111)

	fmt.Println("编码前的数据为：", src_ls111)
	fmt.Println("编码后的数据为：", buffer111)
	fmt.Println("编码后的数据字节为：", buffer111.Bytes())
	//输出：
	//	编码前的数据为： abcd
	//	编码后的数据为：  abcd
	//	编码后的数据字节为： [7 12 0 4 97 98 99 100]

	fmt.Println("------------------Encode()编码数据222---------------------")

	bf_ls222 := make([]byte, 0, 20)
	buffer222 := bytes.NewBuffer(bf_ls222)

	gob_encoder222 := gob.NewEncoder(buffer222)
	src_ls222 := 12340

	encode_err222 := gob_encoder222.Encode(src_ls222)
	checkErr_2(encode_err222)

	fmt.Println("编码前的数据为：", src_ls222)
	fmt.Println("编码后的数据为：", buffer222)
	fmt.Println("编码后的数据字节为：", buffer222.Bytes())
	// FullRune报告p中的字节是否以符文的完整UTF-8编码开头。
	//无效的编码被视为完整的符文，因为它将转换为宽度为1的错误符文。
	buffer_byte := buffer222.Bytes()
	fmt.Println("编码后的数据完全是utf8编码的么？",utf8.Valid(buffer_byte))
	fmt.Println(utf8.FullRune(buffer_byte))
	fmt.Println(utf8.FullRune([]byte{254, 0, 5, 104, 96, 4}))

	// DecodeRune解码p中的第一个UTF-8编码，并返回符文及其宽度（以字节为单位）。 如果p为空，则返回（RuneError，0）。
	// 否则，如果编码无效，则返回（RuneError，1）。 对于正确的非空UTF-8，这都是不可能的结果。
	//如果编码不正确，则编码无效；如果编码不正确，则编码超出范围，或者不是该值的最短UTF-8编码。 不执行其他任何验证。
	for k := 0; k < len(buffer_byte); {
		fmt.Println(utf8.DecodeRune(buffer_byte))
		buffer_byte = buffer_byte[k+1:]
	}
	fmt.Println(buffer_byte)
	//输出：
	//	编码前的数据为： 12340
	//	编码后的数据为：  �`h
	//	编码后的数据字节为： [5 4 0 254 96 104]
	//	true
	//	true
	//	5 1
	//	4 1
	//	0 1
	//	65533 1
	//	96 1
	//	104 1
	//	[]

	fmt.Println("------------------Encode()编码数据333---------------------")

	bf_ls333 := make([]byte, 0, 20)
	buffer333 := bytes.NewBuffer(bf_ls333)

	gob_encoder333 := gob.NewEncoder(buffer333)
	var src_ls333 float64 = 1234.89

	encode_err333 := gob_encoder333.Encode(src_ls333)
	checkErr_2(encode_err333)

	fmt.Println("编码前的数据为：", src_ls333)
	fmt.Println("编码后的数据为：", buffer333)
	fmt.Println("编码后的数据字节为：", buffer333.Bytes())
	// FullRune报告p中的字节是否以符文的完整UTF-8编码开头。
	//无效的编码被视为完整的符文，因为它将转换为宽度为1的错误符文。
	fmt.Println(utf8.FullRune(buffer333.Bytes()))
	buffer_byte333:=buffer333.Bytes()
	fmt.Println("编码后的数据完全是utf8编码的么？",utf8.Valid(buffer_byte333))
	for k := 0; k < len(buffer_byte333); {
		fmt.Println(utf8.DecodeRune(buffer_byte333))
		buffer_byte333 = buffer_byte333[k+1:]
	}
	fmt.Println(buffer_byte333)
	//输出：
	//	编码前的数据为： 1234.89
	//	编码后的数据为：  ���(\�K�@
	//	编码后的数据字节为： [11 8 0 248 195 245 40 92 143 75 147 64]
	//	true
	//	11 1
	//	8 1
	//	0 1
	//	65533 1
	//	65533 1
	//	65533 1
	//	40 1
	//	92 1
	//	65533 1
	//	75 1
	//	65533 1
	//	64 1
	//	[]
	//注释：
	//	从这里看，似乎也不完全是utf8编码的字符，看来是自己特定的编码



	fmt.Println("------------------EncodeValue()编码数据---------------------")
	//其实这个方法跟上面的方法差不了多少的，只是说提供了解决不同输入的参数类型的编码而已，只是增强*Encoder对象的实用性
	bf_ls444 := make([]byte, 0, 20)
	buffer444 := bytes.NewBuffer(bf_ls444)

	gob_encoder444 := gob.NewEncoder(buffer444)
	var src_ls444 interface{} = "abcd"
	// EncodeValue传输由反射值表示的数据项，从而确保所有必需的类型信息均已首先传输。
	//将nil指针传递给EncodeValue会感到恐慌，因为它们无法通过gob传输。
	encode_err444 := gob_encoder444.EncodeValue(reflect.ValueOf(src_ls444))
	checkErr_2(encode_err444)

	fmt.Println("编码前的数据为：", src_ls444)
	fmt.Println("编码后的数据为：", buffer444)
	fmt.Println("编码后的数据字节为：", buffer444.Bytes())
	//输出：
	//	编码前的数据为： abcd
	//	编码后的数据为：  abcd
	//	编码后的数据字节为： [7 12 0 4 97 98 99 100]


	fmt.Println("------------------Decode()解码数据---------------------")

	gob_decoder := gob.NewDecoder(buffer444)//我们直接采用上面编码好的东西进行解码
	//解码从输入流中读取下一个值，并将其存储在由空接口值表示的数据中。
	//如果e为nil，则该值将被丢弃。 否则，e底下的值必须是指向接收到的下一个数据项的正确类型的指针。
	//如果输入在EOF处，则Decode返回io.EOF且不修改e。

	//因为解码必须会读取缓存里面的东西，而缓存编读取会边pop东西出来，所以我们先保存这个缓存原始值
	dec_buffer:=buffer444.Bytes()
	var dec_str string
	decode_err := gob_decoder.Decode(&dec_str) //这个参数是决定解码成什么数据类型,必须使用指针
	checkErr_2(decode_err)

	fmt.Println("解码后的buffer数据为：", buffer444)//这时候的数据已经被读取光了
	fmt.Println("解码后的buffer数据字节为：", buffer444.Bytes())
	fmt.Println("解码前的数据为：", dec_buffer)
	fmt.Println("解码后的数据为：", dec_str)
	//输出：
	//	解码后的buffer数据为：
	//	解码后的buffer数据字节为： []
	//	解码前的数据为： [7 12 0 4 97 98 99 100]
	//	解码后的数据为： abcd

	fmt.Println("------------------EncodeValue()解码数据---------------------")
	//我们不能使用上面的解码器，因为解码器绑定的要解码的数据是一个buffer，而这个buffer已经在上面消耗光了，buffer不可多次读取或者解码的
	fmt.Println(dec_buffer)
	gob_decoder111 := gob.NewDecoder(strings.NewReader(string(dec_buffer)))//我们采用他的副本字节切片，副本有值,原buffer无值

	var dst_str111 string
	decode_err111:= gob_decoder111.DecodeValue(reflect.ValueOf(&dst_str111))
	checkErr_2(decode_err111)

	fmt.Println("解码后的buffer数据为：", buffer444)//这时候的数据已经被读取光了
	fmt.Println("解码后的buffer数据字节为：", buffer444.Bytes())
	fmt.Println("解码前的数据为：", dec_buffer)
	fmt.Println("解码后的数据为：", dst_str111)
	//输出：
	//	[7 12 0 4 97 98 99 100]
	//	解码后的buffer数据为：
	//	解码后的buffer数据字节为： []
	//	解码前的数据为： [7 12 0 4 97 98 99 100]
	//	解码后的数据为： abcd

	fmt.Println("------------------gob.Register记录类型然后编码数据---------------------")
	// GobEncoder是描述数据的接口，该接口提供了自己的表示形式，用于编码值以传输到GobDecoder。
	// 实现GobEncoder和GobDecoder的类型可以完全控制其数据的表示形式，因此可以包含诸如私有字段，
	// 通道和函数之类的东西，这些通常在gob流中是不可传输的。
	//注意：由于gob可以永久存储，因此最好的设计是确保GobEncoder使用的编码随软件的发展而稳定。
	// 例如，对于GobEncode在编码中包含版本号可能很有意义。
	//gob.GobEncoder()
	// Register在其内部类型名称下记录一个由该类型的值标识的类型。 该名称将标识作为接口变量发送或接收的值的具体类型。
	// 仅将要注册的类型将作为接口值的实现进行注册。
	//期望仅在初始化期间使用，如果类型和名称之间的映射不是双射，则会出现恐慌。

	//底层采用的额是map.go文件里面的方法LoadOrStore（）实现的
	// LoadOrStore返回键的现有值（如果存在）。
	//否则，它将存储并返回给定值。
	//如果已加载该值，则加载的结果为true，如果已存储，则为false。

	type person struct {
		Name string
		Age int
		isstudent bool//私有字段并不会编码
	}

	src_ls999:=person{"anko",18,true}
	//src_ls999 := [][]byte{[]byte{'a','b'},[]byte{'c','d','e'}}
	bf_ls999 := make([]byte, 0, 20)
	buffer999 := bytes.NewBuffer(bf_ls999)

	gob_encoder999 := gob.NewEncoder(buffer999)
	gob.Register(&src_ls999)//应该放在初始化之前，似乎并没什么作用

	encode_err999 := gob_encoder999.Encode(src_ls999)
	checkErr_2(encode_err999)

	fmt.Println("编码前的数据为：", src_ls999)
	fmt.Println("编码后的数据为：", buffer999)
	fmt.Println("编码后的数据字节为：", buffer999.Bytes())
	//字节切片的话输出：
	//	编码前的数据为： [[97 98] [99 100 101]]
	//	编码后的数据为： ���� 
	//	�� abcde
	//	编码后的数据字节为： [12 255 129 2 1 2 255 130 0 1 10 0 0 11 255 130 0 2 2 97 98 3 99 100 101]

	//结构体的话输出：
	//	编码前的数据为： {anko 18}
	//	编码后的数据为： %��person�� Name Age   ��anko$
	//	编码后的数据字节为： [37 255 129 3 1 1 6 112 101 114 115 111 110 1 255 130 0 1 2 1 4 78 97 109 101
	//						1 12 0 1 3 65 103 101 1 4 0 0 0 11 255 130 1 4 97 110 107 111 1 36 0]

	//对这个东西的讲解请看下一节
}
func checkErr_2(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
