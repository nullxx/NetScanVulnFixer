package pkg

import (
	"net"

	"github.com/t94j0/nmap"
)

var s nmap.Scan

func init() {
	s = nmap.Init()
}

func DiscoverHosts(ip net.IP) (map[string]nmap.Host, error) {
	sc, err := nmap.Scan.SetHosts(s, ip.String()+"/24").Run()
	if err != nil {
		return nil, err
	}

	return sc.Hosts, nil
}
