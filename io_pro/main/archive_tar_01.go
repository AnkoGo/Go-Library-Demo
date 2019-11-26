package main

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func check_tar_err(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
func ReadTar() {
	//下面是读取压缩文件中的内容
	//打开tar压缩文件以供阅读。
	file888, e_tar888 := os.OpenFile("main/dst_tar/text.tar", os.O_RDWR, 777)
	check_tar_err(e_tar888)

	tr := tar.NewReader(file888) //file本身实现了这个接口io.Reader
	//遍历压缩文件中的文件。
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			// end of tar archive
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("----下面是通过header对象获取os.FileInfo对象的一些信息----")
		fmt.Println("--------FileInfo对象:", hdr.FileInfo())
		fmt.Println("--------FileInfo对象.Name():", hdr.FileInfo().Name())
		fmt.Println("--------FileInfo对象.Size():", hdr.FileInfo().Size())
		fmt.Println("--------FileInfo对象.Mode():", hdr.FileInfo().Mode())
		fmt.Println("--------FileInfo对象.ModTime():", hdr.FileInfo().ModTime())
		fmt.Println("--------FileInfo对象.IsDir():", hdr.FileInfo().IsDir())
		fmt.Println("--------FileInfo对象.Sys():", hdr.FileInfo().Sys())
		fmt.Println()

		fmt.Println("----下面是获取当前tar文件的header对象的一些信息----")
		fmt.Println("Typeflag是标题条目的类型：", hdr.Typeflag) // Typeflag是标题条目的类型。根据Name中是否存在尾部斜杠，零值会自动提升为TypeReg或TypeDir。
		fmt.Println("文件入口名称：", hdr.Name)               //文件入口名称
		fmt.Println("链接的目标名称：", hdr.Linkname)          //链接的目标名称（对TypeLink或TypeSymlink有效）
		fmt.Println("逻辑文件大小：", hdr.Size)               //逻辑文件大小（以字节为单位）
		fmt.Printf("权限和模式位(二进制表示)：%b\n", hdr.Mode)     //权限和模式位
		fmt.Println("所有者的用户标识：", hdr.Uid)              //所有者的用户标识
		fmt.Println("所有者的组ID：", hdr.Gid)               //所有者的组ID
		fmt.Println("所有者的用户名：", hdr.Uname)             //所有者的用户名
		fmt.Println("所有者的组名：", hdr.Gname)              //所有者的组名
		fmt.Println("修改时间：", hdr.ModTime)              //修改时间
		fmt.Println("访问时间：", hdr.AccessTime)           //访问时间（需要PAX或GNU支持）
		fmt.Println("更改时间：", hdr.ChangeTime)           //更改时间（需要PAX或GNU支持）
		fmt.Println("主设备号：", hdr.Devmajor)             //主设备号（对TypeChar或TypeBlock有效）
		fmt.Println("次设备号：", hdr.Devminor)             //次设备号（对TypeChar或TypeBlock有效）
		fmt.Println("PAX扩展头记录的映射：", hdr.PAXRecords)    // PAXRecords是PAX扩展头记录的映射。
		fmt.Println("tar标头的格式：", hdr.Format)           // Format指定tar标头的格式。
		fmt.Println("-----------------------------")

		fmt.Printf("%s的内容是:\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatalln(err)
		}
		fmt.Println()
	}
}

