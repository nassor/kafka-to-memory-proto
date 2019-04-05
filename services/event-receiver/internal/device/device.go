package device

import (
	"time"

	"github.com/vmihailenco/msgpack"
)

type Device struct {
	ID             string
	OrganizationID string
	TZ             *time.Location
}

func (d *Device) Marshal() ([]byte, error) {
	return msgpack.Marshal(d)
}

func (d *Device) Unmarshal(b []byte) error {
	return msgpack.Unmarshal(b, d)
}

func (d *Device) Reset() {
	*d = Device{}
}
