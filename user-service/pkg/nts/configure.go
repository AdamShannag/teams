package nts

import (
	"context"
	"time"
	"user-service/cmd/user-service/config"
)

func ConfigureStreaming() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	con := GetConnection()
	js := NewJetStream(con)
	_ = NewStream(ctx, js, config.USERS_STREAM, config.USERS_SUBJECT_NEW)
}
