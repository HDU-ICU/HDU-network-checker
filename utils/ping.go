package utils

import (
	"time"

	"github.com/go-ping/ping"
	"github.com/ljcbaby/HDU-network-checker/log"
)

func Ping(host string) (int, error) {
	log.Logger.Sugar().Debugf("Ping %s", host)
	pinger, err := ping.NewPinger(host)
	if err != nil {
		return 0, err
	}

	pinger.Count = 4
	pinger.Interval = time.Millisecond * 200
	pinger.Timeout = time.Second * 1
	pinger.SetPrivileged(true)
	pinger.Run()

	stats := pinger.Statistics()
	res := stats.PacketsSent - stats.PacketsRecv

	log.Logger.Sugar().Debug("Ping result:")
	log.Logger.Sugar().Debugf("%+v", stats)

	return res, nil
}
