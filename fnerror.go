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
	trace := fmt.Sprintf(
		"Message: \"%s\"\nFrom: %s()\n --- at %s:%d",
		msg, frame.Function, frame.File, frame.Line)
	return trace
}

//GetChainTrace Get Errors trace
func GetChainTrace(e Apperror) string {
	chainTrace := ""
	cause := ""
	trace := e.GetTrace()
	err := e.Unwrap()
	for {
		if err == nil {
			break
		}
		apperr, ok := err.(Apperror)
		if ok {
			trace += "\n" + apperr.GetTrace()
		} else {
			cause += err.Error()
			break
		}
		err = apperr.Unwrap()
	}

	if cause == "" {
		chainTrace = trace
	} else {
		cause = fmt.Sprintf("Original Error: \"%s\"", cause)
		chainTrace = fmt.Sprintf("%s\n%s", trace, cause)
	}
	return chainTrace
}