func WriteTar() {
	tarfile, e_tar999 := os.OpenFile("main/dst_tar/text.tar", os.O_CREATE, 777)
	check_tar_err(e_tar999)

	//创建一个新的tar压缩文件写入器。
	tw := tar.NewWriter(tarfile)

	//创建结构体存储读取出来的2个文件的信息
	type files struct {
		Name, Body string
		Uid        int
		Mode       int64
		Size       int64
		ModTime    time.Time
		AccessTime time.Time
	}
	var f_files_ls [2]files//这里千万注意，因为你上面创建了新的东西，所以你这里为2的话会报错。
	// 我们将新的东西创建在不同的文件夹以阻止这个错误的发生，当然你也可以通过将创建新文件的代码放到下面。这样的话就不会产生错误了

	i:=0
	var Wf = func(path string, info os.FileInfo, err error) error {
		//如果是文件
		if !info.IsDir() {
			f, err_tar1 := os.Open(path)
			check_tar_err(err_tar1)
			f_dst_byte := make([]byte, info.Size())
			n, Read_f1_err := f.Read(f_dst_byte)
			check_tar_err(Read_f1_err)
			fmt.Println("f1中的字节数是：", n)


			//将打开的2个文件的名字和内容分别存储到对应的结构体列表中去！但是这2个东西的有些信息是一样的，如果你希望不一样的话，那么你可以通过if判断来产生分支
			f_files_one := files{
					Name: info.Name(),
					Body: string(f_dst_byte),
					Uid:  10,
					Mode: int64(info.Mode()),
					Size: info.Size(),
					//ModTime:f1_Info.ModTime(),
					//AccessTime:time.Now(),
					ModTime:    time.Date(2019, 11, 11, 11, 11, 11, 11, time.Local),
					AccessTime: time.Date(2019, 11, 12, 11, 11, 11, 11, time.Local), //虽然我们这里设置了，但是好像看系统的才能起作用的
			}
			f_files_ls[i]=f_files_one//注意i要维护在外围，否则当函数结束的时候i的计数就会废除
			i++
		}
		return nil
	}
	//var walkF filepath.WalkFunc("test",Finfo,err1)//事实上这里我们并不需要自己创建WalkFunc类对象，因为这个filepath.WalkFunc（）
	//的基类是func类，所以我们直接创建func类即可
	walkErr := filepath.Walk("main/tar_data", Wf)
	check_tar_err(walkErr)

	fmt.Printf("---%#v",f_files_ls)
	for _, file := range f_files_ls {
		//按照file中的信息来新建tar.Header元数据
		hdr := &tar.Header{
			Name:       file.Name,
			Size:       int64(len(file.Body)),
			Uid:        file.Uid,
			Mode:       file.Mode,
			ModTime:    file.ModTime,
			AccessTime: file.AccessTime,
		}
		//写入器将元数据写入到新建的压缩文件中去
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatalln(err)
		}

		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatalln(err)//需要注意的是写入后的文件的的最小的大小是4096个字节，似乎打包后的文件更加大，不知道为什么
		}
	}
	//关闭写入器，检查错误。
	if err := tw.Close(); err != nil {
		log.Fatalln(err)
	}



}

