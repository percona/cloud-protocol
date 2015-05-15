package reports

import "github.com/percona/cloud-protocol/proto"

type EmailReport struct {
	Meta   Meta
	Report Report
}

type Meta struct {
	Id       uint
	Uuid     string
	TimeZone string
	Begin    string
	End      string
}

type Report struct {
	Snapshot         *Snapshot
	Chart            *Chart
	SlowQueries      *SlowQueries
	HighLoadQueries  *HighLoadQueries
	FirstSeenQueries *FirstSeenQueries
}

type Snapshot struct {
	Queries              int64
	QueriesPct           float64
	QueryLoad            float64
	QueryLoadPct         float64
	ResponseTimePct95    float64
	ResponseTimePct95Pct float64
	ResponseTimeAvg      float64
	ResponseTimeAvgPct   float64
}

type Chart struct {
	ChartPoints []ChartPoint
}

type ChartPoint struct {
	Datetime int64
	Metric   float32
	QPS      float64
	Count    int64
}

type SlowQueries struct {
	Queries
}

type HighLoadQueries struct {
	Queries
}
type FirstSeenQueries struct {
	Queries
}

type Queries struct {
	Queries []Query
}

type Query struct {
	Distillate        string
	Checksum          string
	Fingerprint       string
	FirstSeen         string
	LastSeen          string
	ClassId           int
	Rank              string // +x, -x, 0, new
	MetricSum         float64
	MetricSumChange   proto.NullFloat64
	MetricPct95       float64
	MetricPct95Change proto.NullFloat64
	MetricAvg         float64
	MetricAvgChange   proto.NullFloat64
	MetricMin         float32
	MetricMinChange   proto.NullFloat64
	MetricMax         float32
	MetricMaxChange   proto.NullFloat64
	MetricCount       int64
	MetricCountChange proto.NullFloat64
	Load              float64           // Load of the query
	LoadPct           float64           // Load of the query vs total load in given period
	LoadChange        proto.NullFloat64 // Change in load vs previous period
}
