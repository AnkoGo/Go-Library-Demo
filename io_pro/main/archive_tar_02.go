package main

import (
	"archive/tar"
	"fmt"
	"os"
)

func main() {

	fmt.Println("-----------快速获取一个文化的元数据------------")
	//创建软连接文件，如果已经存在的话则会报错
	//symlink_err := os.Symlink("main/tar_data/text1.txt", "main/tar_02/text1_Symlink.lnk")
	//checkErr_tar(symlink_err)
	//time.Sleep(2e9)
	file, e := os.Open("main/tar_data/text1.txt")
	checkErr_tar(e)
	info, e := file.Stat()
	checkErr_tar(e)
	//wt := tar.NewWriter(file)
	// FileInfoHeader从fi创建一个部分填充的Header。
	//如果fi描述了符号链接，则FileInfoHeader会将链接记录为链接目标。
	//如果fi描述目录，则在名称后添加斜杠。
	//由于os.FileInfo的Name方法仅返回其描述的文件的基本名称，因此可能有必要修改Header.Name以提供文件的完整路径名。
	//说白了就是快速获取一个文化的元数据
	h, e := tar.FileInfoHeader(info, "")
	//h, e := tar.FileInfoHeader(info, "main/tar_data/text1.txt")//似乎第二个参数写跟没写都一样不起作用的！暂时不知道怎么用！
	checkErr_tar(e)
	printHeaderinfo(h)
	//输出：
	//	-----------快速获取一个文化的元数据------------
	//	----下面是获取当前tar文件的header对象的一些信息----
	//	FileInfo().Sys()总信息： &{48 text1.txt  727 438 0 0   2019-10-16 14:58:40.9768512 +0800 CST 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC 0 0 map[] map[] <unknown>}
	//	Typeflag是标题条目的类型： 48
	//	文件入口名称： text1.txt
	//	链接的目标名称：
	//	逻辑文件大小： 727
	//	权限和模式位(二进制表示)：110110110
	//	所有者的用户标识： 0
	//	所有者的组ID： 0
	//	所有者的用户名：
	//	所有者的组名：
	//	修改时间： 2019-10-16 14:58:40.9768512 +0800 CST
	//	访问时间： 0001-01-01 00:00:00 +0000 UTC
	//	更改时间： 0001-01-01 00:00:00 +0000 UTC
	//	主设备号： 0
	//	次设备号： 0
	//	PAX扩展头记录的映射： map[]
	//	tar标头的格式： <unknown>
	//	-----------------------------
	//关于header的更多信息请参考https://blog.csdn.net/lichangrui2009/article/details/85995620



}
func checkErr_tar(err error)  {
	if err != nil{
		fmt.Println(err)
	}
}

func printHeaderinfo(h *tar.Header)  {//这样表示的是*Header而不是*tar
	fmt.Println("----下面是获取当前tar文件的header对象的一些信息----")
	fmt.Println("FileInfo().Sys()总信息：", h.FileInfo().Sys())
	fmt.Println("Typeflag是标题条目的类型：", h.Typeflag) // Typeflag是标题条目的类型。根据Name中是否存在尾部斜杠，零值会自动提升为TypeReg或TypeDir。
	fmt.Println("文件入口名称：", h.Name)               //文件入口名称
	fmt.Println("链接的目标名称：", h.Linkname)          //链接的目标名称（对TypeLink或TypeSymlink有效）
	fmt.Println("逻辑文件大小：", h.Size)               //逻辑文件大小（以字节为单位）
	fmt.Printf("权限和模式位(二进制表示)：%b\n", h.Mode)     //权限和模式位
	fmt.Println("所有者的用户标识：", h.Uid)              //所有者的用户标识
	fmt.Println("所有者的组ID：", h.Gid)               //所有者的组ID
	fmt.Println("所有者的用户名：", h.Uname)             //所有者的用户名
	fmt.Println("所有者的组名：", h.Gname)              //所有者的组名
	fmt.Println("修改时间：", h.ModTime)              //修改时间
	fmt.Println("访问时间：", h.AccessTime)           //访问时间（需要PAX或GNU支持）
	fmt.Println("更改时间：", h.ChangeTime)           //更改时间（需要PAX或GNU支持）
	fmt.Println("主设备号：", h.Devmajor)             //主设备号（对TypeChar或TypeBlock有效）
	fmt.Println("次设备号：", h.Devminor)             //次设备号（对TypeChar或TypeBlock有效）
	fmt.Println("PAX扩展头记录的映射：", h.PAXRecords)    // PAXRecords是PAX扩展头记录的映射。
	fmt.Println("tar标头的格式：", h.Format)           // Format指定tar标头的格式。
	fmt.Println("-----------------------------")

}