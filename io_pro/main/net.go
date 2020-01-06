package main

import (
	//"bytes"
	//"encoding/binary"
	"fmt"
	//"os"
	"time"

	//"runtime"

	//"syscall"

	//"sync"
	//"time"

	//"sync"

	//"strconv"

	//"reflect"
	//"time"

	//"io/ioutil"//这些包之所以被我注释掉是因为下面的其他测试我注释掉了很多的代码，只是为了减少不必要的麻烦，在你放开这些注释时候需要自己重新导入这些包！
	"net"
	//"reflect"
	//"time"
	//"time"
)




// 在你阅读本包之前，请务必先去了解必须的tcp/ip/udp/http技术，这里推荐一个b站
// 的视频！https://www.bilibili.com/video/av21179285?p=7，别看音质糟糕，这里讲的非常的不错了

//包说明：
//包net为网络I/O提供了一个便携式接口，包括
//TCP/IP，UDP，域名解析和Unix域套接字。
//
//尽管该软件包提供了对低级网络原语的访问，但是大多数客户端仅需要Dial，
//Listen和Accept函数提供的基本接口以及关联的Conn和Listener接口。 crypto / tls软件包使用相同的接口和相似的Dial and Listen功能。

//下面的很多都被我注释掉了，因为我学习下一个api之前都会将上一个api的代码都注释掉，所以你需要知道！在你学习测试时候需要放开这些代码而且还要自行手动导入未能自动导入的包！

