package checker

import (
	"os"

	"github.com/ljcbaby/HDU-network-checker/log"
	"github.com/ljcbaby/HDU-network-checker/utils"
)

func BasicCheck() {
	log.Logger.Info("Basic check start")

	// Check if the IP address is in the campus network
	iface, err := utils.IfaceCheck()
	if err != nil {
		log.Logger.Sugar().Errorf("获取网卡信息失败：%v", err)
	} else {
		if iface {
			log.Logger.Warn("IP 地址不在校园网内，如有路由器请忽略")
		}
	}

	// Check if proxy on
	proxy, err := utils.ProxyCheck()
	if err != nil {
		log.Logger.Sugar().Errorf("代理/虚拟网卡检查失败：%v", err)
	} else {
		if proxy {
			log.Logger.Warn("检测到代理/虚拟网卡，建议先关闭/禁用")
		}
	}

	// Check connection to BAS 10.150.0.1
	p, err := utils.Ping("10.150.0.1")
	if err != nil {
		log.Logger.Sugar().Errorf("Ping 失败：%v", err)
	} else {
		if p == 4 {
			log.Logger.Error("无法连接到 BAS，请检查IP配置或尝试重新进行物理连接")
			os.Exit(0)
		} else if p > 0 {
			log.Logger.Warn("到 BAS 的连接存在丢包")
		} else {
			log.Logger.Info("连接到 BAS 正常")
		}
	}

	// Check connection to AAA 192.168.112.97
	p, err = utils.Ping("192.168.112.97")
	if err != nil {
		log.Logger.Sugar().Errorf("Ping 失败：%v", err)
	} else {
		if p == 4 {
			log.Logger.Error("无法连接到 深澜认证")
			os.Exit(0)
		} else if p > 0 {
			log.Logger.Warn("到 深澜 的连接存在丢包")
		} else {
			log.Logger.Info("连接到 深澜 正常")
		}
	}

	c := 0

	// Check connection to DNS 210.32.32.1
	p, err = utils.Ping("210.32.32.1")
	if err != nil {
		log.Logger.Sugar().Errorf("Ping 失败：%v", err)
	} else {
		if p == 4 {
			log.Logger.Error("无法连接到 DNS")
			c++
		} else if p > 0 {
			log.Logger.Warn("到 DNS 的连接存在丢包")
		} else {
			log.Logger.Info("连接到 主DNS 正常")
		}
	}
}
