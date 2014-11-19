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
	Uuid       string
	Hostname   string
	Version    string
	Running    bool              // connected to API?
	QAN        bool              // has QAN config
	QANRunning bool              // is QAN running or not?
	Configs    []AgentConfig     `json:",omitempty"`
	Links      map[string]string `json:",omitempty"`
}
