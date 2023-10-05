package nts

import (
	"context"
	"github.com/nats-io/nats.go"
	"sync"
	"team-service/config"
	"team-service/pkg/logger"

	"github.com/nats-io/nats.go/jetstream"
)

var once sync.Once
var natsConnection *nats.Conn

func GetConnection() *nats.Conn {
	once.Do(func() {
		l := logger.Get()
		url := config.Get().NatsConnectionUrl
		l.Info().Str("url", url).Msg("connecting to nats-server...")
		nc, err := nats.Connect(url)
		if err != nil {
			l.Fatal().Err(err).Msg("error connecting to nats-server")
		}
		natsConnection = nc
	})

	return natsConnection
}

func NewJetStream(nc *nats.Conn) jetstream.JetStream {
	l := logger.Get()

	js, err := jetstream.New(nc)
	if err != nil {
		l.Fatal().Err(err).Msg("an error has occurred while creating a jet stream context")
	}

	return js
}

func NewStream(ctx context.Context, js jetstream.JetStream, stream string, subjects ...string) jetstream.Stream {
	l := logger.Get()

	streamConfig := jetstream.StreamConfig{
		Name:     stream,
		Subjects: subjects,
		MaxAge:   0,
		Storage:  jetstream.FileStorage,
	}

	strm, err := js.CreateStream(ctx, streamConfig)

	if err != nil {
		l.Fatal().Err(err).Msg("an error has occurred while creating the stream")
	}

	return strm
}

func GetStream(ctx context.Context, js jetstream.JetStream, stream string) jetstream.Stream {
	l := logger.Get()

	s, err := js.Stream(ctx, stream)
	if err != nil {
		l.Fatal().Err(err).Msg("an error has occurred while getting stream")
	}
	return s
}

func CreateOrUpdateConsumer(ctx context.Context, stream jetstream.Stream, name string) jetstream.Consumer {
	l := logger.Get()
	consumer, err := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Durable:       name,
		AckPolicy:     jetstream.AckExplicitPolicy,
		DeliverPolicy: jetstream.DeliverAllPolicy,
	})
	if err != nil {
		l.Fatal().Err(err).Msg("an error has occurred while creating/updating consumer")
	}

	return consumer
}

func GetConsumer(ctx context.Context, js jetstream.JetStream, stream string, name string) jetstream.Consumer {
	l := logger.Get()

	c, err := js.Consumer(ctx, stream, name)
	if err != nil {
		l.Fatal().Err(err).Msg("an error has occurred while getting consumer")
	}
	return c
}

func Publish(ctx context.Context, js jetstream.JetStream, msg *nats.Msg) (*jetstream.PubAck, error) {
	return js.PublishMsg(ctx, msg)
}

func ListenOnConsumer(ctx context.Context, consumer jetstream.Consumer, cb func(*jetstream.Msg) error) {
	l := logger.Get()

	go func() {
		iter, _ := consumer.Messages()
		defer iter.Stop()
		for {
			msg, err := iter.Next()
			if err != nil {
				l.Error().Err(err).Msg("an error occurred")
				continue
			}
			err = cb(&msg)
			if err != nil {
				l.Error().Err(err).Msg("an error occurred")
			}
			msg.Ack()
		}
	}()
}

func Subscribe(nc *nats.Conn, subject string, cb func(*nats.Msg) error) *nats.Subscription {
	l := logger.Get()
	s, err := nc.Subscribe(subject, func(msg *nats.Msg) {
		err := cb(msg)
		if err != nil {
			l.Error().Err(err).Msg("an error occurred")
		}
	})

	if err != nil {
		l.Fatal().Err(err).Msg("an error has occurred while subscribing to subject: " + subject)
	}

	return s
}
