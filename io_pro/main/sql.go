package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //一定要导入这个
	"time"
)

//Goals of the sql and sql/driver packages:
//
//* Provide a generic database API for a variety of SQL or SQL-like
//databases.  There currently exist Go libraries for SQLite, MySQL,
//and Postgres, but all with a very different feel, and often
//a non-Go-like feel.
//
//* Feel like Go.
//
//* Care mostly about the common cases. Common SQL should be portable.
//SQL edge cases or db-specific extensions can be detected and
//conditionally used by the application.  It is a non-goal to care
//about every particular db's extension or quirk.
//
//* Separate out the basic implementation of a database driver
//(implementing the sql/driver interfaces) vs the implementation
//of all the user-level types and convenience methods.
//In a nutshell:
//
//User Code ---> sql package (concrete types) ---> sql/driver (interfaces)
//Database Driver -> sql (to register) + sql/driver (implement interfaces)
//
//* Make type casting/conversions consistent between all drivers. To
//achieve this, most of the conversions are done in the sql package,
//not in each driver. The drivers then only have to deal with a
//smaller set of types.
//
//* Be flexible with type conversions, but be paranoid about silent
//truncation or other loss of precision.
//
//* Handle concurrency well.  Users shouldn't need to care about the
//database's per-connection thread safety issues (or lack thereof),
//and shouldn't have to maintain their own free pools of connections.
//The 'sql' package should deal with that bookkeeping as needed.  Given
//an *sql.DB, it should be possible to share that instance between
//multiple goroutines, without any extra synchronization.
//
//* Push complexity, where necessary, down into the sql+driver packages,
//rather than exposing it to users. Said otherwise, the sql package
//should expose an ideal database that's not finnicky about how it's
//accessed, even if that's not true.
//
//* Provide optional interfaces in sql/driver for drivers to implement
//for special cases or fastpaths.  But the only party that knows about
//those is the sql package.  To user code, some stuff just might start
//working or start working slightly faster.
//sql和sql / driver软件包的目标：
//*为各种SQL或类似SQL的数据库提供通用数据库API。当前存在用于SQLite，MySQL和Postgres的Go库，但它们的感觉却截然不同，并且常常具有非Go的感觉。
//*Feel like Go。
//*主要关心常见情况。通用SQL应该是可移植的。
//应用程序可以检测到SQL边缘情况或特定于db的扩展。关心每个特定数据库的扩展名或怪癖是非目标。
//*分离出数据库驱动程序的基本实现（实现sql / driver接口）与所有用户级别类型和便捷方法的实现。
//  简而言之：
//  用户代码---> sql程序包（具体类型）---> sql/driver（接口）
//  数据库驱动程序-> sql（要注册）+ sql/driver（实现接口）
//*使所有驱动程序之间的类型转换/转换保持一致。为此，大多数转换是在sql包中完成的，而不是在每个驱动程序中完成的。然后，驱动程序只需要处理较小的一组类型。
//*对类型转换要灵活一些，但对无声截断或其他精度损失会产生疑虑。
//*处理好并发性。用户不必关心数据库的每个连接线程的安全性问题（或缺少这些问题），也不必维护自己的空闲连接池。
//“ sql”包应根据需要处理该簿记。给定一个*sql.DB，应该可以在多个goroutine之间共享该实例，而无需任何额外的同步。
//*在必要时将复杂性降低到sql+driver软件包中，而不是向用户公开。否则，sql程序包应该公开一个理想的数据库，即使它不是真的，它也对访问方式并不敏感。
//*在sql/driver中提供可选接口，以供驱动程序针对特殊情况或快速路径实现。但是唯一知道这些的一方是sql软件包。对于用户代码，有些东西可能只是开始工作或开始工作得更快。



