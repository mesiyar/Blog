package util

import (
	//"cloud-notes/src/pkg/logging"
	//"cloud-notes/src/pkg/setting"
	//"encoding/json"
	//"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type J struct {
	Code string `json:"code"`
}

//func SendMsg(phone string, verifyCode string) string  {
//	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", setting.AccessKeyId, setting.AccessSecret)
//
//	request := dysmsapi.CreateSendSmsRequest()
//	request.Scheme = "https"
//	code, _ := json.Marshal(J{Code: verifyCode})
//	request.PhoneNumbers = phone
//	request.SignName = setting.SignName
//	request.TemplateCode = setting.TemplateCode
//	request.TemplateParam = code
//
//	response, err := client.SendSms(request)
//	if err != nil {
//		logging.Error(err)
//	}
//	return response.ok
//}
