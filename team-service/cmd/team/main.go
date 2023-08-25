package main

import (
	"net/http"
	"team-service/api"
	"team-service/pkg/logger"
)

const webPort = "8080"

func main() {
	var (
		l      = logger.Get()
		mux    = api.NewMux()
		server = http.Server{
			Addr:    ":" + webPort,
			Handler: mux,
		}
	)

	l.Info().
		Str("port", webPort).
		Msg("starting teams-service")

	l.Fatal().
		Err(server.ListenAndServe()).
		Msg("teams-service Server Closed")
}
