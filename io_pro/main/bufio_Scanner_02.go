package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
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
	//
	//
	//func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error)
	//	ScanRunes是用于Scanner类型的分割函数（符合SplitFunc），本函数会将每一行文本去掉末尾的换行标记作
	//	为一个token返回。返回的行可以是空字符串。换行标记为一个可选的回车后跟一个必选的换行符。最后一行即
	//	使没有换行符也会作为一个token返回。


	// An artificial input source.
	const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Count the words.
	count := 0
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
		count++//对每个scanner.Text()字符串进行计数
	}
	if err := scanner.Err(); err != nil {
		// Fprintln格式使用其操作数的默认格式并写入w。
		//始终在操作数之间添加空格，并添加换行符。
		//返回写入的字节数以及遇到的任何写入错误。
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("%d\n", count)
	// ScanWords是扫描程序的拆分功能，它返回每个空格分隔的文本单词，
	// 并删除周围的空格。 它永远不会返回空字符串。 空间的定义由unicode.IsSpace设置。
	//因为不同的编码有不同的空格unicode值,所以这里所说的空格是unicode值表示的空格，他包括：
	// 第一：' ', '\t', '\n', '\v', '\f', '\r',
	// 第二：'\u0085', '\u00A0'
	// 第三：'\u2000' <= r && r <= '\u200a'
	// 第四：'\u1680', '\u2028', '\u2029', '\u202f', '\u205f', '\u3000'

	//输出如下：
		//===================================
		//Now
		//is
		//the
		//winter
		//of
		//our
		//discontent,
		//	Made
		//glorious
		//summer
		//by
		//this
		//sun
		//of
		//York.
		//15
}

























