package device

import (
	"fmt"
	"os"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gogo/protobuf/proto"
	"github.com/rs/zerolog/log"

	"github.com/nassor/kafka-compact-pipeline/api"
)

type KafkaSubscriber struct {
	kc        *kafka.Consumer
	loaded    *sync.Once
	loadedCh  chan uint64
	receiveCh chan api.Device
}

// NewKafkaSubscriber creates a new instance of KafkaSubscriber
func NewKafkaSubscriber(kc *kafka.Consumer, topic string) (*KafkaSubscriber, error) {
	kc.Unsubscribe()
	if err := kc.Subscribe(topic, nil); err != nil {
		return nil, err
	}

	return &KafkaSubscriber{
		kc:        kc,
		loaded:    &sync.Once{},
		loadedCh:  make(chan uint64),
		receiveCh: make(chan api.Device),
	}, nil
}

// Listen for kafka events
func (s *KafkaSubscriber) Listen() {
	var total uint64
	for {
		event := s.kc.Poll(100)
		if event == nil {
			continue
		}

		switch e := event.(type) {
		case *kafka.Message:
			total++
			if err := s.receive(e.Value); err != nil {
				log.Error().Err(err).Msg("error on receiving event")
			}
		case kafka.PartitionEOF:
			s.loaded.Do(func() {
				s.loadedCh <- total
			})
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			break
		default:
			fmt.Printf("Ignored %v\n", e)
		}
	}
}

// IsLoaded return a channel that will return the total messages consumed
func (s *KafkaSubscriber) IsLoaded() <-chan uint64 {
	return s.loadedCh
}

// Receive return a channel of messages received
func (s *KafkaSubscriber) Receive() <-chan api.Device {
	return s.receiveCh
}

func (s *KafkaSubscriber) receive(data []byte) error {
	var d api.Device
	if err := proto.Unmarshal(data, &d); err != nil {
		return err
	}
	if d.Id == "1" || d.Id == "2" || d.Id == "3" || d.Id == "4" || d.Id == "5" {
		log.Debug().Msgf("received update for device %s", d.Id)
	}
	s.receiveCh <- d
	return nil
}
