package fnerror

import (
	"net/http"
	"runtime"
)

func NewFnerror(msg string, err error, httpCode int, errCode string) *FinnoError {

	return &FinnoError{
		msg:      msg,
		err:      err,
		frame:    getFrame(),
		errCode:  errCode,
		httpCode: httpCode,
	}
}

type FinnoError struct {
	msg      string
	err      error
	frame    runtime.Frame
	errCode  string
	httpCode int
}

func (e *FinnoError) Error() string {
	if e.msg == "" {
		return http.StatusText(e.GetHTTPCode())
	}
	return e.msg
}

//GetCode return the application error code.
func (e *FinnoError) GetCode() string {
	return e.errCode
}

//GetHTTPCode return http response status code which related to the error.
func (e *FinnoError) GetHTTPCode() int {
	return e.httpCode
}

func (e *FinnoError) GetTrace() string {
	return getTrace(e.Error(), e.frame)
}

//Unwrap return the child error of the error, mostly used for extracting errors chain.
func (e *FinnoError) Unwrap() error {
	return e.err
}
