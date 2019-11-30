package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"
	"text/template/parse"
)

//学习本报之前最好先阅读go中文网的doc

func main() {

	fmt.Println("--------------template.HTMLEscape()-------------------")
	//// HTML escaping.(// HTML转义。)
	//var (
	//	htmlQuot = []byte("&#34;") //双引号转义后的表示// shorter than "&quot;"	(比"&quot;"短)
	//	htmlApos = []byte("&#39;") //单引号转义后的表示// shorter than "&apos;" and apos was not in HTML until HTML5	（比"&quot;"短，直到HTML5才使用HTML）
	//	htmlAmp  = []byte("&amp;")//&字符转义后的表示
	//	htmlLt   = []byte("&lt;")	//<字符转义后的表示
	//	htmlGt   = []byte("&gt;")	//>字符转义后的表示
	//	htmlNull = []byte("\uFFFD")	//非\000字符转义后的表示,向Unicode转化时，如果在某语言中没有该字符，得到的将是Unicode的代码“\uffffd”
	//)

	// HTMLEscape writes to w the escaped HTML equivalent of the plain text data b.
	//HTMLEscape将等同于纯文本数据b的转义HTML写入w。
	//底层是遍历传进来的切片，然后判断是否是上面的字符对应的未转义的字符，如果是，则转义程上面列举出来的相对应的表示，逐个字节写入io
	buffer := new(bytes.Buffer)
	by := []byte(`<head>
<title>我的第&一个     "HTML" 页面</title>
</head>`)
	by = append(by, '\000') //故意在后面加上一个不能由utf8编码转成unicode值的字节

	fmt.Println("转义之前字节：", by)
	fmt.Println("转义之前字符串：", string(by))
	template.HTMLEscape(buffer, by)
	fmt.Println()
	fmt.Printf("转义之后字节：%v\n", buffer.Bytes())
	fmt.Printf("转义之后字符串：%s\n", buffer)

	//输出：
	//转义之前字节： [60 104 101 97 100 62 10 60 116 105 116 108 101 62 230 136 145
	//231 154 132 231 172 172 38 228 184 128 228 184 170 32 32 32 32 32 34 72 84
	//77 76 34 32 233 161 181 233 157 162 60 47 116 105 116 108 101 62 10 60 47 104 101 97 100 62 0]
	//转义之前字符串： <head>
	//<title>我的第&一个     "HTML" 页面</title>
	//</head>
	//
	//转义之后字节：[38 108 116 59 104 101 97 100 38 103 116 59 10 38 108 116 59 116
	//105 116 108 101 38 103 116 59 230 136 145 231 154 132 231 172 172 38 97 109
	//112 59 228 184 128 228 184 170 32 32 32 32 32 38 35 51 52 59 72 84 77 76 38 35
	//51 52 59 32 233 161 181 233 157 162 38 108 116 59 47 116 105 116 108 101 38 103
	//116 59 10 38 108 116 59 47 104 101 97 100 38 103 116 59 239 191 189]
	//转义之后字符串：&lt;head&gt;
	//&lt;title&gt;我的第&amp;一个     &#34;HTML&#34; 页面&lt;/title&gt;
	//&lt;/head&gt;�

	fmt.Println("--------------template.HTMLEscapeString()-------------------")

	// HTMLEscapeString returns the escaped HTML equivalent of the plain text data s.
	// HTMLEscapeString返回等效于纯文本数据s的转义HTML。
	//底层实现：
	//var b bytes.Buffer
	//HTMLEscape(&b, []byte(s))
	src_str := `<head>
<title>我的第&一个     "HTML" 页面</title>
</head>`
	src_byte := append([]byte(src_str), '\000')
	src_str = string(src_byte)
	fmt.Println("转义之前字符串：", src_str)
	dst_str := template.HTMLEscapeString(src_str)

	fmt.Println()
	fmt.Println("转义之后字符串：", dst_str)
	//输出：
	//转义之前字符串： <head>
	//<title>我的第&一个     "HTML" 页面</title>
	//</head>
	//
	//转义之后字符串： &lt;head&gt;
	//&lt;title&gt;我的第&amp;一个     &#34;HTML&#34; 页面&lt;/title&gt;
	//&lt;/head&gt;�

	fmt.Println("--------------template.HTMLEscaper()-------------------")

	// HTMLEscaper returns the escaped HTML equivalent of the textual
	// representation of its arguments.
	// HTMLEscaper返回与其参数的文本表示形式等效的转义HTML。

	//底层实现：
	//HTMLEscapeString(evalArgs(args))
	//evalArgs（）的文档如下：
	// evalArgs将参数列表格式化为字符串。 因此，它等效于
	// fmt.Sprint（args ...）
	//只是每个参数都根据需要使用与模板执行期间默认字符串求值相同的规则进行间接（如果是指针）。
	fmt.Println("转义之前字符串：", src_str)
	dst_str = template.HTMLEscaper(src_str)
	fmt.Println()
	fmt.Println("转义之后字符串：", dst_str)

	fmt.Println()
	fmt.Println("转义之前字符串：", string(by))
	dst_str = template.HTMLEscaper(by)
	fmt.Println()
	fmt.Printf("转义之后字符串：%#v\n", dst_str)

	fmt.Println()
	//如果没有特殊字符（"'\"&<>\000"）的话，他的作用类似下面的Sprint函数：
	//使用默认格式为其操作数设置Sprint格式，并返回结果字符串。
	//如果都不是字符串，则在操作数之间添加空格。
	s := fmt.Sprint(by)
	fmt.Printf("==Sprintf()==:%#v\n", s)

	//输出：
	//转义之前字符串： <head>
	//<title>我的第&一个     "HTML" 页面</title>
	//</head>
	//
	//转义之后字符串： &lt;head&gt;
	//&lt;title&gt;我的第&amp;一个     &#34;HTML&#34; 页面&lt;/title&gt;
	//&lt;/head&gt;�
	//
	//转义之前字符串： <head>
	//<title>我的第&一个     "HTML" 页面</title>
	//</head>
	//
	//转义之后字符串："[60 104 101 97 100 62 10 60 116 105 116 108 101 62 230 136 145 231 154 132 231 172 172 38 228 184 128 228 184 170 32 32 32 32 32 34 72 84 77 76 34 32 233 161 181 233 157 162 60 47 116 105 116 108 101 62 10 60 47 104 101 97 100 62 0]"
	//
	//==Sprintf()==:"[60 104 101 97 100 62 10 60 116 105 116 108 101 62 230 136 145 231 154 132 231 172 172 38 228 184 128 228 184 170 32 32 32 32 32 34 72 84 77 76 34 32 233 161 181 233 157 162 60 47 116 105 116 108 101 62 10 60 47 104 101 97 100 62 0]"
	//转义之后字符串："[60 104 101 97 100 62 10 60 116 105 116 108 101 62 230 136 145 231 154 132 231 172 172 38 228 184 128 228 184 170 32 32 32 32 32 34 72 84 77 76 34 32 233 161 181 233 157 162 60 47 116 105 116 108 101 62 10 60 47 104 101 97 100 62 0]"
	//注意传字符串和传字节时候返回值的形式不同
	//HTMLEscaper(s)如果传入的s不是字符串的话，那么就会转换成为s对应的字符串字面量的形式，并且也是返回这个形式。比如上面传入的是字节形式的话也是这样返回的！

	fmt.Println("--------------template.JSEscape()-------------------")

	//// JavaScript escaping.(// JavaScript转义)
	//
	//var (
	//	jsLowUni = []byte(`\u00`)	//不可识别的unicode值的字符转义为`\u00`
	//	hex      = []byte("0123456789ABCDEF")
	//
	//	jsBackslash = []byte(`\\`)	//'\\'转义为`\\`
	//	jsApos      = []byte(`\'`)	//'\''转义为`\'`
	//	jsQuot      = []byte(`\"`)	//'"'转义为`\"`
	//	jsLt        = []byte(`\x3C`)	//'<'转义为`\x3C`
	//	jsGt        = []byte(`\x3E`)	//'>'转义为`\x3E`
	//)

	// JSEscape writes to w the escaped JavaScript equivalent of the plain text data b.
	// JS Escape使用转义后的JavaScript写入纯文本数据b。
	bf := new(bytes.Buffer)
	src_byte111 := []byte(`<script src="http:\\www.w3school.com.cn\demo/aa//myScript.js"></script>�`)

	fmt.Println("转化之前的js字符串为：", string(src_byte111))
	template.JSEscape(bf, src_byte111)

	fmt.Println("转化之后的js字符串为：", bf)
	//输出：
	//转化之前的js字符串为： <script src="http:\\www.w3school.com.cn\demo/aa//myScript.js"></script>�
	//转化之后的js字符串为： \x3Cscript src=\"http:\\\\www.w3school.com.cn\\demo/aa//myScript.js\"\x3E\x3C/script\x3E�
	//如上，/和//不会转义，但是\和\\会进行转义

	fmt.Println("--------------template.JSEscapeString()-------------------")

	src_str111 := `<script src="http:\\www.w3school.com.cn\demo/aa//myScript.js"></script>�`

	fmt.Println("转化之前的js字符串为：", src_str111)
	JsEscapeString := template.JSEscapeString(src_str111)

	fmt.Println("转化之后的js字符串为：", JsEscapeString)

	//输出：
	//转化之前的js字符串为： <script src="http:\\www.w3school.com.cn\demo/aa//myScript.js"></script>�
	//转化之后的js字符串为： \x3Cscript src=\"http:\\\\www.w3school.com.cn\\demo/aa//myScript.js\"\x3E\x3C/script\x3E�
	//其实是跟上面的api差不多，只不过这里转换的是js字符串的特殊字符，而不是html中的特殊字符，只是转义的规则不一样

	fmt.Println("--------------template.JSEscaper()-------------------")

	fmt.Println()
	fmt.Println("转化之前的js字符串为：", src_str111)
	JSEscaper_str := template.JSEscaper(src_str111)
	fmt.Println("转化之后的js字符串为：", JSEscaper_str)

	fmt.Println()
	fmt.Println("转化之前的js字符串为：", string(src_byte111))
	JSEscaper_str = template.JSEscaper(src_byte111)
	fmt.Printf("转化之后的js字符串为：%#v\n", JSEscaper_str)
	//输出：
	//转化之前的js字符串为： <script src="http:\\www.w3school.com.cn\demo/aa//myScript.js"></script>�
	//转化之后的js字符串为： \x3Cscript src=\"http:\\\\www.w3school.com.cn\\demo/aa//myScript.js\"\x3E\x3C/script\x3E�
	//
	//转化之前的js字符串为： <script src="http:\\www.w3school.com.cn\demo/aa//myScript.js"></script>�
	//转化之后的js字符串为："[60 115 99 114 105 112 116 32 115 114 99 61 34 104 116 116 112 58 92 92 119 119 119 46 119 51 115 99 104 111 111 108 46 99 111 109 46 99 110 92 100 101 109 111 47 97 97 47 47 109 121 83 99 114 105 112 116 46 106 115 34 62 60 47 115 99 114 105 112 116 62 239 191 189]"
	//这个函数跟上面的template.HtmlEscaper()类似，只不过前者是转义html，但是后者是转义js的区别罢了，同样如果没有转义字符的话，也等价于fmt.Sprint()方法

	fmt.Println("--------------template.URLQueryEscaper()-------------------")

	// URLQueryEscaper returns the escaped value of the textual representation of
	// its arguments in a form suitable for embedding in a URL query.
	// URLQueryEscaper以适合嵌入URL查询的形式返回其参数的文本表示形式的转义值。
	//底层实现：
	// url.QueryEscape(evalArgs(args))，目前还没学到net.url包
	// url.QueryEscape对该字符串进行转义，以便可以将其安全地放置在URL查询中。
	//URLQueryEscaper()的参数跟上个api的参数要求相同
	fmt.Println()
	url_str := `http:\\www.w3school.com.cn\demo/aa//myScript.js`
	fmt.Println("转化之后的url查询字符串为：", url_str)

	URLQueryEscaper_str := template.URLQueryEscaper(url_str)
	fmt.Println("转化之后的url查询字符串为：", URLQueryEscaper_str)

	//第二个例子
	fmt.Println()
	url_str = `https://www.w3school.com.cn/js/js_examples.asp`
	fmt.Println("转化之后的url查询字符串为：", url_str)

	URLQueryEscaper_str = template.URLQueryEscaper(url_str)
	fmt.Println("转化之后的url查询字符串为：", URLQueryEscaper_str)
	//输出：
	//转化之后的url查询字符串为： http:\\www.w3school.com.cn\demo/aa//myScript.js
	//转化之后的url查询字符串为： http%3A%5C%5Cwww.w3school.com.cn%5Cdemo%2Faa%2F%2FmyScript.js
	//
	//转化之后的url查询字符串为： https://www.w3school.com.cn/js/js_examples.asp
	//转化之后的url查询字符串为： https%3A%2F%2Fwww.w3school.com.cn%2Fjs%2Fjs_examples.asp
	//等到讲到net.url包时候再继续讲这个api

	fmt.Println("--------------template.IsTrue()-------------------")

	fmt.Println()
	// IsTrue reports whether the value is 'true', in the sense of not the zero of its type,
	// and whether the value has a meaningful truth value. This is the definition of
	// truth used by if and other such actions.
	// IsTrue报告值是否为“ true”（不是其类型的零），以及该值是否具有有意义的true值。 这是if和其他此类操作使用的真相定义。
	// 说白了就是判断是否 不是类型的零值，也就是是否是非类型零值的有效的值，跟isnil和iszero方法作用刚好相反
	fmt.Print("string:")
	fmt.Println(template.IsTrue(""))
	fmt.Print("int:")
	fmt.Println(template.IsTrue(0))
	fmt.Print("int:")
	fmt.Println(template.IsTrue(1))
	fmt.Print("int:")
	fmt.Println(template.IsTrue(-1))

	var ls []byte
	fmt.Print("[]byte:")
	fmt.Println(template.IsTrue(ls))

	var lsptr *[]byte
	fmt.Print("*[]byte:")
	fmt.Println(template.IsTrue(lsptr))

	var arr [2]byte
	fmt.Print("[2]byte:")
	fmt.Println(template.IsTrue(arr))

	var arrptr *[2]byte
	fmt.Print("*[2]byte:")
	fmt.Println(template.IsTrue(arrptr))

	var c chan<- int
	fmt.Print("chan<-int:")
	fmt.Println(template.IsTrue(c))

	var ctr *chan<- int
	fmt.Print("*chan<-int:")
	fmt.Println(template.IsTrue(ctr))
	//输出：
	//string:false true
	//int:false true
	//int:true true
	//int:true true
	//[]byte:false true
	//*[]byte:false true
	//[2]byte:true true
	//*[2]byte:false true
	//chan<-int:false true
	//*chan<-int:false true

	fmt.Println("--------------template.ParseFiles()-------------------")
	fmt.Println()
	// ParseFiles creates a new Template and parses the template definitions from
	// the named files. The returned template's name will have the base name and
	// parsed contents of the first file. There must be at least one file.
	// If an error occurs, parsing stops and the returned *Template is nil.
	//
	// When parsing multiple files with the same name in different directories,
	// the last one mentioned will be the one that results.
	// For instance, ParseFiles("a/foo", "b/foo") stores "b/foo" as the template
	// named "foo", while "a/foo" is unavailable.
	// ParseFiles创建一个新模板，并从命名文件中解析模板定义。 返回的模板名称将具有第一个文件的基本名称和已解析的内容。 必须至少有一个文件。
	//如果发生错误，解析将停止，返回的*Template为nil。
	//在不同目录中解析具有相同名称的多个文件时，最后一个文件名的将是返回值。
	//例如，ParseFiles（“ a / foo”，“ b / foo”）将“ b / foo”存储为名为“ foo”的模板，而“ a / foo”不可用。

	//fmt.Println(template.ParseFiles(""))

	fmt.Println("--------------template.ParseGlob()-------------------")
	fmt.Println()
	// ParseGlob creates a new Template and parses the template definitions from
	// the files identified by the pattern. The files are matched according to the
	// semantics of filepath.Match, and the pattern must match at least one file.
	// The returned template will have the (base) name and (parsed) contents of the
	// first file matched by the pattern. ParseGlob is equivalent to calling
	// ParseFiles with the list of files matched by the pattern.
	//
	// When parsing multiple files with the same name in different directories,
	// the last one mentioned will be the one that results.
	// ParseGlob创建一个新的模板，并从该模式标识的文件中解析模板定义。 这些文件根据filepath.Match的语义进行匹配，并且该模式必须匹配至少一个文件。
	//返回的模板将具有与模式匹配的第一个文件的（基本）名称和（解析的）内容。 ParseGlob等效于使用模式匹配的文件列表调用ParseFiles。
	//在不同目录中解析具有相同名称的多个文件时，最后一个文件名的将是返回值。
	//底层实现：
	//filepath.Glob(pattern)，说明如下：
	//// Glob返回所有匹配模式的文件的名称，如果没有匹配的文件，则返回nil。 模式的语法与Match中的语法相同。 该模式可以描述层次结构名称，例如
	//// / usr / * / bin / ed（假设分隔符为'/'）。
	////
	//// Glob会忽略文件系统错误，例如读取目录的I / O错误。
	////格式错误时，唯一可能返回的错误是ErrBadPattern。
	//其实跟上面的api差不多，只是后者通过正则表达式来进行文件的匹配，而不是给定具体的文件名字，其实这俩都是怼template对象的便捷调用而已

	//fmt.Println(template.ParseGlob(""))

	fmt.Println("--------------template.Must()-------------------")
	fmt.Println()

	// Must is a helper that wraps a call to a function returning (*Template, error)
	// and panics if the error is non-nil. It is intended for use in variable
	// initializations such as
	//	var t = template.Must(template.New("name").Parse("text"))
	// Must是一个帮助程序，它包装对返回(*Template, error)的函数的调用，并且如果error为非nil，则会出现恐慌。 它旨在用于变量初始化，例如
	// var t = template.Must(template.New("name").Parse("text"))
	//fmt.Println(template.Must(template.New("name").Parse("text")))

	fmt.Println("----上面都是这个包中的函数，下面我们来进行讲解这个包中的重要对象Template-------")

	//// common holds the information shared by related templates.
	//// common保存相关模板共享的信息。
	//type common struct {
	//	tmpl   map[string]*Template // Map from name to defined templates.（Map从名称name到定义的模板templates。）
	//	option option
	//	// We use two maps, one for parsing and one for execution.
	//	// This separation makes the API cleaner since it doesn't
	//	// expose reflection to the client.
	//	//我们使用两个映射map，一个用于解析，一个用于执行。
	//	//这种分离使API更加清洁，因为它不会向客户端暴露反射。
	//	muFuncs    sync.RWMutex // protects parseFuncs and execFuncs	(//保护parseFuncs和execFuncs)
	//	parseFuncs FuncMap
	//	execFuncs  map[string]reflect.Value
	//}

	//// FuncMap is the type of the map defining the mapping from names to functions.
	//// Each function must have either a single return value, or two return values of
	//// which the second has type error. In that case, if the second (error)
	//// return value evaluates to non-nil during execution, execution terminates and
	//// Execute returns that error.
	////
	//// When template execution invokes a function with an argument list, that list
	//// must be assignable to the function's parameter types. Functions meant to
	//// apply to arguments of arbitrary type can use parameters of type interface{} or
	//// of type reflect.Value. Similarly, functions meant to return a result of arbitrary
	//// type can return interface{} or reflect.Value.
	//// FuncMap是定义从名称names到函数functions的映射的映射map类型。
	////每个函数必须具有一个返回值，或者具有两个返回值，其中第二个具有类型错误。 在这种情况下，如果第二个（错误）返回值在执行execution过程中评估为非零，则执行终止，并且Execute返回该错误。
	////当模板执行调用带有参数列表的函数时，该列表必须是可分配给函数的参数类型。 适用于任意类型参数的函数可以使用interface{}类型或
	//// reflect.Value类型的参数。 同样，旨在返回任意类型结果的函数可以返回interface{}或reflect.Value。
	//type FuncMap map[string]interface{}
	//
	//下面的每个键的值都是一个内置的函数，比如and是与运算函数，call是调用给定的函数的函数，...
	//var builtins = FuncMap{
	//	"and":      and,	//与运算，and计算其参数的与运算，返回其遇到的第一个错误参数或最后一个参数。
	//	"call":     call,	// call返回将第一个参数作为函数求值的结果。该函数必须返回1个结果或2个结果，其中第二个是错误。
	//	"html":     HTMLEscaper,	// HTMLEscaper返回与其参数的文本表示形式等效的转义HTML。
	//	"index":    index,	// index返回通过以下参数索引其第一个参数的结果。 因此，在Go语法中，“index x 1 2 3”是x[1][2][3]。 每个索引项目必须是映射，切片或数组。
	//	"slice":    slice,	//// slice返回将其第一个参数与其余参数相切片的结果。 因此，按照Go语法，“ slice x 1 2”是x [1：2]，而“ slice x”是x [：]，“ slice x 1”是x [1：]和“ slice x 1 2 3” “是x [1：2：3]。 第一个参数必须是string，slice或array。
	//	"js":       JSEscaper,		// JSEscaper返回与其参数的文本表示形式等效的转义JavaScript。
	//	"len":      length,		// length返回项目的长度，如果没有定义的长度，则返回错误。
	//	"not":      not,	//跟and功能类似，不过这里是非运算
	//	"or":       or,		//跟and功能类似，不过这里是或运算
	//	"print":    fmt.Sprint,		//使用默认格式为其操作数设置Sprint格式，并返回结果字符串。 当都不是字符串时，在操作数之间添加空格。
	//	"printf":   fmt.Sprintf,	//略
	//	"println":  fmt.Sprintln,	//略
	//	"urlquery": URLQueryEscaper,	// URLQueryEscaper以适合嵌入URL查询的形式返回其参数的文本表示形式的转义值。
	//
	//	// Comparisons	（//比较）
	//	"eq": eq, // ==
	//	"ge": ge, // >=
	//	"gt": gt, // >
	//	"le": le, // <=
	//	"lt": lt, // <
	//	"ne": ne, // !=
	//}

	//// Template is the representation of a parsed template. The *parse.Tree
	//// field is exported only for use by html/template and should be treated
	//// as unexported by all other clients.
	////Template是已解析模板的表示。 *parse.Tree字段仅导出供html/template使用，应被所有其他客户端视为未导出。
	//type Template struct {
	//	name string
	//	*parse.Tree			//解析树
	//	*common				//上面有详情定义
	//	leftDelim  string	//左侧分隔符
	//	rightDelim string	//右侧分隔符
	//}

	//// Tree is the representation of a single parsed template.
	//// Tree是单个已解析模板的表示。
	//type Tree struct {
	//	Name      string    // name of the template represented by the tree.	（//树表示的模板的名称。）
	//	ParseName string    // name of the top-level template during parsing, for error messages.	（//解析过程中顶级模板的名称，用于显示错误消息。）
	//	Root      *ListNode // top-level root of the tree.	（最顶层的树根节点）
	//	text      string    // text parsed to create the template (or its parent)	（//要解析的文本，以创建模板（或其父模板））
	//	// Parsing only; cleared after parse.
	//	//仅解析； 解析后清除。或者调用template.copy()复制模板时候也会清楚下面的解析状态信息
	//	funcs     []map[string]interface{}
	//	lex       *lexer
	//	token     [3]item // three-token lookahead for parser.	（//解析器的前一个three-token。）
	//	peekCount int
	//	vars      []string // variables defined at the moment.	（//目前定义的变量。）
	//	treeSet   map[string]*Tree		//树集合
	//}

	fmt.Println("--------------template.New()-------------------")
	//New()也是该包下的一个函数，我们将它和对象Template一起讲
	fmt.Println()
	// New allocates a new, undefined template with the given name.
	// New使用给定名称分配一个新的未定义模板。
	teml := template.New("myTempl")
	// Name returns the name of the template.
	// Name返回模板的名称。
	fmt.Println(teml.Name())

	// Parse parses text as a template body for t.
	// Named template definitions ({{define ...}} or {{block ...}} statements) in text
	// define additional templates associated with t and are removed from the
	// definition of t itself.
	//
	// Templates can be redefined in successive calls to Parse.
	// A template definition with a body containing only white space and comments
	// is considered empty and will not replace an existing template's body.
	// This allows using Parse to add new named template definitions without
	// overwriting the main template body.
	// P将文本解析为t的模板主体。t是teml调用者（Template对象）。
	//文本中的命名模板定义（{{define ...}}或{{block ...}}语句）定义了与t关联的其他模板，并从t本身的定义中删除。
	//可以在连续调用Parse中重新定义模板。
	//具有仅包含空白和注释的主体的模板定义被认为是空的，不会替换现有模板的主体。
	//这允许使用Parse添加新的命名模板定义，而不会覆盖主模板主体。
	//fmt.Println(teml.Parse())

	fmt.Println("--------------template对象-------------------")

	//// state represents the state of an execution. It's not part of the
	//// template so that multiple executions of the same template
	//// can execute in parallel.
	////state表示执行状态。 它不是模板的一部分，因此同一模板的多次执行可以并行执行（就是因为并行执行相同模板时每个执行环境都拥有属于自己的state对象来保存每个执行状态）。
	//type state struct {
	//	tmpl  *Template
	//	wr    io.Writer
	//	node  parse.Node // current node, for errors	（//当前节点，用于错误）
	//	vars  []variable // push-down stack of variable values.	（//保存新建变量的堆栈。）
	//	depth int        // the height of the stack of executing templates.	（//被执行的模板的堆栈的高度。）
	//}

	//// maxExecDepth specifies the maximum stack depth of templates within
	//// templates. This limit is only practically reached by accidentally
	//// recursive template invocations. This limit allows us to return
	//// an error instead of triggering a stack overflow.
	//// maxExecDepth指定模板中模板的最大堆栈深度。 实际上，只有通过意外递归模板调用才能达到此限制。 此限制使我们可以返回错误，而不是触发堆栈溢出。
	//var maxExecDepth = initMaxExecDepth()

	//// variable holds the dynamic value of a variable such as $, $x etc.
	////变量保存变量的动态值，例如$，$ x等。
	////在模板中声明的变量都将保存在这个对象里面
	//type variable struct {
	//	name  string
	//	value reflect.Value
	//}

	//// A Node is an element in the parse tree. The interface is trivial.
	//// The interface contains an unexported method so that only
	//// types local to this package can satisfy it.
	////节点Node是解析树中的元素。 该接口是微不足道的。
	////接口包含一个未导出的方法，因此只有此程序包本地的类型才能满足它。
	//type Node interface {
	//	Type() NodeType
	//	String() string
	//	// Copy does a deep copy of the Node and all its components.
	//	// To avoid type assertions, some XxxNodes also have specialized
	//	// CopyXxx methods that return *XxxNode.
	//	//复制Copy对节点及其所有组件进行深层复制。
	//	//为了避免类型断言，某些XxxNode还具有专门的CopyXxx方法，这些方法返回*XxxNode。
	//	Copy() Node
	//	Position() Pos // byte position of start of node in full original input string	（//完整原始输入字符串中节点开始的字节位置）
	//	// tree returns the containing *Tree.
	//	// It is unexported so all implementations of Node are in this package.
	//	// tree返回包含的* Tree。
	//	//它未导出，因此Node的所有实现都在此包中。这个对象在上面有介绍到
	//	tree() *Tree
	//}

	// Define a template.
	//定义一个模板对象letter

	//模板1
	const letter = `
Dear {{.Name}},
{{if .Attended}}
It was a pleasure to see you at the wedding.{{else}}
It is a shame you couldn't make it to the wedding.{{end}}
{{with .Gift}}Thank you for the lovely {{.}}.
{{end}}
Best wishes,
Josie
`

	//模板2
	//	const letter = `
	//Dear {{.Name}},
	//{{if .Attended}}It was a pleasure to see you at the wedding.{{else}}It is a shame you couldn't make it to the wedding.{{end}}
	//{{with .Gift}}Thank you for the lovely {{.}}.
	//{{end}}
	//Best wishes,
	//Josie
	//`

	////模板3
	//const letter = `Dear {{.Name}},{{if .Attended}}It was a pleasure to see you at the wedding.{{else}}It is a shame you couldn't make it to the wedding.{{end}}{{with .Gift}}Thank you for the lovely {{.}}.{{end}}Best wishes,Josie`

	// Prepare some data to insert into the template.
	//准备一些要插入模板的数据。Recipient的意思是接受者
	type Recipient struct {
		Name, Gift string
		Attended   bool //参与邀请与否
	}
	var recipients = []Recipient{
		{"名字1", "书籍", true},
		{"名字2", "石头", false},
		{"名字3", "", false},
	}

	// Must is a helper that wraps a call to a function returning (*Template, error)
	// and panics if the error is non-nil. It is intended for use in variable
	// initializations such as
	//	var t = template.Must(template.New("name").Parse("text"))

	//Must是一个帮助程序，它包装对参数函数返回值（* Template，err）的调用，
	//如果错误为非nil，则会出现恐慌。 它旨在用于变量初始化，例如：
	//	var t = template.Must(template.New("name").Parse("text"))

	// Parse parses text as a template body for t.
	// Named template definitions ({{define ...}} or {{block ...}} statements) in text
	// define additional templates associated with t and are removed from the
	// definition of t itself.
	//
	// Templates can be redefined in successive calls to Parse.
	// A template definition with a body containing only white space and comments
	// is considered empty and will not replace an existing template's body.
	// This allows using Parse to add new named template definitions without
	// overwriting the main template body.
	//Parse将文本解析为t的模板主体。
	//文本中的命名模板定义（{{define ...}}或{{block ...}}语句）定义了与t关联的其他模板，并从t本身的定义中删除。
	//可以在连续调用Parse中重新定义模板。
	//具有仅包含空白和注释的主体的模板定义被认为是空的，不会替换现有模板的主体。
	//这允许使用Parse添加新的命名模板定义，而不会覆盖主模板主体。

	// Create a new template and parse the letter into it.
	//创建一个新模板并将解析letter。
	t := template.Must(template.New("letter").Parse(letter))
	// Execute the template for each recipient.
	for _, r := range recipients { //遍历填充物
		// Execute applies a parsed template to the specified data object,
		// and writes the output to wr.
		// If an error occurs executing the template or writing its output,
		// execution stops, but partial results may already have been written to
		// the output writer.
		// A template may be executed safely in parallel, although if parallel
		// executions share a Writer the output may be interleaved.
		//
		// If data is a reflect.Value, the template applies to the concrete
		// value that the reflect.Value holds, as in fmt.Print.
		// Execute将已解析的模板应用于第二个参数指定的data数据对象，并将输出写入第一个参数wr。
		//如果执行模板或写入模板输出时发生错误，则执行将停止，但是可能已将部分结果写入输出写入器。
		//尽管并行执行共享一个Writer且输出可能会交错，但是可以安全地并行执行模板。
		//如果参数data是reflect.Value，则该模板将应用于reflect.Value所保存的具体值，如fmt.Print中所示。
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("解析template失败:", err)
		} else {
			fmt.Println("解析template成功！")
		}
		fmt.Println("---------")
	}

	//模板1输出：
	//
	//	Dear 名字1,
	//
	//	It was a pleasure to see you at the wedding.
	//	Thank you for the lovely 书籍.
	//
	//	Best wishes,
	//	Josie
	//	解析template成功！
	//	---------
	//
	//	Dear 名字2,
	//
	//	It is a shame you couldn't make it to the wedding.
	//	Thank you for the lovely 石头.
	//
	//	Best wishes,
	//	Josie
	//	解析template成功！
	//	---------
	//
	//	Dear 名字3,
	//
	//	It is a shame you couldn't make it to the wedding.
	//
	//	Best wishes,
	//	Josie
	//	解析template成功！
	//	---------

	//模板2输出：
	//
	//	Dear 名字1,
	//	It was a pleasure to see you at the wedding.
	//	Thank you for the lovely 书籍.
	//
	//	Best wishes,
	//	Josie
	//	解析template成功！
	//	---------
	//
	//	Dear 名字2,
	//	It is a shame you couldn't make it to the wedding.
	//	Thank you for the lovely 石头.
	//
	//	Best wishes,
	//	Josie
	//	解析template成功！
	//	---------
	//
	//	Dear 名字3,
	//	It is a shame you couldn't make it to the wedding.
	//
	//	Best wishes,
	//	Josie
	//	解析template成功！
	//	---------

	//模板3输出：
	//	Dear 名字1,It was a pleasure to see you at the wedding.Thank you for the lovely 书籍.Best wishes,Josie解析template成功！
	//	---------
	//	Dear 名字2,It is a shame you couldn't make it to the wedding.Thank you for the lovely 石头.Best wishes,Josie解析template成功！
	//	---------
	//	Dear 名字3,It is a shame you couldn't make it to the wedding.Best wishes,Josie解析template成功！
	//	---------

	//对比上面的3个模板可以看出：模板挖空并不会自动换行，但是如果使用反引号括上模板字符串的话则会导致换行

	fmt.Println("--------------template对象111-------------------")

	//// NodeType identifies the type of a parse tree node.
	//// NodeType标识解析树节点的类型。
	//type NodeType int
	//
	//const (
	//	NodeText       NodeType = iota // Plain text.// 纯文本。
	//	NodeAction                     // A non-control action such as a field evaluation.//非控制动作，例如field评估。
	//	NodeBool                       // A boolean constant.//一个布尔常量。
	//	NodeChain                      // A sequence of field accesses.//一系列的字段访问。
	//	NodeCommand                    // An element of a pipeline.//管道的元素。
	//	NodeDot                        // The cursor, dot.//光标，点。
	//	nodeElse                       // An else action. Not added to tree.//其他动作。 没有添加到树中。
	//	nodeEnd                        // An end action. Not added to tree.//结束动作。 没有添加到树中。
	//	NodeField                      // A field or method name.//字段或方法名称。
	//	NodeIdentifier                 // An identifier; always a function name.//标识符； 始终是函数名称。
	//	NodeIf                         // An if action.// if动作
	//	NodeList                       // A list of Nodes.//节点列表。
	//	NodeNil                        // An untyped nil constant.//无类型的nil常数。
	//	NodeNumber                     // A numerical constant.//一个数字常数。
	//	NodePipe                       // A pipeline of commands.//命令管道。
	//	NodeRange                      // A range action.//range范围动作。
	//	NodeString                     // A string constant.//字符串常量。
	//	NodeTemplate                   // A template invocation action.//模板调用动作。
	//	NodeVariable                   // A $ variable.//一个$变量。
	//	NodeWith                       // A with action.//一个with动作
	//)
	//
	//// Nodes.
	//
	//// ListNode holds a sequence of nodes.
	//// ListNode包含一系列节点。
	//type ListNode struct {
	//	NodeType
	//	Pos
	//	tr    *Tree		//属于哪棵树下的节点列表
	//	Nodes []Node // The element nodes in lexical order.	（//元素节点按词法顺序。）
	//}

	//// Pos represents a byte position in the original input text from which
	//// this template was parsed.
	//// Pos表示解析此模板的原始输入文本中的字节位置。
	//type Pos int

	////func (p Pos) Position() Pos {
	////	return p
	////}

	//// NodeType identifies the type of a parse tree node.
	//// NodeType标识解析树节点的类型。
	//type NodeType int

	//// Type returns itself and provides an easy default implementation
	//// for embedding in a Node. Embedded in all non-trivial Nodes.
	////Type返回自身，并提供了一个简单的默认实现以嵌入到Node中。 嵌入所有不重要的节点中。
	//func (t NodeType) Type() NodeType {
	//	return t
	//}

	for _, r := range recipients { //遍历填充物

		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("解析template失败:", err)
		} else {
			fmt.Println("解析template成功！")
		}
		//解析过程中顶级模板的名称，用于显示错误消息。属于Tree的字段
		fmt.Printf("ParseName:%#v\n", t.ParseName)
		//树的顶级根。
		ls_node := t.Root
		//这里千万不要去掉*号，因为如果去掉星号，那么就会输出调用这个对象上面的string()方法返回的值（因为他只实现了指针
		// 类型的String()方法，而没有实现值类型的String()方法），但是他跟我想要输出的ls_node对象上面的每个字段的值是不同的！
		//因为Printf（）打印的参数对象如果实现了String()方法，那么就会输出这个对象上面的String()方法返回的值
		//下面我们就是不想输出String()的返回值
		fmt.Printf("===Root:%+v\n", *ls_node)
		fmt.Printf("===ls_node.String():%+v\n", ls_node.String())
		fmt.Printf("===ls_node.Nodes:%+v\n", ls_node.Nodes)
		fmt.Printf("===ls_node.Pos:%+v\n", ls_node.Pos)
		fmt.Printf("===ls_node.NodeType:%+v\n", ls_node.NodeType)     //11代表NodeList （A list of Nodes.）
		fmt.Printf("===ls_node.Position():%+v\n", ls_node.Position()) //这个方法是type Pos int对象的方法
		fmt.Printf("===ls_node.CopyList():%+v\n", ls_node.CopyList()) //对每个Node对象进行深层的复制，然后组成一个新的ListNode对象
		//底层实现：return l.CopyList()，注意，ListNode对象实现了接口Node，所以下面的Copy虽然要求返回的是Node，但是我们可以返回Node的实现类ListNode对象
		//他跟上面的CopyList（）还是有本质的区别，返回值的类型不同，下面是返回一个接口类型Node,而上面返回的是一个具体的类ListNode对象
		fmt.Printf("===ls_node.Copy():%+v\n", ls_node.Copy())
		fmt.Printf("===ls_node.Type():%+v\n", ls_node.Type()) //这是type NodeType int对象上面的方法，跟上面的NodeType属性是一样的

		//上面都是讲到的ListNode对象或者他的组合对象的方法和属性，下面我们继续讲解Template对象的方法和属性

		// Copy returns a copy of the Tree. Any parsing state is discarded.
		//复制返回树的副本。 任何解析状态都将被丢弃。Template组合了parse.tree树对象,所以Copy（）方法是Tree的方法
		tree := t.Copy()
		fmt.Printf("===t.Copy():%+v\n", tree)
		fmt.Printf("===tree.Name:%+v\n", tree.Name)           ////树表示的模板的名称。
		fmt.Printf("===tree.ParseName:%+v\n", tree.ParseName) ////解析过程中顶级模板的名称，用于显示错误消息。
		fmt.Printf("===tree.Root:%+v\n", tree.Root)           //////树的顶级根。

		letter1 := `亲爱的：{{.Name}},{{if .Attended}}这是一份很好的礼物送给你{{else}}我没有礼物给你，抱歉{{end}}{{with .Gift}}多谢你之前送我的礼物。 {{.}}.{{end}}Best wishes,anko`
		//		letter1 := `
		//Dear {{.Name}},
		//{{if .Attended}}
		//It was a pleasure to see you at the wedding.{{else}}
		//It is a shame you couldn't make it to the wedding.{{end}}
		//{{with .Gift}}Thank you for the lovely {{.}}.
		//{{end}}
		//Best wishes,
		//Josie
		//`
		// Parse parses the template definition string to construct a representation of
		// the template for execution. If either action delimiter string is empty, the
		// default ("{{" or "}}") is used. Embedded template definitions are added to
		// the treeSet map.
		// Parse解析模板定义字符串，以构造要执行的模板的表示形式。 如果任何一个动作分隔符字符串为空，则使用默认值（“ {{”或“}}”）。 嵌入式模板定义将添加到treeSet映射中。

		//parse.Tree没实现任何接口
		tree_cpTempl, err := tree.Parse(letter1, "", "", make(map[string]*parse.Tree), nil)
		check_err_template(err)
		fmt.Printf("===tree.Parse():%+v\n", tree_cpTempl)

		// ErrorContext returns a textual representation of the location of the node in the input text.
		// The receiver is only used when the node does not have a pointer to the tree inside,
		// which can occur in old code.
		// ErrorContext返回输入文本中节点位置的文本表示形式。
		//仅当节点内部没有指向树的指针时才使用接收器，这可能会在旧代码中出现。
		//Node接口组合了parse.Tree类对象的字段
		for k, v := range tree.Root.Nodes {
			location, context := tree.ErrorContext(v)

			//letter1的输入格式会严重影响下面的解析的结果，随着输入格式的不同，下面输出的解析结果也会不同的，比如上面将letter1改成下行的已经注释掉的letter1的话则输出结果会不一样！
			fmt.Printf("===tree.ErrorContext()解析第%v个node后的返回值的第1个参数location:%+v\n", k, location) //一个命令的位置，并不是一行代表一个命令，位置从哪里开始也会标志
			fmt.Printf("===tree.ErrorContext()解析第%v个node后的返回值的第2个参数context:%+v\n", k, context)   //该位置下的内容
		}

		//上面讲解到tree对象，下面我们继续讲解Template对象

		//因为template对象组合了tree对象，所以下面关于通过t调用的tree对象上面的方法不再讲解，这些方法有：
		//t.ErrorContext()
		//t.Copy()

		//下面的2个方法将在下面的testTemplateParseGlobAndParseFiles()方法中讲到，这里不再讲解，
		// 值得注意的是，一个已经解析之后的模板还是可以继续解析的！实质是更新源模板template的字段common,在common对象
		// 的tmpl字段上面追加一个新的模板，但是跟输出结果没有什么关系， common保存相关模板共享的信息。所以输出的模板还是原来的模板，
		// 总之这2个方法不会追加解析树节点，只是会在源节点上面进行更新共享的信息（共享信息中有个字段是模板对象，没错，就是更新这个对象）

		//底层：associate（）方法
		//associate将新模板安装到与t关联的模板组中。 已知这两者共享相同的结构。
		//布尔值返回值报告是否将此树存储为t.Tree。

		//下面的这个方法跟ParseGlob（）方法类似，我们只讲ParseGlob（）方法（事实上我们已经讲了函数类型的ParseFiles()，对象类型的ParseFiles()方法其实是类似的）
		//t.ParseFiles()

		//如果你点开Must源码，你会发现其实他除了捕捉异常并且抛出异常之外，其他东西都没做了
		fmt.Println("***下面是输出再解析的内容****")
		//t如果是已经绑定模板的template对象，则输出的还是原来绑定的模板,而不会是main3\mytemp\*.txt中定义的模板。
		//注意下面的t已经绑定过模板的了，如果我们再次调用t.ParseXxxx的话则会继续绑定模板，这些模板我们都可以查找到
		t1 := template.Must(t.ParseGlob(`main3\mytemp\*.txt`))
		// Lookup returns the template with the given name that is associated with t.
		// It returns nil if there is no such template or the template has no definition.
		//查找返回具有与t关联的给定名称的模板。
		//如果没有这样的模板或模板没有定义，则返回nil。
		T_lookup := t1.Lookup("letter")
		fmt.Printf("===t.Lookup(\"letter\"):%+v\n", T_lookup)

		fmt.Printf("===t.Lookup(\"letter\"):%+v\n", T_lookup.Root.String())
		T_lookup = t1.Lookup("templ1.txt")
		fmt.Printf("===t.Lookup(\"templ1.txt\"):%+v\n", T_lookup)

		fmt.Printf("===t.Lookup(\"templ1.txt\"):%+v\n", T_lookup.Root.String())
		T_lookup = t1.Lookup("templ2.txt")
		fmt.Printf("===t.Lookup(\"templ2.txt\"):%+v\n", T_lookup)

		//通过访问模板下面的节点列表的字符串形式来输出绑定的模板的字符串参数
		fmt.Printf("===t.Lookup(\"templ2.txt\"):%+v\n", T_lookup.Root.String())

		//如果访问不存在的模板的话则会返回nil
		T_lookup = t1.Lookup("templ3.txt")
		fmt.Printf("===t.Lookup(\"templ3.txt\"):%+v\n", T_lookup)
		if T_lookup != nil {
			//千万不要为nil的时候执行下面的语句，会报错
			fmt.Printf("===t.Lookup(\"templ3.txt\"):%+v\n", T_lookup.Root.String())
		}

		//抛出异常信息：template: : "" is an incomplete or empty template，在Execute（）执行时候会抛出异常
		//t1 := template.Must(new(template.Template).ParseGlob(`main3\mytemp\*.txt`))

		//t1 := template.Must(template.New("templ1.txt").ParseGlob(`main3\mytemp\*.txt`))
		//但是，注意了，我们采用new()创建模板时候指定的模板名必须是跟我要解析的模板的名字是相同的，也就是跟ParseGlob（）参数正则
		// 匹配到的名字相同（注意是一定要相同，而不是匹配正则就可以了），否则会报错
		//t1 := template.Must(template.New("templ2.txt").ParseGlob(`main3\mytemp\*.txt`))

		//与上面几乎相同的.ParseFiles方法，同时也是上一个方法的底层,同理，我们在给定.ParseFiles（）参数列表时候必须给定一个跟template.name相同的路径字符串，
		// 否则会报错，参数列表的先后顺序不影响，
		//t1 := template.Must(template.New("templ1.txt").ParseFiles(`main3\mytemp\templ1.txt`,`main3\mytemp\templ2.txt`))
		//参数列表的先后顺序不影响，所以下面也是正确的
		//t1 := template.Must(template.New("templ1.txt").ParseFiles(`main3\mytemp\templ2.txt`,`main3\mytemp\templ1.txt`))
		//下面会报错
		//t1 := template.Must(template.New("templ1.txt").ParseFiles(`main3\mytemp\templ2.txt`)
		//下面这个才是正确的
		//t1 := template.Must(template.New("templ2.txt").ParseFiles(`main3\mytemp\templ2.txt`))

		//通过实例对象调用的new同理，我们不再累叙了

		var recipients1 = []Recipient{
			{"名字11", "书籍11", true},
			{"名字22", "石头22", false},
			{"名字33", "", false},
		}

		fmt.Println("---下面对一个template对象绑定的3个模板对象进行执行填充---")
		for _, r1 := range recipients1 { //遍历填充物
			// Execute applies a parsed template to the specified data object,
			// and writes the output to wr.
			// If an error occurs executing the template or writing its output,
			// execution stops, but partial results may already have been written to
			// the output writer.
			// A template may be executed safely in parallel, although if parallel
			// executions share a Writer the output may be interleaved.
			//
			// If data is a reflect.Value, the template applies to the concrete
			// value that the reflect.Value holds, as in fmt.Print.
			// Execute将已解析的模板应用于第二个参数指定的data数据对象，并将输出写入第一个参数wr。
			//如果执行模板或写入模板输出时发生错误，则执行将停止，但是可能已将部分结果写入输出写入器。
			//尽管并行执行共享一个Writer且输出可能会交错，但是可以安全地并行执行模板。
			//如果参数data是reflect.Value，则该模板将应用于reflect.Value所保存的具体值，如fmt.Print中所示。
			//底层主要实现是：walk()方法
			//myerr := t1.Execute(os.Stdout, r1)

			//除了上面的解析到默认的模板之外，还可以解析到指定名字的模板，如下

			// ExecuteTemplate applies the template associated with t that has the given name
			// to the specified data object and writes the output to wr.
			// If an error occurs executing the template or writing its output,
			// execution stops, but partial results may already have been written to
			// the output writer.
			// A template may be executed safely in parallel, although if parallel
			// executions share a Writer the output may be interleaved.
			// ExecuteTemplate将与具有给定名称的t关联的模板应用于指定的数据对象，并将输出写入wr。
			//如果执行模板或写入模板输出时发生错误，则执行将停止，但是可能已将部分结果写入输出写入器。
			//尽管并行执行共享一个Writer，且输出可能会交错，但仍然可以安全地并行执行模板。

			//为了解析上面t绑定的模板（letter，templ1.txt，templ2.txt），你可以分别松开下面的3条语句中的一条，分别进行尝试：
			myerr := t1.ExecuteTemplate(os.Stdout, "letter", r1)
			//myerr := t1.ExecuteTemplate(os.Stdout,"templ1.txt", r1)
			//myerr := t1.ExecuteTemplate(os.Stdout,"templ2.txt", r1)
			//当然上面我解析的填充对象不一样是r1，也可以另外用其他类似的对象，这个你自行测试即可

			//假如我们给一个未绑定到template对象上的模板名称的话，则会报错
			//myerr := t1.ExecuteTemplate(os.Stdout,"templ3.txt", r1)

			if myerr != nil {
				log.Println("^-^\ttemplate再解析失败:", myerr)
			} else {
				fmt.Println("^-^\ttemplate再解析成功！")
			}

			//tree1:=t1.Copy()
			//for k, v := range tree1.Root.Nodes {
			//	location, context := tree1.ErrorContext(v)
			//
			//	//letter1的输入格式会严重影响下面的解析的结果，随着输入格式的不同，下面输出的解析结果也会不同的，比如上面将letter1改成下行的已经注释掉的letter1的话则输出结果会不一样！
			//	fmt.Printf("===tree1.ErrorContext()解析第%v个node后的返回值的第1个参数location:%+v\n", k, location) //一个命令的位置，并不是一行代表一个命令，位置从哪里开始也会标志
			//	fmt.Printf("===tree1.ErrorContext()解析第%v个node后的返回值的第2个参数context:%+v\n", k, context)   //该位置下的内容
			//}

		}

		fmt.Println("***上面是输出再解析的内容****")

		//以上基于模板1输出：
		//--------------template对象111-------------------
		//			Dear 名字1,
		//
		//				It was a pleasure to see you at the wedding.
		//				Thank you for the lovely 书籍.
		//
		//				Best wishes,
		//				Josie
		//			解析template成功！
		//		ParseName:"letter"
		//			===Root:{NodeType:11 Pos:0 tr:0xc00009e000 Nodes:[
		//			Dear  {{.Name}} ,
		//			{{if .Attended}}
		//			It was a pleasure to see you at the wedding.{{else}}
		//			It is a shame you couldn't make it to the wedding.{{end}}
		//			{{with .Gift}}Thank you for the lovely {{.}}.
		//			{{end}}
		//			Best wishes,
		//				Josie
		//	]}
		//	===ls_node.String():
		//	Dear {{.Name}},
		//	{{if .Attended}}
		//	It was a pleasure to see you at the wedding.{{else}}
		//	It is a shame you couldn't make it to the wedding.{{end}}
		//	{{with .Gift}}Thank you for the lovely {{.}}.
		//	{{end}}
		//	Best wishes,
		//	Josie
		//
		//	===ls_node.Nodes:[
		//	Dear  {{.Name}} ,
		//	{{if .Attended}}
		//	It was a pleasure to see you at the wedding.{{else}}
		//	It is a shame you couldn't make it to the wedding.{{end}}
		//	{{with .Gift}}Thank you for the lovely {{.}}.
		//	{{end}}
		//	Best wishes,
		//	Josie
		//	]
		//	===ls_node.Pos:0
		//	===ls_node.NodeType:11
		//	===ls_node.Position():0
		//	===ls_node.CopyList():
		//	Dear {{.Name}},
		//	{{if .Attended}}
		//	It was a pleasure to see you at the wedding.{{else}}
		//	It is a shame you couldn't make it to the wedding.{{end}}
		//	{{with .Gift}}Thank you for the lovely {{.}}.
		//	{{end}}
		//	Best wishes,
		//	Josie
		//
		//	===ls_node.Copy():
		//	Dear {{.Name}},
		//	{{if .Attended}}
		//	It was a pleasure to see you at the wedding.{{else}}
		//	It is a shame you couldn't make it to the wedding.{{end}}
		//	{{with .Gift}}Thank you for the lovely {{.}}.
		//	{{end}}
		//	Best wishes,
		//	Josie
		//
		//	===ls_node.Type():11
		//	===t.Copy():&{Name:letter ParseName:letter Root:
		//	Dear {{.Name}},
		//	{{if .Attended}}
		//	It was a pleasure to see you at the wedding.{{else}}
		//	It is a shame you couldn't make it to the wedding.{{end}}
		//	{{with .Gift}}Thank you for the lovely {{.}}.
		//	{{end}}
		//	Best wishes,
		//	Josie
		//	text:
		//	Dear {{.Name}},
		//	{{if .Attended}}
		//	It was a pleasure to see you at the wedding.{{else}}
		//	It is a shame you couldn't make it to the wedding.{{end}}
		//	{{with .Gift}}Thank you for the lovely {{.}}.
		//	{{end}}
		//	Best wishes,
		//	Josie
		//	funcs:[] lex:<nil> token:[{typ:0 pos:0 val: line:0} {typ:0 pos:0 val: line:0} {typ:0 pos:0 val: line:0}] peekCount:0 vars:[] treeSet:map[]}
		//	===tree.Name:letter
		//	===tree.ParseName:letter
		//	===tree.Root:
		//	Dear {{.Name}},
		//	{{if .Attended}}
		//	It was a pleasure to see you at the wedding.{{else}}
		//	It is a shame you couldn't make it to the wedding.{{end}}
		//	{{with .Gift}}Thank you for the lovely {{.}}.
		//	{{end}}
		//	Best wishes,
		//	Josie
		//
		//	===tree.Parse():&{Name:letter ParseName:letter Root:亲爱的：{{.Name}},{{if .Attended}}这是一份很好的礼物送给你{{else}}我没有礼物给你，抱歉{{end}}{{with .Gift}}多谢你之前送我的礼物。 {{.}}.{{end}}Best wishes,anko text:亲爱的：{{.Name}},{{if .Attended}}这是一份很好的礼物送给你{{else}}我没有礼物给你，抱歉{{end}}{{with .Gift}}多谢你之前送我的礼物。 {{.}}.{{end}}Best wishes,anko funcs:[] lex:<nil> token:[{typ:7 pos:196 val: line:1} {typ:10 pos:119 val:{{ line:1} {typ:0 pos:0 val: line:0}] peekCount:1 vars:[] treeSet:map[]}
		//	===tree.ErrorContext()解析第0个node后的返回值的第1个参数location:letter:1:0
		//	===tree.ErrorContext()解析第0个node后的返回值的第2个参数context:亲爱的：
		//	===tree.ErrorContext()解析第1个node后的返回值的第1个参数location:letter:1:14
		//	===tree.ErrorContext()解析第1个node后的返回值的第2个参数context:{{.Name}}
		//	===tree.ErrorContext()解析第2个node后的返回值的第1个参数location:letter:1:21
		//	===tree.ErrorContext()解析第2个node后的返回值的第2个参数context:,
		//	===tree.ErrorContext()解析第3个node后的返回值的第1个参数location:letter:1:27
		//	===tree.ErrorContext()解析第3个node后的返回值的第2个参数context:{{if .Attended}}这是一份很好的礼物送给你{{else}}我没有礼物给你，抱歉{{end}}
		//	===tree.ErrorContext()解析第4个node后的返回值的第1个参数location:letter:1:126
		//	===tree.ErrorContext()解析第4个node后的返回值的第2个参数context:{{with .Gift}}多谢你之前送我的礼物。 {{.}}.{{end}}
		//	===tree.ErrorContext()解析第5个node后的返回值的第1个参数location:letter:1:180
		//	===tree.ErrorContext()解析第5个node后的返回值的第2个参数context:Best wishes,anko
		//	***下面是输出再解析的内容****
		//	===t.Lookup("letter"):&{name:letter Tree:0xc00009e000 common:0xc00003e100 leftDelim: rightDelim:}
		//	===t.Lookup("letter"):
		//	Dear {{.Name}},
		//	{{if .Attended}}
		//	It was a pleasure to see you at the wedding.{{else}}
		//	It is a shame you couldn't make it to the wedding.{{end}}
		//	{{with .Gift}}Thank you for the lovely {{.}}.
		//	{{end}}
		//	Best wishes,
		//	Josie
		//
		//	===t.Lookup("templ1.txt"):&{name:templ1.txt Tree:0xc0000d6100 common:0xc00003e100 leftDelim: rightDelim:}
		//	===t.Lookup("templ1.txt"):亲爱的：{{.Name}}
		//	{{if .Attended}}
		//	It was 这是我自定义的模板templ111,a pleasure to see you at the wedding.{{else}}
		//	It is a shame you couldn't make it to the wedding.{{end}}
		//	{{with .Gift}}Thank you for the lovely {{.}}.
		//	{{end}}
		//
		//	===t.Lookup("templ2.txt"):&{name:templ2.txt Tree:0xc00009e300 common:0xc00003e100 leftDelim: rightDelim:}
		//	===t.Lookup("templ2.txt"):亲爱的：{{.Name}}
		//	{{if .Attended}}
		//	It was 这是我自定义的模板templ222,a pleasure to see you at the wedding.{{else}}
		//	It is a shame you couldn't make it to the wedding.{{end}}
		//	{{with .Gift}}Thank you for the lovely {{.}}.
		//	{{end}}
		//
		//	===t.Lookup("templ3.txt"):<nil>
		//	---下面对一个template对象绑定的3个模板对象进行执行填充---
		//
		//	Dear 名字11,
		//
		//	It was a pleasure to see you at the wedding.
		//	Thank you for the lovely 书籍11.
		//
		//	Best wishes,
		//	Josie
		//	^-^	template再解析成功！
		//
		//	Dear 名字22,
		//
		//	It is a shame you couldn't make it to the wedding.
		//	Thank you for the lovely 石头22.
		//
		//	Best wishes,
		//	Josie
		//	^-^	template再解析成功！
		//
		//	Dear 名字33,
		//
		//	It is a shame you couldn't make it to the wedding.
		//
		//	Best wishes,
		//	Josie
		//	^-^	template再解析成功！
		//	***上面是输出再解析的内容****
		//	---------
		//
		//	Dear 名字2,
		//
		//	It is a shame you couldn't make it to the wedding.
		//	Thank you for the lovely 石头.
		//
		//	Best wishes,
		//	Josie
		//	解析template成功！
		//	ParseName:"letter"
		//	===Root:{NodeType:11 Pos:0 tr:0xc00009e000 Nodes:[
		//	Dear  {{.Name}} ,
		//	{{if .Attended}}
		//	It was a pleasure to see you at the wedding.{{else}}
		//	It is a shame you couldn't make it to the wedding.{{end}}
		//	{{with .Gift}}Thank you for the lovely {{.}}.
		//	{{end}}
		//	Best wishes,
		//	Josie
		//	]}
		//	===ls_node.String():
		//	Dear {{.Name}},
		//	{{if .Attended}}
		//	It was a pleasure to see you at the wedding.{{else}}
		//	It is a shame you couldn't make it to the wedding.{{end}}
		//	{{with .Gift}}Thank you for the lovely {{.}}.
		//	{{end}}
		//	Best wishes,
		//	Josie
		//
		//	===ls_node.Nodes:[
		//	Dear  {{.Name}} ,
		//	{{if .Attended}}
		//	It was a pleasure to see you at the wedding.{{else}}
		//	It is a shame you couldn't make it to the wedding.{{end}}
		//	{{with .Gift}}Thank you for the lovely {{.}}.
		//	{{end}}
		//	Best wishes,
		//	Josie
		//	]
		//	===ls_node.Pos:0
		//	===ls_node.NodeType:11
		//	===ls_node.Position():0
		//	===ls_node.CopyList():
		//	Dear {{.Name}},
		//	{{if .Attended}}
		//	It was a pleasure to see you at the wedding.{{else}}
		//	It is a shame you couldn't make it to the wedding.{{end}}
		//	{{with .Gift}}Thank you for the lovely {{.}}.
		//	{{end}}
		//	Best wishes,
		//	Josie
		//
		//	===ls_node.Copy():
		//	Dear {{.Name}},
		//	{{if .Attended}}
		//	It was a pleasure to see you at the wedding.{{else}}
		//	It is a shame you couldn't make it to the wedding.{{end}}
		//	{{with .Gift}}Thank you for the lovely {{.}}.
		//	{{end}}
		//	Best wishes,
		//	Josie
		//
		//	===ls_node.Type():11
		//	===t.Copy():&{Name:letter ParseName:letter Root:
		//	Dear {{.Name}},
		//	{{if .Attended}}
		//	It was a pleasure to see you at the wedding.{{else}}
		//	It is a shame you couldn't make it to the wedding.{{end}}
		//	{{with .Gift}}Thank you for the lovely {{.}}.
		//	{{end}}
		//	Best wishes,
		//	Josie
		//	text:
		//	Dear {{.Name}},
		//	{{if .Attended}}
		//	It was a pleasure to see you at the wedding.{{else}}
		//	It is a shame you couldn't make it to the wedding.{{end}}
		//	{{with .Gift}}Thank you for the lovely {{.}}.
		//	{{end}}
		//	Best wishes,
		//	Josie
		//	funcs:[] lex:<nil> token:[{typ:0 pos:0 val: line:0} {typ:0 pos:0 val: line:0} {typ:0 pos:0 val: line:0}] peekCount:0 vars:[] treeSet:map[]}
		//	===tree.Name:letter
		//	===tree.ParseName:letter
		//	===tree.Root:
		//	Dear {{.Name}},
		//	{{if .Attended}}
		//	It was a pleasure to see you at the wedding.{{else}}
		//	It is a shame you couldn't make it to the wedding.{{end}}
		//	{{with .Gift}}Thank you for the lovely {{.}}.
		//	{{end}}
		//	Best wishes,
		//	Josie
		//
		//	===tree.Parse():&{Name:letter ParseName:letter Root:亲爱的：{{.Name}},{{if .Attended}}这是一份很好的礼物送给你{{else}}我没有礼物给你，抱歉{{end}}{{with .Gift}}多谢你之前送我的礼物。 {{.}}.{{end}}Best wishes,anko text:亲爱的：{{.Name}},{{if .Attended}}这是一份很好的礼物送给你{{else}}我没有礼物给你，抱歉{{end}}{{with .Gift}}多谢你之前送我的礼物。 {{.}}.{{end}}Best wishes,anko funcs:[] lex:<nil> token:[{typ:7 pos:196 val: line:1} {typ:10 pos:119 val:{{ line:1} {typ:0 pos:0 val: line:0}] peekCount:1 vars:[] treeSet:map[]}
		//	===tree.ErrorContext()解析第0个node后的返回值的第1个参数location:letter:1:0
		//	===tree.ErrorContext()解析第0个node后的返回值的第2个参数context:亲爱的：
		//	===tree.ErrorContext()解析第1个node后的返回值的第1个参数location:letter:1:14
		//	===tree.ErrorContext()解析第1个node后的返回值的第2个参数context:{{.Name}}
		//	===tree.ErrorContext()解析第2个node后的返回值的第1个参数location:letter:1:21
		//	===tree.ErrorContext()解析第2个node后的返回值的第2个参数context:,
		//	===tree.ErrorContext()解析第3个node后的返回值的第1个参数location:letter:1:27
		//	===tree.ErrorContext()解析第3个node后的返回值的第2个参数context:{{if .Attended}}这是一份很好的礼物送给你{{else}}我没有礼物给你，抱歉{{end}}
		//	===tree.ErrorContext()解析第4个node后的返回值的第1个参数location:letter:1:126
		//	===tree.ErrorContext()解析第4个node后的返回值的第2个参数context:{{with .Gift}}多谢你之前送我的礼物。 {{.}}.{{end}}
		//	===tree.ErrorContext()解析第5个node后的返回值的第1个参数location:letter:1:180
		//	===tree.ErrorContext()解析第5个node后的返回值的第2个参数context:Best wishes,anko
		//	***下面是输出再解析的内容****
		//	===t.Lookup("letter"):&{name:letter Tree:0xc00009e000 common:0xc00003e100 leftDelim: rightDelim:}
		//	===t.Lookup("letter"):
		//	Dear {{.Name}},
		//	{{if .Attended}}
		//	It was a pleasure to see you at the wedding.{{else}}
		//	It is a shame you couldn't make it to the wedding.{{end}}
		//	{{with .Gift}}Thank you for the lovely {{.}}.
		//	{{end}}
		//	Best wishes,
		//	Josie
		//
		//	===t.Lookup("templ1.txt"):&{name:templ1.txt Tree:0xc0000d6300 common:0xc00003e100 leftDelim: rightDelim:}
		//	===t.Lookup("templ1.txt"):亲爱的：{{.Name}}
		//	{{if .Attended}}
		//	It was 这是我自定义的模板templ111,a pleasure to see you at the wedding.{{else}}
		//	It is a shame you couldn't make it to the wedding.{{end}}
		//	{{with .Gift}}Thank you for the lovely {{.}}.
		//	{{end}}
		//
		//	===t.Lookup("templ2.txt"):&{name:templ2.txt Tree:0xc0000d6400 common:0xc00003e100 leftDelim: rightDelim:}
		//	===t.Lookup("templ2.txt"):亲爱的：{{.Name}}
		//	{{if .Attended}}
		//	It was 这是我自定义的模板templ222,a pleasure to see you at the wedding.{{else}}
		//	It is a shame you couldn't make it to the wedding.{{end}}
		//	{{with .Gift}}Thank you for the lovely {{.}}.
		//	{{end}}
		//
		//	===t.Lookup("templ3.txt"):<nil>
		//	---下面对一个template对象绑定的3个模板对象进行执行填充---
		//
		//	Dear 名字11,
		//
		//	It was a pleasure to see you at the wedding.
		//	Thank you for the lovely 书籍11.
		//
		//	Best wishes,
		//	Josie
		//	^-^	template再解析成功！
		//
		//	Dear 名字22,
		//
		//	It is a shame you couldn't make it to the wedding.
		//	Thank you for the lovely 石头22.
		//
		//	Best wishes,
		//	Josie
		//	^-^	template再解析成功！
		//
		//	Dear 名字33,
		//
		//	It is a shame you couldn't make it to the wedding.
		//
		//	Best wishes,
		//	Josie
		//	^-^	template再解析成功！
		//	***上面是输出再解析的内容****
		//	---------
		//
		//	Dear 名字3,
		//
		//	It is a shame you couldn't make it to the wedding.
		//
		//	Best wishes,
		//	Josie
		//	解析template成功！
		//	ParseName:"letter"
		//	===Root:{NodeType:11 Pos:0 tr:0xc00009e000 Nodes:[
		//	Dear  {{.Name}} ,
		//	{{if .Attended}}
		//	It was a pleasure to see you at the wedding.{{else}}
		//	It is a shame you couldn't make it to the wedding.{{end}}
		//	{{with .Gift}}Thank you for the lovely {{.}}.
		//	{{end}}
		//	Best wishes,
		//	Josie
		//	]}
		//	===ls_node.String():
		//	Dear {{.Name}},
		//	{{if .Attended}}
		//	It was a pleasure to see you at the wedding.{{else}}
		//	It is a shame you couldn't make it to the wedding.{{end}}
		//	{{with .Gift}}Thank you for the lovely {{.}}.
		//	{{end}}
		//	Best wishes,
		//	Josie
		//
		//	===ls_node.Nodes:[
		//	Dear  {{.Name}} ,
		//	{{if .Attended}}
		//	It was a pleasure to see you at the wedding.{{else}}
		//	It is a shame you couldn't make it to the wedding.{{end}}
		//	{{with .Gift}}Thank you for the lovely {{.}}.
		//	{{end}}
		//	Best wishes,
		//	Josie
		//	]
		//	===ls_node.Pos:0
		//	===ls_node.NodeType:11
		//	===ls_node.Position():0
		//	===ls_node.CopyList():
		//	Dear {{.Name}},
		//	{{if .Attended}}
		//	It was a pleasure to see you at the wedding.{{else}}
		//	It is a shame you couldn't make it to the wedding.{{end}}
		//	{{with .Gift}}Thank you for the lovely {{.}}.
		//	{{end}}
		//	Best wishes,
		//	Josie
		//
		//	===ls_node.Copy():
		//	Dear {{.Name}},
		//	{{if .Attended}}
		//	It was a pleasure to see you at the wedding.{{else}}
		//	It is a shame you couldn't make it to the wedding.{{end}}
		//	{{with .Gift}}Thank you for the lovely {{.}}.
		//	{{end}}
		//	Best wishes,
		//	Josie
		//
		//	===ls_node.Type():11
		//	===t.Copy():&{Name:letter ParseName:letter Root:
		//	Dear {{.Name}},
		//	{{if .Attended}}
		//	It was a pleasure to see you at the wedding.{{else}}
		//	It is a shame you couldn't make it to the wedding.{{end}}
		//	{{with .Gift}}Thank you for the lovely {{.}}.
		//	{{end}}
		//	Best wishes,
		//	Josie
		//	text:
		//	Dear {{.Name}},
		//	{{if .Attended}}
		//	It was a pleasure to see you at the wedding.{{else}}
		//	It is a shame you couldn't make it to the wedding.{{end}}
		//	{{with .Gift}}Thank you for the lovely {{.}}.
		//	{{end}}
		//	Best wishes,
		//	Josie
		//	funcs:[] lex:<nil> token:[{typ:0 pos:0 val: line:0} {typ:0 pos:0 val: line:0} {typ:0 pos:0 val: line:0}] peekCount:0 vars:[] treeSet:map[]}
		//	===tree.Name:letter
		//	===tree.ParseName:letter
		//	===tree.Root:
		//	Dear {{.Name}},
		//	{{if .Attended}}
		//	It was a pleasure to see you at the wedding.{{else}}
		//	It is a shame you couldn't make it to the wedding.{{end}}
		//	{{with .Gift}}Thank you for the lovely {{.}}.
		//	{{end}}
		//	Best wishes,
		//	Josie
		//
		//	===tree.Parse():&{Name:letter ParseName:letter Root:亲爱的：{{.Name}},{{if .Attended}}这是一份很好的礼物送给你{{else}}我没有礼物给你，抱歉{{end}}{{with .Gift}}多谢你之前送我的礼物。 {{.}}.{{end}}Best wishes,anko text:亲爱的：{{.Name}},{{if .Attended}}这是一份很好的礼物送给你{{else}}我没有礼物给你，抱歉{{end}}{{with .Gift}}多谢你之前送我的礼物。 {{.}}.{{end}}Best wishes,anko funcs:[] lex:<nil> token:[{typ:7 pos:196 val: line:1} {typ:10 pos:119 val:{{ line:1} {typ:0 pos:0 val: line:0}] peekCount:1 vars:[] treeSet:map[]}
		//	===tree.ErrorContext()解析第0个node后的返回值的第1个参数location:letter:1:0
		//	===tree.ErrorContext()解析第0个node后的返回值的第2个参数context:亲爱的：
		//	===tree.ErrorContext()解析第1个node后的返回值的第1个参数location:letter:1:14
		//	===tree.ErrorContext()解析第1个node后的返回值的第2个参数context:{{.Name}}
		//	===tree.ErrorContext()解析第2个node后的返回值的第1个参数location:letter:1:21
		//	===tree.ErrorContext()解析第2个node后的返回值的第2个参数context:,
		//	===tree.ErrorContext()解析第3个node后的返回值的第1个参数location:letter:1:27
		//	===tree.ErrorContext()解析第3个node后的返回值的第2个参数context:{{if .Attended}}这是一份很好的礼物送给你{{else}}我没有礼物给你，抱歉{{end}}
		//	===tree.ErrorContext()解析第4个node后的返回值的第1个参数location:letter:1:126
		//	===tree.ErrorContext()解析第4个node后的返回值的第2个参数context:{{with .Gift}}多谢你之前送我的礼物。 {{.}}.{{end}}
		//	===tree.ErrorContext()解析第5个node后的返回值的第1个参数location:letter:1:180
		//	===tree.ErrorContext()解析第5个node后的返回值的第2个参数context:Best wishes,anko
		//	***下面是输出再解析的内容****
		//	===t.Lookup("letter"):&{name:letter Tree:0xc00009e000 common:0xc00003e100 leftDelim: rightDelim:}
		//	===t.Lookup("letter"):
		//	Dear {{.Name}},
		//	{{if .Attended}}
		//	It was a pleasure to see you at the wedding.{{else}}
		//	It is a shame you couldn't make it to the wedding.{{end}}
		//	{{with .Gift}}Thank you for the lovely {{.}}.
		//	{{end}}
		//	Best wishes,
		//	Josie
		//
		//	===t.Lookup("templ1.txt"):&{name:templ1.txt Tree:0xc00010c300 common:0xc00003e100 leftDelim: rightDelim:}
		//	===t.Lookup("templ1.txt"):亲爱的：{{.Name}}
		//	{{if .Attended}}
		//	It was 这是我自定义的模板templ111,a pleasure to see you at the wedding.{{else}}
		//	It is a shame you couldn't make it to the wedding.{{end}}
		//	{{with .Gift}}Thank you for the lovely {{.}}.
		//	{{end}}
		//
		//	===t.Lookup("templ2.txt"):&{name:templ2.txt Tree:0xc00010c400 common:0xc00003e100 leftDelim: rightDelim:}
		//	===t.Lookup("templ2.txt"):亲爱的：{{.Name}}
		//	{{if .Attended}}
		//	It was 这是我自定义的模板templ222,a pleasure to see you at the wedding.{{else}}
		//	It is a shame you couldn't make it to the wedding.{{end}}
		//	{{with .Gift}}Thank you for the lovely {{.}}.
		//	{{end}}
		//
		//	===t.Lookup("templ3.txt"):<nil>
		//	---下面对一个template对象绑定的3个模板对象进行执行填充---
		//
		//	Dear 名字11,
		//
		//	It was a pleasure to see you at the wedding.
		//	Thank you for the lovely 书籍11.
		//
		//	Best wishes,
		//	Josie
		//	^-^	template再解析成功！
		//
		//	Dear 名字22,
		//
		//	It is a shame you couldn't make it to the wedding.
		//	Thank you for the lovely 石头22.
		//
		//	Best wishes,
		//	Josie
		//	^-^	template再解析成功！
		//
		//	Dear 名字33,
		//
		//	It is a shame you couldn't make it to the wedding.
		//
		//	Best wishes,
		//	Josie
		//	^-^	template再解析成功！
		//	***上面是输出再解析的内容****





		fmt.Println("---------")

	}

	//下面我们测试Template.ParseGlob()和ParseFiles（）方法,
	testTemplateParseGlobAndParseFiles()
	fmt.Println()
	////输出：
	//dir: main3\template800184863
	//pattern: main3\template800184863\*.tmpl
	//T0 invokes T1: (T1 invokes T2: (This is T2))

	fmt.Println("-----下面我们继续讲解关于template对象下的方法和属性：----")

	// Clone returns a duplicate of the template, including all associated
	// templates. The actual representation is not copied, but the name space of
	// associated templates is, so further calls to Parse in the copy will add
	// templates to the copy but not to the original. Clone can be used to prepare
	// common templates and use them with variant definitions for other templates
	// by adding the variants after the clone is made.
	// Clone返回模板的副本，包括所有关联的模板。 不会复制实际的表示形式，但是会复制关联模板的名称空间，因此在副本中进一步调用Parse会将模板添加到副本中，
	// 而不是原始模板中。 克隆可用于准备通用模板，并将其与其他模板的变体定义一起使用，方法是在完成克隆后添加变体。
	TestEmptyTemplateCloneCrash:=func () {
		t1 := template.New("base")
		_, e := t1.Clone() //以前版本会报错这里，如今不会了
		check_err_template(e)
	}
	TestEmptyTemplateCloneCrash()

	fmt.Println("--------------")

	T_clone, err := t.Clone()//这里我们也可以用Must()方法来接收然后提取出template对象，略了！
	check_err_template(err)
	fmt.Printf("===t.Lookup(\"letter\"):%+v\n", t.Lookup("letter"))
	fmt.Printf("===t.Lookup(\"letter\"):%+v\n", t.Lookup("templ1.txt"))
	fmt.Printf("===t.Lookup(\"letter\"):%+v\n", t.Lookup("templ2.txt"))

	fmt.Println()
	//下面通过克隆模板的形式进行查找，为了探究到底是怎么克隆的
	T_clone_lookup := T_clone.Lookup("letter")
	fmt.Printf("===T_clone.Lookup(\"letter\"):%+v\n", T_clone_lookup)
	if T_clone_lookup != nil {
		fmt.Printf("===T_clone.Lookup(\"letter\"):%+v\n", T_clone_lookup.Root.String())
	}

	T_clone_lookup = T_clone.Lookup("templ1.txt")
	fmt.Printf("===T_clone.Lookup(\"templ1.txt\"):%+v\n", T_clone_lookup)
	if T_clone_lookup != nil {
		fmt.Printf("===T_clone.Lookup(\"templ1.txt\"):%+v\n", T_clone_lookup.Root.String())
	}

	T_clone_lookup = T_clone.Lookup("templ2.txt")
	fmt.Printf("===T_clone.Lookup(\"templ2.txt\"):%+v\n", T_clone_lookup)
	if T_clone_lookup != nil {
		fmt.Printf("===T_clone.Lookup(\"templ2.txt\"):%+v\n", T_clone_lookup.Root.String())
	}

	//同理，给定一个不存在的模板名字的话会返回nil值的*template,也就是T_clone_lookup值为nil
	T_clone_lookup = T_clone.Lookup("templ3.txt")
	fmt.Printf("===T_clone.Lookup(\"templ3.txt\"):%+v\n", T_clone_lookup)
	if T_clone_lookup != nil {
		fmt.Printf("===T_clone.Lookup(\"templ3.txt\"):%+v\n", T_clone_lookup.Root.String())
	}

	//同样我们也可以解析，注意如果我们采用源本进行解析的话，那么模板就会绑定到源本上面去，副本的话，那么就会绑定到副本上面去
	_, err = T_clone.Parse(`{{define "templ3.txt"}}这里是templ3.txt模板内容啊{{end}}`)
	check_err_template(err)
	T_clone_lookup = T_clone.Lookup("templ3.txt")
	fmt.Printf("===T_clone.Lookup(\"templ3.txt\"):%+v\n", T_clone_lookup)
	if T_clone_lookup != nil {
		fmt.Printf("===T_clone.Lookup(\"templ3.txt\"):%+v\n", T_clone_lookup.Root.String())
	}

	//下面我们故意没有结束符{{end}}查看是否会抛出什么错误！
	_, err = T_clone.Parse(`{{define "templ33.txt"}}这里是templ33.txt模板内容啊`)
	check_err_template(err)
	T_clone_lookup = T_clone.Lookup("templ33.txt")
	fmt.Printf("===T_clone.Lookup(\"templ33.txt\"):%+v\n", T_clone_lookup)
	if T_clone_lookup != nil {
		fmt.Printf("===T_clone.Lookup(\"templ33.txt\"):%+v\n", T_clone_lookup.Root.String())
	}

	//输出：
	//	-----下面我们继续讲解关于template对象下的方法和属性：----
	//	===t.Lookup("letter"):&{name:letter Tree:0xc0000ba000 common:0xc0000480c0 leftDelim: rightDelim:}
	//	===t.Lookup("letter"):&{name:templ1.txt Tree:0xc0000ba900 common:0xc0000480c0 leftDelim: rightDelim:}
	//	===t.Lookup("letter"):&{name:templ2.txt Tree:0xc000128300 common:0xc0000480c0 leftDelim: rightDelim:}
	//
	//	===T_clone.Lookup("letter"):&{name:letter Tree:0xc0000ba000 common:0xc000049680 leftDelim: rightDelim:}
	//	===T_clone.Lookup("letter"):
	//	Dear {{.Name}},
	//	{{if .Attended}}
	//	It was a pleasure to see you at the wedding.{{else}}
	//	It is a shame you couldn't make it to the wedding.{{end}}
	//	{{with .Gift}}Thank you for the lovely {{.}}.
	//	{{end}}
	//	Best wishes,
	//	Josie
	//
	//	===T_clone.Lookup("templ1.txt"):&{name:templ1.txt Tree:0xc0000ba900 common:0xc000049680 leftDelim: rightDelim:}
	//	===T_clone.Lookup("templ1.txt"):亲爱的：{{.Name}}
	//	{{if .Attended}}
	//	It was 这是我自定义的模板templ111,a pleasure to see you at the wedding.{{else}}
	//	It is a shame you couldn't make it to the wedding.{{end}}
	//	{{with .Gift}}Thank you for the lovely {{.}}.
	//	{{end}}
	//
	//	===T_clone.Lookup("templ2.txt"):&{name:templ2.txt Tree:0xc000128300 common:0xc000049680 leftDelim: rightDelim:}
	//	===T_clone.Lookup("templ2.txt"):亲爱的：{{.Name}}
	//	{{if .Attended}}
	//	It was 这是我自定义的模板templ222,a pleasure to see you at the wedding.{{else}}
	//	It is a shame you couldn't make it to the wedding.{{end}}
	//	{{with .Gift}}Thank you for the lovely {{.}}.
	//	{{end}}
	//
	//	===T_clone.Lookup("templ3.txt"):<nil>
	//	===T_clone.Lookup("templ3.txt"):&{name:templ3.txt Tree:0xc00009c300 common:0xc00003e340 leftDelim: rightDelim:}
	//	===T_clone.Lookup("templ3.txt"):这里是templ3.txt模板内容啊
	//	template: letter:1: unexpected EOF
	//	===T_clone.Lookup("templ33.txt"):<nil>
	//从上面可以知道，克隆是完完整整的深度copy出来一个新的对象

	fmt.Println("---下面讲解一个非常重要的知识点---")


	//下面我们对模板中的函数进行介绍，下面是go中文网中的翻译：
	//	Functions
	//	执行模板时，函数从两个函数字典中查找：首先是模板函数字典，然后是全局函数字典。一般不在模板内定义函数，而是使用Funcs方法添加函数到模板里。
	//
	//	预定义的全局函数如下：
	//
	//	and
	//	函数返回它的第一个empty参数或者最后一个参数；
	//	就是说"and x y"等价于"if x then y else x"；所有参数都会执行；
	//	or
	//	返回第一个非empty参数或者最后一个参数；
	//	亦即"or x y"等价于"if x then x else y"；所有参数都会执行；
	//	not
	//	返回它的单个参数的布尔值的否定
	//	len
	//	返回它的参数的整数类型长度
	//	index
	//	执行结果为第一个参数以剩下的参数为索引/键指向的值；
	//	如"index x 1 2 3"返回x[1][2][3]的值；每个被索引的主体必须是数组、切片或者字典。
	//	print
	//	即fmt.Sprint
	//	printf
	//	即fmt.Sprintf
	//	println
	//	即fmt.Sprintln
	//	html
	//	返回其参数文本表示的HTML逸码等价表示。
	//	urlquery
	//	返回其参数文本表示的可嵌入URL查询的逸码等价表示。
	//	js
	//	返回其参数文本表示的JavaScript逸码等价表示。
	//	call
	//	执行结果是调用第一个参数的返回值，该参数必须是函数类型，其余参数作为调用该函数的参数；
	//	如"call .X.Y 1 2"等价于go语言里的dot.X.Y(1, 2)；
	//	其中Y是函数类型的字段或者字典的值，或者其他类似情况；
	//	call的第一个参数的执行结果必须是函数类型的值（和预定义函数如print明显不同）；
	//	该函数类型值必须有1到2个返回值，如果有2个则后一个必须是error接口类型；
	//	如果有2个返回值的方法返回的error非nil，模板执行会中断并返回给调用模板执行者该错误；
	//	布尔函数会将任何类型的零值视为假，其余视为真。
	//
	//	下面是定义为函数的二元比较运算的集合：
	//
	//	eq      如果arg1 == arg2则返回真
	//	ne      如果arg1 != arg2则返回真
	//	lt      如果arg1 < arg2则返回真
	//	le      如果arg1 <= arg2则返回真
	//	gt      如果arg1 > arg2则返回真
	//	ge      如果arg1 >= arg2则返回真
	//	为了简化多参数相等检测，eq（只有eq）可以接受2个或更多个参数，它会将第一个参数和其余参数依次比较，返回下式的结果：
	//
	//	arg1==arg2 || arg1==arg3 || arg1==arg4 ...
	//	（和go的||不一样，不做惰性运算，所有参数都会执行）
	//
	//	比较函数只适用于基本类型（或重定义的基本类型，如"type Celsius float32"）。它们实现了go语言规则的值的比较，但具体的类型和大小会忽略掉，
	//	因此任意类型有符号整数值都可以互相比较；任意类型无符号整数值都可以互相比较；等等。但是，整数和浮点数不能互相比较。

	//下面的2个函数是参考内置的and()函数写的，几乎完全一样，当然也有不同的地方

	// indirectInterface returns the concrete value in an interface value,
	// or else the zero reflect.Value.
	// That is, if v represents the interface value x, the result is the same as reflect.ValueOf(x):
	// the fact that x was an interface value is forgotten.
	// indirectInterface返回接口值中的具体值，否则为零reflect.Value。
	//也就是说，如果v表示接口值x，则结果与reflect.ValueOf（x）相同：
	//忘记x是接口值的事实。
	indirectInterface := func(v reflect.Value) reflect.Value {
		if v.Kind() != reflect.Interface {
			return v
		}
		if v.IsNil() {
			return reflect.Value{}
		}
		return v.Elem()
	}

	// and computes the Boolean AND of its arguments, returning
	// the first false argument it encounters, or the last argument.
	//并计算其参数的布尔AND，返回其遇到的第一个错误参数或最后一个参数。
	//这并不是与运算，内置的or也不是或运算，而是返回第一个遇到的0值的索引
	//总的来说，就是返回第一个遇到的零值参数（可为最后一个参数）或者返回最后一个非零值参数
	and := func(arg0 reflect.Value, args ...reflect.Value) reflect.Value {
		//从下面可义看的出，arg0的类型是interface类型，他的子元素.Elem()返回的是接口的值
		fmt.Println("====",arg0,arg0.Kind(),arg0.Elem())
		fmt.Println(template.IsTrue(0))//int类型的零值确实是0
		//0确实不为Value的零值，Value是一个结构体，他是值类型，struct没有零值,即使是struct{}也不会是零值,他的值总是有效的
		//零值的Value是无效值，也就是不存在这个值，所以Value结构体的零值是没有的！注意了
		fmt.Println(template.IsTrue(reflect.ValueOf(0)))
		fmt.Println(template.IsTrue(reflect.Value{}))
		//下面的这句跟内置的and()函数不同，注意了，加不加indirectInterface对arg0进行封装区别可大了，见上面！
		if truth, _ := template.IsTrue(indirectInterface(arg0)); !truth {
			//如果输入的arg0参数是类型的零值的话，则直接返回他
			return arg0
		}
		//arg0非类型的零值
		for i := range args {
			//将arg0赋值给args列表中遍历的索引值
			arg0 = args[i]
			//判断下一个输入的参数是否是零值
			if truth, _ := template.IsTrue(indirectInterface(arg0)); !truth {
				//如果下一个参数输入的args的第i个参数是类型的零值的话，则break然后直接返回当前args中的索引i的值（一定是类型零值）
				break
			}
			//如果2个都不是零值，则迭代
		}
		//如果来到这里的原因不是break,而是完整遍历完后，则此时返回的是最后一个输入参数的值
		//否则，返回的是最近遍历到的一个类型的零值
		return arg0
	}
	// Funcs adds the elements of the argument map to the template's function map.
	// It must be called before the template is parsed.
	// It panics if a value in the map is not a function with appropriate return
	// type or if the name cannot be used syntactically as a function in a template.
	// It is legal to overwrite elements of the map. The return value is the template,
	// so calls can be chained.
	//Funcs方法向模板t的函数字典里加入参数funcMap内的键值对。如果funcMap某个键值对的值不是函数
	//类型或者返回值不符合要求会panic。但是，可以对t函数列表的成员进行重写。方法返回t以便进行链式调用。
	//这个方法是加锁的，底层实现：
	//		addValueFuncs(t.execFuncs, funcMap)
	//		addFuncs(t.parseFuncs, funcMap)

	funcMap1 := template.FuncMap{
		// The name "title" is what the function will be called in the template text.
		//名称“ title”是在模板文本中将调用的函数。
		"title": strings.Title,//这个函数是go库中内置的
		"and1": and,//这个是我们参考内置的FuncMap 的and函数来写的，功能完全一样
	}

	tt := template.New("templ_func")

	//必须在parse之前进行绑定自定义的函数，否则无效
	T_funcs := tt.Funcs(funcMap1)

	const templateText = `
Input: {{printf "%q" .word}}
Output 0: {{title .word}}
Output 1: {{title .word | printf "%q"}}
Output 2: {{printf "%q" .word | title}}
Output 3: {{and1 .OperationArg1 .OperationArg0}}
Output 4: {{and1 .OperationArg0 .OperationArg1}}
Output 4.1: {{and1 .OperationArg0 .OperationArg1 .OperationArg1}}
Output 4.2: {{and1 .OperationArg0 .OperationArg1 .OperationArg0}}
Output 4.3: {{and1 .OperationArg0 .OperationArg1 .OperationArg1 .OperationArg1}}
Output 4.4: {{and1 .OperationArg1 .OperationArg1 .OperationArg0 .OperationArg1}}
Output 4.5: {{and1 .OperationArg1 .OperationArg1 .OperationArg1 .OperationArg0}}
Output 4.7: {{and1 .OperationArg1 .OperationArg2}}
Output 4.8: {{and1 .OperationArg2 .OperationArg1}}
Output 4.9: {{and1 .OperationArg1 .OperationArg0 .OperationArg2}}
Output 5: {{.0}}
Output 6: {{.1}}
Output 7: {{index .ls 0}}
Output 8: {{index .ls 1}}
Output 9: {{len .ls}}
Output 10: {{index . "word"}}
Output 11: {{index . "ls"}}
Output 12: {{html .htmlexp}}
Output 13: {{js .jsexp}}
Output 14: {{urlquery .urlexp}}
Output 15: {{eq .eqArg1 .eqArg2}}
Output 16: {{eq .eqArg1 .eqArg0}}
Output 17: {{lt .eqArg1 .eqArg0}}
Output 18: {{gt .eqArg1 .eqArg0}}
Output 19: {{ge .eqArg1 .eqArg0}}
Output 20: {{with $x := "output" | printf "%q"}}{{$x}}{{end}}
Output 21: {{with $x := "output" | printf "%q"}}{{$x}}{{111}}{{end}}
Output 22: {{with $x := "output" | printf "%q"}}{{111}}{{end}}
Output 23: {{with $x := "output"}}{{printf "%q" $x}}{{end}}
Output 24: {{with $x := "output"}}{{printf "%q%d" $x 111}}{{end}}
Output 25: {{with $x := "output"}}{{printf "%q%q" $x 111}}{{end}}
Output 26: {{with $x := "output"}}{{$x | printf "%q"}}{{end}}
Output 27: {{with $x := "output"}}{{$x | printf "%d%q" 111}}{{end}}
Output 28: {{with $x := "output"}}{{$x | printf "%d%q" 111 | title}}{{end}}
Output 29: {{"output" | printf "%s" }}
Output 30: {{"output" | printf "%s" | printf "%q"}}
Output 31: {{.eqbool0 | printf "%v" }}
Output 32: {{.eqbool1 | printf "%v" }}
Output 33: {{eq .eqstr1 .eqstr2}}
Output 34: {{eq .eqstr2 .eqstr3}}
Output 35: {{eq .eqbool0 .eqbool1}}
Output 36: {{eq .eqbool1 .eqbool2}}
Output 37: {{call .callDividefunc1 5 2}}
Output 38: {{call .callDividefunc2 5 2}}

Output 40: {{call .callDividefunc3 16 4 2}}
Output 40.5: {{call .callAddfunc1 16 4 }}
Output 40.6: {{call .callAddfunc1 16 4 2 3}}



Output 45: {{range $index, $element := .slice}}索引为：{{$index}}，值为：{{$element}}{{println}}{{end}}
Output 46: {{range $index, $element := .array}}索引为：{{$index}}，值为：{{$element}}{{println}}{{end}}
Output 47: {{range $index, $element := .map}}索引为：{{$index}}，值为：{{$element}}{{println}}{{end}}
Output 48: {{range $index, $element := .myslice}}索引为：{{$index}}，值为：{{$element}}{{println}}{{end}}
Output 49: {{range $element:= .myslice}}值为：{{$element}}{{println}}{{end}}
Output 50: {{range $index, $ := .myslice}}索引为：{{$index}}，值为：{{$}}{{println}}{{end}}
Output 51: {{range $index, $_:= .myslice}}索引为：{{$index}}，值为：{{$}}{{println}}{{end}}
Output 52: {{range $index, $_:= .myslice}}索引为：{{$index}}，值为：{{$_}}{{println}}{{end}}
Output 53: {{print $.ls}}
Output 53.5: {{print $}}
Output 53.6: {{print .}}

Output 54: {{print .ls}}
Output 55: {{with $myint:=88}}xx{{$myint}}yy{{end}}
Output 56: {{with 88}}xx{{end}}
Output 57: {{with 0}}xx{{end}}
Output 58: {{with false}}xx{{end}}
Output 59: {{with 'a'}}xx{{end}}
Output 60: {{with "abc"}}xx{{end}}
Output 61: {{with -9}}xx{{end}}
Output 62: {{with 3.14}}xx{{end}}
Output 63: {{with 0b00001000}}xx{{end}}
Output 64: {{with 0xfffd}}xx{{end}}
Output 65: {{with 0O144}}xx{{end}}
Output 66: {{with 3.14e9}}xx{{end}}

Output 67: {{with 3.14e9}}xx{{else}}yy{{end}}
Output 68: {{with 0}}xx{{else}}yy{{end}}

Output 69: {{with 0}}dot为：{{.}}{{else}}else:dot为：{{.}}{{println}}$为：{{$}}{{end}}
Output 70: {{with 111}}dot为：{{.}}{{println}}$为：{{$}}{{end}}
Output 71: {{println .multiStr}}
Output 72: {{.slice1}}
Output 73: {{print .slice1}}
Output 74: {{with $sli := .slice1}}{{slice $sli 1 2 3}}{{end}}
Output 75: {{$sli := .slice1}}
Output 76: {{$sli := .slice1}}{{$sli}}
Output 77: {{$sli := .slice1}}{{.}}
Output 78: {{$sli := .slice1}}{{$}}
Output 79: {{$sli := .slice1}}{{slice $sli 1 2 3}}
{{/* a comment（这是单行注释，不会被解析到模板中去） */}}
{{/* 
a comment（这是多行注释，不会被解析到模板中去） 
a comment（这是多行注释，不会被解析到模板中去）
a comment（这是多行注释，不会被解析到模板中去）
a comment（这是多行注释，不会被解析到模板中去）
*/}}


`

	//Output 75: {{slice []string{"a","b","c","d","e","f"} 1 2}}
	//Output 76: {{slice []string{"a","b","c","d","e","f"} 1}}
	//因为Output 39，41，42会抛出异常，所以我写在下面
