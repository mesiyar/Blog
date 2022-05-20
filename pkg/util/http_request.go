package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"wechatNotify/pkg/logging"
)

// post 请求
func Post(url string, data map[string]interface{}) (string, error) {
	jsonStr, err := JsonEncode(data)
	if err != nil {
		logging.Error(err)
		return "", err
	}
	logging.Info(jsonStr)
	resp, err := http.Post(url, "application/json",
		strings.NewReader(jsonStr))
	if err != nil {
		logging.Error(err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logging.Error(err)
		return "", err
	}
	logging.Info(string(body))
	return string(body), nil
}

// get 请求
func Get(url string) (map[string]interface{}, error) {
	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logging.Warn(err)
	}
	response, err := client.Do(request)
	if err != nil {
		logging.Warn(err)
	}

	//关闭流
	defer response.Body.Close()
	//检出结果集
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logging.Error(fmt.Printf("ioutil.ReadAll failed ,err:%v", err))
	}
	logging.Info(fmt.Printf("返回结果:%s", string(body)))
	return JsonDecode(string(body))

}
