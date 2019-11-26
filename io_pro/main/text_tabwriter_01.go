package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)



//这整个包的原理可以在tabwriter.go的format（）方法里面体现

//在介绍整个包之前很有必要先看下面的对象（强烈建议自己写个例子调试一波，否则即使你看了下面我翻译的东西也是难以懂得）：
//// ----------------------------------------------------------------------------
//// Filter implementation
//
//// A cell represents a segment of text terminated by tabs or line breaks.
//// The text itself is stored in a separate buffer; cell only describes the
//// segment's size in bytes, its width in runes, and whether it's an htab
//// ('\t') terminated cell.
////过滤器实现
////
////单元格代表由制表符或换行符终止的一段文本。
////文本本身存储在单独的缓冲区中。 单元格仅以字节为单位描述段的大小，以符文表示其宽度，以及它是否为以htab（'\ t'）结尾的单元格。
////一个cell是一行中的一个\t分割的左或者右块
////每一line中的最后一个cell都是size和width等于0，且htab等于false,以表示结束一行了
//type cell struct {
//	size  int  // cell size in bytes	（//单元大小（以字节为单位））
//	width int  // cell width in runes	（//符格中的单元格宽度）
//	htab  bool // true if the cell is terminated by an htab ('\t')	（//如果单元格由htab（'\t'）终止，则为true）
//}
//
//// A Writer is a filter that inserts padding around tab-delimited
//// columns in its input to align them in the output.
//// Writer是一个过滤器，在其输入中的制表符分隔的列周围插入填充以使它们在输出中对齐。
////
//// The Writer treats incoming bytes as UTF-8-encoded text consisting
//// of cells terminated by horizontal ('\t') or vertical ('\v') tabs,
//// and newline ('\n') or formfeed ('\f') characters; both newline and
//// formfeed act as line breaks.
////Writer将传入的字节输出为UTF-8编码的文本，此文本由水平（'\t'）或垂直（'\v'）制表符以及换行符（'\n'）或换页符（'\f'）终止的单元格cells组成；换行符和换页符都充当换行符。
////
//// Tab-terminated cells in contiguous lines constitute a column. The
//// Writer inserts padding as needed to make all cells in a column have
//// the same width, effectively aligning the columns. It assumes that
//// all characters have the same width, except for tabs for which a
//// tabwidth must be specified. Column cells must be tab-terminated, not
//// tab-separated: non-tab terminated trailing text at the end of a line
//// forms a cell but that cell is not part of an aligned column.
//// For instance, in this example (where | stands for a horizontal tab):
////连续行中制表符终止的单元格构成一列。 Writer根据需要插入填充以使列中的所有单元格具有相同的宽度，从而有效地对齐列。
//// 它假定所有字符都具有相同的宽度，但必须为其指定制表符宽度的制表符除外。 列单元格必须以制表符结尾，
//// 而不用制表符分隔：行末尾的非制表符结尾的尾随文本构成一个单元格，但该单元格不是对齐列的一部分。
////例如，在此示例中（其中|代表水平制表符）：
////
////	aaaa|bbb|d
////	aa  |b  |dd
////	a   |
////	aa  |cccc|eee
////
//// the b and c are in distinct columns (the b column is not contiguous
//// all the way). The d and e are not in a column at all (there's no
//// terminating tab, nor would the column be contiguous).
//// b和c在不同的列中（b列并非一直是连续的）。 d和e根本不在列中（没有终止标签，列也不是连续的）。
////
//// The Writer assumes that all Unicode code points have the same width;
//// this may not be true in some fonts or if the string contains combining
//// characters.
////Writer假定所有Unicode代码点都具有相同的宽度。 在某些字体fonts中或者字符串string包含组合字符characters时，可能不正确。
////
//// If DiscardEmptyColumns is set, empty columns that are terminated
//// entirely by vertical (or "soft") tabs are discarded. Columns
//// terminated by horizontal (or "hard") tabs are not affected by
//// this flag.
////如果设置了DiscardEmptyColumns，则丢弃完全由垂直(or "soft")制表符终止的空列。 水平(or "hard")制表符终止的列不受此标志flag的影响。
////
//// If a Writer is configured to filter HTML, HTML tags and entities
//// are passed through. The widths of tags and entities are
//// assumed to be zero (tags) and one (entities) for formatting purposes.
////如果将Writer配置为过滤HTML，则会通过HTML标签和实体。 出于格式化目的，假定标签tags和实体entities的宽度分别为零和一。
////
//// A segment of text may be escaped by bracketing it with Escape
//// characters. The tabwriter passes escaped text segments through
//// unchanged. In particular, it does not interpret any tabs or line
//// breaks within the segment. If the StripEscape flag is set, the
//// Escape characters are stripped from the output; otherwise they
//// are passed through as well. For the purpose of formatting, the
//// width of the escaped text is always computed excluding the Escape
//// characters.
////通过用转义字符将其括起来，可以对一段文本进行转义。 tabwriter将转义过的文本段通过原样传递。 特别是，它不会解释段中的任何制表符tabs或换行符line breaks。
////如果设置StripEscape标志，则从输出中删除转义字符；否则，它们也会通过。 为了格式化，总是计算转义文本的宽度（不包括转义字符）。
////
//// The formfeed character acts like a newline but it also terminates
//// all columns in the current line (effectively calling Flush). Tab-
//// terminated cells in the next line start new columns. Unless found
//// inside an HTML tag or inside an escaped text segment, formfeed
//// characters appear as newlines in the output.
////换页符的作用类似于换行符，但它也会终止当前行中的所有列（有效地调用Flush）。
//// 下一行line中以制表符结尾的单元格将开始新列。 除非在HTML标记内或转义的文本段内找到，否则换页符在输出中显示为换行符。
////
//// The Writer must buffer input internally, because proper spacing
//// of one line may depend on the cells in future lines. Clients must
//// call Flush when done calling Write.
////Writer必须在内部缓冲输入，因为一行的适当间距可能取决于将来行中的单元格。 调用完写后，客户端必须调用Flush。
////
////type Writer struct {
////	// configuration	(//配置)
////	output   io.Writer
////	minwidth int
////	tabwidth int
////	padding  int
////	padbytes [8]byte
////	flags    uint
////
////	// current state	(// 当前状态，同时以下的配置一般在reset（）函数中重置信息，同时reset（）函数不仅仅出现在init()中)
////	buf     []byte   // collected text excluding tabs or line breaks	(//收集的文本（要格式化的文本中不包括制表符或换行符，剩下的就是了）)
////	pos     int      // buffer position up to which cell.width of incomplete cell has been computed	（buf中字节的个数）
////	cell    cell     // current incomplete cell; cell.width is up to buf[pos] excluding ignored sections	（单元格，除\t外的字母和数字，符号等字符）
////	endChar byte     // terminating char of escaped sequence (Escape for escapes, '>', ';' for HTML tags/entities, or 0)	（//终止转义序列的字符（用于转义的Escape，用于HTML标签/实体的E>，';'或0））
////	lines   [][]cell // list of lines; each line is a list of cells	（//行列表； 每行是一个单元格列表）
////	widths  []int    // list of column widths in runes - re-used during formatting	（//符文中的列宽列表-在格式化期间重复使用，这个非常重要，整个format（）递归函数求取的就是这个属性的值）
////}








