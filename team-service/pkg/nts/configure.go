package nts

import (
	"context"
	"team-service/config"
	"time"
)

func ConfigureStreaming() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	con := GetConnection()
	js := NewJetStream(con)
	config := config.Get()
	_ = NewStream(ctx, js, config.TeamsStream, config.TeamsSubjectNew)
}
