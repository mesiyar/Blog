package models

import (
	"time"
	"wechatNotify/pkg/logging"
	"wechatNotify/pkg/setting"
	"wechatNotify/pkg/util"

	"github.com/jinzhu/gorm"
)

const (
	IsStatusEnable  = 1
	IsStatusDisable = 0
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

func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id,password").Where(Auth{Username: username, IsStatus: IsStatusEnable}).First(&auth)
	if auth.ID > 0 {
		passwordEncode := encodePassword(password)
		if passwordEncode == auth.Password {
			return true
		}
	}
	return false
}

func encodePassword(s string) string {
	encoder := util.Encoder{}
	password := s + setting.AuthSalt
	return encoder.Md5(password)
}

// 创建账号
func CreateAuthAccount(username string, password string) bool {
	rs := db.Create(&Auth{
		Username: username,
		Password: encodePassword(password),
		IsStatus: IsStatusEnable,
	})
	if rs.Error != nil {
		logging.Error(rs.Error)
		return false
	}
	return true
}

func DisableAuthAccount(username string) bool {
	data := make(map[string]interface{})
	data["username"] = username
	data["is_status"] = IsStatusDisable
	err := db.Model(&Auth{}).Where("username = ?", username).Update(data)
	if err.Error != nil {
		logging.Error("禁用账号失败")
		logging.Error(err.Error)
		return false
	} else {
		logging.Info("禁用账号成功")
		return true
	}
}

func (auth *Auth) BeforeCreated(scope *gorm.Scope) error {
	t := time.Now().Unix()
	logging.Info(t)
	scope.SetColumn("CreatedOn", t)
	return nil
}

func (auth *Auth) BeforeUpdate(scope *gorm.Scope) error {
	t := time.Now().Unix()
	logging.Info(t)
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
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
