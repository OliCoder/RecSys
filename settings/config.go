package settings

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"sync"
)

var configPath string

type Config struct {
	serverHost string
	serverPort string
	
	logDestination string
	
	rpcHost string
	rpcPort string
}

var cfg *Config
var once sync.Once

func init() {
	flag.StringVar(&configPath, "conf", "./conf/default.ini", "config file path")
}

func (config *Config)GetServerHost() string {
	return config.serverHost
}

func (config *Config)GetServerPort() string {
	return config.serverPort
}

func (config *Config)GetRpcHost() string {
	return config.rpcHost
}

func (config *Config)GetRpcPort() string {
	return config.rpcPort
}

func (config *Config)GetLogDestination() string {
	return config.logDestination
}

func GetConfigInstance() *Config {
	once.Do(func() {
		config, err := ini.Load(configPath)
		if err != nil {
			log.Fatalf("Fail to load config file conf/default.ini")
		}
		cfg = &Config{
			serverHost:     config.Section("server").Key("host").MustString("127.0.0.1"),
			serverPort:     config.Section("server").Key("port").MustString("8002"),
			logDestination: config.Section("log").Key("destination").MustString("./conf/default.ini"),
			rpcHost:        config.Section("rpc").Key("host").MustString("127.0.0.1"),
			rpcPort:        config.Section("rpc").Key("port").MustString("8001"),
		}
	})
	return cfg
}