package qan

import "github.com/percona/cloud-protocol/proto"

type QueryPlan struct {
	Metrics []QueryPlanMetric
}

type QueryPlanMetric struct {
	Name  string
	Value proto.NullInt64
}
