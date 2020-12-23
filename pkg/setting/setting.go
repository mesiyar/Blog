package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File

	RunMode string

	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
	AuthSalt  string

	SignName     string
	TemplateCode string
	AccessKeyId  string
	AccessSecret string

	RedisHost        string
	RedisPassword    string
	RedisMaxIdle     int
	RedisMaxActive   int
	RedisIdleTimeout time.Duration
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini' : %v", err)
	}
	LoadBase()
	LoadServer()
	LoadApp()
	LoadAuth()
	LoadAliyunMsgConfig()
	LoadRedis()
}

func LoadRedis() {
	sec, err := Cfg.GetSection("redis")
	if err != nil {
		log.Fatalf("Fail to get section 'redis' : %v", err)
	}
	RedisHost = sec.Key("HOST").MustString("127.0.0.1:6379")
	RedisPassword = sec.Key("PASSWORD").MustString("")
	RedisMaxIdle = sec.Key("MaxIdle").MustInt(30)
	RedisMaxActive = sec.Key("MaxActive").MustInt(30)
	RedisIdleTimeout = time.Duration(sec.Key("IdleTimeout").MustInt(200))
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app' : %v", err)
	}
	JwtSecret = sec.Key("JWT_SECRET ").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}

func LoadAuth() {
	sec, err := Cfg.GetSection("auth")
	if err != nil {
		log.Fatalf("Fail to get section 'auth' : %v", err)
	}
	AuthSalt = sec.Key("AUTH_SALT ").MustString("!@)*#)!@U#@*!@!)")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server' : %v", err)
	}
	HttpPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadAliyunMsgConfig() {
	sec, err := Cfg.GetSection("aliyun_msg_config")
	if err != nil {
		log.Fatalf("Fail to get section 'aliyun_msg_config' : %v", err)
	}
	AccessKeyId = sec.Key("ACCESS_KEY_ID").MustString("")
	AccessSecret = sec.Key("ACCESS_SECRET").MustString("")
	TemplateCode = sec.Key("TEMPLATE_CODE").MustString("")
	SignName = sec.Key("SIGN_NAME").MustString("")
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}
