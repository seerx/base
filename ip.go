package base

import (
	"encoding/binary"
	"net"
)

// IPType IP 地址类型定义
type IPType int

const (
	// IPA A 类
	IPA IPType = 1
	// IPB B 类
	IPB = 1 << 1
	// IPC C 类
	IPC = 1 << 2
	// IPR 保留 IP 地址，如 127.*, 0.0.0.0 等
	IPR = 1 << 3
)

// IPIsAType 判断 IP 地址是否是 A 类地址
func IPIsAType(IP net.IP) bool {
	return (IP[0] >= 1 && IP[0] < 126) ||
		(IP[0] == 126 && IP[1] == 0 && IP[2] == 0 && IP[3] == 0)
}

// IPIsBType 判断 IP 地址是否是 B 类地址
func IPIsBType(IP net.IP) bool {
	return IP[0] >= 128 && IP[0] <= 191
}

// IPIsCType 判断 IP 地址是否是 C 类地址
func IPIsCType(IP net.IP) bool {
	return IP[0] >= 192 && IP[0] <= 223
}

// IPIsReserve 判断是否保留 IP 地址
func IPIsReserve(IP net.IP) bool {
	return IP.IsLoopback() || IP.IsLinkLocalMulticast() || IP.IsLinkLocalUnicast()
}

// GetLocalIP 获取本机 IP 地址，适用于 IPv4
func GetLocalIP(ipType IPType) ([]string, error) {
	var ips []string
	if err := FindIPs(func(ipn *net.IPNet) error {
		ipv4 := ipn.IP.To4()
		if ipv4 != nil {
			match := false
			if ipType&IPA == IPA {
				match = IPIsAType(ipv4)
			}
			if !match && ipType&IPB == IPB {
				match = IPIsBType(ipv4)
			}
			if !match && ipType&IPC == IPC {
				match = IPIsCType(ipv4)
			}
			if !match && ipType&IPR == IPR {
				match = IPIsReserve(ipv4)
			}
			if match {
				ips = append(ips, ipn.IP.String())
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return ips, nil
}

// FindIPs 查找 IP
func FindIPs(fn func(ipn *net.IPNet) error) error {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return err
	}
	// var ips []string
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok {
			if err := fn(ipnet); err != nil {
				return err
			}
		}
	}
	return nil
}

/*GetIP 获取本机 IP 地址
 */
func GetIP() ([]string, error) {
	var ips []string

	if err := FindIPs(func(ipn *net.IPNet) error {
		if ipn.IP.To4() != nil {
			ips = append(ips, ipn.IP.String())
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return ips, nil
}

// IP2long 字符串 IP 地址转 long
func IP2long(ipstr string) uint32 {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}

// Long2IP long 转字符串 IP
func Long2IP(ipLong uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, ipLong)
	ip := net.IP(ipByte)
	return ip.String()
}
