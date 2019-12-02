package main

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("----------big.Int()-----------")

	// NewInt分配并返回一个新的Int设置为x。返回的是指针
	BInt := big.NewInt(16)
	FormatP(BInt)
	BInt111 := big.NewInt(9)
	// Sqrt将z设置为⌊√x⌋，即z²≤x的最大整数，并返回z。
	//如果x为负数则表示恐慌。
	//说白了就是设置z的值为x的开平方值，返回的是一个新的可以被接收的值，但是z一定会改变的
	BInt_Sqrt := BInt.Sqrt(BInt111)

	fmt.Println(BInt_Sqrt)
	FormatP(BInt)
	FormatP(BInt111)
	//输出：
	//*big.Int--16
	//3
	//*big.Int--3
	//*big.Int--9

	// Mod将y！= 0的z设置为模数x％y并返回z。
	//如果y == 0，则会发生除以零的运行时恐慌。
	// Mod实现欧几里得模数（不同于Go）； 有关更多详细信息，请参见DivMod。
	x := big.NewInt(16)
	y := big.NewInt(9)
	BInt_Mod := BInt.Mod(x, y)

	fmt.Println()
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(BInt_Mod)
	fmt.Println(BInt)
	//输出：
	//16
	//9
	//7
	//7

	y = big.NewInt(9)
	//将集合z加到x + y的总和并返回z。
	BInt_Mod = BInt.Add(x, y)
	BInt_Mod = BInt.Add(x, y)
	fmt.Println()
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(BInt_Mod)
	fmt.Println(BInt)
	//输出：
	//16
	//9
	//25
	//25

	x = big.NewInt(16)
	// Sub将z设置为差x-y并返回z。
	BInt_Mod = BInt.Sub(x, y)

	BInt_Mod = BInt.Sub(x, y)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(BInt_Mod)
	fmt.Println(BInt)
	//输出：
	//16
	//9
	//7
	//7

	x = big.NewInt(16)
	y = big.NewInt(9)

	fmt.Println()
	fmt.Println(x)
	BInt_Mod = BInt.Div(x, y)
	fmt.Println(BInt_Mod)
	fmt.Println(BInt)
	//输出：
	//16
	//9
	//1
	//1

	x = big.NewInt(32)
	y = big.NewInt(9)
	m := big.NewInt(2)
	// DivMod将z设置为商x div y，将m设置为模数x mod y，并返回y！= 0的对（z，m）。
	// DivMod实现欧几里得除法和模数（与Go不同）：
	// q = x div y这样
	// m = x-y * q，0 <= m <| y |
	//（请参阅Raymond T.Boute，``函数div和mod的欧几里得定义''.ACM Transactions on Programming Languages and Systems（TOPLAS），14（2）：127-144，New York，NY，USA， 4/1992。ACM出版社。）
	//有关T划分和模数（如Go），请参见QuoRem。
	BInt_DivMod_i, BInt_DivMod_m := BInt.DivMod(x, y, m) //不关m原来是多少，这里都将重新为他赋值为x%y

	fmt.Println()
	fmt.Println(x)
	fmt.Println(y)
	BInt_DivMod_i, BInt_DivMod_m := BInt.DivMod(x, y, m) //不关m原来是多少，这里都将重新为他赋值为x%y
	fmt.Println(BInt_DivMod_m)
	fmt.Println(BInt)
	fmt.Println(m)
	//输出：
	//32
	//9
	//3
	//5
	//3
	//5

	x = big.NewInt(0b11001100)
	y = big.NewInt(0b11001000)
	//                0b11001000
	//And设置z = x＆y并返回z。
	BInt_Mod = BInt.And(x, y)
	fmt.Println()
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(BInt_Mod)
	BInt_Mod = BInt.And(x, y)
	fmt.Println(0b11001000)
	//输出：
	//204
	//200
	//200
	//200
	//200

	x = big.NewInt(0b11001100)
	y = big.NewInt(0b11001000)
	//   ^y           0b00110111
	//   x            0b11001100
	//   x&^y         0b00000100
	// AndNot设置z = x＆^ y并返回z。
	fmt.Println()
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(BInt_Mod)
	fmt.Println(BInt)
	fmt.Println(0b00000100)
	BInt_Mod = BInt.AndNot(x, y)
	//204
	//200
	//4
	//4
	//4

	x = big.NewInt(3)
	y = big.NewInt(3)
	m = big.NewInt(2)
	// Exp set z = x ** y mod | m | （即，忽略m的符号），并返回z。
	//如果m == nil或m == 0，则z = x ** y，除非y <= 0则z =1。如果m> 0，y <0，并且x和n不是相对质数，则z不变 然后返回nil。
	//特定大小的输入的模幂不是密码恒定时间操作。
	BInt_ret := BInt.Exp(x, y, m)

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("m:", m)
	fmt.Println("BInt_ret:", BInt_ret) //9除以2余数1
	fmt.Println("BInt:", BInt)
	//输出：
	BInt_ret := BInt.Exp(x, y, m)
	//y: 3
	//m: 2
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("m:", m)
	fmt.Println("BInt_ret:", BInt_ret) //9除以2余数1
	fmt.Println("BInt:", BInt)
	BInt_ret = BInt.Abs(x)

	fmt.Println()
	fmt.Println("x:", x)
	fmt.Println("BInt_ret:", BInt_ret)
	fmt.Println("BInt:", BInt)
	//输出：
	//BInt_ret: 3
	//BInt: 3
	BInt_ret = BInt.Abs(x)
	x = big.NewInt(3)
	BInt = big.NewInt(2)
	fmt.Println("x:", x)
	fmt.Println("BInt_ret:", BInt_ret)
	fmt.Println("BInt:", BInt)
	//如果x> y则为+1
	BInt_ret_int := BInt.Cmp(x)

	fmt.Println()
	fmt.Println("x:", x)
	fmt.Println("BInt:", BInt)
	BInt = big.NewInt(2)
	BInt = big.NewInt(3)
	BInt_ret_int = BInt.Cmp(x)

	fmt.Println()
	BInt_ret_int := BInt.Cmp(x)
	fmt.Println("BInt_ret_int:", BInt_ret_int)
	fmt.Println("BInt:", BInt)
	fmt.Println("x:", x)
	fmt.Println("BInt_ret_int:", BInt_ret_int)
	fmt.Println("BInt:", BInt)

	BInt = big.NewInt(3)
	BInt_ret_int = BInt.Cmp(x)
	fmt.Println("BInt_ret_int:", BInt_ret_int)
	fmt.Println("BInt:", BInt)
	fmt.Println("x:", x)
	fmt.Println("BInt_ret_int:", BInt_ret_int)
	fmt.Println("BInt:", BInt)
	//BInt: 2
	BInt = big.NewInt(4)
	BInt_ret_int = BInt.Cmp(x)
	//BInt_ret_int: 0
	//BInt: 3
	fmt.Println("x:", x)
	fmt.Println("BInt_ret_int:", BInt_ret_int)
	fmt.Println("BInt:", BInt)
	//BInt: 4

	y = big.NewInt(-3)
	BInt = big.NewInt(2)
	// CmpAbs比较x和y的绝对值并返回：
	// 如果| x | <| y |返回-1，
	// 如果| x |== | y |则为0，
	// 如果| x | > | y |返回+1，
	BInt_ret_int = BInt.CmpAbs(y)

	fmt.Println()
	fmt.Println("y:", y)
	fmt.Println("BInt_ret_int:", BInt_ret_int)
	BInt = big.NewInt(-3)
	BInt = big.NewInt(2)

	fmt.Println()
	fmt.Println("y:", y)
	fmt.Println("BInt_ret_int:", BInt_ret_int)
	BInt_ret_int = BInt.CmpAbs(y)

	BInt = big.NewInt(-4)
	fmt.Println("y:", y)
	fmt.Println("BInt_ret_int:", BInt_ret_int)
	fmt.Println("BInt:", BInt)
	fmt.Println("y:", y)
	BInt = big.NewInt(-3)
	BInt_ret_int = BInt.CmpAbs(y)
	//输出：
	//y: -3
	fmt.Println("y:", y)
	fmt.Println("BInt_ret_int:", BInt_ret_int)
	fmt.Println("BInt:", BInt)
	//y: -3
	BInt = big.NewInt(-4)
	BInt_ret_int = BInt.CmpAbs(y)
	//
	//y: -3
	fmt.Println("y:", y)
	fmt.Println("BInt_ret_int:", BInt_ret_int)
	fmt.Println("BInt:", BInt)
	//x = big.NewInt(0b00000010)
	//x = big.NewInt(0b11000010)
	x = big.NewInt(0b1100_0000_0000_0010) //中间可以用下划线进行分割
	BInt = big.NewInt(3)
	// Lsh设置z = x << n并返回z。
	BInt_ret = BInt.Lsh(x, 2) //左移N位就相当于你x*2的N次方

	fmt.Println()
	fmt.Println("x:", x)
	fmt.Println("BInt_ret:", BInt_ret) //9除以2余数1
	fmt.Println("BInt:", BInt)
	//0b00000010输出：
	//x: 2
	//BInt: 8

	x = big.NewInt(0b1100_0000_0000_0010) //中间可以用下划线进行分割
	//x: 194
	//BInt_ret: 776
	BInt_ret = BInt.Lsh(x, 2) //左移N位就相当于你x*2的N次方

	//0b1100_0000_0000_0010输出：
	fmt.Println("x:", x)
	fmt.Println("BInt_ret:", BInt_ret) //9除以2余数1
	fmt.Println("BInt:", BInt)

	//x = big.NewInt(0b00000010)//右移会丢弃一些位
	//x = big.NewInt(0b11000010)
	x = big.NewInt(0b1100_0000_0000_0010) //中间可以用下划线进行分割
	BInt = big.NewInt(3)
	// Rsh设置z = x >> n并返回z。
	BInt_ret = BInt.Rsh(x, 2) //左移N位就相当于你x*2的N次方

	fmt.Println()
	fmt.Println("x:", x)
	fmt.Println("BInt_ret:", BInt_ret) //9除以2余数1
	fmt.Println("BInt:", BInt)
	//0b00000010输出：
	//x: 2
	//BInt_ret: 0
	//BInt: 0

	x = big.NewInt(0b1100_0000_0000_0010) //中间可以用下划线进行分割
	//x: 194
	//BInt_ret: 48
	BInt_ret = BInt.Rsh(x, 2) //左移N位就相当于你x*2的N次方

	//0b1100_0000_0000_0010输出：
	fmt.Println("x:", x)
	fmt.Println("BInt_ret:", BInt_ret) //9除以2余数1
	fmt.Println("BInt:", BInt)

	x = big.NewInt(2)
	y = big.NewInt(3)
	BInt = big.NewInt(3)
	// Mul将z设置为乘积x * y并返回z。
	BInt_ret = BInt.Mul(x, y)

	fmt.Println()
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("BInt_ret:", BInt_ret)
	fmt.Println("BInt:", BInt)
	//输出：
	//x: 2
	//y: 3

	BInt = big.NewInt(3)
	// MulRange将z设置为[a，b]范围内的所有整数的乘积，并返回z。
	//如果a> b（空范围），则结果为1。
	BInt_ret = BInt.Mul(x, y)

	fmt.Println()
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("BInt_ret:", BInt_ret)
	fmt.Println("BInt:", BInt)
	//BInt: 720

	x = big.NewInt(23)
	//x = big.NewInt(22)
	//x = big.NewInt(21)
	//x = big.NewInt(20)
	//x = big.NewInt(17)
	//x = big.NewInt(16)
	//x = big.NewInt(15)
	BInt_ret = BInt.MulRange(2, 6) //2*3*4*5*6
	//x = big.NewInt(13)
	//x = big.NewInt(12)
	fmt.Println("BInt_ret:", BInt_ret)
	fmt.Println("BInt:", BInt)
	//y = big.NewInt(9)//莫名的出现不能停止的情况
	//y = big.NewInt(8)//invalid 2nd argument to Int.Jacobi: need odd integer but got 8(译为：Int.Jacobi的无效第二个参数：需要奇数但得到8)
	y = big.NewInt(7)
	BInt = big.NewInt(19)
	// 如果x不是平方模p，则ModSqrt保持z不变并返回nil。 如果p不是一个奇整数，则此函数会出现紧急情况。
	BInt_ret = BInt.ModSqrt(x, y)

	fmt.Println()
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("BInt_ret:", BInt_ret)
	fmt.Println("BInt:", BInt)
	//23输出：
	//x: 23
	//y: 7
	//BInt_ret: 4
	//BInt: 4

	//22输出：
	//x: 22
	//y: 7
	//BInt_ret: 1
	//BInt: 1
	BInt_ret = BInt.ModSqrt(x, y)
	//21输出：
	//x: 21
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("BInt_ret:", BInt_ret)
	fmt.Println("BInt:", BInt)
	//20输出：
	//x: 20
	//y: 7
	//BInt_ret: <nil>
	//BInt: 19

	//18输出：
	//x: 18
	//y: 7
	//BInt_ret: 2
	//BInt: 2

	//y=9时候出现bug
	//目前还不清楚这个api的具体的用法到底是怎么样的！

	//x = big.NewInt(23)
	x = big.NewInt(22)
	y = big.NewInt(7)
	BInt = big.NewInt(11)
	// ModInverse将z设置为环ℤ/nℤ中g的乘法逆，然后返回z。
	// 如果g和n不是相对质数，则g在环ℤ/nℤ中没有乘法逆。 在这种情况下，z不变，返回值为nil。
	//将z设为g相对p的模逆（即z、g满足(z * g) % p == 1）。返回值z大于0小于p。
	BInt_ret = BInt.ModInverse(x, y)

	fmt.Println()
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("BInt_ret:", BInt_ret)
	fmt.Println("BInt:", BInt)
	//输出：
	//x: 23
	//y: 7
	//BInt_ret: 4

	BInt = big.NewInt(-11)
	//通过返回x的绝对值作为低端单词切片，Bits提供对x的原始（未经检查但快速的）访问。 结果和x共享相同的基础数组。
	// Bits旨在支持该程序包之外缺少的底层Int功能的实现； 否则应避免。
	BInt_ret111 := BInt.Bits()

	fmt.Println()
	BInt_ret = BInt.ModInverse(x, y)
	fmt.Println("BInt_ret:", BInt_ret111[0])
	fmt.Println("BInt:", BInt)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("BInt_ret:", BInt_ret)
	fmt.Println("BInt:", BInt)
	//输出：
	//BInt_ret: [11]
	//BInt_ret: 11
	//BInt: -11
	//BInt_ret: [5]
	//BInt: -5,底层共享同一个数组

	// Bits旨在支持该程序包之外缺少的底层Int功能的实现； 否则应避免。
	BInt_ret111 = BInt.Bits()

	BInt_ret111 := BInt.Bits()
	fmt.Println("BInt_ret:", BInt_ret111)
	fmt.Println("BInt_ret:", BInt_ret111[0])
	fmt.Println("BInt_ret:", BInt_ret111)
	fmt.Println("BInt_ret:", BInt_ret111[0])
	fmt.Println("BInt:", BInt)
	fmt.Println("BInt_ret:", BInt_ret111)
	BInt_ret111[0] = big.Word(5)
	fmt.Println("BInt_ret:", BInt_ret111)
	fmt.Println("BInt:", BInt)
	//BInt_ret: 1123572578993467889
	//BInt: -1123572578993467889
	//BInt_ret: [5]
	//BInt: -5

	BInt = big.NewInt(0b0000_1110)
	//                   0b0000_0011
	//    &              0b0000_0001，其实就是返回BInt第2位的值
	//Bit返回x的第i位的值。 也就是说，它返回（x >> i）＆1。 位索引i必须大于等于0。
	BInt_ret222 := BInt.Bit(2)

	BInt_ret111 = BInt.Bits()
	fmt.Println("BInt_ret:", BInt_ret222)
	fmt.Println("BInt:", BInt)
	fmt.Println("BInt_ret:", BInt_ret111)
	fmt.Println("BInt_ret:", BInt_ret111[0])
	fmt.Println("BInt:", BInt)

	BInt_ret111[0] = big.Word(5)
	fmt.Println("BInt_ret:", BInt_ret111)
	fmt.Println("BInt:", BInt)

	// BitLen返回x的绝对值的长度（以位为单位）。
	// 0的位长度为0。
	BInt_ret333 := BInt.BitLen()

	fmt.Println()
	fmt.Println("BInt_ret:", BInt_ret333)
	//0b0000_0110输出：
	//BInt_ret: 3
	//BInt: 6

	//0b0000_1110输出：
	BInt_ret222 := BInt.Bit(2)
	//BInt: 14

	fmt.Println("BInt_ret:", BInt_ret222)
	fmt.Println("BInt:", BInt)
	//BInt: 70

	BInt = big.NewInt(16)
	BInt1 := big.NewInt(0b0100_0110)
	fmt.Println()
	fmt.Printf("设置位之前BInt：%b\n", BInt)
	fmt.Printf("设置位之前BInt1：%b\n", BInt1)
	// SetBit将z设置为x，将x的第i位设置为b（0或1）。
	//如果b为1，则SetBit设置z = x | （1 << i）;
	//如果b为0，则SetBit设置z = x＆^（1 << i）。 如果b不为0或1，则SetBit将发生恐慌。
	BInt_ret333 := BInt.BitLen()

	fmt.Printf("设置位之后BInt：%b\n", BInt)
	fmt.Println("BInt_ret:", BInt_ret333)
	fmt.Println("BInt:", BInt)

	BInt = big.NewInt(16)
	fmt.Println()
	fmt.Println("设置位之前BInt：", BInt.Bits())

	var a = big.Word(1)
	var b = big.Word(2)
	var ls = []big.Word{a, b}
	// SetBits通过将z的值设置为abs（解释为小尾数字切片）并返回z，提供对z的原始（未经检查但快速的）访问。 结果和绝对值共享相同的基础数组。
	// SetBits旨在支持该程序包外部缺少的底层Int功能的实现； 否则应避免。
	BInt_ret666 := BInt.SetBits(ls)

	fmt.Println("设置位之后BInt_ret666：", BInt_ret666)
	//输出：
	//设置位之前BInt： [16]
	//设置位之后BInt： 36893488147419103233
	fmt.Printf("设置位之前BInt：%b\n", BInt)
	fmt.Printf("设置位之前BInt1：%b\n", BInt1)
	fmt.Println()
	BInt = big.NewInt(16)
	fmt.Println("设置位之前BInt：", BInt)
	BInt_ret555 := BInt.SetBit(BInt1, 6, 0) //这里的数字从0开始。而上面的那个BitLen是从1开始的，修改的是副本
	//src_ls:=[]byte{0b0000_0000,0b0000_0011}
	fmt.Printf("设置位之后BInt：%b\n", BInt)
	fmt.Printf("设置位之后BInt1：%b\n", BInt1)
	fmt.Printf("BInt_ret:%b\n", BInt_ret555)
	fmt.Println("设置位之后BInt：", BInt)
	//[]byte{0b0000_0000,0b0000_0011}输出：
	//设置位之前BInt： 16
	fmt.Println("设置位之前BInt：", BInt.Bits())
	//设置位之后BInt_ret777： 3
	var a = big.Word(1)
	var b = big.Word(2)
	var ls = []big.Word{a, b}
	//设置位之后BInt： 259
	//设置位之后BInt_ret777： 259
	BInt_ret666 := BInt.SetBits(ls)
	fmt.Println()
	fmt.Println("设置位之后BInt：", BInt)
	fmt.Println("设置位之后BInt_ret666：", BInt_ret666)
	// SetInt64将z设置为x并返回z。
	BInt_ret888 := BInt.SetInt64(-22)

	fmt.Println("设置位之后BInt：", BInt)
	fmt.Println("设置位之后BInt_ret888：", BInt_ret888)
	//输出：
	//设置位之前BInt： 16
	fmt.Println("设置位之前BInt：", BInt)
	//设置位之后BInt_ret888： -22

	src_ls := []byte{0b0000_0001, 0b0000_0011}
	BInt_ret777 := BInt.SetBytes(src_ls) //将字节切片转为big.Int类型
	fmt.Println("设置位之前BInt：", BInt)
	fmt.Println("设置位之后BInt：", BInt)
	fmt.Println("设置位之后BInt_ret777：", BInt_ret777)
	BInt_ret888 = BInt.SetUint64(22)

	fmt.Println("设置位之后BInt：", BInt)
	fmt.Println("设置位之后BInt_ret888：", BInt_ret888)
	//输出：
	//设置位之前BInt： 16
	//设置位之后BInt： 22
	//设置位之后BInt_ret888： 22

	fmt.Println()
	BInt = big.NewInt(16)
	BInt999 := big.NewInt(17)
	fmt.Println("设置位之前BInt：", BInt)
	//将z设置为x并返回z。
	BInt_ret888 := BInt.SetInt64(-22)

	fmt.Println("设置位之后BInt：", BInt)
	fmt.Println("设置位之后BInt_ret888：", BInt_ret888)
	fmt.Println("设置位之后BInt_ret888：", BInt_ret888)
	//输出：
	//设置位之前BInt： 16
	//设置位之后BInt： 17
	//设置位之后BInt999： 17
	fmt.Println()
	BInt = big.NewInt(12)
	fmt.Println("设置位之前BInt：", BInt)
	// SetString将z设置为s的值，以给定的底数进行解释，并返回z和一个指示成功的布尔值。整个字符串（不仅仅是前缀）必须有效才能成功。如果SetString失败，则z的值不确定（不知道怎么理解），但返回的值为nil。
	//基本参数必须为0或2到MaxBase之间的值。
	BInt_ret888 = BInt.SetUint64(22)
	//对于<= 36的基数，小写和大写字母被认为是相同的：
	fmt.Println("设置位之后BInt：", BInt)
	fmt.Println("设置位之后BInt_ret888：", BInt_ret888)
	//对于以0为底的数字，下划线字符``_''可能出现在底前缀和相邻数字之间以及连续数字之间;这样的下划线不会更改数字的值。
	//如果没有其他错误，则将下划线的错误放置报告为错误。如果base！= 0，则下划线将不被识别，并且与其他任何无效数字字符一样起作用。

	//BInt_ret888,bl:=BInt.SetString("16",10)
	//BInt_ret888,bl:=BInt.SetString("17",2)
	fmt.Println("设置之后BInt：", BInt)
	fmt.Println("设置之后BInt：", bl)
	fmt.Println("设置之后BInt_ret888：", BInt_ret888)
	fmt.Println("设置位之前BInt：", BInt)
	//设置之前BInt： 12
	BInt_ret888 = BInt.Set(BInt999)
	//设置之后BInt： true
	fmt.Println("设置位之后BInt：", BInt)
	fmt.Println("设置位之后BInt999：", BInt999)
	fmt.Println("设置位之后BInt_ret888：", BInt_ret888)
	//设置之前BInt： 12
	//设置之后BInt： 1，如果失败这个值是不确定的
	//设置之后BInt： false
	//设置之后BInt_ret888： <nil>

	//SetString("00000011",2)输出：
	//设置之前BInt： 12
	//设置之后BInt： 3
	fmt.Println("设置之前BInt：", BInt)
	//设置之后BInt_ret888： 3

	//其他进制自己尝试

	fmt.Println()
	BInt = big.NewInt(120)
	fmt.Println("设置之前BInt：", BInt)
	// Source表示在[0，1 << 63）范围内均匀分布的伪随机int64值的源。

	// NewSource返回一个以给定值作为种子的新伪随机源。
	//与顶层函数使用的默认Source不同，此Source对于多个goroutine并发使用是不安全的。
	BInt_ret888, bl := BInt.SetString("00000011", 2) //不能带0b开头或者中间下划线
	//source1 := rand.NewSource(122)
	fmt.Println("设置之后BInt：", BInt)
	fmt.Println("设置之后BInt：", bl)
	fmt.Println("设置之后BInt_ret888：", BInt_ret888)
	rd1 := rand.New(source1)

	// Rand在[0，n）中将z设置为伪随机数并返回z。
	//由于它使用math / rand软件包，因此不得用于对安全性敏感的工作。 使用crypto / rand.Int代替。
	randBigInt := BInt.Rand(rd, BInt)
	//randBigInt11 := BInt.Rand(rd, BInt)
	randBigInt11 := BInt.Rand(rd1, BInt)

	go func() {
		source2 := rand.NewSource(120)
		rd2 := rand.New(source2)
		randBigInt22 := BInt.Rand(rd2, BInt)
		fmt.Println("随机的BigInt22是----：", randBigInt22)
	}()

	fmt.Println("随机的BigInt是----：", randBigInt)
	fmt.Println("随机的BigInt11是----：", randBigInt11)
	//source与source1相同时候输出：
	//设置之前BInt： 120
	//随机的BigInt是----： 13
	//随机的BigInt22是----： 7

	fmt.Println("设置之前BInt：", BInt)
	//设置之前BInt： 120
	//随机的BigInt是----： 26
	//随机的BigInt11是----： 26
	//随机的BigInt22是----： 7
	//不大清楚机制，不建议使用
	time.Sleep(1e9)

	fmt.Println()
	BInt = big.NewInt(120)
	fmt.Println("设置之前BInt：", BInt)
	//字节返回x的绝对值作为大端字节切片。
	dst_byte := BInt.Bytes()

	fmt.Println("big.Int转[]byte类型之前是：", BInt)
	fmt.Println("big.Int转[]byte类型后是：", dst_byte)

	//输出：
	//设置之前BInt： 120
	//big.Int转[]byte类型之前是： 120
	//big.Int转[]byte类型后是： [120]

		fmt.Println("随机的BigInt22是----：", randBigInt22)
	BInt = big.NewInt(12)
	fmt.Println("设置之前BInt：", BInt)
	fmt.Println("随机的BigInt是----：", randBigInt)
	fmt.Println("随机的BigInt11是----：", randBigInt11)

	fmt.Println("big.Int转string类型之前是：", BInt)
	fmt.Println("big.Int转string类型后是：", dst_str)
	//输出：
	//设置之前BInt： 12
	//big.Int转string类型之前是： 12
	//big.Int转string类型后是： 12

	fmt.Println()
	BInt = big.NewInt(121)
	//BInt = big.NewInt(1)
	fmt.Println("设置之前BInt：", BInt)
	// JSON封送程序仅在此处用于API向后兼容（程序会明确查找这两种方法）。 JSON仅适用于TextMarshaler。
	// MarshalJSON实现json.Marshaler接口。
	check_err_math_big(e)

	fmt.Println("设置之前BInt：", BInt)
	fmt.Println("byte_MarshalJSON:", byte_MarshalJSON)
	dst_byte := BInt.Bytes()
	//设置之前BInt： 0
	fmt.Println("big.Int转[]byte类型之前是：", BInt)
	fmt.Println("big.Int转[]byte类型后是：", dst_byte)

	//NewInt(1)输出：
	//设置之前BInt： 1
	//BInt: 1
	//byte_MarshalJSON: [49]

	//BInt: 121
	//byte_MarshalJSON: [49 50 49]，121每个位的的unicode值
	fmt.Println("设置之前BInt：", BInt)
	fmt.Println()
	dst_str := BInt.String()
	//BInt = big.NewInt(1)
	fmt.Println("big.Int转string类型之前是：", BInt)
	fmt.Println("big.Int转string类型后是：", dst_str)
	ls111 := []byte{49, 50, 49}
	e = BInt.UnmarshalJSON(ls111)
	check_err_math_big(e)
	fmt.Println("BInt:", BInt)
	//输出：
	//设置之前BInt： 11
	//BInt: 121

	fmt.Println("设置之前BInt：", BInt)
	BInt = big.NewInt(121)
	//BInt = big.NewInt(1)
	fmt.Println("设置之前BInt：", BInt)
	// MarshalText实现encoding.TextMarshaler接口。关于这个接口说明如下：
	// TextUnmarshaler是由对象实现的接口，可以解组自身的文本表示形式。
	// UnmarshalText必须能够解码MarshalText生成的格式。
	fmt.Println("BInt:", BInt)
	fmt.Println("byte_MarshalJSON:", byte_MarshalJSON)
	ls222, e := BInt.MarshalText()
	check_err_math_big(e)
	fmt.Println("BInt:", BInt)
	fmt.Println("ls222:", ls222)
	//输出：
	//设置之前BInt： 121
	//BInt: 121
	//ls222: [49 50 49]，似乎跟MarshalJson差不多也

	fmt.Println()
	BInt = big.NewInt(11)
	//BInt = big.NewInt(1)
	fmt.Println("设置之前BInt：", BInt)
	// UnmarshalText实现encoding.TextUnmarshaler接口。
	ls222 = []byte{49, 50, 49}
	check_err_math_big(e)
	fmt.Println("BInt:", BInt)
	//输出：
	fmt.Println("设置之前BInt：", BInt)
	//BInt: 121
	ls111 := []byte{49, 50, 49}
	fmt.Println()
	BInt = big.NewInt(15)
	fmt.Println("BInt:", BInt)
	//文本返回给定基数中x的字符串表示形式。
	//底数必须介于2到62之间（含2和62）。 结果将小写字母'a'至'z'用于数字值10至35，将大写字母'A'至'Z'用于数字值36至61。
	//没有前缀（例如“ 0x”）添加到字符串中。 如果x是nil指针，则返回“ <nil>”。
	str10 := BInt.Text(10)
	str2 := BInt.Text(2)
	str16 := BInt.Text(16)

	fmt.Println("设置之前BInt：", BInt)
	fmt.Println("str111:", str10)
	fmt.Println("str8:", str8)
	fmt.Println("str2:", str2)
	fmt.Println("str16:", str16)
	//输出：
	ls222, e := BInt.MarshalText()
	//BInt: 15
	fmt.Println("BInt:", BInt)
	fmt.Println("ls222:", ls222)
	//str2: 1111
	//str16: f

	fmt.Println()
	BInt = big.NewInt(15)
	//        ^               0b1111_0110
	fmt.Println("BInt之前:", BInt)
	fmt.Println("BInt_arg之前:", BInt_arg)
	fmt.Println("设置之前BInt：", BInt)
	Int_ret := BInt.Not(BInt_arg)
	ls222 = []byte{49, 50, 49}
	fmt.Println("BInt之后:", BInt)
	fmt.Println("BInt_arg之后:", BInt_arg)
	fmt.Println("BInt:", BInt)
	//输出：
	//BInt之前: 15
	//BInt_arg之前: 9
	//BInt之后: -10
	//返回值Int_ret: -10

	fmt.Println("设置之前BInt：", BInt)
	BInt = big.NewInt(15)
	BInt_arg = big.NewInt(123)

	fmt.Println("BInt之前:", BInt)
	fmt.Println("BInt_arg之前:", BInt_arg)
	// Neg将z设置为-x并返回z。
	Int_ret = BInt.Neg(BInt_arg)

	fmt.Println("BInt:", BInt)
	fmt.Println("str111:", str10)
	fmt.Println("str8:", str8)
	fmt.Println("str2:", str2)
	fmt.Println("str16:", str16)
	//BInt_arg之前: 123
	//BInt之后: -123
	//BInt_arg之后: 123
	//返回值Int_ret: -123

	fmt.Println()
	BInt = big.NewInt(15)

	//如果x不能用uint64表示，则结果不确定。
	uint64_exmp := BInt.Uint64()

	fmt.Printf("BInt之后的类型是:%T,值是:%v\n", BInt, BInt)
	fmt.Println("BInt之前:", BInt)
	fmt.Println("BInt_arg之前:", BInt_arg)
	//BInt之前: 15
	Int_ret := BInt.Not(BInt_arg)
	//返回值uint64_exmp的类型是:uint64,值是:15
	fmt.Println("BInt之后:", BInt)
	fmt.Println("BInt_arg之后:", BInt_arg)
	fmt.Println("返回值Int_ret:", Int_ret)

	fmt.Println("BInt之前:", BInt)
	// Uint64返回x的uint64表示形式。
	//如果x不能用uint64表示，则结果不确定。
	int64_exmp := BInt.Int64()

	fmt.Printf("BInt之后的类型是:%T,值是:%v\n", BInt, BInt)
	//输出：
	//BInt之前: 15
	BInt_arg = big.NewInt(123)
	//返回值int64_exmp的类型是:int64,值是:15
	fmt.Println("BInt之前:", BInt)
	fmt.Println("BInt_arg之前:", BInt_arg)
	BInt = big.NewInt(15)
	Int_ret = BInt.Neg(BInt_arg)

	fmt.Println("BInt之后:", BInt)
	fmt.Println("BInt_arg之后:", BInt_arg)
	fmt.Println("返回值Int_ret:", Int_ret)
	uint64_bool := BInt.IsUint64()

	int64_bool1 := BInt1.IsInt64()
	uint64_bool1 := BInt1.IsUint64()

	fmt.Printf("BInt之后的类型是:%T,值是:%v\n", BInt, BInt)
	fmt.Printf("返回值int64_bool的类型是:%T,值是:%v\n", int64_bool, int64_bool)
	fmt.Printf("返回值uint64_bool1的类型是:%T,值是:%v\n", uint64_bool1, uint64_bool1)
	//输出：
	//BInt之前: 15
	fmt.Println("BInt之前:", BInt)
	//返回值int64_bool的类型是:bool,值是:true
	//返回值uint64_bool的类型是:bool,值是:true
	uint64_exmp := BInt.Uint64()
	//返回值uint64_bool1的类型是:bool,值是:false
	fmt.Printf("BInt之后的类型是:%T,值是:%v\n", BInt, BInt)
	fmt.Printf("返回值uint64_exmp的类型是:%T,值是:%v\n", uint64_exmp, uint64_exmp)
	BInt = big.NewInt(15)

	fmt.Println("BInt之前:", BInt)
	//将由x.Text（base）生成的x的字符串表示形式追加到buf，并返回扩展的缓冲区。
	buf := make([]byte, 3)
	buf = BInt.Append(buf, 10)

	fmt.Println("BInt之后:", BInt)
	fmt.Println("BInt之前:", BInt)
	//fmt.Println("追加到字节切片后的切片为：",ls333)
	//输出：
	int64_exmp := BInt.Int64()
	//追加之前的字节切片buf： [0 0 0]
	fmt.Printf("BInt之后的类型是:%T,值是:%v\n", BInt, BInt)
	fmt.Printf("返回值int64_exmp的类型是:%T,值是:%v\n", int64_exmp, int64_exmp)
	//追加到字节切片后的切片为： [0 0 0 49 53]，说明不是原地追加的，当然也可以原地追加，看接受者

	fmt.Println()
	BInt = big.NewInt(15)
	BInt1 = big.NewInt(13)
	fmt.Println("BInt之前:", BInt)
	//二项式将z设置为（n，k）的二项式系数，并返回z。
	//具体请看：https://baike.baidu.com/pic/%E4%BA%8C%E9%A1%B9%E5%BC%8F%E7%B3%BB%E6%95%B0/6763242/0/377adab44aed2e73b9e960568c01a18b87d6fa4c?fr=lemma&ct=single#aid=0&pic=377adab44aed2e73b9e960568c01a18b87d6fa4c
	BInt1 = BInt.Binomial(4, 2)
	fmt.Println("BInt之前:", BInt)
	fmt.Println("BInt之后:", BInt)
	int64_bool := BInt.IsInt64()
	uint64_bool := BInt.IsUint64()
	//BInt之前: 15
	int64_bool1 := BInt1.IsInt64()
	uint64_bool1 := BInt1.IsUint64()

	fmt.Printf("BInt之后的类型是:%T,值是:%v\n", BInt, BInt)
	fmt.Printf("返回值int64_bool的类型是:%T,值是:%v\n", int64_bool, int64_bool)
	fmt.Printf("返回值uint64_bool的类型是:%T,值是:%v\n", uint64_bool, uint64_bool)
	fmt.Printf("返回值int64_bool1的类型是:%T,值是:%v\n", int64_bool1, int64_bool1)
	fmt.Printf("返回值uint64_bool1的类型是:%T,值是:%v\n", uint64_bool1, uint64_bool1)
	fmt.Println("BInt之前:", BInt)
	// GCD将z设置为a和b的最大公约数（都必须大于0），并返回z。
	//如果x或y不为nil，则GCD会将其值设置为z = a * x + b * y。
	//如果a或b均<= 0，则GCD设置z = x = y = 0。
	BInt2 := BInt_r.GCD(nil, nil, BInt, BInt1)

	fmt.Println("BInt之后:", BInt)
	fmt.Println("BInt1之后:", BInt1)
	//输出：
	//BInt之前: 12
	//BInt之后: 12
	fmt.Println("BInt之前:", BInt)
	//BInt2之后: 4
	buf := make([]byte, 3)
	fmt.Println("追加之前的字节切片buf：", buf)
	fmt.Println()
	buf = BInt.Append(buf, 10)
	BInt1 = big.NewInt(8)
	fmt.Println("BInt之后:", BInt)
	fmt.Println("追加之后的字节切片buf：", buf)
	fmt.Println("BInt之前:", BInt)
	// GCD将z设置为a和b的最大公约数（都必须大于0），并返回z。
	//如果x或y不为nil，则GCD会将其值设置为z = a * x + b * y。
	//如果a或b均<= 0，则GCD设置z = x = y = 0。
	BInt2 = BInt_r.GCD(big.NewInt(-2), big.NewInt(-4), BInt, BInt1)

	fmt.Println("BInt之后:", BInt)
	fmt.Println("BInt1之后:", BInt1)
	//输出：
	//BInt之前: 12
	//BInt之后: 12
	//BInt1之后: 8
	fmt.Println("BInt之前:", BInt)
	//BInt_r之后: 4，似乎x和y并不起作用

	BInt1 = BInt.Binomial(4, 2)
	BInt = big.NewInt(12)
	fmt.Println("BInt之后:", BInt)
	fmt.Println("BInt1之后:", BInt1)
	//或设置z = x | y并返回z。
	BInt2 = BInt.Or(big.NewInt(0b0000_0001), big.NewInt(0b0000_0011))

	fmt.Println("BInt之后:", BInt)
	fmt.Println("BInt2之后:", BInt2)
	//BInt之前: 12
	//BInt之后: 3
	//BInt2之后: 3

	fmt.Println()
	fmt.Println("BInt之前:", BInt)

	fmt.Println("BInt之前:", BInt)
	// Quo将z设为y！= 0时候x / y的商并返回z。
	BInt2 := BInt_r.GCD(nil, nil, BInt, BInt1)
	// Quo实现了截断的除法（例如Go）； 有关更多详细信息，请参见QuoRem。
	fmt.Println("BInt之后:", BInt)
	fmt.Println("BInt1之后:", BInt1)
	fmt.Println("BInt2之后:", BInt2)
	fmt.Println("BInt_r之后:", BInt_r)
	fmt.Println("BInt2之后:", BInt2)
	//2输出：
	//BInt之前: 12
	//BInt之后: 4
	//BInt2之后: 4

	//3输出：
	//BInt之后: 2
	//BInt2之后: 2

	fmt.Println()
	BInt = big.NewInt(12)
	fmt.Println("BInt之前:", BInt)

	fmt.Println("BInt之前:", BInt)
	// QuoRem将z设置为x / y的商，将r设置为x％y的余数，并返回y！= 0的（z，r）对。
	BInt2 = BInt_r.GCD(big.NewInt(-2), big.NewInt(-4), BInt, BInt1)
	// QuoRem实现T划分和模数（如Go）：
	fmt.Println("BInt之后:", BInt)
	fmt.Println("BInt1之后:", BInt1)
	fmt.Println("BInt2之后:", BInt2)
	fmt.Println("BInt_r之后:", BInt_r)
	BInt_re1, BInt_re2 := BInt.QuoRem(big.NewInt(17), big.NewInt(5), BInt1)

	fmt.Println("BInt之后:", BInt)
	fmt.Println("BInt1之后:", BInt1)
	fmt.Println("BInt_re1之后:", BInt_re1)
	fmt.Println("BInt_re2之后:", BInt_re2)
	//输出：
	//BInt1之后: 2
	//BInt_re1之后: 3
	//BInt_re2之后: 2
	fmt.Println("BInt之前:", BInt)
	fmt.Println()
	BInt2 = BInt.Or(big.NewInt(0b0000_0001), big.NewInt(0b0000_0011))

	fmt.Println("BInt之后:", BInt)
	fmt.Println("BInt2之后:", BInt2)
	//如果y == 0，则会发生除以零的运行时恐慌。
	// Rem实现了截断的模数（如Go）； 有关更多详细信息，请参见QuoRem。
	BInt_re1 = BInt.Rem(big.NewInt(17), big.NewInt(5))

	fmt.Println("BInt之后:", BInt)
	//输出：
	//BInt之前: 12
	//BInt之后: 2
	fmt.Println("BInt之前:", BInt)

	fmt.Println()
	BInt = big.NewInt(12)

	BInt2 = BInt.Quo(big.NewInt(8), big.NewInt(3))
	// Xor设置z = x ^ y并返回z。异或的意思，相同为0，不同为1
	fmt.Println("BInt之后:", BInt)
	fmt.Println("BInt2之后:", BInt2)
	//                                  0000_0010
	fmt.Println("BInt之后:", BInt)
	fmt.Println("BInt_re1之后:", BInt_re1)

	//输出：
	//BInt之前: 12
	//BInt之后: 2
	//BInt_re1之后: 2

	fmt.Println()
	BInt = big.NewInt(0)
	fmt.Println("BInt之前:", BInt)
	//符号返回：
	//如果x <0则为-1
	fmt.Println("BInt之前:", BInt)
	//如果x> 0，则+1
	int_ret111 := BInt.Sign()

	fmt.Println("BInt之后:", BInt)
	fmt.Println("int_ret111之后:", int_ret111)
	//12输出：
	//BInt之前: 12
	BInt_re1, BInt_re2 := BInt.QuoRem(big.NewInt(17), big.NewInt(5), BInt1)
	//int_ret111之后: 1
	fmt.Println("BInt之后:", BInt)
	fmt.Println("BInt1之后:", BInt1)
	fmt.Println("BInt_re1之后:", BInt_re1)
	fmt.Println("BInt_re2之后:", BInt_re2)
	//int_ret111之后: -1

	//0输出：
	//BInt之前: 0
	//BInt之后: 0
	//int_ret111之后: 0

	//BInt = big.NewInt(0b0000_1001)//0
	//BInt = big.NewInt(0b0000_1000)//3
	//BInt = big.NewInt(0b1000_1000)//3
	fmt.Println("BInt之前:", BInt)
	BInt = big.NewInt(0b1001_0000) //3
	fmt.Println("BInt之前:", BInt)
	// TrailingZeroBits返回|x|的连续最低有效零位的数目。也就是二进制中为1的最后一个位的索引值
	BInt_re1 = BInt.Rem(big.NewInt(17), big.NewInt(5))

	fmt.Println("BInt之后:", BInt)
	fmt.Println("BInt_re1之后:", BInt_re1)
	//输出：
	//BInt之前: 144
	//BInt之后: 144
	//int_ret222之后: 4

	//BInt = big.NewInt(122)
	//BInt = big.NewInt(122023563)
	fmt.Println("BInt之前:", BInt)
	fmt.Println("BInt之前:", BInt)

	BInt_re1 = BInt.Xor(big.NewInt(0b0000_0001), big.NewInt(0b0000_0011))
	check_err_math_big(e)

	fmt.Println("BInt之后:", BInt)
	fmt.Println("BInt_re1之后:", BInt_re1)
	//2输出:
	//BInt之前: 2
	//BInt之后: 2
	//ls333切片为: [2 2]

	//122输出:
	//BInt之前: 122
	//BInt之后: 122
	//ls333切片为: [2 122]

	fmt.Println("BInt之前:", BInt)
	//BInt之前: 122023563
	//BInt之后: 122023563
	//ls333切片为: [2 7 69 238 139]

	int_ret111 := BInt.Sign()
	BInt = big.NewInt(2)
	fmt.Println("BInt之后:", BInt)
	fmt.Println("int_ret111之后:", int_ret111)
	fmt.Println("BInt之前:", BInt)

	// GobDecode实现gob.GobDecoder接口。
	//e= BInt.GobDecode([]byte{2,122})
	e = BInt.GobDecode([]byte{2, 7, 69, 238, 139})
	check_err_math_big(e)

	fmt.Println("BInt之后:", BInt)
	//[]byte{2,122}输出：
	//BInt之前: 2
	//BInt之后: 122

	//[]byte{2 ,7, 69, 238, 139}输出：
	//BInt之前: 2
	//BInt之后: 122023563
	fmt.Println()
	//BInt = big.NewInt(11)
	BInt = big.NewInt(12)
	fmt.Println("BInt之前:", BInt)

	BInt = big.NewInt(0b1001_0000) //3
	fmt.Println("BInt之前:", BInt)
	//如果x是随机选择而不是素数，则ProbablyPrime可能返回false。
	int_ret222 := BInt.TrailingZeroBits()
	//如果输入小于2⁶⁴，则Prime的准确度为100％。
	fmt.Println("BInt之后:", BInt)
	fmt.Println("int_ret222之后:", int_ret222)
	//从Go 1.8开始，ProbablyPrime（0）被允许，并且仅应用Baillie-PSW测试。
	//在Go 1.8之前，ProbablyPrime仅应用了Miller-Rabin测试，而ProbablyPrime（0）则惊慌失措。
	b111 := BInt.ProbablyPrime(13) //这个n不知道是什么

	fmt.Println("BInt之后:", BInt)
	//BInt之前: 11
	//BInt之后: 11
	//b111之后: true

	fmt.Println("BInt之前:", BInt)
	//BInt之前: 12
	//BInt之后: 12
	ls333, e := BInt.GobEncode()

	//关于big.Int类型下的方法还剩余2/3个方法没展示，
	fmt.Println("BInt之后:", BInt)
	fmt.Println("ls333切片为:", ls333)
	//对于和big.Int有着共同名字的方法不再累叙，下面仅仅列出big.float特有的区别于big.int的方法

	//非零的有限浮点数表示一个多精度浮点数
	//
	//符号×尾数×2 **指数
	//
	// 0.5 <=尾数<1.0，并且MinExp <=指数<= MaxExp。
	//浮点数也可以是零（+0，-0）或无限（+ Inf，-Inf）。
	//所有浮点数都是有序的，两个浮点数x和y的顺序由x.Cmp（y）定义。
	//
	//每个Float值还具有精度位数，舍入模式和精度误差。
	//精度位数是可用于表示该值的最大尾数位数。舍入模式指定如何将结果舍入以适合尾数位，而精度误差描述相对于精确结果的舍入误差。
	//
	//除非另有说明，否则为结果指定* Float变量的所有操作（包括设置器）（通常通过MantExp除外通过接收器）都将根据结果变量的精度位数和舍入模式舍入数值结果。
	//
	// 因此，作为结果参数提供的未初始化的Float会将其精度位数设置为由操作数确定的合理值，并且其模式为RoundingMode（ToNearestEven）的零值。
	//
	//通过将所需的精度设置为24或53并使用匹配的舍入模式（通常为ToNearestEven），对于与普通（即非正常）浮点数相对应的操作数，Float运算产生的结果与相应的float32或float64 IEEE-754运算法则相同或float64数字。
	//指数下溢和上溢导致值为0或无穷大（与IEEE-754不同），因为浮点指数的范围更大。
	fmt.Println("BInt之前:", BInt)
	//可以立即使用浮点数的零（未初始化）值，该值精确地表示数字+0.0，精度为0，舍入模式为ToNearestEven。
	//
	//操作始终使用指针参数（* Float）而不是Float值，并且每个唯一的Float值都需要其自己的唯一* Float指针。要“复制” Float值，必须使用Float.Set方法将现有（或新分配的）Float设置为新值；不支持Float的浅表副本，可能会导致错误。
	e = BInt.GobDecode([]byte{2, 7, 69, 238, 139})
	// NewFloat分配并返回设置为x的新Float，精度为53，舍入模式为ToNearestEven。返回的是指针
	//如果x是NaN，NewFloat会因ErrNaN出现恐慌。
	fmt.Println("BInt之后:", BInt)
	BFloat := big.NewFloat(16.6)
	FormatP(BFloat)

	// Signbit报告x是否是负数（负零和正数返回false）。
	fmt.Println(BFloat.Signbit())
	BFloat = big.NewFloat(-16.6)
	fmt.Println(BFloat.Signbit())
	BFloat = big.NewFloat(-0)
	//输出：
	//*big.Float--16.6
	//false
	fmt.Println("BInt之前:", BInt)
	//false

	fmt.Println()
	BFloat = big.NewFloat(-16.6)
	// Mode返回x的舍入模式。
	fmt.Println(BFloat, BFloat.Mode())

	BFloat = big.NewFloat(16.63467823336762323)
	// Mode返回x的舍入模式。
	fmt.Println(BFloat, BFloat.Mode())
	b111 := BInt.ProbablyPrime(13) //这个n不知道是什么
	BFloat = big.NewFloat(16)
	fmt.Println("BInt之后:", BInt)
	fmt.Println("b111之后:", b111)

	mode := BFloat.Mode()
	fmt.Println(mode.String()) //只有这么一个方法
	//输出：
	//-16.6 ToNearestEven
	//16.634678233367623 ToNearestEven
	//16 ToNearestEven,看来默认的就是这种舍入模式
	//ToNearestEven

	fmt.Println()
	BFloat = big.NewFloat(2.1)
	// SetMode将z的舍入模式设置为mode并返回精确的z。
	fmt.Println(BFloat, BFloat.SetMode(big.ToNearestAway))
	fmt.Println(BFloat, BFloat.Mode(), BFloat.SetPrec(100))

	BFloat = big.NewFloat(2.1)
	fmt.Println(BFloat, BFloat.SetMode(big.ToZero))
	fmt.Println(BFloat, BFloat.Mode(), BFloat.SetPrec(100))

	BFloat = big.NewFloat(2.1)
	fmt.Println(BFloat, BFloat.SetMode(big.ToPositiveInf))
	fmt.Println(BFloat, BFloat.Mode(), BFloat.SetPrec(100))

	BFloat = big.NewFloat(2.1)
	fmt.Println(BFloat, BFloat.SetMode(big.ToNegativeInf))
	fmt.Println(BFloat, BFloat.Mode(), BFloat.SetPrec(100))
	//输出：
	//-1.6700006458e+06 -1.6700006458e+06
	//-1.6e+06 ToNearestAway -1.6e+06
	//-1.6e+06 -1.6e+06
	//-1.6e+06 ToZero -1.6e+06
	//-1.6e+06 -1.6e+06
	//-1.6e+06 ToPositiveInf -1.6e+06
	//-1.6e+06 -1.6e+06
	//-1.6e+06 ToNegativeInf -1.6e+06
	//还是不大清楚这到底是怎么算的！！！但是这个位数基本可以确定是100bit的意思，大概就是采用100bit的长度来存储这个小数

	// MantExp将x分解为其尾数和指数成分，并返回指数。 如果提供了非nil mant参数，则将其值设置为x的尾数，其精度和舍入模式与x相同。 分量满足x == mant×2 ** exp，其中0.5 <= | mant | <1.0。
	//特殊情况是：
	//（±0）.MantExp（mant）= 0，mant设置为±0
	//（±Inf）.MantExp（mant）= 0，且mant设置为±Inf
	// x和mant可以相同，在这种情况下x设置为其尾数值。
	BFloat.SetMode(big.ToNearestEven)
	BFloat = big.NewFloat(13.46)
	BFloat_111 := big.NewFloat(13.46)
	fmt.Println(BFloat, BFloat_111, BFloat.MantExp(BFloat_111))
	//输出
	//13.46 0.84125 4
	//真的不知道什么作用

	fmt.Println()
	BFloat = big.NewFloat(13.46)
	// IsInt报告x是否为整数。也就是小数部分的位数全部为0
	//±Inf值不是整数。
	fmt.Println(BFloat, BFloat.IsInt())

	BFloat = big.NewFloat(13)
	// IsInt报告x是否为整数。
	//±Inf值不是整数。
	fmt.Println(BFloat, BFloat.Mode())

	BFloat = big.NewFloat(13.000)
	// IsInt报告x是否为整数。
	fmt.Println(BFloat, BFloat.Mode())
	fmt.Println(BFloat, BFloat.IsInt())

	//如果符号> = 0，则Inf返回正无穷大；如果符号<= 0，则Inf返回负无穷大。
	fmt.Println(BFloat, BFloat.Mode())
	fmt.Println(BFloat, BFloat.IsInt())
	mode := BFloat.Mode()
	fmt.Println(mode.String()) //只有这么一个方法
	fmt.Println(BFloat, BFloat.IsInt())

	BFloat = big.NewFloat(math.Inf(0))
	fmt.Println(BFloat, BFloat.IsInt())
	//输出：
	//13.46 false
	//13 true
	//+Inf false
	//-Inf false
	//+Inf false

	fmt.Println()
	BFloat = big.NewFloat(2.1)
	fmt.Println(BFloat)
	fmt.Println(BFloat, BFloat.SetMode(big.ToNearestAway))
	fmt.Println(BFloat, BFloat.Mode(), BFloat.SetPrec(100))
	fmt.Println(BFloat)
	// Acc返回最近一次操作产生的x的（此时的）精度。
	fmt.Println(BFloat, BFloat.SetMode(big.ToZero))
	fmt.Println(BFloat, BFloat.Mode(), BFloat.SetPrec(100))
	//2.1
	//2.1 Exact
	fmt.Println(BFloat, BFloat.SetMode(big.ToPositiveInf))
	fmt.Println(BFloat, BFloat.Mode(), BFloat.SetPrec(100))

	fmt.Println()
	fmt.Println(BFloat, BFloat.SetMode(big.ToNegativeInf))
	fmt.Println(BFloat, BFloat.Mode(), BFloat.SetPrec(100))
	// Acc返回最近一次操作产生的x的精度误差描述。
	fmt.Println(BFloat, BFloat.Acc())
	//输出：
	//2.1
	//2.1 Exact
	//2.100000000000000088817841970013
	//2.100000000000000088817841970013 Exact

	fmt.Println()
	BFloat = big.NewFloat(223256.134663344566)
	// MinPrec返回精确表示x所需的最小精度（即，x.SetPrec（prec）开始对x取整之前的最小prec）。
	fmt.Println(BFloat, BFloat.MinPrec())

	BFloat = big.NewFloat(2.1)
	fmt.Println(BFloat, BFloat.MinPrec())

	BFloat = big.NewFloat(0.1)
	fmt.Println(BFloat, BFloat.MinPrec())

	BFloat = big.NewFloat(0.0)
	fmt.Println(BFloat, BFloat_111, BFloat.MantExp(BFloat_111))
	//输出：
	//223256.13466334456 53
	//2.1 53，表示一个小数并不是2bit就可以了
	//0.1 52

	fmt.Println()
	BFloat = big.NewFloat(2.1)
	// Prec返回x的尾数精度（以位为单位）。
	fmt.Println(BFloat, BFloat.IsInt())
	fmt.Println(BFloat, BFloat.Prec())

	BFloat = big.NewFloat(223256.134663344566)
	// Prec返回x的尾数精度（以位为单位）。
	fmt.Println(BFloat, BFloat.IsInt())
	fmt.Println(BFloat, BFloat.Prec())

	BFloat = big.NewFloat(0.0)
	fmt.Println(BFloat, BFloat.Prec())
	fmt.Println(BFloat, BFloat.IsInt())
	//2.1 53
	//223256.13466334456 53
	//0 53，一般为53或者24
	fmt.Println(BFloat, BFloat.IsInt())
	fmt.Println()
	BFloat = big.NewFloat(2.1)
	fmt.Println(BFloat, BFloat.IsInt())
	//如果x不是Inf，则结果为Exact。
	//如果提供了非* Rat参数z，则Rat将结果存储在z中，而不是分配新的Rat。
	fmt.Println(BFloat, BFloat.IsInt())
	fmt.Println(BFloat)
	fmt.Println(BFloat.Rat(BRat))
	//输出：
	//2.1
	//4728779608739021/2251799813685248 Exact

	fmt.Println()
	BFloat = big.NewFloat(2.1)
	BFloat222 := big.NewFloat(22.15)
	//复制将z设置为x，其精度，舍入模式和精度与x相同，并返回z。 即使z和x相同，x也不会改变。
	BFloat222_ret := BFloat.Copy(BFloat222)
	fmt.Println(BFloat, BFloat.Acc())
	fmt.Println(BFloat, BFloat222_ret)
	BFloat.SetPrec(100)
	fmt.Println(BFloat.Prec(), BFloat222_ret.Prec())
	fmt.Println(BFloat, BFloat.Acc())

	BFloat = big.NewFloat(99.22)
	fmt.Println(BFloat.Prec(), BFloat222_ret.Prec())
	fmt.Println(BFloat, BFloat222_ret)
	//输出：
	//53 53
	//22.15 22.15
	//100 100
	//22.1499999999999985789145284798 22.1499999999999985789145284798
	//看来是深复制，底层不共用一个内存
	fmt.Println(BFloat, BFloat.Acc())
	fmt.Println()
	BFloat.SetPrec(53)
	BFloat = big.NewFloat(2.1)
	fmt.Println(BFloat, BFloat.Prec())
	// SetRat将z设置为x的（可能是四舍五入的）值并返回z。
	//如果z的精度为0，则将其更改为a.BitLen（），b.BitLen（）或64中的最大值； x = a / b。
	BRat111 := big.NewRat(2, 5)
	fmt.Println(BFloat, BFloat.SetRat(BRat111))
	fmt.Println(BFloat, BFloat.Prec())
	//输出：
	fmt.Println(BFloat, BFloat.MinPrec())
	//0.4 0.4
	//0.4 53
	fmt.Println(BFloat, BFloat.MinPrec())
	fmt.Println()
	BFloat.SetPrec(53)
	fmt.Println(BFloat, BFloat.MinPrec())
	BFloat333 := big.NewFloat(3)
	fmt.Println(BFloat, BFloat.Prec())
	fmt.Println(BFloat, BFloat.MinPrec())
	// SetMantExp将z设置为mant×2 ** exp并返回z。
	//结果z具有与mant相同的精度和舍入模式。 SetMantExp是MantExp的倒数，但不需要0.5 <= | mant | <1.0。 特别：
	// mant：= new（Float）
	// new（Float）.SetMantExp（mant，x.MantExp（mant））。Cmp（x）== 0
	//特殊情况是：
	// z.SetMantExp（±0，exp）=±0
	// z.SetMantExp（±Inf，exp）=±Inf
	//在将z的指数设置为exp的情况下，z和mant可以相同。
	fmt.Println(BFloat, BFloat.SetMantExp(BFloat333, 4)) //3*2**4=48
	fmt.Println(BFloat, BFloat.Prec())
	fmt.Println(BFloat, BFloat.Prec())
	//输出：
	//2.1 53
	//3 53
	//48 48
	fmt.Println(BFloat, BFloat.Prec())
	//3 53

	fmt.Println(BFloat, BFloat.Prec())
	BFloat = big.NewFloat(2.1)
	fmt.Println(BFloat, BFloat.Prec())
	//如果设置了signbit，则SetInf将z设置为无限Float -Inf；
	// 如果未设置signbit，则将+ Inf设置为+ Inf，并返回z。 z的精度不变，结果始终为“Exact”。
	fmt.Println(BFloat, BFloat.SetInf(true))
	fmt.Println(BFloat, BFloat.SetInf(false))
	fmt.Println(BFloat, BFloat.Prec())
	//输出：
	//2.1 53
	//-Inf -Inf
	BRat := big.NewRat(2, 5)
	//+Inf 53

	fmt.Println()
	//解析s，该s必须包含浮点数的文本表示形式，该浮点数的文本表示形式为给定转换基数（尾数始终为十进制数），或者是表示无穷大值的字符串。
	//对于以0为底的数字，下划线字符``_''可能出现在底前缀和相邻数字之间以及连续数字之间;这样的下划线不会更改数字的值或返回的数字计数。如果没有其他错误，则将下划线的错误放置报告为错误。如果基数！= 0，则不识别下划线，并像其他不是有效小数点或数字的字符一样终止扫描。
	//将z设置为相应浮点值的（可能是四舍五入的）值，并返回z，实际基数b和错误err（如果有）。
	//为了成功，必须使用整个字符串（而不仅仅是前缀）。
	//如果z的精度为0，则在四舍五入生效之前将其更改为64。
	//数字必须采用以下形式：
	//数字= [符号]（float |“ inf” |“ Inf”）。
	BFloat222_ret := BFloat.Copy(BFloat222)
	fmt.Println(BFloat.Prec(), BFloat222_ret.Prec())
	fmt.Println(BFloat, BFloat222_ret)
	//尾数=数字“。” [位数] |数字| “。”数字。
	fmt.Println(BFloat.Prec(), BFloat222_ret.Prec())
	fmt.Println(BFloat, BFloat222_ret)
	//数字=数字{[[“ _”]数字}。
	// digit =“ 0” ...“ 9” | “ a” ...“ z” | “ A” ...“ Z”。
	fmt.Println(BFloat.Prec(), BFloat222_ret.Prec())
	fmt.Println(BFloat, BFloat222_ret)
	//“ p”或“ P”指数表示以2为底（而不是以10为底）的指数；
	//例如，“ 0x1.fffffffffffffpp1023”（使用基数0）表示最大float64值。对于十六进制尾数，指数字符必须为'p'或'P'（如果存在）（不能将“ e”或“ E”指数指示符与尾数位区分开）。
	//返回的* Float f为nil，并且z的值有效，但如果报告错误，则未定义。
	BFloat = big.NewFloat(2.1)
	fmt.Println(BFloat)
	fmt.Println(BFloat.Parse("16", 10))
	fmt.Println(BFloat)

	fmt.Println(BFloat)
	fmt.Println(BFloat.Parse("00000010", 8)) //将八进制的00000010转为十进制表示，下同上同
	fmt.Println(BFloat, BFloat.Prec())

	fmt.Println(BFloat)
	BRat111 := big.NewRat(2, 5)
	fmt.Println(BFloat, BFloat.SetRat(BRat111))
	fmt.Println(BFloat, BFloat.Prec())
	fmt.Println(BFloat)
	fmt.Println(BFloat.Parse("ff", 16))
	fmt.Println(BFloat)
	//输出：
	//2.1
	//16 10 <nil>
	//16
	//16
	//8 8 <nil>
	fmt.Println(BFloat, BFloat.Prec())
	fmt.Println(BFloat333, BFloat333.Prec())
	//2 2 <nil>
	//2
	//2
	//255 16 <nil>
	//255

	fmt.Println("----------big.NewRat-----------")
	// NewRat用分子a和分母b创建一个新的Rat。
	fmt.Println(BFloat, BFloat.SetMantExp(BFloat333, 4)) //3*2**4=48
	fmt.Println(BFloat, BFloat.Prec())
	fmt.Println(BFloat333, BFloat333.Prec())
	// IsInt报告x的分母是否为1。
	fmt.Println("分数的分母是否为1？：", BRat_test1.IsInt())
	BRat_test1 = big.NewRat(2, 1)
	fmt.Println("分数的分母是否为1？：", BRat_test1.IsInt())
	//输出：
	//*big.Rat--2/5
	//分数的分母是否为1？： false
	fmt.Println()
	BRat_test1 = big.NewRat(2, 5)
	fmt.Println(BFloat, BFloat.Prec())
	//结果是对x的分母的引用； 如果将新值分配给x，它可能会更改，反之亦然。
	fmt.Println("分数为：", BRat_test1, "分数的分母为：", BRat_test1.Denom())
	fmt.Println(BFloat, BFloat.SetInf(true))
	fmt.Println(BFloat, BFloat.SetInf(false))
	fmt.Println(BFloat, BFloat.Prec())
	fmt.Println("分数为：", BRat_test1, "分数的分母为：", BRat_test1.Denom())
	BRat_test1 = big.NewRat(2, -1)
	fmt.Println("分数为：", BRat_test1, "分数的分母为：", BRat_test1.Denom())
	//输出：
	//分数为： 2/5 分数的分母为： 5
	//分数为： -1/2 分数的分母为： 2

	fmt.Println()
	BRat_test1 = big.NewRat(2, 5)
	BRat_test2 := big.NewRat(2, 1)
	// Inv将z设置为1 / x并返回z。
	fmt.Println("分数为：", BRat_test1, "分数的分母为：", BRat_test1.Inv(BRat_test2))
	//输出：
	//分数为： 1/2 分数的分母为： 1/2
	//从这里可以看得出BRat_test1.Inv是先于BRat_test1来计算的！

	fmt.Println()
	BRat_test1 = big.NewRat(2, 5)
	// Num返回x的分子； 它可能<= 0。
	//结果是对x分子的引用； 如果将新值分配给x，它可能会更改，反之亦然。
	//分子的符号与x的符号相对应。
	fmt.Println("分数为：", BRat_test1, "分数的分母为：", BRat_test1.Num())
	BRat_test1 = big.NewRat(2, -5)
	fmt.Println("分数为：", BRat_test1, "分数的分母为：", BRat_test1.Num())
	BRat_test1 = big.NewRat(2, -1)
	fmt.Println("分数为：", BRat_test1, "分数的分母为：", BRat_test1.Num())
	BRat_test1 = big.NewRat(2, -4)
	fmt.Println("分数为：", BRat_test1, "分数的分母为：", BRat_test1.Num())
	//输出：
	fmt.Println(BFloat.Parse("16", 10))
	//分数为： -2/5 分数的分母为： -2
	//分数为： -2/1 分数的分母为： -2
	//分数为： -1/2 分数的分母为： -1
	fmt.Println(BFloat.Parse("00000010", 8)) //将八进制的00000010转为十进制表示，下同上同
	fmt.Println()
	BRat_test1 = big.NewRat(2, 5)
	// SetFrac将z设置为a / b并返回z。
	fmt.Println(BFloat.Parse("00000010", 2))
	//fmt.Println("分数为：",BRat_test1,"执行后的分数为：",BRat_test1.SetFrac(big.NewInt(2),big.NewInt(0)))//报错，panic: division by zero
	//输出：
	//分数为： -1/2 执行后的分数为： -1/2
	fmt.Println(BFloat.Parse("ff", 16))
	fmt.Println()
	BRat_test1 = big.NewRat(2, 5)
	// SetFrac64将z设置为a / b并返回z。
	fmt.Println("分数为：", BRat_test1, "执行后的分数为：", BRat_test1.SetFrac64(2, -4))
	//输出：
	//分数为： -1/2 执行后的分数为： -1/2

	fmt.Println()
	BRat_test1 = big.NewRat(47, 3)
	// FloatString以十进制形式返回x的字符串表示形式，其prec是在小数点后的精度数位。 最后一位四舍五入到最接近的位数，一半从零舍入。
	fmt.Println("分数为：", BRat_test1, "执行后的分数字符串表示形式：", BRat_test1.FloatString(3))
	fmt.Println("分数为：", BRat_test1, "执行后的分数字符串表示形式：", BRat_test1.FloatString(2))
	fmt.Println("分数为：", BRat_test1, "执行后的分数字符串表示形式：", BRat_test1.FloatString(1))
	//输出：
	//分数为： 47/3 执行后的分数字符串表示形式： 15.667
	//如果b！= 1，RatString返回x的字符串表示形式，形式为“ a / b”，如果b == 1，则形式形式为“ a”。
	fmt.Println("分数为：", BRat_test1, "执行后的分数字符串表示形式：", BRat_test1.RatString())
	BRat_test1 := big.NewRat(2, 5)
	fmt.Println("分数为：", BRat_test1, "执行后的分数字符串表示形式：", BRat_test1.RatString())
	//输出：
	//分数为： 47/3 执行后的分数字符串表示形式： 47/3
	fmt.Println("分数的分母是否为1？：", BRat_test1.IsInt())
	BRat_test1 = big.NewRat(2, 1)
	fmt.Println("分数的分母是否为1？：", BRat_test1.IsInt())
	fmt.Println()
	//有关于分数的方法特别的就这么多，其他的方法和big.int和big.float是一样的，不再累叙！

	fmt.Println("----------big.ParseFloat()-----------")
	// ParseFloat类似于f.Parse（s，base），其中f设置为给定的精度和舍入模式。就是将字符串转为一定进制和精度的浮点数
	//	ToNearestEven RoundingMode = iota // == IEEE 754-2008 roundTiesToEven
	BRat_test1 = big.NewRat(2, 5)
	//	ToZero                            // == IEEE 754-2008 roundTowardZero
	//	AwayFromZero                      // no IEEE 754-2008 equivalent
	fmt.Println("分数为：", BRat_test1, "分数的分母为：", BRat_test1.Denom())
	BRat_test1 = big.NewRat(2, -4)
	fmt.Println("分数为：", BRat_test1, "分数的分母为：", BRat_test1.Denom())
	BRat_test1 = big.NewRat(2, -5)
	fmt.Println("分数为：", BRat_test1, "分数的分母为：", BRat_test1.Denom())
	BRat_test1 = big.NewRat(2, -1)
	fmt.Println("分数为：", BRat_test1, "分数的分母为：", BRat_test1.Denom())
		check_err_math_big(err)
		fmt.Printf("(%v)%v:%v---%v\n", mode, s, f, b)
	}

	Parse_Float("3.56666", 10, 3, big.ToNearestEven)
	Parse_Float("3.56666", 10, 3, big.ToNearestAway)
	Parse_Float("3.56666", 10, 3, big.ToNegativeInf)
	BRat_test1 = big.NewRat(2, 5)
	BRat_test2 := big.NewRat(2, 1)
	//(ToNearestEven)3.56666:3.5---10
	fmt.Println("分数为：", BRat_test1, "分数的分母为：", BRat_test1.Inv(BRat_test2))
	//(ToZero)3.56666:3.5---10
	//(AwayFromZero)3.56666:4---10
	//(ToNegativeInf)3.56666:3.5---10
	//(ToPositiveInf)3.56666:4---10

	BRat_test1 = big.NewRat(2, 5)
	// Jacobi返回+ 1，-1或0的Jacobi符号（x / y）。
	// y参数必须是一个奇整数。
	//具体请看：https://baike.baidu.com/item/%E9%9B%85%E5%8F%AF%E6%AF%94%E7%9F%A9%E9%98%B5/10753754?fr=aladdin
	fmt.Println("分数为：", BRat_test1, "分数的分母为：", BRat_test1.Num())
	BRat_test1 = big.NewRat(2, -5)
	fmt.Println("分数为：", BRat_test1, "分数的分母为：", BRat_test1.Num())
	BRat_test1 = big.NewRat(2, -1)
	fmt.Println("分数为：", BRat_test1, "分数的分母为：", BRat_test1.Num())
	BRat_test1 = big.NewRat(2, -4)
	fmt.Println("分数为：", BRat_test1, "分数的分母为：", BRat_test1.Num())
	//newInt =big.Jacobi(big.NewInt(2),big.NewInt(-6))
	//fmt.Println(newInt)//panic: big: invalid 2nd argument to Int.Jacobi: need odd integer but got -6
	//输出：
	//-1
	//0
	//-1
	//目前不大清楚这个数学
	BRat_test1 = big.NewRat(2, 5)
}
	fmt.Println("分数为：", BRat_test1, "执行后的分数为：", BRat_test1.SetFrac(big.NewInt(2), big.NewInt(-4)))
	fmt.Printf("%T--%v\n", a, a)
}

