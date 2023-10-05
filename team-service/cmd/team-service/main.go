package main

import (
	"context"
	"github.com/nats-io/nats.go"
	"net/http"
	"team-service/api"
	"team-service/config"
	"team-service/pkg/koj"
	"team-service/pkg/logger"
	"team-service/pkg/nts"
	"team-service/repository/ent"
	u "team-service/repository/user"
	"team-service/service/log"
	"team-service/service/user"
)

func main() {
	loadConfig(".")
	nts.ConfigureStreaming()

	var (
		l           = logger.Get()
		c           = config.Get()
		kc          = getKeycloakClient(c)
		client      = initClient(l, c)
		userService = user.NewServiceCreate(
			u.NewRepository(*client.User),
			log.NewLogService(l),
		)
		mux = api.NewMux(
			client,
			koj.NewKeycloakOfflineJWT(
				kc,
				c.KeyClockRealm,
				koj.KeyclaokMode(c.AuthMode),
			),
		)
		server = http.Server{
			Addr:    ":" + c.ServerPort,
			Handler: mux,
		}
	)

	nts.Subscribe(nts.GetConnection(), c.UsersSubjectNew, func(m *nats.Msg) error {
		l.Info().Str("Data", string(m.Data)).Msg("message received from user service")
		return userService.Create(context.Background(), string(m.Data))
	})

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
		l := logger.Get()
		l.Error().Err(err).Msg("failed to load environment from file, using os environment")
	}
}
