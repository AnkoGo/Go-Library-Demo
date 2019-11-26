package main

import (
	"bytes"
	"fmt"
	"text/scanner"
)




func main() {


	fmt.Println("-----------scanner.TokenString()-------------")

	//// Predefined mode bits to control recognition of tokens. For instance,
	//// to configure a Scanner such that it only recognizes (Go) identifiers,
	//// integers, and skips comments, set the Scanner's Mode field to:
	////
	////	ScanIdents | ScanInts | SkipComments
	////
	//// With the exceptions of comments, which are skipped if SkipComments is
	//// set, unrecognized tokens are not ignored. Instead, the scanner simply
	//// returns the respective individual characters (or possibly sub-tokens).
	//// For instance, if the mode is ScanIdents (not ScanStrings), the string
	//// "foo" is scanned as the token sequence '"' Ident '"'.
	////
	//// Use GoTokens to configure the Scanner such that it accepts all Go
	//// literal tokens including Go identifiers. Comments will be skipped.
	////预定义模式位以控制令牌的识别。 例如，要将扫描仪配置为仅识别（Go）标识符，整数并跳过注释，请将扫描仪的模式字段设置为：
	//// ScanIdents | ScanInts | SkipComments
	////除注释（如果设置了SkipComments会跳过注释）外，不会忽略无法识别的标记。 取而代之的是，扫描程序仅返回相应的单个字符（或可能的子令牌）。
	////例如，如果模式为ScanIdents（而不是ScanStrings），则将字符串“ foo”作为标记序列'"' Ident '"'进行扫描。
	////使用GoTokens配置扫描程序，使其可以接受所有Go文字令牌，包括Go标识符。 注释将被跳过。
	//const (
	//	ScanIdents     = 1 << -Ident
	//	ScanInts       = 1 << -Int
	//	ScanFloats     = 1 << -Float // includes Ints and hexadecimal floats	（//包含整数和十六进制浮点数）
	//	ScanChars      = 1 << -Char
	//	ScanStrings    = 1 << -String
	//	ScanRawStrings = 1 << -RawString
	//	ScanComments   = 1 << -Comment
	//	SkipComments   = 1 << -skipComment // if set with ScanComments, comments become white space	（//如果使用ScanComments设置，则注释变为空白）
	//	GoTokens       = ScanIdents | ScanFloats | ScanChars | ScanStrings | ScanRawStrings | ScanComments | SkipComments
	//)


	//// The result of Scan is one of these tokens or a Unicode character.
	////Scan扫描的结果是这些标记tokens或Unicode字符之一。
	//const (
	//	EOF = -(iota + 1)
	//	Ident
	//	Int
	//	Float
	//	Char
	//	String
	//	RawString
	//	Comment
	//
	//	// internal use only(// 限内部使用)
	//	skipComment
	//)
	//
	//var tokenString = map[rune]string{
	//	EOF:       "EOF",
	//	Ident:     "Ident",
	//	Int:       "Int",
	//	Float:     "Float",
	//	Char:      "Char",
	//	String:    "String",
	//	RawString: "RawString",
	//	Comment:   "Comment",
	//}



	// TokenString returns a printable string for a token or Unicode character.
	// TokenString返回令牌token(如上定义的那些)或Unicode字符的可打印字符串。

	//token
	fmt.Println(scanner.TokenString(scanner.EOF))
	fmt.Println(scanner.TokenString(scanner.Ident))
	fmt.Println(scanner.TokenString(scanner.Float))
	fmt.Println(scanner.TokenString(scanner.Char))
	fmt.Println(scanner.TokenString(scanner.String))
	fmt.Println(scanner.TokenString(scanner.RawString))
	fmt.Println(scanner.TokenString(scanner.Comment))

	//Unicode字符
	fmt.Println(scanner.TokenString(97))
	fmt.Println(scanner.TokenString(98))


	//内置有意义的Unicode字符
	fmt.Println(scanner.TokenString(scanner.ScanIdents))
	fmt.Println(scanner.TokenString(scanner.ScanInts))
	fmt.Println(scanner.TokenString(scanner.ScanFloats))
	fmt.Println(scanner.TokenString(scanner.ScanChars))
	fmt.Println(scanner.TokenString(scanner.ScanStrings))
	fmt.Println(scanner.TokenString(scanner.ScanRawStrings))
	fmt.Println(scanner.TokenString(scanner.SkipComments))
	fmt.Println(scanner.TokenString(scanner.GoTokens))


	//输出：
	//	EOF
	//	Ident
	//	Float
	//	Char
	//	String
	//	RawString
	//	Comment
	//	"a"
	//	"b"
	//	"\x04"
	//	"\b"
	//	"\x10"
	//	" "
	//	"@"
	//	"\u0080"
	//	"Ȁ"
	//	"ϴ"


	fmt.Println("-----------scanner.Scanner对象-------------")

	//// A Scanner implements reading of Unicode characters and tokens from an io.Reader.
	////Scanner扫描器实现了从io.Reader中读取Unicode字符和令牌的功能。
	//type Scanner struct {
	//	// Input
	//	src io.Reader
	//
	//	// Source buffer
	//	srcBuf [bufLen + 1]byte // +1 for sentinel for common case of s.next()	(//对于s.next（）的常见情况，标记为+1)
	//	srcPos int              // reading position (srcBuf index)	(//读取位置（srcBuf索引）)
	//	srcEnd int              // source end (srcBuf index)	(//源端（srcBuf索引）)
	//
	//	// Source position	(//源位置)
	//	srcBufOffset int // byte offset of srcBuf[0] in source	(//源中srcBuf [0]的字节偏移)
	//	line         int // line count	(//行数)
	//	column       int // character count		(//字符数)
	//	lastLineLen  int // length of last line in characters (for correct column reporting)	(//最后一行的长度（以字符为单位）（用于正确的列报告）)
	//	lastCharLen  int // length of last character in bytes	(//最后一个字符的长度（以字节为单位）)
	//
	//	// Token text buffer
	//	// Typically, token text is stored completely in srcBuf, but in general
	//	// the token text's head may be buffered in tokBuf while the token text's
	//	// tail is stored in srcBuf.
	//	// Token令牌文本缓冲区
	//	// 通常，令牌文本完全存储在srcBuf中，但是通常令牌文本的头部可以缓存在tokBuf中，而令牌文本的尾部则存储在srcBuf中。
	//	tokBuf bytes.Buffer // token text head that is not in srcBuf anymore	(//不再位于srcBuf中的令牌文本头)
	//	tokPos int          // token text tail position (srcBuf index); valid if >= 0	(//标记文字的尾部位置（srcBuf索引）； >> 0时有效)
	//	tokEnd int          // token text tail end (srcBuf index)	(//标记文字尾部（srcBuf索引）)
	//
	//	// One character look-ahead		(//上一个被Scanner扫描的字符)
	//	ch rune // character before current srcPos	(//当前srcPos之前的字符)
	//
	//	// Error is called for each error encountered. If no Error
	//	// function is set, the error is reported to os.Stderr.
	//	//为遇到的每个错误都调用Error设置的函数。 如果未设置错误功能，则将错误报告给os.Stderr。
	//	Error func(s *Scanner, msg string)
	//
	//	// ErrorCount is incremented by one for each error encountered.
	//	//对于遇到的每个错误，ErrorCount都会加一。
	//	ErrorCount int
	//
	//	// The Mode field controls which tokens are recognized. For instance,
	//	// to recognize Ints, set the ScanInts bit in Mode. The field may be
	//	// changed at any time.
	//	//Mode模式字段控制识别哪些令牌。 例如，要识别整数，请在模式下将 ScanInts位 置1。 该字段可以随时更改。
	//	Mode uint
	//
	//	// The Whitespace field controls which characters are recognized
	//	// as white space. To recognize a character ch <= ' ' as white space,
	//	// set the ch'th bit in Whitespace (the Scanner's behavior is undefined
	//	// for values ch > ' '). The field may be changed at any time.
	//	//“空白”字段控制将哪些字符识别为空白。 要将字符ch <=' '识别为空白，请在空白中设置第ch位（对于值ch>' '，扫描仪的行为未定义）。 该字段可以随时更改。
	//	Whitespace uint64
	//
	//	// IsIdentRune is a predicate controlling the characters accepted
	//	// as the ith rune in an identifier. The set of valid characters
	//	// must not intersect with the set of white space characters.
	//	// If no IsIdentRune function is set, regular Go identifiers are
	//	// accepted instead. The field may be changed at any time.
	//	// IsIdentRune是一个谓词，用于控制在标识符中作为第i个符文接受的字符。 有效字符集不得与空白字符集相交。
	//	// 如果未设置IsIdentRune函数，则将接受常规的Go标识符。 该字段可以随时更改。
	//	IsIdentRune func(ch rune, i int) bool
	//
	//	// Start position of most recently scanned token; set by Scan.
	//	// Calling Init or Next invalidates the position (Line == 0).
	//	// The Filename field is always left untouched by the Scanner.
	//	// If an error is reported (via Error) and Position is invalid,
	//	// the scanner is not inside a token. Call Pos to obtain an error
	//	// position in that case, or to obtain the position immediately
	//	// after the most recently scanned token.
	//	//最近扫描的令牌的起始位置； 通过扫描设置。
	//	//调用Init或Next使位置无效（线== 0）。
	//	//扫描程序始终保持“文件名”字段不变。
	//	//如果报告了一个错误（通过Error）并且Position无效，则表明扫描仪不在令牌内。 在这种情况下，请调用Pos以获取错误位置，或获取最近扫描的令牌之后的位置。
	//	Position
	//}
	//	// A source position is represented by a Position value.
	//	// A position is valid if Line > 0.
	//	//源位置由位置值Position value表示。
	//	//如果Line> 0，则位置有效。
	//	type Position struct {
	//		Filename string // filename, if any	(//文件名（如果有）)
	//		Offset   int    // byte offset, starting at 0	(//字节偏移量，从0开始)
	//		Line     int    // line number, starting at 1	(//行号，从1开始)
	//		Column   int    // column number, starting at 1 (character count per line)	(//列号，从1开始（每行的字符数）)
	//	}


	var s scanner.Scanner
	// Init initializes a Scanner with a new source and returns s.
	// Error is set to nil, ErrorCount is set to 0, Mode is set to GoTokens,
	// and Whitespace is set to GoWhitespace.
	// Init使用新的源初始化Scanner并返回s。
	// Error设置为nil，ErrorCount设置为0，Mode设置为GoTokens，空白设置为GoWhitespace。
	// 返回值接收不接受都可以，因为他是操作对象的原本

	//语句1
	buffer := bytes.NewBufferString("hello world")
	//语句2
	//	fmt.Println(len(`
	//// This is scanned code.
	//`))
	//	buffer := bytes.NewBufferString(`
	//// This is scanned code.
	//if a > 10 {
	//	someParsable = text
	//}`)
	//s=*s.Init(buffer)//接收值
	s.Init(buffer)//不接收值，两种方式都可以
	// Scan reads the next token or Unicode character from source and returns it.
	// It only recognizes tokens t for which the respective Mode bit (1<<-t) is set.
	// It returns EOF at the end of the source. It reports scanner errors (read and
	// token errors) by calling s.Error, if not nil; otherwise it prints an error
	// message to os.Stderr.
	//Scan从源中读取下一个标记token或Unicode字符并将其返回。
	//仅识别设置了相应模式位（1 <<-t）的令牌t。
	//它在源末尾返回EOF。 它通过调用s.Error（如果不是nil）来报告扫描程序错误（读取和令牌错误）。 否则，它会向os.Stderr打印一条错误消息。
	//-2表示尚未读取char，不是EOF
	tok := s.Scan()
	for i:=1;tok != scanner.EOF;i++ {
		// TokenText returns the string corresponding to the most recently scanned token.
		// Valid after calling Scan and in calls of Scanner.Error.
		// TokenText返回与最近扫描的令牌token相对应的字符串。
		// 在调用Scan之后和Scanner.Error调用中有效。
		fmt.Printf("===》扫描到第%v个token:%v\n",i,s.TokenText())
		//事实上这里并没有扫描到错误，也就是s.Error如果不经过我下面的这样手动设置错误信息的话，是应该会为nil的
		//但是s.ErrorCount不会跟着改变，因为我的手动设置错误信息
		//
		s.Error= func(s *scanner.Scanner, msg string) {
			fmt.Printf("错误处理函数！！！接收到的第%v个错误信息如下：%v\n",i,msg)
		}
		s.Error(&s,"这是我传递进错误处理函数的错误信息")

		fmt.Printf("扫描过程中遇到的错误个数:%v\n",s.ErrorCount)
		// uint是无符号整数类型，其大小至少为32位。 但是，它是一种独特的类型，而不是uint32的别名。
		fmt.Printf("s.Mode:%v\n",s.Mode)
		//下面是继承的position对象的属性值
		fmt.Printf("s.Column:%v\n",s.Column)
		fmt.Printf("s.Line:%v\n",s.Line)
		fmt.Printf("s.Filename:%v\n",s.Filename)
		fmt.Printf("s.Offset:%v\n",s.Offset)

		fmt.Printf("s.Whitespace:%v\n",s.Whitespace)
		fmt.Printf("s.IsIdentRune:%v\n",s.IsIdentRune)//这个方法将在下面进行展示

		//这个Position对象看下源码就知道是什么了，其实是调用了position.string()方法
		fmt.Printf("s.Position:%v\n",s.Position)

		// Pos returns the position of the character immediately after
		// the character or token returned by the last call to Next or Scan.
		// Use the Scanner's Position field for the start position of the most
		// recently scanned token.
		// Pos返回上次调用Next或Scan所返回的字符或标记之后的字符位置，
		// 也就是当前scan扫描到的字符的后一位，比如hello world中，s.Pos()会返回空格的后一位也就是w字符的position对象，
		// 再次scan时候s.Pos()会返回d的后一位的position对象。其实就是scan的postion对象当前的信息，也是scan即将扫描到的位置信息
		// 使用Scanner对象的Position字段作为最近扫描的令牌的起始位置。
		fmt.Printf("s.Pos():%v\n",s.Pos())

		//下面是继承的position的方法(s.String()和s.IsValid())
		fmt.Printf("s.String():%v\n",s.String())

		// IsValid reports whether the position is valid.
		// IsValid报告position是否有效。
		fmt.Printf("s.IsValid():%v\n",s.IsValid())


		// Next reads and returns the next Unicode character.
		// It returns EOF at the end of the source. It reports
		// a read error by calling s.Error, if not nil; otherwise
		// it prints an error message to os.Stderr. Next does not
		// update the Scanner's Position field; use Pos() to
		// get the current position.
		// Next读取并返回下一个Unicode字符。
		// 它在源字符串的末尾返回EOF。 它通过调用s.Error（如果不为nil）报告读取错误。
		// 否则，它会向os.Stderr打印一条错误消息。 Next()不更新扫描仪的位置字段； 使用Pos（）获取当前位置。
		r:=s.Next()
		fmt.Printf("s.Next():%v-----字符形式为:%s\n",r,string(r))


		// Peek returns the next Unicode character in the source without advancing
		// the scanner. It returns EOF if the scanner's position is at the last
		// character of the source.
		// Peek会在源代码中返回下一个Unicode字符，而无需提前扫描程序。 如果扫描仪的位置在源的最后一个字符处，它将返回EOF。
		// 跟上面的next()方法都是不更新扫描的位置，所以他不是真的扫描，只是读取，扫描的话是会更新扫描的位置和状态信息的，但是读取的话是不会更新扫描对象的任何信息的！
		fmt.Printf("s.Peek():%v\n",s.Peek())


		tok = s.Scan()
		fmt.Println("-------")
	}
	//语句1输出：
	//	-----------scanner.Scanner对象-------------
	//	===》扫描到第1个token:hello
	//	错误处理函数！！！接收到的第1个错误信息如下：这是我传递进错误处理函数的错误信息
	//	扫描过程中遇到的错误个数:0
	//	s.Mode:1012
	//	s.Column:1
	//	s.Line:1
	//	s.Filename:
	//	s.Offset:0
	//	s.Whitespace:4294977024
	//	s.IsIdentRune:<nil>
	//	s.Position:<input>:1:1
	//	s.Pos():<input>:1:6
	//	s.String():<input>:1:1
	//	s.IsValid():true
	//	s.Next():32-----字符形式为:
	//	s.Peek():119
	//	-------
	//	===》扫描到第2个token:world
	//	错误处理函数！！！接收到的第2个错误信息如下：这是我传递进错误处理函数的错误信息
	//	扫描过程中遇到的错误个数:0
	//	s.Mode:1012
	//	s.Column:7
	//	s.Line:1
	//	s.Filename:
	//	s.Offset:6
	//	s.Whitespace:4294977024
	//	s.IsIdentRune:<nil>
	//	s.Position:<input>:1:7
	//	s.Pos():<input>:1:12
	//	s.String():<input>:1:7
	//	s.IsValid():true
	//	s.Next():-1-----字符形式为:�
	//	s.Peek():-1
	//	-------

	//语句2输出：
	//	-----------scanner.Scanner对象-------------
	//	26
	//	===》扫描到第1个token:if
	//	错误处理函数！！！接收到的第1个错误信息如下：这是我传递进错误处理函数的错误信息
	//	扫描过程中遇到的错误个数:0
	//	s.Mode:1012
	//	s.Column:1
	//	s.Line:3
	//	s.Filename:
	//	s.Offset:26
	//	s.Whitespace:4294977024
	//	s.IsIdentRune:<nil>
	//		s.Position:<input>:3:1
	//	s.Pos():<input>:3:3
	//	s.String():<input>:3:1
	//	s.IsValid():true
	//	s.Next():32-----字符形式为:
	//	s.Peek():97
	//	-------
	//	===》扫描到第2个token:a
	//	错误处理函数！！！接收到的第2个错误信息如下：这是我传递进错误处理函数的错误信息
	//	扫描过程中遇到的错误个数:0
	//	s.Mode:1012
	//	s.Column:4
	//	s.Line:3
	//	s.Filename:
	//	s.Offset:29
	//	s.Whitespace:4294977024
	//	s.IsIdentRune:<nil>
	//	s.Position:<input>:3:4
	//	s.Pos():<input>:3:5
	//	s.String():<input>:3:4
	//	s.IsValid():true
	//	s.Next():32-----字符形式为:
	//	s.Peek():62
	//	-------
	//	===》扫描到第3个token:>
	//	错误处理函数！！！接收到的第3个错误信息如下：这是我传递进错误处理函数的错误信息
	//	扫描过程中遇到的错误个数:0
	//	s.Mode:1012
	//	s.Column:6
	//	s.Line:3
	//	s.Filename:
	//	s.Offset:31
	//	s.Whitespace:4294977024
	//	s.IsIdentRune:<nil>
	//	s.Position:<input>:3:6
	//	s.Pos():<input>:3:7
	//	s.String():<input>:3:6
	//	s.IsValid():true
	//	s.Next():32-----字符形式为:
	//	s.Peek():49
	//	-------
	//	===》扫描到第4个token:10
	//	错误处理函数！！！接收到的第4个错误信息如下：这是我传递进错误处理函数的错误信息
	//	扫描过程中遇到的错误个数:0
	//	s.Mode:1012
	//	s.Column:8
	//	s.Line:3
	//	s.Filename:
	//	s.Offset:33
	//	s.Whitespace:4294977024
	//	s.IsIdentRune:<nil>
	//	s.Position:<input>:3:8
	//	s.Pos():<input>:3:10
	//	s.String():<input>:3:8
	//	s.IsValid():true
	//	s.Next():32-----字符形式为:
	//	s.Peek():123
	//	-------
	//	===》扫描到第5个token:{
	//	错误处理函数！！！接收到的第5个错误信息如下：这是我传递进错误处理函数的错误信息
	//	扫描过程中遇到的错误个数:0
	//	s.Mode:1012
	//	s.Column:11
	//	s.Line:3
	//	s.Filename:
	//	s.Offset:36
	//	s.Whitespace:4294977024
	//	s.IsIdentRune:<nil>
	//	s.Position:<input>:3:11
	//	s.Pos():<input>:3:12
	//	s.String():<input>:3:11
	//	s.IsValid():true
	//	s.Next():10-----字符形式为: （空字符串）
	//	s.Peek():9
	//	-------
	//	===》扫描到第6个token:someParsable
	//	错误处理函数！！！接收到的第6个错误信息如下：这是我传递进错误处理函数的错误信息
	//	扫描过程中遇到的错误个数:0
	//	s.Mode:1012
	//	s.Column:2
	//	s.Line:4
	//	s.Filename:
	//	s.Offset:39
	//	s.Whitespace:4294977024
	//	s.IsIdentRune:<nil>
	//	s.Position:<input>:4:2
	//	s.Pos():<input>:4:14
	//	s.String():<input>:4:2
	//	s.IsValid():true
	//	s.Next():32-----字符形式为:
	//	s.Peek():61
	//	-------
	//	===》扫描到第7个token:=
	//	错误处理函数！！！接收到的第7个错误信息如下：这是我传递进错误处理函数的错误信息
	//	扫描过程中遇到的错误个数:0
	//	s.Mode:1012
	//	s.Column:15
	//	s.Line:4
	//	s.Filename:
	//	s.Offset:52
	//	s.Whitespace:4294977024
	//	s.IsIdentRune:<nil>
	//	s.Position:<input>:4:15
	//	s.Pos():<input>:4:16
	//	s.String():<input>:4:15
	//	s.IsValid():true
	//	s.Next():32-----字符形式为:
	//	s.Peek():116
	//	-------
	//	===》扫描到第8个token:text
	//	错误处理函数！！！接收到的第8个错误信息如下：这是我传递进错误处理函数的错误信息
	//	扫描过程中遇到的错误个数:0
	//	s.Mode:1012
	//	s.Column:17
	//	s.Line:4
	//	s.Filename:
	//	s.Offset:54
	//	s.Whitespace:4294977024
	//	s.IsIdentRune:<nil>
	//	s.Position:<input>:4:17
	//	s.Pos():<input>:4:21
	//	s.String():<input>:4:17
	//	s.IsValid():true
	//	s.Next():10-----字符形式为:
	//	s.Peek():125
	//	-------
	//	===》扫描到第9个token:}
	//	错误处理函数！！！接收到的第9个错误信息如下：这是我传递进错误处理函数的错误信息
	//	扫描过程中遇到的错误个数:0
	//	s.Mode:1012
	//	s.Column:1
	//	s.Line:5
	//	s.Filename:
	//	s.Offset:59
	//	s.Whitespace:4294977024
	//	s.IsIdentRune:<nil>
	//	s.Position:<input>:5:1
	//	s.Pos():<input>:5:2
	//	s.String():<input>:5:1
	//	s.IsValid():true
	//	s.Next():-1-----字符形式为:�
	//	s.Peek():-1
	//	-------




}



func check_err(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

}
