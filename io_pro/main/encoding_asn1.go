package main

import (
	"encoding/asn1"
	"fmt"
	"os"
	"sort"
	"time"
)

type myPerson struct {
	Name    string
	Age     int
	Country string
	Isfat   bool//这里所有的字段首字母都必须大写
}

func main() {
	fmt.Println("-----------Marshal函数---------------")
	//Marshal函数返回val的ASN.1编码。
	//除了Unmarshal可以识别的struct标记外,此外还提供了供Unmarshal函数识别的结构体标签，可用如下标签：
	//	ia5:           使字符串序列化为ASN.1 IA5String类型
	//	omitempty:     使空切片被跳过
	//	printable:     使字符串序列化为ASN.1 PrintableString类型
	//	utf8:          使字符串序列化为ASN.1 UTF8字符串
	//	utc:     	   使time.Time序列化为ASN.1 UTCTime类型值
	//	generalized:   使time.Time序列化为ASN.1 GeneralizedTime类型值
	fmt.Println("你好啊")
	var P = myPerson{
		Name:    "anko",
		Age:     22,
		Country: "China",
		Isfat:   false,
	}
	bytes, e := asn1.Marshal(P)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(bytes) //[48 19 19 4 97 110 107 111 2 1 22 19 5 67 104 105 110 97 1 1 0]
	fmt.Println(string(bytes))

	//asn1.MarshalWithParams()
	fmt.Println("-----------UnMarshal函数---------------")

	src_ls:=[]byte{48 ,19, 19, 4 ,97, 110, 107, 111, 2 ,1, 22, 19, 5, 67, 104, 105, 110, 97, 1, 1, 0}

	var P1 myPerson

	rest, e1 := asn1.Unmarshal(src_ls, &P1)//这里必须传递的是指针&P1,否则报错
	if e1 != nil{
		fmt.Println(e1)
	}

	fmt.Println(rest)
	fmt.Println(string(rest))
	fmt.Println(P1)
	//输出：
	//	[]
	//（空字符串）
	//	{anko 22 China false}
	//这里我们测试了struct这个复杂类型，但是其实像int和string类型都是可以的！
	fmt.Println("------------------string类型的序列化与反序列化-----------------")
	s := "hello"
	//s := "hello\u00bc"//非ASCII字符会报错，必须只能是ascii字符
	mdata, _ := asn1.Marshal(s)
	fmt.Println(mdata,"----",string(mdata))
	var newstr string
	asn1.Unmarshal(mdata, &newstr)
	fmt.Println(newstr)
	//输出：
	//	[19 5 104 101 108 108 111] ---- hello
	//	hello

	fmt.Println("------------------int类型的序列化与反序列化-----------------")
	s1 := 16
	mdata1, _ := asn1.Marshal(s1)
	fmt.Println(mdata1,"----",string(mdata1))
	var newstr1 int
	asn1.Unmarshal(mdata1, &newstr1)
	fmt.Println(newstr)
	//输出：
	//	[2 1 16] ---- 
	//	hello

	fmt.Println("------------------time类型的序列化与反序列化-----------------")

	t := time.Now()
	mdata, err44 := asn1.Marshal(t)
	checkError(err44)
	fmt.Println(t)

	var newtime = new(time.Time)
	_, err1234 := asn1.Unmarshal(mdata,newtime)//因为这里newtime已经是指针了，所以不要再加上&来取指针了，不然会报错
	checkError(err1234)
	fmt.Println(newtime)
	//输出：
	//	2019-10-10 10:57:05.6999874 +0800 CST m=+0.003908801
	//	2019-10-10 10:57:05 +0800 CST


	fmt.Println("-----------------asn1包的类型系统------------------------")
	// BitString是需要ASN.1 BIT STRING类型时使用的结构。 将位字符串填充到内存中最近的字节，并记录有效位数。 填充位将为零。
	var BS = asn1.BitString{
		Bytes:     []byte{'a', 'b', 'c', 'd'},
		BitLength: 4,
	}
	// RightAlign返回一个切片，其中填充位在开头。 切片可以与BitString共享内存。
	// At返回给定索引处的位。 如果索引超出范围，则返回false。
	fmt.Println(BS.Bytes)        //[97 98 99 100]
	fmt.Println(BS.BitLength)    //4
	fmt.Println(BS.At(2))        //2
	fmt.Println(BS.RightAlign()) //[6 22 38 54]
	fmt.Println(BS)              //{[97 98 99 100] 4}

	fmt.Println("----------------------")
	var i asn1.Enumerated = 1
	fmt.Println(i) //1

	var b asn1.Flag = true
	fmt.Println(b) //true

	var O1 asn1.ObjectIdentifier = []int{1, 2, 3, 4}

	var O2 asn1.ObjectIdentifier = []int{1, 2, 3, 4}
	var O3 asn1.ObjectIdentifier = []int{1, 2, 3}
	var O4 asn1.ObjectIdentifier = []int{1, 2, 3, 0}

	fmt.Println(O1)           //1.2.3.4
	fmt.Println(O1.String())  //1.2.3.4
	fmt.Println(O1.Equal(O2)) //true
	fmt.Println(O1.Equal(O3)) //false
	fmt.Println(len(O1))      //4
	sort.Ints(O4)
	fmt.Println(O4) //0.1.2.3
	O4 = append(O4, 5)
	fmt.Println(O4) //0.1.2.3.5
	//属性基本跟[]int类型完全一样

	fmt.Println("----------------------")

	// NullBytes包含表示DER编码的ASN.1 NULL类型的字节。
	NBytes := asn1.NullBytes
	fmt.Println(NBytes) //[5 0]
	fmt.Println(NBytes) //[5 0]

	// RawContent用于表示需要为结构保留未解码的DER数据。 要使用它，
	// 结构的第一个字段必须具有此类型。 任何其他字段都具有此类型是错误的。
	var RC = asn1.RawContent{'a', 'b', 'c'}
	fmt.Println(RC)      //[97 98 99]
	fmt.Println(len(RC)) //3

	// RawValue表示未解码的ASN.1对象。
	var RV = asn1.RawValue{
		Class:      1,
		Tag:        2,
		IsCompound: false,
		Bytes:      []byte{'a', 'b'},
		FullBytes:  []byte{'c', 'd'},
	}
	fmt.Println(RV.Class)
	fmt.Println(RV.Tag)
	fmt.Println(RV.IsCompound)
	fmt.Println(RV.Bytes)
	fmt.Println(RV.FullBytes)
	//输出：
	//	1
	//	2
	//	false
	//	[97 98]
	//	[99 100]

	fmt.Println(asn1.TagUTF8String)        //12,其实这个数字对应一个类型，被用于上面的asn1.RawValue.Tag
	fmt.Println(asn1.ClassContextSpecific) //2,其实这个数字对应一个类型，被用于上面的asn1.RawValue.Class

	//对asn1还不大熟悉，暂且先讲这么多
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
