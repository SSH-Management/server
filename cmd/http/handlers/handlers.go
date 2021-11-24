package handlers

import (
	"errors"
)

type ErrorResponse struct {
	Message string `json:"message,omitempty"`
}

var ErrInvalidUrl = errors.New("invalid redirect URL")
