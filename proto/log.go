package proto

import (
	"encoding/json"
	"time"
)

// http://en.wikipedia.org/wiki/Syslog#Severity_levels
const (
	LOG_EMERGENCY byte = iota // not used
	LOG_ALERT                 // not used
	LOG_CRITICAL
	LOG_ERROR
	LOG_WARNING
	LOG_NOTICE // not used
	LOG_INFO   // default
	LOG_DEBUG
)

var LogLevels map[string]byte = map[string]byte{
	"emergency": LOG_EMERGENCY,
	"alert":     LOG_ALERT,
	"critical":  LOG_CRITICAL,
	"error":     LOG_ERROR,
	"warning":   LOG_WARNING,
	"notice":    LOG_NOTICE,
	"info":      LOG_INFO,
	"debug":     LOG_DEBUG,
}

type LogEntry struct {
	Ts      time.Time
	Level   byte
	Service string
	Msg     string
	Cmd     string `json:",omitempty"`
}

func (e *LogEntry) String() string {
	bytes, err := json.Marshal(e)
	if err != nil {
		return ""
	}
	return string(bytes)
}

type GetAgentLog struct {
	Uuid   string
	Limit  uint
	Period uint
}
