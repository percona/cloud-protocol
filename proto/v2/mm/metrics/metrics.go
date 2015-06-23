package metrics

import (
	protoV2 "github.com/percona/cloud-protocol/proto/v2"
)

type Metrics []Metric

type Metric struct {
	Id        uint
	Subsystem protoV2.Subsystem
	Name      string
	Title     string
	Type      string
	Unit      string
}

type Configs []Config

type Config struct {
	Id        uint
	Subsystem protoV2.Subsystem
	Name      string
	Title     string
}
