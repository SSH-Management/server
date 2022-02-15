package handlers

type ErrorResponse struct {
	Message interface{} `json:"message,omitempty"`
}
