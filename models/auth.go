package models

import (
	"time"
	"wechatNotify/pkg/logging"
	"wechatNotify/pkg/setting"
	"wechatNotify/pkg/util"

	"github.com/jinzhu/gorm"
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

func CheckAuth(username, password string) (b bool, id int) {
	var auth Auth
	db.Select("id,password").Where(Auth{Username: username, IsStatus: IsStatusEnable}).First(&auth)
	if auth.ID > 0 {
		passwordEncode := encodePassword(password)
		if passwordEncode == auth.Password {
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
	logging.Info(auth.ID)
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

func (auth *Auth) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (auth *Auth) BeforeUpdate(scope *gorm.Scope) error {
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
