percona/cloud-protocol
======================

Shared types between cloud-api and cloud-tools.

## Monitors

To start a monitor:

```
proto.Cmd:
    Service: mm
    Cmd:     StartService
    Data:    mm.Config:
        Name:    db1
        Type:    mysql
        Collect: 1
        Report:  60
        Config: mm/mysql.Config:
            DSN:    ...
            InnoDB: ...
            Status:
                bytes_sent:      counter
                bytes_received:  counter
                threads_running: guage
```

## Service Config Structures

### agent
```json
{
  "ApiKey":      "123",
  "ApiHostname": "https://cloud-api-v2.percona.com",
  "AgentUuid" :  "abc-123-def",
  "PidFile":     ""
}
```

### log
```json
{
  "File":  "/var/log/percona/agent.log",
  "Level": "warn"
}
```

### data
```json
{
  "Dir":          "/var/spool/percona",
  "Encoding":     "gzip",
  "SendInterval": 60
}
```

### qan
```json
{
    "DSN":               "user:pass@(tcp:localhost:3360)/",
    "Start":             [<queries>],
    "Stop":              [<queries>],
    "MaxWorkers":        2,
    "ReportInterval":    300,
    "MaxSlowLogSize":    1073741824,
    "RemoveOldSlowLogs": true,
    "ExampleQueries":    true,
    "WorkerRunTime":     500
}
```

### mm
```json
{
  "Name":            "db1",
  "Type":            "mysql",
  "Config":          "<bytes>",
  "CollectInterval": 1,
  "ReportInterval":  60
}
```

#### mysql
```json
{
    "DSN":               "user:pass@(tcp)/",
    "InstanceName":      "",                  // optional name of MySQL instance
    "InnoDB":            "...",               // SET GLOBAL innodb_monitor_enable=<value> 
    "UserStats":         true,                // SET GLOBAL userstat=ON|OFF
    "UserStatsIgnoreDb": "",                  // dbs to ignore in user stats
    "Status":            {                    // SHOW STATUS variables to collect, lowercase
      "bytes_sent":      "counter",
      "bytes_received":  "counter",
      "threads_running": "guage"
    }
}
