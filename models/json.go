package models

type ResponseJSON struct {
	Code    int         `json:"code"`
	Status  string      `json:"status,omitempty"`
	Error   string      `json:"error,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}
