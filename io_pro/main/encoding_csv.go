package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
)

func main446677() {
	fmt.Println("-----------------如何写csv之Write()------------------------")
	buffer_csv := new(bytes.Buffer)
	writer := csv.NewWriter(buffer_csv)
	Err_write := writer.Write([]string{"aaa", "bbb", "ccc", "eee", "ddd", "ggg",})
	checkErr_1(Err_write)
	writer.Flush()//记得这个，否则不会有效果

	fmt.Printf("%q\n",buffer_csv)//"aaa,bbb,ccc,eee,ddd,ggg\n",必须使用%q才可以输出转义的东西，%v是不可以的
	fmt.Println(buffer_csv.Bytes())
	fmt.Println(buffer_csv.String())
	fmt.Println("=======")
	//输出：
	//	"aaa,bbb,ccc,eee,ddd,ggg\n"
	//	[97 97 97 44 98 98 98 44 99 99 99 44 101 101 101 44 100 100 100 44 103 103 103 10]
	//	aaa,bbb,ccc,eee,ddd,ggg
	//
	//	=======

	fmt.Println("-----------------如何写csv之WriteAll()------------------------")

	buffer_csv111 := new(bytes.Buffer)
	writer111 := csv.NewWriter(buffer_csv111)//底层是采用bufio.Writer来创建的writer
	//指定字段分隔符（由NewWriter设置为'，'），默认是逗号作为单个记录的元素分隔符，
	// 一个csv可以有单条或者多条的记录，单个记录可以包含单个或者多个字符串，记录里面的元素之间可以指定分隔符
	//writer111.Comma='-'

	//writer111.UseCRLF=true//使用\r\n作为行终止符为true,默认是采用\n
	Err_write111 := writer111.WriteAll([][]string{[]string{"aaa", "bbb", "ccc", "eee", "ddd", "ggg"},
												[]string{"111","222","333"},
												[]string{"fff","hhh","444"},
												[]string{"ff我","h的","世界666"},
										})
	checkErr_1(Err_write111)
	writer.Flush()//记得这个，否则不会有效果

	fmt.Printf("%q\n",buffer_csv111)
	fmt.Println(buffer_csv111.Bytes())
	fmt.Println(buffer_csv111.String())
	fmt.Println("=======")
	//输出：
	//	"aaa,bbb,ccc,eee,ddd,ggg\n111,222,333\nfff,hhh,444\nff我,h的,世界666\n"
	//	[97 97 97 44 98 98 98 44 99 99 99 44 101 101 101 44 100 100 100 44 103 103 103 10 49
	//	49 49 44 50 50 50 44 51 51 51 10 102 102 102 44 104 104 104 44 52 52 52 10 102 102 230
	//	136 145 44 104 231 154 132 44 228 184 150 231 149 140 54 54 54 10]
	//	aaa,bbb,ccc,eee,ddd,ggg
	//	111,222,333
	//	fff,hhh,444
	//	ff我,h的,世界666
	//
	//	=======

	//指定分隔符为'-'的话输出如下：
	//	"aaa-bbb-ccc-eee-ddd-ggg\n111-222-333\nfff-hhh-444\nff我-h的-世界666\n"
	//	[97 97 97 45 98 98 98 45 99 99 99 45 101 101 101 45 100 100 100 45 103 103 103 10 49 49 49 45 50 50 50 45 51 51 51 10 102 102 102 45 104 104 104 45 52 52 52 10 102 102 230 136 145 45 104 231 154 132 45 228 184 150 231 149 140 54 54 54 10]
	//	aaa-bbb-ccc-eee-ddd-ggg
	//	111-222-333
	//	fff-hhh-444
	//	ff我-h的-世界666
	//
	//	=======

	//如果指定行终止符为"\r\n"的话的输出如下：
	//	"aaa,bbb,ccc,eee,ddd,ggg\r\n111,222,333\r\nfff,hhh,444\r\nff我,h的,世界666\r\n"
	//	[97 97 97 44 98 98 98 44 99 99 99 44 101 101 101 44 100 100 100 44 103 103 103 13 10 49 49 49 44 50 50 50 44 51 51 51 13 10 102 102 102 44 104 104 104 44 52 52 52 13 10 102 102 230 136 145 44 104 231 154 132 44 228 184 150 231 149 140 54 54 54 13 10]
	//	aaa,bbb,ccc,eee,ddd,ggg
	//	111,222,333
	//	fff,hhh,444
	//	ff我,h的,世界666
	//
	//	=======

	fmt.Println("-----------------如何读取csv之Read()------------------------")

	//阅读器从CSV编码文件中读取记录。
	//由NewReader返回，阅读器期望输入符合RFC 4180。
	//在第一次调用Read或ReadAll之前，可以更改导出的字段以自定义详细信息。
	//阅读器将其输入中的所有\r\n序列转换为普通\n，包括多行字段值在内，以便返回的数据不取决于输入文件使用哪种行尾约定。
	Reader_csv := csv.NewReader(buffer_csv)
	Reader_csv.FieldsPerRecord=-1//这里设置不设置都是一样的，因为只有一条的记录，我们判断读取时候是否是EOF来进行是否读取完成
	//读取从r读取一条记录（一片字段）。
	//如果记录具有意外的字段数，则Read会返回记录以及错误ErrFieldCount。
	//除这种情况外，Read总是返回非空记录或非空错误，但不会同时返回两者。
	//如果没有剩余要读取的数据，则Read返回nil io.EOF。
	//如果ReuseRecord为true，则可以在多个Read调用之间共享返回的slice。

	record, Err_reader := Reader_csv.Read()
	checkErr_1(Err_reader)
	fmt.Println(record)//

	Reader_csv000 := csv.NewReader(buffer_csv)
	record000, Err_reader000 := Reader_csv000.Read()
	checkErr_1(Err_reader000)
	fmt.Println(record000)

	//输出：
	//	[aaa bbb ccc eee ddd ggg]
	//	发生了错误： EOF
	//	[]

	fmt.Println("----上面是单条记录的读取，下面是对多条记录的读取----")

	Reader_csv111 := csv.NewReader(buffer_csv111)
	// FieldsPerRecord是每条记录的预期字段数。
	//如果FieldsPerRecord为正，则Read要求每个记录都具有给定的字段数。如果FieldsPerRecord为0，则Read将其设置为第一条记录中的字段数，
	// 以便将来的记录必须具有相同的字段数。如果FieldsPerRecord为负，则不进行检查，并且记录的字段数可能可变。
	//
	Reader_csv111.FieldsPerRecord=-1//必须设置这个为负数，因为我的每条记录的元素个数都是不一致的，导致下面会出错，假如我们的记录中的字段数都是一样个数的话，那么我们是不用设置这个的
	//record111, Err_reader111 := Reader_csv111.Read()
	//checkErr_1(Err_reader111)
	//fmt.Println(record111)
	//fmt.Println(Reader_csv111.FieldsPerRecord)//看来是固定不变的
	//
	//fmt.Println("----")
	//record222, Err_reader222 := Reader_csv111.Read()
	//checkErr_1(Err_reader222)
	//fmt.Println(record222)
	//fmt.Println(Reader_csv111.FieldsPerRecord)//看来是固定不变的
	//
	//fmt.Println("----")
	//record333, Err_reader333 := Reader_csv111.Read()
	//checkErr_1(Err_reader333)
	//fmt.Println(record333)
	//fmt.Println(Reader_csv111.FieldsPerRecord)
	//
	//fmt.Println("----")
	//record444, Err_reader444 := Reader_csv111.Read()
	//checkErr_1(Err_reader444)
	//fmt.Println(record444)
	//
	//fmt.Println("----")
	//record555, Err_reader555 := Reader_csv111.Read()
	//checkErr_1(Err_reader555)
	//fmt.Println(record555)
	//输出：
	//	----上面是单条记录的读取，下面是对多条记录的读取----
	//	[aaa bbb ccc eee ddd ggg]
	//	----
	//	[111 222 333]
	//	----
	//	[fff hhh 444]
	//	----
	//	[ff我 h的 世界666]
	//	----
	//	发生了错误： EOF
	//	[]

	fmt.Println("下面我们采用ReadAll()来一次性读取多条数据")
	//不过我们需要注释掉上面的read(),因为我们读取的是buf，所以只能读取一次，不能多次读取,
	// 但是同理我们不能注释掉FieldsPerRecord=-1，否则无法正常的读取
	records, Err_readall := Reader_csv111.ReadAll()
	checkErr_1(Err_readall)

	fmt.Println(records)
	//输出：
	//	下面我们采用ReadAll()来一次性读取多条数据
	//	[[aaa bbb ccc eee ddd ggg] [111 222 333] [fff hhh 444] [ff我 h的 世界666]]

	//下面我们对csv.Reader这个类的其他字段进行打印查看
	fmt.Println("下面我们对csv.Reader这个类的其他字段进行打印查看")
	//Comma            rune // 字段分隔符（NewReader将之设为','）
	//Comment          rune // 一行开始位置的注释标识符
	//FieldsPerRecord  int  // 每条记录期望的字段数
	//LazyQuotes       bool // 允许懒引号
	//TrailingComma    bool // 忽略，出于后端兼容性而保留(这个字段不建议使用了)
	//TrimLeadingSpace bool // 去除前导的空白
	fmt.Println(Reader_csv111.Comma)
	fmt.Println(Reader_csv111.Comment)
	fmt.Println(Reader_csv111.ReuseRecord)
	fmt.Println(Reader_csv111.FieldsPerRecord)
	fmt.Println(Reader_csv111.LazyQuotes)
	fmt.Println(Reader_csv111.TrimLeadingSpace)
	fmt.Println(Reader_csv111.TrailingComma)//(这个字段不建议使用了)
	//输出：
	//	44
	//	0
	//	false
	//	-1,无论读取多少次都不会改变这个值的，定死的值
	//	false
	//	false
	//	false

	fmt.Println("---------------csv的东西就这么多---------------------")

}
func checkErr_1(err error)  {
	if err !=nil{
		fmt.Println("发生了错误：",err)
	}
}















