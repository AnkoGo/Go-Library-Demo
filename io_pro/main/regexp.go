package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	//该包下的对象：
	//// Regexp是已编译正则表达式的表示。
	//// Regexp可安全地供多个goroutine并发使用，但配置方法（例如Longest）除外。
	//type Regexp struct {
	//	expr           string       //传递给Compile（即我们输入的正则表达式字符）
	//	prog           *syntax.Prog //编译程序
	//	onepass        *onePassProg // onepass程序或nil
	//	numSubexp      int
	//	maxBitStateLen int
	//	subexpNames    []string
	//	prefix         string         //未锚定匹配中的必需前缀
	//	prefixBytes    []byte         //前缀，为[] byte
	//	prefixRune     rune           //前缀中的第一个符文
	//	prefixEnd      uint32         //前缀中最后一个符文的pc
	//	mpool          int            //匹配池子
	//	matchcap       int            //记录的匹配长度的大小
	//	prefixComplete bool           //前缀是整个正则表达式的开头
	//	cond           syntax.EmptyOp //匹配开始时需要的空白宽度条件（// EmptyOp指定零宽度断言的一种或混合形式。）
	//	minInputLen    int            //输入的最小长度（以字节为单位）
	//
	//	//此字段可以通过Longest方法修改，但其他方法则为只读。
	//	longest bool // regexp是否为最左最长的匹配
	//}
	//
	//// String返回用于编译正则表达式的源文本。
	//func (re *Regexp) String() string {
	//	return re.expr
	//}
	//
	//
	////编译程序
	////可能不属于此程序包，但现在很方便。
	//
	//// Prog是已编译的正则表达式程序。
	//type Prog struct {
	//	Inst   []Inst
	//	Start  int //起始指令的索引（存储接下来将要compile的文本的索引下标）
	//	NumCap int // re中的InstCapture实例数
	//}
	//
	//
	////执行“一次性”正则表达式。
	////可以对某些正则表达式进行分析以确定它们永远不需要回溯：保证它们在字符串上一次运行，而不必费心保存所有常规的NFA状态。
	////检测到它们并更快地执行它们。
	//
	//// onePassProg是已编译的一遍正则表达式程序。
	////除了使用onePassInst之外，它与语法.Prog相同。
	//type onePassProg struct {
	//	Inst   []onePassInst
	//	Start  int //起始指令的索引（存储接下来将要compile的文本的索引下标）
	//	NumCap int // re中的InstCapture实例数
	//}

	//// A onePassInst is a single instruction in a one-pass regular expression program.
	//// It is the same as syntax.Inst except for the new 'Next' field.
	//// onePassInst是单遍正则表达式程序中的一条指令。
	////与语法相同，除了新的“ Next”字段。
	//type onePassInst struct {
	//	syntax.Inst
	//	Next []uint32
	//}
	//
	//
	//// Inst是正则表达式程序中的一条指令。
	//type Inst struct {
	//	Op   InstOp
	//	Out  uint32 // all but InstMatch, InstFail( 除了InstMatch，InstFail之外的所有 )
	//	Arg  uint32 // InstAlt, InstAltMatch, InstCapture, InstEmptyWidth
	//	Rune []rune
	//}
	//
	//// InstOp是指令操作码。
	//type InstOp uint8

	//const (
	//	InstAlt InstOp = iota
	//	InstAltMatch
	//	InstCapture
	//	InstEmptyWidth
	//	InstMatch
	//	InstFail
	//	InstNop
	//	InstRune
	//	InstRune1
	//	InstRuneAny
	//	InstRuneAnyNotNL
	//)


	fmt.Println("-------------regexp包中的零碎函数或者对象方法-----------------")


	//通常在初始化时间编译一次表达式。
	//使用原始字符串，以避免必须引用反斜杠。
	//	^表示 匹配文本开始，标志m为真时，还匹配行首
	//	[a-z]表示 a到z之间的字符族
	//	+表示 至少匹配一次字符，越多越好（优先重复匹配字符）
	//	\[表示 不转义[，使用原始的[的意义
	//	[0-9]表示 0到9之间的字符族
	//	$表示 匹配文本结尾，标志m为真时，还匹配行尾


	//compile(expr string, mode syntax.Flags, longest bool)将解析正则表达式，如果成功，则返回可用于与文本匹配的Regexp对象。
	//当与文本进行匹配时，regexp返回一个匹配项，该匹配项尽早在输入中（最左侧）开始，并且在其中选择从回溯搜索中最先找到的匹配项。
	//这种所谓的“最左边优先”匹配与Perl，Python和其他实现使用的语义相同，尽管此包无需牺牲回溯即可实现它。
	//有关POSIX最左最长的匹配，请参见CompilePOSIX。
	//这个方法一定会返回*Regexp对象，无论找没找到匹配的字符，区别是匹配到字符的话返回的*Regexp对象与匹配不到字符情况返回的*Regexp对象信息不同

	//mode是匹配时候的附加信息，比如是否不区分大小写进行匹配，常见的匹配模式如下：
	////Flags控制解析器的行为并记录有关regexp上下文的信息。
	//type Flags uint16
	//
	//const (
	//	FoldCase      Flags = 1 << iota //不区分大小写的匹配
	//	Literal                         //将模式视为文字字符串
	//	ClassNL                         //允许[^ a-z]和[[：space：]]之类的字符类匹配换行符
	//	DotNL                           // allow "." 匹配换行符
	//	OneLine                         //将^和$视为仅在文本的开头和结尾匹配
	//	NonGreedy                       //使重复运算符默认为非贪婪
	//	PerlX                           //允许Perl扩展
	//	UnicodeGroups                   // allow \p{Han}, \P{Han} for Unicode group and negation（允许\p{Han}，\P{Han}用于Unicode组和非Unicode组）
	//	WasDollar                       // regexp OpEndText was $, not \z （正则表达式OpEndText是$，而不是\z）
	//	Simple                          // regexp contains no counted repetition（正则表达式不包含重复计数）
	//
	//	MatchNL = ClassNL | DotNL		//可以多种flags组合在一起共同起作用，表示既 allow "." 匹配换行符 也可以允许[^ a-z]和[[：space：]]之类的字符类匹配换行符
	//如果是 ClassNL & DotNL	的话，则第二个情况要在第一个情况成立时候才会进行匹配，代表进一步的意思
	//
	//	Perl        = ClassNL | OneLine | PerlX | UnicodeGroups //尽可能接近Perl语言中规定的匹配规则或者匹配模式
	//	POSIX Flags = 0                                         //POSIX语法
	//)



	//type parser struct {
	//	flags       Flags     // parse 解析模式标志
	//	stack       []*Regexp //解析表达式解析后的堆栈
	//	free        *Regexp
	//	numCap      int //看到的捕获组数
	//	wholeRegexp string	//我们输入的整个的正则字符串
	//	tmpClass    []rune //临时char类工作空间
	//}



	//// An Op is a single regular expression operator.（Op是单个的正则表达式运算符）
	//type Op uint8
	//
	//// Operators are listed in precedence order, tightest binding to weakest.
	//// Character class operators are listed simplest to most complex
	//// (OpLiteral, OpCharClass, OpAnyCharNotNL, OpAnyChar).
	////运算符按优先顺序列出，从最紧密的绑定到最弱的绑定。
	////列出的字符类运算符最简单到最复杂（OpLiteral，OpCharClass，OpAnyCharNotNL，OpAnyChar）。
	//
	//const (
	//	OpNoMatch        Op = 1 + iota // matches no strings（不匹配任何字符串）
	//	OpEmptyMatch                   // matches empty string（匹配空字符串）
	//	OpLiteral                      // matches Runes sequence（匹配符文序列）
	//	OpCharClass                    // matches Runes interpreted as range pair list（匹配解释为范围对列表的符文（像【0-9】这种））
	//	OpAnyCharNotNL                 // matches any character except newline（匹配除换行符以外的任何字符）
	//	OpAnyChar                      // matches any character（匹配任何字符）
	//	OpBeginLine                    // matches empty string at beginning of line（匹配行首的空字符串）
	//	OpEndLine                      // matches empty string at end of line（匹配行尾的空字符串）
	//	OpBeginText                    // matches empty string at beginning of text（匹配文本开头的空字符串）
	//	OpEndText                      // matches empty string at end of text（匹配文本末尾的空字符串）
	//	OpWordBoundary                 // matches word boundary `\b`（匹配单词边界`\ b`）
	//	OpNoWordBoundary               // matches word non-boundary `\B`（匹配单词非边界`\ B`）
	//	OpCapture                      // capturing subexpression with index Cap, optional name Name（使用索引Cap捕获子表达式，可选名称Name）
	//	OpStar                         // matches Sub[0] zero or more times（匹配Sub[0]零次或多次，当匹配到*时候会使用）
	//	OpPlus                         // matches Sub[0] one or more times（匹配Sub[0]1次或多次，当匹配到+时候会使用）
	//	OpQuest                        // matches Sub[0] zero or one times（匹配Sub[0]0次或1次，当匹配到？时候会使用）
	//	OpRepeat                       // matches Sub[0] at least Min times, at most Max (Max == -1 is no limit)（至少与Sub [0]匹配Min次，最大Max（Max == -1是没有限制））
	//	OpConcat                       // matches concatenation of Subs（匹配Sub的串联）
	//	OpAlternate                    // matches alternation of Subs（匹配Subs的交替）
	//)
	//const opPseudo Op = 128 // where pseudo-ops start（伪操作开始的地方）

	//其实在匹配过程中自身使用切片生成一个栈结构来维护一个正则解析后的Regexp对象的树，并且随着解析，树的枝叶越来越多


	// MustCompile类似于Compile，但是如果无法解析该表达式，则会发生恐慌。（注意跟上面的方法的区别）
	//它简化了保存已编译正则表达式的全局变量的安全初始化。
	//这个方法应该会比上面的Compile方法更加通用
	//var validID,err = regexp.Compile(`^[a-z]+\[[0-9]+\]$`)
	//check_err_regexp(err)
	var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)
	fmt.Printf("%#v\n",validID)

	// MatchString报告字符串s是否包含正则表达式re的任何匹配项。
	//底层是通过func (re *Regexp) doMatch(r io.RuneReader, b []byte, s string) bool方法实现的
	// doMatch报告r，b或s是否与正则表达式匹配。
	//	但是doMatch底层又是通过func (re *Regexp) doExecute(r io.RuneReader, b []byte, s string, pos int, ncap int, dstCap []int) []int 来实现的。
	// doExecute在输入中找到最左边的匹配项，将其子表达式的位置附加到dstCap并返回dstCap。
	//如果未找到匹配项，则返回nil；如果找到匹配项，则返回非nil。

	//对于记录遍历我们的输入字符时候的状态信息的接口：
	//// input abstracts different representations of the input text. It provides
	//// one-character lookahead.
	//type input interface {
	//	step(pos int) (r rune, width int) // advance one rune（取用下一个rune）
	//	canCheckPrefix() bool             // can we look ahead without losing info?（我们可以在不丢失信息的情况下向前看吗？）
	//	hasPrefix(re *Regexp) bool			//判断是否有前缀
	//	index(re *Regexp, pos int) int		//当前取用到的元素的索引值
	//	context(pos int) lazyFlag			//
	//}


	//// A lazyFlag is a lazily-evaluated syntax.EmptyOp,
	//// for checking zero-width flags like ^ $ \A \z \B \b.
	//// It records the pair of relevant runes and does not
	//// determine the implied flags until absolutely necessary
	//// (most of the time, that means never).
	//// lazyFlag是惰性计算的语法.EmptyOp，用于检查零宽度标志，例如^ $ \ A \ z \ B \ b。
	////它记录一对相关的符文，直到绝对必要时才确定隐含的标志（大多数情况下，从不）。
	//type lazyFlag uint64

	//// An EmptyOp specifies a kind or mixture of zero-width assertions.
	//// EmptyOp指定一种或一种多个类的混合类的零宽度断言。
	//type EmptyOp uint8
	//
	//const (
	//	EmptyBeginLine EmptyOp = 1 << iota	//空开始字符行
	//	EmptyEndLine						//空结束字符行
	//	EmptyBeginText						//空开始文本
	//	EmptyEndText						//空结束文本
	//	EmptyWordBoundary					//空的字边界
	//	EmptyNoWordBoundary					//空的无字边界
	//)





	fmt.Println(validID.MatchString("adam[23]"))
	fmt.Println(validID.MatchString("eve[7]"))
	fmt.Println(validID.MatchString("Job[48]"))
	fmt.Println(validID.MatchString("snakey"))
	//输出：
	//	&regexp.Regexp{expr:"^[a-z]+\\[[0-9]+\\]$",
	//					prog:(*syntax.Prog)(0xc00005a390),
	//					onepass:(*regexp.onePassProg)(0xc00005a3f0),
	//					numSubexp:0,
	//					maxBitStateLen:0,
	//					subexpNames:[]string{""},
	//					prefix:"",
	//					prefixBytes:[]uint8(nil),
	//					prefixRune:0,
	//					prefixEnd:0x1,
	//					mpool:0,
	//					matchcap:2,
	//					prefixComplete:false,
	//					cond:0x4,
	//					minInputLen:4,
	//					longest:false}
	//	true
	//	true
	//	false
	//	false
	fmt.Println("-----------------------下面查看regexp对象的方法和属性------------------------------")

	fmt.Printf("%#v\n",validID)
	// String返回用于编译正则表达式的源文本。
	fmt.Printf("%#v\n",validID.String())
	//Copy返回从re复制的新Regexp对象。
	//在一个副本上调用Longest不会影响另一副本。
	//
	//不推荐使用：在较早的版本中，当在多个goroutine中使用Regexp时，为每个goroutine提供自己的副本有助于避免锁争用。
	//从Go 1.12开始，不再需要使用Copy来避免锁争用。
	//如果使用复制的原因是要使用不同的Longest设置制作两份副本，则复制可能仍然合适。
	fmt.Printf("%#v\n",validID.Copy())
	//将片段s分割成由表达式分隔的子字符串，并返回这些表达式匹配之间的子字符串的片段。
	//此方法返回的切片由FindAllString返回的切片中未包含的s的所有子字符串组成。 在不包含元字符的表达式上调用时，它等效于strings.SplitN。
	//示例：
	// s := regexp.MustCompile("a*").Split("abaabaccadaaae", 5)
	//   // s: ["", "b", "b", "c", "cadaaae"]
	//计数确定要返回的子字符串数：
	// n> 0：最多n个子字符串； 最后一个子字符串将是未拆分的余数。
	// n == 0：结果为nil（零子字符串）
	// n <0：所有子字符串
	s := regexp.MustCompile("a*").Split("abaabaccadaaae", 5)
	fmt.Printf("%#v\n",s)
	//输出：
	//	&regexp.Regexp{expr:"^[a-z]+\\[[0-9]+\\]$", prog:(*syntax.Prog)(0xc00006a390), onepass:(*regexp.onePassProg)(0xc00006a3f0), numSubexp:0, maxBitStateLen:0, subexpNames:[]string{""}, prefix:"", prefixBytes:[]uint8(nil), prefixRune:0, prefixEnd:0x1, mpool:0, matchcap:2, prefixComplete:false, cond:0x4, minInputLen:4, longest:false}
	//	"^[a-z]+\\[[0-9]+\\]$"
	//	&regexp.Regexp{expr:"^[a-z]+\\[[0-9]+\\]$", prog:(*syntax.Prog)(0xc00006a390), onepass:(*regexp.onePassProg)(0xc00006a3f0), numSubexp:0, maxBitStateLen:0, subexpNames:[]string{""}, prefix:"", prefixBytes:[]uint8(nil), prefixRune:0, prefixEnd:0x1, mpool:0, matchcap:2, prefixComplete:false, cond:0x4, minInputLen:4, longest:false}
	//	[]string{"", "b", "b", "c", "cadaaae"}


	fmt.Println()
	// Match reports whether the byte slice b
	// contains any match of the regular expression re.
	// Match报告字节切片b是否包含正则表达式re的任何匹配项。
	s1 := regexp.MustCompile("a+").Match([]byte("abaabaccadaaae"))
	fmt.Printf("%#v\n",s1)
	s1 = regexp.MustCompile("a+").Match([]byte{'b','c'})
	fmt.Printf("%#v\n",s1)
	//输出：
	//	true
	//	false

	fmt.Println()
	s1 = regexp.MustCompile("a+").MatchString("abaabaccadaaae")
	fmt.Printf("%#v\n",s1)
	s1 = regexp.MustCompile("a+").MatchString("bc")
	fmt.Printf("%#v\n",s1)
	//输出：
	//	true
	//	false

	fmt.Println()
	// RuneReader是包装ReadRune方法的接口。
	// ReadRune读取单个UTF-8编码的Unicode字符，并返回符文及其大小（以字节为单位）。 如果没有可用字符，将设置err。

	// MatchReader报告RuneReader返回的文本是否包含正则表达式re的任何匹配项。
	bf:=bytes.NewBuffer([]byte("abaabaccadaaae"))
	bf1:=bytes.NewBuffer([]byte("bc"))
	s1 = regexp.MustCompile("a+").MatchReader(bf)
	fmt.Printf("%#v\n",s1)
	s1 = regexp.MustCompile("a+").MatchReader(bf1)
	fmt.Printf("%#v\n",s1)
	//输出：
	//	true
	//	false


	fmt.Println()
	//Find返回一个切片，其中包含正则表达式b中最左边匹配的文本。
	//返回值nil表示不匹配成功。
	by1 := regexp.MustCompile("a+").Find([]byte("abaabaccadaaae"))
	fmt.Printf("%v\n",by1)
	by1 = regexp.MustCompile("a+").Find([]byte("bc"))
	fmt.Printf("%v\n",by1)
	//输出：
	//	[97]
	//	[]



	fmt.Println()
	// FindAll是Find的“all”版本； 它返回表达式的所有连续匹配的[]byte切片，如程序包注释中的“all”描述所定义。
	//返回值nil表示不匹配。
	//如果'All'出现了，该方法会返回输入中所有互不重叠的匹配结果。如果一个匹配结果的前后（没有间隔字符）存在长度为0的成功匹配，
	// 该空匹配会被忽略。包含All的方法会要求一个额外的整数参数n，如果n>=0，方法会返回最多前n个匹配结果。
	by2 := regexp.MustCompile("a+").FindAll([]byte("abaabaccadaaae"),0)
	fmt.Printf("%v\n",by2)
	by2 = regexp.MustCompile("a+").FindAll([]byte("abaabaccadaaae"),1)
	fmt.Printf("%v\n",by2)
	by2 = regexp.MustCompile("a+").FindAll([]byte("abaabaccadaaae"),2)
	fmt.Printf("%v\n",by2)
	by2 = regexp.MustCompile("a+").FindAll([]byte("bc"),2)
	fmt.Printf("%v\n",by2)
	//输出：
	//	[]
	//	[[97]]
	//	[[97] [97 97]]
	//	[]


	fmt.Println()
	// FindIndex返回一个由两个元素组成的整数切片，该切片定义正则表达式b中最左边的匹配项的位置。 匹配项本身位于b[loc[0]：loc[1]](不包括loc[1])。
	//返回值nil表示不匹配。
	//如果'All'出现了，该方法会返回输入中所有互不重叠的匹配结果。如果一个匹配结果的前后（没有间隔字符）存在长度为0的成功匹配，
	// 该空匹配会被忽略。包含All的方法会要求一个额外的整数参数n，如果n>=0，方法会返回最多前n个匹配结果。
	by3 := regexp.MustCompile("a+").FindIndex([]byte("abaabaccadaaae"))
	fmt.Printf("%v\n",by3)
	by3 = regexp.MustCompile("a+").FindIndex([]byte("bc"))
	fmt.Printf("%v\n",by3)
	//输出：
	//	[0 1]
	//	[]



	fmt.Println()
	// FindAll是Find的“all”版本； 它返回表达式的所有连续匹配的一部分，如程序包注释中的“all”描述所定义。
	//返回值nil表示不匹配。
	//如果'All'出现了，该方法会返回输入中所有互不重叠的匹配结果。如果一个匹配结果的前后（没有间隔字符）存在长度为0的成功匹配，
	// 该空匹配会被忽略。包含All的方法会要求一个额外的整数参数n，如果n>=0，方法会返回最多前n个匹配结果。
	by31 := regexp.MustCompile("a+").FindAllIndex([]byte("abaabaccadaaae"),0)
	fmt.Printf("%v\n",by31)
	by31 = regexp.MustCompile("a+").FindAllIndex([]byte("abaabaccadaaae"),1)
	fmt.Printf("%v\n",by31)
	by31 = regexp.MustCompile("a+").FindAllIndex([]byte("abaabaccadaaae"),2)
	fmt.Printf("%v\n",by31)
	by31 = regexp.MustCompile("a+").FindAllIndex([]byte("bc"),2)
	fmt.Printf("%v\n",by31)
	//输出：
	//	[]
	//	[[0 1]]
	//	[[0 1] [2 4]]
	//	[]


	fmt.Println()
	// FindAll是Find的“all”版本； 它返回表达式的所有连续匹配的[]int切片，如程序包注释中的“all”描述所定义。
	//返回值nil表示不匹配。
	//如果'All'出现了，该方法会返回输入中所有互不重叠的匹配结果。如果一个匹配结果的前后（没有间隔字符）存在长度为0的成功匹配，
	// 该空匹配会被忽略。包含All的方法会要求一个额外的整数参数n，如果n>=0，方法会返回最多前n个匹配结果。

	//下面我们通过compile返回的regexp对象看下跟上面有什么区别
	compile, e := regexp.Compile("a+")
	check_err_regexp(e)
	by32 := compile.FindAllIndex([]byte("abaabaccadaaae"),0)
	fmt.Printf("%v\n",by32)
	by32 = compile.FindAllIndex([]byte("abaabaccadaaae"),1)
	fmt.Printf("%v\n",by32)
	by32 = compile.FindAllIndex([]byte("abaabaccadaaae"),2)
	fmt.Printf("%v\n",by32)
	by32 = compile.FindAllIndex([]byte("bc"),2)
	fmt.Printf("%v\n",by32)
	//输出：
	//	[]
	//	[[0 1]]
	//	[[0 1] [2 4]]
	//	[]
	//似乎没什么区别


	fmt.Println()
	// FindAllString是FindString的“全部”版本； 它返回表达式的所有连续匹配的字符串切片，如程序包注释中的“全部”描述所定义。
	//返回值nil表示不匹配。
	//如果'All'出现了，该方法会返回输入中所有互不重叠的匹配结果。如果一个匹配结果的前后（没有间隔字符）存在长度为0的成功匹配，
	// 该空匹配会被忽略。包含All的方法会要求一个额外的整数参数n，如果n>=0，方法会返回最多前n个匹配结果。
	by33 := regexp.MustCompile("a+").FindAllString("abaabaccadaaae",0)
	fmt.Printf("%v\n",by33)
	by33 = regexp.MustCompile("a+").FindAllString("abaabaccadaaae",1)
	fmt.Printf("%v\n",by33)
	by33 = regexp.MustCompile("a+").FindAllString("abaabaccadaaae",2)
	fmt.Printf("%v\n",by33)
	by33 = regexp.MustCompile("a+").FindAllString("bc",2)
	fmt.Printf("%v\n",by33)
	//输出：
	//	[]
	//	[a]
	//	[a aa]
	//	[]


	fmt.Println()
	// FindAllString是FindString的“全部”版本； 它返回表达式的所有连续匹配的字符串切片，如程序包注释中的“全部”描述所定义。
	//返回值nil表示不匹配。
	//如果'All'出现了，该方法会返回输入中所有互不重叠的匹配结果。如果一个匹配结果的前后（没有间隔字符）存在长度为0的成功匹配，
	// 该空匹配会被忽略。包含All的方法会要求一个额外的整数参数n，如果n>=0，方法会返回最多前n个匹配结果。
	by4 := regexp.MustCompile("a+").FindAllStringIndex("abaabaccadaaae",0)
	fmt.Printf("%v\n",by4)
	by4 = regexp.MustCompile("a+").FindAllStringIndex("abaabaccadaaae",1)
	fmt.Printf("%v\n",by4)
	by4 = regexp.MustCompile("a+").FindAllStringIndex("abaabaccadaaae",2)
	fmt.Printf("%v\n",by4)
	by4 = regexp.MustCompile("a+").FindAllStringIndex("bc",2)
	fmt.Printf("%v\n",by4)
	//输出：
	//	[]
	//	[[0 1]]
	//	[[0 1] [2 4]]
	//	[]


	fmt.Println()
	// FindSubmatch返回一个切片的切片(显然长度最小为2)，该切片包含b中正则表达式最左侧匹配的文本以及其子表达式的匹配项（如果有），这由包注释中的“ Submatch”描述定义。
	//返回值nil表示不匹配。

	//如果'Submatch'出现了，返回值是表示正则表达式中成功的组匹配（子匹配/次级匹配）的切片。组匹配是正则表达式内部的括号包围的次级表达式（也被称为“捕获分组”），
	//从左到右按左括号的顺序编号。，索引0的组匹配为完整表达式的匹配结果，1为第一个分组的匹配结果，依次类推。
	//如果'Index'出现了，匹配/分组匹配会用输入流的字节索引对表示result[2*n:2*n+1]表示第n个分组匹配的的匹配结果。如果没有'Index'，匹配结果表示为匹配到的文本。
	//如果索引为负数，表示分组匹配没有匹配到输入流中的文本。
	by5 := regexp.MustCompile("a+(\\d+)").FindSubmatch([]byte("ccaa1234aa1234"))
	fmt.Printf("%v\n",by5)
	by5 = regexp.MustCompile("a+(\\d+)").FindSubmatch([]byte("ccbc1234bc1234"))
	fmt.Printf("%v\n",by5)
	by5 = regexp.MustCompile("a+(\\d+)(a+)").FindSubmatch([]byte("ccaa1234aa1234"))
	fmt.Printf("%v\n",by5)
	by5 = regexp.MustCompile("a+").FindSubmatch([]byte("ccaa1234aa1234"))
	fmt.Printf("%v\n",by5)
	//输出：
	//	[[97 97 49 50 51 52] [49 50 51 52]]
	//	[]
	//	[[97 97 49 50 51 52 97 97] [49 50 51 52] [97 97]]
	//	[[97 97]]


	fmt.Println()
	// FindSubmatchIndex返回一个切片(显然长度为2)，该切片包含索引对，该索引对标识b中正则表达式的最左匹配及其子表达式的匹配项（如果有），这由包注释中的“ Submatch”和“ Index”描述定义。
	//返回值nil表示不匹配。

	//如果'Submatch'出现了，返回值是表示正则表达式中成功的组匹配（子匹配/次级匹配）的切片。组匹配是正则表达式内部的括号包围的次级表达式（也被称为“捕获分组”），
	//从左到右按左括号的顺序编号。，索引0的组匹配为完整表达式的匹配结果，1为第一个分组的匹配结果，依次类推。
	//如果'Index'出现了，匹配/分组匹配会用输入流的字节索引对表示result[2*n:2*n+1]表示第n个分组匹配的的匹配结果。如果没有'Index'，匹配结果表示为匹配到的文本。
	//如果索引为负数，表示分组匹配没有匹配到输入流中的文本。
	by6 := regexp.MustCompile("a+(\\d+)").FindSubmatchIndex([]byte("ccaa1234aa1234"))
	fmt.Printf("%v\n",by6)
	by6 = regexp.MustCompile("a+(\\d+)").FindSubmatchIndex([]byte("ccbc1234bc1234"))
	fmt.Printf("%v\n",by6)
	by6 = regexp.MustCompile("a+(\\d+)(a+)").FindSubmatchIndex([]byte("ccaa1234aa1234"))
	fmt.Printf("%v\n",by6)
	by6 = regexp.MustCompile("a+").FindSubmatchIndex([]byte("ccaa1234aa1234"))
	fmt.Printf("%v\n",by6)
	//输出：
	//	[2 8 4 8]
	//	[]
	//	[2 10 4 8 8 10]
	//	[2 4]


	fmt.Println()
	// FindAllSubmatchIndex是FindSubmatchIndex的“all”版本； 它返回表达式的所有连续匹配的一部分，如程序包注释中的“all”描述所定义。
	//返回值nil表示不匹配。

	by7 := regexp.MustCompile("a+(\\d+)").FindAllSubmatchIndex([]byte("ccaa1234aa1234"),0)
	fmt.Printf("%v\n",by7)
	by7 = regexp.MustCompile("a+(\\d+)").FindAllSubmatchIndex([]byte("ccaa1234aa1234"),1)
	fmt.Printf("%v\n",by7)
	by7 = regexp.MustCompile("a+(\\d+)").FindAllSubmatchIndex([]byte("ccaa1234aa1234"),2)
	fmt.Printf("%v\n",by7)
	by7 = regexp.MustCompile("a+(\\d+)").FindAllSubmatchIndex([]byte("ccaa1234aa1234"),3)//特地超过次数
	fmt.Printf("%v\n",by7)

	by7 = regexp.MustCompile("a+(\\d+)").FindAllSubmatchIndex([]byte("ccaa1234aa1234"),2)
	fmt.Printf("%v\n",by7)
	by7 = regexp.MustCompile("a+(\\d+)(a+)").FindAllSubmatchIndex([]byte("ccaa1234aa1234"),2)
	fmt.Printf("%v\n",by7)
	by7 = regexp.MustCompile("a+").FindAllSubmatchIndex([]byte("ccaa1234aa1234"),2)
	fmt.Printf("%v\n",by7)
	//输出：
	//	[]
	//	[[2 8 4 8]]
	//	[[2 8 4 8] [8 14 10 14]]
	//	[[2 8 4 8] [8 14 10 14]]
	//	[[2 8 4 8] [8 14 10 14]]
	//	[[2 10 4 8 8 10]]
	//	[[2 4] [8 10]]



	fmt.Println()
	// FindAllSubmatchIndex是FindSubmatchIndex的“all”版本； 它返回表达式的所有连续匹配的一部分，如程序包注释中的“all”描述所定义。
	//返回值nil表示不匹配。

	by8 := regexp.MustCompile("a+(\\d+)").FindAllSubmatch([]byte("ccaa1234aa1234"),0)
	fmt.Printf("%v\n",by8)
	by8 = regexp.MustCompile("a+(\\d+)").FindAllSubmatch([]byte("ccaa1234aa1234"),1)
	fmt.Printf("%v\n",by8)
	by8 = regexp.MustCompile("a+(\\d+)").FindAllSubmatch([]byte("ccaa1234aa1234"),2)
	fmt.Printf("%v\n",by8)
	by8 = regexp.MustCompile("a+(\\d+)").FindAllSubmatch([]byte("ccaa1234aa1234"),3)//特地超过次数
	fmt.Printf("%v\n",by8)

	by8 = regexp.MustCompile("a+(\\d+)").FindAllSubmatch([]byte("ccaa1234aa1234"),2)
	fmt.Printf("%v\n",by8)
	by8 = regexp.MustCompile("a+(\\d+)(a+)").FindAllSubmatch([]byte("ccaa1234aa1234"),2)
	fmt.Printf("%v\n",by8)
	by8 = regexp.MustCompile("a+").FindAllSubmatch([]byte("ccaa1234aa1234"),2)
	fmt.Printf("%v\n",by8)
	//输出：
	//	[]
	//	[[[97 97 49 50 51 52] [49 50 51 52]]]
	//	[[[97 97 49 50 51 52] [49 50 51 52]] [[97 97 49 50 51 52] [49 50 51 52]]]
	//	[[[97 97 49 50 51 52] [49 50 51 52]] [[97 97 49 50 51 52] [49 50 51 52]]]
	//	[[[97 97 49 50 51 52] [49 50 51 52]] [[97 97 49 50 51 52] [49 50 51 52]]]
	//	[[[97 97 49 50 51 52 97 97] [49 50 51 52] [97 97]]]
	//	[[[97 97]] [[97 97]]]




	fmt.Println()
	// FindAllStringSubmatch是FindStringSubmatch的“全部”版本； 它返回表达式的所有连续匹配的一部分，如程序包注释中的“全部”描述所定义。
	//返回值nil表示不匹配。

	by9 := regexp.MustCompile("a+(\\d+)").FindAllStringSubmatch("ccaa1234aa1234",0)
	fmt.Printf("%v\n",by9)
	by9 = regexp.MustCompile("a+(\\d+)").FindAllStringSubmatch("ccaa1234aa1234",1)
	fmt.Printf("%v\n",by9)
	by9 = regexp.MustCompile("a+(\\d+)").FindAllStringSubmatch("ccaa1234aa1234",2)
	fmt.Printf("%v\n",by9)
	by9 = regexp.MustCompile("a+(\\d+)").FindAllStringSubmatch("ccaa1234aa1234",3)//特地超过次数
	fmt.Printf("%v\n",by9)

	by9 = regexp.MustCompile("a+(\\d+)").FindAllStringSubmatch("ccaa1234aa1234",2)
	fmt.Printf("%v\n",by9)
	by9 = regexp.MustCompile("a+(\\d+)(a+)").FindAllStringSubmatch("ccaa1234aa1234",2)
	fmt.Printf("%v\n",by9)
	by9 = regexp.MustCompile("a+").FindAllStringSubmatch("ccaa1234aa1234",2)
	fmt.Printf("%v\n",by9)
	//输出：
	//	[]
	//	[[aa1234 1234]]
	//	[[aa1234 1234] [aa1234 1234]]
	//	[[aa1234 1234] [aa1234 1234]]
	//	[[aa1234 1234] [aa1234 1234]]
	//	[[aa1234aa 1234 aa]]
	//	[[aa] [aa]]



	fmt.Println()
	// FindStringSubmatch返回一个字符串切片，其中包含s中正则表达式最左侧匹配的文本以及其子表达式的匹配项（如果有），这由包注释中的“ Submatch”描述定义。
	//返回值nil表示不匹配。

	by91 := regexp.MustCompile("a+(\\d+)").FindStringSubmatch("ccaa1234aa1234")
	fmt.Printf("%v\n",by91)
	by91 = regexp.MustCompile("a+(\\d+)").FindStringSubmatch("ccaa1234aa1234")
	fmt.Printf("%v\n",by91)
	by91 = regexp.MustCompile("a+(\\d+)").FindStringSubmatch("ccaa1234aa1234")
	fmt.Printf("%v\n",by91)
	by91 = regexp.MustCompile("a+(\\d+)").FindStringSubmatch("ccaa1234aa1234")
	fmt.Printf("%v\n",by91)

	by91 = regexp.MustCompile("a+(\\d+)").FindStringSubmatch("ccaa1234aa1234")
	fmt.Printf("%v\n",by91)
	by91 = regexp.MustCompile("a+(\\d+)(a+)").FindStringSubmatch("ccaa1234aa1234")
	fmt.Printf("%v\n",by91)
	by91 = regexp.MustCompile("a+").FindStringSubmatch("ccaa1234aa1234")
	fmt.Printf("%v\n",by91)
	//输出：
	//	[aa1234 1234]
	//	[aa1234 1234]
	//	[aa1234 1234]
	//	[aa1234 1234]
	//	[aa1234 1234]
	//	[aa1234aa 1234 aa]
	//	[aa]


	fmt.Println()
	// FindStringSubmatch返回一个字符串切片，其中包含s中正则表达式最左侧匹配的文本以及其子表达式的匹配项（如果有），这由包注释中的“ Submatch”描述定义。
	//返回值nil表示不匹配。
	bf2:=bytes.NewBuffer([]byte("ccaa1234aa1234"))
	by92 := regexp.MustCompile("a+(\\d+)").FindReaderSubmatchIndex(bf2)
	fmt.Printf("%v\n",by92)
	bf2=bytes.NewBuffer([]byte("ccaa1234aa1234"))
	by92 = regexp.MustCompile("a+(\\d+)(a+)").FindReaderSubmatchIndex(bf2)
	fmt.Printf("%v\n",by92)
	bf2=bytes.NewBuffer([]byte("ccaa1234aa1234"))
	by92 = regexp.MustCompile("a+").FindReaderSubmatchIndex(bf2)
	fmt.Printf("%v\n",by92)
	//输出：
	//	[2 8 4 8]
	//	[2 10 4 8 8 10]
	//	[2 4]

	fmt.Println()
	//匹配然后填充模板（跟写web时候的填充差不多），具体实现看函数里面
	ExampleRegexp_Expand()
	ExampleRegexp_ExpandString()


	fmt.Println()
	// ReplaceAll返回src的副本，用替换文本repl替换Regexp的匹配项。 在repl内部，$符号的解释方式与Expand相同，
	// 例如$1代表第一个子匹配项的文本。
	//底层其实也是调用了re.expand(Expand()方法也是调用了这个方法)
	by_src:=[]byte("ccaa1234aa1234")
	by_repl:=[]byte("kk66")
	by_ret := regexp.MustCompile("a+(\\d+)").ReplaceAll(by_src,by_repl)
	fmt.Printf("%s---%p\n",by_ret,by_ret)
	fmt.Printf("%s---%p\n",by_src,by_src)
	fmt.Printf("%s---%p\n",by_repl,by_repl)
	//输出：
	//	cckk66kk66---0xc000063090
	//	ccaa1234aa1234---0xc000063040
	//	kk66---0xc000063050
	//从这里看出其实和Expand()方法没什么区别的，主要是少了一个模板，但是灵活性的话肯定是Expand()方法更加灵活，因为可以自定义模板
	//地址不同，说明了返回值是一个新生成的[]byte类型实例对象
	//下面是探究如果替换字符串是空的情况：
	fmt.Println()
	by_rep2:=[]byte("")
	by_ret = regexp.MustCompile("a+(\\d+)").ReplaceAll(by_src,by_rep2)
	fmt.Printf("%s---%p\n",by_ret,by_ret)
	fmt.Printf("%s---%p\n",by_src,by_src)
	fmt.Printf("%s---%p\n",by_rep2,by_rep2)
	//输出：
	//	cc---0xc0000670a8
	//	ccaa1234aa1234---0xc000067040
	//	---0x5dc1f8



	fmt.Println()
	//跟上面的api功能差不多
	// ReplaceAllString返回src的副本，用替换字符串repl替换Regexp的匹配项。 在repl内部，$符号的解释方式与Expand相同，
	// 例如$ 1代表第一个子匹配项的文本。
	str_src :="ccaa1234aa1234"
	str_repl :="kk66"
	str_ret := regexp.MustCompile("a+(\\d+)").ReplaceAllString(str_src,str_repl)
	fmt.Printf("%s---%p\n",str_ret,&str_ret)
	fmt.Printf("%s---%p\n",str_src,&str_src)
	fmt.Printf("%s---%p\n",str_repl,&str_repl)
	//输出：
	//	cckk66kk66---0xc000034ca0
	//	ccaa1234aa1234---0xc000034c80
	//	kk66---0xc000034c90




	fmt.Println()
	// ReplaceAllFunc返回src的副本，其中Regexp的所有匹配项都已替换为应用于匹配的字节片的函数repl的返回值。
	// 由repl返回的替换将直接替换，而无需使用Expand。（说明底层并不通过expand来实现的！）
	by_src=[]byte("ccaa1234aa1234")
	by_repl=[]byte("kk66")
	by_ret = regexp.MustCompile("a+(\\d+)").ReplaceAllFunc(by_src, func(i []byte) []byte {
		fmt.Println("===",string(i))
		return by_repl
	})
	fmt.Printf("%s---%p\n",by_ret,by_ret)
	fmt.Printf("%s---%p\n",by_src,by_src)
	fmt.Printf("%s---%p\n",by_repl,by_repl)
	//输出：
	//	=== aa1234
	//	=== aa1234
	//	cckk66kk66---0xc00000b1a0
	//	ccaa1234aa1234---0xc00000b150
	//	kk66---0xc00000b128

	//事实上他是边匹配边替换的而不是一次性匹配好然后再全部替换！下面说明这点：
	fmt.Println()
	by_src=[]byte("ccaa1234aa1234")
	by_repl=[]byte("kk66")
	by_src_sub:=by_src[2:8]

	by_ret = regexp.MustCompile("a+(\\d+)").ReplaceAllFunc(by_src, func(i []byte) []byte {
		fmt.Println("===",string(i))
		fmt.Printf("%s---%p\n",i,i)
		i=append(i,[]byte("77")...)
		return i
		//或者下面这样写都一样的结果
		//return append(i,[]byte("77")...) //字节只能append，不能用+号来连接字节，只能用来连接字符串的字符，这里的append不用接收值
	})
	fmt.Printf("%s---%p\n",by_src_sub,by_src_sub)
	fmt.Printf("%s---%p\n",by_ret,by_ret)
	fmt.Printf("%s---%p\n",by_src,by_src)
	fmt.Printf("%s---%p\n",by_repl,by_repl)
	//输出：
	//	=== aa1234
	//	aa1234---0xc00000b1b2
	//	aa1234---0xc00000b1b2,这个地址和上面的额这个地址相同则说明了func(i []byte) []byte中的i是和by_src是引用的同一个地址，
	//	ccaa123477771234---0xc00000b1f0
	//	ccaa1234771234---0xc00000b1b0
	//	kk66---0xc00000b198
	//因为func(i []byte) []byte中的i是和by_src是引用的同一个地址，所以在函数func(i []byte) []byte中被append了一次的“77”，
	// 之后又在外层的ReplaceAllFunc（）方法中再次append了一次！所以才会输出结果“ccaa123477771234”



	//ReplaceAllStringFunc方法和上面的这个方法几乎差不多，不再累叙！



	fmt.Println()
	by_src=[]byte("ccaa1234aa1234")
	by_repl=[]byte("kk66")

	// ReplaceAllLiteral返回src的副本，用替换字节repl替换Regexp的匹配项。 repl直接替换对应位置，而无需使用Expand。
	//其实和.ReplaceAll（）差不多，区别在于是否直接替换相应的匹配字节！
	by_ret = regexp.MustCompile("a+(\\d+)").ReplaceAllLiteral(by_src, by_repl)

	fmt.Printf("%s---%p\n",by_ret,by_ret)
	fmt.Printf("%s---%p\n",by_src,by_src)
	fmt.Printf("%s---%p\n",by_repl,by_repl)
	//输出：
	//	cckk66kk66---0xc000063200
	//	ccaa1234aa1234---0xc0000631c0
	//	kk66---0xc0000631d0

	//ReplaceAllLiteralString和ReplaceAllLiteral方法差不多，不再累叙,Literal文本的意思





	fmt.Println()

	// LiteralPrefix返回一个文字字符串，该字符串必须以正则表达式re的任何匹配开头。 如果文字字符串包含整个正则表达式，则返回布尔值true。
	//说白了就是输出正则表达式要匹配的字符串必须以什么字符串开头，其实底层输出的是这个regexp对象的一些字段信息（.prefix, .prefixComplete这2个字段）
	fmt.Println(regexp.MustCompile("a+(\\d+)").LiteralPrefix())
	fmt.Println(regexp.MustCompile("^ab+(\\d+)cd$").LiteralPrefix())
	fmt.Println(regexp.MustCompile("^.+(\\d+)cd$").LiteralPrefix())
	fmt.Println(regexp.MustCompile("abcd").LiteralPrefix())
	fmt.Println(regexp.MustCompile("abcd$").LiteralPrefix())
	fmt.Println(regexp.MustCompile("[xyz]").LiteralPrefix())
	//输出：
	//	a false
	//	ab false
	//	false
	//	abcd true
	//	abcd false
	//	false



	fmt.Println()

	//Longest，使以后的搜索更喜欢最长的匹配。
	//也就是说，当与文本进行匹配时，regexp返回一个匹配项，该匹配项尽早在输入中（最左侧）开始，并且在其中选择一个尽可能长的匹配项。
	//此方法修改了Regexp，不得与其他任何方法同时调用。
	re:=regexp.MustCompile("[a-z]+?")//?表示尽可能少的进行匹配，一但我们在下面进行了re.Longest()的设置的话，那么我们这个?号将相当于没有了额！
	ret :=re.FindAllStringSubmatch("ccaa1234aa1234",-1)
	fmt.Println(ret)

	re.Longest()//设置这个额相当于把MustCompile（没有最左最长）写成了 MustCompilePOSIX（有最左最长），但是还是有一点区别的，在下面会介绍到
	ret =re.FindAllStringSubmatch("ccaa1234aa1234",-1)
	fmt.Println(ret)
	//输出：
	//	[[c] [c] [a] [a] [a] [a]]
	//	[[ccaa] [aa]]




	fmt.Println()

	// NumSubexp返回此Regexp中带括号的子表达式的数量。
	re=regexp.MustCompile("[a-z]+?")
	fmt.Println(re.NumSubexp())
	re=regexp.MustCompile("([a-z]+?)")
	fmt.Println(re.NumSubexp())
	re=regexp.MustCompile("(a+)(\\d+)")
	fmt.Println(re.NumSubexp())
	re=regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)
	fmt.Println(re.NumSubexp())

	//输出：
	//	0
	//	1
	//	2
	//	2



	fmt.Println()

	// SubexpNames返回此Regexp中带括号的子表达式的名称。 第一个子表达式的名称为names [1]，因此，如果m是匹配片，则m [i]的名称为SubexpNames（）[i]。
	//由于不能对整个Regexp进行命名，所以names [0]始终为空字符串。 切片不得修改。
	re=regexp.MustCompile("[a-z]+?")
	fmt.Printf("%#v\n",re.SubexpNames())
	re=regexp.MustCompile("([a-z]+?)")
	fmt.Printf("%#v\n",re.SubexpNames())
	re=regexp.MustCompile("(a+)(\\d+)")
	fmt.Printf("%#v\n",re.SubexpNames())
	re=regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)
	fmt.Printf("%#v\n",re.SubexpNames())
	//输出：
	//	[]string{""}
	//	[]string{"", ""}
	//	[]string{"", "", ""}
	//	[]string{"", "key", "value"}



	fmt.Println()
	// MustCompilePOSIX类似于CompilePOSIX，但如果无法解析该表达式，则会发生恐慌。
	//它简化了保存已编译正则表达式的全局变量的安全初始化。
	re=regexp.MustCompilePOSIX("[a-z]*?")//建立一个最左最长的正则匹配对象，如果是MustCompile（）创建的对象是不会采用这个模式的！
	ret =re.FindAllStringSubmatch("ccaa1234aa1234",-1)
	fmt.Println(ret)
	//输出：
	//	[[ccaa] [] [] [] [aa] [] [] [] []]




	fmt.Println()
	// CompilePOSIX与Compile类似，但是将正则表达式限制为POSIX ERE（egrep）语法，并将match语义更改为最左最长。
	//也就是说，当与文本进行匹配时，regexp返回一个匹配项，该匹配项尽早在输入中（最左侧）开始，并且在其中选择一个尽可能长的匹配项。
	//所谓的最左最长匹配与POSIX指定的早期正则表达式实现所使用的语义相同。
	//但是，可以有多个最左最长的匹配项，具有不同的子匹配项选择，此处此程序包与POSIX不同。
	//在可能的最左最长的匹配项中，此程序包选择一个回溯搜索将首先找到的匹配项，而POSIX指定选择该匹配项以最大化第一个子表达式的长度，然后最大化第二个子表达式的长度，依此类推。在右边。
	// POSIX规则在计算上是禁止的，甚至定义不明确。
	//有关详细信息，请参见https://swtch.com/~rsc/regexp/regexp2.html#posix。
	re,err:=regexp.CompilePOSIX("[a-z]*?")//建立一个最左最长的正则匹配对象，如果是MustCompile（）创建的对象是不会采用这个模式的！
	check_err_regexp(err)
	ret =re.FindAllStringSubmatch("ccaa1234aa1234",-1)
	fmt.Println(ret)
	//输出：
	//	[[ccaa] [] [] [] [aa] [] [] [] []]


	fmt.Println()
	// MatchString reports whether the string s
	// contains any match of the regular expression pattern.
	// More complicated queries need to use Compile and the full Regexp interface.
	// MatchString报告字符串s是否包含正则表达式模式的任何匹配项。
	//更复杂的查询需要使用Compile和完整的Regexp接口。

	//这个函数就相当于是regexp对象的MatchString的方法阉割简洁使用形式，当然底层也是通过Compile(pattern)实例化一个regexp对象，
	// 进而调用这个对象上面的MatchString的方法来实现的！
	matched, e := regexp.MatchString("[a-z]*?", "ccaa1234aa1234")
	check_err_regexp(e)
	fmt.Println(matched)
	//输出：
	//	true
	//还有MatchReader和Match（）这2个函数和上面的函数几乎一样，不再累叙


	fmt.Println()
	//Compile(编译)将解析正则表达式，如果成功，则返回可用于与文本匹配的Regexp对象。
	//当与文本进行匹配时，regexp返回一个匹配项，该匹配项尽早在输入中（最左侧）开始，并且在其中选择从回溯搜索中最先找到的匹配项。
	//这种所谓的“最左边优先”匹配与Perl，Python和其他实现使用的语义相同，尽管此包无需牺牲回溯即可实现它。
	//有关POSIX最左最长的匹配，请参见CompilePOSIX。
	re, e = regexp.Compile("[a-z]*")
	check_err_regexp(e)
	re1, e := regexp.Compile("[a-z]*?")
	check_err_regexp(e)
	dst:=re.FindAllStringSubmatch("ccaa1234aa1234",-1)
	dst1:=re1.FindAllStringSubmatch("ccaa1234aa1234",-1)
	fmt.Printf("%#v\n",re)
	fmt.Println(dst)
	fmt.Println(dst1)
	//输出：
	//	&regexp.Regexp{expr:"[a-z]*",
	//		prog:(*syntax.Prog)(0xc0000c8cc0),
	//		onepass:(*regexp.onePassProg)(nil),
	//		numSubexp:0,
	//		maxBitStateLen:65536,
	//		subexpNames:[]string{""},
	//		prefix:"",
	//		prefixBytes:[]uint8(nil),
	//		prefixRune:0,
	//		prefixEnd:0x0,
	//		mpool:0,
	//		matchcap:2,
	//		prefixComplete:false,
	//		cond:0x0,
	//		minInputLen:0,
	//		longest:false}
	//	[[ccaa] [] [] [] [aa] [] [] [] []]
	//	[[] [] [] [] [] [] [] [] [] [] [] [] [] [] []]


	fmt.Println()
	// CompilePOSIX与Compile类似，但是将正则表达式限制为POSIX ERE（egrep）语法，并将match语义更改为最左最长。
	//也就是说，当与文本进行匹配时，regexp返回一个匹配项，该匹配项尽早在输入中（最左侧）开始，并且在其中选择一个尽可能长的匹配项。
	//所谓的最左最长匹配与POSIX指定的早期正则表达式实现所使用的语义相同。
	//但是，可以有多个最左最长的匹配项，具有不同的子匹配项选择，此处此程序包与POSIX不同。
	//在可能的最左最长的匹配项中，此程序包选择一个回溯搜索将首先找到的匹配项，而POSIX指定选择该匹配项以最大化第一个子表达式的长度，然后最大化第二个子表达式的长度，依此类推。在右边。
	// POSIX规则在计算上是禁止的，甚至定义不明确。
	//有关详细信息，请参见https://swtch.com/~rsc/regexp/regexp2.html#posix。
	re, e = regexp.CompilePOSIX("[a-z]*")
	check_err_regexp(e)
	re1, e = regexp.CompilePOSIX("[a-z]*?")//一旦设置了最长匹配则 正则表达式中的?号将相当于失效了！
	check_err_regexp(e)
	dst=re.FindAllStringSubmatch("ccaa1234aa1234",-1)
	dst1=re1.FindAllStringSubmatch("ccaa1234aa1234",-1)
	fmt.Printf("%#v\n",re)
	fmt.Println(dst)
	fmt.Println(dst1)
	//输出：
	//	&regexp.Regexp{expr:"[a-z]*",
	//		prog:(*syntax.Prog)(0xc0000aed80),
	//		onepass:(*regexp.onePassProg)(nil),
	//		numSubexp:0,
	//		maxBitStateLen:65536,
	//		subexpNames:[]string{""},
	//		prefix:"",
	//		prefixBytes:[]uint8(nil),
	//		prefixRune:0,
	//		prefixEnd:0x0,
	//		mpool:0,
	//		matchcap:2,
	//		prefixComplete:false,
	//		cond:0x0,
	//		minInputLen:0,
	//		longest:true}
	//	[[ccaa] [] [] [] [aa] [] [] [] []]
	//	[[ccaa] [] [] [] [aa] [] [] [] []]


	fmt.Println()
	//Compile(编译)将解析正则表达式，如果成功，则返回可用于与文本匹配的Regexp对象。
	//当与文本进行匹配时，regexp返回一个匹配项，该匹配项尽早在输入中（最左侧）开始，并且在其中选择从回溯搜索中最先找到的匹配项。
	//这种所谓的“最左边优先”匹配与Perl，Python和其他实现使用的语义相同，尽管此包无需牺牲回溯即可实现它。
	//有关POSIX最左最长的匹配，请参见CompilePOSIX。
	re = regexp.MustCompile("[a-z]*")
	fmt.Printf("%#v\n",re)
	re1 = regexp.MustCompile("[a-z]*?")

	dst =re.FindAllStringSubmatch("ccaa1234aa1234",-1)
	dst1 =re1.FindAllStringSubmatch("ccaa1234aa1234",-1)
	fmt.Printf("%#v\n",re)
	fmt.Println(dst)
	fmt.Println(dst1)
	//输出：
	//	&regexp.Regexp{expr:"[a-z]*",
	//		prog:(*syntax.Prog)(0xc0000aee40),
	//		onepass:(*regexp.onePassProg)(nil),
	//		numSubexp:0,
	//		maxBitStateLen:65536,
	//		subexpNames:[]string{""},
	//		prefix:"",
	//		prefixBytes:[]uint8(nil),
	//		prefixRune:0,
	//		prefixEnd:0x0,
	//		mpool:0,
	//		matchcap:2,
	//		prefixComplete:false,
	//		cond:0x0,
	//		minInputLen:0,
	//		longest:false}
	//	&regexp.Regexp{expr:"[a-z]*",
	//		prog:(*syntax.Prog)(0xc0000aee40),
	//		onepass:(*regexp.onePassProg)(nil),
	//		numSubexp:0,
	//		maxBitStateLen:65536,
	//		subexpNames:[]string{""},
	//		prefix:"",
	//		prefixBytes:[]uint8(nil),
	//		prefixRune:0,
	//		prefixEnd:0x0,
	//		mpool:0,
	//		matchcap:2,
	//		prefixComplete:false,
	//		cond:0x0,
	//		minInputLen:0,
	//		longest:false}
	//	[[ccaa] [] [] [] [aa] [] [] [] []]
	//	[[] [] [] [] [] [] [] [] [] [] [] [] [] [] []]



	fmt.Println()
	//QuoteMeta返回将s中所有正则表达式元字符都进行转义后字符串。该字符串可以用在正则表达式中匹配字面值s。
	//例如，QuoteMeta(`[foo]`)会返回`\[foo\]`。
	//需要进行转义的字符包含以下：\.+*?()|[]{}^$

	//字节循环是正确的，因为所有元字符都是ASCII。
	//未找到元字符，因此返回原始字符串。
	Quote_str := regexp.QuoteMeta("[a-z]*")

	fmt.Println(Quote_str)
	//输出：
	//	\[a-z\]\*


}





