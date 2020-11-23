package config

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type DBInfo struct {
	Host     string
	Database string
	Username string
	Password string
}

type LogInfo struct {
	Level     string
	SentryDSN string
}
type PhaseInfo struct {
	Level string
}

type Context struct {
	Timeout int
}

type Configuration struct {
	Database DBInfo    `mapstructure:"database"`
	Log      LogInfo   `mapstructure:"logging"`
	Phase    PhaseInfo `mapstructure:"phase"`
	Context  Context
}

func LoadConfig() *Configuration {
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

	config := &Configuration{}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("config unmarshal error")
		panic(err)
	}

	return config
}
