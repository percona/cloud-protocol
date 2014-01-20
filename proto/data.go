package proto

import (
	"time"
)

type QanData struct {
	AgentUuid     string
	StartTs       time.Time // UTC
	EndTs         time.Time // UTC
	SlowLogFile   string    // not slow_query_log_file if rotated
	StartOffset   uint64    // parsing starts
	EndOffset     uint64    // parsing stops, but...
	StopOffset    uint64    // ...parsing didn't complete if stop < end
	RunTime       uint      // seconds
	GlobalMetrics []byte
	ClassMetrics  []byte
}

type MmData struct {
	AgentUuid string
	Ts        time.Time // UTC
	Duration  uint      // seconds
	Metrics   []byte
}
