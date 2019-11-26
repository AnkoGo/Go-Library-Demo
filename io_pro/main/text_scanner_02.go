package main

import (
	"fmt"
	"strings"
	"text/scanner"
	"unicode"
)




func main() {


	fmt.Println("-----------scanner.IsIdentRune()-------------")

	Example_isIdentRune()

	fmt.Println("-----------scanner.Mode-------------")
	Example_mode()

	fmt.Println("-----------scanner.whitespace-------------")
	Example_whitespace()
}

func Example_isIdentRune() {
	const src = "%var1 var2%"

	var s scanner.Scanner
	s.Init(strings.NewReader(src))
	s.Filename = "default"

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("%s: %s\n", s.Position, s.TokenText())
	}

	fmt.Println()
	//将上面已经设置好的s初始化（跟重置状态信息差不多）
	s.Init(strings.NewReader(src))
	s.Filename = "percent"

	// treat leading '%' as part of an identifier	（//将前导'％'视为标识符的一部分）
	//Ident是Identify（识别）的缩写，IsIdentRune判断是否是rune
	s.IsIdentRune = func(ch rune, i int) bool {
		// IsLetter reports whether the rune is a letter (category L).
		// IsLetter报告该符文是否为字母（类别L）。

		// IsDigit reports whether the rune is a decimal digit.
		// IsDigit报告该符文是否为十进制数字。
		//我们把字母，数字和索引为0时候的%号（索引不为0的%号不识别为rune）都认为是rune
		return ch == '%' && i == 0 || unicode.IsLetter(ch) || unicode.IsDigit(ch) && i > 0
	}

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("%s: %s\n", s.Position, s.TokenText())
	}

	// Output:
	// default:1:1: %
	// default:1:2: var1
	// default:1:7: var2
	// default:1:11: %
	//
	// percent:1:1: %var1
	// percent:1:7: var2
	// percent:1:11: %
}


func Example_mode() {
	const src = `
    // Comment begins at column 5.

This line should not be included in the output.

/*
This multiline comment
should be extracted in
its entirety.
*/
hello world!
`

	var s scanner.Scanner
	s.Init(strings.NewReader(src))
	s.Filename = "comments"
	s.Mode ^= scanner.SkipComments // don't skip comments	（//不要跳过注释，默认不会识别注释）

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		txt := s.TokenText()
		// HasPrefix tests whether the string s begins with prefix.
		// HasPrefix测试字符串s是否以前缀开头。
		if strings.HasPrefix(txt, "//") || strings.HasPrefix(txt, "/*") {
			fmt.Printf("（注释）%s: %s\n", s.Position, txt)
		}else {
			fmt.Println(txt)
		}

	}

	// Output:
	//	（注释）comments:2:5: // Comment begins at column 5.
	//	This
	//	line
	//	should
	//	not
	//	be
	//	included
	//	in
	//	the
	//	output
	//	.
	//	（注释）comments:6:1: /*
	//	This multiline comment
	//	should be extracted in
	//	its entirety.
	//	*/
	//	hello
	//	world
	//	!
}



func Example_whitespace() {
	// tab-separated values	（//制表符分隔的值）
	const src = `aa	ab	ac	ad
ba	bb	bc	bd
ca	cb	cc	cd
da	db	dc	dd`

	var (
		col, row int
		s        scanner.Scanner
		tsv      [4][4]string // large enough for example above	（//足够大，例如以上）
	)
	s.Init(strings.NewReader(src))
	// The Whitespace field controls which characters are recognized
	// as white space. To recognize a character ch <= ' ' as white space,
	// set the ch'th bit in Whitespace (the Scanner's behavior is undefined
	// for values ch > ' '). The field may be changed at any time.
	//“空白”字段控制将哪些字符识别为空白。 要将字符ch <=' '识别为空白，请在Whitespace属性中设置第ch位（对于值ch>' '，
	// 扫描仪Scanner的行为未定义）。 该字段可以随时更改。
	// 注意下面的表现形式
	s.Whitespace ^= 1<<'\t' | 1<<'\n' // don't skip tabs and new lines	（//不要跳过制表符和换行符，也就是识别他们）

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		switch tok {
		case '\n':
			row++
			col = 0
		case '\t':
			col++
		default:
			tsv[row][col] = s.TokenText()
		}
	}

	fmt.Print(tsv)

	// Output:
	// [[aa ab ac ad] [ba bb bc bd] [ca cb cc cd] [da db dc dd]]
}





func check_err(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

}
