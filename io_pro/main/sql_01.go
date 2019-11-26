package main

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql" //一定要导入这个,注册驱动的意思
	"strconv"
	"time"

	//记得在开始本文件的执行之前需要确保github.com/go-sql-driver/mysql里面的文件没被你手动更改过
)

func main() {
	db, err := sql.Open("mysql", "root:mysql@/godb")
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
	PirntQueryRows(rows)




	fmt.Println("-------------------------------------")
	// Drivers returns a sorted list of the names of the registered drivers.(//Drivers返回已注册驱动程序名称的排序列表。)
	fmt.Println(sql.Drivers())




	//// A NamedArg is a named argument. NamedArg values may be used as
	//// arguments to Query or Exec and bind to the corresponding named
	//// parameter in the SQL statement.
	////
	//// For a more concise way to create NamedArg values, see
	//// the Named function.
	//// NamedArg是一个命名参数。 NamedArg值可用作Query或Exec的参数，并绑定到SQL语句中相应的命名参数。
	////有关创建NamedArg值的更简洁方法，请参见Named函数。
	//type NamedArg struct {
	//	_Named_Fields_Required struct{}
	//
	//	// Name is the name of the parameter placeholder.
	//	//
	//	// If empty, the ordinal position in the argument list will be
	//	// used.
	//	//
	//	// Name must omit any symbol prefix.
	//	// Name是参数占位符的名称。
	//	//如果Name为空，将使用参数列表中的序数位置。
	//	//Name必须省略任何符号前缀
	//	Name string
	//
	//	// Value is the value of the parameter.
	//	// It may be assigned the same value types as the query
	//	// arguments.
	//	// Value是参数的值。
	//	//可以为其分配与查询参数相同的值类型。
	//	Value interface{}
	//}


	// Named provides a more concise way to create NamedArg values.
	// Named提供了一种更简洁的方法来创建NamedArg值。
	// Example usage:
	//
	//     db.ExecContext(ctx, `
	//         delete from Invoice
	//         where
	//             TimeCreated < @end
	//             and TimeCreated >= @start;`,
	//         sql.Named("start", startTime),
	//         sql.Named("end", endTime),
	//     )
	namedArg := sql.Named("name", "anko")
	fmt.Printf("%+v\n",namedArg)

	//不知道为什么只能填充一个问号
	//rows, err = db.Query("select * from person where id=? and id=?",2,1)
	rows, err = db.Query("select * from person where id=?",2)
	//go-sql-driver\mysql这个驱动不支持这种写法
	//输出：mysql: driver does not support the use of Named Parameters

	check_err_sql(err)
	PirntQueryRows(rows)
	//输出：
	//	[mysql]
	//	{_Named_Fields_Required:{} Name:name Value:anko}
	//	mysql: driver does not support the use of Named Parameters


	fmt.Println("---------------从连接池中拿出来一个连接进行操作----------------------")

	//// Conn represents a single database connection rather than a pool of database
	//// connections. Prefer running queries from DB unless there is a specific
	//// need for a continuous single database connection.
	////
	//// A Conn must call Close to return the connection to the database pool
	//// and may do so concurrently with a running query.
	////
	//// After a call to Close, all operations on the
	//// connection fail with ErrConnDone.
	//// Conn表示单个数据库连接，而不是数据库连接池。 除非特别需要连续的单个数据库连接，否则最好从DB运行查询。
	//// Conn必须调用Close才能返回到数据库池的连接，并且可以与正在运行的查询同时进行。
	////调用Close之后，对该连接的所有操作都会因ErrConnDone失败。
	//type Conn struct {
	// DB是表示零个或多个已经存在的连接池数据库句柄。 对于多个goroutine并发使用是安全的。
	// sql软件包自动创建并释放连接； 它还维护空闲连接的空闲池。 如果数据库具有按连接状态的概念，则可以在事务（Tx）或连接（Conn）中可靠地观察到这种状态。
	// 调用DB.Begin之后，返回的Tx将绑定到单个连接。 在事务上调用Commit或Rollback后，该事务的连接将返回到DB的空闲连接池。 池大小可以使用SetMaxIdleConns控制。
	// 说白了就是连接池句柄
	//	db *DB
	//
	//	// closemu prevents the connection from closing while there
	//	// is an active query. It is held for read during queries
	//	// and exclusively during close.
	//	// closemu防止在有活动查询时关闭连接。 只在连接查询queries期间和关闭连接close期间上锁。
	//	closemu sync.RWMutex
	//
	//	// dc is owned until close, at which point
	//	// it's returned to the connection pool.
	//	// 拥有驱动连接dc，直到关闭为止，这时它将返回到连接池。代表增删改操作附属在哪个连接上进行
	//	dc *driverConn
	//
	//	// done transitions from 0 to 1 exactly once, on close.
	//	// Once done, all operations fail with ErrConnDone.
	//	// Use atomic operations on value when checking value.
	//	// DONE在CLOSE上从0到1的转换仅一次。
	//	//一旦完成，所有操作都会因ErrConnDone失败。
	//	//检查值时对值使用原子操作。
	//	done int32
	//}



	// Conn returns a single connection by either opening a new connection
	// or returning an existing connection from the connection pool. Conn will
	// block until either a connection is returned or ctx is canceled.
	// Queries run on the same Conn will be run in the same database session.
	//
	// Every Conn must be returned to the database pool after use by
	// calling Conn.Close.
	// Conn通过打开新连接或从连接池返回现有连接来返回单个连接。 Conn将阻塞，直到返回连接或取消ctx。
	// 在同一Conn上运行的查询将在同一数据库会话中运行。
	// 使用后，必须通过调用Conn.Close将每个Conn返回到数据库池。
	// 说白了就是从连接池中get拿到一个连接对象出来,操作完后要close()来put放回连接池中去
	conn, err := db.Conn(context.Background())
	check_err_sql(err)
	background := context.Background()
	rows, err = conn.QueryContext(background, "select * from person")
	check_err_sql(err)
	PirntQueryRows(rows)

	// QueryRowContext executes a query that is expected to return at most one row.
	// QueryRowContext always returns a non-nil value. Errors are deferred until
	// Row's Scan method is called.
	// If the query selects no rows, the *Row's Scan will return ErrNoRows.
	// Otherwise, the *Row's Scan scans the first selected row and discards
	// the rest.
	// QueryRowContext执行的查询预期最多返回一行。
	// QueryRowContext始终返回非null值。 错误将一直延迟到调用Row的Scan方法。
	//如果查询未返回任何行row，则“行扫描*Row's Scan”将返回ErrNoRows。
	//否则，*行扫描*Row's Scan将扫描所选的第一行，并丢弃其余的行。
	row := conn.QueryRowContext(background, "select * from person")
	check_err_sql(err)
	PirntQueryRow(row)

	// Close returns the connection to the connection pool.
	// All operations after a Close will return with ErrConnDone.
	// Close is safe to call concurrently with other operations and will
	// block until all other operations finish. It may be useful to first
	// cancel any used context and then call close directly after.
	// Close将连接返回到连接池。
	// Close后的所有操作都将返回ErrConnDone。
	// Close可以安全地与其他操作同时调用，并且将阻塞直到所有其他操作完成。 最好首先取消任何已经在使用的上下文，然后再直接调用close。
	defer func() {
		//fmt.Println(background.Done())//这个上下文无法取消的！因为这个上下文一直没创建！是个占位的上下文而已
		conn.Close()
	}()

	// PrepareContext creates a prepared statement for later queries or executions.
	// Multiple queries or executions may be run concurrently from the
	// returned statement.
	// The caller must call the statement's Close method
	// when the statement is no longer needed.
	//
	// The provided context is used for the preparation of the statement, not for the
	// execution of the statement.
	// PrepareContext为以后的查询或执行创建准备好的语句。
	//可以从返回的语句statement中同时运行多个查询或执行。
	//当不再需要该语句statement时，调用者必须调用该语句的Close方法。
	//提供的上下文用于语句statement的准备，而不是语句的执行。
	//下面的stmt是statement的简写,为什么要生成语句statement呢？

	//sql.Stmt支持预备表达式，可以用来优化SQL查询提高性能，减少SQL注入的风险, DB.Prepare()和Tx.Prepare()都提供了对于预备表达式的支持。
	//
	//预处理的流程:
	//	step1. 将sql分为2部分.命令部分和数据部分.
	//	step2. 首先将命令部分发送给mysql服务器,mysql进行预处理.(如生成AST)
	//	step3. 然后将数据部分发送给mysql服务器,mysql进行占位符替换.
	//	step4. mysql服务器执行sql语句,把执行结果发送给客户端.
	//
	//预处理的优势:
	//	1.因为发送命令后,在mysql服务器端,就会将AST生成好,所以不需要对每一次值的更换都重新生成一次AST.对同样的数据不同的SQL来讲,只需生成1次AST,并缓存起来即可.
	//	2.避免SQL注入.因为mysql知道再次发送过来的内容为”数据”,因此不会将这些数据解析为SQL,避免了SQL注入.
	//
	//需要注意的点:
	//	使用预处理进行查询操作时,不仅在defer时需要关闭结果集,而且还要关闭命令句柄,否则同样会占用连接,导致阻塞.



	//// Stmt is a prepared statement.
	//// A Stmt is safe for concurrent use by multiple goroutines.
	////
	//// If a Stmt is prepared on a Tx or Conn, it will be bound to a single
	//// underlying connection forever. If the Tx or Conn closes, the Stmt will
	//// become unusable and all operations will return an error.
	//// If a Stmt is prepared on a DB, it will remain usable for the lifetime of the
	//// DB. When the Stmt needs to execute on a new underlying connection, it will
	//// prepare itself on the new connection automatically.
	//// Stmt是一个准备好的语句statement。
	//// 一个Stmt可安全地供多个goroutine并发使用。
	//// 如果在Tx或Conn上prepared准备了Stmt，它将永远绑定到单个基础连接。 如果Tx或Conn关闭，则Stmt将变得不可用，并且所有操作都将返回错误。
	//// 如果在数据库DB上prepared准备了Stmt，它将在数据库DB的整个生命周期内保持可用状态。 当Stmt需要在新的基础连接上执行时，它将自动在新的连接上进行准备。
	//type Stmt struct {
	//	// Immutable:(//不可变：)
	//	db        *DB    // where we came from(//我们来自哪里,指明附属者)
	//	query     string // that created the Stmt（//创造了Stmt的操作字符串）
	//	stickyErr error  // if non-nil, this error is returned for all operations（//如果非nil，则为所有操作返回此错误）
	//
	//	closemu sync.RWMutex // held exclusively during close, for read otherwise.（//仅在关闭期间上锁，其他时候都不会上锁且允许被获取锁）
	//
	//	// If Stmt is prepared on a Tx or Conn then cg is present and will
	//	// only ever grab a connection from cg.
	//	// If cg is nil then the Stmt must grab an arbitrary connection
	//	// from db and determine if it must prepare the stmt again by
	//	// inspecting css.
	//	//如果在Tx或Conn上准备了Stmt，则cg被赋值（不为nil），并且当前只会从cg获取连接。
	//	//如果cg为nil，则Stmt必须从db获取任一连接，并通过检查css确定是否必须再次准备stmt。
	//	cg   stmtConnGrabber
	//	cgds *driverStmt
	//
	//	// parentStmt is set when a transaction-specific statement
	//	// is requested from an identical statement prepared on the same
	//	// conn. parentStmt is used to track the dependency of this statement
	//	// on its originating ("parent") statement so that parentStmt may
	//	// be closed by the user without them having to know whether or not
	//	// any transactions are still using it.
	//	// 当从同一conn上准备的同一语句statement中请求特定事务的语句statement时，将设置parentStmt。
	//	// parentStmt用于跟踪此语句statement对其原始（“父”）语句statement的依赖性，以便用户可以关闭parentStmt，而不必知道是否有任何事务在使用它。
	//	parentStmt *Stmt
	//
	//	mu     sync.Mutex // protects the rest of the fields（//保护其余字段）
	//	closed bool		//此语句关闭与否
	//
	//	// css is a list of underlying driver statement interfaces
	//	// that are valid on particular connections. This is only
	//	// used if cg == nil and one is found that has idle
	//	// connections. If cg != nil, cgds is always used.
	//	// css是在特定连接上有效的基础驱动程序语句接口的列表。 仅当cg == nil并且发现有空闲连接时使用。 如果cg！= nil，则始终使用cgds。
	//	css []connStmt
	//
	//	// lastNumClosed is copied from db.numClosed when Stmt is created
	//	// without tx and closed connections in css are removed.
	//	// 创建不带tx的Stmt并从CSS中关闭的连接删除时，从db.numClosed复制lastNumClosed。
	//	lastNumClosed uint64
	//}

	fmt.Println("******")
	stmt,err := conn.PrepareContext(background,"select * from person where id=? or name=?")
	check_err_sql(err)

	fmt.Printf("%+v\n",stmt)

	id := 2
	name:="anko"
	// Query executes a prepared query statement with the given arguments
	// and returns the query results as a *Rows.
	//Query使用给定的参数执行准备好的查询语句，并将查询结果作为* Rows返回。

	//rows, err = stmt.Query(name, id)//顺序千万不能反了，否则什么错误都不会返回
	rows, err = stmt.Query(id, name)
	check_err_sql(err)
	PirntQueryRows(rows)

	fmt.Println("******")
	rows, err = stmt.QueryContext(background,id, name)
	check_err_sql(err)
	PirntQueryRows(rows)


	fmt.Println("******")
	// QueryRowContext executes a prepared query statement with the given arguments.
	// If an error occurs during the execution of the statement, that error will
	// be returned by a call to Scan on the returned *Row, which is always non-nil.
	// If the query selects no rows, the *Row's Scan will return ErrNoRows.
	// Otherwise, the *Row's Scan scans the first selected row and discards
	// the rest.
	// QueryRowContext使用给定的参数执行准备好的查询语句。
	//如果在执行语句期间发生错误，则该错误将通过在返回的* Row上调用Scan来返回，该行始终为nil。
	//如果查询未选择任何行，则“行扫描”将返回ErrNoRows。
	//否则，*行扫描将扫描所选的第一行，并丢弃其余的行。
	row= stmt.QueryRowContext(background, id, name)
	PirntQueryRow(row)


	// QueryRow executes a prepared query statement with the given arguments.
	// If an error occurs during the execution of the statement, that error will
	// be returned by a call to Scan on the returned *Row, which is always non-nil.
	// If the query selects no rows, the *Row's Scan will return ErrNoRows.
	// Otherwise, the *Row's Scan scans the first selected row and discards
	// the rest.
	//
	// Example usage:
	//
	//  var name string
	//  err := nameByUseridStmt.QueryRow(id).Scan(&name)
	// QueryRow使用给定的参数执行准备好的查询语句。
	//如果在执行语句期间发生错误，则该错误将通过在返回的* Row上调用Scan来返回，该行始终为nil。
	//如果查询未选择任何行，则“行扫描”将返回ErrNoRows。
	//否则，*行扫描将扫描所选的第一行，并丢弃其余的行。
	//
	//示例用法：
	//
	//  var name string
	//  err := nameByUseridStmt.QueryRow(id).Scan(&name)

	fmt.Println("******")
	row= stmt.QueryRow(id, name)
	PirntQueryRow(row)
	//输出：
	//{1 anko 20}
	//{2 gogo 32}
	//{1 anko 20}
	//******
	//&{db:0xc000094000
	//	query:select * from person where id=? or name=?
	//	stickyErr:<nil>
	//	closemu:{w:{state:0 sema:0} writerSem:0 readerSem:0 readerCount:0 readerWait:0}
	//	cg:0xc0000c2000
	//	cgds:0xc0000d4000
	//	parentStmt:<nil>
	//	mu:{state:0 sema:0}
	//	closed:false
	//	css:[]
	//	lastNumClosed:0}
	//{1 anko 20}
	//{2 gogo 32}
	//******
	//{1 anko 20}
	//{2 gogo 32}
	//&{db:0xc000094000
	//	query:select * from person where id=? or name=?
	//	stickyErr:<nil>
	//	closemu:{w:{state:0 sema:0} writerSem:0 readerSem:0 readerCount:0 readerWait:0}
	//	cg:0xc0000b2120
	//	cgds:<nil>
	//	parentStmt:<nil>
	//	mu:{state:0 sema:0}
	//	closed:true		这里跟上面会形成对比
	//	css:[]
	//	lastNumClosed:0}
	//	******
	//	{1 anko 20}
	//	******
	//	{1 anko 20}	说明返回的数据是有序的！




	fmt.Println("========")
	//// A Result summarizes an executed SQL command.(//结果汇总了已执行(准确来说是“已整合”，而不是执行)的SQL命令。)
	//type Result interface {
	//	// LastInsertId returns the integer generated by the database
	//	// in response to a command. Typically this will be from an
	//	// "auto increment" column when inserting a new row. Not all
	//	// databases support this feature, and the syntax of such
	//	// statements varies.
	//	// LastInsertId返回数据库为响应命令而生成的整数。 通常，这将在插入新行时来自“自动增量”列数（代表第几行，而不是影响了多少行或者列，下同）。
	//	// 并非所有数据库都支持此功能，并且此类语句的语法也有所不同。
	//	// 注意这个操作仅限insert操作才不会为0
	//	LastInsertId() (int64, error)
	//
	//	// RowsAffected returns the number of rows affected by an
	//	// update, insert, or delete. Not every database or database
	//	// driver may support this.
	//	// RowsAffected返回受更新，插入或删除影响的行数(这里跟上面的列数不同，这里指的是受影响的行数，而不是第几行或者第几列)。 并非每个数据库或数据库驱动程序都可以支持此功能。
	//	RowsAffected() (int64, error)
	//}


	//// Exec executes a prepared statement with the given arguments and
	//// returns a Result summarizing the effect of the statement.
	//// Exec用给定的参数执行一个准备好的语句，并返回一个受到该总结性语句影响的结果Result对象（该接口包含执行sql命令后
	//// 的受影响的行数和列数,如果是查询的sql语句的话则不会产生影响数据库的行为，即返回影响的行数和列数都为0，但是如果执行的语句是
	//// 增删改sql，则返回的影响的行数和列数均不为0）。
	//stmt,err = conn.PrepareContext(background,"insert into person(name,age) values(?,?)")
	//check_err_sql(err)
	//fmt.Printf("%+v\n",stmt)
	//
	//Sperson := []struct {
	//	name  string
	//	age int
	//}{
	//	{"cpp", 23},
	//	{"java", 24},
	//	{"go", 23},
	//	{"python", 56},
	//}
	//
	//
	//for id, person := range Sperson {
	//	fmt.Printf("---%v---",id)
	//	result, err := stmt.Exec(person.name, person.age)
	//	check_err_sql(err)
	//	fmt.Println(result)
	//	fmt.Println(result.LastInsertId())
	//	fmt.Println(result.RowsAffected())
	//}
	//
	////数据库中的数据如下：
	////	mysql> use godb
	////		Database changed
	////	mysql> show tables;
	////		+----------------+
	////		| Tables_in_godb |
	////		+----------------+
	////		| person         |
	////		+----------------+
	////		1 row in set (0.01 sec)
	////
	////	mysql> select * from person;
	////		+----+--------+-----+
	////		| id | name   | age |
	////			+----+--------+-----+
	////		|  1 | anko   |  20 |
	////		|  2 | gogo   |  32 |，这2行原本就有的！下面才是新添加的！
	////		|  3 | cpp    |  23 |
	////		|  4 | java   |  24 |
	////		|  5 | go     |  23 |
	////		|  6 | python |  56 |
	////		+----+--------+-----+
	////		6 rows in set (0.00 sec)


	//输出如下：
	//	========
	//	&{db:0xc000094000
	//		query:insert into person(name,age) values(?,?)
	//		stickyErr:<nil>
	//		closemu:{w:{state:0 sema:0} writerSem:0 readerSem:0 readerCount:0 readerWait:0}
	//		cg:0xc0000b20f0
	//		cgds:0xc0000dc040
	//		parentStmt:<nil>
	//		mu:{state:0 sema:0}
	//		closed:false
	//		css:[]
	//		lastNumClosed:0}
	//	---0---{0xc0000b0000 0xc0000ce2b0}
	//	3 <nil>
	//	1 <nil>
	//	---1---{0xc0000b0000 0xc0000ce2e0}
	//	4 <nil>
	//	1 <nil>
	//	---2---{0xc0000b0000 0xc0000ce320}
	//	5 <nil>
	//	1 <nil>
	//	---3---{0xc0000b0000 0xc0000ce350}
	//	6 <nil>
	//	1 <nil>



	//在开始之前记得将上面的插入操作代码全部注释掉，否则会再次插入
	//不允许插入2行，不允许sql语句中有;号
	//stmt,err = conn.PrepareContext(background,`insert into person(name,age) values(?,?);insert into person(name,age) values(?,?)`)
	//stmt,err = conn.PrepareContext(background,`UPDATE person SET name = ?, age = ? WHERE id = ?`)//语句1
	stmt,err = conn.PrepareContext(background,`UPDATE person SET name = ?, age = ? WHERE id = ? or id = ?`)//语句2
	check_err_sql(err)
	fmt.Printf("%+v\n",stmt)

	result, err := stmt.Exec("Zhongshan", 44,4,5)
	check_err_sql(err)
	fmt.Println(result)
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
	//语句1输出：
	//	&{db:0xc000094000 query:UPDATE person SET name = ?, age = ? WHERE id = ? stickyErr:<nil> closemu:{w:{state:0 sema:0} writerSem:0 readerSem:0 readerCount:0 readerWait:0} cg:0xc0000b2120 cgds:0xc0000be080 parentStmt:<nil> mu:{state:0 sema:0} closed:false css:[] lastNumClosed:0}
	//	{0xc00009c100 0xc0000ae350}
	//	0 <nil>
	//	1 <nil>

	//语句2输出：
	//	&{db:0xc0000ae000 query:UPDATE person SET name = ?, age = ? WHERE id = ? or id = ? stickyErr:<nil> closemu:{w:{state:0 sema:0} writerSem:0 readerSem:0 readerCount:0 readerWait:0} cg:0xc00007a930 cgds:0xc000048100 parentStmt:<nil> mu:{state:0 sema:0} closed:false css:[] lastNumClosed:0}
	//	{0xc0000b2100 0xc000058580}
	//	0 <nil>
	//	2 <nil>
	//特别注意的是，当数据已经是name=Zhongshan,age=44的时候再进行上面update操作的话不会更改任何东西，相应的行列数会返回如下：
	//	0 <nil>
	//	0 <nil>

	//此时数据库数据如下：
	//	mysql> select * from person;
	//		+----+-----------+-----+
	//		| id | name      | age |
	//			+----+-----------+-----+
	//		|  1 | anko      |  20 |
	//		|  2 | gogo      |  32 |
	//		|  3 | cpp       |  23 |
	//		|  4 | Zhongshan |  44 |
	//		|  5 | Zhongshan |  44 |
	//		|  6 | Zhongshan |  44 |
	//		+----+-----------+-----+
	//		6 rows in set (0.00 sec)

	//stmt.ExecContext()这个方法跟上面几乎一样，不再累叙




	fmt.Println("---------------从连接池中拿出来一个连接进行操作----------------------")
	// ExecContext executes a query without returning any rows.
	// The args are for any placeholder parameters in the query.
	// ExecContext执行查询而不返回任何行。
	// args用于查询中的任何占位符参数。
	// 返回一个影响的行列数result对象
	result,err = conn.ExecContext(background,`UPDATE person SET name = ? WHERE id = ? or id = ?`,"python",4,5)//语句2
	check_err_sql(err)
	fmt.Println(result)
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
	//操作后的数据库为：
	//	mysql> select * from person;
	//		+----+-----------+-----+
	//		| id | name      | age |
	//		+----+-----------+-----+
	//		|  1 | anko      |  20 |
	//		|  2 | gogo      |  32 |
	//		|  3 | cpp       |  23 |
	//		|  4 | python    |  44 |
	//		|  5 | python    |  44 |
	//		|  6 | Zhongshan |  44 |
	//		+----+-----------+-----+
	//		6 rows in set (0.00 sec)

	//控制台输出：
	//	{0xc0000cc080 0xc0000d8350}
	//	0 <nil>
	//	2 <nil>


	// QueryContext执行一个查询，该查询返回行，通常是SELECT。
	// args用于查询中的任何占位符参数。
	rows, err = conn.QueryContext(background, "select * from person where id=? or id=?", 4, 5)
	check_err_sql(err)
	PirntQueryRows(rows)
	//输出：
	//	{4 python 44}
	//	{5 python 44}

	// QueryRowContext执行的查询预期最多返回一行。
	// QueryRowContext始终返回非null值。 错误将一直延迟到调用Row的Scan方法。
	//如果查询未选择任何行，则“行扫描”将返回ErrNoRows。
	//否则，*行扫描将扫描所选的第一行，并丢弃其余的行。
	//底层其实也是调用了 conn.QueryContext（）,然后返回&Row{rows: rows, err: err}即可
	row= conn.QueryRowContext(background, "select * from person where id=? or id=?", 4, 5)

	PirntQueryRow(row)
	//输出：
	//	{4 python 44}




	fmt.Println("----------")

	//// driverConn wraps a driver.Conn with a mutex, to
	//// be held during all calls into the Conn. (including any calls onto
	//// interfaces returned via that Conn, such as calls on Tx, Stmt,
	//// Result, Rows)
	//// driverConn在要调用Conn的所有过程中用一个互斥量mutex包装driver.Conn（包括通过该Conn返回的接口的任何调用，例如Tx，Stmt，Result，Row的调用）
	//type driverConn struct {
	//	db        *DB	//连接池对象
	//	createdAt time.Time		//创建时间
	//
	//	sync.Mutex  // guards following	（//保护以下的数据（上锁，确保同一时刻只有一个g程操作这个driverConn对象））
	//	ci          driver.Conn		//被包装的driver.Conn对象
	//	closed      bool			//连接关闭与否
	//	finalClosed bool // ci.Close has been called（//driver.Conn.Close是否被调用了）
	//	openStmt    map[*driverStmt]bool	//开放的允许被执行的Stmt语句对象的map
	//	lastErr     error // lastError captures the result of the session resetter.（// lastError捕获会话重置器的结果。表示最近出现的错误）
	//
	//	// guarded by db.mu	（//通过db.mu保证以下的数据上锁）
	//	inUse      bool		//当前连接ci正在使用与否
	//	onPut      []func() // code (with db.mu held) run when conn is next returned（//下次返回conn时，运行db.mu的代码，相当于回调）
	//	dbmuClosed bool     // same as closed, but guarded by db.mu, for removeClosedStmtLocked（//与关闭相同，但受db.mu保护，用于已经关闭的stmt语句的锁removeClosedStmtLocked）
	//}



	//// Conn is a connection to a database. It is not used concurrently
	//// by multiple goroutines.
	////
	//// Conn is assumed to be stateful.
	//// Conn是到数据库的连接。 多个goroutine不能同时使用它。
	////假设Conn是有状态的。
	//type Conn interface {
	//	// Prepare returns a prepared statement, bound to this connection.
	//	// Prepare返回绑定到此连接的一条预处理语句Stmt。
	//	Prepare(query string) (Stmt, error)
	//
	//	// Close invalidates and potentially stops any current
	//	// prepared statements and transactions, marking this
	//	// connection as no longer in use.
	//	//
	//	// Because the sql package maintains a free pool of
	//	// connections and only calls Close when there's a surplus of
	//	// idle connections, it shouldn't be necessary for drivers to
	//	// do their own connection caching.
	//	// Close使该连接无效并有可能停止任何当前准备的语句和事务，从而将该连接标记为不再使用。
	//	//因为sql程序包维护一个空闲的连接池，并且仅在有多余的空闲连接时才调用Close，所以驱动程序不必自己进行连接缓存。
	//	Close() error
	//
	//	// Begin starts and returns a new transaction.
	//	//
	//	// Deprecated: Drivers should implement ConnBeginTx instead (or additionally).
	//	//Begin开始并返回一个新的事务。
	//	//不推荐使用：驱动程序应改为（或另外）实现ConnBeginTx。
	//	Begin() (Tx, error)
	//}



	// Raw executes f exposing the underlying driver connection for the
	// duration of f. The driverConn must not be used outside of f.
	//
	// Once f returns and err is nil, the Conn will continue to be usable
	// until Conn.Close is called.
	// Raw执行f，在f函数执行的这段持续时间内才暴露公开基础驱动程序连接driverConn（即拥有f的第一个参数ci字段的结构体，同时他是conn的dc字段）。
	// 不能在f之外使用driverConn。Raw（）方法结束 后会关闭并且释放这个driverConn
	// 一旦f返回且err为nil，则Conn将继续可用，直到调用Conn.Close。如果f中抛出异常，则该driverConn会被关闭，因为driverConn封装了Conn，也就是Conn将不可用
	err=conn.Raw(func(driverConn interface{}) error{
		dc := driverConn.(driver.Conn)
		stmt, e := dc.Prepare("select * from person where id =? or id=?")
		//除了dc.Prepare，还有dc.Begin(),我不打算在这里讲这个方法，下面会讲到

		check_err_sql(e)
		if e !=nil{
			panic(e)
		}

		//下面不能为int类型，必须是int64
		//仅支持以下类型才可以转化为driver.Value类型：
		//   int64
		//   float64
		//   bool
		//   []byte
		//   string
		//   time.Time
		rows, e := stmt.Query(([]driver.Value{int64(4), int64(5)}))//这个rows是driver.rows,不是sql.rows，故不能用下面的PirntQueryRows（）方法,我们在下面新建了一个新的方法
		// NumInput返回占位符参数的数量。
		//如果NumInput返回> = 0，则sql程序包将在调用语句的Exec或Query方法之前从调用方进行完整性检查参数计数，并将错误返回给调用方。
		//如果驱动程序不知道其占位符数量，则NumInput也可能返回-1。 在这种情况下，sql程序包将不会检查Exec或Query参数计数。
		fmt.Println("stmt中的占位符的数量为：",stmt.NumInput())
		//除了上面几个方法之外，还有stmt.Exec(),不再累叙
		defer stmt.Close()

		check_err_sql(e)
		if e !=nil{
			panic(e)
		}

		e = PirntQueryDriverRows(rows)
		if e !=nil{
			panic(e)
		}
		return nil
	} )
	check_err_sql(err)
	//输出：
	//	stmt中的占位符的数量为： 2
	//	[id name age]
	//	[]driver.Value{4, []uint8{0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e}, 44}
	//	[4 [112 121 116 104 111 110] 44]
	//	[%!s(int64=4) python %!s(int64=44)]
	//	4python44
	//	-------
	//	[]driver.Value{5, []uint8{0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e}, 44}
	//	[5 [112 121 116 104 111 110] 44]
	//	[%!s(int64=5) python %!s(int64=44)]
	//	5python44

	fmt.Println("-------探究是否关闭了conn---------")
	rows, err = conn.QueryContext(background, "select * from person")
	check_err_sql(err)
	PirntQueryRows(rows)
	//输出：
	//	{1 anko 20}
	//	{2 gogo 32}
	//	{3 cpp 23}
	//	{4 python 44}
	//	{5 python 44}
	//	{6 Zhongshan 44}
	//从上可知，我们只是关闭了driverConn，但不是关闭conn




	fmt.Println("---------------------验证连接是否有效------------------------")

	//// Pinger is an optional interface that may be implemented by a Conn.
	////
	//// If a Conn does not implement Pinger, the sql package's DB.Ping and
	//// DB.PingContext will check if there is at least one Conn available.
	////
	//// If Conn.Ping returns ErrBadConn, DB.Ping and DB.PingContext will remove
	//// the Conn from pool.
	//// Pinger是Conn可以实现的可选接口。
	//// 如果Conn未实现Pinger，则sql包的DB.Ping和DB.PingContext将检查是否至少有一个Conn可用。
	//// 如果Conn.Ping返回ErrBadConn，则DB.Ping和DB.PingContext将从池中删除Conn。
	//type Pinger interface {
	//	Ping(ctx context.Context) error
	//}

	// PingContext verifies the connection to the database is still alive.
	// PingContext验证与数据库的连接是否仍然有效。如果有效则直接建立连接
	// 这个方法和*DB.ping()方法的底层一样都是调用了func (db *DB) PingContext(ctx context.Context) error{}方法
	// 返回值是上面接口Pinger的error
	err = conn.PingContext(background)
	check_err_sql(err)

	db_test, err := sql.Open("mysql", "root:mysql111@/godb")
	check_err_sql(err)
	//这个方法会尝试进行连接，如果连接出错，会抛出异常！因此在下面的PingContext会直接抛出异常，conn_test为nil,
	// nil.PingContext()肯定会抛出异常的！
	// 事实上Conn也是抛出了异常，不过我们在这里处理了
	test:= func() {
		conn_test, err := db_test.Conn(context.Background())
		if err != nil{
			fmt.Println("===1====",err)
		}

		fmt.Println("111111")
		time.Sleep(1e9)
		err = conn_test.PingContext(background)
		if err != nil{
			fmt.Println("===2====",err)
		}
	}
	fmt.Println(test)
	//test()
	//输出：
	//	0x5f1f70
	//	===1==== Error 1045: Access denied for user 'root'@'localhost' (using password: YES)
	//	111111
	//	panic: runtime error: invalid memory address or nil pointer dereference
	//	[signal 0xc0000005 code=0x0 addr=0x28 pc=0x4d7f7d]
	//
	//	goroutine 1 [running]:
	//	database/sql.(*Conn).grabConn(0x0, 0x6b7580, 0xc000056010, 0xc00003c200, 0xc000028000, 0x8144a0, 0x8144a0)
	//	C:/Go/src/database/sql/sql.go:1798 +0x2d
	//	database/sql.(*Conn).PingContext(0x0, 0x6b7580, 0xc000056010, 0x1, 0x1)
	//	C:/Go/src/database/sql/sql.go:1807 +0x4a
	//	main.main()
	//	C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:752 +0x2022




	fmt.Println("-------------------事务------------------------")
	//我们将上面的test()注释掉才可以进行下面的步骤


	//// TxOptions holds the transaction options to be used in DB.BeginTx.
	//// TxOptions保留要在DB.BeginTx中使用的事务选项。
	//type TxOptions struct {
	//	// Isolation is the transaction isolation level.
	//	// If zero, the driver or database's default level is used.
	//	//Isolation隔离是事务隔离级别。
	//	//如果为零值，则使用驱动程序或数据库的默认级别。
	//	Isolation IsolationLevel
	//	ReadOnly  bool	//是否可读
	//}
	////IsolationLevel对象如下：
	//// IsolationLevel is the transaction isolation level used in TxOptions.
	//// IsolationLevel是TxOptions中使用的事务隔离级别。
	//type IsolationLevel int
	//
	//// Various isolation levels that drivers may support in BeginTx.
	//// If a driver does not support a given isolation level an error may be returned.
	////
	//// See https://en.wikipedia.org/wiki/Isolation_(database_systems)#Isolation_levels.
	////驱动程序可以在BeginTx中支持的各种隔离级别。
	////如果驱动程序不支持给定的隔离级别，则可能会返回错误。
	////参见https://en.wikipedia.org/wiki/Isolation_(database_systems)#Isolation_levels。
	//const (
	//	LevelDefault IsolationLevel = iota
	//	LevelReadUncommitted
	//	LevelReadCommitted
	//	LevelWriteCommitted
	//	LevelRepeatableRead
	//	LevelSnapshot
	//	LevelSerializable
	//	LevelLinearizable
	//)


	// BeginTx starts a transaction.
	//
	// The provided context is used until the transaction is committed or rolled back.
	// If the context is canceled, the sql package will roll back
	// the transaction. Tx.Commit will return an error if the context provided to
	// BeginTx is canceled.
	//
	// The provided TxOptions is optional and may be nil if defaults should be used.
	// If a non-default isolation level is used that the driver doesn't support,
	// an error will be returned.
	// BeginTx开始一个事务。
	// 提供的上下文将一直使用，直到事务被提交或回滚为止。
	// 如果取消上下文，则sql包将回滚该事务。 如果提供给BeginTx的上下文被取消，Tx.Commit将返回错误。
	// 提供的TxOptions是可选的，如果使用默认值，则提供的TxOptions可能为nil。
	// 如果使用了驱动程序不支持的非默认隔离级别，则将返回错误。

	//如果是不想&sql.TxOptions给全字段值，那么就应该使用关键字的形式赋值
	tx, err := conn.BeginTx(background, &sql.TxOptions {Isolation:sql.LevelDefault})
	check_err_sql(err)
	//rows_tx, err := tx.Query(`select * from person where id=? or id=?`, 4, 5)

	result_tx, err := tx.Exec(`update person set name = "anko",age=11 where id=? or id=?`, 4, 5)
	check_err_sql(err)
	if err !=nil{
		//回滚将中止事务。
		err := tx.Rollback()
		check_err_sql(err)
	}
	//还有以下的方法不再演示！同理的！最后的几个方法将会在下面讲解
	//tx.QueryRowContext()
	//tx.QueryRow()
	//tx.QueryContext()
	//tx.ExecContext()
	//tx.PrepareContext()
	//tx.Prepare()
	//tx.Stmt()
	//tx.StmtContext()


	defer rows.Close()

	check_err_sql(err)
	//PirntQueryRows(rows_tx)//语句1的代码

	//下面2行是语句2 的代码
	fmt.Println(result_tx.LastInsertId())
	fmt.Println(result_tx.RowsAffected())


	// Commit commits the transaction.
	// Commit提交事务。
	//必须在PirntQueryRows（）的下面
	err = tx.Commit()
	check_err_sql(err)

	//语句1输出：
	//	{4 python 44}
	//	{5 python 44}
	//语句2输出：
	//	0 <nil>
	//	2 <nil>



	fmt.Println("---------------PingContext----------------------")

	////故意写错账号
	//db1, err := sql.Open("mysql", "root111:mysql@/godb")
	//if err != nil {
	//	panic(err.Error())
	//}
	//
	//// PingContext验证与数据库的连接是否仍然存在，并在必要时建立连接。
	//// 这其实是db1.Ping()的底层实现
	//err= db1.PingContext(context.Background())
	//check_err_sql(err)
	////输出：
	////Error 1045: Access denied for user 'root111'@'localhost' (using password: YES)



	fmt.Println("---------------PingContext----------------------")

	db2, err := sql.Open("mysql", "root:mysql@/godb")
	if err != nil {
		panic(err.Error())
	}

	// BeginTx starts a transaction.
	//
	// The provided context is used until the transaction is committed or rolled back.
	// If the context is canceled, the sql package will roll back
	// the transaction. Tx.Commit will return an error if the context provided to
	// BeginTx is canceled.
	//
	// The provided TxOptions is optional and may be nil if defaults should be used.
	// If a non-default isolation level is used that the driver doesn't support,
	// an error will be returned.
	// BeginTx开始事务。
	//提供的上下文将一直使用，直到事务被提交或回滚为止。
	//如果取消上下文，则sql包将回滚该事务。 如果提供给BeginTx的上下文被取消，Tx.Commit将返回错误。
	//提供的TxOptions是可选的，如果应使用默认值，则可以为nil。
	//如果使用了驱动程序不支持的非默认隔离级别，则将返回错误。
	tx,err= db2.BeginTx(context.Background(),&sql.TxOptions{Isolation: sql.LevelDefault})
	check_err_sql(err)

	//故意用不存在的键age111,好让事务回滚
	//如果我给定错误或者不存在的id的话，程序不会回滚事务，而是无任何响应
	//result, err = tx.ExecContext(context.Background(), `update person set name = "anko",age=22 where id=? or id=?`, 77, 88)
	result, err = tx.ExecContext(context.Background(), `update person set name = "anko",age111=22 where id=? or id=?`, 4, 5)
	check_err_sql(err)
	if err != nil {
		fmt.Println("准备回滚事务。。。")
		err := tx.Rollback()
		if err !=nil{
			fmt.Println("(===1===)",err)
		}
	}else {
		fmt.Println(result.LastInsertId())
		fmt.Println(result.RowsAffected())
		err = tx.Commit()
		if err !=nil{
			fmt.Println("(===2===)",err)
		}
	}

	//输出如下：
	//Error 1054: Unknown column 'age111' in 'field list'
	//准备回滚事务。。。


	//上面的这种写法才是最合适的！下面的写法仅仅是展示一些api的原理



	//必须重新开启事务，上一个事务对象不允许被重用，因为他已经被取消了
	tx,err= db2.BeginTx(context.Background(),&sql.TxOptions{Isolation: sql.LevelDefault})
	check_err_sql(err)
	// Prepare creates a prepared statement for use within a transaction.
	//
	// The returned statement operates within the transaction and can no longer
	// be used once the transaction has been committed or rolled back.
	//
	// To use an existing prepared statement on this transaction, see Tx.Stmt.
	// Prepare创建一个准备好的语句，供在事务中使用。
	// 返回的语句在事务内运行，并且一旦事务已提交或回滚就不能再使用。
	// 要在此事务上使用现有的准备好的语句，请参见Tx.Stmt。
	//底层会先提交过去这个预填充的语句过去给数据库，如果没有当前的字段的话，tx.Prepare会捕捉到错误并且抛出同时stmt=nil。
	//跟上一个api不同，上一个不需要准备语句，而是一次性发送完整的语句过去给数据库了！而不是分次发送语句过去给数据库！
	stmt, err = tx.Prepare(`update person set name = "anko55",age=95 where id=? or id=?`)
	if err !=nil{
		fmt.Println("(===1===)",err)
		//如果出错的话，我们应该在这里 进行结束，不再执行下面的语句，否则下面的语句会抛出异常！
		panic(err)
	}

	// Exec用给定的参数执行一个准备好的语句，并返回一个总结语句效果的结果。
	result, err = stmt.Exec(4, 5)//如果这里提交的id是不存在的话则不会抛出任何的错误，只是没反应！
	if err !=nil{
		fmt.Println("(===2===)",err)
	}
	if err != nil{
		//如果出错直接回滚
		fmt.Println("准备回滚事务。。。。")
		err111 := tx.Rollback()
		check_err_sql(err111)
	}else {
		fmt.Println(result.LastInsertId())
		fmt.Println(result.RowsAffected())
	}


	//我们看下一个事务一个stmt语句能不能提交2次的填充词汇 或者 一个事务能否存在多条stmt语句，很明显是可以的！
	stmt, err = tx.Prepare(`update person set name = "anko55",age=81 where id=? or id=?`)

	result, err = stmt.Exec(2, 3)//如果这里提交的id是不存在的话则不会抛出任何的错误，只是没反应！
	if err !=nil{
		fmt.Println("(===2*2===)",err)
	}
	if err != nil{
		//如果出错直接回滚
		fmt.Println("准备回滚事务。。。。")
		err111 := tx.Rollback()
		check_err_sql(err111)
	}else {
		fmt.Println(result.LastInsertId())
		fmt.Println(result.RowsAffected())
	}



	defer func() {
		//千万不要向下面这样捕获错误，因为如果已经提交了事务的话，err111也会不为nil的！commit（）方法同理
		//err111 := tx.Rollback()
		//if err111 !=nil{
		//	fmt.Println("(===3===)",err111)
		//}
		//假设上面不注释掉的话，无论事务是否成功都会输出：
		//(===3===) sql: transaction has already been committed or rolled back
		//
		//我们应该像下面这样不捕获异常，而是直接Rollback()，因为如果我们已经提交了事务的话，Rollback或者commit都是会检测到然后不进行回滚的或者提交的！
		tx.Rollback()
	}()

	//如果不出错的话直接提交整个更改
	fmt.Println("准备提交事务。。。。")
	err = tx.Commit()
	check_err_sql(err)
	//千万不能写这里，result会抛出空指针
	//fmt.Println(result.LastInsertId())
	//fmt.Println(result.RowsAffected())

	//输出：
	//	0 <nil>
	//	2 <nil>
	//	0 <nil>
	//	2 <nil>
	//	准备提交事务。。。。

	//多次提交相同结果的输出：
	//	0 <nil>
	//	2 <nil>
	//	0 <nil>
	//	0 <nil>，不知道这里为什么这里是0，可能是数据库缓存查询时候检查到相对应的值一样，但是绝对不是go的sql包缓存了（因为我探究过了）
	//	准备提交事务。。。。

	//此时数据库如下：
	//mysql> select * from person;
	//	+----+-----------+-----+
	//	| id | name      | age |
	//	+----+-----------+-----+
	//	|  1 | anko      |  20 |
	//	|  2 | anko55    |  85 |
	//	|  3 | anko55    |  85 |
	//	|  4 | anko55    |  95 |
	//	|  5 | anko55    |  95 |
	//	|  6 | Zhongshan |  44 |
	//	+----+-----------+-----+
	//	6 rows in set (0.00 sec)

	//从上面可以知道我们一个事务tx对象可以创建多个stmt的语句，同样一个语句可以exec解析执行多行的填充参数，
	//只需要一次的commit，可以提交绑定到该tx对象上面的所有与stmt语句！


	//我们验证如果一个事务结束后，事务相对应的语句是否能够继续用！！！
	result, err = stmt.Exec(2, 3)//如果这里提交的id是不存在的话则不会抛出任何的错误，只是没反应！
	if err !=nil{
		fmt.Println("事务已经结束，绑定到相对应事务对象上面的语句对象也是不能用了的！")
		fmt.Println("(===2===)",err)
	}else {
		fmt.Println(result.LastInsertId())
		fmt.Println(result.RowsAffected())
	}

	//以上输出：
	//	准备提交事务。。。。
	//	事务已经结束，绑定到相对应事务对象上面的语句对象也是不能用了的！
	//	(===2===) sql: statement is closed





	fmt.Println("-------------db中产生的语句 封装成 事务的语句------------")

	//必须重新开启事务，上一个事务对象不允许被重用，因为他已经被取消了
	tx,err= db2.BeginTx(context.Background(),&sql.TxOptions{Isolation: sql.LevelDefault})
	check_err_sql(err)

	// Stmt returns a transaction-specific prepared statement from
	// an existing statement.
	//
	// Example:
	//  updateMoney, err := db.Prepare("UPDATE balance SET money=money+? WHERE id=?")
	//  ...
	//  tx, err := db.Begin()
	//  ...
	//  res, err := tx.Stmt(updateMoney).Exec(123.45, 98293203)
	//
	// The returned statement operates within the transaction and will be closed
	// when the transaction has been committed or rolled back.

	// Stmt从现有语句返回特定于事务的预准备语句。
	// Example:
	//  updateMoney, err := db.Prepare("UPDATE balance SET money=money+? WHERE id=?")
	//  ...
	//  tx, err := db.Begin()
	//  ...
	//  res, err := tx.Stmt(updateMoney).Exec(123.45, 98293203)
	//
	// 返回的语句在事务内运行，并且在事务已提交或回滚后将关闭。
	// 底层主要用tx.StmtContext()


	stmt, err = db2.Prepare(`select * from person where id=? or name=?`)
	//这个语句是从db中产生的，还不能用于事务，因此才需要下面的语句来创建一个用于事务的语句
	tx_stmt := tx.Stmt(stmt)
	rows, err = tx_stmt.Query(2, "Zhongshan")//如果tx_stmt语句不是查询的语句的话，则Query（）方法将不会有更改！

	if err !=nil{
		fmt.Println("(===2*2===)",err)
	}

	PirntQueryRows(rows)
	//输出：
	//	{2 anko55 81}
	//	{6 Zhongshan 44}

	//tx.StmtContext()方法同理，不再累叙



	fmt.Println("-------------db2.Begin()开启一个事务------------")



	stmt, err = db2.Prepare(`UPDATE person SET name = ?, age = ? WHERE id = ? or id = ? `)//语句1，会成功提交
	//得到的stmt=nil,因此我们必须在这里就进行了panic，但不是回滚
	//stmt, err = db2.Prepare(`UPDATE person SET name = ?, age111 = ? WHERE id = ? or id = ? `)//语句2，会回滚,


	// Begin starts a transaction. The default isolation level is dependent on
	// the driver.
	// Begin开始一个事务。 默认隔离级别取决于驱动程序。

	if err != nil{
		check_err_sql(err)
		panic(err)

	}else {
		tx, err = db2.Begin()
		check_err_sql(err)
		tx_stmt = tx.Stmt(stmt)

		result, err = tx_stmt.Exec("uupp", 18, 2, 3)
		check_err_sql(err)

		if err != nil{
			fmt.Println("准备进行事务回滚。。。")
			err := tx.Rollback()
			check_err_sql(err)
		}else {
			fmt.Println("准备提交事务！！！")
			err := tx.Commit()
			check_err_sql(err)
		}

	}



	//语句1输出：
	//	准备提交事务！！！

	//语句2输出：
	//	Error 1054: Unknown column 'age111' in 'field list'
	//	panic: Error 1054: Unknown column 'age111' in 'field list'
	//
	//	goroutine 1 [running]:
	//	main.main()
	//		C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:1145 +0x44e2





	fmt.Println("-------------db2.Driver()获取数据库所使用的的驱动------------")


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
	//	// name名称是驱动程序特定格式的字符串。
	//	// open可以一个缓存的连接（以前关闭了一个连接），但是没有必要这样做； sql软件包维护一个空闲连接池，以进行有效的重用。
	//	// 返回的连接一次只能由一个goroutine使用。
	//	Open(name string) (Conn, error)
	//}


	// Driver returns the database's underlying driver.
	//驱动程序返回数据库的基础驱动程序。也就是返回这个db2的mysql驱动，所以在下面的open()中只需要指定相应的参数即可，而不必在指定驱动的类型mysql了
	dri := db2.Driver()
	defer db2.Close()

	//注意下面的driver.conn对象而不是slq.conn对象
	//Open的参数字符串必须是dsn格式，如下：
	//[user[:password]@][net[(addr)]]/dbname[?param1=value1&paramN=valueN]
	dri_conn, err := dri.Open("root:mysql@/godb")//跟sql.open().conn()差不多，只是参数不同
	check_err_sql(err)
	V_stmt, err := dri_conn.Prepare( "select * from person where id = ? or id = ?")
	check_err_sql(err)

	//注意下面是driver.rows而不是sql.rows对象,
	//Value必须为以下类型：
	//	int64
	//	float64
	//  bool
	//  []byte
	//  string
	//  time.Time
	//dri_rows, err := V_stmt.Query([]driver.Value{2, 3})//2,3类型错误
	dri_rows, err := V_stmt.Query([]driver.Value{int64(2), int64(3)})//2,3类型错误
	check_err_sql(err)

	//调用Next，以将下一行数据填充到提供的切片中。 提供的切片将与Columns（）的宽度相同。
	//当没有更多行时，下一步应返回io.EOF。
	//目标不应写入Next之外。 关闭行时应注意不要修改dest中保存的缓冲区。
	err = PirntQueryDriverRows(dri_rows)
	check_err_sql(err)
	//输出：
	//	[id name age]
	//	[]driver.Value{2, []uint8{0x75, 0x75, 0x70, 0x70}, 18}
	//	[2 [117 117 112 112] 18]
	//	[%!s(int64=2) uupp %!s(int64=18)]
	//	2uupp18
	//	-------
	//	[]driver.Value{3, []uint8{0x75, 0x75, 0x70, 0x70}, 18}
	//	[3 [117 117 112 112] 18]
	//	[%!s(int64=3) uupp %!s(int64=18)]
	//	3uupp18




	fmt.Println("-------------设置驱动连接的数目------------")


	//// DBStats contains database statistics.
	//// DBStats包含数据库统计信息。
	//type DBStats struct {
	//	MaxOpenConnections int // Maximum number of open connections to the database.(//与数据库的最大打开连接数。)
	//
	//	// Pool Status
	//	OpenConnections int // The number of established connections both in use and idle.(//正在使用和空闲的已建立连接数。)
	//	InUse           int // The number of connections currently in use.(//当前正在使用的连接数。)
	//	Idle            int // The number of idle connections.(//空闲连接数。)
	//
	//	// Counters(//计数器)
	//	WaitCount         int64         // The total number of connections waited for.(//等待的连接总数。)
	//	WaitDuration      time.Duration // The total time blocked waiting for a new connection.(//等待新连接被阻止的总时间。)
	//	MaxIdleClosed     int64         // The total number of connections closed due to SetMaxIdleConns.(//由于SetMaxIdleConns而关闭的连接总数。)
	//	MaxLifetimeClosed int64         // The total number of connections closed due to SetConnMaxLifetime.(//由于SetConnMaxLifetime而关闭的连接总数。)
	//}



	// Stats returns database statistics.
	//统计信息Stats返回数据库统计信息。
	for i:=0;i<10;i++ {//这是最奇葩的写法么？
		//新建连接
		db, err := sql.Open("mysql", "root:mysql@/godb")
		check_err_sql(err)
		defer db.Close()
		//并发10个连接同时访问数据库
		go func(i int) {
			// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
			//
			// Expired connections may be closed lazily before reuse.
			//
			// If d <= 0, connections are reused forever.
			// SetConnMaxLifetime设置连接可以重用的最长时间。也就是连接取出连接池最长多久
			//过期的连接可能会在重新使用之前延迟关闭。
			//如果d <= 0，则连接将永远重复使用。
			db.SetConnMaxLifetime(-1)

			// SetMaxOpenConns sets the maximum number of open connections to the database.
			//
			// If MaxIdleConns is greater than 0 and the new MaxOpenConns is less than
			// MaxIdleConns, then MaxIdleConns will be reduced to match the new
			// MaxOpenConns limit.
			//
			// If n <= 0, then there is no limit on the number of open connections.
			// The default is 0 (unlimited).
			// SetMaxOpenConns设置到数据库的最大打开连接数（或者叫做 从池子中取出的正在工作的conn数）。
			// 如果MaxIdleConns大于0并且新的MaxOpenConns小于MaxIdleConns，则MaxIdleConns将减少以匹配新的MaxOpenConns限制。
			// 如果n <= 0，则打开的连接数没有限制。
			// 默认值为0（无限制）。
			db.SetMaxOpenConns(6)

			// SetMaxIdleConns设置空闲连接池中的最大连接数。（创建好的但是还在池子中的没取出来工作的conn数）
			//如果MaxOpenConns大于0但小于新的MaxIdleConns，则新的MaxIdleConns将减少以匹配MaxOpenConns限制。
			//如果n <= 0，则不保留任何空闲连接。
			//当前默认的最大空闲连接数为2。在将来的版本中可能会更改。
			db.SetMaxIdleConns(3)

			rows, err = db.Query("select * from person where id =? or id =?", 2, 3)
			check_err_sql(err)
			fmt.Printf("==连接正在使用==%v====%+v\n",i,db.Stats())
			PirntQueryRows(rows)
			fmt.Printf("==连接已经释放==%v====%+v\n",i,db.Stats())
		}(i)

	}


	time.Sleep(3e9)
	//输出：
	//	==连接正在使用==3===={MaxOpenConnections:6 OpenConnections:1 InUse:1 Idle:0 WaitCount:0 WaitDuration:0s MaxIdleClosed:0 MaxLifetimeClosed:0}
	//	{2 uupp 18}
	//	{3 uupp 18}
	//	==连接已经释放==3===={MaxOpenConnections:6 OpenConnections:1 InUse:0 Idle:1 WaitCount:0 WaitDuration:0s MaxIdleClosed:0 MaxLifetimeClosed:0}
	//	==连接正在使用==0===={MaxOpenConnections:6 OpenConnections:1 InUse:1 Idle:0 WaitCount:0 WaitDuration:0s MaxIdleClosed:0 MaxLifetimeClosed:0}
	//	{2 uupp 18}
	//	{3 uupp 18}
	//	==连接正在使用==5===={MaxOpenConnections:6 OpenConnections:1 InUse:1 Idle:0 WaitCount:0 WaitDuration:0s MaxIdleClosed:0 MaxLifetimeClosed:0}
	//	==连接正在使用==2===={MaxOpenConnections:6 OpenConnections:1 InUse:1 Idle:0 WaitCount:0 WaitDuration:0s MaxIdleClosed:0 MaxLifetimeClosed:0}
	//	{2 uupp 18}
	//	{3 uupp 18}
	//	==连接已经释放==0===={MaxOpenConnections:6 OpenConnections:1 InUse:0 Idle:1 WaitCount:0 WaitDuration:0s MaxIdleClosed:0 MaxLifetimeClosed:0}
	//	==连接已经释放==2===={MaxOpenConnections:6 OpenConnections:1 InUse:1 Idle:0 WaitCount:0 WaitDuration:0s MaxIdleClosed:0 MaxLifetimeClosed:0}
	//	==连接已经释放==5===={MaxOpenConnections:6 OpenConnections:1 InUse:0 Idle:1 WaitCount:0 WaitDuration:0s MaxIdleClosed:0 MaxLifetimeClosed:0}
	//	==连接正在使用==4===={MaxOpenConnections:6 OpenConnections:1 InUse:1 Idle:0 WaitCount:0 WaitDuration:0s MaxIdleClosed:0 MaxLifetimeClosed:0}
	//	{2 uupp 18}
	//	{3 uupp 18}
	//	==连接已经释放==4===={MaxOpenConnections:6 OpenConnections:1 InUse:0 Idle:1 WaitCount:0 WaitDuration:0s MaxIdleClosed:0 MaxLifetimeClosed:0}
	//	==连接正在使用==1===={MaxOpenConnections:6 OpenConnections:1 InUse:1 Idle:0 WaitCount:0 WaitDuration:0s MaxIdleClosed:0 MaxLifetimeClosed:0}
	//	{2 uupp 18}
	//	{3 uupp 18}
	//	==连接已经释放==1===={MaxOpenConnections:6 OpenConnections:1 InUse:0 Idle:1 WaitCount:0 WaitDuration:0s MaxIdleClosed:0 MaxLifetimeClosed:0}
	//	==连接正在使用==6===={MaxOpenConnections:6 OpenConnections:1 InUse:1 Idle:0 WaitCount:0 WaitDuration:0s MaxIdleClosed:0 MaxLifetimeClosed:0}
	//	{2 uupp 18}
	//	{3 uupp 18}
	//	==连接已经释放==6===={MaxOpenConnections:6 OpenConnections:1 InUse:0 Idle:1 WaitCount:0 WaitDuration:0s MaxIdleClosed:0 MaxLifetimeClosed:0}
	//	==连接正在使用==8===={MaxOpenConnections:6 OpenConnections:1 InUse:1 Idle:0 WaitCount:0 WaitDuration:0s MaxIdleClosed:0 MaxLifetimeClosed:0}
	//	==连接正在使用==7===={MaxOpenConnections:6 OpenConnections:1 InUse:1 Idle:0 WaitCount:0 WaitDuration:0s MaxIdleClosed:0 MaxLifetimeClosed:0}
	//	{2 uupp 18}
	//	{3 uupp 18}
	//	==连接正在使用==9===={MaxOpenConnections:6 OpenConnections:1 InUse:1 Idle:0 WaitCount:0 WaitDuration:0s MaxIdleClosed:0 MaxLifetimeClosed:0}
	//	{2 uupp 18}
	//	{3 uupp 18}
	//	==连接已经释放==8===={MaxOpenConnections:6 OpenConnections:1 InUse:1 Idle:0 WaitCount:0 WaitDuration:0s MaxIdleClosed:0 MaxLifetimeClosed:0}
	//	==连接已经释放==7===={MaxOpenConnections:6 OpenConnections:1 InUse:0 Idle:1 WaitCount:0 WaitDuration:0s MaxIdleClosed:0 MaxLifetimeClosed:0}
	//	==连接已经释放==9===={MaxOpenConnections:6 OpenConnections:1 InUse:0 Idle:1 WaitCount:0 WaitDuration:0s MaxIdleClosed:0 MaxLifetimeClosed:0}

	//目测好像测试不出来，先搁置







	fmt.Println("------------sql.OpenDB（）（sql.open的底层方法）使用连接器Connector和Driver来连接---------------------")
	//// A Connector represents a driver in a fixed configuration
	//// and can create any number of equivalent Conns for use
	//// by multiple goroutines.
	////
	//// A Connector can be passed to sql.OpenDB, to allow drivers
	//// to implement their own sql.DB constructors, or returned by
	//// DriverContext's OpenConnector method, to allow drivers
	//// access to context and to avoid repeated parsing of driver
	//// configuration.
	////连接器代表固定配置的驱动程序，可以创建任意数量的等效Conns供多个goroutine使用。
	////一个连接器可以传递给sql.OpenDB，以允许驱动程序实现自己的sql.DB构造函数，或由DriverContext的OpenConnector方法返回，
	//// 以允许驱动程序访问上下文并避免重复分析驱动程序配置。
	//type Connector interface {
	//	// Connect returns a connection to the database.
	//	// Connect may return a cached connection (one previously
	//	// closed), but doing so is unnecessary; the sql package
	//	// maintains a pool of idle connections for efficient re-use.
	//	//
	//	// The provided context.Context is for dialing purposes only
	//	// (see net.DialContext) and should not be stored or used for
	//	// other purposes.
	//	//
	//	// The returned connection is only used by one goroutine at a
	//	// time.
	//	// Connect返回到数据库的连接。
	//	// Connect可能会返回一个缓存的连接（以前关闭了一个连接），但是没有必要这样做； sql软件包维护一个空闲连接池，以进行有效的重用。
	//	//提供的context.Context仅用于拨号目的（请参见net.DialContext），不应存储或用于其他目的。
	//	//返回的连接一次只能由一个goroutine使用。
	//	Connect(context.Context) (Conn, error)//这个Conn是driver.Conn
	//
	//	// Driver returns the underlying Driver of the Connector,
	//	// mainly to maintain compatibility with the Driver method
	//	// on sql.DB.
	//	//驱动程序返回连接器的基础驱动程序，主要是为了保持与sql.DB上的Driver方法的兼容性。
	//	Driver() Driver
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




	// OpenDB opens a database using a Connector, allowing drivers to
	// bypass a string based data source name.
	//
	// Most users will open a database via a driver-specific connection
	// helper function that returns a *DB. No database drivers are included
	// in the Go standard library. See https://golang.org/s/sqldrivers for
	// a list of third-party drivers.
	//
	// OpenDB may just validate its arguments without creating a connection
	// to the database. To verify that the data source name is valid, call
	// Ping.
	//
	// The returned DB is safe for concurrent use by multiple goroutines
	// and maintains its own pool of idle connections. Thus, the OpenDB
	// function should be called just once. It is rarely necessary to
	// close a DB.
	// OpenDB使用连接器打开数据库，允许驱动程序绕过基于字符串的数据源名称。
	//大多数用户将通过特定于驱动程序的连接帮助程序函数打开数据库，该函数返回* DB。 Go标准库中没有数据库驱动程序。 有关第三方驱动程序的列表，请参见https://golang.org/s/sqldrivers。
	// OpenDB可能只验证其参数而不创建与数据库的连接。 要验证数据源名称是否有效，请致电Ping。
	//返回的数据库可安全地供多个goroutine并发使用，并维护其自己的空闲连接池。 因此，OpenDB函数应该仅被调用一次。 很少需要关闭数据库。

	//dri = db2.Driver()//获取驱动*mysql.MySQLDriver，返回的是一个驱动接口，下面需要断言成该实例实现的其他的接口，（注意该实例对象实现了2个接口，由接口类型断言成另外一个接口类型是允许的）
	//fmt.Println(dri)
	//fmt.Println(reflect.TypeOf(dri))
	//driverContext := dri.(driver.DriverContext)

	//验证不使用*mysql.MySQLDriver看下能否正常执行，上面跟下面方式是一样的！
	var driverContext driver.DriverContext
	driverContext=&mysql.MySQLDriver{}//这样也是可以的，不过实例化了新的mysql.MySQLDriver对象
	connector, err := driverContext.OpenConnector("root:mysql@/godb")
	check_err_sql(err)
	db = sql.OpenDB(connector)
	rows, err = db.Query("select * from person where id =? or id =?", 4, 5)
	check_err_sql(err)
	PirntQueryRows(rows)
	//输出：
	//	{4 anko55 95}
	//	{5 anko55 95}



	//终于写完了，

	defer func() {
		err = stmt.Close()
		check_err_sql(err)
		fmt.Printf("%+v\n",stmt)
	}()














}









