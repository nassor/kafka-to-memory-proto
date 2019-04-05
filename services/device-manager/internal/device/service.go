package device

import (
	"context"
	"fmt"

	"github.com/nassor/kafka-to-memory-proto/api"
)

// Publisher represents actions to publish device events
type Publisher interface {
	Publish(d *api.Device) error
}

// Storer represents actions to store device events
type Storer interface {
	Upsert(ctx context.Context, d *api.Device) error
}

// Service is the integration of all device systems
type Service struct {
	pub Publisher
	st  Storer
}

// NewService creates a new device-manager Service
func NewService(pub Publisher, st Storer) *Service {
	return &Service{pub: pub, st: st}
}

// RegisterDevice insert or updates a device, publishing the changes after
func (s *Service) RegisterDevice(d *api.Device) error {
	ctx := context.Background()
	if err := s.st.Upsert(ctx, d); err != nil {
		return fmt.Errorf("when storing a device %s", err)
	}
	if err := s.pub.Publish(d); err != nil {
		return fmt.Errorf("when publishing a device %s", err)
	}
	return nil
}
