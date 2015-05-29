package qan

import protoV1 "github.com/percona/cloud-protocol/proto/v1"

type MetricsSummary struct {
    Metrics    []Metric
    QueryCount int64
}

type Metric struct {
	Name        string
	Total       protoV1.NullFloat32
	Avg         protoV1.NullFloat32
	Max         protoV1.NullFloat32
	Min         protoV1.NullFloat32
	Pct_95      protoV1.NullFloat32
	Stddev      protoV1.NullFloat32
	Median      protoV1.NullFloat32
}
