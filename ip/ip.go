package ip

import (
	"log"
	"net"
)

//GetLocalIPv4 获取本机IPv4地址
func GetLocalIPv4() (ip string) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		log.Println("net.Interfaces failed, err:", err.Error())
		return ""
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						log.Println("本机IP：" + ipnet.IP.String())
						return ipnet.IP.String()
					}
				}
			}
		}
	}

	return ""
}
