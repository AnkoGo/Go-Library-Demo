package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	fmt.Println("-------------reflect包中的零碎函数或者对象方法-----------------")
	// Append appends the values x to a slice s and returns the resulting slice.(追加value x到和切片s元素类型相同的切片s中去然后返回新的切片的value封装对象)
	// As in Go, each x's value must be assignable to the slice's element type.

	fmt.Println(reflect.Append(reflect.ValueOf([]int{'a', 'b', 'c'}), reflect.ValueOf(int('d'))))
	//下面报错，不同类型，panic: reflect.Set: value of type int8 is not assignable to type int
	//fmt.Println(reflect.Append(reflect.ValueOf([]int{'a','b','c'}),reflect.ValueOf(int8('d'))))
	//输出：
	//	[97 98 99 100]

	// AppendSlice appends a slice t to a slice s and returns the resulting slice.(跟上面的区别是，这个方法添加的不是元素，而是切片)
	// The slices s and t must have the same element type.
	fmt.Println(reflect.AppendSlice(reflect.ValueOf([]int{'a', 'b', 'c'}), reflect.ValueOf([]int{'d', 'e'})))
	//输出：
	//	[97 98 99 100 101]

	fmt.Println()
	// Swapper returns a function that swaps the elements in the provided
	// slice.(返回一个交换元素的函数，参数的kind类型必须是一个切片)
	// Swapper panics if the provided interface is not a slice.

	//sl_int:=[4]int{'a', 'b', 'c', 'd'}//报错，panic: reflect: call of Swapper on array Value
	sl_int := []int{'a', 'b', 'c', 'd'}
	swapper_func := reflect.Swapper(sl_int)
	fmt.Printf("%#v\n", swapper_func)
	swapper_func(1, 2) //交换索引1和索引2 的值

	fmt.Printf("%#v\n", sl_int)
	swapper_func(0, 3) //交换索引0和索引3 的值,不能超过索引限制
	fmt.Printf("%#v\n", sl_int)
	//输出：
	//	(func(int, int))(0x47b6e0)
	//	[]int{97, 99, 98, 100}
	//	[]int{100, 99, 98, 97}

	// ArrayOf returns the array type with the given count and element type.（根据给出的count长度和子元素的类型返回一个数组类型）
	// For example, if t represents int, ArrayOf(5, t) represents [5]int.（比如，如果给出的子元素是int类型，ArrayOf(5, t)返回一个类型 [5]int）
	//
	// If the resulting type would be larger than the available address space,（如果返回的类型比有效的地址空间要大，那么就会抛出panic）
	// ArrayOf panics.
	fmt.Println()
	fmt.Println(reflect.ArrayOf(5, reflect.TypeOf(2)))
	fmt.Println(reflect.ArrayOf(5, reflect.TypeOf("abc")))
	fmt.Println(reflect.ArrayOf(5, reflect.TypeOf([]byte{1, 2, 4})))
	fmt.Println(reflect.ArrayOf(5, reflect.TypeOf([...]byte{1, 2, 4})))
	//输出：
	//	[5]int
	//	[5]string
	//	[5][]uint8
	//	[5][3]uint8

	// Copy copies the contents of src into dst until either
	// dst has been filled or src has been exhausted.(Copy赋值src序列（仅限slice，string或者array）的元素进到dst（仅限slice,Uint8或者array）序列中去直到任何一方达到容量限制才停止)
	// It returns the number of elements copied.（它会返回复制元素的个数）
	// Dst and src each must have kind Slice or Array, and
	// dst and src must have the same element type.(Dst and src的kind类型都必须是Slice or Array，而且子元素都必须是相同的类型)
	//
	// As a special case, src can have kind String if the element type of dst is kind Uint8.（特别的，如果dst的子元素的kind类型是Uint8的话，那么此时src的kind类型可以是string）
	fmt.Println()
	sl_by_src := []byte{'a', 'b'}
	sl_by_dst := []byte{'c', 'd', 'e'}
	V_sl_by_src := reflect.ValueOf(sl_by_src)
	V_sl_by_dst := reflect.ValueOf(sl_by_dst)
	fmt.Println("复制了多少个元素？", reflect.Copy(V_sl_by_dst, V_sl_by_src))
	fmt.Println("复制了后的源是：", sl_by_src)
	fmt.Println("复制了后的目的是：", sl_by_dst)

	fmt.Println()
	sl_arr_src := [2]byte{'a', 'b'}
	sl_arr_dst := [3]byte{'c', 'd', 'e'}
	V_sl_arr_src := reflect.ValueOf(sl_arr_src)
	V_sl_arr_dst := reflect.ValueOf(&sl_arr_dst).Elem() //必须可寻址，上面之所以不用这样写是因为上面的切片是引用类型，是可以由值取到相对应的地址的类型，但是这里的数组是不可以的！
	fmt.Println("复制了多少个元素？", reflect.Copy(V_sl_arr_dst, V_sl_arr_src))
	fmt.Println("复制了后的源是：", sl_arr_src)
	fmt.Println("复制了后的目的是：", sl_arr_dst)

	fmt.Println()
	sl_str_src := string("abcde")
	sl_str_dst := []uint8{'x', 'y', 'z'}
	V_sl_str_src := reflect.ValueOf(sl_str_src)
	V_sl_str_dst := reflect.ValueOf(sl_str_dst)
	fmt.Println("复制了多少个元素？", reflect.Copy(V_sl_str_dst, V_sl_str_src))
	fmt.Println("复制了后的源是：", sl_str_src)
	fmt.Println("复制了后的目的是：", sl_str_dst)
	//输出：
	//	复制了多少个元素？ 2
	//	复制了后的源是： [97 98]
	//	复制了后的目的是： [97 98 101]
	//
	//	复制了多少个元素？ 2
	//	复制了后的源是： [97 98]
	//	复制了后的目的是： [97 98 101]
	//
	//	复制了多少个元素？ 3
	//	复制了后的源是： abcde
	//	复制了后的目的是： [97 98 99]

	fmt.Println()
	// ChanOf returns the channel type with the given direction and element type.
	// For example, if t represents int, ChanOf(RecvDir, t) represents <-chan int.
	//
	// The gc runtime imposes a limit of 64 kB on channel element types.
	// If t's size is equal to or exceeds this limit, ChanOf panics.
	//ChanOf返回元素类型为t、方向为dir的通道类型。运行时GC强制将通道的元素类型的大小限定为64kb。如果t的尺寸大于或等于该限制，本函数将会panic。
	T_BothDirChan := reflect.ChanOf(reflect.BothDir, reflect.TypeOf(4))
	fmt.Println(T_BothDirChan)

	T_RecvDirChan := reflect.ChanOf(reflect.RecvDir, reflect.TypeOf(4))
	fmt.Println(T_RecvDirChan)

	T_SendDirChan := reflect.ChanOf(reflect.SendDir, reflect.TypeOf(4))
	fmt.Println(T_SendDirChan)

	T_BothDirChan111 := reflect.ChanOf(1|2, reflect.TypeOf(4))
	fmt.Println(T_BothDirChan111)

	T_RecvDirChan111 := reflect.ChanOf(1, reflect.TypeOf(4))
	fmt.Println(T_RecvDirChan111)

	T_SendDirChan111 := reflect.ChanOf(2, reflect.TypeOf(4))
	fmt.Println(T_SendDirChan111)
	//输出：
	//	chan int
	//	<-chan int
	//	chan<- int
	//	chan int
	//	<-chan int
	//	chan<- int

	fmt.Println()
	// MakeChan creates a new channel with the specified type and buffer size.（MakeChan通过使用type类型(通道方向只能是BothDir)和缓存大小来创建一个新的chan实例，）
	V_BothDirChan := reflect.MakeChan(T_BothDirChan, 3)
	fmt.Printf("%#v\n", V_BothDirChan)

	//V_RecvDirChan:= reflect.MakeChan(T_RecvDirChan,3)//报错，panic: reflect.MakeChan: unidirectional channel type
	//fmt.Printf("%#v",V_RecvDirChan)
	//
	//V_SendDirChan:= reflect.MakeChan(T_SendDirChan,3)//报错，panic: reflect.MakeChan: unidirectional channel type
	//fmt.Printf("%#v",V_SendDirChan)

	fmt.Println()
	// MakeMap creates a new map with the specified type.(MakeMap使用给定的参数(只有一个类型参数，无初始化子元素个数值)来创建一个map的value封装对象)
	V_map := reflect.MakeMap(reflect.TypeOf(map[string]int{}))
	fmt.Printf("%#v\n", V_map)
	//fmt.Printf("%#v\n",V_map.Cap())//对于map来说是len,不是cap,cap()的调用者只能是Array, Chan, or Slice.我们这样无法查询cap的大小
	fmt.Printf("%#v\n", V_map.Len())
	//输出：
	//	map[string]int{}
	//	0

	fmt.Println()
	// MakeSlice creates a new zero-initialized slice value(MakeSlice使用给定的参数（slice类型，len和cap大小）来创建一个新的初始化为子元素类型零值的切片类型的value封装对象)
	// for the specified slice type, length, and capacity.

	//V_slice:= reflect.MakeSlice(reflect.TypeOf([2]byte{}),2,5)//报错panic: reflect.MakeSlice of non-slice type，只能是切片
	V_slice := reflect.MakeSlice(reflect.TypeOf([]byte{}), 2, 5)
	fmt.Printf("%#v,长度为：%v,cap容量为：%v\n", V_slice, V_slice.Len(), V_slice.Cap())
	//输出：
	//	[]uint8{0x0, 0x0},长度为：2,cap容量为：5

	fmt.Println()
	// MakeMapWithSize creates a new map with the specified type(MakeMap使用给定的参数(有2个参数，分别是map的类型和初始化空间n---子元素个数值)来创建一个map的value封装对象)
	// and initial space for approximately n elements.
	V_mapSize := reflect.MakeMapWithSize(reflect.TypeOf(map[string]int{}), 5)
	fmt.Printf("%#v\n", V_mapSize)
	fmt.Printf("%v\n", V_mapSize.Len())

	//输出：
	//	map[string]int{}
	//	0，暂时还不知道怎么查看map的cap大小，这个len不是查看容量的！

	fmt.Println()
	// MakeFunc returns a new function of the given Type
	// that wraps the function fn. When called, that new function
	// does the following:
	//
	//	- converts its arguments to a slice of Values.
	//	- runs results := fn(args).
	//	- returns the results as a slice of Values, one per formal result.
	//
	// The implementation fn can assume that the argument Value slice
	// has the number and type of arguments given by typ.
	// If typ describes a variadic function, the final Value is itself
	// a slice representing the variadic arguments, as in the
	// body of a variadic function. The result Value slice returned by fn
	// must have the number and type of results given by typ.
	//
	// The Value.Call method allows the caller to invoke a typed function
	// in terms of Values; in contrast, MakeFunc allows the caller to implement
	// a typed function in terms of Values.
	//
	// The Examples section of the documentation includes an illustration
	// of how to use MakeFunc to build a swap function for different types.
	//
	//MakeFunc返回一个具有给定类型、包装函数fn的函数的Value封装。当被调用时，该函数会：
	//
	//- 将提供给它的参数转化为Value切片
	//- 执行results := fn(args)
	//- 将results中每一个result依次排列作为返回值函数fn的实现可以假设参数Value切片匹配typ类型指定的参数数目和类型。
	// 	如果typ表示一个可变参数函数类型，参数切片中最后一个Value本身必须是一个包含所有可变参数的切片。fn返回的结果Value切片也必须匹配typ类型指定的结果数目和类型。
	//
	//Value.Call方法允许程序员使用Value调用一个有类型约束的函数；反过来，MakeFunc方法允许程序员使用Value实现一个有类型约束的函数。
	//
	//下例是一个用MakeFunc创建一个生成不同参数类型的swap函数的代码及其说明。

	// swap is the implementation passed to MakeFunc.
	// It must work in terms of reflect.Values so that it is possible
	// to write code without knowing beforehand what the types
	// will be.
	//这个函数主要是交换了序列中的2个value值并返回交换后的结果
	swap := func(in []reflect.Value) []reflect.Value {
		return []reflect.Value{in[1], in[0]}
	}
	// makeSwap expects fptr to be a pointer to a nil function.
	// It sets that pointer to a new function created with MakeFunc.
	// When the function is invoked, reflect turns the arguments
	// into Values, calls swap, and then turns swap's result slice
	// into the values returned by the new function.
	makeSwap := func(fptr interface{}) {
		// fptr is a pointer to a function.
		// Obtain the function value itself (likely nil) as a reflect.Value
		// so that we can query its type and then set the value.
		fn := reflect.ValueOf(fptr).Elem() //获取源对象上面的子元素
		// Make a function of the right type.
		//将第一个参数指定的函数类型（各种参数返回值不同的函数的类型也不同，所以需要指定）作为类型，将第二个参数指定的函数体作为函数体来实例化一个函数实例！！
		//返回一个函数体的value值对象
		v := reflect.MakeFunc(fn.Type(), swap)
		// Assign it to the value fn represents.
		fn.Set(v) //在源函数对象上面进行设置值，无论源函数对象上面初始时候是否有函数体了都会进行替换，很明显这里是还没有初始化函数体的！
	}
	// Make and call a swap function for ints.
	var intSwap func(int, int) (int, int) //这个函数还没给定函数体的值，我们将在makeSwap(&intSwap)中的reflect.MakeFunc(fn.Type(), swap)中给定函数体的value值
	makeSwap(&intSwap)                    //这里必须传递指针
	fmt.Println(intSwap(0, 1))
	// Make and call a swap function for float64s.
	var floatSwap func(float64, float64) (float64, float64)
	makeSwap(&floatSwap)
	fmt.Println(floatSwap(2.72, 3.14))
	//输出：
	//	1 0
	//	3.14 2.72

	fmt.Println()
	// NewAt returns a Value representing a pointer to a value of the
	// specified type, using p as that pointer.
	//NewAt返回一个Value类型值，该值持有一个指向类型为typ、地址为p的值的指针。

	//下一行会报错：cannot use &[]byte literal (type *[]byte) as type unsafe.Pointer in argument to reflect.NewAt
	//*[]byte也是一种类型，他和unsafe.Pointer不同，所以需要转换！
	//T_NewAt:=reflect.NewAt(reflect.TypeOf([]byte{}), &[]byte{'a','b','c'})
	T_NewAt := reflect.NewAt(reflect.TypeOf([]byte{}), unsafe.Pointer(&[]byte{'a', 'b', 'c'}))
	fmt.Println(T_NewAt)

	fmt.Println()
	//用来判断两个值是否深度一致：除了类型相同；在可以时（主要是基本类型）会使用==；但还会比较array、slice的成员，map的键值对，
	//结构体字段进行深入比对。map的键值对，对键只使用==，但值会继续往深层比对。DeepEqual函数可以正确处理循环的类型。函数类型只
	//有都会nil时才相等；空切片不等于nil切片；还会考虑array、slice的长度、map键值对数。
	fmt.Println("切片深度对比")
	b_DeepEqual := reflect.DeepEqual([]byte{'a', 'b', 'c'}, []byte{'a', 'b', 'c'})
	fmt.Println(b_DeepEqual)

	b_DeepEqual = reflect.DeepEqual([]byte{'a', 'b', 'c'}, []byte{'a', 'b'})
	fmt.Println(b_DeepEqual)

	b_DeepEqual = reflect.DeepEqual([]byte{'a', 'b', 'c'}, []int{'a', 'b', 'c'})
	fmt.Println(b_DeepEqual)

	b_DeepEqual = reflect.DeepEqual([]byte{'a', 'b', 'c'}, []uint8{'a', 'b', 'c'}) //byte类型是uint8类型的别名
	fmt.Println(b_DeepEqual)

	fmt.Println("数组深度对比")
	b_DeepEqual = reflect.DeepEqual([...]byte{'a', 'b', 'c'}, [...]byte{'a', 'b', 'c'})
	fmt.Println(b_DeepEqual)

	b_DeepEqual = reflect.DeepEqual([...]byte{'a', 'b', 'c'}, [...]uint8{'a', 'b', 'c'}) //byte类型是uint8类型的别名
	fmt.Println(b_DeepEqual)

	fmt.Println("参数化的nil值深度对比")

	b_DeepEqual = reflect.DeepEqual(nil, nil)
	fmt.Println(b_DeepEqual)

	fmt.Println("带类型的的nil值深度对比")

	b_DeepEqual = reflect.DeepEqual(([]byte)(nil), ([]byte)(nil))
	fmt.Println(b_DeepEqual)

	b_DeepEqual = reflect.DeepEqual(([]byte)(nil), ([]uint8)(nil))
	fmt.Println(b_DeepEqual)

	b_DeepEqual = reflect.DeepEqual(([]byte)(nil), ([]int)(nil))
	fmt.Println(b_DeepEqual)

	fmt.Println("map类型深度对比")

	b_DeepEqual = reflect.DeepEqual(map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1, "b": 2})
	fmt.Println(b_DeepEqual)

	b_DeepEqual = reflect.DeepEqual(map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1, "b": 2, "c": 3})
	fmt.Println(b_DeepEqual)

	b_DeepEqual = reflect.DeepEqual(map[string]int{"a": 1, "b": 2}, map[string]int64{"a": 1, "b": 2})
	fmt.Println(b_DeepEqual)
	//更多其他类型的深度对比可以自己尝试！
	//输出：
	//	切片深度对比
	//	true
	//	false
	//	false
	//	true
	//	数组深度对比
	//	true
	//	true
	//	参数化的nil值深度对比
	//	true
	//	带类型的的nil值深度对比
	//	true
	//	true
	//	false
	//	map类型深度对比
	//	true
	//	false
	//	false

	fmt.Println()
	// Indirect returns the value that v points to.(Indirect返回v指向的value值)
	// If v is a nil pointer, Indirect returns a zero Value.(如果v是nil指针类型的话(比如下面的V_nil_arr1对象)，会返回zero Value)
	// If v is not a pointer, Indirect returns v.(如果v不是指针类型的话，会返回原本的v)
	//说白了，就是返回值类型的值，引用类型的引用指向的值，指针类型的指针指向的值，如果指针类型指向的值是一个引用类型的话则继续会探寻到引用类型对应引用的值
	fmt.Println("返回应用类型或者值类型抑或者是指针类型的值")
	ls := []byte{'a', 'b', 'c'}
	arr := [...]byte{'a', 'b', 'c', 'd'}
	//nil_by:=(*[]byte)(nil)//给引用类型赋值nil，但是指针的值是一个无效的value值，对于nil值的指针类型调用IsZero和IsNil一定会报错的！
	nil_by := ([]byte)(nil)
	nil_arr := [2]byte{} //值类型应该这样初始化零值，也可以var nil_arr [2]byte来进行初始化零值，不可以给值类型的对象赋值nil,只能给引用类型的对象赋值nil
	//nil_arr:=(*[2]byte)(nil)//不可以给值类型的对象赋值nil,只能给引用类型的对象赋值nil，对于nil值的指针类型调用IsZero和IsNil一定会报错的！

	V_ls := reflect.ValueOf(ls)
	V_ls1 := reflect.ValueOf(&ls)
	V_arr := reflect.ValueOf(arr)
	V_arr1 := reflect.ValueOf(&arr)
	V_nil_by := reflect.ValueOf(nil_by)
	V_nil_by1 := reflect.ValueOf(&nil_by)
	V_nil_arr := reflect.ValueOf(nil_arr)
	V_nil_arr1 := reflect.ValueOf(&nil_arr).Elem()
	fmt.Println("==Elem===", reflect.ValueOf(&nil_arr).Elem().Kind())

	V_Indirect := reflect.Indirect(V_ls)
	V_Indirect1 := reflect.Indirect(V_ls1)
	V_Indirect2 := reflect.Indirect(V_arr)
	V_Indirect3 := reflect.Indirect(V_arr1)
	V_Indirect4 := reflect.Indirect(V_nil_by)
	V_Indirect5 := reflect.Indirect(V_nil_by1)
	V_Indirect6 := reflect.Indirect(V_nil_arr)
	V_Indirect7 := reflect.Indirect(V_nil_arr1)
	fmt.Println("==Indirect===", reflect.Indirect(V_nil_arr1).Kind())

	fmt.Println(V_Indirect)
	fmt.Println(V_Indirect1)
	fmt.Println(V_Indirect2)
	fmt.Println(V_Indirect3)

	// IsZero reports whether v is the zero value for its type.（IsZero报告v是否是类型的零值（注意不是判断value的零值，如果对此类型进行了判断会抛出异常））
	// It panics if the argument is invalid.（如果这个值是value的零值（或者叫做无效值，无值）的话，那么他会抛出panic）
	// 其实这个判断的类型就比下面的IsNil（）方法的范围要大些，他不一定需要判断引用类型，可以判断任意的类型（除了上面说到的value的零值（或者叫做无效值，无值））的零值！总之区别就是前者是判零值，后者也是判nil值
	// （但是其实也是引用类型的零值，或者叫未声明value的内存时候的初始值）

	// IsNil返回是否是对象的类型的零值，这个对象必须是引用类型：chan, func, interface, map, pointer, or slice 。
	// 如果不是的话，会抛出 异常。注意IsNil并不总是等价于go语言中值与参数nil的常规比较。例如：如果v是通过使用某个值为nil的接口调用ValueOf函数创建的，
	// v.IsNil()返回真，但是如果v是Value零值(注意和零值的value区别，Value零值指的是类型的零值)，会panic。也就是nil必须带有类型！不能是参数(不带类型)的nil
	fmt.Printf("%T---%#v,是zero Value么？%v;是nil value么？%v\n", V_Indirect4, V_Indirect4, V_Indirect4.IsZero(), V_Indirect4.IsNil()) //对于引用类型来说，IsZero()和IsNil()是一样的效果
	fmt.Printf("%T---%#v,是zero Value么？%v;是nil value么？%v\n", V_Indirect5, V_Indirect5, V_Indirect5.IsZero(), V_Indirect5.IsNil())
	fmt.Printf("%T---%#v,是zero Value么？%v;是nil value么？值类型不可调用V_Indirect6.IsNil()\n", V_Indirect6, V_Indirect6, V_Indirect6.IsZero())
	fmt.Printf("%T---%#v,是zero Value么？不可对invalid Value对象调用V_Indirect7.IsZero();是nil value么？不可对invalid Value对象调用V_Indirect7.IsNil()\n", V_Indirect7, V_Indirect7)

	//输出：
	//	返回应用类型或者值类型抑或者是指针类型的值
	//	[97 98 99]
	//	[97 98 99]
	//	[97 98 99 100]
	//	[97 98 99 100]
	//	reflect.Value---[]uint8(nil),是zero Value么？true;是nil value么？true
	//	reflect.Value---[]uint8(nil),是zero Value么？true;是nil value么？true
	//	reflect.Value---[2]uint8{0x0, 0x0},是zero Value么？true;是nil value么？值类型不可调用V_Indirect6.IsNil()
	//	reflect.Value---[2]uint8{0x0, 0x0},是zero Value么？不可对invalid Value对象调用V_Indirect7.IsZero();是nil value么？不可对invalid Value对象调用V_Indirect7.IsNil()
	//更多的请参考上一节的说明！

	fmt.Println()
	// SliceOf returns the slice type with element type t.(SliceOf返回含有子元素类型t的切片类型[]t)
	// For example, if t represents int, SliceOf(t) represents []int.(比如，如果参数t是int类型的话， SliceOf(t)返回[]int)
	fmt.Println(reflect.SliceOf(reflect.TypeOf(2)))
	fmt.Println(reflect.SliceOf(reflect.TypeOf([]byte{})))
	fmt.Println(reflect.SliceOf(reflect.TypeOf("abc")))
	//输出：
	//	[]int
	//	[][]uint8
	//	[]string

	fmt.Println()
	// MapOf returns the map type with the given key and element types.(根据给出的参数中指定的键值类型来返回一个自定义的map的type类型封装对象。)
	// For example, if k represents int and e represents string,（比如，如果给定k是int类型，e是string类型，那么MapOf(k, e)就返回map[int]string的type类型对象）
	// MapOf(k, e) represents map[int]string.
	//
	// If the key type is not a valid map key type (that is, if it does
	// not implement Go's == operator), MapOf panics.
	fmt.Println(reflect.MapOf(reflect.TypeOf(2), reflect.TypeOf("abc")))
	//输出：
	//	map[int]string

	fmt.Println()
	// StructOf returns the struct type containing fields.（利用StructField实例对象来创建一个 struct type类型）
	// The Offset and Index fields are ignored and computed as they would be
	// by the compiler.
	//
	// StructOf currently does not generate wrapper methods for embedded
	// fields and panics if passed unexported StructFields.
	// These limitations may be lifted in a future version.

	//type StructField struct {
	//	// Name是字段的名字。PkgPath是非导出字段的包路径，对导出字段该字段为""。
	//	// 参见http://golang.org/ref/spec#Uniqueness_of_identifiers
	//	Name    string
	//	PkgPath string
	//	Type      Type      // 字段的类型
	//	Tag       StructTag // 字段的标签
	//	Offset    uintptr   // 字段在结构体中的字节偏移量
	//	Index     []int     // 用于Type.FieldByIndex时的索引切片
	//	Anonymous bool      // 是否匿名字段
	//}
	//StructField类型描述结构体中的一个字段的信息。

	struct11 := struct {
		Name string `name`
		age  int    `age`
	}{"anko", 33}
	T_struct11 := reflect.TypeOf(struct11)
	struct_field := T_struct11.Field(0)
	struct_field1 := T_struct11.Field(1)
	fmt.Printf("%T---%#v\n", struct_field, struct_field)
	fmt.Printf("%T---%#v\n", struct_field1, struct_field1)
	//reflect.TypeOf()
	fmt.Println(reflect.StructOf([]reflect.StructField{struct_field}))
	//不可导出的字段不可以被使用在StructOf方法
	//fmt.Println(reflect.StructOf([]reflect.StructField{struct_field,struct_field1}))
	//输出：
	//	reflect.StructField---reflect.StructField{Name:"Name", PkgPath:"", Type:(*reflect.rtype)(0x4f46a0), Tag:"name", Offset:0x0, Index:[]int{0}, Anonymous:false}
	//	reflect.StructField---reflect.StructField{Name:"age", PkgPath:"main", Type:(*reflect.rtype)(0x4f3d20), Tag:"age", Offset:0x10, Index:[]int{1}, Anonymous:false}
	//	struct { Name string "name" }

	fmt.Println()
	// Zero returns a Value representing the zero value for the specified type.(Zero根据给出的type类型返回该类型对应的类型零值的Value封装对象)
	// The result is different from the zero value of the Value struct,
	// which represents no value at all.（返回值是不同于Value结构体的无效值（或者零值，这个值相当于无值）），从这里可以看得出跟Iszero()方法的用法是不一样的！
	// For example, Zero(TypeOf(42)) returns a Value with Kind Int and value 0.（比如，Zero(TypeOf(42))返回Value对象（kind类型为Int，value为0））
	// The returned value is neither addressable nor settable.(返回值value可以寻址也可以设置值)
	fmt.Println(reflect.Zero(reflect.TypeOf(2)))
	fmt.Println(reflect.Zero(reflect.TypeOf("abc")))
	fmt.Println(reflect.Zero(reflect.TypeOf([]byte{'a', 'b'})))
	fmt.Println(reflect.Zero(reflect.TypeOf([2]byte{'a', 'b'})))
	fmt.Println(reflect.Zero(reflect.TypeOf(&[]byte{'a', 'b'})))
	fmt.Println(reflect.Zero(reflect.TypeOf(make(chan int, 2))))
	fmt.Println(reflect.Zero(reflect.TypeOf(make(map[string]int, 2))))
	fmt.Println(reflect.Zero(reflect.TypeOf(true)))
	//fmt.Println(reflect.Zero(reflect.TypeOf(nil)))//panic: reflect: Zero(nil)
	fmt.Println(reflect.Zero(reflect.TypeOf([]byte(nil))))    //panic: reflect: Zero(nil)
	fmt.Println(reflect.Zero(reflect.TypeOf((*[]byte)(nil)))) //panic: reflect: Zero(nil)
	//输出：
	//	0
	//		（空字符串）
	//	[]
	//	[0 0]
	//	<nil>
	//	<nil>
	//	map[]
	//	false
	//	[]
	//	<nil>，如果是方法iszero()的话，这里就不会输出<nil>，而是直接报错！iszero返回的是指针对象指向的内存的对象的类型的零值，而Zero返回的是该内存上的指针对象的类型的零值，注意区分

	fmt.Println()
	// Select executes a select operation described by the list of cases.
	// Like the Go select statement, it blocks until at least one of the cases
	// can proceed, makes a uniform pseudo-random choice,
	// and then executes that case. It returns the index of the chosen case
	// and, if that case was a receive operation, the value received and a
	// boolean indicating whether the value corresponds to a send on the channel
	// (as opposed to a zero value received because the channel is closed).
	//Select函数执行cases切片描述的select操作。 类似go的select语句，它会阻塞直到至少一条case可以执行，
	//从可执行的case中（伪）随机的选择一条，并执行该条case。
	// 它会返回选择执行的case的索引chosen int，以及如果执行的是接收case时，
	//会返回接收到的值recv Value，以及一个布尔值recvOK bool说明该值是否对应于通道中某次发送的值（用以区分通道关闭时接收到的零值，此时recvOK会设为false）。

	//type SelectCase struct {
	//	Dir  SelectDir // case的方向
	//	Chan Value     // 使用的通道（send发送的存储目的地/resv取值时候的取源）
	//	Send Value     // 用于发送的值
	//}
	// SelectCase描述select操作中的单条case。Case的类型由通信方向Dir决定。
	//
	//如果Dir是SelectDefault，该条case代表default case。Chan和Send字段必须是Value零值(无效值)。
	//
	//如果Dir是SelectSend，该条case代表一个发送操作。Chan字段底层必须是一个chan类型，Send的底层必须是可以直接赋值给该chan类型的成员类型的类型。
	// 如果Chan是Value零值(无效值)，则不管Send字段是不是零值，该条case都会被忽略。
	//
	//如果Dir是SelectRecv，该条case代表一个接收操作。Chan字段底层必须是一个chan类型，而Send必须是一个Value零值(无效值)。
	// 如果Chan是Value零值(无效值)，该条case会被忽略，但Send字段仍需是Value零值。当该条case被执行时，接收到的值会被Select返回。

	var chan_send = make(chan int, 10)
	var chan_resv = make(chan int, 11)

	chan_send <- 1
	chan_send <- 2
	chan_send <- 3

	chan_resv <- 4
	chan_resv <- 5
	chan_resv <- 6

	SelectCase_Default := reflect.SelectCase{
		reflect.SelectDefault,
		reflect.Value{},
		reflect.Value{},
	}

	SelectCase_Send := reflect.SelectCase{
		reflect.SelectSend,
		reflect.ValueOf(chan_send),
		reflect.ValueOf(9),
		//表示chan_resv<-9
	}

	SelectCase_Recv := reflect.SelectCase{
		reflect.SelectRecv,
		reflect.ValueOf(chan_resv),
		reflect.Value{},
		//表示resv<-chan_resv(其中resv会在reflect.Select（）方法中返回)
	}

	fmt.Println(reflect.Select([]reflect.SelectCase{SelectCase_Default, SelectCase_Send}))
	fmt.Println(reflect.Select([]reflect.SelectCase{SelectCase_Recv, SelectCase_Default})) //我们故意不让SelectCase_Default在第一位置
	close(chan_send)
	close(chan_resv)
	//记得关闭通道发送一个通道结束发送或者接收值的信息
	//用for range也是可以的！
	for v := range chan_resv {
		fmt.Println("从chan_resv通道取出一个值：", v)
	}

	for v := range chan_send {
		fmt.Println("从chan_send通道取出一个值：", v)
	}
	//for  {
	//	v,ok:=<-chan_resv
	//	if !ok{
	//		break
	//	}
	//	fmt.Println("从chan_resv通道取出一个值：",v)
	//}
	//
	//for  {
	//	v,ok:=<-chan_send
	//	if !ok{
	//		break
	//	}
	//	fmt.Println("从chan_send通道取出一个值：",v)
	//}
	//输出：
	//	1 <invalid reflect.Value> false
	//	0 4 true（chan_resv通道的4在这里看到没有？）
	//	从chan_resv通道取出一个值： 5
	//	从chan_resv通道取出一个值： 6
	//	从chan_send通道取出一个值： 1
	//	从chan_send通道取出一个值： 2
	//	从chan_send通道取出一个值： 3
	//	从chan_send通道取出一个值： 9

	fmt.Println()
	//我去，终于写完了！捂脸！

}

func check_err_reflect(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
