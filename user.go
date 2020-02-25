package fnerror

import (
	"net/http"
)

//NewNotFoundError return a NotFoundError that has error message msg and child error err.
func NewNotFoundError(msg string, err Apperror) *NotFoundError {
	fne := newFnerror(msg, err, http.StatusNotFound, StatusNotFound)
	return &NotFoundError{fne}
}

//NewPayloadTooLargeError return a PayloadTooLargeError that has error message msg and child error err.
func NewPayloadTooLargeError(msg string, err Apperror) *PayloadTooLargeError {
	fne := newFnerror(msg, err, http.StatusRequestEntityTooLarge, StatusPayloadTooLarge)
	return &PayloadTooLargeError{fne}
}

//NewTimoutError return a TimeoutError that has error message msg and child error err.
func NewTimoutError(msg string, err Apperror) *TimeoutError {
	fne := newFnerror(msg, err, http.StatusRequestTimeout, StatusTimeout)
	return &TimeoutError{fne}
}

//NewValidationError return a ValidationError that has error message msg and child error err.
func NewValidationError(msg string, err Apperror) *ValidationError {
	fne := newFnerror(msg, err, http.StatusBadRequest, StatusValidation)
	return &ValidationError{fne}
}

//NewMissingRequiredParamsError return a MissingRequiredParamsError that has error message msg and child error err.
func NewMissingRequiredParamsError(msg string, err Apperror) *MissingRequiredParamsError {
	fne := newFnerror(msg, err, http.StatusBadRequest, StatusMissingRequiredParameter)
	return &MissingRequiredParamsError{fne}
}

//NewUnprocessableEntityError return an UnprocessableEntityError that has error message msg and child error err.
func NewUnprocessableEntityError(msg string, err Apperror) *UnprocessableEntityError {
	fne := newFnerror(msg, err, http.StatusUnprocessableEntity, StatusUnprocessableEntity)
	return &UnprocessableEntityError{fne}
}

//NewTooManyRequestError return a TooManyRequestError that has error message msg and child error err.
func NewTooManyRequestError(msg string, err Apperror) *TooManyRequestError {
	fne := newFnerror(msg, err, http.StatusTooManyRequests, StatusTooManyRequests)
	return &TooManyRequestError{fne}
}

//NewBadRequestError return a BadRequestError that has error message msg and child error err.
func NewBadRequestError(msg string, err Apperror) *BadRequestError {
	fne := newFnerror(msg, err, http.StatusBadRequest, StatusBadRequest)
	return &BadRequestError{fne}
}

//NotFoundError indicates that the requested resource could not be found. | http: 404
type NotFoundError struct {
	*finerror
}

//PayloadTooLargeError indicates that the payload's size exceed the limit of the server. | http: 413
type PayloadTooLargeError struct {
	*finerror
}

//TimeoutError indicates that the process did not receive a response in time.
type TimeoutError struct {
	*finerror
}

//ValidationError indicates that the request could not be understood by the server due to malformed syntax.
type ValidationError struct {
	*finerror
}

//MissingRequiredParamsError is used when the request doesn't provide sufficient parameters to the server.
type MissingRequiredParamsError struct {
	*finerror
}

//UnprocessableEntityError indicates that the request could be understood by the server(syntactically correct), but somehow could not be able to proccess the request.
type UnprocessableEntityError struct {
	*finerror
}

//TooManyRequestError indicates that the user has sent too many request in a given amount of time
type TooManyRequestError struct {
	*finerror
}

//BadRequestError indicates that the server cannot or will not process the request due to something that is perceived
type BadRequestError struct {
	*finerror
}
