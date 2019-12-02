package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"text/tabwriter"
)

func main() {
	//包rand实现伪随机数生成器。
	//随机数由源生成。 顶级函数（例如Float64和Int）使用默认的共享源，该源在每次运行程序时都会产生确定的值序列。 如果每次运行需要不同的行为，请使用种子函数初始化默认的源。
	//默认Source可安全地供多个goroutine并发使用，但不是由NewSource创建的Source。
	//此包的整个文档中都使用数学间隔符号，例如[0，n）。
	//有关适合于对安全敏感的工作的随机数，请参见crypto / rand软件包。

	fmt.Println("----------rand.Intn（）-----------")
	// Seed使用提供的种子值将默认Source初始化为确定性状态。 如果未调用Seed，则生成器的行为就像由Seed（1）设置种子一样。
	// 除以2 ^ 31-1时具有相同余数的种子值会生成相同的伪随机序列。
	// Seed与Rand.Seed方法不同，可以安全地并发使用。

	//rand.Seed(42) //尝试更改此数字！
	//rand.Seed(142) //尝试更改此数字！
	//rand.Seed(1) //尝试更改此数字！
	rand.Seed(0) //尝试更改此数字！
	answers := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}
	// Intn从默认Source返回[0，n）中的非负伪随机数作为int。
	//如果n <= 0，则表示恐慌。

	fmt.Println("Magic 8-Ball says:", answers[rand.Intn(len(answers))])
	fmt.Println("Magic 8-Ball says:", answers[rand.Intn(len(answers))])
	fmt.Println("Magic 8-Ball says:", answers[rand.Intn(len(answers))])
	fmt.Println("Magic 8-Ball says:", answers[rand.Intn(len(answers))])
	//42输出：
	//Magic 8-Ball says: As I see it yes
	//Magic 8-Ball says: Outlook good
	//Magic 8-Ball says: Yes
	//Magic 8-Ball says: Reply hazy try again

	//142输出：
	//Magic 8-Ball says: My reply is no
	//Magic 8-Ball says: Signs point to yes
	//Magic 8-Ball says: Reply hazy try again
	//Magic 8-Ball says: It is certain

	//1输出：
	//Magic 8-Ball says: It is decidedly so
	//Magic 8-Ball says: Outlook good
	//Magic 8-Ball says: Outlook good
	//Magic 8-Ball says: Very doubtful

	//0输出：
	//Magic 8-Ball says: Concentrate and ask again
	//Magic 8-Ball says: Concentrate and ask again
	//Magic 8-Ball says: Cannot predict now
	//Magic 8-Ball says: Most likely

	//不指定seek时候（默认为1）输出：
	//Magic 8-Ball says: It is decidedly so
	//Magic 8-Ball says: Outlook good
	//Magic 8-Ball says: Outlook good
	//Magic 8-Ball says: Very doubtful

	fmt.Println()
	fmt.Println("----------rand.Int（）-----------")
	//种子已经在上面给了！输出的值给种子有关，下同，上同，
	// Int从默认Source返回一个非负的伪随机整数。
	A1 := rand.Int()
	fmt.Println(A1)

	A2 := rand.Int()
	fmt.Println(A2)
	//输出：
	//3390393562759376202
	//2669985732393126063（19位数）

	fmt.Println()
	fmt.Println("----------rand.Int31（）-----------")
	// Int31从默认Source返回一个非负的伪随机31位整数作为int32。
	A3 := rand.Int31()
	fmt.Println(A3)

	A4 := rand.Int31()
	fmt.Println(A4)
	//输出：
	//413258767
	//1407315077

	fmt.Println()
	fmt.Println("----------rand.Int31（）-----------")
	// Int63从默认Source返回一个非负的伪随机63位整数作为int64。
	A5 := rand.Int63()
	fmt.Println(A5)

	A6 := rand.Int63()
	fmt.Println(A6)
	//输出：
	//8274930044578894929
	//1543572285742637646（19位数，跟rand.int方法差不多，只是返回的类型不一样）

	fmt.Println()
	fmt.Println("----------rand.Uint32()-----------")
	// Uint32从默认Source返回一个伪随机的32位值作为uint32。
	A7 := rand.Uint32()
	fmt.Println(A7)

	A8 := rand.Uint32()
	fmt.Println(A8)
	//输出：
	//1239465936
	//3876658295

	fmt.Println()
	fmt.Println("----------rand.Uint64()-----------")
	// Uint32从默认Source返回一个伪随机的32位值作为uint32。
	A9 := rand.Uint64()
	fmt.Println(A9)

	A0 := rand.Uint64()
	fmt.Println(A0)
	//输出：
	//7837839688282259259
	//2518412263346885298（19位数，跟rand.int方法差不多，只是返回的类型不一样）

	fmt.Println()
	fmt.Println("----------rand.Int31n()-----------")

	// Int31n从默认Source返回[0，n）中的非负伪随机数作为int32。
	// 有n跟没n的区别就在于返回的值的范围限不限制最大索引值
	//如果n <= 0，则表示恐慌。
	B0 := rand.Int31n(50)
	fmt.Println(B0)

	B1 := rand.Int31n(50)
	fmt.Println(B1)

	B0 = rand.Int31n(20)
	fmt.Println(B0)

	B1 = rand.Int31n(20)
	fmt.Println(B1)
	//输出：
	//2
	//26
	//11
	//10

	fmt.Println()
	fmt.Println("----------rand.Float32（）-----------")

	// Float64从默认Source返回[0.0,1.0）中的伪随机数作为float64。
	f1 := rand.Float32()
	fmt.Println(f1)

	f1 = rand.Float32()
	fmt.Println(f1)

	f1 = rand.Float32()
	fmt.Println(f1)

	f1 = rand.Float32()
	fmt.Println(f1)
	//输出：
	//0.43753722
	//0.104014836
	//0.3159685
	//0.1512936

	fmt.Println()
	fmt.Println("----------rand.Float64()-----------")

	// Float32从默认Source返回一个[0.0,1.0）中的伪随机数作为float32。
	f1_64 := rand.Float64()
	fmt.Println(f1_64)

	f1_64 = rand.Float64()
	fmt.Println(f1_64)

	f1_64 = rand.Float64()
	fmt.Println(f1_64)

	f1_64 = rand.Float64()
	fmt.Println(f1_64)
	//输出：
	//0.5102423328818813
	//0.24043190328608438
	//0.2092018731282357
	//0.6930700440076261

	fmt.Println()
	fmt.Println("----------rand.Int31n()-----------")

	// Int63n从默认Source返回[0，n）中的非负伪随机数作为int64。
	//如果n <= 0，则表示恐慌。
	B2 := rand.Int63n(50)
	fmt.Println(B2)

	B3 := rand.Int63n(50)
	fmt.Println(B3)

	B2 = rand.Int63n(20)
	fmt.Println(B0)

	B3 = rand.Int63n(20)
	fmt.Println(B1)
	//输出：
	//39
	//48
	//11
	//10

	fmt.Println()
	fmt.Println("----------rand.Int31n()-----------")
	// Intn从默认Source返回[0，n）中的非负伪随机数作为int。返回的值不会大于等于n
	//如果n <= 0，则表示恐慌。
	B4 := rand.Intn(50)
	fmt.Println(B4)

	B5 := rand.Intn(50)
	fmt.Println(B5)

	B4 = rand.Intn(20)
	fmt.Println(B4)

	B5 = rand.Intn(20)
	fmt.Println(B5)
	//输出：
	//16
	//30
	//1
	//6

	fmt.Println()
	fmt.Println("----------rand.Read()-----------")
	//上面是一次性读取一个随机数，而下面则是一次性读取多个字节返回回去
	// Read从默认Source生成len（p）个随机字节，并将其写入p中。 它总是返回len（p）和nil错误。
	//读取与Rand.Read方法不同，可以安全地并发使用。

	ls := make([]byte, 10)
	n0, e := rand.Read(ls)
	check_err_math_rand(e)
	fmt.Println("写入到字节切片中的随机数的数目为：", n0)
	fmt.Println("写入到字节切片中的随机数切片为：", ls)

	ls = make([]byte, 10)
	n0, e = rand.Read(ls)
	check_err_math_rand(e)
	fmt.Println("写入到字节切片中的随机数的数目为：", n0)
	fmt.Println("写入到字节切片中的随机数切片为：", ls)
	//输出：
	//写入到字节切片中的随机数的数目为： 10
	//写入到字节切片中的随机数切片为： [146 94 96 123 224 99 113 111 150 221]
	//写入到字节切片中的随机数的数目为： 10
	//写入到字节切片中的随机数切片为： [205 208 29 117 4 92 63 0 15 138]

	fmt.Println()
	fmt.Println("----------rand.Perm()-----------")

	// Perm以n个int的切片形式返回默认Source中整数[0，n）的伪随机排列。返回的值不会大于等于n
	ls_int := rand.Perm(10)
	fmt.Println(ls_int)

	ls_int = rand.Perm(20)
	fmt.Println(ls_int)

	ls_int = rand.Perm(20)
	fmt.Println(ls_int)

	ls_int = rand.Perm(50)
	fmt.Println(ls_int)
	//输出：
	//[5 8 0 6 4 2 3 7 9 1]
	//[13 0 14 8 18 6 2 9 11 4 5 7 1 15 3 10 16 12 17 19]
	//[0 4 19 14 15 1 13 18 7 12 17 3 9 6 11 10 2 16 5 8]
	//[31 24 43 46 23 10 11 37 42 38 4 0 26 28 35 49 20 36 41 27 48 16 18 1 44 14 6 19 15 13 47 45 3 22 2 25 7 33 29 9 17 30 34 5 12 40 8 32 21 39]

	fmt.Println()
	fmt.Println("----------rand.Perm()-----------")

	// ExpFloat64返回范围（0，+ math.MaxFloat64]范围内的指数分布float64，指数分布的比率参数（lambda）为1，平均值为默认值Source的1/lambda（1）。
	//要生成具有不同速率参数的分布，调用者可以使用以下方法调整输出：
	//样本= ExpFloat64（）/ requiredRateParameter
	ExpF := rand.ExpFloat64()
	fmt.Println(ExpF)

	ExpF = rand.ExpFloat64()
	fmt.Println(ExpF)

	//ExpF =rand.ExpFloat64()
	//fmt.Println(ExpF)
	//输出：
	//2.339417061355766
	//0.5356457145883641
	//0.18133519160443423

	fmt.Println()
	fmt.Println("----------rand.Perm()-----------")

	//使用默认Source随机伪随机化元素的顺序。
	// n是元素数。 如果n <0，则对Shuffle进行洗牌。
	// swap交换具有索引i和j的元素。
	//不应使用不适合32位的n来调用Shuffle。 这不仅会花费很长时间，而且2³¹！ 可能的排列，任何PRNG都不可能拥有足够大的内部状态以生成很小百分比的可能排列。
	//不过，正确的API签名接受一个int n，因此请尽我们最大可能进行处理。
	ExampleShuffle := func() {
		//字段根据unicode.IsSpace的定义，将字符串s围绕一个或多个连续的空白字符的每个实例进行拆分，
		// 返回s的子字符串切片；如果s仅包含空白，则返回空切片。
		words := strings.Fields("ink runs from the corners of my mouth")
		rand.Shuffle(len(words), func(i, j int) {
			words[i], words[j] = words[j], words[i]
		})
		fmt.Println(words)
		// Output:
		// [mouth my the of runs corners from ink]
	}
	ExampleShuffle()

	ExampleShuffle_slicesInUnison := func() {
		numbers := []byte("12345")
		letters := []byte("ABCDE")
		//随机排列数字，同时交换字母中的相应条目。
		rand.Shuffle(len(numbers), func(i, j int) {
			numbers[i], numbers[j] = numbers[j], numbers[i]
			letters[i], letters[j] = letters[j], letters[i]
		})
		for i := range numbers {
			fmt.Printf("%c: %c\n", letters[i], numbers[i])
		}

	}
	ExampleShuffle_slicesInUnison()
	//输出：
	//[my the ink mouth from of corners runs]
	//C: 3
	//E: 5
	//B: 2
	//D: 4
	//A: 1
	//所有的随机（或者说i,j）都受到种子seek的影响，

	fmt.Println()
	fmt.Println("----------rand.New()-----------")
	//此示例显示了* Rand上每种方法的用法。
	//全局函数的用法相同，但没有接收者。
	Example_rand := func() {

		//创建并生成种子。
		//通常应使用非固定种子，例如time.Now（）。UnixNano（）。
		//使用固定种子将在每次运行中产生相同的输出。

		// New返回一个新的Rand，它使用src中的随机值生成其他随机值。
		// NewSource返回一个以给定值作为种子的新伪随机源。
		//与顶层函数使用的默认Source不同，此Source对于多个goroutine并发使用是不安全的。
		r := rand.New(rand.NewSource(99))

		//这里的tabwriter可以帮助我们生成对齐的输出。
		// NewWriter分配并初始化一个新的tabwriter.Writer。
		//这些参数与Init函数的参数相同。
		w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
		defer w.Flush()
		show := func(name string, v1, v2, v3 interface{}) {
			//这些例程以'f'结尾，并采用格式字符串。
			// Fprintf根据格式说明符格式化并写入w。
			//返回写入的字节数以及遇到的任何写入错误。
			fmt.Fprintf(w, "%s\t%v\t%v\t%v\n", name, v1, v2, v3)
		}

		// Float32和Float64值位于[0，1）中。
		show("Float32", r.Float32(), r.Float32(), r.Float32())
		show("Float64", r.Float64(), r.Float64(), r.Float64())

		// ExpFloat64值的平均值为1，但呈指数衰减。
		show("ExpFloat64", r.ExpFloat64(), r.ExpFloat64(), r.ExpFloat64())

		// NormFloat64值的平均值为0，标准偏差为1。
		// NormFloat64返回-math.MaxFloat64到+ math.MaxFloat64（含）范围内的正态分布float64，具有标准正态分布（均值= 0，stddev = 1）。
		//要产生不同的正态分布，调用者可以使用以下方法调整输出：
		// sample = NormFloat64（）* wantedStdDev + wantedMean
		show("NormFloat64", r.NormFloat64(), r.NormFloat64(), r.NormFloat64())

		// Int31，Int63和Uint32生成给定宽度的值。
		//根据'int'的大小，Int方法（未显示）类似于Int31或Int63。至于是哪个要看系统位数！
		show("Int31", r.Int31(), r.Int31(), r.Int31())
		show("Int63", r.Int63(), r.Int63(), r.Int63())
		show("Uint32", r.Uint32(), r.Uint32(), r.Uint32())

		// Intn，Int31n和Int63n将它们的输出限制为<n。
		//他们比使用r.Int（）％n更谨慎。
		show("Intn(10)", r.Intn(10), r.Intn(10), r.Intn(10))
		show("Int31n(10)", r.Int31n(10), r.Int31n(10), r.Int31n(10))
		show("Int63n(10)", r.Int63n(10), r.Int63n(10), r.Int63n(10))

		// Perm生成数字[0，n）的随机排列。
		show("Perm", r.Perm(5), r.Perm(5), r.Perm(5))

	}
	Example_rand()
	//输出：
	//----------rand.New()-----------
	//Float32     0.2635776           0.6358173           0.6718283
	//Float64     0.628605430454327   0.4504798828572669  0.9562755949377957
	//ExpFloat64  0.3362240648200941  1.4256072328483647  0.24354758816173044
	//NormFloat64 0.17233959114940064 1.577014951434847   0.04259129641113857
	//Int31       1501292890          1486668269          182840835
	//Int63       3546343826724305832 5724354148158589552 5239846799706671610
	//Uint32      2760229429          296659907           1922395059
	//Intn(10)    1                   2                   5
	//Int31n(10)  4                   7                   8
	//Int63n(10)  7                   6                   3
	//Perm        [1 4 2 3 0]         [4 2 1 3 0]         [1 2 4 0 3]

	fmt.Println()
	fmt.Println("----------rand.New()1111-----------")
	//rand实现了Source接口
	source := rand.NewSource(100)
	// Source表示在[0，1 << 63）范围内均匀分布的伪随机int64值的源。
	//如果我这里给的不是100，则下面的所有得到输出都将会发生改变，Seed表示可以更改种子，也就是不使用100了，更换为其他的种子值
	//source.Seed(200)
	// Int63返回序列中的一个非负的伪随机63位整数作为int64。
	//一旦这里取值的话下面的所有结果将会发生改变
	fmt.Println(source.Int63())

	r1 := rand.New(source)
	//fmt.Println(r1.Intn(20))
	//fmt.Println(r1.Intn(20))
	//fmt.Println(r1.Intn(20))
	//fmt.Println(r1.Intn(20))
	//fmt.Println(r1.Intn(20))

	fmt.Println(r1.Intn(10))
	fmt.Println(r1.Intn(10))
	fmt.Println(r1.Intn(10))
	fmt.Println(r1.Intn(10))
	fmt.Println(r1.Intn(10))
	//输出：
	//3
	//8
	//0
	//0
	//2
	//将10注释，将20放开输出结果一样的！说明这个序列已经是产生了的

	//ls10:=make([]byte,10)
	ls20 := make([]byte, 10)
	//fmt.Println(r1.Read(ls10))
	//fmt.Println(ls10)

	fmt.Println(r1.Read(ls20))
	fmt.Println(ls20)

	//输出：
	//10 <nil>
	//[24 175 16 25 140 55 52 156 137 163]
	//将10注释，将20放开输出结果一样的！说明这个序列已经是产生了的

	//下面是一个很无聊的数学东西，略
	//fmt.Println()
	//fmt.Println("----------rand.NewZipf()----------")
	//// NewZipf返回一个Zipf变量生成器。
	////生成器生成值k∈[0，imax]，使得P（k）与（v + k）**（-s）成比例。
	////要求：s> 1和v> = 1。
	//rand.NewZipf()

	//关于math模块就讲这么多,有一个模块没讲，复数的模块，不建议花费时间！
}

func check_err_math_rand(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
