package proto

import (
	"time"
)

type Data struct {
	ProtocolVersion string    `json:",omitempty"` // protocol version in use
	Created         time.Time // when Data was spooled (UTC)
	Hostname        string    // OS instance name the agent is running on
	Tool            string    // "qan" or "mm"
	ContentType     string    // of Data ("application/json")
	ContentEncoding string    // of Data ("gzip" or empty)
	Data            []byte    // encoded qan.Report or mm.Report
}

type Response struct {
	Code  uint   // standard HTTP status (http://httpstatus.es/)
	Error string // empty if ok (Code=200)
}
