package clients

import "errors"

var (
	ErrClientNotFound = errors.New("client not found")
	ErrClientNotLimit = errors.New("client not have limit")
)
