package main

import (
	"net/http"
	"project-service/api"
	"project-service/config"
	"project-service/pkg/logger"
	"project-service/repository/ent"
)

const defaultPort = "80"

func main() {
	loadConfig(".")

	var (
		l      = logger.Get()
		c      = config.Get()
		client = initClient(l, c)
		mux    = api.NewMux(client)
		port   = getPort(c)
		server = http.Server{
			Addr:    ":" + port,
			Handler: mux,
		}
	)

	l.Info().
		Str("port", port).
		Msg("starting project-service")

	l.Fatal().
		Err(server.ListenAndServe()).
		Msg("project-service Server Closed")

	defer func(client *ent.Client) {
		_ = client.Close()
	}(client)
}

func getPort(c config.Config) string {
	if c.ServerPort == "" {
		return defaultPort
	}
	return c.ServerPort
}

func loadConfig(path string) {
	err := config.LoadConfig(path)
	if err != nil {
		logger.DefaultInit()
		l := logger.Get()
		l.Fatal().Err(err).Msg("cannot load config")
	} else {
		c := config.Get()
		logger.Init(c.LogLevel, c.Environment)
	}
}
