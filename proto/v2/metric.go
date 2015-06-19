package proto

// Metrics V2
type Metric struct {
	ID          uint64 `json:"-"`
	SubsystemID uint
	Name        string
	Title       string
	Type        string
	Unit        string
}
