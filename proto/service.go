package proto

// Service numbers are only used internally by API.
var ExternalService map[string]uint = map[string]uint{
	"server": 1,
	"mysql":  2,
}

// GET /instances returns a list of these, one for each unique hostname:
type UniqueInstance struct {
	Hostname string // Real or virtual server, Amazon RDS instance, etc.
	ServerId uint   // ServerInstance.Id
	MySQLId  uint   // MySQLInstance.Id
}

type ServiceInstance struct {
	Service    string // internal (agent, log, data) or and ExternalService
	InstanceId uint   // external service instance ID: server-5, mysql-9, etc.
	Instance   []byte `json:",omitempty"` // an *Instance structure
}

type Settings struct {
	AgentUuid string // @todo @obsolete after deploying https://jira.percona.com/browse/PCT-731
	Agent     *Agent
	Running   bool
}

type Monitors struct {
	Metrics *Settings `json:",omitempty"`
	Config  *Settings `json:",omitempty"`
}

type ServerInstance struct {
	Id       uint      // set by API
	Hostname string    `json:",omitempty"`
	Agent    *Agent    `json:",omitempty"`
	Monitors *Monitors `json:",omitempty"`
}

type MySQLInstance struct {
	Id       uint            // set by API
	Hostname string          // @@hostname[.port] if port != 3306
	DSN      string          // [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	Distro   string          `json:",omitempty"` // MySQL Community Edition, Percona Server, etc.
	Version  string          `json:",omitempty"`
	Server   *ServerInstance `json:",omitempty"`
	Monitors *Monitors       `json:",omitempty"`
}
