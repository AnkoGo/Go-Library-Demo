package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println("**********************************")
	//函数报告r是否在rangeTable指定的字符范围内。
	fmt.Println(unicode.Is(unicode.Han,'中')) // true
	fmt.Println(unicode.Is(unicode.Han,'z')) // false

	//判断是否是十六进制
	fmt.Println(unicode.Is(unicode.ASCII_Hex_Digit,'z')) // false
	fmt.Println(unicode.Is(unicode.ASCII_Hex_Digit,'a')) // true
	fmt.Println(unicode.Is(unicode.ASCII_Hex_Digit,0x7F)) // false
	fmt.Println(unicode.Is(unicode.ASCII_Hex_Digit,'0')) // true
	fmt.Println(unicode.Is(unicode.ASCII_Hex_Digit,'1')) // true
	fmt.Println(unicode.Is(unicode.ASCII_Hex_Digit,'b')) // true
	fmt.Println(unicode.Is(unicode.ASCII_Hex_Digit,'e')) // true
	fmt.Println(unicode.Is(unicode.ASCII_Hex_Digit,'f')) // true
	fmt.Println(unicode.Is(unicode.ASCII_Hex_Digit,'g')) // false

	fmt.Printf("%x\n",98) // 62
	fmt.Printf("%x\n",127) // 7f

	fmt.Println(unicode.IsUpper('A'))//true
	fmt.Println(unicode.IsUpper('a'))//false

	fmt.Println(unicode.IsLower('a'))//true
	fmt.Println(unicode.IsLower('A'))//false

	// 判断字符 r 是否为 Unicode 规定的 Title 字符
	// 大部分字符的 Title 格式就是其大写格式
	// 只有少数字符的 Title 格式是特殊字符
	// 这里判断的就是特殊字符
	fmt.Println(unicode.IsTitle('A'))//false
	fmt.Println(unicode.IsTitle('a'))//false
	fmt.Println(unicode.IsTitle('中'))//false
	fmt.Println(unicode.IsTitle('ᾏ'))//true
	fmt.Println(unicode.IsTitle('ᾟ'))//true
	fmt.Println(unicode.IsTitle('ᾯ'))//true


	fmt.Println(unicode.IsDigit('0'))//true
	fmt.Println(unicode.IsDigit(0))//false
	fmt.Println(unicode.IsDigit('9'))//true
	fmt.Println(unicode.IsDigit('f'))//false
	fmt.Println(unicode.IsDigit('e'))//false


	// IsPrint 判断字符 r 是否为 Go 所定义的“可打印字符”
	// “可打印字符”包括字母、标记、数字、标点、符号和 ASCII 空格
	// 他们分别对应于 L, M, N, P, S 类别和 ASCII 空格
	// “可打印字符”和“图形字符”基本是相同的，不同之处在于
	// “可打印字符”只包含 Zs 类别中的 ASCII 空格（U+0020）
	fmt.Println(unicode.IsPrint('e'))//true
	fmt.Println(unicode.IsPrint(' '))//true
	fmt.Println(unicode.IsPrint(127))//false
	fmt.Println(unicode.IsPrint(126))//true

	// IsLetter 判断 r 是否为一个字母字符 (类别 L)
	// 汉字也是一个字母字符
	fmt.Println(unicode.IsLetter(0xFF))//true
	fmt.Println(unicode.IsLetter(0x1FF))//true
	fmt.Println(unicode.IsLetter('0'))//false
	fmt.Println(unicode.IsLetter('1'))//false
	fmt.Println(unicode.IsLetter('a'))//true
	fmt.Println(unicode.IsLetter('中'))//true
	fmt.Println(unicode.IsLetter('中'))//true

	//判断是否是一个数字
	fmt.Println(unicode.IsNumber('中'))//false
	fmt.Println(unicode.IsNumber('a'))//false
	fmt.Println(unicode.IsNumber('0'))//true
	fmt.Println(unicode.IsNumber('5'))//true

	// IsGraphic 判断字符 r 是否为一个“图形字符”
	// “图形字符”包括字母、标记、数字、标点、符号、空格
	// 他们分别对应于 L、M、N、P、S、Zs 类别
	// 这些类别是 RangeTable 类型，存储了相应类别的字符范围
	fmt.Println(unicode.IsGraphic('5'))//true

	s := "Hello　世界！\t"
	for _, r := range s {
		fmt.Printf("%c = %v\n", r, unicode.IsGraphic(r))
	}
	//输出如下：
	//	H = true
	//	e = true
	//	l = true
	//	l = true
	//	o = true
	//	　 = true
	//	世 = true
	//	界 = true
	//	！ = true
	//		 = false


	s1 := "Hello\n\t世界！"
	for _, r := range s1 {
		fmt.Printf("%c = %v\n", r, unicode.IsControl(r))
	}//\n\t才是控制字符
	//输出如下：
	//	H = false
	//	e = false
	//	l = false
	//	l = false
	//	o = false
	//
	//	 = true
	//		 = true
	//	世 = false
	//	界 = false
	//	！ = false

	// IsMark 判断 r 是否为一个 mark 字符 (类别 M)
	s2 := "Hello ៉៊់៌៍！"
	for _, r := range s2 {
		fmt.Printf("%c = %v\n", r, unicode.IsMark(r))
	} // ៉៊់៌៍ = true
	//输出如下：
	//	H = false
	//	e = false
	//	l = false
	//	l = false
	//	o = false
	//	= false
	//	៉ = true
	//	៊ = true
	//	់ = true
	//	៌ = true
	//	៍ = true
	//	！ = false

	//IsOneOf判断rune是否是RangeTable里面的某个集合的范围，这里判断的是否是汉字、标点符号
	tables := []*unicode.RangeTable{unicode.Han, unicode.P}
	fmt.Println(unicode.IsOneOf(tables,'a')) // false
	fmt.Println(unicode.IsOneOf(tables,'中')) // true
	fmt.Println(unicode.IsOneOf(tables,'！')) // true
	fmt.Println(unicode.IsOneOf(tables,'!')) // true

	// IsSpace报告该符文是否为Unicode的White Space属性所定义的空格字符； 在Latin-1空间
	// 这是
	//'\t'，'\n'，'\v'，'\f'，'\r'，''，U+0085（NEL），U+00A0（NBSP）。
	//间隔字符的其他定义由类别Z和属性Pattern_White_Space设置。
	fmt.Println(unicode.IsSpace('\t')) // true
	fmt.Println(unicode.IsSpace('\n')) // true
	fmt.Println(unicode.IsSpace(' ')) // true
	fmt.Println(unicode.IsSpace(0x85)) // true
	fmt.Println(unicode.IsSpace(0xA0)) // true
	fmt.Println(unicode.IsSpace('0')) // false
	fmt.Println(unicode.IsSpace('a')) // false
	fmt.Println(unicode.IsSpace('\\')) // false


	// IsSymbol报告该符文是否为符号字符。（运算符可能）
	s3 := "Hello (<世=界>)"
	for _, r := range s3 {
		fmt.Printf("%c = %v\n", r, unicode.IsSymbol(r))
	} // <=> = true
	//输出：
	//	H = false
	//	e = false
	//	l = false
	//	l = false
	//	o = false
	//	  = false
	//	( = false
	//	< = true
	//	世 = false
	//	= = true
	//	界 = false
	//	> = true
	//	) = false


	// IsPunct报告该符文是否为Unicode标点字符（类别P）。
	s4 := "Hello 世界！!"
	for _, r := range s4 {
		fmt.Printf("%c = %v\n", r, unicode.IsPunct(r))
	} // ！! = true
	//输出如下：
	//	H = false
	//	e = false
	//	l = false
	//	l = false
	//	o = false
	//	  = false
	//	世 = false
	//	界 = false
	//	！ = true
	//	! = true

	fmt.Println("==================================")
	// ToUpper将符文映射为大写。
	fmt.Println(unicode.ToUpper('a'))//65(A)
	fmt.Println(unicode.ToUpper('A'))//65(A)

	fmt.Println(unicode.ToTitle('a'))//65(A)
	fmt.Println(unicode.ToTitle('A'))//65(A)

	fmt.Println(unicode.ToLower('A'))//97(a)
	fmt.Println(unicode.ToLower('a'))//97(a)

	fmt.Println("==================================")
	//其实上面都是调用了下面的这个to函数的
	fmt.Println(unicode.To(unicode.LowerCase,'A'))//97(a)
	fmt.Println(unicode.To(unicode.LowerCase,'a'))//97(a)

	fmt.Println(unicode.To(unicode.UpperCase,'a'))//65(A)
	fmt.Println(unicode.To(unicode.UpperCase,'A'))//65(A)

	fmt.Println(unicode.To(unicode.TitleCase,'A'))//65(A)
	fmt.Println(unicode.To(unicode.TitleCase,'a'))//65(A)

	fmt.Println("==================================")
	//下面仅仅是尝试，用法不对的
	fmt.Println(unicode.To(1,'a'))//97(a)
	fmt.Println(unicode.To(2,'a'))//65(A)
	fmt.Println(unicode.To(3,'a'))//65533
	fmt.Println(unicode.To(4,'a'))//65533
	fmt.Println(unicode.To(0,'a'))//65
	fmt.Println(unicode.To(0,'A'))//65
	fmt.Println(unicode.To(-1,'a'))//65533
	fmt.Println(unicode.To(-2,'a'))//65533
	fmt.Println(unicode.To(-3,'a'))//65533
	fmt.Println(unicode.To(10,'a'))//65533


	fmt.Println("==================================")

	fmt.Println(unicode.TitleCase)//2
	fmt.Println(unicode.UpperCase)//0
	fmt.Println(unicode.LowerCase)//1

	fmt.Println("==================================")
	//各种编码的最大unicode值
	fmt.Println(unicode.MaxASCII)//127
	fmt.Println(unicode.MaxLatin1)//255
	fmt.Println(unicode.MaxRune)//1114111
	fmt.Println(unicode.MaxCase)//3

	fmt.Println("==================================")
	//底层是通过is来实现的
	fmt.Println(unicode.In('a',unicode.Han,unicode.Upper))//false
	fmt.Println(unicode.In('A',unicode.Han,unicode.Upper))//true
	fmt.Println(unicode.In('Z',unicode.Han,unicode.Upper))//true
	fmt.Println(unicode.In('中',unicode.Han,unicode.Upper))//true
	//Other_Uppercase不知道是什么大写
	fmt.Println(unicode.In('中',unicode.Han,unicode.Other_Uppercase))//true
	fmt.Println(unicode.In('国',unicode.Han,unicode.Other_Uppercase))//true
	fmt.Println(unicode.In('A',unicode.Han,unicode.Other_Uppercase))//false
	fmt.Println(unicode.In('a',unicode.Han,unicode.Other_Uppercase))//false

	fmt.Println(unicode.In('a',unicode.Han,unicode.Lower))//true
	fmt.Println(unicode.In('A',unicode.Han,unicode.Lower))//false
	//Other_Uppercase不知道是什么小写
	fmt.Println(unicode.In('A',unicode.Han,unicode.Other_Lowercase))//false
	fmt.Println(unicode.In('a',unicode.Han,unicode.Other_Lowercase))//false
	fmt.Println(unicode.In('中',unicode.Han,unicode.Other_Lowercase))//true
	fmt.Println(unicode.In('国',unicode.Han,unicode.Other_Lowercase))//true

	//SimpleFold函数迭代在unicode标准字符映射中互相对应的unicode码值。在与r对应的码值中（包括r自身），会返回最小的那个大于r的字符（如果有）；否则返回映射中最小的字符。
	//举例：
	//SimpleFold('A') = 'a'
	//SimpleFold('a') = 'A'
	//SimpleFold('K') = 'k'
	//SimpleFold('k') = '\u212A' (Kelvin symbol, K)
	//SimpleFold('\u212A') = 'K'
	//SimpleFold('1') = '1'
	fmt.Println(unicode.SimpleFold('A'))//97('a')
	fmt.Println(unicode.SimpleFold('a'))//65('A')
	fmt.Println(unicode.SimpleFold('K'))//107('k')
	fmt.Println(unicode.SimpleFold('k'))//8490('\u212A')
	fmt.Println(unicode.SimpleFold('\u212A'))//75('K')
	fmt.Println(unicode.SimpleFold('1'))//49('1')
	fmt.Println(unicode.SimpleFold('0'))//48('0')
	fmt.Println(unicode.SimpleFold('2'))//50('2')




}
