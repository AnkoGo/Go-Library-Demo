package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
)

type person struct {
	Name      string
	Age       int
	Isstudent bool
}

func main() {

	fmt.Println("-------------NewEncoder()创建编码器，.Encode编码数据成为json对象-------------------------")

	var P = person{
		Name:      "anko<>&",
		Age:       18,
		Isstudent: true,
	}

	dst_buffer_enc := new(bytes.Buffer)
	enc := json.NewEncoder(dst_buffer_enc)
	fmt.Println("编码前的buffer数据是：", dst_buffer_enc)

	// SetEscapeHTML指定是否应在JSON引用的字符串内转义有问题的HTML字符。
	//默认行为是将＆，<和>转义为\ u0026，\ u003c和\ u003e，以避免在将JSON嵌入HTML时可能出现的某些安全问题。
	//
	//在非HTML设置中，转义会干扰输出的可读性，SetEscapeHTML（false）会禁用此行为。
	//enc.SetEscapeHTML(true)//默认是这个转义的
	enc.SetEscapeHTML(false) //默认是这个转义的

	//编码将v的JSON编码写入流，后跟换行符。
	//有关将Go值转换为JSON的详细信息，请参见Marshal的文档。
	encode_err := enc.Encode(&P)
	check_EncDec_err(encode_err)

	fmt.Println("编码后的buffer数据是：", dst_buffer_enc)
	fmt.Println("编码后的buffer转为字节数据是：", dst_buffer_enc.Bytes())
	fmt.Println("编码后的buffer转为字符串数据是：", dst_buffer_enc.String())
	fmt.Println("====")
	//SetEscapeHTML(false)不转义特殊字符的输出：
	//	编码前的buffer数据是：
	//	编码后的buffer数据是： {"Name":"anko<>&","Age":18,"Isstudent":true}
	//
	//	编码后的buffer转为字节数据是： [123 34 78 97 109 101 34 58 34 97 110 107 111 60 62 38 34 44 34 65 103 101 34 58 49 56 44 34 73 115 115 116 117 100 101 110 116 34 58 116 114 117 101 125 10]
	//	编码后的buffer转为字符串数据是： {"Name":"anko<>&","Age":18,"Isstudent":true}
	//
	//	====
	//注意编码后json会在最后跟上换行符

	//SetEscapeHTML(true)（默认行为）转义特殊字符的输出：
	//	编码前的buffer数据是：
	//	编码后的buffer数据是： {"Name":"anko\u003c\u003e\u0026","Age":18,"Isstudent":true}
	//
	//	编码后的buffer转为字节数据是： [123 34 78 97 109 101 34 58 34 97 110 107 111 92 117 48 48 51 99 92 117 48 48 51 101 92 117 48 48 50 54 34 44 34 65 103 101 34 58 49 56 44 34 73 115 115 116 117 100 101 110 116 34 58 116 114 117 101 125 10]
	//	编码后的buffer转为字符串数据是： {"Name":"anko\u003c\u003e\u0026","Age":18,"Isstudent":true}
	//
	//	====

	fmt.Println("-------------NewDecoder()创建解码器，.Decode解码数据成为go对象-------------------------")

	// NewDecoder返回从r读取的新解码器。
	//解码器引入自己的缓冲，并且可能从r读取超出请求的JSON值的数据。
	dec := json.NewDecoder(dst_buffer_enc)

	var P1 person
	// Decode从输入中读取下一个JSON编码的值，并将其存储在v指向的值中。
	//有关将JSON转换为Go值的详细信息，请参见Unmarshal的文档。
	fmt.Println("编码前的buffer数据是：", dst_buffer_enc)
	deccode_err := dec.Decode(&P1)
	check_EncDec_err(deccode_err)

	fmt.Println("解码后的buffer数据是：", dst_buffer_enc)
	fmt.Printf("解码后的person结构体实例数据是：%#v\n", P1)
	fmt.Println("====")
	//SetEscapeHTML(false)不转义特殊字符的输出：
	//	编码前的buffer数据是： {"Name":"anko<>&","Age":18,"Isstudent":true}
	//
	//	解码后的buffer数据是：
	//	解码后的person结构体实例数据是：main.person{Name:"anko<>&", Age:18, Isstudent:true}
	//	====

	//SetEscapeHTML(true)（默认行为）转义特殊字符的输出：
	//	编码前的buffer数据是： {"Name":"anko\u003c\u003e\u0026","Age":18,"Isstudent":true}
	//
	//	解码后的buffer数据是：
	//	解码后的person结构体实例数据是：main.person{Name:"anko<>&", Age:18, Isstudent:true}
	//	====

	fmt.Println("-------------enc.SetIndent()设置json编码后的格式-------------------------")

	//enc.SetIndent("=","\t")//这样设置是很方便，但是无法解码这样的json格式
	enc.SetIndent(" ", "\t")

	var P3 = person{
		Name:      "anko111",
		Age:       20,
		Isstudent: false,
	}
	fmt.Println("编码前的buffer数据是：", dst_buffer_enc)
	encode_err111 := enc.Encode(&P3)
	check_EncDec_err(encode_err111)

	fmt.Println("编码后的buffer数据是：\n", dst_buffer_enc)
	fmt.Println("编码后的buffer转为字节数据是：", dst_buffer_enc.Bytes())
	fmt.Println("====")
	//enc.SetIndent("=","\t")输出：
	//	编码前的buffer数据是：（空字符串）
	//	编码后的buffer数据是：
	//	 {
	//	=	"Name": "anko111",
	//	=	"Age": 20,
	//	=	"Isstudent": false
	//	=}
	//
	//	编码后的buffer转为字节数据是： [123 10 61 9 34 78 97 109 101 34 58 32 34 97 110 107 111 49 49 49 34 44 10 61 9 34 65 103 101 34 58 32 50 48 44 10 61 9 34 73 115 115 116 117 100 101 110 116 34 58 32 102 97 108 115 101 10 61 125 10]
	//	====

	//enc.SetIndent(" ","\t")输出：
	//	编码前的buffer数据是：
	//	编码后的buffer数据是：
	//	{
	//		"Name": "anko111",
	//		"Age": 20,
	//		"Isstudent": false
	//	}
	//
	//	编码后的buffer转为字节数据是： [123 10 32 9 34 78 97 109 101 34 58 32 34 97 110 107 111 49 49 49 34 44 10 32 9 34 65 103 101 34 58 32 50 48 44 10 32 9 34 73 115 115 116 117 100 101 110 116 34 58 32 102 97 108 115 101 10 32 125 10]
	//	====

	fmt.Println("-------------dec.Buffered()返回解码器缓冲区中剩余数据的读取器-------------------------")

	dec111 := json.NewDecoder(dst_buffer_enc)

	var P111 person
	// Decode从输入中读取下一个JSON编码的值，并将其存储在v指向的值中。
	//有关将JSON转换为Go值的详细信息，请参见Unmarshal的文档。
	//line, _ := dst_buffer_enc.ReadBytes('0')//放一个不存在的值检测全部
	//fmt.Println("解码前的buffer数据是：\n",line)
	fmt.Println("解码前的buffer数据是：\n", dst_buffer_enc)

	//如果上面enc.SetIndent("=","\t")的话，那么这里将解码失败，不知道为什么提供了自定义编码格式的api但是却不提供解
	// 码这种编码格式的api,大佬也会有bug,我们以后写代码不要这样！提供什么东西要记得成双成对出现，不然api用法不对称！
	deccode_err111 := dec111.Decode(&P111)
	//check_EncDec_err(deccode_err111)
	if deccode_err111 != nil {
		fmt.Println("++++", deccode_err111)
	}
	// Buffered返回解码器缓冲区中剩余数据的读取器。 读取器在下一次调用Decode之前一直有效。
	buffered_Rd := dec.Buffered() //如果这句话出现在解码之前则缓存是满的

	ls_dec := make([]byte, 10)
	Rd_n, buffered_Rd_err := buffered_Rd.Read(ls_dec)
	check_EncDec_err(buffered_Rd_err)

	fmt.Println("解码器解码完成后的缓冲区中剩余数据字节数为：", Rd_n)
	fmt.Println("解码器解码完成后的缓冲区中剩余数据为：", ls_dec)
	fmt.Printf("解码后的person结构体实例数据是：%#v\n", P111)
	fmt.Println("====")
	//输出：
	//	解码前的buffer数据是：
	//	{
	//		"Name": "anko111",
	//		"Age": 20,
	//		"Isstudent": false
	//	}
	//
	//	解码器解码完成后的缓冲区中剩余数据字节数为： 1
	//	解码器解码完成后的缓冲区中剩余数据为： [10 0 0 0 0 0 0 0 0 0]
	//	解码后的person结构体实例数据是：main.person{Name:"anko111", Age:20, Isstudent:false}
	//	====

	fmt.Println()
	fmt.Println("-------------不使用编解码类实例而是使用便捷的封装函数json.Marshal()和json.Unmarshal()进行数据的编码解码-------------------------")
	// Marshal返回v的JSON编码。更多规则请点进去看文档,默认是转义的
	func_marshal, func_encode_err := json.Marshal(&P)
	check_EncDec_err(func_encode_err)

	fmt.Printf("要编码类实例为：%#v\n", P)
	fmt.Println("编码后且不转义的json数据为：", string(func_marshal))

	var P222 person
	func_decode_err := json.Unmarshal(func_marshal, &P222)
	check_EncDec_err(func_decode_err)

	fmt.Println("解码前不转义的json数据为：", string(func_marshal))
	fmt.Printf("解码不转义的json数据后的类实例为：%#v\n", P222)

	fmt.Println()
	fmt.Println("--------------json.HTMLEscape将未转义的json数据进行转义----------------------------")
	// HTMLEscape将JSON编码后的src附加到dst中，其中字符串文字中的<，>，＆，U + 2028和U + 2029字符更
	// 改为\u003c，\u003e，\u0026，\u2028，\u2029，以便JSON 可以安全地嵌入HTML<script>标记中。
	//由于历史原因，Web浏览器不支持<script>标记内的标准HTML转义，因此必须使用替代的JSON编码。

	//假如我们有一段没转义过的json数据，这时候我们可以自己转义这个json数据的！
	var P666 = person{
		Name:      "anko666<>&",
		Age:       18,
		Isstudent: true,
	}

	dst_buffer_enc666 := new(bytes.Buffer)
	enc666 := json.NewEncoder(dst_buffer_enc666)
	fmt.Println("编码前的json数据是：", dst_buffer_enc666)
	enc666.SetEscapeHTML(false) //设置输出不转义的json数据

	encode_err666 := enc666.Encode(&P666)
	check_EncDec_err(encode_err666)

	fmt.Println("没转义之前的json数据是：", dst_buffer_enc666)

	bf := new(bytes.Buffer)
	json.HTMLEscape(bf, dst_buffer_enc666.Bytes())

	fmt.Println("转义之后的json数据是：", bf)
	fmt.Println("======")
	//输出：
	//	编码前的buffer数据是：
	//	没转义之前的json数据是： {"Name":"anko666<>&","Age":18,"Isstudent":true}
	//
	//	转义之后的json数据是： {"Name":"anko666\u003c\u003e\u0026","Age":18,"Isstudent":true}
	//
	//	======


	fmt.Println()
	fmt.Println("-----------------json.Compact去除json中多余的空白字符（但是不转义）----------------------------")

	var P777 = person{
		Name:      "anko777<>&  ",
		Age:       18,
		Isstudent: true,
	}

	var P888 = person{
		Name:      "anko888<>&  ",
		Age:       28,
		Isstudent: false,
	}

	unit:=[]person{P777,P888}
	//str:=`[{"Name": "anko777<>&  ", "Age": 18, "Isstudent": true},{"Name": "anko888<>&  ", "Age": 28, "Isstudent": false}]`


	dst_buffer_enc777 := new(bytes.Buffer)
	enc777 := json.NewEncoder(dst_buffer_enc777)
	fmt.Println("编码前的json数据是：", dst_buffer_enc777)
	enc777.SetEscapeHTML(false) //设置输出不转义的json数据
	enc777.SetIndent(" ","\t")//设置前缀和空格

	//encode_err777 := enc777.Encode(&str)
	encode_err777 := enc777.Encode(&unit)
	check_EncDec_err(encode_err777)

	fmt.Printf("没转义且没去除空格之前的json数据是：\n%v", dst_buffer_enc777)

	//Compact函数将json编码的src中无用的空白字符剔除后写入dst。
	dst_buffer_Compact := new(bytes.Buffer)
	encode_err_Compact := json.Compact(dst_buffer_Compact, dst_buffer_enc777.Bytes())//Compact紧凑的意思
	check_EncDec_err(encode_err_Compact)

	fmt.Println("Compact仅去除json空格但是没转义的json数据是：", dst_buffer_Compact)
	//输出：
	//	编码前的json数据是：
	//	没转义且没去除空格之前的json数据是：
	//	[
	//		{
	//			"Name": "anko777<>&  ",
	//			"Age": 18,
	//			"Isstudent": true
	//		},
	//		{
	//			"Name": "anko888<>&  ",
	//			"Age": 28,
	//			"Isstudent": false
	//		}
	//	]
	//	Compact仅去除json空格但是没转义的json数据是： [{"Name":"anko777<>&  ","Age":18,"Isstudent":true},{"Name":"anko888<>&  ","Age":28,"Isstudent":false}]

	fmt.Println()
	fmt.Println("-----------------json.Indent展开缩进json数据，但是不转义----------------------------")
	//缩进将JSON编码的src的缩进形式附加到dst。
	// JSON对象或数组中的每个元素都从一个新的缩进行开始，该行以前缀开头，然后根据缩进嵌套嵌套一个或多个缩进副本。
	//附加到dst的数据不以前缀或任何缩进开头，以便更轻松地嵌入其他格式化的JSON数据中。
	//尽管删除了src开头的前导空格字符（空格，制表符，回车符，换行符），但保留了src末尾的结尾空格字符并将其复制到dst。
	//例如，如果src没有尾随空格，则dst都不会； 如果src以结尾的换行符结尾，则dst也将如此。
	//我们可以指明展开后的前缀和缩进字符
	fmt.Println("没展开之前的json字符串：",dst_buffer_Compact)

	dst_buffer_enc888 := new(bytes.Buffer)
	//encode_err_Indent := json.Indent(dst_buffer_enc888, dst_buffer_Compact.Bytes(), " ", "\t")
	encode_err_Indent := json.Indent(dst_buffer_enc888, dst_buffer_Compact.Bytes(), "=", "-")
	check_EncDec_err(encode_err_Indent)

	fmt.Printf("展开缩进之后的json字符串：\n%v",dst_buffer_enc888)
	//输出1：
	//	没展开之前的json字符串： [{"Name":"anko777<>&  ","Age":18,"Isstudent":true},{"Name":"anko888<>&  ","Age":28,"Isstudent":false}]
	//	展开缩进之后的json字符串：
	//	[
	//		{
	//			"Name": "anko777<>&  ",
	//			"Age": 18,
	//			"Isstudent": true
	//		},
	//		{
	//			"Name": "anko888<>&  ",
	//			"Age": 28,
	//			"Isstudent": false
	//		}
	//	]

	//输出2：
	//	没展开之前的json字符串： [{"Name":"anko777<>&  ","Age":18,"Isstudent":true},{"Name":"anko888<>&  ","Age":28,"Isstudent":false}]
	//	展开缩进之后的json字符串：
	//	[
	//	=-{
	//	=--"Name": "anko777<>&  ",
	//	=--"Age": 18,
	//	=--"Isstudent": true
	//	=-},
	//	=-{
	//	=--"Name": "anko888<>&  ",
	//	=--"Age": 28,
	//	=--"Isstudent": false
	//	=-}
	//	=]

	fmt.Println()
	fmt.Println("-----------------json.MarshalIndent类似于Marshal应用Indent格式化输出----------------------------")

	var P0 =person{
		Name:      "anko000<>&",
		Age:       30,
		Isstudent: false,
	}


	// MarshalIndent类似于Marshal，但应用Indent格式化输出。会转义特殊字符。
	//输出中的每个JSON元素都将从换行开始，以前缀开头，然后根据缩进嵌套嵌套一个或多个缩进副本。
	indent_json, encode_err_indent := json.MarshalIndent(&P0, " ", "\t")
	//indent, encode_err_indent := json.MarshalIndent(&P0, "=", "-")
	check_EncDec_err(encode_err_indent)

	fmt.Println("序列化后的字节切片：",indent_json)
	fmt.Printf("序列化后的字节切片转字符串：\n%v",string(indent_json))
	//输出1：
	//	序列化后的字节切片： [123 10 32 9 34 78 97 109 101 34 58 32 34 97 110 107 111 48 48 48 92 117 48 48 51 99 92 117 48 48 51 101 92 117 48 48 50 54 34 44 10 32 9 34 65 103 101 34 58 32 51 48 44 10 32 9 34 73 115 115 116 117 100 101 110 116 34 58 32 102 97 108 115 101 10 32 125]
	//	序列化后的字节切片转字符串：
	//	{
	//		"Name": "anko000\u003c\u003e\u0026",
	//		"Age": 30,
	//		"Isstudent": false
	//	}

	//输出2：
	//	序列化后的字节切片： [123 10 61 45 34 78 97 109 101 34 58 32 34 97 110 107 111 48 48 48 92 117 48 48 51 99 92 117 48 48 51 101 92 117 48 48 50 54 34 44 10 61 45 34 65 103 101 34 58 32 51 48 44 10 61 45 34 73 115 115 116 117 100 101 110 116 34 58 32 102 97 108 115 101 10 61 125]
	//	序列化后的字节切片转字符串：
	//	{
	//	=-"Name": "anko000\u003c\u003e\u0026",
	//	=-"Age": 30,
	//	=-"Isstudent": false
	//	=}

	fmt.Println()
	fmt.Println("-----------------json.Valid()报告数据是否为有效的JSON编码----------------------------")
	//Valid报告数据是否为有效的JSON编码。
	fmt.Println(json.Valid(indent_json))//true

	fmt.Println()
	fmt.Println("-----------------json.RawMessage延迟JSON解码----------------------------")
	// RawMessage是原始编码的JSON值。
	//它实现Marshaler和Unmarshaler，可用于延迟JSON解码或预计算JSON编码。

	type Color struct {
		Space string
		Point json.RawMessage //延迟解析，直到我们知道色彩空间
	}
	type RGB struct {
		R   uint8
		G   uint8
		B   uint8
	}
	type YCbCr struct {
		Y   uint8
		Cb  int8
		Cr  int8
	}
	var j = []byte(`[
		{"Space": "YCbCr", "Point": {"Y": 255, "Cb": 0, "Cr": -10}},
		{"Space": "RGB",   "Point": {"R": 98, "G": 218, "B": 255}}
	]`)
	var colors []Color
	err := json.Unmarshal(j, &colors)
	if err != nil {
		log.Fatalln("error:", err)
	}
	fmt.Printf("第1次反序列化时候的类实例是:%#v\n",colors)

	for i:=0;i<len(colors);i++ {
		var dst interface{}
		switch colors[i].Space {//根据输入的不同反序列化成为不同的类型
		case "RGB":
			dst = new(RGB)
		case "YCbCr":
			dst = new(YCbCr)
		}
		err := json.Unmarshal(colors[i].Point, dst)
		if err != nil {
			log.Fatalln("error:", err)
		}
		fmt.Println(colors[i].Space, dst)
		fmt.Printf("延迟反序列化的类型是：%#v\n",dst)
	}
	//因为每次反序列化都是作用到一个新的实例上面，所以这里并没有因为第2次的反序列化而影响第一次反序列化后的结果
	fmt.Printf("第2次完全反序列化时候的类实例是:%#v",colors)

	//输出：
	//	第1次序列化时候的类实例是:[]main.Color{main.Color{Space:"YCbCr", Point:json.RawMessage{0x7b, 0x22, 0x59, 0x22, 0x3a, 0x20, 0x32, 0x35, 0x35, 0x2c, 0x20, 0x22,
	//			0x43, 0x62, 0x22, 0x3a, 0x20, 0x30, 0x2c, 0x20, 0x22, 0x43, 0x72, 0x22, 0x3a, 0x20, 0x2d, 0x31, 0x30, 0x7d}},
	//			main.Color{Space:"RGB", Point:json.RawMessage{0x7b, 0x22, 0x52, 0x22, 0x3a, 0x20, 0x39, 0x38, 0x2c, 0x20, 0x22, 0x47, 0x22, 0x3a, 0x20, 0x32, 0x31, 0x38,
	//			0x2c, 0x20, 0x22, 0x42, 0x22, 0x3a, 0x20, 0x32, 0x35, 0x35, 0x7d}}}
	//	YCbCr &{255 0 -10}
	//	延迟序列化的类型是：&main.YCbCr{Y:0xff, Cb:0, Cr:-10}
	//	RGB &{98 218 255}
	//	延迟序列化的类型是：&main.RGB{R:0x62, G:0xda, B:0xff}
	//	第2次完全序列化时候的类实例是:[]main.Color{main.Color{Space:"YCbCr", Point:json.RawMessage{0x7b, 0x22, 0x59, 0x22, 0x3a, 0x20, 0x32, 0x35, 0x35, 0x2c, 0x20, 0x22,
	//			0x43, 0x62, 0x22, 0x3a, 0x20, 0x30, 0x2c, 0x20, 0x22, 0x43, 0x72, 0x22, 0x3a, 0x20, 0x2d, 0x31, 0x30, 0x7d}},
	//			main.Color{Space:"RGB", Point:json.RawMessage{0x7b, 0x22, 0x52, 0x22, 0x3a, 0x20, 0x39, 0x38, 0x2c, 0x20, 0x22, 0x47, 0x22, 0x3a, 0x20, 0x32, 0x31, 0x38,
	//			0x2c, 0x20, 0x22, 0x42, 0x22, 0x3a, 0x20, 0x32, 0x35, 0x35, 0x7d}}}


	fmt.Println()
	fmt.Println("-----------------json.RawMessage预序列化类实例----------------------------")
	//这个例子是源码中的例子
	h := json.RawMessage(`{"precomputed": true}`)
	//这样实例化一个类实例都可以，我的天，不过我觉得这个真的没什么特别大的作用，起码目前没发现
	c := struct {
		Header *json.RawMessage `json:"header"`
		Body   string           `json:"body"`
	}{Header: &h, Body: "Hello Gophers!"}

	b, err := json.MarshalIndent(&c, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)

	//输出:
	//	{
	//		"header": {
	//		"precomputed": true
	//	},
	//		"body": "Hello Gophers!"
	//	}


	fmt.Println()
	fmt.Println("-----------------RawMessage对象的序列化与反序列化----------------------------")
	//h := json.RawMessage(`{"precomputed": true}`)//在上面的定义了
	marshalJSON, encode_err000 := h.MarshalJSON()//上个例子中用到对RawMessage对象的序列化的底层就是通过这个方法来实现序列化的
	check_EncDec_err(encode_err000)
	fmt.Printf("RawMessage对象序列化成json对象是：%#v\n",marshalJSON)
	fmt.Printf("他的值是：%v\n",marshalJSON)

	var h1 json.RawMessage

	encode_err_0000 := h1.UnmarshalJSON(marshalJSON)
	check_EncDec_err(encode_err_0000)
	fmt.Printf("json反序列化成RawMessage对象是：%#v\n",h1)
	fmt.Printf("他的值是：%v\n",h1)
	//RawMessage对象其实底层的类型是[]byte,所以他序列化与反序列化的值是一样的，但是类型是不一样的
	//输出：
	//	RawMessage对象序列化成json对象是：[]byte{0x7b, 0x22, 0x70, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x22, 0x3a, 0x20, 0x74, 0x72, 0x75, 0x65, 0x7d}
	//	他的值是：[123 34 112 114 101 99 111 109 112 117 116 101 100 34 58 32 116 114 117 101 125]
	//	json反序列化成RawMessage对象是：json.RawMessage{0x7b, 0x22, 0x70, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x22, 0x3a, 0x20, 0x74, 0x72, 0x75, 0x65, 0x7d}
	//	他的值是：[123 34 112 114 101 99 111 109 112 117 116 101 100 34 58 32 116 114 117 101 125]，你可以看出这个值上面的值是一样的。



	fmt.Println("-------------interface{}类型的字段的序列化与反序列化以及.UseNumber()在解码时候指定解码的类型为json.Number-------------------------")

	type person struct {
		Name interface{}
		Age interface{}
		Isstudent interface{}
	}

	var P_i =person{
		Name:      "ankoiii<>&",
		Age:       -19,
		Isstudent: false,
	}

	marshal_interface, encode_err_interface := json.Marshal(&P_i)
	check_EncDec_err(encode_err_interface)

	fmt.Println("序列化后的json数据字节切片为：",marshal_interface)
	fmt.Println("序列化后的json数据字符串为：",string(marshal_interface))
	//输出：
	//	序列化后的json数据字节切片为： [123 34 78 97 109 101 34 58 34 97 110 107 111 105 105 105 92 117 48 48 51 99 92 117 48 48 51 101 92
	//					117 48 48 50 54 34 44 34 65 103 101 34 58 49 57 44 34 73 115 115 116 117 100 101 110 116 34 58 102 97 108 115 101 125]
	//	序列化后的json数据字符串为： {"Name":"ankoiii\u003c\u003e\u0026","Age":19,"Isstudent":false}

	//下面是对上面json数据的反序列化

	var P_i111 person
	bf_inter:=bytes.NewBuffer(marshal_interface)

	decoder_Rd := json.NewDecoder(bf_inter)
	// UseNumber使解码器将数字解组到接口{}中，作为number类型而不是float64。但是我不知道这样做有什么用以及是否还可以被再次编码（序列化）
	decoder_Rd.UseNumber()
	encode_err_inter111 := decoder_Rd.Decode(&P_i111)
	check_EncDec_err(encode_err_inter111)

	//encode_err_interface111 := json.Unmarshal(marshal_interface, &P_i111)
	//check_EncDec_err(encode_err_interface111)

	fmt.Printf("json数据的反序列化成go对象的实例为：%#v\n",P_i111)
	fmt.Printf("反序列化后的go对象的字段的值和类型是：%v：%T  %v：%T  %v：%T\n",P_i111.Name,
		P_i111.Name,P_i111.Age,P_i111.Age,P_i111.Isstudent,P_i111.Isstudent)
	// TypeOf返回表示i的动态类型的反射类型。
	//如果i是一个nil接口值，则TypeOf返回nil。
	//上面的%T其实底层是采用的是reflect.TypeOf来实现的，也是通过反射来获取他的反射类型，但记住不是获取他没发射之前的类型。这个要分清楚，记住接口其实也是一个类型。但是因为接口没有父类，所以无法断言出来！
	fmt.Printf("反序列化后的go对象的字段的值和类型是：%v：%v  %v：%v  %v：%v\n",P_i111.Name,
		reflect.TypeOf(P_i111.Name),P_i111.Age,reflect.TypeOf(P_i111.Age),P_i111.Isstudent,reflect.TypeOf(P_i111.Isstudent))

	//f_Age := P_i111.Age.(float64)
	f_Age := P_i111.Age.(json.Number)//一旦使用这个断言成json.Number 的话，那么就无法使用go中float64类型的所有特有的方法，比如math.Abs（）以及reflect.TypeOf（）
	//fmt.Println(math.Abs(f_Age),"断言前类型是：",reflect.TypeOf(P_i111.Age),"断言后的新生成的类型是：",reflect.TypeOf(f_Age))//这个Number类的基类是string
	fmt.Println("断言前类型是其实是interface{}类型，因为无法显示这个类型，所以我们这里的其实断言后的类型或者说是他底层原本的类型：",reflect.TypeOf(P_i111.Age),"断言后的新生成的类型是：",reflect.TypeOf(f_Age))//这个Number类的基类是string
	fmt.Println(len(P_i111.Name.(string)))//这里显示的是断言成功的字节数，而并不是断言后的类型或者类型相对应的值
	//.UseNumber()输出：
	//	序列化后的json数据字节切片为： [123 34 78 97 109 101 34 58 34 97 110 107 111 105 105 105 92 117 48 48 51 99 92 117 48 48 51 101 92 117 48 48 50 54 34 44 34 65 103 101 34 58 45 49 57 44 34 73 115 115 116 117 100 101 110 116 34 58 102 97 108 115 101 125]
	//	序列化后的json数据字符串为： {"Name":"ankoiii\u003c\u003e\u0026","Age":-19,"Isstudent":false}
	//	json数据的反序列化成go对象的实例为：main.person{Name:"ankoiii<>&", Age:"-19", Isstudent:false}
	//	反序列化后的go对象的字段的值和类型是：ankoiii<>&：string  -19：json.Number  false：bool
	//	反序列化后的go对象的字段的值和类型是：ankoiii<>&：string  -19：json.Number  false：bool
	//	断言前类型是： json.Number 断言后的新生成的类型是： json.Number
	//	10

	//不使用.UseNumber()的输出需要你自己打印，我实在没法提供，主要是为了偷懒，我懒得重新建立新代码段来测试不使用.UseNumber()情况下的输出，
	// 不过可以告诉你有什么区别：-19：json.Number---》-19：float64

}

func check_EncDec_err(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
