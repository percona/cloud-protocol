package proto

type MmEntry struct {
	Agent string
	Meta  string
	Data  string
}

type MmAgent struct {
	Versions map[string]string `json:"versions"`
	Hostname string            `json:"hostname"`
	Alias    string            `json:"alias"`
	Uuid     string            `json:"uuid"`
}
