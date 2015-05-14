package proto

import (
	"time"
)

type Data struct {
	Created         time.Time // when Data was spooled (UTC)
	Hostname        string    // that agent is running on
	Service         string    // "qan" or "mm"
	ContentType     string    // of Data ("application/json")
	ContentEncoding string    // of Data ("gzip" or empty)
	Data            []byte    // encoded qan.Report or mm.Report
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
