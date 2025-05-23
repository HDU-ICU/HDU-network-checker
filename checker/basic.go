package checker

import (
	"fmt"
	"net"
	"os"
	"strings"

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
		switch iface {
		case 0:
			log.Logger.Error("未找到有效网卡，请检查网络连接")
			exit()
		case 1:
			log.Logger.Warn("IP 地址不在校园网内，如有路由器请忽略")
		case 2:
			log.Logger.Info("IP 地址正确")
		default:
			log.Logger.Warn("存在多个校园网 IP 地址，请检查网络连接")
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
			log.Logger.Warn("无法连接到 BAS")
			_, err := utils.Get("http://connect.rom.miui.com/generate_204")
			if err != nil {
				if strings.Contains(err.Error(), "ERR_204StatusNoContent") {
					log.Logger.Error("您可能不是在使用校园网，请检查网络连接")
					exit()
				} else if strings.Contains(err.Error(), "no such host") {
					log.Logger.Sugar().Debugf("Curl 请求失败：%v", err)
					log.Logger.Error("请检查IP配置或尝试重新进行物理连接")
					exit()
				} else {
					log.Logger.Sugar().Errorf("Curl 请求失败：%v", err)
				}
			} else {
				log.Logger.Error("请检查IP配置或尝试重新进行物理连接")
				exit()
			}
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
			exit()
		} else if p > 0 {
			log.Logger.Warn("到 深澜 的连接存在丢包")
		} else {
			log.Logger.Info("连接到 深澜 正常")
		}
	}

	aaa := 0
	aaaAddr := net.IPAddr{IP: net.ParseIP("192.168.112.97")}

	// Check DNS resolve
	res, err := utils.Reslove("portal.hdu.edu.cn.", "210.32.32.1")
	if err != nil {
		if strings.Contains(err.Error(), "i/o timeout") ||
			strings.Contains(err.Error(), "unreachable network") {
			log.Logger.Sugar().Debugf("DNS 解析失败：%v", err)
			log.Logger.Error("DNS 解析失败")
		} else {
			log.Logger.Sugar().Errorf("DNS 解析失败：%v", err)
		}
		aaa = 1
	} else {
		if res.IP.String() != aaaAddr.IP.String() {
			log.Logger.Sugar().Errorf("DNS 解析错误：%v", res.IP.String())
			aaa = 1
		} else {
			log.Logger.Info("DNS 解析正常")
		}
	}

	baseUrl := "https://"
	if aaa == 1 {
		log.Logger.Debug("DNS 异常存在，后续访问认证页面使用 IP 进行")
		baseUrl += aaaAddr.IP.String()
	} else {
		baseUrl += "portal.hdu.edu.cn"
	}

	log.Logger.Sugar().Debugf("AAA 地址：%s", baseUrl)

	// Check AAA auth status
	api, err := utils.GetAAAInfo(baseUrl)
	if err != nil {
		log.Logger.Sugar().Errorf("认证状态获取失败：%v", err)
	} else {
		if api.Error == "not_online_error" {
			log.Logger.Error("深澜未认证，请打开下方链接进行认证")
			log.Logger.Sugar().Errorf("认证页面地址：%s", baseUrl)
			exit()
		} else {
			log.Logger.Sugar().Infof("%s 已认证 %s，在线设备数 %s", api.UserName, api.ProductsName, api.OnlineDeviceTotal)
		}
	}

	// ISP gateway check
	switch api.ProductsId {
	case "3": // 3: 联通
		p, err = utils.Ping("172.20.64.1")
		if err != nil {
			log.Logger.Sugar().Errorf("Ping 失败：%v", err)
		} else {
			if p == 4 {
				log.Logger.Warn("无法连接到 联通网关172.20.64.1")
			} else if p > 0 {
				log.Logger.Warn("到 联通网关172.20.64.1 的连接存在丢包")
			} else {
				log.Logger.Info("连接到 联通网关172.20.64.1 正常")
			}
		}
	case "4": // 4: 电信
		p, err = utils.Ping("60.176.40.1")
		if err != nil {
			log.Logger.Sugar().Errorf("Ping 失败：%v", err)
		} else {
			if p == 4 {
				log.Logger.Warn("无法连接到 电信网关60.176.40.1")
			} else if p > 0 {
				log.Logger.Warn("到 电信网关60.176.40.1 的连接存在丢包")
			} else {
				log.Logger.Info("连接到 电信网关60.176.40.1 正常")
			}
		}
	case "5": // 5: 移动
		p, err = utils.Ping("10.106.0.1")
		if err != nil {
			log.Logger.Sugar().Errorf("Ping 失败：%v", err)
		} else {
			if p == 4 {
				log.Logger.Warn("无法连接到 移动网关10.106.0.1")
			} else if p > 0 {
				log.Logger.Warn("到 移动网关10.106.0.1 的连接存在丢包")
			} else {
				log.Logger.Info("连接到 移动网关10.106.0.1 正常")
			}
		}
	}

	// Check connection to Internet
	p, err = utils.Ping("223.5.5.5")
	if err != nil {
		log.Logger.Sugar().Errorf("Ping 失败：%v", err)
	} else {
		if p == 4 {
			log.Logger.Error("无法连接到 阿里DNS主")
			log.Logger.Error("请检查代拨上线情况，或尝试重新绑定")
			log.Logger.Sugar().Errorf("认证页面地址：%s", baseUrl)
			exit()
		} else if p > 0 {
			log.Logger.Warn("到 阿里DNS主 的连接存在丢包")
		} else {
			log.Logger.Info("连接到 阿里DNS主 正常")
		}
	}

	// Check connection to Internet
	resp, err := utils.Get("http://connect.rom.miui.com/generate_204")
	if err != nil {
		if strings.Contains(err.Error(), "204") {
			log.Logger.Info("连接到外网正常")
		} else {
			log.Logger.Sugar().Errorf("连接到外网失败：%v", err)
		}
	} else {
		log.Logger.Sugar().Infof("204 返回异常：%s", resp)
	}

	log.Logger.Info("基本检查完成，无异常")
	exit()
}

func exit() {
	log.Logger.Info("按任意键退出")
	fmt.Scanln()

	os.Exit(0)
}
