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

package fakeapi

import (
	"encoding/json"
	"fmt"
	"github.com/percona/cloud-protocol/proto/v1"
	"github.com/percona/cloud-protocol/proto/v1/agent/mm"
	"github.com/percona/cloud-protocol/proto/v1/agent/mm/mysql"
	"github.com/percona/cloud-protocol/proto/v1/agent/mm/system"
	"github.com/percona/cloud-protocol/proto/v1/agent/sysconfig"
	sysconfigMysql "github.com/percona/cloud-protocol/proto/v1/agent/sysconfig/mysql"
	"io/ioutil"
	"net/http"
)

var (
	ConfigMmDefaultMysql = mysql.Config{
		Config: mm.Config{
			Collect: 1,
			Report:  60,
		},
		UserStats: false,
	}
	ConfigMmDefaultServer = system.Config{
		Config: mm.Config{
			Collect: 1,
			Report:  60,
		},
	}
	ConfigSysconfigDefaultMysql = sysconfigMysql.Config{
		Config: sysconfig.Config{
			Report: 3600,
		},
	}
)

func (f *FakeApi) AppendPing() {
	f.Append("/ping", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(600)
		}
	})
}

func (f *FakeApi) AppendInstancesServer(id uint, serverInstance *proto.ServerInstance) {
	f.Append("/instances/server", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(body, serverInstance)
		if err != nil {
			panic(err)
		}
		serverInstance.Id = id
		w.Header().Set("Location", fmt.Sprintf("%s/instances/server/%d", f.URL(), serverInstance.Id))
		w.WriteHeader(http.StatusCreated)
	})
}
func (f *FakeApi) AppendInstancesServerId(id uint, serverInstance *proto.ServerInstance) {
	f.Append(fmt.Sprintf("/instances/server/%d", id), func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		data, _ := json.Marshal(&serverInstance)
		w.Write(data)
	})
}
func (f *FakeApi) AppendInstancesMysql(id uint, mysqlInstance *proto.MySQLInstance) {
	f.Append("/instances/mysql", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(body, mysqlInstance)
		if err != nil {
			panic(err)
		}
		mysqlInstance.Id = id
		w.Header().Set("Location", fmt.Sprintf("%s/instances/mysql/%d", f.URL(), mysqlInstance.Id))
		w.WriteHeader(http.StatusCreated)
	})
}
func (f *FakeApi) AppendInstancesMysqlId(id uint, mysqlInstance *proto.MySQLInstance) {
	f.Append(fmt.Sprintf("/instances/mysql/%d", id), func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		data, _ := json.Marshal(&mysqlInstance)
		w.Write(data)
	})
}
func (f *FakeApi) AppendConfigsMmDefaultServer() {
	f.Append("/configs/mm/default-server", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		data, _ := json.Marshal(&ConfigMmDefaultServer)
		w.Write(data)
	})
}
func (f *FakeApi) AppendConfigsMmDefaultMysql() {
	f.Append("/configs/mm/default-mysql", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		data, _ := json.Marshal(&ConfigMmDefaultMysql)
		w.Write(data)
	})
}
func (f *FakeApi) AppendSysconfigDefaultMysql() {
	f.Append("/configs/sysconfig/default-mysql", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		data, _ := json.Marshal(&ConfigSysconfigDefaultMysql)
		w.Write(data)
	})
}
func (f *FakeApi) AppendConfigsQanDefault() {
	f.Append("/configs/qan/default", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{ "Service": "mysql", "InstanceId": 0, "Interval": 60}`))
	})
}
func (f *FakeApi) AppendAgents(agent *proto.Agent) {
	f.Append("/agents", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", fmt.Sprintf("%s/agents/%s", f.URL(), agent.Uuid))
		w.WriteHeader(http.StatusCreated)
	})
}
func (f *FakeApi) AppendAgentsUuid(agent *proto.Agent) {
	f.Append(fmt.Sprintf("/agents/%s", agent.Uuid), func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		data, _ := json.Marshal(&agent)
		w.Write(data)
	})
}
