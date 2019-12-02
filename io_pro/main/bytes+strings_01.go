package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("**************bytes********************")
	fmt.Println("bytes库区别于strings库特有的函数如下：")

	//等于报告a和b是否长度相同且包含相同字节。
	// nil参数等效于一个空切片。
	fmt.Println(bytes.Equal([]byte{'a','b','c','d',},[]byte{'a','b','c','d',})) //true
	fmt.Println(bytes.Equal([]byte{'a','b','c','d',},[]byte{'a','b','c',})) //false

	//符文将s解释为UTF-8编码的代码点序列。
	//返回等于s的rune切片（Unicode代码点）。
	fmt.Println(bytes.Runes([]byte{'a','b','c','d',})) //[97 98 99 100]
	fmt.Println(bytes.Runes([]byte{})) //[]
	fmt.Println(bytes.Runes([]byte{' ',})) //[32]
	fmt.Println(bytes.Runes([]byte{'a','A'})) //[97 65]

	fmt.Println("**************strings********************")
	fmt.Println("strings库区别于bytes库特有的函数如下：")
	fmt.Println(strings.ContainsRune("abcdefg",'d'))//true
	fmt.Println(strings.ContainsRune("abcdefg",'z'))//false
	fmt.Println(strings.ContainsRune("abcd中efg",'中'))//true
	fmt.Println(strings.ContainsRune("",'中'))//false
	fmt.Println(strings.ContainsRune(" ",' '))//true


	fmt.Println(strings.ContainsAny("abcd中efg","中g"))//true
	fmt.Println(strings.ContainsAny("abcd中efg","中"))//true
	fmt.Println(strings.ContainsAny("abcd中efg","中国"))//true
	fmt.Println(strings.ContainsAny("abcd中efg","xyz"))//false
	fmt.Println(strings.ContainsAny("abcd中efg",""))//false
	fmt.Println(strings.ContainsAny("",""))//false
	fmt.Println(strings.ContainsAny("","x"))//false
	fmt.Println(strings.ContainsAny(" "," "))//true

	fmt.Println("strings库区别于bytes库特有的类型如下：")
	//Replacer类型进行一系列字符串的替换。
	replacer := strings.NewReplacer("ab","AB","c","C")
	//Replace返回s的所有替换进行完后的拷贝。
	fmt.Println(replacer.Replace("abcdef")) // ABCdef
	fmt.Println(replacer.Replace("cdef")) // Cdef
	fmt.Println(replacer.Replace("def")) // Cdef
	fmt.Println(replacer.Replace("")) // 空字符串
	fmt.Println(replacer.Replace(" ")) // (这里有个空格)

	//WriteString向w中写入s的所有替换进行完后的拷贝。
	ls_Bf:=make([]byte,0,10)//指明往哪里写,我们应该执行他的长度为0，同时应该指定cap为较大的数，
	// 否则会不停的扩容影响效率,如果不指定的话他会根据往里面塞进去多少元素就扩容多少容量
	buffer := bytes.NewBuffer(ls_Bf)
	fmt.Println(replacer.WriteString(buffer,"abcdefgh")) //8 <nil>
	fmt.Println(buffer) //ABCdefgh
	fmt.Println(buffer.Bytes()) //[65 66 67 100 101 102 103 104]
	fmt.Println(buffer.Len()) //8
	fmt.Println(buffer.Cap()) //10
	fmt.Println(ls_Bf) //[]
	//关于Buffer:
		// NewBuffer使用buf作为其初始内容创建并初始化一个新的Buffer。 新的Buffer拥有buf的所有权，并且在此调用之后，
		// 调用方不应使用buf。 NewBuffer旨在准备一个Buffer以读取现有数据。 它也可以用来设置用于写入的内部缓冲区的初始大小。
		// 为此，buf应该具有所需的容量，但长度为零。
		//
		//在大多数情况下，new（Buffer）（或仅声明一个Buffer变量）是
		//足以初始化Buffer。

	fmt.Println("**********************************")
	fmt.Println("strings库和bytes库共有的函数如下：")

	// Compare返回一个按字典顺序比较两个字节片的整数。
	//如果a == b，结果将为0，如果a <b，结果将为-1，如果a> b，结果将为+1。
	// nil参数等效于一个空切片。这个相对与strings中的同名函数更加有相率
	fmt.Println(bytes.Compare([]byte{'a','b','c','d',},[]byte{'a','b','b','d',}))//1
	fmt.Println(bytes.Compare([]byte{'a','b','c','d',},[]byte{'a','b','c','d',}))//0
	fmt.Println(bytes.Compare([]byte{'A','b','c','d',},[]byte{'a','b','c','d',}))//-1  按照unicode值进行比较
	fmt.Println(bytes.Compare([]byte{98,99},[]byte{99}))//-1 按照unicode值进行比较


	// Compare返回一个按字典顺序比较两个字符串的整数。
	//如果a == b，结果将为0，如果a <b，结果将为-1，如果a> b，结果将为+1。
	//
	//仅包含比较，以便与bytes包对称。
	//使用内置的字符串比较运算符==，<，>等通常更清晰，总是更快。

	// NOTE（rsc）：此函数不会调用运行时cmpstring函数，因为我们不想为使用string.Compare提供任何性能证明。 基本上没有人应该使用strings.Compare。
	//如上面的评论所述，这里仅用于与bytes包对称。
	//如果性能很重要，则应更改编译器以识别模式，以便所有进行三向比较的代码（而不仅仅是使用字符串的代码）都可以受益。
	fmt.Println(strings.Compare("abc","abc"))//0
	fmt.Println(strings.Compare("abc","acc"))//-1
	fmt.Println(strings.Compare("abc","aac"))//1
	fmt.Println(strings.Compare("abc","aacsdsdsdsdz"))//1
	fmt.Println(strings.Compare("abc","bacsdsdsdsdz"))//-1
	fmt.Println(strings.Compare("中华人民共和国","中华人民共和"))//1


	//判断两个utf-8编码字符串（将unicode大写、小写、标题三种格式字符视为相同）是否相同。
	fmt.Println(strings.EqualFold("hello", "HellO")) // true
	fmt.Println(strings.EqualFold("hello", "Helll")) // false

	fmt.Println(strings.HasPrefix("hello", "hel")) // true
	fmt.Println(strings.HasPrefix("hello", "Hel")) // false
	fmt.Println(strings.HasSuffix("hello", "hel")) // false
	fmt.Println(strings.HasSuffix("hello", "lo"))  // true

	fmt.Println(strings.Contains("hello world", "hello")) // true
	fmt.Println(strings.Contains("hello world", "Hello")) // false
	fmt.Println(strings.Contains("hello world", ""))      // true
	fmt.Println(strings.Contains("helloworld", ""))       // true
	fmt.Println(strings.Contains("hello world", " "))     // true
	fmt.Println(strings.Contains("helloworld", " "))      // false
	// Count计算s中substr的非重叠实例的数量。
	//如果substr是一个空字符串，则Count返回1 + s中的Unicode代码点数。
	fmt.Println(strings.Count("helloworld", "hel")) // 1
	fmt.Println(strings.Count("helloworld", "l"))   // 3
	fmt.Println(strings.Count("helloworld", ""))    // 11(10+1)
	fmt.Println(strings.Count("helloworld", " "))   // 0
	fmt.Println(strings.Count("helloworld ", " "))  // 1
	fmt.Println(strings.Count("helloworld  ", " ")) // 2

	fmt.Println(strings.Index("helloworld", "d"))   // 9
	fmt.Println(strings.Index("helloworld", "l"))   // 2
	fmt.Println(strings.Index("helloworld", ""))    // 0
	fmt.Println(strings.Index("helloworld", " "))   // -1
	fmt.Println(strings.Index("helloworld ", " "))  // 10
	fmt.Println(strings.Index("helloworld  ", " ")) // 10

	fmt.Println(strings.IndexByte("abchelloworld", 'a'))     // 0
	fmt.Println(strings.IndexByte("abchelloworld", 'b'))     // 1
	fmt.Println(strings.IndexByte("abchelloworld", 99))      // 2
	fmt.Println(strings.IndexByte("abchelloworld", 'v'))     // -1
	fmt.Println(strings.IndexByte("abchelloworld", 0))       // -1
	fmt.Println(strings.IndexByte("abchelloworld", 32))      // -1
	fmt.Println(strings.IndexByte("abchelloworld ", 32))     // 13
	fmt.Println(strings.IndexByte("abchelloworld ", ' '))    // 13
	fmt.Println(strings.IndexByte("abchelloworld", '\n'))    // -1
	fmt.Println(strings.IndexByte("abchelloworld\n", '\n'))  // 13
	fmt.Println(strings.IndexByte("abchelloworld\t", '\t'))  // 13
	fmt.Println(strings.IndexByte("abchelloworld	", '\t'))   // 13(world后面的空格是tab空格)
	fmt.Println(strings.IndexByte("abchelloworld\\t", '\t')) // -1(\被转义了)

	fmt.Println(strings.IndexRune("abchelloworld", 'a'))    // 0
	fmt.Println(strings.IndexRune("abchelloworld", 'c'))    // 2
	fmt.Println(strings.IndexRune("abchelloworld", 98))     // 1
	fmt.Println(strings.IndexRune("abchelloworld中", '中'))   // 13
	fmt.Println(strings.IndexRune("abchelloworld中", '中'))   // 13
	fmt.Println(strings.IndexRune("abchelloworld ", ' '))   // 13
	fmt.Println(strings.IndexRune("abchelloworld\t", ' '))  // -1
	fmt.Println(strings.IndexRune("abchelloworld\t", '\t')) // 13(tab空格)
	fmt.Println(strings.IndexRune("abchelloworld\t", '	'))  // 13(tab空格)
	fmt.Println(strings.IndexRune("", 'a'))                 // -1

	// IndexAny 返回字符串 chars 中的任何一个字符在字符串 s 中第一次出现的位置
	// 如果找不到，则返回 -1，如果 chars 为空，则返回 -1
	fmt.Println(strings.IndexAny("abc-helloworld", "-"))   // 3
	fmt.Println(strings.IndexAny("abc-helloworld", "a"))   // 0
	fmt.Println(strings.IndexAny("abc-helloworld", "abc")) // 0
	fmt.Println(strings.IndexAny("abc-helloworld", "c-"))  // 2
	fmt.Println(strings.IndexAny("abc-helloworld", "w"))   // 9
	fmt.Println(strings.IndexAny("abchelloworld", "w"))    // 8
	fmt.Println(strings.IndexAny("abchelloworld", "wb"))   // 1

	fmt.Println(strings.IndexFunc("abchelloworld", indexD))  // 12
	fmt.Println(strings.IndexFunc("abchelloworld", indexD1)) // -1

	fmt.Println(strings.LastIndex("abchelloworld", "l"))  // 11
	fmt.Println(strings.LastIndex("abchelloworld", "a"))  // 0
	fmt.Println(strings.LastIndex("abchelloworld", "ld")) // 11
	fmt.Println(strings.LastIndex("abchelloworld", "le")) // -1(注意跟IndexAny区分好)
	fmt.Println(strings.LastIndex("abchelloworld", "ow")) // 7
	fmt.Println(strings.LastIndex("abchelloworld", ""))   // 13
	fmt.Println(strings.LastIndex("", ""))                // 0
	fmt.Println(strings.LastIndex("", "a"))               // -1
	fmt.Println(strings.LastIndex("中华人民共和国", "和"))        // 15(其他同原理)

	fmt.Println(strings.LastIndexByte("abchelloworld", 'b')) // 1
	fmt.Println(strings.LastIndexByte("abchelloworld", 'o')) // 9
	fmt.Println(strings.LastIndexByte("abchelloworld", 99))  // 2
	fmt.Println(strings.LastIndexByte("abchelloworld", 32))  // -1

	fmt.Println(strings.LastIndexAny("abchelloworld", "dl")) // 12
	fmt.Println(strings.LastIndexAny("abchelloworld", "la")) // 11
	fmt.Println(strings.LastIndexAny("abchelloworld", "a"))  // 0
	fmt.Println(strings.LastIndexAny("中华国人民共人和国", "人"))      // 18
	fmt.Println(strings.LastIndexAny("中华国人民共人和国", "国人"))     // 24

	fmt.Println(strings.LastIndexFunc("abchelloworld", LastIndFunc)) // 11
	fmt.Println(strings.LastIndexFunc("中华国人民共人和国", LastIndFunc))     // 24

	//返回s中每个单词的首字母都改为标题格式的字符串拷贝。
	//BUG: Title用于划分单词的规则不能很好的处理Unicode标点符号。
	//关于分隔符更详细的请阅读源码参考，这里列出简单的：
	// ASCII字母数字和下划线不是分隔符
	//字母和数字不是分隔符
	//否则，我们现在只能将空格视为分隔符。
	//底层调用的是unicode.ToTitle()方法，
	//底层主要是维护了很多张映射表供非ascii编码的参数操作，如果参数是非ascii编码的话就会多次用到二分查找映射表元素和线性查找元素
	fmt.Println(strings.Title("中华人民共和国"))         // 中华国人民共人和国
	fmt.Println(strings.Title("a中华人民共和国b"))       // A中华人民共和国b
	fmt.Println(strings.Title("中华人民共和国b"))        // 中华人民共和国b
	fmt.Println(strings.Title("abchelloworld"))   // Abchelloworld
	fmt.Println(strings.Title("abc hello world")) // Abc Hello World
	fmt.Println(strings.Title("abc,hello,world")) // Abc,Hello,World
	fmt.Println(strings.Title("abc!hello!world")) // Abc!Hello!World
	fmt.Println(strings.Title("abc中hello中world")) // Abc中hello中world
	fmt.Println(strings.Title("abc·hello·world")) // Abc·hello·world

	//返回将所有字母都转为对应的标题版本的拷贝。
	//底层调用的是unicode.ToTitle()方法
	fmt.Println(strings.ToTitle("abc·hello·world")) // ABC·HELLO·WORLD
	fmt.Println(strings.ToTitle("abchelloworld"))   // ABCHELLOWORLD
	fmt.Println(strings.ToTitle("中华人民共和国"))         // ABCHELLOWORLD
	fmt.Println(strings.ToTitle("a中华人民共和国b"))       // A中华人民共和国B

	//fmt.Println(strings.ToTitleSpecial()) // A中华人民共和国B

	fmt.Println(strings.ToLower("ABC"))       // abc
	fmt.Println(strings.ToLower("A-B-C-"))    // a-b-c-
	fmt.Println(strings.ToLower("A中华人民共和国B")) // a中华人民共和国b
	fmt.Println(strings.ToLower("中华人民共和国"))   // 中华人民共和国
	//fmt.Println(strings.ToLowerSpecial()) // 中华人民共和国
	//fmt.Println(strings.ToLowerSpecial()) // 中华人民共和国

	fmt.Println(strings.ToUpper("a"))         // A
	fmt.Println(strings.ToUpper("ab"))        // AB
	fmt.Println(strings.ToUpper("中华人民共和国"))   // 中华人民共和国
	fmt.Println(strings.ToUpper("a中华人民共和国b")) // A中华人民共和国B

	//fmt.Println(strings.ToUpperSpecial("a中华人民共和国b")) // A中华人民共和国B
	//fmt.Println(strings.ToUpperSpecial("a中华人民共和国b")) // A中华人民共和国B

	fmt.Println(strings.Repeat("a", 2))  // aa
	fmt.Println(strings.Repeat("ab", 2)) // abab
	fmt.Println(strings.Repeat("ab", 4)) // abababab
	fmt.Println(strings.Repeat("ab", 0)) // 空字符串
	//fmt.Println(strings.Repeat("ab",-1)) // panic: strings: negative Repeat count
	fmt.Println(strings.Repeat("ab", 1)) // ab
	var n32 float32 = 2.5
	fmt.Println(strings.Repeat("ab", int(n32))) // abab
	fmt.Println(strings.Repeat("", 10))         // 空字符串
	fmt.Println(strings.Repeat(" ", 10))        //           (前面有很多个空格)
	fmt.Println(strings.Repeat("中", 10))        //中中中中中中中中中中(前面有很多个空格)
	fmt.Println(strings.Repeat("中国", 5))        //中国中国中国中国中国(前面有很多个空格)
	fmt.Println(strings.Repeat("a中国", 5))       //a中国a中国a中国a中国a中国(前面有很多个空格)

	//将s的每一个unicode码值r都替换为mapping(r)，返回这些新码值组成的字符串拷贝。如果mapping返回一个负值，
	//将会丢弃该码值而不会被替换。（返回值中对应位置将没有码值）

	fmt.Println(strings.Map(jin1, "abcdefg")) //bcdefgh
	fmt.Println(strings.Map(upperAndlower, "abcdefg")) //ABCDEFG

	//返回将s前后端所有cutset包含的utf-8码值都去掉的字符串。
	fmt.Println(strings.Trim("abcdefga","a")) //bcdefg
	fmt.Println(strings.Trim("acbcdefgac","ac")) //bcdefg
	fmt.Println(strings.Trim("acbcdeacfg","ac")) //bcdeacfg
	fmt.Println(strings.Trim(" acbcdeacfg   "," ")) //acbcdeacfg
	fmt.Println(strings.Trim("! acbcdeacfg !     ","! ")) //acbcdeacfg
	fmt.Println(strings.Trim("   acbcdeacfg      ","!")) //   acbcdeacfg      (前后有空格)
	fmt.Println(strings.Trim("   acbcdeacfg      ","! ")) //acbcdeacfg
	fmt.Println(strings.Trim("   acbcdeacfg      "," ")) //acbcdeacfg
	fmt.Println(strings.Trim("   acbcdeacfg      ","")) //   acbcdeacfg      (前后有空格)
	fmt.Println(strings.Trim("   acbcdeacfg      ","\t")) //   acbcdeacfg      (前后有空格,但不是tab空格)

	//去掉首位空格
	fmt.Println(strings.TrimSpace("   acbcdeacfg      ")) //acbcdeacfg
	fmt.Println(strings.TrimSpace("   acbc de acfg      ")) //acbc de acfg
	fmt.Println(strings.TrimSpace("!   acbc de acfg    !  ")) //!   acbc de acfg    !

	//返回将s前端所有cutset包含的utf-8码值都去掉的字符串。
	fmt.Println(strings.TrimLeft("   acbcdeacfg      ",""))//   acbcdeacfg      (前后有空格)
	fmt.Println(strings.TrimLeft("   acbcdeacfg      "," "))//acbcdeacfg      (后有空格)
	fmt.Println(strings.TrimLeft("acbcdeacfga","a"))//cbcdeacfga
	fmt.Println(strings.TrimLeft("\tacbcdeacfg\t","\t"))//acbcdeacfg	(后有tab空格)
	fmt.Println("========",strings.TrimLeft("acbcbdeacfgacb","acb"))//deacfgacb

	//和TrimLeft差不多，不再累叙
	fmt.Println(strings.TrimRight("\tacbcdeacfg\t","\t"))//acbcdeacfg	(前有tab空格)


	fmt.Println(strings.TrimFunc("AcbcdeacfgA",delA))//cbcdeacfg
	fmt.Println(strings.TrimLeftFunc("AcbcdeacfgA",delA))//cbcdeacfgA
	fmt.Println(strings.TrimRightFunc("AcbcdeacfgA",delA))//Acbcdeacfg

	//返回去除s可能的前缀prefix的字符串。和trimleft区别好
	fmt.Println(strings.TrimPrefix("AcbcdeacfgAc","Ac"))//bcdeacfgAc
	fmt.Println(strings.TrimPrefix("bcdeacfgAc","Ac"))//bcdeacfgAc
	fmt.Println(strings.TrimPrefix("hello world","hello"))// world
	fmt.Println(strings.TrimPrefix("hello world","hello1"))//hello world
	fmt.Println(strings.TrimPrefix("abcbcdeacfgAcabcbc","abc"))//bcdeacfgAcabcbc
	fmt.Println(strings.TrimPrefix("中abcbcdcbc中","中"))//abcbcdcbc中
	fmt.Println(strings.TrimPrefix(" abcbcdcbc "," "))//abcbcdcbc (后有一个空格)

	//返回去除s可能的后缀suffix的字符串。如果suffix不存在则原样返回
	fmt.Println(strings.TrimSuffix(" abcbcdcbc "," "))// abcbcdcbc(前有一个空格)
	fmt.Println(strings.TrimSuffix("中abcbcdcbc中","中"))//中abcbcdcbc

	//Fields（字段）根据unicode.IsSpace的定义，将字符串s围绕一个或多个连续的空白字符的每个实例进行拆分，
	// 返回s的子字符串切片；如果s仅包含空白，则返回空切片。
	//底层速览：
		//参数为ascii的情况：
			//首先计算字段。
			//如果s为ASCII，则为精确计数，否则为近似值。
			// setBits用于跟踪在s字节中设置了哪些位。
			// ASCII快速路径
			//跳过输入前面的空格。
			//跳过字段之间的空格。
			//最后一个字段可能以EOF结尾。】

		//参数为非ascii的情况：
			// FieldsFunc在每次运行满足f（c）的Unicode代码点c时将字符串s拆分，并返回s的切片数组。 如果s中的所有代码点都满足f（c）或字符串为空，则返回空片。
			// FieldsFunc不保证调用f（c）的顺序。
			//如果f对于给定的c没有返回一致的结果，则FieldsFunc可能会崩溃。

			//跨度用于记录形式为s [start：end]的s的一部分。
			//开始索引是包含的，结束索引是排他的。

			//查找字段的开始和结束索引。

			//最后一个字段可能以EOF结尾。
			//根据记录的字段索引创建字符串。
	fmt.Println(strings.Fields("中 abcb cdcbc中"))//[中 abcb cdcbc中]
	fmt.Println(strings.Fields("    中 abcb cdcbc中    "))//[中 abcb cdcbc中]
	fmt.Println(strings.Fields("    中    abcb    cdcbc中    "))//[中 abcb cdcbc中]
	fmt.Println(strings.Fields(""))//[]
	fmt.Println(strings.Fields(" "))//[]
	fmt.Println(strings.Fields("\n"))//[]
	fmt.Println(strings.Fields("\t"))//[]



	fmt.Println(strings.FieldsFunc("中华人民共和国人中国",SplitByRen))//[中华 民共和国 中国]
	//这个确实是go的不足之处，从这里我们无法得知怎么区分这个列表的元素
	fieldsFunc_s := strings.FieldsFunc("  中华a人  民共和国 a人中 国", SplitByRen)
	fmt.Println(fieldsFunc_s)//[  中华a   民共和国 a 中 国]
	for key, value := range fieldsFunc_s {
		fmt.Println(key,value)//好吧，系统知道怎么区分
		//输出：
		//	0   中华a
		//	1   民共和国 a
		//	2 中 国
	}

	//将slice分割为所有由sep分隔的子字符串，并返回这些分隔符之间的子字符串的一部分。
	//如果s不包含sep且sep不为空，则Split返回长度为1的切片，其唯一元素为s。
	//如果sep为空，则Split在每个UTF-8序列后拆分。 如果s和sep均为空，则Split返回一个空切片。
	//等于SplitN，计数为-1。
	fmt.Println(strings.Split("中华人民共和国人中国","人"))//[中华 民共和国 中国]
	fmt.Println(strings.Split("中华人民共和国人民中国","人民"))//[中华 共和国 中国]
	fmt.Println(strings.Split(" abc "," "))//[ abc ],前面没值补空格
	fmt.Println(strings.Split("abc"," "))//[abc]
	fmt.Println(strings.Split("abca","a"))//[ bc ],前面没值补空格
	fmt.Println(strings.Split("abca","b"))//[a ca]
	fmt.Println(strings.Split("abbbca","b"))//[a   ca]
	for key, value := range strings.Split("abbbca","b") {
		fmt.Println(key,value)
		//输出：
		//	0 a
		//	1
		//	2
		//	3 ca
	}
	fmt.Println(strings.Split("abca",""))//[a b c a]，sep为空的话则按照一个unicode值一个元素组成切片返回
	fmt.Println(strings.Split("只给你abca的go库",""))//[只 给 你 a b c a 的 g o 库]，sep为空的话则按照一个unicode值一个元素组成切片返回，每个unicode值可以拥有不同的编码格式
	fmt.Println(strings.Split("abca","bc"))//[a a]
	//之所以出现首尾被截断的话出现空格的是因为如下原因：
	//ls:=make([]string,3,3)
	//ls[0]=""
	//ls[1]="bc"
	//ls[2]=""
	//fmt.Println("这是我们自己构造的首尾2个空字符串元素的切片：",ls)//这是我们自己构造的首尾2个空字符串元素的切片： [ bc ]

	fmt.Println("========================================")
	//用去掉s中出现的sep的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片（每一个sep都会进行一次切割，
	//即使两个sep相邻，也会进行两次切割）。如果sep为空字符，Split会将s切分成每一个unicode码值一个字符串。
	// 参数n决定返回的切片的数目：
	//
	//n > 0 : 返回的切片最多n个子字符串；最后一个子字符串包含未进行切割的部分。
	//n == 0: 返回nil
	//n < 0 : 返回所有的子字符串组成的切片,也就是从头切到结尾
	fmt.Println(strings.SplitN("cabcdaefgaz","a",2))//[c bcdaefgaz]
	fmt.Println(strings.SplitN("cabcdaefgaz","a",3))//[c bcd efgaz]
	fmt.Println(strings.SplitN("cabcdaefgaz","a",0))//[]
	fmt.Println(strings.SplitN("cabcdaefgaz","a",-1))//[c bcd efg z]

	fmt.Println("========================================")
	//并不是从后往前切，而是将切点置于切符sep的后面，而不是置于切符sep上s,这会导致保留sep切符

	fmt.Println(strings.SplitAfter("cabcdaefgaz","a"))//[ca bcda efga z]
	fmt.Println(strings.SplitAfter("cabcdaefgaz",""))//[c a b c d a e f g a z]
	fmt.Println(strings.SplitAfter("cabc dae fgaz"," "))//[cabc  dae  fgaz]
	fmt.Println(strings.SplitAfter("cabc dae fgaz"," "))//[cabc  dae  fgaz]



	fmt.Println(strings.SplitAfterN("cabc dae fgaz"," ",2))//[cabc  dae fgaz]
	fmt.Println(strings.SplitAfterN("cabcdaefgaz","a",2))//[ca bcdaefgaz]
	fmt.Println(strings.SplitAfterN("中华人民共和国人民国中","中",2))//[中 华人民共和国人民国中]
	fmt.Println(strings.SplitAfterN("中华人民共和国人民国中","中",3))//[中 华人民共和国人民国中 ],注意这里
	fmt.Println(strings.SplitAfterN("中华人民共和国人民国中","人",3))//[中华人 民共和国人 民国中]
	fmt.Println(strings.SplitAfterN("中华人民共和国人民国中","",3))//[中 华 人民共和国人民国中]，注意
	fmt.Println(strings.SplitAfterN(" 中华人民共和国人民国中 "," ",30))//[  中华人民共和国人民国中  ]
	fmt.Println(strings.SplitAfterN("中华人民共和国人民国中","人",30))//[中华人 民共和国人 民国中]

	// Join连接切片的元素以创建单个字符串。 分隔符字符串sep放置在结果字符串中的元素之间。
	ls111:=[]string{"a", "b", "c",}
	fmt.Println(strings.Join(ls111,"人"))//a人b人c
	fmt.Println(strings.Join(ls111," "))//a b c
	fmt.Println(strings.Join(ls111,"-"))//a-b-c
	fmt.Println(strings.Join(ls111,""))//abc
	fmt.Println(strings.Join(ls111,"\n"))//"a\nb\nc",所以输出的时候会显示换行
	//输出如下：
	//	a
	//	b
	//	c
	ls222:=[]string{}
	fmt.Println(strings.Join(ls222,"人"))//空字符串

	ls333:=[]string{"a"}
	fmt.Println(strings.Join(ls333,"人"))//a



}



func SplitByRen(r rune)bool  {
	if r=='人'{
		return true
	}
	return false
}


func delA(r rune) bool {
	if r=='A'{
		return true
	}
	return false
}


func jin1(r rune) rune {
	switch {
	case r >= 'A' && r <= 'Z':
		return 'A' + (r-'A'+1)%26
	case r >= 'a' && r <= 'z':
		return 'a' + (r-'a'+1)%26
	}
	return r
}

func upperAndlower(r rune) rune {
	switch {
	case r >= 'A' && r <= 'Z':
		return 'a' + (r-'a'+32)%26
	case r >= 'a' && r <= 'z':
		return 'A' + (r-'A'-32)%26
	}
	return r
}

func indexD(r rune) bool {
	if r == 100 {
		return true
	}
	return false
}
func indexD1(r rune) bool {
	if r == 10000 {
		return true
	}
	return false
}

func LastIndFunc(r rune) bool {
	if r == 'l' || r == '国' {
		return true
	}
	return false
}
