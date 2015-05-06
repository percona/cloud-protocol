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

// QAN data struct, sent by the agent
type QANReport struct {
	UUID    string              // UUID of MySQL instance
	StartTs time.Time           // of interval, UTC
	EndTs   time.Time           // of interval, UTC
	RunTime float64             // seconds parsing data
	Global  *event.GlobalClass  // metrics for all data
	Class   []*event.QueryClass // per-class metrics
	// slow log:
	SlowLogFile     string `json:",omitempty"` // not slow_query_log_file if rotated
	SlowLogFileSize int64  `json:",omitempty"`
	StartOffset     int64  `json:",omitempty"` // parsing starts
	EndOffset       int64  `json:",omitempty"` // parsing stops, but...
	StopOffset      int64  `json:",omitempty"` // ...parsing didn't complete if stop < end
}
