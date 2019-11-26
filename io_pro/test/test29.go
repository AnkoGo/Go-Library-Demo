package main

import (
	"fmt"
)

type jiekou interface {
	meth()
}

type zifu string

//对于go,我们不说zifu这个类实现了接口jiekou，我们说zifu的实例zifu实现了这个接口jiekou,如果下面是*zifu的话，我们就说
//zifu的*zifu实例实现了接口jiekou，但是不能说zifu的实例zifu实现了接口！他们是完全不同的东西！切记！这也是导致了下面将
//实例zifu（a）还是实例*zifu（&a）赋给接口j！
func (t zifu) meth() {
	fmt.Println("meth start.....")
	fmt.Printf("p(t.meth)===%p\n", (&t).meth)
	fmt.Printf("p(t)====%p\n", &t)
	fmt.Printf("%T调用方法成功！\n", t)
	fmt.Println("meth end.....")
}

func main88769() {
	a := zifu("a test")

	fmt.Println("测试zifu的类型：")
	a.meth()
	(&a).meth()

	fmt.Println()
	fmt.Println("测试接口：")

	// var j &jiekou//不可以对接口进行取值，因为此时的接口类型为nil，不可以对类
	// 型nil进行取内存，因为类型nil代表的是没申请内存，既然不存在内存，那就更加不存在去内存地址的可能了！
	// 所以不允许对其取指针！同理我们也不能写&nil!
	var j jiekou
	fmt.Println("j赋值之前的地址：", &j)
	fmt.Printf("j赋值之前的类型T(j)---%T\n", j)

	j = a
	fmt.Println("j赋值之后的地址：", &j)
	fmt.Printf("j赋值之后的类型T(j)---%T\n", j)
	fmt.Printf("p(j.meth)===%p\n", j.meth)
	fmt.Println("--------")
	j.meth()

}
