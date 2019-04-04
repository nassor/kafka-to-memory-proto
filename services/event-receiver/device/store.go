package device

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/nassor/kafka-compact-pipeline/api"
)

// InMemoryStore keep device data in a map
type InMemoryStore struct {
	data  *sync.Map
	total uint64
}

func NewInMemoryStore() *InMemoryStore {
	st := &InMemoryStore{data: &sync.Map{}}
	go func() {
		for {
			log.Debug().Msgf("total stored: %d", atomic.LoadUint64(&st.total))
			time.Sleep(20 * time.Second)
		}
	}()
	return st
}

// Upsert stores or update device information
func (st *InMemoryStore) Upsert(dev *api.Device) error {
	tz, err := time.LoadLocation(dev.Tz)
	if err != nil {
		return err
	}
	d := &Device{
		ID:             dev.Id,
		OrganizationID: dev.OrganizationId,
		TZ:             tz,
	}
	if _, ok := st.data.Load(d.ID); !ok {
		st.total = atomic.AddUint64(&st.total, 1)
	}
	st.data.Store(d.ID, d)
	return nil
}

// Get retrieve device information by id
func (st *InMemoryStore) Get(id string) (*Device, error) {
	v, ok := st.data.Load(id)
	if !ok {
		return nil, errors.New("device not found")
	}
	d, ok := v.(*Device)
	if !ok {
		return nil, errors.New("failed on convert in memory data to device")
	}
	return d, nil
}
