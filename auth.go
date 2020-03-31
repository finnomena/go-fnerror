package fnerror

import (
	"net/http"
)

//NewUnauthorizedError return an UnauthorizedError that has error message msg and child error err.
func NewUnauthorizedError(msg string, err error) *UnauthorizedError {
	fne := NewFnerror(msg, err, http.StatusUnauthorized, StatusUnauthorized)
	return &UnauthorizedError{fne}
}

//NewWrongCredentialError return a WrongCredentialError that has error message msg and child error err.
func NewWrongCredentialError(msg string, err error) *WrongCredentialError {
	fne := NewFnerror(msg, err, http.StatusUnauthorized, StatusWrongCredential)
	return &WrongCredentialError{fne}
}

//NewForbiddenError return a ForbiddenError that has error message msg and child error err
func NewForbiddenError(msg string, err error) *ForbiddenError {
	fne := NewFnerror(msg, err, http.StatusForbidden, StatusForbidden)
	return &ForbiddenError{fne}
}

//UnauthorizedError indicates that the user has not been yet authenticated.
type UnauthorizedError struct {
	*FinnoError
}

//WrongCredentialError indicates that the credentials (Ex. username or password) provided by the user is not correct.
type WrongCredentialError struct {
	*FinnoError
}

//ForbiddenError indicates that the user has no permission to perform the specific action.
type ForbiddenError struct {
	*FinnoError
}