//Output 39: {{call .callDividefunc2 5 0}}
//Output 41: {{call .callDividefunc3 16 4 0}}
//Output 42: {{call .callDividefunc3 16 0 4}}
//Output 43: {{call .callDividefunc4 16 4}}此函数会抛出异常function called with 3 args; should be 1 or 2
//Output 44: {{call .callDividefunc5 16 4}}此函数会抛出异常function called with 3 args; should be 1 or 2
//Output 40.7: {{call .callAddfunc1 16}}此函数会抛出异常error calling call: wrong number of args: got 1 want at least 2
//$和.可以同时存在，但是不能仅仅存在.和$这两个元素，而没有其他元素
//Output 53.7: {{print $.}}
//Output 53.8: {{print .$}}

//下面均会产生报错，无法识别%和[]字符或者byte类型
//Output 67: {{with 3.14%%}}xx{{end}}
//Output 67: {{with byte(97)}}xx{{end}}
//Output 67: {{with []byte{'a','b'}}}xx{{end}}

//range函数仅适用于map,slice和array,不适用于string
//Output 72: {{range $k, $v:= .multiStr}}索引为：{{$k}}，值为：{{$v}}{{end}}

//从下面的2句可以看得出，无论有没有print函数，事实上并不是默认是print函数，而是我们下面有Execute(os.Stdout,m )
//Output 72: {{.slice1}}
//Output 73: {{print .slice1}}

