package proto

type QanEntry struct {
	Agent string
	Meta  string
	Data  string
}

type QanAgent struct {
	Versions map[string]string `json:"versions"`
	Hostname string		   `json:"hostname"`
	Alias	 string		   `json:"alias"`
	Uuid	 string		   `json:"uuid"`
}
