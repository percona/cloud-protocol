/*
   Copyright (c) 2014-2015, Percona LLC and/or its affiliates. All rights reserved.

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>
*/

package qan

import (
	"github.com/percona/cloud-protocol/proto/v1"
	"github.com/percona/cloud-protocol/proto/v1/agent/mysql"
)

type Config struct {
	proto.ServiceInstance
	// Manager
	CollectFrom       string // "slowlog" or "perfschema"
	Start             []mysql.Query
	Stop              []mysql.Query
	MaxWorkers        int
	Interval          uint  // minutes, "How often to report"
	MaxSlowLogSize    int64 // bytes, 0 = no max
	RemoveOldSlowLogs bool  // after rotating for MaxSlowLogSize
	// Worker
	ExampleQueries bool // only fingerprints if false
	WorkerRunTime  uint // seconds
	// Report
	ReportLimit uint
}
