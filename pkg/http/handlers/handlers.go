package handlers

import "errors"

type ErrorResponse struct {
	Message interface{} `json:"message,omitempty"`
}

var ErrInvalidPayload = errors.New("Invalid Payload")
