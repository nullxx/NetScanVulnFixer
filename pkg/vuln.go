package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/t94j0/nmap"
)

type VulnsJSON []struct {
	Service    string `json:"Service"`
	Version    string `json:"Version"`
	Port       string `json:"Port"`
	Metasploit struct {
		Path     string `json:"path"`
		Doc      string `json:"doc"`
		Tutorial struct {
			Commands []struct {
				Cmd  string `json:"cmd"`
				Type string `json:"type"`
			} `json:"commands"`
		} `json:"tutorial"`
	} `json:"metasploit"`
	Fix string `json:"fix"`
}

func GetVulns() (VulnsJSON, error) {
	// read json file "vulns.json"
	jsonFile, err := os.Open("vulns.json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	// read json file as byte array
	byteValue, _ := io.ReadAll(jsonFile)

	var vulns VulnsJSON
	json.Unmarshal(byteValue, &vulns)

	return vulns, nil
}

type PortInfo struct {
	ID          uint32
	State       string
	Service     Service
	Method      string
	HostAddress string
}

func MatchForVulns(portInfos []PortInfo) {
	vulns, err := GetVulns()
	if err != nil {
		panic(err)
	}

	for _, portInfo := range portInfos {
		fmt.Printf("- Host: %s Port: %d State: %s Service: %s(v%s) Method: %s\n", portInfo.HostAddress, portInfo.ID, portInfo.State, portInfo.Service.Product, portInfo.Service.Version, portInfo.Method)
		vulnFound := false
		for _, vuln := range vulns {
			if strings.HasPrefix(portInfo.Service.Product, vuln.Service) {
				if portInfo.Service.Version == vuln.Version {
					vulnFound = true
					fmt.Printf("\tVuln found: %s %s %s %s\n", portInfo.Service.Product, portInfo.Service.Version, portInfo.Service.Name, portInfo.Service.Product)
					fmt.Printf("\tMetasploit path: %s\n", vuln.Metasploit.Path)
					fmt.Printf("\tMetasploit doc: %s\n", vuln.Metasploit.Doc)
					fmt.Printf("\tMetasploit tutorial:\n")
					fmt.Printf("\t\tuse %s\n", vuln.Metasploit.Path)
					for _, command := range vuln.Metasploit.Tutorial.Commands {
						if command.Type == "rhost" {
							fmt.Printf("\t\t"+command.Cmd+"\n", portInfo.HostAddress)
						} else if command.Type == "rport" {
							fmt.Printf("\t\t"+command.Cmd+"\n", portInfo.ID)
						} else {
							fmt.Printf("\t\t" + command.Cmd + "\n")
						}
					}

					fmt.Printf("\nTo fix: %s\n", vuln.Fix)
				}

			}
		}

		if !vulnFound {
			fmt.Printf("\tNo vuln found\n")
		}
	}
}

func filterHostsForPorts(hosts map[string]nmap.Host, ports []uint32) []PortInfo {
	portInfos := []PortInfo{}

	for _, host := range hosts {
		for _, port := range host.Ports {
			for _, httpPort := range ports {
				if port.ID == httpPort {
					service := GetHostPortService(host, port.ID)
					portInfos = append(portInfos, PortInfo{ID: port.ID, State: port.State, Service: service, Method: port.Method, HostAddress: host.Address})
				}
			}
		}
	}

	return portInfos
}

func scanWebServers(hosts map[string]nmap.Host) {
	httpPorts := []uint32{80, 443, 8080, 8081, 8082}
	fmt.Printf("Scanning for web servers on ports: %v\n", httpPorts)

	portInfos := filterHostsForPorts(hosts, httpPorts)

	MatchForVulns(portInfos)
}

func ScanForVulns(hosts map[string]nmap.Host) {
	// scan for vulns
	scanWebServers(hosts)
}
