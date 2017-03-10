package usuario

// Request -> struct
type Request struct{}

// Response -> struct
type Response struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}
