package main

import (
	"fmt"
	"time"
)

func main6666() {
	fmt.Println("**********************************")
	//time包提供了时间的显示和测量用的函数。日历的计算采用的是公历。

	fmt.Println("***************time库中的类型*******************")
	t:=time.Now()
	fmt.Println(t)  //2019-10-08 00:47:18.0702494 +0800 CST m=+0.005861001
	fmt.Println(t.Year())  //2019
	fmt.Println(t.Month())  //October
	fmt.Println(t.Day())  //8
	fmt.Println(t.Hour())  //0
	fmt.Println(t.Minute()) //47
	fmt.Println(t.Second())  //18

	fmt.Println(t.Date())  //2019 October 8
	fmt.Println(t.String())  //2019-10-08 00:49:00.9213353 +0800 CST m=+0.004883101,其实这里写不写String()都是一样的会自动调用这个函数
	fmt.Println(t.UTC())  //2019-10-07 16:51:44.0672758 +0000 UTC
	//Zone()计算在时间t生效的时区，并返回该区域的缩写名称（例如“ CET”）及其在UTC以东的秒数内的偏移量。
	fmt.Println(t.Zone())  //CST 28800

	//相等报告t和u是否代表同一时刻。
	//即使它们位于不同的位置，两次也可以相等。
	//例如，6:00 +0200 CEST和4:00 UTC相等。
	//请参阅有关时间类型的文档，以了解将==与
	//时间值； 大多数代码应改用Equal。
	fmt.Println(t.Equal(t))  //true
	fmt.Println(t.Equal(t.UTC()))  //true


	t1 := time.Date(2009, time.November, 10, 23, 2, 3, 6, time.UTC)
	//Local()返回t，并将位置设置为本地时间。
	fmt.Printf("Go launched at %s\n", t1.Local())//Go launched at 2009-11-11 07:02:03.000000006 +0800 CST
	fmt.Println(t1)//2009-11-10 23:02:03.000000006 +0000 UTC
	fmt.Println(t.Equal(t1.Local()))  //false
	fmt.Println(t.Equal(t1))  //false
	fmt.Println(t1.Equal(t1.Local()))  //true


	// Parse解析格式化的字符串并返回它表示的时间值。
	//布局通过显示参考时间来定义格式，
	//定义为
	// 2006年1月2日星期一1:04:05 -0700
	//如果是值则将被解释；它用作输入格式的示例。然后将对输入字符串进行相同的解释。
	//
	//预定义的布局ANSIC，UnixDate，RFC3339等描述了参考时间的标准和便捷表示形式。有关格式和参考时间的定义的更多信息，请参见ANSIC文档以及此程序包定义的其他常量。
	//另外，Time.Format的可执行示例详细说明了布局字符串的工作原理，是一个很好的参考。
	//
	//从值中省略的元素假定为零，或者当不可能为零时为1，因此解析“ 3:04 pm”将返回与世界标准时间1年1月1日，15：44：00对应的时间（请注意，因为年为0，此时间早于零时间）。
	//年必须在0000..9999的范围内。将检查星期几的语法，否则将忽略该语法。
	//
	//在没有时区指示符的情况下，Parse以UTC返回时间。
	//
	//解析时间偏移为-0700的时间时，如果偏移量对应于当前位置（本地）使用的时区，则Parse在返回的时间中使用该位置和时区。否则，它将时间记录为处于伪造位置，时间固定在给定的区域偏移量。
	//
	//解析带有MST等区域缩写的时间时，如果该区域缩写在当前位置具有已定义的偏移量，则使用该偏移量。
	//区域缩写“ UTC”被识别为UTC，与位置无关。
	//如果区域缩写未知，则Parse会将时间记录为在给定区域缩写和零偏移量的虚构位置中。
	//此选择意味着可以使用相同的布局无损地解析和重新格式化这样的时间，但是表示中使用的确切瞬间将因实际区域偏移而有所不同。为避免此类问题，请首选使用数字区域偏移量的时间布局或使用ParseInLocation。


	// longFormLayout通过示例显示了如何在所需布局中表示参考时间。
	const longFormLayout = "Jan 2, 2006 at 3:04pm (MST)"
	t2, _ := time.Parse(longFormLayout, "Feb 3, 2013 at 7:54pm (PST)")
	fmt.Println(t2)  //2013-02-03 19:54:00 +0000 PST
	//shortFormLayout是参考时间以所需布局表示的另一种方式； 它没有时区。
	//注意：如果没有显式区域，则返回UTC时间。
	const shortFormLayout = "2006-Jan-02"
	t3, _ := time.Parse(shortFormLayout, "2013-Feb-03")
	fmt.Println(t3)  //2013-02-03 00:00:00 +0000 UTC

	const Layout02 = "15:04:05"
	t5, _ := time.Parse(Layout02, "15:04:05")
	fmt.Println(t5)  //0001-01-01 00:00:00 +0000 UTC，没给出的值为0，不能为0的就为1，比如这里的0001-01-01

	const Layout03 = "2006-Jan-02"
	t6, err:= time.Parse(Layout03, "2017-Feb-03")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(t6)  //2017-02-03 00:00:00 +0000 UTC

	const Layout04 = "2006-Jan-02 15:04:05"
	t7, err:= time.Parse(Layout04, "2017-Feb-03 15:04:05")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(t7)  //2017-02-03 15:04:05 +0000 UTC

	const Layout05 = "2006-1-02 15:04:05"
	t8, err:= time.Parse(Layout05, "2017-7-03 15:04:05")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(t8)  //2017-07-03 15:04:05 +0000 UTC

	const Layout06 = "2006-01-02 15:04:05"//只能是这个字符串的这个时间，下同，上同
	t9, err:= time.Parse(Layout06, "2017-08-03 15:04:05")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(t9)  //2017-08-03 15:04:05 +0000 UTC

	const Layout07 = "2006-01-02"
	t0, err:= time.Parse(Layout07, "2017-09-03")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(t0)  //2017-09-03 00:00:00 +0000 UTC

	const Layout08 = "15:04:05"
	t91, err:= time.Parse(Layout08, "15:04:05")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(t91)  //0000-01-01 15:04:05 +0000 UTC

	const Layout91 = "03:04pm"
	t92, err03 := time.Parse(Layout91, "07:18pm")
	if err03 != nil{
		fmt.Println(err03)
	}
	fmt.Println(t92)  //0000-01-01 19:18:00 +0000 UTC


	//下面是错误示例
	const Layout01 = "03:04:03pm"
	t4, err02 := time.Parse(Layout01, "11:04:05am")
	if err02 != nil{
		fmt.Println(err02)//
	}
	fmt.Println(t4)  //0000-01-01 19:18:00 +0000 UTC,虽然没报错，但是输出的时间非常的诡异
	//下面为了测试这个时间是否是utc时间
	tt:=time.Date(1,1,1,11,04,05,0,time.Local)
	fmt.Println(tt)//0001-01-01 11:04:05 +0800 CST
	fmt.Println(tt.UTC())//0001-01-01 03:04:05 +0000 UTC



	//下面是错误示例
	const Layout_11 = "03:04:03pm"
	t933, err0222 := time.Parse(Layout_11, "11:04:05")
	if err0222 != nil{
		fmt.Println(err0222)  //parsing time "11:04:05" as "03:04:03pm": cannot parse "" as "pm"
	}
	fmt.Println(t933)  //0001-01-01 00:00:00 +0000 UTC



	fmt.Println("***************time库中的函数*******************")
	//单调时间报告为相对于startNano的偏移量。
	//我们将startNano初始化为runtimeNano（）-1，以便在单调时间分辨率相当低的系统上（例如，Windows 2008的默认分辨率为15ms），我们避免将单调时间报告为0。
	//（调用者可能希望将0用作“未设置时间”。）
	fmt.Println(time.Now())  //2019-10-07 22:48:11.3174417 +0800 CST m=+0.004882201
	//+0800也就是时差，这里表示在中国北京或者同纬度上的时间
	//CST =Central Standard Time (美国)中部标准时间
	//m=+0.004882201表示自程序启动以来进行了多少s的时间


	// runtimeNano返回运行时时钟的当前值（以纳秒为单位）。

}
