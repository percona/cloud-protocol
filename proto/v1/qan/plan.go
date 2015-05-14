package qan

import "github.com/percona/cloud-protocol/proto/v1"

type QueryPlan struct {
	Metrics []QueryPlanMetric
}

type QueryPlanMetric struct {
	Name  string
	Value proto.NullInt64
}
