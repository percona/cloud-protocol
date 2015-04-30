package proto

import (
	"time"
)

type AgentConfig struct {
	Service string
	UUID    string
	Config  string // JSON
	Running bool
	Updated time.Time
}
