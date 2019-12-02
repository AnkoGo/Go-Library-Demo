package main

import (
	"flag"
	"fmt"
	"net/url"
	"time"
)

func init() {
	fmt.Println("test31.....")
} //如果这里我们定义此文件的init()函数的话，我们将使用main包下的其他文件中定义的init()函数进行执行，如果我们定义了的话，
//那么我们执行这个文件时候将不会采用同包下的其他文件中的init()函数（假设存在）进行执行，因此我们要特别注意了

type v1 struct {
	//String() string//不可以在这里指定
	//Set(string) error//不可以在这里指定
	str string
}

//flag.Value是存储在标志中的动态值的接口。
//（默认值表示为字符串。）
//如果Value的IsBoolFlag（）bool方法返回true，则命令行解析器将-name等效于-name = true，而不使用下一个命令行参数。
//对于每个存在的标志，按命令行顺序调用一次Set。
//标志包可以使用零值接收器（例如nil指针）调用String方法。
//下面2个方法都实现了flag.Value接口
func (v1 *v1) String() string {
	return v1.str
}

func (v1 *v1) Set(s string) error {
	v1.str = s
	return nil
}

//---------------------------------------------
type URLValue struct {
	URL *url.URL
}
//下面的2个方法都实现了flag.Value这个接口
func (v URLValue) String() string {
	if v.URL != nil {
		return v.URL.String()
	}
	return ""
}

func (v URLValue) Set(s string) error {
	//url.Parse将rawurl解析为URL结构。
	// rawurl可以是相对的（无主机的路径）或绝对的（以方案开头）。 尝试在不使用方案的情况下解析主机名和路径是无效的，
	// 但由于解析不明确，可能不一定会返回错误。
	if u, err := url.Parse(s); err != nil {
		return err
	} else {
		*v.URL = *u
	}
	return nil
}

var u = &url.URL{}//这里定义一个空实例对象

