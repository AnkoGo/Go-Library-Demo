package main

import (
	"fmt"
	"net"
	//"time"
)

func main() {
	la, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:9999")
	check_err_net(err)
	tcpListener, err := net.ListenTCP("tcp4", la) //监听9999端口，从这个端口接收发送端发送过来的数据！
	check_err_net(err)
	fmt.Println("tcpListener:", tcpListener)

	//SetDeadline设置与侦听器关联的截止日期。
	//零时间值将禁用截止日期。
	//这个我们不设置的话就是默认的无限期！
	//他的底层是调用poll.FD.SetDeadline（）方法，这个方法的文档如下：
	//SetDeadline设置与fd相关的读取和写入截止时间。
	//fmt.Println("tcpListener:",tcpListener.SetDeadline())

	defer tcpListener.Close()
	Addr := tcpListener.Addr()

	fmt.Println("Addr:", Addr)
	fmt.Println("Addr.String():", Addr.String())
	fmt.Println("Addr.Network():", Addr.Network())

	fmt.Println("准备进行沉睡44s。。。")
	//在这里设置的延迟是为了故意让客户端连接服务器时候超时连接，在你需要测试客户端中的
	//tcpConn.SetKeepAlivePeriod(3 * time.Second)时候可以把下面的这行代码放开注释！
	//不过在你放开这个注释时候，最好把下面的time.Sleep(4e9)注释掉！如果已经注释则无需理会！
	//我先注释掉这个，因为我不想每次测试都等待这么久，可是事实上这行应该是要打开的！
	//time.Sleep(44e9)
	fmt.Println("准备进行苏醒。。。")
	// AcceptTCP accepts the next incoming call and returns the new
	// connection.
	// AcceptTCP接受下一个来电并返回新的连接。如果没有连接进来的话，那么这个方法会阻塞直到有连接连进来！
	//事实上这里我们正常的应该使用并发去处理tcpConn，每次创建一个连接后都应该用一个g程去操作read()和write().并且应该用一个循环来包住这个tcpListener.AcceptTCP()
	//而且tcpListener一般都不会关闭，除非整个文件代码执行完毕，但是这个是很少出现的！应该只要你的代码一开始运行，结束代码则表示你的 整个服务器都暂停了！也就是无法正常的让客户访问了！
	//我们会在介绍完整个net包后写一个完整的tcp或者http的服务器代码！但是目前来说，我们还是主要以讲解api用法和api结构体为主！
	tcpConn, err := tcpListener.AcceptTCP()
	check_err_net(err)

	defer tcpConn.Close()
	////其实这个Close相当于下面的 2句代码了：
	//// CloseRead shuts down the reading side of the TCP connection.
	//// Most callers should just use Close.
	////CloseRead关闭TCP连接的读取端。
	////大多数呼叫者应该只使用Close。
	//tcpConn.CloseRead()
	//
	//// CloseWrite shuts down the writing side of the TCP connection.
	//// Most callers should just use Close.
	////CloseWrite关闭TCP连接的写入端。
	////大多数呼叫者应该只使用Close。
	//tcpConn.CloseWrite()

	////这句话完全是为了测试连接超时，在你需要测试时候可以放开这行注释！
	////因为我们在客户端中通过e=tcpConn.SetDeadline(time.Now().Add(someTimeout))或者tcpConn.SetReadDeadline(time.Now().Add(someTimeout))
	////或者tcpConn.SetWriteDeadline(time.Now().Add(someTimeout))设置了3s的读写时间 ，而我们这里服
	//// 务器设置了延迟4s才发送响应给客户端，所以我们这里一定会抛出错误的！
	////需要注意的是，我们这里延迟是设置了连接之后的读写延迟，而不是像上面的连接延迟（上面也有一个time.Sleep(4e9)注意了）！他们是不一样的！
	//time.Sleep(4e9)

	fmt.Println("tcpConn.LocalAddr()", tcpConn.LocalAddr())
	fmt.Println("tcpConn.RemoteAddr()", tcpConn.RemoteAddr())
	// SyscallConn returns a raw network connection.
	// This implements the syscall.Conn interface.
	//SyscallConn返回原始网络连接。
	//这实现了syscall.Conn接口。

	//// A RawConn is a raw network connection.// RawConn是原始网络连接。
	//type RawConn interface {
	//	// Control invokes f on the underlying connection's file
	//	// descriptor or handle.
	//	// The file descriptor fd is guaranteed to remain valid while
	//	// f executes but not after f returns.
	//	//Control(控件)在基础连接的文件描述符或句柄上调用f。
	//	//确保文件描述符fd在f执行期间保持有效，但在f返回之后无效。
	//	Control(f func(fd uintptr)) error
	//
	//	// Read invokes f on the underlying connection's file
	//	// descriptor or handle; f is expected to try to read from the
	//	// file descriptor.
	//	// If f returns true, Read returns. Otherwise Read blocks
	//	// waiting for the connection to be ready for reading and
	//	// tries again repeatedly.
	//	// The file descriptor is guaranteed to remain valid while f
	//	// executes but not after f returns.
	//	//读取调用基础连接的文件描述符或句柄上的f； f应该尝试从文件描述符中读取。
	//	//如果f返回true，则Read返回。 否则，“读取”将阻止等待连接准备好进行读取，然后再次尝试。
	//	//确保文件描述符在f执行期间保持有效，但在f返回之后无效。
	//	Read(f func(fd uintptr) (done bool)) error
	//
	//	// Write is like Read but for writing.//Write就像Read，但是是写操作。
	//	Write(f func(fd uintptr) (done bool)) error
	//}

	rawConn, err := tcpConn.SyscallConn()
	check_err_net(err)
	fmt.Println("rawConn:", rawConn)

	Read_by := make([]byte, 100)
	i, err := tcpConn.Read(Read_by)
	check_err_net(err)

	fmt.Println("tcp服务器读取来自客户端的字节数为：", i)
	fmt.Println("tcp服务器读取来自客户端的字节数据为：", Read_by)
	fmt.Println("tcp服务器读取来自客户端的字节转字符串数据为：", string(Read_by))
	fmt.Println("服务器接收到了客户端的数据完毕！！")

	write_i, err := tcpConn.Write([]byte("这是tcp服务器响应给客户端的内容2222"))

	fmt.Println()
	fmt.Println("tcp服务器响应给客户端的字节数为：", write_i)
	fmt.Println("tcp服务器响应给客户端的字节转字符串数据为：", "这是tcp服务器响应给客户端的内容2222")
	fmt.Println("服务器响应完毕！！")

	//// SetReadBuffer sets the size of the operating system's
	//// receive buffer associated with the connection.
	////SetReadBuffer设置与连接关联的操作系统的 接收缓冲区 的大小。
	//tcpConn.SetReadBuffer()
	//
	//// SetWriteBuffer sets the size of the operating system's
	//// transmit buffer associated with the connection.
	////SetWriteBuffer设置与连接关联的操作系统的 传输缓冲区 的大小。
	//tcpConn.SetWriteBuffer()

	//结果1
	//本文件作为服务端输出：（如果你想要查看客户端的代码，请到compress_zlib.go文件中去搜索“net.DialTCP”就能看到了，或者搜索“ListenTcp.go”更加准确！）
	//	tcpListener: &{0xc00007ca00 {<nil> 0}}
	//	Addr: 127.0.0.1:9999
	//	Addr.String(): 127.0.0.1:9999
	//	Addr.Network(): tcp
	//	tcpConn.LocalAddr() 127.0.0.1:9999，这个地址是服务端的主机和端口，当然你我们这里就是本机！
	//	tcpConn.RemoteAddr() 127.0.0.1:52169，这个地址是客户端的主机和端口，当然我们这里也是采用本机了！你可以自己建立一个虚拟机来进行尝试！
	//	rawConn: &{0xc00007cc80}
	//	tcp服务器读取来自客户端的字节数为： 64
	//	tcp服务器读取来自客户端的字节数据为： [232 191 153 230 152 175 229 174 162 230 136 183 231 171 175 229 143 145 233 128
	//	129 231 187 153 230 156 141 229 138 161 231 171 175 231 154 132 228 191 161 230 129 175 230 149 176 230 141 174 49
	//	49 49 49 84 67 80 67 79 78 78 32 84 69 83 84 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	//	tcp服务器读取来自客户端的字节转字符串数据为： 这是客户端发送给服务端的信息数据1111TCPCONN TEST
	//	服务器接收到了客户端的数据完毕！！
	//
	//	tcp服务器响应给客户端的字节数为： 49
	//	tcp服务器响应给客户端的字节转字符串数据为： 这是tcp服务器响应给客户端的内容2222
	//	服务器响应完毕！！

	//结果2
	//这是第二个结果的输出：请看客户端中的额第二个结果的注释即可知道这里输出什么了！其实 跟上面的输出结果是一样的！只是客户端和服务器端都会同时延迟44s的时间！

	//结果3
	//如果是放开上面//time.Sleep(44e9)这行代码的注释的话，则输出结果如下：
	//客户端的操作请看客户端中的注释，他是对应的，比如上面的第一个结果对应客户端的第一个输出结果，这里的第3个结果则对应客户端中的第3个结果！
	//	tcpListener: &{0xc00007ca00 {<nil> 0}}
	//	Addr: 127.0.0.1:9999
	//	Addr.String(): 127.0.0.1:9999
	//	Addr.Network(): tcp
	//	事实上本文件没执行完成，因为即使到了44s过后的时间，也不会有新的连接到来，如果是真正上线的代码则会有新的客户端请求连接到来
	//	所以，我们这里其实是自己结束了本文件程序的运行！

	//延时4s的话则会输出（也就是放开上面//time.Sleep(4e9)这行代码的注释即可）：（如果你想要查看客户端的代码，请到compress_zlib.go文件中去搜索“net.DialTCP”就能看到了，或者搜索“ListenTcp.go”更加准确！）
	//	tcpListener: &{0xc00007ca00 {<nil> 0}}
	//	Addr: 127.0.0.1:9999
	//	Addr.String(): 127.0.0.1:9999
	//	Addr.Network(): tcp
	//	tcpConn.LocalAddr() 127.0.0.1:9999
	//	tcpConn.RemoteAddr() 127.0.0.1:52275，这个端口是不固定的，注意了
	//	rawConn: &{0xc00007cc80}
	//	出错了，错误信息为： read tcp4 127.0.0.1:9999->127.0.0.1:52275: wsarecv: An existing connection was forcibly closed by the remote host.
	//	panic: read tcp4 127.0.0.1:9999->127.0.0.1:52275: wsarecv: An existing connection was forcibly closed by the remote host.
	//
	//	goroutine 1 [running]:
	//	main.check_err_net(0x5286a0, 0xc000098000)
	//		C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/ListenTcp.go:147 +0xdd
	//	main.main()
	//		C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/ListenTcp.go:96 +0x6d1

}
func check_err_net(err error) {
	if err != nil {
		//fmt.Fprintln(os.Stderr,err)
		//上面的这种方式会导致输出顺序不确定，虽然他可以输出红色的字体，但是由于顺序不确定，我们不采用他！
		fmt.Println("出错了，错误信息为：", err)
		panic(err)
	}
}
