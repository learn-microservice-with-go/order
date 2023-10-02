package config

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
)

type Config struct {
	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     string `mapstructure:"REDIS_PORT"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	MySQLHost     string `mapstructure:"MYSQL_HOST"`
	MySQLPort     string `mapstructure:"MYSQL_PORT"`
	MySQLUser     string `mapstructure:"MYSQL_USER"`
	MySQLPassword string `mapstructure:"MYSQL_PASSWORD"`
	MySQLDbName   string `mapstructure:"MYSQL_DB_NAME"`
}

var Provider = wire.NewSet(NewConfig)

func NewConfig() (*Config, error) {
	viper.SetConfigFile("../../.env")
	viper.AutomaticEnv()
	var config Config
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
