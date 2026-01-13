package ip

import "net"

func Local() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, addr := range addrs {
		ipnet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipnet.IP.IsLoopback() {
			continue
		}
		if ipnet.IP.To4() != nil {
			return ipnet.IP.String()
		}
	}

	return ""
}
