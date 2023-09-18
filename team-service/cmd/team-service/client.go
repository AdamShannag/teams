package main

import (
	"context"
	"database/sql"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	_ "github.com/sijms/go-ora/v2"
	"log"
	"team-service/config"
	"team-service/repository/ent"
)

func initClient(l zerolog.Logger, config config.Config) *ent.Client {

	client := Open(config)

	//Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		l.Fatal().Msgf("failed creating schema resources: %v", err)
	}
	return client
}

func Open(config config.Config) *ent.Client {
	db, err := sql.Open(config.DBDriver, config.DSN)
	if err != nil {
		log.Fatal(err)
	}

	drv := entsql.OpenDB(config.DBDriver, db)
	return ent.NewClient(ent.Driver(drv))
}
