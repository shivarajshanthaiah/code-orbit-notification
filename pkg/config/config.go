package config

import "github.com/spf13/viper"

type Config struct {
	Email    string `mapstructure:"EMAIL"`
	Password string `mapstructure:"PASSWORD"`
}

func LoadConfig() *Config {
	var config Config
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.Unmarshal(&config)
	return &config
}