func main() {

	w := new(tabwriter.Writer)
	// Format in tab-separated columns with a tab stop of 8.
	// 用8个制表符停止位格式化制表符分隔的列。
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "a\tb\tc\td\t.")
	fmt.Fprintln(w, "123\t12345\t12345678900\t123456789\t.")
	fmt.Fprintln(w)
	w.Flush()


	//// A Writer must be initialized with a call to Init. The first parameter (output)
	//// specifies the filter output. The remaining parameters control the formatting:
	////
	////	minwidth	minimal cell width including any padding
	////	tabwidth	width of tab characters (equivalent number of spaces)
	////	padding		padding added to a cell before computing its width
	////	padchar		ASCII char used for padding
	////			if padchar == '\t', the Writer will assume that the
	////			width of a '\t' in the formatted output is tabwidth,
	////			and cells are left-aligned independent of align_left
	////			(for correct-looking results, tabwidth must correspond
	////			to the tab width in the viewer displaying the result)
	////	flags		formatting control
	////
	////必须通过调用Init初始化Writer。 第一个参数（output）指定过滤器过滤后的文本输出到哪里去。 其余参数控制格式：
	//// minwidth		最小单元格宽度，包括任何填充（下面的padding字段）
	//// tabwidth		制表符的宽度（等效空格）
	//// padding		在计算单元格宽度之前将填充添加到单元格（cell的宽度小于minwidth时候，单元格的总宽度=minwidth；否则，单元格的总宽度=cell.width+padding)
	//// padchar		用于填充（上面的padding字段）的ASCII字符，padbytes [8]byte的8个元素都是这个padchar,而且必须是byte类型，不能是其他编码的字符，因为他
	// 					一律按照一个byte来进行读取,如果你输入的是其他编码的字符会识别出错！
	////				如果padchar =='\t'，Writer将假定格式化输出中'\t'的宽度为tabwidth，并且单元格独立于align_left左对齐（为获得外观正确的结果，
	//// 				制表符宽度必须与显示结果的查看器中的制表符宽度相对应）
	//// flags			格式控制（主要是包含转义字符的处理，文本对齐方式，空列column处理方式，是否对制表符列(!!)使用制表符代替填充字符padchar和是否在列之间（格式化后）打印竖线）
	//
	//
	//// Internal representation (current state):
	////
	//// - all text written is appended to buf; tabs and line breaks are stripped away
	//// - at any given time there is a (possibly empty) incomplete cell at the end
	////   (the cell starts after a tab or line break)
	//// - cell.size is the number of bytes belonging to the cell so far
	//// - cell.width is text width in runes of that cell from the start of the cell to
	////   position pos; html tags and entities are excluded from this width if html
	////   filtering is enabled
	//// - the sizes and widths of processed text are kept in the lines list
	////   which contains a list of cells for each line
	//// - the widths list is a temporary list with current widths used during
	////   formatting; it is kept in Writer because it's re-used
	////
	////                    |<---------- size ---------->|
	////                    |                            |
	////                    |<- width ->|<- ignored ->|  |
	////                    |           |             |  |
	//// [---processed---tab------------<tag>...</tag>...]
	//// ^                  ^                         ^
	//// |                  |                         |
	//// buf                start of incomplete cell  pos
	//
	//
	////内部表示（当前状态）：
	////-所有写入的文本都附加到buf后面；制表符和换行符被剥离
	////-在任何给定时间，结尾处都有一个（可能为空）不完整的单元格（该单元格在制表符或换行符之后开始）
	////-cell.size是到目前为止属于该单元cell的字节数
	////-cell.width是从单元格的开始到位置pos的文本宽度，以该单元格的符文表示；如果启用了html过滤，则从此宽度中排除html标签和实体
	////-处理后的文本的大小和宽度保留在行列表中，该列表包含每行的单元格列表
	////-宽度列表是一个临时列表，其中包含格式化formatting时使用的当前宽度；它被保留在Writer中，因为它会被重复使用
	////
	////                    |<---------- size ---------->|
	////                    |                            |
	////                    |<- width ->|<- ignored ->|  |
	////                    |           |             |  |
	//// [---processed---tab------------<tag>...</tag>...]
	//// ^                  ^                         ^
	//// |                  |                         |
	//// buf                start of incomplete cell  pos

	////可以使用这些标志控制格式。
	//const (
	//	// Ignore html tags and treat entities (starting with '&'
	//	// and ending in ';') as single characters (width = 1).
	//	//忽略html标签，并将实体entities（以'＆'开头，以';'结尾）视为单个字符（宽度width= 1）。
	//	FilterHTML uint = 1 << iota
	//
	//	// Strip Escape characters bracketing escaped text segments
	//	// instead of passing them through unchanged with the text.
	//	//删除转义字符，将转义的文本段括起来，而不是将其与文本保持不变。
	//	StripEscape
	//
	//	// Force right-alignment of cell content.
	//	// Default is left-alignment.
	//	//强制将单元格内容右对齐。
	//	//默认为左对齐。
	//	AlignRight
	//
	//	// Handle empty columns as if they were not present in
	//	// the input in the first place.
	//	//清除空列（\v会导致空列），就好像它们首先不在输入中一样。
	//	DiscardEmptyColumns
	//
	//	// Always use tabs for indentation columns (i.e., padding of
	//	// leading empty cells on the left) independent of padchar.
	//	//始终对制表符列(!!,其他列不符合)使用制表符（即，左侧的前导空白单元格（\t等）的填充），而不是使用padchar填充。
	//	TabIndent
	//
	//	// Print a vertical bar ('|') between columns (after formatting).
	//	// Discarded columns appear as zero-width columns ("||").
	//	//在列之间（格式化后）打印竖线（'|'）。
	//	//舍弃的列显示为零宽度的列（"||"）。
	//	Debug
	//)


	// Format right-aligned in space-separated columns of minimal width 5
	// and at least one blank of padding (so wider column entries do not
	// touch each other).
	//用最小宽度为5且空格至少为空白的空格分隔的形式来格式化制表符分隔的列向右对齐（因此，较宽的列条目不会相互接触，语句1是默认的左对齐）。
	//w.Init(os.Stdout, 5, 0, 2, '+', 0)//语句0
	//w.Init(os.Stdout, 5, 0, 1, '+', 0)//语句1
	//w.Init(os.Stdout, 5, 0, 2, '+', tabwriter.AlignRight)//语句2
	//语句3，如果没空列，那么跟语句1没什么本质的区别,
	//w.Init(os.Stdout, 5, 0, 2, '+', tabwriter.DiscardEmptyColumns | tabwriter.Debug)//语句3
	//w.Init(os.Stdout, 5, 0, 2, '+', tabwriter.TabIndent)//语句4
	//下面的init会覆盖上面的init()
	//w.Init(os.Stdout, 5, 0, 2, '+', tabwriter.Debug)//语句5
	//w.Init(os.Stdout, 5, 0, 2, '+', tabwriter.FilterHTML | tabwriter.Debug )//语句6
	//w.Init(os.Stdout, 5, 0, 2, '+', tabwriter.AlignRight|tabwriter.FilterHTML | tabwriter.Debug )//语句7
	//w.Init(os.Stdout, 5, 0, 2, '+',  tabwriter.Debug)//语句8
	//w.Init(os.Stdout, 5, 0, 2, '.',  tabwriter.AlignRight|tabwriter.Debug)//语句9
	w.Init(os.Stdout, 5, 0, 2, '+', tabwriter.AlignRight|tabwriter.DiscardEmptyColumns | tabwriter.Debug)//语句10


	//fmt.Fprintln(w, "\ta\tb\tc\td\t.")
	//fmt.Fprintln(w, "\t001234\t345\t001234567\t00123456789\t.")//格式化内容1
	//下面的\t<body>要和\ta对齐，\t</html>要和\tb
	//fmt.Fprintln(w, "<html>\t<body><h1>My First Heading</h1><p>My first paragraph.</p></body>\t</html>")//格式化内容2
	//fmt.Fprintln(w, "<html>\t<body><h1>\tMy First Heading</h1><p>My first paragraph.</p></body>\t</html>")//格式化内容3

	//下面3行代码对应语句8才需要用到
	//fmt.Fprintln(w, "abc\xff\tdef")//格式化内容4
	//fmt.Fprintln(w, "\xff\"foo\t\n\tbar\"\xff")//格式化内容4


	//fmt.Fprintln(w, "g) f&lt;o\t<b>bar</b>\t non-terminated entity &amp")//格式化内容5，还是不大懂

	//fmt.Fprintln(w, "1\t2\t3\t4\n" + "11\t222\t3333\t44444\n")//格式化内容6

	//fmt.Fprintln(w, "本\tb\tc\n" +
	//	"aa\t\u672c\u672c\u672c\tcccc\tddddd\n" +
	//	"aaa\tbbbb\n")//格式化内容7

	//fmt.Fprintln(w, ".0\t.3\t2.4\t-5.1\t\n" +
	//					"23.0\t12345678.9\t2.4\t-989.4\t\n" +
	//					"5.1\t12.0\t2.4\t-7.0\t\n" +
	//					".0\t0.0\t332.0\t8908.0\t\n" +
	//					".0\t-.3\t456.4\t22.1\t\n" +
	//					".0\t1.2\t44.4\t-13.3\t\t",)//格式化内容8



	fmt.Fprintln(w, "a\t\tb")//格式化内容9,与下面的一行代码不可同时不被注释，否则会相互影响
	//fmt.Fprintln(w, "a\v\vb")//格式化内容9



	fmt.Fprintln(w)//这一行主要是为了换行，在之前我们换行没写过参数进println中，但是假如我们指定了
	// 写入到哪里去的话（也就是采用Fprintln这个api），那么我们就要指定目的地

	// Flush should be called after the last call to Write to ensure
	// that any data buffered in the Writer is written to output. Any
	// incomplete escape sequence at the end is considered
	// complete for formatting purposes.
	//应该在最后一次调用Write之后调用Flush，以确保Writer中缓冲的所有数据都被写入输出output。
	// 最后任何不完整的转义序列都出于格式化目的被视为完整。
	w.Flush()

	//语句0内容2时候输出：
	//	a	b	c		d		.
	//	123	12345	12345678900	123456789	.
	//
	//	++++++++a+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++b++++c++++d++++.
	//	<html>++<body><h1>My First Heading</h1><p>My first paragraph.</p></body>++</html>

	//语句0内容3时候输出：
	//省略
	//
	//	++++++++a+++++++++++b+++++++++++++++++++++++++++++++++++++++++++++++++++++++c++++d++++.
	//	<html>++<body><h1>++My First Heading</h1><p>My first paragraph.</p></body>++</html>


	//语句6内容3时候输出：
	//省略
	//
	//	+++++|a++++|b++++++++++++++++++++++++++++++++++++|c++++|d++++|.
	//	<html>+++++|<body><h1>+++++|My First Heading</h1><p>My first paragraph.</p></body>++|</html>
	//总体来说没看出对齐的标准是什么

	//语句7内容3时候输出：
	//省略
	//
	//	+++++|++++a|++++++++++++++++++++++++++++++++++++b|++++c|++++d|.
	//	+++++<html>|+++++<body><h1>|++My First Heading</h1><p>My first paragraph.</p></body>|</html>
	//总体来说没看出对齐的标准是什么

	//语句8内容3时候输出：
	//省略
	//
	//++++++++++++++|a++++|b++++|c++++|d++++|.
	//abc	def
	//"foo++|
	//++++++++++++++|bar"
	//假设去除语句8中的tabwriter.StripEscape，则输出：
	//++++++++++++++|a++++|b++++|c++++|d++++|.
	//abc�	def
	//�"foo++|
	//++++++++++++++|bar"�

	//语句6内容5时候输出：
	//省略
	//
	//g) f&lt;o++|<b>bar</b>++| non-terminated entity &amp

	//语句6内容6时候输出：
	//省略
	//
	//1++++|2++++|3+++++|4
	//11+++|222++|3333++|44444

	//语句5内容7时候输出：
	//省略
	//
	//本++++|b++++|c
	//aa+++|本本本++|cccc++|ddddd
	//aaa++|bbbb
	//很明显这里把中文字符也算作一个宽度了，但他不是一个宽度，所以看上去不像是对齐，其实go代码认为是对齐的！
	// 但是我们实际上无法用go外国人的思维来思考我们中文的宽度，所以实际上没达到对齐的效果


	//语句9内容8时候输出：
	//省略
	//.....0|...........3|....2.4|....-5.1|
	//..23.0|..12345678.9|....2.4|..-989.4|
	//...5.1|........12.0|....2.4|....-7.0|
	//.....0|.........0.0|..332.0|..8908.0|
	//.....0|.........-.3|..456.4|....22.1|
	//.....0|.........1.2|...44.4|...-13.3|.....|



	//语句3内容9时候输出：
	//省略
	//a++++|+++++|b
	//a++++||b

	//语句10内容9时候输出：
	//省略
	//++++a|+++++|b
	//++++a||b



	//以下均是基于内容1：
	//语句0左对齐padding=2输出：
	//	a       b       c       d               .
	//	123     12345   1234567 123456789       .
	//
	//	+++++a+++++++b++++c++++++++++d++++++++++++.
	//	+++++001234++345++001234567++00123456789++.

	//语句4(!!)左对齐padding=2,tabwriter.TabIndent输出：(省略上半部分)
	//
	//	a+++++++b++++c++++++++++d++++++++++++.
	//	001234++345++001234567++00123456789++.
	//通过对比上面的语句0，可以看到每一行的最开头的+++++没有了，


	//语句1左对齐padding=1输出：(省略上半部分)
	//
	//	+++++a++++++b++++c+++++++++d+++++++++++.
	//	+++++001234+345++001234567+00123456789+.(通过补右来对齐左)
	//语句1为了和语句0进行对比（padding不同）



	//语句5左对齐padding=2输出：(省略上半部分)
	//
	//	+++++|a+++++++|b++++|c++++++++++|d++++++++++++|.
	//	+++++|001234++|345++|001234567++|00123456789++|.
	//语句1为了和语句0进行对比（flags不同）



	//语句2右对齐padding=2, tabwriter.AlignRight输出：
	//	a       b       c       d               .
	//	123     12345   1234567 123456789       .
	//	（如果你是采用的goland运行的话，会输出结果跟我的不一样，请采用系统自带的cmd命令go run xx.go运行,上同）
	//	如果padding是\t的话，则填充和字符加起来的总长度一定要求是参数tabwidth的最小倍数（上面我们设置了tabwidth为8）
	//  如果padding是非\t的话，则宽度是最小的字符+padding宽度，如下面
	//
	//	++++++++++++a++++b++++++++++c++++++++++++d.
	//	+++++++001234++345++001234567++00123456789.(通过补左来对齐右,minwidth最小宽度为5,加上padding的宽度2，所以001234前面是+++++++一共7个+号)
	//	+号补左（默认），但是内容必须保证右对齐（a和4，b和5，c和7以及d和9必须右对齐），上面刚好相反




}







func check_err(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

}
