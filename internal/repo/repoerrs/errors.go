package repoerrs

import "errors"

var (
	ErrNotFound      = errors.New("не нашли")
	ErrAlreadyExists = errors.New("уже существует")
)
