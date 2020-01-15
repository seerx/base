package base

import (
	"encoding/binary"
	"fmt"
	"net"
)

/*GetIP 获取本机 IP 地址
 */
func GetIP() ([]string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	var ips []string
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP.String())
				//return ipnet.IP.String()
			}
		}
	}
	if len(ips) == 0 {
		return nil, fmt.Errorf("Unable to determine local IP address (non loopback)")
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
