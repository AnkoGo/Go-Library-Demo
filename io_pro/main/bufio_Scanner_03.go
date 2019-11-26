package main

import (
	"bufio"
	"fmt"
	"os"
)

func main54533() {
	fmt.Println("===================================")

	//对于分割函数go内置了4个常用的分割函数，当然你也可以自定义，下面的4个函数分别是按照字节，字符rune,空白空格
	//和换行符进行分割
	//
	//func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error)
	//	ScanBytes是用于Scanner类型的分割函数（符合SplitFunc），本函数会将每个字节作为一个token返回。
	//
	//
	//func ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error)
	//	ScanRunes是用于Scanner类型的分割函数（符合SplitFunc），本函数会将每个utf-8编码的unicode码值作
	//	为一个token返回。本函数返回的rune序列和range一个字符串的输出rune序列相同。错误的utf-8编码会翻译为
	//	U+FFFD = "\xef\xbf\xbd"，但只会消耗一个字节。调用者无法区分正确编码的rune和错误编码的rune。
	//
	//
	//func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error)
	//	ScanRunes是用于Scanner类型的分割函数（符合SplitFunc），本函数会将空白（参见unicode.IsSpace）
	//	分隔的片段（去掉前后空白后）作为一个token返回。本函数永远不会返回空字符串。
	//  // ScanWords是扫描程序的拆分功能，它返回每个空格分隔的文本单词，
	//	// 并删除周围的空格。 它永远不会返回空字符串。 空间的定义由unicode.IsSpace设置。
	//	//因为不同的编码有不同的空格unicode值,所以这里所说的空格是unicode值表示的空格，他包括：
	//	// 第一：' ', '\t', '\n', '\v', '\f', '\r',
	//	// 第二：'\u0085', '\u00A0'
	//	// 第三：'\u2000' <= r && r <= '\u200a'
	//	// 第四：'\u1680', '\u2028', '\u2029', '\u202f', '\u205f', '\u3000'
	//
	//func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error)
	//	ScanRunes是用于Scanner类型的分割函数（符合SplitFunc），本函数会将每一行文本去掉末尾的换行标记作
	//	为一个token返回。返回的行可以是空字符串。换行标记为一个可选的回车后跟一个必选的换行符。最后一行即
	//	使没有换行符也会作为一个token返回。


	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'，默认是这种按行分割模式
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	//第一次输出：
	//	你好啊（单数行都是我自己输入的字符串）
	//	你好啊（双数行都是程序自己打印的字符串）
	//	nnn
	//	nnn
	//	sdsd
	//	sdsd
	//第二次输出：
	//	你好 迭代 sdsd \n（单数行都是我自己输入的字符串）
	//	你好 迭代 sdsd \n（双数行都是程序自己打印的字符串）
	//	sdsd \n sdsds \n  sdsdsd
	//	sdsd \n sdsds \n  sdsdsd
}

























