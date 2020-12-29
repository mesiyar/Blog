package api

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"wechatNotify/models"
	"wechatNotify/pkg/e"
	"wechatNotify/pkg/logging"
	"wechatNotify/pkg/redisKey"
	"wechatNotify/pkg/util"
)

type auth struct {
	Username string `valid:"Required; MinSize(6);MaxSize(20)"`
	Password string `valid:"Required; MinSize(6);MaxSize(32)"`
}

// @Summary 获取认证
// @Produce  json
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Success 200 {string} json "{"code":200,"data":{"token":""},"msg":"ok"}"
// @Router /auth [get]
func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)
	data := make(map[string]interface{})
	code := e.InvalidParams
	errMsg := ""
	if ok {
		isExist, id := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password, id)
			if err != nil {
				code = e.ErrorAuthToken
			} else {
				logging.Info(username, "获取token token \n", token)
				data["token"] = token

				code = e.SUCCESS
			}
			redis := util.Redis{}
			err2 := redis.HSet(redisKey.KeyAccountInfo, username, token)
			if err2 != nil {
				logging.Error(err2)
			}
		} else {
			code = e.ErrorAuth
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
			errMsg = errMsg + err.Key + err.Message
		}
	}
	if errMsg == "" {
		errMsg = e.GetMsg(code)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errMsg,
		"data": data,
	})
}

func CreateAuth(c *gin.Context) {

	username := c.Query("username")
	password := c.Query("password")
	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(a)
	data := "ok"
	code := e.InvalidParams
	errMsg := ""
	if ok {
		isExist := models.CheckAuthName(username)
		if isExist {
			createAuth := models.CreateAuthAccount(username, password)
			if createAuth {
				code = e.SUCCESS
			} else {
				data = "创建失败"
				code = e.ERROR
			}
		} else {
			code = e.ErrorAuthnameExist
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
			errMsg += err.Key + " " + err.Message
		}
	}

	if errMsg == "" {
		errMsg = e.GetMsg(code)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errMsg,
		"data": data,
	})
}

func DisableAuth(c *gin.Context) {
	username := c.Query("username")

	data := "ok"
	code := e.InvalidParams

	isExist := models.CheckAuthName(username)
	if !isExist {
		disableAuth := models.DisableAuthAccount(username)
		if disableAuth {
			code = e.SUCCESS
		} else {
			data = "禁用失败"
			code = e.ERROR
		}
	} else {
		code = e.ErrorAccountNotExist
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func Logout(c *gin.Context) {
	logging.Info(fmt.Sprintf("%s 退出登录", util.UserInfo.Username))
	r := util.Redis{}
	r.HDel(redisKey.KeyAccountInfo, util.UserInfo.Username)
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": "",
	})
}
