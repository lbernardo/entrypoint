package entrypoint

type Request struct {
	Headers map[string]string `json:"header"`
	Body    string            `json:"body"`
	Vars    map[string]string `json:"vars"`
}

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message,omitempty"`
}