func main() {
	// Open opens a database specified by its database driver name and a driver-specific data source name, usually consisting of at least a database name and connection information.
	// Most users will open a database via a driver-specific connection helper function that returns a *DB. No database drivers are included in the Go standard library. See https://golang.org/s/sqldrivers for a list of third-party drivers.
	// Open may just validate its arguments without creating a connection to the database. To verify that the data source name is valid, call Ping.
	// The returned DB is safe for concurrent use by multiple goroutines and maintains its own pool of idle connections. Thus, the Open function should be called just once. It is rarely necessary to close a DB. // Open打开一个由其数据库驱动程序名称和特定于驱动程序的数据源名称指定的数据库，该名称通常至少由数据库名称和连接信息组成。

	//大多数用户将通过特定于驱动程序的连接帮助程序函数打开数据库，该函数返回* DB。 Go标准库中没有数据库驱动程序。 有关第三方驱动程序的列表，请参见https://golang.org/s/sqldrivers。

	// Open只会验证其参数而不创建与数据库的连接。 要验证数据源名称是否有效，请调用Ping()。

	//返回的数据库可安全地供多个goroutine并发使用(因为这个open函数的实现里面对驱动连接的访问和读取加了读写锁)，并维护其自己的空闲连接池。 因此，Open函数应仅被调用一次。 很少需要关闭数据库。
	//"mysql"：指定的是操作的数据库程序类型
	//"root:mysql@/person":root是数据库用户名，mysql是数据库密码，person指定操作的是数据库哪个database
	db, err := sql.Open("mysql", "root:mysql@/godb")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	// Close closes the database and prevents new queries from starting.
	// Close then waits for all queries that have started processing on the server to finish.
	//
	// It is rare to Close a DB, as the DB handle is meant to be long-lived and shared between many goroutines.
	// Close关闭数据库并阻止启动新查询。
	//然后Close等待已经在服务器上开始处理的所有查询完成之后再完全关闭连接。
	//
	//很少关闭数据库，因为该数据库句柄是长期存在的并且在许多goroutine之间共享。
	defer db.Close()

	// Ping verifies a connection to the database is still alive,
	// establishing a connection if necessary.
	// Ping验证与数据库的连接是否仍然存在，并在必要时（数据库的验证信息完全正确时）建立连接。

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Use the DB normally, execute the querys etc

	// Query executes a query that returns rows, typically a SELECT.
	// The args are for any placeholder parameters in the query.
	// Query执行的查询返回行，通常是SELECT。
	// args用于查询中的任何占位符参数。
	//底层调用db.QueryContext(context.Background(), query, args...)


	//// Rows is the result of a query. Its cursor starts before the first row
	//// of the result set. Use Next to advance from row to row.
	////行是查询的结果。 它的光标从结果集的第一行开始。 使用“Next()”在每一行中前进。
	//type Rows struct {
	//	dc          *driverConn // owned; must call releaseConn when closed to release(//数据库连接对象； 关闭时必须调用releaseConn才能释放)
	//	releaseConn func(error)	//释放数据库连接的函数字段
	//	rowsi       driver.Rows
	//	cancel      func()      // called when Rows is closed, may be nil.（//当Rows关闭时调用，可能为nil。）
	//	closeStmt   *driverStmt // if non-nil, statement to Close on close（//如果非nil，则在关闭时声明Close）
	//
	//	// closemu prevents Rows from closing while there
	//	// is an active streaming result. It is held for read during non-close operations
	//	// and exclusively during close.
	//	// closemu防止在有活动流结果时关闭Rows。 在非关闭操作期间 和 仅在关闭期间 将其保留以供读取。
	//	//
	//	// closemu guards lasterr and closed.（closemu守卫lasterr并关闭。）
	//	closemu sync.RWMutex
	//	closed  bool	//标志连接关闭与否
	//	lasterr error // non-nil only if closed is true（//仅当close为true时才为非null）
	//
	//	// lastcols is only used in Scan, Next, and NextResultSet which are expected
	//	// not to be called concurrently.
	//	// lastcols仅在不希望同时调用的Scan，Next和NextResultSet中使用。
	//	lastcols []driver.Value
	//}


	//// driverConn wraps a driver.Conn with a mutex, to
	//// be held during all calls into the Conn. (including any calls onto
	//// interfaces returned via that Conn, such as calls on Tx, Stmt,
	//// Result, Rows)
	//// driverConn将具有互斥量的driver.Conn包装在要调用Conn的所有过程中（包括通过该Conn返回的接口的任何调用，例如Tx，Stmt，Result，Row的调用）
	//// 说白了就是数据库连接的执行者，但不是数据库连接对象DB（关于DB在下面介绍到），他是数据库连接对象的操作者（增删改查操作等）
	//type driverConn struct {
	//	db        *DB			//指定在哪个连接上进行操作
	//	createdAt time.Time		//指定创建数据库连接对象的操作者driverConn的时间
	//
	//	sync.Mutex  // guards following（守卫跟随）
	//	ci          driver.Conn	//组合这个 driver.Conn接口
	//	closed      bool		//数据库的操作流关闭与否
	//	finalClosed bool // ci.Close has been called（ci.Close是否已被调用，决定需不需要阻塞等待锁的释放）
	//	openStmt    map[*driverStmt]bool
	//	lastErr     error // lastError captures the result of the session resetter.（lastError捕获会话重置器的结果。）
	//
	//	// guarded by db.mu（由db.mu保护）
	//	inUse      bool
	//	onPut      []func() // code (with db.mu held) run when conn is next returned（下次返回conn时，（运行db.mu的）代码保持运行）
	//	dbmuClosed bool     // same as closed, but guarded by db.mu, for removeClosedStmtLocked（与closed相同，但受db.mu保护，用于removeClosedStmtLocked）
	//}



	//// Conn is a connection to a database. It is not used concurrently
	//// by multiple goroutines.
	//// Conn是到数据库的连接。 多个goroutine不能同时使用它。
	////
	//// Conn is assumed to be stateful.（Conn被假定为有状态的。）
	//type Conn interface {
	//	// Prepare returns a prepared statement, bound to this connection.
	//	//Prepare返回绑定到此连接的已准备语句。
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


	//// Rows is an iterator over an executed query's results.
	//// Rows是执行查询结果的迭代器。
	//type Rows interface {
	//	// Columns returns the names of the columns. The number of
	//	// columns of the result is inferred from the length of the
	//	// slice. If a particular column name isn't known, an empty
	//	// string should be returned for that entry.
	//	//Columns返回列的名称。 从切片的长度推断出结果的列数。 如果不知道特定的列名，则应为该条目返回一个空字符串。
	//	Columns() []string
	//
	//	// Close closes the rows iterator.(// Close关闭行迭代器。)
	//	Close() error
	//
	//	// Next is called to populate the next row of data into
	//	// the provided slice. The provided slice will be the same
	//	// size as the Columns() are wide.
	//	//
	//	// Next should return io.EOF when there are no more rows.
	//	//
	//	// The dest should not be written to outside of Next. Care
	//	// should be taken when closing Rows not to modify
	//	// a buffer held in dest.
	//	//调用Next，以将下一行数据填充到提供的切片中。 提供的切片将与Columns（）的宽度相同。
	//	//当没有更多行时，下一步应返回io.EOF。
	//	//目标不应写入Next之外。 关闭行时应注意不要修改dest中保存的缓冲区。
	//	Next(dest []Value) error
	//}


	//// DB is a database handle representing a pool of zero or more
	//// underlying connections. It's safe for concurrent use by multiple
	//// goroutines.
	////
	//// The sql package creates and frees connections automatically; it
	//// also maintains a free pool of idle connections. If the database has
	//// a concept of per-connection state, such state can be reliably observed
	//// within a transaction (Tx) or connection (Conn). Once DB.Begin is called, the
	//// returned Tx is bound to a single connection. Once Commit or
	//// Rollback is called on the transaction, that transaction's
	//// connection is returned to DB's idle connection pool. The pool size
	//// can be controlled with SetMaxIdleConns.
	//// DB是代表零个或多个基础连接池的数据库句柄。 对于多个goroutine并发使用是安全的。
	//// sql软件包自动创建并释放连接； 它还维护空闲连接的空闲池。 如果数据库具有按连接状态的概念，
	//// 则可以在事务（Tx）或连接（Conn）中可靠地观察到这种状态。 调用DB.Begin之后，返回的Tx将绑定到单个连接。
	//// 在事务上调用Commit或Rollback后，该事务的连接将返回到DB的空闲连接池。 池大小可以使用SetMaxIdleConns控制。
	//type DB struct {
	//	// Atomic access only. At top of struct to prevent mis-alignment
	//	// on 32-bit platforms. Of type time.Duration.
	//	//仅限原子访问。 防止在32位平台上出现未对齐问题的结构的顶部。 类型为time.Duration。
	//	waitDuration int64 // Total time waited for new connections.（等待新连接的总时间。）
	//
	//	connector driver.Connector	//组合接口（下面会讲到）
	//	// numClosed is an atomic counter which represents a total number of
	//	// closed connections. Stmt.openStmt checks it before cleaning closed
	//	// connections in Stmt.css.
	//	// numClosed是一个原子计数器，表示关闭的连接总数。 在清除Stmt.css中的已关闭连接之前，Stmt.openStmt会对其进行检查。
	//	numClosed uint64
	//
	//	mu           sync.Mutex // protects following fields（保护以下字段）
	//	freeConn     []*driverConn
	//	connRequests map[uint64]chan connRequest
	//	nextRequest  uint64 // Next key to use in connRequests.（在connRequests中使用的下一个键。）
	//	numOpen      int    // number of opened and pending open connections（已打开和等候打开的连接数）
	//	// Used to signal the need for new connections
	//	// a goroutine running connectionOpener() reads on this chan and
	//	// maybeOpenNewConnections sends on the chan (one send per needed connection)
	//	// It is closed during db.Close(). The close tells the connectionOpener
	//	// goroutine to exit.
	//	//用于表示需要建立新连接的goroutine在此chan上读取正在运行的goroutine（），也许OpenNewConnections在chan上发送（每个所需的连接发送一次）。
	//	// 在db.Close（）期间关闭。 关闭将通知connectionOpener goroutine退出。
	//	openerCh          chan struct{}
	//	resetterCh        chan *driverConn
	//	closed            bool
	//	dep               map[finalCloser]depSet
	//	lastPut           map[*driverConn]string // stacktrace of last conn's put; debug only（最后一个conn的stacktrace； 仅调试）
	//	maxIdle           int                    // zero means defaultMaxIdleConns; negative means 0（零表示defaultMaxIdleConns; 负数表示0）
	//	maxOpen           int                    // <= 0 means unlimited（<= 0表示无限制）
	//	maxLifetime       time.Duration          // maximum amount of time a connection may be reused（连接可以被重用的最长时间）
	//	cleanerCh         chan struct{}
	//	waitCount         int64 // Total number of connections waited for.（等待的连接总数。）
	//	maxIdleClosed     int64 // Total number of connections closed due to idle.（由于空闲而关闭的连接总数。）
	//	maxLifetimeClosed int64 // Total number of connections closed due to max free limit.（由于最大可用限制而关闭的连接总数。）
	//
	//	stop func() // stop cancels the connection opener and the session resetter.（stop取消连接打开程序和会话重置程序。）
	//}



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
	////一个连接器可以传递给sql.OpenDB，以允许驱动程序实现自己的sql.DB构造函数，或由DriverContext的OpenConnector方法返回，以允许驱动程序访问上下文并避免重复分析驱动程序配置。
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
	//	Connect(context.Context) (Conn, error)
	//
	//	// Driver returns the underlying Driver of the Connector,
	//	// mainly to maintain compatibility with the Driver method
	//	// on sql.DB.
	//	//驱动程序返回连接器的基础驱动程序，主要是为了保持与sql.DB上的Driver方法的兼容性。
	//	Driver() Driver
	//}



	//// driverStmt associates a driver.Stmt with the
	//// *driverConn from which it came, so the driverConn's lock can be
	//// held during calls.
	//// driverStmt将一个driver.Stmt与它来自的* driverConn关联，因此可以在调用过程中保持driverConn的锁。
	//type driverStmt struct {
	//	sync.Locker // the *driverConn
	//	si          driver.Stmt
	//	closed      bool
	//	closeErr    error // return value of previous Close call（上一个Close调用的返回值）
	//}


	//age := 20
	//rows, err := db.Query("SELECT * FROM person WHERE age=?", age)//语句1
	rows, err := db.Query("SELECT * FROM person ")//语句2
	check_err_sql(err)
	fmt.Printf("%+v\n",rows)
	// Next准备下一个结果行，以使用Scan方法读取。 如果成功，它将返回true；如果没有下一个结果行或在准备它时发生错误，则返回false。 应咨询Err来区分这两种情况。
	//每次调用Scan，即使是第一个调用，都必须先调用Next。
	fmt.Println(rows.Next())
	//var dst =make([]string,3)
	var dst =make([]string,6)
	//扫描将当前行中的列复制到dest指向的值中。 dest中的值数必须与“行”中的列数相同。
	//
	//扫描将从数据库读取的列转换为sql包提供的以下常见Go类型和特殊类型：
	//
	//    *string
	//    *[]byte
	//    *int, *int8, *int16, *int32, *int64
	//    *uint, *uint8, *uint16, *uint32, *uint64
	//    *bool
	//    *float32, *float64
	//    *interface{}
	//    *RawBytes
	//    *Rows (cursor value)
	//	  任何实现Scanner的类型（请参阅Scanner文档）
	//
	//在最简单的情况下，如果源列中的值的类型是整数，布尔型或字符串类型T，而dest的类型是* T，则Scan仅通过指针分配值。
	//
	//只要没有信息丢失，扫描也会在字符串和数字类型之间转换。扫描将所有从数字数据库列扫描的数字都字符串化为* string时，将检查对数字类型的扫描是否溢出。例如，尽管float64（255）或“ 255”可以扫描到uint8，但是值300的float64或值“ 300”的字符串可以扫描到uint16，但不能扫描到uint8。一个例外是，将一些float64数字扫描到字符串可能会在字符串化时丢失信息。通常，将浮点列扫描到* float64中。
	//
	//如果dest参数的类型为* [] byte，则Scan在该参数中保存相应数据的副本。该副本归调用方所有，可以无限期修改和保存。可以通过使用* RawBytes类型的参数来避免复制。有关其使用限制，请参见RawBytes文档。
	//
	//如果参数的类型为* interface {}，则Scan会复制基础驱动程序提供的值，而无需进行转换。从[] byte类型的源值扫描到* interface {}时，将生成切片的副本，并且调用者拥有结果。
	//
	//可以将类型为time.Time的源值扫描为* time.Time，* interface {}，* string或* [] byte类型的值。转换为后两者时，将使用time.RFC3339Nano。
	//
	// bool类型的源值可以扫描为* bool，* interface {}，* string，* [] byte或* RawBytes类型。
	//
	//为了扫描到*bool中，源可能是strconv.ParseBool可以解析的true，false，1、0或字符串输入。
	//
	//扫描还可以将查询返回的游标（例如“从double中选择游标（从my_table中选择*）”）转换为* Rows值，该值可以从中进行扫描。如果父* Rows已关闭，则父选择查询将关闭所有游标* Rows。

	err = rows.Scan(&dst[0],&dst[1],&dst[2],)//必须传指针
	check_err_sql(err)
	//如果是只有一条数据的话则不用写下面的345,如果你需要rows.ColumnTypes()见到效果的话则需要保证下面被注释掉，也就是不scan获取光所有查询到的数据
	//fmt.Println(rows.Next())//前进到下一条的数据
	//err = rows.Scan(&dst[3],&dst[4],&dst[5],)//必须传指针
	//check_err_sql(err)

	fmt.Println("-----")
	fmt.Println(rows.Next())//查看是否还有下一行的数据，跟下面的NextResultSet()不同在于NextResultSet()返回的是下一个结果集而不是返回当前结果集中的下一行数据
	//Columns返回列名称。
	//如果行被关闭，则列返回错误。
	fmt.Println(rows.Columns())


	//// ColumnType contains the name and type of a column.(// ColumnType包含列的名称和类型。)
	//type ColumnType struct {
	//	name string
	//
	//	hasNullable       bool
	//	hasLength         bool
	//	hasPrecisionScale bool
	//
	//	nullable     bool
	//	length       int64
	//	databaseType string
	//	precision    int64
	//	scale        int64
	//	scanType     reflect.Type
	//}
	// ColumnTypes返回列信息，例如列类型，长度和可为空。 某些驱动程序可能无法提供某些信息。
	ColumnTypes, err := rows.ColumnTypes()
	check_err_sql(err)
	fmt.Println(ColumnTypes)
	for key, value := range ColumnTypes {
		fmt.Printf("%v:%+v\n",key,value)
	}
	// Err返回在迭代过程中遇到的错误（如果有）。
	//可以在显式或隐式Close之后调用Err。
	fmt.Println(rows.Err())
	// NextResultSet prepares the next result set for reading. It reports whether
	// there is further result sets, or false if there is no further result set
	// or if there is an error advancing to it. The Err method should be consulted
	// to distinguish between the two cases.
	//
	// After calling NextResultSet, the Next method should always be called before
	// scanning. If there are further result sets they may not have rows in the result
	// set.
	// NextResultSet准备下一个要读取的结果集。 它报告是否还有其他结果集，如果没有其他结果集或前进时出现错误，则报告false。 应该使用Err方法来区分这两种情况。
	//调用NextResultSet之后，应始终在扫描之前调用Next方法。 如果还有其他结果集，则结果集中可能没有行。

	//// RowsNextResultSet extends the Rows interface by providing a way to signal
	//// the driver to advance to the next result set.
	//// RowsNextResultSet通过提供一种向驱动程序发出信号以前进到下一个结果集的方式，扩展了Rows接口。
	//type RowsNextResultSet interface {
	//	Rows
	//
	//	// HasNextResultSet is called at the end of the current result set and
	//	// reports whether there is another result set after the current one.
	//	// HasNextResultSet在当前结果集的末尾被调用，并报告在当前结果集之后是否还有另一个结果集。
	//	HasNextResultSet() bool
	//
	//	// NextResultSet advances the driver to the next result set even
	//	// if there are remaining rows in the current result set.
	//	//
	//	// NextResultSet should return io.EOF when there are no more result sets.
	//	//即使当前结果集中有剩余的行，NextResultSet也会使驱动程序前进到下一个结果集。
	//	//当没有更多结果集时，NextResultSet应该返回io.EOF。
	//	NextResultSet() error
	//}


	fmt.Println(rows.NextResultSet())
	fmt.Println(dst)//注意interface{}类型不能用string()进行转型，同时[]string不能转为[]byte类型
	// Close关闭行，防止进一步枚举。 如果调用Next并返回false，并且没有其他结果集，则将自动关闭行，这将足以检查Err的结果。
	// 关闭是幂等的，不会影响Err的结果。
	defer rows.Close()
	fmt.Println(rows.Err())
	//语句1输出：
	//	&{dc:0xc0000b2100
	//	releaseConn:0x4d7a70
	//	rowsi:0xc0000a00f0
	//	cancel:<nil>
	//	closeStmt:0xc000046080
	//	closemu:{w:{state:0 sema:0} writerSem:0 readerSem:0 readerCount:0 readerWait:0}
	//	closed:false
	//	lasterr:<nil>
	//	lastcols:[]}
	//	true
	//	-----
	//	false
	//	[] sql: Rows are closed
	//	[] sql: Rows are closed
	//	<nil>
	//	false
	//	[1 anko 20]

	//假设语句2捕获一条的数据时候的输出：
	//	&{dc:0xc00009a100
	//	releaseConn:0x4d8bb0
	//	rowsi:0xc0000880f0
	//	cancel:<nil>
	//	closeStmt:<nil>
	//	closemu:{w:{state:0 sema:0} writerSem:0 readerSem:0 readerCount:0 readerWait:0}
	//	closed:false
	//	lasterr:<nil>
	//	lastcols:[]}
	//	true
	//	-----
	//	true
	//	[id name age] <nil>
	//	[0xc000088140 0xc000088190 0xc0000881e0]
	//	0:&{name:id hasNullable:true hasLength:false hasPrecisionScale:false nullable:false length:0 databaseType:INT precision:0 scale:0 scanType:0x60f2e0}
	//	1:&{name:name hasNullable:true hasLength:false hasPrecisionScale:false nullable:false length:0 databaseType:VARCHAR precision:0 scale:0 scanType:0x618240}
	//	2:&{name:age hasNullable:true hasLength:false hasPrecisionScale:false nullable:false length:0 databaseType:INT precision:0 scale:0 scanType:0x60f2e0}
	//	<nil>
	//	false
	//	[1 anko 20   ]

	//假设语句2捕获2条的数据（一共才2条）时候的输出：
	//	&{dc:0xc00009a100
	//	releaseConn:0x4d8bb0
	//	rowsi:0xc0000880f0
	//	cancel:<nil>
	//	closeStmt:<nil>
	//	closemu:{w:{state:0 sema:0} writerSem:0 readerSem:0 readerCount:0 readerWait:0}
	//	closed:false
	//	lasterr:<nil>
	//	lastcols:[]}
	//	true
	//	true
	//	-----
	//	false
	//	[] sql: Rows are closed
	//	[] sql: Rows are closed
	//	<nil>
	//	false
	//	[1 anko 20 2 gogo 32]


	fmt.Println("----------------------------------------------")
	//rows, err = db.Query("select * from person where id=1")//语句1
	rows, err = db.Query("select * from person ")//语句2
	check_err_sql(err)
	defer rows.Close()
	type person struct {
		id int
		name string
		age int
	}
	//下面仅仅是针对一条结果集，如果有多条结果集的话这样遍历会出错
	for b:=rows.Next(); b != false;b=rows.Next() {
		var p =person{}
		err := rows.Scan(&p.id, &p.name, &p.age)
		check_err_sql(err)
		fmt.Println(p)
	}

	//语句1输出：
	//	{1 anko 20}

	//语句2输出：
	//	{1 anko 20}
	//	{2 gogo 32}


	fmt.Println("-------------------关于Context在sql中的应用---------------------------")

	//// A Context carries a deadline, a cancellation signal, and other values across
	//// API boundaries.
	////
	//// Context's methods may be called by multiple goroutines simultaneously.
	////Context上下文在API边界上带有期限，取消信号和其他值。
	////多个goroutine可以同时调用Context的方法。
	//type Context interface {
	//	// Deadline returns the time when work done on behalf of this context
	//	// should be canceled. Deadline returns ok==false when no deadline is
	//	// set. Successive calls to Deadline return the same results.
	//	// 截止日期返回应取消代表该上下文完成的工作的时间。 如果未设置截止日期，则截止日期返回ok == false。
	//	// 连续调用Deadline会返回相同的结果。
	//	Deadline() (deadline time.Time, ok bool)
	//
	//	// Done returns a channel that's closed when work done on behalf of this
	//	// context should be canceled. Done may return nil if this context can
	//	// never be canceled. Successive calls to Done return the same value.
	//	//
	//	// WithCancel arranges for Done to be closed when cancel is called;
	//	// WithDeadline arranges for Done to be closed when the deadline
	//	// expires; WithTimeout arranges for Done to be closed when the timeout
	//	// elapses.
	//	//
	//	// Done is provided for use in select statements:
	//	//
	//	//  // Stream generates values with DoSomething and sends them to out
	//	//  // until DoSomething returns an error or ctx.Done is closed.
	//	//  func Stream(ctx context.Context, out chan<- Value) error {
	//	//  	for {
	//	//  		v, err := DoSomething(ctx)
	//	//  		if err != nil {
	//	//  			return err
	//	//  		}
	//	//  		select {
	//	//  		case <-ctx.Done():
	//	//  			return ctx.Err()
	//	//  		case out <- v:
	//	//  		}
	//	//  	}
	//	//  }
	//	//
	//	// See https://blog.golang.org/pipelines for more examples of how to use
	//	// a Done channel for cancellation.
	//
	//	// Done返回一个通道，当取消代表该上下文的工作时，该通道已关闭。 如果此上下文永远无法取消，则可能会返回nil。 连续调用Done将返回相同的值。
	//	// WithCancel安排在调用cancel时关闭Done；
	//	// WithDeadline安排在截止日期到期时关闭“Done”； WithTimeout安排超时后关闭“Done”。
	//	// Done被提供用于select语句：
	//	// //Stream流使用DoSomething生成值并将其发送出去，直到DoSomething返回错误或ctx.Done关闭为止。
	//	//  func Stream(ctx context.Context, out chan<- Value) error {
	//	//  	for {
	//	//  		v, err := DoSomething(ctx)
	//	//  		if err != nil {
	//	//  			return err
	//	//  		}
	//	//  		select {
	//	//  		case <-ctx.Done():
	//	//  			return ctx.Err()
	//	//  		case out <- v:
	//	//  		}
	//	//  	}
	//	//  }
	//	//有关如何使用“Done”管道进行取消的更多示例，请参见https://blog.golang.org/pipelines。
	//
	//
	//	Done() <-chan struct{}
	//
	//	// If Done is not yet closed, Err returns nil.
	//	// If Done is closed, Err returns a non-nil error explaining why:
	//	// Canceled if the context was canceled
	//	// or DeadlineExceeded if the context's deadline passed.
	//	// After Err returns a non-nil error, successive calls to Err return the same error.
	//	// 如果尚未关闭Done，则Err返回nil。
	//	// 如果Done关闭，Err将返回非nil错误，解释原因：
	//	// 如果上下文已取消，则返回取消Canceled；如果上下文的截止日期已过，则返回截止日期已过DeadlineExceeded。
	//	// Err返回非nil错误后，对Err的连续调用将返回相同的错误。
	//	Err() error
	//
	//	// Value returns the value associated with this context for key, or nil
	//	// if no value is associated with key. Successive calls to Value with
	//	// the same key returns the same result.
	//	//
	//	// Use context values only for request-scoped data that transits
	//	// processes and API boundaries, not for passing optional parameters to
	//	// functions.
	//	//
	//	// A key identifies a specific value in a Context. Functions that wish
	//	// to store values in Context typically allocate a key in a global
	//	// variable then use that key as the argument to context.WithValue and
	//	// Context.Value. A key can be any type that supports equality;
	//	// packages should define keys as an unexported type to avoid
	//	// collisions.
	//	//
	//	// Packages that define a Context key should provide type-safe accessors
	//	// for the values stored using that key:
	//	//
	//	// 	// Package user defines a User type that's stored in Contexts.
	//	// 	package user
	//	//
	//	// 	import "context"
	//	//
	//	// 	// User is the type of value stored in the Contexts.
	//	// 	type User struct {...}
	//	//
	//	// 	// key is an unexported type for keys defined in this package.
	//	// 	// This prevents collisions with keys defined in other packages.
	//	// 	type key int
	//	//
	//	// 	// userKey is the key for user.User values in Contexts. It is
	//	// 	// unexported; clients use user.NewContext and user.FromContext
	//	// 	// instead of using this key directly.
	//	// 	var userKey key
	//	//
	//	// 	// NewContext returns a new Context that carries value u.
	//	// 	func NewContext(ctx context.Context, u *User) context.Context {
	//	// 		return context.WithValue(ctx, userKey, u)
	//	// 	}
	//	//
	//	// 	// FromContext returns the User value stored in ctx, if any.
	//	// 	func FromContext(ctx context.Context) (*User, bool) {
	//	// 		u, ok := ctx.Value(userKey).(*User)
	//	// 		return u, ok
	//	// 	}
	//
	//	// Value返回与此键的上下文关联的值；如果没有值与key关联，则返回nil。使用相同的键连续调用Value会返回相同的结果。
	//	//仅将上下文值用于传递过程和API边界的请求范围数据，而不用于将可选参数传递给函数。
	//	//关键字标识上下文中的特定值。希望在Context中存储值的函数通常会在全局变量中分配一个键，然后将该键用作context.WithValue和Context.Value的参数。密钥可以是任何支持相等性的类型。软件包应将键定义为未导出的类型，以免发生冲突。
	//	//定义上下文关键字的包应为使用该关键字存储的值提供类型安全的访问器：
	//	// //包用户定义了存储在上下文中的用户类型。
	//	// 	package user
	//	//
	//	// 	import "context"
	//	//
	//	// 	// User is the type of value stored in the Contexts.（用户是存储在上下文中的值的类型。）
	//	// 	type User struct {...}
	//	//
	//	// 	// key is an unexported type for keys defined in this package.（// key是此包中定义的密钥的未导出类型。）
	//	// 	// This prevents collisions with keys defined in other packages.（//这样可以防止与其他包中定义的键冲突。）
	//	// 	type key int
	//	//
	//	// 	// userKey是user.User上下文中的值的键。 它是未导出的； 客户端使用user.NewContext和user.FromContext而不是直接使用此键。
	//	// 	var userKey key
	//	//
	//	// 	// NewContext returns a new Context that carries value u.(// NewContext返回带有值u的新Context。)
	//	// 	func NewContext(ctx context.Context, u *User) context.Context {
	//	// 		return context.WithValue(ctx, userKey, u)
	//	// 	}
	//	//
	//	// 	// FromContext returns the User value stored in ctx, if any.(// FromContext返回存储在ctx中的User值（如果有）。)
	//	// 	func FromContext(ctx context.Context) (*User, bool) {
	//	// 		u, ok := ctx.Value(userKey).(*User)
	//	// 		return u, ok
	//	// 	}
	//	Value(key interface{}) interface{}
	//}


	// Background returns a non-nil, empty Context. It is never canceled, has no
	// values, and has no deadline. It is typically used by the main function,
	// initialization, and tests, and as the top-level Context for incoming
	// requests.
	// Background返回一个非空的Context。 它永远不会被取消，没有价值，也没有期限。 它通常由主要功能，初始化和测试使用，并用作传入请求的顶级上下文。
	backgroundContext := context.Background()
	//var backgroundContext context.Context//不能用空接口
	fun:= func() {
		fmt.Println("------")
		fmt.Println(backgroundContext.Value("id"))
		//截止日期返回应取消代表该上下文完成的工作的时间。 如果未设置截止日期，则截止日期返回ok == false。 连续调用Deadline会返回相同的结果。
		fmt.Println(backgroundContext.Deadline())
		//返回发送（存入值）完成的取消管道。
		fmt.Println(backgroundContext.Done())

		fmt.Println(backgroundContext.Err())
		fmt.Println("======")
	}
	fun()
	// QueryContext executes a query that returns rows, typically a SELECT.
	// The args are for any placeholder parameters in the query.
	// QueryContext执行一个查询，该查询返回行，通常是SELECT。
	// args用于查询中的任何占位符参数。
	rows, err = db.QueryContext(*&backgroundContext,"select * from person ")//语句2
	check_err_sql(err)
	defer rows.Close()

	//下面仅仅是针对一条结果集，如果有多条结果集的话这样遍历会出错
	for b:=rows.Next(); b != false;b=rows.Next() {
		fun()
		var p =person{}
		err := rows.Scan(&p.id, &p.name, &p.age)
		check_err_sql(err)
		fmt.Println(p)
	}
	//输出：
	//	------
	//	<nil>
	//	0001-01-01 00:00:00 +0000 UTC false
	//	<nil>
	//	<nil>
	//	======
	//	------
	//	<nil>
	//	0001-01-01 00:00:00 +0000 UTC false
	//	<nil>
	//	<nil>
	//	======
	//	{1 anko 20}
	//	------
	//	<nil>
	//	0001-01-01 00:00:00 +0000 UTC false
	//	<nil>
	//	<nil>
	//	======
	//	{2 gogo 32}



	fmt.Println("-------------------关于Context在sql中的应用---------------------------")

	rows, err = db.QueryContext(backgroundContext,"select * from person ")//语句2
	check_err_sql(err)
	defer rows.Close()

	//下面仅仅是针对一条结果集，如果有多条结果集的话这样遍历会出错
	for b:=rows.Next(); b != false;b=rows.Next() {
		fun()
		var p =person{}
		err := rows.Scan(&p.id, &p.name, &p.age)
		check_err_sql(err)
		fmt.Println(p)
	}


	fmt.Println("-------------------QueryRow返回一条（行）数据集（注意不是一条或者一行数据）而丢弃其他查询到的数据（底层QueryRowContext（））---------------------------")

	//// Row is the result of calling QueryRow to select a single row.
	////Row是调用QueryRow选择单个行的结果。
	//type Row struct {
	//	// One of these two will be non-nil:
	//	//这两个之一为非零：
	//	err  error // deferred error for easy chaining（延迟错误，便于链接）
	//	rows *Rows	//一个结果集中的多行（条）数据的对象
	//}

	// QueryRow执行的查询预期最多返回一行。
	// QueryRow始终返回非null值。 错误将一直延迟到调用Row的Scan方法。
	//如果查询未选择任何行，则“行扫描”将返回ErrNoRows。
	//否则，*行扫描将扫描所选的第一行，并丢弃其余的行。
	//底层使用的是db.QueryRowContext(context.Background(), query, args...)
	row:= db.QueryRow("select * from person ")//注意这里返回的是*Row对象，而不是像上面的*Rows对象

	//下面其实跟上面是一样的

	// QueryRowContext执行的查询预期最多返回一行。
	// QueryRowContext始终返回非null值。 错误将一直延迟到调用Row的Scan方法。
	//如果查询未选择任何行，则“行扫描”将返回ErrNoRows。
	//否则，*行扫描将扫描所选的第一行，并丢弃其余的行。
	//mybackgroundContext := new(emptyCtxMy)
	//row1:= db.QueryRowContext(mybackgroundContext,"select * from person ")//使用自己自定义的对象也是可以的
	row1:= db.QueryRowContext(backgroundContext,"select * from person where id=2")//既然我们上面已经定义了，我们就使用原来的对象好了
	// Scan copies the columns from the matched row into the values
	// pointed at by dest. See the documentation on Rows.Scan for details.
	// If more than one row matches the query,
	// Scan uses the first row and discards the rest. If no row matches
	// the query, Scan returns ErrNoRows.
	//Scan将匹配行中的列复制到dest指向的值中。 有关详细信息，请参见Rows.Scan上的文档。
	//如果与查询匹配的行超过一个，则Scan使用第一行并丢弃其余行。 如果没有与查询匹配的行，则Scan返回ErrNoRows。
	var p person
	var p1 person
	err = row.Scan(&p.id,&p.name,&p.age)
	check_err_sql(err)
	err = row1.Scan(&p1.id,&p1.name,&p1.age)
	check_err_sql(err)
	fmt.Println(p)
	fmt.Println(p1)
	//输出：
	//	{1 anko 20}
	//	{2 gogo 32}






}



type emptyCtxMy int

func (*emptyCtxMy) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*emptyCtxMy) Done() <-chan struct{} {
	return nil
}

func (*emptyCtxMy) Err() error {
	return nil
}

func (*emptyCtxMy) Value(key interface{}) interface{} {
	return nil
}

func check_err_sql(err error) {
	if err != nil {
		fmt.Println(err)
	}
}