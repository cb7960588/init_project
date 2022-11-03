package utils

import (
	"fmt"
	"github.com/koding/multiconfig"
)

var Config *AdConfig

type AdConfig struct {
	HttpConfig  *HttpConfig
	RedisConfig *RedisConfig
	LogConfig   *LogConfig
	MysqlConfig *DbConfig
}
type HttpConfig struct {
	Port int
}
type RedisConfig struct {
	Addr           string
	ConnectTimeout int
	ReadTimeout    int
	WriteTimeout   int
	PoolSize       int
	Password       string
	DB             int
}
type LogConfig struct {
	RuntimeConf string `required`
	WorkConf     string `required`
}
type DbConfig struct {
	Driver   string
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
	Status   bool
}

func LoadConfigFile(confPath string) error {
	m := &multiconfig.TOMLLoader{Path: confPath}
	if err := m.Load(Config); err != nil {
		return err
	}
	return nil
}

func init() {
	Config = new(AdConfig)
}
func (adConfig *AdConfig) GetMysqlConfig() string {
	config := adConfig.MysqlConfig
	name := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Dbname)
	return name
}
