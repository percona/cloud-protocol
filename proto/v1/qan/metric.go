package qan

import protoV1 "github.com/percona/cloud-protocol/proto/v1"

type MetricsSummary struct {
	Metrics    []Metric
	QueryCount protoV1.NullInt64
}

type Metric struct {
	Name        string
	Total       float32
	Avg         float32
	Max         float32
	Min         float32
	Pct_95      float32
	Stddev      float32
	Median      float32
	Query_count int
}
