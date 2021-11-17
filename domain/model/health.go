package model

import "time"

type Health struct {
	Uptime     time.Duration
	StatusCode int
}
