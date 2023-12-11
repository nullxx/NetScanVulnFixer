package pkg

import (
	"fmt"
	"net"
	"os/exec"

	"encoding/xml"

	"github.com/t94j0/nmap"
)

var s nmap.Scan

func init() {
	s = nmap.Init()
}

type NmapRun struct {
	XMLName xml.Name `xml:"nmaprun"`
	Hosts   []Host   `xml:"host"`
}

type Host struct {
	Status Status `xml:"status"`
	Ports  Ports  `xml:"ports"`
}

type Status struct {
	State string `xml:"state,attr"`
}

type Ports struct {
	Port Port `xml:"port"`
}

type Port struct {
	ID      string  `xml:"portid,attr"`
	State   State   `xml:"state"`
	Service Service `xml:"service"`
}

type State struct {
	State string `xml:"state,attr"`
}

type Service struct {
	Name    string `xml:"name,attr"`
	Product string `xml:"product,attr"`
	Version string `xml:"version,attr"`
}

func DiscoverHosts(ip net.IP) (map[string]nmap.Host, error) {
	sc, err := nmap.Scan.SetHosts(s, ip.String()+"/24").Run()
	if err != nil {
		return nil, err
	}

	return sc.Hosts, nil
}

func Cmd(cmd string, shell bool) ([]byte, error) {

	if shell {
		return exec.Command("bash", "-c", cmd).Output()
	}
	return exec.Command(cmd).Output()
}

func GetHostPortService(host nmap.Host, port uint32) Service {
	cmdStr := fmt.Sprintf("/opt/homebrew/bin/nmap -oX - -sV -p %d %s", port, host.Address)

	xmlBytes, err := Cmd(cmdStr, true)
	if err != nil {
		panic(err)
	}

	var nmapRun NmapRun
	err = xml.Unmarshal(xmlBytes, &nmapRun)
	if err != nil {
		panic(err)
	}

	return nmapRun.Hosts[0].Ports.Port.Service
}
