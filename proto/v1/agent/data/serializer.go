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

package data

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
)

type Serializer interface {
	ToBytes(data interface{}) ([]byte, error)
	Encoding() string
	Concurrent() bool
}

type JsonGzipSerializer struct {
	e *json.Encoder
	g *gzip.Writer
	b *bytes.Buffer
}

func NewJsonGzipSerializer() *JsonGzipSerializer {
	b := &bytes.Buffer{}    // 4. buffer
	g := gzip.NewWriter(b)  // 3. gzip
	e := json.NewEncoder(g) // 2. encode
	// ....................... 1. data

	s := &JsonGzipSerializer{
		e: e,
		g: g,
		b: b,
	}
	return s
}

func (s *JsonGzipSerializer) ToBytes(data interface{}) ([]byte, error) {
	s.b.Reset()
	s.g.Reset(s.b)
	if err := s.e.Encode(data); err != nil {
		return nil, err
	}
	s.g.Close()
	return s.b.Bytes(), nil
}

func (s *JsonGzipSerializer) Encoding() string {
	return "gzip"
}

func (s *JsonGzipSerializer) Concurrent() bool {
	return false
}

// --------------------------------------------------------------------------

type JsonSerializer struct {
}

func NewJsonSerializer() *JsonSerializer {
	j := &JsonSerializer{}
	return j
}

func (j *JsonSerializer) ToBytes(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func (s *JsonSerializer) Encoding() string {
	return ""
}

func (s *JsonSerializer) Concurrent() bool {
	return true
}
