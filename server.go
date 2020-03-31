package fnerror

import (
	"net/http"
)

//NewUnexpectedError return a UnexpectedError that has error message msg and child error err.
//Noted that msg should be empty string("") as it is not the best practice to expose the unhandled message to the caller.
func NewUnexpectedError(msg string, err error) *UnexpectedError {
	fne := NewFnerror(msg, err, http.StatusInternalServerError, StatusUnexpected)
	return &UnexpectedError{fne}
}

//NewServiceUnavailableError return a ServiceUnavailableError that has error message msg and child error err.
func NewServiceUnavailableError(msg string, err error) *ServiceUnavailableError {
	fne := NewFnerror(msg, err, http.StatusServiceUnavailable, StatusServiceUnavailable)
	return &ServiceUnavailableError{fne}
}

//UnexpectedError is used when something is not expected or unhandled by the server.
type UnexpectedError struct {
	*FinnoError
}

//ServiceUnavailableError indicates the the service is not ready to handle the request or it is down for maintenance or overloaded.
type ServiceUnavailableError struct {
	*FinnoError
}
