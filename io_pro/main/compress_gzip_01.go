package main

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main3472() {
	fmt.Println("-----下面的代码都是为了验证是否可以压缩多个文件------")
	dst_file, err := os.Create("main/gzip/dst2.gz")
	check_err_gzip2(err)
	wt1 := gzip.NewWriter(dst_file)
	wt2:= gzip.NewWriter(dst_file)
	ii:=0

	err=filepath.Walk("main/gzip", func(path string, info os.FileInfo, err error) error {
		// Ext返回路径使用的文件扩展名。
		//扩展名是从path的最后一个元素中的最后一个点开始的后缀； 如果没有点，则为空。
		//info.Name()返回文件名，不包含路径，如下：
		//dst.gz
		//gzip.txt
		//gzip02.txt
		if !info.IsDir() && filepath.Ext(info.Name())==".txt" {
			fmt.Println("当前操作的文件名是：",info.Name())

			//path是相对文件路径
			file, err := os.Open(path)
			check_err_gzip2(err)
			//这个方法比read()好用很多倍，而且不用我们手动创建[]byte来存储读取到的字节数据
			src_byte, err := ioutil.ReadAll(file)
			check_err_gzip2(err)

			if ii != 0{
				i, err := wt2.Write(src_byte)
				check_err_gzip2(err)

				//写不写都没什么所谓，因为后面调用了close()（包含flush()方法）
				//err = wt.Flush()
				//check_err_gzip2(err)
				fmt.Println("写入压缩文件中得到字节数是222：",i)
				err = wt2.Close()
				check_err_gzip2(err)
			}else {
				i, err := wt1.Write(src_byte)
				check_err_gzip2(err)

				//写不写都没什么所谓，因为后面调用了close()（包含flush()方法）
				//err = wt.Flush()
				//check_err_gzip2(err)
				fmt.Println("写入压缩文件中得到字节数是111：",i)
				ii++
				err = wt1.Close()
				check_err_gzip2(err)
			}



		}
		return nil
	})
	check_err_gzip2(err)

	err = dst_file.Close()
	check_err_gzip2(err)

	//验证结果显示，但凡是使用了flate算法的压缩（gzip和zlib压缩就是）都是无法压缩多个文件的，除非将多个文件写入到单个文件中然后再进行压缩，上面的方式就是这样！


	//下面的api个人认为不怎么常用

	//多流控制阅读器是否支持多流文件。
	//如果已启用（默认设置），则Reader希望输入是一系列单独压缩的数据流，每个数据流都有自己的标题和结尾，以EOF结尾。 效果是，压缩文件序列的串联被视为等同于该序列串联的gzip。 这是gzip阅读器的标准行为。
	//调用Multistream（false）将禁用此行为； 当读取区分单个gzip数据流或将gzip数据流与其他数据流区分开的文件格式时，禁用该行为可能很有用。
	//在此模式下，当Reader到达数据流的末尾时，Read返回io.EOF。 基础阅读器必须实现io.ByteReader才能将其定位在gzip流之后。
	//要开始下一个流，请调用z.Reset（r），然后调用z.Multistream（false）。
	//如果没有下一个流，则z.Reset（r）将返回io.EOF。
	//reader.Multistream(true)

	// Reset丢弃Reader z的状态，使其等效于从NewReader读取其原始状态的结果，但改为从r读取。
	//这允许重用Reader，而不是分配新的Reader。
	//reader.Reset()

}

func check_err_gzip2(err error)  {
	if err != nil{
		fmt.Println(err)
	}
}