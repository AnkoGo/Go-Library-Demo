package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main1689() {
	// r := strings.NewReader("Hello, Reader!")

	bytes, e := ioutil.ReadFile(`main\test.txt`) //读取整个文件，第一个参数跟打开的主文件目录有关，感觉很奇怪
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(bytes)
	fmt.Println(string(bytes))
	fmt.Println(len(bytes), cap(bytes))
	fmt.Println("------------------------------------")

	Fi, err000 := ioutil.ReadDir("main") //按字母排序返回一个os.FileInfo对象列表
	if err000 != nil {
		fmt.Println(err000)
	}
	fmt.Println(Fi)
	for _, v := range Fi {
		fmt.Println(v.Name())
	}
	fmt.Println("------------------------------------")

	r := strings.NewReader("Hello, Reader!")
	RClose := ioutil.NopCloser(r) //完全是为了让io.Reader方便转为io.ReadCloser对象而不用通过实现io.ReadCloser来实例化这个对象
	ls := make([]byte, 10)
	RClose.Read(ls)
	fmt.Println(string(ls))
	RClose.Close() //什么都没做，所以一般不会用这个方法的
	fmt.Println("------------------------------------")

	err222 := ioutil.WriteFile("main/test3.txt", []byte{99, 100, 101, 102, 103, 104, 105, 106}, 0777) //直接将字节写入到文件中去，需要指定文件权限名
	if err222 != nil {
		fmt.Println(err222)
	}

	// fmt.Println("------------------------------------")

	// newDir, err333 := ioutil.TempDir("main", "temp_") //在main里面创建一个以temp_前缀为开始的随机名字的文件夹,
	// // 即使之前已经创建过一次也再次创建，因为每次创建的文件目录名字是不一样的
	// fmt.Println(newDir)
	// fmt.Println(err333)

	// fmt.Println("------------------------------------")

	// newFile, err444 := ioutil.TempFile("main", "tempF_") //在main里面创建一个以temp_前缀为开始的随机名字的文件夹
	// fmt.Println(newFile)
	// fmt.Println(err444)

	fmt.Println("------------------------------------")

	ls111 := []byte{99, 100, 101, 102, 103, 104, 105, 106}
	n, err555 := ioutil.Discard.Write(ls111) //事实上这里什么都没写入的！不过不知道到底有什么用
	if err555 != nil {
		fmt.Println(err555)
	}
	fmt.Println(n)
	fmt.Println(ioutil.Discard) //0
	// ioutil.devNull              //这个名字是小写，所以不能外用！

}
