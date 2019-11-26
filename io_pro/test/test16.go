package main
import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"math"
)
//接口
type Pythagoras interface {
	Hypotenuse() float64
}

//实现接口的结构体
type Point struct {
	X, Y int
}
func (p Point) Hypotenuse() float64 {//求
	return math.Hypot(float64(p.X), float64(p.Y))
}

// interfaceEncode编码器函数 将接口值编码到编码器中。
func interfaceEncode(enc *gob.Encoder, p Pythagoras) {
	//除非已注册具体类型（或者说结构体），否则编码将失败。 我们在调用函数Example_interface()中注册了它。
	//将指针传递给接口，以便Encode看到（并因此发送）接口类型的值。 如果我们直接传递p，它将看到具体类型。
	//有关背景，请参见博客文章“反射定律(The Laws of Reflection)”。
	err := enc.Encode(&p)//必须是指针
	if err != nil {
		log.Fatal("encode:", err)
	}
}
// interfaceDecode解码器函数 解码流中的下一个接口值并返回它。
func interfaceDecode(dec *gob.Decoder) Pythagoras {
	//除非导线上的具体类型（或者说结构体）已注册，否则解码将失败。 我们在调用函数Example_interface()中注册了它。
	var p Pythagoras
	err := dec.Decode(&p)//必须是指针，指明解码成什么数据结构
	if err != nil {
		log.Fatal("decode:", err)
	}
	fmt.Printf("解码后的数据类型为：%T\n",p)
	fmt.Printf("解码后的数据为：%#v\n",p)
	return p//返回解码后的数据，这个数据不是接口，是接口的实现struct---Point类实例
}


//此示例说明如何编码接口值。 与常规类型的主要区别是注册实现接口的具体类型。
func Example_interface() {
	var network bytes.Buffer //网络的替身。装载编码后的数据的容器
	//我们必须注册编码器和解码器的具体类型（通常与编码器在不同的机器上）。 在每一端，这告诉引擎正在发送实现该接口的具体类型。
	gob.Register(Point{})//指明要编码哪个结构体
	//创建一个编码器并发送一些值。
	enc := gob.NewEncoder(&network)//编码到哪里去存着
	for i := 1; i <= 3; i++ {
		interfaceEncode(enc, Point{3 * i, 4 * i})//在这里实例化并且将实例传递过去
	}
	fmt.Println("编码后的数据为：",network)

	//创建一个解码器并接收一些值。
	dec := gob.NewDecoder(&network)//指定解码什么，指明编码什么数据结构
	for i := 1; i <= 3; i++ {
		result := interfaceDecode(dec)
		fmt.Printf("%T---%v\n",result,result.Hypotenuse())//调用struct的方法
	}

	//输出：
	//	编码后的数据为： {[44 16 0 10 109 97 105 110 46 80 111 105 110 116 255 129 3 1 1 5 80 111
	//					105 110 116 1 255 130 0 1 2 1 1 88 1 4 0 1 1 89 1 4 0 0 0 8 255 130 5 1
	//					6 1 8 0 21 16 0 10 109 97 105 110 46 80 111 105 110 116 255 130 5 1 12 1
	//					16 0 21 16 0 10 109 97 105 110 46 80 111 105 110 116 255 130 5 1 18 1 24 0] 0 0}
	//	解码后的数据类型为：main.Point
	//	解码后的数据为：main.Point{X:3, Y:4}
	//	main.Point---5
	//	解码后的数据类型为：main.Point
	//	解码后的数据为：main.Point{X:6, Y:8}
	//	main.Point---10
	//	解码后的数据类型为：main.Point
	//	解码后的数据为：main.Point{X:9, Y:12}
	//	main.Point---15
}


func main232546() {
	Example_interface()
}