//slice函数只能作用于slice类型，不能作用于interface{}类型（我们的m拿出来的值就是interface{}类型，所以我们需要对这个类型
// 进行类型断言成切片，数组或者字符串类型或者他们的拓展类型才可以，当然也可以Output 74或者Output 79那样），从下面的map[string]interface{}
// 中取出来的值都是interface{}类型的。总的来说就是要先接收这个interface{}类型，其实接收的同时会自动断言成相对应的类型的！所以我们没必要自己手动写一个函数
// 来进行断言并且调用！而且我们也没法实现，因为{{call func .slice1 | slice 1}},管道符的前一个命令得到结果 作为后一个子命令的最后一个参数来传递的，但是我们
// 必须作为slice的第一个参数才可以，所以这是go的不足之处。
// 下面的都会报错
//Output 77: {{slice .slice1 1}}
//Output 77.5: {{call slice1fun .slice1 | slice 1}}

	//上面的.代表单个的m类型数据直接传递进去,{{.1}}中的.1代表的是浮点数0.1而不是map中的第一个索引值,
	//{{index .ls 1}}千万不能写成{{.ls 1}}或者{{.ls.1}}的形式，不然均会报错！
	//{{index . "word"}}不能去除word的引号
	//{{$x}}{{111}}千万不能合并2个变量为{{$x 111}}或者{{$x,111}}的形式
	//{{$x | printf "%d%q" 111}}的第一个参数是111，第二个参数才是$x
	//{{$x | printf "%d%q" 111 | title}}title只接受一个参数，不要再title后面再加上任意一个参数，比如"abc",这样就是2个参数了
	//Output 45: {{range $index, $element := .slice}}索引为：{{$index}}，值为：{{$element}}{{println}}{{end}}中的{{println}}不能换成\n,因为上面不是双引号括着，而是单引号

	//模板开始执行时，$会设置为传递给Execute方法的参数，就是说，dot的初始值。(golang中文网)
	//模板执行时会遍历结构并将指针表示为'.'（称之为"dot"）指向运行过程中数据结构的当前位置的值。(golang中文网)
	//{{$}}和dot对象都等于刚开始传递进来的对象，当然我们也可以对这个变量取属性或者调用方法，或者对$重新赋值给另外一个对象。$.ls==.ls
	//$和.刚开始时候都会等于传递进来的data对象,他们都可以被重新赋值给其他对象，with语句中是把.赋值给为true的对象，但为false时候不会赋值，with语句
	//不会影响$的值，而$在range或者其他任何语句中都可以赋值给其他的任何对象，$和.都是任意的接口类型而非具体的某一个类型

	//eq.lt等等比较运算符只能比较基本数据类型的值，不能比较像[]byte类型的值,能作为比较函数的类型如下：
	//		reflect.Bool
	//		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64
	//		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr
	//		reflect.Float32, reflect.Float64
	//		reflect.Complex64, reflect.Complex128
	//		reflect.String
	//		其中，有符号整数类型可以无符号整形类型可以进行跨类型相互比较（事实上底层会进行强制转换），除此之外的其他类型的比较都必须类型相同
	//		为了简化多参数相等检测，eq（只有eq）可以接受2个或更多个参数，它会将第一个参数和其余参数依次比较，返回下式的结果：
	//		arg1==arg2 || arg1==arg3 || arg1==arg4 ...（和go的||不一样，不做惰性运算，所有参数都会执行）

	//事实证明，上面的Output 1和Output 2中的2个函数title和printf的顺序不是死的，可以随意更换，对结果无影响

	//template包下的call()文档（其实调用的reflect.Value.Call()方法，包括其他的任何eq等等函数都会调用这个函数来执行函数的！）
	// call返回将第一个参数作为函数求值的结果。
	//该函数必须返回1个结果或2个结果，其中第二个是错误。

	//关于reflect.Value.Call()方法：
	// Call调用带有输入参数的函数v。
	//例如，如果len（in）== 3，则v.Call（in）表示Go调用v（in [0]，in [1]，in [2]）。
	//如果v的Kind不是Func，则呼叫恐慌。
	//将输出结果作为值返回。
	//和Go一样，每个输入参数必须可分配给函数相应输入参数的类型。
	//如果v是可变参数函数，则Call会自己创建可变参数slice参数，并复制相应的值。



	T_funcs_parsed, err := T_funcs.Parse(templateText)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	type myslice []byte

	//事实上我们的map[string]interface{}的值不一定要设置interface{}类型，可以是任意类型，因为下面的Execute()
	//的第二个参数是任意类型
	m:=map[string]interface{}{"word":"the go programming language",
		"OperationArg0":10,
		"OperationArg1":55,
		"OperationArg2":0,
		"ls":[]byte{'a','b','c'},
		"htmlexp":"<html><body><h1>我的第一个标题</h1><p>我的第一个段落。</p></body></html>",
		"jsexp":`<script src="http://www.w3school.com.cn/a\\b\demo/myScript.js"></script>`,
		"urlexp":`http://www.w3school.com.cn/a\\b\demo/myScript.js`,
		"eqArg0":0,
		"eqArg1":1,
		"eqArg2":1,
		"callexp":"title",
		"eqbool0":false,
		"eqbool1":true,
		"eqbool2":true,
		"eqstr1":"anko",
		"eqstr2":"anko",
		"eqstr3":"anko33",
		"callDividefunc1": func(i,j int) int{return i/j},
		"callDividefunc2": func(i,j int) (int,error){
			if j==0{
				return 0,errors.New("除数不能为零")//只要你在这里返回的第二个参数的值不是nil的话，go包都会panic,不用你手动panic
			}
			return i/j,nil
		},
		"callDividefunc3": func(i,j,z int) (int,error){
			if j==0||z==0{
				return 0,errors.New("第二第三个参数作为除数不能为零")
			}
			return i/j/z,nil
		},

		//下面的函数callDividefunc4和callDividefunc5会抛出异常function called with 3 args; should be 1 or 2
		//自定义的函数自能返回最多2个值，如果返回2个值，第二个值一定为error接口类型才可以
		"callDividefunc4": func(i,j int) (int,int,error){
			return i,j,nil
		},
		"callDividefunc5": func(i,j int) (int,int){
			return i,j
		},
		//当然也可以设置可变参数的函数
		"callAddfunc1": func(i,j int,z ...int) int{
			sum:=i+j
			for _, v := range z {
				sum+=v
			}
			return sum

		},
		"slice":[]byte{'a','b','c'},
		"array":[3]byte{'a','b','c'},
		"map":map[string]byte{"key1":'a',"key2":'b',"key3":'c'},
		"myslice":myslice{'x','y','z'},
		"multiStr":`当美容院向你推出一种价格不菲的新服务——蜗牛爬脸美容，
你敢试吗？蜗牛爬脸美容，意在让肌肤吸取蜗牛粘液，
从而达到美容的效果。在你大胆一试之前，不妨先了解一下蜗牛吧！`,
		//
		//字符串，slice和array类型的切片操作：
		// slice返回将其第一个参数与其余参数相切片的结果。 因此，按照Go语法，“ slice x 1 2”是x [1：2]，而“ slice x”是x [：]，
		// “ slice x 1”是x [1：]和“ slice x 1 2 3” “是x [1：2：3]。 第一个参数必须是字符串，切片或数组。
		//对于给3个索引时候需要满足以下的条件：
		// Slice3是slice操作的3索引形式：它返回v [i：j：k]。
		//如果v的Kind不是Array或Slice，或者v是不可寻址的数组，或者索引超出范围，它就会发生混乱。
		"slice1":[]string{"a","b","c","d","e","f"},
		"slice1fun": func(sli_in interface{})[]string {
			//我们就不检查异常了
			slice := sli_in.([]string)
			return slice
		},

	}
	// Run the template to verify the output.
	//注意这里接收的第二个参数的类型是任意的
	err = T_funcs_parsed.Execute(os.Stdout,m )
	if err != nil {
		log.Fatalf("execution: %s", err)
	}


	//输出：
	//	dir: main3\template896667435
	//	pattern: main3\template896667435\*.tmpl
	//	T0 invokes T1: (T1 invokes T2: (This is T2))
	//	-----下面我们继续讲解关于template对象下的方法和属性：----
	//	===t.Lookup("letter"):&{name:letter Tree:0xc00009e000 common:0xc00003e100 leftDelim: rightDelim:}
	//	===t.Lookup("letter"):&{name:templ1.txt Tree:0xc0000e8800 common:0xc00003e100 leftDelim: rightDelim:}
	//	===t.Lookup("letter"):&{name:templ2.txt Tree:0xc0000e8900 common:0xc00003e100 leftDelim: rightDelim:}
	//
	//	===T_clone.Lookup("letter"):&{name:letter Tree:0xc00009e000 common:0xc00003e3c0 leftDelim: rightDelim:}
	//	===T_clone.Lookup("letter"):
	//	Dear {{.Name}},
	//	{{if .Attended}}
	//	It was a pleasure to see you at the wedding.{{else}}
	//	It is a shame you couldn't make it to the wedding.{{end}}
	//	{{with .Gift}}Thank you for the lovely {{.}}.
	//	{{end}}
	//	Best wishes,
	//		Josie
	//
	//	===T_clone.Lookup("templ1.txt"):&{name:templ1.txt Tree:0xc0000e8800 common:0xc00003e3c0 leftDelim: rightDelim:}
	//	===T_clone.Lookup("templ1.txt"):亲爱的：{{.Name}}
	//	{{if .Attended}}
	//	It was 这是我自定义的模板templ111,a pleasure to see you at the wedding.{{else}}
	//	It is a shame you couldn't make it to the wedding.{{end}}
	//	{{with .Gift}}Thank you for the lovely {{.}}.
	//	{{end}}
	//
	//	===T_clone.Lookup("templ2.txt"):&{name:templ2.txt Tree:0xc0000e8900 common:0xc00003e3c0 leftDelim: rightDelim:}
	//	===T_clone.Lookup("templ2.txt"):亲爱的：{{.Name}}
	//	{{if .Attended}}
	//	It was 这是我自定义的模板templ222,a pleasure to see you at the wedding.{{else}}
	//	It is a shame you couldn't make it to the wedding.{{end}}
	//	{{with .Gift}}Thank you for the lovely {{.}}.
	//	{{end}}
	//
	//	===T_clone.Lookup("templ3.txt"):<nil>
	//	---下面讲解一个非常重要的知识点---
	//
	//		Input: "the go programming language"
	//	Output 0: The Go Programming Language
	//	Output 1: "The Go Programming Language"
	//	Output 2: "The Go Programming Language"
	//	Output 3: ==== 55 interface 55
	//	false true
	//	true true
	//	true true
	//	10
	//	Output 4: ==== 10 interface 10
	//	false true
	//	true true
	//	true true
	//	55
	//	Output 4.1: ==== 10 interface 10
	//	false true
	//	true true
	//	true true
	//	55
	//	Output 4.2: ==== 10 interface 10
	//	false true
	//	true true
	//	true true
	//	10
	//	Output 4.3: ==== 10 interface 10
	//	false true
	//	true true
	//	true true
	//	55
	//	Output 4.4: ==== 55 interface 55
	//	false true
	//	true true
	//	true true
	//	55
	//	Output 4.5: ==== 55 interface 55
	//	false true
	//	true true
	//	true true
	//	10
	//	Output 4.7: ==== 55 interface 55
	//	false true
	//	true true
	//	true true
	//	0
	//	Output 4.8: ==== 0 interface 0
	//	false true
	//	true true
	//	true true
	//	55
	//	Output 4.9: ==== 55 interface 55
	//	false true
	//	true true
	//	true true
	//	0
	//	Output 5: 0
	//	Output 6: 0.1
	//	Output 7: 97
	//	Output 8: 98
	//	Output 9: 3
	//	Output 10: the go programming language
	//	Output 11: [97 98 99]
	//	Output 12: &lt;html&gt;&lt;body&gt;&lt;h1&gt;我的第一个标题&lt;/h1&gt;&lt;p&gt;我的第一个段落。&lt;/p&gt;&lt;/body&gt;&lt;/html&gt;
	//	Output 13: \x3Cscript src=\"http://www.w3school.com.cn/a\\\\b\\demo/myScript.js\"\x3E\x3C/script\x3E
	//	Output 14: http%3A%2F%2Fwww.w3school.com.cn%2Fa%5C%5Cb%5Cdemo%2FmyScript.js
	//	Output 15: true
	//	Output 16: false
	//	Output 17: false
	//	Output 18: true
	//	Output 19: true
	//	Output 20: "output"
	//	Output 21: "output"111
	//	Output 22: 111
	//	Output 23: "output"
	//	Output 24: "output"111
	//	Output 25: "output"'o'
	//	Output 26: "output"
	//	Output 27: 111"output"
	//	Output 28: 111"Output"
	//	Output 29: output
	//	Output 30: "output"
	//	Output 31: false
	//	Output 32: true
	//	Output 33: true
	//	Output 34: false
	//	Output 35: false
	//	Output 36: true
	//	Output 37: 2
	//	Output 38: 2
	//
	//	Output 40: 2
	//	Output 40.5: 20
	//	Output 40.6: 25
	//
	//
	//
	//	Output 45: 索引为：0，值为：97
	//	索引为：1，值为：98
	//	索引为：2，值为：99
	//
	//	Output 46: 索引为：0，值为：97
	//	索引为：1，值为：98
	//	索引为：2，值为：99
	//
	//	Output 47: 索引为：key1，值为：97
	//	索引为：key2，值为：98
	//	索引为：key3，值为：99
	//
	//	Output 48: 索引为：0，值为：120
	//	索引为：1，值为：121
	//	索引为：2，值为：122
	//
	//	Output 49: 值为：120
	//	值为：121
	//	值为：122
	//
	//	Output 50: 索引为：0，值为：120
	//	索引为：1，值为：121
	//	索引为：2，值为：122
	//
	//	Output 51: 索引为：0，值为：map[OperationArg0:10 OperationArg1:55 OperationArg2:0 array:[97 98 99] callAddfunc1:0x516760 callDividefunc1:0x516530 callDividefunc2:0x516580 callDividefunc3:0x516640 callDividefunc4:0x516720 callDividefunc5:0x516740 callexp:title eqArg0:0 eqArg1:1 eqArg2:1 eqbool0:false eqbool1:true eqbool2:true eqstr1:anko eqstr2:anko eqstr3:anko33 htmlexp:<html><body><h1>我的第一个标题</h1><p>我的第一个段落。</p></body></html> jsexp:<script src="http://www.w3school.com.cn/a\\b\demo/myScript.js"></script> ls:[97 98 99] map:map[key1:97 key2:98 key3:99] multiStr:当美容院向你推出一种价格不菲的新服务——蜗牛爬脸美容，
	//	你敢试吗？蜗牛爬脸美容，意在让肌肤吸取蜗牛粘液，
	//	从而达到美容的效果。在你大胆一试之前，不妨先了解一下蜗牛吧！ myslice:[120 121 122] slice:[97 98 99] slice1:[a b c d e f] slice1fun:0x516790 urlexp:http://www.w3school.com.cn/a\\b\demo/myScript.js word:the go programming language]
	//	索引为：1，值为：map[OperationArg0:10 OperationArg1:55 OperationArg2:0 array:[97 98 99] callAddfunc1:0x516760 callDividefunc1:0x516530 callDividefunc2:0x516580 callDividefunc3:0x516640 callDividefunc4:0x516720 callDividefunc5:0x516740 callexp:title eqArg0:0 eqArg1:1 eqArg2:1 eqbool0:false eqbool1:true eqbool2:true eqstr1:anko eqstr2:anko eqstr3:anko33 htmlexp:<html><body><h1>我的第一个标题</h1><p>我的第一个段落。</p></body></html> jsexp:<script src="http://www.w3school.com.cn/a\\b\demo/myScript.js"></script> ls:[97 98 99] map:map[key1:97 key2:98 key3:99] multiStr:当美容院向你推出一种价格不菲的新服务——蜗牛爬脸美容，
	//	你敢试吗？蜗牛爬脸美容，意在让肌肤吸取蜗牛粘液，
	//	从而达到美容的效果。在你大胆一试之前，不妨先了解一下蜗牛吧！ myslice:[120 121 122] slice:[97 98 99] slice1:[a b c d e f] slice1fun:0x516790 urlexp:http://www.w3school.com.cn/a\\b\demo/myScript.js word:the go programming language]
	//	索引为：2，值为：map[OperationArg0:10 OperationArg1:55 OperationArg2:0 array:[97 98 99] callAddfunc1:0x516760 callDividefunc1:0x516530 callDividefunc2:0x516580 callDividefunc3:0x516640 callDividefunc4:0x516720 callDividefunc5:0x516740 callexp:title eqArg0:0 eqArg1:1 eqArg2:1 eqbool0:false eqbool1:true eqbool2:true eqstr1:anko eqstr2:anko eqstr3:anko33 htmlexp:<html><body><h1>我的第一个标题</h1><p>我的第一个段落。</p></body></html> jsexp:<script src="http://www.w3school.com.cn/a\\b\demo/myScript.js"></script> ls:[97 98 99] map:map[key1:97 key2:98 key3:99] multiStr:当美容院向你推出一种价格不菲的新服务——蜗牛爬脸美容，
	//	你敢试吗？蜗牛爬脸美容，意在让肌肤吸取蜗牛粘液，
	//	从而达到美容的效果。在你大胆一试之前，不妨先了解一下蜗牛吧！ myslice:[120 121 122] slice:[97 98 99] slice1:[a b c d e f] slice1fun:0x516790 urlexp:http://www.w3school.com.cn/a\\b\demo/myScript.js word:the go programming language]
	//
	//	Output 52: 索引为：0，值为：120
	//	索引为：1，值为：121
	//	索引为：2，值为：122
	//
	//	Output 53: [97 98 99]
	//	Output 53.5: map[OperationArg0:10 OperationArg1:55 OperationArg2:0 array:[97 98 99] callAddfunc1:0x516760 callDividefunc1:0x516530 callDividefunc2:0x516580 callDividefunc3:0x516640 callDividefunc4:0x516720 callDividefunc5:0x516740 callexp:title eqArg0:0 eqArg1:1 eqArg2:1 eqbool0:false eqbool1:true eqbool2:true eqstr1:anko eqstr2:anko eqstr3:anko33 htmlexp:<html><body><h1>我的第一个标题</h1><p>我的第一个段落。</p></body></html> jsexp:<script src="http://www.w3school.com.cn/a\\b\demo/myScript.js"></script> ls:[97 98 99] map:map[key1:97 key2:98 key3:99] multiStr:当美容院向你推出一种价格不菲的新服务——蜗牛爬脸美容，
	//	你敢试吗？蜗牛爬脸美容，意在让肌肤吸取蜗牛粘液，
	//	从而达到美容的效果。在你大胆一试之前，不妨先了解一下蜗牛吧！ myslice:[120 121 122] slice:[97 98 99] slice1:[a b c d e f] slice1fun:0x516790 urlexp:http://www.w3school.com.cn/a\\b\demo/myScript.js word:the go programming language]
	//	Output 53.6: map[OperationArg0:10 OperationArg1:55 OperationArg2:0 array:[97 98 99] callAddfunc1:0x516760 callDividefunc1:0x516530 callDividefunc2:0x516580 callDividefunc3:0x516640 callDividefunc4:0x516720 callDividefunc5:0x516740 callexp:title eqArg0:0 eqArg1:1 eqArg2:1 eqbool0:false eqbool1:true eqbool2:true eqstr1:anko eqstr2:anko eqstr3:anko33 htmlexp:<html><body><h1>我的第一个标题</h1><p>我的第一个段落。</p></body></html> jsexp:<script src="http://www.w3school.com.cn/a\\b\demo/myScript.js"></script> ls:[97 98 99] map:map[key1:97 key2:98 key3:99] multiStr:当美容院向你推出一种价格不菲的新服务——蜗牛爬脸美容，
	//	你敢试吗？蜗牛爬脸美容，意在让肌肤吸取蜗牛粘液，
	//	从而达到美容的效果。在你大胆一试之前，不妨先了解一下蜗牛吧！ myslice:[120 121 122] slice:[97 98 99] slice1:[a b c d e f] slice1fun:0x516790 urlexp:http://www.w3school.com.cn/a\\b\demo/myScript.js word:the go programming language]
	//
	//	Output 54: [97 98 99]
	//	Output 55: xx88yy
	//	Output 56: xx
	//	Output 57:
	//	Output 58:
	//	Output 59: xx
	//	Output 60: xx
	//	Output 61: xx
	//	Output 62: xx
	//	Output 63: xx
	//	Output 64: xx
	//	Output 65: xx
	//	Output 66: xx
	//
	//	Output 67: xx
	//	Output 68: yy
	//
	//	Output 69: else:dot为：map[OperationArg0:10 OperationArg1:55 OperationArg2:0 array:[97 98 99] callAddfunc1:0x516760 callDividefunc1:0x516530 callDividefunc2:0x516580 callDividefunc3:0x516640 callDividefunc4:0x516720 callDividefunc5:0x516740 callexp:title eqArg0:0 eqArg1:1 eqArg2:1 eqbool0:false eqbool1:true eqbool2:true eqstr1:anko eqstr2:anko eqstr3:anko33 htmlexp:<html><body><h1>我的第一个标题</h1><p>我的第一个段落。</p></body></html> jsexp:<script src="http://www.w3school.com.cn/a\\b\demo/myScript.js"></script> ls:[97 98 99] map:map[key1:97 key2:98 key3:99] multiStr:当美容院向你推出一种价格不菲的新服务——蜗牛爬脸美容，
	//	你敢试吗？蜗牛爬脸美容，意在让肌肤吸取蜗牛粘液，
	//	从而达到美容的效果。在你大胆一试之前，不妨先了解一下蜗牛吧！ myslice:[120 121 122] slice:[97 98 99] slice1:[a b c d e f] slice1fun:0x516790 urlexp:http://www.w3school.com.cn/a\\b\demo/myScript.js word:the go programming language]
	//	$为：map[OperationArg0:10 OperationArg1:55 OperationArg2:0 array:[97 98 99] callAddfunc1:0x516760 callDividefunc1:0x516530 callDividefunc2:0x516580 callDividefunc3:0x516640 callDividefunc4:0x516720 callDividefunc5:0x516740 callexp:title eqArg0:0 eqArg1:1 eqArg2:1 eqbool0:false eqbool1:true eqbool2:true eqstr1:anko eqstr2:anko eqstr3:anko33 htmlexp:<html><body><h1>我的第一个标题</h1><p>我的第一个段落。</p></body></html> jsexp:<script src="http://www.w3school.com.cn/a\\b\demo/myScript.js"></script> ls:[97 98 99] map:map[key1:97 key2:98 key3:99] multiStr:当美容院向你推出一种价格不菲的新服务——蜗牛爬脸美容，
	//	你敢试吗？蜗牛爬脸美容，意在让肌肤吸取蜗牛粘液，
	//	从而达到美容的效果。在你大胆一试之前，不妨先了解一下蜗牛吧！ myslice:[120 121 122] slice:[97 98 99] slice1:[a b c d e f] slice1fun:0x516790 urlexp:http://www.w3school.com.cn/a\\b\demo/myScript.js word:the go programming language]
	//	Output 70: dot为：111
	//	$为：map[OperationArg0:10 OperationArg1:55 OperationArg2:0 array:[97 98 99] callAddfunc1:0x516760 callDividefunc1:0x516530 callDividefunc2:0x516580 callDividefunc3:0x516640 callDividefunc4:0x516720 callDividefunc5:0x516740 callexp:title eqArg0:0 eqArg1:1 eqArg2:1 eqbool0:false eqbool1:true eqbool2:true eqstr1:anko eqstr2:anko eqstr3:anko33 htmlexp:<html><body><h1>我的第一个标题</h1><p>我的第一个段落。</p></body></html> jsexp:<script src="http://www.w3school.com.cn/a\\b\demo/myScript.js"></script> ls:[97 98 99] map:map[key1:97 key2:98 key3:99] multiStr:当美容院向你推出一种价格不菲的新服务——蜗牛爬脸美容，
	//	你敢试吗？蜗牛爬脸美容，意在让肌肤吸取蜗牛粘液，
	//	从而达到美容的效果。在你大胆一试之前，不妨先了解一下蜗牛吧！ myslice:[120 121 122] slice:[97 98 99] slice1:[a b c d e f] slice1fun:0x516790 urlexp:http://www.w3school.com.cn/a\\b\demo/myScript.js word:the go programming language]
	//	Output 71: 当美容院向你推出一种价格不菲的新服务——蜗牛爬脸美容，
	//	你敢试吗？蜗牛爬脸美容，意在让肌肤吸取蜗牛粘液，
	//	从而达到美容的效果。在你大胆一试之前，不妨先了解一下蜗牛吧！
	//
	//	Output 72: [a b c d e f]
	//	Output 73: [a b c d e f]
	//	Output 74: [b]
	//	Output 75:
	//	Output 76: [a b c d e f]
	//	Output 77: map[OperationArg0:10 OperationArg1:55 OperationArg2:0 array:[97 98 99] callAddfunc1:0x516760 callDividefunc1:0x516530 callDividefunc2:0x516580 callDividefunc3:0x516640 callDividefunc4:0x516720 callDividefunc5:0x516740 callexp:title eqArg0:0 eqArg1:1 eqArg2:1 eqbool0:false eqbool1:true eqbool2:true eqstr1:anko eqstr2:anko eqstr3:anko33 htmlexp:<html><body><h1>我的第一个标题</h1><p>我的第一个段落。</p></body></html> jsexp:<script src="http://www.w3school.com.cn/a\\b\demo/myScript.js"></script> ls:[97 98 99] map:map[key1:97 key2:98 key3:99] multiStr:当美容院向你推出一种价格不菲的新服务——蜗牛爬脸美容，
	//	你敢试吗？蜗牛爬脸美容，意在让肌肤吸取蜗牛粘液，
	//	从而达到美容的效果。在你大胆一试之前，不妨先了解一下蜗牛吧！ myslice:[120 121 122] slice:[97 98 99] slice1:[a b c d e f] slice1fun:0x516790 urlexp:http://www.w3school.com.cn/a\\b\demo/myScript.js word:the go programming language]
	//	Output 78: map[OperationArg0:10 OperationArg1:55 OperationArg2:0 array:[97 98 99] callAddfunc1:0x516760 callDividefunc1:0x516530 callDividefunc2:0x516580 callDividefunc3:0x516640 callDividefunc4:0x516720 callDividefunc5:0x516740 callexp:title eqArg0:0 eqArg1:1 eqArg2:1 eqbool0:false eqbool1:true eqbool2:true eqstr1:anko eqstr2:anko eqstr3:anko33 htmlexp:<html><body><h1>我的第一个标题</h1><p>我的第一个段落。</p></body></html> jsexp:<script src="http://www.w3school.com.cn/a\\b\demo/myScript.js"></script> ls:[97 98 99] map:map[key1:97 key2:98 key3:99] multiStr:当美容院向你推出一种价格不菲的新服务——蜗牛爬脸美容，
	//	你敢试吗？蜗牛爬脸美容，意在让肌肤吸取蜗牛粘液，
	//	从而达到美容的效果。在你大胆一试之前，不妨先了解一下蜗牛吧！ myslice:[120 121 122] slice:[97 98 99] slice1:[a b c d e f] slice1fun:0x516790 urlexp:http://www.w3school.com.cn/a\\b\demo/myScript.js word:the go programming language]
	//	Output 79: [b]


	fmt.Println("-------------继续探讨template对象之Templates()输出绑定的模板列表---------------")
	// Templates returns a slice of defined templates associated with t.
	//模板返回与t关联的已定义模板的切片。
	for key, value := range t.Templates() {
		fmt.Printf("key:value=%v:%v,模板字符串如下：\n%v\n",key,value.Name(),value.Root.String())
	}

	fmt.Println("-------------继续探讨template对象之Option()设置选项---------------")

	//// missingKeyAction defines how to respond to indexing a map with a key that is not present.
	//// missingKeyAction定义了如何响应 使用不存在的键索引map值时候的情况。
	//type missingKeyAction int
	//
	//const (
	//	mapInvalid   missingKeyAction = iota // Return an invalid reflect.Value.//返回无效的reflect.Value。
	//	mapZeroValue                         // Return the zero value for the map element.//返回map元素的零值。
	//	mapError                             // Error out//错误抛出
	//)
	//
	//type option struct {
	//	missingKey missingKeyAction
	//}


	// Option sets options for the template. Options are described by
	// strings, either a simple string or "key=value". There can be at
	// most one equals sign in an option string. If the option string
	// is unrecognized or otherwise invalid, Option panics.
	//
	// Known options:
	//
	// missingkey: Control the behavior during execution if a map is
	// indexed with a key that is not present in the map.
	//	"missingkey=default" or "missingkey=invalid"
	//		The default behavior: Do nothing and continue execution.
	//		If printed, the result of the index operation is the string
	//		"<no value>".
	//	"missingkey=zero"
	//		The operation returns the zero value for the map type's element.
	//	"missingkey=error"
	//		Execution stops immediately with an error.

	// Option设置模板的选项。 选项由字符串（简单字符串或"key=value"）描述。 选项字符串中最多可以有一个等号。
	// 如果选项字符串无法识别或无效，则Option()方法会抛出异常。
	//已知选项：
	// missingkey: 如果使用映射中不存在的键索引了映射，则在执行期间控制行为。
	//	"missingkey=default" or "missingkey=invalid"
	//		默认行为：不执行任何操作并继续执行。
	//		如果打印，则索引操作的结果是字符串"<no value>".
	//	"missingkey=zero"
	//		该操作将返回map类型元素的零值。
	//	"missingkey=error"
	//		执行立即停止并出现错误。
	//说白了就是在填充词中找不到模板中的键时候该如何处理！！！
	//底层实现1：
	//		switch elems[0] {
	//		case "missingkey":
	//			switch elems[1] {
	//			case "invalid", "default":
	//				t.option.missingKey = mapInvalid
	//				return
	//			case "zero":
	//				t.option.missingKey = mapZeroValue
	//				return
	//			case "error":
	//				t.option.missingKey = mapError
	//				return
	//			}
	//		}
	//再深度的底层实现请看template.exec.evalField()方法(不在这里做分析，因为没什么意义)

