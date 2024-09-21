package utils

import (
	"net"

	"github.com/ljcbaby/HDU-network-checker/log"
)

func IfaceCheck() (int, error) {
	flag := 0

	interfaces, err := net.Interfaces()
	if err != nil {
		return flag, err
	}

	for _, iface := range interfaces {
		// if flags without up or lookback, skip
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		if flag == 0 {
			flag = 1
		}

		log.Logger.Sugar().Infof("Interface %s", iface.Name)
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
					flag += 1
				}
			}
		}
	}

	return flag, nil
}
