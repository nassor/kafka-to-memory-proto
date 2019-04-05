package device

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/nassor/kafka-to-memory-proto/api"
)

func TestService_Start(t *testing.T) {
	t.Run("receive and upsert data successfully", func(t *testing.T) {
		sub := &subscriber{}
		ch := make(chan api.Device)
		sub.On("Receive").Return(ch).Once()

		st := &storer{}
		st.On("Upsert", mock.AnythingOfType("*api.Device")).Return(nil).Once()

		s := NewService(sub, st)
		go s.Start()
		ch <- api.Device{}
		s.Stop()

		sub.AssertExpectations(t)
		st.AssertExpectations(t)
	})

	t.Run("receive and but upsert data can fail", func(t *testing.T) {
		sub := &subscriber{}
		ch := make(chan api.Device)
		sub.On("Receive").Return(ch).Once()

		st := &storer{}
		st.On("Upsert", mock.AnythingOfType("*api.Device")).Return(errors.New("an error")).Once()

		s := NewService(sub, st)
		go s.Start()
		ch <- api.Device{}
		s.Stop()

		sub.AssertExpectations(t)
		st.AssertExpectations(t)
	})
}

type subscriber struct {
	mock.Mock
}

func (sub *subscriber) Receive() <-chan api.Device {
	args := sub.Called()
	return args.Get(0).(chan api.Device)
}

type storer struct {
	mock.Mock
}

func (st *storer) Upsert(d *api.Device) error {
	args := st.Called(d)
	return args.Error(0)
}
