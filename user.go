package fnerror

import (
	"net/http"
	"runtime"
)

//NewNotFoundError return a NotFoundError that has error message msg and child error err.
func NewNotFoundError(msg string, err Apperror) *NotFoundError {
	return &NotFoundError{msg: msg, err: err, frame: getFrame()}
}

//NotFoundError indicates that the requested resource could not be found. | http: 404
type NotFoundError struct {
	msg   string
	err   Apperror
	frame runtime.Frame
}

func (e *NotFoundError) Error() string {
	if e.msg == "" {
		return http.StatusText(e.GetHTTPCode())
	}
	return e.msg
}

func (NotFoundError) GetCode() int {
	return StatusNotFound
}

func (NotFoundError) GetHTTPCode() int {
	return http.StatusNotFound
}

func (e *NotFoundError) GetTrace() string {
	return getTrace(e.Error(), e.frame)
}

func (e *NotFoundError) Unwrap() error {
	return e.err
}

//NewPayloadTooLargeError return a PayloadTooLargeError that has error message msg and child error err.
func NewPayloadTooLargeError(msg string, err Apperror) *PayloadTooLargeError {
	return &PayloadTooLargeError{msg: msg, err: err, frame: getFrame()}
}

//PayloadTooLargeError indicates that the payload's size exceed the limit of the server. | http: 413
type PayloadTooLargeError struct {
	msg   string
	err   Apperror
	frame runtime.Frame
}

func (e *PayloadTooLargeError) Error() string {
	if e.msg == "" {
		return http.StatusText(e.GetHTTPCode())
	}
	return e.msg
}

func (PayloadTooLargeError) GetCode() int {
	return StatusPayloadTooLarge
}

func (PayloadTooLargeError) GetHTTPCode() int {
	return http.StatusRequestEntityTooLarge
}

func (e *PayloadTooLargeError) GetTrace() string {
	return getTrace(e.Error(), e.frame)
}

func (e *PayloadTooLargeError) Unwrap() error {
	return e.err
}

//NewTimoutError return a TimeoutError that has error message msg and child error err.
func NewTimoutError(msg string, err Apperror) *TimeoutError {
	return &TimeoutError{msg: msg, err: err, frame: getFrame()}
}

//TimeoutError indicates that the process did not receive a response in time.
type TimeoutError struct {
	msg   string
	err   Apperror
	frame runtime.Frame
}

func (e *TimeoutError) Error() string {
	if e.msg == "" {
		return http.StatusText(e.GetHTTPCode())
	}
	return e.msg
}

func (TimeoutError) GetCode() int {
	return StatusTimeout
}

func (TimeoutError) GetHTTPCode() int {
	return http.StatusRequestTimeout
}

func (e *TimeoutError) GetTrace() string {
	return getTrace(e.Error(), e.frame)
}

func (e *TimeoutError) Unwrap() error {
	return e.err
}

//NewValidationError return a ValidationError that has error message msg and child error err.
func NewValidationError(msg string, err Apperror) *ValidationError {
	return &ValidationError{msg: msg, err: err, frame: getFrame()}
}

//ValidationError indicates that the request could not be understood by the server due to malformed syntax.
type ValidationError struct {
	msg   string
	err   Apperror
	frame runtime.Frame
}

func (e *ValidationError) Error() string {
	if e.msg == "" {
		return "Validation Failed"
	}
	return e.msg
}

func (ValidationError) GetCode() int {
	return StatusValidation
}

func (ValidationError) GetHTTPCode() int {
	return http.StatusBadRequest
}

func (e *ValidationError) GetTrace() string {
	return getTrace(e.Error(), e.frame)
}

func (e *ValidationError) Unwrap() error {
	return e.err
}

//NewMissingRequiredParamsError return a MissingRequiredParamsError that has error message msg and child error err.
func NewMissingRequiredParamsError(msg string, err Apperror) *MissingRequiredParamsError {
	return &MissingRequiredParamsError{msg: msg, err: err, frame: getFrame()}
}

//MissingRequiredParamsError is used when the request doesn't provide sufficient parameters to the server.
type MissingRequiredParamsError struct {
	msg   string
	err   Apperror
	frame runtime.Frame
}

func (e *MissingRequiredParamsError) Error() string {
	if e.msg == "" {
		return http.StatusText(e.GetHTTPCode())
	}
	return e.msg
}

func (MissingRequiredParamsError) GetCode() int {
	return StatusMissingRequiredParameter
}

func (MissingRequiredParamsError) GetHTTPCode() int {
	return http.StatusBadRequest
}

func (e *MissingRequiredParamsError) GetTrace() string {
	return getTrace(e.Error(), e.frame)
}

func (e *MissingRequiredParamsError) Unwrap() error {
	return e.err
}

//NewUnprocessableEntityError return an UnprocessableEntityError that has error message msg and child error err.
func NewUnprocessableEntityError(msg string, err Apperror) *UnprocessableEntityError {
	return &UnprocessableEntityError{msg: msg, err: err, frame: getFrame()}
}

//UnprocessableEntityError indicates that the request could be understood by the server(syntactically correct), but somehow could not be able to proccess the request.
type UnprocessableEntityError struct {
	msg   string
	err   Apperror
	frame runtime.Frame
}

func (e *UnprocessableEntityError) Error() string {
	if e.msg == "" {
		return http.StatusText(e.GetHTTPCode())
	}
	return e.msg
}

func (UnprocessableEntityError) GetCode() int {
	return StatusUnprocessableEntity
}

func (UnprocessableEntityError) GetHTTPCode() int {
	return http.StatusUnprocessableEntity
}

func (e *UnprocessableEntityError) GetTrace() string {
	return getTrace(e.Error(), e.frame)
}

func (e *UnprocessableEntityError) Unwrap() error {
	return e.err
}

//NewTooManyRequestError return a TooManyRequestError that has error message msg and child error err.
func NewTooManyRequestError(msg string, err Apperror) *TooManyRequestError {
	return &TooManyRequestError{msg: msg, err: err, frame: getFrame()}
}

//TooManyRequestError indicates that the user has sent too many request in a given amount of time
type TooManyRequestError struct {
	msg   string
	err   Apperror
	frame runtime.Frame
}

func (e *TooManyRequestError) Error() string {
	if e.msg == "" {
		return http.StatusText(e.GetHTTPCode())
	}
	return e.msg
}

func (TooManyRequestError) GetCode() int {
	return StatusTooManyRequest
}

func (TooManyRequestError) GetHTTPCode() int {
	return http.StatusTooManyRequests
}

func (e *TooManyRequestError) GetTrace() string {
	return getTrace(e.Error(), e.frame)
}

func (e *TooManyRequestError) Unwrap() error {
	return e.err
}

//NewBadRequestError return a BadRequestError that has error message msg and child error err.
func NewBadRequestError(msg string, err Apperror) *BadRequestError {
	return &BadRequestError{msg: msg, err: err, frame: getFrame()}
}

//BadRequestError indicates that the server cannot or will not process the request due to something that is perceived
type BadRequestError struct {
	msg   string
	err   Apperror
	frame runtime.Frame
}

func (e *BadRequestError) Error() string {
	if e.msg == "" {
		return http.StatusText(e.GetHTTPCode())
	}
	return e.msg
}

func (BadRequestError) GetCode() int {
	return StatusBadRequest
}

func (BadRequestError) GetHTTPCode() int {
	return http.StatusBadRequest
}

func (e *BadRequestError) GetTrace() string {
	return getTrace(e.Error(), e.frame)
}

func (e *BadRequestError) Unwrap() error {
	return e.err
}
