package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

type DBInfo struct {
	Host     string
	Database string
	Username string
	Password string
}

type Context struct {
	Timeout int
}

type Config struct {
	Database DBInfo `mapstructure:"database"`
	Context  Context
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("config read error")
		panic(err)
	}

	config := &Config{}

	err = viper.Unmarshal(&config)

	if err != nil {
		log.Fatal("config unmarshal error")
		panic(err)
	}

	return config
}
