package main

import (
	"fmt"
	"net"
	//"time"
)

func main() {
	//la, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:9996")
	//check_err_net(err)
	//如上，如果是采用 net.Listen（）的话则不用封装特殊的tcp或者udp的addr类型地址了，直接写字符串即可，看下面：

	//// A Listener is a generic network listener for stream-oriented protocols.
	////
	//// Multiple goroutines may invoke methods on a Listener simultaneously.
	////侦听器是面向流协议的通用网络侦听器。
	////多个goroutine可以同时调用侦听器上的方法。
	//type Listener interface {
	//	// Accept waits for and returns the next connection to the listener.
	//	//接受等待，并将下一个连接返回给侦听器。
	//	Accept() (Conn, error)
	//
	//	// Close closes the listener.
	//	// Any blocked Accept operations will be unblocked and return errors.
	//	//Close关闭监听器。
	//	//任何阻塞的接受操作将被解除阻塞并返回错误。
	//	Close() error
	//
	//	// Addr returns the listener's network address.
	//	// Addr返回侦听器的网络地址。
	//	Addr() Addr
	//}

	// Listen announces on the local network address.
	//
	// The network must be "tcp", "tcp4", "tcp6", "unix" or "unixpacket".
	//
	// For TCP networks, if the host in the address parameter is empty or
	// a literal unspecified IP address, Listen listens on all available
	// unicast and anycast IP addresses of the local system.
	// To only use IPv4, use network "tcp4".
	// The address can use a host name, but this is not recommended,
	// because it will create a listener for at most one of the host's IP
	// addresses.
	// If the port in the address parameter is empty or "0", as in
	// "127.0.0.1:" or "[::1]:0", a port number is automatically chosen.
	// The Addr method of Listener can be used to discover the chosen
	// port.
	//
	// See func Dial for a description of the network and address
	// parameters.
	//	Listen在本地网络地址上宣布。
	//	网络必须是“ tcp”，“ tcp4”，“ tcp6”，“ unix”或“ unixpacket”。
	//	对于TCP网络，如果address参数中的主机为空或文字未指定的IP地址，则Listen侦听本地系统的所有可用单播和Anycast IP地址。
	//	要仅使用IPv4，请使用网络“ tcp4”。
	//	该地址可以使用主机名，但是不建议使用该名称，因为它会为主机的IP地址之一最多创建一个侦听器。
	//	如果地址参数中的端口为空或“ 0”，例如“ 127.0.0.1：”或“ [:: 1]：0”，则会自动选择端口号。
	//	侦听器的Addr方法可用于发现所选端口。
	//	有关网络和地址参数的说明，请参见func Dial。
	tcpListener, err := net.Listen("tcp4", "127.0.0.1:9996") //监听9996端口，从这个端口接收发送端发送过来的数据！
	check_err_net(err)
	fmt.Println("tcpListener:", tcpListener)

	defer tcpListener.Close()
	Addr := tcpListener.Addr()

	fmt.Println("Addr:", Addr)
	fmt.Println("Addr.String():", Addr.String())
	fmt.Println("Addr.Network():", Addr.Network())

	fmt.Println("正在等待客户端请求。。。")
	Conn, err := tcpListener.Accept()
	check_err_net(err)

	defer Conn.Close()

	fmt.Println("tcpConn.LocalAddr()", Conn.LocalAddr())
	fmt.Println("tcpConn.RemoteAddr()", Conn.RemoteAddr())

	//同样没有下面的这个方法，我们注释掉她！
	//rawConn, err := Conn.SyscallConn()
	//check_err_net(err)
	//fmt.Println("rawConn:",rawConn)

	Read_by := make([]byte, 100) //我们不知道我们要接收的数据的大小，但是我们可以先给个比较大的容量大小！
	i, err := Conn.Read(Read_by)
	check_err_net(err)

	fmt.Println("tcp服务器读取来自客户端的字节数为：", i)
	fmt.Println("tcp服务器读取来自客户端的字节数据为：", Read_by)
	fmt.Println("tcp服务器读取来自客户端的字节转字符串数据为：", string(Read_by))
	fmt.Println("服务器接收到了客户端的数据完毕！！")

	//这里不用再指明对方的ip:端口了，因为默认就是发送到那个端口和ip的，和谁在通信已经知道了
	write_i, err := Conn.Write([]byte("这是tcp服务器响应给客户端的内容2222"))

	fmt.Println()
	fmt.Println("tcp服务器响应给客户端的字节数为：", write_i)
	fmt.Println("tcp服务器响应给客户端的字节转字符串数据为：", "这是tcp服务器响应给客户端的内容2222")
	fmt.Println("服务器响应完毕！！")

	//输出：
	//	tcpListener: &{0xc00007ca00 {<nil> 0}}
	//	Addr: 127.0.0.1:9996
	//	Addr.String(): 127.0.0.1:9996
	//	Addr.Network(): tcp
	//	正在等待客户端请求。。。
	//	tcpConn.LocalAddr() 127.0.0.1:9996
	//	tcpConn.RemoteAddr() 127.0.0.1:9994
	//	tcp服务器读取来自客户端的字节数为： 42
	//	tcp服务器读取来自客户端的字节数据为： [232 191 153 230 152 175 229 174 162 230 136 183 231 171 175 229 143 145 233 128 129 231 154 132 116 99 112 230 149 176 230 141 174 227 128 130 227 128 130 227 128 130 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	//	tcp服务器读取来自客户端的字节转字符串数据为： 这是客户端发送的tcp数据。。。
	//	服务器接收到了客户端的数据完毕！！
	//
	//	tcp服务器响应给客户端的字节数为： 49
	//	tcp服务器响应给客户端的字节转字符串数据为： 这是tcp服务器响应给客户端的内容2222
	//	服务器响应完毕！！
}
func check_err_net(err error) {
	if err != nil {
		//fmt.Fprintln(os.Stderr,err)
		//上面的这种方式会导致输出顺序不确定，虽然他可以输出红色的字体，但是由于顺序不确定，我们不采用他！
		fmt.Println("出错了，错误信息为：", err)
		panic(err)
	}
}
