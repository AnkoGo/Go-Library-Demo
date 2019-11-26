package main

import (
	"fmt"
	"os"
	"path/filepath"
)




func main555() {
	//这个包会根据平台输出不同的斜杠

	fmt.Println("------------------------------")
	// Abs返回路径的绝对表示。
	//如果该路径不是绝对路径，它将与当前工作目录合并，以将其转换为绝对路径。 给定文件的绝对路径名不能保证唯一。
	// Abs对结果调用Clean。

	fmt.Println(filepath.Abs("aa/bb"))//C:\Users\Administrator\Desktop\go_pro\src\io_pro\aa\bb <nil>
	fmt.Println(filepath.Abs("/aa/bb"))//C:\aa\bb <nil>
	fmt.Println(filepath.Abs("\\aa\\bb"))//C:\aa\bb <nil>
	fmt.Println(filepath.Abs("C:\\aa\\bb"))//C:\aa\bb <nil>
	fmt.Println(filepath.Abs(`C:\aa\bb`))//C:\aa\bb <nil>
	fmt.Println(filepath.Abs(`\aa\bb`))//C:\aa\bb <nil>
	fmt.Println(filepath.Abs(`\aa\bb\`))//C:\aa\bb <nil>
	fmt.Println(filepath.Abs(`..\aa\bb\`))//C:\Users\Administrator\Desktop\go_pro\src\aa\bb <nil>
	fmt.Println(filepath.Abs(`\..aa\bb\`))//C:\..aa\bb <nil>
	fmt.Println(filepath.Abs(`\aa\bb\..`))//C:\aa <nil>
	fmt.Println(filepath.Abs(`\aa\bb\cc\..`))//C:\aa\bb <nil>
	fmt.Println(filepath.Abs(`\aa\bb\cc\..\`))//C:\aa\bb <nil>
	fmt.Println(filepath.Abs("\\aa\\bb\\cc\\..\\"))//C:\aa\bb <nil>
	fmt.Println(filepath.Abs("/aa/bb/cc/../"))//C:\aa\bb <nil>
	fmt.Println(filepath.Abs("/"))//C:\ <nil>
	fmt.Println(filepath.Abs("//"))// 返回\ <nil>
	fmt.Println(filepath.Abs("///"))// 返回\ <nil>
	fmt.Println(filepath.Abs("../"))// 返回C:\Users\Administrator\Desktop\go_pro\src <nil>
	fmt.Println(filepath.Abs("../../"))// 返回C:\Users\Administrator\Desktop\go_pro <nil>
	fmt.Println(filepath.Abs("."))// 返回C:\Users\Administrator\Desktop\go_pro\src\io_pro <nil>
	fmt.Println(filepath.Abs("/."))// 返回C:\ <nil>
	fmt.Println(filepath.Abs("/"))// 返回C:\ <nil>
	fmt.Println(filepath.Abs("/../"))// 返回C:\ <nil>


	fmt.Println("------------------------------")

	fmt.Println(filepath.IsAbs(`C:\Users\Administrator\Desktop\go_pro\src\io_pro\aa\bb`))//true
	fmt.Println(filepath.IsAbs(`C:\Users\Administrator\Desktop\go_pro\src\io_pro\aa\bb.html`))//true
	fmt.Println(filepath.IsAbs(`aa\bb`))//false
	fmt.Println(filepath.IsAbs(`\aa\bb`))//false
	fmt.Println(filepath.IsAbs(`..\aa\bb`))//false
	fmt.Println(filepath.IsAbs(`../../../../../../../../`))//false
	fmt.Println(filepath.IsAbs(`/`))//false
	fmt.Println(filepath.IsAbs(`.`))//false
	fmt.Println(filepath.IsAbs(`C:\Users\`))//true
	fmt.Println(filepath.IsAbs(`C:\Users\..\..\`))//true
	fmt.Println(filepath.IsAbs(`C:\Users\..\..\..\..\`))//true
	fmt.Println(filepath.IsAbs(`C:\Users\哈哈\`))//true
	fmt.Println(filepath.IsAbs(`C:`))//false
	fmt.Println(filepath.IsAbs(`C:\`))//true
	fmt.Println(filepath.IsAbs(`H:\..\..\`))//true
	fmt.Println(filepath.IsAbs(`中:\..\..\`))//false
	fmt.Println(filepath.IsAbs(`1:\..\..\`))//false
	fmt.Println(filepath.IsAbs(`c:\..\..\`))//true
	fmt.Println(filepath.IsAbs(`z:\..\..\`))//true
	fmt.Println(filepath.IsAbs(`zz:\..\..\`))//false
	fmt.Println(filepath.IsAbs(`c:c:\..\..\`))//false
	fmt.Println(filepath.IsAbs(`c:.`))//false
	fmt.Println(filepath.IsAbs(`c:\.`))//true
	fmt.Println(filepath.IsAbs(`c:\\.`))//true
	fmt.Println(filepath.IsAbs(`c:\\\.`))//true
	fmt.Println(filepath.IsAbs(`c:\\\\\\.`))//true
	fmt.Println(filepath.IsAbs(`c:\\\\\\//.`))//true
	fmt.Println(filepath.IsAbs(`c://.`))//true
	fmt.Println(filepath.IsAbs(`c:a//.`))//false


	fmt.Println("------------------------------")
	// Base返回路径的最后一个元素。
	//在提取最后一个元素之前，先删除尾随路径分隔符。
	//如果路径为空，则Base返回“.”。
	//如果路径完全由分隔符组成，则Base返回单个分隔符。
	fmt.Println(filepath.Base(`C:\Users\Administrator\Desktop\go_pro\src`))//src
	fmt.Println(filepath.Base(`C:\Users\Administrator\Desktop\go_pro\src.html`))//src.html
	fmt.Println(filepath.Base(`C:\Users\Administrator\Desktop\go_pro\\src.html`))//src.html
	fmt.Println(filepath.Base(`C:\Users\Administrator\Desktop\go_pro\\\src.html`))//src.html
	fmt.Println(filepath.Base(`C:\Users\Administrator\Desktop\go_pro\src.html\`))//src.html
	fmt.Println(filepath.Base(`C:\Users\Administrator\Desktop\go_pro\src\`))//src
	fmt.Println(filepath.Base(`C:\Users\Administrator\Desktop\go_pro\src\..`))//返回..
	fmt.Println(filepath.Base(`C:\\Users\\Administrator\\Desktop\\go_pro\\src\\..`))//返回..
	fmt.Println(filepath.Base(`C:\\Users\\Administrator\\Desktop/go_pro/src\\..`))//返回..
	fmt.Println(filepath.Base(`C:\\Users\\Administrator\\Desktop/go_pro/src\\..中`))//返回..中

	fmt.Println("------------------------------")
	// Ext返回路径使用的文件扩展名。
	//扩展名是从path的最后一个元素中的最后一个点开始的后缀； 如果没有点，则为空。
	fmt.Println(filepath.Ext(`C:\Users\Administrator\Desktop\go_pro\src.html`))//.html
	fmt.Println(filepath.Ext(`C:\Users\Administrator\Desktop\go_pro\src.`))//返回.
	fmt.Println(filepath.Ext(`C:\Users\Administrator\Desktop\go_pro\src`)) //返回空字符串
	fmt.Println(filepath.Ext(`go_pro\src`)) //返回空字符串
	fmt.Println(filepath.Ext(`go_pro\src.html`)) //返回.html
	fmt.Println(filepath.Ext(`go_pro\src.htm.l`)) //返回.l
	fmt.Println(filepath.Ext(`go_pro\src.\html`)) //返回空字符串
	fmt.Println(filepath.Ext(``)) //返回空字符串
	fmt.Println(filepath.Ext(` `)) //返回空字符串
	fmt.Println(filepath.Ext(` .html`)) //返回.html
	fmt.Println(filepath.Ext(` .   html`)) //返回.   html
	fmt.Println(filepath.Ext(` .   html   `)) //返回.   html   (ml后面还有空格)
	fmt.Println(filepath.Ext(` .   html   `)) //返回.   html   (ml后面还有空格)
	fmt.Println(filepath.Ext(` .   `)) //返回.   (.后面还有空格)
	fmt.Println(filepath.Ext(` ...aaa`)) //返回.aaa
	fmt.Println(filepath.Ext(`././././`)) //返回空字符串
	fmt.Println(filepath.Ext(`.//.//./z/./z/`)) //返回空字符串
	fmt.Println(filepath.Ext(`.//.//./z/.z/z/`)) //返回空字符串
	fmt.Println(filepath.Ext(`.//.//./z/.z/z/.z`)) //返回.z

	fmt.Println("------------------------------")
	// Dir返回除路径的最后一个元素以外的所有元素，通常是路径的目录。
	//删除最后一个元素后，Dir会在路径上调用Clean并删除斜杠。
	//如果路径为空，则Dir返回“。”。
	//如果路径完全由分隔符组成，则Dir返回单个分隔符。
	//除非它是根目录，否则返回的路径不会以分隔符结尾。
	fmt.Println(filepath.Dir(`C:\Users\Administrator\Desktop\go_pro\src.html`))//C:\Users\Administrator\Desktop\go_pro
	fmt.Println(filepath.Dir(`C:\Users\Administrator\Desktop\go_pro\src`))//C:\Users\Administrator\Desktop\go_pro
	fmt.Println(filepath.Dir(`C:\Users\Administrator\Desktop\go_pro\src\`))//C:\Users\Administrator\Desktop\go_pro\src
	fmt.Println(filepath.Dir(`C:\Users\Administrator\Desktop\go_pro\src\`))//C:\Users\Administrator\Desktop\go_pro\src
	fmt.Println(filepath.Dir(`\C:\Users\Administrator\Desktop\go_pro\src\`))// 返回\C:\Users\Administrator\Desktop\go_pro\src
	fmt.Println(filepath.Dir(`\go_pro\src\`))// 返回\go_pro\src
	fmt.Println(filepath.Dir(`/go_pro/src/`))// 返回\go_pro\src
	fmt.Println(filepath.Dir(`/go_pro`))// 返回\
	fmt.Println(filepath.Dir(`/go_pro/`))// 返回\go_pro
	fmt.Println(filepath.Dir(`/go_pro//`))// 返回\go_pro
	fmt.Println(filepath.Dir(`//go_pro//`))// 返回\go_pro
	fmt.Println(filepath.Dir(`///go_pro///`))// 返回\go_pro
	fmt.Println(filepath.Dir(`\\\go_pro\\\`))// 返回\go_pro
	fmt.Println(filepath.Dir(`\\\go_pro///`))// 返回\go_pro
	fmt.Println(filepath.Dir(`/`))// 返回\
	fmt.Println(filepath.Dir(`/.`))// 返回\
	fmt.Println(filepath.Dir(`/./`))// 返回\
	fmt.Println(filepath.Dir(`C:\Users\Administrator\Desktop\go_pro\src.html`))// 返回C:\Users\Administrator\Desktop\go_pro
	fmt.Println(filepath.Dir(`C:\Users\Administrator\Desktop\go_pro\src.html\..\`))// 返回C:\Users\Administrator\Desktop\go_pro
	fmt.Println(filepath.Dir(`C:\Users\Administrator\Desktop\go_pro\src.html\..\..\`))// 返回C:\Users\Administrator\Desktop
	fmt.Println(filepath.Dir(`C:\Users\Administrator\Desktop\go_pro\src.html\..\..\..\`))// 返回C:\Users\Administrator

	fmt.Println(filepath.Dir(`aa/bb/cc/ee/gg/hh.html`))// 返回aa\bb\cc\ee\gg
	fmt.Println(filepath.Dir(`aa/bb/cc/ee/gg/hh.html/../../`))// 返回aa\bb\cc\ee
	fmt.Println(filepath.Dir(`aa/bb/cc/ee/gg/hh.html/../`))// 返回aa\bb\cc\ee\gg
	fmt.Println(filepath.Dir(`..aa/bb/`))// 返回..aa\bb
	fmt.Println(filepath.Dir(`../aa/bb/`))// 返回..\aa\bb
	fmt.Println(filepath.Dir(`../../aa/bb/`))// 返回..\..\aa\bb
	fmt.Println(filepath.Dir(``))// 返回.
	fmt.Println(filepath.Dir(`///`))// 返回\
	fmt.Println(filepath.Dir(`/a//`))// 返回\a
	fmt.Println(filepath.Dir(`/a`))// 返回\


	fmt.Println("------------------------------")
	// Join将任意数量的路径元素连接到一条路径中，并在必要时添加一个分隔符。 加入通话清理结果； 特别是，所有空字符串都将被忽略。
	//在Windows上，且仅当第一个path元素是UNC路径时，结果才是UNC路径(\\unix网络路径)。
	fmt.Println(filepath.Join(`C:\Users\Administrator`,`a`))//返回C:\Users\Administrator\a
	fmt.Println(filepath.Join(`C:\Users\Administrator\`,`a`))//返回C:\Users\Administrator\a
	fmt.Println(filepath.Join(`C:\Users\Administrator\`,`\a`))//返回C:\Users\Administrator\a
	fmt.Println(filepath.Join(`C:\Users\Administrator`,`a.html`))//返回C:\Users\Administrator\a.html
	fmt.Println(filepath.Join(`C:\Users\Administrator`,`a.html\`))//返回C:\Users\Administrator\a.html
	fmt.Println(filepath.Join(`\C:\Users\Administrator`,`a.html\`))//返回\C:\Users\Administrator\a.html
	fmt.Println(filepath.Join(`\\C:\Users\Administrator`,`a.html\`))//返回\\C:\Users\Administrator\a.html
	fmt.Println(filepath.Join(`\\\C:\Users\Administrator`,`a.html\`))//返回\C:\Users\Administrator\a.html
	fmt.Println(filepath.Join(`C:\\Users\\Administrator`,`a.html\`))//返回C:\Users\Administrator\a.html
	fmt.Println(filepath.Join(`   C:\\Users\\Administrator`,`\\a.html\`))//返回   C:\Users\Administrator\a.html(C前面有空格)
	fmt.Println(filepath.Join(`   C:\\Users\\Administrator`,`\\   a.html\`))//返回   C:\Users\Administrator\   a.html(有空格)
	fmt.Println(filepath.Join(`   C:\\Users\\Administrator`,`   a.html\`))//返回   C:\Users\Administrator\   a.html(有空格)
	fmt.Println(filepath.Join(`   `,`   a.html\`))//返回   \   a.html(有空格)
	fmt.Println(filepath.Join(`   `,`   `))//返回   \   (有空格)
	fmt.Println(filepath.Join(`   `,`   a    `))//返回   \   a    (有空格)
	fmt.Println(filepath.Join(`a`,`bb   `))//返回a\bb   (有空格)
	fmt.Println(filepath.Join(``,`bb   `))//返回bb   (有空格)
	fmt.Println(filepath.Join(`a`,``))//返回a
	fmt.Println(filepath.Join(``,``))//返回空字符串
	fmt.Println(filepath.Join(`\`,``))//返回\
	fmt.Println(filepath.Join(`\\`,``))//返回\
	fmt.Println(filepath.Join(`\\`,`\`))//返回\
	fmt.Println(filepath.Join(`\\`,`\a`))//返回\a
	fmt.Println(filepath.Join(`\\aa`,`\a`))//返回\aa\a
	fmt.Println(filepath.Join(`\\aa\b`,`\a`))//返回\\aa\b\a
	fmt.Println(filepath.Join(`\\`,`\\`))//返回\
	fmt.Println(filepath.Join(`\\\`,`\\\`))//返回\
	fmt.Println(filepath.Join(`\\t\`,`\t\\`))//返回\
	fmt.Println(filepath.Join("\taa",`b`))//返回	aa\b(aa前面有tab空格)


	fmt.Println("------------------------------")
	// Split会在最终的Separator之后立即拆分路径，将其分为目录和文件名部分。
	//如果路径中没有分隔符，则Split返回一个空的目录，并将文件设置为path。
	//返回的值具有path = dir + file的属性。
	fmt.Println(filepath.Split(`C:\Users\Administrator`))// 返回C:\Users\    Administrator(这里 有2个值注意)
	fmt.Println(filepath.Split(`C:\Users\Administrator\a`))// 返回C:\Users\    Administrator(这里 有2个值注意)
	fmt.Println(filepath.Split(`C:\Users\Administrator\a`))// 返回C:\Users\Administrator\    a(这里 有2个值注意)
	fmt.Println(filepath.Split(`C:\Users\Administrator\a.html`))// 返回C:\Users\Administrator\    a.html(这里 有2个值注意)
	fmt.Println(filepath.Split(`C:\Users\Administrator\a.html\`))// 返回C:\Users\Administrator\a.html\ (这里 有1个值注意，另外一个是空字符串)
	fmt.Println(filepath.Split(`C:\Users\Administrator\a.html\..\`))// 返回C:\Users\Administrator\a.html\..\ (这里 有1个值注意，另外一个是空字符串)
	fmt.Println(filepath.Split(`C:\Users\Administrator\a.html\..\..\`))// 返回C:\Users\Administrator\a.html\..\..\  (这里 有1个值注意，另外一个是空字符串)
	fmt.Println(filepath.Split(`C:\Users\Administrator\  a.html\..\..\`))// 返回C:\Users\Administrator\  a.html\..\..\ (这里 有1个值注意，另外一个是空字符串)
	fmt.Println(filepath.Split(`\..\..\`))// 返回\..\..\ (这里 有1个值注意,另外一个是空字符串)
	fmt.Println(filepath.Split(`\`))// 返回\ (这里 有1个值注意,另外一个是空字符串)
	fmt.Println(filepath.Split(`\\`))// 返回\\ (这里 有1个值注意,另外一个是空字符串)
	fmt.Println(filepath.Split(`\\\`))// 返回\\\ (这里 有1个值注意,另外一个是空字符串)
	fmt.Println(filepath.Split(`\\\a`))// 返回\\\    a(这里 有2个值注意)
	fmt.Println(filepath.Split(`\\\a\`))// 返回\\\a\ (这里 有1个值注意)
	fmt.Println(filepath.Split(``))// 返回 (这里 有1个值注意)
	fmt.Println(filepath.Split(`\C:\Users\Administrator\a.html`))// 返回\C:\Users\Administrator\    a.html(这里 有2个值注意)

	fmt.Println("------------------------------")
	//FromSlash函数将path中的斜杠（'/'）替换为路径分隔符并返回替换结果，多个斜杠会替换为多个路径分隔符。也就是一律替换成为当前系统的路径分隔符"\"，而不是"/"
	fmt.Println(filepath.FromSlash(`C:\\Users\\Administrator\\a.html`))//返回C:\\Users\\Administrator\\a.html
	fmt.Println(filepath.FromSlash(`C:\\\Users\\\Administrator\\\a.html`))//返回C:\\\Users\\\Administrator\\\a.html
	fmt.Println(filepath.FromSlash(`C:///Users///Administrator///a.html`))//返回C:\\\Users\\\Administrator\\\a.html
	fmt.Println(filepath.FromSlash(`C:UsersAdministratora.html`))//返回C:UsersAdministratora.html
	fmt.Println(filepath.FromSlash(`C:/Users/Administrator/a.html`))//返回C:\Users\Administrator\a.html
	fmt.Println(filepath.FromSlash(`C:/Users/Administrator/a.html`))//返回C:\Users\Administrator\a.html
	fmt.Println(filepath.FromSlash(`C: Users Administrator a.html`))//返回C: Users Administrator a.html

	fmt.Println("------------------------------")
	//ToSlash函数将path中的路径分隔符替换为斜杠（'/'）并返回替换结果，多个路径分隔符会替换为多个斜杠。也就是一律替换成为分隔符"/"，而不是当前系统的路径分隔符"\"
	fmt.Println(filepath.ToSlash("C:/Users/Administrator/a.html"))// C:/Users/Administrator/a.html
	fmt.Println(filepath.ToSlash("C://Users//Administrator//a.html"))// C://Users//Administrator//a.html
	fmt.Println(filepath.ToSlash("C:\\Users\\Administrator\\a.html"))// C:/Users/Administrator/a.html
	fmt.Println(filepath.ToSlash(`C:\\Users\\Administrator\\a.html`))// C://Users//Administrator//a.html
	fmt.Println(filepath.ToSlash(`C:\\\Users\\\Administrator\\\a.html`))// C:///Users///Administrator///a.html


	fmt.Println("------------------------------")
	//Match returns true if name matches the shell file name pattern. The pattern syntax is:
	//
	//pattern:
	//	{ term }
	//term:
	//	'*'                                  匹配0或多个非路径分隔符的字符
	//	'?'                                  匹配1个非路径分隔符的字符
	//	'[' [ '^' ] { character-range } ']'  字符组（必须非空）
	//	c                                    匹配字符c（c != '*', '?', '\\', '['）
	//	'\\' c                               匹配字符c
	//	character-range:
	//	c           匹配字符c（c != '\\', '-', ']'）
	//	'\\' c      匹配字符c
	//	lo '-' hi   匹配区间[lo, hi]内的字符
	//	Match要求匹配整个name字符串，而不是它的一部分。只有pattern语法错误时，会返回ErrBadPattern。
	//
	//	Windows系统中，不能进行转义：'\\'被视为路径分隔符。
	fmt.Println(filepath.Match("a","a"))//true <nil>
	fmt.Println(filepath.Match("a","abc"))//false <nil>
	fmt.Println(filepath.Match("a*","abc"))//true <nil>
	fmt.Println(filepath.Match("a[b-d]","ab"))//true <nil>
	fmt.Println(filepath.Match("a[b-d]","abc"))//false <nil>
	fmt.Println(filepath.Match("a[b-d]*","abc"))//true <nil>
	fmt.Println(filepath.Match("a[b-d]*","abcbc"))//true <nil>
	fmt.Println(filepath.Match("a[b-d]*","abcbcd"))//true <nil>
	fmt.Println(filepath.Match("a[b-d]*","abcbcdefg"))//true <nil>
	fmt.Println(filepath.Match("a*[b-d]","abcbcd"))//true <nil>
	fmt.Println(filepath.Match("a*[b-d]","aeeeebcd"))//true <nil>
	fmt.Println(filepath.Match("a[^b-d]","aeeeebcd"))//false <nil>
	fmt.Println(filepath.Match("a[^b-d]","ab"))//false <nil>
	fmt.Println(filepath.Match("a[^b-d]","af"))//true <nil>
	fmt.Println(filepath.Match("a[^b-d]","aff"))//false <nil>
	fmt.Println(filepath.Match("a[^b-d]","afg"))//false <nil>
	fmt.Println(filepath.Match("a[^b-d][^b-d]","afg"))//true <nil>
	fmt.Println(filepath.Match("*","afg"))//true <nil>
	fmt.Println(filepath.Match("*","af/g"))//true <nil>
	fmt.Println(filepath.Match("*","af//g"))//true <nil>
	fmt.Println(filepath.Match("*","af\\g"))//false <nil>
	fmt.Println(filepath.Match("*",`af\\g`))//false <nil>
	fmt.Println(filepath.Match("*",`af\g`))//false <nil>
	fmt.Println(filepath.Match(`*\*`,`af\g`))//true <nil>
	fmt.Println(filepath.Match(`*\*`,`af/g`))//false <nil>
	fmt.Println(filepath.Match(`*\*`,`af//g`))//false <nil>
	fmt.Println(filepath.Match(`*\*`,`af\\g`))//false <nil>
	fmt.Println(filepath.Match(`*`,`af*g`))//true <nil>
	fmt.Println(filepath.Match(`*`,`af-g`))//true <nil>

	fmt.Println("------------------------------")
	//将PATH或GOPATH等环境变量里的多个路径分割开（这些路径被OS特定的表分隔符连接起来）。与strings.Split函数的不同之处是：对""，
	// SplitList返回[]string{}，而strings.Split返回空字符串。
	//
	//Example
	//fmt.Println("On Unix:", filepath.SplitList("/a/b/c:/usr/bin"))
	//Output:
	//
	//On Unix: [/a/b/c /usr/bin]
	//在linux中是“:”作为分隔符，在windows中是“;”作为分隔符
	fmt.Println(filepath.SplitList(`\a\b\c;\usr\bin;\cc\dd\ee`))//[\a\b\c \usr\bin \cc\dd\ee]
	fmt.Println(filepath.SplitList(`\\a\\b\\c;\\usr\\bin;\\cc\\dd\\ee`))//[\\a\\b\\c \\usr\\bin \\cc\\dd\\ee]
	fmt.Println(filepath.SplitList(`/a/b/c;/usr/bin;/cc/dd/ee`))//[/a/b/c /usr/bin /cc/dd/ee]
	fmt.Println(filepath.SplitList(`abc;usrbin;ccddee`))//[abc usrbin ccddee]
	fmt.Println(filepath.SplitList(`abc/;usrbin;ccddee`))//[abc/ usrbin ccddee]
	fmt.Println(filepath.SplitList(`abc/;usrbin;ccddee`))//[abc/ usrbin ccddee]
	fmt.Println(filepath.SplitList(``))//[]
	fmt.Println(filepath.Split(``))//返回空字符串
	fmt.Println(filepath.SplitList(`;;;`))//返回[   ]

	fmt.Println("------------------------------")

	//Clean函数通过单纯的词法操作返回和path代表同一地址的最短路径。
	//
	//它会不断的依次应用如下的规则，直到不能再进行任何处理：
	//
	//1. 将连续的多个路径分隔符替换为单个路径分隔符
	//2. 剔除每一个.路径名元素（代表当前目录）
	//3. 剔除每一个路径内的..路径名元素（代表父目录）和它前面的非..路径名元素
	//4. 剔除开始一个根路径的..路径名元素，即将路径开始处的"/.."替换为"/"（假设路径分隔符是'/'）
	//返回的路径只有其代表一个根地址时才以路径分隔符结尾，如Unix的"/"或Windows的`C:\`。
	//
	//如果处理的结果是空字符串，Clean会返回"."。
	fmt.Println(filepath.Clean("/a/b/c"))// 返回\a\b\c
	fmt.Println(filepath.Clean("a/b/c"))// 返回a\b\c
	fmt.Println(filepath.Clean(".a/b/c"))// 返回.a\b\c
	fmt.Println(filepath.Clean("..a/b/c"))// 返回..a\b\c
	fmt.Println(filepath.Clean("../a/b/c"))// 返回..\a\b\c
	fmt.Println(filepath.Clean("../../a/b/c"))// 返回..\..\a\b\c
	fmt.Println(filepath.Clean("./b/c"))// 返回b\c
	fmt.Println(filepath.Clean("../b/c"))// 返回..\b\c
	fmt.Println(filepath.Clean("a/../b/c"))// 返回b\c
	fmt.Println(filepath.Clean("a/../../b/c"))// 返回..\b\c
	fmt.Println(filepath.Clean("a/../"))// 返回.
	fmt.Println(filepath.Clean("a/.."))// 返回.
	fmt.Println(filepath.Clean("/.."))// 返回\
	fmt.Println(filepath.Clean("/"))// 返回\
	fmt.Println(filepath.Clean("/."))// 返回\
	fmt.Println(filepath.Clean("./"))// 返回.
	fmt.Println(filepath.Clean("../"))// 返回..
	fmt.Println(filepath.Clean(""))// 返回.
	fmt.Println(filepath.Clean("."))// 返回.
	fmt.Println(filepath.Clean(".."))// 返回..
	fmt.Println(filepath.Clean("////"))// 返回\
	fmt.Println(filepath.Clean(`C:\`))// 返回C:\
	fmt.Println(filepath.Clean(`C:\a\`))// 返回C:\a
	fmt.Println(filepath.Clean(`C:`))// 返回C:.
	fmt.Println(filepath.Clean(`\C:\`))// 返回\C:
	fmt.Println(filepath.Clean(`\\\C:\`))// 返回\C:


	fmt.Println("------------------------------")
	//EvalSymlinks函数返回path指向的符号链接（软链接）所包含的路径。如果path和返回值都是相对路径，
	// 会相对于当前目录；除非两个路径其中一个是绝对路径。如果不是软连接那就返回真实的路径
	fmt.Println(filepath.EvalSymlinks(`test1.txt`))//main\test.txt <nil>
	fmt.Println(filepath.EvalSymlinks(`main\test.txt`))//main\test.txt <nil>
	fmt.Println(filepath.EvalSymlinks(`main\test.txt`))//main\test.txt <nil>
	fmt.Println(filepath.EvalSymlinks(`C:\Users\Administrator\Desktop\Wireshark.lnk`))//报错不知道为什么，也许windows中的快捷方式不是软连接
	fmt.Println(filepath.EvalSymlinks(`C:\Users\Administrator\Desktop\go_pro\src\io_pro\test1.txt`))//C:\Users\Administrator\Desktop\go_pro\src\io_pro\main\test.txt <nil>
	fmt.Println(filepath.EvalSymlinks(`a/b/c`))//报错,没有这个文件目录或者文件

	fmt.Println("------------------------------")

	// Glob返回所有匹配pattern或nil的文件的名称
	//如果没有匹配的文件。 模式的语法与Match中的语法相同。 该模式可以描述分层名称，例如/ usr / * / bin / ed（假设分隔符为'/'）。
	//
	// Glob会忽略文件系统错误，例如读取目录的I / O错误。
	//格式错误时，唯一可能返回的错误是ErrBadPattern。
	//似乎只在当前的项目路径进行匹配
	fmt.Println(filepath.Glob("test1.txt"))//[test1.txt] <nil>
	fmt.Println(filepath.Glob("test?.txt"))//[test1.txt test2.txt test3.txt test4.txt] <nil>
	fmt.Println(filepath.Glob(`C:\Users\Administrator\Desktop\go_pro\src\io_pro\test1.txt`))//[C:\Users\Administrator\Desktop\go_pro\src\io_pro\test1.txt] <nil>
	fmt.Println(filepath.Glob(`C:\Users\Administrator\Desktop\go_pro\src\io_pro\test?.txt`))//[C:\Users\Administrator\Desktop\go_pro\src\io_pro\test1.txt
	// C:\Users\Administrator\Desktop\go_pro\src\io_pro\test2.txt C:\Users\Administrator\Desktop\go_pro\src\io_pro\test3.txt C:\Users\Administrator\Desktop\go_pro\src\io_pro\test4.txt] <nil>
	fmt.Println(filepath.Glob(`C:\Users\Administrator\Desktop\*.jpg`))//[C:\Users\Administrator\Desktop\001.jpg C:\Users\Administrator\Desktop\th.jpg C:\Users\Administrator\Desktop\timg.jpg] <nil>


	fmt.Println("------------------------------")
	//Rel函数返回一个相对路径，将basepath和该路径用路径分隔符连起来的新路径在词法上等价于targpath。也就是说，
	//Join(basepath, Rel(basepath, targpath))等价于targpath本身。如果成功执行，返回值总是相对于basepath的，
	//即使basepath和targpath没有共享的路径元素。如果两个参数一个是相对路径而另一个是绝对路径，或者targpath无法表示为相对于basepath的路径，将返回错误。
	paths := []string{
		"/a/b/c",
		"/b/c",
		"./b/c",
	}
	base := "/a"
	fmt.Println("On Unix:")
	for _, p := range paths {
		rel, err := filepath.Rel(base, p)
		fmt.Printf("%q: %q %v\n", p, rel, err)
	}
	// 输出：
	// On Unix:
	// "/a/b/c": "b\\c" <nil>
	// "/b/c": "..\\b\\c" <nil>
	// "./b/c": "" Rel: can't make ./b/c relative to /a

	fmt.Println("------------------------------")
	//一些常量
	fmt.Println(filepath.ListSeparator)//59
	fmt.Println(filepath.Separator)//92
	fmt.Println(filepath.ErrBadPattern)//syntax error in pattern
	fmt.Println(filepath.SkipDir)//skip this directory
	//var a rune=59
	var a rune=92
	fmt.Println(string(a))// ;和\

	fmt.Println("------------------------------")

	//WalkFunc():
	//Walk函数对每一个文件/目录都会调用WalkFunc函数类型值。调用时path参数会包含Walk的root参数作为前缀；
	//就是说，如果Walk函数的root为"dir"，该目录下有文件"a"，将会使用"dir/a"调用walkFn参数。walkFn参数
	//被调用时的info参数是path指定的地址（文件/目录）的文件信息，类型为os.FileInfo。
	//
	//如果遍历path指定的文件或目录时出现了问题，传入的参数err会描述该问题，WalkFunc类型函数可以决定如何去
	//处理该错误（Walk函数将不会深入该目录）；如果该函数返回一个错误，Walk函数的执行会中止；只有一个例外，
	//如果Walk的walkFn返回值是SkipDir，将会跳过该目录的内容而Walk函数照常执行处理下一个文件。

	//Walk():
	//Walk函数会遍历root指定的目录下的文件树，对每一个该文件树中的目录和文件都会调用walkFn，包括root自身。
	//所有访问文件/目录时遇到的错误都会传递给walkFn过滤。文件是按词法顺序遍历的，这让输出更漂亮，但也导致
	//处理非常大的目录时效率会降低。Walk函数不会遍历文件树中的符号链接（快捷方式）文件包含的路径。

	//获取fileInfo的方式，但是很明显是不需要的！
	//f,err1:=os.Open("main/test.txt")
	//if err1 != nil{
	//	fmt.Println(err1)
	//}
	//Finfo, err2 := f.Stat()
	//if err2 != nil{
	//	fmt.Println(err2)
	//}

	var wf func(path string, info os.FileInfo, err error) error
	wf=fmtprint
	//var walkF filepath.WalkFunc("test",Finfo,err1)//事实上这里我们并不需要自己创建WalkFunc类对象，因为这个filepath.WalkFunc（）
	//的基类是func类，所以我们直接创建func类即可
	walkErr := filepath.Walk("test", wf)
	if walkErr != nil{
		fmt.Println(walkErr)
	}
	fmt.Println("结束。。。。。。。。。")
	//输出如下:
	//	test=== true 4096
	//	test\test01.go=== false 541
	//	test\test02.go=== false 886
	//	test\test03.go=== false 105
	//	test\test04.go=== false 2006
	//	test\test05.go=== false 413
	//	test\test06.go=== false 923
	//	test\test07.go=== false 826
	//	test\test08.go=== false 3535
	//	test\testdir=== true 0
	//	test\testdir\1.txt=== false 0
	//	test\testdir\2.txt=== false 0
	//	结束。。。。。。。。。


}

func fmtprint(path string, info os.FileInfo, err error) error {
	fmt.Println(path+"===",info.IsDir(),info.Size())
	return nil
}





















