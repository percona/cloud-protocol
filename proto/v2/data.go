package proto

import (
	"time"

	"github.com/percona/go-mysql/event"
)

// Data struct to send an encoded QANReport or MMReport
type Data struct {
	ProtocolVersion string    `json:",omitempty"` // protocol version in use
	Created         time.Time // when Data was spooled (UTC)
	Hostname        string    // OS instance name the agent is running on
	Service         string    // "qan", "mm", "log"
	ContentType     string    // of Data ("application/json")
	ContentEncoding string    // of Data ("gzip" or empty)
	Data            []byte    // encoded QANReport or MMReport
}

type Response struct {
	Code  uint   // standard HTTP status (http://httpstatus.es/)
	Error string // empty if ok (Code=200)
}

type DataSpoolLimits struct {
	MaxAge   uint   // seconds
	MaxSize  uint64 // bytes
	MaxFiles uint
}

// QAN data struct, sent by the agent
type QANReport struct {
	UUID    string              // UUID of MySQL instance
	StartTs time.Time           // Start time of interval, UTC
	EndTs   time.Time           // Stop time of interval, UTC
	RunTime float64             // Time parsing data, seconds
	Global  *event.GlobalClass  // Metrics for all data
	Class   []*event.QueryClass // per-class metrics
	// slow log:
	SlowLogFile     string `json:",omitempty"` // not slow_query_log_file if rotated
	SlowLogFileSize int64  `json:",omitempty"`
	StartOffset     int64  `json:",omitempty"` // parsing starts
	EndOffset       int64  `json:",omitempty"` // parsing stops, but...
	StopOffset      int64  `json:",omitempty"` // ...parsing didn't complete if stop < end
}

// MM data struct, sent by the agent
type MMReport struct {
	Ts       time.Time // start, UTC
	Duration uint      // seconds
	Stats    []*InstanceStats
}

// Stats for each metric from a service instance, computed at each report interval and sent by agent as part of MMReport.
type InstanceStats struct {
	UUID  string
	Stats map[string]*Stats // keyed on metric name
}

// MM stats for interval, sent by agent as part or InstanceStats
type Stats struct {
	Cnt   int
	Min   float64
	Pct5  float64
	Avg   float64
	Med   float64
	Pct95 float64
	Max   float64
}
