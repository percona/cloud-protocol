package proto

type ServiceInstance struct {
	Service  string // tool name
	UUID     string // instance UUID.
	Instance []byte `json:",omitempty"` // an *Instance structure
}
