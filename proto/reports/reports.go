package reports

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
	ResponseTimeAvg      float64
	ResponseTimeAvgPct   float64
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
	Distillate        string
	Checksum          string
	Fingerprint       string
	FirstSeen         string
	LastSeen          string
	ClassId           int
	Rank              string // +x, -x, 0, new
	MetricSum         float64
	MetricSumChange   float64
	MetricPct95       float64
	MetricPct95Change float64
	MetricAvg         float64
	MetricAvgChange   float64
	MetricMin         float64
	MetricMinChange   float64
	MetricMax         float64
	MetricMaxChange   float64
	MetricCount       int64
	MetricCountChange float64
	Load              float64 // Load of the query
	LoadPct           float64 // Load of the query vs total load in given period
	LoadChange        float64 // Change in load vs previous period
}
