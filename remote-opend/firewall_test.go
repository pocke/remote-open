package main

import (
	"net"
	"testing"
)

func TestNewFierwall(t *testing.T) {
	assert := func(ip string, ok bool) {
		_, err := NewFirewall(ip)
		if ok != (err == nil) {
			var msg string
			if ok {
				msg = "success"
			} else {
				msg = "failure"
			}
			t.Errorf("%s should %s, but not", ip, msg)
		}
	}

	assert("192.168.0.1", true)
	assert("192.168.0.1,192.168.1.0/24", true)
	assert("0.0.0.0/0,::/0", true)
	assert("a::b", true)
	assert("0.0.0.0/0, ::/0", false)
	assert("300.3.5.6", false)
}

func TestFirewallIsAllowed(t *testing.T) {
	assert := func(allow, ip string, ok bool) {
		f, err := NewFirewall(allow)
		if err != nil {
			t.Fatal(err)
		}
		addr := &net.TCPAddr{IP: net.ParseIP(ip)}

		if f.IsAllowed(addr) != ok {
			var msg string
			if ok {
				msg = "be allowed"
			} else {
				msg = "not be allowed"
			}
			t.Errorf("When filter is %s, %s should %s. But not.", allow, addr, msg)
		}
	}

	assert("192.168.0.1", "192.168.0.1", true)
	assert("192.168.0.1", "192.168.0.2", false)
	assert("192.168.0.0/24", "192.168.0.2", true)
	assert("192.168.0.0/24", "192.168.1.2", false)
	assert("172.0.0.0/16", "172.0.243.134", true)
	assert("172.0.0.0/16", "172.10.243.134", false)
	assert("0.0.0.0/0", "172.10.243.134", true)

	assert("192.168.0.1,192.168.0.2", "192.168.0.1", true)
	assert("192.168.0.1,192.168.0.2", "192.168.0.2", true)

	assert("2001:0db8:bd05:01d2:288a:1fc0:0001:10ee", "2001:0db8:bd05:01d2:288a:1fc0:0001:10ee", true)
	assert("::/0", "2001:0db8:bd05:01d2:288a:1fc0:0001:10ee", true)
}
