package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"wechatNotify/pkg/logging"
	"wechatNotify/pkg/setting"
)

const (
	GetWeatherCityId   = 1 // 通过city id   获取天气
	GetWeatherCityName = 2 // 通过city name 获取天气
)

// 获取天气
func WeatherGet(t int, s string) (w map[string]interface{}) {
	client := &http.Client{}

	url := fmt.Sprintf("%s?version=v61&appid=%s&appsecret=%s", setting.WeatherUrl, setting.WeatherAppID, setting.WeatherAppSecret)
	switch t {
	case GetWeatherCityId:
		url = fmt.Sprintf("%s&cityid=%s", url, s)
		break
	case GetWeatherCityName:
		url = fmt.Sprintf("%s&city=%s", url, s)
		break
	}
	logging.Info("获取天气信息地址", url)
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
	w, _ = JsonDecode(string(body))
	logging.Info(w)
	return
}
