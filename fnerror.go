package fnerror

import (
	"fmt"
	"runtime"
)

//Apperror interface which is used as a standard error across entire application.
type Apperror interface {
	Error() string
	//GetCode return the application error code.
	GetCode() int
	//GetHTTPCode return http response status code which related to the error.
	GetHTTPCode() int
	//GetTrace return
	GetTrace() string
	//Unwrap return the child error of the error, mostly used for extracting errors chain.
	Unwrap() error
}

//Client Error Code
const (
	StatusBadRequest = iota + 100
	StatusNotFound
	StatusMissingRequiredParameter
	StatusValidation
	StatusTimeout
	StatusUnprocessableEntity
	StatusTooManyRequests
	StatusPayloadTooLarge
)

//Server Error Code
const (
	StatusServiceUnavailable = iota + 200
	StatusUnexpected         = 999
)

//Authorization Error Code
const (
	StatusUnauthorized = iota + 300
	StatusWrongCredential
	StatusForbidden
)

func getFrame() runtime.Frame {
	stackBuf := make([]uintptr, 50)
	length := runtime.Callers(3, stackBuf[:])
	stack := stackBuf[:length]
	frames := runtime.CallersFrames(stack)
	frame, _ := frames.Next()
	return frame
}

func getTrace(msg string, frame runtime.Frame) string {
	trace := fmt.Sprintf("\n%s() \"%s\"\n%s:%d", frame.Function, msg, frame.File, frame.Line)
	return trace
}

//GetChainTrace Get Errors trace
func GetChainTrace(e Apperror) string {
	trace := e.GetTrace()
	for childErr := e.Unwrap(); childErr != nil; childErr = childErr.(Apperror).Unwrap() {
		trace = trace + "\n" + childErr.(Apperror).GetTrace()
	}
	return trace
}
