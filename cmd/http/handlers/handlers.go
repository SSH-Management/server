package handlers

import (
	"errors"
	"net/url"
)

type ErrorResponse struct {
	Message string `json:"message,omitempty"`
}

var ErrInvalidUrl = errors.New("invalid redirect URL")

func Redirect(redirect string, message ...string) (string, error) {
	uri, err := url.Parse(redirect)

	if err != nil {
		return "", ErrInvalidUrl
	}

	query := uri.Query()

	query.Add("status", "success")

	if len(message) > 0 {
		query.Add("message", message[0])
	} else {
		query.Add("message", "Message sent")
	}

	uri.RawQuery = query.Encode()

	return uri.String(), nil
}
