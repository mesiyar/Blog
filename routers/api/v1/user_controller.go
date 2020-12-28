package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wechatNotify/models"
	"wechatNotify/pkg/e"
	"wechatNotify/pkg/setting"
	"wechatNotify/pkg/util"
)

type UserController struct {
}

func (u *UserController)GetUsers(c *gin.Context) {
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

func (u *UserController) GetWeather(c *gin.Context)  {
	ret := util.WeatherGet(util.GetWeatherCityName, "深圳")
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": ret,
	})
}
