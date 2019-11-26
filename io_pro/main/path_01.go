package main

import (
	"fmt"
	"path"
)

func main111() {

	//path实现了对斜杠分隔的路径的实用操作函数。

	fmt.Println("------------------------")
	//项目路径判断依据：长度不为零并且以'/'开头
	fmt.Println(path.IsAbs("main/os_exec_01.go"))//false
	fmt.Println(path.IsAbs("/main/os_exec_01.go"))//true
	fmt.Println(path.IsAbs(`C:\Users\Administrator\Desktop\go_pro\src\io_pro\main\os_exec_01.go`))//false



	fmt.Println("------------------------")
	// Dir返回除路径的最后一个元素以外的所有元素，通常是路径的目录。
	//使用Split删除最终元素后，路径将被清理，尾部的斜杠将被删除。
	//如果路径为空，则Dir返回“.”。
	//如果路径完全由斜杠后跟非斜杠字节组成，则Dir返回单个斜杠。 在任何其他情况下，返回的路径都不以斜杠结尾。
	fmt.Println(path.Dir(""))// 返回.
	fmt.Println(path.Dir(" "))// 返回.
	fmt.Println(path.Dir("    "))// 返回.
	fmt.Println(path.Dir("/aa"))//  返回/
	fmt.Println(path.Dir("/aa/bb"))//  返回/aa
	fmt.Println(path.Dir("/aa/bb/"))//  返回/aa/bb
	fmt.Println(path.Dir("/aa/bb/cc"))//  返回/aa/bb
	fmt.Println(path.Dir("///"))//  返回/
	fmt.Println(path.Dir("//"))//  返回/
	fmt.Println(path.Dir("/"))//  返回/
	fmt.Println(path.Dir("/  "))//  返回/
	fmt.Println(path.Dir("aa/bb"))//  返回aa
	fmt.Println(path.Dir("aa/bb/"))//  返回aa/bb
	fmt.Println(path.Dir("aa/b b/"))//  返回aa/b b
	fmt.Println(path.Dir("  aa/b b/  "))//  返回  aa/b b(aa前面有空格)
	fmt.Println(path.Dir("a  aa/b b/  "))//  返回a  aa/b b
	fmt.Println(path.Dir("../aa/bb "))//  返回../aa
	fmt.Println(path.Dir("../../aa/bb "))//  返回../../aa
	fmt.Println(path.Dir(".. /../aa/bb "))//  返回aa
	fmt.Println(path.Dir("../.. /aa/bb "))//  返回../.. /aa
	fmt.Println(path.Dir("../. ./aa/bb "))//  返回../. ./aa
	fmt.Println(path.Dir(". ./../aa/bb "))//  返回aa
	fmt.Println(path.Dir("  ../../aa/bb "))//  返回aa
	fmt.Println(path.Dir("./../aa/bb "))//  返回../aa
	fmt.Println(path.Dir("././aa/bb "))//  返回aa
	fmt.Println(path.Dir(".././aa/bb "))//  返回../aa
	fmt.Println(path.Dir("../  .  .  /aa/bb "))//  返回../  .  .  /aa
	//总结，返回最后一个/之前的所有东西，path只含有单个或者多个斜杠的话只返回单个斜杠，path为空或者path只含有单个或者多个空格只返回一个.

	fmt.Println("------------------------")
	// Base返回路径的最后一个元素。
	//在提取最后一个元素之前，先删除斜杠。
	//如果路径为空，则Base返回“。”。
	//如果路径完全由斜杠组成，则Base返回“ /”。
	//跟path.Dir()完全相反
	fmt.Println(path.Base("aa/bb"))//bb
	fmt.Println(path.Base("/aa/bb"))//bb
	fmt.Println(path.Base("/aa/bb/"))//bb
	fmt.Println(path.Base("/aa/"))//aa
	fmt.Println(path.Base("/aa"))//aa
	fmt.Println(path.Base("aa"))//aa
	fmt.Println(path.Base("aa/"))//aa
	fmt.Println(path.Base(""))//.
	fmt.Println(path.Base("/"))// 返回/
	fmt.Println(path.Base("//"))// 返回/
	fmt.Println(path.Base("///"))// 返回/
	fmt.Println(path.Base("///aa"))// 返回aa
	fmt.Println(path.Base("//bb/aa"))// 返回aa
	fmt.Println(path.Base("cc//bb/aa"))// 返回aa
	fmt.Println(path.Base("/cc//bb/aa"))// 返回aa
	fmt.Println(path.Base("/cc//bb/aa/"))// 返回aa
	fmt.Println(path.Base("/cc//bb/aa//"))// 返回aa
	fmt.Println(path.Base("/cc//bb/中//"))// 返回中
	fmt.Println(path.Base("/c c/bb"))// 返回bb
	fmt.Println(path.Base("/c c/b b"))// 返回b b
	fmt.Println(path.Base("/c c/b b    "))// 返回b b    (后面有几个空格)
	fmt.Println(path.Base("   /c c/b b    "))// 返回b b    (后面有几个空格)
	fmt.Println(path.Base("   /"))// 返回   (后面有几个空格)
	fmt.Println(path.Base("   /          "))// 返回          (后面有几个空格)
	fmt.Println(path.Base("   /          /"))// 返回          (后面有几个空格)
	fmt.Println(path.Base("aa/bb/.."))// 返回..
	fmt.Println(path.Base("aa/bb/..aa"))// 返回..aa
	fmt.Println(path.Base("aa/bb/../aa"))// 返回aa
	fmt.Println(path.Base("aa/bb/../aa.."))// 返回aa..
	fmt.Println(path.Base("../aa/bb/aa"))// 返回aa
	//总结，返回最后的字符串，完全由/组成的path则返回单个/,空的path返回.


	fmt.Println("------------------------")
	// Clean通过纯词法处理返回与path等效的最短路径名。 它反复应用以下规则，直到无法进行进一步处理为止：
	// 1.用一个斜杠替换多个斜杠。
	// 2.消除每个.路径名元素（当前目录）。
	// 3.删除每个内部..路径名元素（父目录）以及在其前面的non-..元素。
	// 4.消除以..元素开头的根路径：
	//，即在路径的开头将“/..”替换为“/”。
	//
	//返回的路径仅当其根为“/”时才以斜杠结尾。
	//
	//如果此过程的结果为空字符串，则Clean返回字符串“.”。
	fmt.Println(path.Clean(""))//.
	fmt.Println(path.Clean("/"))//返回/
	fmt.Println(path.Clean("//"))//返回/
	fmt.Println(path.Clean("///"))// 返回/
	fmt.Println(path.Clean("\n"))//返回空字符串
	fmt.Println(path.Clean("\t"))//返回	(这里有一个tab空格)
	fmt.Println(path.Clean("\taa"))//返回	aa(这里有一个tab空格,不知道为什么会这样)
	fmt.Println(path.Clean("/aa/bb"))// 返回/aa/bb
	fmt.Println(path.Clean("../aa/bb"))// 返回../aa/bb
	fmt.Println(path.Clean("../../aa/bb"))// 返回../../aa/bb
	fmt.Println(path.Clean("../../aa/bb/"))// 返回../../aa/bb
	fmt.Println(path.Clean("/../../aa/bb/"))// 返回/aa/bb
	fmt.Println(path.Clean("/../../aa/bb"))// 返回/aa/bb
	fmt.Println(path.Clean("/../../aa/../bb"))// 返回/bb
	fmt.Println(path.Clean("/../../"))// 返回/
	fmt.Println(path.Clean("bb/../../"))// 返回..
	fmt.Println(path.Clean("./bb/aa"))// 返回bb/aa
	fmt.Println(path.Clean("./bb/aa."))// 返回bb/aa.
	fmt.Println(path.Clean("../bb/aa."))// 返回../bb/aa.
	fmt.Println(path.Clean("..bb/aa"))// 返回..bb/aa
	fmt.Println(path.Clean(".bb/aa"))// 返回.bb/aa
	fmt.Println(path.Clean(".bb/aa"))// 返回.bb/aa
	//总结，消除../../这样的东西，但是如果最前面有/的话则返回值加上/，空字符串返回. ，全部是斜杠的话返回单个斜杠

	fmt.Println("------------------------")
	// Join将任意数量的路径元素连接到单个路径中，并添加了
	//必要时分隔斜线。 结果是Cleaned； 特别是，所有空字符串都将被忽略。
	fmt.Println(path.Join("a","b"))// 返回a/b
	fmt.Println(path.Join("a","/b"))// 返回a/b
	fmt.Println(path.Join("a/","/b"))// 返回a/b
	fmt.Println(path.Join("a//","/b"))// 返回a/b
	fmt.Println(path.Join("a///","/b"))// 返回a/b
	fmt.Println(path.Join("a//","//b"))// 返回a/b
	fmt.Println(path.Join("a","b/c"))// 返回a/b/c
	fmt.Println(path.Join("a/d","b/c"))// 返回a/d/b/c
	fmt.Println(path.Join("a/..","b/c"))// 返回b/c
	fmt.Println(path.Join("a/../","b/c"))// 返回b/c
	fmt.Println(path.Join("../a","b/c"))// 返回../a/b/c
	fmt.Println(path.Join("../../a/","b/c"))// 返回../../a/b/c
	fmt.Println(path.Join("../../a","b/c"))// 返回../../a/b/c
	fmt.Println(path.Join("../../a../","b/c"))// 返回../../a../b/c
	fmt.Println(path.Join("../../a..","b/c"))// 返回../../a../b/c
	fmt.Println(path.Join("../../a/","../b/c"))// 返回../../b/c
	fmt.Println(path.Join("a/","../b/c"))// b/c
	fmt.Println(path.Join("a/..","../b/c"))//返回../b/c
	fmt.Println(path.Join("a/../","../b/c"))//返回../b/c
	fmt.Println(path.Join("a/../","/../b/c"))//返回../b/c
	fmt.Println(path.Join("a/../","/../b../c"))//返回../b../c
	fmt.Println(path.Join("a/","/..b/c"))//返回a/..b/c
	fmt.Println(path.Join("a/","/.b/c"))//返回a/.b/c
	fmt.Println(path.Join("a/","/.b/c.."))//返回a/.b/c..
	fmt.Println(path.Join("a/","/b/..c"))//返回a/b/..c
	fmt.Println(path.Join("a/","/b/.c"))//返回a/b/.c
	fmt.Println(path.Join("a/","/b/.c.."))//返回a/b/.c..
	fmt.Println(path.Join("a/","/b/.c..///"))//返回a/b/.c..
	fmt.Println(path.Join("a/","/"))//返回a
	fmt.Println(path.Join("/","/"))//返回/
	fmt.Println(path.Join("","/"))//返回/
	fmt.Println(path.Join("",""))//返回空字符串
	fmt.Println(path.Join("////",""))//返回/
	fmt.Println(path.Join("//a//",""))//返回/a
	fmt.Println(path.Join("//a/",""))//返回/a
	fmt.Println(path.Join("a/",""))//返回a
	fmt.Println(path.Join("..a/",""))//返回..a
	fmt.Println(path.Join("/..a/",""))//返回/..a
	fmt.Println(path.Join("a","\t"))//返回a/	(有tab空格)
	fmt.Println(path.Join("a","\tbb"))//返回a/	bb(有tab空格)
	fmt.Println(path.Join("a\t","\tbb"))//返回a	/	bb(有tab空格)
	fmt.Println(path.Join("\ta","\tbb"))//返回	a/	bb(有tab空格)
	fmt.Println(path.Join("/\ta","\tbb"))//返回/	a/	bb(有tab空格)
	fmt.Println(path.Join("a","\nbb"))//返回a/(换行)bb



	fmt.Println("------------------------")
	// Split会在最后一个斜杠之后立即分割路径，将其分为目录和文件名部分。
	//如果路径中没有斜杠，则Split返回一个空的目录，并将文件设置为path。
	//返回的值具有path = dir + file的属性。
	dir, file := path.Split(`C:\Users\Public\Downloads`)
	fmt.Println(dir)//空字符串
	fmt.Println(file)//C:\Users\Public\Downloads

	dir1, file1 := path.Split("C:\\Users\\Public\\Downloads")
	fmt.Println(dir1)//空字符串
	fmt.Println(file1)//C:\Users\Public\Downloads

	dir2, file2 := path.Split("C:/Users/Public/Downloads")
	fmt.Println(dir2)//C:/Users/Public/
	fmt.Println(file2)//Downloads,之所以会出现跟上面的不同是因为这个/让go知道了这是linux环境然后在linux中不用后缀名字也是文件的！

	dir3, file3 := path.Split("C:/Users/Public/Downloads.html")
	fmt.Println(dir3)//C:/Users/Public/
	fmt.Println(file3)//Downloads.html

	dir4, file4 := path.Split("C:\\Users\\Public\\Downloads.html")
	fmt.Println(dir4)//空字符串
	fmt.Println(file4)//C:\Users\Public\Downloads.html

	dir5, file5 := path.Split(`C:\Users\Public\Downloads.html`)
	fmt.Println(dir5)//空字符串
	fmt.Println(file5)//C:\Users\Public\Downloads.html

	//总结：
	//i := strings.LastIndex(path, "/")获取最后斜杠，而且还是只是这种斜杠/
	//return path[:i+1], path[i+1:]返回斜杠的前后2部分



	fmt.Println("------------------------")

//如果name匹配shell文件名模式匹配字符串，Match函数返回真。该模式匹配字符串语法为：
//
//pattern:
//	{ term }
//term:
//	'*'                                  匹配0或多个非/的字符
//	'?'                                  匹配1个非/的字符
//	'[' [ '^' ] { character-range } ']'  字符组（必须非空）
//	c                                    匹配字符c（c != '*', '?', '\\', '['）
//	'\\' c                               匹配字符c
//	character-range:
//	c           匹配字符c（c != '\\', '-', ']'）
//	'\\' c      匹配字符c
//	lo '-' hi   匹配区间[lo, hi]内的字符

	fmt.Println(path.Match("","C:/Users/Public/Downloads.html"))//false <nil>
	fmt.Println(path.Match("*","C:/Users/Public/Downloads.html"))//false <nil>
	fmt.Println(path.Match("^[0-9]*$","01223/2323/"))//false <nil>
	fmt.Println(path.Match("^[0-9]*$","0"))//false <nil>
	//似乎并不关正则的事
	fmt.Println(path.Match("*","CDownloadshtml"))//true <nil>
	fmt.Println(path.Match("*","C/Downloadshtml"))//false <nil>
	fmt.Println(path.Match("*","CDownloadshtml222"))//true <nil>
	fmt.Println(path.Match("*","CDownloads\\html222"))//true <nil>
	fmt.Println(path.Match("?","CDownloads\\html222"))//false <nil>
	fmt.Println(path.Match("?","C"))//true <nil>
	fmt.Println(path.Match("?","2"))//true <nil>
	fmt.Println(path.Match("?","C2"))//false <nil>
	fmt.Println(path.Match("?","/"))//false <nil>
	fmt.Println(path.Match("a","b"))//false <nil>
	fmt.Println(path.Match("b","b"))//true <nil>
	fmt.Println(path.Match("b+","b"))//false <nil>,根本无法用到正则，不过有点像正则而已
	fmt.Println(path.Match("b+","b+"))//true <nil>
	fmt.Println(path.Match("/","/"))//true <nil>
	fmt.Println(path.Match("\\","\\"))//false syntax error in pattern,报错
	fmt.Println(path.Match("?","?"))//true <nil>
	fmt.Println(path.Match("[","["))//false syntax error in pattern
	fmt.Println(path.Match("*","*"))//true <nil>
	fmt.Println(path.Match("*",""))//true <nil>
	fmt.Println(path.Match("*","/"))//false <nil>
	//fmt.Println(path.Match("{*[^{0-9}]}","2"))//false <nil>
	fmt.Println(path.Match("*x", "xxx"))//true <nil>
	fmt.Println(path.Match("*x", "sdx"))//true <nil>
	fmt.Println(path.Match("*x", "/dx"))//false <nil>


	fmt.Println("------------------------")
	// Ext返回路径使用的文件扩展名。
	//扩展名是从path的最后斜杠分隔的元素中的最后一个点开始的后缀；
	//如果没有点，则为空。Ext是extend的缩写，是拓展名的意思
	fmt.Println(path.Ext("C:/Users/Public/Downloads.html"))//返回.html
	fmt.Println(path.Ext("C:/Users/Public/Downloads"))//返回空字符串
	fmt.Println(path.Ext("C://Users//Public//Downloads"))//返回空字符串
	fmt.Println(path.Ext("C://Users//Public//Downloads.html"))//返回.html
	fmt.Println(path.Ext("C:\\Users\\Public\\Downloads.html"))//返回.html
	fmt.Println(path.Ext("C:\\Users\\Public\\Downloads"))//返回空字符串
	fmt.Println(path.Ext("C:\\Users\\Public\\Downloads."))//返回.
	fmt.Println(path.Ext("C:\\Users\\Public\\Downloads..."))//返回.
	fmt.Println(path.Ext("C:\\Users\\Public\\Downloads...a"))//返回.a
	fmt.Println(path.Ext("C:\\Users\\Public\\Downloads...a."))//返回.
	fmt.Println(path.Ext(".html"))//返回.html
	fmt.Println(path.Ext("/.html"))//返回.html
	fmt.Println(path.Ext("//.html"))//返回.html
	fmt.Println("====")
	fmt.Println(path.Ext(".//html"))//返回空字符串
	fmt.Println(path.Ext("./html"))//返回空字符串
	fmt.Println(path.Ext(".h/tml"))//返回空字符串
	fmt.Println(path.Ext(".\\html"))//返回.\html
	fmt.Println(path.Ext("\\.\\html"))//返回.\html
	fmt.Println("====")
	fmt.Println(path.Ext(".html/"))//返回空字符串
	fmt.Println(path.Ext(".html//"))//返回空字符串
	fmt.Println(path.Ext(".html\\/"))//返回空字符串
	fmt.Println(path.Ext(".html.\\/"))//返回空字符串

	fmt.Println("---------------------------")
	fmt.Println(path.ErrBadPattern)//syntax error in pattern
}


























