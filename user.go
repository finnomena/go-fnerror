package fnerror

import (
	"net/http"
)

//NewNotFoundError return a NotFoundError that has error message msg and child error err.
func NewNotFoundError(msg string, err error) *NotFoundError {
	fne := NewFnerror(msg, err, http.StatusNotFound, StatusNotFound)
	return &NotFoundError{fne}
}

//NewPayloadTooLargeError return a PayloadTooLargeError that has error message msg and child error err.
func NewPayloadTooLargeError(msg string, err error) *PayloadTooLargeError {
	fne := NewFnerror(msg, err, http.StatusRequestEntityTooLarge, StatusPayloadTooLarge)
	return &PayloadTooLargeError{fne}
}

//NewTimoutError return a TimeoutError that has error message msg and child error err.
func NewTimoutError(msg string, err error) *TimeoutError {
	fne := NewFnerror(msg, err, http.StatusRequestTimeout, StatusTimeout)
	return &TimeoutError{fne}
}

//NewValidationError return a ValidationError that has error message msg and child error err.
func NewValidationError(msg string, err error) *ValidationError {
	fne := NewFnerror(msg, err, http.StatusBadRequest, StatusValidation)
	return &ValidationError{fne}
}

//NewMissingRequiredParamsError return a MissingRequiredParamsError that has error message msg and child error err.
func NewMissingRequiredParamsError(msg string, err error) *MissingRequiredParamsError {
	fne := NewFnerror(msg, err, http.StatusBadRequest, StatusMissingRequiredParameter)
	return &MissingRequiredParamsError{fne}
}

//NewUnprocessableEntityError return an UnprocessableEntityError that has error message msg and child error err.
func NewUnprocessableEntityError(msg string, err error) *UnprocessableEntityError {
	fne := NewFnerror(msg, err, http.StatusUnprocessableEntity, StatusUnprocessableEntity)
	return &UnprocessableEntityError{fne}
}

//NewTooManyRequestError return a TooManyRequestError that has error message msg and child error err.
func NewTooManyRequestError(msg string, err error) *TooManyRequestError {
	fne := NewFnerror(msg, err, http.StatusTooManyRequests, StatusTooManyRequests)
	return &TooManyRequestError{fne}
}

//NewBadRequestError return a BadRequestError that has error message msg and child error err.
func NewBadRequestError(msg string, err error) *BadRequestError {
	fne := NewFnerror(msg, err, http.StatusBadRequest, StatusBadRequest)
	return &BadRequestError{fne}
}

//NotFoundError indicates that the requested resource could not be found. | http: 404
type NotFoundError struct {
	*FinnoError
}

//PayloadTooLargeError indicates that the payload's size exceed the limit of the server. | http: 413
type PayloadTooLargeError struct {
	*FinnoError
}

//TimeoutError indicates that the process did not receive a response in time.
type TimeoutError struct {
	*FinnoError
}

//ValidationError indicates that the request could not be understood by the server due to malformed syntax.
type ValidationError struct {
	*FinnoError
}

//MissingRequiredParamsError is used when the request doesn't provide sufficient parameters to the server.
type MissingRequiredParamsError struct {
	*FinnoError
}

//UnprocessableEntityError indicates that the request could be understood by the server(syntactically correct), but somehow could not be able to proccess the request.
type UnprocessableEntityError struct {
	*FinnoError
}

//TooManyRequestError indicates that the user has sent too many request in a given amount of time
type TooManyRequestError struct {
	*FinnoError
}

//BadRequestError indicates that the server cannot or will not process the request due to something that is perceived
type BadRequestError struct {
	*FinnoError
}
