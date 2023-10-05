package config

import (
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"log"
	"sync"
	"team-service/constant"
)

var once sync.Once

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	Environment       string `mapstructure:"APP_ENV"`
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DSN               string `mapstructure:"DSN"`
	ServerPort        string `mapstructure:"SERVER_PORT"`
	LogLevel          int    `mapstructure:"Log_LEVEL"`
	NatsConnectionUrl string `mapstructure:"NATS_CONNECTION_URL"`
	TeamsStream       string `mapstructure:"TEAMS_STREAM"`
	TeamsSubjectNew   string `mapstructure:"TEAMS_SUBJECT_NEW"`
	UsersSubjectNew   string `mapstructure:"USERS_SUBJECT_NEW"`
	KeyClockBaseUrl   string `mapstructure:"KEYCLOAK_BASE_URL"`
	KeyClockRealm     string `mapstructure:"KEYCLOAK_REALM"`
	AuthMode          string `mapstructure:"AUTH_MODE"`
}

var config Config

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (err error) {
	once.Do(func() {
		viper.AddConfigPath(path)
		viper.SetConfigName("app")
		viper.SetConfigType("env")
		viper.AutomaticEnv()

		viper.SetDefault("LOG_LEVEL", zerolog.InfoLevel)
		viper.SetDefault("SERVER_PORT", constant.DEFAULT_PORT)

		viper.BindEnv("SERVER_PORT")
		viper.BindEnv("APP_ENV")
		viper.BindEnv("LOG_LEVEL")
		viper.BindEnv("DSN")
		viper.BindEnv("DB_DRIVER")
		viper.BindEnv("NATS_CONNECTION_URL")
		viper.BindEnv("TEAMS_STREAM")
		viper.BindEnv("TEAMS_SUBJECT_NEW")
		viper.BindEnv("USERS_SUBJECT_NEW")
		viper.BindEnv("KEYCLOAK_BASE_URL")
		viper.BindEnv("KEYCLOAK_REALM")
		viper.BindEnv("AUTH_MODE")

		err = viper.ReadInConfig()
		errUnmarshal := viper.Unmarshal(&config)
		if err == nil {
			err = errUnmarshal
		}
		return
	})
	return
}

func Get() Config {
	err := LoadConfig(".")
	if err != nil {
		log.Println("failed to load environment from file, using os environment")
	}
	return config
}
