package main
//
//import (
//	"fmt"
//	"github.com/axgle/mahonia"
//)
//
////src为要转换的字符串，srcCode为待转换的编码格式，targetCode为要转换的编码格式
//func ConvertToByte(src string, srcCode string, targetCode string) []byte {
//	srcCoder := mahonia.NewDecoder(srcCode)
//	srcResult := srcCoder.ConvertString(src)
//	tagCoder := mahonia.NewDecoder(targetCode)
//	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
//	return cdata
//}
//
//func main() {
//	//r := '世'
//
//	ls := []byte{34, 11, 8, 146, 147, 214, 236, 5, 16, 152, 148, 189, 4, 42, 0, 50, 10, 8, 12, 18, 6, 49, 50, 51, 52, 53, 54}
//	//对string由gbk转码为utf8
//	fmt.Println("******************")
//	response := ConvertToByte(string(ls), "gbk", "utf8")
//	fmt.Println(response)
//	fmt.Println(string(response))
//}
