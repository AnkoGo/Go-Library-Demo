package main

import (
	"fmt"
	"time"
)

func main99661() {
	fmt.Println("**********************************")
	//time包提供了时间的显示和测量用的函数。日历的计算采用的是公历。

	fmt.Println("***************time库中的类型*******************")
	t := time.Now()
	fmt.Println(t) //2019-10-08 02:04:48.1684409 +0800 CST m=+0.004922801

	// IsZero报告t是否表示零时刻，即1月1日，1年1月00:00:00 UTC。
	fmt.Println(t.IsZero()) //false

	// Location返回与t关联的时区信息。
	fmt.Println(t.Location())          //Local
	fmt.Println(t.UTC().Location())    //UTC
	fmt.Println(t.Location().String()) //Local

	// Unix返回t作为Unix时间，这是自1970年1月1日UTC以来经过的秒数。 结果不取决于与t关联的位置。
	fmt.Println(t.Unix()) //1570472202

	// UnixNano返回t作为Unix时间，即自UTC 1970年1月1日起经过的纳秒数。
	// 如果Unix时间（以纳秒为单位）不能用int64（日期在1678年之前或2262年之后）表示，
	// 则结果是不确定的。 注意，这意味着在零时间调用UnixNano的结果是不确定的。 结果不取决于与t关联的位置。
	fmt.Println(t.UnixNano()) //1570472202377239600(int64所能存储的最大位数为19位，所以这个值也是这么多位数)

	fmt.Println("**********************************")
	//Duration表示两个瞬间之间经过的时间段，以int64纳秒计数。 该表示将最大可表示持续时间限制为大约290年。
	var Dur time.Duration = 3600e9
	// String()返回以“ 72h3m0.5s”形式表示持续时间的字符串。
	//省略前导零单位。 作为一种特殊情况，持续时间少于一秒的格式使用较小的单位（毫秒，微秒或纳秒），
	// 以确保前导数字不为零。 零持续时间格式为0s。
	fmt.Println(Dur.String())       //1h0m0s
	fmt.Println(Dur.Hours())        //1
	fmt.Println(Dur.Minutes())      //60
	fmt.Println(Dur.Seconds())      //3600
	fmt.Println(Dur.Milliseconds()) //3600000(就是3600e3)
	fmt.Println(Dur.Microseconds()) //3600000000(就是3600e6)
	fmt.Println(Dur.Nanoseconds())  //3600000000000(就是3600e9)

	// Truncate返回将t向下舍入为d的倍数的结果（从零开始）。
	//如果d <= 0，则Truncate返回t，去除任何单调时钟读数，否则不改变。
	//
	//截断以自零时间开始的绝对持续时间对时间进行操作； 它不能按时间的显示形式进行操作。
	// 因此，取决于时间的位置，Truncate（Hour）可能会返回非零分钟的时间。
	tt1, _ := time.Parse("2006 Jan 02 15:04:05", "2012 Dec 07 12:15:30.918273645")
	trunc := []time.Duration{
		time.Nanosecond,  //1
		time.Microsecond, //1000
		time.Millisecond, //1e6
		time.Second,      //1e9
		2 * time.Second,  //2e9
		time.Minute,      //60e9
		10 * time.Minute, //600e9
		time.Hour,        //3600e9
	}
	fmt.Println(tt1.UnixNano()) //1354882530918273645
	for _, d := range trunc {
		fmt.Printf("t.Truncate(%6s) = %s\n", d, tt1.Truncate(d).Format("15:04:05.999999999"))
		//fmt.Printf("t.Truncate(%6s) = %s\n", d, tt1.Truncate(d).String())
	}
	//Format输出：
	//	t.Truncate(   1ns) = 12:15:30.918273645
	//	t.Truncate(   1µs) = 12:15:30.918273,向下取整
	//	t.Truncate(   1ms) = 12:15:30.918,向下取整
	//	t.Truncate(    1s) = 12:15:30,向下取整
	//	t.Truncate(    2s) = 12:15:30,向下取整为2的倍数，你可以尝试把上面的2改为4然后30会变为28的
	//	t.Truncate(  1m0s) = 12:15:00,向下取整
	//	t.Truncate( 10m0s) = 12:10:00,向下取整为10的倍数
	//	t.Truncate(1h0m0s) = 12:00:00,向下取整为1的倍数

	//String输出：
	//	t.Truncate(   1ns) = 2012-12-07 12:15:30.918273645 +0000 UTC
	//	t.Truncate(   1µs) = 2012-12-07 12:15:30.918273 +0000 UTC
	//	t.Truncate(   1ms) = 2012-12-07 12:15:30.918 +0000 UTC
	//	t.Truncate(    1s) = 2012-12-07 12:15:30 +0000 UTC
	//	t.Truncate(    2s) = 2012-12-07 12:15:30 +0000 UTC
	//	t.Truncate(  1m0s) = 2012-12-07 12:15:00 +0000 UTC
	//	t.Truncate( 10m0s) = 2012-12-07 12:10:00 +0000 UTC
	//	t.Truncate(1h0m0s) = 2012-12-07 12:00:00 +0000 UTC

	fmt.Println("**********************************")
	//下面讲解Duration中的trucate()
	var Dur111 time.Duration = 9530.918273645e9
	fmt.Println(Dur111.String())       //2h38m50.918273645s
	fmt.Println(Dur111.Hours())        //2.647477298234722
	fmt.Println(Dur111.Minutes())      //158.84863789408334
	fmt.Println(Dur111.Milliseconds()) //9530918
	fmt.Println(Dur111.Microseconds()) //9530918273
	fmt.Println(Dur111.Nanoseconds())  //9530918273645

	fmt.Println("**********************************")
	//截断返回将d朝零舍入到m的倍数的结果。
	//如果m <= 0，则Truncate不变地返回d。
	trunc111 := []time.Duration{
		time.Nanosecond,  //1
		time.Microsecond, //1000
		time.Millisecond, //1e6
		time.Second,      //1e9
		2 * time.Second,  //2e9
		3 * time.Second,  //3e9
		time.Minute,      //60e9
		10 * time.Minute, //600e9
		time.Hour,        //3600e9
	}
	for _, d := range trunc111 {
		fmt.Printf("t.Truncate(%6s) = %s\n", d, Dur111.Truncate(d))
	}
	//输出：
	//	t.Truncate(   1ns) = 2h38m50.918273645s,向下取整为1的倍数
	//	t.Truncate(   1µs) = 2h38m50.918273s,向下取整为1的倍数
	//	t.Truncate(   1ms) = 2h38m50.918s,向下取整为1的倍数
	//	t.Truncate(    1s) = 2h38m50s,向下取整为1的倍数
	//	t.Truncate(    2s) = 2h38m50s,向下取整为2的倍数
	//	t.Truncate(    3s) = 2h38m48s,向下取整为3的倍数
	//	t.Truncate(  1m0s) = 2h38m0s,向下取整为1的倍数
	//	t.Truncate( 10m0s) = 2h30m0s,向下取整为10的倍数
	//	t.Truncate(1h0m0s) = 2h0m0s,向下取整为1的倍数
	fmt.Println(trunc111) //[1ns 1µs 1ms 1s 2s 3s 1m0s 10m0s 1h0m0s],string（）定义了字符串和如何输出的格式和

	fmt.Println("**********************************")
	//下面讲解time中的round（），为了引出和便于理解Duration中的round()方法

	//舍入运算将t舍入到d的最接近倍数（从零开始）的结果。
	//中途值的舍入行为是向上舍入。
	//如果d <= 0，则Round返回t，去除了任何单调时钟读数，但未更改。
	//
	//Round以从零时间开始的绝对持续时间为准； 它不能按时间的显示形式进行操作。
	// 因此，Round（Hour）可能会返回非零分钟的时间，具体取决于时间的位置。
	t2244 := time.Date(0, 0, 0, 12, 15, 30, 918273645, time.UTC)
	round := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		4 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}
	for _, d := range round {
		fmt.Printf("t.Round(%6s) = %s\n", d, t2244.Round(d).Format("15:04:05.999999999"))
	}
	fmt.Println(round)
	//输出：
	//	t.Round(   1ns) = 12:15:30.918273645，四舍五入且向上取整为1的倍数
	//	t.Round(   1µs) = 12:15:30.918274，四舍五入且向上取整为1的倍数
	//	t.Round(   1ms) = 12:15:30.918，四舍五入且向上取整为1的倍数
	//	t.Round(    1s) = 12:15:31，四舍五入且向上取整为1的倍数
	//	t.Round(    2s) = 12:15:30，四舍五入且向上取整为2的倍数
	//	t.Round(    4s) = 12:15:32，四舍五入且向上取整为1的倍数
	//	t.Round(  1m0s) = 12:16:00，四舍五入且向上取整为1的倍数
	//	t.Round( 10m0s) = 12:20:00，四舍五入且向上取整为10的倍数
	//	t.Round(1h0m0s) = 12:00:00，四舍五入且向上取整为1的倍数
	//	[1ns 1µs 1ms 1s 2s 3s 1m0s 10m0s 1h0m0s]


	fmt.Println("**********************************")
	//截断返回将d朝零舍入到m的倍数的结果。
	//如果m <= 0，则Truncate不变地返回d。
	round111 := []time.Duration{
		time.Nanosecond,  //1
		time.Microsecond, //1000
		time.Millisecond, //1e6
		time.Second,      //1e9
		2 * time.Second,  //2e9
		3 * time.Second,  //3e9
		4 * time.Second,  //3e9
		time.Minute,      //60e9
		10 * time.Minute, //600e9
		time.Hour,        //3600e9
	}
	for _, d := range round111 {
		fmt.Printf("t.Truncate(%6s) = %s\n", d, Dur111.Round(d))
	}
	fmt.Println(round)
	//输出：
	//	t.Truncate(   1ns) = 2h38m50.918273645s，四舍五入且向上取整为1的倍数
	//	t.Truncate(   1µs) = 2h38m50.918274s，四舍五入且向上取整为1的倍数
	//	t.Truncate(   1ms) = 2h38m50.918s，四舍五入且向上取整为1的倍数
	//	t.Truncate(    1s) = 2h38m51s，四舍五入且向上取整为1的倍数
	//	t.Truncate(    2s) = 2h38m50s，四舍五入且向上取整为2的倍数
	//	t.Truncate(    3s) = 2h38m51s，四舍五入且向上取整为3的倍数
	//	t.Truncate(    4s) = 2h38m52s，四舍五入且向上取整为1的倍数，不取48则说明了向上取整
	//	t.Truncate(  1m0s) = 2h39m0s，四舍五入且向上取整为1的倍数
	//	t.Truncate( 10m0s) = 2h40m0s，四舍五入且向上取整为10的倍数
	//	t.Truncate(1h0m0s) = 3h0m0s，四舍五入且向上取整为1的倍数
	//	[1ns 1µs 1ms 1s 2s 4s 1m0s 10m0s 1h0m0s]

	fmt.Println("**********************************")
	//下面讲解time类

	var tt=time.Now()
	fmt.Println(tt)//2019-10-08 03:39:13.8313688 +0800 CST m=+0.027346501

	//这些是在Time.Format和time.Parse中使用的预定义布局。
	//布局中使用的参考时间是特定时间：
	// 2006年1月2日星期一1:04:05
	//这是Unix时间1136239445。由于MST是GMT-0700，因此参考时间可以视为01/02 03:04:05 PM '06 -0700
	//要定义自己的格式，请记下参考时间的格式设置；有关示例，请参见ANSIC，StampMicro或Kitchen等常数的值。
	//             该模型将演示参考时间的外观，以便Format和Parse方法可以将相同的转换应用于一般时间值。
	//
	//一些有效的布局对于time.Parse来说是无效的时间值，这是因为格式，例如_表示空格填充，Z表示区域信息。
	//
	//在格式字符串中，下划线_表示如果后面的数字（一天）有两位数字，可以用一位数字代替空格；与固定宽度的Unix时间格式兼容。
	//
	//小数点后跟一个或多个零表示小数秒，打印到给定的小数位数。小数点后跟一个或多个9代表一个小数秒，打印到给定的小数位数，并删除了尾随的零。
	//（仅）分析时，即使布局不表示其存在，输入也可能在秒字段之后紧随其后的分数字段包含小数。在这种情况下，小数点后跟最大的数字序列被解析为小数秒。
	//
	//数字时区偏移格式如下：
	// -0700±hhmm
	// -07：00±hh：mm
	// -07±hh
	//用Z替换格式中的符号会触发ISO 8601打印Z的行为，而不是UTC区域的偏移量。从而：
	// Z0700 Z或±hhmm
	// Z07：00 Z或±hh：mm
	// Z07 Z或±hh
	//
	//公认的星期几格式为"Mon" 和 "Monday"
	//公认的月份格式为“ Jan”和“ January”。
	//
	//格式2，_2和02分别是每月的无填充，有空格和零填充。格式__2和002是一年中的三个字符的空格和零填充；没有一年中未填充的日期格式。
	//
	//格式字符串中未被识别为参考时间的文本将在Format期间逐字回显，并有望逐字出现在Parse的输入中。
	//
	// Time.Format的可执行示例详细说明了布局字符串的工作原理，是一个很好的参考。
	//
	//请注意，RFC822，RFC850和RFC1123格式应仅应用于本地时间。将它们应用于UTC时间将使用“ UTC”作为时区缩写，而严格地说，在这种情况下，那些RFC要求使用“ GMT”。
	//对于坚持使用该格式的服务器，通常应使用RFC1123Z代替RFC1123，并且对于新协议，应首选RFC3339。
	// RFC3339，RFC822，RFC822Z，RFC1123和RFC1123Z对于格式化非常有用；
	//与time一起使用时，解析不接受RFC允许的所有时间格式。
	// RFC3339Nano格式从seconds字段中删除了结尾的零，因此一旦格式化就可能无法正确排序。

	//下面是time为格式化定义的一些内置常量
	//const (
	//	ANSIC       = "Mon Jan _2 15:04:05 2006"
	//	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	//	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	//	RFC822      = "02 Jan 06 15:04 MST"
	//	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	//	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	//	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	//	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	//	RFC3339     = "2006-01-02T15:04:05Z07:00"
	//	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	//	Kitchen     = "3:04PM"
	//	// Handy time stamps.
	//	Stamp      = "Jan _2 15:04:05"
	//	StampMilli = "Jan _2 15:04:05.000"
	//	StampMicro = "Jan _2 15:04:05.000000"
	//	StampNano  = "Jan _2 15:04:05.000000000"
	//)
	//
	//const (
	//	_                        = iota
	//	stdLongMonth             = iota + stdNeedDate  // "January"
	//	stdMonth                                       // "Jan"
	//	stdNumMonth                                    // "1"
	//	stdZeroMonth                                   // "01"
	//	stdLongWeekDay                                 // "Monday"
	//	stdWeekDay                                     // "Mon"
	//	stdDay                                         // "2"
	//	stdUnderDay                                    // "_2"
	//	stdZeroDay                                     // "02"
	//	stdUnderYearDay                                // "__2"
	//	stdZeroYearDay                                 // "002"
	//	stdHour                  = iota + stdNeedClock // "15"
	//	stdHour12                                      // "3"
	//	stdZeroHour12                                  // "03"
	//	stdMinute                                      // "4"
	//	stdZeroMinute                                  // "04"
	//	stdSecond                                      // "5"
	//	stdZeroSecond                                  // "05"
	//	stdLongYear              = iota + stdNeedDate  // "2006"
	//	stdYear                                        // "06"
	//	stdPM                    = iota + stdNeedClock // "PM"
	//	stdpm                                          // "pm"
	//	stdTZ                    = iota                // "MST"
	//	stdISO8601TZ                                   // "Z0700"  // prints Z for UTC
	//	stdISO8601SecondsTZ                            // "Z070000"
	//	stdISO8601ShortTZ                              // "Z07"
	//	stdISO8601ColonTZ                              // "Z07:00" // prints Z for UTC
	//	stdISO8601ColonSecondsTZ                       // "Z07:00:00"
	//	stdNumTZ                                       // "-0700"  // always numeric
	//	stdNumSecondsTz                                // "-070000"
	//	stdNumShortTZ                                  // "-07"    // always numeric
	//	stdNumColonTZ                                  // "-07:00" // always numeric
	//	stdNumColonSecondsTZ                           // "-07:00:00"
	//	stdFracSecond0                                 // ".0", ".00", ... , trailing zeros included
	//	stdFracSecond9                                 // ".9", ".99", ..., trailing zeros omitted
	//
	//	stdNeedDate  = 1 << 8             // need month, day, year
	//	stdNeedClock = 2 << 8             // need hour, minute, second
	//	stdArgShift  = 16                 // extra argument in high bits, above low stdArgShift
	//	stdMask      = 1<<stdArgShift - 1 // mask out argument
	//)


	tt_1:=time.ANSIC
	fmt.Println(tt_1)//Mon Jan _2 15:04:05 2006

	// ParseDuration解析一个持续时间字符串。
	//持续时间字符串是可能带符号的十进制数字序列，每个序列都有可选的分数和一个单位后缀，例如“ 300ms”，“-1.5h”或“ 2h45m”。
	//有效时间单位为“ ns”，“ us”（或“ µs”），“ ms”，“ s”，“ m”，“ h”。
	//说白了就是字符串转Duration
	fmt.Println(Dur111)//2h38m50.918273645s
	tt_2,err:=time.ParseDuration(Dur111.String())
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(tt_2)//2h38m50.918273645s

	//自1970年1月1日UTC以来，Unix返回与给定的Unix时间，秒秒和nsec纳秒相对应的本地时间。
	//将nsec传递到[0，999999999]范围之外是有效的。
	//并非所有秒值都具有对应的时间值。 这样的值之一是1 << 63-1（最大的int64值）。
	fmt.Println(tt.Unix())//1570479317
	fmt.Println(tt.UnixNano())//1570479317209147400
	T1:=time.Unix(tt.Unix(),tt.UnixNano())
	T2:=time.Unix(tt.Unix(),0)
	fmt.Println(T1)//2069-07-14 00:30:34.2091474 +0800 CST
	fmt.Println(T2)//2019-10-08 04:15:17 +0800 CST
	fmt.Println(T2.Add(time.Duration(tt.UnixNano())))//2069-07-14 00:30:34.2091474 +0800 CST

	fmt.Println("-----------------------------------------------")

	// Sub返回持续时间t-u。 如果结果超过可以在“持续时间”中存储的最大（或最小）值，则将返回最大（或最小）持续时间。
	//要计算持续时间d的t-d，请使用t.Add（-d）。
	var tt_now=time.Now()
	var tt_sub=time.Date(2019,10,8,4,20,0,0,time.Local)

	dur_sub:=tt_now.Sub(tt_sub)
	fmt.Println(dur_sub)//4m52.2310735s
	new_now:=tt_sub.Add(dur_sub)//可以用负数，这样就是减了
	fmt.Println(new_now)//2019-10-08 04:24:52.2310735 +0800 CST
	fmt.Println(tt_now.Equal(new_now))//true

	fmt.Println("-----------------------------------------------")

	var now_time=time.Now()
	//Location()会将时刻映射到当时使用的区域。
	//通常，Location()表示在某个地理区域（例如中欧的CEST和CET）中使用的时间偏移的集合。

	//struct如下：
	//	type Location struct {
	//		name string
	//		zone []zone
	//		tx   []zoneTrans
	//
	//		//大多数查询将针对当前时间。
	//		//为了避免通过tx进行二进制搜索，请保留一个静态的单元素缓存，该缓存在创建位置时会提供正确的区域。
	//		//如果cacheStart <= t <cacheEnd，则查找可以返回cacheZone。
	//		// cacheStart和cacheEnd的单位是自UTC 1970年1月1日以来的秒数，以将参数与lookup匹配。
	//		cacheStart int64
	//		cacheEnd   int64
	//		cacheZone  *zone
	//	}
	//
	//	//区域代表单个时区，例如CEST或CET。
	//	type zone struct {
	//		name   string //缩写名称“ CET”
	//		offset int    // UTC以东的秒钟
	//		isDST  bool   // 该区域是夏时制吗？
	//	}
	//
	//	// zoneTrans表示单个时区转换。
	//	type zoneTrans struct {
	//		when         int64 // 转换时间，自1970 GMT以来的秒数
	//		index        uint8 // 当时生效的区域的索引
	//		isstd, isutc bool  // 被忽略-不知道这些是什么意思
	//	}

	// FixedZone返回一个位置，该位置始终使用给定的区域名称和偏移量（UTC以东的秒数）。

	fmt.Println(now_time.Location())//Local(struct中的name字段的值)
	fmt.Println(now_time.Location().String())//Local

	// LoadLocation返回具有给定名称的Location。
	//
	//如果名称是“”或“ UTC”，则LoadLocation返回UTC。
	//如果名称为“ Local”，则LoadLocation返回Local。
	//
	//否则，该名称将被视为与IANA时区数据库中的文件相对应的位置名称，例如“ America / New_York”。
	//
	// LoadLocation所需的时区数据库可能并非在所有系统上都存在，尤其是在非Unix系统上。
	// LoadLocation查找由ZONEINFO环境变量命名的目录或未压缩的zip文件（如果有），
	// 然后查找Unix系统上的已知安装位置，最后查找$ GOROOT / lib / time / zoneinfo.zip。
	fmt.Println(time.LoadLocation("UTC"))//Local <nil>
	fmt.Println(time.LoadLocation("Local"))//Local <nil>

	// LoadLocationFromTZData返回一个具有给定名称的位置，该位置是从IANA时区数据库格式的数据中初始化的。
	//数据应采用标准IANA时区文件的格式（例如，Unix系统上/ etc / localtime的内容）。
	ls:=make([]byte,10,20)
	fmt.Println(time.LoadLocationFromTZData("Local",ls))//UTC malformed time zone information(时区信息格式错误)
	//不知道怎么用

	//睡眠至少在持续时间d内暂停当前goroutine。
	//持续时间为负数或零会导致休眠立即返回。
	time.Sleep(time.Duration(1e9))
	time.Sleep(1e9)
	fmt.Println("2s结束")

	fmt.Println("-----------------------------------------------")
	fmt.Println(time.Now())//2019-10-08 05:15:50.3213659 +0800 CST m=+2.029073301
	Ctime:=<-time.After(2e9)
	fmt.Println(time.Now())//2019-10-08 05:15:52.3216225 +0800 CST m=+4.029329901
	fmt.Println(Ctime)//2019-10-08 05:15:52.3216225 +0800 CST m=+4.029329901

	fmt.Println("-----------------------------------------------")
	TT:=time.AfterFunc(3e9,fmtPrintln)
	fmt.Println(TT)//&{<nil> {7085184 0 36224267820800 0 0x495ec0 0x4b8df0 0}}
	fmt.Println(TT.C)//<nil>
	fmt.Println("等待1s准备fmtPrintln函数")
	//{7085184 0 36224267820800 0 0x495ec0 0x4b8df0 0}对应如下：
	//		tb uintptr
	//		i  int
	//		when   int64
	//		period int64
	//		f      func(interface{}, uintptr) // NOTE: must not be closure
	//		arg    interface{}
	//		seq    uintptr


	//重置将计时器更改为持续时间d。
	//如果计时器处于活动状态，则返回true；如果计时器已过期或已停止，则返回false。
	//
	//只能在通道耗尽的停止或到期计时器上调用复位。
	//如果程序已经从t.C接收到一个值，则已知计时器已到期且通道已耗尽，因此可以直接使用t.Reset。
	//如果程序尚未从t.C接收到值，则必须停止计时器，并且-如果Stop报告停止之前计时器已到期，则该通道显式耗尽：
	//
	//如果！t.Stop（）{
	// <-t.C
	//}
	// t.Reset（d）
	//
	//此操作不应与定时器通道的其他接收同时进行。
	//
	// //请注意，无法正确使用Reset的返回值，因为在耗尽通道和新计时器到期之间存在竞争条件。
	//如上所述，应始终在停止或过期的通道上调用复位。
	//返回值的存在是为了保持与现有程序的兼容性。
	fmt.Println(TT.Reset(1e9))//true,这个时间是重置时间，不要超过主协程的睡眠时间6s

	//停止可阻止计时器触发。
	//如果调用停止了计时器，则返回true；如果计时器已过期或已停止，则返回false。
	// Stop不会关闭通道，以防止从通道读取失败。
	//为确保在调用Stop之后通道为空，请检查返回值并清空通道。
	//例如，假设尚未从t.C收到程序：
	//
	//如果！t.Stop（）{
	// <-t.C
	//}
	//
	//这不能与定时器的其他接收同时进行
	//频道。
	//对于使用AfterFunc（d，f）创建的计时器，如果t.Stop返回false，则说明该计时器已经到期，并且函数f已在其自己的goroutine中启动；
	//停止不等待f完成后再返回。
	//如果调用者需要知道f是否完成，则必须与f明确协调。
	fmt.Println(TT.Stop())//true,这个时间是重置时间，不要超过主协程的睡眠时间6s

	time.Sleep(6e9)//等待fmtPrintln（）函数执行
	fmt.Println("主协程结束")

	fmt.Println("------------------------------------")
	//Since()返回从t开始经过的时间。
	//这是time.Now().Sub（t）的简写。
	tt_sub111:=time.Date(2019,10,8,5,58,0,0,time.Local)
	fmt.Println(time.Since(tt_sub111))//3m8.9043571s

	//Until返回t之前的持续时间。
	//它是t.Sub（time.Now（））的简写。
	fmt.Println(time.Until(tt_sub111))//-3m8.9043571s


	// FixedZone返回一个位置，该位置始终使用给定的区域名称和偏移量（UTC以东的秒数）。
	//位置会将时刻映射到当时使用的区域。
	//通常，位置表示在某个地理区域（例如中欧的CEST和CET）中使用的时间偏移的集合。
	//time类型中保存着*Location的字段，下面是这个字段的解释：
	//	loc指定应用于确定与此时间对应的分钟，小时，月，日和年的位置。
	//	nil位置表示UTC。
	//	所有UTC时间都以loc == nil表示，从不loc ==＆utcLoc表示。
	fmt.Println(time.Now())//2019-10-08 06:04:12.4773693 +0800 CST m=+10.034079001
	locat:=time.FixedZone("TTT",800)
	fmt.Println(locat)//TTT

	//可以设置相同名字的Location，我去
	locat111:=time.FixedZone(time.Local.String(),800)
	fmt.Println(locat111)//Local

	fmt.Println(time.Local)//Local
	fmt.Println(time.Local==locat111)//判断对象而不是对象的值（字符串）

	// Tick是NewTicker的便捷包装，仅提供对滴答通道的访问。 尽管Tick对于不需要关闭Ticker的客户端很有用，
	// 但是请注意，没有办法关闭它，底层的Ticker无法被垃圾收集器恢复； 它“泄漏”。
	//与NewTicker不同，如果d <= 0，Tick将返回nil。
	n_time:=<-time.Tick(2e9)
	fmt.Println(n_time)//2019-10-08 06:16:41.993895 +0800 CST m=+12.032315901


	// NewTicker返回一个新的Ticker，其中包含一个通道，该通道将每隔一个时间t发送当前的时间，该时间t由duration参数指定。
	//调整间隔或滴答滴答声以弥补接收速度慢的问题。
	//持续时间d必须大于零； 如果没有，NewTicker将惊慌。
	//Stop() ticker以释放关联的资源。即使在时间到后也要这样关闭。
	//Ticker是心脏的意思，他的功能也如同心脏一样搏动
	ticker := time.NewTicker(2e9)
	//ticker.Stop()//会导致下面的无限等待
	fmt.Println(<-ticker.C)
	//Stop()关闭一个Ticker。 停止后，将不再发送任何ticks。
	// Stop不会关闭通道，以防止同时从通道读取goroutine时看到错误的“滴答”。
	fmt.Println("===",<-ticker.C)

	fmt.Println("===",<-ticker.C)
	ticker.Stop()
	//输出如下：
	//	2019-10-08 06:26:20.9176751 +0800 CST m=+12.034265101
	//	2019-10-08 06:26:22.9184231 +0800 CST m=+14.035013101
	//	=== 2019-10-08 06:26:24.9186796 +0800 CST m=+16.035269601
	//	=== 2019-10-08 06:26:26.9179633 +0800 CST m=+18.034553301


	// ParseInLocation类似于Parse，但在两个重要方面有所不同。
	//首先，在没有时区信息的情况下，Parse将时间解释为UTC；
	// ParseInLocation将时间解释为给定位置。
	//第二，当给定区域偏移量或缩写时，Parse尝试将其与本地位置进行匹配； ParseInLocation使用给定的位置。
	//location_T, err99 := time.ParseInLocation("2006-01-02 15:04:06", "2017-08-03 15:04:05", time.UTC)//2005-08-03 15:04:00 +0000 UTC, layout只能使用这个数字"2006-01-02 15:04:06"
	location_T, err99 := time.ParseInLocation("2006-01-02 15:04:05", "2019-10-08 15:04:08", time.UTC)//2019-10-08 不能写成2019-10-8
	if err99 != nil{
		fmt.Println(err99)
	}
	fmt.Println(location_T)//2019-10-08 15:04:08 +0000 UTC,


	location_T1, err991 := time.ParseInLocation("2006-01-02 15:04:05", "2019-10-08 15:04:08", time.Local)//
	if err991 != nil{
		fmt.Println(err991)
	}
	fmt.Println(location_T1)//2019-10-08 15:04:08 +0800 CST

	location_T2, err992 := time.ParseInLocation("2006-01-02 15:04:05", "2019-10-08 15:04:08", locat)//
	if err992 != nil{
		fmt.Println(err992)
	}
	fmt.Println(location_T2)//2019-10-08 15:04:08 +0013 TTT

	fmt.Println("============================================")
	// NewTimer创建一个新的Timer，它将至少在持续时间d之后在其通道上发送当前时间。
	//Timer结构体如下：
		//Timer类型代表一个事件。
		//计时器到期时，当前时间将在C上发送，除非该计时器是由AfterFunc创建的。
		//必须使用NewTimer或AfterFunc创建一个计时器。
		//type Timer struct {
		//	C <-chan Time
		//	r runtimeTimer
		//}

	timer := time.NewTimer(2e9)
	fmt.Println(time.Now())//2019-10-08 06:50:36.5608996 +0800 CST m=+18.034065201
	//重置将计时器更改为持续时间d。
	//如果计时器处于活动状态，则返回true；如果计时器处于活动状态，则返回false
	//已过期或已停止。
	//
	//只能在通道耗尽的停止或到期计时器上调用复位。
	//如果程序已经从t.C接收到一个值，则已知计时器已到期且通道已耗尽，因此可以直接使用t.Reset。
	//如果程序尚未从t.C接收到值，则必须停止计时器，并且-如果Stop报告停止之前计时器已到期，则该通道显式耗尽：
	//
	//如果！t.Stop（）{
	// <-t.C
	//}
	// t.Reset（d）
	//
	//此操作不应与定时器通道的其他接收同时进行。
	// //请注意，无法正确使用Reset的返回值，因为在耗尽通道和新计时器到期之间存在竞争条件。
	//如上所述，应始终在停止或过期的通道上调用复位。
	//返回值的存在是为了保持与现有程序的兼容性。
	fmt.Println(timer.Reset(5e9))//true
	fmt.Println(<-timer.C)//2019-10-08 06:50:41.5634982 +0800 CST m=+23.036663801
	//fmt.Println(<-timer.C)//一直等待，无值取，阻塞在这里，后面无法执行
	//fmt.Println(<-timer.C)
	//time.Sleep(6e9)

	//停止可阻止计时器触发。
	//如果调用停止了计时器，则返回true；如果计时器已过期或已停止，则返回false。
	// Stop不会关闭通道，以防止从通道读取失败。
	//为确保在调用Stop之后通道为空，请检查返回值并清空通道。
	//例如，假设尚未从t.C收到程序：
	//if!t.Stop(){
	// 	<-t.C
	//}
	//这不能与定时器通道的其他接收同时进行。
	//对于使用AfterFunc（d，f）创建的计时器，如果t.Stop返回false，则说明该计时器已经到期，并且函数f已在其自己的goroutine中启动；
	//停止不等待f完成后再返回。
	//如果调用者需要知道f是否完成，则必须与f明确协调。
	fmt.Println(timer.Stop())



	//一个工作日指定一周中的某一天（星期日= 0，...）。
	var i time.Weekday//默认是0
	fmt.Println(i)//Sunday
	for _, value := range []int{0,1,2,3,4,5,6}{
		var i1 time.Weekday=time.Weekday(value)
		fmt.Println(i1)
	}
	//输出如下：
	//	Sunday
	//	Monday
	//	Tuesday
	//	Wednesday
	//	Thursday
	//	Friday
	//	Saturday


	fmt.Println("============================================")
	t66:=time.Now()
	// Clock返回由t指定的日期内的小时，分钟和秒。
	fmt.Println(t66.Clock())  //7 13 9
	// Weekday返回由t指定的星期几。
	fmt.Println(t66.Weekday())  //Tuesday,底层调用abs()
	fmt.Println(t66.UTC())  //2019-10-07 23:13:09.9757082 +0000 UTC
	fmt.Println(t66.Local())  //2019-10-08 07:13:09.9757082 +0800 CST(和UTC相差8小时)
	fmt.Println(t66)  //2019-10-08 07:13:09.9757082 +0800 CST m=+23.040147301
	// YearDay返回由t指定的一年中的日期，平年的范围为[1,365]，闰年的范围为[1,366]。
	fmt.Println(t66.YearDay())  //281

	// In返回代表同一时刻的t的副本，但副本的位置信息设置为loc以便显示。
	//如果loc为nil则处于恐慌状态。
	fmt.Println(t66.In(time.Local))  //2019-10-08 07:13:09.9757082 +0800 CST
	fmt.Println(t66.In(time.UTC))  //2019-10-07 23:13:09.9757082 +0000 UTC
	fmt.Println(t66.In(locat))  //2019-10-07 23:26:29.9757082 +0013 TTT

	//之前报告时间t是否早于u。
	fmt.Println(t66.Before(tt_sub))//false
	fmt.Println(tt_sub.Before(t66))//true
	// After报告时刻t是否在u之后。
	fmt.Println(t66.Before(tt_sub))//true
	fmt.Println(tt_sub.Before(t66))//false

	fmt.Println(t66.AddDate(1,1,1))//2020-11-09 07:21:24.8716063 +0800 CST

	// AppendFormat类似于Format，但将文本表示形式附加到b并返回扩展的缓冲区。
	ls567:=make([]byte,0,16)
	fmt.Println(t66.AppendFormat(ls567,"2006-01-02 15:04:05"))//[50 48 49 57 45 49 48 45 48 56 32 48 55 58 50 55 58 51 54]
	ls5677:=make([]byte,0)
	format_ls := t66.AppendFormat(ls5677, "2006-01-02 15:04:05")
	fmt.Println(format_ls)//[50 48 49 57 45 49 48 45 48 56 32 48 55 58 50 55 58 51 54]
	fmt.Println(string(format_ls))//2019-10-08 07:27:36

	// Format返回根据布局格式化的时间值的文本表示形式，该时间值通过显示参考时间（定义为Mon Jan 2 15:04:05 -0700 MST 2006）
	// （如果是该值）来显示，从而定义了格式 ; 它作为所需输出的示例。 然后，将相同的显示规则应用于时间值。
	//
	//通过在布局字符串的秒部分的末尾添加句点和零来表示小数秒，如“ 15：04：05.000”中所示，以毫秒精度设置时间戳。
	//
	//预定义的布局ANSIC，UnixDate，RFC3339等描述了参考时间的标准和便捷表示形式。 有关格式和参考时间的定义的更多信息，
	// 请参见ANSIC文档以及此程序包定义的其他常量。
	fmt.Println(t66.Format("2006-01-02 15:04:05"))//2019-10-08 07:27:36
	fmt.Println(string(49))//1


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


	fmt.Println("==================利用god模块对时间的编码解码==========================")
	//纳秒返回t所指定的秒内的纳秒偏移，范围[0，999999999]。
	fmt.Println(t66.Nanosecond())//430666700
	ls969,_:=t66.GobEncode()//这个方法实现了god中的接口，将在go2删除，如果你想知道怎么学习实现这个接口的话可以查看这个方法怎么写
	fmt.Println(ls969)//[1 0 0 0 14 213 45 198 21 25 171 115 204 1 224]
	fmt.Println(string(ls969))//   �-��s��

	fmt.Println(t66.GobDecode(ls969))//<nil>，这个方法实现了god中的接口，将在go2删除，
	fmt.Println(ls969)//[1 0 0 0 14 213 45 198 21 25 171 115 204 1 224]
	fmt.Println(string(ls969))//   �-��s��

	fmt.Println("============================================")
	// ISOWeek返回出现t的ISO 8601年和周号。
	//周的范围是1到53。n年的1月1日到1月03日可能属于n-1年的52周或53周，12月29日到12月31日可能属于n + 1年的第1周。
	fmt.Println(t66.ISOWeek())//2019 41


}
func fmtPrintln()  {
	fmt.Println("执行fmtPrintln中，需要2s")
	time.Sleep(2e9)
}
