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
	Service    string // one of ExternalService
	InstanceId uint   // unique for Service: mysql-1, mysql-2, memcached-1, etc.c
	Instance   []byte `json:",omitempty"` // one of the structures below
}

type Agent struct {
	UUID string
	QAN  bool
}

type Settings struct {
	AgentUUID string
	Running   bool
}

type ServerInstance struct {
	Id       uint     // set by API
	Hostname string   `json:",omitempty"`
	Agent    Agent    `json:",omitempty"`
	Monitors Monitors `json:",omitempty"`
}

type Monitors struct {
	Metrics Settings
	Config  Settings `json:",omitempty"`
}

type MySQLInstance struct {
	Id       uint           // set by API
	Hostname string         // @@hostname[.port] if port != 3306
	DSN      string         // [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	Distro   string         `json:",omitempty"` // MySQL Community Edition, Percona Server, etc.
	Version  string         `json:",omitempty"`
	Server   ServerInstance `json:",omitempty"`
	Monitors Monitors       `json:",omitempty"`
}

