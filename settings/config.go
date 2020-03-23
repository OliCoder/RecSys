package settings

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"os"
	"sync"
	"time"
)

var configPath string

type Config struct {
	ServerHost         string
	ServerPort         string
	ServerWriteTimeout time.Duration
	ServerReadTimeout  time.Duration
	IdentityKey        string

	LogDestination string
	logFile        *os.File

	RpcHost string
	RpcPort string

	DBType      string
	User        string
	Password    string
	TablePrefix string
	DBHost      string
	DBPort      string
	DBName      string
}

var cfg *Config
var once sync.Once

func init() {
	flag.StringVar(&configPath, "conf", "./conf/default.ini", "config file path")
}

func GetConfigInstance() *Config {
	once.Do(func() {
		config, err := ini.Load(configPath)
		if err != nil {
			log.Fatalf("Fail to load config file conf/default.ini")
		}
		cfg = &Config{
			ServerHost:         config.Section("server").Key("host").MustString("127.0.0.1"),
			ServerPort:         config.Section("server").Key("port").MustString("8002"),
			ServerWriteTimeout: time.Duration(config.Section("server").Key("write_timeout").MustInt(100)) * time.Second,
			ServerReadTimeout:  time.Duration(config.Section("server").Key("read_timeout").MustInt(100)) * time.Second,
			IdentityKey:        config.Section("server").Key("identity_key").MustString("idname"),
			LogDestination:     config.Section("log").Key("destination").MustString("./conf/default.ini"),
			RpcHost:            config.Section("rpc").Key("host").MustString("127.0.0.1"),
			RpcPort:            config.Section("rpc").Key("port").MustString("8001"),
			DBType:             config.Section("db").Key("dbtype").MustString("mysql"),
			User:               config.Section("db").Key("user").MustString("root"),
			Password:           config.Section("db").Key("password").MustString("123456"),
			TablePrefix:        config.Section("db").Key("table_prefix").MustString("recsys_"),
			DBHost:             config.Section("db").Key("server").MustString("127.0.0.1"),
			DBPort:             config.Section("db").Key("port").MustString("3306"),
		}

		// log config
		cfg.logFile, err = os.OpenFile(cfg.LogDestination, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Error("Fail to log to file")
		} else {
			log.SetOutput(cfg.logFile)
		}
	})
	return cfg
}
