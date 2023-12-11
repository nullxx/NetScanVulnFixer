package main

import (
	"fmt"
	"net"
	"netscanvulnfixer/main/pkg"
)

func main() {
	ip := net.IPv4(192, 168, 1, 1)
	fmt.Printf("Scaning for hosts on %s\n", ip.String())
	hosts, err := pkg.DiscoverHosts(ip)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Found %d hosts\n", len(hosts))

	for _, host := range hosts {
		for _, port := range host.Ports {
			println("Host:", host.Address, "Port:", port.ID, "State:", port.State, "Service:", port.Service, "Method:", port.Method)
		}
	}

	// hosts := map[string]nmap.Host{}
	// hosts["192.168.1.78"] = nmap.Host{
	// 	Address: "192.168.1.78",
	// 	Ports: []nmap.Port{
	// 		nmap.Port{ID: 22, State: "open", Service: "ssh", Method: "syn-ack"},
	// 		nmap.Port{ID: 80, State: "open", Service: "http", Method: "syn-ack"},
	// 		nmap.Port{ID: 111, State: "open", Service: "rpcbind", Method: "syn-ack"},
	// 		nmap.Port{ID: 8080, State: "open", Service: "http", Method: "syn-ack"},
	// 		nmap.Port{ID: 8081, State: "open", Service: "http", Method: "syn-ack"},
	// 		nmap.Port{ID: 8082, State: "open", Service: "http", Method: "syn-ack"},
	// 	},
	// }
	pkg.ScanForVulns(hosts)
}
