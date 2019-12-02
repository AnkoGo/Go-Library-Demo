package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strings"
)

func main() {
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
	//enc.Indent("=", "-")//即使这样指明之后还是可以正常无误的进行解码，这个跟json就不一样了，json如果这样指明之后是无法解码的！

	if err := enc.Encode(v); err != nil {//进行编码
		fmt.Printf("error: %v\n", err)
	}

	//my_token(enc)


	// Println格式使用其操作数的默认格式并写入标准输出。
	//始终在操作数之间添加空格，并添加换行符。
	//返回写入的字节数以及遇到的任何写入错误。
	//底层调用的是Fprintln,关于这个方法的说明如下：
		// Fprintln格式使用其操作数的默认格式并写入w。
		//始终在操作数之间添加空格，并添加换行符。
		//返回写入的字节数以及遇到的任何写入错误。
	//fmt.Println("编码后的xml数据是\n",buffer)
	fmt.Printf("编码后的xml数据是：\n%v\n", buffer)//为了避免空格的产生，我们最好通过Printf来输出而不是Println


	fmt.Println("-----------------xml解码-----------------------")

	var P Person
	xml_str:=`
  <person id="13">
      <name>
          <first>John</first>
          <last>Doe</Llast>
      </name>
      <age>42</age>
      <Married>false</Married>
      <City>Hanga Roa</City>
      <State>Easter Island</State>
      <!-- Need more details. -->
  </person>
			`

	reader := strings.NewReader(xml_str)
	decoder_xml := xml.NewDecoder(reader)//这个是自定义字符串来解码
	//decoder_xml := xml.NewDecoder(buffer)//这个是采用上面编码的缓存来进行解码

	// Strict默认为true，以强制执行XML规范的要求。
	//如果设置为false，则解析器允许包含常见错误的输入：
	// *如果元素缺少结束标记，则解析器会根据需要发明结束标记，以使Token的返回值保持适当的平衡。
	// *在属性值和字符数据中，未知或格式不正确的字符实体（以＆开头的序列）将保留下来。
	// 设置：
	//
	// d.Strict = false
	// d.AutoClose = xml.HTMLAutoClose
	// d.Entity = xml.HTMLEntity
	//
	//创建一个可以处理典型HTML的解析器。
	//
	//严格模式不会强制执行XML命名空间TR的要求。
	//特别是它不会拒绝使用未定义前缀的名称空间标签。
	//此类标签以未知前缀记录为名称空间URL。
	//decoder_xml.Strict=true//默认就是true，这样的话遇到错误的标签元素就无法就无法解码当前以及往后的元素了
	decoder_xml.Strict=false//为false的话，遇到错误的标签元素的话当前元素也一样会被解析，但是往后的元素一样不会被解析，比如上面的“<last>Doe</Llast>”就是错误的，
	// 但是如果是第一个标签错误的话是无法解析的，比如<Llast>Doe</last>，这样的话是不会报错的，但是当前元素和往后的元素都将不会被解码，
	//再者，假设<last>Doe</Llast>把后面的标签去掉的情况下解码会输出：
	//main.Person{XMLName:xml.Name{Space:"", Local:"person"}, Id:13, FirstName:"John", LastName:"Doe\n      ", Age:42, Height:0, Married:false, Address:main.Address{City:"Hanga Roa", State:"Easter Island"}, Comment:" Need more details. ", school:""}

	//AutoClose说明：
	//当Strict == false时，AutoClose指示一组元素在打开后立即视为已关闭，无论是否存在结束元素。
	//HTMLAutoClose说明：
	// HTMLAutoClose是应视为自动关闭的HTML元素集。
	//请参阅Decoder.Strict和Decoder.Entity字段的文档。
	decoder_xml.AutoClose = xml.HTMLAutoClose

	//Entity说明：
	//Entity可用于将非标准实体名称映射为字符串替换。
	//无论实际的映射内容如何，解析器的行为就像这些标准映射出现在映射中一样：
	//
	//	"lt": "<",
	//	"gt": ">",
	//	"amp": "&",
	//	"apos": "'",
	//	"quot": `"`,
	//HTMLEntity说明：
	// HTMLEntity是一个实体映射，其中包含标准HTML实体字符的翻译。
	//请参阅Decoder.Strict和Decoder.Entity字段的文档。
	decoder_xml.Entity = xml.HTMLEntity//不清楚作用
	dec_err := decoder_xml.Decode(&P)//不清楚作用
	check_EncDec_err(dec_err)


	//decoder_xml.DecodeElement()
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

	//指明其他的缩进以及前缀的输出：
	//	-----------------xml编码-----------------------
	//	编码后的xml数据是：
	//	=<person id="13">
	//	=-<name>
	//	=--<first>John</first>
	//	=--<last>Doe</last>
	//	=-</name>
	//	=-<age>42</age>
	//	=-<Married>false</Married>
	//	=-<City>Hanga Roa</City>
	//	=-<State>Easter Island</State>
	//	=-<!-- Need more details. -->
	//	=</person>-----------------xml解码-----------------------
	//	解码后的go类实例是：
	//	main.Person{XMLName:xml.Name{Space:"", Local:"person"}, Id:13, FirstName:"John", LastName:"Doe", Age:42, Height:0, Married:false, Address:main.Address{City:"Hanga Roa", State:"Easter Island"}, Comment:" Need more details. ", school:""}

	//decoder_xml.Strict=true时候解码xml_str的输出：
	//	-----------------xml编码-----------------------
	//	编码后的xml数据是：
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
	//	-----------------xml解码-----------------------
	//	XML syntax error on line 5: element <last> closed by </Llast>
	//	解码后的go类实例是：
	//	main.Person{XMLName:xml.Name{Space:"", Local:"person"}, Id:13, FirstName:"John", LastName:"", Age:0, Height:0, Married:false, Address:main.Address{City:"", State:""}, Comment:"", school:""}

}




