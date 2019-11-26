package main

import (
	"fmt"
	"log"
	"os"
)

func main0989() {
	file, err := os.Open("main/test.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 101)
	count, err := file.Read(data) //读取一个字节切片并返回字节数
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
	fmt.Println(os.Hostname())             //返回主机名
	fmt.Println(os.Getpagesize())          //返回内存页的尺寸
	for key, value := range os.Environ() { //返回环境变量的切片
		fmt.Println(key, ":", value)
	}
	fmt.Println(os.Getenv("GOROOT")) //获取环境变量的key的value值
	os.Setenv("my", "GOGO")          //即使存在也不会报错，如果不存在该环境变量会返回空字符串
	fmt.Println(os.Getenv("my"))
	//Clearenv()//清空所有的环境变量，前面记得加上os
	//defer fmt.Println("22222222222222222")
	//Exit让当前程序以给出的状态码code退出。一般来说，状态码0表示成功，
	// 非0表示出错。程序会立刻终止，defer的函数不会被执行。
	//os.Exit(0)
	//fmt.Println("11111111111111111")
	fmt.Println(os.Geteuid())
	fmt.Println(os.Getuid())
	fmt.Println(os.Getgid())
	fmt.Println(os.Getgroups()) //[] getgroups: not supported by windows
	fmt.Println(os.Getpid())
	fmt.Println(os.Getppid()) //他的父进程是goland进程
	var fm os.FileMode
	//fm=os.ModeDir

	fm = os.ModePerm
	fmt.Println(fm.IsDir())
	fmt.Println(fm.IsRegular())
	fmt.Println(fm.Perm())
	fmt.Println(fm.String())
	fmt.Println("=====================")
	fmt.Println(os.Getwd())

	var ModePerm111 os.FileMode = 0555
	os.Mkdir("cc1", ModePerm111)
	os.MkdirAll(`cc2\cc3\cc4`, ModePerm111)
	os.Rename("aa.txt", "aa1.txt")
	os.Remove("aa2.txt")
	os.RemoveAll("cc2")
	os.Symlink("main/test.txt", "test1.txt") //创建软连接
	os.Link("main/test.txt", "test2.txt")    //在windows中没有硬链接的跟原文件保持更改同步
	fmt.Println(os.Readlink(`test1.txt`))
	fmt.Println(os.TempDir()) //C:\Users\ADMINI~1\AppData\Local\Temp

	os.Create("test3.txt")
	//fl,_:=os.Open("main/test.txt")
	fl, _ := os.OpenFile("main/test.txt", 0, 0777) //如果文件不存在不会报错
	data111 := make([]byte, 50)
	n, _ := fl.Read(data111)
	fmt.Println(string(data111[:n]))
	fmt.Printf("%q", data111[:n])
	fmt.Println()
	fmt.Printf(fl.Name()) //返回的是相对的路径名字

	fmt.Println("-------------------")

	//具备多个权限，就把相应的 4、2、1 相加就可以了：
	//若要 rwx 则 4+2+1=7
	//若要 rw- 则 4+2=6
	//若要 r-x 则 4+1=5
	//若要 r-- 则 =4
	//若要 -wx 则 2+1=3
	//若要 -w- 则 =2
	//若要 --x 则 =1
	//若要 --- 则 =0
	//os.Chmod("main/test111.txt",0555)
	//无论下面以什么权限打开，只要上面的那句改变了文件的访问读写权限的话，那么你就无法写入文件的
	f2, _ := os.OpenFile("main/test111.txt", os.O_CREATE, 0555) //创建
	//f2,_:=os.OpenFile("main/test111.txt",os.O_APPEND,0777)//追加
	data222 := []byte{97, 98, 99, 228, 188, 159}
	n1, err333 := f2.Write(data222) //将字节切片中的东西写到文件中去，不管有没有这个文件了都会覆盖原文件
	if err333 != nil {
		fmt.Println(err333)
	}
	fmt.Println(n1)
	fmt.Println(data222)
	fmt.Printf("%q\n", data222[:n1]) //%q会将字节按照解码成utf8显示出来，但是并不是说将字节原地解码了

	fmt.Println(f2.Fd()) //每次运行的Unix文件描述符都不一样的，随机分配的
	fmt.Println(f2.Stat()) //文件对象的标志识别码0x4afe20
	//fmt.Printf("%p",f2)

	f3, _ := os.OpenFile("main", os.O_RDONLY, 0777)
	finfo, _ := f3.Readdir(-2) //<0读取剩余的，>0读取n个
	fmt.Println(finfo)
	fmt.Println(finfo[0].Name(), finfo[1].Name())       //返回文件名
	fmt.Println(finfo[0].IsDir(), finfo[1].Mode())      //false -rw-rw-rw-
	fmt.Println(finfo[0].Size(), "---", finfo[1].Sys()) //173 --- &{32 {293080951 30766355} {1047502017 30766413} {1047502017 30766413} 0 3716}

	finfo111, _ := os.Stat("main")
	//main
	//2019-09-28 00:09:56.4975268 +0800 CST
	fmt.Println(finfo111.Name())
	fmt.Println(finfo111.ModTime()) //文件或者文件目录修改时间

	fmt.Println(os.IsPathSeparator(92)) // 92是\的unicode值的十进制，这里似乎只能传十进制
	fmt.Println(os.IsPathSeparator(93))

	fmt.Println("````````````")
	dir_Exit := os.IsExist(os.ErrInvalid)
	fmt.Println(dir_Exit)
	dir_Exit111 := os.IsExist(os.ErrExist)
	fmt.Println(dir_Exit111)
	dir_Exit222 := os.IsExist(os.ErrNotExist)
	fmt.Println(dir_Exit222)

	fmt.Println("````````````")

	err00 := os.Chmod("main/test222.txt", 0222) //不大明白为什么只有0000会导致权限问题
	//time.Sleep(2e9)
	f4, err444 := os.OpenFile("main/test222.txt", os.O_RDONLY, 0777) //创建
	if err00 != nil {
		fmt.Println(err00)
	}
	if err444 != nil {
		fmt.Println(err444)
	}
	if os.IsPermission(err444) {
		fmt.Println("IsPermission")
	}
	data333 := []byte{97, 98, 99, 228, 188, 159}
	n4, err444 := f4.Write(data333) //将字节切片中的东西写到文件中去，不管有没有这个文件了都会覆盖原文件
	if err444 != nil {
		fmt.Println(err444)
	}
	fmt.Println(n4)

	fmt.Println("````````````")
	f5, _ := os.OpenFile("main", os.O_RDONLY, 0777)
	fi, _ := f5.Readdirnames(-2) //<0读取剩余的，>0读取n个
	fmt.Println(fi)

	fmt.Println("````````````")
	//必须是这种含有读功能的flag才可以进行截断,比如创建，写，读写，但是O_APPEND不行，因为他不具有读的功能
	f6, _ := os.OpenFile("main/test222.txt", os.O_RDWR, 0777)
	//err666:=f6.Truncate(2)
	//if err666 != nil{
	//	fmt.Println(err666)
	//}
	data444 := make([]byte, 4)
	//n7,_:=f6.Read(data444)//从文件开头开始读
	//1234567....会变成这样--->12 AAA,不知道为什么前面会加个空格
	n7, _ := f6.ReadAt(data444, 4) //文件数数从1开始的
	fmt.Printf("%q\n", data444[:n7])
	//f6.Close()
	data555 := []byte{66, 66, 66}
	n8, err888 := f6.WriteAt(data555, 3)
	if err888 != nil {
		fmt.Println(err888)
	}
	fmt.Println(n8)
	fmt.Println("````````````")
	//n9,err999:=f6.Seek(5,0)
	//f6.WriteString("你好啊")//默认是从开头写，会覆盖原文件,除非seek()
	//
	//if err999 !=nil{
	//	fmt.Println(err999)
	//}
	//
	//fmt.Println(n9)

	fmt.Println("````````````")
	//err01:=f6.Sync()//不知道这个怎么用
	//if err01 !=nil{
	//	fmt.Println(err01)
	//}
	p1, err02 := os.FindProcess(7492) //获取不到相应pid的程序就会报错
	if err02 != nil {
		fmt.Println(p1)
	}
	fmt.Println(p1) //&{1244 292 0 {{0 0} 0 0 0 0}}
	//p1.Kill()
	//fmt.Println(p1.Release())//不知道怎么用
	//fmt.Println(p1.Signal())//不知道怎么用
	fmt.Println(p1.Pid)
	PS, err03 := p1.Wait()
	if err03 != nil {
		fmt.Println(err03)
	}
	fmt.Println("********")
	fmt.Println(PS.Pid()) //Pi返回一个已退出的进程的进程id。6244
	//Sys返回该已退出进程系统特定的退出信息。
	//需要将其类型转换为适当的底层类型，如Unix里转换为*syscall.WaitStatus类型以获取其内容。
	fmt.Println(PS.Sys())        //{0}
	fmt.Println(PS.String())     //exit status 0,这个是对退出码的描述
	fmt.Println(PS.Exited())     //Exited报告进程是否已退出。如true
	fmt.Println(PS.Success())    //Success报告进程是否成功退出，如在Unix里以状态码0退出。如true
	fmt.Println(PS.SystemTime()) //SystemTime返回已退出进程及其子进程耗费的系统CPU时间。如1.4375s，这个时间比下面的用户的时间要大
	//SysUsage返回关于退出进程的系统独立的资源使用信息。主要用来获取进程使用系统资源
	fmt.Println(PS.SysUsage()) //&{{3706799625 30766427} {4021950816 30766427} {14375000 0} {6406250 0}}
	fmt.Println(PS.UserTime()) //UserTime返回已退出进程及其子进程耗费的用户CPU时间。如640.625ms
	// ExitCode返回已退出进程的退出代码，或-1
	//如果进程尚未退出或被信号终止。
	fmt.Println(PS.ExitCode()) //0
	fmt.Println()              //等待这个p1退出才解阻塞

}
