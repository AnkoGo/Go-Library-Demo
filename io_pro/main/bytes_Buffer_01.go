package main

import (
	"bytes"
	"fmt"
)

func main4444() {
	byte_slice := []byte{97, 98, 99, 100, 101, 102, 126, 127}
	Bf:=bytes.NewBuffer(byte_slice)
	fmt.Println(Bf)//会在控制台打印出字面值，最大126可显示出字面值，超过的就会显示乱码了
	//fmt.Println(Bf[0])//buffer不支持索引

	//ls:=make([]byte,5)
	//fmt.Println(Bf.Read(ls))
	//fmt.Println(Bf)//索引会接着上面的读取指针进行读取到结尾
	//fmt.Println(ls)//[97 98 99 100 101]

	fmt.Println("-----------------")


	fmt.Println(Bf.ReadByte())//读取下一个索引的值
	fmt.Println(Bf)//bcdef~(每次读取一个字节，Bf就会pop一个字节出来)

	fmt.Println(Bf.UnreadByte())//不但指针回退，pop出去的索引值被重新push进来
	fmt.Println(Bf)//abcdef~(每次读取一个字节，Bf就会pop一个字节出来)
	//因为上面的UnreadByte将读取指针调前一位了，所以这里还是读取到和上面的ReadByte读取一样的值
	fmt.Println(Bf.ReadByte())
	fmt.Println(Bf)//bcdef~(每次读取一个字节，Bf就会pop一个字节出来)


	//索引会接着上面的读取指针进行读取下一个并且返回下一个索引值并且pop出一个索引值，如果是从最开始读取的话，那么就是最开始的索引值
	fmt.Println(Bf.ReadRune())
	fmt.Println(Bf)//cdef~

	fmt.Println(Bf.UnreadRune())//跟上面的函数作用相反
	fmt.Println(Bf)//bcdef~

	fmt.Println(Bf.ReadRune())//跟UnreadByte同理
	fmt.Println(Bf)//cdef~

	fmt.Println("-----------------")
	ls:=[]byte{65,66,67,68,69}
	reader := bytes.NewReader(ls)
	fmt.Println("******",reader)//&{[65 66 67 68 69] 0 -1}
	n, _ := Bf.ReadFrom(reader)///从io.Reader中append到Buffer中去原地扩充元素，Buffer的容量跟随需要扩大
	fmt.Println(n)//5,表示append到Buffer多少个值
	fmt.Println(Bf)//cdef~ABCDE
	fmt.Println("******",reader)// &{[65 66 67 68 69] 5 -1}

	fmt.Println("-----------------")
	b:=byte('e')
	line, _ := Bf.ReadString(b)//读取到元素的相应值为b的值的话就返回从开始到这个元素的这段相应的字符串
	fmt.Println(line)//cde
	fmt.Println("-----------------")
	fmt.Println(Bf)//f~ABCDE

	fmt.Println("-----------------")
	b111:=byte('C')
	lines, _ := Bf.ReadBytes(b111)//注意这里返回的是字节切片的类型，而上面那个函数返回的是字符串的类型
	fmt.Println(lines)
	fmt.Println(Bf)

	fmt.Println("-----------------")
	fmt.Printf("%p--%v\n",Bf,Bf)
	Bf.Reset() //通过原地重置切片的索引值和长度值重置了Bf
	fmt.Printf("%p--%v\n",Bf,Bf)//空的Buffer会返回空的什么都不显示

	fmt.Println("-----------------")
	ls111:=[]byte{65,66,67,68,69}
	reader111 := bytes.NewReader(ls111)
	fmt.Println(Bf.ReadFrom(reader111))
	//fmt.Println(reader)//&{[65 66 67 68 69] 5 -1}
	//fmt.Println(Bf.ReadFrom(reader))//0 <nil>,重复的reader不能被读取2次进buffer,因为reader里面的索引值为最尾端的值了,比如上面的值5
	fmt.Println(Bf)

	fmt.Println("-----------------")
	fmt.Println(Bf.Write([]byte{77,78,79,80}))
	fmt.Println(Bf)

	fmt.Println("-----------------")
	fmt.Println(Bf.Len())//9
	fmt.Println(Bf.String())//ABCDEMNOP
	fmt.Println(Bf.Bytes())//[65 66 67 68 69 77 78 79 80]
	fmt.Println(Bf.Cap())//528

	fmt.Println("-----------------")
	fmt.Println(Bf.WriteString("xyz"))
	fmt.Println(Bf)
	fmt.Println(Bf.Bytes())//[65 66 67 68 69 77 78 79 80 120 121 122]
	//一个中文占3个字节，而且这2个中文字是按照utf8编码后的字节存进去的，
	// 非原本的unicode字符的字节数组
	fmt.Println(Bf.WriteString("你好"))
	fmt.Println(Bf)
	fmt.Println(Bf.Bytes())//[65 66 67 68 69 77 78 79 80 120 121 122 228 189 160 229 165 189]

	fmt.Println("-----------------")
	r, size, _ := Bf.ReadRune()
	fmt.Println(r, size)//65 1
	Bf.Reset()
	Bf.WriteString("你好")
	fmt.Println(Bf)
	fmt.Println(Bf.Bytes())
	r111, size111, _ := Bf.ReadRune()//rune会将前3个字节一起编码为uft8的单个字符串类型，具体编码格式如下：
	fmt.Println(r111, size111)//20320 3(20320是)
	fmt.Println(string(r111))//对二进制（十进制为20320）的值进行解析转义。输出：你
	//假设前3个字节：228 189 160
	//字节值：二进制值
	//228:11100100
	//189:10111101
	//160:10100000
	//去掉每个字节的前2位，然后分别为：
	//228:100100
	//189:111101
	//160:100000
	//然后连接这3个二进制为：100100111101100000
	//最后将二进制存进去内存中
	//其实我们发现100100111101100000十进制是20320，也就是对应utf8字典中的你这个中文单字符
	//我们输出到屏幕上面就是需要解码的过程，这里说的解码并不是由utf8类型的二进制转成单字节的二进制，
	//而是翻译二进制为字符串在屏幕上显示出来成为人眼能够识别的字符串的过程，我们也成为解码，其实解码的概念很大的！

	// string是所有8位字节的字符串的集合，按常规，但不是
	//必须代表UTF-8编码的文本。 字符串可以为空，但不能为nil。 字符串类型的值是不可变的。

	fmt.Println("-----------------")
	//新建立一个Buffer好了。
	byte_slice111 := []byte{97, 98, 99, 100, 101, 102, 103,104}
	Bf111:=bytes.NewBuffer(byte_slice111)
	fmt.Println(Bf111)
	//截断会丢弃缓冲区中除前n个未读取字节以外的所有字节，但会继续使用相同的已分配存储。
	//如果n为负数或大于缓冲区的长度，则会发生恐慌。
	Bf111.Truncate(3)//保留前3个字节,其他去除
	fmt.Println(Bf111)


	//如有必要，Grow会增加缓冲区的容量，以保证另外n个字节的空间。 在Grow（n）之后，至少可以将n个字节写入缓冲区，而无需进行其他分配。
	//如果n为负数，Grow会惊慌。
	//如果缓冲区无法增长，则会因ErrTooLarge感到恐慌。
	fmt.Println(Bf111.Cap(),"--",Bf111.Len())//8 -- 3
	Bf111.Grow(20)//计算方法是当前的cap长度*2+n,也就是8*2+20=36
	fmt.Println(Bf111.Cap(),"--",Bf111.Len())//36 -- 3


	// Next返回一个切片，其中包含缓冲区中的下n个字节，
	//前进缓冲区，就好像字节已由Read返回。
	//如果缓冲区中的字节数少于n个，则Next返回整个缓冲区。
	//切片仅在下一次调用read或write方法之前才有效。
	fmt.Println("-----------------")
	fmt.Println(Bf111)//abc
	fmt.Println(Bf111.Next(2))//返回字节切片[97 98]，读取接下来的2个，同时读取指针前进2，源缓存会改变
	fmt.Println(Bf111)//c
	fmt.Println(Bf111)//c
	Bf111.WriteString("defg")//加点东西
	fmt.Println(Bf111)//cdefg
	fmt.Println(Bf111.Next(6))//[99 100 101 102 103],很明显只有5个，但是我们却要取6个，于是全部返回
	fmt.Println(Bf111)//这里输出空

	fmt.Println("-----------------")
	// WriteString将s的内容附加到缓冲区，根据需要增大缓冲区。
	// 返回值n是s的长度； 错误始终为零。 如果缓冲区太大，WriteString将对ErrTooLarge感到恐慌。
	//Bf111.WriteString("abcdef世")//写入中文也可以
	Bf111.WriteString("abcdefg")//内部是通过copy()实现的，copy()的用法如下：
	//复制内置函数将元素从源切片复制到目标切片。
	// （在特殊情况下，它还会将字节从字符串复制到字节切片。）
	// 源和目标可能会重叠。 复制返回复制的元素数量，该数量将是len（src）和len（dst）的最小值。
	fmt.Println(Bf111)//abcdefg


	// WriteRune将Unicode代码点r的UTF-8编码附加到缓冲区，返回其长度和错误，
	// 该错误始终为nil，但包含在其中以匹配bufio.Writer的WriteRune。 缓冲区根据需要增长；
	//如果太大，WriteRune会因ErrTooLarge感到恐慌。
	fmt.Println("-----------------")
	var a rune=104//abcdefgh
	//var a rune=19990//abcdefg世
	fmt.Println(Bf111.WriteRune(a))//往切片里面写入utf8编码的rune
	fmt.Println(Bf111)//abcdefgh


}
