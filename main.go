package main

import (
	"net"
	"netscanvulnfixer/main/pkg"
)

func main() {
	ip := net.IPv4(192, 168, 1, 1)
	hosts, err := pkg.DiscoverHosts(ip)
	if err != nil {
		panic(err)
	}

	for _, host := range hosts {
		for _, port := range host.Ports {
			println("Host:", host.Address, "Port:", port.ID, "State:", port.State, "Service:", port.Service, "Method:", port.Method)
		}
	}
	/*
		Host: 192.168.1.108 Port: 22 State: open
		Host: 192.168.1.108 Port: 80 State: open
		Host: 192.168.1.108 Port: 111 State: open
		Host: 192.168.1.108 Port: 443 State: open
		Host: 192.168.1.179 Port: 80 State: open
		Host: 192.168.1.179 Port: 9503 State: open
		Host: 192.168.1.105 Port: 22 State: open
		Host: 192.168.1.105 Port: 80 State: open
		Host: 192.168.1.105 Port: 111 State: open
		Host: 192.168.1.105 Port: 443 State: open
		Host: 192.168.1.105 Port: 49153 State: open
		Host: 192.168.1.106 Port: 22 State: open
		Host: 192.168.1.106 Port: 80 State: open
		Host: 192.168.1.106 Port: 111 State: open
		Host: 192.168.1.106 Port: 443 State: open
		Host: 192.168.1.107 Port: 22 State: open
		Host: 192.168.1.107 Port: 80 State: open
		Host: 192.168.1.107 Port: 111 State: open
		Host: 192.168.1.107 Port: 443 State: open
		Host: 192.168.1.200 Port: 1234 State: open
		Host: 192.168.1.200 Port: 8090 State: open
		Host: 192.168.1.1 Port: 22 State: open
		Host: 192.168.1.1 Port: 80 State: open
		Host: 192.168.1.1 Port: 443 State: open
		Host: 192.168.1.36 Port: 80 State: open
		Host: 192.168.1.36 Port: 443 State: open
		Host: 192.168.1.36 Port: 515 State: open
		Host: 192.168.1.36 Port: 631 State: open
		Host: 192.168.1.36 Port: 9100 State: open
		Host: 192.168.1.171 Port: 80 State: open
		Host: 192.168.1.192 Port: 53 State: open
		Host: 192.168.1.192 Port: 80 State: open
		Host: 192.168.1.192 Port: 443 State: open
	*/
}
