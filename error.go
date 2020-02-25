package fnerror

import (
	"net/http"
	"runtime"
)

func newFnerror(msg string, err Apperror, errCode, httpCode int) *finerror {
	return &finerror{
		msg:      msg,
		err:      err,
		frame:    getFrame(),
		errCode:  errCode,
		httpCode: httpCode,
	}
}

type finerror struct {
	msg      string
	err      Apperror
	frame    runtime.Frame
	errCode  int
	httpCode int
}

func (e *finerror) Error() string {
	if e.msg == "" {
		return http.StatusText(e.GetHTTPCode())
	}
	return e.msg
}

//GetCode return the application error code.
func (e *finerror) GetCode() int {
	return e.errCode
}

//GetHTTPCode return http response status code which related to the error.
func (e *finerror) GetHTTPCode() int {
	return e.httpCode
}

func (e *finerror) GetTrace() string {
	return getTrace(e.Error(), e.frame)
}

//Unwrap return the child error of the error, mostly used for extracting errors chain.
func (e *finerror) Unwrap() error {
	return e.err
}
