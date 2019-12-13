package main

import (
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"time"
)


/*
	软件包fmt使用与C的printf和scanf类似的功能实现格式化的I / O。 格式“动词”派生自C，但更简单。


	列印

	The verbs:

	一般（通用）:
		%v 打印结构时，以默认格式％v值，加号（%+v）添加字段名称
		%#v 值的Go语法表示形式
		$T  值类型的Go语法表示形式
		%%  文字百分号； 不消耗任何价值

	Boolean:
		%t	值只能为true或false
	Integer:
		%b	2进制表示数
		%c	相应的Unicode代码点表示的字符
		%d	10进制表示数
		%o	8进制表示数
		%O	以0o为前缀的8进制表示数
		%q	使用Go语法安全地转义的单引号字符文字。
		%x	16进制表示数（a-f小写字母）
		%X	16进制表示数（A-F大写字母）
		%U	Unicode格式：U + 1234； 与“ U +％04X”相同
	浮点和复数成分:
		%b	指数为2的幂的无小数科学计数法，采用strconv.FormatFloat的形式为'b'格式，例如 -123456p-78
		%e	科学记数法, e.g. -1.234456e+78
		%E	科学记数法, e.g. -1.234456E+78
		%f	小数点但无指数, e.g. 123.456
		%F	％f的同义词
		%g	％e用于大指数，否则％f。 精度将在下面讨论。
		%G	％E用于大指数，否则为％F
		%x	十六进制表示法（具有两个指数的十进制幂）， e.g. -0x1.23abcp+20
		%X	大写十六进制符号， e.g. -0X1.23ABCP+20
	字符串和字节切片（与这些动词等效地对待）：
		%s	字符串或切片的未解释字节
		%q	使用Go语法安全地转义的双引号字符串
		%x	基数16，小写，每字节两个字符
		%X	基数16，大写，每字节两个字符
	Slice:
		%p	以基数16表示的第0个元素的地址，开头为0x
	Pointer:
		%p	以16为基础的表示法，前导0x
		％b，％d，％o，％x和％X动词也可以与指针一起使用（并不是一起使用，需要获取返回值还需要转换等等），将值格式化为完全等于整数的形式。

	通用的%v的默认格式是:
		bool:                    %t
		int, int8 etc.:          %d
		uint, uint8 etc.:        ％d，如果使用%#v打印，则为%#x
		float32, complex64, etc: %g
		string:                  %s
		chan:                    %p
		pointer:                 %p
	对于复合对象，使用以下规则递归地打印元素，其布局如下：
		struct:             {field0 field1 ...}
		array, slice:       [elem0 elem1 ...]
		maps:               map[key1:value1 key2:value2 ...]
		pointer to above:   &{}, &[], &map[]

	宽度由动词前的可选十进制数字指定。
	如果不存在，则宽度是表示值所必需的。
	在（可选）宽度之后由句点（或者叫点）和十进制数字指定精度。 如果不存在句点（或者叫点），则使用默认精度。
	没有以下数字的句点（或者叫点）指定精度为零。
	例子：
		%f     默认宽度，默认精度
		%9f    宽度9，默认精度
		%.2f   默认宽度，精度2
		%9.2f  宽度9，精度2
		%9.f   宽度9，精度0

	宽度和精度以Unicode代码点（即，符文）为单位进行度量。
	（这与C的printf不同，后者总是以字节为单位。）两个标志中的一个或两个都可以用字符'*'替换，
	从而使它们的值从下一个操作数获得（在格式化之前）， 操作数必须是int类型。

	对于大多数值，width是要输出的最小符文数，必要时用空格填充格式化的表单。

	但是，对于字符串，字节片和字节数组，精度会限制要格式化的输入的长度（而不是输出的大小），
	并在必要时将其截断。 通常，它以字符为单位进行度量，但是对于这些类型，使用％x或％X格式进行格式化时，将以字节为单位进行度量。

	对于浮点值，width设置字段的最小宽度，而precision设置小数点后的位数（如果适用），除了％g /％G精度设置最大有效位数（除去零位） 。
	例如，给定12.345，格式％6.3f可打印12.345，而％.3g格式可打印12.3（看后面会解释）。 ％e，％f和％＃g的默认精度为6； 对于％g，它是唯一标识该值所需的最少位数。

	对于复数，宽度和精度分别应用于两个分量，并且将结果括在括号中，因此将％f应用于1.2 + 3.4i生成（1.200000 + 3.400000i）。

	Other flags:
		+	始终为数字值打印符号； 保证％q（％+ q）的纯ASCII输出
		-	在右边而不是左边填充空格（左对齐字段）
		#	备用格式：为二进制（％＃b）添加前导0b，为八进制（％＃o）添加0，对于十六进制（％＃x或％＃X）添加0x或0X； 为％p（％＃p）抑制0x; 对于％q，
			如果strconv.CanBackquote返回true，则输出一个原始（带反引号）的字符串； 始终为％e，％E，％f，％F，％g和％G打印小数点； 不要删除％g和％G的结尾零；
			写例如 U + 0078'x'（如果字符可用于％U（％＃U）打印）。
		' '	（空格）在数字（％d）处留有省略号的空格； 在打印字符串或切片的字节之间放置十六进制空格（％x，％X）
		0	填充前导零而不是空格；对于数字，这会将填充移到符号后

	Flags 被不期望使用它们的动词verb忽略。
	例如，如果没有替代的十进制格式，此时％＃d和％d的行为相同。

	对于每个类似于Printf的函数，还有一个Print函数，该函数不采用任何格式，等效于为每个操作数说％v。 Println的另一种变体在操作数之间插入空格，并添加换行符。

	无论动词如何，如果操作数是接口值，那么将使用内部具体值，而不是接口本身。
	因此：
		var i interface{} = 23
		fmt.Printf("%v\n", i)
	将打印出：23.

	除非使用动词％T和％p进行打印，否则特殊的格式化注意事项适用于实现某些接口的操作数。 按申请顺序：

	1.如果操作数是reflect.Value，则将操作数替换为其所保存的具体值，并使用下一个规则继续打印。

	2.如果一个操作数实现了Formatter接口，它将被调用。 Formatter提供了对格式的精细控制。

	3.如果将％v动词与＃标志（%#v,go的语法表示）一起使用，并且操作数实现GoStringer接口，则将调用该接口。

	如果格式（对于Println等，隐式为％v）对字符串（％s％q％v％x％X）有效，则以下两个规则适用：

	4.如果操作数实现错误接口，则将调用Error方法将对象转换为字符串，然后将根据动词的要求对其进行格式化（如果有）。

	5.如果操作数实现String（）字符串方法，则将调用该方法将对象转换为字符串，然后将根据动词的要求对其进行格式化（如果有）。

	对于切片和结构之类的复合操作数（或者说是序列结构），格式递归地应用于每个操作数的元素，而不是整个操作数。
	因此％q将引用字符串切片中的每个元素，而％6.2f将控制浮点数组中每个元素的格式。

	但是，在打印带有类似字符串的动词（％s％q％x％X）的字节片时，将其与字符串等同地视为一个项。

	在诸如以下情况下避免递归
		type X string
		func (x X) String() string { return Sprintf("<%s>", x) }
	在重复之前转换值：
		func (x X) String() string { return Sprintf("<%s>", string(x)) }
	无限递归还可以由自引用数据结构触发，例如，包含自身作为元素的切片（如果该类型具有String方法）。 但是，这种病理情况很少见，而且该软件包也无法防止它们。

	打印结构时，fmt无法并且因此不会在未导出的字段上调用诸如Error或String之类的格式化方法。

	显式参数索引：

	在Printf，Sprintf和Fprintf中，默认行为是为每个格式化动词格式化在调用中传递的连续参数。
	但是，动词前的符号[n]表示将改为格式化第n个单索引参数。 宽度或精度的'*'之前的相同符号选择保存该值的参数索引。 在处理了带括号的表达式[n]后，除非另外指出，否则后续动词将使用自变量n + 1，n + 2等。

	例如，
		fmt.Sprintf("%[2]d %[1]d\n", 11, 22)
	将产生“ 22 11”，而
		fmt.Sprintf("%[3]*.[2]*[1]f", 12.0, 2, 6)
	相当于
		fmt.Sprintf("%6.2f", 12.0)
	将产生“ 12.00”。 因为显式索引会影响后续动词，所以可以通过为第一个要重设的参数重置索引来多次使用相同的符号来打印相同的值：
		fmt.Sprintf("%d %d %#[1]x %#x", 16, 17)
	将返回 "16 17 0x10 0x11".

	格式错误：

	如果为动词给出了无效的参数，例如为％d提供了一个字符串，则生成的字符串将包含问题的描述，如以下示例所示：

		错误类型或未知动词： %!verb(type=value)
			Printf("%d", "hi"):        %!d(string=hi)
		参数太多: %!(EXTRA type=value)
			Printf("hi", "guys"):      hi%!(EXTRA string=guys)
		参数太少: %!verb(MISSING)
			Printf("hi%d"):            hi%!d(MISSING)
		宽度或精度为非整数: %!(BADWIDTH) or %!(BADPREC)
			Printf("%*s", 4.5, "hi"):  %!(BADWIDTH)hi
			Printf("%.*s", 4.5, "hi"): %!(BADPREC)hi
		参数索引的无效或无效使用： %!(BADINDEX)
			Printf("%*[2]d", 7):       %!d(BADINDEX)
			Printf("%.[2]d", 7):       %!d(BADINDEX)

	所有错误均以字符串“％！”开头。 有时后面跟一个字符（动词），并以括号括起来。

	如果Error或String方法在由打印例程调用时触发了紧急情况，则fmt软件包会重新格式化来自紧急情况的错误消息，
	并用表明它来自fmt软件包的指示来装饰它。 例如，如果String方法调用panic（“ bad”），则生成的格式化消息看起来像
	%!s(PANIC=bad)

	%!s仅显示发生故障时正在使用的打印动词。 但是，如果panic是由nil接收器导致Error或String方法引起的，则输出为未经修饰的字符串"<nil>"。

	Scanning

	一组类似的功能会扫描格式化的文本以产生值。 从os.Stdin读取的Scan，Scanf和Scanln; Fscan，Fscanf和Fscanln从指定的io.Reader读取； 从参数字符串读取Sscan，Sscanf和Sscanln。

	Scan，Fscan，Sscan将输入中的换行符视为空格。当且仅当遇到输入中的结尾EOF时候才会结束扫描

	Scanln，Fscanln和Sscanln在换行符处停止进行扫描，并要求在这些项目之后加上换行符或EOF。

	Scanf，Fscanf和Sscanf根据类似于Printf的格式字符串解析参数。 在下面的文本中，“空格”表示除换行符外的任何Unicode空格字符。

	在格式字符串中，由％字符引入的动词消耗并解析输入； 这些动词将在下面更详细地描述。
	格式中除％，空格或换行符以外的其他字符会完全消耗，字符必须存在的该输入字符。 格式字符串中包含零个或多个
	空格的换行符会在输入中消耗零个或多个空格，后跟一个换行符或输入结尾。 格式字符串中换行符后的空格在输入中
	消耗零个或多个空格。 否则，格式字符串中任何一个或多个空格的运行都会在输入中消耗尽可能多的空格。 除非格式字
	符串中的空格行与换行符相邻，否则该行必须占用输入中至少一个空格或找到输入的结尾。

	空格和换行的处理不同于C的scanf系列：在C中，换行被视为任何其他空格，并且当格式字符串中的空格运行在输入中找不到要使用的空格时，这绝不是错误。

	这些动词的行为类似于Printf的行为。
	例如，％x将扫描一个十六进制整数，％v将扫描默认表示形式的值。
	未实现Printf动词％p和％T以及标志＃和+。
	对于浮点和复数值，所有有效的格式化动词（％b％e％E％E％f％F％g％G％x％X和％v）都是等效的，
	并且接受十进制和十六进制表示法（例如：“ 2.3 e + 7”，“ 0x4.5p-8”）和数字分隔的下划线（例如：“ 3.14159_26535_89793”）。

	动词处理的输入是隐式用空格分隔的：除％c外，每个动词的实现都从丢弃其余输入中的前导空格开始，而％s动词（和％v读入字符串）停止在第一个空格或第二个空格处使用输入 换行符。

	扫描不带格式或带有％v动词的整数时，可以接受熟悉的基本设置前缀0b（二进制），0o和0（八进制）和0x（十六进制），以数字分隔的下划线也是如此。

	宽度在输入文本中解释，但是没有用于精确扫描的语法（没有％5.2f，只有％5f）。
	如果提供了width，它将在修剪前导空格后应用，并指定要满足动词阅读的最大符文数。 例如，
	   Sscanf(" 1234567 ", "%5s%d", &s, &i)
	会将s设置为“ 12345”，而将i设置为67，但是当
	   Sscanf(" 12 34 567 ", "%5s%d", &s, &i)
	将s设置为“ 12”，将i设置为34。

	在所有扫描功能中，回车符后紧跟换行符被视为普通换行符（\r\n表示与\n相同）。

	在所有扫描功能中，如果操作数实现了Scan方法（即，它实现了Scanner接口），则该方法将用于扫描该操作数的文本。 另外，如果扫描的参数数量小于提供的参数数量，则会返回错误。

	所有要扫描的参数都必须是指向基本类型或Scanner接口实现的指针。

	与Scanf和Fscanf一样，Sscanf不需要消耗其整个输入。
	无法恢复Sscanf使用了多少输入字符串。

	注意：Fscan等可以读取返回的输入后的一个字符（符文），这意味着调用扫描例程的循环可能会跳过某些输入。
	仅当输入值之间没有空格时，这通常会出现问题。 如果提供给Fscan的阅读器实现了ReadRune，则该方法将用于读取字符。
	如果阅读器还实现了UnreadRune，则将使用该方法保存字符，并且后续调用不会丢失数据。 要将ReadRune和UnreadRune方法
	附加到没有该功能的阅读器，请使用bufio.NewReader。
*/


