package proto

import (
	"time"
)

type AgentConfig struct {
	InternalService string
	ExternalService ServiceInstance
	Config          string // JSON
	Running         bool
	Updated         time.Time
}

type Agent struct {
	Uuid     string
	Hostname string
	Version  string
	Configs  []AgentConfig `json:",omitempty"`
	QAN      bool
	Links    map[string]string `json:",omitempty"`
	Running  bool
}
