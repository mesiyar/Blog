package api

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"wechatNotify/models"
	"wechatNotify/pkg/e"
	"wechatNotify/pkg/logging"
	"wechatNotify/pkg/redisKey"
	"wechatNotify/pkg/setting"
	"wechatNotify/pkg/util"
)

type auth struct {
	Username string `valid:"Required; MinSize(6);MaxSize(20)"`
	Password string `valid:"Required; MinSize(6);MaxSize(32)"`
}

// @Summary 获取认证
// @Accept application/x-www-form-urlencoded
// @Produce  json
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {string} json "{"code":200,"data":{"token":"", "user":"","expire":""},"msg":"ok"}"
// @Router /auth [post]
func GetAuth(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	logging.Info(username, password)
	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)
	data := make(map[string]interface{})
	code := e.InvalidParams
	errMsg := ""
	if ok {
		isExist, id := models.CheckAuth(username, password, c.ClientIP())
		if isExist {
			redis := util.Redis{}
			token, err := redis.HGet(redisKey.KeyAccountInfo, username)
			if err == nil {
				var dat map[string]interface{}
				json.Unmarshal(token, &dat)
				data= dat
				code = e.SUCCESS
			} else {
				token, expire, err := util.GenerateToken(username, password, id)
				if err != nil {
					code = e.ErrorAuthToken
				} else {
					logging.Info(username, "获取token token \n", token)
					data["token"] = token
					data["expire"] = expire
					data["user"] = username

					code = e.SUCCESS
				}
				err2 := redis.Set(redisKey.KeyAccountInfo, data, int(setting.JwtExpireTime))
				if err2 != nil {
					logging.Error(err2)
				}
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

// @Summary 创建账号
// @Produce  json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Param token header string true "token"
// @Success 200 {string} json "{"code":200,"data":"ok","msg":"ok"}"
// @Router /api/v1/create_account [post]
func CreateAuth(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
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
	if _,err := r.Delete(redisKey.KeyAccountInfo); err != nil {
		logging.Info("清理redis缓存失败, ", err)
	}
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": "",
	})
}
