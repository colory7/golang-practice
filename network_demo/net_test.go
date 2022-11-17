package network_demo

import (
	"fmt"
	"net"
	"testing"
)

func TestNet(t *testing.T) {
	//InterfaceAddrs 返回该系统的网络接口的地址列表。
	addr, _ := net.InterfaceAddrs()
	fmt.Println(addr)

	//Interfaces 返回该系统的网络接口列表
	interfaces, _ := net.Interfaces()
	fmt.Println(interfaces)

	//LookupAddr 查询某个地址，返回映射到该地址的主机名序列
	lt, _ := net.LookupAddr("www.alibaba.com")
	fmt.Println(lt)

	fmt.Println("==============================")

	//LookupCNAME函数查询name的规范DNS名（但该域名未必可以访问）。
	cname, _ := net.LookupCNAME("www.baidu.com")
	fmt.Println(cname)

	//LookupHost函数查询主机的网络地址序列。
	host, _ := net.LookupHost("www.baidu.com")
	fmt.Println(host)

	//LookupIP函数查询主机的ipv4和ipv6地址序列。
	ip, _ := net.LookupIP("www.baidu.com")
	fmt.Println(ip)
}

func TestNet2(t *testing.T) {
	//函数将host和port合并为一个网络地址。一般格式为"host:port"；如果host含有冒号或百分号，格式为"[host]:port"。
	//Ipv6的文字地址或者主机名必须用方括号括起来，如"[::1]:80"、"[ipv6-host]:http"、"[ipv6-host%zone]:80"。
	hp := net.JoinHostPort("127.0.0.1", "8080")
	fmt.Println(hp)

	//函数将格式为"host:port"、"[host]:port"或"[ipv6-host%zone]:port"的网络地址分割为host或ipv6-host%zone和port两个部分。
	shp, port, _ := net.SplitHostPort("127.0.0.1:8080")
	fmt.Println(shp, " _ ", port)

}