//上面的翻译全部是采用了谷歌翻译而来，少数是通过我调整翻译的！下面的所有并不包含上面所有的规则！仔细阅读上面所有的规则会更加明白！




// Errorf函数使我们能够使用格式设置功能来创建描述性错误消息。
func ExampleErrorf() {
	const name, id = "bueller", 17
	// Errorf根据格式说明符进行格式化，然后将字符串作为满足错误的值返回。
	//如果格式说明符包含带有错误操作数的％w动词，则返回的错误将实现Unwrap方法，返回操作数。
	// 包含多个％w动词或向其提供未实现错误接口的操作数是无效的。 另外，％w动词是％v的同义词。

	// 这个方法主要是将格式化的字符串变为一个error对象,假如我们使用 errors.New（）是无法做到自定义格式化的字符串的，除非事先将此字符串格式化了
	// ，但是如果像下面这样写的话可以一步到位，不用先格式化字符串然后再new一个error对象了。
	err := fmt.Errorf("user %q (id %d) not found", name, id)
	//错误内置接口类型是用于表示错误情况的常规接口，其中nil值表示没有错误。
	fmt.Println(err.Error())

	var e1 =err.Error()//这个err是上面的
	e1="你好啊"			//可以变更
	fmt.Println(e1)

	//var e2 error
	err111 := errors.New("abc你好啊abc")
	fmt.Println(err111)
	fmt.Println(err111.Error())//从这里可以看得出无论是调不调用这个方法都一样会默认调用的！
	// Output: user "bueller" (id 17) not found
}