func ExampleValue() {
	// NewFlagSet返回带有指定名称name和错误处理属性的新的标志flag空集合FlagSet对象。 如果名称不为空，则会在默认用法消息和错误消息中打印该名称。
	//关于集合FlagSet对象的信息说明如下：

	// FlagSet代表一组已定义的标志。 FlagSet的零值没有名称，并且具有ContinueOnError错误处理。
	//type FlagSet struct {
	//	//Usage函数是在解析标志时发生错误时调用的函数。
	//	//该字段是一个函数（不是方法），可以更改为指向自定义错误处理程序。 调用“用法Usage”后会发生什么情况取决于ErrorHandling设置。
	//	// 对于命令行，默认情况下为ExitOnError，它将在调用Usage之后退出程序。
	//	Usage func()
	//
	//	name          string	//该FlagSet对应的名字是什么
	//	parsed        bool	//解析与否
	//	actual        map[string]*Flag	//实际
	//	formal        map[string]*Flag	//真正存储flag的地方，下面的fs.Var就是会设置到这里去的，这里的key是和flag的name相同
	//	args          []string 		// 要解析的参数都有哪些，里面的元素按照键，值的顺序排着
	//	errorHandling ErrorHandling
	//	output        io.Writer // nil表示stderr; 使用out（）输出
	//}

	//这里的name定义的是FlagSet的名字，而不是flag的name
	fs := flag.NewFlagSet("ExampleValue", flag.ExitOnError)
	// Var使用指定的名称和用法字符串定义一个flag。 flag的类型和值由类型为Value的第一个参数表示，该参数通常保存用户定义的Value接口的实现类。
	// 例如，调用者可以创建一个标志，该标志可以通过给切片使用Value方法来将逗号分隔的字符串转换为切片。 特别是，Set会将逗号分隔的字符串分解为切片。
	// 他的作用和函数flag.Var()一样，其实这个方法是指定了解析到哪个对象上面去的以及使用哪个对象上面的set和string方法来进行解析！
	// 表明了数据的去向！同时指定了需要解析的数据是什么和如何解析！这里并没有指定value值（或者说是默认值，在flag.string()等函
	// 数中一般都要指定的），没有默认值的指定，我们却可以在下面的fs.Parse（）方法中直接设置value!随意下面相当于设置了键和键的帮助信息以及解
	// 析所必须的相关的上下文，但是真正需要我们解析的数据却在fs.Parse（）中指定。
	// 这里的name才是指明了flag的name
	fs.Var(&URLValue{u}, "url", "URL帮助信息")
	//虽然我们设置了多个解析接受者，但是如果我们解析的数据只有一条的话，那么也是无法给这些接受者的键赋值的！
	//但是由于我们要解析到的对象&URLValue{u}是一个对象，而且里面只有一个u的对象，所以我们解析多条数据是不可行的！他会覆盖原来的u,所以除非
	// 我们将u设置成为u的序列或者在set方法中做相对应的设定才会起效果！
	fs.Var(&URLValue{u}, "url1", "URL1帮助信息")
	fs.Var(&URLValue{u}, "url2", "URL2帮助信息")


	////这里是设置已有的键的值，如果不存在这个键的话会抛出异常！
	//set_err := fs.Set("S", "S_value")
	//if set_err != nil {
	//	fmt.Println("###",set_err)
	//}
	////输出：
	////### no such flag -S

	go func() {
		time.Sleep(5e9)
		fmt.Println()
		//fs对象下的方法其实和flag下的函数有着很大的相同之处！
		fmt.Println("------------------fs对象下的方法其实和flag下的函数有着很大的相同之处！下面仅展示一部分用法------------------------")
		// Name返回flag集的名称。
		fmt.Println(fs.Name())//ExampleValue
		// ErrorHandling返回flag集的错误处理行为。
		fmt.Println(fs.ErrorHandling())//1,flag.ExitOnError的值
		// NFlag返回已设置的flag数目。
		fmt.Println(fs.NFlag())//3
		// NArg是在处理flag后剩余的参数个数。
		fmt.Println(fs.NArg())//0，不知道这里为什么一直是0，即使我只解析不足3条数据还是显示0
		// Arg返回第i个参数。 Arg（0）是flag已处理后的第一个剩余参数。 如果请求的元素不存在，则Arg返回一个空字符串。
		fmt.Println(fs.Arg(0))
		fmt.Println(fs.Arg(1))
		fmt.Println(fs.Arg(2))
		fmt.Println(fs.Arg(3))//不知道为什么这里的0,1,2,3全部返回空字符串
		// Args返回非flag参数。
		fmt.Println(fs.Args())//[]
		fmt.Println(fs.Usage)//0x4a7880，Usage是一个函数
		fmt.Println("flag的Name:",fs.Lookup("url").Name)
		fmt.Println("flag的Value:",fs.Lookup("url").Value)
		fmt.Println("flag的DefValue:",fs.Lookup("url").DefValue)

		fmt.Println("flag1的Name:",fs.Lookup("url1").Name)
		fmt.Println("flag1的Value:",fs.Lookup("url1").Value)
		fmt.Println("flag1的DefValue:",fs.Lookup("url1").DefValue)

		fmt.Println(fs.Lookup("S"))
		fmt.Println(fs.Duration("T",time.Hour,"Duration(T)的帮助信息！！！"))


		//这里是设置已有的键的值，如果不存在这个键的话会抛出异常！
		set_err := fs.Set("url", "url_value")
		if set_err != nil {
			fmt.Println("###",set_err)
		}

		fmt.Println("flag的Name:",fs.Lookup("url").Name)
		fmt.Println("flag的Value:",fs.Lookup("url").Value)
		fmt.Println("flag的DefValue:",fs.Lookup("url").DefValue)

		var t time.Duration
		fs.DurationVar(&t,"TV",time.Hour,"Duration(TV)的帮助信息！！！")

		var U uint
		fs.UintVar(&U,"U",12,"Uint64Var(U)的帮助信息！！！")

		var U64 uint64
		fs.Uint64Var(&U64,"U64",12,"Uint64Var(U64)的帮助信息！！！")

		fmt.Println(fs.Int("I",12,"Int(I)的帮助信息！！！"))

		var I int
		fs.IntVar(&I,"IV",12,"Int(IV)的帮助信息！！！")


		fmt.Println(fs.Int64("I64",12,"Int64(I64)的帮助信息！！！"))

		var Ui uint
		fs.UintVar(&Ui,"UI",12,"Int(UI)的帮助信息！！！")

		var Ui64 uint64
		fs.Uint64Var(&Ui64,"UI64",12,"Int(UI64)的帮助信息！！！")

		time.Sleep(1e9)
		fs.Usage()
		fs.PrintDefaults()


		//输出如下：
		//	ExampleValue
		//	1
		//	3
		//	0
		//
		//
		//
		//
		//	[]
		//	0x4a7880
		//	flag的Name: url
		//	flag的Value: https://golang.org/pkg/flag2/ ,这里多次解析数据时候被覆盖了
		//	flag的DefValue:
		//	flag1的Name: url1
		//	flag1的Value: https://golang.org/pkg/flag2/ ,这里多次解析数据时候被覆盖了
		//	flag1的DefValue:	（从这里可以看得出这个的默认只是一个空字符串，下同！可以看出这个flag.NewFlagSet（）创建的flag是与flag.string()函数的形式还是有一些区别的额！）
		//	<nil>
		//	1h0m0s
		//	flag的Name: url
		//	flag的Value: url_value
		//	flag的DefValue:
		//	0xc00000a0a8
		//	0xc00000a0b8
		//	Usage of ExampleValue:
		//		-I int
		//			Int(I)的帮助信息！！！ (default 12)
		//		-I64 int
		//			Int64(I64)的帮助信息！！！ (default 12)
		//		-IV int
		//			Int(IV)的帮助信息！！！ (default 12)
		//		-T duration
		//			Duration(T)的帮助信息！！！ (default 1h0m0s)
		//		-TV duration
		//			Duration(TV)的帮助信息！！！ (default 1h0m0s)
		//		-U uint
		//			Uint64Var(U)的帮助信息！！！ (default 12)
		//		-U64 uint
		//			Uint64Var(U64)的帮助信息！！！ (default 12)
		//		-UI uint
		//			Int(UI)的帮助信息！！！ (default 12)
		//		-UI64 uint
		//			Int(UI64)的帮助信息！！！ (default 12)
		//		-url value
		//			URL帮助信息
		//		-url1 value
		//			URL1帮助信息
		//		-url2 value
		//			URL2帮助信息
		//		-I int
		//			Int(I)的帮助信息！！！ (default 12)
		//		-I64 int
		//			Int64(I64)的帮助信息！！！ (default 12)
		//		-IV int
		//			Int(IV)的帮助信息！！！ (default 12)
		//		-T duration
		//			Duration(T)的帮助信息！！！ (default 1h0m0s)
		//		-TV duration
		//			Duration(TV)的帮助信息！！！ (default 1h0m0s)
		//		-U uint
		//			Uint64Var(U)的帮助信息！！！ (default 12)
		//		-U64 uint
		//			Uint64Var(U64)的帮助信息！！！ (default 12)
		//		-UI uint
		//			Int(UI)的帮助信息！！！ (default 12)
		//		-UI64 uint
		//			Int(UI64)的帮助信息！！！ (default 12)
		//		-url value
		//			URL帮助信息
		//		-url1 value
		//			URL1帮助信息
		//		-url2 value
		//			URL2帮助信息




	}()

	//Parse从参数列表解析flag定义，该定义不包括命令名称。 必须在定义了FlagSet中的所有flag之后并且在程序访问flag之前必须调用它。
	//如果设置了-help或-h，但未定义，则返回值为ErrHelp。
	//关于这个方法看源码会更加清晰，这里做简单介绍：主要是指明了命令行中的参数有哪些，表明了要解析数据的来源，当然这里我们是直接给了一个[]string对象！
	//parse_err := fs.Parse([]string{"-url", "https://golang.org/pkg/flag/"})
	//为了给上面的多个flag的定义解析，我们最好设置与定义的flag的数量相同的要解析数据
	//虽然我们这里解析了多条数据，但是由于我们只是指定了解析数据的一个对象同时我们的set中没做相对应的步骤，所以我们的数据会覆盖原来的u，如果你需要
	//多条数据分别解析到不同的数据的话，我们可以相对应的设置set方法！这里不做拓展了！
	parse_err := fs.Parse([]string{"-url", "https://golang.org/pkg/flag/","-url1",
		"https://golang.org/pkg/flag1/","-url2", "https://golang.org/pkg/flag2/"})

	//这里必须要在url前面加上-，表示完整的命令名字,但是如果出错的话要么返回nil（注意这点，目前觉得这是一个逻辑bug）或者不再返回，直接中断，
	//并且打印出相关的flag的使用信息，他有几种错误的处理，但是有一种错误没处理而是直接退出程序，所以压根没返回值赋给parse_err ，因此更不存在被check_err_flag(parse_err)
	// 捕捉到相关的异常信息的可能了！这也是go中无法捕捉到的异常信息的原理：在函数执行过程中被直接中断退出了，而还没有返回这个函数的值！
	//parse_err := fs.Parse([]string{"url", "https://golang.org/pkg/flag/"})
	check_err_flag(parse_err)

	//parse_err = fs.Parse([]string{"-url111", "https://golang.org/pkg/flag111/"})
	//check_err_flag(parse_err)
	//指明不存在的键名会报错：
	//	flag provided but not defined: -url111
	//	Usage of ExampleValue:
	//		-url value
	//			URL帮助信息

	//多次指定会把前面的解析值进行覆盖，因为上面fs.Parse()的机制，导致了这里u.Scheme等不一定有值，
	// 如果上面出错的话，这里会导致u.Scheme等为空字符串！
	//fs.Parse([]string{"-url", "https://golang.org/pkg/flag111/"})
	fmt.Printf(`{scheme: %q, host: %q, path: %q}`, u.Scheme, u.Host, u.Path)
	time.Sleep(8e9)
	// Output:
	// {scheme: "https", host: "golang.org", path: "/pkg/flag/"}
}




