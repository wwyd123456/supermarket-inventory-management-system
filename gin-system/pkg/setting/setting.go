package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Port 		string
	Name        string
}

var DatabaseSetting = &Database{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
}

var RedisSetting = &Redis{}

var cfg *ini.File

func Setup()  {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}
	mapTo("database", DatabaseSetting)
	mapTo("server", ServerSetting)
	mapTo("redis", RedisSetting)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}