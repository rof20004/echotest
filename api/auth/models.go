package auth

// Request struct
type Request struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
