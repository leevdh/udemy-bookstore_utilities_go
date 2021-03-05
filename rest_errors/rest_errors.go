package rest_errors

import (
	"fmt"
	"net/http"
)

type RestErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

type restErr struct {
	errMessage string        `json:"message"`
	errStatus  int           `json:"status"`
	errError   string        `json:"error"`
	errCauses  []interface{} `json:"causes"`
}

func (e restErr) Error() string {
	return fmt.Sprintf("errMessage: %s - errStatus: %d - errError: %s - errCauses: [ %v ]", e.errMessage, e.errStatus, e.errError, e.errCauses)
}

func (e restErr) Message() string {
	return e.errMessage
}

func (e restErr) Status() int {
	return e.errStatus
}

func (e restErr) Causes() []interface{} {
	return e.errCauses
}

func NewRestError(message string, status int, err string, causes []interface{}) restErr {
	return restErr{
		errMessage: message,
		errStatus:  status,
		errError:   err,
		errCauses:  causes,
	}
}

func NewBadRequestError(message string) restErr {
	return restErr{
		errMessage: message,
		errStatus:  http.StatusBadRequest,
		errError:   "bad_request",
	}
}

func NewNotFoundError(message string) restErr {
	return restErr{
		errMessage: message,
		errStatus:  http.StatusNotFound,
		errError:   "not_found",
	}
}

func NewUnauthorizedError(message string) restErr {
	return restErr{
		errMessage: message,
		errStatus:  http.StatusUnauthorized,
		errError:   "unauthorized",
	}
}

func NewInternalServiceError(message string, err error) restErr {
	result := restErr{
		errMessage: message,
		errStatus:  http.StatusInternalServerError,
		errError:   "internal_server_error",
	}

	if err != nil {
		result.errCauses = append(result.errCauses, err.Error())
	}

	return result
}
