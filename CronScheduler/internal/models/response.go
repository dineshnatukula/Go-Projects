package models

type Response struct {
	Status           string `json:"status,omitempty"`
	StatusMessage    string `json:"statusMessage,omitempty"`
	ErrorCode        string `json:"errorCode,omitempty"`
	ErrorDescription string `json:"errorDescription,omitempty"`
}
