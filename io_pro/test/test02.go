package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"strings"
// )

// func main() {
// 	r := strings.NewReader("Hello, Reader!")

// 	bytes, e := ioutil.ReadAll(r)
// 	if e != nil {
// 		fmt.Println(e)
// 	}
// 	fmt.Println(bytes)
// 	fmt.Println(len(bytes), cap(bytes))

// 	bytes111, e111 := ioutil.ReadAll(r)
// 	if e111 != nil {
// 		fmt.Println(e111)
// 	}
// 	fmt.Println(bytes111)
// 	fmt.Println(len(bytes111), cap(bytes111))

// 	bytes222, e222 := ioutil.ReadAll(r) //只能读取一次，多次调用一律不再重新读
// 	if e222 != nil {
// 		fmt.Println(e222)
// 	}
// 	fmt.Println(bytes222)
// 	fmt.Println(len(bytes222), cap(bytes222))

// 	//b := make([]byte, 8)
// 	//for {
// 	//	n, err := r.Read(b)
// 	//	fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
// 	//	fmt.Printf("b[:n] = %q\n", b[:n])
// 	//	if err == io.EOF {
// 	//		break
// 	//	}
// 	//}
// }
