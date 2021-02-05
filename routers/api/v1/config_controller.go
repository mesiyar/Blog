package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
	"wechatNotify/models"
	"wechatNotify/pkg/e"
	"wechatNotify/pkg/setting"
	"wechatNotify/pkg/util"
)

type ConfigController struct {
}

var configModel = new(models.ConfigModel)

//获取多个文章
func (a ConfigController) GetConfigs(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}
	code := e.InvalidParams
	if ! valid.HasErrors() {
		code = e.SUCCESS
		data["lists"] = configModel.GetConfigs(util.GetPage(c), setting.PageSize, maps)
		data["total"] = configModel.GetTotal(maps)

	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//新增文章
func (a ConfigController) AddConfig(c *gin.Context) {
	configName := c.PostForm("config_name")
	configValue := c.PostForm("config_value")

	valid := validation.Validation{}
	valid.Required(configName, "config_name").Message("name不能为空")
	valid.Required(configValue, "config_value").Message("value不能为空")

	code := e.InvalidParams
	if ! valid.HasErrors() {
		data := make(map[string]string)
		data["config_name"] = configName
		data["config_value"] = configValue

		configModel.CreateConfig(data)
		code = e.SUCCESS
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

//修改文章
func (a ConfigController) EditConfig(c *gin.Context) {
	valid := validation.Validation{}

	id := com.StrTo(c.PostForm("id")).MustInt()
	configName := c.PostForm("config_name")
	configValue := c.PostForm("config_value")

	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.Required(configName, "config_name").Message("name不能为空")
	valid.Required(configValue, "config_value").Message("value不能为空")
	modifiedBy := util.UserInfo.Username

	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")

	code := e.InvalidParams
	if ! valid.HasErrors() {
		if configModel.ExistByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			data["config_name"] = configName
			data["config_value"] = configValue
			configModel.EditConfig(id, data)
			code = e.SUCCESS
		} else {
			code = e.ErrorArticleNotExist
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

//删除文章
func (a ConfigController) DeleteConfig(c *gin.Context) {
	id := com.StrTo(c.Query("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.InvalidParams
	if ! valid.HasErrors() {
		if configModel.ExistByID(id) {
			configModel.DeleteConfig(id)
			code = e.SUCCESS
		} else {
			code = e.ErrorArticleNotExist
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func (a ConfigController) SiteConfig(c *gin.Context) {
	configs := configModel.GetAll()
	data := make(map[string]string)
	for _,v := range configs {
		data[v.ConfigName] = v.ConfigValue
	}
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}


