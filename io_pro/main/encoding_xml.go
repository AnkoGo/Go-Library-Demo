package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

func main34674() {
	fmt.Println("-----------------xml编码-----------------------")

	type Address struct {
		City, State string
	}
	type Person struct {
		//xml.Name表示带有名称空间标识符（Space）的XML名称（本地）。
		//在Decoder.Token返回的令牌中，Space标识符以规范的URL形式给出，而不是在要解析的文档中使用的短前缀。
		XMLName   xml.Name `xml:"person"`  //指明xml的根元素，双引号必须写上，不写的话不能被编码，而且该字段必须大写，小写为私有字段，下同
		Id        int      `xml:"id,attr"` //标签中含有“attr”字眼的表示申明一个属性并且作为当前父类的属性而存在
		FirstName string   `xml:"name>first"`
		LastName  string   `xml:"name>last"` // >的左边当前字段的命名父节点，右边是当前字段的命名子节点，
		Age       int      `xml:"age"`
		Height    float32  `xml:"height,omitempty"` //标签中含有“omitempty”属性的字段如果没值(空值为false、0、nil指针、nil接口、长度为0的数组、切片、映射)的话将不会被编码
		Married   bool     //没有标签的字段会按照字段名字来编码xml
		Address            //匿名字段(只写类型而没写字段名字)也能被解析，匿名字段（其标签无效）会被处理为其字段是外层结构体的字段
		Comment   string   `xml:",comment"`
		school    string   `xml:"school"` //私有字段的tag标签是无效的，无论任何时候都将不会被解析
		//Father string 	`xml:"father,chardata"`//具有标签",chardata"的字段会作为字符数据写入，而非XML元素，会报错
		//Mother string	`xml:"mother,innerxml"`//具有标签",innerxml"的字段会原样写入，而不会经过正常的序列化过程,会报错
	}
	v := &Person{Id: 13, FirstName: "John", LastName: "Doe", Age: 42, school: "anko_school"} //实例化一个类
	v.Comment = " Need more details. "                                                       //实例化之后还可以指明字段的值
	v.Address = Address{"Hanga Roa", "Easter Island"}                                        //再实例化一个其他的类的实例作为这个类的字段

	buffer := new(bytes.Buffer)   //编码后的xml装载容器
	enc := xml.NewEncoder(buffer) //实例化一个编码器
	enc.Indent("  ", "    ")      //指明编码器的格式---前缀以及缩进采用什么字符

	if err := enc.Encode(v); err != nil {//进行编码
		fmt.Printf("error: %v\n", err)
	}


	//下面的一堆是我对底层的一些api的试用，不过还没知道怎么使用

	// StartElement代表XML开始元素。
	//var startEle = xml.StartElement{
	//	//xml.Name表示带有名称空间标识符（Space）的XML名称（本地）。
	//	//在Decoder.Token返回的令牌中，Space标识符以规范的URL形式给出，而不是在要解析的文档中使用的短前缀。
	//	//Name: xml.Name{"space111","person111"},//这个会解码失败，
	//	Name: xml.Name{"space111", "person"}, //这个不会解码失败，我们可以随意指定Space，但是不能随意指定Local，
	//	// 这个值必须与Person结构体XMLName字段标签xml:"person"中相同，当前如果Person结构体中没有类型为xml.Name的XMLName字段的话，那么我们这里
	//	//可以随意指定值
	//	Attr: nil, //Attr表示XML元素（名称=值）中的属性。
	//}
	//startEle.Copy()////复制创建StartElement的新副本。
	//startEle.End()// End返回相应的XML end元素。

	// EncodeElement使用start作为编码中最外面的标记，将v的XML编码写入流中。
	//有关将Go值转换为XML的详细信息，请参见Marshal的文档。
	// EncodeElement在返回之前会调用Flush。.Encode()也一样，跟.encode()的底层几乎是一样的实现
	//enc.EncodeElement(v,startEle)

	// EncodeToken将给定的XML令牌写入流。
	//如果StartElement和EndElement标记未正确匹配，则返回错误。
	//
	// EncodeToken不会调用Flush，因为它通常是更大的操作的一部分，例如Encode或EncodeElement（或在此过程中调用的自定义Marshaler的MarshalXML），并且这些操作将在完成时调用Flush。
	//创建一个Encoder然后直接调用EncodeToken而不使用Encode或EncodeElement的调用者，在完成时需要调用Flush以确保将XML写入基础编写器。
	//
	// EncodeToken仅允许将Target设置为“ xml”的ProcInst写入流中，作为第一个令牌。目前还不知道token怎么用，先搁置

	// StartElement代表XML开始元素。
	//_ = enc.EncodeToken(xml.StartElement{
	//	Name: xml.Name{"", "person"},
	//	Attr: []xml.Attr{xml.Attr{
	//		Name:  xml.Name{"", "name111"},
	//		Value: "anko",
	//	}, xml.Attr{
	//		Name:  xml.Name{"", "id"},
	//		Value: "16",
	//	}},
	//}) //可以采用上面定义的StartElement类实例
	//
	//_ = enc.EncodeToken(xml.StartElement{
	//	Name: xml.Name{"", "name"},
	//	Attr:nil,
	//})
	//
	//_ = enc.EncodeToken(xml.StartElement{
	//	Name: xml.Name{"", "first"},
	//	Attr:nil,
	//})
	//// CharData表示XML字符数据（原始文本），其中XML转义序列已替换为它们所表示的字符。
	//enc.EncodeToken(xml.CharData("John"))
	//
	//
	//_ = enc.EncodeToken(xml.StartElement{
	//	Name: xml.Name{"", "/first"},
	//	Attr:nil,
	//})
	//
	//
	//
	//
	//_ = enc.EncodeToken(xml.StartElement{
	//	Name: xml.Name{"", "/name"},
	//	Attr:nil,
	//})
	//
	//_ = enc.EncodeToken(xml.StartElement{
	//	Name: xml.Name{"", "age"},
	//	Attr:nil,
	//})
	//// CharData表示XML字符数据（原始文本），其中XML转义序列已替换为它们所表示的字符。
	//enc.EncodeToken(xml.CharData("42"))
	//
	//
	//_ = enc.EncodeToken(xml.StartElement{
	//	Name: xml.Name{"", "/age"},
	//	Attr:nil,
	//})
	//
	//
	//
	////指令表示形式为<！text>的XML指令。
	////字节不包含<！ 和>标记。
	////enc.EncodeToken(xml.Directive{'a', 'b'})
	//// ProcInst表示<？target inst？>形式的XML处理指令
	//
	////注释表示<！-comment->形式的XML注释。
	////字节不包含<！-和->注释标记。
	//enc.EncodeToken(xml.Comment("Need more details."))
	//
	////enc.EncodeToken(xml.ProcInst{"name", nil})
	//
	////enc.EncodeToken(xml.ProcInst{"Married", []byte{byte(false)}})
	//_ = enc.EncodeToken(xml.StartElement{
	//	Name: xml.Name{"", "/person"},
	//	Attr: nil,
	//})



	//刷新将所有缓冲的XML刷新到基础编写器。
	//有关何时需要的详细信息，请参见EncodeToken文档。
	enc.Flush() //这个方法对于.encode()方法和.EncodeElement（）类型的编码是不必的和多余的，因为encode()本身就会调用这个flush()方法
	//flush方法的真正用法是针对.EncodeToken()方法的，因为这个方法编码结束后不会调用flush(),而其他的2个编码方法是会的！

	fmt.Println("编码后的xml数据是：\n", buffer)



	fmt.Println("-----------------xml解码-----------------------")

	var P Person
	//bytes.NewBuffer()
	decoder_xml := xml.NewDecoder(buffer)

	dec_err := decoder_xml.Decode(&P)
	check_EncDec_err(dec_err)

	fmt.Printf("解码后的go类实例是：\n%#v\n", P)
	//.encode()输出：
	//-----------------xml编码-----------------------
	//编码后的xml数据是：
	//	<person id="13">
	//		<name>
	//			<first>John</first>
	//			<last>Doe</last>
	//		</name>
	//		<age>42</age>
	//		<Married>false</Married>
	//		<City>Hanga Roa</City>
	//		<State>Easter Island</State>
	//		<!-- Need more details. -->
	//	</person>
	//-----------------xml解码-----------------------
	//解码后的go类实例是：
	//	main.Person{XMLName:xml.Name{Space:"", Local:"person"}, Id:13, FirstName:"John", LastName:"Doe",
	//		Age:42, Height:0, Married:false, Address:main.Address{City:"Hanga Roa", State:"Easter Island"},
	//		Comment:" Need more details. ", school:""}
	//Space:""表示当前的目录下，xml.Name指明根元素

	//.EncodeElement()输出：
	//-----------------xml编码-----------------------
	//编码后的xml数据是：
	//	<person111 xmlns="space111" id="13">
	//		<name>
	//			<first>John</first>
	//			<last>Doe</last>
	//		</name>
	//		<age>42</age>
	//		<Married>false</Married>
	//		<City>Hanga Roa</City>
	//		<State>Easter Island</State>
	//		<!-- Need more details. -->
	//	</person111>
	//-----------------xml解码-----------------------
	//expected element type <person> but have <person111>
	//解码后的go类实例是：
	//main.Person{XMLName:xml.Name{Space:"", Local:""}, Id:0, FirstName:"", LastName:"", Age:0, Height:0, Married:false,
	//			Address:main.Address{City:"", State:""}, Comment:"", school:""}
	//从上面可以看得出，虽然我们可以指定一些属性，但是如果解码出错的话将会影响后面字段的解码


	fmt.Println("-----------------使用函数.MarshalIndent()和.Marshal()和进行xml编码-----------------------")
	//indent_byte, e1 := xml.MarshalIndent(v, "  ", "    ")
	indent_byte, e1 := xml.Marshal(v)
	check_EncDec_err(e1)
	fmt.Println("编码后的字节数据是：",indent_byte)
	fmt.Println("编码后的字符串数据是：\n",string(indent_byte))
	//输出：
	//	-----------------使用函数.MarshalIndent()和.Marshal()进行xml编码-----------------------
	//	编码后的字节数据是： [32 32 60 112 101 114 115 111 110 32 105 100 61 34 49 51 34 62 10 32 32 32 32 32 32 60 110 97 109 101 62
	//						10 32 32 32 32 32 32 32 32 32 32 60 102 105 114 115 116 62 74 111 104 110 60 47 102 105 114 115 116 62 10 32 32 32 32 32 32 32
	//						32 32 32 60 108 97 115 116 62 68 111 101 60 47 108 97 115 116 62 10 32 32 32 32 32 32 60 47 110 97 109 101 62 10 32 32 32 32 32
	//						32 60 97 103 101 62 52 50 60 47 97 103 101 62 10 32 32 32 32 32 32 60 77 97 114 114 105 101 100 62 102 97 108 115 101 60 47 77
	//						97 114 114 105 101 100 62 10 32 32 32 32 32 32 60 67 105 116 121 62 72 97 110 103 97 32 82 111 97 60 47 67 105 116 121 62 10 32
	//						32 32 32 32 32 60 83 116 97 116 101 62 69 97 115 116 101 114 32 73 115 108 97 110 100 60 47 83 116 97 116 101 62 10 32 32 32 32
	//						32 32 60 33 45 45 32 78 101 101 100 32 109 111 114 101 32 100 101 116 97 105 108 115 46 32 45 45 62 10 32 32 60 47 112 101 114
	//						115 111 110 62]
	//	编码后的字符串数据是：
	//		<person id="13">
	//			<name>
	//				<first>John</first>
	//				<last>Doe</last>
	//			</name>
	//			<age>42</age>
	//			<Married>false</Married>
	//			<City>Hanga Roa</City>
	//			<State>Easter Island</State>
	//			<!-- Need more details. -->
	//		</person>

	//输出：
	//	-----------------使用函数.MarshalIndent()和.Marshal()进行xml编码-----------------------
	//	编码后的字节数据是： [60 112 101 114 115 111 110 32 105 100 61 34 49 51 34 62 60 110 97 109 101 62 60 102 105 114 115 116 62
	//						74 111 104 110 60 47 102 105 114 115 116 62 60 108 97 115 116 62 68 111 101 60 47 108 97 115 116 62 60 47 110 97 109 101 62 60
	//						97 103 101 62 52 50 60 47 97 103 101 62 60 77 97 114 114 105 101 100 62 102 97 108 115 101 60 47 77 97 114 114 105 101 100 62
	//						60 67 105 116 121 62 72 97 110 103 97 32 82 111 97 60 47 67 105 116 121 62 60 83 116 97 116 101 62 69 97 115 116 101 114 32 73
	//						115 108 97 110 100 60 47 83 116 97 116 101 62 60 33 45 45 32 78 101 101 100 32 109 111 114 101 32 100 101 116 97 105 108 115 46
	//						32 45 45 62 60 47 112 101 114 115 111 110 62]
	//	编码后的字符串数据是：
	//	<person id="13"><name><first>John</first><last>Doe</last></name><age>42</age><Married>false</Married><City>Hanga Roa</City><State>Easter Island</State><!-- Need more details. --></person>




	fmt.Println("-----------------使用函数.MarshalIndent()和.Marshal()进行xml解码-----------------------")

	var Per Person
	Unm_err := xml.Unmarshal(indent_byte, &Per)

	check_EncDec_err(Unm_err)

	fmt.Printf("解码后的go结构体的实例是：\n%#v\n",Per)
	//输出：
	//	-----------------使用函数.MarshalIndent()和.Marshal()进行xml解码-----------------------
	//	解码后的go结构体的实例是：
	//	main.Person{XMLName:xml.Name{Space:"", Local:"person"}, Id:13, FirstName:"John", LastName:"Doe", Age:42, Height:0, Married:false, Address:main.Address{City:"Hanga Roa", State:"Easter Island"}, Comment:" Need more details. ", school:""}

	fmt.Println("-----------------使用函数.EscapeText()和.Escape()进行xml解码-----------------------")

	fmt.Println("在转义之前的xml字符串是：",string(indent_byte))
	bf:=new(bytes.Buffer)

	// Escape类似于EscapeText，但省略了错误返回值。
	//提供它是为了与Go 1.0向后兼容。
	//定位到Go 1.1或更高版本的代码应使用EscapeText。底层是调用了EscapeText（）方法，这个方法不应该再被使用
	//xml.Escape()

	// EscapeText向w写入等效于纯文本数据的XML。
	// 底层是用到了escapeText()函数，关于这个函数的说明如下
	// escapeText向w写入与纯文本数据s相同的正确转义的XML。 如果escapeNewline为true，则换行符将被转义。但是绝对是True,因为这个方法在被EscapeText（）
	//函数调用的时候设置了一定为true,所以这个输出的结果一定是连换行符都被转义了的纯文本
	EscapeText_err := xml.EscapeText(bf, indent_byte)
	check_EncDec_err(EscapeText_err)
	fmt.Println("在转义之后的纯文本字符串是：",bf)
	//输出：
	//	在转义之前的xml字符串是： <person id="13"><name><first>John</first><last>Doe</last></name><age>42</age><Married>false</Married><City>Hanga Roa</City><State>Easter Island</State><!-- Need more details. --></person>
	//	在转义之后的纯文本字符串是： &lt;person id=&#34;13&#34;&gt;&lt;name&gt;&lt;first&gt;John&lt;/first&gt;&lt;last&gt;Doe&lt;/last&gt;&lt;/name&gt;&lt;age&gt;42&lt;/age&gt;&lt;Married&gt;false&lt;/Married&gt;&lt;City&gt;Hanga Roa&lt;/City&gt;&lt;State&gt;Easter Island&lt;/State&gt;&lt;!-- Need more details. --&gt;&lt;/person&gt;




}











































