package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	levelFlag = flag.Int("level", 11, "级别")
	bnFlag int
)

func init() {
	flag.IntVar(&bnFlag, "bn", 12, "份数")
}

func main() {

	flag.Parse()
	// Args保留以程序名称开头的命令行参数。按照空格进行分割参数
	count := len(os.Args)
	fmt.Println("参数总个数:",count)

	fmt.Println("参数详情:")
	for i := 0 ; i < count ;i++{
		fmt.Println(i,":",os.Args[i])
	}

	fmt.Println("\n参数值:")
	fmt.Println("级别:", *levelFlag)
	fmt.Println("份数:", bnFlag)
}


/*
运行结果:(compress_zlib.go改为相应的文件名即可)

C:\Users\Administrator\Desktop\go_pro\src\io_pro\main3>go run compress_zlib.go
参数总个数: 1
参数详情:
0 : C:\Users\Administrator\AppData\Local\Temp\___go_build_compress_zlib_go.exe

参数值:
级别: 11
份数: 12


C:\Users\Administrator\Desktop\go_pro\src\io_pro\main3>go run compress_zlib.go  -level 111 -bn=122
参数总个数: 4
参数详情:
0 : C:\Users\ADMINI~1\AppData\Local\Temp\go-build779318678\b001\exe\compress_zlib.exe
1 : -level
2 : 111
3 : -bn=122

参数值:
级别: 111
份数: 122


C:\Users\Administrator\Desktop\go_pro\src\io_pro\main3>go run compress_zlib.go  -level 111 -bn 122
参数总个数: 5
参数详情:
0 : C:\Users\ADMINI~1\AppData\Local\Temp\go-build387111486\b001\exe\compress_zlib.exe
1 : -level
2 : 111
3 : -bn
4 : 122

参数值:
级别: 111
份数: 122



*/