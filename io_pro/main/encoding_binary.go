package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

//关于binary包的简单说明：
//	二进制程序包实现了数字和字节序列之间的简单转换以及varint的编码和解码。（关于varint编码需要了解的自己百度）
//
//	通过读取和写入固定大小的值来转换数字。
//	固定大小的值可以是固定大小的算术类型（bool，int8，uint8，int16，float32，complex64等），也可以是仅包含固定大小值的数组或结构。
//
//	varint函数使用可变长度编码对单个整数值进行编码和解码； 较小的值需要较少的字节。
//	有关规范，请参见https://developers.google.com/protocol-buffers/docs/encoding。
//
//	此软件包优先考虑简单性而不是效率。 需要高性能序列化的客户，特别是对于大型数据结构的客户，应查看更高级的解决方案，例如编码/目标包或协议缓冲区。




func main235342() {
	// 接口ByteOrder指定如何将字节序列转换为16位，32位或64位无符号整数。
	//binary.ByteOrder()

	//binary.Read():
	//读取将结构化的binary二进制的数据从r读取到data中。
	//数据必须是指向固定大小值的指针或固定大小值的切片。
	//从r读取的字节使用指定的字节顺序解码，并写入数据的连续字段。
	//解码布尔值时，零字节解码为false，其他任何非零字节解码为true。
	//读入结构时，将跳过具有空白（_）字段名称的字段的字段数据； 即，空白字段名称可用于填充。
	//读入结构时，必须导出所有非空白字段，否则“读”可能会出现紧急情况。
	//
	//仅当未读取任何字节时，err才是EOF。
	//如果在读取部分但不是全部字节后发生EOF，则Read返回ErrUnexpectedEOF。
	fmt.Println("--------------binary.Read111解码二进制数据------------------")
	var pi float64
	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	buf := bytes.NewReader(b)
	err := binary.Read(buf, binary.LittleEndian, &pi)//binary.LittleEndian实例的类型littleEndian实现了binary.ByteOrder{}接口
	checkErr(err)
	//if err != nil {
	//	fmt.Println("binary.Read failed:", err)
	//}
	fmt.Println(pi)//3.141592653589793

	fmt.Println("--------------binary.Read222解码二进制数据------------------")
	//var pi1 int//没有int
	var pi1 int64
	//b1 := []byte{'a','b','c','d'}//unexpected EOF,也许需要更多的字节才可以
	//b1 := []byte{'a','b','c','d','e','f','g','h'}////7523094288207667809(已经达到了最大位宽19位)
	//b1 := []byte{'a','b','c','d','e','f','g','h','a','b','c','d','e','f','g','h'}////7523094288207667809(已经达到了最大位宽19位)
	//b1 := []byte{1,2,3,4,5,6,7,8}//578437695752307201(已经达到了最大位宽19位)
	//b1 := []byte{1,2,3,4,5,6,7,8,9}//578437695752307201(已经达到了最大位宽19位)
	b1 := []byte{11,22,3,4,5,6,7,8}//578437695752312331(已经达到了最大位宽19位)
	//b1 := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	buf1 := bytes.NewReader(b1)
	err1 := binary.Read(buf1, binary.LittleEndian, &pi1)
	checkErr(err1)

	fmt.Println(pi1)

	fmt.Println("--------------binary.Write编码还原数据------------------")
	// Write将数据的二进制表示形式写入w。
	//数据必须是固定大小的值或固定大小的值的切片，或者是指向此类数据的指针。
	//布尔值编码为一个字节：1表示true，0表示false。
	//写入w的字节使用指定的字节顺序进行编码，并从数据的连续字段中读取。
	//写入结构时，将为具有空白（_）字段名称的字段写入零值。

	// Size返回Write将生成多少字节以对值v进行编码，该值必须是固定大小的值或固定大小的值的切片，或者是指向此类数据的指针。
	//如果v都不是，Size返回-1。
	//先讲这个方法是因为在下面的解码后的容器切片的初始化中要用到
	size := binary.Size(pi1)
	fmt.Println(size)

	b2 := make([]byte,0,size)//在这里用到，不用也可以，但是会浪费内存
	buf2:= bytes.NewBuffer(b2)
	err222 := binary.Write(buf2, binary.LittleEndian, &pi1)//我们将上面编码后的东西还原成原来的二进制数据
	checkError(err222)

	fmt.Println("编码后的数据是：",pi1)
	fmt.Println("还原为二进制数据是：",buf2.Bytes())
	//输出：
	//	8
	//	编码后的数据是： 578437695752312331
	//	还原为二进制数据是： [11 22 3 4 5 6 7 8]

	fmt.Println("-----------------从这开始到下面都是varint编码编码，跟上面似乎没半毛钱关系！！！---------------------")
	fmt.Println("-----------------varint编码Varint---------------------")

	//此文件实现64位整数的“ varint”编码。
	//编码为：
	//-无符号整数一次序列化7位，从最低有效位开始
	//-每个输出字节中的最高有效位（msb）指示是否存在连续字节（msb = 1）
	//-使用“ zig-zag”编码将有符号整数映射为无符号整数：正值x表示为2 * x + 0，负值表示为2 *（^ x）+1；即，对负数进行补码，并在位0中编码是否补码。
	//
	//设计说明：
	// 64位值最多需要10个字节。编码可能会更密集：完整的64位值需要一个额外的字节才能保存位63。
	//相反，由于我们知道不能超过64位，因此可以使用前一个字节的msb来保存位63。这是一个微不足道的改进，
	//   可以将最大编码长度减少到9个字节。但是，它打破了msb始终是“连续位”的不变性，因此使格式与较大数字（例如128位）的varint编码不兼容。


	// Varint从buf解码int64并返回该值和读取的字节数（> 0）。 如果发生错误，则值为0，字节数n <= 0，含义如下：
	// n == 0：buf太小
	// n <0：值大于64位（溢出），-n为读取的字节数
	//其实底层是采用Uvarint（）
	b3 := []byte{11,22,3,4,5,6,7,8}
	//b3 := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}//输出1和12
	int64_my, n := binary.Varint(b3)
	fmt.Println("读取的字节数为：",n)
	fmt.Println("解码后的数据为：",int64_my)//解码后的数据
	//输出：
	//	读取的字节数为： 1
	//	解码后的数据为： -6

	fmt.Println("-----------------varint编码Uvarint---------------------")
	// Uvarint从buf解码uint64并返回该值和
	//读取的字节数（> 0）。 如果发生错误，则值为0
	//字节数n <= 0表示：
	//
	// n == 0：buf太小
	// n <0：值大于64位（溢出），-n为读取的字节数

	int64_my111, n111 := binary.Uvarint(b3)
	fmt.Println("读取的字节数为：",n111)
	fmt.Println("解码后的数据为：",int64_my111)//解码后的数据
	//输出：
	//	读取的字节数为： 1，不知道为什么每次只是读取一个字节
	//	解码后的数据为： 11

	fmt.Println("-----------------varint编码ReadVarint---------------------")
	// ReadVarint从r读取一个编码的带符号整数，并将其作为int64返回。其实底层是采用ReadUvarint方法实现的
	//其实跟上面的2个方法差不多，只是参数不同而已

	buffer := bytes.NewBuffer([]byte{11, 22, 3, 4, 5, 6, 7, 8})
	i64, err_b := binary.ReadVarint(buffer)//不知道为什么只是读取一个字节
	checkError(err_b)

	fmt.Println("buffer中还剩余的字节数为：",buffer.Len())
	fmt.Println("解码后的数据为：",i64)//解码后的数据
	//输出：
	//	buffer中还剩余的字节数为： 7
	//	解码后的数据为： -6

	fmt.Println("-----------------varint编码ReadUvarint---------------------")

	// ReadUvarint从r读取一个编码的无符号整数，并将其作为int64返回。
	//其实跟上面的2个方法差不多，只是参数不同而已

	buffer111 := bytes.NewBuffer([]byte{11, 22, 3, 4, 5, 6, 7, 8})
	i64111, err_b111 := binary.ReadUvarint(buffer111)//不知道为什么只是读取一个字节，也许varint编码就是这样的
	checkError(err_b111)

	fmt.Println("buffer中还剩余的字节数为：",buffer111.Len())
	fmt.Println("解码后的数据为：",i64111)//解码后的数据
	//输出：
	//	buffer中还剩余的字节数为： 7
	//	解码后的数据为： 11

	//为了输出全部的字节，我们这次来读取真个切片并且编码他
	fmt.Println("-----------------varint编码ReadUvarint整个字节切片---------------------")
	src222:=[]byte{11, 22, 3, 4, 5, 6, 7, 8}
	buffer222 := bytes.NewBuffer(src222)
	len_ls:=binary.Size(src222)

	dst_ls:=make([]uint64,len_ls)
	buf_len:=buffer222.Len()

	//for i:=0;i<buffer222.Len() ; i++ {//又犯同样的错，不要将buf的长度作为循环退出的条件，因为他是随时会变化的,而是应该先获取他的长度
	for i:=0;i<buf_len ; i++ {//又犯同样的错，不要将buf的长度作为循环退出的条件，因为他是随时会变化的,而是应该先获取他的长度
		i64222, err_b222 := binary.ReadUvarint(buffer222)
		checkError(err_b222)
		fmt.Println("此次解码单个字节后的数据为：",i64222)//解码后的数据
		fmt.Println("buffer中还剩余的字节数为：",buffer222.Len())
		dst_ls[i]=i64222
	}
	fmt.Println("完全解码字节切片后的数据为：",dst_ls)
	//输出：
	//	此次解码单个字节后的数据为： 11
	//	buffer中还剩余的字节数为： 7
	//	此次解码单个字节后的数据为： 22
	//	buffer中还剩余的字节数为： 6
	//	此次解码单个字节后的数据为： 3
	//	buffer中还剩余的字节数为： 5
	//	此次解码单个字节后的数据为： 4
	//	buffer中还剩余的字节数为： 4
	//	此次解码单个字节后的数据为： 5
	//	buffer中还剩余的字节数为： 3
	//	此次解码单个字节后的数据为： 6
	//	buffer中还剩余的字节数为： 2
	//	此次解码单个字节后的数据为： 7
	//	buffer中还剩余的字节数为： 1
	//	此次解码单个字节后的数据为： 8
	//	buffer中还剩余的字节数为： 0
	//	完全解码字节切片后的数据为： [11 22 3 4 5 6 7 8]

	fmt.Println("-----------------binary.PutUvarint编码数据-----------------------------")
	// PutUvarint将uint64编码为buf并返回写入的字节数。
	//如果缓冲区太小，PutUvarint会惊慌。

	new_byte:=make([]byte,20)
	putUv_num:= binary.PutUvarint(new_byte, 578437695752312331)

	fmt.Println("编码后写入到切片中的字节数为：",putUv_num)
	fmt.Println("编码后的切片为：",new_byte)
	//输出：
	//	编码后写入到切片中的字节数为： 9
	//	编码后的切片为： [139 172 140 160 208 192 193 131 8 0 0 0 0 0 0 0 0 0 0 0]

	fmt.Println("-----------------binary.PutVarint编码数据-----------------------------")
	new_byte111:=make([]byte,20)
	putUv_num111:= binary.PutVarint(new_byte111, 578437695752312331)

	fmt.Println("编码后写入到切片中的字节数为：",putUv_num111)
	fmt.Println("编码后的切片为：",new_byte111)
	//输出：
	//	编码后写入到切片中的字节数为： 9
	//	编码后的切片为： [150 216 152 192 160 129 131 135 16 0 0 0 0 0 0 0 0 0 0 0]

	//对这种编码我也不大熟悉，只知道怎么用

}
func checkErr(err error)  {
	if err != nil {
		fmt.Println(err)
	}
}

















