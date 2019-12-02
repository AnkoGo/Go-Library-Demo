package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("**********************************")

	//有效报告p是否完全由有效的UTF-8编码的符文组成。
	fmt.Println(utf8.Valid([]byte{'a','b','c','z'}))//true
	fmt.Println(utf8.Valid([]byte{'a','b','c'}))//true
	fmt.Println(utf8.Valid([]byte{97,98,99}))//true
	fmt.Println(utf8.Valid([]byte{255,254,253}))//false
	fmt.Println(utf8.Valid([]byte{128,}))//false
	fmt.Println(utf8.Valid([]byte{127,}))//true(似乎ascii是最大的是127,utf8可以编码ascii和gbk系列)

	valid := []byte("Hello, 世界")
	invalid := []byte{0xff, 0xfe, 0xfd}
	fmt.Println(utf8.Valid(valid))//true
	fmt.Println(utf8.Valid(invalid))//false


	// ValidRune报告r是否可以合法编码为UTF-8。
	//超出范围或替代一半的代码点是非法的。
	//下面输出全部为true
	fmt.Println(utf8.ValidRune('a'))
	fmt.Println(utf8.ValidRune(128))
	fmt.Println(utf8.ValidRune(127))
	fmt.Println(utf8.ValidRune(254))
	fmt.Println(utf8.ValidRune('世'))
	fmt.Println(utf8.ValidRune(0xff))
	fmt.Println(utf8.ValidRune(0xfd))

	valid1 := 'a'
	invalid1 := rune(0xfffffff)
	valid2 := rune(0xfffff)
	fmt.Println(utf8.ValidRune(valid1))//true
	fmt.Println(utf8.ValidRune(invalid1))//false
	fmt.Println(utf8.ValidRune(valid2))//true
	//范围：
	// [0,0xD800)
	//(0xDFFF,\U0010FFFF]

	//报告s是否包含完整且合法的utf-8编码序列。
	fmt.Println(utf8.ValidString("0xfd"))//true
	fmt.Println(utf8.ValidString("adsdsd"))//true
	fmt.Println(utf8.ValidString("adsdsd中"))//true
	fmt.Println(utf8.ValidString(" 中  "))//true
	fmt.Println(utf8.ValidString(" 中  "))//true

	valid4 := "Hello, 世界"
	invalid4 := string([]byte{0xff, 0xfe, 0xfd})//自己组成字符串
	fmt.Println(utf8.ValidString(valid4))//true
	fmt.Println(utf8.ValidString(invalid4))//false

	// EncodeRune将符文的UTF-8编码写入p（必须足够大）。
	//返回写入的字节数。
	ls:=make([]byte,10)
	ls1:=make([]byte,10)
	ls2:=make([]byte,10)
	ls3:=make([]byte,10)
	fmt.Println(utf8.EncodeRune(ls,'a'))//1
	fmt.Println(utf8.EncodeRune(ls1,'中'))//3
	fmt.Println(utf8.EncodeRune(ls2,'b'))//1
	fmt.Println(utf8.EncodeRune(ls3,-10))//3
	fmt.Println(ls)//[97 0 0 0 0 0 0 0 0 0]
	fmt.Println(ls1)//[228 184 173 0 0 0 0 0 0 0]
	fmt.Println(ls2)//[98 0 0 0 0 0 0 0 0 0]
	fmt.Println(ls3)//[239 191 189 0 0 0 0 0 0 0]

	// RuneLen返回编码符文所需的字节数。
	//如果该符文不是要在UTF-8中编码的有效值，则返回-1。
	fmt.Println(utf8.RuneLen('b'))//1
	fmt.Println(utf8.RuneLen('中'))//3
	fmt.Println(utf8.RuneLen('a'))//1
	fmt.Println(utf8.RuneLen(-10))//-1

	// RuneCount返回p中的runes数。 错误和短编码被视为宽度为1字节的单个符文。
	fmt.Println(utf8.RuneCount([]byte{'a','b','c'}))//3
	fmt.Println(utf8.RuneCount([]byte{'a','b',255}))//3
	fmt.Println(utf8.RuneCount([]byte{'a','b',0}))//3
	//fmt.Println(utf8.RuneCount([]byte{'a','b',-1}))//3，报错，不能为负数
	//fmt.Println(utf8.RuneCount([]byte{'a','b',256}))//3,报错，最大255

	//跟上面的差不多，不过输入的是string
	fmt.Println(utf8.RuneCountInString("abc"))//3
	fmt.Println(utf8.RuneCountInString("abc中"))//4
	fmt.Println(utf8.RuneCountInString("abc中"))//4
	fmt.Println(utf8.RuneCountInString("中国"))//2
	fmt.Println(utf8.RuneCountInString("-+"))//2
	fmt.Println(utf8.RuneCountInString(" "))//1

	// RuneStart报告该字节是否可能是编码的（可能是无效的）符文的第一个字节。
	// 第二个和后续字节始终将高两位设置为10。
	fmt.Println(utf8.RuneStart(97))//true
	fmt.Println(utf8.RuneStart(255))//true
	fmt.Println(utf8.RuneStart(0))//true
	//fmt.Println(utf8.RuneStart(256))//constant 256 overflows byte
	//fmt.Println(utf8.RuneStart(19990))//constant 19990 overflows byte

	// FullRune报告p中的字节是否以符文的完整UTF-8编码开头。
	//无效的编码被视为完整的符文，因为它将转换为宽度为1的错误符文。
	fmt.Println(utf8.FullRune([]byte{'a','b','c'}))//true
	fmt.Println(utf8.FullRune([]byte{'a','b',0xff}))//true
	//fmt.Println(utf8.FullRune([]byte{'a','b',0x100}))//constant 256 overflows byte
	//fmt.Println(utf8.FullRune([]byte{'a','b',0xfff}))//constant 4095 overflows byte

	buf := []byte{228, 184, 150} // '世'的utf8编码
	fmt.Println(utf8.FullRune(buf))//true
	fmt.Println(utf8.FullRune(buf[:2]))//false

	fmt.Println(utf8.FullRuneInString("abc"))//true
	fmt.Println(utf8.FullRuneInString(string([]byte{'a','b',0xff})))//true
	fmt.Println(utf8.FullRuneInString("中国"))//true

	buf1 := []byte{228, 184, 150} // '世'的utf8编码
	fmt.Println(utf8.FullRuneInString(string(buf1)))//true
	str:=string(buf1[:2])
	fmt.Println(utf8.FullRuneInString(str))//false


	// DecodeRune解压缩p中的第一个UTF-8编码，并返回符文rune及其宽度（以字节为单位）。 如果p为空，则返回（RuneError，0）。
	// 否则，如果编码无效，则返回（RuneError，1）。 对于正确的非空UTF-8，这都是不可能的结果。
	//
	//如果编码不正确，则编码无效；如果编码不正确，则编码超出范围，或者不是该值的最短UTF-8编码。 不执行其他任何验证。
	fmt.Println(utf8.DecodeRune([]byte{'a','b','c'}))//97 1
	fmt.Println(utf8.DecodeRune([]byte{'a','b',0xff}))//97 1
	fmt.Println(utf8.DecodeRune([]byte{0xff,'a','b'}))//65533 1
	fmt.Println(utf8.DecodeRune([]byte{10,'a','b'}))//10 1
	//fmt.Println(utf8.DecodeRune([]byte{-10,'a','b'}))// constant -10 overflows byte


	fmt.Println(utf8.DecodeLastRune([]byte{10,'a','b'}))//98 1
	fmt.Println(utf8.DecodeLastRune([]byte{'a','b',10}))//10 1
	fmt.Println(utf8.DecodeLastRune([]byte{'a','b',0xff}))//65533 1（这里其实是表示无效的编码）

	fmt.Println("***********************************s")

	fmt.Println(utf8.FullRuneInString(str))//false
	fmt.Println(utf8.DecodeRuneInString("abc"))//97 1
	fmt.Println(utf8.DecodeRuneInString("中abc"))//20013 3
	fmt.Println(utf8.DecodeRuneInString(""))//65533 0（0这里其实说是表示空字符串，DecodeRune()同理）
	buf2 := []byte{228, 184, 150} // '世'的utf8编码
	str2:=string(buf2[:2])
	fmt.Println(utf8.DecodeRuneInString(str2))//65533 1（1这里其实是表示无效的编码，DecodeRune()同理）
	fmt.Println(utf8.DecodeRuneInString(string(0xff)))//255 2
	fmt.Println(utf8.DecodeRuneInString(string(0xfffff)))//1048575 4
	fmt.Println(utf8.DecodeRuneInString(string(0xffffff)))//65533 3


	fmt.Println(utf8.DecodeLastRuneInString("abc"))//99 1
	fmt.Println(utf8.DecodeLastRuneInString("abc中"))//20013 3
	fmt.Println(utf8.DecodeLastRuneInString(string(0xfffff)))//1048575 4
	fmt.Println(utf8.DecodeLastRuneInString(""))//65533 0
	fmt.Println(utf8.DecodeLastRuneInString("\\"))//92 1


	fmt.Println(utf8.RuneError)//65533
	fmt.Println(utf8.RuneSelf)//128
	fmt.Println(utf8.MaxRune)//1114111
	fmt.Println(utf8.UTFMax)//4


	buf3 := []byte("Hello, 世界")
	fmt.Println("bytes =", len(buf3))//13
	fmt.Println("runes =", utf8.RuneCount(buf3))//9


	buf4 := []byte("a界")
	fmt.Println(buf4)//[97 231 149 140]
	fmt.Println(utf8.RuneStart(buf4[0]))//true
	fmt.Println(utf8.RuneStart(buf4[1]))//true
	fmt.Println(utf8.RuneStart(buf4[2]))//false



	b := []byte("Hello, 世界")

	for len(b) > 0 {
		r, size := utf8.DecodeLastRune(b)
		fmt.Printf("%c %v\n", r, size)

		b = b[:len(b)-size]
	}
	//输出：
	//	界 3
	//	世 3
	//	  1
	//	, 1
	//	o 1
	//	l 1
	//	l 1
	//	e 1
	//	H 1

}
