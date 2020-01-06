package main

import (
	"fmt"
	"net"
)

//这个文件作为服务端
func main() {

	//// UnixListener is a Unix domain socket listener. Clients should
	//// typically use variables of type Listener instead of assuming Unix
	//// domain sockets.
	//// UnixListener是Unix域套接字侦听器。 客户端通常应使用Listener类型的变量，而不是使用Unix域套接字。
	//type UnixListener struct {
	//	fd         *netFD
	//	path       string
	//	unlink     bool
	//	unlinkOnce sync.Once
	//}

	//// Network file descriptor.//网络文件描述符。
	//type netFD struct {
	//	pfd poll.FD
	//
	//	// immutable until Close//不可变，直到关闭
	//	family      int
	//	sotype      int
	//	isConnected bool // handshake completed or use of association with peer//握手完成或使用与对等方的关联
	//	net         string
	//	laddr       Addr
	//	raddr       Addr
	//}

	//// FD is a file descriptor. The net and os packages embed this type in
	//// a larger type representing a network connection or OS file.
	//// FD是文件描述符。 net和os软件包将这种类型嵌入更大的类型中，表示网络连接或OS文件。
	//type FD struct {
	//	// Lock sysfd and serialize access to Read and Write methods.//锁定sysfd并序列化对Read和Write方法的访问。
	//	fdmu fdMutex
	//
	//	// System file descriptor. Immutable until Close.//系统文件描述符。 直到关闭为止不变。
	//	Sysfd syscall.Handle
	//
	//	// Read operation.//读取操作。
	//	rop operation
	//	// Write operation.//写入操作。
	//	wop operation
	//
	//	// I/O poller.// I / O轮询器。
	//	pd pollDesc
	//
	//	// Used to implement pread/pwrite.//用于实现pread / pwrite。
	//	l sync.Mutex
	//
	//	// For console I/O.//用于控制I/O。
	//	lastbits       []byte   // first few bytes of the last incomplete rune in last write//最后写入中最后一个不完整符文的前几个字节
	//	readuint16     []uint16 // buffer to hold uint16s obtained with ReadConsole//缓冲区以保存通过ReadConsole获得的uint16
	//	readbyte       []byte   // buffer to hold decoding of readuint16 from utf16 to utf8//缓冲区以保存从utf16到utf8的readuint16的解码
	//	readbyteOffset int      // readbyte[readOffset:] is yet to be consumed with file.Read// readbyte [readOffset：]尚未与file.Read一起使用
	//
	//	// Semaphore signaled when file is closed.//关闭文件时发出信号量。
	//	csema uint32
	//
	//	skipSyncNotif bool
	//
	//	// Whether this is a streaming descriptor, as opposed to a
	//	// packet-based descriptor like a UDP socket.
	//	//与基于分组的描述符（如UDP套接字）相反，这是否是流描述符。
	//	IsStream bool
	//
	//	// Whether a zero byte read indicates EOF. This is false for a
	//	// message based socket connection.
	//	//读取的零字节是否表示EOF。 对于基于消息的套接字连接，这是错误的。
	//	ZeroReadIsEOF bool
	//
	//	// Whether this is a file rather than a network socket.
	//	//这是文件还是网络套接字。
	//	isFile bool
	//
	//	// The kind of this file.//此文件的种类。
	//	kind fileKind
	//}

	// ListenUnix acts like Listen for Unix networks.
	//
	// The network must be "unix" or "unixpacket".
	// ListenUnix的行为类似于侦听Unix网络。
	//
	//网络必须是“ unix”或“ unixpacket”。
	//其实这个net作为网络名字可以随意给，虽然文档上说应该给的是unix，unixgram或者unixpacket，我觉得给空都是可以的！我就不测试了！
	//而且从下面的打印出来的数据可以看出，即使你给了remotenet字符串，但是net.UnixAddr对象还是设置了第一个参数“unix”这个字符串，而没有采用第三个参数的"remotenet"。具体原因还不知道，因为我无法调试！先搁置！
	unixListener, e := net.ListenUnix("unix", &net.UnixAddr{"./raddrUnixSocket", "remotenet"})
	check_err_net(e) //这个函数在同一个包中的其他文件被定义了，但是我想要这个文件单独运行，所以我在这里在定义一次，这样的话 单独运行也不会相互影响了！

	// Close stops listening on the Unix address. Already accepted
	// connections are not closed.
	// Close停止监听Unix地址。 已接受的连接未关闭。

	//底层调用了UnixListener.close()方法，关于这个方法的文档如下：
	// The operating system doesn't clean up
	// the file that announcing created, so
	// we have to clean it up ourselves.
	// There's a race here--we can't know for
	// sure whether someone else has come along
	// and replaced our socket name already--
	// but this sequence (remove then close)
	// is at least compatible with the auto-remove
	// sequence in ListenUnix. It's only non-Go
	// programs that can mess us up.
	// Even if there are racy calls to Close, we want to unlink only for the first one.
	//操作系统不会清理宣布创建的文件，因此我们必须自己清理它。
	//这里有一场比赛-我们无法确定是否有人已经来代替我们的套接字名称-但是此序列（先删除再关闭）至少与ListenUnix中的自动删除序列兼容。 只有非Go程序会使我们搞砸。
	//即使对Close发出了积极的呼吁，我们也只希望第一个取消链接。

	//从上面可以知道这个Close()方法会先删除由于通信需要的socket文件然后再关闭socket连接的！
	defer unixListener.Close()

	// Addr returns the listener's network address.
	// The Addr returned is shared by all invocations of Addr, so
	// do not modify it.
	// Addr返回侦听器的网络地址。
	//返回的Addr由Addr的所有调用共享，因此请勿对其进行修改。

	//// Addr represents a network end point address.
	////
	//// The two methods Network and String conventionally return strings
	//// that can be passed as the arguments to Dial, but the exact form
	//// and meaning of the strings is up to the implementation.
	//// Addr表示网络端点地址。
	//// Network和String这两种方法通常会返回可作为Dial参数传递的字符串，但是字符串的确切形式和含义取决于实现方式。
	//type Addr interface {
	//	Network() string // name of the network (for example, "tcp", "udp")//网络名称（例如“ tcp”，“ udp”）
	//	String() string  // string form of address (for example, "192.0.2.1:25", "[2001:db8::1]:80")//地址的字符串形式（例如，"192.0.2.1:25", "[2001:db8::1]:80"）
	//}

	Addr := unixListener.Addr()

	fmt.Printf("Addr：%v---%T\n", Addr, Addr)
	fmt.Println("Addr.String():", Addr.String())
	fmt.Println("Addr.Network():", Addr.Network())

	// AcceptUnix accepts the next incoming call and returns the new
	// connection.
	// AcceptUnix接受下一个来电并返回新的连接。
	unixConn, e := unixListener.AcceptUnix()
	check_err_net(e)

	src_data := make([]byte, 100)
	i_src, e := unixConn.Read(src_data)
	check_err_net(e)
	fmt.Println("服务端接收到客户端的字节数为：", i_src)
	fmt.Println("服务端接收到客户端的字节数据为：", src_data)
	fmt.Println("服务端接收到客户端的字节转字符串的数据为：", string(src_data))

	// Write implements the Conn Write method.// Write实现Conn Write方法。
	i, e := unixConn.Write([]byte("套接字域socket的服务端响应给客户端的消息"))
	check_err_net(e)
	fmt.Println("服务端响应的字节数为：", i)
	fmt.Println("服务端响应的字符串为：", "套接字域socket的服务端响应给客户端的消息")
	fmt.Println("服务器响应完毕！！！")
	//务必使用命令行先go run这个文件
	//C:\Users\Administrator\Desktop\go_pro\src\io_pro\main3>go run ListenUnix.go
	//输出：
	//	Addr：./raddrUnixSocket---*net.UnixAddr
	//	Addr.String(): ./raddrUnixSocket
	//	Addr.Network(): unix
	//	服务端接收到客户端的字节数为： 52
	//	服务端接收到客户端的字节数据为： [232 191 153 230 152 175 229 174 162 230 136 183 231 171
	//	175 229 143 145 233 128 129 232 191 135 229 142 187 231 187 153 230 156 141 229 138 161
	//	229 153 168 231 154 132 230 149 176 230 141 174 49 49 49 49 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	//	服务端接收到客户端的字节转字符串的数据为： 这是客户端发送过去给服务器的数据1111
	//	服务端响应的字节数为： 57
	//	服务端响应的字符串为： 套接字域socket的服务端响应给客户端的消息
	//	服务器响应完毕！！！



	//下面这几哥方法暂时不讲，需要的请自行了解,下面列出相应的文档：
	//也不是不讲，只是暂时没找到好的例子！

	// SetUnlinkOnClose sets whether the underlying socket file should be removed
	// from the file system when the listener is closed.
	//
	// The default behavior is to unlink the socket file only when package net created it.
	// That is, when the listener and the underlying socket file were created by a call to
	// Listen or ListenUnix, then by default closing the listener will remove the socket file.
	// but if the listener was created by a call to FileListener to use an already existing
	// socket file, then by default closing the listener will not remove the socket file.

	//SetUnlinkOnClose设置在关闭侦听器时是否应从文件系统中删除基础套接字文件。
	//默认行为是仅在程序包net创建套接字文件时才取消其链接。
	//也就是说，当通过调用Listen或ListenUnix创建侦听器和基础套接字文件时，默认情况下，关闭侦听器将删除该套接字文件。
	//但是，如果侦听器是通过调用FileListener来使用现有的套接字文件而创建的，则默认情况下，关闭侦听器不会删除该套接字文件。

	//unixListener.SetUnlinkOnClose(true)

	//Accept()其实跟上面讲到的 AcceptUnix几乎用法一样的！
	//conn, e := unixListener.Accept()
	//check_err_net(e)



}
func check_err_net(err error) {
	if err != nil {
		//fmt.Fprintln(os.Stderr,err)
		//上面的这种方式会导致输出顺序不确定，虽然他可以输出红色的字体，但是由于顺序不确定，我们不采用他！
		fmt.Println("出错了，错误信息为：", err)
		panic(err)
	}
}
