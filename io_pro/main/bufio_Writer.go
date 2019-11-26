package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func main1357() {
	// Writer为io.Writer对象实现缓冲。
	//如果在写入Writer时发生错误，将不再接受更多数据，并且所有后续写入和Flush都将返回错误。
	//写入所有数据后，客户端应调用Flush方法以确保所有数据均已转发到基础io.Writer。

	ls:=make([]byte,0)//作为buffer的参数千万不要给长度
	buffer := bytes.NewBuffer(ls)//buffer实现了io.Reader和io.Writer
	writer := bufio.NewWriter(buffer)
	ls111:=[]byte{97,98,99,100,101,102,103}
	nn, err111 := writer.Write(ls111)//将ls111写入writer的底层字节切片ls指向的地址
	if err111 != nil{
		fmt.Println(err111)
	}
	fmt.Println(nn)//7
	fmt.Println(ls111)//[97 98 99 100 101 102 103]
	fmt.Println(writer)//&{<nil> [97 98 99 100 101 102 103 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0....](4096字节)
	//没有方法读取才死
	fmt.Println(writer.WriteRune(104))
	fmt.Println(writer)//&{<nil> [97 98 99 100 101 102 103 104 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0....](4096字节)

	fmt.Println(writer.WriteString("ijk"))
	fmt.Println(writer)//&{<nil> [97 98 99 100 101 102 103 104 105 106 107 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0....](4096字节)

	fmt.Println(writer.WriteByte(108))
	fmt.Println(writer)//&{<nil> [97 98 99 100 101 102 103 104 105 106 107 108 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0....](4096字节)
	//err := writer.Flush()//一般都调用这个做结尾，似乎不调用也可以，除非大数据
	//if err != nil{
	//	fmt.Println(err)
	//}

	// Buffered返回已写入当前缓冲区的字节数。
	fmt.Println(writer.Buffered())//12,如果上面调用了flush()则会输出0，flush相当于置零了writer的计数n
	fmt.Println(writer.Size())//4096
	fmt.Println(writer.Available())//4084(4096-12),表示余下空间多少字节没用
	writer.Flush()
	fmt.Println("buffer:",buffer)//buffer: abcdefghijkl


	fmt.Println("-------------------------------------")
	ls333:=make([]byte,0)
	buffer333 := bytes.NewBuffer(ls333)
	// Reset丢弃所有未刷新的缓冲数据，清除所有错误，然后
	//重置b将其输出写入w。
	writer.Reset(buffer333)//重置,只是重新跟换了写入器而已，底层缓存切片没换的
	writer.WriteString("ppp")
	writer.Flush()
	fmt.Println("==",buffer333)
	fmt.Println(writer)//&{<nil> [97 98 99 100 101 102 103 104 105 106 107 108 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0...](4096字节，不知道为什么还是保留以前的数据)
	fmt.Println(writer.Size())//4096
	fmt.Println(writer.Available())//4096
	fmt.Println(writer.Buffered())//0

	fmt.Println(writer.WriteString("mno"))//3 <nil>
	fmt.Println(writer)//&{<nil> [109 110 111 100 101 102 103 104 105 106 107 108 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 ...]

	fmt.Println("-------------------------------------")
	reader := strings.NewReader("pq")
	fmt.Println(writer.ReadFrom(reader))//2 <nil>
	fmt.Println(writer)//&{<nil> [109 110 111 112 113 102 103 104 105 106 107 108 0 0 0 0 0 0 0 0 0 0 0 0 0...]
	writer.Flush()
	fmt.Println("--",buffer333)//-- mnopq,不加上面的flush则是空字符串，说明数据还在写入器里面，没更新到真的接收器里面去
	//写入器里面的东西随时可能被覆盖，我们要记得flush,取数据也要到buffer中去取才对、
	//另外一个鲜明的例子:
	fmt.Println("-------------------------------------")
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	bw.WriteString("123")
	fmt.Println("--",b)
	c := bytes.NewBuffer(make([]byte, 0))
	bw.Reset(c)
	bw.WriteString("456")
	bw.Flush()
	fmt.Println("--",b)
	fmt.Println(c)
	//输出：
	//--
	//--
	//456

	fmt.Println("-------------------------------------")
	ls6:=make([]byte,0)//作为buffer的参数千万不要给长度
	buffer6 := bytes.NewBuffer(ls6)//buffer实现了io.Reader和io.Writer
	writer6 := bufio.NewWriter(buffer6)
	writer6.WriteString("abc")
	writer6.Flush()
	fmt.Println(writer6)
	fmt.Println("--",buffer6)
	// Len返回缓冲区未读部分的字节数；
	// b.Len（）== len（b.Bytes（））。
	fmt.Println("--",buffer6.Len())

	fmt.Println("--",buffer6.Cap())// Cap返回缓冲区基础字节片的容量，即为缓冲区数据分配的总空间。
	fmt.Println("--",ls6)
	//输出：
	//&{<nil> [97 98 99 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0...]
	//-- abc
	//-- 3
	//-- 3
	//-- []


	fmt.Println("-------------------------------------")
	//newReader("abcd")==NewReaderSize("abcd",4096)

	newReader := strings.NewReader("abcd")
	newReaderSize := bufio.NewReaderSize(newReader, 20)
	fmt.Println(newReaderSize.Buffered())//0
	fmt.Println(newReaderSize.Size())//100
	ls777:=make([]byte,10)
	fmt.Println(newReaderSize.Read(ls777))
	fmt.Println(newReaderSize)//&{[97 98 99 100 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] 0xc000004520 4 4 <nil> 100 -1}
	fmt.Println(ls777)//[97 98 99 100 0 0 0 0 0 0]

	//newWriter(newBuffer)==NewWriterSize(newBuffer,4096)

	ls999:=make([]byte,0)
	newBuffer := bytes.NewBuffer(ls999)
	Wsize := bufio.NewWriterSize(newBuffer, 20)
	i, _ := Wsize.WriteString("dcba")
	Wsize.Flush()
	fmt.Println(i)//4
	fmt.Println(Wsize.Size())//20
	fmt.Println(Wsize.Buffered())//4
	fmt.Println(Wsize)//&{<nil> [100 99 98 97 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] 0 0xc00005c420}
	fmt.Println(newBuffer)//dcba
	fmt.Println(ls999)//[]

	fmt.Println("-------------------------------------")

	// NewReadWriter分配一个新的ReadWriter，该新的ReadWriter调度到r和w。
	newRd := strings.NewReader("ABCD")//读取来源处
	bufio_R := bufio.NewReaderSize(newRd, 20)
	ls555:=make([]byte,0)//写入接收处
	newBf := bytes.NewBuffer(ls555)
	bufio_W :=bufio.NewWriterSize(newBf,20)
	readWriter := bufio.NewReadWriter(bufio_R, bufio_W)//读取写入器
	p_ls:=make([]byte,10)
	readWriter.Read(p_ls)
	fmt.Println(p_ls)//[65 66 67 68 0 0 0 0 0 0]

	w_n, _ := readWriter.WriteString("EFG")
	fmt.Println(w_n)//3
	readWriter.Flush()
	fmt.Println(newBf)//EFG
	fmt.Println(ls555)//[]

	fmt.Println("***************************")
	//将读取器和接收器(或者说是读取写入器的两头)对接
	R1 := strings.NewReader("aabbcc")
	bufio_R1 := bufio.NewReaderSize(R1, 15)

	ls_W:=make([]byte,0)//写入接收器
	newBf_W := bytes.NewBuffer(ls_W)
	bufio_W1 := bufio.NewWriter(newBf_W)

	newReadWriter := bufio.NewReadWriter(bufio_R1, bufio_W1)
	ls_R:=make([]byte,10)//中间承接的切片
	newReadWriter.Read(ls_R)
	fmt.Println("ls_R:",ls_R)//ls_R: [97 97 98 98 99 99 0 0 0 0]

	nn2, err22 := newReadWriter.Write(ls_R)
	if err22 != nil{
		fmt.Println(err22)
	}
	fmt.Println(nn2)//10
	newReadWriter.Flush()
	fmt.Println("ls_W:",ls_W)//ls_W: []
	fmt.Println("newBf_W:",newBf_W)//newBf_W: aabbcc




}

























