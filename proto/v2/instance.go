package proto

import "time"

// Keep the ID and ParentID fields because they are being used
// by the subsystems cache
type Subsystem struct {
	ID       uint `json:"-"`
	ParentID uint `json:"-"`
	Type     string
	Prefix   string
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
	// --
	V2 interface{}
}

type SystemTreeSync struct {
	Added   []string // UUID of added instances
	Deleted []string // UUID of deleted instances
	Updated []string // UUID of updated instances
	Version uint     // Version of instance tree
	Tree    Instance // Full instance tree
}
