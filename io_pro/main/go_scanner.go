package main

import (
	"bytes"
	"errors"
	"fmt"
	"go/scanner"
	//"go/token"	//不知道为什么导包这里会出错，事实上我们一定要导包，可能是跟其他文件冲突了，其他文件已经导入了这个包
)




func main34534() {

	//// ErrorList is a list of *Errors.
	//// The zero value for an ErrorList is an empty ErrorList ready to use.
	////
	//// ErrorList是*Error的切片列表。
	//// ErrorList的零值是可以使用的空ErrorList。
	//type ErrorList []*Error


	//// In an ErrorList, an error is represented by an *Error.
	//// The position Pos, if valid, points to the beginning of
	//// the offending token, and the error condition is described
	//// by Msg.
	////
	////在ErrorList中，错误由*Error表示。
	////位置Pos（如果有效）指向有问题的令牌token的开头，并且错误状态由Msg描述。
	//type Error struct {
	//	Pos token.Position
	//	Msg string
	//}

	//// Position describes an arbitrary source position
	//// including the file, line, and column location.
	//// A Position is valid if the line number is > 0.
	////
	////Position描述任意的源位置，包括文件，行和列的位置。
	////如果行号Line> 0，则位置有效。
	//type Position struct {
	//	Filename string // filename, if any	(//文件名（如果有）)
	//	Offset   int    // offset, starting at 0	(//偏移量，从0开始)
	//	Line     int    // line number, starting at 1	(//行号，从1开始)
	//	Column   int    // column number, starting at 1 (byte count)	(//列号，从1（字节数）开始)
	//}


	// PrintError is a utility function that prints a list of errors to w,
	// one error per line, if the err parameter is an ErrorList. Otherwise
	// it prints the err string.
	//
	// PrintError是一个实用程序函数，如果err参数是ErrorList，则将错误列表打印到w，每行一个错误。 否则，它将输出err字符串。

	buffer := new(bytes.Buffer)

	e1 := errors.New("这是一个自定义的错误信息111！")
	e2 := errors.New("这是一个自定义的错误信息222！")
	scanner.PrintError(buffer,e1)
	scanner.PrintError(buffer,e2)

	//token包还学习，目前先认识个别的api
	scan_err1:=scanner.Error{
		Pos: token.Position{"main3/compress_zlib.go",3,4,5},
		Msg: "scanner.Error错误信息111！",
	}//这里的信息都可以自定义，而无需验证是否是正确的！他就相当于一个容器而已

	scan_err2:=scanner.Error{
		Pos: token.Position{"main3/compress_zlib.go",6,7,8},
		Msg: "scanner.Error错误信息222！",
	}



	var err_ls=scanner.ErrorList{&scan_err1,&scan_err2}
	//ErrorList也实现了error接口
	scanner.PrintError(buffer,err_ls)
	fmt.Println(buffer.String())
	//输出：
	//	这是一个自定义的错误信息111！
	//	这是一个自定义的错误信息222！
	//	main3/compress_zlib.go:4:5: scanner.Error错误信息111！
	//	main3/compress_zlib.go:7:8: scanner.Error错误信息222！
	//每一个错误信息都自动会加上换行符




	fmt.Println("----------scanner.ErrorList对象下的一些方法--------------")
	// An ErrorList implements the error interface.
	//这个方法只会输出ErrorList切片中的第一个错误，如果不知一个错误，而会在后面提示(and xx more errors),xx表示ErrorList切片中还有多少个错误没打印出来的！
	//假如错误切片中没有任何的错误对象的话则会输出："no errors"
	fmt.Println(err_ls.Error())
	fmt.Println(err_ls.Error())
	var err_ls1=scanner.ErrorList{}
	fmt.Println(err_ls1.Error())

	//输出：
	//	main3/compress_zlib.go:4:5: scanner.Error错误信息111！ (and 1 more errors)
	//	main3/compress_zlib.go:4:5: scanner.Error错误信息111！ (and 1 more errors)
	//	no errors


	// Err returns an error equivalent to this error list.
	// If the list is empty, Err returns nil.
	// Err返回与该错误列表等效的错误。
	//	如果不为空，则等效于上面的.Error()方法
	//	如果列表为空，则Err返回nil。
	fmt.Println(err_ls.Err())
	fmt.Println(err_ls1.Err())
	//输出：
	//	main3/compress_zlib.go:4:5: scanner.Error错误信息111！ (and 1 more errors)
	//	<nil>




	fmt.Println(err_ls.Len())
	fmt.Println(err_ls1.Len())
	//输出：
	//	2
	//	0


	// Add adds an Error with given position and error message to an ErrorList.
	// Add将具有给定位置pos和错误消息msg的错误对象scanner.Error添加到scanner.ErrorList错误切片中去。
	Pos:=token.Position{"main3/compress_zlib.go",1,2,3}
	Msg:="scanner.Error错误信息333！"
	err_ls.Add(Pos,Msg)

	range_err:= func(err_ls scanner.ErrorList) {
		for _, v:= range err_ls {
			fmt.Println(v)
		}
		fmt.Println()
	}
	range_err(err_ls)
	//输出：
	//	main3/compress_zlib.go:4:5: scanner.Error错误信息111！
	//	main3/compress_zlib.go:7:8: scanner.Error错误信息222！
	//	main3/compress_zlib.go:2:3: scanner.Error错误信息333！
	//fmt.Println(err_ls)会输出err_ls.Error()的返回值，也就是第一个值


	//// Reset resets an ErrorList to no errors.
	//// Reset会将ErrorList重置为没有错误。
	//fmt.Printf("%p----%v\n",err_ls,err_ls.Len())
	//err_ls.Reset()
	//fmt.Printf("%p----%v\n",err_ls,err_ls.Len())
	////输出：
	////	0xc000004500----3
	////	0xc000004500----0



	//我们注释掉上面的代码，以防止对下面进行干扰

	// ErrorList implements the sort Interface.
	range_err(err_ls)
	err_ls.Swap(1,2)
	range_err(err_ls)
	//输出：
	//	main3/compress_zlib.go:4:5: scanner.Error错误信息111！
	//	main3/compress_zlib.go:7:8: scanner.Error错误信息222！
	//	main3/compress_zlib.go:2:3: scanner.Error错误信息333！
	//
	//	main3/compress_zlib.go:4:5: scanner.Error错误信息111！
	//	main3/compress_zlib.go:2:3: scanner.Error错误信息333！
	//	main3/compress_zlib.go:7:8: scanner.Error错误信息222！



	range_err(err_ls)
	// Sort sorts an ErrorList. *Error entries are sorted by position,
	// other errors are sorted by error message, and before any *Error
	// entry.
	// Sort对ErrorList进行排序。 *Error条目按位置position排序，其他错误则按错误消息排序，且在任何*Error条目之前。
	err_ls.Sort()
	range_err(err_ls)
	//输出：
	//	main3/compress_zlib.go:4:5: scanner.Error错误信息111！
	//	main3/compress_zlib.go:2:3: scanner.Error错误信息333！
	//	main3/compress_zlib.go:7:8: scanner.Error错误信息222！
	//
	//	main3/compress_zlib.go:2:3: scanner.Error错误信息333！
	//	main3/compress_zlib.go:4:5: scanner.Error错误信息111！
	//	main3/compress_zlib.go:7:8: scanner.Error错误信息222！




	//判断2个错误对象是否是升序排序
	//建议直接看源码就知道怎么比较了
	//	e := &p[i].Pos
	//	f := &p[j].Pos
	//	//注意，仅比较文件偏移量是不够的，因为偏移量不能反映修改后的行信息（通过//行注释）。所以下面没有进行比较偏移量而是比较了其他信息。
	//	if e.Filename != f.Filename {
	//		return e.Filename < f.Filename
	//	}
	//	if e.Line != f.Line {
	//		return e.Line < f.Line
	//	}
	//	if e.Column != f.Column {
	//		return e.Column < f.Column
	//	}
	//	return p[i].Msg < p[j].Msg
	fmt.Println(err_ls.Less(0,1))
	err_ls.Swap(0,1)
	fmt.Println(err_ls.Less(0,1))
	//输出：
	//	true
	//	false



	// RemoveMultiples sorts an ErrorList and removes all but the first error per line.
	// RemoveMultiples对ErrorList进行排序，并删除每行中除第一个错误以外的所有错误。
	fmt.Println()
	scan_err5:=scanner.Error{
		Pos: token.Position{"main3/compress_zlib.go",5,4,5},
		Msg: "scanner.Error错误信息第四行555！",
	}//这里的信息都可以自定义，而无需验证是否是正确的！他就相当于一个容器而已

	scan_err6:=scanner.Error{
		Pos: token.Position{"main3/compress_zlib.go",6,4,6},
		Msg: "scanner.Error错误信息第四行666！",
	}

	scan_err7:=scanner.Error{
		Pos: token.Position{"main3/compress_zlib.go",7,5,7},
		Msg: "scanner.Error错误信息第五行666！",
	}
	var err_ls_sameline=scanner.ErrorList{&scan_err5,&scan_err6,&scan_err7}

	range_err(err_ls_sameline)
	err_ls_sameline.RemoveMultiples()
	range_err(err_ls_sameline)
	//输出：
	//	main3/compress_zlib.go:4:5: scanner.Error错误信息第四行555！
	//	main3/compress_zlib.go:4:6: scanner.Error错误信息第四行666！
	//	main3/compress_zlib.go:5:7: scanner.Error错误信息第五行666！
	//
	//	main3/compress_zlib.go:4:5: scanner.Error错误信息第四行555！
	//	main3/compress_zlib.go:5:7: scanner.Error错误信息第五行666！



	//这个东西需要结合init()方法来讲解，先搁置
	//scanner.ErrorHandler()



	fmt.Println("------------scanner.Scanner对象------------------")
	//
	////// A Scanner holds the scanner's internal state while processing
	////// a given text. It can be allocated as part of another data
	////// structure but must be initialized via Init before use.
	//////
	//////Scanner扫描器在处理给定文本时会保留扫描器的内部状态。 可以将Scanner分配为另一个数据结构的一部分，但必须在使用前通过Init进行初始化。
	////type Scanner struct {
	////	// immutable state（//不可变状态）
	////	file *token.File  // source file handle（//源文件句柄）
	////	dir  string       // directory portion of file.Name()（// file.Name（）的目录部分）
	////	src  []byte       // source	（//来源）
	////	err  ErrorHandler // error reporting; or nil	(//错误报告或零)
	////	mode Mode         // scanning mode	(//扫描模式)
	////
	////	// scanning state(//扫描状态)
	////	ch         rune // current character	(//当前字符)
	////	offset     int  // character offset		(//字符偏移)
	////	rdOffset   int  // reading offset (position after current character)	(//读取偏移量（当前字符之后的位置）)
	////	lineOffset int  // current line offset	(//当前行的偏移量)
	////	insertSemi bool // insert a semicolon before next newline	(//在下一行之前插入分号)
	////
	////	// public state - ok to modify	(//公共状态-可以修改)
	////	ErrorCount int // number of errors encountered	(//遇到的错误数)
	////}
	//
	//
	//scan:=scanner.Scanner{
	//	ErrorCount: 1,
	//}
	//
	//
	////// A File is a handle for a file belonging to a FileSet.
	////// A File has a name, size, and line offset table.
	////// File是属于FileSet的文件的句柄。
	////// 文件具有名称，大小和行偏移量表。
	//////
	////type File struct {
	////	set  *FileSet
	////	name string // file name as provided to AddFile	（//提供给AddFile的文件名）
	////	base int    // Pos value range for this file is [base...base+size]	（//此文件的Pos值范围为[base ... base + size]）
	////	size int    // file size as provided to AddFile		（//提供给AddFile的文件大小）
	////
	////	// lines and infos are protected by mutex	（//行和信息受互斥锁保护）
	////	mutex sync.Mutex
	////	lines []int // lines contains the offset of the first character for each line (the first entry is always 0)（//行包含每行第一个字符的偏移量（第一个条目始终为0））
	////	infos []lineInfo
	////}
	//
	//
	//// Init通过将扫描器设置在src的开头来准备扫描器以标记文本src。 扫描仪使用文件集文件获取位置信息，并为每行添加行信息。
	////重新扫描同一文件时可以重新使用同一文件，因为已存在的行信息将被忽略。 如果文件大小与src大小不匹配，则Init会引起恐慌。
	////
	////如果遇到语法错误并且err不是nil，则调用Scan会调用错误处理程序err。 此外，对于遇到的每个错误，“扫描器”字段ErrorCount都会增加1。 模式参数确定如何处理注释。
	////
	////注意，如果文件的第一个字符有错误，则Init可能会调用err。
	//
	//scan.Init()

	//上面先注释掉，需要先学习go/token包


}



func check_err(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

}
