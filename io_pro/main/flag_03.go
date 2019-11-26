package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func init() {
	flagSet.Var(&addrFlag, "a", "b")
}

type StringArray []string

func (s *StringArray) String() string {
	return fmt.Sprint([]string(*s))
}

func (s *StringArray) Set(value string) error {
	*s = append(*s, value)//可以看到这里确实可以存储多个值，而不是替代旧值！
	return nil
}

var (
	/*
		参数解析出错时错误处理方式
		switch f.errorHandling {
			case ContinueOnError:
				return err
			case ExitOnError:
				os.Exit(2)
			case PanicOnError:
				panic(err)
			}
	*/

	//flagSet = flag.NewFlagSet(os.Args[0],flag.PanicOnError)//以抛出异常的形式处理错误
	flagSet = flag.NewFlagSet(os.Args[0],flag.ExitOnError)//以退出整个程序的形式处理错误
	//flagSet = flag.NewFlagSet("xcl",flag.ExitOnError)
	verFlag = flagSet.String("ver", "", "version")//版本号，这里的值是一个空字符串，表示没给！
	xtimeFlag  = flagSet.Duration("time", 10*time.Minute, "time Duration")//发行时间

	addrFlag = StringArray{}//联系地址
)




func main345792() {
	fmt.Println("os.Args[0]:", os.Args[0])

	fmt.Println("-------------")
	fmt.Println(flagSet.NFlag())
	fmt.Println(flagSet.Arg(1))
	fmt.Println(flagSet.Arg(2))
	fmt.Println(flagSet.Args())
	fmt.Println("-------------")

	flagSet.Parse(os.Args[1:]) //flagSet.Parse(os.Args[0:])，从第一个参数开始才是要被解析的参数！

	fmt.Println("-------------")
	fmt.Println(flagSet.NFlag())
	fmt.Println(flagSet.Arg(1))
	fmt.Println(flagSet.Arg(2))
	fmt.Println(flagSet.Args())
	fmt.Println("-------------")
	//以上会输出：
	//	-------------
	//	0
	//
	//
	//	[]
	//	-------------
	//	-------------
	//	3
	//
	//
	//	[]
	//	-------------
	//  还是不大懂这几个方法到底是怎么用的！

	// NFlag返回已设置的标志数。
	fmt.Println("当前命令行参数类型个数（NFlag()）:", flagSet.NFlag())
	// NArg()是在处理标志flag后剩余的参数个数。
	for i := 0; i != flagSet.NArg(); i++ {
		// Arg返回第i个命令行参数。 Arg（0）是标志已处理后的第一个剩余参数。 如果请求的元素不存在，则Arg返回一个空字符串。
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}

	fmt.Println("\n参数值:")
	fmt.Println("ver:", *verFlag)
	fmt.Println("xtimeFlag:", *xtimeFlag)
	fmt.Println("addrFlag:",addrFlag.String())
	// Args返回非标志命令行参数。
	for i,param := range flag.Args(){
		fmt.Printf("---#%d :%s\n",i,param)
	}
}



/*
运行结果:
C:\Users\Administrator\Desktop\go_pro\src\io_pro\main3>go run compress_zlib.go -ver 9.0 -a ba -a ca -a d2 -ver 10.0 -time 2m0s
os.Args[0]: C:\Users\ADMINI~1\AppData\Local\Temp\go-build062115438\b001\exe\compress_zlib.exe
当前命令行参数类型个数（NFlag()）: 3

参数值:
ver: 10.0
xtimeFlag: 2m0s
addrFlag: [ba ca d2]

C:\Users\Administrator\Desktop\go_pro\src\io_pro\main3>go run compress_zlib.go -ver 9.0 -a ba -a ca -a d2 -ver 10.0
os.Args[0]: C:\Users\ADMINI~1\AppData\Local\Temp\go-build887043822\b001\exe\compress_zlib.exe
当前命令行参数类型个数（NFlag()）: 2

参数值:
ver: 10.0
xtimeFlag: 10m0s
addrFlag: [ba ca d2]

C:\Users\Administrator\Desktop\go_pro\src\io_pro\main3>go run compress_zlib.go
os.Args[0]: C:\Users\ADMINI~1\AppData\Local\Temp\go-build940038518\b001\exe\compress_zlib.exe
当前命令行参数类型个数（NFlag()）: 0

参数值:
ver:
xtimeFlag: 10m0s
addrFlag: []

C:\Users\Administrator\Desktop\go_pro\src\io_pro\main3>

//下面是一些出错时候的错误的处理方式


-- flagSet = flag.NewFlagSet(os.Args[0],flag.PanicOnError) 结果:
C:\TEMP\testflag>go run tfs.go -ver 9.0 -a ba -a ca -a d2 -ver 10.0 -time 2m0s33
os.Args[0]: C:\DOCUME~1\ADMINI~1\LOCALS~1\Temp\go-build841833143\command-line-arguments\_obj\exe\tfs.exe
invalid value "2m0s33" for flag -time: time: missing unit in duration 2m0s33
Usage of C:\DOCUME~1\ADMINI~1\LOCALS~1\Temp\go-build841833143\command-line-arguments\_obj\exe\tfs.exe:
  -a=[]: b
  -time=10m0s: time Duration
  -ver="": version
panic: invalid value "2m0s33" for flag -time: time: missing unit in duration 2m0s33

goroutine 1 [running]:
flag.(*FlagSet).Parse(0x10b18180, 0x10b42008, 0xc, 0xc, 0x0, 0x0)
        c:/go/src/flag/flag.go:814 +0xee
main.main()
        C:/TEMP/testflag/tfs.go:41 +0x163
exit status 2


-- flagSet = flag.NewFlagSet(os.Args[0],flag.ExitOnError) 结果:
C:\TEMP\testflag>go run tfs.go -ver 9.0 -a ba -a ca -a d2 -ver 10.0 -time 2m0s33
os.Args[0]: C:\DOCUME~1\ADMINI~1\LOCALS~1\Temp\go-build501686683\command-line-arguments\_obj\exe\tfs.exe
invalid value "2m0s33" for flag -time: time: missing unit in duration 2m0s33
Usage of C:\DOCUME~1\ADMINI~1\LOCALS~1\Temp\go-build501686683\command-line-arguments\_obj\exe\tfs.exe:
  -a=[]: b
  -time=10m0s: time Duration
  -ver="": version
exit status 2


*/