package proto

type ToolInstance struct {
	Tool     string // tool name
	UUID     string // instance UUID.
	Instance []byte `json:",omitempty"` // an *Instance structure
}
