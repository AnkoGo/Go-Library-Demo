package main

import (
	"fmt"
	"net"
)

//这个文件作为udp的服务器端来使用
func main() {
	la, err := net.ResolveUDPAddr("udp4", "127.0.0.1:8888") //这个端口千万年不要和上部分测试的tcp使用的端口9999相同，否则会报错
	check_err_net(err)
	//net.ListenTCP()返回值是一个tcp监听器*TCPListener对象，但是为什么下面的net.ListenUDP（）确实返回一个*UDPConn呢？
	//因为tcp才需要建立连接，而udp是不需要建立连接的额！所以不存在监听器，直接创建一个*UDPConn对象直接读取和写入即可！
	//*UDPConn对象的更多方法将会在客户端中进行探索而不是在这里的服务端进行探索！请到客户端在去查看！
	udpListener, err := net.ListenUDP("udp4", la) //监听9999端口，从这个端口接收发送端发送过来的数据！
	check_err_net(err)
	fmt.Println("tcpListener:", udpListener)

	defer udpListener.Close()

	//直接进行读取即可，不存在监听的过程！这就是udp有别于tcp的地方！这里会阻塞到一直接收到数据为止！
	read_by := make([]byte, 100)
	i, err := udpListener.Read(read_by)
	check_err_net(err)

	fmt.Println("udp服务器接收到客户端的字节数为：", i)
	fmt.Println("udp服务器接收到客户端的字节数据为：", read_by)
	fmt.Println("udp服务器接收到客户端的字节转字符串数据为：", string(read_by))
	fmt.Println("udp服务器接收完毕！！")

	fmt.Println()
	//write_i, err := udpListener.Write([]byte("这是udp服务器发送给客户端的数据2222"))
	//注意在tcp是长连接所以不会指定对方的地址都是可以直接发送数据的，只要客户端请求进来，那么服务器发送数据时候是不用再指定
	//对方的地址的了，但是udp是短连接，他虽然也是全双工通信，但是还是需要指定对方的地址才可以进行发送数据回头的！
	//事实上这个write()方法是UDPConn组合下的conn对象上面的方法，这是go库 api设计需要改进的地方，既然有这个方法，但是却不能被调用，
	//这是非常不人道的行为！切记我们在设计api时候不要这样做，如果他们的行为不完全一样，那么就不要生搬硬套，而是因为抽象出共同的行为来新建
	// 一个接口或者结构体，最后组合这个结构体或者接口即可！
	//假如采用上面的方式的话则会抛出这个错误：
	//	出错了，错误信息为： write udp4 127.0.0.1:8888: wsasend: A request to send or receive data was disallowed because the socket is not connected and (when sending on a datagram socket using a sendto call) no address was supplied.
	//	panic: write udp4 127.0.0.1:8888: wsasend: A request to send or receive data was disallowed because the socket is not connected and (when sending on a datagram socket using a sendto call) no address was supplied.
	//
	//	goroutine 1 [running]:
	//	main.check_err_net(0x524880, 0xc0000a4000)
	//		C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/ListenUdp.go:56 +0xdd
	//	main.main()
	//		C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/ListenUdp.go:33 +0x53c
	//错误信息的意思是：
	//	不允许发送或接收数据的请求，因为未连接套接字，并且（当使用sendto调用在数据报套接字上发送时）未提供地址。

	//这里我们不能写回这个8888端口，这个端口一定要是客户端中监听的7777端口！一个端口在udp中不能享有读和写功能，但是在tcp中是可以的！
	//然而udp中要分开来！
	//
	ra, err := net.ResolveUDPAddr("udp4", "127.0.0.1:7777")
	check_err_net(err)

	write_i, err := udpListener.WriteTo([]byte("这是udp服务器发送给客户端的数据2222"), ra)
	check_err_net(err)
	fmt.Println("udp服务器响应给客户端的字节个数为：", write_i)
	if write_i == 0 {
		fmt.Println("未发送任何数据给客户端，发生了错误！请检查！")
	}
	fmt.Println("udp服务器响应给客户端的字节数据为：", "这是udp服务器发送给客户端的数据2222")
	fmt.Println("udp服务器响应完毕！！")

	//此服务器端中的输出为：（注意先要运行服务器代码，然后才能运行客户端的代码）
	//	tcpListener: &{{0xc00007ca00}}
	//	udp服务器接收到客户端的字节数为： 60
	//	udp服务器接收到客户端的字节数据为： [232 191 153 230 152 175 229 174 162 230 136 183 231 171 175 229 143 145 233 128 129 231 187 153 230
	//	156 141 229 138 161 229 153 168 231 171 175 231 154 132 228 191 161 230 129 175 49 49 49 117 100 112 67 111 110 110 32 116 101 115 116
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	//	udp服务器接收到客户端的字节转字符串数据为： 这是客户端发送给服务器端的信息111udpConn test
	//	udp服务器接收完毕！！
	//
	//	udp服务器响应给客户端的字节个数为： 49
	//	udp服务器响应给客户端的字节数据为： 这是udp服务器发送给客户端的数据2222
	//	udp服务器响应完毕！！

}
func check_err_net(err error) {
	if err != nil {
		//fmt.Fprintln(os.Stderr,err)
		//上面的这种方式会导致输出顺序不确定，虽然他可以输出红色的字体，但是由于顺序不确定，我们不采用他！
		fmt.Println("出错了，错误信息为：", err)
		panic(err)
	}
}
