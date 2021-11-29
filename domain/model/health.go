package model

import "time"

// Health - Contains the following fields:
// - Uptime: Uptime of project since last build in ms.
// - StatusCode: HTTPStatus.
type Health struct {
	Uptime     time.Duration
	StatusCode int
}
