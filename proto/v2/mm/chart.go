package mm

import (
	protoV1 "github.com/percona/cloud-protocol/proto/v1"
	"time"
)

type ChartArgs struct {
	Id        uint
	Metrics   []Metric // {Name:"bytes_read", Stat:"pct95"}
	Sysconfig []Sysconfig
	Range     string // "1 week" (see timeseries.ParsePeriod())
	Begin     string // ... 2014-01-01 00:00:00 UTC
	End       string // ... 2014-01-07 00:00:00 UTC
	Group     string // "1 day" (see timeseries.ParsePeriod())
	// --
	begin    time.Time
	end      time.Time
	realtime bool
}

type ChartData struct {
	Id           uint
	Error        string
	Series       []Series
	ConnectionId string
	ResponseTime float64
}

// A metric is name and a stat, e.g. avg bytes_read, requested by user.
type Metric struct {
	protoV1.ServiceInstance
	Name   string // bytes_read
	Stat   string // avg
	Error  error
	Values []Value
}

type Sysconfig struct {
	protoV1.ServiceInstance
	System string // mysql global variables
	Name   string // max_connections
	Error  error
	Values []Value
}

// Raw metric values are fetched from the database.
type Value struct {
	Ts       int64   // UTC Unix timestamp
	Duration uint    // seconds
	Value    float64 // value of min, max, etc.
}

// For now, data points are timestamps on the x-axis.
// Other chart types, like histograms, would use differnet x-axis values.
type DataPoint struct {
	Begin    int64   // UTC Unix timestamp
	End      int64   // UTC Unix timestamp
	Count    uint    // number of data values
	Duration uint    // sum of Value durations
	Y        float64 // y-axis value, i.e. summarized value of metric stat @ x
}

type Series struct {
	Error string
	Data  []DataPoint
}
