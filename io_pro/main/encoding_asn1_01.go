/* ASN1 DaytimeServer
 */
package main

import (
	"encoding/asn1"
	"fmt"
	"net"
	"os"
	"time"
)

//服务端：等待客户端连接并且连接后发送服务器的时间
func main() {

	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError1(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)//在这里阻塞监听端口1200
	checkError1(err)

	for {
		conn, err := listener.Accept()//解阻塞后来到这里创建新的连接
		if err != nil {
			continue
		}

		//daytime := time.Now()//获取服务器当前的时间
		daytime := time.Date(1990,12,11,2,3,4,5,time.Local)//自己创建一个时间发送过去
		// Ignore return network errors.
		mdata, _ := asn1.Marshal(daytime)//按照asn1来编码服务器的时间
		// Write将数据写入连接。
		//可以使写入超时并在固定的时间限制后使用Timeout（）== true返回错误； 请参见SetDeadline和SetWriteDeadline。
		conn.Write(mdata)//将编码后的东西写进入conn里面
		conn.Close() // 关闭连接
	}
}

func checkError1(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}