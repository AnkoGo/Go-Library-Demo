package main

import (
	"fmt"
	"math"
)

func main3452() {
	fmt.Println("----------求绝对值-----------")

	fmt.Println(math.Abs(1.1))
	fmt.Println(math.Abs(-1.1))
	fmt.Println(math.Abs(-+1.1))
	fmt.Println(math.Abs(+1.1))
	//输出：
	//	1.1
	//	1.1
	//	1.1
	//	1.1

	fmt.Println("----------求余弦与反余弦-----------")
	//角A的邻边比斜边 叫做角A的余弦，记作 cosA（由余弦英文cosine简写 ）
	// Acos返回弧度的x的反余弦值。
	//特殊情况是：
	//如果x <-1或x> 1则Acos（x）= NaN
	fmt.Println(math.Acos(0.8))

	// Cos返回弧度参数x的余弦值。
	//特殊情况是：
	// Cos（±Inf）= NaN
	// Cos（NaN）= NaN
	fmt.Println(math.Cos(120))
	//输出：
	//	0.6435011087932843
	//	0.8141809705265618

	fmt.Println("----------求正弦与反正弦-----------")
	// Asin以弧度返回x的反正弦值。
	//特殊情况是：
	// Asin（±0）=±0
	//如果x <-1或x> 1则Asin（x）= NaN
	fmt.Println(math.Asin(0.5))//0.5235987755982989

	// Sin返回弧度参数x的正弦值。
	//特殊情况是：
	// Sin（±0）=±0
	// Sin（±Inf）= NaN
	// Sin（NaN）= NaN
	fmt.Println(math.Sin(120))//0.5806111842123143
	fmt.Println(math.Sin(45))//0.8509035245341183(17位小数)

	fmt.Println("----------求反双曲正弦和余弦值-----------")

	// Acosh返回x的反双曲余弦值。
	//特殊情况是：
	// Acosh（+ Inf）= + Inf
	//如果x <1，则Acosh（x）= NaN
	// Acosh（NaN）= NaN
	fmt.Println(math.Acosh(30))//4.0940666686320855

	// ASINH返回x的反双曲正弦。
	//特殊情况是：
	// ASINH（±0）=±0
	// Asinh（±Inf）=±Inf
	// ASINH（NAN）= NaN的
	fmt.Println(math.Asinh(30))//4.09462222433053


	fmt.Println("---------返回x的双曲正弦值。----------")
	// Sinh返回x的双曲正弦值。
	//特殊情况是：
	// Sinh（±0）=±0
	// Sinh（±Inf）=±Inf
	// Sinh（NaN）= NaN
	fmt.Println(math.Sinh(120))
	fmt.Println(math.Sinh(60))
	fmt.Println(math.Cosh(120))
	fmt.Println(math.Cosh(60))
	//输出：
	//	6.520904391968161e+51
	//	5.710036949078421e+25
	//	6.520904391968161e+51
	//	5.710036949078421e+25



	fmt.Println("----------求正切值,反正切值与反双曲正切值-----------")
	// Atan返回弧度的x的反正切值。
	//特殊情况是：
	// Atan（±0）=±0
	// Atan（±Inf）=±Pi / 2
	fmt.Println(math.Atan(30))
	fmt.Println(math.Atan(0.5))

	// Tan返回弧度参数x的正切值。
	//特殊情况是：
	// Tan（±0）=±0
	// Tan（±Inf）= NaN
	// Tan（NaN）= NaN
	fmt.Println(math.Tan(30))
	fmt.Println(math.Tan(45))
	fmt.Println(math.Tan(0.5))


	// Atanh返回x的反双曲正切值。
	//特殊情况是：
	// Atanh（1）= + Inf
	// Atanh（±0）=±0
	// Atanh（-1）= -Inf
	//如果x <-1或x> 1则Atanh（x）= NaN
	// Atanh（NaN）= NaN
	fmt.Println(math.Atanh(0.5))//0.5493061443340548
	fmt.Println(math.Atanh(30))//NaN


	//输出：
	//	1.5374753309166493
	//	0.4636476090008061
	//	-6.405331196646276
	//	1.6197751905438615
	//	0.5463024898437905

	fmt.Println("--------返回x的双曲正切值。----------")
	// Tanh返回x的双曲正切值。
	//特殊情况是：
	// Tanh（±0）=±0
	// Tanh（±Inf）=±1
	// Tanh（NaN）= NaN
	fmt.Println(math.Tanh(120))
	fmt.Println(math.Tanh(-120))
	fmt.Println(math.Tanh(60))
	fmt.Println(math.Tanh(-60))

	fmt.Println(math.Tanh(30))
	fmt.Println(math.Tanh(-30))

	fmt.Println(math.Tanh(30.34))
	fmt.Println(math.Tanh(-30.34))
	//输出：
	//	1
	//	-1
	//	1
	//	-1
	//	1
	//	-1
	//	1
	//	-1
	//不了解双曲线的话去了解，无论如何都会是1或者-1，没有第三个值了

	fmt.Println("----------求y/x的反正切-----------")
	// Atan2返回y / x的反正切，使用两者的符号确定返回值的象限。
	//特殊情况是（按顺序）：
	// Atan2（y，NaN）= NaN
	// Atan2（NaN，x）= NaN
	// Atan2（+0，x> = 0）= +0
	// Atan2（-0，x> = 0）= -0
	// Atan2（+0，x <=-0）= + Pi
	// Atan2（-0，x <=-0）= -Pi
	// Atan2（y> 0，0）= + Pi / 2
	// Atan2（y <0，0）= -Pi / 2
	// Atan2（+ Inf，+ Inf）= + Pi / 4
	// Atan2（-Inf，+ Inf）= -Pi / 4
	// Atan2（+ Inf，-Inf）= 3Pi / 4
	// Atan2（-Inf，-Inf）= -3Pi / 4
	// Atan2（y，+ Inf）= 0
	// Atan2（y> 0，-Inf）= + Pi
	// Atan2（y <0，-Inf）= -Pi
	// Atan2（+ Inf，x）= + Pi / 2
	// Atan2（-Inf，x）= -Pi / 2
	fmt.Println(math.Atan2(2.0,4.0))
	fmt.Println(math.Atan2(2,4))
	fmt.Println(math.Atan2(1,3))
	fmt.Println(math.Atan2(10,2))
	//输出：
	//	0.4636476090008061
	//	0.4636476090008061
	//	0.3217505543966422
	//	1.3734007669450157



	fmt.Println("----------IsNaN（）-----------")

	// IsNaN报告f是否为IEEE 754``非数字''值。
	// IEEE 754说，只有NaN满足f！= f。
	//为了避免浮点硬件，可以使用：
	// x：= Float64bits（f）;
	//返回uint32（x >> shift）＆mask == mask && x！= uvinf && x！= uvneginf
	//具体请看http://c.biancheng.net/view/314.html
	//底层就一句话return f != f
	fmt.Println(math.IsNaN(2.0))
	fmt.Println(math.IsNaN(2))
	fmt.Println(math.IsNaN(1.52))
	fmt.Println(math.IsNaN(0.01))
	fmt.Println(math.IsNaN(00.01))
	fmt.Println(math.IsNaN(6.6666*100))
	fmt.Println(math.IsNaN(34.6))
	fmt.Println(math.IsNaN(0/1))
	fmt.Println(math.IsNaN(1.4533071302642137507696589e+02))
	fmt.Println(math.IsNaN(8.2537402562185562902577219e-01))
	fmt.Println(math.IsNaN(2.0))
	fmt.Println(math.IsNaN(-1.5))
	fmt.Println(math.IsNaN(0x7FF8000000000001))
	//输出：
	//	false
	//	false
	//	false
	//	false
	//	false
	//	false
	//	false
	//	false
	//	false
	//	false
	//	false

	fmt.Println("----------求最大值-----------")


	fmt.Println(math.Max(1.0,2.0))
	fmt.Println(math.Max(10.0,2.0))

	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	fmt.Println(math.SmallestNonzeroFloat32)
	fmt.Println(math.SmallestNonzeroFloat64)

	fmt.Println(math.MaxUint8)
	fmt.Println(math.MaxUint16)
	fmt.Println(math.MaxUint32)
	//fmt.Println(math.MaxUint64)//这个会报错，超过长度了,constant 18446744073709551615 overflows int

	fmt.Println(math.MaxInt8)
	fmt.Println(math.MaxInt16)
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MaxInt64)
	//输出：
	//	2
	//	10
	//	1.401298464324817e-45
	//	5e-324
	//	3.4028234663852886e+38
	//	1.7976931348623157e+308
	//	255
	//	65535
	//	4294967295
	//	127
	//	32767
	//	2147483647
	//	9223372036854775807

	fmt.Println("----------求最小值-----------")

	fmt.Println(math.Min(1.0,2.0))
	fmt.Println(math.Min(10.0,2.0))
	fmt.Println(math.MinInt8)
	fmt.Println(math.MinInt16)
	fmt.Println(math.MinInt32)
	fmt.Println(math.MinInt64)
	//输出：
	//	1
	//	2
	//	-128
	//	-32768
	//	-2147483648
	//	-9223372036854775808

	fmt.Println("----------NaN()-----------")

	// NaN返回IEEE 754``非数字''值。底层用到Float64frombits（）方法
	// Float64frombits返回与IEEE 754二进制表示形式b相对应的浮点数，其符号位b和结果位于相同的位位置。
	// Float64frombits（Float64bits（x））== x。
	fmt.Println(math.NaN())//NaN


	fmt.Println("----------四舍五入到最近值且向上取整-----------")
	// Round返回最接近的整数，从零开始四舍五入。
	//特殊情况是：
	//舍入（±0）=±0
	//舍入（±Inf）=±Inf
	// Round（NaN）= NaN

	num:=math.Round(3.14)//在这里已经转为3了
	num1:=math.Round(3.56)
	num2:=math.Round(3.50)
	fmt.Println(3.14,num,num1,num2)


	fmt.Println("---------返回最接近的整数，四舍五入为偶数-----------")

	// RoundToEven返回最接近的整数，四舍五入为偶数。
	//特殊情况是：
	// RoundToEven（±0）=±0
	// RoundToEven（±Inf）=±Inf
	// RoundToEven（NaN）= NaN
	fmt.Println(math.RoundToEven(3.14))
	fmt.Println(math.RoundToEven(3))
	fmt.Println(math.RoundToEven(3.5))
	fmt.Println(math.RoundToEven(2.14))
	fmt.Println(math.RoundToEven(2.5))
	fmt.Println(math.RoundToEven(1.14))
	fmt.Println(math.RoundToEven(1.5))
	//输出：
	//	3
	//	3
	//	4
	//	2
	//	2
	//	1
	//	2

	fmt.Println("----------返回大于或等于x的最小整数值-----------")
	// Ceil返回大于或等于x的最小整数值。
	//特殊情况是：
	// Ceil（±0）=±0
	// Ceil（±Inf）=±Inf
	// Ceil（NaN）= NaN
	fmt.Println(math.Ceil(3.14))
	fmt.Println(math.Ceil(2.96))
	fmt.Println(math.Ceil(2.45))
	fmt.Println(math.Ceil(2))
	//输出：
	//	4
	//	3
	//	3
	//	2


	fmt.Println("---------返回小于或等于x的最大整数值-----------")

	// Floor返回小于或等于x的最大整数值。
	//特殊情况是：
	//地板（±0）=±0
	//底数（±Inf）=±Inf
	// Floor（NaN）= NaN
	fmt.Println(math.Floor(3.14))
	fmt.Println(math.Floor(2.96))
	fmt.Println(math.Floor(2.0))

	//输出：
	//	3
	//	2
	//	2



	fmt.Println("----------返回立方根-----------")
	// Cbrt返回x的立方根(Cubic root)。
	//特殊情况是：
	// Cbrt（±0）=±0
	// Cbrt（±Inf）=±Inf
	// Cbrt（NaN）= NaN
	fmt.Println(math.Cbrt(8))
	fmt.Println(math.Cbrt(27))
	fmt.Println(math.Cbrt(-27))
	fmt.Println(math.Cbrt(0))
	//输出：
	//	2
	//	3
	//	-3
	//	0


	fmt.Println("---------Copysign（）复制标志位-----------")
	//返回拥有x的量值（绝对值）和y的标志位（正负号）的浮点数。
	fmt.Println(math.Copysign(0,9))
	fmt.Println(math.Copysign(0,-9))
	fmt.Println(math.Copysign(0,+9))
	fmt.Println(math.Copysign(8,9))
	fmt.Println(math.Copysign(8,0))
	fmt.Println(math.Copysign(9,8))
	fmt.Println(math.Copysign(9,-8))
	fmt.Println(math.Copysign(4,3))
	fmt.Println(math.Copysign(-4,-3))
	fmt.Println(math.Copysign(-4,-30))
	fmt.Println(math.Copysign(-4,-300))

	//输出：
	//	0
	//	-0
	//	0
	//	8
	//	8
	//	9
	//	-9
	//	4
	//	-4
	//	-4
	//	-4


	fmt.Println("---------返回IEEE 754二进制表示形式以及逆向转换-----------")
	// Float32bits返回f的IEEE 754二进制表示形式，其中f的符号位和结果位于相同的位位置。
	// Float32bits（Float32frombits（x））== x。
	//函数返回浮点数f的IEEE 754格式二进制表示对应的4字节无符号整数。

	fmt.Println(math.Float32bits(0))
	fmt.Println(math.Float32bits(1))
	fmt.Println(math.Float32bits(2))
	//下面是逆向转换
	// Float32frombits返回对应于IEEE 754二进制表示形式b的浮点数，其符号位b和结果位于相同的位位置。
	// Float32frombits（Float32bits（x））== x。
	//函数返回无符号整数b对应的IEEE 754格式二进制表示的8字节浮点数。
	fmt.Println(math.Float32frombits(1065353216))
	fmt.Println(math.Float32frombits(1073741824))
	fmt.Println(math.Float32frombits(0))

	//输出：
	//	0
	//	1065353216
	//	1073741824
	//	1
	//	2
	//	0

	//函数返回浮点数f的IEEE 754格式二进制表示对应的8字节无符号整数。
	fmt.Println(math.Float64bits(0))
	fmt.Println(math.Float64bits(1))
	fmt.Println(math.Float64bits(2))
	//下面是逆向转换
	fmt.Println(math.Float64frombits(1065353216))
	fmt.Println(math.Float64frombits(1073741824))
	fmt.Println(math.Float64frombits(0))
	//输出：
	//	0
	//	4607182418800017408
	//	4611686018427387904
	//	5.263544247e-315
	//	5.304989477e-315
	//	0


	fmt.Println("---------Dim返回x-y或0的最大值-----------")

	// Dim返回x-y或0的最大值。(要么返回0，要么返回x-y的差)
	//特殊情况是：
	// Dim（+ Inf，+ Inf）= NaN
	// Dim（-Inf，-Inf）= NaN
	// Dim（x，NaN）= Dim（NaN，x）= NaN
	fmt.Println(math.Dim(2,5))
	fmt.Println(math.Dim(5,2))
	//输出：
	//	0
	//	3

	fmt.Println("---------Erf（x）求x在误差函数对应的值-----------")
	fmt.Println(math.E)
	//关于误差函数请看：https://baike.baidu.com/pic/%E8%AF%AF%E5%B7%AE%E5%87%BD%E6%95%B0/5890875/0/34fae6cd7b899e5101b3d83d42a7d933c8950ddd?fr=lemma&ct=single#aid=0&pic=34fae6cd7b899e5101b3d83d42a7d933c8950ddd
	//或者：https://baike.baidu.com/item/%E8%AF%AF%E5%B7%AE%E5%87%BD%E6%95%B0/5890875?fr=aladdin

	//Erf返回错误函数或x。
	//特殊情况是：
	//Erf（+ Inf）= 1
	//Erf（-Inf）= -1
	//Erf（NaN）= NaN
	fmt.Println(math.Erf(3.14))
	// Erfc返回x的互补误差函数。
	//
	//特殊情况是：
	// Erfc（+ Inf）= 0
	// Erfc（-Inf）= 2
	// Erfc（NaN）= NaN
	fmt.Println(math.Erfc(3.14))
	// Erfinv返回x的反误差函数。Erf的逆向运算
	//
	//特殊情况是：
	// Erfinv（1）= + Inf
	// Erfinv（-1）= -Inf
	//如果x <-1或x> 1则Erfinv（x）= NaN
	// Erfinv（NaN）= NaN
	fmt.Println(math.Erfinv(0.9999910304344467))
	// Erfcinv返回Erfc（x）的逆向。
	//
	//特殊情况是：
	// Erfcinv（0）= + Inf
	// Erfcinv（2）= -Inf
	//如果x <0或x> 2则Erfcinv（x）= NaN
	// Erfcinv（NaN）= NaN
	fmt.Println(math.Erfcinv(8.969565553264983e-06))

	//输出：
	//	2.718281828459045
	//	0.9999910304344467
	//	8.969565553264983e-06
	//	3.1399999999994708
	//	3.1399999999994708


	fmt.Println("---------关于E的相关运算函数----------")
	// Exp返回e ** x，即x的底e指数。
	//特殊情况是：
	// Exp（+ Inf）= + Inf
	// Exp（NaN）= NaN
	//非常大的值溢出到0或+ Inf。
	//非常小的值下溢到1。
	fmt.Println(math.Exp(8))

	// Exp2返回2 ** x，x的以2为底的指数。
	//特殊情况与Exp相同。
	fmt.Println(math.Exp2(3))

	// Expm1返回(e ** x)-1，x的基数e减1。
	//当x接近零时，它比Exp（x）-1更准确。
	//特殊情况是：
	// Expm1（+ Inf）= + Inf
	// Expm1（-Inf）= -1
	// Expm1（NaN）= NaN
	//非常大的值溢出到-1或+ Inf。
	fmt.Println(math.Expm1(8))


	// Frexp将f分解为归一化的分数和2的整数次幂。它返回frac和exp满足f == frac×(2 ** exp)，且frac的绝对值在[1/2，1）区间内。
	//特殊情况是：
	// Frexp（±0）=±0，0
	// Frexp（±Inf）=±Inf，0
	// Frexp（NaN）= NaN，0
	fmt.Println(math.Frexp(8))

	//Frexp（）方法的逆运算
	fmt.Println(math.Ldexp(0.5,4))

	//上面所有输出如下：
	//	2980.9579870417283
	//	8
	//	2979.9579870417283
	//	0.5 4
	//	8

	fmt.Println("---------返回x的整数值-----------")
	// Trunc返回x的整数值。
	//特殊情况是：
	// Trunc（±0）=±0
	// Trunc（±Inf）=±Inf
	// Trunc（NaN）= NaN
	//底层主要是Modf（）方法，下面是说明：
	// Modf返回总和为f的整数和小数浮点数。 这两个值与f具有相同的符号。
	//特殊情况是：
	// Modf（±Inf）=±Inf，NaN
	// Modf（NaN）= NaN，NaN
	fmt.Println(math.Trunc(3.14))
	fmt.Println(math.Trunc(3))
	fmt.Println(math.Trunc(2.96))
	fmt.Println(math.Trunc(0.5))
	//输出：
	//	3
	//	3
	//	2
	//	0


	fmt.Println("---------返回总和为f的整数和小数浮点数-----------")
	// Modf返回总和为f的整数和小数浮点数。 这两个值与f具有相同的符号。
	//说白了就是返回浮点数的整数部分和小数部分
	//特殊情况是：
	// Modf（±Inf）=±Inf，NaN
	// Modf（NaN）= NaN，NaN
	fmt.Println(math.Modf(3.14))
	fmt.Println(math.Modf(3.15))
	fmt.Println(math.Modf(3.16))
	fmt.Println(math.Modf(3.1))
	//输出：
	//	3 0.14000000000000012（这里是2个数，下同，同时由于计算机的原因无法完全精准的表示某些浮点数）
	//	3 0.1499999999999999
	//	3 0.16000000000000014
	//	3 0.10000000000000009（保留18位小数）

	fmt.Println("---------返回x / y的浮点余数-----------")
	// Mod返回x / y的浮点余数。
	//结果的大小小于y，其符号与x一致。
	//特殊情况是：
	// Mod（±Inf，y）= NaN
	// Mod（NaN，y）= NaN
	// Mod（x，0）= NaN
	// Mod（x，±Inf）= x
	// Mod（x，NaN）= NaN
	fmt.Println(math.Mod(6,3))
	fmt.Println(math.Mod(-6,3))

	fmt.Println(math.Mod(5,3))
	fmt.Println(math.Mod(-5,3))
	//输出：
	//	0
	//	-0
	//	2
	//	-2


	fmt.Println("---------返回x的Gamma函数对应的值-----------")
	//对于伽马函数请看：https://baike.baidu.com/item/%E4%BC%BD%E7%8E%9B%E5%87%BD%E6%95%B0/3540177?fromtitle=gamma%E5%87%BD%E6%95%B0&fromid=10819772&fr=aladdin

	// Gamma返回x的Gamma函数。
	//特殊情况是：
	//伽玛（+ Inf）= + Inf
	//伽玛（+0）= + Inf
	//伽玛（-0）= -Inf
	// Gamma（x）= NaN（整数x <0）
	//伽玛（-Inf）= NaN
	//伽玛（NaN）= NaN
	fmt.Println(math.Gamma(6))//5!
	fmt.Println(math.Gamma(6.1))//5.1!
	fmt.Println(math.Gamma(5))//4!
	fmt.Println(math.Gamma(5.1))//4.1!=4.1*3.1*2.1*1.1-xxx(这个xxx是一定的值，但是这个值必须使用Gamma函数才能算出来的)

	//输出：
	//	120
	//	142.45194406567867
	//	24
	//	27.93175373836837


	fmt.Println("---------返回Gamma（x）的自然对数和符号（-1或+1）-----------")
	// Lgamma返回Gamma（x）的自然对数和符号（-1或+1）。
	//特殊情况是：
	// Lgamma（+ Inf）= + Inf
	// Lgamma（0）= + Inf
	// Lgamma（-integer）= + Inf
	// Lgamma（-Inf）= -Inf
	// Lgamma（NaN）= NaN
	fmt.Println(math.Lgamma(6.1))
	fmt.Println(math.Lgamma(6))
	fmt.Println(math.Lgamma(5.1))
	fmt.Println(math.Lgamma(5))

	fmt.Println(math.Lgamma(-5.1))
	fmt.Println(math.Lgamma(-5))

	fmt.Println(math.Lgamma(-2.5))
	fmt.Println(math.Lgamma(-2))
	//输出：
	//	4.959004708205504 1
	//	4.787491742782046 1
	//	3.3297641684752244 1
	//	3.1780538303479458 1
	//	-2.639915816736552 1
	//	+Inf 1
	//	-0.05624371649767412 -1
	//	+Inf 1

	fmt.Println("---------返回正无穷大和负无穷大-----------")
	//如果符号> = 0，则Inf返回正无穷大；如果符号<= 0，则Inf返回负无穷大。
	fmt.Println(math.Inf(6))
	fmt.Println(math.Inf(-6))
	fmt.Println(math.Inf(0))
	fmt.Println(math.Inf(-0))

	//输出：
	//	+Inf
	//	-Inf
	//	+Inf
	//	+Inf

	fmt.Println("---------求2个数的平方和的开平方-----------")

	//Hypot-sqrt（p * p + q * q），但仅在结果允许时溢出。

	// Hypot返回Sqrt（p * p + q * q），请注意避免不必要的上溢和下溢。
	//特殊情况是：
	// Hypot（±Inf，q）= + Inf
	// Hypot（p，±Inf）= + Inf
	// Hypot（NaN，q）= NaN
	// Hypot（p，NaN）= NaN
	fmt.Println(math.Hypot(3,4))
	fmt.Println(math.Hypot(6,8))
	fmt.Println(math.Hypot(8,10))
	//输出：
	//	5
	//	10
	//	12.806248474865697




	fmt.Println("---------返回第一种零阶贝塞尔函数-----------")
	//更多信息请看：https://baike.baidu.com/item/%E8%B4%9D%E5%A1%9E%E5%B0%94%E5%87%BD%E6%95%B0/3431101?fr=aladdin
	//或者：https://baike.baidu.com/pic/%E8%B4%9D%E5%A1%9E%E5%B0%94%E5%87%BD%E6%95%B0/3431101/0/a1ec08fa513d26978d6e6a7152fbb2fb4316d859?fr=lemma&ct=single#aid=0&pic=a1ec08fa513d26978d6e6a7152fbb2fb4316d859
	// J0返回第一种零阶贝塞尔函数。
	//特殊情况是：
	// J0（±Inf）= 0
	// J0（0）= 1
	// J0（NaN）= NaN
	fmt.Println(math.J0(0.0000001))
	fmt.Println(math.J0(0.1))
	fmt.Println(math.J0(1))
	fmt.Println(math.J0(16))
	fmt.Println(math.J0(-16))
	fmt.Println(math.J0(25))
	fmt.Println(math.J0(250))


	// J1返回第一类的一阶贝塞尔函数。
	//特殊情况是：
	// J1（±Inf）= 0
	// J1（NaN）= NaN
	fmt.Println()
	fmt.Println(math.J1(0.0000001))
	fmt.Println(math.J1(0.1))
	fmt.Println(math.J1(16))

	// Jn返回第一种类型的n阶贝塞尔函数。
	//特殊情况是：
	// Jn（n，±Inf）= 0
	// Jn（n，NaN）= NaN
	fmt.Println()
	fmt.Println(math.Jn(2,0.0000001))
	fmt.Println(math.Jn(2,0.1))
	fmt.Println(math.Jn(2,16))

	//输出：
	//	0.9999999999999974
	//	0.99750156206604
	//	0.7651976865579666
	//	-0.1748990739836292
	//	-0.1748990739836292
	//	0.09626678327595813
	//	-0.026053373425204234
	//
	//	4.999999999999994e-08
	//	0.049937526036242
	//	0.09039717566130417
	//
	//	1.249999999999999e-15
	//	0.001248958658799919
	//	0.18619872094129222


	fmt.Println("---------返回第二种零阶贝塞尔函数-----------")
	// Y0返回第二种零阶贝塞尔函数。
	//特殊情况是：
	// Y0（+ Inf）= 0
	// Y0（0）= -Inf
	// Y0（x <0）= NaN
	// Y0（NaN）= NaN
	fmt.Println(math.Y0(0.00000001))
	fmt.Println(math.Y0(0.1))
	fmt.Println(math.Y0(1))
	fmt.Println(math.Y0(4))
	fmt.Println(math.Y0(16))
	// Y1返回第二种的一阶贝塞尔函数。
	//特殊情况是：
	// Y1（+ Inf）= 0
	// Y1（0）= -Inf
	// Y1（x <0）= NaN
	// Y1（NaN）= NaN
	fmt.Println()
	fmt.Println(math.Y1(0.00000001))
	fmt.Println(math.Y1(0.1))
	fmt.Println(math.Y1(1))
	fmt.Println(math.Y1(4))
	fmt.Println(math.Y1(16))

	// Yn返回第二种序数n的贝塞尔函数。
	//特殊情况是：
	// Yn（n，+ Inf）= 0
	// Yn（n≥0，0）= -Inf
	// Yn（n <0，0）= + Inf如果n为奇数，-Inf如果n为偶数
	// Yn（n，x <0）= NaN
	// Yn（n，NaN）= NaN
	fmt.Println()
	fmt.Println(math.Yn(2,0.00000001))
	fmt.Println(math.Yn(2,0.1))
	fmt.Println(math.Yn(2,1))
	fmt.Println(math.Yn(2,4))
	fmt.Println(math.Yn(2,16))

	//输出：
	//	-11.800773877179532
	//	-1.5342386513503667
	//	0.08825696421567697
	//	-0.016940739325064996
	//	0.0958109970807124
	//
	//	-6.3661977236758195e+07
	//	-6.458951094702027
	//	-0.7812128213002887
	//	0.3979257105571
	//	0.17797516893941684
	//
	//	-1.2732395447351626e+16
	//	-127.64478324269018
	//	-1.6506826068162543
	//	0.215903594603615
	//	-0.0735641009632853


	fmt.Println("---------返回x的以各种底数的指数-----------")
	//自然对数以常数e为底数的对数。记作lnN(N>0)。在物理学，生物学等自然科学中有重要的意义。一般表示方法为lnx。数学中也常见以logx表示自然对数。



	// Log返回x的自然对数。
	//特殊情况是：
	// Log（+ Inf）= + Inf
	// Log（0）= -Inf
	// Log（x <0）= NaN
	// Log（NaN）= NaN
	fmt.Println(math.Log(100))

	// Log1p返回自然对数1及其参数x。
	//当x接近零时，它比Log（1 + x）更准确。
	//特殊情况是：
	// Log1p（+ Inf）= + Inf
	// Log1p（±0）=±0
	// Log1p（-1）= -Inf
	// Log1p（x <-1）= NaN
	// Log1p（NaN）= NaN
	fmt.Println(math.Log1p(100))


	// Log2返回某个数以2为底的指数。
	//特殊情况与Log相同。
	fmt.Println(math.Log2(64))
	fmt.Println(math.Log2(64.4))

	fmt.Println(math.Log10(100))
	fmt.Println(math.Log10(100.3))
	//输出：
	//	4.605170185988092
	//	4.61512051684126
	//	6
	//	6.008988783227255
	//	2
	//	2.0013009330204183


	fmt.Println("---------返回某个数以2为底的指数的整形和浮点数-----------")
	//其实跟上面的api差不多
	//返回x的二进制指数值，可以理解为Trunc(Log2(x))；特例如下：
	//Logb(±Inf) = +Inf
	//Logb(0) = -Inf
	//Logb(NaN) = NaN
	fmt.Println(math.Logb(16))
	fmt.Println(math.Logb(16.14))
	fmt.Println(math.Logb(32))
	fmt.Println(math.Logb(32.14))

	//类似Logb，但返回值是整型；
	// Ilogb以整数形式返回x的二进制指数。
	//特殊情况是：
	// Ilogb（±Inf）= MaxInt32
	// Ilogb（0）= MinInt32
	// Ilogb（NaN）= MaxInt32
	fmt.Println(math.Ilogb(32))
	fmt.Println(math.Ilogb(32.14))
	//输出：
	//	4
	//	4
	//	5
	//	5
	//	5
	//	5

	fmt.Println("---------返回某个数是否是无穷大-----------")
	// IsInf根据符号报告f是否为无穷大。
	//如果符号> 0，则IsInf报告f是否为正无穷大。
	//如果符号<0，则IsInf报告f是否为负无穷大。
	//如果符号== 0，则IsInf报告f是否为无穷大。
	fmt.Println(math.IsInf(32.14,1))
	fmt.Println(math.IsInf(32.14,10))
	fmt.Println(math.IsInf(32.14,-1))
	fmt.Println(math.IsInf(math.MaxFloat64,1))
	//输出：
	//	false
	//	false
	//	false
	//	false

	fmt.Println("---------返回x之后到y的下一个可表示的float64值或者float32的值-----------")
	// Nextafter返回x之后到y的下一个可表示的float64值。
	//特殊情况是：
	// Nextafter（x，x）= x
	// Nextafter（NaN，y）= NaN
	// Nextafter（x，NaN）= NaN
	fmt.Println(math.Nextafter(2,10))
	fmt.Println(math.Nextafter(3,10))
	fmt.Println(math.Nextafter(3,11))
	fmt.Println(math.Nextafter(3,1))

	fmt.Println(math.Nextafter(3,-3))
	fmt.Println(math.Nextafter(3,3))
	fmt.Println(math.Nextafter(0,3))
	//输出：
	//	2.0000000000000004
	//	3.0000000000000004
	//	3.0000000000000004
	//	2.9999999999999996
	//	3
	//	2.9999999999999996
	//	5e-324

	fmt.Println(math.Nextafter32(2,10))
	fmt.Println(math.Nextafter32(3,10))
	fmt.Println(math.Nextafter32(3,11))
	fmt.Println(math.Nextafter32(3,1))

	fmt.Println(math.Nextafter32(3,-3))
	fmt.Println(math.Nextafter32(3,3))
	fmt.Println(math.Nextafter32(0,3))

	//输出：
	//	2.0000002
	//	3.0000002
	//	3.0000002
	//	2.9999998
	//	2.9999998
	//	3
	//	1e-45



	fmt.Println("---------求一个数x的y次方-----------")
	// Pow返回x ** y，y的基数x指数。
	//特殊情况是（按顺序）：
	//	Pow(x, ±0) = 1 for any x
	//	Pow(1, y) = 1 for any y
	//	Pow(x, 1) = x for any x
	//	Pow(NaN, y) = NaN
	//	Pow(x, NaN) = NaN
	//	Pow(±0, y) = ±Inf for y an odd integer < 0
	//	Pow(±0, -Inf) = +Inf
	//	Pow(±0, +Inf) = +0
	//	Pow(±0, y) = +Inf for finite y < 0 and not an odd integer
	//	Pow(±0, y) = ±0 for y an odd integer > 0
	//	Pow(±0, y) = +0 for finite y > 0 and not an odd integer
	//	Pow(-1, ±Inf) = 1
	//	Pow(x, +Inf) = +Inf for |x| > 1
	//	Pow(x, -Inf) = +0 for |x| > 1
	//	Pow(x, +Inf) = +0 for |x| < 1
	//	Pow(x, -Inf) = +Inf for |x| < 1
	//	Pow(+Inf, y) = +Inf for y > 0
	//	Pow(+Inf, y) = +0 for y < 0
	//	Pow(-Inf, y) = Pow(-0, -y)
	// Pow（x，y）= NaN表示有限x <0和有限非整数y
	fmt.Println(math.Pow(2,16))
	fmt.Println(math.Pow(2,3))
	fmt.Println(math.Pow(2,4))
	fmt.Println(math.Pow(2,-4))
	fmt.Println(math.Pow(2,-2))

	fmt.Println(math.Pow10(2))
	fmt.Println(math.Pow10(-2))

	//输出：
	//	65536
	//	8
	//	16
	//	0.0625
	//	0.25
	//	100
	//	0.01


	fmt.Println("---------求x/y之后的余数-----------")
	// Remainder返回x / y的IEEE 754浮点余数。
	//特殊情况是：
	//余数（±Inf，y）= NaN
	//余数（NaN，y）= NaN
	//余数（x，0）= NaN
	//余数（x，±Inf）= x
	//余数（x，NaN）= NaN
	fmt.Println(math.Remainder(8,2))
	fmt.Println(math.Remainder(8.1,2))
	fmt.Println(math.Remainder(8.2,2))
	//输出：
	//0
	//0.09999999999999964
	//0.1999999999999993

	fmt.Println("---------同时求一个角度的Sin（x）和Cos（x）----------")
	//系数_sin []和_cos []可在pkg / math / sin.go中找到。
	// Sincos返回Sin（x），Cos（x）。
	//特殊情况是：
	// Sincos（±0）=±0，1
	// Sincos（±Inf）= NaN，NaN
	// Sincos（NaN）= NaN，NaN
	fmt.Println(math.Sincos(120))
	fmt.Println(math.Sin(120))
	fmt.Println(math.Cos(120))
	//输出：
	//	0.5806111842123143 0.8141809705265618
	//	0.5806111842123143
	//	0.8141809705265618


	fmt.Println("---------报告x是否是真正的负数----------")
	// 报告x是否是真正的负数（主要是为了区别负零和负数的区别，负零不是负数）
	fmt.Println(math.Signbit(-120))
	fmt.Println(math.Signbit(-1))
	fmt.Println(math.Signbit(-0))
	fmt.Println(math.Signbit(1))
	fmt.Println(math.Signbit(120))//给正数的话也会返回false
	//输出：
	//	true
	//	true
	//	false
	//	false
	//	false

	fmt.Println("---------求一个数的开平方的值（负数开平方会得到NaN）----------")
	// 报告x是否是真正的负数（主要是为了区别负零和负数的区别，负零不是负数）
	fmt.Println(math.Sqrt(2))
	fmt.Println(math.Sqrt(4))
	fmt.Println(math.Sqrt(-4))

	fmt.Println(math.IsNaN(math.Sqrt(-4)))
	fmt.Println(math.Sqrt(100))
	//输出：
	//	1.4142135623730951
	//	2
	//	NaN
	//	true
	//	10


	//fmt.Println("---------课外----------")
	//feeStr := ""
	//fee,_ := strconv.ParseFloat(feeStr,32)
	//fee2 := fee + 123
	//fmt.Println(fee)
	//fmt.Println(fee2)
	////输出：
	////0
	////123



}




