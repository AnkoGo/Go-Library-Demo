//这整个文件都是为了说明sql.Register()的用法
package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"

	//_ "github.com/go-sql-driver/mysql" //一定要导入这个
)
/*

在进行下面的测试之前，我们必须将外置库\src\github.com\go-sql-driver\mysql\driver.go中的2行注释掉：
第一行是开头的导入包行：
import (
"context"
//"database/sql"//将这行注释掉
"database/sql/driver"
"net"
"sync"
)
第2行是本文件中的最后一个init()函数：
//func init() {
//	sql.Register("mysql", &MySQLDriver{})
//}

你可以看到上面的init函数其实跟我们要写的函数是一样的！因为如果我们要测试sql.Register（）函数的话则只能这样测试，这个函数
就是用来注册驱动的！如果你不注册的话，例如将我们这本文件中init()函数注释掉然后执行本文件的话会抛出以下的异常：
panic: sql: unknown driver "mysql" (forgotten import?)

goroutine 1 [running]:
main.main()
	C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:14 +0x473


其实报错就是这样报错：db, err := sql.Open("mysql", "root:mysql@/godb")，因为mysql这个驱动还没注册（注册不是指简单的指定名称
	即可，我们之所以不自己写注册驱动代码是因为目前来说还不会写（在后面会讲解到怎么写），所以既然有第三方写好的了，我们就拿来测试sql.Register（）函数的功能。）
*/
func main() {
	db, err := sql.Open("mysql111", "root:mysql@/godb")//这个mysql111必须跟下面的init函数中注册的name相同才可以
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Query("select * from person")
	check_err_sql(err)

	type person struct {
		id int
		name string
		age int
	}

	for b:=rows.Next(); b != false;b=rows.Next() {
		var p =person{}
		err := rows.Scan(&p.id, &p.name, &p.age)
		check_err_sql(err)
		fmt.Println(p)
	}
	//输出：
	//	{1 anko 20}
	//	{2 gogo 32}

}

func check_err_sql(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
//本init函数会在所有本文件的代码执行之前执行
func init() {
	fmt.Println("-------------------注册驱动---------------------------")

	//// Driver is the interface that must be implemented by a database
	//// driver.
	////
	//// Database drivers may implement DriverContext for access
	//// to contexts and to parse the name only once for a pool of connections,
	//// instead of once per connection.
	////驱动程序是必须由数据库驱动程序实现的接口。
	////数据库驱动程序可以实现DriverContext来访问上下文，并且对于连接池仅解析一次名称，而不是对每个连接解析一次。
	//type Driver interface {
	//	// Open returns a new connection to the database.
	//	// The name is a string in a driver-specific format.
	//	//
	//	// Open may return a cached connection (one previously
	//	// closed), but doing so is unnecessary; the sql package
	//	// maintains a pool of idle connections for efficient re-use.
	//	//
	//	// The returned connection is only used by one goroutine at a
	//	// time.
	//	// Open返回到数据库的新连接。
	//	//name是驱动程序特定格式的字符串。
	//	//Open可以返回一个缓存的连接（以前关闭了一个连接），但是没有必要这样做； sql软件包维护一个空闲连接池，以进行有效的重用。
	//	//返回的连接一次只能由一个goroutine使用。
	//	Open(name string) (Conn, error)
	//}


	//// If a Driver implements DriverContext, then sql.DB will call
	//// OpenConnector to obtain a Connector and then invoke
	//// that Connector's Conn method to obtain each needed connection,
	//// instead of invoking the Driver's Open method for each connection.
	//// The two-step sequence allows drivers to parse the name just once
	//// and also provides access to per-Conn contexts.
	////如果驱动程序实现了DriverContext，则sql.DB将调用OpenConnector以获得连接器，然后调用该连接器的Conn方法以获取每个所需的连接，而不是为每个连接调用驱动程序的Open方法。
	////两步序列允许驱动程序仅解析一次名称，还提供对每个Conn上下文的访问。
	//type DriverContext interface {
	//	// OpenConnector must parse the name in the same format that Driver.Open
	//	// parses the name parameter.
	//	// OpenConnector必须以与Driver.Open解析name参数相同的格式解析名称。
	//	OpenConnector(name string) (Connector, error)
	//}


	// Register makes a database driver available by the provided name.
	// If Register is called twice with the same name or if driver is nil,
	// it panics.
	// Register通过提供的名称name使数据库驱动程序driver可用。
	// 如果使用相同的名称两次调用Register或驱动程序为nil，则会出现紧急情况。
	sql.Register("mysql111", &mysql.MySQLDriver{})
	//sql.Register("mysql111", &mysql.MySQLDriver{})//重复注册相同的名称两次会报错
}