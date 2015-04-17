package proto

import (
	"time"
)

type Data struct {
	ProtocolVersion string    `json:",omitempty"` // protocol version in use
	Created         time.Time // when Data was spooled (UTC)
	OSUUID          string    // OS instance the agent is running on
	Tool            string    // "qan" or "mm"
	ContentType     string    // of Data ("application/json")
	ContentEncoding string    // of Data ("gzip" or empty)
	Data            []byte    // encoded qan.Report or mm.Report
	AgentVersion    string    `json:",omitempty"` // version of the agent producing the data
}

type Response struct {
	Code  uint   // standard HTTP status (http://httpstatus.es/)
	Error string // empty if ok (Code=200)
}