func ExampleFscanf() {
	var (
		i int
		b bool
		s string
	)

	r := strings.NewReader("5 true gophers")
	// Fscanf扫描从r读取的文本，将连续的以空格分隔的值存储到由格式确定的连续的参数中。 它返回成功解析的项目数。
	//输入中的换行符必须与格式中的换行符匹配。
	// Fscanf名称说明：凡是带有F开头的方法都是要么是从reader中读取然后赋值，要么是将格式化的数据写入writer,
	// 凡是不带F开头的方法都是默认的写入到输出控制台os.Stdout中去。至于reader的格式数据则由后面的f指定，带f的方法
	// 表示格式化，否则表示非格式化。
	//下面的方法就是从r中按照格式"%d %t %s"的形式读取以空格分隔的字符串，然后依次赋值给&i, &b, &s这3个对象，
	// 如果是writer的话则刚好相反，writer是按照什么格式写入，reader则是按照什么格式读取出来
	//Fprintf的名称说明如上原理
	n, err := fmt.Fscanf(r, "%d %t %s", &i, &b, &s)
	if err != nil {
		// Stdin，Stdout和Stderr是打开的文件，它们指向标准输入，标准输出和标准错误文件描述符。
		//请注意，Go运行时会为恐慌和崩溃写入标准错误；
		//关闭Stderr可能会导致这些消息转到其他地方，甚至可能到达以后打开的文件中。
		fmt.Fprintf(os.Stderr, "Fscanf: %v\n", err)
	}
	fmt.Println(i, b, s)
	fmt.Println(n)
	// Output:
	// 5 true gophers
	// 3
}

