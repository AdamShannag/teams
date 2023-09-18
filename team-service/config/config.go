package config

import (
	"github.com/spf13/viper"
	"sync"
)

var once sync.Once

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	Environment string `mapstructure:"ENVIRONMENT"`
	DBDriver    string `mapstructure:"DB_DRIVER"`
	DSN         string `mapstructure:"DSN"`
	ServerPort  string `mapstructure:"SERVER_PORT"`
	LogLevel    int    `mapstructure:"Log_LEVEL"`
}

var config Config

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (err error) {
	once.Do(func() {
		viper.AddConfigPath(path)
		viper.SetConfigName("app")
		viper.SetConfigType("env")

		viper.AutomaticEnv()

		err = viper.ReadInConfig()
		if err != nil {
			return
		}

		err = viper.Unmarshal(&config)
		return
	})
	return
}

func Get() Config {
	return config
}
