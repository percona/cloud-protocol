package proto

// Keep the ID and ParentID fields because they are being used
// by the subsystems cache
type Subsystem struct {
	ID       uint `json:"-"`
	ParentID uint `json:"-"`
	Type     string
	Prefix   string
}
