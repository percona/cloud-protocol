package proto

import (
	"time"
)

type Transmission struct {
	Created   time.Time // when Data was spool (UTC)
	Hostname  string
	Service   string    // qan or mm
    Data      []byte    // encoded qan.Report or mm.Report
}
