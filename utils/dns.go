package utils

import (
	"errors"
	"net"

	"github.com/ljcbaby/HDU-network-checker/log"
	"github.com/miekg/dns"
)

func Reslove(domin string, server string) (*net.IPAddr, error) {
	if server == "" {
		server = "223.5.5.5"
	}
	log.Logger.Sugar().Debugf("Reslove %s by %s", domin, server)
	c := dns.Client{}
	m := dns.Msg{}
	m.SetQuestion(domin, dns.TypeA)
	r, _, err := c.Exchange(&m, server+":53")
	if err != nil {
		return nil, err
	}
	log.Logger.Sugar().Debugf("%v", r)
	if len(r.Answer) == 0 {
		return nil, errors.New("no_answer")
	}
	return &net.IPAddr{IP: r.Answer[len(r.Answer)-1].(*dns.A).A}, nil
}
