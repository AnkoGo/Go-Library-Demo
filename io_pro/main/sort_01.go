package main

import (
	"fmt"
	"sort"
)

func main5647() {
	//整个sort包的排序都是采用的堆排序,sort.go里面的siftDown()函数就是具体的实现

	fmt.Println("============================================")
	//Ints()按递增顺序原地对一片整数进行排序。底层调用Sort()
	ls_int:=[]int{9,81,2,4,12}
	fmt.Println(sort.IntsAreSorted(ls_int))//false
	sort.Ints(ls_int)
	fmt.Println(ls_int)//[2 4 9 12 81]
	// IntsAreSorted测试一个整数切片是否按升序排序。
	fmt.Println(sort.IntsAreSorted(ls_int))//true

	// SearchInts在按整数排序的切片中搜索x，并返回Search指定的索引。 返回值是在x不存在的情况下插入x的索引（它可能是len（a））。
	//切片必须按升序排序。
	i:=sort.SearchInts(ls_int,3)
	fmt.Println(i)//1
	fmt.Println(ls_int)//[2 4 9 12 81],并没有真的插入，仅仅是假如插入的话应该插入的位置索引


	fmt.Println("============================================")

	fmt.Println("同理的还有float64类型和string类型，下面进行展示：")
	ls_float64:=[]float64{9.1,81.1,2.2,4.4,12.5}
	fmt.Println(sort.Float64sAreSorted(ls_float64))//false
	sort.Float64s(ls_float64)//
	fmt.Println(ls_float64)//[2.2 4.4 9.1 12.5 81.1]

	fmt.Println(sort.Float64sAreSorted(ls_float64))//true

	i1:=sort.SearchFloat64s(ls_float64,3.1)//如果这里填写3的话也行，默认就是3.0了
	fmt.Println(i1)//1
	fmt.Println(ls_float64)//[2.2 4.4 9.1 12.5 81.1]

	fmt.Println("============================================")
	ls_string:=[]string{"x","a","y","b","aa"}
	fmt.Println(sort.StringsAreSorted(ls_string))//false
	sort.Strings(ls_string)
	fmt.Println(ls_string)//[a aa b x y]

	fmt.Println(sort.StringsAreSorted(ls_string))//true

	i2:=sort.SearchStrings(ls_string,"k")//
	fmt.Println(i2)//3
	fmt.Println(ls_string)//[a aa b x y]


	fmt.Println("=================通用的排序函数，是上面排序函数的基石111===========================")


	//Sort()对数据进行排序。
	//一次调用data.Len确定n，然后调用O（n * log（n））调用data.Less和data.Swap。 不能保证排序是稳定的。
	ls_Array:=ArraySlice{[]byte{23,34,56,99,34,24},[]byte{23,34},[]byte{23,34,3,6}}
	fmt.Println(sort.SliceIsSorted(ls_Array,ls_Array.Less))//false
	sort.Sort(ls_Array)//这种方式不稳定的排序
	//sort.Stable(ls_Array)//稳定的排序，参数必须是sort.Interface{}的实现(注意不是interface{}任意类型，这是2种不同的接口)
	//sort.SliceStable(ls_Array,ls_Array.Less)//稳定的排序，参数任意类型,不过要指定比较的函数，对于实现了sort.Interface{}的类型是不大适合这种比较方式的！

	fmt.Println(ls_Array)//[[23 34] [23 34 3 6] [23 34 56 99 34 24]]

	//fmt.Println(sort.SliceIsSorted(ls_Array,ls_Array.Less))//true，这种方式必须传函数，对于我们实现了这个接口的实例的情况我们最好采用下面的方式
	fmt.Println(sort.IsSorted(ls_Array))//true

	//这个search其实也是上面所有searchxxx的基石，完全可以封装成为一个SearchArray（）来实现简单的调用而不是像下面那样写这么多
	//但是我们为了展示sort.Search()函数就不封装了
	i3:=sort.Search(ls_Array.Len(),func(i int) bool { return len(ls_Array[i]) >= len([]byte{34,123,99,35,57}) })
	//下面是对于search()的一些底层实现文档：
	//	func Search
	//	func Search(n int, f func(int) bool) int
	//	Search函数采用二分法搜索找到[0, n)区间内最小的满足f(i)==true的值i。也就是说，Search函数希望f在输入位于区间[0, n)的前面
	//	某部分（可以为空）时返回假，而在输入位于剩余至结尾的部分（可以为空）时返回真；Search函数会返回满足f(i)==true的最小值i。如果
	//	没有该值，函数会返回n。注意，未找到时的返回值不是-1，这一点和strings.Index等函数不同。Search函数只会用区间[0, n)内的值调用f。
	//
	//	一般使用Search找到值x在插入一个有序的、可索引的数据结构时，应插入的位置。
	//	这种情况下，参数f（通常是闭包）会捕捉应搜索的值和被查询的数据集。
	//
	//	例如，给定一个递增顺序的切片，调用Search(len(data), func(i int) bool { return data[i] >= 23 })会返回data中
	//	最小的索引i满足data[i] >= 23。如果调用者想要知道23是否在切片里，它必须另外检查data[i] == 23。
	//
	//	搜索递减顺序的数据时，应使用<=运算符代替>=运算符。
	fmt.Println(i3)//2
	fmt.Println(ls_Array)//[[23 34] [23 34 3 6] [23 34 56 99 34 24]]


	fmt.Println("=================通用的排序函数，不实现接口也能对任意序列类型排序==========================")
	//上面我们通过了创建类然后实现3个方法来实现interface这个接口，其实go也提供了不用我们实现这个接口的api也能实现对非string,int和float64类型之外的序列类型进行排序
	//还是通过[]byte类型的长度来排序
	//interface{}表示任意的类型
	ls_Array111:=[]interface{}{[]byte{23,34,56,99,34,24},[]byte{23,34},[]byte{23,34,3,6}}
	//ls_Array111:=[][]byte{[]byte{23,34,56,99,34,24},[]byte{23,34},[]byte{23,34,3,6}}

	//还需要指定这个类的两元素比较什么，但是不用指定如何就行交换元素
	Less_slice:=func(i, j int) bool {
		return len(ls_Array111[i].([]byte)) < len(ls_Array111[j].([]byte))//如果元素是interface{}的话就采用这种方式来进行比较
		//return len(ls_Array111[i]) < len(ls_Array111[j])//如果元素是具体的类型就采用这种方法来 进行比较
	}
	//判断是否升序排序
	fmt.Println(sort.SliceIsSorted(ls_Array111,Less_slice))//这个函数可以接受任意类型
	//fmt.Println(sort.IsSorted(ls_Array111,Less_slice))//这个函数可以仅仅可以接受实现了Interface{}接口的实现类
	//进行排序
	sort.Slice(ls_Array111,Less_slice)//flase
	fmt.Println(ls_Array111)//[[23 34] [23 34 3 6] [23 34 56 99 34 24]]
	//判断是否升序排序
	fmt.Println(sort.SliceIsSorted(ls_Array111,Less_slice))//true

	i4:=sort.Search(len(ls_Array111),func(i int) bool { return len(ls_Array111[i].([]byte)) >= len([]byte{34,123,99,35,57}) })
	//i4:=sort.Search(len(ls_Array111),func(i int) bool { return len(ls_Array111[i]) >= len([]byte{34,123,99,35,57}) })//如果元素是具体的类型就采用这种方法来 进行比较

	fmt.Println(i4)//2
	fmt.Println(ls_Array)//[[23 34] [23 34 3 6] [23 34 56 99 34 24]]

	fmt.Println("------------对sort.Interface{}的实现进行反序--------------------------")
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)//[6 5 4 3 2 1]

	fmt.Println("-------------------sort包里面的类型----------------------------")
	//下面仅仅介绍sort.IntSlice类型，其他的类型不再列举，同理的！
	var int_ls =sort.IntSlice{11,3,45,235,75}
	fmt.Println(int_ls)
	fmt.Println(int_ls.Less(0,1))//判断第1是否小于第2个元素
	int_ls.Swap(0,1)//交换第1第2个元素的值
	fmt.Println(int_ls)

	int_ls.Sort()//排序
	fmt.Println(int_ls)
	fmt.Println(int_ls.Len())
	fmt.Println(int_ls.Search(13))//判断这个13能插入这个序列的哪个位置
	//输出：
	//[11 3 45 235 75]
	//false
	//[3 11 45 235 75]
	//[3 11 45 75 235]
	//5
	//2

}


type ArraySlice [][]byte
//通过[]byte类型的长度来排序
func (p ArraySlice) Len() int           { return len(p) }//长度
func (p ArraySlice) Less(i, j int) bool { return len(p[i]) < len(p[j]) }//两元素比较什么
func (p ArraySlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }//两元素比较后做什么

