func main45762() {

	fmt.Println("---------------------将文件进行压缩并且写入压缩文件------------------------")
	WriteTar()

	fmt.Println("----------读取压缩文件中的内容------------")
	ReadTar()
	//输出：
	//	---------------------将文件进行压缩并且写入压缩文件------------------------
	//	f1中的字节数是： 727
	//	f1中的字节数是： 550
	//	---[2]main.files{main.files{Name:"text1.txt", Body:"\ufeff1111111adsdsdsd李克强在致辞中代表中国政府和人民祝贺本届世园会成功举办，并对支持和参与北京世园会的各国朋友致以谢意。他表示，本届世园会以“绿色生活，美丽家园”为主题，精彩纷呈、成果丰硕。\r\n在开幕式上，中国国家主席习近平倡导共同建设美丽地球家sdsdsdsd园、构建人类命运共同体。这是一场文明互鉴的绿色盛会，促进了各国文明交流、民心相通和绿色合作。这是一场创新荟萃的科技盛会，展现了绿色科技应用的美好前景。\r\n这是一场走进自然的体验盛会，中外访客用心感受环保与发展相互促进、人与自然和谐共处的美好。sdsdsds", Uid:10, Mode:438, Size:727, ModTime:time.Time{wall:0xb, ext:63709038671, loc:(*time.Location)(0x5cef40)}, AccessTime:time.Time{wall:0xb, ext:63709125071, loc:(*time.Location)(0x5cef40)}}, main.files{Name:"text2.txt", Body:"2222222222“保持中美关aaabbbb系健康稳定发展，对两国对世界都有利。双方要坚持协abc调、合作、稳定的基调，在相互尊重基础上管控分歧，在互惠互利基础上拓展合作，推动两国关系沿着正确轨道向前发展。”\r\n习近平主席在致特朗普总统的口信中强调了中方主张，特朗普总统请刘鹤副总理转达对习近平主席的感谢，并明确表示“美中经贸磋商取得了实质性的第一阶段成果，这对美中两国和世界都是重大利好”。abc", Uid:10, Mode:438, Size:550, ModTime:time.Time{wall:0xb, ext:63709038671, loc:(*time.Location)(0x5cef40)}, AccessTime:time.Time{wall:0xb, ext:63709125071, loc:(*time.Location)(0x5cef40)}}}----------读取压缩文件中的内容------------
	//	----下面是通过header对象获取os.FileInfo对象的一些信息----
	//	--------FileInfo对象: {0xc0000ca000}
	//	--------FileInfo对象.Name(): text1.txt
	//	--------FileInfo对象.Size(): 727
	//	--------FileInfo对象.Mode(): -rw-rw-rw-
	//	--------FileInfo对象.ModTime(): 2019-11-11 11:11:11 +0800 CST
	//	--------FileInfo对象.IsDir(): false
	//	--------FileInfo对象.Sys(): &{48 text1.txt  727 438 10 0   2019-11-11 11:11:11 +0800 CST 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC 0 0 map[] map[] USTAR}
	//
	//	----下面是获取当前tar文件的header对象的一些信息----
	//	Typeflag是标题条目的类型： 48
	//	文件入口名称： text1.txt
	//	链接的目标名称：
	//	逻辑文件大小： 727
	//	权限和模式位(二进制表示)：110110110，十进制是438，110分别代表的意思是可读，可写，不可执行
	//	所有者的用户标识： 10
	//	所有者的组ID： 0
	//	所有者的用户名：
	//	所有者的组名：
	//	修改时间： 2019-11-11 11:11:11 +0800 CST
	//	访问时间： 0001-01-01 00:00:00 +0000 UTC
	//	更改时间： 0001-01-01 00:00:00 +0000 UTC
	//	主设备号： 0
	//	次设备号： 0
	//	PAX扩展头记录的映射： map[]
	//	tar标头的格式： USTAR
	//	-----------------------------
	//	text1.txt的内容是:
	////1111111adsdsdsd李克强在致辞中代表中国政府和人民祝贺本届世园会成功举办，并对支持和参与北京世园会的各国朋友致以谢意。他表示，本届世园会以“绿色生活，美丽家园”为主题，精彩纷呈、成果丰硕。
	//	在开幕式上，中国国家主席习近平倡导共同建设美丽地球家sdsdsdsd园、构建人类命运共同体。这是一场文明互鉴的绿色盛会，促进了各国文明交流、民心相通和绿色合作。这是一场创新荟萃的科技盛会，展现了绿色科技应用的美好前景。
	//	这是一场走进自然的体验盛会，中外访客用心感受环保与发展相互促进、人与自然和谐共处的美好。sdsdsds
	//	----下面是通过header对象获取os.FileInfo对象的一些信息----
	//	--------FileInfo对象: {0xc0000ca1c0}
	//	--------FileInfo对象.Name(): text2.txt
	//	--------FileInfo对象.Size(): 550
	//	--------FileInfo对象.Mode(): -rw-rw-rw-
	//	--------FileInfo对象.ModTime(): 2019-11-11 11:11:11 +0800 CST
	//	--------FileInfo对象.IsDir(): false
	//	--------FileInfo对象.Sys(): &{48 text2.txt  550 438 10 0   2019-11-11 11:11:11 +0800 CST 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC 0 0 map[] map[] USTAR}
	//
	//	----下面是获取当前tar文件的header对象的一些信息----
	//	Typeflag是标题条目的类型： 48
	//	文件入口名称： text2.txt
	//	链接的目标名称：
	//	逻辑文件大小： 550
	//	权限和模式位(二进制表示)：110110110，十进制是438，110分别代表的意思是可读，可写，不可执行
	//	所有者的用户标识： 10
	//	所有者的组ID： 0
	//	所有者的用户名：
	//	所有者的组名：
	//	修改时间： 2019-11-11 11:11:11 +0800 CST
	//	访问时间： 0001-01-01 00:00:00 +0000 UTC
	//	更改时间： 0001-01-01 00:00:00 +0000 UTC
	//	主设备号： 0
	//	次设备号： 0
	//	PAX扩展头记录的映射： map[]
	//	tar标头的格式： USTAR
	//	-----------------------------
	//	text2.txt的内容是:
	//	2222222222“保持中美关aaabbbb系健康稳定发展，对两国对世界都有利。双方要坚持协abc调、合作、稳定的基调，在相互尊重基础上管控分歧，在互惠互利基础上拓展合作，推动两国关系沿着正确轨道向前发展。”
	//	习近平主席在致特朗普总统的口信中强调了中方主张，特朗普总统请刘鹤副总理转达对习近平主席的感谢，并明确表示“美中经贸磋商取得了实质性的第一阶段成果，这对美中两国和世界都是重大利好”。abc

}