func main() {

	fmt.Println("------------flag.String--------------")
	//字符串定义具有指定名称，默认值和用法字符串的字符串标志。
	//返回值是存储标志值的字符串变量的地址。
	//第一个参数是键，第二个参数是值，第三个参数是帮助信息，下同
	flag_str := flag.String("S", "/home/default_dir", "String帮助信息。。。")
	// Bool使用指定的名称，默认值和用法字符串定义一个bool标志。
	//返回值是存储标记值的bool变量的地址。
	flag_bool := flag.Bool("B", false, "Bool帮助信息。。。")
	// Float64使用指定的名称，默认值和用法字符串定义一个float64标志。
	//返回值是一个float64变量的地址，该变量存储标志的值。
	flag_float64 := flag.Float64("F", 3.14, "Float64帮助信息。。。")
	// Int使用指定的名称，默认值和用法字符串定义一个int标志。
	//返回值是存储该标志值的int变量的地址。
	flag_int := flag.Int("I", 3, "Int帮助信息。。。")
	// Int64使用指定的名称，默认值和用法字符串定义一个int64标志。
	//返回值是存储该标志值的int64变量的地址。
	flag_int64 := flag.Int64("I64", 3, "Int64帮助信息。。。")
	// Uint定义了具有指定名称，默认值和用法字符串的uint标志。
	//返回值是存储标志值的uint变量的地址。
	flag_uint := flag.Uint("Ui", 3, "Uint帮助信息。。。")
	// Uint64使用指定的名称，默认值和用法字符串定义uint64标志。
	//返回值是存储标志值的uint64变量的地址。
	flag_uint64 := flag.Uint64("U64", 3, "Uint64帮助信息。。。")

	//持续时间定义一个time.Duration标志，具有指定的名称，默认值和用法字符串。
	//返回值是time.Duration变量的地址，该变量存储标志的值。
	//标志接受time.ParseDuration可接受的值。

	//duration, e := time.ParseDuration("346899654543323478854")
	//check_err_flag(e)
	//fmt.Println(duration)

	flag_duration := flag.Duration("Dur", time.Second, "Duration帮助信息。。。")
	// Var使用指定的名称和用法字符串定义一个标志。 标志的类型和值由类型为Value的第一个参数表示，
	// 该参数通常保存用户定义的Value实现。 例如，调用者可以创建一个标志，该标志可以通过给分片使用
	// Value方法来将逗号分隔的字符串转换为分片。 特别是，Set会将逗号分隔的字符串分解为切片。
	var v v1
	v.str = "var"
	flag.Var(&v, "V", "Var帮助信息。。。")
	//下面是无法输出这个值的，我们最好采用上面的方式来新建一个变量v,这样的话，在下面我们就可以通过v.str来访问到我们设置好
	// 或者默认的值，
	//flag.Var(&v1{"var"}, "V", "Var帮助信息。。。")

	//下面其实跟上面差不了多少，作用完全相同

	//错误示范：
	//var str1 *string
	//// StringVar定义具有指定名称，默认值和用法字符串的字符串标志。
	////参数p指向一个字符串变量，用于在其中存储标志的值。
	//flag.StringVar(str1,"S1","abc","StringVar111帮助信息。。。")
	//输出：
	//紧急：运行时错误：无效的内存地址或nil指针取消引用
	//[信号0xc0000005代码= 0x0地址= 0x0 pc = 0x4eb7de]
	//错误原因，此时还未初始化啊，指针指向nil,下面进行验证，指针和引用类型的零值是nil,但是值类型的零值不是nil,
	//我们对值类型进行取指针最后得到的也是指针类型，因此我们要特别注意了！
	var str3 *string
	fmt.Println(str3)

	var str1 string
	fmt.Println(&str1)
	//输出：
	//<nil>
	//0xc00004d330

	//正确示范111：
	//var str2="sdsdsd"
	//var str1 *string =&str2
	//flag.StringVar(str1,"S1","abc","StringVar111帮助信息。。。")

	//正确示范222：

	flag.StringVar(&str1, "S1", "abc", "StringVar111帮助信息。。。")

	var bool1 bool
	flag.BoolVar(&bool1, "B1", true, "BoolVar111帮助信息。。。")

	var float1 float64
	flag.Float64Var(&float1, "F1", 3.14, "Float64Var111帮助信息。。。")

	var int1 int
	flag.IntVar(&int1, "I1", 3, "IntVar111帮助信息。。。")

	var i641 int64
	flag.Int64Var(&i641, "I641", 3, "Int64Var111帮助信息。。。")

	var u641 uint64
	flag.Uint64Var(&u641, "U641", 3, "Uint64Var111帮助信息。。。")

	var u1 uint
	flag.UintVar(&u1, "U1", 3, "UintVar111帮助信息。。。")

	var t1 time.Duration
	flag.DurationVar(&t1, "T1", time.Minute, "DurationVar帮助信息。。。")

	// Set设置命名命令行标志的值。
	//set_err := flag.Set("S", "SSS_set")
	//check_err_flag(set_err)
	//虽然我们这里传递的是一个字符串，但是真正的应用到我们的对象F时候是会转为对应的类型的，这个无需我们自己动手做！
	//其他类型类似！底层是通过NewFlagSet()这个函数来进行实现的！下面会讲到这个函数！
	//set_err := flag.Set("F", "3.141")

	//假设命令不传参数的话的输出(省略了其他相同的地方)：
	// ...
	//	flag_str:  SSS_set
	// ...
	//假设传递参数的命令go run test31.go -S /home/abcabc，此时会输出：
	//flag_str:  /home/abcabc

	//// Args返回非标志命令行参数。但是不应该放在这里，而是应该放在flag.Parse()之后才能起作用
	//fmt.Println(flag.Args())//[]
	fmt.Println("---", flag.Arg(1))

	fmt.Println("===", flag.Parsed()) //false

	// NArg是在处理标志后剩余的参数个数。暂时还不知道怎么用
	fmt.Println("~~~", flag.NArg())

	// NFlag返回已设置的命令行标志的数量。
	fmt.Println("^^^", flag.NFlag())

	//Parse()从os.Args [1:]解析命令行标志。 必须在定义所有标志之后并且在程序访问标志之前调用。
	flag.Parse()
	fmt.Println("===", flag.Parsed()) //true

	fmt.Println("---", flag.Arg(1))

	// NArg是在处理标志后剩余的参数个数。
	fmt.Println("~~~", flag.NArg())

	// NFlag返回已设置的命令行标志的数量。
	fmt.Println("^^^", flag.NFlag())

	fmt.Println("flag_str: ", *flag_str) //因为flag_str是指针，所以这里要取值
	fmt.Println("flag_bool: ", *flag_bool)
	fmt.Println("flag_float64: ", *flag_float64)
	fmt.Println("flag_int: ", *flag_int)
	fmt.Println("flag_int64: ", *flag_int64)
	fmt.Println("flag_uint: ", *flag_uint)
	fmt.Println("flag_uint64: ", *flag_uint64)
	fmt.Println("flag_duration: ", *flag_duration)
	//fmt.Println("flag_Var: ", *flag_Var)

	fmt.Println(flag.Args())
	fmt.Println("---", flag.Arg(1))

	fmt.Println("flag_str1: ", str1) //因为flag_str是值，所以这里不用再写*号了
	fmt.Println("flag_bool1: ", bool1)
	fmt.Println("flag_float641: ", float1)
	fmt.Println("flag_int1: ", int1)
	fmt.Println("flag_int641: ", i641)
	fmt.Println("flag_uint1: ", u1)
	fmt.Println("flag_uint641: ", u641)
	fmt.Println("flag_duration1: ", t1)
	fmt.Println("flag_var: ", v.str)

	// Args返回非标志命令行参数。暂时还不知道怎么用
	fmt.Println(flag.Args()) //[]
	// Arg返回第i个命令行参数。 Arg（0）是标志已处理后的第一个剩余参数。 如果请求的元素不存在，则Arg返回一个空字符串。
	//真的不知道什么意思！
	fmt.Println("---", flag.Arg(1))
	//Parsed（）报告是否已解析命令行标志。
	fmt.Println("===", flag.Parsed())

	//查找返回命名命令行标志的*Flag对象，如果不存在则返回nil。
	// Flag表示一个Flag的状态信息。
	//type Flag struct {
	//	Name     string //出现在命令行上的名称，也称为键，当然整个flag列表都是通过字典map的形式来进行存储和查询的！
	//	Usage    string //帮助信息
	//	Value    Value  //名字对应的且需要设定的值，当你设定一个值时候并不会替换默认值，而是存储在这个字段中，多次set都是存储到这个字段中！
	//	DefValue string //默认值（以文本形式），这个值几乎不会随着set进行更改的！
	//}
	fmt.Println("***", flag.Lookup("S"))
	fmt.Println("***", flag.Lookup("V"))
	//假设输入命令时候不传递任何的参数的话，输出：
	//*** &{S String帮助信息。。。 SSS_set /home/default_dir}
	//*** &{V Var帮助信息。。。 var var}
	//假设输入命令时候传递任何的参数：go run test31.go -S /home/abcabc，输出：
	//*** &{S String帮助信息。。。 /home/abcabc /home/default_dir}
	//*** &{V Var帮助信息。。。 var var}

	// NArg是在处理标志后剩余的参数个数。
	fmt.Println("~~~", flag.NArg())

	// NFlag返回已设置的命令行标志的数量。暂时还不知道怎么用
	fmt.Println("^^^", flag.NFlag())

	//除非另行配置，否则PrintDefaults会显示使用情况相关的信息，显示所有已定义的命令行标志的默认设置，除非出现其他错误，否则会以错误的形式抛出相关的信息。
	//对于整数值标志x，默认输出形式为
	// -x int
	// x的用法消息（默认值为7）
	// 用法消息将显示在单独的行上，除了带有一个字节名称的布尔标志外，其他任何消息都不会显示。对于布尔标志，将省略类型，
	// 并且如果标志名称为一个字节，则用法消息将显示在同一行上。如果类型的默认值为零，则省略括号默认值。可以通过在标志
	// 的用法字符串中加上反引号来更改列出的类型（此处为int）。消息中的第一个此类项目被视为要在消息中显示的参数名称，
	// 并且在显示时从消息中删除反引号。例如，给定
	// flag.String（“ I”，“”，“在`directory`中搜索包含文件”）
	// 输出将是
	// 		-I directory
	//			在directory目录中搜索包含文件。
	//要更改标志消息的目标，请调用CommandLine.SetOutput。
	flag.PrintDefaults()
	//上面这条代码输出如下：
	//	-B	Bool帮助信息。。。
	//	-B1
	//	BoolVar111帮助信息。。。 (default true)
	//	-Dur duration
	//	Duration帮助信息。。。 (default 1s)
	//	-F float
	//	Float64帮助信息。。。 (default 3.14)
	//	-F1 float
	//	Float64Var111帮助信息。。。 (default 3.14)
	//	-I int
	//	Int帮助信息。。。 (default 3)
	//	-I1 int
	//	IntVar111帮助信息。。。 (default 3)
	//	-I64 int
	//	Int64帮助信息。。。 (default 3)
	//	-I641 int
	//	Int64Var111帮助信息。。。 (default 3)
	//	-S string
	//	String帮助信息。。。 (default "/home/default_dir")
	//	-S1 string
	//	StringVar111帮助信息。。。 (default "abc")
	//	-T1 duration
	//	DurationVar帮助信息。。。 (default 1m0s)
	//	-U1 uint
	//	UintVar111帮助信息。。。 (default 3)
	//	-U64 uint
	//	Uint64帮助信息。。。 (default 3)
	//	-U641 uint
	//	Uint64Var111帮助信息。。。 (default 3)
	//	-Ui uint
	//	Uint帮助信息。。。 (default 3)
	//	-V value
	//	Var帮助信息。。。 (default var)
	//	-deltaT value
	//	comma-separated list of intervals to use between events
	//	-g string
	//	the variety of gopher (shorthand) (default "pocket")
	//	-gopher_type string
	//	the variety of gopher (default "pocket")
	//	-species string
	//	the species we are studying (default "gopher")
	//  从上面可以知道go默认有4个命令，同时细心发现bool类型的B1和B是没有打印出来类型bool的！
	//  而且，你会发现B1比B多了一个后面括号和信息，这个信息表示的是该键给的默认值，但是如果这个默认值是该类型的零值的话
	//  会省略这个默认值，也就是后面不用写明默认值了，如上B的显示！
	//  在上面我们int64类型的键也被显示成了int类型，这也是应该注意的！

	// Visit按字典顺序访问命令行标志，为每个标志调用fn。 它仅访问已设置的那些标志。
	fmt.Println("flag.Visit111...")
	flag.Visit(func(f *flag.Flag) {
		fmt.Println("flag.Visit222...")
		fmt.Println("flag的Name：", f.Name)
		fmt.Println("flag的Value：", f.Value)
		fmt.Println("flag的DefValue：", f.DefValue)
		fmt.Println("flag的Usage：", f.Usage)
		//输出如下：
		//flag.Visit222...
		//flag的Name： S
		//flag的Value： SSS_set
		//flag的DefValue： /home/default_dir
		//flag的Usage： String帮助信息。。。
		//因为我们在上面有flag.set(xx),所以这里的话我们是可以看到有值的！

		//假设我们去除上面的flag.set(xx)的话，输出如下：
		//空值，因为不执行上面参数中的函数

		//假设我们在命令行中传递参数：go run test31.go -S /home/abcabc，则输出：
		//flag.Visit222...
		//flag的Name： S
		//flag的Value： /home/abcabc
		//flag的DefValue： /home/default_dir
		//flag的Usage： String帮助信息。。。

	})

	time.Sleep(1e9)
	// VisitAll以字典顺序访问命令行标志，为每个标志调用fn。 它访问所有标志，甚至没有设置的标志。
	i:=1
	flag.VisitAll(func(f *flag.Flag) {

		fmt.Printf("flag.Visit 第%v个值:%v\n",i,f)
		// UnquoteUsage从用法字符串中提取一个带反引号的标志，并返回它和未引用的用法。
		//给定“要显示的名称”，它会返回（“名称”，“要显示的名称”）。
		//如果没有反引号，则该名称是对标记值类型的有根据的猜测，如果标记为布尔值，则为空字符串。
		name, usage := flag.UnquoteUsage(f)
		fmt.Println("该flag的类型name：",name)
		fmt.Println("该flag的usage：",usage)
		i++
		//输出如下：
		//	flag.Visit 第1个值:&{B Bool帮助信息。。。 false false}
		//	该flag的name：
		//	该flag的usage： Bool帮助信息。。。
		//	flag.Visit 第2个值:&{B1 BoolVar111帮助信息。。。 true true}
		//	该flag的name：
		//	该flag的usage： BoolVar111帮助信息。。。
		//	flag.Visit 第3个值:&{Dur Duration帮助信息。。。 1s 1s}
		//	该flag的name： duration
		//	该flag的usage： Duration帮助信息。。。
		//	flag.Visit 第4个值:&{F Float64帮助信息。。。 3.14 3.14}
		//	该flag的name： float
		//	该flag的usage： Float64帮助信息。。。
		//	flag.Visit 第5个值:&{F1 Float64Var111帮助信息。。。 3.14 3.14}
		//	该flag的name： float
		//	该flag的usage： Float64Var111帮助信息。。。
		//	flag.Visit 第6个值:&{I Int帮助信息。。。 3 3}
		//	该flag的name： int
		//	该flag的usage： Int帮助信息。。。
		//	flag.Visit 第7个值:&{I1 IntVar111帮助信息。。。 3 3}
		//	该flag的name： int
		//	该flag的usage： IntVar111帮助信息。。。
		//	flag.Visit 第8个值:&{I64 Int64帮助信息。。。 3 3}
		//	该flag的name： int
		//	该flag的usage： Int64帮助信息。。。
		//	flag.Visit 第9个值:&{I641 Int64Var111帮助信息。。。 3 3}
		//	该flag的name： int
		//	该flag的usage： Int64Var111帮助信息。。。
		//	flag.Visit 第10个值:&{S String帮助信息。。。 /home/default_dir /home/default_dir}
		//	该flag的name： string
		//	该flag的usage： String帮助信息。。。
		//	flag.Visit 第11个值:&{S1 StringVar111帮助信息。。。 abc abc}
		//	该flag的name： string
		//	该flag的usage： StringVar111帮助信息。。。
		//	flag.Visit 第12个值:&{T1 DurationVar帮助信息。。。 1m0s 1m0s}
		//	该flag的name： duration
		//	该flag的usage： DurationVar帮助信息。。。
		//	flag.Visit 第13个值:&{U1 UintVar111帮助信息。。。 3 3}
		//	该flag的name： uint
		//	该flag的usage： UintVar111帮助信息。。。
		//	flag.Visit 第14个值:&{U64 Uint64帮助信息。。。 3 3}
		//	该flag的name： uint
		//	该flag的usage： Uint64帮助信息。。。
		//	flag.Visit 第15个值:&{U641 Uint64Var111帮助信息。。。 3 3}
		//	该flag的name： uint
		//	该flag的usage： Uint64Var111帮助信息。。。
		//	flag.Visit 第16个值:&{Ui Uint帮助信息。。。 3 3}
		//	该flag的name： uint
		//	该flag的usage： Uint帮助信息。。。
		//	flag.Visit 第17个值:&{V Var帮助信息。。。 var var}
		//	该flag的name： value
		//	该flag的usage： Var帮助信息。。。
		//	flag.Visit 第18个值:&{deltaT comma-separated list of intervals to use between events [] []}
		//	该flag的name： value
		//	该flag的usage： comma-separated list of intervals to use between events
		//	flag.Visit 第19个值:&{g the variety of gopher (shorthand) pocket pocket}
		//	该flag的name： string
		//	该flag的usage： the variety of gopher (shorthand)
		//	flag.Visit 第20个值:&{gopher_type the variety of gopher pocket pocket}
		//	该flag的name： string
		//	该flag的usage： the variety of gopher
		//	flag.Visit 第21个值:&{species the species we are studying gopher gopher}
		//	该flag的name： string
		//	该flag的usage： the species we are studying
		// 从上面可知，flag.UnquoteUsage输出的信息和flag.PrintDefaults()这个方法输出的信息是一致的！所以不再累叙！

	})

	time.Sleep(1e9)
	//注意：Usage不仅是defaultUsage（CommandLine），因为它（通过godoc标志“用法”）用作如何编写自己的用法函数的示例。
	//Usage将使用情况消息打印出来，记录所有已定义的命令行标志到CommandLine的输出，默认情况下为os.Stderr。
	//解析标志时发生错误时调用。
	//该函数是一个变量，可以更改为指向自定义函数。
	//默认情况下，它会打印一个简单的标题并调用PrintDefaults; 有关输出格式及其控制方式的详细信息，请参见PrintDefaults的文档。
	//自定义使用功能可以选择退出程序； 默认情况下，无论如何退出都会发生，因为命令行的错误处理策略设置为ExitOnError。
	//说白了就是展示当前程序的命令行参数和用法
	fmt.Println(flag.Usage)//0x4ea700
	flag.Usage()
	//输出如下：
	//	0x4ea700
	//	Usage of C:\Users\Administrator\AppData\Local\Temp\___2go_build_io_pro_test.exe:
	//	-B	Bool帮助信息。。。
	//	-B1
	//	BoolVar111帮助信息。。。 (default true)
	//	-Dur duration
	//	Duration帮助信息。。。 (default 1s)
	//	-F float
	//	Float64帮助信息。。。 (default 3.14)
	//	-F1 float
	//	Float64Var111帮助信息。。。 (default 3.14)
	//	-I int
	//	Int帮助信息。。。 (default 3)
	//	-I1 int
	//	IntVar111帮助信息。。。 (default 3)
	//	-I64 int
	//	Int64帮助信息。。。 (default 3)
	//	-I641 int
	//	Int64Var111帮助信息。。。 (default 3)
	//	-S string
	//	String帮助信息。。。 (default "/home/default_dir")
	//	-S1 string
	//	StringVar111帮助信息。。。 (default "abc")
	//	-T1 duration
	//	DurationVar帮助信息。。。 (default 1m0s)
	//	-U1 uint
	//	UintVar111帮助信息。。。 (default 3)
	//	-U64 uint
	//	Uint64帮助信息。。。 (default 3)
	//	-U641 uint
	//	Uint64Var111帮助信息。。。 (default 3)
	//	-Ui uint
	//	Uint帮助信息。。。 (default 3)
	//	-V value
	//	Var帮助信息。。。 (default var)
	//	-deltaT value
	//	comma-separated list of intervals to use between events
	//	-g string
	//	the variety of gopher (shorthand) (default "pocket")
	//	-gopher_type string
	//	the variety of gopher (default "pocket")
	//	-species string
	//	the species we are studying (default "gopher")


	//接下来的这个东西是比较重要的东西
	fmt.Println("------------------------")
	time.Sleep(3e9)
	//这个方法在上面被定义了，请跳到上面去查看！
	ExampleValue()

	////不重要的error对象！
	//fmt.Println(flag.ErrHelp)
	//fmt.Println(flag.CommandLine)
	//fmt.Println(flag.ContinueOnError)
	//fmt.Println(flag.PanicOnError)
	//fmt.Println(flag.ExitOnError)
	//// ErrorHandling定义在解析失败时FlagSet.Parse的行为。
	//fmt.Println(flag.ErrorHandling(1))
	////输出：
	////0
	////2
	////1
	////1



	//在命令行中输入：go run test31.go -b /home/backup
	//注意上面的b键是不存在的
	//输出：
	//	C:\Users\Administrator\Desktop\go_pro\src\io_pro\test>go run test31.go -b /home/backup
	//	test31.....
	//	------------flag.String--------------
	//	<nil>
	//	0xc000032250
	//	flag provided but not defined: -b
	//	Usage of C:\Users\ADMINI~1\AppData\Local\Temp\go-build624475282\b001\exe\test31.exe:
	//	-B    Bool帮助信息。。。
	//	-B1
	//	BoolVar111帮助信息。。。 (default true)
	//	-Dur duration
	//	Duration帮助信息。。。 (default 1s)
	//	-F float
	//	Float64帮助信息。。。 (default 3.14)
	//	-F1 float
	//	Float64Var111帮助信息。。。 (default 3.14)
	//	-I int
	//	Int帮助信息。。。 (default 3)
	//	-I1 int
	//	IntVar111帮助信息。。。 (default 3)
	//	-I64 int
	//	Int64帮助信息。。。 (default 3)
	//	-I641 int
	//	Int64Var111帮助信息。。。 (default 3)
	//	-S string
	//	String帮助信息。。。 (default "/home/default_dir")
	//	-S1 string
	//	StringVar111帮助信息。。。 (default "abc")
	//	-T1 duration
	//	DurationVar帮助信息。。。 (default 1m0s)
	//	-U1 uint
	//	UintVar111帮助信息。。。 (default 3)
	//	-U64 uint
	//	Uint64帮助信息。。。 (default 3)
	//	-U641 uint
	//	Uint64Var111帮助信息。。。 (default 3)
	//	-Ui uint
	//	Uint帮助信息。。。 (default 3)
	//	-V value
	//	Var帮助信息。。。 (default var)
	//上面即使是输入了不存在的命令也一样会执行init()方法

	//接着我们输入：go run test31.go -S /home/abcabc
	//此时输出如下：
	//	C:\Users\Administrator\Desktop\go_pro\src\io_pro\test>go run test31.go -S /home/abcabc
	//	test31.....
	//	------------flag.String--------------
	//	<nil>
	//	0xc00004a240
	//	flag_str:  /home/abcabc
	//	flag_bool:  false
	//	flag_float64:  3.14
	//	flag_int:  3
	//	flag_int64:  3
	//	flag_uint:  3
	//	flag_uint64:  3
	//	flag_duration:  1s
	//	flag_str1:  abc
	//	flag_bool1:  true
	//	flag_float641:  3.14
	//	flag_int1:  3
	//	flag_int641:  3
	//	flag_uint1:  3
	//	flag_uint641:  3
	//	flag_duration1:  1m0s

	//接着我们输入：go run test31.go -V value_abc_bac
	//输出如下：
	//	C:\Users\Administrator\Desktop\go_pro\src\io_pro\test>go run test31.go -V value_abc_bac
	//	test31.....
	//	------------flag.String--------------
	//	<nil>
	//	0xc00004a240
	//	flag_str:  /home/default_dir
	//	flag_bool:  false
	//	flag_float64:  3.14
	//	flag_int:  3
	//	flag_int64:  3
	//	flag_uint:  3
	//	flag_uint64:  3
	//	flag_duration:  1s
	//	flag_str1:  abc
	//	flag_bool1:  true
	//	flag_float641:  3.14
	//	flag_int1:  3
	//	flag_int641:  3
	//	flag_uint1:  3
	//	flag_uint641:  3
	//	flag_duration1:  1m0s
	//  因为我们不知道如何输出这个V，所以我们无法打印出来

	//假如我们想要查询一下命令的使用方法，则应该如下：
	//	C:\Users\Administrator\Desktop\go_pro\src\io_pro\test>go run test31.go -help(或者go run test31.go -h)
	//	test31.....
	//	------------flag.String--------------
	//	<nil>
	//		0xc000032250
	//	<nil>
	//	===
	//	=== false
	//	Usage of C:\Users\ADMINI~1\AppData\Local\Temp\go-build333212630\b001\exe\test31.exe:
	//	-B    Bool帮助信息。。。
	//	-B1
	//	BoolVar111帮助信息。。。 (default true)
	//	-Dur duration
	//	Duration帮助信息。。。 (default 1s)
	//	-F float
	//	Float64帮助信息。。。 (default 3.14)
	//	-F1 float
	//	Float64Var111帮助信息。。。 (default 3.14)
	//	-I int
	//	Int帮助信息。。。 (default 3)
	//	-I1 int
	//	IntVar111帮助信息。。。 (default 3)
	//	-I64 int
	//	Int64帮助信息。。。 (default 3)
	//	-I641 int
	//	Int64Var111帮助信息。。。 (default 3)
	//	-S string
	//	String帮助信息。。。 (default "/home/default_dir")
	//	-S1 string
	//	StringVar111帮助信息。。。 (default "abc")
	//	-T1 duration
	//	DurationVar帮助信息。。。 (default 1m0s)
	//	-U1 uint
	//	UintVar111帮助信息。。。 (default 3)
	//	-U64 uint
	//	Uint64帮助信息。。。 (default 3)
	//	-U641 uint
	//	Uint64Var111帮助信息。。。 (default 3)
	//	-Ui uint
	//	Uint帮助信息。。。 (default 3)
	//	-V value
	//	Var帮助信息。。。 (default var)
	// 从上面可以知道我们确实可以查看到帮助信息或者说是使用信息，但是我们一样会执行了这个文件！

}
func check_err_flag(err error) {
	if err != nil {
		fmt.Println("$$$$$",err)
	}
}
