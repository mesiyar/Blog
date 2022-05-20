package jwt

import (
	"net/http"
	"time"
	"wechatNotify/pkg/e"
	"wechatNotify/pkg/util"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.Request.Header.Get("Token") // 拿到token
		if token == "" {
			code = e.InvalidParams
		} else {
			claims, err := util.ParseToken(token)

			if err != nil {
				code = e.ErrorAuthCheckTokenFail
			}
			if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			}

		}

		if code != e.SUCCESS {
			c.JSON(http.StatusOK, gin.H{
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
