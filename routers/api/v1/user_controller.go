package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"wechatNotify/models"
	"wechatNotify/pkg/e"
	"wechatNotify/pkg/setting"
	"wechatNotify/pkg/util"
)

type UserController struct {
}

func (u *UserController) GetUsers(c *gin.Context) {
	name := c.Query("username")
	data := make(map[string]interface{})
	if name != "" {
		data["username"] = name
	}

	ret := make(map[string]interface{})
	usModel := models.UserModel{}
	ret["list"] = usModel.GetUsers(util.GetPage(c), setting.PageSize, data)
	ret["total"] = usModel.GetTotal(data)
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": ret,
	})
}

func (u *UserController) GetWeather(c *gin.Context) {
	ret := util.WeatherGet(util.GetWeatherCityName, "深圳")
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": ret,
	})
}

func (u *UserController) GetUserInfo(c *gin.Context) {
	id := util.UserInfo.ID
	usModel := models.UserModel{}
	userInfo := usModel.GetUser(id)
	code := e.SUCCESS
	data := make(map[string]interface{})
	if userInfo.ID != id {
		code = e.ERROR
	} else {
		data["username"] = userInfo.Username
		ip ,_ := util.Long2IPString(uint32(userInfo.LastLoginIp))
		data["last_login_ip"] = ip
		data["last_login_time"] = time.Unix(int64(userInfo.LastLoginTime), 0).Format("2006-01-02 15:04:05")
	}


	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
