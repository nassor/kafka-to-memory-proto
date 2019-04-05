package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/golang/protobuf/ptypes"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/ksuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/nassor/kafka-to-memory-proto/api"
	"github.com/nassor/kafka-to-memory-proto/services/device-manager/internal/device"
)

func main() {
	ctx := context.Background()

	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	// Closing signal
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh,
		syscall.SIGKILL,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	// MongoDB
	mc, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Panic().AnErr("creating mongodb client", err)
	}
	if err = mc.Connect(ctx); err != nil {
		log.Panic().AnErr("connecting to mongodb", err)
	}
	defer func() {
		if err = mc.Disconnect(ctx); err != nil {
			log.Error().AnErr("disconnecting mongodb", err)
		}
	}()
	st, err := device.NewMongoStore(ctx, mc.Database("example").Collection("device"))
	if err != nil {
		log.Panic().AnErr("creating mongo store", err)
	}
	log.Info().Msg("mongodb store initialized")

	// Kafka Producer
	kp, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "127.0.0.1:9092"})
	if err != nil {
		log.Panic().AnErr("connecting to kafka", err)
	}
	pub := device.NewKafkaPublisher(kp, "example.device.v1.pb")
	log.Info().Msg("kafka publisher initialized")

	// starting the service
	s := device.NewService(pub, st)
	log.Info().Msg("producer service started")

	log.Info().Msg("adding 5 updatable devices")
	for i := 0; i < 5; i++ {
		d := generateRandomDevice()
		d.Id = fmt.Sprint(i + 1)
		err := s.RegisterDevice(d)
		if err != nil {
			log.Error().Err(err).Msg("error on adding device")
		}
	}
	go func() {
		log.Info().Msg("adding 1000000 random devices")
		for i := 0; i < 1000000; i++ {
			if err := s.RegisterDevice(generateRandomDevice()); err != nil {
				log.Error().Err(err).Msg("error on adding device")
			}
		}
		log.Info().Msg("========= all 1000000 random devices are added ==========")
	}()
	go func() {
		for {
			if err := s.RegisterDevice(generateRandomDevice()); err != nil {
				log.Error().Err(err).Msg("error on adding device")
			}
			time.Sleep(time.Millisecond * time.Duration(500+rand.Intn(1000)))
		}
	}()
	go func() {
		for {
			d := generateRandomDevice()
			d.Id = fmt.Sprint(rand.Intn(5) + 1)
			log.Info().Msgf("updating device %s", d.Id)
			err := s.RegisterDevice(d)
			if err != nil {
				log.Error().Err(err).Msg("error on adding device")
			}
			time.Sleep(time.Second * time.Duration(10+rand.Intn(10)))
		}
	}()

	sig := <-stopCh
	log.Info().Msgf("stopping the device service. Signal: %s", sig)
	defer close(stopCh)

}

func generateRandomDevice() *api.Device {
	id := ksuid.New().String()
	oid := ksuid.New().String()
	i, _ := ptypes.TimestampProto(time.Now().AddDate(0, 0, -3*rand.Intn(100)))
	u, _ := ptypes.TimestampProto(time.Now())
	d := &api.Device{
		Id:             string(id[:]),
		OrganizationId: string(oid[:]),
		Tz:             getRandomTimezone(),
		InstalledAt:    i,
		UpdatedAt:      u,
	}
	return d
}

func getRandomTimezone() string {
	tz := []string{"America/Los_Angeles", "America/Montreal", "America/Sao_Paulo", "Europe/Helsinki"}
	return tz[rand.Intn(len(tz))]
}
