package utils

import (
	"net"

	"github.com/ljcbaby/HDU-network-checker/log"
)

func IfaceCheck() (bool, error) {
	flag := true

	interfaces, err := net.Interfaces()
	if err != nil {
		return false, err
	}

	for _, iface := range interfaces {
		log.Logger.Sugar().Infof("Interface %s", iface.Name)
		log.Logger.Sugar().Infof("  Interface status: %v", iface.Flags.String())
		log.Logger.Sugar().Infof("  Hardware addr: %s", iface.HardwareAddr.String())

		addrs, err := iface.Addrs()
		if err != nil {
			log.Logger.Sugar().Errorf("Failed to get addresses : %v", err)
			continue
		}

		for _, addr := range addrs {
			log.Logger.Sugar().Infof("  %s", addr.String())
			_, subnet, _ := net.ParseCIDR("10.150.0.0/16")
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if subnet.Contains(ipnet.IP) {
					flag = false
				}
			}
		}
	}

	return flag, nil
}