func ExampleFscanln() {
	s := `dmr 1771 1.61803398875
	ken 271828 3.14159`
	r := strings.NewReader(s)
	var a string
	var b int
	var c float64
	for {
		// Fscan扫描从r读取的文本，并将连续的以空格分隔的值存储到连续的参数中。
		// 换行符算作空格。 它返回成功扫描的项目数。 如果该数目少于参数数目，则err将报告原因。

		// Fscanln与Fscan类似，但是Fscanln在换行符处停止扫描并且返回了，并且在最后一个项目之后必须有换行符或EOF。
		// Fscanln名称说明：F的说明同上，在我们过往的学习中都可以发现：ln结尾要么表示每次读取或者写入完成后是否在
		// 末尾数据加上换行符，要么表示是否在换行处进行结束该次读取(会将换行符读取出来)！很明显，这个scan对象表示
		// 的都是是第2种情况（结束该次读取）！
		n, err := fmt.Fscanln(r, &a, &b, &c)
		if err == io.EOF {//判断是否是文件的末尾了
			break
		}
		if err != nil {//判断读取是否出现了异常
			panic(err)
		}
		// Printf根据格式说明符设置格式并写入标准输出。
		//返回写入的字节数以及遇到的任何写入错误。
		// 因为Printf这个名称没有F开头，我们可以推断出这个print是打印到os.stdout标准输出控制台界面的。f则制定了输出到控制台界面时候的格式
		fmt.Printf("%d: %s, %d, %f\n", n, a, b, c)
	}
	// Output:
	// 3: dmr, 1771, 1.618034
	// 3: ken, 271828, 3.141590
}

func testHavedLnOrNot(){
	
	//下面对scan中是否带ln进行讲解
	var a byte
	var b int
	var b1 int
	var c int
	stdin := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("请输入一个byte值：")
		count, err := fmt.Fscan(stdin, &a)
		//stdin.ReadString('\n')

		//count, err := fmt.Fscan(stdin, &a)
		//stdin.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			fmt.Printf("count:%d\n", count)
			fmt.Println("您的输入有误，请重新输入")
			stdin.ReadString('\n')
			continue
		}
		fmt.Println("输出：", a)
		break
	}



	for {
		fmt.Println("请输入一个int值：")
		//没ln说明了填充满所有的变量为止，而不是以换行符为停止符！同时遇到换行符也当做间隔多个输入值的空格！
		count, err := fmt.Fscan(stdin, &b, &b1)
		//下面这这种方式也可以，有ln说明了是以换行符为停止符来停止scan的！而不是以填满变量为止！即时没填满也会停止scan
		//count, err := fmt.Fscanln(stdin, &b)
		//stdin.ReadString('\n')//不应该放在这里

		if err != nil {
			fmt.Println(err)
			fmt.Printf("count:%d\n", count)
			fmt.Println("您的输入有误，请重新输入")
			//如果输错了，对于scan来说，那么\n肯定没被读取扫描出来的,因此我们要消耗掉\n之前的所有因错误而停止扫描的字符串，否则我们在输入时候会被要求多输入一个换行符
			//对于有ln的scan,则不需要下面的这行代码，因为scanln会自动读取换行符，而scan不会！
			stdin.ReadString('\n')
			continue
		}
		//如果输对了，那么Fscan会把\n当做空格，每个空格前面的一个字符串被当做输入值，空格分隔的所有输入的字符串值被赋值给所有变量后会自动结束scan
		//所以，\n还是被消耗掉了。
		//对于Fscanln不会把换行符当做空格，同时会在换行符处停止扫描，即使变量还没被完全填充！！
		//不过既然能来到这里，说明了输入对了，那么我们会在读取完换行符后就 会停止整个扫描，所以这里就不必写
		//stdin.ReadString('\n')了，否则我们在输入时候会被要求多输入一个换行符
		fmt.Println("输出1：", b)
		fmt.Println("输出2：", b1)
		break
	}


	for {
		fmt.Println("请输入一个int值：")
		count, err := fmt.Fscan(stdin, &c)
		//stdin.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			fmt.Printf("count:%d\n", count)
			fmt.Println("您的输入有误，请重新输入")
			stdin.ReadString('\n')
			continue
		}
		fmt.Println("输出：", c)
		break
	}
	//输出：
	//请输入一个byte值：
	//a
	//expected integer
	//count:0
	//您的输入有误，请重新输入
	//请输入一个byte值：
	//v
	//expected integer
	//count:0
	//您的输入有误，请重新输入
	//请输入一个byte值：
	//97
	//输出： 97
	//请输入一个int值：
	//11
	//
	//
	//
	//22
	//输出1： 11
	//输出2： 22
	//请输入一个int值：
	//33
	//输出： 33
	
}


