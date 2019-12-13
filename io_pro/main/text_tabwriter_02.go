package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)


func Example_elastic() {
	// Observe how the b's and the d's, despite appearing in the
	// second cell of each line, belong to different columns.
	//观察b和d尽管出现在每行的第二个单元格中，但它们如何属于不同的列。
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, '.', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "a\tb\tc")
	fmt.Fprintln(w, "aa\tbb\tcc")
	fmt.Fprintln(w, "aaa\t") // trailing tab（//尾随制表符），这行会导致下面的行中的列不再对齐上面了
	fmt.Fprintln(w, "aaaa\tdddd\teeee")
	fmt.Fprintln(w, "a\tb\tc")
	fmt.Fprintln(w, "aa\tbb\tcc")
	w.Flush()

	// output:
	//	....a|..b|c
	//	...aa|.bb|cc
	//	..aaa|
	//	.aaaa|.dddd|eeee
	//	....a|....b|c
	//	...aa|...bb|cc

	//假设注释掉第三行"aaa\t"，则会输出：
	//	....a|....b|c
	//	...aa|...bb|cc
	//	.aaaa|.dddd|eeee
	//	....a|....b|c
	//	...aa|...bb|cc
	//对比以上两次的输出可以发现，假设出现更宽的列的话，那么接下来的列宽的将会扩展，但是之前的不会扩展，还是以之前的宽度进行对齐
}

func Example_trailingTab() {
	// Observe that the third line has no trailing tab,
	// so its final cell is not part of an aligned column.
	//请注意，第三行没有尾随制表符，因此其最后一个单元格不是对齐列的一部分。
	const padding = 3


	// NewWriter allocates and initializes a new tabwriter.Writer.
	// The parameters are the same as for the Init function.
	// NewWriter分配并初始化一个新的tabwriter.Writer。
	//这些参数与Init函数的参数相同。
	//底层：return new(Writer).Init(output, minwidth, tabwidth, padding, padchar, flags)

	//多个flags标志属性之间用|老进行连接，表示一起发挥作用
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, '-', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "a\tb\taligned\t")//1
	fmt.Fprintln(w, "aa\tbb\taligned\t")//2
	fmt.Fprintln(w, "aaa\tbbb\tunaligned") //3// no trailing tab(//没有尾随制表符),这行会导致下面的行中的列不再对齐上面了
	fmt.Fprintln(w, "aaaa\tbbbb\taligneddddd\t")//4
	fmt.Fprintln(w, "aaaa\tbbbb\tali\t")//5
	w.Flush()

	// output:
	//	------a|------b|---aligned|
	//	-----aa|-----bb|---aligned|
	//	----aaa|----bbb|unaligned
	//	---aaaa|---bbbb|---aligneddddd|
	//	---aaaa|---bbbb|-----------ali|
	//事实上对不对其行中的列，底层是用一个width的宽度值随着扫描每个单元格的长度同步更新的，也就是width保持最大的值，其实别不存在出现断层的概念，
	//我只是为了便于表述才这样说的，比如，为了表述，我简称width为w(w并不是整个单元格的宽度，还要加上padding)
	//扫描第1行中的第三列时候：w=len("aligned")，因为w初始化时候为0，所以这里开始扫描时候肯定赋值给w,所以更新w,
	//扫描第2行中的第三列时候：w=len("aligned"),因为len("aligned")=len("aligned"),不更新w
	//扫描第3行中的第三列时候：w=len("unaligned"),因为len("aligned")<len("unaligned")，更新w
	//扫描第4行中的第三列时候：w=len("aligneddddd"),因为len("unaligned")<len("aligneddddd")，更新w
	//扫描第5行中的第三列时候：w=len("aligneddddd"),因为len("ali")<len("aligneddddd"),不更新w
	//上面的Example_elastic()函数同理，不再累叙
}

func main() {

	Example_elastic()
	fmt.Println()
	Example_trailingTab()


}









func check_err(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

}
