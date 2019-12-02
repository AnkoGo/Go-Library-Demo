package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"log"
	"os"
	"time"
)

func main(){

//创建一个缓冲区以将我们的压缩文件内容写入其中
	buf := new(bytes.Buffer)
	// NewWriter返回一个新的Writer，将一个zip文件写入buf。
	w := zip.NewWriter(buf)
	//添加一些信息到压缩文件中去
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}
	for _, file := range files {
		// Create使用提供的名称将文件添加到zip文件中。
		//返回一个Writer，文件内容应写入该Writer中。
		//文件内容将使用Deflate方法压缩。
		//名称必须是相对路径：不得以驱动器号（例如C :）或前斜杠开头，并且只能使用正斜杠。 要创建目录而不是文件，请在名称后添加斜杠。
		//在下一次调用Create，CreateHeader或Close之前，必须将文件的内容写入io.Writer。
		//说白了相当于创建一个压缩文件的元数据信息准备接受写入
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		//元数据文件写入压缩文件的相关过程以及状态信息
		//type fileWriter struct {
		//	*header
		//	zipw      io.Writer	//真正的压缩写入器
		//	rawCount  *countWriter
		//	comp      io.WriteCloser
		//	compCount *countWriter
		//	crc32     hash.Hash32
		//	closed    bool
		//}

		//type countWriter struct {
		//	w     io.Writer	//写入器
		//	count int64		//写入的字节数
		//}

		//type Writer struct {
		//	cw          *countWriter	//写入器以及写入的字节数相关的对象
		//	dir         []*header		//写入器的写入哪些元数据
		//	last        *fileWriter		//元数据文件写入压缩文件的相关过程以及状态信息
		//	closed      bool			//关闭写入器的状态信息
		//	compressors map[uint16]Compressor//压缩器类
		//	comment     string			//有关压缩的注释信息
		//
		//}
		_, err = f.Write([]byte(file.Body)) //这里才是压缩文件真正写入的操作
		if err != nil {
			log.Fatal(err)
		_, err = f.Write([]byte(file.Body)) //这里才是压缩文件真正写入的操作

	}
	//	关闭压缩写入器
	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.Bytes())
	fmt.Println(buf.String())
	//输出：
	//	[80 75 3 4 20 0 8 0 8 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 10 0 0 0 114 101 97 100 109 101 46 116 120 116 10 201 200
	//	44 86 72 44 74 206 200 44 75 85 72 206 207 43 73 204 204 43 86 40 206 207 77 85 40 73 173 40 81 72 203 204 73 45
	//	214 3 4 0 0 255 255 80 75 7 8 208 190 233 29 44 0 0 0 38 0 0 0 80 75 3 4 20 0 8 0 8 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 10 0 0 0 103 111 112 104 101 114 46 116 120 116 114 207 47 200 72 45 82 200 75 204 77 45 182 226 114 79 205 47
	//	74 79 5 81 105 105 69 169 149 92 238 249 121 85 249 128 0 0 0 255 255 80 75 7 8 83 121 181 152 39 0 0 0 35 0 0 0 80
	//	75 3 4 20 0 8 0 8 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 8 0 0 0 116 111 100 111 46 116 120 116 114 79 45 81 72 204 203
	//	204 77 204 81 200 72 204 75 201 201 204 75 87 200 201 76 78 205 75 78 213 227 10 47 202 44 73 85 200 205 47 74 85
	//	72 173 72 204 45 200 73 45 214 3 4 0 0 255 255 80 75 7 8 159 125 221 59 55 0 0 0 49 0 0 0 80 75 1 2 20 0 20 0 8 0 8
	//	0 0 0 0 0 208 190 233 29 44 0 0 0 38 0 0 0 10 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 114 101 97 100 109 101 46 116 120 116
	//	80 75 1 2 20 0 20 0 8 0 8 0 0 0 0 0 83 121 181 152 39 0 0 0 35 0 0 0 10 0 0 0 0 0 0 0 0 0 0 0 0 0 100 0 0 0 103 111 112
	//	104 101 114 46 116 120 116 80 75 1 2 20 0 20 0 8 0 8 0 0 0 0 0 159 125 221 59 55 0 0 0 49 0 0 0 8 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 195 0 0 0 116 111 100 111 46 116 120 116 80 75 5 6 0 0 0 0 3 0 3 0 166 0 0 0 48 1 0 0 0 0]
	//	PK
	//	readme.txt
	//	��,VH,J��,KUH��+I��+V(��MU(I�(QH��I-�  ��PKо�,   &   PK
	//	gopher.txtr�/�H-R�K�M-��rO�/JOQiiE��\��yU��   ��PKSy��'   #   PK                   todo.txtrO-QH���M�Q�H�K���KW��LN�KN��
	//	/�,IU��/JUH�H�-�I-�  ��PK�}�;7   1   PK      о�,   &
	//	readme.txtPK      Sy��'   #
	//	d   gopher.txtPK      �}�;7   1               �   todo.txtPK      �   0

	fmt.Println()
	fmt.Println("--------压缩字符串内容，写入到文件中去----------")

	//// 如果成功，则可以将返回的File上的方法用于I/O。 关联的文件描述符的模式为O_RDWR。
	////如果有错误，它将是* PathError类型。
	//file, err := os.Create("main/zip/zip_01.zip")
	//w1 := zip.NewWriter(file)
	//
	////添加一些信息到压缩文件中去
	//var files1 = []struct {
	//	Name, Body string
	//}{
	//	{"readme.txt", "This archive contains some text files."},
	//	{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
	//	{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	//}
	//for _, file := range files1 {
	//
	//	f, err := w1.Create(file.Name)
	//	check_err_zip(err)
	//	_, err = f.Write([]byte(file.Body))//这里才是压缩文件真正写入的操作
	//	check_err_zip(err)
	//	//err = w.Flush()
	//}
	////	关闭压缩写入器
	//err1 := w1.Close()
	//check_err_zip(err1)
	//defer func() {file.Close()}()
	//
	//
	////输出：
	////请看路径上面生成了一个zip_01.zip的文件

	//fmt.Println()
	//fmt.Println("--------压缩文件a，b的内容，写入到c压缩文件中去（相当于我们日常压缩某个文件）----------")
	//
	////因为下面的需求所以我们先将上面的注释掉
	//file2, err := os.Create("main/zip/zip_02.zip")
	//info1, err := f1.Stat()
	//check_err_zip(err)
	//f1_src_byte:=make([]byte,info1.Size())
	//n1, err := f1.Read(f1_src_byte)
	//check_err_zip(err)
	//fmt.Println("读取到第一个文件的字节数为：",n1)
	//
	//f2, err := os.Open("main/tar_data/text2.txt")
	//check_err_zip(err)
	//info2, err := f2.Stat()
	//check_err_zip(err)
	//f2_src_byte:=make([]byte,info2.Size())
	//n2, err := f2.Read(f2_src_byte)
	//check_err_zip(err)
	//fmt.Println("读取到第一个文件的字节数为：",n2)
	//
	//
	////新建结构体接收打开的文件内容，当然你也可以不用新建结构体来接收，而是直接传进去Create（）和Write（）中去，上同下同
	//var files2 = []struct {
	//	Name, Body string
	//}{
	//	{info1.Name(), string(f1_src_byte)},
	//	{info2.Name(), string(f2_src_byte)},
	//}
	//for _, file := range files2 {
	//	// Create使用提供的名称将文件添加到zip文件中。
	//	//返回一个Writer，文件内容应写入该Writer中。
	//	//文件内容将使用Deflate方法压缩。
	//	//名称必须是相对路径：不得以驱动器号（例如C :）或前斜杠开头，并且只能使用正斜杠。 要创建目录而不是文件，请在名称后添加斜杠。
	//	//在下一次调用Create，CreateHeader或Close之前，必须将文件的内容写入io.Writer。
	//
	//	//大白话：这个相当于在面对压缩文件中新建一个空的元数据文件做好写入的准备，同时返回写入器
	//	w, err := w2.Create(file.Name)
	//	check_err_zip(err)
	//	_, err = w.Write([]byte(file.Body))//写入的操作，但是他并没有真正的写入，真正写入的是close()方法，这里相当于读取源文件的部分字节到zip
	//	// 模块的类型type writeBuf []byte中去，并做好写入的准备，writeBuf相当于搬运工。
	//	check_err_zip(err)
	//
	//}
	////刷新将所有缓冲的数据刷新到基础写入器。
	////通常不需要调用Flush； 调用Close就足够了。
	////err = w.Flush()
	////check_err_zip(err)
	//
	//
	//defer func() {
	//	// Close通过写入中央目录完成zip文件的写入。
	//	//它不会关闭基础编写器。
	//	//	关闭压缩写入器，这个东西不能放到defer的下面中去，因为close()刷新所有的缓存到对应的目的压缩文件中去，如果把它先于file2，f1，f2关闭的话，那么就无法获取数据了
	//	//	如果不调用这个close()的话，那么压缩文件中将不会有任何的数据，w.Flush()会自动调用flush()，所以上面的flush()不是必须的
	//	err2 := w2.Close()//这里从zip模块的类型type writeBuf []byte中获取到数据进行真正的写入到目的压缩文件中去，相对于下面的文件关闭，这个的关闭必须在最上面
	//	check_err_zip(err2)
	//	e := file2.Close()
	//	check_err_zip(e)
	//	e = f1.Close()
	//	check_err_zip(e)
	//	e = f2.Close()
	//	check_err_zip(e)
	//
	//}()//在上面我们没有进行关闭一些资源，这是不大好的，
	//
	////输出：
	////请看路径上面生成了一个zip_02.zip的文件，通过对比我们发现了原来的额text1和text2文件的总大小为1.24KB，
	//// 现在我们看生成的zip_02.zip是1.10KB，确是达到了压缩的效果

	//fmt.Println()
	//fmt.Println("--------【通过CreateHeader()方法实现】压缩文件a，b的内容，写入到c压缩文件中去（相当于我们日常压缩某个文件）----------")
	////注释掉上面的额代码，
	//
	//file3, err_header := os.Create("main/zip/aa.zip")//如果是采用下面的方式会覆盖这个zip文件
	////file3, err_header := os.OpenFile("main/zip/aa.zip",os.O_APPEND,777)// main/zip/aa.zip已经存在
	//check_err_zip(err_header)
	//w3 := zip.NewWriter(file3)
	////info_file3, err := file3.Stat()
	////check_err_zip(err)
	////fmt.Println(info_file3.Size())
	////w3.SetOffset(info_file3.Size())//移动指针到末尾，似乎并不起作用
	//f11_src_byte:=make([]byte,info11.Size())
	//n11, err := f11.Read(f11_src_byte)
	//check_err_zip(err)
	//fmt.Println("读取到第一个文件的字节数为：",n11)
	//
	//f22, err := os.Open("main/tar_data/text2.txt")
	//check_err_zip(err)
	//info22, err := f22.Stat()
	//check_err_zip(err)
	//f22_src_byte:=make([]byte,info22.Size())
	//n22, err := f22.Read(f22_src_byte)
	//check_err_zip(err)
	//fmt.Println("读取到第一个文件的字节数为：",n22)
	//
	////从2个文件中获取到header信息
	//fhd11, e := zip.FileInfoHeader(info11)
	//check_err_zip(e)
	//
	//fhd22, e := zip.FileInfoHeader(info22)
	//check_err_zip(e)
	//
	//// CreateHeader使用提供的FileHeader作为文件元数据将文件添加到zip存档中。 作家拥有fh的所有权，并可以对其字段进行更改。 调用CreateHeader后，调用者不得修改fh。
	////这将返回一个Writer，文件内容应写入该Writer中。
	////在下一次调用Create，CreateHeader或Close之前，必须将文件的内容写入io.Writer。
	//writer11, e := w3.CreateHeader(fhd11)
	//check_err_zip(e)
	//_, err = writer11.Write(f11_src_byte)//这句话必须紧跟上面的那句话，不可以分开，
	//check_err_zip(err)
	//
	//writer22, e := w3.CreateHeader(fhd22)//因为这里会刷新上面的w3.CreateHeader(fhd11)
	//check_err_zip(e)
	//_, err = writer22.Write(f22_src_byte)
	//check_err_zip(err)
	//
	//defer func() {
	//	err2 := w3.Close()
	//	check_err_zip(err2)
	//	e := file3.Close()
	//	check_err_zip(e)
	//	e = f11.Close()
	//	check_err_zip(e)
	//	e = f22.Close()
	//	check_err_zip(e)
	//
	//}()//
	////输出：
	////a.zip中的源文件中的文件会被覆盖，不知道为什么，总之压缩之后的zip文件反而变大了，然后zip里面的文件是你新写入的文件
	////通过CreateHeader（）方法反而产生压缩文件会变大
	////目前还不知道怎么对一个已经存在的压缩文件进行追加文件，比如压缩文件里面有2个文件，我们希望通过追加一个文件来达到压缩文件里面有3个文件的效果

	fmt.Println()
	fmt.Println("--------关于获取FileHeader对象的一些信息----------")

	// FileHeader描述zip文件中的文件。
	//有关详细信息，请参见zip规范。
	//type FileHeader struct {
	//	// Name是文件名。
	//	//它必须是相对路径，不能以驱动器号（例如“ C：”）开头，并且必须使用正斜杠而不是反斜杠。 斜杠末尾表示此文件是目录，不应包含任何数据。
	//	//读取zip文件时，直接从zip文件填充“名称”字段，并且未验证其正确性。
	//	//调用者有责任对其进行适当的清理，包括规范化斜杠方向，验证路径是否相对以及防止通过文件名（“ ../../../”）遍历。
	//	Name string
	//
	//	//Comment是任何短于64KiB的用户定义的字符串。
	//	Comment string
	//
	//	// NonUTF8表示名称和注释未使用UTF-8编码。
	//	//根据规范，允许的唯一其他编码应该是CP-437，但是从历史上看，许多ZIP读取器都将Name和Comment解释为系统的本地字符编码。
	//	//仅当用户打算为特定的本地化区域编码不可移植的ZIP文件时，才应设置此标志。 否则，Writer会自动为有效的UTF-8字符串设置ZIP格式的UTF-8标志。
	//	NonUTF8 bool
	//
	//	//读取时，扩展时间戳记优于旧版MS-DOS日期字段，并且时间之间的偏移量用作时区。
	//	//如果仅存在MS-DOS日期，则将时区假定为UTC。
	//	//写入时，始终会发出扩展的时间戳（与时区无关）。 旧版MS-DOS日期字段是根据修改时间的位置进行编码的。
	//	ModifiedTime uint16 //不推荐使用：旧版MS-DOS日期； 改用Modified。
	//	ModifiedDate uint16 //不推荐使用：旧版MS-DOS时间； 改用Modified。
	//
	//	CRC32              uint32
	//	CompressedSize     uint32 //不推荐使用：改用CompressedSize64。
	//	UncompressedSize   uint32 //不推荐使用：请改用UncompressedSize64。
	//	CompressedSize64   uint64
	//	UncompressedSize64 uint64
	//	Extra              []byte
	//	ExternalAttrs      uint32 //含义取决于CreatorVersion
	//}

	fileinfo, err := os.Stat("main/zip/text1.txt") //我们试下给一个正常的文件，而不是一个压缩过后的文件看下
	//fileinfo, err := os.Stat("main/zip/aa.zip")//一个压缩过后的文件，应该是没什么不同的！都是文件信息。压缩文件也是文件。不过记住Fileheader一般是应用于压缩文件才是对的！
	if err != nil {
		fmt.Println(err)
	}
	// FileHeader描述zip文件中的文件。
	//有关详细信息，请参见zip规范。
	fileheader, err := zip.FileInfoHeader(fileinfo)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Name:", fileheader.Name)
	fmt.Println("Comment:", fileheader.Comment) //不知道为什么无法输出注释，aa.zip中是有注释的！
	fmt.Println("CompressedSize64:", fileheader.CompressedSize64)
	fmt.Println("NonUTF8:", fileheader.NonUTF8)
	fmt.Println("CRC32:", fileheader.CRC32)
	fmt.Println("Method:", fileheader.Method)
	fmt.Println("CreatorVersion:", fileheader.CreatorVersion)
	fmt.Println("Extra:", fileheader.Extra)
	//上面仅仅是展示了很少的一些东西，还有更多的属性我们没展示

	// Mode返回FileHeader的权限和模式位。
	fmt.Println("Mode：", fileheader.Mode())

	// SetMode更改FileHeader的权限和模式位。
	// FileMode表示文件的模式和权限位。
	//这些位在所有系统上都具有相同的定义，因此有关文件的信息可以从一个系统移动到另一个系统。 并非所有位都适用于所有系统。
	//唯一需要的位是目录的ModeDir。

	fileinfo, err := os.Stat("main/zip/text1.txt") //我们试下给一个正常的文件，而不是一个压缩过后的文件看下
	//这9个最低有效位是标准的Unix rwxrwxrwx权限。
	//这些位的值应被视为公共API的一部分，可以在有线协议或磁盘表示中使用：尽管可以添加新的位，但不得更改它们。

	//下面列出的是内置的权限和模式位
	//const (
	//	//单个字母是String方法的格式使用的缩写。
	//	ModeDir        FileMode = 1 << (32 - 1 - iota) // d: is a directory
	//	ModeAppend                                     // a: append-only
	//	ModeExclusive                                  // l: exclusive use 专用
	//	ModeTemporary                                  // T: temporary file; Plan 9 only  临时文件； 仅计划9
	fmt.Println("Name:", fileheader.Name)
	fmt.Println("Comment:", fileheader.Comment) //不知道为什么无法输出注释，aa.zip中是有注释的！
	fmt.Println("CompressedSize64:", fileheader.CompressedSize64)
	fmt.Println("NonUTF8:", fileheader.NonUTF8)
	fmt.Println("CRC32:", fileheader.CRC32)
	fmt.Println("Method:", fileheader.Method)
	fmt.Println("CreatorVersion:", fileheader.CreatorVersion)
	fmt.Println("Extra:", fileheader.Extra)
	//	ModeIrregular                                  // ?: non-regular file; nothing else is known about this file  非常规文件； 该文件没有其他信息
	//
	//	//类型位的掩码。 对于常规文件，将不会设置任何文件。
	fmt.Println("Mode：", fileheader.Mode())
	//
	//	ModePerm FileMode = 0777 // Unix权限位
	//)

	fileheader.SetMode(os.FileMode(0b111110110)) //我这里没采用内置的，而是自己定义
	fmt.Println("Mode：", fileheader.Mode())

	fmt.Println("-----通过header获取zip文件的FileInfo信息----")
	//下面是通过header获取文件的FileInfo信息
	fmt.Println(".FileInfo().Name()：", fileheader.FileInfo().Name())
	fmt.Println(".FileInfo().Size()：", fileheader.FileInfo().Size())
	fmt.Println(".FileInfo().Mode()：", fileheader.FileInfo().Mode())
	fmt.Println(".FileInfo().IsDir()：", fileheader.FileInfo().IsDir())
	fmt.Println(".FileInfo().ModTime()：", fileheader.FileInfo().ModTime().Local())
	fmt.Println(".FileInfo().Sys()：", fileheader.FileInfo().Sys())

	// ModTime使用旧版ModifiedDate和ModifiedTime字段以UTC返回修改时间。
	//不推荐使用：改用Modified。
	//fileheader.Modified和fileheader.SetMode()才是现在和未来的api,现在完全替代了下面即将弃用的api,他们的用法完全一致，不再展示。
	fmt.Println(fileheader.ModTime().Local()) //UTC时间,必须调用.Local()生成本地的时间
	// SetModTime将Modified，ModifiedTime和ModifiedDate字段设置为UTC中的给定时间。
	//不推荐使用：改用Modified。
	fileheader.SetModTime(time.Now().AddDate(1, 1, 1))
	fmt.Println(fileheader.ModTime().Local()) //UTC时间,必须调用.Local()生成本地的时间

	//如果是正常文件的输出：
	//	--------关于获取FileHeader对象的一些信息----------
	//	Name: text1.txt
	//	Comment:
	//	CompressedSize64: 0
	//	NonUTF8: false
	//	CRC32: 0
	//	Method: 0
	fileheader.SetMode(os.FileMode(0b111110110)) //我这里没采用内置的，而是自己定义
	fmt.Println("Mode：", fileheader.Mode())
	//	Mode： -rw-rw-rw-
	//	-----通过header获取zip文件的FileInfo信息----
	//	.FileInfo().Name()： text1.txt
	fmt.Println(".FileInfo().Name()：", fileheader.FileInfo().Name())
	fmt.Println(".FileInfo().Size()：", fileheader.FileInfo().Size())
	fmt.Println(".FileInfo().Mode()：", fileheader.FileInfo().Mode())
	fmt.Println(".FileInfo().IsDir()：", fileheader.FileInfo().IsDir())
	fmt.Println(".FileInfo().ModTime()：", fileheader.FileInfo().ModTime().Local())
	fmt.Println(".FileInfo().Sys()：", fileheader.FileInfo().Sys())
	//	2020-11-17 23:13:00 +0800 CST
	//如果是压缩文件的输出：
	//	--------关于获取FileHeader对象的一些信息----------
	//	Name: aa.zip
	//	Comment:（，不知道为什么获取不到注释信息）
	//	CompressedSize64: 0
	//	NonUTF8: false
	//	CRC32: 0
	//	Method: 0
	//	CreatorVersion: 768
	//	Mode： -rw-rw-rw-
	//	Mode： -rwxrw-rw-
	//	-----通过header获取zip文件的FileInfo信息----
	//	.FileInfo().Name()： aa.zip
	//	.FileInfo().Size()： 1582
	//	.FileInfo().Mode()： -rwxrw-rw-
	//	.FileInfo().IsDir()： false
	//	.FileInfo().ModTime()： 2019-10-16 19:00:47.2557733 +0800 CST（UTC时间是2019-10-16 11:00:47.2557733 +0000 UTC）
	//	.FileInfo().Sys()： &{aa.zip  false 768 0 0 0 2019-10-16 11:00:47.2557733 +0000 UTC 22551 20304 0 0 1582 0 1582 [] 2180382720}
	//	2019-10-16 19:00:46 +0800 CST
	//	2020-11-17 23:08:50 +0800 CST

}

func check_err_zip(err2 error) {
	if err != nil {
		log.Fatal(err)
	}
}
func check_err_zip(err2 error) {
