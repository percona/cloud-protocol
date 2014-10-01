package reports


type EmailReport struct {
	Meta Meta
	Report Report
}

type Meta struct {
	Id uint
	TimeZone string
	Begin string
	End string
}

type Report struct {
	Snapshot Snapshot
	SlowQueries []Query
	HighLoadQueries []Query
	NewQueries []Query
}

type Snapshot struct {
	Queries int64
	QueriesPct float64
	QueryLoad float64
	QueryLoadPct float64
	ResponseTimePct95 float64
	ResponseTimePct95Pct float64
}

type Query struct {
	ClassId string
	MetricSum float64
	MetricPct95 float64
	MetricAvg float64
	MetricMax float64
	MetricCount int64
}

