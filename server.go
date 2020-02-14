package fnerror

import (
	"net/http"
	"runtime"
)

//NewUnexpectedError return a UnexpectedError that has error message msg and child error err.
//Noted that msg should be empty string("") as it is not the best practice to expose the unhandled message to the caller.
func NewUnexpectedError(msg string, err Apperror) *UnexpectedError {
	return &UnexpectedError{msg: msg, err: err, frame: getFrame()}
}

//UnexpectedError is used when something is not expected or unhandled by the server.
type UnexpectedError struct {
	msg   string
	err   Apperror
	frame runtime.Frame
}

func (e *UnexpectedError) Error() string {
	if e.msg == "" {
		return http.StatusText(e.GetHTTPCode())
	}
	return e.msg
}

func (UnexpectedError) GetCode() int {
	return StatusUnexpected
}

func (UnexpectedError) GetHTTPCode() int {
	return http.StatusInternalServerError
}

func (e *UnexpectedError) GetTrace() string {
	return getTrace(e.Error(), e.frame)
}

func (e *UnexpectedError) Unwrap() error {
	return e.err
}

//NewServiceUnavailableError return a ServiceUnavailableError that has error message msg and child error err.
func NewServiceUnavailableError(msg string, err Apperror) *ServiceUnavailableError {
	return &ServiceUnavailableError{msg: msg, err: err, frame: getFrame()}
}

//ServiceUnavailableError indicates the the service is not ready to handle the request or it is down for maintenance or overloaded.
type ServiceUnavailableError struct {
	msg   string
	err   Apperror
	frame runtime.Frame
}

func (e *ServiceUnavailableError) Error() string {
	if e.msg == "" {
		return http.StatusText(e.GetHTTPCode())
	}
	return e.msg
}

func (ServiceUnavailableError) GetCode() int {
	return StatusServiceUnavailable
}

func (ServiceUnavailableError) GetHTTPCode() int {
	return http.StatusServiceUnavailable
}

func (e *ServiceUnavailableError) GetTrace() string {
	return getTrace(e.Error(), e.frame)
}

func (e *ServiceUnavailableError) Unwrap() error {
	return e.err
}
