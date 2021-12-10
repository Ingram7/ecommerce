package app

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"path/filepath"
)

var loader = viper.New()

var (
	AddConfigPath = loader.AddConfigPath
	SetConfigName = loader.SetConfigName
	SetConfigType = loader.SetConfigType
	ReadInConfig  = loader.ReadInConfig
	Unmarshal     = loader.Unmarshal
)

var Path = ""

type configOption struct {
	configName string
	configType string
}

type configOptFunc func(opt *configOption)

func withConfigName(configName string) configOptFunc {
	return func(opt *configOption) {
		opt.configName = configName
	}
}

func withConfigType(configType string) configOptFunc {
	return func(opt *configOption) {
		opt.configType = configType
	}
}

type serverConfig struct {
	Host string
	Port string
}

func (c *serverConfig) String() string {
	return fmt.Sprintf("{Host:%s Port:%s}", c.Host, c.Port)
}

type config struct {
	Mode string
	//Log  *LogConfig
	//DB   *db.Conn

	HashIdSalt string
	TokenSalt  string
	Server     *serverConfig

	//Cert *Cert
}

func loadConfig() *config {
	conf := new(config)
	conf.load()
	return conf
}

func (conf *config) load(options ...configOptFunc) {

	opt := configOption{
		configName: "config",
		configType: "yaml",
	}

	for _, f := range options {
		f(&opt)
	}

	loader.AddConfigPath(Path)
	loader.SetConfigName(opt.configName)
	loader.SetConfigType(opt.configType)
	if err := loader.ReadInConfig(); err != nil {
		log.Println(filepath.Abs(Path))
		log.Fatalf("load config file error: %s, path: %s", err.Error(), Path)
	}

	if err := loader.Unmarshal(conf); err != nil {
		log.Fatalf("config unmarshal error: %s", err.Error())
	}
}

func (conf *config) print() {
	fmt.Printf("config=%+v\n", conf)
}
