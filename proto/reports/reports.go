package reports

type EmailReport struct {
	Meta   Meta
	Report Report
}

type Meta struct {
	Id       uint
	Uuid	 string
	TimeZone string
	Begin    string
	End      string
}

type Report struct {
	Snapshot         Snapshot
	Chart            Chart
	SlowQueries      SlowQueries
	HighLoadQueries  HighLoadQueries
	FirstSeenQueries FirstSeenQueries
}

type Snapshot struct {
	Queries              int64
	QueriesPct           float64
	QueryLoad            float64
	QueryLoadPct         float64
	ResponseTimePct95    float64
	ResponseTimePct95Pct float64
}

type Chart struct {
	ChartPoints []ChartPoint
}

type ChartPoint struct {
	Datetime int64
	Metric   float64
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
	Distillate  string
	Checksum    string
	Fingerprint string
	FirstSeen   string
	LastSeen    string
	ClassId     string
	MetricSum   float64
	MetricPct95 float64
	MetricAvg   float64
	MetricMin   float64
	MetricMax   float64
	MetricCount int64
}
