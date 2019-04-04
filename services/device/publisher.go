package device

import (
	"errors"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gogo/protobuf/proto"
	"github.com/nassor/kafka-compact-pipeline/api"
)

type KafkaPublisher struct {
	kp    *kafka.Producer
	topic string
}

// NewKafkaPublisher creates a new instance of Producer
func NewKafkaPublisher(kp *kafka.Producer, topic string) *KafkaPublisher {
	return &KafkaPublisher{
		kp:    kp,
		topic: topic,
	}
}

// Publish send storeinfo data to a kafka topic
func (p KafkaPublisher) Publish(d *api.Device) error {
	b, err := proto.Marshal(d)
	if err != nil {
		return fmt.Errorf("marshaling error: %s", err)
	}

	event := make(chan kafka.Event)
	defer close(event)
	err = p.kp.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &p.topic, Partition: 0},
		Value:          b,
		Key:            []byte(d.Id),
	}, event)
	if err != nil {
		return fmt.Errorf("producing kafka msg: %s", err)
	}

	select {
	case e := <-event:
		km := e.(*kafka.Message)
		if km.TopicPartition.Error != nil {
			return fmt.Errorf("delivery failed: %s", km.TopicPartition.Error)
		}
	case <-time.After(10 * time.Second):
		return errors.New("delivery timed out")
	}

	return nil
}
