package device

import (
	"errors"
	"sync"
	"time"

	"github.com/nassor/kafka-to-memory-proto/api"
)

// InMemoryStore keep device data in a map
type InMemoryStore struct {
	store      *sync.Map
	devicePool *sync.Pool
}

func NewInMemoryStore() *InMemoryStore {
	st := &InMemoryStore{
		store: &sync.Map{},
		devicePool: &sync.Pool{
			New: func() interface{} {
				return new(Device)
			},
		},
	}
	return st
}

// Upsert stores or update device information
func (st *InMemoryStore) Upsert(dev *api.Device) error {
	tz, err := time.LoadLocation(dev.Tz)
	if err != nil {
		return err
	}

	d := st.devicePool.Get().(*Device)
	d.Reset()
	d.ID = dev.Id
	d.OrganizationID = dev.OrganizationId
	d.TZ = tz

	data, err := d.Marshal()
	if err != nil {
		return err
	}
	st.store.Store(d.ID, data)
	return nil
}

// Get retrieve device information by id
func (st *InMemoryStore) Get(id string) (*Device, error) {
	v, ok := st.store.Load(id)
	if !ok {
		return nil, errors.New("device not found")
	}
	data, ok := v.([]byte)
	if !ok {
		return nil, errors.New("mapped data is not an byte slice")
	}

	d := st.devicePool.Get().(*Device)
	d.Reset()
	if err := d.Unmarshal(data); err != nil {
		return nil, err
	}
	return d, nil
}
