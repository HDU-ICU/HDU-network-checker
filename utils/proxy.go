package utils

import "github.com/ljcbaby/HDU-network-checker/log"

func ProxyCheck() (bool, error) {
	log.Logger.Debug("Proxy check start")
	// ToDo: http proxy check
	log.Logger.Debug("http proxy check not implemented")

	// check TUN/TAP device by ping 198.18.0.1
	p, err := Ping("198.18.0.1")
	if err != nil {
		return false, err
	}

	if p == 0 {
		return true, nil
	}

	return false, nil
}
