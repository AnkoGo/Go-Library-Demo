package main

import (
	"fmt"
)

//log包日志实现一个简单的日志包。 它定义了Logger类型，并带有格式化输出的方法。
// 它还具有可通过帮助函数Print [f | ln]，Fatal [f | ln]和Panic [f | ln]访问的预定义“标准”记录器，比手动创建记录器更容易使用。
//该记录器将写入标准错误，并打印每个记录（log）的消息的日期和时间。
//每条日志消息都在单独的行上输出：如果要打印的消息未以换行符结尾，则记录器将添加一行。
//写入日志消息后，Fatal(致命)函数将调用os.Exit（1）。
// Panic函数在写入日志消息后调用panic。


func main34793() {
	fmt.Println("-----------log包下的函数--------------")
	//Fatal(致命)等同于Print（），然后调用os.Exit（1）。
	//log.Fatal("log.Fatal。。。")
	//输出：
	//2019/10/29 16:01:31 log.Fatal。。。

	//这里之所以显示unreachable code（无法访问的代码是因为上面已经退出去了os.Exit（1））
	//所以我们必须将上面先注释掉，下同
	//str:="log.Fatal。。。"
	// Fatalf等效于Printf（），然后调用os.Exit（1）。
	//log.Fatalf("format:%v",str)
	//输出：
	//2019/10/29 16:08:24 format:log.Fatal。。。

	// Fatalln等效于Println（），然后调用os.Exit（1）。
	//log.Fatalln(str)
	//输出：
	//2019/10/29 16:09:17 log.Fatal。。。




	// Logger表示一个活动的日志记录对象，该对象生成到io.Writer的输出行。
	// 每个记录操作都会调用Writer的Write方法。 一个Logger可以同时从多个goroutines中使用。
	// 它保证序列化对Writer的访问。
	//type Logger struct {
	//	mu     sync.Mutex //确保原子写入； 保护以下字段
	//	prefix string     //在每行开头写入的前缀
	//	flag   int        //属性
	//	out    io.Writer  //输出目标
	//	buf    []byte     //用于累积要写入的文本
	//}

	file, e := os.Create("main3/log/log.txt")
	check_err_log(e)
	// New创建一个新的Logger。 out变量(第一个参数)设置将日志数据写入的目的地。
	// prefix前缀出现在每个生成的日志行的开头。
	// flag参数定义日志记录属性。
	logger := log.New(file, "p303", log.Ldate)//这里我们也可以写入到标准输出控制台上面去，像底层一样！
	//Flags标志返回记录器的输出标志。即Logger结构体中的字段flag的值(int类型)！而这个值就是log.New（）中的第三个参数的值（log.LstdFlags），这个值等于3！
	fmt.Println(logger.Flags())
	fmt.Println(log.Flags())//这里始终为3，底层默认有个logger := log.New(os.Stderr, "", log.LstdFlags)，而不用我们手动创建！
	logger.Fatal("这是一条日志")
	//输出：
	//1
	//3
	//同时查看main3/log/log.txt会发现出现了一条日志信息：
	//p3032019/10/29 这是一条日志


	logger1 := log.New(os.Stderr, "p303\t", log.Ldate)
	Fatalln等效于l.Println（），然后调用os.Exit（1）。
	logger1.Fatalln("这是一条日志(Fatalln)")
	//输出：
	//p303	2019/10/29 这是一条日志(Fatalln)


	str1:="一条日志（Fatalf）"
	logger1.Fatalf("这是%v",str1)
	//输出：
	//p303	2019/10/29 这是一条日志（Fatalf）


	//Print调用l.Output打印到记录器。
	//参数以fmt.Print的方式处理。
	logger1.Print("这是一条打印到控制台的日志(Print)")
	//输出：
	//p303	2019/10/29 这是一条打印到控制台的日志(Print)

	logger1.Println("这是一条打印到控制台的日志(Println)")
	//输出：
	//p303	2019/10/29 这是一条打印到控制台的日志(Println)

	str2 :="一条打印到控制台的日志(Println)"
	logger1.Printf("这是%v",str2)
	//输出：
	//p303	2019/10/29 这是一条打印到控制台的日志(Println)

	// SetPrefix设置记录器的输出前缀。
	logger1.SetPrefix("p202\t")
	//Prefix返回记录器的输出前缀。
	fmt.Println(logger1.Prefix())
	logger1.Println("这是一条日志（Println）")
	//输出：
	//p202	（这里还有一个tab空格）
	//p202	2019/10/29 这是一条日志（Println）


	// Panic等效于l.Print（），然后调用panic（）。
	logger1.Panic("这是一条日志，输出日志后会抛出异常panic")
	//输出：
	//p303	2019/10/29 这是一条日志，输出日志后会抛出异常panic
	//panic: 这是一条日志，输出日志后会抛出异常panic
	//
	//goroutine 1 [running]:
	//log.(*Logger).Panic(0xc000096050, 0xc00008bf40, 0x1, 0x1)
	//	C:/Go/src/log/log.go:212 +0xb1
	//main.main()
	//	C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:107 +0x12b
	从上面可以知道，panic的异常信息也会跟日志完全相同，下同


	str3:="输出日志后会抛出异常panic"
	// Panicf等效于l.Printf（），然后调用panic（）。
	logger1.Panicf("这是一条日志，%v",str3)
	//输出：
	//p303	2019/10/29 这是一条日志，输出日志后会抛出异常panic
	//panic: 这是一条日志，输出日志后会抛出异常panic
	//
	//goroutine 1 [running]:
	//log.(*Logger).Panicf(0xc000082050, 0x4ded72, 0x17, 0xc00007bf40, 0x1, 0x1)
	//C:/Go/src/log/log.go:219 +0xc8
	//main.main()
	//C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:119 +0x165


	// Panicln等效于l.Println（），然后调用panic（）。
	logger1.Panicln("这是一条日志，输出日志后会抛出异常panic")
	//输出：
	//p303	2019/10/29 这是一条日志，输出日志后会抛出异常panic
	//panic: 这是一条日志，输出日志后会抛出异常panic
	//
	//
	//goroutine 1 [running]:
	//log.(*Logger).Panicln(0xc000096050, 0xc00008bf40, 0x1, 0x1)
	//C:/Go/src/log/log.go:226 +0xb1
	//main.main()
	//C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:135 +0x12b
	//可以看到上面多了一个空行


	//这个Output方法是上面所有的fatal系列和panic系列的底层实现！
	// Output写入日志事件的输出。 字符串s包含要在Logger的标志指定的前缀之后打印的文本。
	// 如果s的最后一个字符还不是换行符，则会附加一个换行符。 Calldepth用于恢复PC并用于一般性，
	// 尽管目前在所有预定义路径上为2（ Calldepth一般设置为2即可）。Calldepth主要是被runtime.Caller(calldepth)调用，
	// 下面是runtime.Caller(calldepth)函数的说明：
	
	//runtime.Caller报告有关调用goroutine堆栈上的函数调用的文件和行号信息。 参数skip是要提升的堆栈帧数，
	// 其中0标识Caller的调用者。 （由于历史原因，在调用方和调用方之间，跳过的含义有所不同。）返回
	// 值报告相应调用文件中的程序计数器，文件名和行号。 如果无法恢复信息，则ok布尔值为False。
	
	//output_err := logger1.Output(2, "这是一条Output方法输出的日志111\n")
	//末尾如果有换行了。那么Output方法将不会给他添加换行符了，假如没有换行符则会自动给他加上
	output_err := logger1.Output(2, "这是一条Output方法输出的日志111")
	check_err_log(output_err)
	//输出：
	//p303	2019/10/29 这是一条Output方法输出的日志111

	file, e := os.Create("main3/log/output.txt")
	check_err_log(e)
	//上面的输出是直接输出到控制台，假如我们设置输出到文件的话，那么可以通过以下方法进行设置
	logger1.SetOutput(file)
	//logger1.Fatalln("这是一条日志哈哈！")
	//通过查看文件main3/log/output.txt，可以得到以下信息：
	//p303	2019/10/29 这是一条日志哈哈！
	
	e = logger1.Output(2, "这是一条Output方法输出到文件的日志111")
	check_err_log(e)
	//通过查看文件main3/log/output.txt，可以得到以下信息：
	//p303	2019/10/29 这是一条Output方法输出到文件的日志111
	//这条日志会覆盖上面打印出来的日志，注意，！
	//panic系列的方法也一样不会打印到控制台而是直接打印到文件中去！这里不再做展示！


	//io.writer接口说明：
	// Writer是包装基本Write方法的接口。
	// Write将p的len（p）个字节写入基础数据流。
	//返回从p（0 <= n <= len（p））写入的字节数，以及遇到的任何导致写入提前停止的错误。
	//如果写入返回n <len（p），则必须返回一个非nil错误。
	//写操作不得修改切片数据，即使是临时的也不行。
	//实现不得保留p。
	
	// Writer返回记录器logger的输出目标。
	writer := logger1.Writer()
	//虽然我们知道这个writer是一个接口io.writer的实现，但是我们在下面可以打印出这个实现的具体类型！
	fmt.Printf("%#v\n",writer)
	//之所以要延迟，是因为我们打印到控制台的顺序是不受程序控制的！为了输出好看同时符号我们的逻辑，我们就延迟一会！
	time.Sleep(2e9)
	n, err := writer.Write([]byte("这是一条调用writer写入的日志！\n"))
	//写完上面的一条日志后，我们还是可以继续写入的！下面就是
	//logger1.Fatalln("你好（Fatalln）")
	check_err_log(err)
	time.Sleep(2e9)
	fmt.Println("写入日志的字节数是：",n)
	//writer不用close()(也没有这个方法)
	//输出：
	//&os.File{file:(*os.file)(0xc00007c500)}
	//这是一条调用writer写入的日志！
	//p303	2019/10/29 你好（Fatalln）
	//写入日志的字节数是： 43




	//这些标志定义在Logger生成的每个日志条目之前添加哪些文本。
	//将位进行或运算以控制要打印的内容。
	//无法控制它们的显示顺序（此处列出的顺序）或它们显示的格式（如注释中所述）。
	//仅当指定Llongfile或Lshortfile时，前缀后才带有冒号。
	//例如，标记Ldate | Ltime（或LstdFlags）产生2009/01/23 01:23:23消息，
	// 同时标记Ldate | Ltime | 微秒| Llongfile产生，2009/01/23 01：23：23.123123 /a/b/c/d.go:23：message
	//const (
	//	Ldate         = 1 << iota     //当地时区中的日期：2009/01/23
	//	Ltime                         //当地时区的时间：01：23：23
	//	Lmicroseconds                 //微秒分辨率：01：23：23.123123。 假设Ltime。
	//	Llongfile                     //完整的文件名和行号：/a/b/c/d.go:23
	//	Lshortfile                    //最终文件名元素和行号：d.go：23。 覆盖Llongfile
	//	LUTC                          //如果设置了Ldate或Ltime，请使用UTC而不是本地时区来合成到日志中去，使用这个必须和Ldate或者Ltime一起使用（一起或运算），否则的话不起作用
	//	LstdFlags     = Ldate | Ltime //标准记录器（logger）的初始值
	//)
	fmt.Println(log.Ldate)
	fmt.Println(log.Ltime)
	fmt.Println(log.Lmicroseconds)
	fmt.Println(log.Llongfile)
	fmt.Println(log.Lshortfile)
	fmt.Println(log.LUTC)
	fmt.Println(log.LstdFlags)
	//输出：
	//1
	//2
	//4
	//8
	//16
	//32
	//3
	//除了之后一个之外，上面的所有都是乘以2，也就是左移1位


	//一旦你弄懂上面的logger对象的所有方法后，下面的所有函数都将会完全懂了，因为他们的功能是完全相同的！
	output_err := log.Output(2, "这是一条日志（Output）")
	check_err_log(output_err)
	//输出：
	//2019/10/29 17:40:34 这是一条日志（Output）

	// SetFlags设置标准记录器的输出标志。
	//log.SetFlags(log.Ldate)
	//log.SetFlags(log.Ltime)
	//log.SetFlags(log.Lmicroseconds)
	//log.SetFlags(log.Ldate|log.Lmicroseconds)//任意2个都可以进行或运算代表取这2个信息合成到log中去！
	//log.SetFlags(log.Llongfile)
	//log.SetFlags(log.Lshortfile)
	//log.SetFlags(log.LUTC)
	//log.SetFlags(log.Ldate|log.Ltime|log.LUTC)
	//log.SetFlags(log.LstdFlags)
	log.SetFlags(log.LstdFlags|log.LUTC)
	log.Fatalln("这是一条日志（Fatalln）")
	//上面依次输出：
	//2019/10/29 这是一条日志（Fatalln）
	//17:44:50 这是一条日志（Fatalln）
	//17:45:09.791484 这是一条日志（Fatalln）
	//2019/10/29 17:45:45.366780 这是一条日志（Fatalln）
	//C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:256: 这是一条日志（Fatalln）
	//compress_zlib.go:257: 这是一条日志（Fatalln）
	//这是一条日志（Fatalln）
	//2019/10/29 09:50:34 这是一条日志（Fatalln）
	//2019/10/29 17:51:36 这是一条日志（Fatalln）
	//2019/10/29 09:51:55 这是一条日志（Fatalln）


	log.Panic("打印完这条日志后引发panic")
	//log.Panicf和log.Panicln同理，不再累叙
	
	//输出：
	//2019/10/29 17:53:43 打印完这条日志后引发panic
	//panic: 打印完这条日志后引发panic
	//
	//goroutine 1 [running]:
	//log.Panic(0xc00007bf30, 0x1, 0x1)
	//C:/Go/src/log/log.go:338 +0xb3
	//main.main()
	//C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:275 +0xc0

	//其他的log包中的函数其实是logger对象的方法的简化版本，功能完全一样，在此不再展示了！

}

func check_err_log(err error)  {
	if err != nil {
		fmt.Println(err)
	}
}



























