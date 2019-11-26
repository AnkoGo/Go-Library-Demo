package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main2468() {
	fmt.Println("===================================")
	//Scanner类型提供了方便的读取数据的接口，如从换行符分隔的文本里读取每一行。
	//
	//成功调用的Scan方法会逐步提供文件的token令牌字节切片，跳过token令牌字节切片之间的字节。token令牌字节切片由SplitFunc类型的分割函数指定；
	//默认的分割函数会将输入分割为多个行，并去掉行尾的换行标志。本包预定义的分割函数可以将文件分割为行、字节、
	//unicode码值、空白分隔的word。调用者可以定制自己的分割函数。
	//
	//scan扫描会在抵达输入流结尾、遇到的第一个I/O错误、token过大不能保存进缓冲时，不可恢复的停止。当扫描停止后，
	//当前读取位置可能会远在最后一个获得的token后面。需要更多对错误管理的控制或token很大，或必须从reader连续扫描的程序，
	//应使用bufio.Reader代替。

	//下面几句是错误的示范
	//sReader := strings.NewReader("wvzz")
	//scanner := bufio.NewScanner(sReader)
	//fmt.Println(scanner.Bytes())
	//fmt.Println(scanner.Text())

	// An artificial input source.
	const input = "1234 5678 123456789"
	// NewScanner返回一个新的Scanner以从r中读取。
	//拆分功能默认为ScanLines。
	scanner := bufio.NewScanner(strings.NewReader(input))
	// SplitFunc是用于对输入进行标记化的split函数的签名。
	// 参数data是剩余的未处理数据的初始子字符串(不是token令牌字节切片)
	// 参数atEOF标志， 该标志报告Reader是否没有更多数据可提供。
	// 返回值advance是推进输入的字节数，以及返回给用户的下一个tokon（如果有）以及err（如果有）。
	//如果函数返回错误，扫描将停止，在这种情况下，某些输入可能会被丢弃。
	//否则，扫描仪将前进输入（advance+1）。如果token令牌字节切片不是nil(说明遇到空格了)，则扫描程序会将其返回给用户。如果token令牌字节切片为nil，则扫描程序将读取更多数据并继续扫描；
	//检查参数atEOF，如果没有更多数据-如果atEOF为true（说明没有更多数据了）-扫描程序将返回。如果数据还没有完整的令牌，例如，如果在扫描行时没有换行符，
	// 则SplitFunc可以返回（0，nil，nil），以指示扫描程序将更多数据读取到切片中，并尝试使用更长的时间切片从输入中的同一点开始。
	//除非atEOF为true，否则永远不要使用空数据片调用该函数。但是，如果atEOF为true，则数据可能为非空，并且一如既往地保留未处理的文本。

	//通过包装现有的ScanWords函数来创建自定义拆分函数。
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// ScanWords是扫描程序的拆分功能，它返回每个空格分隔的文本单词，并删除周围的空格。
		// 它永远不会返回空字符串。 空间的定义由unicode.IsSpace设置。
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			//最多能转化10位数字（二进制是31位），如果参数改为64的话最多可以容纳19位数字
			_, err = strconv.ParseInt(string(token), 10, 64)
		}
		return
	}
	//Split（）设置扫描仪的分割功能。
	//默认的拆分功能是ScanLines。
	//如果在扫描开始后调用了panic，则将其拆分。

	// 指定分割函数
	scanner.Split(split)

	//Scan()会将扫描程序前进到下一个令牌，然后可以通过Bytes或Text方法使用它。 当扫描停止（到达输入结尾或错误）时，它将返回false。
	//在Scan返回false之后，Err方法将返回扫描期间发生的任何错误，但如果是io.EOF，Err将返回nil。
	//如果split函数在不提前输入的情况下返回了太多的空令牌，请扫描恐慌。 这是扫描仪的常见错误模式。
	//下面是验证输入
	for scanner.Scan() {
		//Bytes方法返回最近一次Scan调用生成的token。底层数组指向的数据可能会被下一次Scan的调用重写。
		fmt.Println(scanner.Bytes())
		//Text（）返回通过调用Scan生成的最新令牌，作为保留其字节的新分配字符串。
		fmt.Printf("%s\n", scanner.Text())
	}
	//Bytes方法返回最近一次Scan调用生成的token。底层数组指向的数据可能会被下一次Scan的调用重写。
	fmt.Println(scanner.Bytes())//[],读完了，所以空
	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)//Invalid input: strconv.ParseInt: parsing "1234567901234567890": value out of range
	}

	//附录Scanner结构体字段参数：
	//	r 				io.Reader //客户端提供的读取器。
	//	split 			SplitFunc //拆分令牌字节切片的函数。
	//	maxTokenSize 	int //令牌字节切片的最大大小； 通过测试修改。
	//	token 			[]byte //拆分返回的最后一个token令牌字节切片。
	//	buf 			[]byte //缓冲区用作拆分参数。
	//	start 			int // buf中的第一个未处理字节索引。这个参数会跟随读取器不停的变化。
	//	end 			int // buf中的数据结束索引。
	//	err 			err //粘性错误。
	//	empties 		int //连续空令牌字节切片的计数。
	//	scanCalled 		bool //扫描已被调用； 缓冲区正在使用中。
	//	done 			bool //扫描已完成。


	//Bytes方法返回最近一次Scan调用生成的token。底层数组指向的数据可能会被下一次Scan的调用重写。
	fmt.Println(scanner.Bytes())//[],读完了，所以空

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
}

























