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

package mm

import (
	"time"

	"github.com/percona/cloud-protocol/proto/v1"
)

// Stats for each metric from a service instance, computed at each report interval.
type InstanceStats struct {
	proto.ServiceInstance
	Stats map[string]*Stats // keyed on metric name
}

type Report struct {
	Ts       time.Time // start, UTC
	Duration uint      // seconds
	Stats    []*InstanceStats
}
