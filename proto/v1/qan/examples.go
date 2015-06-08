package qan

import (
	protoV1 "github.com/percona/cloud-protocol/proto/v1"
)

type Example struct {
	Db        string
	Query     string
	QueryTime protoV1.NullFloat64
	Ts        protoV1.NullInt64
}