func ExampleSscanf() {
	var name string
	var age int
	// Sscanf扫描参数字符串，将连续的以空格分隔的值存储到由格式确定的连续的参数中。 它返回成功解析的项目数。
	//输入中的换行符必须与格式中的换行符匹配。
	//Sscanf名称说明：f的说明同上，表示时候按照什么格式进行读取，S要么表示要读取的源数据不是从io.reader中读取，而是从字符串中读取的！
	// writer同理！S还表示返回值是否返回string，下面api会展示，但不是在这里展示！
	n, err := fmt.Sscanf("Kim is 22 years old", "%s is %d years old", &name, &age)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d: %s, %d\n", n, name, age)

	// Output:
	// 2: Kim, 22
}

func ExamplePrint() {
	const name, age = "Kim", 22
	//使用其操作数的默认格式打印格式并写入标准输出。
	//如果都不是字符串，则在操作数之间添加空格。
	//返回写入的字节数以及遇到的任何写入错误。
	fmt.Print(name, " is ", age, " years old.\n")

	//通常不要担心Print返回的任何错误。因此这里没处理任何的错误

	// Output:
	// Kim is 22 years old.
}



func ExamplePrintln() {
	const name, age = "Kim", 22
	//通过对比这个api和上一个的api,你会发现这里没输入\n符号。因为这个api默认会帮你加上的！
	//这就属于第一种情况了！（如果这里读不明白请看上面）
	fmt.Println(name, "is", age, "years old.")

	//通常不用担心Println返回的任何错误。

	// Output:
	// Kim is 22 years old.
}

func ExamplePrintf() {
	const name, age = "Kim", 22
	// Printf根据格式说明符设置格式并写入标准输出。
	//返回写入的字节数以及遇到的任何写入错误。
	// f指明了按照什么格式（或者说是指明了在哪个位置）进行插入（而不是随便插入相应的变量的值）
	fmt.Printf("%s is %d years old.\n", name, age)

	//通常不要担心Printf返回的任何错误。

	// Output:
	// Kim is 22 years old.
}



func ExampleSprint() {
	const name, age = "Kim", 22
	//使用默认格式为其操作数设置Sprint格式，并返回结果字符串。
	//如果都不是字符串，则在操作数之间添加空格。
	//Sprint名称说明，S表示的就是第二种情况，即表示需不需要返回默认格式化后的字符串，这跟打印还真没什么关系。
	s := fmt.Sprint(name, " is ", age, " years old.\n")
	// WriteString将字符串s的内容写入w，w接受一个字节切片。
	//如果w实现了StringWriter，则直接调用其WriteString方法。
	//否则，w.Write只会被调用一次。
	io.WriteString(os.Stdout, s) //为简单起见，忽略错误。平常来说，我们这里都需要处理错误的！

	// Output:
	// Kim is 22 years old.
}

func ExampleSprintln() {
	const name, age = "Kim", 22
	// Sprintln格式使用其操作数的默认格式，并返回结果字符串。
	//始终在操作数之间添加空格，并在最后一个操作数后添加换行符。这跟print还是没什么关系！
	s := fmt.Sprintln(name, "is", age, "years old.")

	io.WriteString(os.Stdout, s) // Ignoring error for simplicity.

	// Output:
	// Kim is 22 years old.
}

func ExampleSprintf() {
	const name, age = "Kim", 22
	// Sprintf根据格式说明符设置格式，并返回结果字符串。
	//fmt.Sprint只能按照默认的格式进行拼接字符串，但是我们这里可以自定义格式进行字符串的拼接，只因为后面加了个f。同样这个api也跟打印没什么关系！
	s := fmt.Sprintf("%s is %d years old.\n", name, age)

	io.WriteString(os.Stdout, s) // Ignoring error for simplicity.

	// Output:
	// Kim is 22 years old.
}



func ExampleFprint() {
	const name, age = "Kim", 22
	// Fprint格式使用其操作数的默认格式并写入w。
	//如果都不是字符串，则在操作数之间添加空格。
	//返回写入的字节数以及遇到的任何写入错误。
	//Fprint名称说明：同上，也是允许自定义结果字符串的去向，即赋值给谁！如果没有F的话，我们默认是输出到os.Stdout，
	//下面我们也是指明了输出到os.Stdout。因此有无F此时是不影响结果的！下面的fmt.Print就说明了这点！
	n, err := fmt.Fprint(os.Stdout, name, " is ", age, " years old.\n")

	// Fprint的n和err返回值是基础io.Writer返回的值。
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprint: %v\n", err)
	}
	//使用其操作数的默认格式打印格式并写入标准输出。
	//如果都不是字符串，则在操作数之间添加空格。
	//返回写入的字节数以及遇到的任何写入错误。
	fmt.Print(n, " bytes written.\n")

	// Output:
	// Kim is 22 years old.
	// 21 bytes written.
}

func ExampleFprintln() {
	const name, age = "Kim", 22
	//关于ln命令的api的说明：这些例程以'ln'结尾，不使用格式字符串，始终在操作数之间添加空格，并在最后一个操作数之后添加换行符。

	// Fprintln格式使用其操作数的默认格式并写入w。
	//始终在操作数之间添加空格，并添加换行符。
	//返回写入的字节数以及遇到的任何写入错误。
	//通过对比这个api和上一个api(fmt.Fprint)的区别可以发现我们这里不需要手动码入\n,因为这个api默认会在结尾加上的！
	n, err := fmt.Fprintln(os.Stdout, name, "is", age, "years old.")

	// Fprintln的n和err返回值是基础io.Writer返回的值。
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintln: %v\n", err)
	}
	fmt.Println(n, "bytes written.")

	// Output:
	// Kim is 22 years old.
	// 21 bytes written.
}

