package proto

import (
	"time"
)

type QanReport struct {
	AgentUuid     string
	StartTs       time.Time // UTC
	EndTs         time.Time // UTC
	SlowLogFile   string    // not slow_query_log_file if rotated
	StartOffset   int64    // parsing starts
	EndOffset     int64    // parsing stops, but...
	StopOffset    int64    // ...parsing didn't complete if stop < end
	RunTime       float64 // seconds
	Global []byte
	Class []byte
}

type MmReport struct {
	AgentUuid string
	Ts        time.Time // UTC
	Duration  uint      // seconds
	Metrics   []byte
}

type Transmission struct {
	Created   time.Time // when Data was spool (UTC)
	Hostname  string
	Service   string    // qan or mm
    Data      []byte    // encoded qan.Report or mm.Report
}
