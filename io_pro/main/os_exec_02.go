package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	r := "世"
	fmt.Println(bytes.NewBufferString(r).Bytes())

	b1:='世'
	fmt.Println("----",strconv.QuoteRune(b1))
	buf := make([]byte, 3)
	n := utf8.EncodeRune(buf, b1)
	fmt.Println(buf)
	fmt.Println(n)


	//b := []byte{0x00, 0x00, 0x03, 0xe8}
	//b := []byte{ 0x96, 0xb8,0xe4 }
	//bytesBuffer := bytes.NewBuffer(b)
	////var x int32
	//var x rune//两种都可以
	//binary.Read(bytesBuffer, binary.BigEndian, &x)
	//fmt.Println(x)

	//int8类型的rune转[]byte类型
	var x int32
	//x = 97两种都可以
	x = 'a'
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	fmt.Println(bytesBuffer.Bytes())
	fmt.Println(rune(bytesBuffer.Bytes()[3]))


	//var x int32
	//x = '世'
	//bytesBuffer := bytes.NewBuffer([]byte{})
	//binary.Write(bytesBuffer, binary.BigEndian, x)
	//fmt.Println(bytesBuffer.Bytes())
	//by:=bytesBuffer.Bytes()
	//fmt.Println(by[2],by[3])
	//fmt.Printf("%T---%T\n",Dec2Bin(int64(by[2])),Dec2Bin(int64(by[3])))
	//fmt.Println(Dec2Bin(int64(by[2])),Dec2Bin(int64(by[3])))
	//str111:=Dec2Bin(int64(by[2])) + Dec2Bin(int64(by[3]))
	//fmt.Println(str111)
	////strconv.Atoi(str111)
	//str222:=Bin2Dec(str111)
	//fmt.Println(str222)
	//fmt.Printf("%T",str222)
	//fmt.Printf("%q\n",str222)
	//fmt.Println(string(str222))//string值得是编码int类型，但是不是将int字面值原样转成字符串

}




//256以下的十进制转二进制，并且在绝对保证8位形式的二进制格式
func Dec2Bin(n int64) string {
	if n < 0 {
		log.Println("Decimal to binary error: the argument must be greater than zero.")
		return ""
	}
	if n == 0 {
		return "0"
	}
	s := ""
	for q := n; q > 0; q = q / 2 {
		m := q % 2
		s = fmt.Sprintf("%v%v", m, s)
	}
	//------------------------------------------------
	//这里是我自己添加的。因此限制了能转换的十进制只能是256的最大值
	fmt.Println(s,len(s))
	if len(s)<8{
		length:=8-len(s)
		//编程误区，因为在for的里面修改了s的长度，同时我们每次循环都会去判断s的长度，导致循环的次数不对
		//所以我们要先获取到对象的长度然后再遍历
		for i:=0 ;i<length; i++ {
			s="0"+s
		}
	}
	//------------------------------------------------
	//fmt.Println("=====:",s)
	return s
}
//二进制转十进制
func Bin2Dec(b string) (n int64) {
	s := strings.Split(b, "")
	l := len(s)
	i := 0
	d := float64(0)
	for i = 0; i < l; i++ {
		f, err := strconv.ParseFloat(s[i], 10)
		if err != nil {
			log.Println("Binary to decimal error:", err.Error())
			return -1
		}
		d += f * math.Pow(2, float64(l-i-1))
	}
	return int64(d)
}








