package main
import "fmt"
type inter interface {

}
type Person struct {
	Name string
	Age int
}
//func (*Person)Say() Person {
//	fmt.Println("某人在say....")
//	return Person{
//		Name: "anko111",
//		Age:  22,
//	}
//}

func (*Person)Say() {
	fmt.Println("某人在say....")
}

type money float64

func (n *money)Discrib() string {
	return "money is xxxx"
}

func main346763() {
	var m money=16
	fmt.Println(m.Discrib())

	var f float64 =18.8
	var f2m= money(f)

	fmt.Println(f2m)
	fmt.Println(&f2m)
	//fmt.Println(&(money(f)))//虽然money(f)确实是一个实例，但是可能是go的一个bug，也可能是go不允许这样写，不过前者的原因大些
	fmt.Println(&(struct {}{}))//struct {}{}是一个实例

	P:=Person{
		Name:  "anko",
		Age:   20,
	}
	//&P.Say()//这样代表取say()返回值的指针。报错,即使我在Say（）中返回一个Person{ Name: "anko111", Age:  22, }对象也会报错，go的词法识别并没有那么智能
	(&P).Say()//这样代表取P的指针
	//fmt.Println(&Person)//Person是类型名，不允许对其取指针
	//
	//fmt.Println(&(inter{}))//不允许对接口进行取指针
}
