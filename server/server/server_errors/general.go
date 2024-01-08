package server_errors

import "errors"

var (
	EntityNotFound        = errors.New("EntityNotFound")
	NotProcessiableEntity = errors.New("NotProcessiableEntity")
)
