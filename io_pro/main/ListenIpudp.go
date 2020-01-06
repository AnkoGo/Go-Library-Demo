package main

import (
	//"bytes"
	"fmt"
	"net"
)

func main() {
	//laddr := &net.IPAddr{IP: net.IPv4(0,0,0,0).To4()}
	//laddr, err := net.ResolveIPAddr("ip:tcp", "0.0.0.0")
	//check_err_net(err)

	// ListenIP acts like ListenPacket for IP networks.
	//
	// The network must be an IP network name; see func Dial for details.
	//
	// If the IP field of laddr is nil or an unspecified IP address,
	// ListenIP listens on all available IP addresses of the local system
	// except multicast IP addresses.
	//ListenIP的行为类似于IP网络的ListenPacket。
	//该网络必须是IP网络名称； 有关详细信息，请参见func Dial。
	//如果laddr的IP字段为nil或未指定IP地址，则ListenIP侦听本地系统的所有可用IP地址（多播IP地址除外）。
	//var laddr net.IPAddr = net.IPAddr{IP: net.ParseIP("192.168.1.102")}
	conn, e := net.ListenIP("ip:1", nil)
	check_err_net(e)

	//Readby := make([]byte, 300)
	var msg [512]byte
	// ReadFromIP acts like ReadFrom but returns an IPAddr.
	i, e := conn.Read(msg[0:])
	//i, raddr,e := conn.ReadFromIP(Readby)
	check_err_net(e)

	//fmt.Println("客户端的地址为：", raddr)
	fmt.Println("服务器读取的字节数为：", i)
	fmt.Println("服务器读取的字节数据为：", msg)
	//fmt.Println("服务器读取的字节转字符串数据为：", string(msg))//数组不能够转成字符串
	//不知道为什么无法read,先搁置DialIp（）函数吧！

}

func check_err_net(err error) {
	if err != nil {
		//fmt.Fprintln(os.Stderr,err)
		//上面的这种方式会导致输出顺序不确定，虽然他可以输出红色的字体，但是由于顺序不确定，我们不采用他！
		fmt.Println("出错了，错误信息为：", err)
		panic(err)
	}
}
