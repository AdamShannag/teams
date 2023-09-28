package main

import (
	"net/http"
	"team-service/api"
	"team-service/config"
	"team-service/pkg/logger"
	"team-service/repository/ent"
)

func main() {
	loadConfig(".")
	var (
		l      = logger.Get()
		c      = config.Get()
		client = initClient(l, c)
		mux    = api.NewMux(client)
		server = http.Server{
			Addr:    ":" + c.ServerPort,
			Handler: mux,
		}
	)

	l.Info().
		Str("port", c.ServerPort).
		Msg("starting teams-service")

	l.Fatal().
		Err(server.ListenAndServe()).
		Msg("teams-service Server Closed")

	defer func(client *ent.Client) {
		_ = client.Close()
	}(client)
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
