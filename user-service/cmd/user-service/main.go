package main

import (
	"net/http"
	"user-service/api"
	"user-service/cmd/user-service/config"
	"user-service/pkg/koj"
	"user-service/pkg/logger"
	"user-service/pkg/nts"

	"github.com/nats-io/nats.go"
)

func main() {
	config.LoadEnv()

	nts.ConfigureStreaming()

	var (
		l   = logger.Get()
		kc  = getKeycloakClient()
		mux = api.NewMux(
			kc,
			koj.NewKeycloakOfflineJWT(
				kc,
				config.KEYCLOAK_REALM,
				koj.KeyclaokMode(config.AUTH_MODE),
			),
		)
		server = http.Server{
			Addr:    ":" + config.WEB_PORT,
			Handler: mux,
		}
	)

	nts.Subscribe(nts.GetConnection(), config.USERS_SUBJECT_NEW, func(m *nats.Msg) error {
		l.Info().Str("Data", string(m.Data)).Msg("message received")
		return nil
	})

	l.Info().
		Str("port", config.WEB_PORT).
		Msg("starting user-service")

	l.Fatal().
		Err(server.ListenAndServe()).
		Msg("user-service Server Closed")
}
