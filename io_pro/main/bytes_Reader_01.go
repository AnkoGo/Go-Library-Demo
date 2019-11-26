package main

import (
	"bytes"
	"fmt"
)
func main24342() {
	byte_slice:=[]byte{99,98,99,99,101,102,103,107}
	//byte_slice:=[]byte{00,24,00,35,00,56,00,34}
	//byte_slice:=[]byte{99,98,99,99,100,100,100,107}
	ls:=make([]byte,6)
	reader := bytes.NewReader(byte_slice)
	fmt.Println(reader)
	fmt.Println(reader.Read(ls))
	fmt.Println(reader)
	fmt.Println(ls)
	ls111:=make([]byte,6)
	fmt.Println(reader.ReadAt(ls111,2))
	fmt.Println(ls111)
	//ReadRune读取并返回Reader中的下一个utf-8码值。如果没有数据可用，
	// 返回值err为io.EOF。如果缓冲中的数据是错误的utf-8编码，本方法会吃掉一字节并返回(U+FFFD, 1, nil)。
	ch, size, _ := reader.ReadRune()//ch是读取结束时候的指针值，size是单前的字节宽度，这表明里面没有任何一个rune类型的字节
	//如果读取成功的话各个参数的意义就不一样了
	fmt.Println("----",ch)
	fmt.Printf("%T\n",ch)
	fmt.Println(size)
	fmt.Printf("%T---%v\n",byte_slice[0],byte_slice[0])

	fmt.Println("--------------------")
	// Size返回基础字节片的原始长度。
	// Size是可通过ReadAt读取的字节数。
	//返回的值始终相同，并且不受任何其他方法的调用影响。
	fmt.Println(reader.Size())
	fmt.Println(reader.Len())//Len返回r包含的切片中还没有被读取的部分。
	fmt.Println("--------------------")

	//fmt.Println(reader)
	//fmt.Println(reader.Seek(4,0))//read的指针移动到索引4再读
	//ls222:=make([]byte,3)
	//fmt.Println(reader)
	//fmt.Println(reader.Read(ls222))
	////因为上面的读取，导致这里开始读取的索引为左后一个了，所以直接读取最后一个单个字节，注意他只读取单个字节
	//fmt.Println(reader.ReadByte())

	//ls333:=[]byte{1,2,3,4,5}
	//ls333:=[]byte{}
	//fmt.Printf("%p---%v\n",reader,reader)//0xc00005a330---&{[99 98 99 99 100 100 100 107] 8 -1}
	//reader.Reset(ls333)//0xc00005a330---&{[] 0 -1}
	////reader.Reset(ls333)
	////地址没变，相当于原来切片的切片
	//fmt.Printf("%p---%v\n",reader,reader)//0xc00005a330---&{[1 2 3 4 5] 0 -1}

	fmt.Println(reader.UnreadByte())
	fmt.Println(reader)

	fmt.Println("--------------------")
	fmt.Println(reader.ReadRune())//也受读取指针的影响，如果指针在最后了，那就返回最后一个索引值
	fmt.Println(reader.UnreadRune())//似乎是ReadRune()的反向。调用这个之前必须要调用ReadRune（）才可以

	//UnreadRune吐出最近一次调用ReadRune方法读取的unicode码值。
	// 如果最近一次读写操作不是ReadRune，本方法会返回错误。
	// （这里就能看出来UnreadRune比UnreadByte严格多了）
	fmt.Println("--------------------")
	fmt.Println(reader)//&{[99 98 99 99 101 102 103 107] 6 -1}
	fmt.Println(reader.UnreadByte())//读取指针发生变化
	fmt.Println(reader)//&{[99 98 99 99 101 102 103 107] 5 -1}
}

