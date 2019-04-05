package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/ksuid"

	"github.com/nassor/kafka-compact-pipeline/services/event-receiver/internal/device"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	// Closing signal
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh,
		syscall.SIGKILL,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	// Kafka Consumer
	kc, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":    "127.0.0.1:9092",
		"session.timeout.ms":   10000,
		"enable.auto.commit":   false,
		"enable.partition.eof": true,
		"group.id":             ksuid.New().String(),
		"default.topic.config": kafka.ConfigMap{"auto.offset.reset": "earliest"},
	})
	if err != nil {
		log.Panic().AnErr("connecting to kafka", err)
	}
	sub, err := device.NewKafkaSubscriber(kc, "example.device.v1.pb")
	if err != nil {
		log.Panic().AnErr("subscribing to the topic", err)
	}
	log.Info().Msg("kafka subscriber initialized")
	defer kc.Close()

	// In-Memory store
	st := device.NewInMemoryStore()

	// Device Service
	s := device.NewService(sub, st)
	go s.ReceiveData()

	// starting receive data
	now := time.Now()
	log.Info().Msgf("starting to load evens")
	go sub.Listen()
	timeout := time.NewTimer(5 * time.Minute)
	select {
	case <-timeout.C:
		log.Error().Msgf("load timed out after 5 minutes")
	case total := <-sub.IsLoaded():
		log.Info().Msgf("Finished loading %d events in %s", total, time.Since(now))
	}

	// stopping the service
	sig := <-stopCh
	defer close(stopCh)
	log.Info().Msgf("stopping the consumer service. Signal: %s", sig)
}