func ExampleRegexp_Expand() {
	content := []byte(`
	# comment line
	option1: value1
	option2: value2

	# another comment line
	option3: value3`)


	//标志的语法为xyz（设置）、-xyz（清楚）、xy-z（设置xy，清楚z），标志如下：
	//
	//I              大小写敏感（默认关闭）
	//m              ^和$在匹配文本开始和结尾之外，还可以匹配行首和行尾（默认开启）
	//s              让.可以匹配\n（默认关闭）,和.差不多，但是会匹配\n
	//U              非贪婪的：交换x*和x*?、x+和x+?……的含义（默认关闭）

	// Regex pattern captures "key: value" pair from the content.
	//正则表达式模式从content中捕获“键：值”对
	//\s是指空白，包括空格、换行、tab缩进等类型的空白，而\S刚好相反
	//这样一正一反下来，就表示所有的字符，完全的，一字不漏的。
	//\s== [\t\n\f\r ]
	//\S== [^\t\n\f\r ]
	//(?P<name>re)   命名并编号的捕获分组
	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)

	// Template to convert "key: value" to "key=value" by
	// referencing the values captured by the regex pattern.
	//通过引用正则表达式模式捕获的值，将“键：值”转换为“键=值”的模板。
	template := []byte("$key=$value\n")//不一定全得字母，也可以数字+字母，可以数字开头，如$1x等价于${1x}而非${1}x
	template1 := []byte("${key}=${value}\n")
	template2 := []byte("$1=$2\n")//完全由数字开头的才会看成是索引
	template3 := []byte("${1}=${2}\n")
	template4 := []byte("${}=${}\n")
	template5 := []byte("$$-${1}=$$-${2}\n")//如果要在输出中插入一个字面值'$'，在template里可以使用$$。
	template6 := []byte("$$-${key111}=$$-${key222}\n")//匹配不到的值会返回空切片
	template7 := []byte("${-1}=${-2}\n")//试验索引是否可以是负数
	template8 := []byte("${10}=${20}\n")//试验索引是否可以超过最大值

	result := []byte{}

	// For each match of the regex in the content.
	//对于内容中正则表达式的每个匹配项。
	//FindAllSubmatchIndex的第二个参数是负数的话则会默认匹配全部
	for _, submatches := range pattern.FindAllSubmatchIndex(content, -1) {



		//Expand返回新生成的将template添加到dst后面的切片。在添加时，Expand会将template中的变量替换为从src匹配的结果。match应该是被FindSubmatchIndex返回的匹配结果起止位置索引。
		// （通常就是匹配src，除非你要将匹配得到的位置用于另一个[]byte）
		//
		//在template参数里，一个变量表示为格式如：$name或${name}的字符串，其中name是长度>0的字母、数字和下划线的序列。一个单纯的数字字符名如$1会作为捕获分组的数字索引；
		// 其他的名字对应(?P<name>...)语法产生的命名捕获分组的名字。超出范围的数字索引、索引对应的分组未匹配到文本、正则表达式中未出现的分组名，都会被替换为空切片。
		//
		//$name格式的变量名，name会尽可能取最长序列：$1x等价于${1x}而非${1}x，$10等价于${10}而非${1}0。因此$name适用在后跟空格/换行等字符的情况，${name}适用所有情况。
		//
		//如果要在输出中插入一个字面值'$'，在template里可以使用$$。

		// Apply the captured submatches to the template and append the output
		// to the result.
		//将捕获的子匹配项应用于模板，并将输出附加到结果。
		result = pattern.Expand(result, template, content, submatches)
	}
	result1 := []byte{}
	for _, submatches := range pattern.FindAllSubmatchIndex(content, -1) {
		result1 = pattern.Expand(result1, template1, content, submatches)
	}
	result2 := []byte{}
	for _, submatches := range pattern.FindAllSubmatchIndex(content, -1) {
		result2 = pattern.Expand(result2, template2, content, submatches)
	}
	result3 := []byte{}
	for _, submatches := range pattern.FindAllSubmatchIndex(content, -1) {
		result3 = pattern.Expand(result3, template3, content, submatches)
	}
	result4 := []byte{}
	for _, submatches := range pattern.FindAllSubmatchIndex(content, -1) {
		result4 = pattern.Expand(result4, template4, content, submatches)
	}
	result5 := []byte{}
	for _, submatches := range pattern.FindAllSubmatchIndex(content, -1) {
		result5 = pattern.Expand(result5, template5, content, submatches)
	}
	result6 := []byte{}
	for _, submatches := range pattern.FindAllSubmatchIndex(content, -1) {
		result6 = pattern.Expand(result6, template6, content, submatches)
	}
	result7 := []byte{}
	for _, submatches := range pattern.FindAllSubmatchIndex(content, -1) {
		result7 = pattern.Expand(result7, template7, content, submatches)
	}
	result8 := []byte{}
	for _, submatches := range pattern.FindAllSubmatchIndex(content, -1) {
		result8 = pattern.Expand(result8, template8, content, submatches)
	}

	fmt.Println(string(result1))
	fmt.Println(string(result2))
	fmt.Println(string(result3))
	fmt.Println(string(result4))
	fmt.Println(string(result5))
	fmt.Println(string(result6))
	fmt.Println(string(result7))
	fmt.Println(string(result8))
	// Output:
	//	option1=value1
	//	option2=value2
	//	option3=value3
	//
	//	option1=value1
	//	option2=value2
	//	option3=value3
	//
	//	option1=value1
	//	option2=value2
	//	option3=value3
	//
	//	${}=${}
	//	${}=${}
	//	${}=${}
	//
	//	$-option1=$-value1
	//	$-option2=$-value2
	//	$-option3=$-value3
	//
	//	$-=$-
	//	$-=$-
	//	$-=$-
	//
	//	${-1}=${-2}
	//	${-1}=${-2}
	//	${-1}=${-2}
	//
	//	=
	//	=
	//	=
	//其实这个函数和下面的这个函数可以应用到html文件的模板填充中去
}
//这个函数跟上面的差不多！
func ExampleRegexp_ExpandString() {
	content := `
	# comment line
	option1: value1
	option2: value2

	# another comment line
	option3: value3
`

	// Regex pattern captures "key: value" pair from the content.
	//正则表达式模式从内容中捕获“键：值”对。

	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)

	// Template to convert "key: value" to "key=value" by
	// referencing the values captured by the regex pattern.
	//通过引用正则表达式模式捕获的值，将“键：值”转换为“键=值”的模板。
	template := "$key=$value\n"

	result := []byte{}

	// For each match of the regex in the content.
	//对于内容中正则表达式的每个匹配项。
	for _, submatches := range pattern.FindAllStringSubmatchIndex(content, -1) {

		// ExpandString类似于Expand，但是模板和源是字符串。
		//它附加并返回一个字节片，以使调用代码可以控制分配。

		// Apply the captured submatches to the template and append the output
		// to the result.
		//将捕获的子匹配项应用于模板，并将输出附加到结果。
		result = pattern.ExpandString(result, template, content, submatches)
	}
	fmt.Println(string(result))
	// Output:
	// option1=value1
	// option2=value2
	// option3=value3
}



func check_err_regexp(err error) {
	if err != nil {
		fmt.Println(err)
	}
}



