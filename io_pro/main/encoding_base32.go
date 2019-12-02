package main

import (
	"bytes"
	"encoding/base32"
	"fmt"
	"strings"
)

func main() {

	fmt.Println("--------------------编码-------------------")
	//双向的编码/解码协议，根据一个32字符的字符集定义，RFC 4648标准化了两种字符集。默认字符集用于SASI和GSSAPI，另一种用于DNSSEC。
	// NewEncoding返回由给定字母定义的新Encoding，该字母必须是32字节的字符串。
	str:="ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
	encoding := base32.NewEncoding(str)//给出映射的编码表
	src_ls:=[]byte{1,2,3,4,5,6,7,8,9,10,11,12,13}
	//src_ls:=[]byte{'A','B','C','D','E','F','G','H','I','J','L','M','N'}
	dst_ls:=make([]byte,encoding.EncodedLen(len(src_ls)))//必须给这个长度，否则会导致编码报错（即使给的更大的空间也会报错）,通过encoding.EncodedLen(len(src))来获取长度从而不浪费空间



	//使用编码enc对src进行编码，将EncodedLen（len（src））字节写入dst。
	//编码会将输出填充为8字节的倍数，因此Encode不适合用于大型数据流的各个块。 请改用NewEncoder（）。
	encoding.Encode(dst_ls,src_ls)
	fmt.Println(encoding)
	fmt.Println(src_ls)
	fmt.Println(dst_ls)
	fmt.Println(string(61))
	//输出：
	//	&{[65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 50 51 52 53 54 55]
	//	[255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255
	//	255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255
	//	26 27 28 29 30 31 255 255 255 255 255 255 255 255 255 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18
	//	19 20 21 22 23 24 25 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255
	//	255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255
	//	255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255
	//	255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255
	//	255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255
	//	255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255
	//	255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255] 61}
	//	[1 2 3 4 5 6 7 8 9 10 11 12 13]
	//	[65 69 66 65 71 66 65 70 65 89 68 81 81 67 73 75 66 77 71 65 50 61 61 61]
	//	=


	fmt.Println("--------------------解码-------------------")
	dec_dst_ls:=make([]byte,encoding.DecodedLen(len(dst_ls)))
	//解码使用编码enc解码src。 它最多将DecodedLen（len（src））个字节写入dst，并返回写入的字节数。
	// 如果src包含无效的base32数据，它将返回成功写入的字节数和CorruptInputError。
	//换行符（\ r和\ n）将被忽略。
	n, err := encoding.Decode(dec_dst_ls, dst_ls)
	CheckErr(err)
	fmt.Println("解码的字节数是：",n)
	fmt.Println("解码前的字节切片是：",dst_ls)
	fmt.Println("解码后的字节切片是：",dec_dst_ls)
	//输出：
	//	解码的字节数是： 13
	//	解码前的字节切片是： [65 69 66 65 71 66 65 70 65 89 68 81 81 67 73 75 66 77 71 65 50 61 61 61]
	//	解码后的字节切片是： [1 2 3 4 5 6 7 8 9 10 11 12 13 0 0]


	fmt.Println("--------------------编码解码其他方式-------------------")
	//返回n字节数据进行base32编码后的最大长度。
	fmt.Println(encoding.EncodedLen(7))//16
	//返回n字节base32编码的数据解码后的最大长度。
	fmt.Println(encoding.DecodedLen(16))//10

	//src:="abcdefghijk"
	//en_ls:=make([]byte,encoding.EncodedLen(len(src)))//这里用到了EncodedLen非常好，这样的话不浪费空间
	dst_str:=encoding.EncodeToString([]byte{1,2,3,4,5,6,7})
	fmt.Println(dst_str)//AEBAGBAFAYDQ====，内部自己维护接收器

	dst_bytes, err222 := encoding.DecodeString(dst_str)
	checkError(err222)
	fmt.Println(dst_bytes)//[1 2 3 4 5 6 7],这种方式居然正常输出


	fmt.Println("--------------------编码的填充与无填充(指定填充字节是什么，默认是填充等号，可以指定无填充)-------------------")
	// WithPadding创建一个与enc相同的但是带有指定的填充字符的新编码，或者使用NoPadding禁用填充。
	//填充字符不得为'\ r'或'\ n'，且不得包含在编码映射字母中，且其符文必须等于或小于'\ xff'。
	//enc_padding := encoding.WithPadding('A')//panic: padding contained in alphabet，不可填充在映射表中存在的值或者等于大于255的值
	//enc_padding := encoding.WithPadding(128)//因为enc_padding是由encoding产生的，所以应该在范围上面也许还更小些，具体原理不大懂
	//enc_padding := encoding.WithPadding(255)//
	//enc_padding := encoding.WithPadding(254)//
	enc_padding := encoding.WithPadding('a')
	//enc_padding := encoding.WithPadding(base32.NoPadding)//无填充
	//enc_padding := encoding.WithPadding(base32.StdPadding)//默认的填充方式
	src111:=[]byte{1,2,3,4,5,6,7,8,9,10,11,12,13}
	dst_ls111:=make([]byte,enc_padding.EncodedLen(len(src111)))
	enc_padding.Encode(dst_ls111,src111)

	fmt.Println(enc_padding)
	fmt.Println(src111)
	fmt.Println(dst_ls111)
	//填充为a的输出：
	//	&{[65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 50 51 52 53 54 55]
	//	[255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255
	//	255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255
	//	26 27 28 29 30 31 255 255 255 255 255 255 255 255 255 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18
	//	19 20 21 22 23 24 25 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255
	//	255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255
	//	255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255
	//	255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255
	//	255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255
	//	255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255
	//	255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255] 97}
	//	[1 2 3 4 5 6 7 8 9 10 11 12 13]
	//	[65 69 66 65 71 66 65 70 65 89 68 81 81 67 73 75 66 77 71 65 50 97 97 97],最后的3个补a
	//若填充的编码数据为：[65 69 66 65 71 66 65 70 65 89 68 81 81 67 73 75 66 77 71 65 50 61 61 61],最后的3个补=号


	//无填充的输出：
	//	&{[65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 50 51 52 53 54 55]
	//	[255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255
	//	255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255
	//	26 27 28 29 30 31 255 255 255 255 255 255 255 255 255 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18
	//	19 20 21 22 23 24 25 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255
	//	255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255
	//	255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255
	//	255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255
	//	255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255
	//	255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255
	//	255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 255] -1}
	//	[1 2 3 4 5 6 7 8 9 10 11 12 13]
	//	[65 69 66 65 71 66 65 70 65 89 68 81 81 67 73 75 66 77 71 65 50]，这里是无填充

	fmt.Println("--------------------编码的填充与无填充的解码(指定填充字节是什么，默认是填充等号，可以指定无填充)-------------------")

	dst_ls222:=make([]byte,enc_padding.DecodedLen(len(dst_ls111)))//采用上面编码的数据进行解码
	i, err333 := enc_padding.Decode(dst_ls222, dst_ls111)
	checkError(err333)

	fmt.Println("解码的字节数是：",i)
	fmt.Println("解码后的字节数是：",enc_padding.DecodedLen(len(dst_ls111)))
	fmt.Println("解码前的字节切片是：",dst_ls111)
	fmt.Println("解码后的字节切片是：",dst_ls222)

	//无填充的输出如下：
	//	解码的字节数是： 13
	//	解码前的字节切片是： [65 69 66 65 71 66 65 70 65 89 68 81 81 67 73 75 66 77 71 65 50]
	//	解码后的字节切片是： [1 2 3 4 5 6 7 8 9 10 11 12 13]

	//填充为a的输出如下
	//	解码的字节数是： 13
	//	解码后的字节数是： 15
	//	解码前的字节切片是： [65 69 66 65 71 66 65 70 65 89 68 81 81 67 73 75 66 77 71 65 50 97 97 97]
	//	解码后的字节切片是： [1 2 3 4 5 6 7 8 9 10 11 12 13 0 0]，这里多了2个0,是因为在创建切片时候制定了长度，
	//	                                              但是编码之后的长度并不总是会填充完这个切片容器的长度的

	fmt.Println("--------------------填充与无填充的编码解码的其他方式-------------------")
	src_str:=[]byte{1,2,3,4,5,6,7,8,9,10,11,12,13}
	s := enc_padding.EncodeToString(src_str)
	fmt.Println(s)

	bytes99, err444 := enc_padding.DecodeString(s)
	checkError(err444)
	fmt.Println("解码后的：",bytes99)
	//输出：
	//	AEBAGBAFAYDQQCIKBMGA2aaa
	//	解码后的： [1 2 3 4 5 6 7 8 9 10 11 12 13]

	fmt.Println("--------------------无填充的encoding也可以产生无填充的encoding,但是只能编码，不可用于解码-------------------")
	enc_padding222 := enc_padding.WithPadding('b')

	src_str222:=[]byte{1,2,3,4,5,6,7,9,10,11,12,13}
	s2 := enc_padding222.EncodeToString(src_str222)
	fmt.Println(s2)
	//如果是采用原来的东西进行解码的话还是可以解码的,但是如果采用由padding_encoding衍生出来的padding_encoding的话解码是会出错的，因为2个padding_encoding的填充字符是不一样的
	//原因看源码就知道了，因为每次创建padding_encoding都是拿到原来的编码指针并且声称相应的指针对象，并且将填充字符串设置为相对应的值，但是这个值是保存在新建的指针中，所以无论
	// 产生多少的padding_encoding都不会相互影响，也就是&enc，但是如果你由&enc再次产生padding_encoding的话就会返回&&enc了。当然一样可以拿到对象，
	// 对一个内存地址值进行&拿地址还是会输出这个地址的。下面我们采用enc_padding来解码s2是不会解码成功的，enc_padding只能解码按照自己编码规则编码的数据（他的编码规则就是填充a）
	// 而enc_padding222的编码规则是完全不一样的，当且仅当2中编码规则一样的话才可以相互解码数据！不然的话会报错的！比如这个错*** illegal base32 data at input byte 20，显示无法编码的数据的意思
	bytes2, err4442 := enc_padding222.DecodeString(s2)//
	//bytes2, err4442 := enc_padding.DecodeString(s2)//会报错

	//CheckErr(err4442)
	if err4442 !=nil{
		fmt.Println("***",err4442)
	}
	fmt.Println("解码后的：",bytes2)

	//输出：
	//	AEBAGBAFAYDQSCQLBQGQbbbb，很明显和上面的第一代的padding_encoding是不太一样的，但是却可以解码出来同一个结果，非常奇怪
	//	解码后的： [1 2 3 4 5 6 7 8 9 10 11 12 13]

	fmt.Println("-----------------------方便快捷的base32标准编码解码（也就是默认采用go里面给出的编码表来进行编码解码）-------------------------------------")

	//编码
	data := []byte("any + old & data")
	str_new := base32.StdEncoding.EncodeToString(data)
	fmt.Println(str_new)

	//解码
	decodeByte, err555 := base32.StdEncoding.DecodeString(str_new)
	CheckErr(err555)
	fmt.Println(decodeByte)
	fmt.Println(string(decodeByte))
	//输出：
	//	MFXHSIBLEBXWYZBAEYQGIYLUME======
	//	[97 110 121 32 43 32 111 108 100 32 38 32 100 97 116 97]
	//	any + old & data
	//更多的操作不再展示


	fmt.Println("-----------将数据编码并且写入io.Writer-----------")

	// NewEncoder返回一个新的base32流编码器。 写入返回的写入器的数据将使用enc进行编码，然后写入w。
	// Base32编码以5字节块为单位； 完成写入后，调用者必须关闭返回的编码器以刷新任何部分写入的块。
	dst_bf:=make([]byte,0,30)
	buffer := bytes.NewBuffer(dst_bf)//创建一个缓存器来装解码后的东西
	writeCloser111 := base32.NewEncoder(base32.StdEncoding, buffer)//要指明写入到哪里去，这里我采用的是内置的标准的编码表
	dst_byte111:=[]byte{1,2,3,4,5,6,7,9,10,11,12,13}//要编码的数据，这里漏了8，没事，记得一共才12个数据
	n111, err1111 := writeCloser111.Write(dst_byte111)//写什么
	checkError(err1111)
	writeCloser111.Close()//必须关闭流才可以把全部的字节写入到切片中去

	fmt.Println("编码的字节个数为：",n111)
	fmt.Println("编码前的数据为：",dst_byte111)
	fmt.Println("编码后的数据为：",buffer)
	fmt.Println("buffer底层切片为：",dst_bf)
	//输出：
	//	编码的字节个数为：12
	//	编码前的数据为： [1 2 3 4 5 6 7 9 10 11 12 13]
	//	编码后的数据为： AEBAGBAFAYDQSCQLBQGQ====
	//	buffer底层切片为： []

	fmt.Println("-----------从io.Reader读取要编码的数据来进行编码-----------")

	str_R:="AEBAGBAFAYDQSCQLBQGQ===="//这个是解码后的数据
	enc_reader := strings.NewReader(str_R)
	new_Reader:=base32.NewDecoder(base32.StdEncoding,enc_reader)//记得这里要采用和编码时候相同的编码解码器，否则无法正确解码。

	fmt.Println("读取前的读取器里面的数据：",enc_reader)

	dst_ls2223:=make([]byte,20)
	n9, err99 := new_Reader.Read(dst_ls2223)
	checkError(err99)
	fmt.Println("解码了的字节个数：",n9)
	fmt.Println("解码前的数据：",str_R)
	fmt.Println("读取后的读取器里面的数据：",enc_reader)
	fmt.Println("解码后的数据：",dst_ls2223)
	//输出：
	//	读取前的读取器里面的数据： &{AEBAGBAFAYDQSCQLBQGQ==== 0 -1}
	//	解码了的字节个数： 12
	//	解码前的数据： AEBAGBAFAYDQSCQLBQGQ====
	//		读取后的读取器里面的数据： &{AEBAGBAFAYDQSCQLBQGQ==== 24 -1}
	//	解码后的数据： [1 2 3 4 5 6 7 9 10 11 12 13 0 0 0 0 0 0 0 0]

	fmt.Println("-----------------------------")
}

func CheckErr(err error)  {
	if err != nil{
		//似乎并不是所有的err值都是有类型的，所以不建议在这里采用获取类型的方式，会导致报错的位置很奇怪，而且这里也会出错
		fmt.Printf("%T---%v\n",err,err)
	}
}

















