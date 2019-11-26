package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// func main() {
// 	json1 := `[{"aa":"astr1"},{"aa":"bstr1"}]`
// 	var j1 []interface{}
// 	json.Unmarshal([]byte(json1), &j1)

// 	fmt.Println(j1)
// 	// vv := j1[0] //现在编译器认为他是interface {}类型而不是map类型，所以不能使用map的键来寻找相应的值，总之不能调用map的任何方法
// 	a11 := j1[0]
// 	vv, ok := a11.(map[string]interface{}) //为了使interface {}类型缩小范围，我们使用断言来转换类型为map,如果出错的话，那么我们将断言失败
// 	fmt.Println(ok)
// 	fmt.Println("==", vv["aa"])
// 	fmt.Println("----------1---------")

// 	vvv111, ok111 := vv["aa"].(string) //因为上面的断言已经使得interface {}类型转成了map类型，所以我们这时候可以使用map["aa"]来取值，
// 	// 但是注意我们上面断言的是map[string]interface{}类型（不能断言map[string]string类型，每次只能断言一个数据结构的类型） ，但是我们
// 	// 还没断言map里面的interface{}类型，因此我们这时候有需要再来一次对map的键值中的值的断言
// 	fmt.Println(ok111)
// 	fmt.Printf("%T\n", vvv111)
// 	fmt.Println("----------2---------")

// 	//以下是原型，上面是我的探究
// 	// for k, v := range j1 {
// 	// 	fmt.Println(k, reflect.TypeOf(v).String())

// 	// 	if v2, ok := v.(map[string]string); ok {
// 	// 		println("string断言成功")
// 	// 		println(v2["aa"])
// 	// 	} else if v2, ok := v.(map[string]interface{}); ok {
// 	// 		fmt.Println("键值为interface:")
// 	// 		fmt.Println(v2["aa"])
// 	// 	} else {
// 	// 		fmt.Println("失败:")
// 	// 	}
// 	// }

// 	//------------------------------------------------
// 	// a := 8
// 	// if a > 6 {
// 	// 	fmt.Println("大于6")
// 	// } else if a > 7 {//else if表示上面的条件成功的话则不再执行其他的if或者else
// 	// 	fmt.Println("大于7")
// 	// } else {
// 	// 	fmt.Println("小于等于6")
// 	// }
// }
