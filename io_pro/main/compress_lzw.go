package main

import (
	"compress/lzw"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//对于这种算法建议阅读https://baike.baidu.com/reference/7293853/524dhTCm7Hnt8jSdM7ycM4Gd81AvPcI0GG_T94Jzzkd1qM_Mm05Z_I7EhSedUJYIspbGWi1segdqT65xzSUHAUvns38us4pu2XGFFw
	//或者https://segmentfault.com/a/1190000011425787?utm_source=tag-newest
	//这种算法多用于重复性比较大的数据，如压缩GIF文件、图象、文本文件是最理想的，对于像我下面的读取少字的pdf反而因压缩会变大！

	fmt.Println("---------压缩文件------------")
	dst_file, e := os.Create("main/lzw/lzw_dst.txt")
	check_err_lzw(e)

	////顺序指定LZW数据流中的位顺序。
	//type Order int
	//const (
	//	// LSB表示GIF文件格式中使用的最低有效位优先。
	//	LSB Order = iota
	//	// MSB首先表示最高有效位，如TIFF和PDF文件格式所用。
	//	MSB
	//)

	// NewWriter创建一个新的io.WriteCloser。
	//写入返回的io.WriteCloser被压缩并写入w。
	//完成写入后，调用者有责任在WriteCloser上调用Close。
	//用于文字代码的位数litWidth必须在[2,8]范围内，通常为8。输入字节必须小于1 << litWidth。
	wt := lzw.NewWriter(dst_file, lzw.MSB, 8)//如果是中文的pdf只能是8，

	src_bytes, e := ioutil.ReadFile("main/lzw/lzw_src.pdf")
	check_err_lzw(e)
	n, e := wt.Write(src_bytes)
	check_err_lzw(e)

	fmt.Println("写入压缩文件的字节数是：",n)
	e = wt.Close()
	check_err_lzw(e)

	fmt.Println("---------读取压缩文件------------")

	ret, e := dst_file.Seek(0, 0)
	check_err_lzw(e)


	fmt.Println("将压缩文件的读写指针重置：",ret)

	dst_bytes, e := ioutil.ReadAll(dst_file)

	fmt.Println("读取到的压缩文件的字节数是：",len(dst_bytes))
	fmt.Println("读取到的压缩文件的字节是：",dst_bytes)
	fmt.Println("读取到的压缩文件的字符串是：",string(dst_bytes))



	fmt.Println("---------解压压缩文件并且读取解压后的数据------------")
	i, e := dst_file.Seek(0, 0)
	check_err_lzw(e)

	fmt.Println("重置压缩文件的读写指针到：",i)

	dst_bytes111:=make([]byte,0,512*91)//不给长度，记得，这样的话append才会从切片的开头开始追加
	fmt.Printf("总装载器的地址是：%p ",dst_bytes111)
	// 返回的rd实际上是bufio.NewReader(r)，而bufio.NewReader(r)就是限制了4096个字节的。
	rd := lzw.NewReader(dst_file, lzw.MSB, 8)
	aready_rd_num :=0
	for {
		//每次都重新申请内存，我们在这里模仿缓存，如果你想要节省空间的话可以使用真正的缓存buf
		//buffer := make([]byte,4096)
		buffer := make([]byte,512)//为了不让解压后的数据生成的文件过大，我们不用4096，因为如果一开始给太大，
		// 在最后那里几次性申请的内存会非常的大，具体表现需要你自己打印输出看看,如果你希望不产生多余的字节，那么可以通过buf来实现的


		m, e := rd.Read(buffer)//这里千万不要read大于4096字节的切片，否则会出错(数据不会真的全部解压)
		if m < 0 {
			panic("errNegativeRead")
		}
		aready_rd_num += m
		fmt.Println("已读到的压缩文件的字节数为",aready_rd_num)
		fmt.Println("buf的长度是：",len(buffer))

		dst_bytes111=append(dst_bytes111,buffer...)
		fmt.Printf("总装载器的地址是：%p\t",dst_bytes111)
		//n=copy(dst_bytes111,buffer)//千万不要用copy，因为长度容量什么的都会被复制过去了，这并不是我们想要的结果
		fmt.Println("总容量为：",len(dst_bytes111))

		if e == io.EOF {
			break
		}
		if e != nil {
			fmt.Println("读取解压发生了错误。。。")
		}
	}

	//下面的注释掉的一大段是之前的失败尝试，之所以失败是因为我们 rd.Read(dst_bytes111)时候dst_bytes111大于4096个字节了
	//-------------------------------------11111111111111--------------------------------------------
	//dst_bytes111:=make([]byte,45358,45358)
	//n2, e1 := rd.Read(dst_bytes111)
	//fmt.Println("遍历读取，每次读取的字节数为：",n2)
	//// rd实际上是bufio.NewReader(r)，而bufio.NewReader(r)就是限制了4096个字节的。
	//// 因为文件大于4096,所以不可以使用这个方法进行单次读取完全部，必须使用buf缓存来进行读取，因为缓存会自动扩展容量的，
	//// 经典的有缓存读取的api是ioutil.ReadAll（），一般我们都会优先使用这个api,但是这个缓存大小初始化仅为512个字节，因此
	//// 如果你不希望多次申请内存的话，那么就采用下面的循环遍历读取来提升性能，而不是用便捷的ioutil.ReadAll（）方法！但是不知道
	//// 为什么读取出来的数据像是没解压的，因为我打开pdf是没正确显示中文的！
	//
	//var num int=n2
	//fmt.Println("num第一次的值为：",num)
	//for e1 != io.EOF{
	//	//defer里面的recoverb并不会捕获文件读取到结尾抛出来的信息，因此不会执行recover下面的东西
	//	//defer func() {
	//	//	if err := recover(); err != nil {
	//	//		fmt.Println(err) // 这里的err其实就是panic传入的内容,
	//	//	}
	//	//}()
	//
	//	n2, e1 = rd.Read(dst_bytes111[n2:])//每次只能读取4096个字节。读取到最后时候读取不到字节因此n2最后值为0
	//	// 下面的ReadAll的内部实现就是通过缓存进行读取的！我们可以看到这个Read方法有一定的缺陷。
	//	if e1 !=nil{
	//		fmt.Println("读取到了文件结尾，抛出读取到文件结尾信息为：",e1)
	//		break
	//	}
	//	num+=n2
	//	fmt.Println("num当前值为：",num)
	//}
	////事实上这个Close并没做任何事情，只是设置了标志位为close而已
	//e = rd.Close()
	//
	//check_err_lzw(e)
	//-------------------------------------111111111111111------------------------------------------

	//-------------------------------------222222222222--------------------------------------------
	// readAll从r读取直到出现错误或EOF，并从分配给指定容量的内部缓冲区中返回读取的数据。
	//dst_bytes111, e := ioutil.ReadAll(rd)
	//check_err_lzw(e)
	//-------------------------------------222222222222--------------------------------------------
	//fmt.Println("压缩文件解压之后的字节数是：",num)//循环读取需要放开这行注释同时注释下一行的代码
	fmt.Println("压缩文件解压之后的字节数是：",len(dst_bytes111))
	fmt.Println("压缩文件解压之后的字节是：",dst_bytes111)
	fmt.Printf("压缩文件解压之后的字符串是：\n%v\n",string(dst_bytes111))

	//不知道上面的中文为什么输出乱码

	fmt.Println("---------将解压后的数据写入到文件中去------------")
	dst_file111, e := os.Create("main/lzw/lzw_dst.pdf")
	check_err_lzw(e)

	n3, e := dst_file111.Write(dst_bytes111)
	check_err_lzw(e)

	fmt.Println("解压后的数据写入到新文件中的字节数是：",n3)

	//输出如下：自己执行，因为输出太多了！此处略！


}

func check_err_lzw(err error)  {
	if err != nil {
		fmt.Println(err)
	}
}















