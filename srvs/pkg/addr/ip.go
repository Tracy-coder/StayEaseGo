package addr

import (
	"errors"
	"net"
)

func GetLocalIPv4Address() (string, error) {
	addr, err := net.InterfaceAddrs()
	// log.Debug("addr", addr)
	if err != nil {
		return "", err
	}

	for _, addr := range addr {
		ipNet, isIpNet := addr.(*net.IPNet)
		if isIpNet && !ipNet.IP.IsLoopback() {
			ipv4 := ipNet.IP.To4()
			if ipv4 != nil {
				return ipv4.String(), nil
			}
		}
	}
	return "", errors.New("not found ipv4 address")
}
