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

type Stats struct {
	metricType string    `json:"-"` // ignore
	str        string    `json:",omitempty"`
	firstVal   bool      `json:"-"`
	prevTs     int64     `json:"-"`
	penuTs     int64     `json:"-"`
	prevVal    float64   `json:"-"` // last value
	penuVal    float64   `json:"-"` // 2nd to last (penultimate) value
	vals       []float64 `json:"-"`
	sum        float64   `json:"-"`
	Cnt        int
	Min        float64
	Pct5       float64
	Avg        float64
	Med        float64
	Pct95      float64
	Max        float64
}
