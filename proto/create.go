package proto

import (
	"time"
)

type AgentConfig struct {
	ServiceInstance
	Config  string // JSON
	Running bool
	Updated time.Time
}

type Agent struct {
	Uuid     string
	Hostname string
	Alias    string
	Version  string
	Configs  map[string]AgentConfig `json:",omitempty"`
	QAN      bool
}
