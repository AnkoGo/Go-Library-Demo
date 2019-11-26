package main

import (
	"fmt"
	"time"
)

func main6893(){
	t661:=time.Date(2018,12,11,13,05,06,99,time.Local)
	fmt.Println("===========序列化后的结果=============")

	// MarshalBinary实现encoding.BinaryMarshaler接口。
	byte1, _ := t661.MarshalBinary()
	fmt.Println(byte1)
	fmt.Println(string(byte1))
	byte2, _ := t661.MarshalJSON()
	fmt.Println(byte2)
	fmt.Println(string(byte2))
	byte3, _ := t661.MarshalText()
	fmt.Println(byte3)
	fmt.Println(string(byte3))
	//输出如下：
	//	===========序列化后的结果=============
	//		[1 0 0 0 14 211 161 60 130 0 0 0 99 1 224]
	//	   ӡ<�   c�
	//	[34 50 48 49 56 45 49 50 45 49 49 84 49 51 58 48 53 58 48 54 46 48 48 48 48 48 48 48 57 57 43 48 56 58 48 48 34]
	//	"2018-12-11T13:05:06.000000099+08:00"
	//	[50 48 49 56 45 49 50 45 49 49 84 49 51 58 48 53 58 48 54 46 48 48 48 48 48 48 48 57 57 43 48 56 58 48 48]
	//	2018-12-11T13:05:06.000000099+08:00
	fmt.Println("===========================")

	var Dt=time.Date(2017,12,11,13,05,06,99,time.Local)//仅仅年份不同
	//var t time.Time//在有数据的time类实例里面序列化也是可以的
	fmt.Println(Dt.UnmarshalBinary(byte1))//<nil>，二进制原始字节
	fmt.Println(Dt.UnmarshalJSON(byte2))//<nil>，加引号的utf8编码自己，就是json了
	fmt.Println(Dt.UnmarshalText(byte3))//<nil>，utf8编码字节
	fmt.Println(byte1)
	fmt.Println(string(byte1))
	fmt.Println(byte2)
	fmt.Println(string(byte2))
	fmt.Println(byte3)
	fmt.Println(string(byte3))
	fmt.Println("===========反序列化后的结果=============")
	fmt.Println(Dt)
	fmt.Println(Dt.String())
	//输出如下：
	//	===========================
	//	<nil>
	//	<nil>
	//	<nil>
	//	[1 0 0 0 14 211 161 60 130 0 0 0 99 1 224]
	//	   ӡ<�   c�
	//	[34 50 48 49 56 45 49 50 45 49 49 84 49 51 58 48 53 58 48 54 46 48 48 48 48 48 48 48 57 57 43 48 56 58 48 48 34]
	//	"2018-12-11T13:05:06.000000099+08:00"
	//	[50 48 49 56 45 49 50 45 49 49 84 49 51 58 48 53 58 48 54 46 48 48 48 48 48 48 48 57 57 43 48 56 58 48 48]
	//	2018-12-11T13:05:06.000000099+08:00
	//	===========反序列化后的结果=============
	//	2018-12-11 13:05:06.000000099 +0800 CST
	//	2018-12-11 13:05:06.000000099 +0800 CST
}