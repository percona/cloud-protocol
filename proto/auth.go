package proto

type UserAuth struct {
	ApiKey string
}

type AuthResponse struct {
	Code  uint   // standard HTTP status (http://httpstatus.es/)
	Error string // empty if auth ok (Code=200)
}
