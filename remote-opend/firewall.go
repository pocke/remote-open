package main

import (
	"fmt"
	"net"
	"strings"
)

type Firewall struct {
	allows []*net.IPNet
}

func NewFirewall(ipStr string) (*Firewall, error) {
	IPs := strings.Split(ipStr, ",")
	f := &Firewall{
		allows: make([]*net.IPNet, 0, len(IPs)),
	}

	for _, i := range IPs {
		if !strings.Contains(i, "/") {
			if strings.Contains(i, ".") { // IPv4
				i += "/32"
			} else {
				i += "/128"
			}
		}

		_, mask, err := net.ParseCIDR(i)
		if err != nil {
			return nil, err
		}
		f.allows = append(f.allows, mask)
	}

	return f, nil
}

func (f *Firewall) IsAllowed(addr net.Addr) bool {
	ip, ok := addr.(*net.TCPAddr)
	if !ok {
		panic(fmt.Sprintf("%s(%T) is not an IP", addr, addr))
	}

	for _, m := range f.allows {
		masked := ip.IP.Mask(m.Mask)
		if masked.Equal(m.IP) {
			return true
		}
	}

	return false
}