func check_err_math_big(err error) {
	}
	BRat_test1 = big.NewRat(2, 5)
	fmt.Println("分数为：", BRat_test1, "执行后的分数为：", BRat_test1.SetFrac64(2, -4))
	BRat_test1 = big.NewRat(47, 3)
	fmt.Println("分数为：", BRat_test1, "执行后的分数字符串表示形式：", BRat_test1.FloatString(3))
	fmt.Println("分数为：", BRat_test1, "执行后的分数字符串表示形式：", BRat_test1.FloatString(2))
	fmt.Println("分数为：", BRat_test1, "执行后的分数字符串表示形式：", BRat_test1.FloatString(1))
	BRat_test1 = big.NewRat(47, 3)
	fmt.Println("分数为：", BRat_test1, "执行后的分数字符串表示形式：", BRat_test1.RatString())
	BRat_test1 = big.NewRat(48, 3)
	fmt.Println("分数为：", BRat_test1, "执行后的分数字符串表示形式：", BRat_test1.RatString())
	Parse_Float := func(s string, base int, prec uint, mode big.RoundingMode) {
		f, b, err := big.ParseFloat(s, base, prec, mode)
		fmt.Printf("(%v)%v:%v---%v\n", mode, s, f, b)
	newInt := big.Jacobi(big.NewInt(2), big.NewInt(5))
	newInt = big.Jacobi(big.NewInt(15), big.NewInt(3))
	newInt = big.Jacobi(big.NewInt(2), big.NewInt(-5))
func FormatP(a interface{}) {
	fmt.Printf("%T--%v\n", a, a)
func check_err_math_big(err error) {
