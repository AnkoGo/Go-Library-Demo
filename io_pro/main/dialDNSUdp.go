//转载请注明出处：http://blog.csdn.net/gophers/article/details/22942457

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strings"
	"time"
)

type DNSHeader struct {
	ID            uint16
	Flag          uint16
	QuestionCount uint16
	AnswerRRs     uint16 //RRs is Resource Records
	AuthorityRRs  uint16
	AdditionalRRs uint16
}

func (header *DNSHeader) SetFlag(QR uint16, OperationCode uint16, AuthoritativeAnswer uint16, Truncation uint16, RecursionDesired uint16, RecursionAvailable uint16, ResponseCode uint16) {
	header.Flag = QR<<15 + OperationCode<<11 + AuthoritativeAnswer<<10 + Truncation<<9 + RecursionDesired<<8 + RecursionAvailable<<7 + ResponseCode
}

type DNSQuery struct {
	QuestionType  uint16
	QuestionClass uint16
}

func ParseDomainName(domain string) []byte {
	//要将域名解析成相应的格式，例如：
	//"www.google.com"会被解析成"0x03www0x06google0x03com0x00"
	//就是长度+内容，长度+内容……最后以0x00结尾
	var (
		buffer   bytes.Buffer
		segments []string = strings.Split(domain, ".")
	)
	for _, seg := range segments {
		binary.Write(&buffer, binary.BigEndian, byte(len(seg)))
		binary.Write(&buffer, binary.BigEndian, []byte(seg))
	}
	binary.Write(&buffer, binary.BigEndian, byte(0x00))

	return buffer.Bytes()
}

func DialDnsUdp() {
	var (
		dns_header   DNSHeader
		dns_question DNSQuery
	)

	//填充dns首部
	dns_header.ID = 0xFFFF
	dns_header.SetFlag(0, 0, 0, 0, 1, 0, 0)
	dns_header.QuestionCount = 1
	dns_header.AnswerRRs = 0
	dns_header.AuthorityRRs = 0
	dns_header.AdditionalRRs = 0

	//填充dns查询首部
	dns_question.QuestionType = 1 //IPv4
	dns_question.QuestionClass = 1
	go func() {
		//同样不知道为什么不可以读取
		udpConn, e := net.ListenUDP("udp", nil)
		if e != nil {
			fmt.Println(e)
			return
		}
		by := make([]byte, 512)
		i, e := udpConn.Read(by)
		if e != nil {
			fmt.Println(e)
			return
		}
		fmt.Println("接收到的字节数为：", i)
		fmt.Println("接收到的字节数据为：", by)
		fmt.Println("接收到的字节数据为：", string(by))
	}()
	var (
		conn net.Conn
		err  error

		buffer bytes.Buffer
	)

	//DNS服务器的端口一般是53，dns的服务器ip地址你自己ipconfig查一下
	//别忘了DNS是基于UDP协议的
	if conn, err = net.Dial("udp", "120.196.165.24:53"); err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()

	//buffer中是我们要发送的数据，里面的内容是DNS首部+查询内容+DNS查询首部
	binary.Write(&buffer, binary.BigEndian, dns_header)
	binary.Write(&buffer, binary.BigEndian, ParseDomainName("www.baidu.com"))
	binary.Write(&buffer, binary.BigEndian, dns_question)
	fmt.Println(buffer.Bytes())
	fmt.Println(buffer.String())

	if _, err := conn.Write(buffer.Bytes()); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("send success.")
	//请自己抓包查看效果，这里查询的是baidu.com的dns地址
	time.Sleep(2e9)
}
