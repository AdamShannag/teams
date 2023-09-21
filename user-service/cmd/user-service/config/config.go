package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	WEB_PORT string = os.Getenv("WEB_PORT")

	LOG_LEVEL string = os.Getenv("LOG_LEVEL")
	APP_ENV   string = os.Getenv("APP_ENV")

	KEYCLOAK_BASE_URL      string = os.Getenv("KEYCLOAK_BASE_URL")
	KEYCLOAK_REALM         string = os.Getenv("KEYCLOAK_REALM")
	KEYCLOAK_CLINET_ID     string = os.Getenv("KEYCLOAK_CLINET_ID")
	KEYCLOAK_CLIENT_SECRET string = os.Getenv("KEYCLOAK_CLIENT_SECRET")

	AUTH_MODE string = os.Getenv("AUTH_MODE")

	NATS_CONNECTION_URL string = os.Getenv("NATS_CONNECTION_URL")
	USERS_STREAM        string = os.Getenv("USERS_STREAM")
	USERS_SUBJECT_NEW   string = os.Getenv("USERS_SUBJECT_NEW")
)

func LoadEnv() {
	err := godotenv.Load("user-service.env")
	if err != nil {
		log.Println("failed to load environment from file, using os environment")
	}

	initializeEnvs()
}

func initializeEnvs() {
	if WEB_PORT = os.Getenv("WEB_PORT"); WEB_PORT == "" {
		WEB_PORT = "80"
	}

	LOG_LEVEL = os.Getenv("LOG_LEVEL")
	APP_ENV = os.Getenv("APP_ENV")

	KEYCLOAK_BASE_URL = os.Getenv("KEYCLOAK_BASE_URL")
	KEYCLOAK_REALM = os.Getenv("KEYCLOAK_REALM")
	KEYCLOAK_CLINET_ID = os.Getenv("KEYCLOAK_CLINET_ID")
	KEYCLOAK_CLIENT_SECRET = os.Getenv("KEYCLOAK_CLIENT_SECRET")

	AUTH_MODE = os.Getenv("AUTH_MODE")

	NATS_CONNECTION_URL = os.Getenv("NATS_CONNECTION_URL")
	USERS_STREAM = os.Getenv("USERS_STREAM")
	USERS_SUBJECT_NEW = os.Getenv("USERS_SUBJECT_NEW")
}
