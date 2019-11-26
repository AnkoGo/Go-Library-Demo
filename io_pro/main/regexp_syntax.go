package main

import (
	"fmt"
	"regexp/syntax"
)

func main() {
	//这个包下应该学习的对象如下：

	//// A frag represents a compiled program fragment.
	////frag（fragment简写，分段，片段的意思）表示已编译的程序片段。
	//type frag struct {
	//	i   uint32    // index of first instruction（//第一条指令的索引）
	//	out patchList // where to record end instruction（//在哪里记录结束指令）
	//}
	//
	// A patchList is a list of instruction pointers that need to be filled in (patched).
	// Because the pointers haven't been filled in yet, we can reuse their storage
	// to hold the list. It's kind of sleazy, but works well in practice.
	// See https://swtch.com/~rsc/regexp/regexp1.html for inspiration.
	//
	// These aren't really pointers: they're integers, so we can reinterpret them
	// this way without using package unsafe. A value l denotes
	// p.inst[l>>1].Out (l&1==0) or .Arg (l&1==1).
	// l == 0 denotes the empty list, okay because we start every program
	// with a fail instruction, so we'll never want to point at its output link.
	// patchList是需要填写（修补）的指令指针的列表。
	//因为尚未填充指针，所以我们可以重复使用它们的存储来保存列表。 这有点琐，但在实践中效果很好。
	//有关启发，请参见https://swtch.com/~rsc/regexp/regexp1.html。
	//这些并不是真正的指针：它们是整数，因此我们可以通过这种方式重新解释它们，而无需使用不安全的包。 值l表示p.inst [l >> 1] .Out（l＆1 == 0）或.Arg（l＆1 == 1）。
	// l == 0表示为空列表，好吧，因为我们以失败指令开始每个程序，所以我们永远也不想指向它的输出链接。
	type patchList uint32

	//type compiler struct {
	//	p *Prog	//p是编译的程序的指针
	//}




	//// Compiled program.
	//// May not belong in this package, but convenient for now.
	////编译程序
	////可能不属于此程序包，但现在很方便。
	//
	//// A Prog is a compiled regular expression program.（// Prog是已编译的正则表达式程序。）
	//type Prog struct {
	//	Inst   []Inst	//instruction指令集
	//	Start  int // index of start instruction（开始指令的索引）
	//	NumCap int // number of InstCapture insts in re（在正则对象中捕获分组指令的数目）
	//}


	//// An Inst is a single instruction in a regular expression program.（// Inst是正则表达式程序中的一条指令比如“a+c*”中的+和*都是一条指令。）
	//type Inst struct {
	//	Op   InstOp
	//	Out  uint32 // all but InstMatch, InstFail（//除了InstMatch，InstFail之外的所有）
	//	Arg  uint32 // InstAlt, InstAltMatch, InstCapture, InstEmptyWidth
	//	Rune []rune
	//}

	//// An InstOp is an instruction opcode.（// InstOp是指令操作码。）
	//type InstOp uint8
	//
	//const (
	//	InstAlt InstOp = iota
	//	InstAltMatch
	//	InstCapture		//捕获分组的指令
	//	InstEmptyWidth	//\b或者\B的指令
	//	InstMatch
	//	InstFail		//不匹配的指令
	//	InstNop			//什么都不做的指令
	//	InstRune		//匹配某个字符rune的指令
	//	InstRune1		//匹配某个字符rune的指令
	//	InstRuneAny
	//	InstRuneAnyNotNL
	//)



	//// An EmptyOp specifies a kind or mixture of zero-width assertions.（// EmptyOp指定零宽度断言的一种或混合形式。）
	//type EmptyOp uint8
	//
	//const (
	//	EmptyBeginLine EmptyOp = 1 << iota
	//	EmptyEndLine
	//	EmptyBeginText
	//	EmptyEndText
	//	EmptyWordBoundary
	//	EmptyNoWordBoundary
	//)




	fmt.Println("-----------------------")
	// IsWordChar reports whether r is consider a ``word character''
	// during the evaluation of the \b and \B zero-width assertions.
	// These assertions are ASCII-only: the word characters are [A-Za-z0-9_].
	//IsWordChar报告在\ b和\ B零宽度断言的评估期间r是否被视为“单词字符”。
	//这些断言仅是ASCII：字字符为[A-Za-z0-9_]。
	fmt.Println(syntax.IsWordChar('a'))
	fmt.Println(syntax.IsWordChar(110))
	fmt.Println(syntax.IsWordChar(122))//97+26=123
	fmt.Println(syntax.IsWordChar(123))
	fmt.Println(syntax.IsWordChar(126))
	fmt.Println(syntax.IsWordChar(127))


	fmt.Println()
	// Compile compiles the regexp into a program to be executed.
	// The regexp should have been simplified already (returned from re.Simplify).
	// Compile将正则表达式编译为要执行的程序。
	// regexp应该已经简化了（从re.Simplify返回）。

	//// A Regexp is a node in a regular expression syntax tree.
	//// Regexp是正则表达式语法树中的一个节点。(跟regexp包中的Regexp不是同一个对象)
	//type Regexp struct {
	//	Op       Op // operator
	//	Flags    Flags
	//	Sub      []*Regexp  // subexpressions, if any（//子表达式（如果有））
	//	Sub0     [1]*Regexp // storage for short Sub（//存储简短的Sub）
	//	Rune     []rune     // matched runes, for OpLiteral, OpCharClass（//匹配的符文，用于OpLiteral，OpCharClass）
	//	Rune0    [2]rune    // storage for short Rune（//储存短符文）
	//	Min, Max int        // min, max for OpRepeat（// OpRepeat的最小值，最大值）
	//	Cap      int        // capturing index, for OpCapture（//捕获索引，用于OpCapture）
	//	Name     string     // capturing name, for OpCapture（//为OpCapture捕获名称）
	//}

	//// An Op is a single regular expression operator.（// Op是单个正则表达式运算符。）
	//type Op uint8
	//
	//// Operators are listed in precedence order, tightest binding to weakest.
	//// Character class operators are listed simplest to most complex
	//// (OpLiteral, OpCharClass, OpAnyCharNotNL, OpAnyChar).
	////运算符按优先顺序列出，从最紧密的绑定到最弱的绑定。
	////列出的字符类运算符最简单到最复杂（OpLiteral，OpCharClass，OpAnyCharNotNL，OpAnyChar）。
	//
	//const (
	//	OpNoMatch        Op = 1 + iota // matches no strings（匹配无字符串）
	//	OpEmptyMatch                   // matches empty string（匹配空字符串）
	//	OpLiteral                      // matches Runes sequence（匹配runes序列）
	//	OpCharClass                    // matches Runes interpreted as range pair list（匹配解释为范围对列表的runes）
	//	OpAnyCharNotNL                 // matches any character except newline(匹配除换行符以外的任何字符)
	//	OpAnyChar                      // matches any character(匹配任何字符)
	//	OpBeginLine                    // matches empty string at beginning of line(在行开头匹配空字符串)
	//	OpEndLine                      // matches empty string at end of line(在行尾匹配空字符串)
	//	OpBeginText                    // matches empty string at beginning of text(在文本开头匹配空字符串)
	//	OpEndText                      // matches empty string at end of text(在文本尾匹配空字符串)
	//	OpWordBoundary                 // matches word boundary `\b`（匹配单词边界`\b`）
	//	OpNoWordBoundary               // matches word non-boundary `\B`（匹配单词非边界`\B`）
	//	OpCapture                      // capturing subexpression with index Cap, optional name Name（使用索引Cap捕获子表达式，可选名称Name）
	//	OpStar                         // matches Sub[0] zero or more times（匹配Sub [0]零次或多次）
	//	OpPlus                         // matches Sub[0] one or more times（匹配Sub [0]一次或多次）
	//	OpQuest                        // matches Sub[0] zero or one times（匹配Sub [0]零次或1次）
	//	OpRepeat                       // matches Sub[0] at least Min times, at most Max (Max == -1 is no limit)（至少与Sub [0]匹配Min次，最大Max（Max == -1是没有限制即匹配全部））
	//	OpConcat                       // matches concatenation of Subs（匹配Sub的串联）
	//	OpAlternate                    // matches alternation of Subs（匹配Subs的交替）
	//)



	// Flags control the behavior of the parser and record information about regexp context.
	//Flags控制解析器的行为并记录有关regexp上下文的信息。说白了就是设置一些零碎的项
	//type Flags uint16
	//
	//const (
	//	FoldCase      Flags = 1 << iota // case-insensitive match(不区分大小写的匹配)
	//	Literal                         // treat pattern as literal string(将正则表达式视为文字字符串，也就是进行转义)
	//	ClassNL                         // allow character classes like [^a-z] and [[:space:]] to match newline（允许像[^a-z]和[[:space:]]这样的字符类匹配换行符）
	//	DotNL                           // allow . to match newline(允许 . 匹配换行符)
	//	OneLine                         // treat ^ and $ as only matching at beginning and end of text(将^和$视为仅在文本的开头和结尾匹配)
	//	NonGreedy                       // make repetition operators default to non-greedy(使重复运算符默认为非贪婪)
	//	PerlX                           // allow Perl extensions(允许Perl扩展)
	//	UnicodeGroups                   // allow \p{Han}, \P{Han} for Unicode group and negation(允许\p{Han}，\P{Han}用于Unicode组和非Unicode组)
	//	WasDollar                       // regexp OpEndText was $, not \z（设置正则表达式OpEndText是$，而不是\z）
	//	Simple                          // regexp contains no counted repetition（正则表达式不包含重复计数）
	//
	//	MatchNL = ClassNL | DotNL
	//
	//	Perl        = ClassNL | OneLine | PerlX | UnicodeGroups // as close to Perl as possible(//尽可能接近Perl)
	//	POSIX Flags = 0                                         // POSIX syntax(// POSIX语法)
	//)


	//Re:=syntax.Regexp{
	//	Op:    syntax.OpPlus,
	//	Flags: syntax.FoldCase,
	//	Sub:   nil,
	//	Sub0:  [1]*syntax.Regexp{},
	//	Rune:  nil,
	//	Rune0: [2]rune{},
	//	Min:   0,
	//	Max:   0,
	//	Cap:   0,
	//	Name:  "",
	//}
	//fmt.Println(syntax.Compile(&Re))
	OnePassCutoff()
}

