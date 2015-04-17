package proto

import (
	"time"
)

type AgentConfig struct {
	Tool    string
	UUID    string
	Config  string // JSON
	Running bool
	Updated time.Time
}
