package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main11122() {
	fmt.Println("你好啊", strconv.ErrRange) //sdsdsd
	//a := make([]byte, 10)
	//fmt.Println(strconv.AppendInt(a,1,10))
	a:=rune(97)
	fmt.Println(strconv.IsPrint(a))//里面的接收值类型只能是rune
	fmt.Println(strconv.IsPrint(97))//可打印的类型包括：字母（广义）、数字、标点、符号、ASCII空格。
	fmt.Println(strconv.IsPrint(97))
	fmt.Println("-----------------------------")
	fmt.Println(strconv.CanBackquote("'"))//判断是否是单引号，但不是反单引号
	// Quote返回带双引号的Go字符串文字，表示s。 返回的字符串使用Go转义序列
	// （\t，\n，\xFF，\u0100）来控制字符和IsPrint定义的不可打印字符。
	fmt.Println(strconv.Quote("aaa"))
	fmt.Println(strconv.Quote("'sdsd`"))
	fmt.Println(strconv.Quote("\t\n\xFF"))//进行转义
	fmt.Println(strconv.Quote("\u0100"))//进行转义
	fmt.Println(strconv.Quote("\u0097"))//进行转义
	fmt.Println(strconv.Quote(" "))
	fmt.Println(strconv.Quote("中"))


	fmt.Println("-----------------------------")
	// QuoteToASCII返回表示s的双引号Go字符串文字。
	//非ASCII字符和不可打印的字符会使用Go转义序列（\ t，\ n，\ xFF，\ u0100）。
	fmt.Println(strconv.QuoteToASCII("中"))
	fmt.Println(strconv.QuoteToASCII("\t\n\xFF"))
	fmt.Println(strconv.QuoteToASCII("a"))
	fmt.Println(strconv.QuoteToASCII(" "))
	fmt.Println(strconv.Quote("\u0100"))//进行转义
	fmt.Println(strconv.Quote("\u0097"))//进行转义

	fmt.Println("-----------------------------")
	//b:=rune(98)
	//b:=rune('t')
	//b:=rune('中')
	//b:=rune('\\')
	//b:=rune('\\\')//报错
	//b:=rune(100000)//'𘚠'
	b:=rune(19990)//'世'
	//b:=rune(1999000000)//'�',
	//b:=rune(199900000000)//报错，最多10位数字，也就是最多31位的二进制
	fmt.Println(strconv.QuoteRune(b))


	fmt.Println("-----------------------------")
	//c:=rune(19990)//'\u4e16',只能打印出ascii
	//c:=rune(19)//'\x13',只能打印出ascii
	//c:=rune(97)//'a',只能打印出ascii
	//c:=rune('中')//'\u4e2d'
	//c:=rune(' ')//' '
	//c:=rune('\\')//'\\'
	//c:=rune('\\\')//报错
	//c:=rune(1999000000)//'\ufffd'
	c:=rune('\t')//'\t'
	//c:=rune(19990000000)//报错，最多10位数字，也就是最多31位的二进制
	fmt.Println(strconv.QuoteRuneToASCII(c))

	fmt.Println("-----------------------------")
	// QuoteRuneToGraphic返回表示符文的单引号Go字符文字。
	// 返回的字符串对IsGraphic定义的非ASCII字符和不可打印的字符使用Go转义序列（\t，\n，\xFF，\u0100）。
	//d:=rune(1999000000)//'�'
	//d:=rune(97)//'a'
	//d:=rune('t')//'t'
	//d:=rune('\\')//'\\'
	//d:=rune(' ')//' '
	//d:=rune('\n')//'\t'
	//d:=rune('\n')//'\n'
	//d:=rune('\xFF')//'ÿ'
	//d:=rune('\u0100')//'Ā'
	//d:=rune('\u0097')//'\u0097'
	d:=rune('中')//'中',说白了就是对能编码的unicode值进行转码显示出来
	fmt.Println(strconv.QuoteRuneToGraphic(d))

	fmt.Println("-----------------------------")
	// Unquote将s解释为单引号，双引号或反引号的Go字符串文字，并返回s引用的字符串值。
	// （如果s用单引号引起来，它将是Go字符文字； Unquote返回相应的单字符字符串。）

	fmt.Println(strconv.Unquote("`var`"))
	fmt.Println(strconv.Unquote("`fmt.Println(d)`"))
	fmt.Println(strconv.Unquote("`2*2`"))
	fmt.Println(strconv.Unquote("`a\tb`"))
	fmt.Println(strconv.Unquote("`a\nb`"))
	fmt.Println(strconv.Unquote("`a\xFFb`"))
	fmt.Println(strconv.Unquote("d"))
	fmt.Println(strconv.Unquote(`d`))
	fmt.Println(strconv.Unquote(string('d')))
	//输出如下：
	//  var <nil>
	//	fmt.Println(d) <nil>
	//	2*2 <nil>
	//	a	b <nil>
	//	a
	//  b <nil>
	//	a�b <nil>
	//  invalid syntax
	//  invalid syntax
	//  invalid syntax

	fmt.Println("-----------------------------")

	// IsGraphic报告是否通过Unicode将符文定义为图形。 这些字符包括字母，标记，数字，标点符号，
	// 符号和空格，来自类别L，M，N，P，S和Zs。
	//f:=rune(97)//true
	//f:=rune(19990)//true
	//f:=rune(19900)//false...unicode表没有这个值
	//f:=rune(1990000)//false...unicode表没有这个值
	//f:=rune(199000)//false...unicode表没有这个值
	//f:=rune(199900)//false...unicode表没有这个值。一般不大于5W好像，且不小于32
	//f:=rune(30)//false...unicode表没有这个值。一般不大于5W好像，且不小于32
	//f:=rune(31)//false...unicode表没有这个值。一般不大于5W好像，且不小于32
	//f:=rune(32)//true
	//f:=rune('\t')//false
	//f:=rune('\n')//false
	//f:=rune('\xFF')//true
	f:=rune(' ')//true
	fmt.Println(strconv.IsGraphic(f))//判断是否unicode表里面有相应的值


	//函数假设s是一个表示字符的go语法字符串，解析它并返回四个值：
	//
	//1) value，表示一个rune值或者一个byte值
	//2) multibyte，表示value是否是一个多字节的utf-8字符
	//3) tail，表示字符串剩余的部分
	//4) err，表示可能存在的语法错误
	//quote参数为单引号时，函数认为单引号是语法字符，不接受未转义的单引号；双引号时，
    //函数认为双引号是语法字符，不接受未转义的双引号；如果是零值，函数把单引号和双引号当成普通字符。
	fmt.Println("-----------------------------")
	var g byte ='"'
	//var g byte ='\''
	//var g byte ='`'
	//var g byte ='i'//输入其他值会默认是零值，函数把单引号和双引号当成普通字符。
	//var g byte =97
	//var g byte ='中'// 报错,constant 20013 overflows byte
	fmt.Println(g)//34
	fmt.Printf("%T---%c\n",g,g)//uint8---"
	fmt.Println(strconv.UnquoteChar("a",g))//97 false  <nil>,判断a的字符串形式是否是utf8多字节字符
	fmt.Println(strconv.UnquoteChar("\t",g))//9 false  <nil>
	fmt.Println(strconv.UnquoteChar("中",g))//20013 true  <nil>
	fmt.Println(strconv.UnquoteChar("中国",g))//20013 true 国 <nil>,剩余一个国字
	fmt.Println(strconv.UnquoteChar("中国人们是最棒的！！",g))//20013 true 国人们是最棒的！！ <nil>

	fmt.Println("-----------------------------")
	//返回字符r在go语法下的单引号字面值表示，控制字符、
	//不可打印字符、非ASCII字符会进行转义。
	fmt.Println(strconv.QuoteRuneToASCII('中'))//'\u4e2d'
	fmt.Println(strconv.QuoteRuneToASCII('a'))//'a'
	fmt.Println(strconv.QuoteRuneToASCII('\t'))//'\t'
	fmt.Println(strconv.QuoteRuneToASCII('\n'))//'\n'
	fmt.Println(strconv.QuoteRuneToASCII('\xFF'))//'\u00ff'

	fmt.Println("-----------------------------")
	// AppendQuote将由Quote生成的表示s的双引号Go
	// 字符串文字追加到dst字节切片并返回扩展缓冲区字节切片中去。
	ls:=make([]byte,10)
	//var aabyte byte =34
	//var aabyte byte ='"'
	//var aabyte rune =39
	var aabyte rune ='\''
	fmt.Println(aabyte)
	fmt.Println(string(aabyte))
	//返回字符r在go语法下的单引号字面值表示，控制字符、不可打印字符会进行转义。（如\t，\n，\xFF，\u0100）并且会追加到字节切片后
	fmt.Println(strconv.AppendQuote(ls,"abcd"))//[0 0 0 0 0 0 0 0 0 0 34 97 98 99 100 34]
	fmt.Println(strconv.AppendQuoteRune(ls,'a'))//[0 0 0 0 0 0 0 0 0 0 39 97 39]
	fmt.Println(strconv.AppendQuoteRune(ls,'中'))//[0 0 0 0 0 0 0 0 0 0 39 228 184 173 39]
	//返回字符r在go语法下的单引号字面值表示，控制字符、不可打印字符、非ASCII字符会进行转义。并且会追加到字节切片后
	fmt.Println(strconv.AppendQuoteRuneToASCII(ls,'中'))//[0 0 0 0 0 0 0 0 0 0 39 92 117 52 101 50 100 39]，非ascii转义成unicode值
	fmt.Println(strconv.AppendQuoteRuneToASCII(ls,'a'))//[0 0 0 0 0 0 0 0 0 0 39 97 39]


	fmt.Println("-----------------------------")
	fmt.Println(strconv.AppendQuoteRuneToGraphic(ls,'a'))//[0 0 0 0 0 0 0 0 0 0 39 97 39],utf8值
	fmt.Println(strconv.AppendQuoteRuneToGraphic(ls,'中'))//[0 0 0 0 0 0 0 0 0 0 39 228 184 173 39],utf8值
	fmt.Println(strconv.AppendQuoteRuneToGraphic(ls,19990))//[0 0 0 0 0 0 0 0 0 0 39 228 184 150 39],utf8值
	fmt.Println(strconv.AppendQuoteRuneToGraphic(ls,19900))//[0 0 0 0 0 0 0 0 0 0 39 92 117 52 100 98 99 39],unicode值
	g2:=strconv.AppendQuoteRuneToGraphic(ls,19900)
	fmt.Println(strconv.IsGraphic(rune(g2[11])))//true

	fmt.Println("-----------------------------")
	fmt.Println(strconv.AppendBool(ls,true))// [0 0 0 0 0 0 0 0 0 0 116 114 117 101]
	fmt.Println(strconv.AppendBool(ls,false))//[0 0 0 0 0 0 0 0 0 0 102 97 108 115 101]

	//utf8编码前瞻
	ss1:=make([]byte,10)
	//n:=utf8.EncodeRune(sl,'中')
	//n:=utf8.EncodeRune(sl,'a')
	//n:=utf8.EncodeRune(sl,19990)//[228 184 150 0 0 0 0 0 0 0]
	n1:=utf8.EncodeRune(ss1,19900)//[228 182 188 0 0 0 0 0 0 0]
	fmt.Println(n1)
	fmt.Println(ss1)

	fmt.Println("-----------------------------")
	ss:=make([]byte,10)
	ss2:=strconv.AppendInt(ss,16,16)//[0 0 0 0 0 0 0 0 0 0 49 48]
	fmt.Println(ss2)
	fmt.Println(strconv.IsGraphic(rune(ss2[10])))//true

	fmt.Println(strconv.FormatInt(16,16))//10,返回16的16进制表示的字符串形式
	fmt.Println(strconv.FormatInt(16,8))//20
	fmt.Println(strconv.FormatInt(16,2))//10000
	fmt.Println(strconv.FormatInt(16,10))//10000


	fmt.Println("-----------------------------")
	fmt.Println(strconv.FormatBool(true))//true

	fmt.Println(strconv.FormatBool(false))//	false

	fmt.Println("-----------------------------")
	fmt.Println(strconv.FormatFloat(16.3456,'f',2,32))//16.35
	fmt.Println(strconv.FormatFloat(16.3456,'f',2,64))//16.35
	fmt.Println(strconv.FormatFloat(16.3456,'f',4,64))//16.3456
	fmt.Println(strconv.FormatFloat(16.3456,'f',6,64))//16.345600
	fmt.Println(strconv.FormatFloat(16.3456789,'b',2,64))//4600899587697361p-48
	fmt.Println(strconv.FormatFloat(160.3456789,'e',2,64))//1.60e+02
	fmt.Println(strconv.FormatFloat(160.3456789,'e',4,64))//1.6035e+02
	fmt.Println(strconv.FormatFloat(160.3456789,'E',4,64))//1.6035E+02
	fmt.Println(strconv.FormatFloat(160.3456789,'g',4,64))//160.3,此时4是总位数
	fmt.Println(strconv.FormatFloat(160345678923.235,'g',4,64))//1.603e+11,此时4是总位数


	fmt.Println("-----------------------------")
	fmt.Println(strconv.FormatUint(16,10))//16
	fmt.Println(strconv.FormatInt(-16,10))//-16
	//fmt.Println(strconv.FormatUint(-16,10))//报错，constant -16 overflows uint64

	fmt.Println("-----------------------------")
	//Atoi是ParseInt(s, 10, 0)的简写。
	fmt.Println(strconv.Atoi("13"))
	//fmt.Println(strconv.Atoi("a"))//报错，必须数字字面量的字符串

	fmt.Println("-----------------------------")
	//Itoa是FormatInt(i, 10) 的简写。
	fmt.Println(strconv.Itoa(12))//12,字符串的12

	fmt.Println("-----------------------------")
	fmt.Println(strconv.ParseBool("true"))//false <nil>
	fmt.Println(strconv.ParseBool("false"))//true <nil>
	fmt.Println(strconv.ParseBool("false111"))//false strconv.ParseBool: parsing "false111": invalid syntax

	fmt.Println("-----------------------------")
	fmt.Println(strconv.ParseFloat("12.45",64))//12.45 <nil>
	fmt.Println(strconv.ParseFloat("12.45",32))//12.449999809265137 <nil>
	fmt.Println(strconv.ParseFloat("12.10",32))//12.100000381469727 <nil>
	fmt.Println(strconv.ParseFloat("12.6",32))//12.600000381469727 <nil>
	//fmt.Println(strconv.ParseFloat("12.6aaa",32))//0 strconv.ParseFloat: parsing "12.6aaa": invalid syntax


	fmt.Println("-----------------------------")
	fmt.Println(strconv.ParseInt("10",16,64))//16 <nil>,这里的16表示我写的10是16进制的10，而不是10进制的10大小
	fmt.Println(strconv.ParseInt("10",16,32))//16 <nil>
	fmt.Println(strconv.ParseInt("-10",16,32))//-16 <nil>
	// bitSize参数指定结果必须适合的整数类型。 位大小0、8、16、32和64对应于int，int8，int16，int32和int64。
	//如果bitSize小于0或大于64，则返回错误。
	fmt.Println(strconv.ParseInt("-10",16,0))//-16 <nil>

	fmt.Println("-----------------------------")
	fmt.Println(strconv.ParseUint("10",16,32))//16 <nil>
	fmt.Println(strconv.ParseUint("-10",16,32))//0 strconv.ParseUint: parsing "-10": invalid syntax


	fmt.Println("-----------------------------")
	// QuoteToGraphic返回表示s的双引号Go字符串文字。
	//返回的字符串对IsGraphic定义的非ASCII字符和不可打印的字符使用Go转义序列（\ t，\ n，\ xFF，\ u0100）。
	fmt.Println(strconv.QuoteToGraphic("aa"))//"aa"
	fmt.Println(strconv.QuoteToGraphic("`aa`"))//"`aa`"
	fmt.Println(strconv.QuoteToGraphic(`"aa"`))//"\"aa\""
	fmt.Println(strconv.QuoteToGraphic(`中`))//"中"
	fmt.Println(strconv.QuoteToGraphic(`a\tb`))//"a\\tb",不可打印的东西会转义
	fmt.Println(strconv.QuoteToGraphic(`a\nb`))//"a\\nb"
	fmt.Println(strconv.QuoteToGraphic(`a\u0100b`))//"a\\u0100b"
	fmt.Println(strconv.QuoteToGraphic(`a b`))//"a b"
	fmt.Println(strconv.QuoteToGraphic(`\`))//"\\"
	fmt.Println(strconv.QuoteToGraphic(`\\`))//"\\\\"
	fmt.Println(strconv.IntSize)//64






}

























