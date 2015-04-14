package proto

import "time"

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
