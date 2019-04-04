package device

import (
	"time"
)

type Device struct {
	ID             string
	OrganizationID string
	TZ             *time.Location
}
