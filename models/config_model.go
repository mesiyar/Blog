package models

import "wechatNotify/pkg/logging"

type Config struct {
	Model

	ConfigName  string `json:"config_name"`
	ConfigValue string `json:"config_value"`
}

type ConfigModel struct {
}

// 用户列表
func (c *ConfigModel) GetConfigs(page int, pageSize int, data map[string]interface{}) (configs []Config) {

	db.Where(data).Offset(page).Limit(pageSize).Find(&configs)
	return
}



// 获取总数
func (c *ConfigModel) GetTotal(maps interface{}) (count int) {
	db.Model(&Config{}).Where(maps).Count(&count)

	return
}

func (c *ConfigModel) ExistByID(id int) bool {
	var t Config
	db.Select("id").Where("id = ?", id).First(&t)
	if t.ID > 0 {
		return true
	}

	return false
}


// 创建账号
func (c *ConfigModel)CreateConfig(data map[string]string) bool {
	var config Config
	config.ConfigName = data["config_name"]
	config.ConfigValue = data["config_value"]

	rs := db.Create(&config)
	if rs.Error != nil {
		logging.Error(rs.Error)
		return false
	}
	logging.Info("创建配置", data["config_name"], "id ", config.ID)
	return true
}

func (c *ConfigModel) EditConfig(id int, data interface{}) bool {
	db.Model(&Config{}).Where("id = ?", id).Updates(data)

	return true
}

func (c *ConfigModel) DeleteConfig(id int) bool {
	data := make(map[string]interface{})
	data["state"] = IsStatusDisable
	db.Model(&Config{}).Where("id = ?", id).Update(data)

	return true
}

func (c *ConfigModel) GetAll() (configs []Config) {
	db.Find(&configs)
	return
}