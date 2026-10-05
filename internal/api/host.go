package api

import (
	"log"
	"net"
	"os"
)

// HostInfo identifies the machine (or container) serving the request.
type HostInfo struct {
	Hostname    string   `json:"hostname"`
	IPAddresses []string `json:"ip_addresses"`
}

// host is resolved once at startup; it does not change while the app runs.
var host = loadHostInfo()

func loadHostInfo() HostInfo {
	hostname, err := os.Hostname()
	if err != nil {
		log.Printf("get hostname: %v", err)
		hostname = "unknown"
	}

	ips := []string{}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Printf("get interface addresses: %v", err)
	}
	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if !ok || ipNet.IP.IsLoopback() || ipNet.IP.IsLinkLocalUnicast() {
			continue
		}
		ips = append(ips, ipNet.IP.String())
	}

	return HostInfo{Hostname: hostname, IPAddresses: ips}
}