func ExampleFprintf() {
	const name, age = "Kim", 22
	//这些例程以'f'结尾，并采用格式字符串。

	// Fprintf根据格式说明符格式化并写入w。
	//返回写入的字节数以及遇到的任何写入错误。
	n, err := fmt.Fprintf(os.Stdout, "%s is %d years old.\n", name, age)

	// Fprintf的n和err返回值是基础io.Writer返回的值。
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintf: %v\n", err)
	}
	fmt.Printf("%d bytes written.\n", n)

	// Output:
	// Kim is 22 years old.
	// 21 bytes written.
}



// Print，Println和Printf以不同的方式布置其参数。 在此示例中，我们可以比较它们的行为。 Println总是在其打印的项目之间添加空格，而Print仅在非字符串参数之间添加空格，而Printf完全按照所告诉的内容进行添加。
// Sprint，Sprintln，Sprintf，Fprint，Fprintln和Fprintf的行为与此处显示的相应Print，Println和Printf函数相同。
func Example_printers() {
	a, b := 3.0, 4.0
	/// *
	//	Hypot-sqrt（p * p + q * q），但仅在结果允许时溢出。
	//* /
	// Hypot返回Sqrt（p * p + q * q），请注意避免不必要的上溢和下溢。
	//特殊情况是：
	// Hypot（±Inf，q）= + Inf
	// Hypot（p，±Inf）= + Inf
	// Hypot（NaN，q）= NaN
	// Hypot（p，NaN）= NaN
	//说白了就是返回2个数的平方和然后再开方的值
	h := math.Hypot(a, b)

	//当两个都不是字符串时，Print在参数之间插入空格。
	//它不会在输出中添加换行符，因此我们显式添加了一个换行符。
	fmt.Print("The vector (", a, b, ") has length ", h, ".\n")
	fmt.Print("The","vector\n")//可以对比上面，如果是字符串确实没在字符串之间添加空格

	// Println总是在其参数之间插入空格，因此在这种情况下，它不能用于产生与Print相同的输出； 它的输出有多余的空格。
	//另外，Println总是向输出添加换行符。
	fmt.Println("The vector (", a, b, ") has length", h, ".")
	fmt.Println("The","vector")//可以对比上面，无论是字符串还是其他任何对象都在参数对象之间添加空格

	// Printf提供了完整的控制，但使用起来更复杂。
	//它不会在输出中添加换行符，因此我们在格式说明符字符串的末尾显式添加一个换行符。
	fmt.Printf("The vector (%g %g) has length %g.\n", a, b, h)
	fmt.Printf("--The","vector")//如果你误用了，他会报错的！

	// Output:
	//	The vector (3 4) has length 5.
	//	Thevector
	//	The vector ( 3 4 ) has length 5 .
	//	The vector
	//	The vector (3 4) has length 5.
	//	--The%!(EXTRA string=vector)
}



