package main

import (
	"fmt"

	"github.com/seerx/base"
)

func main() {
	// c, _ := sshx.NewSFTP("", "", "", 11)
	// c.Close()
	ips, err := base.GetLocalIP(base.IPA | base.IPR)
	// ips, err := base.GetIP()
	if err != nil {
		panic(err)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
