package proto

type Version struct {
	PerconaAgent string
	MySQL        string
}

type Configuration struct {
	Agent string
	Log   string
	Data  string
}

type AgentData struct {
	Name string
	Config Configuration
}

type CreateAgentData struct {
	Hostname string
	Versions Version
	Configs []AgentData
}
