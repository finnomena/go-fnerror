package fnerror

import (
	"net/http"
	"runtime"
)

//NewUnauthorizedError return an UnauthorizedError that has error message msg and child error err.
func NewUnauthorizedError(msg string, err Apperror) *UnauthorizedError {
	return &UnauthorizedError{msg: msg, err: err, frame: getFrame()}
}

//UnauthorizedError indicates that the user has not been yet authenticated.
type UnauthorizedError struct {
	msg   string
	err   Apperror
	frame runtime.Frame
}

//Error return error message as a string
func (e *UnauthorizedError) Error() string {
	if e.msg == "" {
		return http.StatusText(e.GetHTTPCode())
	}
	return e.msg
}

//GetCode return the application error code.
func (UnauthorizedError) GetCode() int {
	return StatusUnauthorized
}

//GetHTTPCode return http response status code which related to the error.
func (UnauthorizedError) GetHTTPCode() int {
	return http.StatusUnauthorized
}

func (e *UnauthorizedError) GetTrace() string {
	return getTrace(e.Error(), e.frame)
}

//Unwrap return the child error of the error, mostly used for extracting errors chain.
func (e *UnauthorizedError) Unwrap() error {
	return e.err
}

//NewWrongCredentialError return a WrongCredentialError that has error message msg and child error err.
func NewWrongCredentialError(msg string, err Apperror) *WrongCredentialError {
	return &WrongCredentialError{msg: msg, err: err, frame: getFrame()}
}

//WrongCredentialError indicates that the credentials (Ex. username or password) provided by the user is not correct.
type WrongCredentialError struct {
	msg   string
	err   Apperror
	frame runtime.Frame
}

func (e *WrongCredentialError) Error() string {
	if e.msg == "" {
		return http.StatusText(e.GetHTTPCode())
	}
	return e.msg
}

func (WrongCredentialError) GetCode() int {
	return StatusWrongCredential
}

func (WrongCredentialError) GetHTTPCode() int {
	return http.StatusUnauthorized
}

func (e *WrongCredentialError) GetTrace() string {
	return getTrace(e.Error(), e.frame)
}

func (e *WrongCredentialError) Unwrap() error {
	return e.err
}

//NewForbiddenError return a ForbiddenError that has error message msg and child error err
func NewForbiddenError(msg string, err Apperror) *ForbiddenError {
	return &ForbiddenError{msg: msg, err: err, frame: getFrame()}
}

//ForbiddenError indicates that the user has no permission to perform the specific action.
type ForbiddenError struct {
	msg   string
	err   Apperror
	frame runtime.Frame
}

func (e *ForbiddenError) Error() string {
	if e.msg == "" {
		return http.StatusText(e.GetHTTPCode())
	}
	return e.msg
}

func (ForbiddenError) GetCode() int {
	return StatusForbidden
}

func (ForbiddenError) GetHTTPCode() int {
	return http.StatusForbidden
}

func (e *ForbiddenError) GetTrace() string {
	return getTrace(e.Error(), e.frame)
}

func (e *ForbiddenError) Unwrap() error {
	return e.err
}
