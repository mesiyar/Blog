package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"net"
	"net/http"
	"wechatNotify/pkg/logging"
)

func Ip2Int(ip string) (uint32, error) {
	logging.Info("解析IP", ip)
	ipObj := net.ParseIP(ip)
	if ipObj != nil {
		return uint32(ipObj[12])<<24 | uint32(ipObj[13])<<16 | uint32(ipObj[14])<<8 | uint32(ipObj[15]), nil
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

type IpAddress struct {
	Ret  string   `json:"ret"`
	Ip   string   `json:"Ip"`
	Data []string `json:"data"`
}

// 获取IP地址
func GetIpAddress(ip string) (IpAddress, error) {
	client := &http.Client{}

	url := fmt.Sprintf("https://api.ip138.com/ip/?ip=%s&datatype=jsonp&token=%s", ip, "09f5f36bfddefd653ef9de6b981a6b29")
	logging.Info("获取ip地址", url)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logging.Warn(err)
	}
	response, _ := client.Do(request)
	//关闭流
	defer response.Body.Close()
	//检出结果集
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logging.Error(fmt.Printf("ioutil.ReadAll failed ,err:%v", err))
	}
	logging.Info("返回结果")
	logging.Info(string(body))
	var ipAddress IpAddress
	err = json.Unmarshal(body, &ipAddress)
	return ipAddress, err
}
