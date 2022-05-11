package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	RunMode string

	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PidFile      string

	PageSize      int
	JwtSecret     string
	JwtExpireTime time.Duration // 过期时间
	AuthSalt      string

	SignName     string
	TemplateCode string
	AccessKeyId  string
	AccessSecret string

	RedisHost        string
	RedisPassword    string
	RedisMaxIdle     int
	RedisMaxActive   int
	RedisIdleTimeout time.Duration
	// 天气预报配置
	WeatherUrl       string
	WeatherAppID     string
	WeatherAppSecret string
	ImagePrefixUrl   string
	ImageSavePath    string
	RuntimeRootPath  string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini' : %v", err)
	}
	ImagePrefixUrl = "http://127.0.0.1:8000"
	ImageSavePath = "upload/images/"
	RuntimeRootPath = "runtime/"
	LoadBase()
	LoadServer()
	LoadApp()
	LoadAuth()
	LoadAliyunMsgConfig()
	LoadRedis()
	LoadWeather()
}

// 天气预报配置
func LoadWeather() {
	sec, err := Cfg.GetSection("weather")
	if err != nil {
		log.Fatalf("Fail to get section 'weather' : %v", err)
	}
	WeatherUrl = sec.Key("API_URL").MustString("https://v0.yiketianqi.com/api")
	WeatherAppID = sec.Key("API_ID").MustString("")
	WeatherAppSecret = sec.Key("API_SECRET").MustString("")
}

// 加载redis配置
func LoadRedis() {
	sec, err := Cfg.GetSection("redis")
	if err != nil {
		log.Fatalf("Fail to get section 'redis' : %v", err)
	}
	RedisHost = sec.Key("Host").MustString("127.0.0.1:6379")
	RedisPassword = sec.Key("Password").MustString("")
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
	JwtExpireTime = time.Duration(sec.Key("JWT_EXPIRE_TIME").MustInt(1800))
	PidFile = sec.Key("PID_FILE").MustString("")
}

func LoadAuth() {
	sec, err := Cfg.GetSection("auth")
	if err != nil {
		log.Fatalf("Fail to get section 'auth' : %v", err)
	}
	AuthSalt = sec.Key("AUTH_SALT ").MustString("!@)*#)!@U#@*!@!)")
}

// 加载服务配置
func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server' : %v", err)
	}
	HttpPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

// 加载短信配置
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

// 加载运行模式
func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}
