package main

import (
	"bytes"
	base642 "encoding/base64"
	"fmt"
	"io"
	"time"

	"os"
	"reflect"
	"strconv"
)

func main() {

	fmt.Println("-------------reflect.Type对象-----------------")

	//type Type 对象说明：
	//type Type interface {
	//	// Kind返回该接口的具体分类
	//	Kind() Kind
	//	// Name返回该类型在自身包内的类型名，如果是未命名类型会返回""
	//	Name() string
	//	// PkgPath返回类型的包路径，即明确指定包的import路径，如"encoding/base64"
	//	// 如果类型为内建类型(string, error)或未命名类型(*T, struct{}, []int)，会返回""
	//	PkgPath() string
	//	// 返回类型的字符串表示。该字符串可能会使用短包名（如用base64代替"encoding/base64"）
	//	// 也不保证每个类型的字符串表示不同。如果要比较两个类型是否相等，请直接用Type类型比较。
	//	String() string
	//	// 返回要保存一个该类型的值需要多少字节；类似unsafe.Sizeof
	//	Size() uintptr
	//	// 返回当从内存中申请一个该类型值时，会对齐的字节数
	//	Align() int
	//	// 返回当该类型作为结构体的字段时，会对齐的字节数
	//	FieldAlign() int
	//	// 如果该类型实现了u代表的接口，会返回真
	//	Implements(u Type) bool
	//	// 如果该类型的值可以直接赋值给u代表的类型，返回真
	//	AssignableTo(u Type) bool
	//	// 如该类型的值可以转换为u代表的类型，返回真
	//	ConvertibleTo(u Type) bool
	//	// 返回该类型的字位数。如果该类型的Kind不是Int、Uint、Float或Complex，会panic
	//	Bits() int
	//	// 返回array类型的长度，如非数组类型将panic
	//	Len() int
	//	// 返回该类型的元素类型，如果该类型的Kind不是Array、Chan、Map、Ptr或Slice，会panic
	//	Elem() Type
	//	// 返回map类型的键的类型。如非映射类型将panic
	//	Key() Type
	//	// 返回一个channel类型的方向，如非通道类型将会panic
	//	ChanDir() ChanDir
	//	// 返回struct类型的字段数（匿名字段算作一个字段），如非结构体类型将panic
	//	NumField() int
	//	// 返回struct类型的第i个字段的类型，如非结构体或者i不在[0, NumField())内将会panic
	//	Field(i int) StructField
	//	// 返回索引序列指定的嵌套字段的类型，
	//	// 等价于用索引中每个值链式调用本方法，如非结构体将会panic
	//	FieldByIndex(index []int) StructField
	//	// 返回该类型名为name的字段（会查找匿名字段及其子字段），
	//	// 布尔值说明是否找到，如非结构体将panic
	//	FieldByName(name string) (StructField, bool)
	//	// 返回该类型第一个字段名满足函数match的字段，布尔值说明是否找到，如非结构体将会panic
	//	FieldByNameFunc(match func(string) bool) (StructField, bool)
	//	// 如果函数类型的最后一个输入参数是"..."形式的参数，IsVariadic返回真
	//	// 如果这样，t.In(t.NumIn() - 1)返回参数的隐式的实际类型（声明类型的切片）
	//	// 如非函数类型将panic
	//	IsVariadic() bool
	//	// 返回func类型的参数个数，如果不是函数，将会panic
	//	NumIn() int
	//	// 返回func类型的第i个参数的类型，如非函数或者i不在[0, NumIn())内将会panic
	//	In(i int) Type
	//	// 返回func类型的返回值个数，如果不是函数，将会panic
	//	NumOut() int
	//	// 返回func类型的第i个返回值的类型，如非函数或者i不在[0, NumOut())内将会panic
	//	Out(i int) Type
	//	// 返回该类型的方法集中方法的数目
	//	// 匿名字段的方法会被计算；主体类型的方法会屏蔽匿名字段的同名方法；
	//	// 匿名字段导致的歧义方法会滤除
	//	NumMethod() int
	//	// 返回该类型方法集中的第i个方法，i不在[0, NumMethod())范围内时，将导致panic
	//	// 对非接口类型T或*T，返回值的Type字段和Func字段描述方法的未绑定函数状态
	//	// 对接口类型，返回值的Type字段描述方法的签名，Func字段为nil
	//	Method(int) Method
	//	// 根据方法名返回该类型方法集中的方法，使用一个布尔值说明是否发现该方法
	//	// 对非接口类型T或*T，返回值的Type字段和Func字段描述方法的未绑定函数状态
	//	// 对接口类型，返回值的Type字段描述方法的签名，Func字段为nil
	//	MethodByName(string) (Method, bool)
	//	// 内含隐藏或非导出方法
	//}Type类型用来表示一个go类型。
	//
	//不是所有go类型的Type值都能使用所有方法。请参见每个方法的文档获取使用限制。在调用有分类限定的方法时，
	// 应先使用Kind方法获知类型的分类。调用该分类不支持的方法会导致运行时的panic。

	str := "hello world!"
	var str_I interface{} = "interface string"
	//TypeOf返回接口中保存的值的类型，TypeOf(nil)会返回nil。
	T := reflect.TypeOf(str)
	Ti := reflect.TypeOf(str_I)
	fmt.Println("str对象的底层类型为：", T)
	fmt.Println("str对象的底层类型为：", Ti)
	// Name返回该类型在自身包内的类型名，如果是未命名类型会返回""
	fmt.Println(T.Name())
	fmt.Println(Ti.Name())
	//输出：
	//str对象的底层类型为： string
	//str对象的底层类型为： string
	//string
	//string

	// String（）返回类型的字符串表示。该字符串可能会使用短包名（如用base64代替"encoding/base64"）
	// 也不保证每个类型的字符串表示不同。如果要比较两个类型是否相等，请直接用Type类型比较。
	fmt.Println(T.String())
	var buf = new(bytes.Buffer)
	T1 := reflect.TypeOf(buf)
	fmt.Println(T1.String())

	var base64 = new(base642.Encoding)
	T2 := reflect.TypeOf(base64)
	fmt.Println(T2.String())
	//输出：
	//string
	//*bytes.Buffer
	//*base64.Encoding

	//// Elem returns a type's element type.（elem返回一个“包含类型”（可以容纳且仅容纳一种其他类型的子元素的类型）包含的的子元素类型）
	//// It panics if the type's Kind is not Array, Chan, Map, Ptr, or Slice.（如果底层类型不是Array, Chan, Map, Ptr, or Slice的话就会抛出异常）
	//fmt.Println(T.Elem())
	////报错，字符串不是“包含类型”
	////panic: reflect: Elem of invalid type
	////
	////goroutine 1 [running]:
	////reflect.(*rtype).Elem(0x4a9f20, 0xc000006018, 0xc00007bec0)
	////	C:/Go/src/reflect/type.go:920 +0x14a
	////main.main()
	////	C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:130 +0x58d

	//var struct1= struct {
	//	string
	//	int
	//	bool
	//}{"anko",15,true}
	//struc_type:=reflect.TypeOf(struct1)
	//fmt.Println(struc_type.Elem())
	////报错：
	////panic: reflect: Elem of invalid type
	////
	////goroutine 1 [running]:
	////reflect.(*rtype).Elem(0x4bbca0, 0xc00007bef0, 0x4bbca0)
	////	C:/Go/src/reflect/type.go:920 +0x14a
	////main.main()
	////	C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:146 +0x601

	m := map[string]int{"anko": 14, "jerry": 34, "gogo": 20}
	T_map := reflect.TypeOf(m)
	fmt.Println(T_map.Elem())
	//输出：
	//int

	c := make(chan string, 3)
	c <- "这是chan中的第1个元素"
	c <- "这是chan中的第2个元素"
	c <- "这是chan中的第3个元素"
	T_chan := reflect.TypeOf(c)
	fmt.Println(T_chan.Elem())
	//输出：
	//string

	array := [3]bool{true, false, true}
	T_array := reflect.TypeOf(array)
	fmt.Println(T_array.Elem())
	//输出：
	//bool

	array1 := [3][]bool{[]bool{true, false}, []bool{false, true, true}}
	T_array1 := reflect.TypeOf(array1)
	fmt.Println(T_array1.Elem())
	//输出：
	//[]bool
	//由上可知，返回的类型只会返回第一层级的类型

	ls := []byte{97, 98, 99, 100}
	T_ls := reflect.TypeOf(ls)
	fmt.Println(T_ls.Elem())
	//输出：
	//uint8

	ls1 := []rune{97, 98, 99, 100}
	T_ls1 := reflect.TypeOf(ls1)
	fmt.Println(T_ls1.Elem())
	//输出：
	//int32

	ls2 := []int16{97, 98, 99, 100}
	T_ls2 := reflect.TypeOf(ls2)
	fmt.Println(T_ls2.Elem())
	//输出：
	//int16

	ls3 := []string{"ab", "cd", "ef"}
	T_ls3 := reflect.TypeOf(ls3)
	fmt.Println(T_ls3.Elem())
	//输出：
	//string

	str5 := "ab"
	str6 := "cd"
	str7 := "ef"
	ls4 := []*string{&str5, &str6, &str7}
	T_ls4 := reflect.TypeOf(ls4)
	fmt.Println(T_ls4.Elem())
	//输出：
	//*string

	ls5 := &([]string{"ab", "cd", "ef"})
	T_ls5 := reflect.TypeOf(ls5)
	fmt.Println(T_ls5)
	fmt.Println(T_ls5.Elem())
	//输出：
	//*[]string
	//[]string

	//// Bits returns the size of the type in bits.(Bits返回数字类型type的存储位数大小)
	//// It panics if the type's Kind is not one of the
	//// sized or unsized Int, Uint, Float, or Complex kinds.（如果类型的kind类型不是unsized Int, Uint, Float, or Complex 的话会跑出异常）
	//fmt.Println(T.Bits())
	////报错：
	////panic: reflect: Bits of non-arithmetic Type string
	////
	////goroutine 1 [running]:
	////reflect.(*rtype).Bits(0x4ab140, 0xc000006018)
	////	C:/Go/src/reflect/type.go:776 +0xd1
	////main.main()
	////	C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:232 +0x118f

	i := int(15)
	T_int := reflect.TypeOf(i)
	fmt.Println(T_int.Bits())
	//输出：
	//64

	i32 := int32(15)
	T_int32 := reflect.TypeOf(i32)
	fmt.Println(T_int32.Bits())
	//输出：
	//32

	i8 := int8(15)
	T_int8 := reflect.TypeOf(i8)
	fmt.Println(T_int8.Bits())
	//输出：
	//8

	f32 := float32(3.14)
	T_float32 := reflect.TypeOf(f32)
	fmt.Println(T_float32.Bits())
	//输出：
	//32

	f64 := float64(3.14)
	T_float64 := reflect.TypeOf(f64)
	fmt.Println(T_float64.Bits())
	//输出：
	//64

	fu8 := uint8(3)
	T_uint8 := reflect.TypeOf(fu8)
	fmt.Println(T_uint8.Bits())
	//输出：
	//8

	Comp64 := complex64(3 + 2i)
	T_complex64 := reflect.TypeOf(Comp64)
	fmt.Println(Comp64)
	fmt.Println(T_complex64.Bits())
	//输出：
	//(3+2i)
	//64

	Comp128 := complex128(3 + 2i)
	T_complex128 := reflect.TypeOf(Comp128)
	fmt.Println(Comp128)
	fmt.Println(T_complex128.Bits())
	//输出：
	//(3+2i)
	//128

	//Bits()是返回一个数字类型的位长，但是下面的是返回任意类型的存储所需要的字节个数！一个字节等于8位（bit）
	// 返回要保存一个该类型的值需要多少字节；类似unsafe.Sizeof
	fmt.Println(T.Size())
	//输出：
	//16

	str111 := "abc"
	fmt.Println(reflect.TypeOf(str111).Size())
	//输出：
	//16

	str111 = "123456789abcdef012345678903444444444444444444444444"
	fmt.Println(reflect.TypeOf(str111).Size())
	//输出：
	//16，注意，他不是返回值的长度，而是返回值的类型名的长度，也就是len(string),所以不关值的长度大小！

	int111 := 123456
	fmt.Println(reflect.TypeOf(int111).Size())
	//输出：
	//8
	int111 = 1234567890
	T_int111 := reflect.TypeOf(int111)
	fmt.Println(T_int111)
	fmt.Println(T_int111.Size())
	//输出：
	//int
	//8，这表示保存int这个类型名所需要的长度！

	int222 := int64(123456)
	fmt.Println(reflect.TypeOf(int222).Size())
	//输出：
	//8

	b1 := true
	fmt.Println(reflect.TypeOf(b1).Size())
	//输出：
	//1

	//// NumIn returns a function type's input parameter count.(NumIn返回一个函数的参数个数)
	//// It panics if the type's Kind is not Func.（如果调用者不是一个函数类型的话会抛出异常）
	//fmt.Println(T.NumIn())
	////报错：
	////panic: reflect: NumIn of non-func type
	////
	////goroutine 1 [running]:
	////reflect.(*rtype).NumIn(0x4ac140, 0xc000006018)
	////	C:/Go/src/reflect/type.go:989 +0x65
	////main.main()
	////	C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:341 +0x1ed7

	fun := func(str1, str2 string, i int, b bool) string {
		return "abc"
	}
	T_fun := reflect.TypeOf(fun)
	fmt.Println(T_fun.NumIn())
	//上面的 方法是获取函数的参数个数，但是下面的方法则是获取第几个参数的类型！
	// In returns the type of a function type's i'th input parameter.（in返回第几个参数的类型）
	// It panics if the type's Kind is not Func.（如果底层不是函数类型的话会抛出异常的）
	// It panics if i is not in the range [0, NumIn()).（参数i必须在[0, NumIn())之间）
	fmt.Println(T_fun.In(0))
	fmt.Println(T_fun.In(1))
	fmt.Println(T_fun.In(2))
	fmt.Println(T_fun.In(3))

	//这里千万不能复用fun变量了，因为2个函数的参数不同，也就是是不同的类型了，函数的参数不同的表示的就是不同的函数类型
	fun1 := func(str1 string, i int, b bool) string {
		return "abc"
	}
	T_fun1 := reflect.TypeOf(fun1)
	fmt.Println(T_fun1.NumIn())
	fmt.Println(T_fun1.In(0))
	fmt.Println(T_fun1.In(1))
	fmt.Println(T_fun1.In(2))
	//fmt.Println(T_fun1.In(3))//这个超过了函数的个数数目
	//输出：
	//	4
	//	string
	//	string
	//	int
	//	bool
	//	3
	//	string
	//	int
	//	bool
	//	panic: runtime error: index out of range [3] with length 3
	//
	//	goroutine 1 [running]:
	//	reflect.(*rtype).In(0x4b5620, 0x3, 0xc00007bbc0, 0x1)
	//		C:/Go/src/reflect/type.go:960 +0xaf
	//	main.main()
	//		C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:377 +0x2452

	// NumMethod returns the number of exported methods in the type's method set.
	// (NumMethod返回在类型（这个类型一般是接口）的方法集合中可导出（第一个字母大写）的方法的数目)
	fmt.Println(T.NumMethod())
	//输出：
	//0

	ls_b := []byte{97, 98, 99, 100, 101}
	T_lsByte := reflect.TypeOf(ls_b)
	fmt.Println(T_lsByte.NumMethod())
	//输出：
	//0

	var p = person{"NAKO", 20, true}
	p.Setname111("anko")
	p.Setage111()
	fmt.Println(p)
	T_person := reflect.TypeOf(p)
	fmt.Println(T_person.NumMethod())

	var behav behavior
	behav = &p
	T_behavior := reflect.TypeOf(behav)
	fmt.Println(T_behavior)
	fmt.Println(T_behavior.NumMethod())
	//输出：
	//{anko 18 true}
	//0
	//*main.person
	//2

	// Method returns the i'th method in the type's method set.
	// It panics if i is not in the range [0, NumMethod()).
	//
	// For a non-interface type T or *T, the returned Method's Type and Func
	// fields describe a function whose first argument is the receiver.
	//
	// For an interface type, the returned Method's Type field gives the
	// method signature, without a receiver, and the Func field is nil.
	//
	// Only exported methods are accessible and they are sorted in
	// lexicographic order.
	fmt.Println(T_behavior.Method(0))
	fmt.Println(T_behavior.Method(1))
	//fmt.Println(T_behavior.Method(2))//超出了接口类型的方法个数
	//输出：
	//{Setage111  func(*main.person) <func(*main.person) Value> 0}
	//{Setname111  func(*main.person, string) error <func(*main.person, string) error Value> 1}
	//panic: reflect: Method index out of range
	//
	//goroutine 1 [running]:
	//reflect.(*rtype).Method(0x4f0680, 0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, ...)
	//	C:/Go/src/reflect/type.go:814 +0x76c
	//main.main()
	//	C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:438 +0x2ad4
	//从上面可以知道他的返回值是一个很复杂的结构，其实是type Method struct类型，下面是这个类型的说明：


	fmt.Println("---------reflect.Method对象----------")
	//// Method represents a single method.
	////Method代表一个方法。
	//type Method struct {
	//	// Name is the method name.
	//	// PkgPath is the package path that qualifies a lower case (unexported)
	//	// method name. It is empty for upper case (exported) method names.
	//	// The combination of PkgPath and Name uniquely identifies a method
	//	// in a method set.
	//	// See https://golang.org/ref/spec#Uniqueness_of_identifiers
	//	// Name是方法名称。
	//	// PkgPath是限定小写（未导出）方法名称的程序包路径。 大写（导出）的方法名称为空。
	//	// PkgPath和Name的组合唯一标识 方法集 中的方法。
	//	//参见https://golang.org/ref/spec#Uniqueness_of_identifiers
	//	Name    string
	//	PkgPath string
	//
	//	Type  Type  // method type//方法的类型
	//	Func  Value // func with receiver as first argument	//以接收者为第一个参数的func
	//	Index int   // index for Type.Method	// Type.Method的索引
	//}

	fmt.Println("---------通过类实例调用方法----------")
	var Cal =Calculate{}
	ret := reflect.ValueOf(&Cal).MethodByName("Add").Call([]reflect.Value{reflect.ValueOf(3),reflect.ValueOf(4)})
	fmt.Println(ret)
	for key, value := range ret {
		fmt.Println(key,value)
	}


	fmt.Println("---------通过类调用方法----------")

	method, ok := reflect.TypeOf(&Cal).MethodByName("Add")
	if ok{
		fmt.Println("method.Name",method.Name)
		fmt.Println("method.Type",method.Type)
		fmt.Println("method.PkgPath",method.PkgPath)
		fmt.Println("method.Index",method.Index)
		fmt.Println("method.Func",method.Func)

		ret = method.Func.Call([]reflect.Value{reflect.ValueOf(&Cal),reflect.ValueOf(3),reflect.ValueOf(4)})
		fmt.Println(ret)
		for key, value := range ret {
			fmt.Println(key,value)
		}
	}
	//输出：
	//	[<int Value>]
	//	0 7
	//	-------------------
	//	method.Name Add
	//	method.Type func(*main.Calculate, int, int) int
	//	method.PkgPath
	//	method.Index 0
	//	method.Func 0x4c8480
	//	[<int Value>]
	//	0 7


	// MethodByName returns the method with that name in the type's
	// method set and a boolean indicating if the method was found.
	//
	// For a non-interface type T or *T, the returned Method's Type and Func
	// fields describe a function whose first argument is the receiver.
	//
	// For an interface type, the returned Method's Type field gives the
	// method signature, without a receiver, and the Func field is nil.
	fmt.Println(T_behavior.MethodByName("sdsdsds"))
	fmt.Println(T_behavior.MethodByName("Setage111"))
	fmt.Println(T_behavior.MethodByName("Setname111"))
	//输出：
	//{  <nil> <invalid Value> 0} false
	//{Setage111  func(*main.person) <func(*main.person) Value> 0} true
	//{Setname111  func(*main.person, string) error <func(*main.person, string) error Value> 1} true

	// Len returns an array type's length.（Len返回一个基类为array类型（不是切片）实例的值的长度）
	// It panics if the type's Kind is not Array.(如果不是array类型的话会抛出异常)
	arr := [...]string{"hello", "world", "hahah"}
	//slice :=[]string{"hello","world","hahah"}
	fmt.Println(reflect.TypeOf(arr).Len())
	//fmt.Println(reflect.TypeOf(slice).Len())//会报错
	//输出：
	//3
	//panic: reflect: Len of non-array type
	//
	//goroutine 1 [running]:
	//reflect.(*rtype).Len(0x4e30a0, 0x3)
	//	C:/Go/src/reflect/type.go:973 +0x65
	//main.main()
	//	C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:496 +0x30b1

	// Align returns the alignment in bytes of a value of
	// this type when allocated in memory.
	//说白了就是返回分配内存时相应类型所分配的类型内存字节数
	fmt.Println()
	fmt.Println(T.Align())
	fmt.Println(reflect.TypeOf(arr).Align())
	fmt.Println(reflect.TypeOf(int16(34)).Align())
	fmt.Println(reflect.TypeOf(int16(340)).Align())
	fmt.Println(reflect.TypeOf(int16(3400)).Align())
	fmt.Println(reflect.TypeOf(int8(34)).Align())
	fmt.Println(reflect.TypeOf(int64(34)).Align())
	fmt.Println(reflect.TypeOf(int64(34000)).Align())
	fmt.Println(reflect.TypeOf(interface{}(34)).Align())
	fmt.Println(reflect.TypeOf(string("abc")).Align())
	fmt.Println(reflect.TypeOf(string("abcsdsdsdsdddddddddddddddddddddddddddddddd")).Align())
	fmt.Println(reflect.TypeOf(string("")).Align())
	fmt.Println(reflect.TypeOf(struct{}{}).Align())
	fmt.Println(reflect.TypeOf(struct {
		string
		int
	}{"anko", 15}).Align())
	//注意在创建匿名结构体的 时候两两字段的类型不能相同，也就是结构体中的某一个类型只允许出现一次，同时字段不允许是切片或者数组类型
	fmt.Println(reflect.TypeOf(struct {
		string
		bool
		int
		byte
	}{"anko", true, 90, 'a'}).Align())
	fmt.Println(reflect.TypeOf(struct {
		string
		bool
		int
		byte
	}{}).Align())
	//输出：
	//	8
	//	8
	//	2
	//	2
	//	2
	//	1
	//	8
	//	8
	//	8
	//	8
	//	8
	//	8
	//	1
	//	8
	//	8
	//	8

	// PkgPath returns a defined type's package path, that is, the import path
	// that uniquely identifies the package, such as "encoding/base64".
	// If the type was predeclared (string, error) or not defined (*T, struct{},
	// []int, or A where A is an alias for a non-defined type), the package path
	// will be the empty string.
	//翻译：PkgPath返回类型的包路径，即明确指定包的import路径，如"encoding/base64"
	//    如果类型为内建类型(string, error)或未预先定义类型(*T, struct{}, []int)，会返回空字符串
	//Pkg：packeage的简写
	fmt.Println(T.PkgPath())
	fmt.Println(reflect.TypeOf("bs64").PkgPath())
	fmt.Println(reflect.TypeOf(int(14)).PkgPath())
	var ii = int64(43)
	fmt.Println(reflect.TypeOf(ii).PkgPath())
	//输出：以上都是内建的类型，返回值都是空字符串

	var ls_i []int
	ls_i = []int{1, 2, 34}
	fmt.Println(reflect.TypeOf(ls_i).PkgPath())
	fmt.Println(reflect.TypeOf(struct{}{}).PkgPath())
	fmt.Println(reflect.TypeOf(&struct{}{}).PkgPath())
	//输出：以上都是未预先定义类型，返回值都是空字符串

	var bs64 base642.Encoding
	fmt.Println(reflect.TypeOf(bs64).PkgPath())
	//输出：
	//encoding/base64

	var buf1 bytes.Buffer
	fmt.Println(reflect.TypeOf(buf1).PkgPath())
	//输出：
	//bytes

	var p1 = person{"p1", 23, false}
	T_p1 := reflect.TypeOf(p1)
	// NumField returns a struct type's field count.（NumField返回一个结构体的字段总个数）
	// It panics if the type's Kind is not Struct.（如果Field的底层类型不是Struct的话会panic）
	fmt.Println(T_p1.NumField())
	// Field returns a struct type's i'th field.（Field返回一个结构体的第i个字段的类型）
	// It panics if the type's Kind is not Struct.（如果Field的底层类型不是Struct的话会panic）
	// It panics if i is not in the range [0, NumField()).（如果i不在【0, NumField()】之间的话也会抛出异常，即索引不能超过结构体字段的个数）
	//这个方法返回一个对象，的说明如下：

	//// A StructField describes a single field in a struct.（A StructField描述在结构体中的单一字段）
	//type StructField struct {
	//	// Name is the field name.（Name是字段的名字）
	//	Name string
	//	// PkgPath is the package path that qualifies a lower case (unexported)
	//	// field name. It is empty for upper case (exported) field names.（PkgPath是非导出字段的包路径，对导出字段该字段为""。）
	//	// See https://golang.org/ref/spec#Uniqueness_of_identifiers
	//	PkgPath string
	//
	//	Type      Type      // field type（字段的类型）
	//	Tag       StructTag // field tag string（字段的标签tag）
	//	Offset    uintptr   // offset within struct, in bytes(字段在结构体的字节水平的下标索引位置，struct的第几个字节是该字段的起始位置)
	//	Index     []int     // index sequence for Type.FieldByIndex（字段在结构体中的索引，表示按照声明时候的是该结构体中的第几个字段，跟上面的Offset完全不是一个东西,主要用于Type.FieldByIndex方法）
	//	Anonymous bool      // is an embedded field（是否是匿名字段？）
	//}

	fmt.Printf("%+v\n", T_p1.Field(0))
	fmt.Printf("%+v\n", T_p1.Field(1))
	fmt.Printf("%+v\n", T_p1.Field(2))
	//fmt.Printf("%+v\n",T_p1.Field(3))//报错，超出了结构体字段的个数
	//输出：
	//3
	//{Name:Name PkgPath: Type:string Tag: Offset:0 Index:[0] Anonymous:false}
	//{Name:Age PkgPath: Type:int Tag: Offset:16 Index:[1] Anonymous:false}
	//{Name:isstudent PkgPath:main Type:bool Tag: Offset:24 Index:[2] Anonymous:false}

	ExampleStructTag_Lookup() //这个函数main外的下面被定义
	//输出：
	// field_0
	// (blank)
	// (not specified)

	// FieldAlign returns the alignment in bytes of a value of
	// this type when used as a field in a struct.
	// 返回当该类型作为结构体的字段时，会对齐（占领）的字节数
	fmt.Println(T_p1.FieldAlign())
	fmt.Println(reflect.TypeOf(struct{}{}).FieldAlign())
	fmt.Println(reflect.TypeOf([]byte{'a', 'b'}).FieldAlign())
	fmt.Println(reflect.TypeOf([]byte{'a', 'b', 'c'}).FieldAlign())
	fmt.Println(reflect.TypeOf([]int{'a', 'b', 'c'}).FieldAlign())
	fmt.Println(reflect.TypeOf([]int64{'a', 'b', 'c'}).FieldAlign())
	fmt.Println(reflect.TypeOf([]string{"wo", "de"}).FieldAlign())
	fmt.Println(reflect.TypeOf([]string{"我的世界", "很大"}).FieldAlign())
	fmt.Println(reflect.TypeOf("我的世界很大").FieldAlign())
	//输出：
	//	8
	//	1
	//	8
	//	8
	//	8
	//	8
	//	8
	//	8
	//	8

	// FieldByIndex returns the nested field corresponding
	// to the index sequence. It is equivalent to calling Field
	// successively for each index i.
	// It panics if the type's Kind is not Struct.
	// FieldByIndex返回字段序列中的第i个嵌套的字段对应的StructField对象和是否找到该字段的bool值，如非结构体将会panic
	fmt.Printf("%+v\n", T_p1.FieldByIndex([]int{0}))
	fmt.Printf("%+v\n", T_p1.FieldByIndex([]int{1}))
	fmt.Printf("%+v\n", T_p1.FieldByIndex([]int{2}))
	//fmt.Printf("%+v",T_p1.FieldByIndex([]int{0,1}))//报错，继续看下面会解释原因
	//输出：
	//Name:Name PkgPath: Type:string Tag: Offset:0 Index:[0] Anonymous:false}
	//{Name:Age PkgPath: Type:int Tag: Offset:16 Index:[1] Anonymous:false}
	//{Name:isstudent PkgPath:main Type:bool Tag: Offset:24 Index:[2] Anonymous:false}

	var p11 = person11{p1}
	T_p11 := reflect.TypeOf(p11)
	fmt.Printf("%+v\n", T_p11.FieldByIndex([]int{0}))    //第一层级的第0索引字段
	fmt.Printf("%+v\n", T_p11.FieldByIndex([]int{0, 0})) //第一层级的第0索引字段包含的第二层级的第0索引的字段
	fmt.Printf("%+v\n", T_p11.FieldByIndex([]int{0, 1})) //第一层级的第0索引字段包含的第二层级的第1索引的字段
	fmt.Printf("%+v\n", T_p11.FieldByIndex([]int{0, 2})) //第一层级的第0索引字段包含的第二层级的第2索引的字段
	//fmt.Printf("%+v\n",T_p11.FieldByIndex([]int{0,2,0}))//报错，最多2层级而已
	//输出：
	//{Name:person PkgPath:main Type:main.person Tag: Offset:0 Index:[0] Anonymous:true}
	//{Name:Name PkgPath: Type:string Tag: Offset:0 Index:[0] Anonymous:false}
	//{Name:Age PkgPath: Type:int Tag: Offset:16 Index:[1] Anonymous:false}
	//{Name:isstudent PkgPath:main Type:bool Tag: Offset:24 Index:[2] Anonymous:false}

	// FieldByName returns the struct field with the given name
	// and a boolean indicating if the field was found.
	// 返回该类型名为name的字段（会查找匿名字段及其子字段），
	// 布尔值说明是否找到，如非结构体将panic
	fmt.Println()
	fmt.Println(T_p1.FieldByName("name"))
	fmt.Println(T_p1.FieldByName("Name"))
	fmt.Println(T_p1.FieldByName("Age"))
	fmt.Println(T_p1.FieldByName("isstudent"))

	fmt.Println()
	fmt.Println(T_p11.FieldByName("name"))
	fmt.Println(T_p11.FieldByName("person"))
	fmt.Println(T_p11.FieldByName("Name"))
	fmt.Println(T_p11.FieldByName("Age"))
	fmt.Println(T_p11.FieldByName("isstudent"))
	//输出：
	//{  <nil>  0 [] false} false
	//{Name  string  0 [0] false} true
	//{Age  int  16 [1] false} true
	//{isstudent main bool  24 [2] false} true
	//
	//{  <nil>  0 [] false} false
	//{person main main.person  0 [0] true} true
	//{Name  string  0 [0 0] false} true
	//{Age  int  16 [0 1] false} true
	//{isstudent main bool  24 [0 2] false} true

	// FieldByNameFunc returns the struct field with a name
	// that satisfies the match function and a boolean indicating if
	// the field was found.（返回结构体中第一个字段的name满足函数func的字段对应的StructField对象和是否找到该字段的bool值，如非结构体将会panic（会查找匿名字段及其子字段））
	//
	// FieldByNameFunc considers the fields in the struct itself
	// and then the fields in any embedded structs, in breadth first order,
	// stopping at the shallowest nesting depth containing one or more
	// fields satisfying the match function. If multiple fields at that depth
	// satisfy the match function, they cancel each other
	// and FieldByNameFunc returns no match.
	// This behavior mirrors Go's handling of name lookup in
	// structs containing embedded fields.(没翻译)
	fmt.Println()
	fmt.Println(T_p1.FieldByNameFunc(func(s string) bool {
		if s == "Name" { //其实这里可以按照某个正则规则来进行查找的！但是返回的却永远是第一个符合该条件的字段，下同
			return true
		}
		return false
	}))

	fmt.Println(T_p1.FieldByNameFunc(func(s string) bool {
		if s == "Age" {
			return true
		}
		return false
	}))

	fmt.Println(T_p1.FieldByNameFunc(func(s string) bool {
		if s == "isstudent" {
			return true
		}
		return false
	}))

	fmt.Println()
	fmt.Println(T_p11.NumField())
	fmt.Println(T_p11.FieldByNameFunc(func(s string) bool {
		if s == "person" {
			return true
		}
		return false
	}))

	fmt.Println(T_p11.FieldByNameFunc(func(s string) bool {
		if s == "Name" {
			return true
		}
		return false
	}))

	fmt.Println(T_p11.FieldByNameFunc(func(s string) bool {
		if s == "Age" {
			return true
		}
		return false
	}))

	fmt.Println(T_p11.FieldByNameFunc(func(s string) bool {
		if s == "isstudent" {
			return true
		}
		return false
	}))
	//以上输出：
	//{Name  string  0 [0] false} true
	//{Age  int  16 [1] false} true
	//{isstudent main bool  24 [2] false} true
	//
	//1
	//{person main main.person  0 [0] true} true
	//{Name  string  0 [0 0] false} true
	//{Age  int  16 [0 1] false} true
	//{isstudent main bool  24 [0 2] false} true

	//// A Kind represents the specific kind of type that a Type represents.
	//// The zero Kind is not a valid kind.
	//type Kind uint
	//const (
	//	Invalid Kind = iota
	//	Bool
	//	Int
	//	Int8
	//	Int16
	//	Int32
	//	Int64
	//	Uint
	//	Uint8
	//	Uint16
	//	Uint32
	//	Uint64
	//	Uintptr
	//	Float32
	//	Float64
	//	Complex64
	//	Complex128
	//	Array
	//	Chan
	//	Func
	//	Interface
	//	Map
	//	Ptr
	//	Slice
	//	String
	//	Struct
	//	UnsafePointer
	//)

	fmt.Println()
	fmt.Println(T_p1.Kind())
	fmt.Println(T_p11.Kind())
	fmt.Println(T.Kind())
	fmt.Println(reflect.TypeOf([]byte{'a', 'b'}).Kind())
	fmt.Println(reflect.TypeOf([2]byte{'a', 'b'}).Kind())
	fmt.Println(reflect.TypeOf([2]int{'a', 'b'}).Kind())
	fmt.Println(reflect.TypeOf(int('a')).Kind())
	fmt.Println(reflect.TypeOf(uint('a')).Kind())
	fmt.Println(reflect.TypeOf(float32('a')).Kind())
	fmt.Println(reflect.TypeOf(float64('a')).Kind())
	fmt.Println(reflect.TypeOf(complex64('a')).Kind())
	fmt.Println(reflect.TypeOf(complex128('a')).Kind())
	fmt.Println(reflect.TypeOf(true).Kind())
	fmt.Println(reflect.TypeOf(map[string]int{"a": 97, "b": 98}).Kind())
	fmt.Println(reflect.TypeOf(interface{}(16)).Kind())
	fmt.Println(reflect.TypeOf(interface{}("aaa")).Kind())
	fmt.Println(reflect.TypeOf((*int)(nil)).Kind())
	fmt.Println(reflect.TypeOf((func())(nil)).Kind())
	fmt.Println(reflect.TypeOf((chan int)(nil)).Kind())
	fmt.Println(reflect.TypeOf((map[string]int)(nil)).Kind())
	fmt.Println(reflect.TypeOf(([]byte)(nil)).Kind())
	fmt.Println(reflect.TypeOf(struct{}{}).Kind())
	fmt.Println(reflect.TypeOf(person{}).Kind())
	//fmt.Println(reflect.TypeOf(([2]byte)(nil)).Kind())//报错，nil可以转化为引用类型，但是不能转化为值类型
	//fmt.Println(reflect.TypeOf(interface{}(nil)).Kind())//报错，nil可以转化为引用类型，但是不能转化为值类型
	//输出：
	//struct
	//struct
	//string
	//slice
	//array
	//array
	//int
	//uint
	//float32
	//float64
	//complex64
	//complex128
	//bool
	//map
	//int
	//string
	//ptr
	//func
	//chan
	//map
	//slice
	//struct

	// Key returns a map type's key type.
	// It panics if the type's Kind is not Map.
	// Key返回一个map的键key的类型，非map类型的话会抛出异常
	//fmt.Println(T_p1.Key())//报错
	fmt.Println()
	m1 := map[string]int{"a": 97, "b": 98}
	m2 := map[interface{}]int{nil: 97, nil: 98} //nil也可以作为键
	m3 := map[person]int{person{"a", 13, true}: 97, person{"b", 13, true}: 98}
	var beh1 behavior
	var beh2 behavior
	m4 := map[behavior]int{beh1: 76, beh2: 87}
	T_m1 := reflect.TypeOf(m1)
	T_m2 := reflect.TypeOf(m2)
	T_m3 := reflect.TypeOf(m3)
	T_m4 := reflect.TypeOf(m4)
	fmt.Println(T_m1.Key())
	fmt.Println(T_m2.Key())
	fmt.Println(T_m3.Key())
	fmt.Println(T_m4.Key())
	//输出：
	//string
	//interface {}
	//main.person
	//main.behavior

	fmt.Println()

	T_f2 := reflect.TypeOf(f2)
	T_f3 := reflect.TypeOf(f3)
	// NumOut returns a function type's output parameter count.(返回函数的返回值个数，非函数类型会报错)
	// It panics if the type's Kind is not Func.
	fmt.Println(T_f2.NumOut())
	fmt.Println(T_f3.NumOut())

	// Out returns the type of a function type's i'th output parameter.（返回函数的第i个返回值的类型，非函数类型会报错，i必须属于[0, NumOut())之中的整数）
	// It panics if the type's Kind is not Func.
	// It panics if i is not in the range [0, NumOut()).
	fmt.Println(T_f2.Out(0))
	fmt.Println(T_f2.Out(1))
	//输出：
	//2
	//0
	//string
	//error

	fmt.Println()
	f4 := func(int, ...float64) {}
	f5 := func(int, float64) {}
	T_f4 := reflect.TypeOf(f4)
	T_f5 := reflect.TypeOf(f5)
	// IsVariadic reports whether a function type's final input parameter
	// is a "..." parameter. If so, t.In(t.NumIn() - 1) returns the parameter's
	// implicit actual type []T.
	//
	// For concreteness, if t represents func(x int, y ... float64), then
	//
	//	t.NumIn() == 2
	//	t.In(0) is the reflect.Type for "int"
	//	t.In(1) is the reflect.Type for "[]float64"
	//	t.IsVariadic() == true
	//
	// IsVariadic panics if the type's Kind is not Func.
	//IsVariadic判断函数的参数中是否有...int(或者其他类型，不一定是int)这样的类型，非函数类型的话会报错
	fmt.Println(T_f4.IsVariadic())
	fmt.Println(T_f4.NumIn())
	fmt.Println(T_f4.In(0))
	fmt.Println(T_f4.In(1))

	fmt.Println(T_f5.IsVariadic())
	fmt.Println(T_f5.NumIn())
	fmt.Println(T_f5.In(0))
	fmt.Println(T_f5.In(1))
	//fmt.Println(T.IsVariadic())//报错，类型不对
	//输出：
	//	true
	//	2
	//	int
	//	[]float64,虽然在这里输出的是[]float64，但是我们在定义函数的参数时候不能这样写，这跟...float64这样定义参数的函数的调用方式是不一样的！
	//	false
	//	2
	//	int
	//	float64

	fmt.Println()
	c_all := make(chan int, 2)   ///定义存取int类型值的通道（chan和int之间必须空格）
	c_in := make(chan<- int, 2)  //定义存int类型值的通道（chan和int之间可以空格，非必须，但是建议）
	c_out := make(<-chan int, 2) //定义取int类型值的通道（chan和int之间必须空格）
	//var c_out <-chan
	T_c_all := reflect.TypeOf(c_all)
	T_c_in := reflect.TypeOf(c_in)
	T_c_out := reflect.TypeOf(c_out)
	// ChanDir returns a channel type's direction.（ChanDir返回一个chan的具体类型，是取值通道 还是 存值通道 或者是既可以取值又可以存值的通道）
	// It panics if the type's Kind is not Chan.（非通道类型会报错）
	fmt.Println(T_c_all.ChanDir())
	fmt.Println(T_c_in.ChanDir())
	fmt.Println(T_c_out.ChanDir())
	//输出：
	//chan
	//chan<-
	//<-chan

	fmt.Println()
	// Comparable reports whether values of this type are comparable.（判断类型的值是否可以比较）
	fmt.Println(T.Comparable())
	fmt.Println(reflect.TypeOf(int(24)).Comparable())
	fmt.Println(reflect.TypeOf(int64(24)).Comparable())
	fmt.Println(reflect.TypeOf(string(24)).Comparable())
	fmt.Println(reflect.TypeOf(bool(true)).Comparable())
	fmt.Println(reflect.TypeOf(struct{}{}).Comparable())

	var k interface{}
	k = "abc" //必须赋值之后下面才不会报错，不赋值的话代表未申请内存
	fmt.Println(reflect.TypeOf(k).Comparable())
	var pp1 person
	fmt.Println(reflect.TypeOf(pp1).Comparable())

	var bb1 behavior
	bb1 = &pp1 //必须赋值之后下面才不会报错，不赋值的话代表未申请内存
	fmt.Println(reflect.TypeOf(bb1).Comparable())

	fmt.Println(reflect.TypeOf(&struct{}{}).Comparable())
	fmt.Println(reflect.TypeOf("sdsds"[0]).Comparable())
	fmt.Println(reflect.TypeOf([2]byte{'a', 'b'}).Comparable())

	//含有map和切片的类型都不可以比较，而不是说数组就一定可以比较的！
	fmt.Println(reflect.TypeOf([2]map[string]int{map[string]int{"a": 1, "b": 2}, map[string]int{"c": 11, "d": 22}}).Comparable())
	fmt.Println(reflect.TypeOf(map[string]int{"a": 0, "b": 1}).Comparable())
	fmt.Println(reflect.TypeOf([]byte{'a', 'b'}).Comparable())
	fmt.Println(reflect.TypeOf([]interface{}{'a', 'b'}).Comparable())
	//输出：
	//	true
	//	true
	//	true
	//	true
	//	true
	//	true
	//	true
	//	true
	//	true
	//	true
	//	true
	//	true
	//	false
	//	false
	//	false
	//	false

	fmt.Println()

	var r *io.Reader
	fmt.Println(reflect.TypeOf(r))
	fmt.Println(reflect.TypeOf(&os.Stdout))
	//报错，panic: reflect: non-interface type passed to Type.Implements，我们必须获取到r接口类型而不是接口对应的指针类型,而
	//要获取对应的接口类型的话我们只能使用elem()方法来对指针类型进行获取！包括下面的AssignableTo（）和ConvertibleTo（）方法都是如此！
	//fmt.Println(reflect.TypeOf(&os.Stdout).Implements(reflect.TypeOf(r)))

	// Elem returns a type's element type.(Elem返回一个类型对象的子元素对象的类型)
	// It panics if the type's Kind is not Array, Chan, Map, Ptr, or Slice.（如果此类型的基类不是Array, Chan, Map, Ptr, or Slice的话会报错）
	T_R_elem := reflect.TypeOf(r).Elem()
	fmt.Println(T_R_elem)
	fmt.Println(T_R_elem.Kind())
	T_stdout_elem := reflect.TypeOf(&os.Stdout).Elem()
	fmt.Println(T_stdout_elem)
	fmt.Println(T_stdout_elem.Kind())
	// Implements reports whether the type implements the interface type u.（判断一个类型是否实现了某个接口类型）
	fmt.Println(T_stdout_elem.Implements(T_R_elem))
	//输出：
	//	*io.Reader
	//	**os.File
	//	io.Reader
	//	interface
	//	*os.File
	//	ptr
	//	true

	fmt.Println()
	// AssignableTo reports whether a value of the type is assignable to type u.
	// 如果该类型的值可以直接赋值给u代表的类型，返回真
	fmt.Println(T_stdout_elem.AssignableTo(T_R_elem))

	fmt.Println()
	var beha *behavior
	//beha=&p1//报错，接口的指针类型的不允许被赋值的！接口类型允许被赋值！
	var p0 *person
	//var p0 *person=&person{"anko",57,true}//无论初始化与否都是可以的！

	T_p0 := reflect.TypeOf(&p0) //必须取指针才可以有子元素（指针类型的子元素才是p0）
	T_beha := reflect.TypeOf(beha)
	elem_p0 := T_p0.Elem()
	elem_beha := T_beha.Elem()
	fmt.Println(T_p0)
	fmt.Println(T_beha)
	fmt.Println(elem_p0)
	fmt.Println(elem_beha)

	fmt.Println(elem_p0.Kind())
	fmt.Println(elem_beha.Kind())
	// ConvertibleTo reports whether a value of the type is convertible to type u.
	// 如该类型的值可以转换为u代表的类型，返回真
	fmt.Println(elem_p0.ConvertibleTo(elem_beha))
	fmt.Println(T_stdout_elem.ConvertibleTo(T_R_elem))
	//输出：
	//	**main.person
	//	*main.behavior
	//	*main.person
	//	main.behavior
	//	ptr
	//	interface
	//	true
	//	true

	fmt.Println()
	fmt.Println("-------------reflect.Value对象-----------------")

	//ValueOf返回一个初始化为i接口保管的具体值的Value，ValueOf(nil)返回Value零值。
	V := reflect.ValueOf(str)
	fmt.Println("str对象的值为：", V)
	//输出：
	//	str对象的值为： hello world!

	// Kind returns v's Kind.(Kind返回v的类型)
	// If v is the zero Value (IsValid returns false), Kind returns Invalid.（如果v是零值（IsValid返回false的值），Kind返回Invalid）
	fmt.Println(V.Kind())

	V1 := reflect.ValueOf(nil)
	fmt.Println(V1.Kind())
	// IsValid reports whether v represents a value.(IsValid返回v是否是零值即是否初始化了)
	// It returns false if v is the zero Value.(如果是零值的话他会返回false)
	// If IsValid returns false, all other methods except String panic.
	// Most functions and methods never return an invalid value.
	// If one does, its documentation states the conditions explicitly.
	//IsValid返回v是否持有一个值。如果v是Value零值(和零值的value是不同的，前者是代表value对象无值，后者代表的是value对象相对应的实例的类型的初始值)会返回假，此时v除了IsValid、String、Kind之外的方法都会导致panic。
	//绝大多数函数和方法一般不返回Value零值。如果某个函数/方法返回了非法的Value，它的文档必须显式的说明具体情况。
	fmt.Println(V1.IsValid())
	//输出：
	//string
	//invalid
	//false

	// Elem returns the value that the interface v contains
	// or that the pointer v points to.（Elem返回接口v包含的Value对象或者指针v指向的Value对象）
	// It panics if v's Kind is not Interface or Ptr.（如果v的kind类型不是Interface or Ptr，就会抛出异常）
	//fmt.Println(V.Elem())//报错，panic: reflect: call of reflect.Value.Elem on string Value
	fmt.Println()
	ptr_LsB := reflect.ValueOf(&[]byte{'a', 'b'})
	fmt.Println(ptr_LsB.Kind())
	fmt.Println(ptr_LsB.Elem().Kind())
	fmt.Println(ptr_LsB.Elem())
	//输出：
	//ptr
	//slice
	//[97 98]

	fmt.Println()
	//ptr_LsB = reflect.ValueOf(&[]byte{'a', 'b'})//千万不可用这个类型来获取长度，他的kind是ptr
	ptr_LsB = reflect.ValueOf([]byte{'a', 'b'})
	ptr_LsB1 := reflect.ValueOf([]byte{'a', 'b', 'c'})
	ptr_LsB2 := reflect.ValueOf([...]byte{'a', 'b', 'c'})
	ptr_LsB3 := reflect.ValueOf("hello")
	ptr_LsB4 := reflect.ValueOf(make(chan int, 2))
	c1 := make(chan int, 2)
	c1 <- 2
	c1 <- 3
	ptr_LsB41 := reflect.ValueOf(c1)
	ptr_LsB5 := reflect.ValueOf(make(map[string]int, 2))

	map1 := make(map[string]int, 2) //最开始时候申请多少个元素的内存，如果我们下面给出超过了初始化时候的内存的话，map会自动在原地上扩增内存的！
	fmt.Printf("%p\n", map1)
	map1["a"] = 1
	map1["b"] = 2
	map1["c"] = 3
	fmt.Printf("%p\n", map1)
	ptr_LsB51 := reflect.ValueOf(map1)
	// Len returns v's length.(Len返回v的长度)
	// It panics if v's Kind is not Array, Chan, Map, Slice, or String.(调用者的kind类型如果不是Array, Chan, Map, Slice, or String则会抛出异常)
	fmt.Println(ptr_LsB.Len())
	fmt.Println(ptr_LsB1.Len())
	fmt.Println(ptr_LsB2.Len())
	fmt.Println(ptr_LsB3.Len())
	fmt.Println(ptr_LsB4.Len())
	fmt.Println(ptr_LsB41.Len())
	fmt.Println(ptr_LsB5.Len())
	fmt.Println(ptr_LsB51.Len())
	//输出：
	//0xc00005e8d0
	//0xc00005e8d0
	//2
	//3
	//3
	//5
	//0
	//2
	//0
	//2

	String_v2 := reflect.ValueOf("hello world")
	String_v3 := reflect.ValueOf([]byte{'a', 'b', 'c'})
	String_v4 := reflect.ValueOf([2]string{"hello", "world"})
	// String returns the string v's underlying value, as a string.
	// String is a special case because of Go's String method convention.
	// Unlike the other getters, it does not panic if v's Kind is not String.
	// Instead, it returns a string of the form "<T value>" where T is v's type.
	// The fmt package treats Values specially. It does not call their String
	// method implicitly but instead prints the concrete values they hold.
	//返回v持有的值的字符串表示。因为go的String方法的惯例，Value的String方法比较特别。和其他获取v持有值的方法不同：
	//v的Kind是String时，返回该字符串；v的Kind不是String时也不会panic而是返回格式为"<T value>"的字符串，其中T是v持有值的类型。
	fmt.Println(String_v2.String())
	fmt.Println(String_v3.String())
	fmt.Println(String_v4.String())
	//输出：
	//hello world
	//<[]uint8 Value>
	//<[2]string Value>

	fmt.Println()
	int_v2 := reflect.ValueOf(14)
	int_v3 := reflect.ValueOf(int8(15))
	//int_v4 := reflect.ValueOf("hello world")//抛出异常
	// Int returns v's underlying value, as an int64.(Int返回整数类型的v的值，返回值类型是int64)
	// It panics if v's Kind is not Int, Int8, Int16, Int32, or Int64.（如果v的kind类型不是Int, Int8, Int16, Int32, or Int64则会抛出异常）
	fmt.Println(int_v2.Int())
	fmt.Println(int_v3.Int())
	//fmt.Println(int_v4.Int())
	//输出：
	//14
	//15

	fmt.Println()
	interface_v2 := reflect.ValueOf(14)
	interface_v3 := reflect.ValueOf("sdsdsd")
	interface_v4 := reflect.ValueOf(struct{}{})
	interface_v5 := reflect.ValueOf(p1)
	interface_v6 := reflect.ValueOf(pp1)
	interface_v7 := reflect.ValueOf(interface{}(24))
	var bb11 *behavior
	interface_v8 := reflect.ValueOf(bb11) //ValueOf(nil)返回Value零值
	// Interface returns v's current value as an interface{}.
	// It is equivalent to:
	//	var i interface{} = (v's underlying value)
	// It panics if the Value was obtained by accessing
	// unexported struct fields.
	//本方法返回v当前持有的值（表示为保管在interface{}类型），等价于：
	//var i interface{} = (v's underlying value)
	//如果v是通过访问非导出结构体字段获取的，会导致panic。
	i2 := interface_v2.Interface()
	fmt.Println(i2)
	fmt.Println(reflect.TypeOf(i2))
	fmt.Println(reflect.TypeOf(&i2))
	fmt.Println(reflect.TypeOf(&i2).Elem())
	fmt.Println(reflect.TypeOf(&i2).Elem().Kind())

	fmt.Println(interface_v3.Interface())
	fmt.Println(interface_v4.Interface())
	fmt.Println(interface_v5.Interface())
	fmt.Println(interface_v6.Interface())
	i3 := interface_v7.Interface()
	fmt.Println(i3)
	fmt.Println(reflect.TypeOf(i2))
	fmt.Println(reflect.TypeOf(&i2))
	fmt.Println(reflect.TypeOf(&i2).Elem())
	fmt.Println(reflect.TypeOf(&i2).Elem().Kind())

	fmt.Println(interface_v8)
	i4 := interface_v8.Interface()
	fmt.Println(i4)
	fmt.Println(reflect.TypeOf(&i4))
	fmt.Println(reflect.TypeOf(&i4).Elem())
	fmt.Println(reflect.TypeOf(&i4).Elem().Kind())

	//输出：
	//	14
	//	int
	//	*interface {}
	//	interface {}
	//	interface
	//	sdsdsd
	//	{}
	//	{p1 23 false}
	//	{ 0 false}
	//	24
	//	int
	//	*interface {}
	//	interface {}
	//	interface
	//	<nil>
	//	<nil>
	//	*interface {}
	//	interface {}
	//	interface

	fmt.Println()
	bool_v2 := reflect.ValueOf(true)
	//bool_v3 := reflect.ValueOf(14)
	// Bool returns v's underlying value.(返回bool类型对象的值，非bool类型对象的话会报错)
	// It panics if v's kind is not Bool.
	fmt.Println(bool_v2.Bool())
	//fmt.Println(bool_v3.Bool())//报错

	fmt.Println()
	ls_v2 := reflect.ValueOf([]byte{'a', 'b', 'c'}) //注意，字节切片类型，字节切片值（多个元素整体来看，不是拆开元素来看的！）
	//ls_v3 := reflect.ValueOf([3]byte{'a','b','c'})
	// Bytes returns v's underlying value.(返回字节切片v的值，v必须是字节切片，否则会panic)
	// It panics if v's underlying value is not a slice of bytes.
	fmt.Println(ls_v2.Bytes())
	fmt.Println(ls_v2.Kind())
	//fmt.Println(ls_v3.Bytes())//报错，数组的kind不是slice类型，是array类型
	//输出：
	//[97 98 99]
	//slice

	fmt.Println()
	float64_v2 := reflect.ValueOf(23.4)
	//float64_v3:= reflect.ValueOf(23)
	// Float returns v's underlying value, as a float64.
	// It panics if v's Kind is not Float32 or Float64
	fmt.Println(float64_v2.Float())
	fmt.Println(float64_v2.Kind())
	//fmt.Println(float64_v3.Float())//报错
	//输出：
	//23.4
	//float64

	fmt.Println()
	uint_v2 := reflect.ValueOf(uint(34))
	uint_v4 := reflect.ValueOf(uint8(34))
	uint_v5 := reflect.ValueOf(uint64(34))
	//uint_v3:= reflect.ValueOf(23)
	// Float returns v's underlying value, as a float64.
	// It panics if v's Kind is not Float32 or Float64
	fmt.Println(uint_v2.Uint())
	fmt.Println(uint_v2.Kind())
	fmt.Println(uint_v4.Uint())
	fmt.Println(uint_v4.Kind())
	fmt.Println(uint_v5.Uint())
	fmt.Println(uint_v5.Kind())
	//fmt.Println(uint_v3.Uint())//报错
	//输出：
	//34
	//uint
	//34
	//uint8
	//34
	//uint64

	fmt.Println()
	ls_b9 := []byte{'a', 'b'}

	hex_str := fmt.Sprintf("%p", ls_b9)
	dig_int64, err := strconv.ParseInt(hex_str[2:], 16, 64)
	check_err_reflect(err)
	fmt.Println("字节切片的10进制地址值为：", dig_int64)

	Pointer_v2 := reflect.ValueOf(&(ls_b9))
	Pointer_v3 := reflect.ValueOf(ls_b9)
	// Pointer returns v's value as a uintptr.
	// It returns uintptr instead of unsafe.Pointer so that
	// code using reflect cannot obtain unsafe.Pointers
	// without importing the unsafe package explicitly.
	// It panics if v's Kind is not Chan, Func, Map, Ptr, Slice, or UnsafePointer.
	//
	// If v's Kind is Func, the returned pointer is an underlying
	// code pointer, but not necessarily enough to identify a
	// single function uniquely. The only guarantee is that the
	// result is zero if and only if v is a nil func Value.
	//
	// If v's Kind is Slice, the returned pointer is to the first
	// element of the slice. If the slice is nil the returned value
	// is 0.  If the slice is empty but non-nil the return value is non-zero.
	//将v持有的值作为一个指针返回。本方法返回值不是unsafe.Pointer类型，以避免程序员不显式导入unsafe包却得到unsafe.Pointer类型表示的指针。
	//如果v的Kind不是Chan、Func、Map、Ptr、Slice或UnsafePointer会panic。
	//
	//如果v的Kind是Func，返回值是底层代码的指针，但并不足以用于区分不同的函数；只能保证当且仅当v持有函数类型零值nil时，返回值为0。
	//
	//如果v的Kind是Slice，返回值是指向切片第一个元素的指针。如果持有的切片为nil，返回值为0；如果持有的切片没有元素但不是nil，返回值不会是0。

	fmt.Println(Pointer_v2.Pointer()) //这个值是指针的地址值，不是切片的地址值
	fmt.Println(Pointer_v2.Kind())
	fmt.Println(Pointer_v2.Elem().Kind())

	fmt.Println(Pointer_v3.Pointer()) //这个值应该和上面的切片地址值相同
	fmt.Println(Pointer_v3.Kind())
	//fmt.Println(Pointer_v3.Elem().Kind())//报错，Pointer_v3必须是指针等

	ls_b91 := []byte{}
	ls_b92 := make([]byte, 2)
	var ls_b93 []byte //未初始化(未声明或者开辟内存)
	Pointer_v4 := reflect.ValueOf(ls_b91)
	Pointer_v5 := reflect.ValueOf(ls_b92)
	Pointer_v6 := reflect.ValueOf(ls_b93)

	fmt.Println(Pointer_v4.Pointer())
	fmt.Println(Pointer_v5.Pointer())
	fmt.Println(Pointer_v6.Pointer())

	ls_b78 := func() {}
	ls_b79 := func() {} //和上面一个函数完全相同
	ls_b80 := func(string, int) {}
	ls_b81 := func(string, bool) {}
	ls_b82 := func(string, bool) string { return "" }
	ls_b83 := func(string, bool) int { return 0 }
	ls_b84 := func(string, bool) int { return 0 } //和上面一个函数完全相同

	Pointer_v68 := reflect.ValueOf(ls_b78)
	Pointer_v69 := reflect.ValueOf(ls_b79)
	Pointer_v7 := reflect.ValueOf(ls_b80)
	Pointer_v8 := reflect.ValueOf(ls_b81)
	Pointer_v9 := reflect.ValueOf(ls_b82)
	Pointer_v10 := reflect.ValueOf(ls_b83)
	Pointer_v11 := reflect.ValueOf(ls_b84)

	fmt.Println(Pointer_v68.Pointer())
	fmt.Println(Pointer_v69.Pointer())
	fmt.Println(Pointer_v7.Pointer())
	fmt.Println(Pointer_v8.Pointer())
	fmt.Println(Pointer_v9.Pointer())
	fmt.Println(Pointer_v10.Pointer())
	fmt.Println(Pointer_v11.Pointer())
	//输出：
	//	字节切片的10进制地址值为： 824633764432
	//	824633740224
	//	ptr
	//	slice
	//	824633764432
	//	slice
	//	6518296
	//	824633764496
	//	0
	//	5107984
	//	5108000
	//	5108016
	//	5108032
	//	5108048
	//	5108064
	//	5108080
	//(以上结果不一定跟我的完全相同)从上面可以知道即使是完全相同的函数，但是如果是不同的内存地址的话也会显示不同的地址值，
	// 说明了通过判断内存地址来判断2个函数相同与否是行不通的！

	com1 := complex64(13 + 2i)
	com2 := complex128(13 + 2i)
	V_complex64 := reflect.ValueOf(com1)
	V_complex128 := reflect.ValueOf(com2)
	// Complex returns v's underlying value, as a complex128.
	// It panics if v's Kind is not Complex64 or Complex128
	fmt.Println("=========", V_complex64.Complex())
	fmt.Println("=========", V_complex128.Complex())
	//输出：
	//========= (13+2i)
	//========= (13+2i)

	fmt.Println()
	// Type returns v's type.底层是返回v.typ
	fmt.Println("=========", V_complex64.Type())
	fmt.Println("=========", V_complex128.Type())
	//输出：
	//========= complex64
	//========= complex128

	////下面先说下关于Value这个结构体的相关信息：
	//// Value is the reflection interface to a Go value.
	////
	//// Not all methods apply to all kinds of values. Restrictions,
	//// if any, are noted in the documentation for each method.
	//// Use the Kind method to find out the kind of value before
	//// calling kind-specific methods. Calling a method
	//// inappropriate to the kind of type causes a run time panic.
	////
	//// The zero Value represents no value.
	//// Its IsValid method returns false, its Kind method returns Invalid,
	//// its String method returns "<invalid Value>", and all other methods panic.
	//// Most functions and methods never return an invalid value.
	//// If one does, its documentation states the conditions explicitly.
	////
	//// A Value can be used concurrently by multiple goroutines provided that
	//// the underlying Go value can be used concurrently for the equivalent
	//// direct operations.
	////
	//// To compare two Values, compare the results of the Interface method.
	//// Using == on two Values does not compare the underlying values
	//// they represent.
	////Value为go值提供了反射接口。
	////
	////不是所有go类型值的Value表示都能使用所有方法。请参见每个方法的文档获取使用限制。在调用有分类限定的方法时，应先使用Kind方法获知该值的分类。调用该分类不支持的方法会导致运行时的panic。
	////
	////Value类型的零值表示不持有某个值。零值的IsValid方法返回false，其Kind方法返回Invalid，而String方法返回"<invalid Value>"，所有其它方法都会panic。绝大多数函数和方法都永远不返回Value零值。如果某个函数/方法返回了非法的Value，它的文档必须显式的说明具体情况。
	////
	////如果某个go类型值可以安全的用于多线程并发操作，它的Value表示也可以安全的用于并发。
	////要比较2个值的话，可以通过比较他们接口方法的返回值来比较他们。在2个值上面使用==号不会比较他们代表的底层值
	//
	//type Value struct {
	//	// typ holds the type of the value represented by a Value.(typ存储着Value对象的值的类型信息)
	//	typ *rtype
	//
	//	// Pointer-valued data or, if flagIndir is set, pointer to data.
	//	// Valid when either flagIndir is set or typ.pointers() is true.
	//	ptr unsafe.Pointer
	//
	//	// flag holds metadata about the value.
	//	// The lowest bits are flag bits:
	//	//	- flagStickyRO: obtained via unexported not embedded field, so read-only
	//	//	- flagEmbedRO: obtained via unexported embedded field, so read-only
	//	//	- flagIndir: val holds a pointer to the data
	//	//	- flagAddr: v.CanAddr is true (implies flagIndir)
	//	//	- flagMethod: v is a method value.
	//	// The next five bits give the Kind of the value.
	//	// This repeats typ.Kind() except for method values.
	//	// The remaining 23+ bits give a method number for method values.
	//	// If flag.kind() != Func, code can assume that flagMethod is unset.
	//	// If ifaceIndir(typ), code can assume that flagIndir is set.
	//	flag
	//
	//	// A method value represents a curried method invocation
	//	// like r.Read for some receiver r. The typ+val+flag bits describe
	//	// the receiver r, but the flag's Kind bits say Func (methods are
	//	// functions), and the top bits of the flag give the method number
	//	// in r's type's method table.
	//}
	//
	//type flag uintptr

	//下面是rtype结构体的信息：
	//// rtype is the common implementation of most values.（rtype是大多数类型的值的公共实现（多数类型的值中内嵌的字段中基本都有rtype字段））
	//// It is embedded in other struct types.（他在其他结构体中是匿名字段）
	////
	//// rtype must be kept in sync with ../runtime/type.go:/^type._type.（rtype字段同步通过../runtime/type.go:/^type._type来实现）
	//type rtype struct {
	//	size       uintptr	//类型的大小
	//	ptrdata    uintptr  // number of bytes in the type that can contain pointers（指针的字节数）
	//	hash       uint32   // hash of type; avoids computation in hash tables（类型的哈希值）
	//	tflag      tflag    // extra type information flags（额外的类型信息标志）
	//	align      uint8    // alignment of variable with this type（类型中有效的对齐字节长度）
	//	fieldAlign uint8    // alignment of struct field with this type（类型中结构字段的有效的对齐字节长度）
	//	kind       uint8    // enumeration for C（c枚举类型）
	//	alg        *typeAlg // algorithm table	（）
	//	gcdata     *byte    // garbage collection data
	//	str        nameOff  // string form（类型名的字符串形式在内存序列中的索引）
	//	ptrToThis  typeOff  // type for pointer to this type, may be zero（类型指针值，可能是零）
	//}

	fmt.Println()
	var str1 = "hello"
	var str2 = "world"
	var p001 = person{
		Name:      "p001",
		Age:       10,
		isstudent: false,
	}
	var p002 = person{
		Name:      "p002",
		Age:       20,
		isstudent: true,
	}

	type PP struct {
		Name string
		age  int
	}
	var pp = PP{"anko", 13}

	a11 := []string{str1, str2}
	a12 := []person{p001, p002}

	Pointer_a10 := reflect.ValueOf(ls_b9[0])
	Pointer_a11 := reflect.ValueOf(a11[0])
	Pointer_a12 := reflect.ValueOf(&(a12[1].Name))
	Pointer_a13 := reflect.ValueOf(pp.Name)
	// Addr returns a pointer value representing the address of v.
	// It panics if CanAddr() returns false.
	// Addr is typically used to obtain a pointer to a struct field
	// or slice element in order to call a method that requires a
	// pointer receiver.
	//函数返回一个持有指向v持有者的指针的Value封装。如果v.CanAddr()返回false，
	//调用本方法会panic。Addr一般用于获取结构体字段的指针或者切片的元素（的Value封装）以便调用需要指针类型接收者的方法。

	// CanAddr reports whether the value's address can be obtained with Addr.
	// Such values are called addressable. A value is addressable if it is
	// an element of a slice, an element of an addressable array,
	// a field of an addressable struct, or the result of dereferencing a pointer.
	// If CanAddr returns false, calling Addr will panic.
	//返回是否可以获取v持有值的指针。可以获取指针的值被称为可寻址的。如果一个值是切片或可寻址数组的元素、
	//可寻址结构体的字段、或从指针解引用得到的，该值即为可寻址的。

	fmt.Println(Pointer_a10.CanAddr())
	fmt.Println(Pointer_a11.CanAddr())
	fmt.Println(Pointer_a12.CanAddr())
	fmt.Println(Pointer_a13.CanAddr())
	//fmt.Println(Pointer_a10.Addr())
	//fmt.Println(Pointer_a11.Addr())
	//fmt.Println(Pointer_a12.Addr())
	//暂时还不大清楚怎么用

	fmt.Println()
	ls_1 := []byte{'a', 'b', 'c'}
	ls_2 := []byte{'a', 'b', 'c', 'd'}
	arr_1 := [...]byte{'a', 'b', 'c', 'd'}
	chan_1 := make(chan int, 5)
	chan_1 <- 11
	chan_1 <- 22
	chan_1 <- 33
	chan_1 <- 44
	chan_1 <- 55

	V_ls_1 := reflect.ValueOf(ls_1)
	V_ls_2 := reflect.ValueOf(ls_2)
	V_arr_1 := reflect.ValueOf(arr_1)
	V_chan_1 := reflect.ValueOf(chan_1)
	// Cap returns v's capacity.
	// It panics if v's Kind is not Array, Chan, or Slice.
	fmt.Println("=========", V_ls_1.Cap())
	fmt.Println("=========", V_ls_2.Cap())
	fmt.Println("=========", V_arr_1.Cap())
	fmt.Println("=========", V_chan_1.Cap())
	//输出：
	//========= 3
	//========= 4
	//========= 4
	//========= 5

	fmt.Println()
	chan_2 := make(chan int, 2)
	chan_2 <- 10
	fmt.Println(chan_2)
	V_chan_2 := reflect.ValueOf(chan_2)
	// Close closes the channel v.
	// It panics if v's Kind is not Chan.
	V_chan_2.Close()
	//chan_2 <- 11//通道已经关闭，此处会报错
	for ok := range chan_2 {
		fmt.Println("取通道值：", ok)
	} //这里等同于下面的循环
	//for{
	//	v1,ok :=<-chan_2
	//	if !ok{
	//		break
	//	}
	//	fmt.Println("取通道值：",v1)
	//}
	//输出:
	//0xc000018310
	//取通道值： 10

	fmt.Println()
	fun10 := func(s string, b ...byte) string {
		ss := ""
		for _, value := range b {
			ss += string(value) + s
		}
		return ss
	}
	V_fun10 := reflect.ValueOf(fun10)
	// CallSlice calls the variadic function v with the input arguments in,
	// assigning the slice in[len(in)-1] to v's final variadic argument.
	// For example, if len(in) == 3, v.CallSlice(in) represents the Go call v(in[0], in[1], in[2]...).
	// CallSlice panics if v's Kind is not Func or if v is not variadic.
	// It returns the output results as Values.
	// As in Go, each input argument must be assignable to the
	// type of the function's corresponding input parameter.
	//CallSlice调用v持有的可变参数函数，会将切片类型的in[len(in)-1]（的成员）分配给v的最后的可变参数。例如，
	//如果len(in) == 3，v.Call(in)代表调用v(in[0], in[1], in[2])（其中Value值表示其持有值，可变参数函数
	//的可变参数位置提供一个切片并跟三个点号代表"解切片"）。如果v的Kind不是Func或者v的持有值不是可变参数函数，
	//会panic。它返回函数所有输出结果的Value封装的切片。和go代码一样，每一个输入实参的持有值都必须可以直接赋值
	//给函数对应输入参数的类型。
	fmt.Println(fun10("-", 'a', 'b', 'c'))
	ls_v := V_fun10.CallSlice([]reflect.Value{
		reflect.ValueOf("-"),
		reflect.ValueOf([]byte{'a', 'b', 'c'}),
		//每一行中的一个参数代表一个函数的参数，千万不要像下面这样！在go中很多都是这样让你出乎意料的！即使你看上面的中文也许都不会使用！
		//我也是想了很久才反应过来！
		//reflect.ValueOf('a'),
		//reflect.ValueOf('b'),
		//reflect.ValueOf('c'),
	})
	fmt.Println(ls_v)
	//输出：
	//a-b-c-
	//[a-b-c-]

	// Call calls the function v with the input arguments in.
	// For example, if len(in) == 3, v.Call(in) represents the Go call v(in[0], in[1], in[2]).
	// Call panics if v's Kind is not Func.
	// It returns the output results as Values.
	// As in Go, each input argument must be assignable to the
	// type of the function's corresponding input parameter.
	// If v is a variadic function, Call creates the variadic slice parameter
	// itself, copying in the corresponding values.
	//Call方法使用输入的参数in调用v持有的函数。例如，如果len(in) == 3，v.Call(in)代表调用v(in[0], in[1], in[2])
	//（其中Value值表示其持有值）。如果v的Kind不是Func会panic。它返回函数所有输出结果的Value封装的切片。和go代码一样，
	//每一个输入实参的持有值都必须可以直接赋值给函数对应输入参数的类型。如果v持有值是可变参数函数，Call方法会自行创建一个
	//代表可变参数的切片，将对应可变参数的值都拷贝到里面。
	//这个方法其实是上面方法的底层实现,不过这个方法不只是应用于可变参数的函数，几乎全部函数都可以！
	fmt.Println()
	ls_v1 := V_fun10.Call([]reflect.Value{
		reflect.ValueOf("-"),
		//reflect.ValueOf('a'),//如果你像这样写的话默认是当做rune(int32)类型而不是byte类型
		reflect.ValueOf(byte('a')),
		reflect.ValueOf(byte('b')),
		reflect.ValueOf(byte('c')),
	})
	fmt.Println(ls_v1)

	fun11 := func(s string, b rune) string {
		ss := ""
		ss += string(b) + s
		return ss
	}
	V_fun11 := reflect.ValueOf(fun11)

	ls_v3 := V_fun11.Call([]reflect.Value{
		reflect.ValueOf("-"),
		reflect.ValueOf('a'), //如果你像这样写的话默认是当做rune(int32)类型而不是byte类型
	})
	fmt.Println(ls_v3)
	//输出：
	//[a-b-c-]
	//[a-]

	fmt.Println()
	var beh_1 behavior
	beh_1 = &p1
	T_beh_1 := reflect.ValueOf(beh_1)
	// CanInterface reports whether Interface can be used without panicking.
	CanInterface_1 := T_beh_1.CanInterface()
	fmt.Println(CanInterface_1)
	type inter interface {
	}
	//var inter_1 behavior//报错,panic: reflect: call of reflect.Value.CanInterface on zero Value，必须使用指针
	var inter_1 *behavior
	T_inter_1 := reflect.ValueOf(inter_1)
	CanInterface_2 := T_inter_1.CanInterface()
	fmt.Println(T_inter_1)
	fmt.Println(CanInterface_2)
	//输出：
	//true
	//<nil>
	//true
	//不大清楚接口在什么情况下回抛出异常！

	// CanSet reports whether the value of v can be changed.
	// A Value can be changed only if it is addressable and was not
	// obtained by the use of unexported struct fields.
	// If CanSet returns false, calling Set or any type-specific
	// setter (e.g., SetBool, SetInt) will panic.
	//如果v持有的值可以被修改，CanSet就会返回真。只有Value持有值可以被寻址同时又不是来自非导出字段时，它才可以被修改。
	//如果CanSet返回假，调用Set或任何限定类型的设置函数（如SetBool、SetInt64）都会panic。

	fmt.Println()
	b001 := T_inter_1.CanSet()
	b002 := T_beh_1.CanSet()

	var str_001 = "abcdef"
	T_str_001 := reflect.ValueOf(&str_001).Elem() //必须使用Elem()，否则无论什么类型都是不可寻址和不可设置的！
	b003 := T_str_001.CanSet()

	var str_002 = [][]byte{{'a', 'b'}, {'c', 'd'}}
	fmt.Println(str_002)
	T_str_002 := reflect.ValueOf(&str_002).Elem()
	b004 := T_str_002.CanSet()

	type student struct {
		Name string
		age  int
	}
	var Stu = student{"anko", 66}
	//var Stu Student
	T_Stu1 := reflect.ValueOf(&Stu.Name).Elem()
	T_Stu2 := reflect.ValueOf(&Stu.age).Elem()
	T_Stu3 := reflect.ValueOf(&Stu).Elem()
	fmt.Printf("%#v--%v\n", T_Stu1, T_Stu1.Kind())
	fmt.Printf("%#v--%v\n", T_Stu2, T_Stu2.Kind())
	fmt.Printf("%#v--%v\n", T_Stu3, T_Stu3.Kind())

	b005 := T_Stu1.CanSet()
	b006 := T_Stu2.CanSet()
	b007 := T_Stu3.CanSet()

	fmt.Println(b001)
	fmt.Println(b002)
	fmt.Println(b003)
	fmt.Println(b004)
	fmt.Println(b005)
	fmt.Println(b006)
	fmt.Println(b007)
	//输出：
	//	[[97 98] [99 100]]
	//	"anko"--string
	//	66--int
	//	main.student{Name:"anko", age:66}--struct
	//	false
	//	false
	//	true
	//	true
	//	true
	//	true
	//	true
	//真的不大了解这个到底是怎么算的！！

	fmt.Println()
	//判断是否可以被寻址
	bn1 := T_Stu1.CanAddr()
	vv_1 := T_Stu1.Addr()
	fmt.Println(bn1)
	fmt.Println(vv_1)
	//输出：
	//true
	//0xc000004f00

	fmt.Println()
	fmt.Println("设置之前的值：", T_str_001)
	fmt.Println("设置之前的str_001：", str_001)
	T_str_001.Set(reflect.ValueOf("ABC"))
	fmt.Println("设置之后的值：", T_str_001)
	fmt.Println("设置之后的str_001：", str_001)

	fmt.Println()
	fmt.Println("设置之前的值：", T_Stu1)
	fmt.Println("设置之前的Stu：", Stu)
	T_Stu1.Set(reflect.ValueOf("anko111"))
	fmt.Println("设置之后的值：", T_Stu1)
	fmt.Println("设置之后的Stu：", Stu)

	fmt.Println()
	fmt.Println("设置之前的值：", T_Stu2)
	fmt.Println("设置之前的Stu：", Stu)
	T_Stu2.Set(reflect.ValueOf(33))
	fmt.Println("设置之后的值：", T_Stu2)
	fmt.Println("设置之后的Stu：", Stu)
	//不知道为什么不可导出的字段也可以被设置
	//输出：
	//	设置之前的值： abcdef
	//	设置之前的str_001： abcdef
	//	设置之后的值： ABC
	//	设置之后的str_001： ABC
	//
	//	设置之前的值： anko
	//	设置之前的Stu： {anko 66}
	//	设置之后的值： anko111
	//	设置之后的Stu： {anko111 66}
	//
	//	设置之前的值： 66
	//	设置之前的Stu： {anko111 66}
	//	设置之后的值： 33
	//	设置之后的Stu： {anko111 33}
	//还有很多setXxx系列的api不再展示了，跟上面的差不多的！
	//T_Stu2.SetInt()
	//T_Stu2.SetBool()
	//...(省略)

	//// MapRange returns a range iterator for a map.
	//// It panics if v's Kind is not Map.
	////
	//// Call Next to advance the iterator, and Key/Value to access each entry.
	//// Next returns false when the iterator is exhausted.
	//// MapRange follows the same iteration semantics as a range statement.
	////
	//// Example:
	////
	////	iter := reflect.ValueOf(m).MapRange()
	//// 	for iter.Next() {
	////		k := iter.Key()
	////		v := iter.Value()
	////		...
	////	}
	//
	//// A MapIter is an iterator for ranging over a map.
	//// See Value.MapRange.
	//type MapIter struct {
	//	m  Value	//map的值
	//	it unsafe.Pointer	//map的指针
	//}
	fmt.Println()
	map2 := map[string]int{"a": 1, "b": 2, "c": 3}
	V_map2 := reflect.ValueOf(map2)
	mapRange1 := V_map2.MapRange()
	fmt.Printf("%+v\n", mapRange1)

	//不可以遍历
	//for key, value := range &mapRange1 {
	//	fmt.Println(key)
	//	fmt.Println(value)
	//}
	//虽然不可以通过range遍历，但是可以通过for条件(mapRange1.Next()是一个很好的条件)来循环遍历

	// Next advances the map iterator and reports whether there is another
	// entry. It returns false when the iterator is exhausted; subsequent
	// calls to Key, Value, or Next will panic.(返回是否有下一个)
	fmt.Println(mapRange1.Next())
	// Key returns the key of the iterator's current map entry.
	fmt.Println(mapRange1.Key())
	// Value returns the value of the iterator's current map entry.
	fmt.Println(mapRange1.Value())

	fmt.Println(mapRange1.Next())
	fmt.Println(mapRange1.Key())
	fmt.Println(mapRange1.Value())

	fmt.Println(mapRange1.Next())
	fmt.Println(mapRange1.Key())
	fmt.Println(mapRange1.Value())

	fmt.Println(mapRange1.Next())
	//fmt.Println(mapRange1.Key())//报错，panic: MapIter.Key called on exhausted iterator
	//fmt.Println(mapRange1.Value())//报错，panic: MapIter.Key called on exhausted iterator
	//输出：
	//	&{m:{typ:0x504f40 ptr:0xc00005cb10 flag:21} it:<nil>}
	//	true
	//	b
	//	2
	//	true
	//	c
	//	3
	//	true
	//	a
	//	1
	//	false

	fmt.Println()
	// MapKeys returns a slice containing all the keys present in the map,(MapKeys返回包含所有键（不保证顺序）的切片，)
	// in unspecified order.
	// It panics if v's Kind is not Map.（调用者的kind类型必须是Map类型）
	// It returns an empty slice if v represents a nil map.（如果map是一个nil的初始值的话，那么MapKeys()将返回空切片）
	fmt.Println(V_map2.MapKeys())

	// MapIndex returns the value associated with key in the map v.(MapIndex返回键的值)
	// It panics if v's Kind is not Map.（调用者的kind类型必须是Map类型）
	// It returns the zero Value if key is not found in the map or if v represents a nil map.（如果map是一个nil的初始值或者map不存在这个键的话，那么MapKeys()将返回零值的Value）
	// As in Go, the key's value must be assignable to the map's key type(在go中，键的值必须保证跟键的类型一致)
	fmt.Println(V_map2.MapIndex(reflect.ValueOf("a")))
	fmt.Println(V_map2.MapIndex(reflect.ValueOf("b")))
	fmt.Println(V_map2.MapIndex(reflect.ValueOf("x")))
	// IsValid reports whether v represents a value.(判断Value是否有除零值（无效值）之外的值)
	// It returns false if v is the zero Value.（如果v是一个零值的value的话会返回false）
	// If IsValid returns false, all other methods except String panic.(如果返回false则这个value除了能调用String()方法之外，其他的方法都不能调用了，否则会发生panic)
	// Most functions and methods never return an invalid value.(大对数的方法（或者函数）不会返回一个无效的值)
	// If one does, its documentation states the conditions explicitly.（如果某个函数这么做了，那么他的文档会明确标注的！）
	fmt.Println(V_map2.MapIndex(reflect.ValueOf("x")).IsValid())
	fmt.Println(V_map2.MapIndex(reflect.ValueOf("x")).String())

	// IsZero reports whether v is the zero value for its type.（IsZero报告v是否是类型的零值（注意不是判断value的零值））
	// It panics if the argument is invalid.（如果这个值是value的零值（或者叫做无效值，无值）的话，那么他会抛出panic）
	// 其实这个判断的类型就比下面的IsNil（）方法的范围要大些，他不一定需要判断引用类型，可以判断任意的类型的零值！总之区别就是前者是判零值，后者也是判nil值
	// （但是其实也是引用类型的零值，或者叫未声明value的内存时候的初始值）
	//fmt.Println(V_map2.MapIndex(reflect.ValueOf("x")).IsZero())//panic: reflect: call of reflect.Value.IsZero on zero Value
	fmt.Println("是否是切片类型的零值？", reflect.ValueOf(([]byte)(nil)).IsZero())
	fmt.Println("是否是指针类型的零值？", reflect.ValueOf((*[]byte)(nil)).IsZero())
	var zero_int = 0
	var zero_string = ""
	var zero_array [2]byte
	var zero_slice []byte
	var zero_chan chan int
	var zero_InChan chan<- int
	var zero_OutChan <-chan int
	var zero_func func()
	var zero_struct struct{}
	var zero_interface behavior
	//千万不要采用这个类型的value来获取子元素类型的value，这个当做是整体来看的 ！一个指针类型的子元素是nil,
	//一个值类型之后获取指针再求子元素才可以正确的获取到该类型的value,下面会有展示！
	var zero_interface1 *behavior
	var zero_map map[string]int
	var zero_bool bool
	var zero_float64 float64
	var zero_complex64 complex64
	var zero_rune rune
	var zero_uint8 uint8
	var zero_uintptr uintptr
	var zero_person person
	fmt.Println("是否是int类型的零值？", reflect.ValueOf(zero_int).IsZero())
	fmt.Println("是否是string类型的零值？", reflect.ValueOf(zero_string).IsZero())
	fmt.Println("是否是array类型的零值？", reflect.ValueOf(zero_array).IsZero())
	fmt.Println("是否是slice类型的零值？", reflect.ValueOf(zero_slice).IsZero())
	fmt.Println("是否是chan类型的零值？", reflect.ValueOf(zero_chan).IsZero())
	fmt.Println("是否是chan<- int类型的零值？", reflect.ValueOf(zero_InChan).IsZero())
	fmt.Println("是否是<-chan int类型的零值？", reflect.ValueOf(zero_OutChan).IsZero())
	fmt.Println("是否是func类型的零值？", reflect.ValueOf(zero_func).IsZero())
	fmt.Println("是否是struct类型的零值？", reflect.ValueOf(zero_struct).IsZero())

	fmt.Println("---------")
	//panic: reflect: call of reflect.Value.IsZero on zero Value
	//fmt.Println("是否是interface类型的零值？",reflect.ValueOf(zero_interface).IsZero())
	//我们一定要明确调用者是谁，像上面的调用者是一个接口的指针，但是接口是没有实例内存存在的！
	fmt.Printf("reflect.ValueOf(&zero_interface)类型？%v\n", reflect.ValueOf(&zero_interface).Kind())
	fmt.Printf("reflect.ValueOf(&zero_interface).Elem()类型？%v\n", reflect.ValueOf(&zero_interface).Elem().Kind())
	fmt.Println("reflect.ValueOf(&zero_interface).Elem()的值是？", reflect.ValueOf(&zero_interface).Elem())
	fmt.Println("是否是interface类型的零值？", reflect.ValueOf(&zero_interface).Elem().IsZero())

	fmt.Println("==========")
	fmt.Printf("reflect.ValueOf(zero_interface1)类型？%v\n", reflect.ValueOf(zero_interface1).Kind())
	fmt.Printf("reflect.ValueOf(zero_interface1).Elem()类型？%v\n", reflect.ValueOf(zero_interface1).Elem().Kind())
	fmt.Println("reflect.ValueOf(zero_interface1).Elem()的值是？", reflect.ValueOf(zero_interface1).Elem())
	//fmt.Println("是否是interface类型的零值？",reflect.ValueOf(zero_interface1).Elem().IsZero())//会报错panic: reflect: call of reflect.Value.IsZero on zero Value
	fmt.Println("---------")

	fmt.Println("是否是map类型的零值？", reflect.ValueOf(zero_map).IsZero())
	fmt.Println("是否是bool类型的零值？", reflect.ValueOf(zero_bool).IsZero())
	fmt.Println("是否是float64类型的零值？", reflect.ValueOf(zero_float64).IsZero())
	fmt.Println("是否是complex64类型的零值？", reflect.ValueOf(zero_complex64).IsZero())
	fmt.Println("是否是rune类型的零值？", reflect.ValueOf(zero_rune).IsZero())
	fmt.Println("是否是uint8类型的零值？", reflect.ValueOf(zero_uint8).IsZero())
	fmt.Println("是否是uintptr类型的零值？", reflect.ValueOf(zero_uintptr).IsZero())
	fmt.Println("是否是自定义person类型的零值？", reflect.ValueOf(zero_person).IsZero())

	// IsNil reports whether its argument v is nil. The argument must be
	// a chan, func, interface, map, pointer, or slice value; if it is
	// not, IsNil panics. Note that IsNil is not always equivalent to a
	// regular comparison with nil in Go. For example, if v was created
	// by calling ValueOf with an uninitialized interface variable i,
	// i==nil will be true but v.IsNil will panic as v will be the zero
	// Value.
	// IsNil返回是否是对象的类型的零值，这个对象必须是引用类型：chan, func, interface, map, pointer, or slice 。
	// 如果不是的话，会抛出 异常。注意IsNil并不总是等价于go语言中值与参数nil的常规比较。例如：如果v是通过使用某个值为nil的接口调用ValueOf函数创建的，
	// v.IsNil()返回真，但是如果v是Value零值(注意和零值的value区别，Value零值指的是类型的零值)，会panic。也就是nil必须带有类型！不能是参数(不带类型)的nil
	//fmt.Println(V_map2.MapIndex(reflect.ValueOf("x")).IsNil())//panic: reflect: call of reflect.Value.IsZero on zero Value
	fmt.Println("是否是切片类型的零值？", reflect.ValueOf((*[]byte)(nil)).IsNil())
	fmt.Println("是否是指针类型的零值？", reflect.ValueOf(([]byte)(nil)).IsNil())
	//fmt.Println(V_map2.MapIndex(reflect.ValueOf(nil)).IsNil())//报错panic: runtime error: invalid memory address or nil pointer dereference
	//输出：
	//	[b c a]
	//	1
	//	2
	//	<invalid reflect.Value>
	//	false
	//	<invalid Value>
	//	是否是切片类型的零值？ true
	//	是否是指针类型的零值？ true
	//	是否是int类型的零值？ true
	//	是否是string类型的零值？ true
	//	是否是array类型的零值？ true
	//	是否是slice类型的零值？ true
	//	是否是chan类型的零值？ true
	//	是否是chan<- int类型的零值？ true
	//	是否是<-chan int类型的零值？ true
	//	是否是func类型的零值？ true
	//	是否是struct类型的零值？ true
	//	---------
	//	reflect.ValueOf(&zero_interface)类型？ptr
	//	reflect.ValueOf(&zero_interface).Elem()类型？interface
	//	reflect.ValueOf(&zero_interface).Elem()的值是？ <nil>
	//	是否是interface类型的零值？ true
	//	==========
	//	reflect.ValueOf(zero_interface1)类型？ptr
	//	reflect.ValueOf(zero_interface1).Elem()类型？invalid
	//	reflect.ValueOf(zero_interface1).Elem()的值是？ <invalid reflect.Value>
	//	---------
	//	是否是map类型的零值？ true
	//	是否是bool类型的零值？ true
	//	是否是float64类型的零值？ true
	//	是否是complex64类型的零值？ true
	//	是否是rune类型的零值？ true
	//	是否是uint8类型的零值？ true
	//	是否是uintptr类型的零值？ true
	//	是否是自定义person类型的零值？ true
	//	是否是切片类型的零值？ true
	//	是否是指针类型的零值？ true

	fmt.Println()
	map3 := map[string][]byte{"a": {'1', '2', '3'},
		"b": {'4', '5', '6'},
		"c": {'7', '8', '9'}}
	V_map3 := reflect.ValueOf(map3)
	fmt.Println("设置之前的map值是：", map3)
	// SetMapIndex sets the element associated with key in the map v to elem.(新设置map中的某个键的值)
	// It panics if v's Kind is not Map.（如果不是map会抛出异常）
	// If elem is the zero Value, SetMapIndex deletes the key from the map.（如果给的参数中的elem是一个零值的value的话，SetMapIndex方法会从map中删除该键）
	// Otherwise if v holds a nil map, SetMapIndex will panic.（另外如果map的值是一个零值的话，会抛出异常）
	// As in Go, key's elem must be assignable to the map's key type,
	// and elem's value must be assignable to the map's elem type.
	V_map3.SetMapIndex(reflect.ValueOf("a"), reflect.ValueOf([]byte{'a', 'b', 'c'}))
	V_map3.SetMapIndex(reflect.ValueOf("b"), reflect.ValueOf([]byte{'d', 'e', 'f'}))
	V_map3.SetMapIndex(reflect.ValueOf("c"), reflect.ValueOf(nil))
	fmt.Println("第一次设置之后的map值是：", map3)
	V_map3.SetMapIndex(reflect.ValueOf("b"), reflect.ValueOf([]byte(nil)))
	fmt.Println("再设置之后的map值是：", map3)
	//输出：
	//设置之前的map值是： map[a:[49 50 51] b:[52 53 54] c:[55 56 57]]
	//第一次设置之后的map值是： map[a:[97 98 99] b:[100 101 102]]
	//再设置之后的map值是： map[a:[97 98 99] b:[]]
	//注意区分不同点

	fmt.Println()
	// Slice returns v[i:j].(返回索引[i,j]之间的序列（包含切片，数组，字符串）的元素)
	// It panics if v's Kind is not Array, Slice or String, or if v is an unaddressable array,
	// or if the indexes are out of bounds.（如果序列类型的不是切片，数组或者字符串,或者是不可寻址的数组，抑或者给的索引超过了序列的长度限制的话都会抛出panic）
	slice_by := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	slice_arr := [...]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	//slice_arr:=&([...]byte{'a','b','c'})
	slice_str := "abcdefg"

	T_slice_by := reflect.ValueOf(slice_by)
	//对于引用类型我们也可以采用下面的传递指针的形式来进行调用
	//T_slice_by:=reflect.ValueOf(&slice_by).Elem()

	//数组是值类型，通过值找不到对应的地址，必须要先获取指针然后获取子元素的形式来获取到对应的数组value,当然我们也可以slice_arr:=&([3]byte{'a','b','c'})在这里获取指针
	//而没有通过&slice_arr获取指针！当然这2种方式都是可以的！其实有很多种写法,但是至于采用哪种写法真的是取决于个人！
	//不过需要注意的是，上面说的仅仅是针对于值类型才是这样的！如果是引用类型的话，那就不同了~
	//T_slice_arr:=reflect.ValueOf(slice_arr).Elem()
	T_slice_arr := reflect.ValueOf(&slice_arr).Elem()

	T_slice_str := reflect.ValueOf(slice_str)
	//对于字符串，是可以通过值来获取到对应的地址的，因为字符串的传递其实是按照指针来进行传递然后共享的的！所以字符串实质上是一个引用类型！如下也是合法的！
	//T_slice_str:=reflect.ValueOf(&slice_str).Elem()
	//上面的名字原本应该是V_slice_xxx,但是由于T_slice_xxx已经这样起名字，懒得改了

	fmt.Println(T_slice_by.Slice(1, 3))
	fmt.Println(T_slice_arr.Slice(1, 3))
	fmt.Println(T_slice_str.Slice(1, 3))

	// Slice3 is the 3-index form of the slice operation: it returns v[i:j:k].
	// It panics if v's Kind is not Array or Slice, or if v is an unaddressable array,
	// or if the indexes are out of bounds.
	//其实这个方法跟上面的方法差不多的，只是多了一个设置cap容量的参数而已,同时这个api不能用于字符串

	var new_by []byte
	V_new_by := reflect.ValueOf(&new_by).Elem()
	fmt.Println("V_new_by是否允许被设置：", V_new_by.CanSet())
	fmt.Println("设置之前:", new_by, "\tnew_by切片的长度：", V_new_by.Len(), "\tnew_by切片的容量：", V_new_by.Cap())

	//对于T_slice_by那是没有影响的！第三个参数我们一般赋值给新的额切片时候才需要设置，说明了这个api都是在需要赋值给新的切片时候才需要使用到，
	//否则的话都应该使用上一个更加通用的api（Slice（）方法）
	//	1代表元素的start索引,4代表切片切到索引4但不包含索引4在内的元素作为返回切片的元素,同时可以通过这个4-start来确定了返回切片的长度len，
	//  6代表的是底层数组的索引6作为new_by切片的end索引，同时可以通过这个end-start来确定了返回切片的容量
	V_ret := T_slice_by.Slice3(1, 4, 6)
	V_new_by.Set(V_ret) //V_new_by必须是指针或者接口类型，参数V_ret的话则不限制

	fmt.Println()
	fmt.Println("设置之后:", new_by, "\tnew_by切片的长度：", V_new_by.Len(), "\tnew_by切片的容量：", V_new_by.Cap())
	//输出：
	//	[98 99]
	//	[98 99]
	//	bc
	//	V_new_by是否允许被设置： true
	//	设置之前: [] 	new_by切片的长度： 0 	new_by切片的容量： 0
	//
	//	设置之后: [98 99 100] 	new_by切片的长度： 3 	new_by切片的容量： 5

	fmt.Println()
	// OverflowInt reports whether the int64 x cannot be represented by v's type.
	// It panics if v's Kind is not Int, Int8, Int16, Int32, or Int64.
	//如果v持有值的类型不能无溢出的表示x，会返回真。如果v的Kind不是Int、Int8、Int16、Int32、Int64会panic
	//Int64是所有int系列类型中允许申请最大的内存大小了（但是他的实例一定会申请那么大，但是真正被用到的有效位是远远没这么多的，除非你的数字特别大），
	// 但是如果采用Int、Int8、Int16、Int32、这几个类型的实例来装Int64的有效位话则不一定能够装的下，所以才要进行判断！
	//到目前来说还不知道到底用在哪里！
	fmt.Println(reflect.ValueOf(int(200000)).OverflowInt(int64(1888888888888880000)))
	fmt.Println(reflect.ValueOf(int8(20)).OverflowInt(int64(1888888888888880000)))
	fmt.Println(reflect.ValueOf(int16(200)).OverflowInt(int64(1888888888888880000)))
	fmt.Println(reflect.ValueOf(int32(20000)).OverflowInt(int64(1888888888888880000)))
	fmt.Println(reflect.ValueOf(int64(2000000000000000000)).OverflowInt(int64(6888888888888880000)))

	type myint int64
	var mi myint
	fmt.Println("====", reflect.ValueOf(mi).Kind())
	fmt.Println("====", reflect.ValueOf(mi).OverflowInt(int64(6888888888888880000)))

	id2 := int64(2)
	fmt.Println(reflect.ValueOf(&id2).Elem().OverflowInt(int64(6888888888888880000)))
	//输出：
	//	false
	//	true
	//	true
	//	true
	//	false
	//	==== int64
	//	==== false
	//	false
	//	除了上面的OverflowInt之外还有:(不再累叙，用法几乎一样的！)
	//		.OverflowComplex()
	//		.OverflowFloat()
	//		.OverflowUint()

	fmt.Println()
	// SetCap sets v's capacity to n.(设置kind类型为Slice(只有切片。数组和字符串不是！)的序列对象实例的容量为n)
	// It panics if v's Kind is not Slice or if n is smaller than the length or
	// greater than the capacity of the slice.（如果调用者v的kind类型不是slice 或者 n<len(v) 或者 n>cap(v)这些情况下都会抛出异常）
	ls_88 := make([]byte, 5, 10)
	fmt.Printf("设置之前的切片为：%v,地址为：%p,长度为：%v,容量为：%v\n", ls_88, ls_88, len(ls_88), cap(ls_88))
	//如果像下面这样的话，我们获取到的是指针对象的值，然后进行set是设置指针的值，但是我们希望设置的是切片ls_88的值，很明显我们应该获取指针对象指向
	// 的对象上的值（这种关系，我们称之为获取指针对象的子元素）
	//V_ls_88:=reflect.ValueOf(&ls_88)
	V_ls_88 := reflect.ValueOf(&ls_88).Elem()
	//只要是涉及setXxx相关的东西的调用者都必须拿到对象的地址才可以进行，也就是必须是可寻址的对象！当然，
	// 可寻址的对象一定是可设置的对象！这是毋庸置疑的！
	//注意我们不能像下面这样获取Value值的指针，我们要获取的是原本对象上面的指针，也就是上面的ls_88的指针，为什么呢？
	//因为我们要设置的目标是ls_88对象，而不是Value对象！
	//(&V_ls_88).SetCap(8)
	V_ls_88.SetCap(8)
	fmt.Printf("设置之后的切片为：%v,地址为：%p,长度为：%v,容量为：%v\n", ls_88, ls_88, len(ls_88), cap(ls_88))
	//V_ls_88.SetCap(4)//不能小于len(也就是5)，panic: reflect: slice capacity out of range in SetCap
	V_ls_88.SetCap(7) //不能小于len(也就是5)，panic: reflect: slice capacity out of range in SetCap
	fmt.Printf("再次设置之后的切片为：%v,地址为：%p,长度为：%v,容量为：%v\n", ls_88, ls_88, len(ls_88), cap(ls_88))
	//V_ls_88.SetCap(8)//不能大于当前的cap,panic: reflect: slice capacity out of range in SetCap
	//输出：
	//	设置之前的切片为：[0 0 0 0 0],地址为：0xc00000acf0,长度为：5,容量为：10
	//	设置之后的切片为：[0 0 0 0 0],地址为：0xc00000acf0,长度为：5,容量为：8
	//	再次设置之后的切片为：[0 0 0 0 0],地址为：0xc00000acf0,长度为：5,容量为：7

	fmt.Println()
	// SetLen sets v's length to n.（设置切片的值v的长度为n）
	// It panics if v's Kind is not Slice or if n is negative or
	// greater than the capacity of the slice.（如果调用者不是切片或者n是负数，或者大于cap的话则会抛出panic）
	V_ls_88.SetLen(6)
	fmt.Printf("再次设置之后的切片为：%v,地址为：%p,长度为：%v,容量为：%v\n", ls_88, ls_88, len(ls_88), cap(ls_88))
	V_ls_88.SetLen(0)
	fmt.Printf("再次设置之后的切片为：%v,地址为：%p,长度为：%v,容量为：%v\n", ls_88, ls_88, len(ls_88), cap(ls_88))
	//V_ls_88.SetLen(-1)//panic: reflect: slice length out of range in SetLen
	//V_ls_88.SetLen(8)//panic: reflect: slice length out of range in SetLen
	//再次设置之后的切片为：[0 0 0 0 0 0],地址为：0xc000068cb0,长度为：6,容量为：7
	//再次设置之后的切片为：[],地址为：0xc000068cb0,长度为：0,容量为：7

	fmt.Println()
	// UnsafeAddr returns a pointer to v's data.（UnsafeAddr返回指针的值）
	// It is for advanced clients that also import the "unsafe" package.
	// It panics if v is not addressable.（如果是不可寻址的对象的话会抛出异常）

	//返回值uintptr的说明：
	// uintptr is an integer type that is large enough to hold the bit pattern of
	// any pointer.
	//type uintptr uintptr

	//方式1：---------------
	var slice *[]byte
	slice = &[]byte{'a', 'b', 'c'}
	V_slice := reflect.ValueOf(slice).Elem()
	//---------------

	//方式2：---------------
	//var slice =[]byte{'a','b','c'}
	//V_slice :=reflect.ValueOf(&slice).Elem()
	//---------------
	fmt.Printf("%p---%p\n", &V_slice, &V_slice)
	fmt.Printf("%p---%p\n", slice, &slice)
	fmt.Printf("%#x\n", V_slice.UnsafeAddr())

	//方式1的输出：（此时V_slice是一个切片，所以UnsafeAddr()相当于取切片的指针）
	//	0xc000005760---0xc000005760
	//	0xc000005740---0xc000006098
	//	0xc000005740（对比这个值跟上面2行的值的迥异）

	//方式2的输出：（此时V_slice是一个指针,所以UnsafeAddr()相当于取切片指针的指针）
	//	0xc000005760---0xc000005760
	//	0xc00000ad40---0xc000005740
	//	0xc000005740（对比这个值跟上面2行的值的迥异）

	fmt.Println()
	// TrySend attempts to send x on the channel v but will not block.(TrySend尝试往通道里面发送一个值进去,但是无缓存的通道发送后也不会阻塞且发送进去的值在执行完这个函数后会被立马作废)
	// It panics if v's Kind is not Chan.(调用者如果不是Chan类型会抛出panic)
	// It reports whether the value was sent.（返回值是bool：参数也就是Value对象是否被发送进去通道了）
	// As in Go, x's value must be assignable to the channel's element type.(在go中，发送进去通道的值必须和通道规定的子元素的类型符合，否则会抛出异常)
	chan111 := make(chan int)
	//chan111:=make(chan int,2)
	go func() {
		fmt.Println("新g程执行中")
		for {
			v, ok := <-chan111
			if !ok {
				break
			}
			fmt.Println("新g程取到一个值为：", v)
		}
		fmt.Println("新g程执行完毕")
	}()
	V_chan111 := reflect.ValueOf(chan111)
	time.Sleep(2e9) //这里必须睡眠，好让g程准备取值，不然让下面的TrySend（）一旦执行完成的话，
	// 发送进去通道的值也会立马作废，因此我们最好让g程时刻准备好取值，
	V_chan111.TrySend(reflect.ValueOf(22))
	fmt.Println("第一个值已经发送进去无缓存通道了而且没阻塞")
	time.Sleep(2e9)
	V_chan111.TrySend(reflect.ValueOf(23))
	fmt.Println("第一个值已经发送进去无缓存通道了而且没阻塞")
	time.Sleep(2e9)
	//chan111如果是无缓存通道且没有其他对象取通道的值的话，下面的操作会导致死锁，因为根本没值取，
	// 在TrySend（）执行完后如果没人取值的话会立马作废，他不会等待直到有人来取值之后才解阻塞！
	//那么为什么我还要写下面的2句代码呢？因为对于有缓存的通道是没阻塞的，塞进去的值也不会作废！，所以下面的2句代码是给有缓存的通道取值用的！
	//当然你也可以通过上面的g程来进行取值，但是其实我上面的g程以及睡眠都是为了无缓存的通道取值用的，下面的代码才是为有缓存的通道取值用的！
	//所以如果你测试的是无缓存的额通道的话，那么请将下面的2句话注释掉即可
	//fmt.Println(<-chan111)
	//fmt.Println(<-chan111)
	//还有一个需要注意的是因为我们测试的发送值，所以我们的通道类型不能是只能取值的通道，我们必须是存取值的通道或者是只能存值的通道
	//输出：
	//	新g程执行中
	//	第一个值已经发送进去无缓存通道了而且没阻塞
	//	新g程取到一个值为： 22
	//	第一个值已经发送进去无缓存通道了而且没阻塞
	//	新g程取到一个值为： 23

	fmt.Println()
	// TryRecv attempts to receive a value from the channel v but will not block.
	// It panics if v's Kind is not Chan.
	// If the receive delivers a value, x is the transferred value and ok is true.
	// If the receive cannot finish without blocking, x is the zero Value and ok is false.
	// If the channel is closed, x is the zero value for the channel's element type and ok is false.
	//TryRecv尝试从v持有的通道接收一个值，但不会阻塞。如果v的Kind不是Chan会panic。如果方法成功接收到一个值，会返回该值（的Value封装）和true；
	//如果不能无阻塞的接收到值，返回Value零值和false；如果因为通道关闭而返回，返回值x是持有通道元素类型的零值的Value和false。
	x, ok := V_chan111.TryRecv()
	fmt.Printf("取到通道中的一个值了？%v,该值为：%v\n", ok, x)
	if !ok {
		fmt.Println("此时娶不到值，但是没阻塞直接有值取才返回！而是立马返回了！")
	}
	//下面跟上面没什么特别的联系
	//为了让他可以取到值，我们创建一个有缓存的通道来让他取值，但是这个跟我们平常没什么区别！如果我们要取到值的话可以循环他！
	// 我们设置一个循环取值和存值，但是即使取不到值却没阻塞下面的代码！
	chan222 := make(chan int, 15)
	V_chan222 := reflect.ValueOf(chan222)

	var num = 0
	go func() {
		var sum,n int
		for sum<70{
			n=0
			x, ok = V_chan222.TryRecv() //不阻塞
			fmt.Printf("取到通道中的一个值了？%v,该值为：%v\n", ok, x)
			if !ok {
				fmt.Println("此时取不到值，但是没阻塞直到有值取才解阻塞！而是立马解阻塞！")
			}
			n++
			sum+=num+n
		}
		fmt.Println("准备退出g程。。。")
	}()

	for num < 7 {
		fmt.Printf("第%v次循环存值\n", num+1)
		//for range chan222 {}//这样写的会死锁！没有接收值的对象
		V_chan222.TrySend(reflect.ValueOf(num))
		num++
	}
	time.Sleep(5e9)
	//输出：
	//	取到通道中的一个值了？false,该值为：<invalid reflect.Value>
	//	此时娶不到值，但是没阻塞直接有值取才返回！而是立马返回了！
	//	第1次循环存值
	//	第2次循环存值
	//	第3次循环存值
	//	第4次循环存值
	//	第5次循环存值
	//	第6次循环存值
	//	第7次循环存值
	//	取到通道中的一个值了？true,该值为：0
	//	取到通道中的一个值了？true,该值为：1
	//	取到通道中的一个值了？true,该值为：2
	//	取到通道中的一个值了？true,该值为：3
	//	取到通道中的一个值了？true,该值为：4
	//	取到通道中的一个值了？true,该值为：5
	//	取到通道中的一个值了？true,该值为：6
	//	取到通道中的一个值了？false,该值为：<invalid reflect.Value>
	//	此时取不到值，但是没阻塞直到有值取才解阻塞！而是立马解阻塞！
	//	取到通道中的一个值了？false,该值为：<invalid reflect.Value>
	//	此时取不到值，但是没阻塞直到有值取才解阻塞！而是立马解阻塞！
	//	准备退出g程。。。

	// Convert returns the value v converted to type t.
	// If the usual Go conversion rules do not allow conversion
	// of the value v to type t, Convert panics.
	//Convert将v持有的值转换为类型为t的值，并返回该值的Value封装。如果go转换规则不支持这种转换，会panic。
	fmt.Println()
	var in32 = int32(32)
	var in64 = int64(32)
	V_in32 := reflect.ValueOf(in32)
	V_in64 := reflect.ValueOf(in64)
	V_ret1 := V_in32.Convert(V_in64.Type())
	fmt.Printf("类型装换后的结果类型是：%v,值是：%v\n", V_ret1.Type(), V_ret1)
	//输出：
	//类型装换后的结果类型是：int64,值是：32

	fmt.Println()
	// InterfaceData returns the interface v's value as a uintptr pair.
	// It panics if v's Kind is not Interface.
	//返回v持有的接口类型值的数据。如果v的Kind不是Interface会panic
	var inter11 behavior
	inter11 = &p1
	V_inter11 := reflect.ValueOf(&inter11).Elem()
	fmt.Println(V_inter11.InterfaceData())
	//下行报错,只能通过接口类型值来调用，panic: reflect: call of reflect.Value.InterfaceData on ptr Value
	//fmt.Println(reflect.ValueOf(&inter11).InterfaceData())
	//输出：
	//[5587584 824633739040]
	//假如不给接口一个实例值即不写上面的inter11=&p1的话会输出：
	//[0,0]
	//目前还不大清楚这[2]uintptr到底是表示什么！

	fmt.Println()
	// Index returns v's i'th element.（返回序列的第i个元素）
	// It panics if v's Kind is not Array, Slice, or String or i is out of range.
	//fmt.Println(V_inter11.Index(0))//报错，panic: reflect: call of reflect.Value.Index on interface Value
	fmt.Println(reflect.ValueOf([]byte{'a', 'b', 'c'}).Index(1))
	fmt.Println(reflect.ValueOf([3]byte{'a', 'b', 'c'}).Index(1))
	fmt.Println(reflect.ValueOf("abcdef").Index(1))
	var sss = "abcdef"
	V_bbb := reflect.ValueOf(sss).Index(1).CanAddr() //Index方法限制了不允许写成&sss，否则编译错误！
	fmt.Println(V_bbb)
	//fmt.Println(reflect.ValueOf("abcdef").Index(10))//超过了索引限制，panic: reflect: string index out of range
	//输出：
	//	98
	//	98
	//	98
	//	false

	fmt.Println()
	var sss0 = "abcdef"
	V_bbb00 := reflect.ValueOf(&sss0).Elem() //可寻址的写法大体上必须是这样写！也必须调用Elem()！
	V_bbb00.SetString("xxx")
	fmt.Println("sss0:", sss0)
	fmt.Println("V_bbb00:", V_bbb00)
	//输出：
	//	sss0: xxx
	//	V_bbb00: xxx

	//下面是错误示例：
	//fmt.Println()
	//var sss1="abcdef"
	//V_bbb11:=reflect.ValueOf(sss1).Index(1)
	//// SetString sets v's underlying value to x.
	//// It panics if v's Kind is not String or if CanSet() is false.(调用者必须是可被设置的或者说是可寻址的对象)
	//V_bbb11.SetString("xxx")
	//fmt.Println("sss1:",sss1)
	//fmt.Println("V_bbb11:",V_bbb11)
	//错误示例输出：
	//panic: reflect: reflect.flag.mustBeAssignable using unaddressable value

	fmt.Println()
	chan_3 := make(chan int, 8)//给足够大的缓存，但是我们并不准备用光它
	V_chan_3 := reflect.ValueOf(chan_3)
	go func() {

		j := 0
		for j < 5 { //循环发送5个值到通道中去

			// Send sends x on the channel v.
			// It panics if v's kind is not Chan or if x's type is not the same type as v's element type.
			// As in Go, x's value must be assignable to the channel's element type.
			//方法向v持有的通道发送x持有的值。如果v的Kind不是Chan，或者x的持有值不能直接赋值给v持有通道的元素类型，会panic。
			V_chan_3.Send(reflect.ValueOf(j))
			fmt.Printf("发送了第%v个值到通道中去了。。。\n", j)
			j++
			if  j==5{
				fmt.Println("发送了5个值，准备关闭通道")
				close(chan_3)//一定要关闭通道，关闭通道时候会发送一个false值到通道中去，下面取值的时候才知道什么时候结束了！
			}
		}

	}()
	time.Sleep(2e9)
	for {
		// Recv receives and returns a value from the channel v.
		// It panics if v's Kind is not Chan.
		// The receive blocks until a value is ready.
		// The boolean value ok is true if the value x corresponds to a send
		// on the channel, false if it is a zero value received because the channel is closed.
		//方法从v持有的通道接收并返回一个值（的Value封装）。如果v的Kind不是Chan会panic。方法会阻塞直到获取到值。
		//如果返回值x对应于某个发送到v持有的通道的值，ok为真；如果因为通道关闭而返回，x为Value零值而ok为假。
		v, ok := V_chan_3.Recv()
		if !ok {
			fmt.Println("通道中没有值了，准备结束接收器。。。")
			break
		}
		fmt.Printf("接收到一个值：%v,kind类型为：%T\n", v, v)
		//time.Sleep(1e9)//我们完全可以在没值的时候沉睡，除非通道是无缓存的，如此的话就不能沉睡了！
	}
	//输出：
	//	发送了第0个值到通道中去了。。。
	//	发送了第1个值到通道中去了。。。
	//	发送了第2个值到通道中去了。。。
	//	发送了第3个值到通道中去了。。。
	//	发送了第4个值到通道中去了。。。
	//	发送了5个值，准备关闭通道
	//	接收到一个值：0,kind类型为：reflect.Value
	//	接收到一个值：1,kind类型为：reflect.Value
	//	接收到一个值：2,kind类型为：reflect.Value
	//	接收到一个值：3,kind类型为：reflect.Value
	//	接收到一个值：4,kind类型为：reflect.Value
	//	通道中没有值了，准备结束接收器。。。




	fmt.Println()
	//好了，几乎讲光了这个反射包的重点了
	
	
	//2019.12.14更新以下内容：
	//主要是为了探究go的可寻址和不可寻址 以及 可设置和不可设置的真正理解
	type Person1 struct {
		name string
		sex uint8
		age uint8
		list [][]string
	}
	testCanAddr := func() {
		p2:=Person1{"luck",0,100,[][]string{{"2"}}}
		p3:=Person1{"luck1",1,200,[][]string{{"22"}}}
		s:=map[int]Person1{}//存的是Person1类型的实例对象的值，copy一份参数值存到map中去，这个参数值既不是内存地址值，也
		// 没任何的变量索引着他，只能通过键来查询到它！但是却不能通过直接的内存地址来查询到它，所以他确实是不可寻址的参数值！好比如下面：
		// &"hello"这样是错误的！"hello"这个参数值确实是存在一个内存地址上面，但是这个内存地址并没有被你用变量来存着，导致go编译器找不到对应的变量来索引你的内存地址！
		// var s="hello";&s这样是正确的，s代表的是一个内存地址！

		//s:=map[int]*Person1{}//按照上面思路可进行推理：第一，*Person1是一个内存地址值么？如果是，那么就是可寻址的！很明显这里是一个内存地址值！
		// 											第二，*Person1有变量索引着他么？很明显这里没有！事实上第二点跟第一点是一样的！变量其实就是内存地址值！编译后运行时候
		// 											不存在变量！只存在随机的内存地址值
		s[1]=p2
		s[2]=p3
		//fmt.Println(reflect.ValueOf(&(s[2])).Elem().CanAddr())//错误，s[2]不是一个内存地址值，而是一个参数
		s2:=s[2]
		fmt.Println(reflect.ValueOf(&(s2)).Elem().CanAddr())//正确，s2是一个内存地址值，不是一个参数
		//s[2].name="sdsdsd"//因为是map[int]Person1，而不是map[int]*Person1类型，所以前者是不可寻址的！也就是不可设置值的！
		//s[2].list=[][]string{{"张三"}}
		fmt.Println(s[2])



		//下面是验证字符串是否可被寻址！其实某个变量的（或者说某个内存地址上的）字符串是可以被更改的！不能更改的是字符串上面的某个元素，因为这些元素没有任何的变量索引着他！
		fmt.Println()
		var p string="hello"
		fmt.Println(p)
		fmt.Println(reflect.ValueOf(&p).Elem().CanAddr())
		fmt.Println(reflect.ValueOf(&p).Elem().Addr())
		p="word"//语句2
		fmt.Println(reflect.ValueOf(&p).Elem().Addr())
		fmt.Println(p)
		
	}
	testCanAddr()
	//输出：
	//	true
	//	{luck1 1 200 [[22]]}
	//
	//	hello
	//	true
	//	0xc000032210
	//	0xc000032210
	//	word

}
type Calculate struct {

}


func (C *Calculate) Add(x,y int) int {
	return x+y
}

func ExampleStructTag_Lookup() {
	type S struct {
		F0 string `alias:"field_0"`
		F1 string `alias:""`
		F2 string
	}

	s := S{}
	st := reflect.TypeOf(s)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		//除了tag字段之外还有其他的字段，需要了解的自己动手查看field.xxx()即可
		if alias, ok := field.Tag.Lookup("alias"); ok {
			if alias == "" {
				fmt.Println("(blank)")
			} else {
				fmt.Println(alias)
			}
		} else {
			fmt.Println("(not specified)")
		}
	}

	// Output:
	// field_0
	// (blank)
	// (not specified)
}

func f2(string, int, bool) (string, error) {
	return "abc", nil
}

func f3(string, int, bool) {

}

type behavior interface {
	Setname111(s string) error
	Setage111()
}

type person struct {
	Name      string
	Age       int
	isstudent bool
}

func (p *person) Setname111(s string) error {
	p.Name = s
	return nil
}

func (p *person) Setage111() {
	p.Age = 18
}

type person11 struct {
	person
}

func check_err_reflect(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
