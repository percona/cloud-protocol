package proto

type AgentData struct {
	Hostname string
	Versions map[string]string
	Configs map[string]string
}
