package server_errors

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrEntityNotFound        = errors.New("EntityNotFound")
	ErrNotProcessiableEntity = errors.New("NotProcessiableEntity")
)

type ServerError struct {
	Error error
	StatusCode int
}


func HandleServerError(c *fiber.Ctx, err error) ServerError {
	switch err {
	case ErrEntityNotFound:
		c.Status(http.StatusNotFound).SendString(ErrEntityNotFound.Error())
		return ServerError{Error: ErrEntityNotFound, StatusCode: http.StatusNotFound}
	case ErrNotProcessiableEntity:
		c.Status(http.StatusUnprocessableEntity).SendString(ErrNotProcessiableEntity.Error())
		return ServerError{Error: ErrNotProcessiableEntity, StatusCode: http.StatusUnprocessableEntity}
	}
	return ServerError{Error: nil, StatusCode: http.StatusOK}
}