package main

import (
	// "fmt"
	"fmt"
	"reflect"
	// "time"
)

func main45893() {
	// start := time.Microsecond.Nanoseconds()
	// slice_dst := make([][]string, 0, 7)

	// str1 := "1"
	// str2 := "anko"
	// str := &[]string{str1, str2}
	// for i := 0; i < cap(slice_dst); i++ {
	// 	slice_dst = append(slice_dst, *str)
	// }
	// end := time.()
	// fmt.Printf("%+v----%v", slice_dst,end-start)
	// fmt.Println(reflect.TypeOf((interface{})(nil)).Kind())
	fmt.Println(reflect.TypeOf((*interface{})(nil)).Elem().Kind())

}
