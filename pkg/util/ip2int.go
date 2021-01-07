package util

import (
	"errors"
	"math"
	"net"
	"wechatNotify/pkg/logging"
)

func Ip2Int(ip string) (uint32, error) {
	logging.Info("解析IP", ip)
	ipObj := net.ParseIP(ip)
	if ipObj != nil {
		return uint32(ipObj[12]) << 24 | uint32(ipObj[13])<<16 | uint32(ipObj[14])<<8 | uint32(ipObj[15]), nil
	} else {
		return 0, errors.New("解析失败")
	}
}

// intip 转换成 string
func Long2IPString(i uint32) (string, error) {
	if i > math.MaxUint32 {
		return "", errors.New("beyond the scope of ipv4")
	}

	ip := make(net.IP, net.IPv4len)
	ip[0] = byte(i >> 24)
	ip[1] = byte(i >> 16)
	ip[2] = byte(i >> 8)
	ip[3] = byte(i)

	return ip.String(), nil
}