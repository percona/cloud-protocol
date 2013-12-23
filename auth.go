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

// Combined set of fields to be able to distinguish Agent or User connections, e.g. for "agent log" connection
type AgentOrUserAuth struct {
	ApiKey		string
	Uuid		string
	Hostname	string
	Username	string
	Email		string
	Provider	string
}

func (au *AgentOrUserAuth) AgentAuth() *AgentAuth {
	return &AgentAuth{ApiKey: au.ApiKey, Uuid: au.Uuid, Hostname: au.Hostname, Username: au.Username}
}

func (au *AgentOrUserAuth) UserAuth() *UserAuth {
	return &UserAuth{Email: au.Email, Provider: au.Provider, ApiKey: au.ApiKey}
}

func (au *AgentOrUserAuth) ContainsAgentCredentials() bool {
	return au.ApiKey != "" && au.Uuid != ""
}

func (au *AgentOrUserAuth) ContainsUserCredentials() bool {
	return au.Email != "" && au.Provider != "" && au.ApiKey != ""
}

type AuthResponse struct {
	Error string
}
