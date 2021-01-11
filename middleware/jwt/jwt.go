package jwt

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"wechatNotify/pkg/e"
	"wechatNotify/pkg/redisKey"
	"wechatNotify/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.Request.Header.Get("token") // 拿到token
		if token == "" {
			code = e.InvalidParams
		} else {
			claims, err := util.ParseToken(token)

			if err != nil {
				code = e.ErrorAuthCheckTokenFail
			} else {
				redis := util.Redis{}
				t, _ := redis.HGet(redisKey.KeyAccountInfo, claims.Username)
				t1 := make(map[string]interface{})
				json.Unmarshal(t, &t1)

				if t1["token"] != token {
					code = e.ErrorAuthCheckTokenFail
				} else {
					if time.Now().Unix() > claims.ExpiresAt {
						code = e.ErrorAuthCheckTokenTimeout
					}
				}
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
