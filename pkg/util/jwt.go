package util

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
	"wechatNotify/pkg/logging"
	"wechatNotify/pkg/setting"
)

var jwtSecret = []byte(setting.JwtSecret)

// 存储接口信息
var UserInfo *Claims

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ID       int    `json:"id"`
	jwt.StandardClaims
}

// @Summary 生成token
func GenerateToken(username, password string, id int) (string, int64, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(setting.JwtExpireTime * time.Second)
	claims := Claims{
		username,
		password,
		id,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "weChat_notify",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, expireTime.Unix(), err
}

// @Summary 解析token
func ParseToken(token string) (*Claims, error) {
	logging.Info("解析token", token)
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			logging.Info("解析成功")
			UserInfo = claims
			return claims, nil
		}
	}
	logging.Info("解析失败", err.Error())
	return nil, err
}
