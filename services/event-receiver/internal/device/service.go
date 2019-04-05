package device

import (
	"github.com/rs/zerolog/log"

	"github.com/nassor/kafka-to-memory-proto/api"
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
	sub  DeviceSubscriber
	st   DeviceStorer
	stop chan struct{}
}

// NewService creates a new event-receiver Service
func NewService(sub DeviceSubscriber, st DeviceStorer) *Service {
	return &Service{sub: sub, st: st, stop: make(chan struct{})}
}

// Start receive device data from the subscriber and send to the store
func (s *Service) Start() {
	rcvCh := s.sub.Receive()
	for {
		select {
		case d := <-rcvCh:
			if err := s.st.Upsert(&d); err != nil {
				log.Error().Err(err).Msg("on receiving device information")
			}
		case <-s.stop:
			return
		}
	}
}

// Stop the device service
func (s *Service) Stop() {
	s.stop <- struct{}{}
}