//语句1
	optionStr:=`
亲爱的：{{.Name}}
{{if .Attended}}
It was 这是我自定义的模板optionStr3333,a pleasure to see you at the wedding.{{else}}
It is a shame you couldn't make it to the wedding.{{end}}
{{with .Gift}}Thank you for the lovely {{.}}.
{{end}}
`

////语句2
////注意，下面跟上面的字符是有区别的额，多了一个{{.invalidKey}}，而这个东西是在填充词中没有的
//	optionStr:=`
//亲爱的：{{.Name}}
//{{if .Attended}}
//It was 这是我自定义的模板optionStr3333,{{.invalidKey}}a pleasure to see you at the wedding.{{else}}
//It is a shame you couldn't make it to the wedding.{{end}}
//{{with .Gift}}Thank you for the lovely {{.}}.
//{{end}}
//`
	T_new := template.New("anko")
	//T_new=T_new.Option("missingkey=default")//默认不设置的话也是这种情况
	//T_new=T_new.Option("missingkey=invalid")//默认不设置的话也是这种情况
	//T_new=T_new.Option("missingkey=zero")//默认不设置的话也是这种情况
	//T_new=T_new.Option("missingkey=mapError")//只有这种情况看出来了效果，似乎上面的3种情况设置了跟没设置并没有什么不同，也许是我还没探究出来
	//T_new=T_new.Option("missingkey=zero","missingkey=invalid")//还可以设置多个配置,但是后一个会把前一个的值给覆盖掉，不大明报为什么要这样写
	T_option := template.Must(T_new.Parse(optionStr))
	fmt.Println("与T_option关联的模板有：")
	for key, value := range T_option.Templates() {
		fmt.Printf("key:value=%v:%v,模板字符串为：\n%v\n",key,value,T_option.Root.String())
	}
	fmt.Println()
	//只要在Execute（）方法执行之前设置Option都是可以的，而不必是未解析之前，比如下面：
	//T_option=T_option.Option("missingkey=mapError")

	exeFun:= func(t *template.Template) {
		fmt.Println("----执行模板，准备输出填充后的完整模板----")
		err = T_option.Execute(os.Stdout, Recipient{"anko","洋娃娃",true})//Recipient对象在上面被定义了，非常上面
		check_err_template(err)
	}
	exeFun(T_option)
	//语句1输出：
	//
	//亲爱的：anko
	//
	//It was 这是我自定义的模板optionStr3333,template: letter:4:50: executing "letter"
	//at <.invalidKey>: can't evaluate field invalidKey in type main.Recipient
	//语句2输出：
	//
	//亲爱的：anko
	//
	//It was 这是我自定义的模板optionStr3333,template: letter:4:50: executing "letter"
	//at <.invalidKey>: can't evaluate field invalidKey in type main.Recipient

	//下面是设置了option为各个值时候的输出：
	//"missingkey=default"，"missingkey=invalid"和"missingkey=zero"的输出:
	//
	//亲爱的：anko
	//It was 这是我自定义的模板optionStr3333,template: letter:4:50: executing "letter"
	//at <.invalidKey>: can't evaluate field invalidKey in type main.Recipient
	//"missingkey=mapError"的输出:
	//panic: unrecognized option: missingkey=mapError
	//
	//goroutine 1 [running]:
	//text/template.(*Template).setOption(0xc00017a0c0, 0x567227, 0x13)
	//	C:/Go/src/text/template/option.go:73 +0x24b
	//text/template.(*Template).Option(0xc00017a0c0, 0xc00008d7e0, 0x1, 0x1, 0x0)
	//	C:/Go/src/text/template/option.go:45 +0x7c
	//main.main()
	//	C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:2110 +0x5ba4


	fmt.Println("-------------继续探讨template对象之Delims定界符---------------")
	//注意松开上面的语句1，语句2是错误的演示
	optionStr=`
亲爱的：{{.Name}}
{{if .Attended}}
It was 这是我^^^自定义的模板optionStr3333,a pleasure to see you at the wedding.{{else}}^
It is a shame you couldn't make it to the wedding.{{end}}^
{{with .Gift}}Thank you for the lovely {{.}}.^
{{end}}
`
	// Delims sets the action delimiters to the specified strings, to be used in
	// subsequent calls to Parse, ParseFiles, or ParseGlob. Nested template
	// definitions will inherit the settings. An empty delimiter stands for the
	// corresponding default: {{ or }}.
	// The return value is the template, so calls can be chained.
	// Delims将Action动作定界符设置为指定的字符串，以在后续对Parse，ParseFiles或ParseGlob的调用中使用。
	// 嵌套模板定义将继承设置。 空定界符代表相应的默认值：{{ or }}.
	//返回值是模板，因此可以链接调用。
	//目测输出模板中会忽略下面设置的*和^,但是不知道有什么用
	T_new1 := template.New("anko")
	T_new1=T_new1.Delims("*","^")
	T_option= template.Must(T_new1.Parse(optionStr))
	//似乎要在解析之前设置，
	//T_option=T_option.Delims("*","^")
	exeFun(T_option)
	//输出：
	//----执行模板，准备输出填充后的完整模板----
	//
	//亲爱的：{{.Name}}
	//{{if .Attended}}
	//It was 这是我^^^自定义的模板optionStr3333,a pleasure to see you at the wedding.{{else}}^
	//It is a shame you couldn't make it to the wedding.{{end}}^
	//{{with .Gift}}Thank you for the lovely {{.}}.^
	//{{end}}
	//目前不知道有什么作用，先搁置


	fmt.Println("-------------继续探讨template对象之DefinedTemplates()输入解析后的模板.command字段的值（command这就是你理解的模板容器）---------------")


	// DefinedTemplates returns a string listing the defined templates,
	// prefixed by the string "; defined templates are: ". If there are none,
	// it returns the empty string. For generating an error message here
	// and in html/template.
	// DefinedTemplates返回一个字符串，其中列出了已定义的模板，并以字符串"; defined templates are: "作为前缀。(注意这里的表达)
	// 如果没有，则返回空字符串。 用于在此处和html/template中生成错误消息。

	fmt.Println(t.DefinedTemplates())
	fmt.Println(T_new1.DefinedTemplates())
	fmt.Println(T_option.DefinedTemplates())
	T_new2 := template.New("anko112344")
	fmt.Println(T_new2.DefinedTemplates(),"==")
	T_new2_parsed := template.Must(T_new2.Parse(optionStr))
	fmt.Println(T_new2.DefinedTemplates(),"==")
	fmt.Println(T_new2_parsed.DefinedTemplates(),"==")
	//输出：
	//; defined templates are: "templ1.txt", "templ2.txt", "letter"
	//; defined templates are: "anko"
	//; defined templates are: "anko"
	//==
	//; defined templates are: "anko112344" ==
	//; defined templates are: "anko112344" ==
	//由此可知，模板必须解析后（或者叫绑定模板字符串）才有输出结果,但是在go中专业根据DefinedTemplates可知，go把它称为已定义的模板列表
	//当然我们理解的话就是已经解析的模板


	fmt.Println("-------------继续探讨template对象之AddParseTree()------------")

	// AddParseTree adds parse tree for template with given name and associates it with t.
	// If the template does not already exist, it will create a new one.
	// If the template does exist, it will be replaced.
	// AddParseTree为具有给定名称的模板添加解析树，并将其与t关联。
	//如果该模板尚不存在，它将创建一个新模板。注意了，他会新创建的tree的模板，所以我们必须需要接收返回值。以此来解析源模板上面的树节点和新添加的模板下面的树节点。
	//如果模板确实存在，它将被替换并且返回新的模板。

	const testTemplates = `{{define "one"}}one{{end}}{{define "two"}}two{{end}}`//define是定义一个模板，后面接模板名字
	TestMessageForExecuteEmpty:=func () {
		// Test a truly empty template.
		//这个模板名字作为不存在的模板，我们准备往这个根模板中加进去东西，空模板下的子模板可以不为空，一样可以解析
		tmpl := template.New("empty")
		//var b bytes.Buffer//用来接收输出的模板字节信息
		//没绑定任何模板直接运行会报错template: empty: "empty" is an incomplete or empty template
		err := tmpl.Execute(os.Stdout, 0)
		if err == nil {
			fmt.Println("expected initial error")
		}
		//报错才会往下执行
		got := err.Error()
		want := `template: empty: "empty" is an incomplete or empty template`
		if got != want {
			fmt.Errorf("expected error %s got %s", want, got)
		}

		// Add a non-empty template to check that the error is helpful.
		//添加非空模板以检查该错误是否有帮助。
		//注意这里的two必须是要解析的树节点的名字(树的名字会决定解析哪个字符串中定义模板)，同时他是返回的新模板的name,这2个必须同时相符才能解析出东西来
		//也就是假设模板有one和two,这里只能指定名字one或者名字two,除此之外的任何字符串都不能正常解析出东西！
		tests, err := template.New("one").Parse(testTemplates)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n",tests.Tree)
		//AddParseTree()第一个参数指定返回的新模板的name而不是解析后tree的名字,第二个参数是要返回的模板绑定的tree
		//指定的是map[string]{*text/template.Template}下的key 和 value下的template的name，但是
		//value下的template的tree.name和tree.parsename还是上面指定的one,而这个one指定了解析时候到底去解析什么名字的模板

		//T_AddParseTree, err := tmpl.AddParseTree("two", tests.Tree)
		//T_AddParseTree, err := tmpl.AddParseTree("AddParseTree", tests.Tree)
		//如果我们输入了跟tmpl.name相同的名字，则会将树添加到tmpl源本上面去！
		T_AddParseTree, err := tmpl.AddParseTree("empty", tests.Tree)
		//这里默认是运行root模板（也就是刚开始创建的模板 的名字"empty",我们叫他根模板）
		//注意上面的AddParseTree并不是在源本上面进行绑定，而是新建一个新的模板树，所以这里的templ实际上什么都没绑定
		//下面的这3行代码当且仅当tmpl.AddParseTree("empty", tests.Tree)时候才会有效果
		//默认采用和new（）参数相同的模板来执行
		err = tmpl.Execute(os.Stdout, 0)
		check_err_template(err)
		fmt.Println()


		//ExecuteTemplate（）第一个参数是指定运行的模板的名字
		//err = T_AddParseTree.ExecuteTemplate(os.Stdout, "AddParseTree", 0)
		err = T_AddParseTree.ExecuteTemplate(os.Stdout, "empty", 0)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println("==",b.String())
	}
	TestMessageForExecuteEmpty()
	//输出：
	//&{Name:one ParseName:one Root:one text:{{define "one"}}one{{end}}{{define "two"}}two{{end}} funcs:[] lex:<nil> token:[{typ:15 pos:24 val:}} line:1} {typ:0 pos:0 val: line:0} {typ:0 pos:0 val: line:0}] peekCount:0 vars:[] treeSet:map[]}
	//one
	//one


	fmt.Println()
	fmt.Println("-------------继续探讨template对象之AddParseTree()1111------------")

	const (
		cloneText1 = `{{define "a"}}{{template "b"}}{{template "c"}}{{end}}`
		cloneText2 = `{{define "b"}}b{{end}}`
		cloneText3 = `{{define "c"}}root{{end}}`
		//cloneText3 = `{{define "c"}}{{and .Root .Root1}}{{end}}`//测试能够使用内置的函数
		cloneText4 = `{{define "c"}}clone{{end}}`
	)
	TestAddParseTree:=func () {
		//模板root绑定模板a,并且将它作为根模板，注意这里是创建了2个模板了，new一个模板root,parse又是一个模板a
		//tree.name当前树节点的的名字，parsename是依附的最顶层根节点的名字！
		root, err := template.New("root1").Parse(cloneText1)
		check_err_template(err)

		_, err = root.Parse(cloneText2)//绑定模板b
		check_err_template(err)

		//这里同样是2个模板，cloneText3模板下绑定了c模板，理论上我们应该给定最后一个参数的！但是我们这里省略定义函数了，代而取之的是nil,
		//所以我们不能再c模板中含有内置的一些函数，比如and,or,等等，若果这样做了，会报错的！我这里就不展示了。
		tree, err := parse.Parse("cloneText3", cloneText3, "", "", nil, nil)
		check_err_template(err)
		//这里将c模板添加到root模板下的子节点去，即使c模板的parsename还是指向之前的cloneText3模板的名字
		added, err := root.AddParseTree("c", tree["c"])
		check_err_template(err)

		var b bytes.Buffer
		//------------------------------测试能够使用内置的函数------------------------
		//var m1=map[string]int{
		//	"n0":0,
		//	"n1":1,
		//}
		//
		//err = added.ExecuteTemplate(&b, "a", m1)
		//------------------------------测试能够使用内置的函数------------------------

		//指定模板名字"a"的话，则在公共区域common（map类型）下查找到key="a"的value值假设是T（也是一个template对象），
		// T.Tree.Nodes就会被解析出来,common存储着众多的解析过的模板！理论上可以添加无限个模板，只要你喜欢！
		err = added.ExecuteTemplate(&b, "a", 0)
		check_err_template(err)
		fmt.Println("root.AddParseTree（）解析之后",b.String())
		if b.String() != "broot" {
			fmt.Printf("expected %q got %q", "broot", b.String())
		}

		b.Reset()
		//不指定模板名字的话，则会解析调用者的模板名字，即c，而c.Tree下的Nodes就会被解析出来
		err = added.Execute(&b,  0)
		check_err_template(err)
		fmt.Println("root.AddParseTree（）解析之后",b.String())

	}
	TestAddParseTree()
	//输出：
	//root.AddParseTree（）解析之后 broot
	//root.AddParseTree（）解析之后 root

	fmt.Println("-------------继续探讨template对象之AddParseTree()222------------")

	//上面我们都是在一个模板上面添加另外一个模板，但是下面我们直接添加模板看下行不行，也就是在没任何模板的template
	//对象上面进行继续添加模板，在这里你需要知道的是，假如没有任何模板的话，也就是没有任何的根节点来供新节点依附！
	//这在过去是确实会导致出错，而如今的新版本的go不会报错了
	TestAddParseTreeToUnparsedTemplate:=func () {
		master := "{{define \"master\"}}master11{{end}}"
		tmpl := template.New("master")
		tree, err := parse.Parse("master", master, "", "", nil)
		check_err_template(err)
		masterTree := tree["master"]
		T_AddParseTree, err := tmpl.AddParseTree("master", masterTree)// used to panic
		check_err_template(err)
		err = T_AddParseTree.Execute(os.Stdout, 0)
		check_err_template(err)

	}

	TestAddParseTreeToUnparsedTemplate()
	//输出：
	//master11



	fmt.Println()
	fmt.Println("-------------继续探讨template对象之解析stylesheet------------")

	TestIssue19294:=func () {
		// The empty block in "xhtml" should be replaced during execution
		// by the contents of "stylesheet", but if the internal map associating
		// names with templates is built in the wrong order, the empty block
		// looks non-empty and this doesn't happen.
		var inlined = map[string]string{
			"stylesheet": `{{define "stylesheet"}}stylesheet{{end}}`,
			"xhtml":      `{{block "stylesheet" .}}{{end}}`,
		}
		all := []string{"stylesheet", "xhtml"}
		for i := 0; i < 100; i++ {
			res, err := template.New("title.xhtml").Parse(`{{template "xhtml" .}}`)
			check_err_template(err)
			for _, name := range all {
				//new一个新的模板然后绑定到res下的公共区域common中去，返回的模板除了名字不同之外，
				// 其他的东西完全相同，也就是，其实等价于：在源模板的common中加上新模板后，将源模板的name赋值给新模板！
				//从这里看的出，这个其实就是新的模板了，但是同时影响了旧的模板，这种做法完全是为了省内存！唯一缺点是会影响到源模板。
				_, err := res.New(name).Parse(inlined[name])
				check_err_template(err)
			}
			var buf bytes.Buffer
			err = res.Execute(&buf, 0)
			check_err_template(err)
			fmt.Println(buf.String())

			if buf.String() != "stylesheet" {
				fmt.Println("iteration %d: got %q; expected %q", i, buf.String(), "stylesheet")
			}
		}
	}
	TestIssue19294()
	//暂时没明白这个由什么用！
	//输出：
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet
	//stylesheet

}