//这些示例演示了使用格式字符串进行打印的基础。 Printf，Sprintf和Fprintf都采用格式字符串，该字符串指定如何格式化后续参数。
// 例如，％d（我们称其为“动词”）说要打印相应的参数，该参数必须是整数（或包含整数的东西，例如int切片），以十进制表示。
// 动词％v（“ v”代表“值”）始终以默认格式设置参数格式，就像Print或Println会如何显示它一样。 特殊动词％T（“ T”表示“类型”）输出自变量的类型，
// 而不是其值。 这些例子并不详尽。 有关所有详细信息，请参见软件包注释。
func Example_formats() {
	//基本示例集，显示％v是默认格式，在本例中为整数的十进制，可以使用％d明确请求； 输出就是Println生成的。
	integer := 23
	//这些打印都显示为“ 23”（不带引号）。
	fmt.Println(integer)
	fmt.Printf("%v\n", integer)//%v等同于fmt.Print，都是默认的格式
	fmt.Printf("%v\n", "sdsdds")
	fmt.Printf("%d\n", integer)//%d等同于fmt.Print，都是默认的格式,但是跟%v的区别在于必须跟数字，如下会报错：
	fmt.Printf("%d\n", "sdsdds")//报错，%!d(string=sdsdds)

	//特殊动词％T显示项目的类型，而不是其值。%T其实是求取反射类型
	fmt.Printf("%T %T\n", integer, &integer)
	// Result: int *int

	// Println（x）与Printf（“％v \ n”，x）相同，因此在以下示例中我们将仅使用Printf。
	// 每一篇都演示了如何格式化特定类型的值，例如整数或字符串。 我们以％v开头每个格式字符串以显示默认输出，
	// 然后以一种或多种自定义格式输出。

	//使用％v或％t将布尔值打印为“ true”或“ false”。
	truth := true
	fmt.Printf("%v %t\n", truth, truth)
	// Result: true true

	//整数以％v和％d的形式显示为小数，或以％x的形式以十六进制显示，以％o的形式以八进制显示，或以％b的形式输出为二进制。
	answer := 42
	fmt.Printf("%v %d %x %o %b\n", answer, answer, answer, answer, answer)
	// Result: 42 42 2a 52 101010

	//浮点数具有多种格式：％v和％g打印一个紧凑的表示形式，而％f打印一个小数点，而％e使用指数表示法。
	// 此处使用的格式％6.2f显示了如何设置宽度和精度以控制浮点值的外观。 在这种情况下，6是该值的打
	// 印文本的总宽度（请注意输出中的多余空格），2是要显示的小数位数。
	pi := math.Pi
	fmt.Printf("%v %g %.2f (%6.2f) %e\n", pi, pi, pi, pi, pi)
	// Result: 3.141592653589793 3.141592653589793 3.14 (  3.14) 3.141593e+00
	//从这里可以知道默认是保留15位的小数，%6.2f表示必须2位小数而且总长度为6，不够长度的在整数前面加上空格。如(  3.14)，
	//结果3.141593e+00说明了如果是以指数来表示的话，则最大的非指数部分的小数只能是len(141593)即6位

	//复数格式为带括号的浮点对，在虚部后带有“ i”。打印出来的字符串默认会放到括号中括着！
	point := 110.7 + 22.5i
	fmt.Printf("%v %g %.2f %.2e\n", point, point, point, point)
	// Result: (110.7+22.5i) (110.7+22.5i) (110.70+22.50i) (1.11e+02+2.25e+01i)
	//%.2f表示实部和虚部都保留2位小数！%.2e同样也同时作用于虚部和实部！


	//rune符文是整数基类，但是当用％c打印时，将显示具有该Unicode值的字符，即解码出字符形式。
	// ％q动词将它们显示为带引号的字符，％U显示为十六进制Unicode代码点，
	// 而％＃U显示为代码点（如果是不可打印符文的话）或者是带引号的可打印形式（如果是可打印符文的话）。
	smile := '😀'
	fmt.Printf("%v %d %c %q %U %#U\n", smile, smile, smile, smile, smile, smile)
	// Result: 128512 128512 😀 '😀' U+1F600 U+1F600 '😀'
	// 16进制的1F600的十进制就是128512。

	//字符串按％v和％s的原样进行格式化，％q为带引号的字符串，％＃q为反引号的字符串。下面就是反引号，''这个是引号（或者说是正引号）
	placeholders := `foo "bar"`
	fmt.Printf("%v %s %q %#q\n", placeholders, placeholders, placeholders, placeholders)
	// Result: foo "bar" foo "bar" "foo \"bar\"" `foo "bar"`

	//以％v格式化的map会以其默认格式显示键和值。
	//％＃v形式（在此上下文中，＃被称为“标志(flag)”）以Go源格式显示map。 map以一致的顺序打印，并按键值排序。
	isLegume := map[string]bool{
		"peanut":    true,
		"dachshund": false,
	}
	fmt.Printf("%v %#v\n", isLegume, isLegume)
	// Result: map[dachshund:false peanut:true] map[string]bool{"dachshund":false, "peanut":true}

	//以％v格式化的struct以其默认格式显示字段值。
	//％+v表示显示struct字段名称和struct字段名称对应的值，而％＃v以Go源格式格式化struct。
	person := struct {
		Name string
		Age  int
	}{"Kim", 22}
	fmt.Printf("%v %+v %#v\n", person, person, person)
	// Result: {Kim 22} {Name:Kim Age:22} struct { Name string; Age int }{Name:"Kim", Age:22}

	//指针的默认格式显示基础值，并在其前加上“＆”号。 ％p动词以十六进制打印指针值。
	// 我们在这里为％p的参数使用类型为nil的参数，因为任何非nil指针的值都会在运行之间变化。
	// 运行注释掉的Printf自己看看。
	pointer := &person
	//(*int)(nil)，nil本身是不带任何类型的参数而已，无论如何获取都是一个相同的nil对象，我们这里给他加上一个类型的话，
	//就不会每次获取nil都会产生共同的指针值了，这里其实是新建了一个*int对象，他的值为nil,所以我们获取的就不是nil参数的
	//内存地址了，我们获取的是*int对象的内存地址值，或者说是指针值！注意我们不能去掉*int，因为在go中当一个参数拥有类型才
	//可以被获取地址值！一个参数值是不可以被获取地址值的！
	fmt.Printf("%v %p %T\n", pointer, (*int)(nil),(*int)(nil))
	//报错，sdsdsd %!p(string=sdsdsd)，"sdsdsd"是一个参数值而不是一个类型实例的对象！他不带有类型（因为不存在类型变量索引它）！虽然看上去是一个字符串！
	//但是这个字符串确实不可被索引的！不可被索引的参数是不可以输出指针值的！在这里我们表示：不被变量名接收的值都是不可被索引的值！
	fmt.Printf("%v %p\n", "sdsdsd", "sdsdsd")
	//输出:
	//&{Kim 22} 0x0 *int，这个0x0也是一个地址值！
	//sdsdsd %!p(string=sdsdsd)

	//fmt.Printf("%v %p\n", pointer, pointer)
	// 输出: &{Kim 22} 0xc000004500 //看上面的注释！

	//通过将格式应用于每个元素来格式化数组和切片。
	greats := [5]string{"Kitano", "Kobayashi", "Kurosawa", "Miyazaki", "Ozu"}
	fmt.Printf("%v %q\n", greats, greats)
	// Result: [Kitano Kobayashi Kurosawa Miyazaki Ozu] ["Kitano" "Kobayashi" "Kurosawa" "Miyazaki" "Ozu"]

	kGreats := greats[:3]
	fmt.Printf("%v %q %#v\n", kGreats, kGreats, kGreats)
	// 输出: [Kitano Kobayashi Kurosawa] ["Kitano" "Kobayashi" "Kurosawa"] []string{"Kitano", "Kobayashi", "Kurosawa"}

	//字节片是特殊的。 像％d这样的整数动词以这种格式打印元素。
	// ％s和％q形式将切片视为字符串。 ％x动词具有带空格标志的特殊形式，该空格标志在字节之间放置一个空格。
	cmd := []byte("a⌘")
	fmt.Printf("%v %d %s %q %x % x %  x\n", cmd, cmd, cmd, cmd, cmd, cmd, cmd)
	// Result: [97 226 140 152] [97 226 140 152] a⌘ "a⌘" 61e28c98 61 e2 8c 98
	//% x，中间有个空格，表示输出的16进制的之间插入一个空格来做分隔符！但是只能是空格作为间隔符而不能是其他字符了，多个空白字符都是
	//当作一个空格字符来对待！

	//time.Unix说明：
	//自1970年1月1日UTC以来，Unix返回与给定的Unix时间，秒秒和nsec纳秒相对应的本地时间。
	//将nsec传递到[0，999999999]范围之外是有效的。
	//并非所有秒值都具有对应的时间值。 这样的值之一是1 << 63-1（最大的int64值）。

	//实现Stringer的类型与字符串的打印方式相同。 由于Stringers返回字符串，因此我们可以使用特定于字符串的动词（例如％q）来打印它们。
	now := time.Unix(123456789, 0).UTC() // time.Time实现fmt.Stringer。
	fmt.Printf("%v %q\n", now, now)
	// Result: 1973-11-29 21:33:09 +0000 UTC "1973-11-29 21:33:09 +0000 UTC"

	// 全部输出如下:
	//	23
	//	23
	//	sdsdds
	//	23
	//	%!d(string=sdsdds)
	//	int *int
	//	true true
	//	42 42 2a 52 101010
	//	3.141592653589793 3.141592653589793 3.14 (  3.14) 3.141593e+00
	//	(110.7+22.5i) (110.7+22.5i) (110.70+22.50i) (1.11e+02+2.25e+01i)
	//	128512 128512 😀 '😀' U+1F600 U+1F600 '😀'
	//	foo "bar" foo "bar" "foo \"bar\"" `foo "bar"`
	//	map[dachshund:false peanut:true] map[string]bool{"dachshund":false, "peanut":true}
	//	{Kim 22} {Name:Kim Age:22} struct { Name string; Age int }{Name:"Kim", Age:22}
	//	&{Kim 22} 0x0 *int
	//	sdsdsd %!p(string=sdsdsd)
	//	[Kitano Kobayashi Kurosawa Miyazaki Ozu] ["Kitano" "Kobayashi" "Kurosawa" "Miyazaki" "Ozu"]
	//	[Kitano Kobayashi Kurosawa] ["Kitano" "Kobayashi" "Kurosawa"] []string{"Kitano", "Kobayashi", "Kurosawa"}
	//	[97 226 140 152] [97 226 140 152] a⌘ "a⌘" 61e28c98 61 e2 8c 98
	//	61 e2 8c 98
	//	1973-11-29 21:33:09 +0000 UTC "1973-11-29 21:33:09 +0000 UTC"
}

