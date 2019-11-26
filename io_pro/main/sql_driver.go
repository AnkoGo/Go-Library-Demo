package main

import (
	"database/sql/driver"
	"fmt"
)

func main() {

	//// Value is a value that drivers must be able to handle.
	//// It is either nil, a type handled by a database driver's NamedValueChecker
	//// interface, or an instance of one of these types:
	////
	////   int64
	////   float64
	////   bool
	////   []byte
	////   string
	////   time.Time
	////
	//// If the driver supports cursors, a returned Value may also implement the Rows interface
	//// in this package. This is used, for example, when a user selects a cursor
	//// such as "select cursor(select * from my_table) from dual". If the Rows
	//// from the select is closed, the cursor Rows will also be closed.
	//// 值是驱动程序必须能够处理的值。
	//// 它要么是nil，由数据库驱动程序的NamedValueChecker接口处理的类型，或者是以下类型之一的实例：
	////
	////   int64
	////   float64
	////   bool
	////   []byte
	////   string
	////   time.Time
	////
	//// 如果驱动程序支持游标，则返回的Value也可以在此包中实现Rows接口。
	//// 例如，当用户选择一个游标，例如“从双精度选择游标（从my_table中选择*）”时，将使用此功能。 如果选择中的行被关闭，则游标行也将被关闭。
	//type Value interface{}



	// IsValue reports whether v is a valid Value parameter type.
	// IsValue报告v是否为有效的type Value interface{}参数类型。
	var i int=5
	fmt.Println(driver.IsValue(i))
	var i1 int32=5
	fmt.Println(driver.IsValue(i1))
	var i2 int64=5
	fmt.Println(driver.IsValue(i2))

	var i3 bool=true
	fmt.Println(driver.IsValue(i3))
	//输出：
	//	false
	//	false
	//	true
	//	true
	//更多类型请自行测试

	// IsScanValue is equivalent to IsValue.
	// It exists for compatibility.
	// IsScanValue等效于IsValue。
	//它是为了兼容而存在。
	//底层完全等价于IsValue（）
	fmt.Println(driver.IsScanValue(i))
	fmt.Println(driver.IsScanValue(i1))
	fmt.Println(driver.IsScanValue(i2))
	fmt.Println(driver.IsScanValue(i3))
	//输出：
	//	false
	//	false
	//	true
	//	true
	//更多类型请自行测试


	//这整个包都是定义了接口而已，如果需要自己写实现的话会非常的复杂，如果需要探究的，请自行查看"github.com/go-sql-driver/mysql"包里面的实现
	//其实很多接口类型的用法都已经在sql包中讲到了



}



func check_err_sql(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
