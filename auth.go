package proto

type AgentAuth struct {
	ApiKey		string
	Uuid		string
	Hostname	string
	Username	string
}

type UserAuth struct {
	Email		string
	Provider	string
	ApiKey		string
}

type AuthResponse struct {
	Error string
}