func testVerb() {
	arr:=[]byte{'a','b'}

	fmt.Printf("%#p\n",&arr)//自动获取地址值并输出16进制
	fmt.Printf("%#v\n",&arr)//默认格式（十进制）
	fmt.Printf("%#b\n",&arr)//二进制
	fmt.Printf("%#x\n",&arr)//16进制
	fmt.Printf("%#d\n",&arr)//十进制
	fmt.Printf("%#v\n",unsafe.Pointer(&arr))//只能先获取地址值然后再用%v打印出来，此时%v底层调用%p获取
	fmt.Printf("%#d\n",unsafe.Pointer(&arr))//只能先获取地址值然后再用%d打印出来上面16进制转换后的10进制地址值。


	fmt.Println("---------")
	c:=make(chan int,2)
	fmt.Printf("%#v\n",&c)
	fmt.Printf("%#p\n",&c)
	//从上面可知，填充%v的值类型为unsafe.Pointer或者chan（chan类族，不是单一类型）时候才会等同(注意我说的是等同，不是等效)于%p打
	// 印出16进制的地址值，其他类型都不会打印出地址值

	//#号对于%d来说无效，对于%p来说省略0x开头，对于进制%b...等动词来说打印出对应的0b等,对于其他类型则直接打印出类型名+后面的
	// 动词打印出来的东西
	
	//输出：
	//	c0000044a0
	//	&[]uint8{0x61, 0x62}
	//	&[]uint8{0x61, 0x62}
	//	&[0b1100001 0b1100010]
	//	&0x6162
	//	&[97 98]
	//	(unsafe.Pointer)(0xc0000044a0)
	//	824633738400
	//	---------
	//	(*chan int)(0xc000006030)
	//	c000006030
}
func main()  {
	//ExampleErrorf()
	//ExampleFscanf()
	//ExampleFscanln()
	//testHavedLnOrNot()
	//ExampleSscanf()
	//ExamplePrint()
	//ExamplePrintln()
	//ExamplePrintf()
	//ExampleSprint()
	//ExampleSprintln()
	//ExampleSprintf()
	//ExampleFprint()
	//ExampleFprintln()
	//ExampleFprintf()
	//Example_printers()
	//Example_formats()
	testVerb()
}