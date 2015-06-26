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
	"time"
)

// An Interval represents a period during which queries are fetched,
// aggregated, and reported. Intervals are created and sent by an IntervalIter.
// Analyzers use Intervals to set up and run Workers.
type Interval struct {
	Number    int
	StartTime time.Time // UTC
	StopTime  time.Time // UTC
	// slowlog
	Filename    string // slow_query_log_file
	StartOffset int64  // bytes @ StartTime
	EndOffset   int64  // bytes @ StopTime
}
