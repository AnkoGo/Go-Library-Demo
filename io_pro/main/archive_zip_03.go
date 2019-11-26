package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

func main235() {
	//Reader结构体说明
	//type Reader struct {
	//	r             io.ReaderAt //表示从哪里开始读
	//	File          []*File  //看下面，相当于压缩文件中的1个源数据（或者键一个文件），这个相当于File的管理器
	//	Comment       string	//压缩包的注释
	//	decompressors map[uint16]Decompressor //解压器对象的map
	//}

	//File结构体说明
	//type File struct {
	//	FileHeader		//这个是FileHeader结构体,所以他相当于FileHeader的读取管理器
	//	zip          *Reader	//读取器接口
	//	zipr         io.ReaderAt	//读取接口
	//	zipsize      int64		//压缩包的大小
	//	headerOffset int64		//读取文件的指针
	//}

	//-------------------------第一种声明读取器的方式-------------------------------
	//打开一个zip存档以供阅读。
	// OpenReader将打开按名称指定的Zip文件，并返回ReadCloser。
	//底层是通过os.open以及os.stat和init()3个方法来进行实现的（init方法才是真正的OpenReader，不过跟函数名字不大搭配）
	//r, err := zip.OpenReader("main/zip/zip_02.zip")
	//-----------------------------第2种声明读取器的方式---------------------------
	//上面的方法是下面的方法的封装，多数情况下采用上面的方法会更便捷
	f, err := os.Open("main/zip/zip_02.zip")
	check_err_zip(err)
	fi, err := f.Stat()
	if err != nil {
		fmt.Println(err)
		f.Close()//如果出错直接关闭并且停止main函数
		return
	}
	// NewReader从r返回一个新的Reader读数，假定具有给定的字节大小（zip的字节大小）。
	r, err := zip.NewReader(f, fi.Size())//os.File对象实现了io.ReaderAt()方法
	if err != nil {
		log.Fatal(err)
	}
	//--------------------------------------------------------
	//因为上面提供了2中方式的创建压缩读取器，所以下面的关闭也展示了2中不同的关闭方式，除此之外无他区别了！
	//close关闭Zip文件，使其无法用于I/O。
	//底层是调用了File.close()方法，下面是这个方法的说明。说明这个方法会关闭处于打开状态的的压缩文件和压缩文件里面的文件。
	// Close关闭文件，使其无法用于I / O。
	//在支持SetDeadline的文件上，所有挂起的I / O操作都将被取消并立即返回错误。
	//如果已调用Close，则将返回错误。
	//defer r.Close()//这是ReadCloser对象才需要写这句代码，如果是Reader的话是不需要写的

	defer f.Close()//这句话就相当于上面的底层的主要实现，这句代码是Reader才需要写的

	//遍历档案中的文件，打印其中的一些内容。
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		// DataOffset返回文件的可能经过压缩的数据（压缩包里面的文件不是绝对会压缩数据的）相对于zip文件开头的偏移量。
		//大多数调用者应该改用Open，它透明地解压缩数据并验证校验和。
		//主要是通过findBodyOffset()方法实现的
		fmt.Println(f.DataOffset())

		// Open返回一个ReadCloser，它提供对文件内容的访问。
		//可以同时读取多个文件。
		//底层实现主要方法是io.NewSectionReader，findBodyOffset()，zip.decompressor方法,
		// open根本不是通过DataOffset（）方法来实现的，取而代之的是findBodyOffset()方法（DataOffset（）的底层就是通过这个方法实现的）
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		// CopyN将n个字节（或直到出错）从src复制到dst。
		//返回复制的字节数以及复制时遇到的最早错误。
		//返回时，当且仅当err == nil时才writter== n。
		//如果dst实现了ReaderFrom接口，则使用该接口实现副本。
		//底层是通过Copy(dst, LimitReader(src, n))来实现的
		_, err = io.CopyN(os.Stdout, rc, f.FileInfo().Size())//基本在Fileheader里面的方法都是可以在File中调用的
		if err != nil {
			log.Fatal(err)
		}
		rc.Close()//关闭io的读取关闭器，注意不是关不压缩文件读取关闭器
		fmt.Println()
		//输出：
		//	Contents of text1.txt:
		//	39 <nil>，说明从第39个字节开始是text1.txt的数据
		//	a1111111adsdsdsd李克强在致辞中代表中国政府和人民祝贺本届世园会成功举办，并对支持和参与北京世园会的各国朋友致以谢意。他表示，本届世园会以“绿色生活，美丽家园”为主题，精彩纷呈、成果丰硕。
		//	在开幕式上，中国国家主席习近平倡导共同建设美丽地球家sdsdsdsd园、构建人类命运共同体。这是一场文明互鉴的绿色盛会，促进了各国文明交流、民心相通和绿色合作。这是一场创新荟萃的科技盛会，展现了绿色科技应用的美好前景。
		//	这是一场走进自然的体验盛会，中外访客用心感受环保与发展相互促进、人与自然和谐共处的美好。sdsdsds
		//	Contents of text2.txt:
		//	586 <nil>，说明从第586个字节开始是text2.txt的数据
		//	2222222222“保持中美关aaabbbb系健康稳定发展，对两国对世界都有利。双方要坚持协abc调、合作、稳定的基调，在相互尊重基础上管控分歧，在互惠互利基础上拓展合作，推动两国关系沿着正确轨道向前发展。”
		//	习近平主席在致特朗普总统的口信中强调了中方主张，特朗普总统请刘鹤副总理转达对习近平主席的感谢，并明确表示“美中经贸磋商取得了实质性的第一阶段成果，这对美中两国和世界都是重大利好”。abc

	}
}