func main() {

	//下面的 代码我并不打算放开注释，你需要学习测试哪部分请自行放开注释然后运行测试
	//在测试本文件代码之前，必须知道本文件运行有些代码依赖外部文件有：
	//  main/dialDNSUdp.go
	//	main/index2.go
	//	main/ListenIpudp.go
	//	main/ListenTcp.go
	//	main/ListenTcp1.go
	//	main/ListenUdp.go
	//	main/ListenUnix.go
	//请到相应的文件中去查看代码，和api的使用用法！这些文件我将不会放到readme.md的导航目录中去快捷显示！请自行查找.本报会感觉有点乱和懵逼，后期会整理，请见谅！

	//
	//fmt.Println("----------net包下的函数之.Dial()--------------")
	//
	//// Dial connects to the address on the named network.
	////
	//// Known networks are "tcp", "tcp4" (IPv4-only), "tcp6" (IPv6-only),
	//// "udp", "udp4" (IPv4-only), "udp6" (IPv6-only), "ip", "ip4"
	//// (IPv4-only), "ip6" (IPv6-only), "unix", "unixgram" and
	//// "unixpacket".
	////
	//// For TCP and UDP networks, the address has the form "host:port".
	//// The host must be a literal IP address, or a host name that can be
	//// resolved to IP addresses.
	//// The port must be a literal port number or a service name.
	//// If the host is a literal IPv6 address it must be enclosed in square
	//// brackets, as in "[2001:db8::1]:80" or "[fe80::1%zone]:80".
	//// The zone specifies the scope of the literal IPv6 address as defined
	//// in RFC 4007.
	//// The functions JoinHostPort and SplitHostPort manipulate a pair of
	//// host and port in this form.
	//// When using TCP, and the host resolves to multiple IP addresses,
	//// Dial will try each IP address in order until one succeeds.
	////
	//// Examples:
	////	Dial("tcp", "golang.org:http")
	////	Dial("tcp", "192.0.2.1:http")
	////	Dial("tcp", "198.51.100.1:80")
	////	Dial("udp", "[2001:db8::1]:domain")
	////	Dial("udp", "[fe80::1%lo0]:53")
	////	Dial("tcp", ":80")
	////
	//// For IP networks, the network must be "ip", "ip4" or "ip6" followed
	//// by a colon and a literal protocol number or a protocol name, and
	//// the address has the form "host". The host must be a literal IP
	//// address or a literal IPv6 address with zone.
	//// It depends on each operating system how the operating system
	//// behaves with a non-well known protocol number such as "0" or "255".
	////
	//// Examples:
	////	Dial("ip4:1", "192.0.2.1")
	////	Dial("ip6:ipv6-icmp", "2001:db8::1")
	////	Dial("ip6:58", "fe80::1%lo0")
	////
	//// For TCP, UDP and IP networks, if the host is empty or a literal
	//// unspecified IP address, as in ":80", "0.0.0.0:80" or "[::]:80" for
	//// TCP and UDP, "", "0.0.0.0" or "::" for IP, the local system is
	//// assumed.
	////
	//// For Unix networks, the address must be a file system path.
	//
	////拨号连接到指定网络上的地址。
	////
	////已知的网络是“ tcp”，“ tcp4”（仅IPv4），“ tcp6”（仅IPv6），“ udp”，“ udp4”（仅IPv4），“ udp6”（仅IPv6），“ ip”，“ ip4”（仅IPv4），“ ip6”（仅IPv6），“ unix”，“ unixgram”和“ unixpacket”。
	////
	////对于TCP和UDP网络，该地址的格式为“ host：port”。
	////主机必须是原义IP地址，或者是可以解析为IP地址的主机名。
	////端口必须是文字端口号或服务名称。
	////如果主机是文本IPv6地址，则必须将其括在方括号中，例如"[2001:db8::1]:80" 或者 "[fe80::1%zone]:80"。
	////zone区域指定RFC 4007中定义的文字IPv6地址的范围。
	////函数JoinHostPort和SplitHostPort以这种形式操纵一对主机和端口。
	////当使用TCP，并且主机解析为多个IP地址时，Dial将依次尝试每个IP地址，直到成功为止。
	////
	//// Examples:
	////	Dial("tcp", "golang.org:http")
	////	Dial("tcp", "192.0.2.1:http")
	////	Dial("tcp", "198.51.100.1:80")
	////	Dial("udp", "[2001:db8::1]:domain")
	////	Dial("udp", "[fe80::1%lo0]:53")
	////	Dial("tcp", ":80")
	////
	////对于IP网络，网络必须为“ ip”，“ ip4”或“ ip6”，后跟冒号和文字协议号或协议名称，并且地址的格式为“主机”。主机必须是文本IP地址或带区域的文本IPv6地址。
	////取决于每个操作系统，操作系统指明了如何以非众所周知的协议编号（例如“ 0”或“ 255”）运行。
	////
	//// Examples:
	////	Dial("ip4:1", "192.0.2.1")
	////	Dial("ip6:ipv6-icmp", "2001:db8::1")
	////	Dial("ip6:58", "fe80::1%lo0")
	////
	////对于TCP，UDP和IP网络，如果主机为空或未指定字面的IP地址，例如对于TCP和UDP，为":80", "0.0.0.0:80" 或者 "[::]:80"，对于IP，使用"", "0.0.0.0" 或者 "::" 作为本地系统。
	////
	////对于Unix网络，该地址必须是文件系统路径。
	//
	////底层实现：
	////	var d Dialer
	////	return d.Dial(network, address)
	////而Dial()方法底层又是调用了DialContext()方法，这个方法的文档如下：
	//
	//// DialContext connects to the address on the named network using
	//// the provided context.
	////
	//// The provided Context must be non-nil. If the context expires before
	//// the connection is complete, an error is returned. Once successfully
	//// connected, any expiration of the context will not affect the
	//// connection.
	////
	//// When using TCP, and the host in the address parameter resolves to multiple
	//// network addresses, any dial timeout (from d.Timeout or ctx) is spread
	//// over each consecutive dial, such that each is given an appropriate
	//// fraction of the time to connect.
	//// For example, if a host has 4 IP addresses and the timeout is 1 minute,
	//// the connect to each single address will be given 15 seconds to complete
	//// before trying the next one.
	////
	//// See func Dial for a description of the network and address
	//// parameters.
	//// DialContext使用提供的上下文连接到命名网络上的地址。
	////
	////提供的Context必须为非null。 如果上下文在连接完成之前到期，则返回错误。 成功连接后，上下文的任何到期都不会影响连接。
	////
	////当使用TCP且address参数中的主机解析为多个网络地址时，任何拨号超时（来自d.Timeout或ctx）都分布在每个连续的拨号上，从而为每个拨号分配了适当的连接时间 。
	////例如，如果主机有4个IP地址，并且超时时间为1分钟，则在尝试下一个地址之前，将为每个地址提供15秒的连接完成时间。
	////
	////有关网络和地址参数的说明，请参见func Dial。
	//
	////关于Dialer对象结构如下：
	////
	////// A Dialer contains options for connecting to an address.
	//////
	////// The zero value for each field is equivalent to dialing
	////// without that option. Dialing with the zero value of Dialer
	////// is therefore equivalent to just calling the Dial function.
	//////拨号程序包含用于连接到地址的选项。
	//////
	//////每个字段的零值等效于不使用该选项的拨号。 因此，使用Dialer的零值进行拨号等效于仅调用Dial函数。
	////type Dialer struct {
	////	// Timeout is the maximum amount of time a dial will wait for
	////	// a connect to complete. If Deadline is also set, it may fail
	////	// earlier.
	////	//
	////	// The default is no timeout.
	////	//
	////	// When using TCP and dialing a host name with multiple IP
	////	// addresses, the timeout may be divided between them.
	////	//
	////	// With or without a timeout, the operating system may impose
	////	// its own earlier timeout. For instance, TCP timeouts are
	////	// often around 3 minutes.
	////	//超时是拨号等待连接完成的最长时间。 如果还设置了截止日期，则它可能会更早失败。
	////	//
	////	//默认值是没有超时。
	////	//
	////	//当使用TCP并拨打具有多个IP地址的主机名时，可能会在其中分配超时。
	////	//
	////	//无论有没有超时，操作系统都可以施加自己的更早超时。 例如，TCP超时通常约为3分钟。
	////	Timeout time.Duration
	////
	////	// Deadline is the absolute point in time after which dials
	////	// will fail. If Timeout is set, it may fail earlier.
	////	// Zero means no deadline, or dependent on the operating system
	////	// as with the Timeout option.
	////	//截止日期是拨号失败的绝对时间点。 如果设置了超时，它可能会更早失败。
	////	//零表示没有截止日期，或者与超时选项一样取决于操作系统。
	////	Deadline time.Time
	////
	////	// LocalAddr is the local address to use when dialing an
	////	// address. The address must be of a compatible type for the
	////	// network being dialed.
	////	// If nil, a local address is automatically chosen.
	////	// LocalAddr是拨打地址时要使用的本地地址。 该地址必须是所拨打网络的兼容类型。
	////	//如果为nil，则会自动选择一个本地地址。
	////	//说明了本地发送端是可以设置地址的！
	////	LocalAddr Addr
	////
	////	// DualStack previously enabled RFC 6555 Fast Fallback
	////	// support, also known as "Happy Eyeballs", in which IPv4 is
	////	// tried soon if IPv6 appears to be misconfigured and
	////	// hanging.
	////	//
	////	// Deprecated: Fast Fallback is enabled by default. To
	////	// disable, set FallbackDelay to a negative value.
	////	// DualStack以前启用了RFC 6555快速回退支持，也称为“开心眼球”，如果IPv6似乎配置错误并挂起，则会很快尝试使用IPv4。
	////	//
	////	//不推荐使用：默认情况下，启用快速回退。 要禁用，请将FallbackDelay设置为负值。
	////	DualStack bool
	////
	////	// FallbackDelay specifies the length of time to wait before
	////	// spawning a RFC 6555 Fast Fallback connection. That is, this
	////	// is the amount of time to wait for IPv6 to succeed before
	////	// assuming that IPv6 is misconfigured and falling back to
	////	// IPv4.
	////	//
	////	// If zero, a default delay of 300ms is used.
	////	// A negative value disables Fast Fallback support.
	////	// FallbackDelay指定在生成RFC 6555快速回退连接之前等待的时间。 也就是说，这是在假设IPv6配置错误并退回到IPv4之前等待IPv6成功的时间。
	////	//
	////	//如果为零，则使用默认延迟300ms。
	////	//负值将禁用快速回退支持。
	////	FallbackDelay time.Duration
	////
	////	// KeepAlive specifies the interval between keep-alive
	////	// probes for an active network connection.
	////	// If zero, keep-alive probes are sent with a default value
	////	// (currently 15 seconds), if supported by the protocol and operating
	////	// system. Network protocols or operating systems that do
	////	// not support keep-alives ignore this field.
	////	// If negative, keep-alive probes are disabled.
	////	// KeepAlive指定活动网络连接的保持活动探测之间的间隔。
	////	//如果为零，则在协议和操作系统支持的情况下，将使用默认值（当前为15秒）发送保持活动的探测。 不支持保持活动状态的网络协议或操作系统将忽略此字段。
	////	//如果为负，则禁用保持活动探测。
	////	KeepAlive time.Duration
	////
	////	// Resolver optionally specifies an alternate resolver to use.
	////	//解析器（可选）指定要使用的备用解析器。
	////	Resolver *Resolver
	////
	////	// Cancel is an optional channel whose closure indicates that
	////	// the dial should be canceled. Not all types of dials support
	////	// cancellation.
	////	//
	////	// Deprecated: Use DialContext instead.
	////	//取消是一个可选通道，其关闭指示应取消拨号。 并非所有类型的拨盘都支持取消。
	////	//
	////	//不推荐使用：改用DialContext。
	////	Cancel <-chan struct{}
	////
	////	// If Control is not nil, it is called after creating the network
	////	// connection but before actually dialing.
	////	//
	////	// Network and address parameters passed to Control method are not
	////	// necessarily the ones passed to Dial. For example, passing "tcp" to Dial
	////	// will cause the Control function to be called with "tcp4" or "tcp6".
	////	//如果Control不为nil，则在创建网络连接之后但实际拨号之前将调用它。
	////	//
	////	//传递给Control方法的网络和地址参数不一定是传递给Dial的参数。 例如，将“ tcp”传递给Dial将导致使用“ tcp4”或“ tcp6”调用控制功能。
	////	Control func(network, address string, c syscall.RawConn) error
	////}
	//
	////// Addr represents a network end point address.
	//////
	////// The two methods Network and String conventionally return strings
	////// that can be passed as the arguments to Dial, but the exact form
	////// and meaning of the strings is up to the implementation.
	////// Addr表示网络端点地址。
	//////
	//////Network()和String()这两种方法通常会返回可以作为Dial参数传递的字符串，但是字符串的确切形式和含义取决于实现方式。
	////type Addr interface {
	////	Network() string // name of the network (for example, "tcp", "udp")//网络名称（例如“ tcp”，“ udp”）
	////	String() string  // string form of address (for example, "192.0.2.1:25", "[2001:db8::1]:80")//地址的字符串形式（例如，"192.0.2.1:25", "[2001:db8::1]:80"）
	////}
	//
	////// DefaultResolver is the resolver used by the package-level Lookup
	////// functions and by Dialers without a specified Resolver.
	////// DefaultResolver是程序包级别的Lookup函数和没有指定解析程序的Dialers使用的解析程序。
	////var DefaultResolver = &Resolver{}
	////
	////// A Resolver looks up names and numbers.
	//////
	////// A nil *Resolver is equivalent to a zero Resolver.
	//////解析程序查找名称和数字。
	//////
	////// 一个为nil 的*Resolver实例对象等于零个Resolver。
	////type Resolver struct {
	////	// PreferGo controls whether Go's built-in DNS resolver is preferred
	////	// on platforms where it's available. It is equivalent to setting
	////	// GODEBUG=netdns=go, but scoped to just this resolver.
	////	// PreferGo控制Go内置的DNS解析器是否在可用的平台上首选。 它等效于设置GODEBUG = netdns = go，但仅限于此解析器。
	////	PreferGo bool
	////
	////	// StrictErrors controls the behavior of temporary errors
	////	// (including timeout, socket errors, and SERVFAIL) when using
	////	// Go's built-in resolver. For a query composed of multiple
	////	// sub-queries (such as an A+AAAA address lookup, or walking the
	////	// DNS search list), this option causes such errors to abort the
	////	// whole query instead of returning a partial result. This is
	////	// not enabled by default because it may affect compatibility
	////	// with resolvers that process AAAA queries incorrectly.
	////	// StrictErrors控制使用Go的内置解析器时的临时错误（包括超时，套接字错误和SERVFAIL）的行为。 对于由多个子查询组成的查询（例如A + AAAA地址查找或遍历DNS搜索列表），
	////	// 此选项会导致此类错误中止整个查询，而不是返回部分结果。 默认情况下未启用此功能，因为它可能会影响与错误处理AAAA查询的解析器的兼容性。
	////	StrictErrors bool
	////
	////	// Dial optionally specifies an alternate dialer for use by
	////	// Go's built-in DNS resolver to make TCP and UDP connections
	////	// to DNS services. The host in the address parameter will
	////	// always be a literal IP address and not a host name, and the
	////	// port in the address parameter will be a literal port number
	////	// and not a service name.
	////	// If the Conn returned is also a PacketConn, sent and received DNS
	////	// messages must adhere to RFC 1035 section 4.2.1, "UDP usage".
	////	// Otherwise, DNS messages transmitted over Conn must adhere
	////	// to RFC 7766 section 5, "Transport Protocol Selection".
	////	// If nil, the default dialer is used.
	////	//拨号（可选）指定备用拨号器，供Go的内置DNS解析器使用，以建立与DNS服务的TCP和UDP连接。 address参数中的主机将始终是原义IP地址而不是主机名，address参数中的端口将是原义端口号而不是服务名。
	////	//如果返回的Conn也是PacketConn，则发送和接收的DNS消息必须遵守RFC 1035第4.2.1节“ UDP使用”。
	////	//否则，通过Conn传输的DNS消息必须遵守RFC 7766第5节“传输协议选择”。
	////	//如果为nil，则使用默认拨号程序。
	////	Dial func(ctx context.Context, network, address string) (Conn, error)
	////
	////	// lookupGroup merges LookupIPAddr calls together for lookups for the same
	////	// host. The lookupGroup key is the LookupIPAddr.host argument.
	////	// The return values are ([]IPAddr, error).
	////	// lookupGroup将LookupIPAddr调用合并在一起，以查找同一主机。 lookupGroup项是LookupIPAddr.host参数。
	////	// 返回值为（[]IPAddr，error）。
	////	lookupGroup singleflight.Group
	////
	////	// TODO(bradfitz): optional interface impl override hook
	////	// TODO(bradfitz): Timeout time.Duration?
	////}
	//
	////// A RawConn is a raw network connection.// RawConn是原始网络连接。
	////type RawConn interface {
	////	// Control invokes f on the underlying connection's file
	////	// descriptor or handle.
	////	// The file descriptor fd is guaranteed to remain valid while
	////	// f executes but not after f returns.
	////	//Control控件在基础连接的文件描述符或句柄上调用f函数。
	////	//确保文件描述符fd在f执行期间保持有效，但在f返回之后无效。
	////	Control(f func(fd uintptr)) error
	////
	////	// Read invokes f on the underlying connection's file
	////	// descriptor or handle; f is expected to try to read from the
	////	// file descriptor.
	////	// If f returns true, Read returns. Otherwise Read blocks
	////	// waiting for the connection to be ready for reading and
	////	// tries again repeatedly.
	////	// The file descriptor is guaranteed to remain valid while f
	////	// executes but not after f returns.
	////	//Read在基础连接的文件描述符或句柄上调用f函数； f函数应该尝试从文件描述符中读取。
	////	//如果f返回true，则Read返回。 否则，“读取”将阻止等待连接准备好进行读取，然后再次尝试。
	////	//确保文件描述符在f执行期间保持有效，但在f返回之后无效。
	////	Read(f func(fd uintptr) (done bool)) error
	////
	////	// Write is like Read but for writing.//Write就像上个字段Read，但是是用于写操作。
	////	Write(f func(fd uintptr) (done bool)) error
	////}
	//
	////// Conn is a generic stream-oriented network connection.
	//////
	////// Multiple goroutines may invoke methods on a Conn simultaneously.
	////// Conn是面向流的通用网络连接。
	//////
	//////多个goroutine可以同时在Conn上调用方法。
	////type Conn interface {
	////	// Read reads data from the connection.
	////	// Read can be made to time out and return an Error with Timeout() == true
	////	// after a fixed time limit; see SetDeadline and SetReadDeadline.
	////	//读取从连接读取数据。
	////	//可以使读取超时并在固定的时间限制后使用Timeout（）== true返回错误； 请参见SetDeadline和SetReadDeadline。
	////	Read(b []byte) (n int, err error)
	////
	////	// Write writes data to the connection.
	////	// Write can be made to time out and return an Error with Timeout() == true
	////	// after a fixed time limit; see SetDeadline and SetWriteDeadline.
	////	// Write将数据写入连接。
	////	//可以使写入超时并在固定的时间限制后使用Timeout（）== true返回错误； 请参见SetDeadline和SetWriteDeadline。
	////	Write(b []byte) (n int, err error)
	////
	////	// Close closes the connection.
	////	// Any blocked Read or Write operations will be unblocked and return errors.
	////	//Close将关闭连接。
	////	//任何阻塞的读或写操作都将被解除阻塞并返回错误。
	////	Close() error
	////
	////	// LocalAddr returns the local network address.
	////	// LocalAddr返回本地网络地址。
	////	LocalAddr() Addr
	////
	////	// RemoteAddr returns the remote network address.
	////	// RemoteAddr返回远程网络地址。
	////	RemoteAddr() Addr
	////
	////	// SetDeadline sets the read and write deadlines associated
	////	// with the connection. It is equivalent to calling both
	////	// SetReadDeadline and SetWriteDeadline.
	////	//
	////	// A deadline is an absolute time after which I/O operations
	////	// fail with a timeout (see type Error) instead of
	////	// blocking. The deadline applies to all future and pending
	////	// I/O, not just the immediately following call to Read or
	////	// Write. After a deadline has been exceeded, the connection
	////	// can be refreshed by setting a deadline in the future.
	////	//
	////	// An idle timeout can be implemented by repeatedly extending
	////	// the deadline after successful Read or Write calls.
	////	//
	////	// A zero value for t means I/O operations will not time out.
	////	//
	////	// Note that if a TCP connection has keep-alive turned on,
	////	// which is the default unless overridden by Dialer.KeepAlive
	////	// or ListenConfig.KeepAlive, then a keep-alive failure may
	////	// also return a timeout error. On Unix systems a keep-alive
	////	// failure on I/O can be detected using
	////	// errors.Is(err, syscall.ETIMEDOUT).
	////	// SetDeadline设置与连接关联的读写期限。 这等效于调用SetReadDeadline和SetWriteDeadline。
	////	//
	////	//截止期限是一个绝对时间，在该绝对时间之后，I / O操作将因超时（请参阅错误类型）而不是阻塞而失败。 截止日期适用于所有将来和未决的I / O，而不仅仅是紧接在其后的读取或写入调用。 超过期限后，可以通过设置将来的期限来刷新连接。
	////	//
	////	//通过在成功的Read或Write调用后重复延长截止期限，可以实现空闲超时。
	////	//
	////	// t的值为零表示I / O操作不会超时。
	////	//
	////	// //请注意，如果TCP连接打开了保持活动状态（除非被Dialer.KeepAlive或ListenConfig.KeepAlive覆盖，否则这是默认设置），则保持活动失败还可能返回超时错误。
	////	// 在Unix系统上，可以使用error.Is（err，syscall.ETIMEDOUT）检测到I / O上的保持活动失败。
	////	SetDeadline(t time.Time) error
	////
	////	// SetReadDeadline sets the deadline for future Read calls
	////	// and any currently-blocked Read call.
	////	// A zero value for t means Read will not time out.
	////	// SetReadDeadline设置将来的Read调用和任何当前阻止的Read调用的截止日期。
	////	// t的值为零表示读取不会超时。
	////	SetReadDeadline(t time.Time) error
	////
	////	// SetWriteDeadline sets the deadline for future Write calls
	////	// and any currently-blocked Write call.
	////	// Even if write times out, it may return n > 0, indicating that
	////	// some of the data was successfully written.
	////	// A zero value for t means Write will not time out.
	////	// SetWriteDeadline设置将来的Write调用和任何当前阻止的Write调用的截止日期。
	////	//即使写入超时，它也可能返回n> 0，表示某些数据已成功写入。
	////	// t的值为零表示写入不会超时。
	////	SetWriteDeadline(t time.Time) error
	////}
	////conn接口继承了众多的io模块的各种读写相关的接口，具体的请看源码可知！
	//
	////conn, e := net.Dial("tcp", "baidu.com:http")//下面在这几种方式也可以
	////conn, e := net.Dial("tcp", "www.baidu.com:80")
	//conn, e := net.Dial("tcp", "baidu.com:80")//语句1
	////conn, e := net.Dial("tcp", "127.0.0.1:8080")//语句2
	//
	//
	//check_err_net(e)
	//
	//fmt.Println(reflect.TypeOf(conn))
	//fmt.Println(conn)
	//defer conn.Close()
	//
	//fmt.Println("---------conn接口连接对象下的read()方法或write()方法---------")
	//
	////下面可以指定什么协议，比如仅访问协议头则指定访问的方式为HEAD，如果需要访问整个页面，则指定访问的方式为GET，
	////但因为仅仅发送了一次请求，所以返回值也是这一次请求中返回的内容，但是需要明白的是，在我们平常访问baidu.com时候不
	////仅仅发起了一次的请求，而是发起了很多次的请求！
	////发送请求数据报，除了这些东西需要指定之外，其他东西都是可省略的！当然你也可以发送其他可省略的数据！
	////末尾之所以要2个\r\n是因为请求头和请求body之间需要空格，这里 仅仅是写上了一行请求头，而请求body为空
	////关于http请求报文格式可以查看https://blog.csdn.net/xyx107/article/details/80436261
	//
	//// n, e := fmt.Fprintf(conn, "HEAD / HTTP/1.0\r\n\r\n")
	////n, e := fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")//GET空格/空格,其中空格不可省略，“/”是要访问的路径，我们也可以改成其他的url
	//n, e := conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))//语句1必须代码，采用这种调用对象上面的额方法来写入，或者采用上面的那种fmt.Fprintf()的形式对conn进行写入都是可以的！底层原理是一样的额！
	////n, e := conn.Write([]byte("GET /index HTTP/1.0\r\n\r\n"))//语句2必须代码
	//
	//check_err_net(e)
	//fmt.Println(n)
	//
	////// TCPConn is an implementation of the Conn interface for TCP network
	////// connections.
	////// TCPConn是用于TCP网络连接的Conn接口的实现。
	////type TCPConn struct {
	////	conn		组合这个结构体，而这个结构体实现了Conn这个接口，也就是TCPConn也实现了这个接口，因为组合就是这样的，一旦组合，意味着着拥有
	////				全部的特性，除了不能像继承那样被认为是conn类之外，几乎都拥有继承的其他特性！比如共享属性和方法！
	////}
	//
	////type conn struct {
	////	fd *netFD
	////}
	////这个结构体实现了read方法，也就是实现了Conn这个接口
	//
	////bytes, e := ioutil.ReadAll(conn)//可以采用这种方式来进行读取，也可以采用下面的这种方式来进行读取！但是目测下面的这种方式会更加麻烦
	//bytes:=make([]byte,1024)
	//n2, e := conn.Read(bytes)
	//fmt.Println(n2)
	//check_err_net(e)
	//
	////如果是ioutil.ReadAll(conn)则采用下面的 这种方式打印输出：
	////fmt.Println(bytes)
	////fmt.Println(string(bytes))
	////fmt.Println(len(bytes))
	////如果是ioutil.ReadAll(conn)则采用上面的 这种方式打印输出
	//
	////如果是conn.Read(bytes)则采用下面的 这种方式打印输出：
	//fmt.Println(bytes[:n2])
	//fmt.Println(string(bytes[:n2]))
	//fmt.Println(len(bytes[:n2]))
	//
	////语句1需要再读取一次，原因在io.reader接口的文档中有，但是语句2不需要，反而会抛出异常！如下：
	//
	//// Reader is the interface that wraps the basic Read method.
	////
	//// Read reads up to len(p) bytes into p. It returns the number of bytes
	//// read (0 <= n <= len(p)) and any error encountered. Even if Read
	//// returns n < len(p), it may use all of p as scratch space during the call.
	//// If some data is available but not len(p) bytes, Read conventionally
	//// returns what is available instead of waiting for more.
	////
	//// When Read encounters an error or end-of-file condition after
	//// successfully reading n > 0 bytes, it returns the number of
	//// bytes read. It may return the (non-nil) error from the same call
	//// or return the error (and n == 0) from a subsequent call.
	//// An instance of this general case is that a Reader returning
	//// a non-zero number of bytes at the end of the input stream may
	//// return either err == EOF or err == nil. The next Read should
	//// return 0, EOF.
	////
	//// Callers should always process the n > 0 bytes returned before
	//// considering the error err. Doing so correctly handles I/O errors
	//// that happen after reading some bytes and also both of the
	//// allowed EOF behaviors.
	////
	//// Implementations of Read are discouraged from returning a
	//// zero byte count with a nil error, except when len(p) == 0.
	//// Callers should treat a return of 0 and nil as indicating that
	//// nothing happened; in particular it does not indicate EOF.
	////
	//// Implementations must not retain p.
	//
	//// Reader是包装基本Read方法的接口。
	////
	////读取最多将len（p）个字节读入p中。它返回读取的字节数（0 <= n <= len（p））和遇到的任何错误。即使Read返回n <len（p），也可能在调用过程中将所有p用作暂存空间。
	////如果某些数据可用但不是len（p）个字节，则按常规方式，Read将返回可用数据，而不是等待更多数据。
	////
	////成功读取n> 0个字节后，当Read遇到错误或文件结束条件时，它将返回读取的字节数。它可能从同一调用返回（非nil）错误，或者从后续调用返回错误（n == 0）。
	////此一般情况的一个实例是，阅读器在输入流的末尾返回非零字节数可能会返回err == EOF或err == nil。下一次读取应返回0，EOF。
	////
	////在考虑错误err之前，调用者应始终处理返回的n> 0个字节。这样做可以正确处理在读取某些字节后发生的I / O错误，以及两种允许的EOF行为。
	////
	////除非len（p）== 0，否则不建议使用Read的实现返回零字节计数且错误为nil的错误。
	////调用者应将返回0和nil视为没有任何反应；特别是它并不表示EOF。
	////
	////实现不得保留p。
	//
	////事实上我们对于read的实现各不相同，但是在这里的话，我们read()方法啥时候返回的话还真说不准，假设你debug的话，他会一次性读取完缓冲池中的额数据报数据，但是
	////假如我们直接run的话，则会立马读取缓冲池中的数据，然后立马返回，他会另外一个线程来进行读取！而不是在当前的调用者所在的线程来进行读取操作，所以，另外一个线程啥时候返回真的说不准，
	////但是我们可以肯定的是，读取一定会一次性读完缓冲池，由于网络延迟到额原因，导致我们第一次读取时候不能完全读取完response，因为此时response还没完全响应光全部的数据，因为tcp是以数据流的
	////形式来进行传送的！事实上，更加准确来说，是一块一块数据传送的，用流来形容也不大准确！每一块都有一定的大小，但是这个大小不能比路经的路由器最小传送大小大！具体的网络技术请自行查找！
	//
	//bytes=make([]byte,1024)
	//n2, e = conn.Read(bytes)
	//fmt.Println(n2)
	//check_err_net(e)
	//
	//fmt.Println(bytes[:n2])
	//fmt.Println(string(bytes[:n2]))
	//fmt.Println(len(bytes[:n2]))
	////如果是conn.Read(bytes)则采用上面的 这种方式打印输出
	//
	//
	////语句1输出
	////如果是采用的io.ReadAll（）的话则输出：
	////	----------net包下的函数--------------
	////	*net.TCPConn，这里说明了虽然返回的是接口，但是实际上的对象是一个类
	////	&{{0xc00007aa00}}
	////	---------试验conn接口连接对象下的方法直接调用是否能输出什么---------
	////	18
	////	[72 84 84 80 47 49 46 49 32 50 48 48 32 79 75 13 10 68 97 116 101 58 32 70 114 105 44 32 48 54
	////	32 68 101 99 32 50 48 49 57 32 49 52 58 51 56 58 52 57 32 71 77 84 13 10 83 101 114 118 101 114
	////	58 32 65 112 97 99 104 101 13 10 76 97 115 116 45 77 111 100 105 102 105 101 100 58 32 84 117
	////	101 44 32 49 50 32 74 97 110 32 50 48 49 48 32 49 51 58 52 56 58 48 48 32 71 77 84 13 10 69 84
	////	97 103 58 32 34 53 49 45 52 55 99 102 55 101 54 101 101 56 52 48 48 34 13 10 65 99 99 101 112
	////	116 45 82 97 110 103 101 115 58 32 98 121 116 101 115 13 10 67 111 110 116 101 110 116 45 76 101
	////	110 103 116 104 58 32 56 49 13 10 67 97 99 104 101 45 67 111 110 116 114 111 108 58 32 109 97 120
	////	45 97 103 101 61 56 54 52 48 48 13 10 69 120 112 105 114 101 115 58 32 83 97 116 44 32 48 55 32 68
	////	101 99 32 50 48 49 57 32 49 52 58 51 56 58 52 57 32 71 77 84 13 10 67 111 110 110 101 99 116 105
	////	111 110 58 32 67 108 111 115 101 13 10 67 111 110 116 101 110 116 45 84 121 112 101 58 32 116 101
	////	120 116 47 104 116 109 108 13 10 13 10 60 104 116 109 108 62 10 60 109 101 116 97 32 104 116 116
	////	112 45 101 113 117 105 118 61 34 114 101 102 114 101 115 104 34 32 99 111 110 116 101 110 116 61
	////	34 48 59 117 114 108 61 104 116 116 112 58 47 47 119 119 119 46 98 97 105 100 117 46 99 111 109 47
	////	34 62 10 60 47 104 116 109 108 62 10]
	////	HTTP/1.1 200 OK
	////	Date: Fri, 06 Dec 2019 14:38:49 GMT
	////	Server: Apache
	////	Last-Modified: Tue, 12 Jan 2010 13:48:00 GMT
	////	ETag: "51-47cf7e6ee8400"
	////	Accept-Ranges: bytes
	////	Content-Length: 81
	////	Cache-Control: max-age=86400
	////	Expires: Sat, 07 Dec 2019 14:38:49 GMT
	////	Connection: Close
	////	Content-Type: text/html
	////
	////	<html>
	////	<meta http-equiv="refresh" content="0;url=http://www.baidu.com/">
	////	</html>
	////
	////	381
	////因为他不会多次请求页面，所以这里只是返回第一次请求时候返回的页面信息！
	//
	////如果是调用的conn.read()方法读取的话则输出：
	////	18
	////	300，这里仅仅是读取了300个字节就返回了，还未读取完的。
	////	[72 84 84 80 47 49 46 49 32 50 48 48 32 79 75 13 10 68 97 116 101 58 32 83 97 116 44 32 49 52 32 68
	////	101 99 32 50 48 49 57 32 49 48 58 51 57 58 51 52 32 71 77 84 13 10 83 101 114 118 101 114 58 32 65
	////	112 97 99 104 101 13 10 76 97 115 116 45 77 111 100 105 102 105 101 100 58 32 84 117 101 44 32 49 50
	////	32 74 97 110 32 50 48 49 48 32 49 51 58 52 56 58 48 48 32 71 77 84 13 10 69 84 97 103 58 32 34 53 49
	////	45 52 55 99 102 55 101 54 101 101 56 52 48 48 34 13 10 65 99 99 101 112 116 45 82 97 110 103 101 115
	////	58 32 98 121 116 101 115 13 10 67 111 110 116 101 110 116 45 76 101 110 103 116 104 58 32 56 49 13 10
	////	67 97 99 104 101 45 67 111 110 116 114 111 108 58 32 109 97 120 45 97 103 101 61 56 54 52 48 48 13 10
	////	69 120 112 105 114 101 115 58 32 83 117 110 44 32 49 53 32 68 101 99 32 50 48 49 57 32 49 48 58 51 57
	////	58 51 52 32 71 77 84 13 10 67 111 110 110 101 99 116 105 111 110 58 32 67 108 111 115 101 13 10 67 111
	////	110 116 101 110 116 45 84 121 112 101 58 32 116 101 120 116 47 104 116 109 108 13 10 13 10]
	////	HTTP/1.1 200 OK
	////	Date: Sat, 14 Dec 2019 10:39:34 GMT
	////	Server: Apache
	////	Last-Modified: Tue, 12 Jan 2010 13:48:00 GMT
	////	ETag: "51-47cf7e6ee8400"
	////	Accept-Ranges: bytes
	////	Content-Length: 81
	////	Cache-Control: max-age=86400
	////	Expires: Sun, 15 Dec 2019 10:39:34 GMT
	////	Connection: Close
	////	Content-Type: text/html
	////
	////
	////	300
	////	81，这里读取剩下的81个字节就遇到了eof然后就返回了，此时已经读取完的。
	////	[60 104 116 109 108 62 10 60 109 101 116 97 32 104 116 116 112 45 101 113 117 105 118 61 34 114 101 102
	////	114 101 115 104 34 32 99 111 110 116 101 110 116 61 34 48 59 117 114 108 61 104 116 116 112 58 47 47 119
	////	119 119 46 98 97 105 100 117 46 99 111 109 47 34 62 10 60 47 104 116 109 108 62 10]
	////	<html>
	////	<meta http-equiv="refresh" content="0;url=http://www.baidu.com/">
	////	</html>
	////
	////	81
	//
	//
	////下面是语句2输出：
	////	23
	////	162
	////	[72 84 84 80 47 49 46 48 32 50 48 48 32 79 75 13 10 67 111 110 116 101 110 116 45 84 121 112 101 58 32 116 101 120 116 47 104 116 109 108 59 32 99 104 97 114 115 101 116 61 117 116 102 45 56 13 10 68 97 116 101 58 32 83 97 116 44 32 49 52 32 68 101 99 32 50 48 49 57 32 49 49 58 49 56 58 50 50 32 71 77 84 13 10 67 111 110 116 101 110 116 45 76 101 110 103 116 104 58 32 52 54 13 10 13 10 60 104 116 109 108 62 13 10 9 60 104 49 62 13 10 9 9 77 97 105 110 32 119 101 98 115 105 116 101 13 10 9 60 47 104 49 62 13 10 60 47 104 116 109 108 62]
	////	HTTP/1.0 200 OK
	////	Content-Type: text/html; charset=utf-8
	////	Date: Sat, 14 Dec 2019 11:18:22 GMT
	////	Content-Length: 46
	////
	////	<html>
	////	<h1>
	////	Main website
	////	</h1>
	////	</html>
	////	162
	////	0
	////	出错了，错误信息为： EOF
	////	panic: EOF
	////
	////	goroutine 1 [running]:
	////	main.check_err_net(0x5351c0, 0xc000032060)
	////	C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:1010 +0xdd
	////	main.main()
	////	C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:548 +0x76a
	////从上面可知，语句2一次性就读完了，所以当你在实现read()方法时候，需要需要一次性读取完缓冲池的话，最好需要有一个比较明确的判断与标准！不过一般都
	////是直接读取光缓冲池中的数据！
	//
	//
	//fmt.Println("----------net包下的函数之.Dial()111--------------")
	////http协议（应用层）基于tcp协议（传输层），具体的报文格式和网络分层可以查看这个网址http://www.023wg.com/message/message/cd_feature_cover.html
	//conn, e = net.Dial("tcp", "localhost:8080")
	//check_err_net(e)
	//fmt.Println(reflect.TypeOf(conn))
	//fmt.Println(conn)
	//defer conn.Close()
	////当然你也可以访问自己写的网站服务器，像上面那句代码就是访问本机的8080端口，但是前提是你必须先创建自己的服务器代码.go文件，文件代码如下：
	//// package main
	//// import (
	//// 	"io"
	//// 	"log"
	//// 	"net/http"
	//// )
	//// func helloHandler(w http.ResponseWriter, r *http.Request) {
	//// 	io.WriteString(w, "Hello, world!")
	//// }
	//// func main() {
	//// 	http.HandleFunc("/hello/world", helloHandler)
	//// 	err := http.ListenAndServe(":8080", nil)
	//// 	if err != nil {
	//// 		log.Fatal("ListenAndServe: ", err.Error())
	//// 	}
	//// }
	//// 然后在控制台中命令：go run xxx.go即可，接着你可以访问到数据了，数据为：
	//// HTTP/1.0 200 OK
	//// Date: Fri, 06 Dec 2019 15:23:37 GMT
	//// Content-Length: 13
	//// Content-Type: text/plain; charset=utf-8
	//
	//// Hello, world!
	//
	////通过查看上面的代码，我们可以看到我们访问的url是自定义的/hello/world
	////GET空格/空格,其中空格不可省略，“/”是要访问的路径，我们也可以改成其他的url,比如下面的/hello/world，注意/hello/world和Get或者HTTP/1.0必须是有单个空格的
	////\r\n\r\n这里之所以不是单个\r\n，而是2个\r\n是因为请求头和请求体之间需要一个空行，这里我们发送的请求body为空！请求body可以为空，但是请求header必须不为空！这
	////是http协议所必需的的
	////n, e = fmt.Fprintf(conn, "GET /hello/world HTTP/1.0\r\n\r\n")//请求1
	//n, e = fmt.Fprintf(conn, "GET /hello/world1 HTTP/1.0\r\n\r\n")//请求2，注意这个路径是不存在的，我们故意让网页返回404，以便查看404的页面
	//check_err_net(e)
	//fmt.Println(n)
	//
	//
	//bytes, e = ioutil.ReadAll(conn)
	//check_err_net(e)
	//
	//
	//fmt.Println(bytes)
	//fmt.Println(string(bytes))
	//fmt.Println(len(bytes))
	////请求1输出：
	////	*net.TCPConn
	////	&{{0xc0000aa000}}
	////	29
	////	[72 84 84 80 47 49 46 48 32 50 48 48 32 79 75 13 10 68 97 116 101 58 32 70 114 105 44 32 49 51
	////	32 68 101 99 32 50 48 49 57 32 48 56 58 51 48 58 51 56 32 71 77 84 13 10 67 111 110 116 101 110
	////	116 45 76 101 110 103 116 104 58 32 49 51 13 10 67 111 110 116 101 110 116 45 84 121 112 101 58
	////	32 116 101 120 116 47 112 108 97 105 110 59 32 99 104 97 114 115 101 116 61 117 116 102 45 56 13
	////	10 13 10 72 101 108 108 111 44 32 119 111 114 108 100 33]
	////	HTTP/1.0 200 OK
	////	Date: Fri, 13 Dec 2019 08:30:38 GMT
	////	Content-Length: 13
	////	Content-Type: text/plain; charset=utf-8
	////
	////	Hello, world!
	////	130
	//
	//
	////请求2输出：
	////	*net.TCPConn
	////	&{{0xc00009c000}}
	////	30
	////	[72 84 84 80 47 49 46 48 32 52 48 52 32 78 111 116 32 70 111 117 110 100 13 10 67 111 110 116 101
	////	110 116 45 84 121 112 101 58 32 116 101 120 116 47 112 108 97 105 110 59 32 99 104 97 114 115 101
	////	116 61 117 116 102 45 56 13 10 88 45 67 111 110 116 101 110 116 45 84 121 112 101 45 79 112 116 105
	////	111 110 115 58 32 110 111 115 110 105 102 102 13 10 68 97 116 101 58 32 70 114 105 44 32 49 51 32
	////	68 101 99 32 50 48 49 57 32 48 56 58 51 52 58 51 52 32 71 77 84 13 10 67 111 110 116 101 110 116 45
	////	76 101 110 103 116 104 58 32 49 57 13 10 13 10 52 48 52 32 112 97 103 101 32 110 111 116 32 102 111
	////	117 110 100 10]
	////	HTTP/1.0 404 Not Found
	////	Content-Type: text/plain; charset=utf-8
	////	X-Content-Type-Options: nosniff
	////	Date: Fri, 13 Dec 2019 08:34:34 GMT
	////	Content-Length: 19
	////
	////	404 page not found
	////
	////	176
	//
	////可以对比请求1和请求2中返回的字段信息是不完全一样的！
	////如果你通过在浏览器中访问该url，可以看到浏览器中接收到的数据比上面多的多，其实，浏览器也是接收到了上面显示出来的那些数据而已，但是
	////浏览器会自动帮你填充一些额外的数据（包括response header和response body）以供展示页面所使用的！
	//
	//
	//
	//
	//
	//fmt.Println("----------net包下的函数之.Dial()222之返回html页面内容--------------")
	////再开始这个测试之前，务必做以下的额工作，在命令行中单独打开index2.go文件，然后执行命令go run index2.go,相当于开启一个服务器，关于index2.go中的
	////作为响应客户端的服务器代码和templates/index.tmpl中的模板文件内容可以自己查看，都是要先安装gin框架才可以进行的！上述说到的2个文件中的代码来自：https://gin-gonic.com/zh-cn/docs/examples/html-rendering/
	////事实上我们也可以不使用gin,而使用net.listen()函数来接收，但是我们还没学到那个函数，就先这样！
	//
	//conn, e = net.Dial("tcp", "localhost:8080")
	//check_err_net(e)
	//fmt.Println(reflect.TypeOf(conn))
	//fmt.Println(conn)
	//defer conn.Close()
	//
	//n, e = fmt.Fprintf(conn, "GET /index HTTP/1.0\r\n\r\n")//请求1
	//check_err_net(e)
	//fmt.Println(n)
	//
	//
	//bytes, e = ioutil.ReadAll(conn)
	//check_err_net(e)
	//
	//
	//fmt.Println(bytes)
	//fmt.Println(string(bytes))
	//fmt.Println(len(bytes))
	////在做完以上的工作后，请打开控制台，然后在控制台中输入go run compress_zlib.go,接着你会看到很多输出，但是上部分的输出我们省略，我们只看当前讨论内容
	////部分的输出，如下：
	////*net.TCPConn
	////&{{0xc0000dc000}}
	////23
	////[72 84 84 80 47 49 46 48 32 50 48 48 32 79 75 13 10 67 111 110 116 101 110 116 45 84 121 112 101 58 32 116 101 120 116 47 104 116 109 108 59 32 99 104 97 114 115 101 116 61 117 116 102 45 56 13 10 68 97 116 101 58 32 83 97 116 44 32 49 52 32 68 101 99 32 50 48 49 5
	////7 32 48 54 58 52 48 58 50 51 32 71 77 84 13 10 67 111 110 116 101 110 116 45 76 101 110 103 116 104 58 32 52 54 13 10 13 10 60 104 116 109 108 62 13 10 9 60 104 49 62 13 10 9 9 77 97 105 110 32 119 101 98 115 105 116 101 13 10 9 60 47 104 49 62 13 10 60 47 104 116
	////109 108 62]
	////HTTP/1.0 200 OK
	////Content-Type: text/html; charset=utf-8
	////Date: Sat, 14 Dec 2019 06:40:23 GMT
	////Content-Length: 46
	////
	////<html>
	////		<h1>
	////				Main website
	////		</h1>
	////</html>
	////162
	////上面的文字格式可能跟控制台中输出的不一样！具体以控制台中的输出为准，之所以复制到这里会发生格式改变是因为goland的原因！
	//
	//
	////关于http的post方式的请求,以及udp协议和ipv6的展示，我不再一一举例了！以后可能会补充这块！
	//
	//
	//fmt.Println("----------net包下的函数之.DialTimeout()，限时建立连接--------------")
	//
	//
	////  》》DialTimeout acts like Dial but takes a timeout.（ 》》是我多加的东西，为了方便醒目阅读）
	////
	//// The timeout includes name resolution, if required.
	//// When using TCP, and the host in the address parameter resolves to
	//// multiple IP addresses, the timeout is spread over each consecutive
	//// dial, such that each is given an appropriate fraction of the time
	//// to connect.
	////
	//// See func Dial for a description of the network and address
	//// parameters.
	//// 》》DialTimeout的作用类似于Dial，但是多需要一个超时时间参数。（ 》》是我多加的东西，为了方便醒目阅读）
	////
	////超时包括名称解析（如果需要）。
	////当使用TCP且address参数中的主机解析为多个IP地址时，超时分布在每个连续的拨号上，从而为每个拨号分配了适当的连接时间。
	////
	////有关网络和地址参数的说明，请参见func Dial。
	////底层几乎跟net.Dial()函数完全一样，区别也仅仅是多了一个timeout参数
	////conn1, e := net.DialTimeout("tcp", "baidu.com:80", 1e4)//这个时间足够短，好让他超时
	//conn1, e := net.DialTimeout("tcp", "baidu.com:80", 1e9)//这个时间足够长，可以让他足够时间返回响应！
	//check_err_net(e)
	//
	//i, e := fmt.Fprintf(conn1, "GET / HTTP/1.0\r\n\r\n")
	//check_err_net(e)
	//fmt.Println("从客户端发送http请求到服务器的字节个数为：",i)
	//
	////在上面，我们是往conn1对象写入，获取数据时候也是从conn1读取获取的！
	//resv_data, e := ioutil.ReadAll(conn1)
	//check_err_net(e)
	//
	//fmt.Println("服务器返回来的字节切片是：",resv_data)
	//fmt.Println("服务器返回来的字节切片转字符串是：",string(resv_data))
	////参数timeout为1e4时输出：
	////出错了，错误信息为： dial tcp 39.156.69.79:80: i/o timeout
	////panic: dial tcp 39.156.69.79:80: i/o timeout
	////
	////goroutine 1 [running]:
	////main.check_err_net(0x533f40, 0xc0000e80a0)
	////	C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:721 +0xdd
	////main.main()
	////	C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:700 +0x1136
	//
	////参数timeout为1e9时输出：
	////从客户端发送http请求到服务器的字节个数为： 18
	////服务器返回来的字节切片是： [72 84 84 80 47 49 46 49 32 50 48 48 32 79 75 13 10 68 97 116 101 58 32 83 97 116
	////44 32 49 52 32 68 101 99 32 50 48 49 57 32 48 55 58 50 51 58 50 54 32 71 77 84 13 10 83 101 114 118 101
	////114 58 32 65 112 97 99 104 101 13 10 76 97 115 116 45 77 111 100 105 102 105 101 100 58 32 84 117 101 44
	////32 49 50 32 74 97 110 32 50 48 49 48 32 49 51 58 52 56 58 48 48 32 71 77 84 13 10 69 84 97 103 58 32 34 53
	////49 45 52 55 99 102 55 101 54 101 101 56 52 48 48 34 13 10 65 99 99 101 112 116 45 82 97 110 103 101 115 58
	////32 98 121 116 101 115 13 10 67 111 110 116 101 110 116 45 76 101 110 103 116 104 58 32 56 49 13 10 67 97 99
	////104 101 45 67 111 110 116 114 111 108 58 32 109 97 120 45 97 103 101 61 56 54 52 48 48 13 10 69 120 112 105
	////114 101 115 58 32 83 117 110 44 32 49 53 32 68 101 99 32 50 48 49 57 32 48 55 58 50 51 58 50 54 32 71 77 84
	////13 10 67 111 110 110 101 99 116 105 111 110 58 32 67 108 111 115 101 13 10 67 111 110 116 101 110 116 45 84
	////121 112 101 58 32 116 101 120 116 47 104 116 109 108 13 10 13 10 60 104 116 109 108 62 10 60 109 101 116 97
	////32 104 116 116 112 45 101 113 117 105 118 61 34 114 101 102 114 101 115 104 34 32 99 111 110 116 101 110 116
	////61 34 48 59 117 114 108 61 104 116 116 112 58 47 47 119 119 119 46 98 97 105 100 117 46 99 111 109 47 34 62 10
	////60 47 104 116 109 108 62 10]
	////服务器返回来的字节切片转字符串是： HTTP/1.1 200 OK
	////Date: Sat, 14 Dec 2019 07:23:26 GMT
	////Server: Apache
	////Last-Modified: Tue, 12 Jan 2010 13:48:00 GMT
	////ETag: "51-47cf7e6ee8400"
	////Accept-Ranges: bytes
	////Content-Length: 81
	////Cache-Control: max-age=86400
	////Expires: Sun, 15 Dec 2019 07:23:26 GMT
	////Connection: Close
	////Content-Type: text/html
	////
	////<html>
	////<meta http-equiv="refresh" content="0;url=http://www.baidu.com/">
	////</html>
	//
	//
	//
	//fmt.Println("----------net包下的函数之.DialTimeout()，限时建立连接111--------------")
	//
	//
	////这部分的谈论跟上面几乎完全一样，我们探究gin创建的服务器返回的数据，亲体是我们先要运行命令：go run index2.go
	////conn1, e = net.DialTimeout("tcp", "127.0.0.1:8080", time.Duration(-1))//因为是本地，所以几乎不需要时间，我这里给0都无法抛出timeout错误，只能给个负数了，好让他超时
	//conn1, e = net.DialTimeout("tcp", "127.0.0.1:8080", 1e9)//这个时间足够长，可以让他足够时间返回响应！
	////需要注意的是，如果一个主机上面有多个ip地址的话，那么这个timeout时间会被平均分割成一定的时间片，然后按照这个时间片来访问这台主机上面的每一个ip，
	//check_err_net(e)
	//
	//i, e = fmt.Fprintf(conn1, "GET /index HTTP/1.0\r\n\r\n")
	//check_err_net(e)
	//fmt.Println("从客户端发送http请求到服务器的字节个数为：",i)
	//
	////在上面，我们是往conn1对象写入，获取数据时候也是从conn1读取获取的！
	//resv_data, e = ioutil.ReadAll(conn1)
	//check_err_net(e)
	//
	//fmt.Println("服务器返回来的字节切片是：",resv_data)
	//fmt.Println("服务器返回来的字节切片转字符串是：",string(resv_data))
	//
	////参数timeout为-1时输出：
	////出错了，错误信息为： dial tcp 127.0.0.1:8080: i/o timeout
	////panic: dial tcp 127.0.0.1:8080: i/o timeout
	////
	////goroutine 1 [running]:
	////main.check_err_net(0x5340e0, 0xc00008e190)
	////	C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:807 +0xdd
	////main.main()
	////	C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:761 +0x1512
	//
	//
	//
	////参数timeout为1e9时输出：
	////从客户端发送http请求到服务器的字节个数为： 23
	////服务器返回来的字节切片是： [72 84 84 80 47 49 46 48 32 50 48 48 32 79 75 13 10 67 111 110 116 101 110 116 45 84
	//// 121 112 101 58 32 116 101 120 116 47 104 116 109 108 59 32 99 104 97 114 115 101 116 61 117 116 102 45 56
	//// 13 10 68 97 116 101 58 32 83 97 116 44 32 49 52 32 68 101 99 32 50 48 49 57 32 48 55 58 51 48 58 49 54 32
	//// 71 77 84 13 10 67 111 110 116 101 110 116 45 76 101 110 103 116 104 58 32 52 54 13 10 13 10 60 104 116 109
	//// 108 62 13 10 9 60 104 49 62 13 10 9 9 77 97 105 110 32 119 101 98 115 105 116 101 13 10 9 60 47 104 49 62
	//// 13 10 60 47 104 116 109 108 62]
	////服务器返回来的字节切片转字符串是： HTTP/1.0 200 OK
	////Content-Type: text/html; charset=utf-8
	////Date: Sat, 14 Dec 2019 07:30:16 GMT
	////Content-Length: 46
	////
	////<html>
	////<h1>
	////Main website
	////</h1>
	////</html>
	//
	//
	//fmt.Println("----------net包下的函数之DialUnix()，同一台主机不同程序之间通信的嵌套字--------------")
	////在开始学习这个api之前务必先去了解什么是unix域套接字，可以参考这里https://www.jianshu.com/p/78103b2a74be
	////这个机制不仅限于linux系统,也可以用于windows系统
	//
	//// BUG(mikio): On JS, NaCl and Plan 9, methods and functions related
	//// to UnixConn and UnixListener are not implemented.
	//// BUG(mikio): On Windows, methods and functions related to UnixConn
	//// and UnixListener don't work for "unixgram" and "unixpacket".
	//
	//// BUG（mikio）：在JS，NaCl和Plan 9上，未实现与UnixConn和UnixListener相关的方法和功能。
	//// BUG（mikio）：在Windows上，与UnixConn和UnixListener相关的方法和功能不适用于“ unixgram”和“ unixpacket”。
	//
	//// DialUnix acts like Dial for Unix networks.
	////
	//// The network must be a Unix network name; see func Dial for details.
	////
	//// If laddr is non-nil, it is used as the local address for the
	//// connection.
	//// DialUnix的行为类似于Unix网络的Dial。
	//// 网络必须是Unix网络名称； 有关详细信息，请参见func Dial。
	//// 如果laddr不为nil，则将其用作连接的本地地址。
	//
	//// 对于Unix网络，该地址必须是文件系统路径。
	//
	//
	//// UnixAddr represents the address of a Unix domain socket end point.
	////// UnixAddr表示Unix域套接字端点的地址。
	////type UnixAddr struct {
	////	Name string
	////	Net  string
	////}
	////这个结构体实现了net.Addr这个接口
	//
	////// OpError is the error type usually returned by functions in the net
	////// package. It describes the operation, network type, and address of
	////// an error.
	////// OpError是通常由net包中的函数返回的错误类型。 它描述了错误的操作，网络类型和地址。
	////type OpError struct {
	////	// Op is the operation which caused the error, such as
	////	// "read" or "write".
	////	// Op是导致错误的操作，例如“ read”或“ write”。
	////	Op string
	////
	////	// Net is the network type on which this error occurred,
	////	// such as "tcp" or "udp6".
	////	// Net是发生此错误的网络类型，例如“ tcp”或“ udp6”。
	////	Net string
	////
	////	// For operations involving a remote network connection, like
	////	// Dial, Read, or Write, Source is the corresponding local
	////	// network address.
	////	//对于涉及远程网络连接的操作，例如Dial，Read或Write，Source是相应的本地网络地址。
	////	Source Addr
	////
	////	// Addr is the network address for which this error occurred.
	////	// For local operations, like Listen or SetDeadline, Addr is
	////	// the address of the local endpoint being manipulated.
	////	// For operations involving a remote network connection, like
	////	// Dial, Read, or Write, Addr is the remote address of that
	////	// connection.
	////	// Addr是发生此错误的网络地址。
	////	//对于本地操作（如Listen或SetDeadline），Addr是要操作的本地端点的地址。
	////	//对于涉及远程网络连接的操作，例如Dial，Read或Write，Addr是该连接的远程地址。
	////	Addr Addr
	////
	////	// Err is the error that occurred during the operation.
	////	// Err是操作期间发生的错误。
	////	Err error
	////}
	////这个结构体实现了很多的有关错误的接口，错误接口并不只是errors模块中才有的，net包中也有自定义区别于errors包的专属的错误接口！
	//
	////// sysDialer contains a Dial's parameters and configuration.
	////// sysDialer包含一个Dial的参数和配置。
	////type sysDialer struct {
	////	Dialer		组合了这个结构体，指明了下面字段信息或者叫配置参数附属于哪一个Dial实例对象
	////	network, address string
	////}
	////这个结构体有3个重要的内部不公开的方法（dialParallel，dialSerial，dialSingle），文档分别如下：
	////
	//// dialParallel races two copies of dialSerial, giving the first a head start. It returns the first established connection and closes the others. Otherwise it returns an error from the first primary address.
	//
	//// dialSerial connects to a list of addresses in sequence, returning either the first successful connection, or the first error.
	//
	//// dialSingle attempts to establish and returns a single connection to the destination address.
	//
	//// DialParallel争夺DialSerial的两个副本，使第一个副本抢先一步。 它返回第一个建立的连接，并关闭其他连接。 否则，它将从第一个主地址返回错误。
	//
	//// DialSerial依次连接到地址列表，返回第一个成功的连接或第一个错误。
	//
	//// DialSingle尝试建立并返回到目标地址的单个连接。
	////当然你也可以自己去看源码，我这里仅仅是提示下
	//
	//
	////下面是返回值类型的介绍，跟上面的不同，他不返回conn对象，而是返回了UnixConn对象
	//
	////// UnixConn is an implementation of the Conn interface for connections
	////// to Unix domain sockets.
	////// UnixConn是Conn接口的实现，用于连接到Unix域套接字。
	////type UnixConn struct {
	////	conn
	////}
	////下面的Name字段（./laddrUnixSocket）不一定要给相对路径，也可以给完整的绝对路径比如C:\\Users\\Administrator\\Desktop\\go_pro\\src\\io_pro\\main3\\laddrUnixSocket，我就不测试了！
	////其实这个net作为网络名字可以随意给，虽然文档上说应该给的是unix，unixgram或者unixpacket，我觉得给空都是可以的！我就不测试了！
	////但是需要注意的是raddrUnixSocket必须和 服务器中定义的socket套接字的文件名相同，否则会找寻不到！因为这是命名的unix socket通信，是按照名字来指定连接的socket来通信的！
	//var lUnixAddr= &net.UnixAddr{Name: "./laddrUnixSocket", Net:  "localnet"}
	//var rUnixAddr= &net.UnixAddr{Name: "./raddrUnixSocket", Net:  "remotenet"}
	//
	//printInfo:= func() {
	//	fmt.Println("----------")
	//	fmt.Println("lUnixAddr:",lUnixAddr)
	//	fmt.Println("lUnixAddr.Name:",lUnixAddr.Name)
	//	fmt.Println("lUnixAddr.Net:",lUnixAddr.Net)
	//	fmt.Println("lUnixAddr.String():",lUnixAddr.String())
	//	fmt.Println("lUnixAddr.Network():",lUnixAddr.Network())
	//
	//	fmt.Println()
	//	fmt.Println("rUnixAddr:",rUnixAddr)
	//	fmt.Println("rUnixAddr.Name:",rUnixAddr.Name)
	//	fmt.Println("rUnixAddr.Net:",rUnixAddr.Net)
	//	fmt.Println("rUnixAddr.String():",rUnixAddr.String())//其实就是返回rUnixAddr.Name
	//	fmt.Println("rUnixAddr.Network():",rUnixAddr.Network())//其实就是返回rUnixAddr实现了Addr接口，
	//	fmt.Println("----------")
	//}
	////之所以放到函数里面是因为在下面还要复用，在调用一次
	//printInfo()
	//
	//unixConn, e := net.DialUnix("unix",lUnixAddr, rUnixAddr )
	//check_err_net(e)
	//
	//printInfo()
	//
	//// LocalAddr returns the local network address.
	//// The Addr returned is shared by all invocations of LocalAddr, so
	//// do not modify it.
	//// LocalAddr返回本地网络地址。
	////返回的Addr由LocalAddr的所有调用共享，因此请勿修改它。
	//fmt.Println("unixConn.LocalAddr()",unixConn.LocalAddr())
	//fmt.Println("unixConn.LocalAddr().String():",unixConn.LocalAddr().String())
	//fmt.Println("unixConn.LocalAddr().Network():",unixConn.LocalAddr().Network())
	//
	//// RemoteAddr returns the remote network address.
	//// The Addr returned is shared by all invocations of RemoteAddr, so
	//// do not modify it.
	//// RemoteAddr返回远程网络地址。
	////返回的Addr由RemoteAddr的所有调用共享，因此请勿修改它。
	//fmt.Println("unixConn.RemoteAddr()",unixConn.RemoteAddr())
	//fmt.Println("unixConn.RemoteAddr().String():",unixConn.RemoteAddr().String())
	//fmt.Println("unixConn.RemoteAddr().Network():",unixConn.RemoteAddr().Network())
	//
	//// Close closes the connection.
	////这个关闭仅仅会关闭连接，但是不会删除为socket通信而创建的laddrUnixSocket文件，
	////但是，你可以查看ListenUnix.go文件中的unixListener.Close()方法不但是close连接，还会先remove为socket通信而创建的raddrUnixSocket文件
	////这就是他们之间的不同，所以我们可以手动删除！
	//defer unixConn.Close()
	//
	//write_i, e := unixConn.Write([]byte("这是客户端发送过去给服务器的数据1111"))
	//check_err_net(e)
	//fmt.Println("客户端发送了的字节数为：",write_i)
	//fmt.Println("客户端发送了的字节数据为：","这是客户端发送过去给服务器的数据1111")
	//
	//resv_data_Client:=make([]byte,100)
	//read_i, e := unixConn.Read(resv_data_Client)
	//
	//fmt.Println("客户端接收到服务器响应回来的字节数为：",read_i)
	//fmt.Println("客户端接收到服务器响应回来的字节数据为：",resv_data_Client)
	//fmt.Println("客户端接收到服务器响应回来的字节转字符串数据为：",string(resv_data_Client))
	////上面的代码是作为客户端代码来运行的，在你看完这里后请去看文件ListenUnix.go中的服务端的代码，他们是相互配合的！代表2个不同的程序之间的socket的通信
	////然后现在控制要命令行中输入go run ListenUnix.go,接着再开一个命令行窗口，这时候才是运行上面的代码的时候，在新打开的命令行窗口中输入命令
	////go run compress_zlib.go(当前文件的文件名)，注意务必使用控制台命令行的方式单独运行，不可使用其他的编辑器！否则会出错！
	////需要注意的是，当你go run ListenUnix.go后，你会发现在当前的目录下多了一个raddrUnixSocket文件，然后当你go run compress_zlib.go时候raddrUnixSocket被删除，但是
	////又新创建了一个laddrUnixSocket文件，原因在上面已经说了，跟各个对象的close()处理方式有关！
	//
	////输出：
	////	----------
	////	lUnixAddr: ./laddrUnixSocket
	////	lUnixAddr.Name: ./laddrUnixSocket
	////	lUnixAddr.Net: localnet
	////	lUnixAddr.String(): ./laddrUnixSocket
	////	lUnixAddr.Network(): localnet
	////
	////	rUnixAddr: ./raddrUnixSocket
	////	rUnixAddr.Name: ./raddrUnixSocket
	////	rUnixAddr.Net: remotenet
	////	rUnixAddr.String(): ./raddrUnixSocket
	////	rUnixAddr.Network(): remotenet
	////	----------
	////	----------
	////	lUnixAddr: ./laddrUnixSocket
	////	lUnixAddr.Name: ./laddrUnixSocket
	////	lUnixAddr.Net: localnet
	////	lUnixAddr.String(): ./laddrUnixSocket
	////	lUnixAddr.Network(): localnet
	////
	////	rUnixAddr: ./raddrUnixSocket
	////	rUnixAddr.Name: ./raddrUnixSocket
	////	rUnixAddr.Net: remotenet
	////	rUnixAddr.String(): ./raddrUnixSocket
	////	rUnixAddr.Network(): remotenet
	////	----------
	////	unixConn.LocalAddr() ./laddrUnixSocket
	////	unixConn.LocalAddr().String(): ./laddrUnixSocket
	////	unixConn.LocalAddr().Network(): unix,这里之所以跟上面的都不相同，原因在下面被说到！
	////	unixConn.RemoteAddr() ./raddrUnixSocket
	////	unixConn.RemoteAddr().String(): ./raddrUnixSocket
	////	unixConn.RemoteAddr().Network(): unix,这里之所以跟上面的都不相同，原因在下面被说到！
	////	客户端发送了的字节数为： 52
	////	客户端发送了的字节数据为： 这是客户端发送过去给服务器的数据1111
	////	客户端接收到服务器响应回来的字节数为： 57
	////	客户端接收到服务器响应回来的字节数据为： [229 165 151 230 142 165 229 173 151 229 159 159 115 111 99 107 101 116 231 154 132 230 156 141 229 138 161 231 171 175 229 147 141 229 186 148 231 187 153 229 174 162 230 136 183 231 171 175 231 154 132 230 182 136 230 129 175 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	////	客户端接收到服务器响应回来的字节转字符串数据为： 套接字域socket的服务端响应给客户端的消息
	//
	////从上面可以看得出，这个函数并没有更改net.UnixAddr实例对象的参数，但是，获取函数的返回值UnixConn.LocalAddr()和UnixConn.RemoteAddr()获取
	//// 到的net.UnixAddr实例对象不同于前面说的那个对象了，而是这个函数自己新建的对象实例！所以才会出现不同。
	////反而，他们是下面的这个对象上面的字段laddr和raddr的属性的值：
	////// Network file descriptor.
	////type netFD struct {
	////	pfd poll.FD
	////
	////	// immutable until Close
	////	family      int
	////	sotype      int
	////	isConnected bool // handshake completed or use of association with peer
	////	net         string
	////	laddr       Addr//虽然这个Addr是接口，但是在这个例子中是这个接口的实现类对象net.UnixAddr
	////	raddr       Addr//虽然这个Addr是接口，但是在这个例子中是这个接口的实现类对象net.UnixAddr
	////}
	//
	////运行完须知，因为每次运行完DialUnix()函数后都会在项目目录或者子目录（取决于你用什么编辑器运行，goland和cmd运行的产生位置不一样），
	////所以每次运行完后都要手动删除laddrUnixSocket文件，如果你没有看到这个文件，那么请自行刷新一下项目目录（goland中鼠标放到项目目录文件夹，右键
	//// ，最后哪里有个刷新的图标就是了！），一定要先删除了laddrUnixSocket文件然后才能再次运行！当然你也可以自己写个删除函数！我这里就不写了！
	//
	//
	//fmt.Println("----------net包下的函数之.DialTCP()发送tcp数据或者接收tcp数据--------------")
	////其实这个api类似.Dial()函数，只不过.Dial()函数可以用于多种协议，而这个函数仅限tcp协议！
	//
	//// DialTCP acts like Dial for TCP networks.
	////
	//// The network must be a TCP network name; see func Dial for details.
	////
	//// If laddr is nil, a local address is automatically chosen.
	//// If the IP field of raddr is nil or an unspecified IP address, the
	//// local system is assumed.
	//// DialTCP的作用类似于TCP网络的Dial。
	////
	////网络必须是TCP网络名称； 有关详细信息，请参见func Dial。
	////
	////如果laddr为nil，则会自动选择一个本地地址。
	////如果raddr的IP字段为nil或未指定IP地址，则使用本地系统。
	//
	//
	//		//// BUG(mikio): On JS, NaCl and Windows, the File method of TCPConn and
	//		//// TCPListener is not implemented.
	//		//
	//		//// TCPAddr represents the address of a TCP end point.
	//		//// BUG（mikio）：在JS，NaCl和Windows上，未实现TCPConn和TCPListener的File方法。
	//		//// TCPAddr表示TCP端点的地址。
	//		//type TCPAddr struct {
	//		//	IP   IP
	//		//	Port int
	//		//	Zone string // IPv6 scoped addressing zone
	//		//}
	//		//
	//		//// IP address lengths (bytes).// IP地址长度（字节）。
	//		//const (
	//		//	IPv4len = 4
	//		//	IPv6len = 16
	//		//)
	//		//
	//		//// An IP is a single IP address, a slice of bytes.
	//		//// Functions in this package accept either 4-byte (IPv4)
	//		//// or 16-byte (IPv6) slices as input.
	//		////
	//		//// Note that in this documentation, referring to an
	//		//// IP address as an IPv4 address or an IPv6 address
	//		//// is a semantic property of the address, not just the
	//		//// length of the byte slice: a 16-byte slice can still
	//		//// be an IPv4 address.
	//		//// IP是一个IP地址，本质是一个字节切片
	//		////此包中的函数接受4字节（IPv4）或16字节（IPv6）片作为输入。
	//		//// //请注意，在本文档中，将IP地址称为IPv4地址或IPv6地址是该地址的语义属性，而不仅仅是字节片的长度：16字节片仍然可以是IPv4地址。
	//		//type IP []byte
	//
	//TestTCPConnSpecificMethods:=func () {
	//	//注意了，这个函数写的是tcp的客户端代码，服务段的代码请自行到ListenTcp.go中去查看。
	//	//如果你已经看了listenUnix的代码，那么你应该知道，要先启动服务端的代码，然后再独自运行当前文件中客户端中的代码！注意文件要单独运行！
	//
	//	ra, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:9999")
	//	check_err_net(err)
	//
	//	//第一个参数可以为tcp,tcp4或者tcp6
	//	//事实上这里是在底层新建一个连接对象，然后开启一个g程去每隔一定的时间就发起连接请求，直到连接成功。
	//	tcpConn, err := net.DialTCP("tcp4", nil, ra)//指定要发送数据到服务器端的9999端口，服务器从这个端口接收发送端发送过来的数据！
	//	check_err_net(err)
	//
	//	defer tcpConn.Close()
	//
	//	// SetKeepAlive sets whether the operating system should send
	//	// keep-alive messages on the connection.
	//	// SetKeepAlive设置操作系统是否应在连接上发送保持活动消息。
	//	e=tcpConn.SetKeepAlive(false)
	//	check_err_net(e)
	//
	//	// SetKeepAlivePeriod sets period between keep-alives.
	//	// SetKeepAlivePeriod设置保持活动之间的时间间隔。也就是每隔多久向服务器发送一次确认客户端还活跃的请求！而不是设置连接的时间！
	//	//这个方法 和上面的SetKeepAlive（）都是为了确保客户端或者服务端还活跃的一种确认方式，这个方式可以是客户端去探测服务器端是否还活跃，
	//	//也可以是服务器端去探测客户端是否还处于活跃的状态，如果不设置的话，则默认是2h的时间！这里我们是客户端发送数据然后来决定是否继续发送数据
	//	//到服务器端，事实上，一般情况下都是在服务端去探测客户端的！但是这里既然我们写了代码，我就懒得改为将下面的代码copy到服务端去了！希望
	//	//读者注意！具体的tcp keepalive知识可以查看这个文章：https://blog.csdn.net/chrisnotfound/article/details/80111559
	//	e=tcpConn.SetKeepAlivePeriod(3 * time.Second)
	//	check_err_net(e)
	//
	//	// SetLinger sets the behavior of Close on a connection which still
	//	// has data waiting to be sent or to be acknowledged.
	//	//
	//	// If sec < 0 (the default), the operating system finishes sending the
	//	// data in the background.
	//	//
	//	// If sec == 0, the operating system discards any unsent or
	//	// unacknowledged data.
	//	//
	//	// If sec > 0, the data is sent in the background as with sec < 0. On
	//	// some operating systems after sec seconds have elapsed any remaining
	//	// unsent data may be discarded.
	//	//SetLinger在仍有等待发送或确认的数据的连接上 设置“Close(关闭)”的行为。
	//	//如果sec <0（默认值），则操作系统完成在后台发送数据。也即使调用close()方法后仍然在后台发送遗留的要发送的数据。
	//	//如果sec == 0，则操作系统将丢弃所有未发送或未确认的数据。
	//	//如果sec> 0，则与sec <0一样在后台发送数据。在某些操作系统上，经过sec秒后，所有剩余的未发送数据都可能会被丢弃。
	//	// 这个跟sec <0（默认值）的区别在于在某些系统上面会经过sec秒后所有剩余的未发送数据都可能会被丢弃。但是sec <0（默认值）不会丢弃这些数据，他会直到数据发送完毕才结束！
	//	//Linger：延迟的意思
	//	e=tcpConn.SetLinger(0)
	//	check_err_net(e)
	//
	//	// SetNoDelay controls whether the operating system should delay
	//	// packet transmission in hopes of sending fewer packets (Nagle's
	//	// algorithm).  The default is true (no delay), meaning that data is
	//	// sent as soon as possible after a Write.
	//	//SetNoDelay控制操作系统是否应延迟数据包传输，以希望发送更少的数据包（Nagle算法）。 默认值为true（无延迟），这意味着在写操作之后尽快发送数据。
	//	//Nagle算法可参考这个链接：https://baike.baidu.com/item/Nagle%E7%AE%97%E6%B3%95/5645172?fr=aladdin
	//	e=tcpConn.SetNoDelay(false)
	//	check_err_net(e)
	//
	//	// LocalAddr returns the local network address.
	//	// The Addr returned is shared by all invocations of LocalAddr, so
	//	// do not modify it.
	//	//LocalAddr返回本地网络地址。
	//	//返回的Addr由LocalAddr的所有调用共享，因此请勿修改它。
	//	fmt.Println("tcpConn.LocalAddr():",tcpConn.LocalAddr())
	//
	//	// RemoteAddr returns the remote network address.
	//	// The Addr returned is shared by all invocations of RemoteAddr, so
	//	// do not modify it.
	//	//RemoteAddr返回远程网络地址。
	//	//返回的Addr由RemoteAddr的所有调用共享，因此请勿对其进行修改。
	//	fmt.Println("tcpConn.RemoteAddr():",tcpConn.RemoteAddr())
	//
	//	// SetDeadline implements the Conn SetDeadline method.
	//	//SetDeadline实现Conn SetDeadline方法。注意不是设置整个连接生存的时间（从开始建立连接到结束连接所限定的一共花费的时间），而是设置的读写的io阻塞所花费的时间，
	//	// 他会同时设置SetReadDeadline和SetWriteDeadline的时间！具体请查看下面的Conn SetDeadline方法的文档说明！
	//	//SetDeadline设置与连接关联的读写期限。这等效于调用SetReadDeadline和SetWriteDeadline。
	//	//截止期限是一个绝对时间，在该绝对时间之后，I / O操作将因超时（请参阅错误类型）而不是阻塞而失败。截止日期适用于所有将来和未决的I / O，而不仅仅是紧接在其后的读取或写入调用。超过期限后，可以
	//	//通过设置将来的期限来刷新连接。
	//	//空闲超时可以通过在成功进行Read或Write调用后重复延长截止期限来实现。
	//	//t的值为零表示I / O操作不会超时。
	//	//请注意，如果TCP连接启用了保持活动状态（除非被Dialer.KeepAlive或ListenConfig.KeepAlive覆盖，否则这是默认设置），则保持活动失败还可能返回超时错误。在Unix系统上，可以使用error.Is（err，syscall.ETIMEDOUT）检
	//	//测到I / O上的保持活动失败。
	//	//var someTimeout time.Duration=3e9
	//
	//	//var someTimeout1 int64=3e9
	//	//time.Sleep(3e9)//这个值会用time.duration类型来存，但是并不是默认的类型转换，没有类型转换，而是从一开始就采用了time.duration类型来存。
	//	//time.Sleep(someTimeout1)//虽然 time.Duration是基于基类int64类型来创建的，但是并不是说明在传递参数的时候都可以这样做！go的类型是不会隐式转换的！
	//	//上面的3行已注释掉的代码是为了说明go中的类型是不会自动隐式转换的！他不是本部分的必须代码！time.Now().Add（）函数也是接收的time.duration类型的参数！
	//			//事实上我们这里设置了的话，其实等价于下面2行代码这样的设置
	//			//tcpConn.SetReadDeadline(time.Now().Add(someTimeout))
	//			//tcpConn.SetWriteDeadline(time.Now().Add(someTimeout))
	//	//e=tcpConn.SetDeadline(time.Now().Add(someTimeout))
	//	check_err_net(e)
	//
	//	// SetReadDeadline implements the Conn SetReadDeadline method.
	//	//SetReadDeadline实现Conn SetReadDeadline方法。设置连接读取的时间，也就是从开始接收对方发送过来的数据到接收完毕所限定的花费的时间，经过这么多时间后仍然接收不完数据的话则会把连接关闭！注意不是指建立连接的时间
	//	//SetReadDeadline设置将来的Read呼叫和任何当前阻止的Read呼叫的截止日期。
	//	//t的值为零表示读取不会超时。
	//	//e=tcpConn.SetReadDeadline(time.Now().Add(someTimeout))//这行代码不是必须的甚至是多余的！因为上面已经有了tcpConn.SetDeadline(time.Now().Add(someTimeout))
	//	check_err_net(e)
	//
	//	// SetWriteDeadline implements the Conn SetWriteDeadline method.
	//	//SetWriteDeadline实现Conn SetWriteDeadline方法。设置连接写入的时间，也就是从开始发送到发送完毕数据所限定的花费的时间，经过这么多时间后仍然发送不完数据的话则会把连接关闭！注意不是指建立连接的时间
	//	//SetWriteDeadline设置将来的Write调用和任何当前阻止的Write调用的截止日期。
	//	//即使写入超时，它也可能返回n> 0，这表明某些数据已成功写入。
	//	//t的值为零表示写不会超时。
	//	//e=tcpConn.SetWriteDeadline(time.Now().Add(someTimeout))//这行代码不是必须的甚至是多余的！因为上面已经有了tcpConn.SetDeadline(time.Now().Add(someTimeout))
	//	check_err_net(e)
	//
	//	i, err := tcpConn.Write([]byte("这是客户端发送给服务端的信息数据1111TCPCONN TEST"))
	//	check_err_net(err)
	//
	//	fmt.Println("客户端发送给服务端的字节数为：",i)
	//	fmt.Println("客户端发送给服务端的字节数据为：","这是客户端发送给服务端的信息数据1111TCPCONN TEST")
	//	fmt.Println("客户端发送完毕！！")
	//
	//	rb := make([]byte, 100)
	//	//事实上如果你不设置上面的deadline的话，这里会一直阻塞的！上面的 write（）则不会如此！
	//	//因为write是往内核写入要发送的数据，然后就不管了，我们上层的代码不关心内核是否和服务器连接上了，也不关心是否发送成功了，
	//	//而read()方法则完全不同，这个方法一定需要服务器端给我们发送数据，我们才可以获取到数据，但是，在接收到服务器数据之前，我们是不是要
	//	//先创建与服务器的连接，然后才能读取服务器响应的数据！所以，你可以从上面我的分析可以看出为什么tcpConn.Read(rb)会阻塞当前g程而tcpConn.Write()
	//	//不会阻塞代码了！
	//	i, err = tcpConn.Read(rb)
	//	check_err_net(err)
	//
	//	fmt.Println()
	//	fmt.Println("客户端接收到服务器响应回来的字节数为：",i)
	//	fmt.Println("客户端接收到服务器响应回来的字节数据为：",rb)
	//	fmt.Println("客户端接收到服务器响应回来的字节转字符串数据为：",string(rb))
	//	fmt.Println("客户端接收完毕！！")
	//}
	//TestTCPConnSpecificMethods()
	//
	////运行须知：
	////1. go run ListenTCP.go
	////2. go run ListenUnix.go,为什么要运行这个东西，因为我们上面讨论DialUnix()函数时候的代码没注释掉，当然我也不打算注释掉！
	////3. go run compress_zlib.go
	//
	//
	////结果1
	////客户端输出结果：（服务端的输出结果请到服务段文件 ListenTCP.go中去查看）
	////	tcpConn.LocalAddr(): 127.0.0.1:52169，这个端口是随意的！每次执行都不一样！但是下面的999是一定的！
	////	tcpConn.RemoteAddr(): 127.0.0.1:9999
	////	客户端发送给服务端的字节数为： 64
	////	客户端发送给服务端的字节数据为： 这是客户端发送给服务端的信息数据1111TCPCONN TEST
	////	客户端发送完毕！！
	////
	////	客户端接收到服务器响应回来的字节数为： 49
	////	客户端接收到服务器响应回来的字节数据为： [232 191 153 230 152 175 116 99 112 230 156 141 229 138 161 229 153 168 229 147 141 229 186 148 231 187 153 229 174 162 230 136 183 231 171 175 231 154 132 229 134 133 229 174 185 50 50 50 50 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	////	客户端接收到服务器响应回来的字节转字符串数据为： 这是tcp服务器响应给客户端的内容2222
	////	客户端接收完毕！！
	//
	////运行完须知，因为每次运行完DialUnix()函数后都会在项目目录或者子目录（取决于你用什么编辑器运行，goland和cmd运行的产生位置不一样），
	////所以每次运行完后都要手动删除laddrUnixSocket文件，如果你没有看到这个文件，那么请自行刷新一下项目目录（goland中鼠标放到项目目录文件夹，右键
	//// ，最后哪里有个刷新的图标就是了！），一定要先删除了laddrUnixSocket文件然后才能再次运行！当然你也可以自己写个删除函数！我这里就不写了！
	//
	//
	////结果2
	////如果注释掉上面的SetDeadline（），SetReadDeadline(),SetWriteDeadline() ，同时服务器端的ListenTCP.go文件放开//time.Sleep(44e9)这行代
	////码的注释的话，则此时客户端和服务器段的输出结果跟上面的额那个结果是一样的！只是客户端和服务器端都会同时延迟44s的时间！这44s是由于服务端延迟导致的！
	//
	//
	//
	////结果3
	////但是如果没注释掉上面的SetDeadline（），SetReadDeadline(),SetWriteDeadline()方法，同时服务器端的ListenTCP.go文件放开//time.Sleep(44e9)这
	//// 行代码的注释的话，则此时会输出：
	////	tcpConn.LocalAddr(): 127.0.0.1:52275
	////	tcpConn.RemoteAddr(): 127.0.0.1:9999
	////	客户端发送给服务端的字节数为： 64
	////	客户端发送给服务端的字节数据为： 这是客户端发送给服务端的信息数据1111TCPCONN TEST
	////	客户端发送完毕！！
	////	出错了，错误信息为： read tcp4 127.0.0.1:52275->127.0.0.1:9999: i/o timeout
	////	panic: read tcp4 127.0.0.1:52275->127.0.0.1:9999: i/o timeout
	////
	////	goroutine 1 [running]:
	////	main.check_err_net(0x539900, 0xc0000be0f0)
	////		C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:1378 +0xdd
	////	main.main.func2()
	////		C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:1333 +0x797
	////	main.main()
	////		C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:1341 +0x27e8
	//
	//
	////从上面的tcp连接，我们测试了属于tcp连接下的方法，在unixconn和udpconn中都是类似的方法的！我们既然在这里讲到了，那么在讲到其他对象时候我就不再举例了！

	//为了避免上面的代码带来的影响和繁琐操作，我只好将上面代码先注释掉！其实每次研究一个模块，我都希望能注释掉上面的代码（从1405行到main函数开始的地方），
	// 或者分割成一个额外的.go文件来单独运行，但是我又不像这样做，如果你在测试学习时候可以将每个api相关的代码来分割到别的文件中去里测试！下同

	//fmt.Println("----------net包下的函数之.DialUDP(),发送udp数据 或者 接收udp数据--------------")
	//
	////其实这个udp发送数据在很多情况下跟上面讲到的tcp发送和接收数据是类似的，对于上面已经讲到的方法我们不再累叙，我们只会讲解有别于tcp对象下的方法和属性！
	//
	//// DialUDP acts like Dial for UDP networks.
	////
	//// The network must be a UDP network name; see func Dial for details.
	////
	//// If laddr is nil, a local address is automatically chosen.
	//// If the IP field of raddr is nil or an unspecified IP address, the
	//// local system is assumed.
	////DialUDP的作用类似于UDP网络的Dial。
	////该网络必须是UDP网络名称； 有关详细信息，请参见func Dial。
	////如果laddr为nil，则会自动选择一个本地地址。
	////如果raddr的IP字段为nil或未指定IP地址，则使用本地系统。
	////跟tcp连接类似，这个也是区分"udp", "udp4", "udp6"
	//
	////这个方法的参数对象文档如下：
	//		//// BUG(mikio): On NaCl and Plan 9, the ReadMsgUDP and
	//		//// WriteMsgUDP methods of UDPConn are not implemented.
	//		//
	//		//// BUG(mikio): On Windows, the File method of UDPConn is not
	//		//// implemented.
	//		//
	//		//// BUG(mikio): On NaCl, the ListenMulticastUDP function is not
	//		//// implemented.
	//		//
	//		//// BUG(mikio): On JS, methods and functions related to UDPConn are not
	//		//// implemented.
	//		//
	//		//// UDPAddr represents the address of a UDP end point.
	//		////BUG（mikio）：在NaCl和计划9上，未实现UDPConn的ReadMsgUDP和WriteMsgUDP方法。
	//		////BUG（mikio）：在Windows上，未实现UDPConn的File方法。
	//		////BUG（mikio）：在NaCl上，未实现ListenMulticastUDP函数。
	//		////BUG（mikio）：在JS上，未实现与UDPConn相关的方法和功能。
	//		////UDPAddr表示UDP端点的地址。
	//		//type UDPAddr struct {
	//		//	IP   IP
	//		//	Port int
	//		//	Zone string // IPv6 scoped addressing zone
	//		//}
	//
	////这个方法的返回值对象文档如下：
	//		//// UDPConn is the implementation of the Conn and PacketConn interfaces
	//		//// for UDP network connections.
	//		//// UDPConn是用于UDP网络连接的Conn和PacketConn接口的实现。
	//		//type UDPConn struct {
	//		//	conn
	//		//}
	//
	//TestUDPConnSpecificMethods:=func () {
	//	//注意了，这个函数写的是udp的客户端代码，服务段的代码请自行到ListenUdp.go中去查看。
	//	//如果你已经看了listenUnix的代码，那么你应该知道，要先启动服务端的代码，然后再独自运行当前文件中客户端中的代码！注意文件要单独运行！
	//
	//
	//	// ResolveUDPAddr returns an address of UDP end point.
	//	//
	//	// The network must be a UDP network name.
	//	//
	//	// If the host in the address parameter is not a literal IP address or
	//	// the port is not a literal port number, ResolveUDPAddr resolves the
	//	// address to an address of UDP end point.
	//	// Otherwise, it parses the address as a pair of literal IP address
	//	// and port number.
	//	// The address parameter can use a host name, but this is not
	//	// recommended, because it will return at most one of the host name's
	//	// IP addresses.
	//	//
	//	// See func Dial for a description of the network and address
	//	// parameters.
	//	//ResolveUDPAddr返回UDP端点的地址。
	//	//该网络必须是UDP网络名称。
	//	//如果address参数中的主机不是文字IP地址或端口不是文字端口号，则ResolveUDPAddr将该地址解析为UDP端点的地址，否则会将地址解析为一对文字IP地址和端口 数。
	//	//address参数可以使用主机名，但是不建议这样做，因为它最多返回主机名的IP地址之一。
	//	//有关网络和地址参数的说明，请参见func Dial。
	//	//请注意上面的文档，仔细看！因为我不会什么条件都进行测试的！这样的话我会需要写更多的代码！
	//	ra, err := net.ResolveUDPAddr("udp4", "127.0.0.1:8888")//采用8888端口发送数据给服务器
	//	check_err_net(err)
	//	la, err := net.ResolveUDPAddr("udp4", "127.0.0.1:7777")//采用7777端口从服务器接收数据，不可给8888
	//	check_err_net(err)
	//	//上面的la千万不要绑定8888端口，否则会报多个程序绑定同一个端口的错误：
	//	//	出错了，错误信息为： dial udp4 127.0.0.1:8888->127.0.0.1:8888: bind: Only one usage of each socket address (protocol/network address/port) is normally permitted.
	//	//	panic: dial udp4 127.0.0.1:8888->127.0.0.1:8888: bind: Only one usage of each socket address (protocol/network address/port) is normally permitted.
	//	//
	//	//	goroutine 1 [running]:
	//	//	main.check_err_net(0x525b20, 0xc000088000)
	//	//		C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:1586 +0xdd
	//	//	main.main.func1()
	//	//		C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:1498 +0x137
	//	//	main.main()
	//	//		C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:1574 +0x90
	//
	//	//本地地址直接给nil即可，这样的话他就会每次执行代码时候，默认就是本机的随机一个端口！当然如果你想指定一个特定的端口的话也是可以的，但是前提是这个端口是没被其他程序正在使用的
	//	//端口就行了！这个跟tcp是不一样的！tcp可以给nil,这里也可以给nil！但是却需要在对方中采用udpConn.ReadFrom（）方法获取这个地址la,具体的就不展示了
	//	udpConn, err := net.DialUDP("udp4", la, ra)
	//	check_err_net(err)
	//	defer udpConn.Close()
	//	//下面的几个方法跟tcp中的几乎一模一样的！不再累叙，仅仅列出来而已
	//	//	udpConn.SetDeadline()
	//	//	udpConn.SetReadDeadline()
	//	//	udpConn.SetWriteDeadline()
	//	//	udpConn.SyscallConn()
	//	//	udpConn.LocalAddr()
	//	//	udpConn.RemoteAddr()
	//
	//	//// File returns a copy of the underlying os.File.
	//	//// It is the caller's responsibility to close f when finished.
	//	//// Closing c does not affect f, and closing f does not affect c.
	//	////
	//	//// The returned os.File's file descriptor is different from the connection's.
	//	//// Attempting to change properties of the original using this duplicate
	//	//// may or may not have the desired effect.
	//	////File返回基础os.File的副本。
	//	////完成后，关闭f是调用者的责任。
	//	////关闭c不会影响f，关闭f不会影响c。
	//	////返回的os.File的文件描述符与连接的文件描述符不同。
	//	////尝试使用此副本来更改原件的属性可能会或可能不会产生预期的效果。
	//	//f, err := udpConn.File()
	//	//fmt.Println("f:",f)
	//	//check_err_net(err)
	//
	//	//事实上上面的方法并没有在windows中被实现，所以会报错的！报错如下：
	//	//	f: <nil>
	//	//	出错了，错误信息为： file udp4 127.0.0.1:59459->127.0.0.1:8888: not supported by windows
	//	//	panic: file udp4 127.0.0.1:59459->127.0.0.1:8888: not supported by windows
	//	//
	//	//	goroutine 1 [running]:
	//	//	main.check_err_net(0x525980, 0xc000088000)
	//	//		C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:1549 +0xdd
	//	//	main.main.func1()
	//	//		C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:1519 +0x246
	//	//	main.main()
	//	//		C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:1537 +0x90
	//
	//	//所以我们想要正常执行的话则要注释掉上面的 udpConn.File()方法！
	//
	//	write_i, err := udpConn.Write([]byte("这是客户端发送给服务器端的信息111udpConn test"))
	//	check_err_net(err)
	//
	//	fmt.Println("客户端发送给服务器端的字节数为：",write_i)
	//	fmt.Println("客户端发送给服务器端的字节转字符串数据为：","这是客户端发送给服务器端的信息111udpConn test")
	//	fmt.Println("客户端发送完毕！！")
	//
	//	fmt.Println()
	//	fmt.Println("正在接收服务器的响应数据。。。")
	//	Read_by:=make([]byte,100)
	//
	//	//// ReadMsgUDP reads a message from c, copying the payload into b and
	//	//// the associated out-of-band data into oob. It returns the number of
	//	//// bytes copied into b, the number of bytes copied into oob, the flags
	//	//// that were set on the message and the source address of the message.
	//	////
	//	//// The packages golang.org/x/net/ipv4 and golang.org/x/net/ipv6 can be
	//	//// used to manipulate IP-level socket options in oob.
	//	////ReadMsgUDP从c读取一条消息，将有效负载复制到b，并将关联的带外数据复制到oob。 它返回复制到b中的字节数，复制到oob中的字节数，
	//	////在消息上设置的标志以及消息的源地址。
	//	////软件包golang.org/x/net/ipv4和golang.org/x/net/ipv6可用于处理oob中的IP级套接字选项。
	//	//Read_by1:=make([]byte,100)
	//	//fmt.Println(udpConn.ReadMsgUDP(Read_by, Read_by1))
	//	////如果采用上面的读取方式则输出：
	//	////	49 0 0 127.0.0.1:8888 <nil>
	//	////目前还不大了解其中的结果数据是什么意思，先搁置！
	//
	//	// ReadFrom reads a packet from the connection,
	//	// copying the payload into p. It returns the number of
	//	// bytes copied into p and the return address that
	//	// was on the packet.
	//	// It returns the number of bytes read (0 <= n <= len(p))
	//	// and any error encountered. Callers should always process
	//	// the n > 0 bytes returned before considering the error err.
	//	// ReadFrom can be made to time out and return
	//	// an Error with Timeout() == true after a fixed time limit;
	//	// see SetDeadline and SetReadDeadline.
	//	//ReadFrom从连接读取数据包，将有效负载复制到p中。 它返回复制到p中的字节数以及该数据包上的返回地址。
	//	//它返回读取的字节数（0 <= n <= len（p））和遇到的任何错误。 在考虑错误err之前，调用者应始终处理返回的n> 0个字节。
	//	//可以使ReadFrom超时并在固定的时间限制后使用Timeout（）== true返回错误； 请参见SetDeadline和SetReadDeadline。
	//
	//	read_i, addr,err := udpConn.ReadFrom(Read_by)
	//	//如果你需要获取对方的地址，那么就可以采用上面的方式，如果不需要则采用下面的方式都是可以的！
	//	//read_i, err := udpConn.Read(Read_by)
	//
	//	//这个地址是显示对方的地址！事实上我们应该在服务器端使用udpConn.ReadFrom（）的，以便客户返回数据给我们时候
	//	// 可以往这个地址上面进行返回！！这里我们在这里使用了，就懒得折腾了！
	//	fmt.Println("addr:",addr)
	//	fmt.Println("客户端接收到服务器端响应回来的字节数为：",read_i)
	//	fmt.Println("客户端接收到服务器端响应回来的字节数据为：",Read_by)
	//	fmt.Println("客户端接收到服务器端响应回来的字节转字符串数据为：",string(Read_by))
	//	fmt.Println("客户端接收完毕！")
	//
	//}
	//
	//TestUDPConnSpecificMethods()
	//
	////先单独运行ListenUdp.go中的代码，然后在运行本文件的代码！
	////输出：
	////	客户端发送给服务器端的字节数为： 60
	////	客户端发送给服务器端的字节转字符串数据为： 这是客户端发送给服务器端的信息111udpConn test
	////	客户端发送完毕！！
	////
	////	正在接收服务器的响应数据。。。
	////	addr: 127.0.0.1:8888，这个地址是对方的地址，注意了！
	////	客户端接收到服务器端响应回来的字节数为： 49
	////	客户端接收到服务器端响应回来的字节数据为： [232 191 153 230 152 175 117 100 112 230 156 141 229 138 161 229 153 168 229 143 145
	////	233 128 129 231 187 153 229 174 162 230 136 183 231 171 175 231 154 132 230 149 176 230 141 174 50 50 50 50 0 0 0 0 0 0 0 0 0
	////	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	////	客户端接收到服务器端响应回来的字节转字符串数据为： 这是udp服务器发送给客户端的数据2222
	////	客户端接收完毕！

	//fmt.Println("----------net包下的函数之.DialIP(),构造原始套接发送ip数据 或者 接收ip数据--------------")
	////请百度了解原始套接字的含义！
	////同理，这个api跟上面的unix套接字,tcp以及udp的连接api是同理的用法！
	//
	//
	//// DialIP acts like Dial for IP networks.
	////
	//// The network must be an IP network name; see func Dial for details.
	////
	//// If laddr is nil, a local address is automatically chosen.
	//// If the IP field of raddr is nil or an unspecified IP address, the
	//// local system is assumed.
	////DialIP的作用类似于IP网络的拨号函数Dial（）。
	////该网络必须是IP网络名称； 有关详细信息，请参见func Dial。
	////如果laddr为nil，则会自动选择一个本地地址。
	////如果raddr的IP字段为nil或未指定IP地址，则使用本地系统。
	//
	////请参见func Dial对ip的讲解:
	//// For IP networks, the network must be "ip", "ip4" or "ip6" followed
	//// by a colon and a literal protocol number or a protocol name, and
	//// the address has the form "host". The host must be a literal IP
	//// address or a literal IPv6 address with zone.
	//// It depends on each operating system how the operating system
	//// behaves with a non-well known protocol number such as "0" or "255".
	////对于IP网络，网络必须为"ip", "ip4" or "ip6"，后跟冒号和协议号数字字面量或协议名称，比如“ip:tcp”或者“ip:6”，并且地址的格式为“主机host”。 主机host必须是文本IP地址或带区域zone的文本IPv6地址。
	////取决于每个操作系统，操作系统如何以不知名的协议编号（例如“ 0”或“ 255”）运行。
	//
	//// Examples:
	////	Dial("ip4:1", "192.0.2.1")
	////	Dial("ip6:ipv6-icmp", "2001:db8::1")
	////	Dial("ip6:58", "fe80::1%lo0")
	//
	//// For TCP, UDP and IP networks, if the host is empty or a literal
	//// unspecified IP address, as in ":80", "0.0.0.0:80" or "[::]:80" for
	//// TCP and UDP, "", "0.0.0.0" or "::" for IP, the local system is
	//// assumed.
	////对于TCP，UDP和IP网络，如果主机为空或文字未指定的IP地址，例如对于TCP和UDP，为":80"，"0.0.0.0:80" 或"[::]:80"，对于IP为"", "0.0.0.0" or "::"，则假定为本地系统。

	//// BUG(mikio): On every POSIX platform, reads from the "ip4" network
	//// using the ReadFrom or ReadFromIP method might not return a complete
	//// IPv4 packet, including its header, even if there is space
	//// available. This can occur even in cases where Read or ReadMsgIP
	//// could return a complete packet. For this reason, it is recommended
	//// that you do not use these methods if it is important to receive a
	//// full packet.
	////
	//// The Go 1 compatibility guidelines make it impossible for us to
	//// change the behavior of these methods; use Read or ReadMsgIP
	//// instead.
	//
	//// BUG(mikio): On JS, NaCl and Plan 9, methods and functions related
	//// to IPConn are not implemented.
	//
	//// BUG(mikio): On Windows, the File method of IPConn is not
	//// implemented.
	//
	//// IPAddr represents the address of an IP end point.
	////BUG（mikio）：在每个POSIX平台上，即使有可用空间，使用ReadFrom或ReadFromIP方法从“ ip4”网络中进行的读取也可能不会返回完整的IPv4数据包，包括其报头。
	////	即使在Read或ReadMsgIP可以返回完整数据包的情况下，也可能发生这种情况。 因此，如果重要的是接收完整的数据包，建议您不要使用这些方法。
	////Go 1兼容性准则使我们无法更改这些方法的行为。 请改用Read或ReadMsgIP。
	////BUG（mikio）：在JS，NaCl和计划9上，未实现与IPConn相关的方法和功能。
	////BUG（mikio）：在Windows上，未实现IPConn的File方法。
	////IPAddr表示IP端点的地址。
	//type IPAddr struct {
	//	IP   IP
	//	Zone string // IPv6 scoped addressing zone// IPv6范围内的寻址区
	//}

	//// An IP is a single IP address, a slice of bytes.
	//// Functions in this package accept either 4-byte (IPv4)
	//// or 16-byte (IPv6) slices as input.
	////
	//// Note that in this documentation, referring to an
	//// IP address as an IPv4 address or an IPv6 address
	//// is a semantic property of the address, not just the
	//// length of the byte slice: a 16-byte slice can still
	//// be an IPv4 address.
	////IP是单个IP地址，是字节的一部分。
	////此程序包中的功能接受4字节（IPv4）或16字节（IPv6）切片作为输入。
	////请注意，在本文档中，将IP地址称为IPv4地址或IPv6地址是该地址的语义属性，而不仅仅是字节片的长度：16字节的片仍可以是IPv4地址。
	//type IP []byte

	//IP这个类实现了string()方法，下面是这个方法的文档：
	// String returns the string form of the IP address ip.
	// It returns one of 4 forms:
	//   - "<nil>", if ip has length 0
	//   - dotted decimal ("192.0.2.1"), if ip is an IPv4 or IP4-mapped IPv6 address
	//   - IPv6 ("2001:db8::1"), if ip is a valid IPv6 address
	//   - the hexadecimal form of ip, without punctuation, if no other cases apply
	// String返回IP地址ip的字符串形式。
	//返回以下4种形式之一：
	//	 -"<nil>"，如果ip的长度为0
	//	 -点分割的十进制数（"192.0.2.1"），如果ip是IPv4或IP4映射的IPv6地址
	//	 -IPv6（"2001:db8::1"），如果ip是有效的IPv6地址
	//	 -ip的十六进制形式，如果没有其他情况，则不带标点

	//// IPConn is the implementation of the Conn and PacketConn interfaces
	//// for IP network connections.
	////IPConn是IP网络连接的Conn和PacketConn接口的实现。
	//type IPConn struct {
	//	conn
	//}

	////// IPv4 returns the IP address (in 16-byte form) of the
	////// IPv4 address a.b.c.d.
	//////IPv4返回IPv4地址a.b.c.d的IP地址（16字节格式）。
	////
	////// To4 converts the IPv4 address ip to a 4-byte representation.
	////// If ip is not an IPv4 address, To4 returns nil.
	//////To4将IPv4地址ip转换为4字节表示形式。
	//////如果ip不是IPv4地址，则To4返回nil。
	//IP16:=net.IPv4(192,168,63,132)
	//IP4:=IP16.To4()
	//raddr := &net.IPAddr{IP:IP4}
	//fmt.Printf("net.IPv4(192,168,63,132)：%v,字节为：%v,长度为：%v\n",IP16,[]byte(IP16),len(IP16))
	//fmt.Printf("net.IPv4(192,168,63,132).To4()：%v,字节为：%v,长度为：%v\n",IP4,[]byte(IP4),len(IP4))

	// ResolveIPAddr returns an address of IP end point.
	//
	// The network must be an IP network name.
	//
	// If the host in the address parameter is not a literal IP address,
	// ResolveIPAddr resolves the address to an address of IP end point.
	// Otherwise, it parses the address as a literal IP address.
	// The address parameter can use a host name, but this is not
	// recommended, because it will return at most one of the host name's
	// IP addresses.
	//
	// See func Dial for a description of the network and address
	// parameters.
	//ResolveIPAddr返回IP端点的地址。
	//该网络必须是IP网络名称。
	//如果address参数中的主机不是字面量IP地址，则ResolveIPAddr将该地址解析为IP端点地址。 否则，它将地址解析为原义IP地址。
	//address参数可以使用主机名host name，但是不建议这样做，因为它最多返回主机名的IP地址之一。
	//有关网络和地址参数的说明，请参见func Dial。

	////raddr, e := net.ResolveIPAddr("ip", "39.156.69.79")//这个地址是百度的地址
	////raddr, e := net.ResolveIPAddr("ip", "192.168.63.132")//这个是我本地的虚拟机的地址
	////check_err_net(e)
	////
	////laddr, e := net.ResolveIPAddr("ip", "127.0.0.1")
	////check_err_net(e)
	//raddr, err := net.ResolveIPAddr("ip", "39.156.69.79")
	////
	////raddr := &net.IPAddr{IP: net.IPv4(127,0,0,1).To4()}
	//check_err_net(err)
	////下面写"ip4:udp"或者"ip4:17"都一样的，udp的ip协议号为17，17代表udp,6代表tcp,但是在这里不知道为什么进行不了tcp协议，先搁置。
	////icmp为1！必须写上协议号或者协议名，否则会报错！但是"ip4:udp"中可以写"ip4"，也可以写"ip".
	////更多的ip协议号请参考：https://wenku.baidu.com/view/d2c7404733d4b14e852468e1.html或者
	////https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml
	//ipUdpConn, err := net.DialIP("ip4:tcp", nil, raddr)
	////注意千万不能像下面那样本地地址使用&net.IPAddr{IP: net.IPv4(127, 0, 0, 1)}，因为127.0.0.1这个地址是本地使用的！但是如果你需要访问外网的话，则本地的地址应该是外网地址才对的！
	////ipUdpConn, err := net.DialIP("ip:udp", &net.IPAddr{IP: net.IPv4(127, 0, 0, 1)}, raddr)
	////ipTcpConn, err := net.DialIP("ip:", laddr, raddr)
	//check_err_net(err)
	//defer ipUdpConn.Close()
	//fmt.Println("ipTcpConn：",ipUdpConn)
	//fmt.Println("ipUdpConn.LocalAddr()：",ipUdpConn.LocalAddr())
	//fmt.Println("ipUdpConn.RemoteAddr()：",ipUdpConn.RemoteAddr())
	//
	////访问百度的地址则输出：
	////	ipTcpConn： &{{0xc000090000}}
	////	ipUdpConn.LocalAddr()： 192.168.1.102，是我的主机上面跟外网通信的地址
	////	ipUdpConn.RemoteAddr()： 39.156.69.79，百度地址
	//
	////访问我本地的一个虚拟机则输出：
	////	ipTcpConn： &{{0xc000090000}}
	////	ipUdpConn.LocalAddr()： 192.168.63.1，是我的主机上面跟虚拟机通信的地址，局域网地址！
	////	ipUdpConn.RemoteAddr()： 192.168.63.132，虚拟机地址
	////通过对比上面的2个ipUdpConn.LocalAddr()，你会发现他们不同！
	//
	////访问我本地的主机自己则输出：
	////	ipTcpConn： &{{0xc000090000}}
	////	ipUdpConn.LocalAddr()： 127.0.0.1
	////	ipUdpConn.RemoteAddr()： 127.0.0.1
	//
	////下面进行消息的发送
	//i, e := ipUdpConn.Write([]byte("这是ip:udp通信客户端发送的消息"))
	//check_err_net(e)
	//fmt.Println("发送的字节数为：",i)
	//fmt.Println("发送的字节转字符串数据为：","这是ip:udp通信客户端发送的消息")
	//fmt.Println("发送的字节数据为：",[]byte("这是ip:udp通信客户端发送的消息"))

	//TestIPConnRemoteName:=func() {
	//	//if !testableNetwork("ip:tcp") {
	//	//	t.Skip("ip:tcp test")
	//	//}
	//
	//	//raddr := &net.IPAddr{IP: net.IPv4(39,156,69,79).To4()}
	//
	//	//然而，并不能发送给自己！即使是127.0.0.1也不行
	//	raddr := &net.IPAddr{IP: net.IPv4(192,168,1,102).To4()}
	//	c, err := net.DialIP("ip:1", nil, raddr)//本地地址为nil即可，除非你本地机子有多个ip能通网，此时可以提供自定义选择本机ip
	//	check_err_net(err)
	//	defer c.Close()
	//	if !reflect.DeepEqual(raddr, c.RemoteAddr()) {
	//		fmt.Printf("got %#v; want %#v", c.RemoteAddr(), raddr)
	//	}
	//
	//
	//	//i, err := c.WriteTo([]byte("abcdefg!!"),raddr)
	//	//i, err := c.WriteToIP([]byte("abcdefg!!"),raddr)
	//	//因为c已经被绑定了拨号地址并且已经创建了准备发送信息的对象，此时再像上面这样使用WriteTo或者WriteToIP都是会报错的，如下：
	//	//use of WriteTo with pre-connected connection（握手完成或使用与对等方的关联）
	//	//这里并不能发送单纯的协议内容，还需要其他的协议头等等，所以这是错误的用法展示！总之要包含完整的协议！
	//	i, err := c.Write([]byte("abcdefg!!"))
	//
	//	check_err_net(err)
	//	fmt.Println("发送了字节数：",i)
	//	fmt.Println("发送了字节数据：",[]byte("abcdefg!!"))
	//	fmt.Println("发送了字节转字符串数据：","abcdefg!!")
	//
	//
	//
	//}
	//
	//TestIPConnRemoteName()
	//
	////下面的这些全部都是输入格式错误的参数的测试，也就是说下面的测试都会报错才对！
	//TestDialListenIPArgs:=func() {
	//	type test struct {
	//		argLists   [][2]string
	//		shouldFail bool
	//	}
	//	tests := []test{
	//		{
	//			argLists: [][2]string{
	//				{"ip", "127.0.0.1"},
	//				{"ip:", "127.0.0.1"},
	//				{"ip::", "127.0.0.1"},
	//				{"ip", "::1"},
	//				{"ip:", "::1"},
	//				{"ip::", "::1"},
	//				{"ip4", "127.0.0.1"},
	//				{"ip4:", "127.0.0.1"},
	//				{"ip4::", "127.0.0.1"},
	//				{"ip6", "::1"},
	//				{"ip6:", "::1"},
	//				{"ip6::", "::1"},
	//			},
	//			shouldFail: true,
	//		},
	//	}
	//	if true {
	//		priv := test{shouldFail: false}
	//		for _, tt := range []struct {
	//			network, address string
	//			args             [2]string
	//		}{
	//			{"ip4:47", "127.0.0.1", [2]string{"ip4:47", "127.0.0.1"}},
	//			{"ip6:47", "::1", [2]string{"ip6:47", "::1"}},
	//		} {
	//			c, err := net.ListenPacket(tt.network, tt.address)
	//			if err != nil {
	//				continue
	//			}
	//			c.Close()
	//			priv.argLists = append(priv.argLists, tt.args)
	//		}
	//		if len(priv.argLists) > 0 {
	//			tests = append(tests, priv)
	//		}
	//	}
	//
	//	for _, tt := range tests {
	//		for _, args := range tt.argLists {
	//			_, err := net.Dial(args[0], args[1])
	//			if tt.shouldFail != (err != nil) {
	//				fmt.Printf("Dial(%q, %q) = %v; want (err != nil) is %t", args[0], args[1], err, tt.shouldFail)
	//			}
	//			_, err = net.ListenPacket(args[0], args[1])
	//			if tt.shouldFail != (err != nil) {
	//				fmt.Printf("ListenPacket(%q, %q) = %v; want (err != nil) is %t", args[0], args[1], err, tt.shouldFail)
	//			}
	//			a, err := net.ResolveIPAddr("ip", args[1])
	//			if err != nil {
	//				fmt.Printf("ResolveIPAddr(\"ip\", %q) = %v", args[1], err)
	//				continue
	//			}
	//			_, err = net.DialIP(args[0], nil, a)
	//			if tt.shouldFail != (err != nil) {
	//				fmt.Printf("DialIP(%q, %v) = %v; want (err != nil) is %t", args[0], a, err, tt.shouldFail)
	//			}
	//			_, err = net.ListenIP(args[0], a)
	//			if tt.shouldFail != (err != nil) {
	//				fmt.Printf("ListenIP(%q, %v) = %v; want (err != nil) is %t", args[0], a, err, tt.shouldFail)
	//			}
	//		}
	//	}
	//}
	//TestDialListenIPArgs()
	//
	//
	//TestICMPconn:= func() {
	//下面icmp的测试转载出处，有更改：http://blog.csdn.net/gophers/article/details/21481447
	//	var (
	//		icmp  ICMP
	//		// ParseIP parses s as an IP address, returning the result.
	//		// The string s can be in dotted decimal ("192.0.2.1")
	//		// or IPv6 ("2001:db8::68") form.
	//		// If s is not a valid textual representation of an IP address,
	//		// ParseIP returns nil.
	//		//ParseIP将s解析为IP地址，并返回结果。
	//		//字符串s可以是点分十进制（“ 192.0.2.1”）或IPv6（“ 2001：db8 :: 68”）形式。
	//		//如果s不是IP地址的有效文本表示形式，则ParseIP返回nil。
	//		laddr net.IPAddr = net.IPAddr{IP: net.ParseIP("192.168.1.102")}  //***IP地址改成你自己的网段***
	//		raddr net.IPAddr = net.IPAddr{IP: net.ParseIP("192.168.1.1")}
	//	)
	//	//如果你要使用网络层的其他协议还可以设置成 ip:ospf、ip:arp 等
	//	conn, err := net.DialIP("ip4:icmp", &laddr, &raddr)
	//	if err != nil {
	//		fmt.Println(err.Error())
	//		return
	//	}
	//	defer conn.Close()
	//
	//	//开始填充数据包
	//	icmp.Type = 8 //8->echo message  0->reply message
	//	icmp.Code = 0
	//	icmp.Checksum = 0
	//	icmp.Identifier = 0
	//	icmp.SequenceNum = 0
	//
	//
	//	go func() {
	//		//不知道为什么无法read，先搁置
	//		readICMP(conn)
	//	}()
	//	var (
	//		buffer bytes.Buffer
	//	)
	//	//先在buffer中写入icmp数据报求去校验和
	//	binary.Write(&buffer, binary.BigEndian, icmp)
	//	icmp.Checksum = CheckSum(buffer.Bytes())
	//	//然后清空buffer并把求完校验和的icmp数据报写入其中准备发送
	//	buffer.Reset()
	//	binary.Write(&buffer, binary.BigEndian, icmp)
	//
	//	if _, err := conn.Write(buffer.Bytes()); err != nil {
	//		fmt.Println(err.Error())
	//		return
	//	}
	//	fmt.Printf("send icmp packet success!\n")
	//
	//	time.Sleep(3e9)
	//}
	//
	//TestICMPconn()
	//请自行抓包验证icmp的包是否发送以及是否得到了response
	////DialIp()还不是很懂，先搁置吧！

	////注释掉上面的代码
	//fmt.Println("----------net包下的函数之.Dial(DNS),发送DNS数据 或者 接收DNS数据--------------")
	////这个函数在另外一个文件里面，如果单独运行本文件会报undefined: DialDnsUdp
	//DialDnsUdp()

	//fmt.Println("----------net包下的函数之ip掩码mask--------------")
	//
	//
	//
	//// IPv4Mask returns the IP mask (in 4-byte form) of the
	//// IPv4 mask a.b.c.d.
	////IPv4Mask返回IPv4掩码a.b.c.d的IP掩码（4字节格式）。
	////mask := net.IPv4Mask(192, 168, 1, 102)//ip1
	//mask := net.IPv4Mask(255, 255, 255, 255)//ip2
	////mask := net.IPv4Mask(255, 0, 0, 0)//ip3
	////mask := net.IPv4Mask(255, 0, 0, 1)//ip4
	////mask := net.IPv4Mask(224, 0, 0, 1)//ip5
	////mask := net.IPv4Mask(224, 0, 0, 2)//ip6
	////mask := net.IPv4Mask(0, 0, 0, 0)//ip7
	////mask := net.IPv4Mask(0, 0, 0, 1)//ip8
	//fmt.Println("mask:",mask)//此时相当于执行mask.String()
	//fmt.Printf("mask的二进制切片:%b\n",mask)
	//i, e := strconv.ParseInt(mask.String(),16,64)
	//check_err_net(e)
	//fmt.Printf("mask的二进制:%b\n",i)
	//fmt.Println("mask转字节切片为:",[]byte(mask))
	//// String returns the hexadecimal form of m, with no punctuation.
	//// String返回m的十六进制形式，不带标点。
	//fmt.Println("mask.String():",mask.String())
	//// Size returns the number of leading ones and total bits in the mask.
	//// If the mask is not in the canonical form--ones followed by zeros--then
	//// Size returns 0, 0.
	////Size返回掩码中前导1的数量和总位数。
	////如果掩码不是规范形式-1后跟零-那么Size返回0，0。
	////规范形式：正确的掩码应该是255（255的二进制就是多个1）开头或者没有0开头才对的，比如0.0.0.0或者255.0.0.0或者255.255.0.0，总之最后一个255或者0后面不能再有1了，
	////比如255.255.0.1和0.0.0.1就不是正确的掩码
	//ones, bits := mask.Size()
	//fmt.Printf("mask ones:%v\nmask bits:%v\n",ones, bits )
	////ip1输出：
	////	mask: c0a80166
	////	mask的二进制切片:[11000000 10101000 1 1100110]
	////	mask的二进制:11000000101010000000000101100110
	////	mask转字节切片为: [192 168 1 102]
	////	mask.String(): c0a80166
	////	mask ones:0
	////	mask bits:0
	//
	//
	////ip2输出：
	////	mask: ffffffff
	////	mask的二进制切片:[11111111 11111111 11111111 11111111]
	////	mask的二进制:11111111111111111111111111111111
	////	mask转字节切片为: [255 255 255 255]
	////	mask.String(): ffffffff
	////	mask ones:32
	////	mask bits:32
	//
	//
	////ip3输出：
	////	mask: ff000000
	////	mask的二进制切片:[11111111 0 0 0]
	////	mask的二进制:11111111000000000000000000000000
	////	mask转字节切片为: [255 0 0 0]
	////	mask.String(): ff000000
	////	mask ones:8
	////	mask bits:32
	//
	//
	////ip4输出：
	////	mask: ff000001
	////	mask的二进制切片:[11111111 0 0 1]
	////	mask的二进制:11111111000000000000000000000001
	////	mask转字节切片为: [255 0 0 1]
	////	mask.String(): ff000001
	////	mask ones:0
	////	mask bits:0
	//
	//
	////ip5输出：
	////	mask: e0000001
	////	mask的二进制切片:[11100000 0 0 1]
	////	mask的二进制:11100000000000000000000000000001
	////	mask转字节切片为: [224 0 0 1]
	////	mask.String(): e0000001
	////	mask ones:0
	////	mask bits:0
	//
	//
	////ip6输出：
	////	mask: e0000002
	////	mask的二进制切片:[11100000 0 0 10]
	////	mask的二进制:11100000000000000000000000000010
	////	mask转字节切片为: [224 0 0 2]
	////	mask.String(): e0000002
	////	mask ones:0
	////	mask bits:0
	//
	//
	////ip7输出：
	////	mask: 00000000
	////	mask的二进制切片:[0 0 0 0]
	////	mask的二进制:0
	////	mask转字节切片为: [0 0 0 0]
	////	mask.String(): 00000000
	////	mask ones:0
	////	mask bits:32
	//
	//
	////ip8输出：
	////	mask: 00000001
	////	mask的二进制切片:[0 0 0 1]
	////	mask的二进制:1
	////	mask转字节切片为: [0 0 0 1]
	////	mask.String(): 00000001
	////	mask ones:0
	////	mask bits:0
	//
	//
	//
	//fmt.Println()
	//fmt.Println("----------net包下的函数之ip掩码mask111--------------")
	//// CIDRMask returns an IPMask consisting of `ones' 1 bits
	//// followed by 0s up to a total length of `bits' bits.
	//// For a mask of this form, CIDRMask is the inverse of IPMask.Size.
	////CIDRMask返回一个IPMask，该IPMask由“ ones” 1位和后跟0s组成，总长度为“ bits”位。
	////对于这种形式的蒙版，CIDRMask是IPMask.Size的逆。
	//
	////// An IP mask is an IP address.// IP掩码是IP地址。
	////type IPMask []byte
	//
	////// Well-known IPv4 addresses//知名的IPv4地址
	////var (
	////	IPv4bcast     = IPv4(255, 255, 255, 255) // limited broadcast//有限广播
	////	IPv4allsys    = IPv4(224, 0, 0, 1)       // all systems//所有系统
	////	IPv4allrouter = IPv4(224, 0, 0, 2)       // all routers//所有路由器
	////	IPv4zero      = IPv4(0, 0, 0, 0)         // all zeros//全零
	////)
	//
	////cidrMask := net.CIDRMask(0, 32)//Mask1
	////cidrMask := net.CIDRMask(8, 32)//Mask2
	////cidrMask := net.CIDRMask(16, 32)//Mask3
	////cidrMask := net.CIDRMask(24, 32)//Mask4
	////cidrMask := net.CIDRMask(32, 32)//Mask5
	//cidrMask := net.CIDRMask(31, 32)//Mask6，一般是8的倍数
	//
	//fmt.Println("cidrMask.String():",cidrMask.String())
	//
	//i, e = strconv.ParseInt(cidrMask.String(),16,64)
	//check_err_net(e)
	//fmt.Printf("mask的二进制:%b\n",i)
	//fmt.Println(cidrMask.Size())
	////Mask1输出：
	////	cidrMask.String(): 00000000
	////	mask的二进制:0
	////	0 32
	//
	////Mask2输出：
	////	cidrMask.String(): ff000000
	////	mask的二进制:11111111000000000000000000000000
	////	8 32
	//
	////Mask3输出：
	////	cidrMask.String(): ffff0000
	////	mask的二进制:11111111111111110000000000000000
	////	16 32
	//
	////Mask4输出：
	////cidrMask.String(): ffffff00
	////mask的二进制:11111111111111111111111100000000
	////24 32
	//
	////Mask5输出：
	////	cidrMask.String(): ffffffff
	////	mask的二进制:11111111111111111111111111111111
	////	32 32
	//
	////Mask6输出：
	////	cidrMask.String(): fffffffe
	////	mask的二进制:11111111111111111111111111111110
	////	31 32
	//
	//
	//
	//fmt.Println()
	//fmt.Println("----------net包下的函数之listen（）函数--------------")
	////下面写的是客户端请求的代码，如果要看listen()函数的用法请到main3/ListenTcp1.go中去看
	////同样如果运行的话则想要执行main3/ListenTcp1.go这个文件，然后在运行本文件！
	////listen（）函数其实和我们上面讲到的下面的函数是类似的，只是有些方法没有！
	////net.ListenUDP()
	////net.ListenIP()
	////net.ListenTCP()
	////net.ListenUnix()
	//
	//laddr, e := net.ResolveTCPAddr("tcp4", "127.0.0.1:9994")
	//raddr, e := net.ResolveTCPAddr("tcp4", "127.0.0.1:9996")
	////上面是本机地址的9994端口和9996端口通信
	//check_err_net(e)
	//tcpconn, e := net.DialTCP("tcp4", laddr, raddr)
	//check_err_net(e)
	//writeNum, e := tcpconn.Write([]byte("这是客户端发送的tcp数据。。。"))
	//check_err_net(e)
	//
	//fmt.Println("发送了的字节数为：",writeNum)
	//fmt.Println("发送了的字节转字符串数据为：","这是客户端发送的tcp数据。。。")
	//
	//defer tcpconn.Close()
	//
	//
	//fmt.Println("正在等待服务器响应。。。")
	////下面是读取服务端响应回来的数据
	//fmt.Println()
	//var readby =make([]byte,100)
	//readnum, e := tcpconn.Read(readby)
	//check_err_net(e)
	//
	//fmt.Println("响应回来的字节数为：",readnum)
	//fmt.Println("响应回来的字节数据为：",readby)
	//fmt.Println("响应回来的字节转字符串数据为：",string(readby))
	//
	////输出：
	////	发送了的字节数为： 42
	////	发送了的字节转字符串数据为： 这是客户端发送的tcp数据。。。
	////	正在等待服务器响应。。。
	////
	////	响应回来的字节数为： 49
	////	响应回来的字节数据为： [232 191 153 230 152 175 116 99 112 230 156 141 229 138 161 229 153 168 229 147 141
	////	229 186 148 231 187 153 229 174 162 230 136 183 231 171 175 231 154 132 229 134 133 229 174 185 50 50 50 50
	////	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	////	响应回来的字节转字符串数据为： 这是tcp服务器响应给客户端的内容2222




	////在开始下面之前我们先注释掉上面的代码！
	//fmt.Println()
	//fmt.Println("----------net包下的函数之net.ParseCIDR()函数以及对ip和ipnet对象下的方法的讲解--------------")
	//// ParseCIDR parses s as a CIDR notation IP address and prefix length,
	//// like "192.0.2.0/24" or "2001:db8::/32", as defined in
	//// RFC 4632 and RFC 4291.
	////
	//// It returns the IP address and the network implied by the IP and
	//// prefix length.
	//// For example, ParseCIDR("192.0.2.1/24") returns the IP address
	//// 192.0.2.1 and the network 192.0.2.0/24.
	////ParseCIDR将s解析为CIDR表示法IP地址和前缀长度，例如RFC 4632和RFC 4291中定义的“ 192.0.2.0/24”或“ 2001：db8 :: / 32”。
	////它返回IP地址和IP和前缀长度所隐含的网络。
	////例如，ParseCIDR（“ 192.0.2.1/24”）返回IP地址192.0.2.1和网络192.0.2.0/24。
	//
	////// An IP is a single IP address, a slice of bytes.
	////// Functions in this package accept either 4-byte (IPv4)
	////// or 16-byte (IPv6) slices as input.
	//////
	////// Note that in this documentation, referring to an
	////// IP address as an IPv4 address or an IPv6 address
	////// is a semantic property of the address, not just the
	////// length of the byte slice: a 16-byte slice can still
	////// be an IPv4 address.
	//////IP是单个IP地址，是字节的一部分。
	//////此程序包中的功能接受4字节（IPv4）或16字节（IPv6）切片作为输入。
	//////请注意，在本文档中，将IP地址称为IPv4地址或IPv6地址是该地址的语义属性，而不仅仅是字节片的长度：16字节的片仍可以是IPv4地址。
	////type IP []byte
	////
	////// An IP mask is an IP address.// IP掩码是IP地址。
	////type IPMask []byte
	////
	////// An IPNet represents an IP network.// IPNet代表IP网络。
	////type IPNet struct {
	////	IP   IP     // network number//网络号
	////	Mask IPMask // network mask//网络掩码
	////}
	//
	//var once sync.Once //必须放外面，不能放函数里面，否则多次执行函数会init这个对象！
	//f := func(s string) {
	//	ip, ipNet, e := net.ParseCIDR(s)
	//	check_err_net(e)
	//	fmt.Println("net.ParseCIDR()返回值ips, ipNet：", ip, ipNet)
	//
	//	//我们不希望每次执行这个函数都运行下面的额代码
	//	once.Do(func() {
	//		fmt.Println("下面对ip这个对象下的方法进行例子展示：")
	//		fmt.Println("ip.String():", ip.String())
	//		fmt.Println("ip.To4():", ip.To4())
	//
	//		//var ip1 net.IP
	//		//copy的dst切片必须有长度，否则无法进行复制
	//		var ip1 = make(net.IP, len(ip)) //net.IP其实也是切片类族
	//		i := copy(ip1, ip)
	//		fmt.Println("对源ip进行复制，复制了的字节数为：", i)
	//		fmt.Println("ip.Equal(ip1):", ip.Equal(ip1))
	//		fmt.Println("ip.Equal(net.ParseIP(\"127.0.0.1\")):", ip.Equal(net.ParseIP("127.0.0.1")))
	//
	//		fmt.Println("ip.To16():", ip.To16())
	//
	//		// MarshalText implements the encoding.TextMarshaler interface.
	//		// The encoding is the same as returned by String, with one exception:
	//		// When len(ip) is zero, it returns an empty slice.
	//		//	MarshalText实现encoding.TextMarshaler接口。
	//		//	编码与String返回的编码相同，但有一个例外：
	//		//	当len（ip）为零时，它返回一个空切片。
	//		MarshalTextBytes, e := ip.MarshalText()
	//		check_err_net(e)
	//		fmt.Println("ip.MarshalText()返回值切片:", MarshalTextBytes)
	//		fmt.Println("ip.MarshalText()返回值切片转字符串:", string(MarshalTextBytes))
	//
	//		// UnmarshalText implements the encoding.TextUnmarshaler interface.
	//		// The IP address is expected in a form accepted by ParseIP.
	//		//	UnmarshalText实现encoding.TextUnmarshaler接口。
	//		//	IP地址应采用ParseIP接受的格式。
	//		var ip2 net.IP
	//		fmt.Println("ip.UnmarshalText(MarshalTextBytes)之前:", ip2)
	//		e = ip2.UnmarshalText(MarshalTextBytes)
	//		fmt.Println("ip.UnmarshalText(MarshalTextBytes)之后:", ip2)
	//		fmt.Println("ip.Equal(ip2):", ip.Equal(ip2))
	//
	//		// DefaultMask returns the default IP mask for the IP address ip.
	//		// Only IPv4 addresses have default masks; DefaultMask returns
	//		// nil if ip is not a valid IPv4 address.
	//		//DefaultMask返回IP地址ip的默认IP掩码。
	//		//只有IPv4地址具有默认掩码； 如果ip不是有效的IPv4地址，则DefaultMask返回nil。
	//		mask := ip.DefaultMask()
	//		fmt.Println("ip.DefaultMask()返回值mask:", mask)
	//		fmt.Println("ip.DefaultMask()返回值mask转切片:", []byte(mask))
	//		fmt.Println("ip.DefaultMask()返回值mask.String():", mask.String())
	//
	//		// IsGlobalUnicast reports whether ip is a global unicast
	//		// address.
	//		//
	//		// The identification of global unicast addresses uses address type
	//		// identification as defined in RFC 1122, RFC 4632 and RFC 4291 with
	//		// the exception of IPv4 directed broadcast addresses.
	//		// It returns true even if ip is in IPv4 private address space or
	//		// local IPv6 unicast address space.
	//		//	IsGlobalUnicast报告ip是否为全局单播地址。
	//		//	全局单播地址的标识使用RFC 1122，RFC 4632和RFC 4291中定义的地址类型标识，但IPv4定向广播地址除外。
	//		//	即使ip位于IPv4专用地址空间或本地IPv6单播地址空间中，它也会返回true。
	//		fmt.Println("ip.IsGlobalUnicast():", ip.IsGlobalUnicast())
	//
	//		//239.0.0.0~239.255.255.255的地址使用要求被限制在特定的多播域内
	//		var ip4 = net.ParseIP("127.0.0.1")
	//		var ip5 = net.ParseIP("255.255.255.255")
	//		var ip6 = net.ParseIP("255.255.255.0")
	//		var ip7 = net.ParseIP("0.0.0.0")
	//		var ip8 = net.ParseIP("239.0.0.0")
	//		var ip9 = net.ParseIP("239.0.0.1")
	//		var ip10 = net.ParseIP("239.0.1.1")
	//		var ip11 = net.ParseIP("239.255.255.255")
	//
	//		var ipby = []net.IP{ip4, ip5, ip6, ip7, ip8,ip9,ip10,ip11}
	//		//这个方法的底层可以自己去看看
	//
	//		for _, ip := range ipby {
	//			fmt.Println("ip.IsGlobalUnicast():", ip.IsGlobalUnicast())
	//			// IsMulticast reports whether ip is a multicast address.
	//			// IsMulticast报告ip是否为组播(多播)地址。
	//			fmt.Println("ip.IsMulticast():",ip.IsMulticast())
	//
	//		}
	//
	//
	//
	//		fmt.Println()
	//
	//
	//		//// Well-known IPv4 addresses
	//		//var (
	//		//	IPv4bcast     = IPv4(255, 255, 255, 255) // limited broadcast//广播有限
	//		//	IPv4allsys    = IPv4(224, 0, 0, 1)       // all systems//所有系统
	//		//	IPv4allrouter = IPv4(224, 0, 0, 2)       // all routers//所有路由器
	//		//	IPv4zero      = IPv4(0, 0, 0, 0)         // all zeros//全零
	//		//)
	//
	//		//// Well-known IPv6 addresses
	//		//var (
	//		//	IPv6zero                   = IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	//		//	IPv6unspecified            = IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	//		//	IPv6loopback               = IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	//		//	IPv6interfacelocalallnodes = IP{0xff, 0x01, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01}
	//		//	IPv6linklocalallnodes      = IP{0xff, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01}
	//		//	IPv6linklocalallrouters    = IP{0xff, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x02}
	//		//)
	//
	//		// IsLoopback reports whether ip is a loopback address.
	//		// IsLoopback报告ip是否为回送地址。ipv4为127.x.x.x形式，ipv6则为IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	//		fmt.Println("ip.IsLoopback()",ip.IsLoopback())
	//		fmt.Println("ip4.IsLoopback()",ip4.IsLoopback())
	//		var ip_loop1  = net.ParseIP("127.1.0.1")
	//		var ip_loop2  = net.ParseIP("127.1.0.255")
	//		fmt.Println("ip_loop1.IsLoopback()",ip_loop1.IsLoopback())
	//		fmt.Println("ip_loop2.IsLoopback()",ip_loop2.IsLoopback())
	//
	//
	//		fmt.Println()
	//		// IsUnspecified reports whether ip is an unspecified address, either
	//		// the IPv4 address "0.0.0.0" or the IPv6 address "::".
	//		// IsUnspecified报告ip是否为未指定地址，即IPv4地址“ 0.0.0.0”还是IPv6地址“ ::”。
	//		ip0:=net.ParseIP("0.0.0.0")
	//		fmt.Println("ip.IsUnspecified()",ip.IsUnspecified())
	//		fmt.Println("ip0.IsUnspecified()",ip0.IsUnspecified())
	//
	//
	//		fmt.Println()
	//
	//		// IsInterfaceLocalMulticast reports whether ip is
	//		// an interface-local multicast address.
	//		//IsInterfaceLocalMulticast报告ip是否为接口本地多播地址。这个我也不是很懂
	//		fmt.Println("net.IPv6interfacelocalallnodes.IsInterfaceLocalMulticast()",net.IPv6interfacelocalallnodes.IsInterfaceLocalMulticast())
	//
	//
	//		fmt.Println()
	//		// IsLinkLocalMulticast reports whether ip is a link-local
	//		// multicast address.
	//		// IsLinkLocalMulticast报告ip是否为本地链接的多播地址。
	//		fmt.Println("ip.IsLinkLocalMulticast()",ip.IsLinkLocalMulticast())
	//		fmt.Println("net.IPv6linklocalallnodes.IsLinkLocalMulticast()",net.IPv6linklocalallnodes.IsLinkLocalMulticast())
	//		fmt.Println("net.IPv6linklocalallrouters.IsLinkLocalMulticast()",net.IPv6linklocalallrouters.IsLinkLocalMulticast())
	//
	//		//下面是验证ipv4的是否为本地链接的多播地址。
	//		//看底层可以知道，ipv4的本地连接的多播地址指的是224.0.0.x,但是ipv6判断条件不同！
	//		fmt.Println()
	//		var myIp0=net.IP{224,0,0,1}
	//		var myIp1=net.IP{224,0,0,2}
	//		var myIp2=net.IP{224,0,1,2}
	//		var myIp3=net.IP{224,1,0,2}
	//
	//		var myby=[]net.IP{myIp0,myIp1,myIp2,myIp3}
	//		for _, ip := range myby {
	//			fmt.Println("ip.IsLinkLocalMulticast():",ip.IsLinkLocalMulticast())
	//		}
	//
	//
	//		fmt.Println()
	//		// IsLinkLocalUnicast reports whether ip is a link-local
	//		// unicast address.
	//		// IsLinkLocalUnicast报告ip是否为本地链接单播地址。ipv4是类似于169.254.x.x,ipv6判断条件则不相同！
	//		fmt.Println("ip.IsLinkLocalMulticast()",ip.IsLinkLocalUnicast())
	//
	//
	//		var myIp4=net.IP{169,254,0,1}
	//		var myIp5=net.IP{169,254,1,1}
	//		var myIp6=net.IP{169,255,1,1}
	//		var myby1=[]net.IP{myIp4,myIp5,myIp6}
	//		for _, ip := range myby1 {
	//			fmt.Println("ip.IsLinkLocalMulticast():",ip.IsLinkLocalUnicast())
	//		}
	//
	//
	//		fmt.Println()
	//		// Mask returns the result of masking the IP address ip with mask.
	//		// Mask返回使用mask屏蔽IP地址ip的结果。
	//
	//		fmt.Println("len(ip.DefaultMask())：",len(ip.DefaultMask()))
	//		//ip:192.0.2.1
	//		fmt.Println("ip.Mask(ip.DefaultMask()):",ip.Mask(ip.DefaultMask()))
	//
	//		fmt.Println("ip.Mask([]byte(\"255.255.255.255\")):",ip.Mask(net.IPMask([]byte("255.255.255.255"))))
	//		//千万不能像上面这样得到4个字节的掩码地址，因为其中的.号也已经写入了字节切片里面去了，所以我们要像下面这样先得到ip地址！
	//		fmt.Println("ip.Mask([]byte(\"255.255.255.255\")):",ip.Mask(net.IPMask(net.ParseIP("255.255.255.255"))))
	//		fmt.Println("ip.Mask([]byte(\"255.255.255.0\")):",ip.Mask(net.IPMask(net.ParseIP("255.255.255.0"))))
	//		fmt.Println("ip.Mask([]byte(\"255.255.0.0\")):",ip.Mask(net.IPMask(net.ParseIP("255.255.0.0"))))
	//		fmt.Println("ip.Mask([]byte(\"255.0.0.0\")):",ip.Mask(net.IPMask(net.ParseIP("255.0.0.0"))))
	//		fmt.Println("ip.Mask([]byte(\"0.0.0.0\")):",ip.Mask(net.IPMask(net.ParseIP("0.0.0.0"))))
	//
	//
	//		fmt.Println("好了，ip对象下的方法介绍到此，下面我们来介绍ipnet对象下面的方法：")
	//
	//		fmt.Println("ipNet.Mask:",ipNet.Mask)
	//		fmt.Println("ipNet.IP:",ipNet.IP)
	//		fmt.Println("ipNet.String():",ipNet.String())
	//		fmt.Println("ipNet.Contains(net.ParseIP(\"127.0.0.1\")):",ipNet.Contains(net.ParseIP("192.0.2.1")))
	//		fmt.Println("ipNet.Contains(net.ParseIP(\"127.0.0.1\")):",ipNet.Contains(net.ParseIP("127.0.0.1")))
	//
	//
	//	})
	//
	//}
	//
	//f("192.0.2.1/24")
	//
	//fmt.Println()
	//f("192.0.2.1/16")
	//f("192.0.2.1/8")
	//f("192.0.2.1/0")
	//
	////输出：
	////	net.ParseCIDR()返回值ips, ipNet： 192.0.2.1 192.0.2.0/24
	////	下面对ip这个对象下的方法进行例子展示：
	////	ip.String(): 192.0.2.1
	////	ip.To4(): 192.0.2.1
	////	对源ip进行复制，复制了的字节数为： 16
	////	ip.Equal(ip1): true
	////	ip.Equal(net.ParseIP("127.0.0.1")): false
	////	ip.To16(): 192.0.2.1
	////	ip.MarshalText()返回值切片: [49 57 50 46 48 46 50 46 49]
	////	ip.MarshalText()返回值切片转字符串: 192.0.2.1
	////	ip.UnmarshalText(MarshalTextBytes)之前: <nil>
	////	ip.UnmarshalText(MarshalTextBytes)之后: 192.0.2.1
	////	ip.Equal(ip2): true
	////	ip.DefaultMask()返回值mask: ffffff00
	////	ip.DefaultMask()返回值mask转切片: [255 255 255 0]
	////	ip.DefaultMask()返回值mask.String(): ffffff00
	////	ip.IsGlobalUnicast(): true
	////	ip.IsGlobalUnicast(): false
	////	ip.IsMulticast(): false
	////	ip.IsGlobalUnicast(): false
	////	ip.IsMulticast(): false
	////	ip.IsGlobalUnicast(): true
	////	ip.IsMulticast(): false
	////	ip.IsGlobalUnicast(): false
	////	ip.IsMulticast(): false
	////	ip.IsGlobalUnicast(): false
	////	ip.IsMulticast(): true
	////	ip.IsGlobalUnicast(): false
	////	ip.IsMulticast(): true
	////	ip.IsGlobalUnicast(): false
	////	ip.IsMulticast(): true
	////	ip.IsGlobalUnicast(): false
	////	ip.IsMulticast(): true
	////
	////	ip.IsLoopback() false
	////	ip4.IsLoopback() true
	////	ip_loop1.IsLoopback() true
	////	ip_loop2.IsLoopback() true
	////
	////	ip.IsUnspecified() false
	////	ip0.IsUnspecified() true
	////
	////	net.IPv6interfacelocalallnodes.IsInterfaceLocalMulticast() true
	////
	////	ip.IsLinkLocalMulticast() false
	////	net.IPv6linklocalallnodes.IsLinkLocalMulticast() true
	////	net.IPv6linklocalallrouters.IsLinkLocalMulticast() true
	////
	////	ip.IsLinkLocalMulticast(): true
	////	ip.IsLinkLocalMulticast(): true
	////	ip.IsLinkLocalMulticast(): false
	////	ip.IsLinkLocalMulticast(): false
	////
	////	ip.IsLinkLocalMulticast() false
	////	ip.IsLinkLocalMulticast(): true
	////	ip.IsLinkLocalMulticast(): true
	////	ip.IsLinkLocalMulticast(): false
	////
	////	len(ip.DefaultMask())： 4
	////	ip.Mask(ip.DefaultMask()): 192.0.2.0
	////	ip.Mask([]byte("255.255.255.255")): <nil>
	////	ip.Mask([]byte("255.255.255.255")): 192.0.2.1
	////	ip.Mask([]byte("255.255.255.0")): 192.0.2.0
	////	ip.Mask([]byte("255.255.0.0")): 192.0.0.0
	////	ip.Mask([]byte("255.0.0.0")): 192.0.0.0
	////	ip.Mask([]byte("0.0.0.0")): 0.0.0.0
	////	好了，ip对象下的方法介绍到此，下面我们来介绍ipnet对象下面的方法：
	////	ipNet.Mask: ffffff00
	////	ipNet.IP: 192.0.2.0
	////	ipNet.String(): 192.0.2.0/24
	////	ipNet.Contains(net.ParseIP("127.0.0.1")): true
	////	ipNet.Contains(net.ParseIP("127.0.0.1")): false
	////
	////	net.ParseCIDR()返回值ips, ipNet： 192.0.2.1 192.0.0.0/16
	////	net.ParseCIDR()返回值ips, ipNet： 192.0.2.1 192.0.0.0/8
	////	net.ParseCIDR()返回值ips, ipNet： 192.0.2.1 0.0.0.0/0



	//
	////注释掉上面的代码！
	//fmt.Println()
	//fmt.Println("----------net包下的函数之net.InterfaceXxx系列函数返回本机的网卡ip信息--------------")
	//
	//
	////// BUG(mikio): On JS and NaCl, methods and functions related to
	////// Interface are not implemented.
	////
	////// BUG(mikio): On AIX, DragonFly BSD, NetBSD, OpenBSD, Plan 9 and
	////// Solaris, the MulticastAddrs method of Interface is not implemented.
	//////BUG（mikio）：在JS和NaCl上，未实现与Interface相关的方法和功能。
	//////BUG（mikio）：在AIX，DragonFly BSD，NetBSD，OpenBSD，Plan 9和Solaris上，未实现Interface的MulticastAddrs方法。
	////
	////var (
	////	errInvalidInterface         = errors.New("invalid network interface")//无效的网络接口
	////	errInvalidInterfaceIndex    = errors.New("invalid network interface index")//无效的网络接口索引号
	////	errInvalidInterfaceName     = errors.New("invalid network interface name")//无效的网络接口命名
	////	errNoSuchInterface          = errors.New("no such network interface")//没有此网络接口
	////	errNoSuchMulticastInterface = errors.New("no such multicast network interface")//没有此类网络接口多播地址
	////)
	////
	////// Interface represents a mapping between network interface name
	////// and index. It also represents network interface facility
	////// information.
	//////接口表示网络接口名称和索引之间的映射。 它还代表网络接口设施信息。
	////type Interface struct {
	////	Index        int          // positive integer that starts at one, zero is never used//从1开始的正整数，从不使用零
	////	MTU          int          // maximum transmission unit//最大传输单位
	////	Name         string       // e.g., "en0", "lo0", "eth0.100"// er,"en0", "lo0", "eth0.100"
	////	HardwareAddr HardwareAddr // IEEE MAC-48, EUI-48 and EUI-64 form// IEEE MAC-48，EUI-48和EUI-64形式
	////	Flags        Flags        // e.g., FlagUp, FlagLoopback, FlagMulticast//例如，FlagUp，FlagLoopback，FlagMulticast
	////}
	////// A HardwareAddr represents a physical hardware address.//HardwareAddr代表物理硬件地址。
	////type HardwareAddr []byte
	////
	////type Flags uint
	////
	////const (
	////	FlagUp           Flags = 1 << iota // interface is up//被启动的接口
	////	FlagBroadcast                      // interface supports broadcast access capability//接口支持广播访问功能
	////	FlagLoopback                       // interface is a loopback interface//接口是一个环回接口
	////	FlagPointToPoint                   // interface belongs to a point-to-point link//接口属于点对点链接
	////	FlagMulticast                      // interface supports multicast access capability//接口支持多播访问功能
	////)
	////
	////var flagNames = []string{
	////	"up",
	////	"broadcast",
	////	"loopback",
	////	"pointtopoint",
	////	"multicast",
	////}
	//
	//
	//
	//
	//
	//
	//// Interfaces returns a list of the system's network interfaces.
	////接口返回系统网络接口的列表。
	//interfaces, e := net.Interfaces()
	//check_err_net(e)
	//fmt.Printf("net.Interfaces():%+v\n",interfaces)
	//for k, v := range interfaces {
	//	fmt.Printf("----------%v---------\n",k)
	//	fmt.Println("interface.Name:",v.Name)
	//	fmt.Println("interface.Index:",v.Index)
	//	fmt.Println("interface.MTU:",v.MTU)
	//	fmt.Println("interface.Flags:",v.Flags)
	//	fmt.Println("interface.HardwareAddr:",v.HardwareAddr)
	//	// Addrs returns a list of unicast interface addresses for a specific
	//	// interface.
	//	// Addrs返回特定接口的单播接口地址列表。
	//	addrs, e := v.Addrs()
	//	check_err_net(e)
	//	fmt.Println("interface.Addrs():",addrs)
	//	// MulticastAddrs returns a list of multicast, joined group addresses
	//	// for a specific interface.
	//	// MulticastAddrs返回特定接口的多播，已加入组地址的列表。
	//	multicastAddrs, e := v.MulticastAddrs()
	//	check_err_net(e)
	//	fmt.Println("interface.MulticastAddrs():",multicastAddrs)
	//}
	////输出：
	////	net.Interfaces():[{Index:18//网卡的id号
	////						MTU:65536//网卡能发送的路由的最大数据大小，一般是1500
	////						Name:Npcap Loopback Adapter//网卡名字
	////						HardwareAddr:02:00:4c:4f:4f:50//网卡硬件地址
	////						Flags:up|broadcast|multicast}//-------up表示正在使用，broadcast表示可发广播，multicast表示可发多播，下同，不再累叙----------------------------
	////						{Index:5
	////						MTU:1500
	////						Name:以太网
	////						HardwareAddr:00:f1:f3:8f:40:9a
	////						Flags:up|broadcast|multicast}//-----------------------------------
	////						{Index:12
	////						MTU:1500
	////						Name:VMware Network Adapter VMnet1
	////						HardwareAddr:00:50:56:c0:00:01
	////						Flags:up|broadcast|multicast}//-----------------------------------
	////						{Index:10
	////						MTU:1500
	////						Name:VMware Network Adapter VMnet8
	////						HardwareAddr:00:50:56:c0:00:08
	////						Flags:up|broadcast|multicast}//-----------------------------------
	////						{Index:16
	////						MTU:1500
	////						Name:以太网 3
	////						HardwareAddr:00:ff:81:59:68:62
	////						Flags:broadcast|multicast}//-----------------------------------
	////						{Index:20
	////						MTU:1500
	////						Name:以太网 4
	////						HardwareAddr:00:ff:f4:40:2e:17
	////						Flags:broadcast|multicast}//-----------------------------------
	////						{Index:3
	////						MTU:1500
	////						Name:以太网 2
	////						HardwareAddr:00:ff:0c:82:d6:8c
	////						Flags:broadcast|multicast}//-----------------------------------
	////						{Index:1
	////						MTU:-1
	////						Name:Loopback Pseudo-Interface 1
	////						HardwareAddr: Flags:up|loopback|multicast}]
	////	----------0---------
	////	interface.Name: Npcap Loopback Adapter
	////	interface.Index: 18
	////	interface.MTU: 65536
	////	interface.Flags: up|broadcast|multicast
	////	interface.HardwareAddr: 02:00:4c:4f:4f:50
	////	interface.Addrs(): [fe80::9c8a:2936:7699:9434/64 169.254.148.52/16]
	////	interface.MulticastAddrs(): [ff01::1 ff02::1 ff02::c ff02::fb ff02::1:3 ff02::1:ff99:9434 224.0.0.1 224.0.0.251 224.0.0.252 239.255.255.250]
	////	----------1---------
	////	interface.Name: 以太网
	////	interface.Index: 5
	////	interface.MTU: 1500
	////	interface.Flags: up|broadcast|multicast
	////	interface.HardwareAddr: 00:f1:f3:8f:40:9a
	////	interface.Addrs(): [fe80::69af:7253:ebb9:bacd/64 192.168.1.102/24]
	////	interface.MulticastAddrs(): [ff01::1 ff02::1 ff02::c ff02::fb ff02::1:3 ff02::1:ffb9:bacd 224.0.0.1 224.0.0.251 224.0.0.252 239.255.255.250]
	////	----------2---------
	////	interface.Name: VMware Network Adapter VMnet1
	////	interface.Index: 12
	////	interface.MTU: 1500
	////	interface.Flags: up|broadcast|multicast
	////	interface.HardwareAddr: 00:50:56:c0:00:01
	////	interface.Addrs(): [fe80::4491:928a:684:d057/64 192.168.170.1/24]
	////	interface.MulticastAddrs(): [ff01::1 ff02::1 ff02::c ff02::fb ff02::1:3 ff02::1:ff84:d057 224.0.0.1 224.0.0.251 224.0.0.252 239.255.255.250]
	////	----------3---------
	////	interface.Name: VMware Network Adapter VMnet8
	////	interface.Index: 10
	////	interface.MTU: 1500
	////	interface.Flags: up|broadcast|multicast
	////	interface.HardwareAddr: 00:50:56:c0:00:08
	////	interface.Addrs(): [fe80::cd27:9f16:5158:406c/64 192.168.63.1/24]
	////	interface.MulticastAddrs(): [ff01::1 ff02::1 ff02::c ff02::fb ff02::1:3 ff02::1:ff58:406c 224.0.0.1 224.0.0.251 224.0.0.252 239.255.255.250]
	////	----------4---------
	////	interface.Name: 以太网 3
	////	interface.Index: 16
	////	interface.MTU: 1500
	////	interface.Flags: broadcast|multicast
	////	interface.HardwareAddr: 00:ff:81:59:68:62
	////	interface.Addrs(): [fe80::f468:5294:7e:f7d0/64 169.254.247.208/16 198.18.0.0/15]
	////	interface.MulticastAddrs(): [ff01::1 ff02::1 ff02::1:ff7e:f7d0 224.0.0.1]
	////	----------5---------
	////	interface.Name: 以太网 4
	////	interface.Index: 20
	////	interface.MTU: 1500
	////	interface.Flags: broadcast|multicast
	////	interface.HardwareAddr: 00:ff:f4:40:2e:17
	////	interface.Addrs(): [fe80::f1b6:1cf0:389b:5f0d/64 169.254.95.13/16]
	////	interface.MulticastAddrs(): [ff01::1 ff02::1 ff02::1:ff9b:5f0d 224.0.0.1]
	////	----------6---------
	////	interface.Name: 以太网 2
	////	interface.Index: 3
	////	interface.MTU: 1500
	////	interface.Flags: broadcast|multicast
	////	interface.HardwareAddr: 00:ff:0c:82:d6:8c
	////	interface.Addrs(): [fe80::10ae:422f:5a17:8f82/64 169.254.143.130/16]
	////	interface.MulticastAddrs(): [ff01::1 ff02::1 ff02::1:ff17:8f82 224.0.0.1]
	////	----------7---------
	////	interface.Name: Loopback Pseudo-Interface 1
	////	interface.Index: 1
	////	interface.MTU: -1
	////	interface.Flags: up|loopback|multicast
	////	interface.HardwareAddr:
	////	interface.Addrs(): [::1/128 127.0.0.1/8]
	////	interface.MulticastAddrs(): [ff02::c 239.255.255.250]
	//
	//
	//fmt.Println()
	//// InterfaceAddrs returns a list of the system's unicast interface
	//// addresses.
	////
	//// The returned list does not identify the associated interface; use
	//// Interfaces and Interface.Addrs for more detail.
	//// InterfaceAddrs返回系统单播接口地址的列表。
	////返回的列表不标识关联的接口； 有关更多详细信息，请使用上面的Interfaces和Interface.Addrs。
	//addrs, e := net.InterfaceAddrs()
	//check_err_net(e)
	//fmt.Println(addrs)
	//for _, v := range addrs {
	//	fmt.Println(v)
	//}
	//
	////输出：
	////	[fe80::9c8a:2936:7699:9434/64
	////	169.254.148.52/16
	////	fe80::69af:7253:ebb9:bacd/64
	////	192.168.1.102/24
	////	fe80::4491:928a:684:d057/64
	////	192.168.170.1/24
	////	fe80::cd27:9f16:5158:406c/64
	////	192.168.63.1/24
	////	fe80::f468:5294:7e:f7d0/64
	////	169.254.247.208/16
	////	198.18.0.0/15
	////	fe80::f1b6:1cf0:389b:5f0d/64
	////	169.254.95.13/16
	////	fe80::10ae:422f:5a17:8f82/64
	////	169.254.143.130/16
	////	::1/128
	////	127.0.0.1/8]
	////除了198.18.0.0/15，上面每个ipv4地址对应一个ipv6地址，所以输出了8*2+1=17个地址
	//
	////	fe80::9c8a:2936:7699:9434/64
	////	169.254.148.52/16
	////	fe80::69af:7253:ebb9:bacd/64
	////	192.168.1.102/24
	////	fe80::4491:928a:684:d057/64
	////	192.168.170.1/24
	////	fe80::cd27:9f16:5158:406c/64
	////	192.168.63.1/24
	////	fe80::f468:5294:7e:f7d0/64
	////	169.254.247.208/16
	////	198.18.0.0/15
	////	fe80::f1b6:1cf0:389b:5f0d/64
	////	169.254.95.13/16
	////	fe80::10ae:422f:5a17:8f82/64
	////	169.254.143.130/16
	////	::1/128
	////	127.0.0.1/8
	//
	//
	//fmt.Println()
	//// InterfaceByName returns the interface specified by name.
	//// InterfaceByName返回由名称指定的网卡接口。名称请看上面有。
	//interfacePointer, e := net.InterfaceByName("Npcap Loopback Adapter")
	//check_err_net(e)
	//fmt.Printf("%+v\n",interfacePointer)
	//
	//interfacePointer, e = net.InterfaceByName("以太网")
	//check_err_net(e)
	//fmt.Printf("%+v\n",interfacePointer)
	////输出：
	////	&{Index:18
	////	MTU:65536
	////	Name:Npcap Loopback Adapter
	////	HardwareAddr:02:00:4c:4f:4f:50
	////	Flags:up|broadcast|multicast}
	////	&{Index:5
	////	MTU:1500
	////	Name:以太网
	////	HardwareAddr:00:f1:f3:8f:40:9a
	////	Flags:up|broadcast|multicast}
	//
	//fmt.Println()
	//// InterfaceByIndex returns the interface specified by index.
	////
	//// On Solaris, it returns one of the logical network interfaces
	//// sharing the logical data link; for more precision use
	//// InterfaceByName.
	//// InterfaceByIndex返回由index指定的接口。
	////在Solaris上，它返回共享逻辑数据链接的逻辑网络接口之一； 为了更精确，请使用InterfaceByName。
	//interfacePointer, e = net.InterfaceByIndex(18)
	//check_err_net(e)
	//fmt.Printf("%+v\n",interfacePointer)
	//
	//interfacePointer, e = net.InterfaceByIndex(5)
	//check_err_net(e)
	//fmt.Printf("%+v\n",interfacePointer)
	//
	////interfacePointer, e = net.InterfaceByIndex(555)//故意给个不存在的网卡id号
	////check_err_net(e)
	////fmt.Printf("%+v\n",interfacePointer)
	////输出：
	////	&{Index:18
	////	MTU:65536
	////	Name:Npcap Loopback Adapter
	////	HardwareAddr:02:00:4c:4f:4f:50
	////	Flags:up|broadcast|multicast}
	////	&{Index:5
	////	MTU:1500
	////	Name:以太网
	////	HardwareAddr:00:f1:f3:8f:40:9a
	////	Flags:up|broadcast|multicast}
	////	出错了，错误信息为： route ip+net: no such network interface
	////	panic: route ip+net: no such network interface
	////
	////	goroutine 1 [running]:
	////	main.check_err_net(0x5069a0, 0xc0000ee000)
	////	C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:2880 +0xdd
	////	main.main()
	////	C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:2823 +0xd06



	////注释掉上面的代码
	//fmt.Println()
	//fmt.Println("----------net包下的函数之net.ParseMAC()函数解析其他mac地址成为“:”间隔的的mac地址--------------")
	//
	////// A HardwareAddr represents a physical hardware address.
	////// HardwareAddr代表物理硬件地址。
	////type HardwareAddr []byte
	//
	//
	//// ParseMAC parses s as an IEEE 802 MAC-48, EUI-48, EUI-64, or a 20-octet
	//// IP over InfiniBand link-layer address using one of the following formats:
	//// ParseMAC使用以下格式之一将s解析为IEEE 802 MAC-48，EUI-48，EUI-64或InfiniBand链路层地址上的20个八位位组IP：
	////	00:00:5e:00:53:01
	////	02:00:5e:10:00:00:00:01
	////	00:00:00:00:fe:80:00:00:00:00:00:00:02:00:5e:10:00:00:00:01
	////	00-00-5e-00-53-01
	////	02-00-5e-10-00-00-00-01
	////	00-00-00-00-fe-80-00-00-00-00-00-00-02-00-5e-10-00-00-00-01
	////	0000.5e00.5301
	////	0200.5e10.0000.0001
	////	0000.0000.fe80.0000.0000.0000.0200.5e10.0000.0001
	//hw, err := net.ParseMAC("02:00:4c:4f:4f:50")
	//check_err_net(err)
	//fmt.Println(hw)
	//
	//hw, err = net.ParseMAC("02-00-4C-4F-4F-50")
	//check_err_net(err)
	//fmt.Println(hw)
	//
	//hw, err = net.ParseMAC("0200.4C4F.4F50")
	//check_err_net(err)
	//fmt.Println(hw)
	//
	////hw, err = net.ParseMAC("fe80::9c8a:2936:7699:9434")//这个是ipv6地址，但不是mac地址
	////check_err_net(err)
	////fmt.Println(hw)
	//
	////输出：
	////	02:00:4c:4f:4f:50
	////	02:00:4c:4f:4f:50
	////	02:00:4c:4f:4f:50
	////	出错了，错误信息为： address fe80::9c8a:2936:7699:9434: invalid MAC address
	////	panic: address fe80::9c8a:2936:7699:9434: invalid MAC address
	////
	////	goroutine 1 [running]:
	////	main.check_err_net(0x4fb960, 0xc000004560)
	////		C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:2940 +0xdd
	////	main.main()
	////		C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:2883 +0x37a




	////注释掉上面的代码
	//fmt.Println()
	//fmt.Println("----------net包下的函数之net.SplitHostPort()函数拆分host,post--------------")
	//
	//// SplitHostPort splits a network address of the form "host:port",
	//// "host%zone:port", "[host]:port" or "[host%zone]:port" into host or
	//// host%zone and port.
	////
	//// A literal IPv6 address in hostport must be enclosed in square
	//// brackets, as in "[::1]:80", "[::1%lo0]:80".
	////
	//// See func Dial for a description of the hostport parameter, and host
	//// and port results.
	//// SplitHostPort将"host:port"，"host%zone:port"，"[host]:port"或"[host%zone]:port"形式的网络地址拆分为host或host%zone 和port。
	////主机端口中的字面量IPv6地址必须用方括号括起来，例如"[::1]:80", "[::1%lo0]:80"。
	////有关hostport参数以及主机和端口结果的说明，请参见func Dial。
	//host, port, err := net.SplitHostPort("192.168.1.102/24:8080")
	//check_err_net(err)
	//fmt.Println("host, port:",host, port)
	//
	//host, port, err = net.SplitHostPort("[192.168.1.102/24]:8080")//[]对于ipv4是可选的，但是对于ipv6是必须的，
	//check_err_net(err)
	//fmt.Println("host, port:",host, port)
	//
	//host, port, err = net.SplitHostPort("192.168.1.102:8080")
	//check_err_net(err)
	//fmt.Println("host, port:",host, port)
	//
	//host, port, err = net.SplitHostPort(":8080")//这里不给host则为空字符串，应该是可以省略host的
	//check_err_net(err)
	//fmt.Println("host, port:",host, port)
	//
	//host, port, err = net.SplitHostPort("[]:8080")//这里不给host则为空字符串，应该是可以省略host的
	//check_err_net(err)
	//fmt.Println("host, port:",host, port)
	//
	//host, port, err = net.SplitHostPort("[fe80::69af:7253:ebb9:bacd/64]:8080")
	//check_err_net(err)
	//fmt.Println("host, port:",host, port)
	//
	//host, port, err = net.SplitHostPort("[fe80::69af:7253:ebb9:bacd]:8080")
	//check_err_net(err)
	//fmt.Println("host, port:",host, port)
	//
	//host, port, err = net.SplitHostPort("[fe80::69af:7253:ebb9:bacd]:")//不给端口，这应该是个bug
	//check_err_net(err)
	//fmt.Println("host, port:",host, port)
	//
	//
	//host, port, err = net.SplitHostPort("192.168.1.102:")//不给端口，这应该是个bug
	//check_err_net(err)
	//fmt.Println("host, port:",host, port)
	//
	//host, port, err = net.SplitHostPort("[192.168.1.102]:")//不给端口，这应该是个bug
	//check_err_net(err)
	//fmt.Println("host, port:",host, port)
	//
	////下面的均会报错，ipv6地址的host必须用[]包住，同时无论是哪种ip地址都不能缺少port!
	////host, port, err = net.SplitHostPort("fe80::69af:7253:ebb9:bacd")
	////check_err_net(err)
	////fmt.Println("host, port:",host, port)
	//
	////host, port, err = net.SplitHostPort("192.168.1.102")
	////check_err_net(err)
	////fmt.Println("host, port:",host, port)
	//
	////输出：
	////	host, port: 192.168.1.102/24 8080
	////	host, port: 192.168.1.102/24 8080
	////	host, port: 192.168.1.102 8080
	////	host, port:  8080
	////	host, port:  8080
	////	host, port: fe80::69af:7253:ebb9:bacd/64 8080
	////	host, port: fe80::69af:7253:ebb9:bacd 8080
	////	host, port: fe80::69af:7253:ebb9:bacd
	////	host, port: 192.168.1.102
	////	host, port: 192.168.1.102
	////	出错了，错误信息为： address 192.168.1.102: missing port in address
	////	panic: address 192.168.1.102: missing port in address
	////
	////	goroutine 1 [running]:
	////	main.check_err_net(0x4fcba0, 0xc00005c4a0)
	////		C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:2999 +0xdd
	////	main.main()
	////		C:/Users/Administrator/Desktop/go_pro/src/io_pro/main3/compress_zlib.go:2943 +0x703






	////注释掉上面的代码
	//fmt.Println()
	//fmt.Println("----------net包下的函数之net.JoinHostPort()函数合并host,post--------------")
	//
	//// JoinHostPort combines host and port into a network address of the
	//// form "host:port". If host contains a colon, as found in literal
	//// IPv6 addresses, then JoinHostPort returns "[host]:port".
	////
	//// See func Dial for a description of the host and port parameters.
	//// JoinHostPort将给定的参数主机host和端口port组合成"host:port"形式的网络地址。 如果host包含冒号（如在文字IPv6地址中找到的），则JoinHostPort返回"[host]:port"。
	////有关主机和端口参数的说明，请参见func Dial。
	//
	////按照上面的文档，host是可选的，不给则为空，但是由于有bug,导致port也是可选的，但是go设计api逻辑上是不可选的！
	//hostPort := net.JoinHostPort("192.168.1.102/24", "8080")
	//fmt.Println(hostPort)
	//
	//hostPort = net.JoinHostPort("[192.168.1.102/24]", "8080")//错误的参数，但是不检查
	//fmt.Println(hostPort)
	//
	//hostPort = net.JoinHostPort("192.168.1.102", "8080")
	//fmt.Println(hostPort)
	//
	////这应该也算是个bug,没做任何的检查，直接接收参数连接合并起来了，虽然这样确实可以减少代码，加快执行速度，但是这设计确实是一个bug,考虑不周导致的！
	////当然也可能在其他的api中会进行检查！如果不符合是会抛出错误的！
	//hostPort = net.JoinHostPort("192.168.1.102/", "8080")
	//fmt.Println(hostPort)
	//
	//hostPort = net.JoinHostPort("192.168.1", "8080")//即使给错误的host地址也是不会做任何的检查的！
	//fmt.Println(hostPort)
	//
	//hostPort = net.JoinHostPort("", "8080")//即使不给host地址也是不会做任何的检查的！
	//fmt.Println(hostPort)
	//
	//
	//hostPort = net.JoinHostPort("192.168.1.102", "")//即使给错误的port地址也是不会做任何的检查的！
	//fmt.Println(hostPort)
	//
	//hostPort = net.JoinHostPort("", "")//2个都为空都不会报错，可能go设计原则是用到时候才检查（即必须检查时候），这里没用到host:port所以没检查！
	//fmt.Println(hostPort)
	//
	//hostPort = net.JoinHostPort("fe80::69af:7253:ebb9:bacd/64", "8080")
	//fmt.Println(hostPort)
	//
	//hostPort = net.JoinHostPort("fe80::69af:7253:ebb9:bacd", "8080")
	//fmt.Println(hostPort)
	//
	//hostPort = net.JoinHostPort("fe80::69af:7253:ebb9:bacd/64", "")
	//fmt.Println(hostPort)
	//
	//hostPort = net.JoinHostPort("[fe80::69af:7253:ebb9:bacd/64]", "8080")//这样给都可以，服了
	//fmt.Println(hostPort)
	//
	////从上面可知，他真的不会检查参数的输入正确与否，只是负责连接起来！
	//
	////输出：
	////	192.168.1.102/24:8080
	////	[192.168.1.102/24]:8080
	////	192.168.1.102:8080
	////	192.168.1.102/:8080
	////	192.168.1:8080
	////	:8080
	////	192.168.1.102:
	////	:
	////	[fe80::69af:7253:ebb9:bacd/64]:8080
	////	[fe80::69af:7253:ebb9:bacd]:8080
	////	[fe80::69af:7253:ebb9:bacd/64]:
	////	[[fe80::69af:7253:ebb9:bacd/64]]:8080






	////注释掉上面的代码
	//fmt.Println()
	//fmt.Println("----------net包下的函数之net.LookupXxx()系列函数--------------")
	//
	//
	//// LookupAddr performs a reverse lookup for the given address, returning a list
	//// of names mapping to that address.
	////
	//// When using the host C library resolver, at most one result will be
	//// returned. To bypass the host resolver, use a custom Resolver.
	//// LookupAddr对给定地址执行反向查找，返回本主机映射到该地址的名称列表。
	////使用主机C库解析器时，最多将返回一个结果。 要绕过主机解析器，请使用自定义解析器。
	////
	////只要是loolupXxx函数的底层实现基本和下面的一样的：
	////return DefaultResolver.lookupXxx(context.Background(), addr)
	//
	////DefaultResolver对象的文档如下：
	//
	////// DefaultResolver is the resolver used by the package-level Lookup
	////// functions and by Dialers without a specified Resolver.
	////// DefaultResolver是程序包级别的Lookup函数和没有指定解析程序的Dialers使用的解析程序。
	////var DefaultResolver = &Resolver{}
	//
	//names, err := net.LookupAddr("120.196.165.24")//语句1//这个和下面的那个都是会输出同样的结果，都是我本主机的dns服务器地址
	////names, err := net.LookupAddr("211.136.192.6")//语句2
	////names, err := net.LookupAddr("192.168.1.102")//语句3//因为我本机没做什么域名解析到地址，所以无法试验
	////names, err := net.LookupAddr("192.168.170.1")//语句4
	//check_err_net(err)
	//fmt.Println(names)
	//
	////输出：
	////	[ns6.gd.cnmobile.net.]
	////	[ns6.gd.cnmobile.net.]
	////	[]
	////	[]
	//
	////cname, err := net.LookupCNAME("")//这里的参数应该填写域名按照cname解析的某个ip地址，我这里无法验证，搁置
	////cname, err := net.LookupCNAME("baidu.com")
	////cname, err := net.LookupCNAME("souhu.com")
	//cname, err := net.LookupCNAME("qq.com")
	//check_err_net(err)
	//fmt.Println(cname)
	//
	//
	//// LookupHost looks up the given host using the local resolver.
	//// It returns a slice of that host's addresses.
	//// LookupHost使用本地解析器查找给定的主机。
	////返回该主机地址的切片。
	////这个api会请求本地的dns服务器来解析域名对应的ip地址切片，必须保证联网才能有本地的dns服务器，上同，下同
	//addrs, err := net.LookupHost("baidu.com")
	//check_err_net(err)
	//fmt.Println(addrs)
	//
	//
	//addrs, err = net.LookupHost("qq.com")
	//check_err_net(err)
	//fmt.Println(addrs)
	//
	//addrs, err = net.LookupHost("192.168.1.102")
	//check_err_net(err)
	//fmt.Println(addrs)
	//
	//ips, err := net.LookupIP("192.168.1.102")
	//check_err_net(err)
	//fmt.Println(ips)
	//
	//ips, err = net.LookupIP("baidu.com")//不应该填写ip地址，而是应该填写域名才对
	//check_err_net(err)
	//fmt.Println(ips)
	//
	//// LookupPort looks up the port for the given network and service.
	//// LookupPort查找给定网络和服务的端口。
	//port, err := net.LookupPort("baidu.com:tcp4", "80")
	//check_err_net(err)
	//fmt.Println(port)
	//
	//port, err = net.LookupPort(":tcp4", "8080")
	//check_err_net(err)
	//fmt.Println(port)
	//
	//
	//port, err = net.LookupPort("baidu.com:", "80")
	//check_err_net(err)
	//fmt.Println(port)
	//
	//
	//port, err = net.LookupPort("baidu.com", "80")
	//check_err_net(err)
	//fmt.Println(port)
	//
	//port, err = net.LookupPort("", "80")
	//check_err_net(err)
	//fmt.Println(port)
	//
	//port, err = net.LookupPort("baidu.com:tcp4", "")
	//check_err_net(err)
	//fmt.Println(port)
	////不知道怎么用，先搁置
	//
	//
	////// An MX represents a single DNS MX record.// MX代表单个DNS MX记录。
	////type MX struct {
	////	Host string
	////	Pref uint16
	////}
	//
	//// LookupMX returns the DNS MX records for the given domain name sorted by preference.
	//// LookupMX返回给定域名的DNS MX记录（按首选项排序）。
	//mxes, err := net.LookupMX("qq.com")
	//check_err_net(err)
	//fmt.Printf("%#v\n",mxes)
	//for _, v := range mxes {
	//	fmt.Println()
	//	fmt.Printf("MX:%#v\n",v)
	//	fmt.Printf("MX.Host:%#v\n",v.Host)
	//	fmt.Printf("MX.Pref:%#v\n",v.Pref)
	//
	//}
	//
	//
	//
	////// An NS represents a single DNS NS record.// NS代表单个DNS NS记录。
	////type NS struct {
	////	Host string
	////}
	//
	//// LookupNS returns the DNS NS records for the given domain name.
	//// LookupNS返回给定域名的DNS NS记录。
	//ns, err := net.LookupNS("qq.com")
	//check_err_net(err)
	//fmt.Printf("%#v\n",ns)
	//
	//for _, v := range ns {
	//	fmt.Println()
	//	fmt.Printf("NS:%#v\n",v)
	//	fmt.Printf("NS.Host:%#v\n",v.Host)
	//
	//}
	//
	//
	//
	//// LookupTXT returns the DNS TXT records for the given domain name.
	//// LookupTXT返回给定域名的DNS TXT记录。
	//DNSTXTStrings, err := net.LookupTXT("qq.com")
	//check_err_net(err)
	//fmt.Printf("%q\n",DNSTXTStrings)
	//
	//
	//DNSTXTStrings, err = net.LookupTXT("baidu.com")
	//check_err_net(err)
	//fmt.Printf("%q\n",DNSTXTStrings)
	//
	//
	//
	//fmt.Println()
	////
	//////SRV记录是DNS服务器的数据库中支持的一种资源记录的类型，它记录了哪台计算机提供了哪个服务这么一个简单的信息。
	////
	//////// An SRV represents a single DNS SRV record.// SRV表示单个DNS SRV记录。
	//////type SRV struct {
	//////	Target   string
	//////	Port     uint16
	//////	Priority uint16
	//////	Weight   uint16
	//////}
	////
	////
	////// LookupSRV tries to resolve an SRV query of the given service,
	////// protocol, and domain name. The proto is "tcp" or "udp".
	////// The returned records are sorted by priority and randomized
	////// by weight within a priority.
	//////
	////// LookupSRV constructs the DNS name to look up following RFC 2782.
	////// That is, it looks up _service._proto.name. To accommodate services
	////// publishing SRV records under non-standard names, if both service
	////// and proto are empty strings, LookupSRV looks up name directly.
	////// LookupSRV尝试解析给定服务，协议和域名的SRV查询。 原型是“ tcp”或“ udp”。
	//////返回的记录按优先级排序，并在优先级内按权重随机分配。
	////// LookupSRV构造DNS名称以遵循RFC 2782进行查找。
	//////也就是说，它查找_service._proto.name。 为了容纳以非标准名称发布SRV记录的服务，如果service和proto均为空字符串，则LookupSRV直接查找名称。
	////s, srvs, err := net.LookupSRV("80", "tcp4", "qq.com")
	////check_err_net(err)
	////fmt.Println(s,srvs)
	////
	//////不知道怎么用，先搁置
	//
	////我的主机上面的输出（每个人主机不一样应该）：
	////	[ns6.gd.cnmobile.net.]
	////	qq.com.
	////	[39.156.69.79 220.181.38.148]
	////	[58.250.137.36 125.39.52.26 58.247.214.47]
	////	[192.168.1.102]
	////	[192.168.1.102]
	////	[39.156.69.79 220.181.38.148]
	////	80
	////	8080
	////	80
	////	80
	////	80
	////	0
	////	[]*net.MX{(*net.MX)(0xc00005c640), (*net.MX)(0xc00005c620), (*net.MX)(0xc00005c600)}
	////
	////	MX:&net.MX{Host:"mx3.qq.com.", Pref:0xa}
	////	MX.Host:"mx3.qq.com."
	////	MX.Pref:0xa
	////
	////	MX:&net.MX{Host:"mx2.qq.com.", Pref:0x14}
	////	MX.Host:"mx2.qq.com."
	////	MX.Pref:0x14
	////
	////	MX:&net.MX{Host:"mx1.qq.com.", Pref:0x1e}
	////	MX.Host:"mx1.qq.com."
	////	MX.Pref:0x1e
	////	[]*net.NS{(*net.NS)(0xc00004c440), (*net.NS)(0xc00004c450), (*net.NS)(0xc00004c460), (*net.NS)(0xc00004c470)}
	////
	////	NS:&net.NS{Host:"ns2.qq.com."}
	////	NS.Host:"ns2.qq.com."
	////
	////	NS:&net.NS{Host:"ns4.qq.com."}
	////	NS.Host:"ns4.qq.com."
	////
	////	NS:&net.NS{Host:"ns3.qq.com."}
	////	NS.Host:"ns3.qq.com."
	////
	////	NS:&net.NS{Host:"ns1.qq.com."}
	////	NS.Host:"ns1.qq.com."
	////	["v=spf1 include:spf.mail.qq.com -all"]
	////	["google-site-verification=GHb98-6msqyx_qqjGl5eRatD3QTHyVB6-xQ3gJB5UwM" "v=spf1 include:spf1.baidu.com include:spf2.baidu.com include:spf3.baidu.com a mx ptr -all"]




	////注释掉上面的代码
	//fmt.Println()
	//fmt.Println("----------net包下的函数之net.Pipe()函数建立socket间无缓存全双工连接通信--------------")
	////可以本机两个不同的socket之间通信，应该不可以在不同的主机之间的socket之间的通信
	//
	////type pipe struct {
	////	wrMu sync.Mutex // Serialize Write operations//序列化写操作
	////
	////	// Used by local Read to interact with remote Write.
	////	// Successful receive on rdRx is always followed by send on rdTx.
	////	//由本地读取用于与远程写入交互。
	////	//总是往rdTx上塞值，紧跟着在rdRx上进行读取
	////	rdRx <-chan []byte
	////	rdTx chan<- int
	////
	////	// Used by local Write to interact with remote Read.
	////	// Successful send on wrTx is always followed by receive on wrRx.
	////	//由本地Write用于与远程Read交互。
	////	//总是在wrRx上取值，紧跟着在wrTx上塞值
	////	wrTx chan<- []byte
	////	wrRx <-chan int
	////
	////	once       sync.Once // Protects closing localDone//保护关闭localDone
	////	localDone  chan struct{}//本地连接关闭与否，
	////	remoteDone <-chan struct{}//远程的连接关闭与否，
	////
	////	readDeadline  pipeDeadline//读操作的生命周期
	////	writeDeadline pipeDeadline//写操作的生命周期
	////}
	////这个pipe结构体实现了conn接口
	//
	//		//// pipeDeadline is an abstraction for handling timeouts.
	//		//// pipeDeadline是用于处理超时的抽象。
	//		//type pipeDeadline struct {
	//		//	mu     sync.Mutex // Guards timer and cancel//保护计时器并取消
	//		//	timer  *time.Timer	//生命周期时间到达与否可以通过该对象来进行通知
	//		//	cancel chan struct{} // Must be non-nil//必须为非零，读或者写入操作取消与否
	//		//}
	//
	//
	//
	//// Pipe creates a synchronous, in-memory, full duplex
	//// network connection; both ends implement the Conn interface.
	//// Reads on one end are matched with writes on the other,
	//// copying data directly between the two; there is no internal
	//// buffering.
	//// Pipe创建一个同步的内存中全双工网络连接； 两端均实现Conn接口。
	////一端的读取与另一端的写入匹配，直接在两端之间复制数据； 没有内部缓冲。
	//conn1, conn2 := net.Pipe()
	//defer conn1.Close()
	//defer conn2.Close()
	//var connSlice=[]net.Conn{conn1, conn2 }
	////fmt.Println("conn1, conn2:",conn1, conn2)
	//
	//for _, v := range connSlice {
	//	fmt.Printf("%#v\n",v)
	//	fmt.Printf("v.LocalAddr():%#v\n",v.LocalAddr())
	//	fmt.Printf("v.RemoteAddr():%#v\n",v.RemoteAddr())
	//	fmt.Println()
	//}
	//
	////输出：
	////	&net.pipe{wrMu:sync.Mutex{state:0, sema:0x0},
	////		rdRx:(<-chan []uint8)(0xc000012180),//从哪里读取数据，请对比下面第2次循环的信息
	////		rdTx:(chan<- int)(0xc000012240),
	////		wrTx:(chan<- []uint8)(0xc0000121e0),//数据写到哪里去，请对比下面第2次循环的信息
	////		wrRx:(<-chan int)(0xc0000122a0),
	////		once:sync.Once{done:0x0, m:sync.Mutex{state:0, sema:0x0}},
	////		localDone:(chan struct {})(0xc000012300),
	////		remoteDone:(<-chan struct {})(0xc000012360),
	////		readDeadline:net.pipeDeadline{mu:sync.Mutex{state:0, sema:0x0}, timer:(*time.Timer)(nil), cancel:(chan struct {})(0xc0000123c0)},
	////		writeDeadline:net.pipeDeadline{mu:sync.Mutex{state:0, sema:0x0}, timer:(*time.Timer)(nil), cancel:(chan struct {})(0xc000012420)}}
	////
	////	&net.pipe{wrMu:sync.Mutex{state:0, sema:0x0},
	////		rdRx:(<-chan []uint8)(0xc0000121e0),//从哪里读取数据，请对比上面第一次循环的信息
	////		rdTx:(chan<- int)(0xc0000122a0),
	////		wrTx:(chan<- []uint8)(0xc000012180),//数据写到哪里去，请对比上面第一次循环的信息
	////		wrRx:(<-chan int)(0xc000012240),
	////		once:sync.Once{done:0x0, m:sync.Mutex{state:0, sema:0x0}},
	////		localDone:(chan struct {})(0xc000012360),
	////		remoteDone:(<-chan struct {})(0xc000012300),
	////		readDeadline:net.pipeDeadline{mu:sync.Mutex{state:0, sema:0x0}, timer:(*time.Timer)(nil), cancel:(chan struct {})(0xc000012480)},
	////		writeDeadline:net.pipeDeadline{mu:sync.Mutex{state:0, sema:0x0}, timer:(*time.Timer)(nil), cancel:(chan struct {})(0xc0000124e0)}}
	//
	//
	//var wg sync.WaitGroup
	//
	//wg.Add(1)
	////并发读取
	//go func() {
	//	rd_by:=make([]byte,100)//这里给的cap必须是足够大的。不够容量会报错！
	//	i, err := conn2.Read(rd_by)
	//	check_err_net(err)
	//	fmt.Println("conn2接收到conn1传递过来的字节数为：",i)
	//	fmt.Println("conn2接收到conn1传递过来的字节数据为：",rd_by)
	//	fmt.Println("conn2接收到conn1传递过来的字节转字符串数据为：",string(rd_by))
	//	wg.Add(-1)
	//}()
	//
	//wg.Add(1)
	////并发写入
	//go func() {
	//	n, err := conn1.Write([]byte("====这是conn1写给conn2的信息"))
	//	check_err_net(err)
	//	fmt.Println("====conn1写给conn2的字节数为：",n)
	//	wg.Add(-1)
	//}()
	//
	//wg.Wait()
	//
	////输出：
	////	====conn1写给conn2的字节数为： 35
	////	conn2接收到conn1传递过来的字节数为： 35
	////	conn2接收到conn1传递过来的字节数据为： [61 61 61 61 232 191 153 230 152 175 99 111 110 110 49 229 134 153 231 187 153 99 111
	////	110 110 50 231 154 132 228 191 161 230 129 175 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	////	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	////	conn2接收到conn1传递过来的字节转字符串数据为： ====这是conn1写给conn2的信息






	////注释掉上面的代码
	//fmt.Println()
	//fmt.Println("----------net包下的函数之net.ListenPacket()函数监听数据包--------------")
	//
	//
	////这里我就不再像上面的那样新开一个文件来弄服务器了，直接在下面新开一个新的g程来作为服务器，而main程则当做请求客户端
	//go func() {
	//
	//	//// PacketConn is a generic packet-oriented network connection.
	//	////
	//	//// Multiple goroutines may invoke methods on a PacketConn simultaneously.
	//	////PacketConn是通用的面向数据包的网络连接。
	//	////多个goroutine可以同时调用PacketConn上的方法。
	//	////其实跟其他的conn的文档几乎一样的！其他的很多个conn接口或者结构体都组合（或者成为实现）这个接口
	//	//type PacketConn interface {
	//	//	// ReadFrom reads a packet from the connection,
	//	//	// copying the payload into p. It returns the number of
	//	//	// bytes copied into p and the return address that
	//	//	// was on the packet.
	//	//	// It returns the number of bytes read (0 <= n <= len(p))
	//	//	// and any error encountered. Callers should always process
	//	//	// the n > 0 bytes returned before considering the error err.
	//	//	// ReadFrom can be made to time out and return
	//	//	// an Error with Timeout() == true after a fixed time limit;
	//	//	// see SetDeadline and SetReadDeadline.
	//	//	//ReadFrom从连接读取数据包，将有效负载复制到p中。 它返回复制到p中的字节数以及该数据包上的返回地址。
	//	//	//它返回读取的字节数（0 <= n <= len（p））和遇到的任何错误。 在考虑错误err之前，调用者应始终处理返回的n> 0个字节。
	//	//	//可以使ReadFrom超时并在固定的时间限制后使用Timeout（）== true返回错误； 请参见SetDeadline和SetReadDeadline。
	//	//	ReadFrom(p []byte) (n int, addr Addr, err error)
	//	//
	//	//	// WriteTo writes a packet with payload p to addr.
	//	//	// WriteTo can be made to time out and return
	//	//	// an Error with Timeout() == true after a fixed time limit;
	//	//	// see SetDeadline and SetWriteDeadline.
	//	//	// On packet-oriented connections, write timeouts are rare.
	//	//	//WriteTo将有效负载为p的数据包写入addr。
	//	//	//在固定的时间限制后，可以使WriteTo超时并使用Timeout（）== true返回错误； 请参见SetDeadline和SetWriteDeadline。
	//	//	//在面向数据包的连接上，写超时很少见。
	//	//	WriteTo(p []byte, addr Addr) (n int, err error)
	//	//
	//	//	// Close closes the connection.
	//	//	// Any blocked ReadFrom or WriteTo operations will be unblocked and return errors.
	//	//	//Close关闭连接。
	//	//	//任何阻塞的ReadFrom或WriteTo操作将被解除阻塞并返回错误。
	//	//	Close() error
	//	//
	//	//	// LocalAddr returns the local network address.
	//	//	// LocalAddr返回本地网络地址。
	//	//	LocalAddr() Addr
	//	//
	//	//	// SetDeadline sets the read and write deadlines associated
	//	//	// with the connection. It is equivalent to calling both
	//	//	// SetReadDeadline and SetWriteDeadline.
	//	//	//
	//	//	// A deadline is an absolute time after which I/O operations
	//	//	// fail with a timeout (see type Error) instead of
	//	//	// blocking. The deadline applies to all future and pending
	//	//	// I/O, not just the immediately following call to ReadFrom or
	//	//	// WriteTo. After a deadline has been exceeded, the connection
	//	//	// can be refreshed by setting a deadline in the future.
	//	//	//
	//	//	// An idle timeout can be implemented by repeatedly extending
	//	//	// the deadline after successful ReadFrom or WriteTo calls.
	//	//	//
	//	//	// A zero value for t means I/O operations will not time out.
	//	//	//SetDeadline设置与连接关联的读写期限。 这等效于调用SetReadDeadline和SetWriteDeadline。
	//	//	//截止期限是一个绝对时间，在该绝对时间之后，I / O操作将因超时（请参阅错误类型）而不是阻塞而失败。 截止日期适用于所有将来和未决的I / O，而不仅仅是紧接在其后的ReadFrom或WriteTo调用。 超过期限后，可以通过设置将来的期限来刷新连接。
	//	//	//空闲超时可以通过在成功执行ReadFrom或WriteTo调用后重复延长截止期限来实现。
	//	//	//t的值为零表示I / O操作不会超时。
	//	//	SetDeadline(t time.Time) error
	//	//
	//	//	// SetReadDeadline sets the deadline for future ReadFrom calls
	//	//	// and any currently-blocked ReadFrom call.
	//	//	// A zero value for t means ReadFrom will not time out.
	//	//	//SetReadDeadline设置将来的ReadFrom调用和任何当前阻止的ReadFrom调用的截止日期。 t的值为零表示ReadFrom不会超时。
	//	//	SetReadDeadline(t time.Time) error
	//	//
	//	//	// SetWriteDeadline sets the deadline for future WriteTo calls
	//	//	// and any currently-blocked WriteTo call.
	//	//	// Even if write times out, it may return n > 0, indicating that
	//	//	// some of the data was successfully written.
	//	//	// A zero value for t means WriteTo will not time out.
	//	//	//SetWriteDeadline设置将来的WriteTo调用和任何当前阻止的WriteTo调用的截止日期。
	//	//	//即使写入超时，它也可能返回n> 0，这表明某些数据已成功写入。
	//	//	//t的值为零表示WriteTo将不会超时。
	//	//	SetWriteDeadline(t time.Time) error
	//	//}
	//
	//
	//
	//
	//
	//	// ListenPacket announces on the local network address.
	//	//
	//	// The network must be "udp", "udp4", "udp6", "unixgram", or an IP
	//	// transport. The IP transports are "ip", "ip4", or "ip6" followed by
	//	// a colon and a literal protocol number or a protocol name, as in
	//	// "ip:1" or "ip:icmp".
	//	//
	//	// For UDP and IP networks, if the host in the address parameter is
	//	// empty or a literal unspecified IP address, ListenPacket listens on
	//	// all available IP addresses of the local system except multicast IP
	//	// addresses.
	//	// To only use IPv4, use network "udp4" or "ip4:proto".
	//	// The address can use a host name, but this is not recommended,
	//	// because it will create a listener for at most one of the host's IP
	//	// addresses.
	//	// If the port in the address parameter is empty or "0", as in
	//	// "127.0.0.1:" or "[::1]:0", a port number is automatically chosen.
	//	// The LocalAddr method of PacketConn can be used to discover the
	//	// chosen port.
	//	//
	//	// See func Dial for a description of the network and address
	//	// parameters.
	//	//	ListenPacket在本地网络地址上起作用。
	//	//	网络必须是"udp", "udp4", "udp6", "unixgram"或IP传输。 IP传输为"ip", "ip4", or "ip6"，后跟冒号和文字协议号或协议名称，如"ip:1" or "ip:icmp"。
	//	//	对于UDP和IP网络，如果address参数中的主机host为空或为字面量未指定的IP地址，ListenPacket将侦听本地系统的所有可用IP地址（组播IP地址除外）。
	//	//	要仅使用IPv4，请使用网络"udp4" 或者 "ip4:proto"。
	//	//	该地址address可以使用主机名，但是不建议使用该名称，因为它会为主机的IP地址之一最多创建一个侦听器。
	//	//	如果地址address参数中的端口为空或"0"，例如"127.0.0.1:" 或者 "[::1]:0"，则会自动选择端口号。
	//	//	PacketConn的LocalAddr方法可用于发现所选端口。
	//	//	有关网络network和地址address参数的说明，请参见func Dial。
	//
	//				//// ListenConfig contains options for listening to an address.
	//				//// ListenConfig包含用于侦听地址的选项。
	//				//type ListenConfig struct {
	//				//	// If Control is not nil, it is called after creating the network
	//				//	// connection but before binding it to the operating system.
	//				//	//
	//				//	// Network and address parameters passed to Control method are not
	//				//	// necessarily the ones passed to Listen. For example, passing "tcp" to
	//				//	// Listen will cause the Control function to be called with "tcp4" or "tcp6".
	//				//	//如果Control不为nil，则在创建网络连接之后但将其绑定到操作系统之前将其调用。
	//				//	//传递给Control方法的网络和地址参数不一定是传递给侦听的参数。 例如，将“ tcp”传递给Listen将导致使用“ tcp4”或“ tcp6”调用Control函数。
	//				//	Control func(network, address string, c syscall.RawConn) error
	//				//
	//				//	// KeepAlive specifies the keep-alive period for network
	//				//	// connections accepted by this listener.
	//				//	// If zero, keep-alives are enabled if supported by the protocol
	//				//	// and operating system. Network protocols or operating systems
	//				//	// that do not support keep-alives ignore this field.
	//				//	// If negative, keep-alives are disabled.
	//				//	//KeepAlive指定此侦听器接受的网络连接的保持活动时间。
	//				//	//如果为零，则在协议和操作系统支持的情况下启用保持活动状态。 不支持保持活动状态的网络协议或操作系统将忽略此字段。
	//				//	//如果为负，则保持活动被禁用。
	//				//	KeepAlive time.Duration
	//				//}
	//
	//
	//	//PacketConn, e := net.ListenPacket("tcp4", "127.0.0.1:9889")
	//	//windows不允许发或者接ip层的tcp或者udp等等原始套接字包，如果你要测试，请到linux系统测试
	//	//关于原始套接字的构造我暂时不讲
	//	//PacketConn, e := net.ListenPacket("ip4:17", "127.0.0.1")
	//	PacketConn, e := net.ListenPacket("udp4", "127.0.0.1:9889")
	//	check_err_net(e)
	//	defer PacketConn.Close()
	//
	//	//ReadFrom文档请看该g程最上面接口的文档
	//	rd_by:=make([]byte,100)
	//	n, raddr, e := PacketConn.ReadFrom(rd_by)
	//	check_err_net(e)
	//	fmt.Println("服务端接收到了的字节数为：",n)
	//	fmt.Println("客户端的地址为：",raddr)
	//	fmt.Println("服务端接收到了的字节数据为：",rd_by)
	//	fmt.Println("服务端接收到了的字节转字符串数据为：",string(rd_by))
	//	fmt.Println("接收完毕，准备发送响应回去给客户端！！！")
	//
	//	i, e := PacketConn.WriteTo([]byte("服务器响应的内容信息2222"), raddr)
	//	check_err_net(e)
	//	fmt.Println("服务器响应的字节数为：",i)
	//	fmt.Println("服务器响应完毕！！！")
	//
	//}()
	//
	////conn, e := net.Dial("tcp4", "127.0.0.1:9889")
	////windows不允许发或者接ip层的tcp或者udp等等原始套接字包，如果你要测试，请到linux系统测试
	////conn, e := net.Dial("ip4:17", "127.0.0.1")
	//conn, e := net.Dial("udp4", "127.0.0.1:9889")
	//check_err_net(e)
	//defer conn.Close()
	////故意在客户端所有的输出打印信息的前面加上===，以便执行输出结果的查看！
	////关于原始套接字的构造我暂时不讲
	//fmt.Println("===conn.LocalAddr():",conn.LocalAddr())
	//fmt.Println("===conn.RemoteAddr():",conn.RemoteAddr())
	//i, e := conn.Write([]byte("===这是客户端发送过去的消息1111"))
	//check_err_net(e)
	//fmt.Println("===客户端已经发送给服务器端的字节数为：",i)
	//
	//
	//fmt.Println("===客户端发送完毕，准备接受服务器的响应！！！")
	//
	//rdBy:=make([]byte,100)
	//n, e := conn.Read(rdBy)
	//check_err_net(e)
	//
	//fmt.Println("===读取到来自服务器端的字节数为：",n)
	//fmt.Println("===读取到来自服务器端的字节数据为：",rdBy)
	//fmt.Println("===读取到来自服务器端的字节转字符串数据为：",string(rdBy))
	////输出：
	////	===conn.LocalAddr(): 127.0.0.1:50933，这个地址是系统随机给的！我们不必关心他！
	////	===conn.RemoteAddr(): 127.0.0.1:9889
	////	===客户端已经发送给服务器端的字节数为： 43
	////	===客户端发送完毕，准备接受服务器的响应！！！
	////	服务端接收到了的字节数为： 43
	////	客户端的地址为： 127.0.0.1:50933
	////	服务端接收到了的字节数据为： [61 61 61 232 191 153 230 152 175 229 174 162 230 136 183 231 171 175 229 143
	////	145 233 128 129 232 191 135 229 142 187 231 154 132 230 182 136 230 129 175 49 49 49 49 0 0 0 0 0 0 0
	////	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	////	服务端接收到了的字节转字符串数据为： ===这是客户端发送过去的消息1111
	////	接收完毕，准备发送响应回去给客户端！！！
	////	服务器响应的字节数为： 34
	////	===读取到来自服务器端的字节数为： 34
	////	服务器响应完毕！！！
	////	===读取到来自服务器端的字节数据为： [230 156 141 229 138 161 229 153 168 229 147 141 229 186 148 231 154
	////	132 229 134 133 229 174 185 228 191 161 230 129 175 50 50 50 50 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	////	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	////	===读取到来自服务器端的字节转字符串数据为： 服务器响应的内容信息2222





	//如果你看到了这里，那么下面的都可以不看了，下面主要是多播，广播的代码，目前还不知道怎么写！先搁置！
	//注释掉上面的代码
	fmt.Println("如果你看到了这里，那么下面的都可以不看了，下面主要是多播，广播的代码，目前还不知道怎么写！先搁置！")
	fmt.Println("----------net包下的函数之net.ListenMulticastUDP()函数监听多播数据包--------------")



	// ListenMulticastUDP acts like ListenPacket for UDP networks but
	// takes a group address on a specific network interface.
	//
	// The network must be a UDP network name; see func Dial for details.
	//
	// ListenMulticastUDP listens on all available IP addresses of the
	// local system including the group, multicast IP address.
	// If ifi is nil, ListenMulticastUDP uses the system-assigned
	// multicast interface, although this is not recommended because the
	// assignment depends on platforms and sometimes it might require
	// routing configuration.
	// If the Port field of gaddr is 0, a port number is automatically
	// chosen.
	//
	// ListenMulticastUDP is just for convenience of simple, small
	// applications. There are golang.org/x/net/ipv4 and
	// golang.org/x/net/ipv6 packages for general purpose uses.
	//	ListenMulticastUDP的行为类似于UDP网络的ListenPacket，但在特定的网络接口上使用组地址。
	//	该网络必须是UDP网络名称； 有关详细信息，请参见func Dial。
	//	ListenMulticastUDP侦听本地系统的所有可用IP地址，包括组，组播IP地址。
	//	如果ifi为nil，则ListenMulticastUDP使用系统分配的多播接口，尽管不建议这样做，因为分配取决于平台，有时可能需要路由配置。
	//	如果gaddr的端口字段为0，则会自动选择一个端口号。
	//	ListenMulticastUDP只是为了方便简单的小型应用程序。 有通用用途的golang.org/x/net/ipv4和golang.org/x/net/ipv6软件包。


	//224.0.0.0~239.255.255.255的地址使用要求被限制在特定的多播域内

	//读取多播端,新开g程
	//往这个地址可以读取多播数据224.0.0.254:9981
	go func() {

		//如果第二参数为nil,它会使用系统指定多播接口，但是不推荐这样使用
		//224.0.0.0~239.255.255.255的地址使用要求被限制在特定的多播域内
		addr, err := net.ResolveUDPAddr("udp", "224.0.0.254:9981")
		check_err_net(err)

		//interfaces, err := net.Interfaces()
		////interfaces, err := net.InterfaceAddrs()
		//check_err_net(err)
		//fmt.Printf("%+v\n",interfaces)

		Interface, err := net.InterfaceByName("以太网")
		check_err_net(err)
		//listener, err := net.ListenMulticastUDP("udp", Interface, addr)
		listener, err := net.ListenMulticastUDP("udp", Interface, addr)
		check_err_net(err)

		fmt.Println("读取多播的地址为:", listener.LocalAddr())
		fmt.Println("发送的多播地址为:", listener.RemoteAddr())
		data := make([]byte, 1024)
		for {
			fmt.Println("正在监听读取。。。。")
			//不知道为什么无法读取到，先搁置吧！
			n, err := listener.Read(data)
			check_err_net(err)
			fmt.Println("接收到的多播字节数为:", n)
			fmt.Println("接收到的多播字节数据为:", data)
			fmt.Println("接收到的多播字节转字符串数据为:", string(data))
		}

	}()



	//发送多播端
	//往这个地址可以发送多播数据224.0.0.254:9981
	//224.0.0.0~239.255.255.255的地址使用要求被限制在特定的多播域内
	ip := net.ParseIP("224.0.0.254")
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}//让本机系统自动选择一个端口
	dstAddr := &net.UDPAddr{IP: ip, Port: 9981}
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	check_err_net(err)
	defer conn.Close()
	i, err := conn.Write([]byte("hello"))
	check_err_net(err)
	fmt.Println("===客户端发送了udp多播的字节数为：",i)
	fmt.Println("===客户端发送了udp多播的字节转字符串数据为：","hello")
	fmt.Println("===发送的多播的源地址为:", conn.LocalAddr())
	fmt.Println("===发送的多播的目的地址为:", conn.RemoteAddr())


	time.Sleep(1e9)


	//下面并不是上面代码的输出：
	//	[{Index:18
	//	MTU:65536
	//	Name:Npcap Loopback Adapter
	//	HardwareAddr:02:00:4c:4f:4f:50
	//	Flags:up|broadcast|multicast}
	//
	//	{Index:5
	//	MTU:1500
	//	Name:以太网
	//	HardwareAddr:00:f1:f3:8f:40:9a
	//	Flags:up|broadcast|multicast}
	//
	//	{Index:12
	//	MTU:1500
	//	Name:VMware Network Adapter VMnet1
	//	HardwareAddr:00:50:56:c0:00:01
	//	Flags:up|broadcast|multicast}
	//
	//	{Index:10
	//	MTU:1500
	//	Name:VMware Network Adapter VMnet8
	//	HardwareAddr:00:50:56:c0:00:08
	//	Flags:up|broadcast|multicast}
	//
	//	{Index:16
	//	MTU:1500
	//	Name:以太网 3
	//	HardwareAddr:00:ff:81:59:68:62
	//	Flags:broadcast|multicast}
	//
	//	{Index:20
	//	MTU:1500
	//	Name:以太网 4
	//	HardwareAddr:00:ff:f4:40:2e:17
	//	Flags:broadcast|multicast}
	//
	//	{Index:3
	//	MTU:1500
	//	Name:以太网 2
	//	HardwareAddr:00:ff:0c:82:d6:8c
	//	Flags:broadcast|multicast}
	//
	//	{Index:1
	//	MTU:-1
	//	Name:Loopback Pseudo-Interface 1
	//	HardwareAddr:
	//	Flags:up|loopback|multicast}]



	//[fe80::9c8a:2936:7699:9434/64
	//169.254.148.52/16
	//fe80::69af:7253:ebb9:bacd/64
	//192.168.1.102/24
	//fe80::4491:928a:684:d057/64
	//192.168.170.1/24
	//fe80::cd27:9f16:5158:406c/64
	//192.168.63.1/24
	//fe80::f468:5294:7e:f7d0/64
	//169.254.247.208/16
	//198.18.0.0/15
	//fe80::f1b6:1cf0:389b:5f0d/64
	//169.254.95.13/16
	//fe80::10ae:422f:5a17:8f82/64
	//169.254.143.130/16
	//::1/128
	//127.0.0.1/8]






	//// ListenUnixgram acts like ListenPacket for Unix networks.
	////
	//// The network must be "unixgram".
	////	ListenUnixgram的行为类似于Unix网络的ListenPacket。
	////	网络必须是“ unixgram”。
	//net.ListenUnixgram()

	//除此之外，还有几个api暂时没写，有以下：
	net.FileConn()
	net.FileListener()
	net.FilePacketConn()

	fmt.Println("本报除了几个api（5,6个）没写之外，其他已经写完，卒！")
}








