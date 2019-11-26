package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main32467783() {
	//提前体验下自定义错误
	err := errors.New("这是我自定义的一个错误")
	fmt.Println("你好啊", err)

	fmt.Println("---------------------------")

	reader := bytes.NewReader([]byte{97, 98, 99, 100, 101, 102, 103, 104})
	newReader := bufio.NewReader(reader)
	fmt.Println(newReader.Size()) //4096

	ls := make([]byte, 10)
	n, _ := newReader.Read(ls)
	fmt.Println(n)  //8
	fmt.Println(ls) //[97 98 99 100 101 102 103 104 0 0]//这里必须要注视掉然后下面才可以正常读取，因为不可多次读取

	// ReadBytes读取直到输入中第一次出现delim为止，并返回一个切片，该切片包含直到定界符（包括定界符）的数据。
	//如果ReadBytes在找到定界符之前遇到错误，它将返回错误之前读取的数据和错误本身（通常为io.EOF）。
	//只有且仅当返回的数据不以delim结尾时，ReadBytes才会返回err！= nil。
	//对于简单的用途，扫描仪可能更方便。
	fmt.Println("---------------------------")
	reader1 := bytes.NewReader([]byte{97, 98, 99, 100, 101, 102, 103, 104})
	newReader1 := bufio.NewReader(reader1)
	fmt.Println(newReader1.Size()) //4096

	fmt.Println(newReader1.ReadBytes(100)) //[97 98 99 100] <nil>

	fmt.Println("---------------------------")
	reader2 := bytes.NewReader([]byte{97, 98, 99, 100, 101, 102, 103, 104})
	newReader2 := bufio.NewReader(reader2)
	fmt.Println(newReader2.Size()) //4096

	// NewBuffer使用buf作为其初始内容创建并初始化一个新的Buffer。 新的Buffer拥有buf的所有权，
	// 并且在此调用之后，调用方不应使用buf。 NewBuffer旨在准备一个Buffer以读取现有数据。
	// 它也可以用来设置用于写入的内部缓冲区的初始大小。 为此，buf应该具有所需的容量，但长度为零。
	//在大多数情况下，new（Buffer）（或仅声明一个Buffer变量）足以初始化Buffer。
	lsls := make([]byte, 0, 15) //注意，初始化的长度只能写0，否则会预存数据，达不到空长度的效果，但是我们的总容量必须要给一个，不给的话也会自动扩容的！但是效率不高。
	buffer := bytes.NewBuffer(lsls)
	fmt.Println("写入之前的buffer:", buffer) //写入之前的buffer:

	fmt.Println(newReader2.WriteTo(buffer)) //这里不一定要给io.wirter接口的，我们给这个接口的实现也是可以的
	fmt.Println(lsls)
	fmt.Println("写入之后的buffer:", buffer) //写入之后的buffer: abcdefgh
	ls_b := buffer.Bytes()
	fmt.Println("写入之后的buffer里面的字节切片:", ls_b) //写入之后的buffer里面的字节切片: [97 98 99 100 101 102 103 104]

	fmt.Println(lsls)                  //[]，看来buffer.Bytes()和原始的字节切片是不同个对象的,
	fmt.Printf("%p\n%p\n", lsls, ls_b) //0xc00000a140
	//0xc00000a140
	//由此可以的出，lsls和ls_b是不同的对象，但是却引用同一个底层的内存数组，只是因为lsls的长度为0而无法输出而已

	fmt.Println("---------------------------")
	reader3 := bytes.NewReader([]byte{97, 98, 99, 100, 101, 102, 103, 104})
	newReader3 := bufio.NewReader(reader3)
	fmt.Println(newReader3.Size()) //4096

	//读取单个字节
	fmt.Println(newReader3.ReadByte()) //97 <nil>

	fmt.Println(newReader3.ReadByte()) //98 <nil>

	fmt.Println(newReader3.ReadByte()) //99 <nil>

	fmt.Println("---------------------------")

	reader4 := bytes.NewReader([]byte{97, 98, 99, 100, 101, 102, 103, 104})
	newReader4 := bufio.NewReader(reader4)
	fmt.Println(newReader4.Size()) //4096

	// Peek返回下一个n个字节，而不会使阅读器前进。 字节在下一个读取调用时停止有效。
	// 如果Peek返回的字节数少于n个字节，则它还会返回一个错误，说明读取短的原因。 如果n大于b的缓冲区大小，则错误为ErrBufferFull。
	//
	//调用Peek会阻止UnreadByte或UnreadRune调用成功，直到下一次读取操作为止。
	peek, _ := newReader4.Peek(2)
	fmt.Println(peek) //[97 98]
	ls1 := make([]byte, 15)
	fmt.Println(newReader4.ReadByte()) //97 <nil>,像这里的话是会影响读取器的
	// Size返回底层缓冲区的总大小（以字节为单位）。不关读取的事
	fmt.Println(newReader4.Size())
	fmt.Println(newReader4.Read(ls1)) //7 <nil>
	fmt.Println(ls1)                  //[97 98 99 100 101 102 103 104 0 0 0 0 0 0 0],确实不会影响读取器
	fmt.Println(peek)                 //[97 98],如果不接收的话，那么他的作用就只能是下一次读取之前
	fmt.Println(newReader4.Peek(2))   //[] EOF,读取完了，所以这里没值

	fmt.Println("---------------------------")

	reader5 := bytes.NewReader([]byte{97, 98, 99, 100, 101, 102, 103, 104})
	newReader5 := bufio.NewReader(reader5)
	fmt.Println(newReader5.ReadRune()) //97 1 <nil>,没rune所以返回第一个值

	reader6 := bytes.NewBufferString("你好啊") //这里创建的底层的字节切片就是rune类型的了
	newReader6 := bufio.NewReader(reader6)
	fmt.Println(reader6.Bytes()) //[228 189 160 229 165 189 229 149 138],输出底层的[]byte，但是其实里面的是rune类型的byte
	r, size, _ := newReader6.ReadRune()
	fmt.Println(r, size, string(r)) //20320 3 你

	r1, size1, _ := newReader6.ReadRune()
	fmt.Println(r1, size1, string(r1)) //22909 3 好

	r2, size2, _ := newReader6.ReadRune()
	fmt.Println(r2, size2, string(r2)) //21834 3 啊

	r3, size3, err3 := newReader6.ReadRune()
	fmt.Println(r3, size3, string(r3), err3) //0 0   EOF,读取完了，所以r3应该是0，然后string转义就输出空字符串

	fmt.Println("---------------------------")

	// ReadString读取直到输入中第一次出现delim为止，
	//返回一个字符串，其中包含直到定界符（包括定界符）的数据。
	//如果ReadString在找到定界符之前遇到错误，它将返回错误之前读取的数据和错误本身（通常为io.EOF）。
	//仅当返回的数据不以delim结尾时，ReadString才返回err！= nil。
	//对于简单的用途，扫描仪可能更方便。
	reader7 := bytes.NewReader([]byte{97, 98, 99, 100, 101, 102, 103, 104})
	newReader7 := bufio.NewReader(reader7)
	fmt.Println(newReader7.ReadString(101)) //abcde <nil>

	reader8 := bytes.NewBufferString("你好啊") //这里创建的底层的字节切片就是rune类型的了
	newReader8 := bufio.NewReader(reader8)
	fmt.Println(reader8.Bytes()) //[228 189 160 229 165 189 229 149 138]
	//fmt.Println(newReader8.ReadString(165))//你� <nil>，对于rune类型的东西必须没3个为切割字节才行，否则会乱码，而且会影响后面
	fmt.Println(newReader8.ReadString(160)) //你 <nil>
	fmt.Println(newReader8.ReadString(189)) //好 <nil>，
	fmt.Println(newReader8.ReadString(0))   //啊 EOF，因为不存在所以就返回剩余的全部3个字节，所以如果是最后3个字节的话，我们的参数可以随便给

	//通过上面这个原理我们可以一次性读取所有字节切片并且string转义一步到位
	reader9 := bytes.NewBufferString("你好啊") //这里创建的底层的字节切片就是rune类型的了
	newReader9 := bufio.NewReader(reader9)
	fmt.Println(reader9.Bytes())          //[228 189 160 229 165 189 229 149 138]
	fmt.Println(newReader9.ReadString(0)) //你好啊 EOF,一步全部输出

	fmt.Println("---------------------------")

	reader91 := bytes.NewBufferString("你好啊") //这里创建的底层的字节切片就是rune类型的了
	newReader91 := bufio.NewReader(reader91)
	fmt.Println(reader91.Bytes())           //[228 189 160 229 165 189 229 149 138]
	fmt.Println(newReader91.ReadSlice(229)) //[228 189 160 229] <nil>//读取切片并且返回子切片

	fmt.Println("---------------------------")
	//ReadLine是一个低水平的行数据读取原语。大多数调用者应使用ReadBytes('\n')或ReadString('\n')代替，或者使用Scanner。
	//
	//ReadLine尝试返回一行数据，不包括行尾标志的字节。如果行太长超过了缓冲，返回值isPrefix会被设为true，并返回行的前面一部分。
	//该行剩下的部分将在之后的调用中返回。返回值isPrefix会在返回该行最后一个片段时才设为false。返回切片是缓冲的子切片，只在下
	//一次读取操作之前有效。ReadLine要么返回一个非nil的line，要么返回一个非nil的err，两个返回值至少一个非nil。
	//
	//返回的文本不包含行尾的标志字节（"\r\n"或"\n"）。如果输入流结束时没有行尾标志字节，方法不会出错，也不会指出这一情况。
	//在调用ReadLine之后调用UnreadByte会总是吐出最后一个读取的字节（很可能是该行的行尾标志字节），即使该字节不是ReadLine返回值的一部分。

	reader92 := bytes.NewBufferString("你好啊") //这里创建的底层的字节切片就是rune类型的了
	newReader92 := bufio.NewReader(reader92)
	fmt.Println(reader92.Bytes())       //[228 189 160 229 165 189 229 149 138]
	fmt.Println(newReader92.ReadLine()) //[228 189 160 229 165 189 229 149 138] false <nil>
	fmt.Println(newReader9.Size())      //4096,超过这个大小就会折次读

	file, _ := os.Open("main/test.txt")
	ls6 := make([]byte, 6000)
	i, _ := file.Read(ls6)
	fmt.Println(i)
	fmt.Println(ls6) //[97 98 99 228 188 159 229 164 167 231 154 132 230 151 182 228 187 163 229 136]
	fmt.Println("******************************************")
	fmt.Println(string(ls6)) //abc伟大的时代�
	fmt.Println("******************************************")

	reader93 := bytes.NewBuffer(ls6) //这里创建的底层的字节切片就是rune类型的了
	newReader93 := bufio.NewReader(reader93)
	fmt.Println(string(reader93.Bytes())) //

	fmt.Println("ReadLine....1111")
	//fmt.Println(newReader93.ReadLine())//[228 189 160 229 165 189 229 149 138] false <nil>
	line, isPrefix, err888 := newReader93.ReadLine()
	if err888 != nil {
		fmt.Println(err888)
	}
	fmt.Println(len(line))                   //136
	fmt.Println(string(line), "+", isPrefix) //[228 189 160 229 165 189 229 149 138] false <nil>
	fmt.Println("ReadLine....1111")

	fmt.Println("ReadLine....2222")
	//fmt.Println(newReader93.ReadLine())//[228 189 160 229 165 189 229 149 138] false <nil>
	line1, isPrefix1, err8881 := newReader93.ReadLine()
	if err8881 != nil {
		fmt.Println(err8881)
	}
	fmt.Println(len(line1))                    //151
	fmt.Println(string(line1), "+", isPrefix1) //[228 189 160 229 165 189 229 149 138] false <nil>
	fmt.Println("ReadLine....2222")

	fmt.Println("ReadLine....3333")
	//fmt.Println(newReader93.ReadLine())//[228 189 160 229 165 189 229 149 138] false <nil>
	line2, isPrefix2, err8882 := newReader93.ReadLine()
	if err8882 != nil {
		fmt.Println(err8882)
	}
	fmt.Println(len(line2))                    //99
	fmt.Println(string(line2), "+", isPrefix2) //[228 189 160 229 165 189 229 149 138] false <nil>
	fmt.Println("ReadLine....3333")

	fmt.Println("ReadLine....4444")
	//fmt.Println(newReader93.ReadLine())//[228 189 160 229 165 189 229 149 138] false <nil>
	line3, isPrefix3, err8883 := newReader93.ReadLine()
	if err8883 != nil {
		fmt.Println("错误：", err8883)
	}
	fmt.Println(len(line3))                    //4096,可是还没读完一行，于是把剩余的字节在下次读取
	fmt.Println(string(line3), "+", isPrefix3) //[228 189 160 229 165 189 229 149 138] false <nil>
	fmt.Println("ReadLine....4444")

	fmt.Println("ReadLine....5555")
	//fmt.Println(newReader93.ReadLine())//[228 189 160 229 165 189 229 149 138] false <nil>
	line4, isPrefix4, err8884 := newReader93.ReadLine()
	if err8884 != nil {
		fmt.Println("错误：", err8884)
	}
	fmt.Println(len(line4))                    //1512
	fmt.Println(string(line4), "+", isPrefix4) //[228 189 160 229 165 189 229 149 138] false <nil>
	fmt.Println("ReadLine....5555")
	//不知道为什么后面的会导致出错
	fmt.Println("ReadLine....6666")
	line5, isPrefix5, err8885 := newReader93.ReadLine()
	if err8885 != nil {
		fmt.Println("错误：", err8885) //错误： EOF
	}
	fmt.Println(len(line5))
	fmt.Println(string(line5), "+", isPrefix5) //[228 189 160 229 165 189 229 149 138] false <nil>
	fmt.Println("ReadLine....6666")

	fmt.Println("ReadLine....7777")
	line6, isPrefix6, err8886 := newReader93.ReadLine()
	if err8886 != nil {
		fmt.Println("错误：", err8886) //错误： EOF
	}
	fmt.Println(len(line6))                    //EOF
	fmt.Println(string(line6), "+", isPrefix6) //0（换行）+ false
	fmt.Println("ReadLine....7777")

	fmt.Println(newReader93.Size()) //4096,超过这个大小就会折次读,但是会发生莫名其妙的错误以及乱码，下一次读取会出现很奇怪的东西

	fmt.Println("--------------------------------------")

	reader94 := bytes.NewBuffer([]byte{97, 98, 99, 100, 101, 102, 103, 104})
	newReader94 := bufio.NewReader(reader94)
	fmt.Println(newReader94.ReadBytes(0)) //[97 98 99 100 101 102 103 104] EOF

	reader95 := bytes.NewReader([]byte{99, 99, 99, 100, 101, 102, 103, 104})
	newReader94.Reset(reader95)           //说白了就是更换底层的io.Reader对象
	fmt.Println(newReader94.ReadBytes(0)) //[99 99 99 100 101 102 103 104] EOF
	fmt.Println("--------------------------------------")

	reader96 := bytes.NewBuffer([]byte{97, 98, 99, 100, 101, 102, 103, 104})
	newReader96 := bufio.NewReader(reader96)
	fmt.Println(newReader96.Buffered()) //0,不知道为什么必须要开始读取之后才可以获取可读取的字节数目
	fmt.Println(newReader96.ReadByte()) //97 <nil>
	// Buffered返回可从当前缓冲区读取的字节数。
	fmt.Println(newReader96.Buffered()) //7
	fmt.Println(newReader96.Buffered()) //7
	fmt.Println(newReader96.ReadByte()) //98 <nil>
	fmt.Println(newReader96.Buffered()) //6

	fmt.Println("--------------------------------------")
	// Discard跳过接下来的n个字节，返回丢弃的字节数。
	//如果Discard跳过少于n个字节，则它还会返回错误。
	//如果0 <= n <= b.Buffered（），则确保Discard成功执行，而无需从基础io.Reader中读取。
	//说白了就是移动读取器的指针
	reader97 := bytes.NewBuffer([]byte{97, 98, 99, 100, 101, 102, 103, 104})
	newReader97 := bufio.NewReader(reader97)
	fmt.Println(newReader97.ReadByte()) //97 <nil>
	fmt.Println(newReader97.Discard(2)) //2 <nil>
	fmt.Println(newReader97.Buffered()) //5
	fmt.Println(newReader97.ReadByte()) //100 <nil>

	fmt.Println("--------------------------------------")
	//strings.NewReader：
	//NewReader创建一个从s读取数据的Reader。本函数类似bytes.NewBufferString，但是更有效率，且底层的子串为只读的，但是并不是说这个Reader只读
	reader98 := strings.NewReader("abcdefg")
	newReader98 := bufio.NewReader(reader98)
	fmt.Println(newReader98.ReadBytes(0)) //[97 98 99 100 101 102 103] EOF
	reader99 := strings.NewReader("gfedcba")
	//fmt.Println(newReader98.Reset(reader99))//没返回值的函数不允许这样写。
	newReader98.Reset(reader99) //[103 102 101 100 99 98 97] EOF
	//fmt.Println(newReader98.ReadBytes(0))
	//fmt.Println()
	//fmt.Println(reader99)
	//下面是视图去更改这个strings.Reader,所以必须先注释上面的几行
	slice := make([]byte, 10)
	fmt.Println(reader99.Read(slice))
	fmt.Println(slice)
	reader99.Reset("aaaaaa")
	slice1 := make([]byte, 10) //千万给长度
	fmt.Println(reader99.Read(slice1))
	fmt.Println(slice1)

	//fmt.Println(reader99[0])//不允许下标索引

	reader991 := strings.NewReader("gfedcba")
	slice2 := make([]byte, 10)
	fmt.Println(reader991.ReadAt(slice2, 2)) //获取子片
	fmt.Println(slice2)                      //[101 100 99 98 97 0 0 0 0 0]

}
