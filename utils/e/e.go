package e

import "errors"

var (
	ErrNoIdHeader       = errors.New("yo'qol bu yerdan!")
	ErrContentsRequired = errors.New("bo'm-bo'shku!")
)