// templateFile defines the contents of a template to be stored in a file, for testing.
// templateFile定义要存储在文件中以供测试的模板的内容。
type templateFile struct {
	name     string
	contents string
}

func createTestDir(files []templateFile) string {
	// TempDir creates a new temporary directory in the directory dir
	// with a name beginning with prefix and returns the path of the
	// new directory. If dir is the empty string, TempDir uses the
	// default directory for temporary files (see os.TempDir).
	// Multiple programs calling TempDir simultaneously
	// will not choose the same directory. It is the caller's responsibility
	// to remove the directory when no longer needed.
	// TempDir在目录dir中创建一个新的临时目录，名称以前缀开头，并返回新目录的路径。
	// 如果dir是空字符串，则TempDir使用默认目录存储临时文件（请参见os.TempDir）。
	//多个同时调用TempDir的程序将不会选择同一目录。 不再需要该目录时，调用方有责任删除它。
	//1.创建临时目录
	dir, err := ioutil.TempDir("main3", "template")
	if err != nil {
		log.Fatal(err)
	}
	//2.按照templateFile结构体中的信息 创建临时文件
	for _, file := range files {
		f, err := os.Create(filepath.Join(dir, file.name))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		//3.按照templateFile结构体中的信息 对创建临时文件进行内容写入
		_, err = io.WriteString(f, file.contents)
		if err != nil {
			log.Fatal(err)
		}
	}
	//4.返回临时目录的路径
	return dir
}



