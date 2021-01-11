package models

import (
	"time"
	"wechatNotify/pkg/logging"
	"wechatNotify/pkg/redisKey"
	"wechatNotify/pkg/setting"
	"wechatNotify/pkg/util"
)

type Auth struct {
	Model

	Username string `json:"username"`
	Password string `json:"password"`
	IsStatus int    `json:"is_status"`
}

type AuthAccount struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password, ip string) (b bool, id int) {
	var auth Auth
	db.Select("id,password").Where(Auth{Username: username, IsStatus: IsStatusEnable}).First(&auth)
	if auth.ID > 0 {
		passwordEncode := encodePassword(password)
		if passwordEncode == auth.Password {
			logging.Info("验证成功")
			// 更新最后登录时间 ip
			data := make(map[string]interface{})
			data["last_login_time"] = time.Now().Unix()
			sIp, err := util.Ip2Int(ip)
			if err != nil {
				logging.Error("解析ip失败")
			}
			data["last_login_ip"] = sIp
			db.Table("t_auth").Where("id = ?", auth.ID).Update(data)
			id = auth.ID
			b = true
			return
		}
	}
	id = 0
	b = false
	return
}

func encodePassword(s string) string {
	password := s + setting.AuthSalt
	return util.Md5(password)
}

// 创建账号
func CreateAuthAccount(username string, password string) bool {
	var auth Auth
	auth.Username = username
	auth.Password = encodePassword(password)
	auth.IsStatus = IsStatusEnable

	rs := db.Create(&auth)
	if rs.Error != nil {
		logging.Error(rs.Error)
		return false
	}
	logging.Info("创建新账号", username, "id ", auth.ID)
	return true
}

// 禁用账号
func DisableAuthAccount(username string) bool {
	data := make(map[string]interface{})
	data["username"] = username
	data["is_status"] = IsStatusDisable
	err := db.Model(&Auth{}).Where("username = ?", username).Update(data)
	redis := util.Redis{}
	redis.HDel(redisKey.KeyAccountInfo, username)
	if err.Error != nil {
		logging.Error("禁用账号失败")
		logging.Error(err.Error)
		return false
	} else {
		logging.Info("禁用账号成功")
		return true
	}
}

func CheckAuthName(username string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, IsStatus: IsStatusEnable}).First(&auth)
	if auth.ID > 0 {
		logging.Info("验证成功")
		return false
	}

	return true
}