// Check that one-pass cutoff does trigger.(//检查一次通过截止是否触发。)
func OnePassCutoff() {
	// Parsing.

	// Parse parses a regular expression string s, controlled by the specified
	// Flags, and returns a regular expression parse tree. The syntax is
	// described in the top-level comment.
	// Parse解析由指定的Flags控制的正则表达式字符串s，并返回一个正则表达式解析树。 语法在顶级注释中描述。
	//注意这个syntax包的regexp跟regexp包的regexp对象时不同的！
	re, err := syntax.Parse(`^a+b{3}$`, syntax.Perl)
	fmt.Printf("%#v\n",re)
	fmt.Println("-----")
	// MaxCap walks the regexp to find the maximum capture index.(// MaxCap使用正则表达式查找最大捕获索引。)
	fmt.Printf("re.MaxCap()==%#v\n",re.MaxCap())
	fmt.Printf("re.String()==%#v\n",re.String())
	re1, err := syntax.Parse(`^a+b{2}$`, syntax.Perl)
	// Equal reports whether x and y have identical structure.(//Equal报告x和y是否具有相同的结构。)
	fmt.Printf("re.Equal(re1)==%#v\n",re.Equal(re1))
	// CapNames walks the regexp to find the names of capturing groups.(// CapNames使用正则表达式查找捕获组的名称。)
	fmt.Printf("re.CapNames()==%#v\n",re.CapNames())
	fmt.Println("-----")
	check_err_regexp_syntax(err)
	//Compile将正则表达式编译为要执行的程序。
	// regexp应该已经简化了（从re.Simplify返回）。

	// Simplify返回与re等效的正则表达式，但不计算重复次数，并具有其他各种简化形式，例如将/(?:a+)+/ 重写为/a+/。
	//生成的正则表达式将正确执行，但是其字符串表示形式将不会产生相同的解析树，因为捕获括号可能已被复制或删除。 例如， /(x){1,2}/的简化形式为 /(x)(x)?/，但是两个括号都捕获为$1。
	//返回的正则表达式可以与原始结构共享或为原始结构。
	p, err := syntax.Compile(re.Simplify())
	check_err_regexp_syntax(err)
	fmt.Println("==",p.String())
	fmt.Printf("p.Inst==%#v\n",p.Inst)
	fmt.Println("p.Start==",p.Start)
	fmt.Println("p.NumCap==",p.NumCap)//目前不知道这个东西到底是什么
	// Prefix returns a literal string that all matches for the
	// regexp must start with. Complete is true if the prefix
	// is the entire match.
	//Prefix(前缀)返回一个文字字符串，所有与regexp匹配的字符串都必须以该字符串开头。 如果前缀是整个匹配项，则Complete为true。
	fmt.Println(p.Prefix())
	// StartCond returns the leading empty-width conditions that must
	// be true in any match. It returns ^EmptyOp(0) if no matches are possible.
	// StartCond返回在任何匹配中都必须为true的前导空白宽度条件。 如果没有匹配项，则返回^ EmptyOp（0）。

	//// An EmptyOp specifies a kind or mixture of zero-width assertions.
	//type EmptyOp uint8
	//
	//const (
	//	EmptyBeginLine EmptyOp = 1 << iota
	//	EmptyEndLine
	//	EmptyBeginText
	//	EmptyEndText
	//	EmptyWordBoundary
	//	EmptyNoWordBoundary
	//)
	fmt.Println("p.StartCond()==",p.StartCond())
	//输出：
	//&syntax.Regexp{Op:0x12,
	//				Flags:0x0,
	//				Sub:[]*syntax.Regexp{(*syntax.Regexp)(0xc000088000),
	//									(*syntax.Regexp)(0xc0000880e0),
	//									(*syntax.Regexp)(0xc0000881c0),
	//									(*syntax.Regexp)(0xc000088230)},
	//				Sub0:[1]*syntax.Regexp{(*syntax.Regexp)(0xc000088000)},(简化形式就是记录第一个元素的开始地址)
	//				Rune:[]int32(nil),
	//				Rune0:[2]int32{0, 0},
	//				Min:0,
	//				Max:0,
	//				Cap:0,
	//				Name:""}
	//-----
	//re.MaxCap()==0
	//re.String()=="\\Aa+b{3}(?-m:$)"
	//re.Equal(re1)==false
	//re.CapNames()==[]string{""}
	//-----
	//==   0	fail
	//	1*	empty 4 -> 2(empty代表指令的字符串名称)
	//	2	rune1 "a" -> 3
	//	3	alt -> 2, 4（ Out:0x2,Arg:0x4,（out表示输出第几个了，Arg表示一次指令输出的字符串切片，他包含out））
	//	4	rune1 "b" -> 5（ Out:0x5）
	//	5	rune1 "b" -> 6
	//	6	rune1 "b" -> 7
	//	7	empty 8 -> 8
	//	8	match（上面输出的是*Prog假设执行的话会执行的指令信息或者叫做指令句法，这也是本包的作用）
	//
	//==[]syntax.Inst{syntax.Inst{Op:0x5,
	// 							  Out:0x0,
	// 							  Arg:0x0,
	// 							  Rune:[]int32(nil)},
	//      		  syntax.Inst{Op:0x3,
	//      	 				  Out:0x2,
	//      	 				  Arg:0x4,
	//      	 				  Rune:[]int32(nil)},
	//      	 	  syntax.Inst{Op:0x8,
	//      	 				  Out:0x3,
	//      	 				  Arg:0x0,
	//      	 				  Rune:[]int32{97}},
	//      	 	  syntax.Inst{Op:0x0,
	//      	 				  Out:0x2,
	//      	 				  Arg:0x4,
	//      	 				  Rune:[]int32(nil)},
	//      	 	  syntax.Inst{Op:0x8,
	//      	 				  Out:0x5,
	//      	 				  Arg:0x0,
	//      	 				  Rune:[]int32{98}},
	//      	 	  syntax.Inst{Op:0x8, Out:0x6, Arg:0x0, Rune:[]int32{98}}, syntax.Inst{Op:0x8, Out:0x7, Arg:0x0, Rune:[]int32{98}}, syntax.Inst{Op:0x3, Out:0x8, Arg:0x8, Rune:[]int32(nil)}, syntax.Inst{Op:0x4, Out:0x0, Arg:0x0, Rune:[]int32(nil)}}== 1
	//p.Start== 1
	//p.NumCap== 2
	// false（前面有个空字符串）
	//p.StartCond()== 4(就是EmptyBeginText)

}

func check_err_regexp_syntax(err error) {
	if err != nil {
		fmt.Println(err)
	}
}



