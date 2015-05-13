package proto

type Subsystem struct {
	ID       uint `json:"-"`
	ParentID uint `json:"-"`
	Type     string
	Prefix   string
}
