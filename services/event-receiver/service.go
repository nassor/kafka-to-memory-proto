package receiver

import (
	"github.com/rs/zerolog/log"

	"github.com/nassor/kafka-compact-pipeline/api"
)

// DeviceSubscriber represent action for subscribe events
type DeviceSubscriber interface {
	Receive() <-chan api.Device
}

// DeviceStorer represents actions to store device events
type DeviceStorer interface {
	Upsert(d *api.Device) error
}

// Service is the integration of all event-receiver systems
type Service struct {
	sub DeviceSubscriber
	st  DeviceStorer
}

// NewService creates a new event-receiver Service
func NewService(sub DeviceSubscriber, st DeviceStorer) *Service {
	return &Service{sub: sub, st: st}
}

func (s *Service) ReceiveData() {
	for d := range s.sub.Receive() {
		if err := s.st.Upsert(&d); err != nil {
			log.Error().Err(err).Msg("on receiving device information")
		}
	}
}