func PirntQueryRow(row *sql.Row)  {
	type person struct {
		id int
		name string
		age int
	}
	var p =person{}
	err := row.Scan(&p.id,&p.name, &p.age)
	check_err_sql(err)
	fmt.Println(p)
}

func PirntQueryRows(rows *sql.Rows)  {
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
}

func PirntQueryDriverRows(rows driver.Rows) error {

	defer rows.Close()
	// Columns returns the names of the columns. The number of
	// columns of the result is inferred from the length of the
	// slice. If a particular column name isn't known, an empty
	// string should be returned for that entry.
	// Columns返回列的名称。 从切片的长度推断出结果的列数。 如果不知道特定的列名，则应为该条目返回一个空字符串。
	fmt.Println(rows.Columns())
	// Next is called to populate the next row of data into
	// the provided slice. The provided slice will be the same
	// size as the Columns() are wide.
	//
	// Next should return io.EOF when there are no more rows.
	//
	// The dest should not be written to outside of Next. Care
	// should be taken when closing Rows not to modify
	// a buffer held in dest.
	// 调用Next，以将下一行数据填充到提供的切片中。 提供的切片将与Columns（）的宽度相同。
	// 当没有更多行时，下一步应返回io.EOF。
	// 目标不应写入Next之外。 关闭行时应注意不要修改dest中保存的缓冲区。
	var dest =make([]driver.Value,len(rows.Columns()))//这里必须指明长度，否则无法存值！这跟初始化和没初始化没关系
	err := rows.Next(dest)
	check_err_sql(err)//我们在这里并没有抛出异常，我们只是输出一些东西
	if err !=nil{
		return err
	}
	P_func:= func() {
		fmt.Printf("%#v\n",dest)
		fmt.Printf("%v\n",dest)
		fmt.Printf("%s\n",dest)//因为driver.Value的基类是interface{}类型，所以我们这里无法使用%s的形式进行转化输出字符串
		//"%s%s%s\n"也可以直接不换为"%d%s%d\n"
		fmt.Printf("%s%s%s\n",strconv.Itoa(int(dest[0].(int64))),dest[1],strconv.Itoa(int(dest[2].(int64))))//因为driver.Value的基类是interface{}类型，所以我们这里无法使用%s的形式进行转化输出字符串
	}
	P_func()

	fmt.Println("-------")
	err = rows.Next(dest)
	check_err_sql(err)//我们在这里并没有抛出异常，我们只是输出一些东西
	if err !=nil{
		return err
	}
	P_func()

	//fmt.Println("-------")
	//err = rows.Next(dest)
	//check_err_sql(err)//我们在这里并没有抛出异常，我们只是输出一些东西
	//if err !=nil{
	//	return err
	//}
	//P_func()
	//由于我们只是查询了2条数据，继续Next（）将会抛出异常EOF
	//panic: EOF

	return nil

}



func check_err_sql(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
