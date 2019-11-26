/* ASN.1 DaytimeClient
 */
package main

import (
	"bytes"
	"encoding/asn1"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)
//客户端：接收到asn1的值，解码数据并且打印出来
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])//在控制台需要用户输出连接的host和port
		os.Exit(1)
	}
	service := os.Args[1]//获取用户输入的host地址

	conn, err := net.Dial("tcp", service)//发送请求到相应的host地址
	checkError(err)

	result, err := readFully(conn)//读取服务器传递过来的ans1编码后的数据
	checkError(err)

	var newtime time.Time
	_, err1 := asn1.Unmarshal(result, &newtime)//解码ans1数据
	checkError(err1)

	fmt.Println("After marshal/unmarshal: ", newtime.String())//打印出解码后的数据

	os.Exit(0)//正常退出程序
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()//读完后关闭连接，不是必须的

	result := bytes.NewBuffer(nil)//创建一个缓存容器
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])//读取conn里面的数据到buf容器中去
		result.Write(buf[0:n])//将buf容器中的数据转存到result这个容器中去
		if err != nil {
			if err == io.EOF {//判断流传递到结尾与否
				break
			}
			return nil, err//读取出错，退出
		}
	}
	return result.Bytes(), nil//读取成功，返回相应的切片数组
}
//在控制台中输入：go run encoding_asn1_02.go localhost:1200
//输出如下：
//After marshal/unmarshal:  1990-12-11 02:03:04 +0800 CST，这数据仅仅会保留到s的精度，即使你传递更加微小的精度也不会解码出来的
