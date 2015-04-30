package proto

import "time"

type Subsystem struct {
	Type   string
	Prefix string
}

type Instance struct {
	Subsystem
	ParentUUID string
	UUID       string
	Name       string
	DSN        string
	Created    time.Time
	Deleted    time.Time
	Properties map[string]string
	Subsystems []Instance
}

type SystemTreeSync struct {
	Added   []string // UUID of added instances
	Deleted []string // UUID of deleted instances
	Updated []string // UUID of updated instances
	Version uint     // Version of instance tree
	Tree    Instance // Full instance tree
}
