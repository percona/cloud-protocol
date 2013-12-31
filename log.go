package proto

import (
	"fmt"
	"time"
)

// http://en.wikipedia.org/wiki/Syslog#Severity_levels
const (
	LOG_EMERGENCY int = iota // not used
	LOG_ALERT                // not used
	LOG_CRITICAL
	LOG_ERROR
	LOG_WARNING
	LOG_NOTICE // not used
	LOG_INFO   // default
	LOG_DEBUG
)

var LogLevels map[string]int = map[string]int{
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
	Level   int
	Service string
	Msg     string
	Cmd     string `json:",omitempty"`
}

func (e *LogEntry) String() string {
	return fmt.Sprintf("level:%d service:%s msg: %s",
		e.Level, e.Service, e.Msg)
}
