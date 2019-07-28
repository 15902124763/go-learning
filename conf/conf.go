package conf

import (
	"errors"
	"github.com/BurntSushi/toml"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"os"
)

var (
	Conf              config // holds the global app config.
	defaultConfigFile = "conf.toml"
)

type config struct {
	ReleaseMode bool   `toml:"release_mode"`
	LogLevel    string `toml:"log_level"`
	CacheStore  string `toml:"cache_store"`
	App         app
	Server      server
	//Redis       redis
	//Kafka       kafka
}

type app struct {
	Name string `toml:"name"`
}

//type kafka struct {
//	Enable     bool     `toml:"kafka_enable"`
//	Topics     []string `toml:"topics"`
//	Servers    []string `toml:"servers"`
//	Ak         string   `toml:"username"`
//	Password   string   `toml:"password"`
//	ConsumerId string   `toml:"consumerGroup"`
//	CertFile   string   `toml:"cert_file"`
//}

type server struct {
	Port string `toml:"port"`
}

//type redis struct {
//	Server         string        `toml:"server"`
//	Pwd            string        `toml:"pwd"`
//	MaxIdle        int           `toml:"maxIdle"`
//	Wait           bool          `toml:"wait"`
//	MaxActive      int           `toml:"maxActive"`
//	IdleTimeout    time.Duration `toml:"idleTimeout"`
//	ReadTimeout    time.Duration `toml:"readTimeout"`
//	WriteTimeout   time.Duration `toml:"writeTimeout"`
//	ConnectTimeout time.Duration `toml:"connectTimeout"`
//	DbIndex        int           `toml:"dbIndex"`
//}

func init() {
	if err := InitConfig(""); err != nil {
		log.Panic(err)
	}
	log.Info("---------配置文件初始化完毕---------")
}

// initConfig initializes the app configuration by first setting defaults,
// then overriding settings from the app config file, then overriding
// It returns an error if any.
func InitConfig(configFile string) error {
	/*默认配置文件在同级的conf下*/
	if configFile == "" {
		configFile = defaultConfigFile
	}

	// Set defaults.
	Conf = config{
		ReleaseMode: false,
		LogLevel:    "DEBUG",
	}

	if _, err := os.Stat(configFile); err != nil {
		return errors.New("config file err:" + err.Error())
	} else {
		log.Infof("load config from file:" + configFile)
		configBytes, err := ioutil.ReadFile(configFile)
		if err != nil {
			return errors.New("config load err:" + err.Error())
		}
		_, err = toml.Decode(string(configBytes), &Conf)
		if err != nil {
			return errors.New("config decode err:" + err.Error())
		}
	}

	// @TODO 配置检查
	log.Infof("config data:%v", Conf)

	return nil
}

func GetLogLvl() log.Lvl {
	//DEBUG INFO WARN ERROR OFF
	switch Conf.LogLevel {
	case "DEBUG":
		return log.DEBUG
	case "INFO":
		return log.INFO
	case "WARN":
		return log.WARN
	case "ERROR":
		return log.ERROR
	case "OF":
		return log.OFF
	}

	return log.DEBUG
}

const (
	// File
	FILE = "FILE"

	// Redis
	REDIS = "REDIS"
)
