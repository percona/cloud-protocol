package proto

import "time"

type QueryClass struct {
	QueryClassID    uint `json:"-"`
	Checksum        string
	Distillate      string
	Fingerprint     string
	ShowCreateTable string
	ShowTableStatus string
	FirstSeen       time.Time
	LastSeen        time.Time
	Status          string
}