func my_token(enc *(xml.Encoder))  {
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

	//StartElement代表XML开始元素。
	_ = enc.EncodeToken(xml.StartElement{
		Name: xml.Name{"", "person"},
		Attr: []xml.Attr{xml.Attr{
			Name:  xml.Name{"", "name111"},
			Value: "anko",
		}, xml.Attr{
			Name:  xml.Name{"", "id"},
			Value: "16",
		}},
	}) //可以采用上面定义的StartElement类实例

	_ = enc.EncodeToken(xml.StartElement{
		Name: xml.Name{"", "name"},
		Attr:nil,
	})

	_ = enc.EncodeToken(xml.StartElement{
		Name: xml.Name{"", "first"},
		Attr:nil,
	})
	// CharData表示XML字符数据（原始文本），其中XML转义序列已替换为它们所表示的字符。
	enc.EncodeToken(xml.CharData("John"))

	enc.Indent("", "")
	_ = enc.EncodeToken(xml.StartElement{
		Name: xml.Name{"", "/first"},
		Attr:nil,
	})
	enc.Indent("\t", "  ")



	_ = enc.EncodeToken(xml.StartElement{
		Name: xml.Name{"", "/name"},
		Attr:nil,
	})

	_ = enc.EncodeToken(xml.StartElement{
		Name: xml.Name{"", "age"},
		Attr:nil,
	})
	// CharData表示XML字符数据（原始文本），其中XML转义序列已替换为它们所表示的字符。
	enc.EncodeToken(xml.CharData("42"))


	_ = enc.EncodeToken(xml.StartElement{
		Name: xml.Name{"", "/age"},
		Attr:nil,
	})



	//指令表示形式为<！text>的XML指令。
	//字节不包含<！ 和>标记。
	//enc.EncodeToken(xml.Directive{'a', 'b'})
	// ProcInst表示<？target inst？>形式的XML处理指令

	//注释表示<！-comment->形式的XML注释。
	//字节不包含<！-和->注释标记。
	enc.EncodeToken(xml.Comment("Need more details."))

	//enc.EncodeToken(xml.ProcInst{"name", nil})

	//enc.EncodeToken(xml.ProcInst{"Married", []byte{byte(false)}})
	_ = enc.EncodeToken(xml.EndElement{Name:xml.Name{"person", "person"},})


	//刷新将所有缓冲的XML刷新到基础编写器。
	//有关何时需要的详细信息，请参见EncodeToken文档。
	enc.Flush() //这个方法对于.encode()方法和.EncodeElement（）类型的编码是不必的和多余的，因为encode()本身就会调用这个flush()方法
	//flush方法的真正用法是针对.EncodeToken()方法的，因为这个方法编码结束后不会调用flush(),而其他的2个编码方法是会的！
}