//func readICMP(conn *net.IPConn)  {
//	//var by =make([]byte,512)
//	//err := binary.Read(conn, binary.BigEndian, &by)
//	//if err != nil {
//	//	fmt.Println(err.Error())
//	//	return
//	//}
//	//fmt.Println("read success:",by)
//	//不知道为什么无法read，先搁置
//}

//这些代码注释掉不是因为不重要，而是他的调用代码也被注释掉了！在你执行它的调用代码时候需要方法下面的函数的注释！上同
//type ICMP struct {
//	Type        uint8
//	Code        uint8
//	Checksum    uint16
//	Identifier  uint16
//	SequenceNum uint16
//}
//
////验证协议校验和，同时返回值（检验和）要写入协议中去一起发送的！
//func CheckSum(data []byte) uint16 {
//	var (
//		sum    uint32
//		length int = len(data)
//		index  int
//	)
//	for length > 1 {
//		sum += uint32(data[index])<<8 + uint32(data[index+1])
//		index += 2
//		length -= 2
//	}
//	if length > 0 {
//		sum += uint32(data[index])
//	}
//	sum += (sum >> 16)
//
//	return uint16(^sum)
//}

func check_err_net(err error) {
	if err != nil {
		//fmt.Fprintln(os.Stderr,err)
		//上面的这种方式会导致输出顺序不确定，虽然他可以输出红色的字体，但是由于顺序不确定，我们不采用他！
		fmt.Println("出错了，错误信息为：", err)
		panic(err)
	}
}
