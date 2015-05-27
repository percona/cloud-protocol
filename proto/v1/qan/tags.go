package qan

type Tag struct {
	ID    uint64 `json:"Id"`
	Name  string
	Color string
}

type ClassTag struct {
	TagId uint
	Name  string
	Color string
}