func testTemplateParseGlobAndParseFiles() {
	// Here we demonstrate loading a set of templates from a directory.
	//在这里，我们演示从目录加载一组模板。

	// Here we create a temporary directory and populate it with our sample
	// template definition files; usually the template files would already
	// exist in some location known to the program.
	//在这里，我们创建一个临时目录，并使用示例模板定义文件填充该目录； 通常，模板文件将已经存在于程序已知的某个位置。
	//dir是临时目录的路径
	dir := createTestDir([]templateFile{
		// T0.tmpl is a plain template file that just invokes T1.（invokes调用的意思）
		// T0.tmpl是一个纯模板文件，仅调用T1
		{"T0.tmpl", `T0 invokes T1: ({{template "T1"}})`},
		// T1.tmpl defines a template, T1 that invokes T2.
		// T1.tmpl定义了一个调用T2的模板T1。
		{"T1.tmpl", `{{define "T1"}}T1 invokes T2: ({{template "T2"}}){{end}}`},
		// T2.tmpl defines a template T2.
		// T2.tmpl定义模板T2。
		{"T2.tmpl", `{{define "T2"}}This is T2{{end}}`},
	})
	fmt.Println("dir:", dir)

	// pattern is the glob pattern used to find all the template files.
	// pattern是用于查找所有模板文件的全局模式。生成正则匹配规则
	pattern := filepath.Join(dir, "*.tmpl")
	fmt.Println("pattern:", pattern)
	// ParseGlob parses the template definitions in the files identified by the
	// pattern and associates the resulting templates with t. The files are matched
	// according to the semantics of filepath.Match, and the pattern must match at
	// least one file. ParseGlob is equivalent to calling t.ParseFiles with the
	// list of files matched by the pattern.
	//
	// When parsing multiple files with the same name in different directories,
	// the last one mentioned will be the one that results.
	//ParseGlob解析正则模式识别的文件中的模板定义，并将结果模板与t关联。 这些文件根据filepath.Match的语义进行匹配，并且该正则模式必须匹配至少一个文件。
	// ParseGlob等效于使用正则模式匹配的文件列表调用t.ParseFiles。
	//在不同目录中解析具有相同名称的多个文件时，最后一个文件解析后才是结果返回值。

	// Here starts the example proper.
	// T0.tmpl is the first name matched, so it becomes the starting template,
	// the value returned by ParseGlob.
	//从这里开始正确的示例。
	// T0.tmpl是匹配的名字，因此它成为起始模板，即ParseGlob返回的值。
	//无论是下面的template.ParseFiles，还是template.ParseGlob(pattern)都会自动创建一个新的template对象并且调用初始化函数init()，
	//这两个方法都不用实例对象来进行调用，而是直接通过template包调用 即可!当然你也可以通过某个实例模板对象来进行调用，但是假如你的
	// template是已经解析parse过的话，此时再进行第二次或者多次解析到其他模板的话，则无效，也就是还是填充并且输出之前解析过的模板，但是假如你的实例对象是没绑定过模板（或者叫解析到某模板）的，也就是通过
	//template.new(模板名)的形式的话，则会新建一个新的模板并且会填充该模板（又名解析到模板）然后输出结果完整模板，假如你的实例对象是类型为（*template）nil的话，则会在执行Execute（）方法时抛出异常
	//因此大部分时候都最好通过下面的这种函数的形式调用.ParseFiles和.ParseGlob才是最稳妥的！.ParseGlob()和.ParseFiles()是.Parse()方法的拓展而已，实际上功效是一样的，都是填充词解析到模板（或者叫做绑定模板）
	tmpl := template.Must(template.ParseGlob(pattern))

	//其实.ParseGlob(pattern)等价于下面：
	//-----------------------以下等价区--------------------------
	//filenames, err1 := filepath.Glob(pattern)
	//check_err_template(err1)
	//if len(filenames) == 0 {
	//	// Errorf根据格式说明符进行格式化，然后将字符串作为满足错误的值返回。
	//	//如果格式说明符包含带有错误操作数的％w动词，则返回的错误将实现Unwrap方法，返回操作数。
	//	// 包含多个％w动词或向其提供未实现错误接口的操作数是无效的。 另外，％w动词是％v的同义词。
	//	check_err_template(fmt.Errorf("template: pattern matches no files: %#q", pattern))
	//}
	//
	//// ParseFiles creates a new Template and parses the template definitions from
	//// the named files. The returned template's name will have the base name and
	//// parsed contents of the first file. There must be at least one file.
	//// If an error occurs, parsing stops and the returned *Template is nil.
	////
	//// When parsing multiple files with the same name in different directories,
	//// the last one mentioned will be the one that results.
	//// For instance, ParseFiles("a/foo", "b/foo") stores "b/foo" as the template
	//// named "foo", while "a/foo" is unavailable.
	//
	//// ParseFiles创建一个新模板，并从命名文件中解析模板定义。 返回的模板名称将具有第一个文件的基本名称和已解析的内容。 必须至少有一个文件。
	////如果发生错误，解析将停止，返回的* Template为nil。
	////在不同目录中解析具有相同名称的多个文件时，最后提到的将是结果文件。
	////例如，ParseFiles（"a/foo", "b/foo"）将"b/foo"存储为名为"foo"的模板，而"a/foo"不可用。
	//
	//tmpl := template.Must(template.ParseFiles(filenames...))
	//-----------------------以上等价区--------------------------

	err := tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}

	// RemoveAll删除路径及其包含的所有子代。
	//删除所有可能的内容，但返回遇到的第一个错误。 如果路径不存在，则RemoveAll返回nil（无错误）。
	//如果有错误，它将是* PathError类型。

	// Clean up after the test; another quirk of running as an example.
	//测试后清理临时目录以及目录下的所有东西； 以运行为例的另一个怪癖。
	//如果你想要看到创建的目录和文件，请将下面的这行代码注释掉
	defer os.RemoveAll(dir)
}




func check_err_template(err error) {
	if err != nil {
		//fmt.Fprintln(os.Stderr,err)
		//上面的这种方式会导致输出顺序不确定，虽然他可以输出红色的字体，但是由于顺序不确定，我们不采用他！
		fmt.Println(err)
	}
}
