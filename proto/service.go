package proto

// Service numbers are only used internally by API.
var ExternalService map[string]uint = map[string]uint{
	"server": 1,
	"mysql":  2,
}

type ServiceInstance struct {
	Service    string // one of ExternalService
	InstanceId uint   // unique for Service: mysql-1, mysql-2, memcached-1, etc.c
	Instance   []byte `json:",omitempty"` // one of the structures below
}

type ServerInstance struct {
	Id       uint // set by API
	Hostname string
}

type MySQLInstance struct {
	Id      uint //set by API
	Name    string
	DSN     string // [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	Distro  string `json:",omitempty"`
	Version string `json:",omitempty"`
}

