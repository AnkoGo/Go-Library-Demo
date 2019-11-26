package main

import (
	"encoding/json"
	"fmt"
)


//多重json的序列化和反序列化


func main() {

	////id：1 ，name：张三，parent：0；
	////id：2 ，name：李四，parent：1；
	////id：2 ，name：王五，parent：2；
	//str:=`
	//	[{id:1,
	//	name:张三,
	//	parent:0,
	//	children:[
	//				id:2 ,
	//				name:李四,
	//				parent:1,
	//				children:[{
	//						  id:3,
	//						  name:王五,
	//						  parent:2]
	//				]}
	//			]
	//	},]
	//`

	//------------------序列化-----------------------

	type dst_struct struct {
		Id int	`json:"id"`
		Name string	`json:"name"`
		Parent int	`json:"parent"`
		Children *dst_struct	`json:"children"`
	}

	var dst_obj2 =dst_struct{
		Id:       3,
		Name:     "王五",
		Parent:   2,
		Children: nil,
	}

	var dst_obj1 =dst_struct{
		Id:       2,
		Name:     "李四",
		Parent:   1,
		Children: &dst_obj2,
	}

	var dst_obj =dst_struct{
		Id:       1,
		Name:     "张三",
		Parent:   0,
		Children: &dst_obj1,
	}

	json_bytes, e := json.Marshal(&dst_obj)
	check_err(e)
	fmt.Printf("%v\n",json_bytes)
	fmt.Printf("%s\n",string(json_bytes))
	//输出：
	//{"id":1,
	//	"name":"张三",
	//	"parent":0,
	//	"children":{"id":2,
	//				"name":"李四",
	//				"parent":1,
	//				"children":{"id":3,
	//							"name":"王五",
	//							"parent":2,
	//							"children":null}}}



	//------------------反序列化-----------------------

	//方式1：
	//var Unmarshal_obj2 =dst_struct{
	//	Children:nil,
	//}
	//var Unmarshal_obj1 =dst_struct{
	//	Children:&Unmarshal_obj2,
	//}
	//
	//var Unmarshal_obj =dst_struct{
	//	Children:&Unmarshal_obj1,
	//}

	//方式2：
	//var Unmarshal_obj =dst_struct{
	//	Children:&dst_struct{
	//		Children:&dst_struct{
	//			Children:nil,
	//		},
	//	},
	//}

	//方式3：
	var Unmarshal_obj =dst_struct{

	}

	e = json.Unmarshal(json_bytes, &Unmarshal_obj)
	check_err(e)
	fmt.Printf("%#v\n",Unmarshal_obj)
	fmt.Printf("%#v\n",Unmarshal_obj.Children)
	fmt.Printf("%#v\n",Unmarshal_obj.Children.Children)
	//输出：
	//main.dst_struct{Id:1, Name:"张三", Parent:0, Children:(*main.dst_struct)(0xc00005c540)}
	//&main.dst_struct{Id:2, Name:"李四", Parent:1, Children:(*main.dst_struct)(0xc00005c510)}
	//&main.dst_struct{Id:3, Name:"王五", Parent:2, Children:(*main.dst_struct)(nil)}
}



func check_err(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

}
