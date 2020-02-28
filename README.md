# FINNOMENA ERROR
[![MIT License](https://img.shields.io/apm/l/atomic-design-ui.svg?)](https://github.com/finnomena/go-fnerror/blob/master/LICENSE) [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/finnomena/go-fnerror?tab=doc)

FINNOMENA Error is an extension of standard golang error which provides more flexibility to use with additional information about an error (Ex. stack tracing)

### Prerequisites

Golang 1.13 or further version must be installed, in order to use the standard library functions such as [errors.Is()](https://pkg.go.dev/errors?tab=doc#Is "errors.Is()") or [errors.As()](https://pkg.go.dev/errors?tab=doc#As "errors.As()") to examine a child error in the error chain.

### Installation
```bash
go get github.com/finnomena/go-fnerror
```

### Example
```golang
package main

import (
	"errors"
	"fmt"

	"github.com/finnomena/go-fnerror"
)

func someFunction() fnerror.Apperror {
	return fnerror.NewUnexpectedError("CUSTOM MESSAGE", nil)
}

func main() {
	unexpectedErr := fnerror.NewUnexpectedError("", nil)
	notFoundErr := fnerror.NewNotFoundError("", unexpectedErr)
	fmt.Println(notFoundErr.Error())
	// Not Found
	fmt.Println(notFoundErr.GetCode())
	// 101 --> application error code
	fmt.Println(notFoundErr.GetHTTPCode())
	// 404 --> http response code
	fmt.Println(notFoundErr.GetTrace())
	// {$dir}/main.go:16
	fmt.Println(fnerror.GetChainTrace(notFoundErr))
	/*
		main.main() "Not Found"
		{$dir}/main.go:16

		main.main() "CUSTOM MESSAGE"
		{$dir}/main.go:15
	*/
	standardErr := errors.New("Error")
	unexpectedErr2 := new(fnerror.UnexpectedError) //allocate with new only for type assertion
	forbiddenErr := new(fnerror.ForbiddenError)

	fmt.Println(errors.Is(notFoundErr, standardErr))
	//false --> has no standard error (compare to a value)
	fmt.Println(errors.Is(notFoundErr, unexpectedErr))
	//true --> has unexpectedErr
	fmt.Println(errors.Is(notFoundErr, unexpectedErr2))
	//false --> doesn't has unexpectedErr2
	fmt.Println(errors.Is(notFoundErr, forbiddenErr))
	//false --> doesnt has forbiddenErr

	fmt.Println(errors.As(notFoundErr, &standardErr))
	//true --> contains error type (compare to a type)
	fmt.Println(errors.As(notFoundErr, &unexpectedErr2))
	//true --> contains UnexpectedError type
	fmt.Println(errors.As(notFoundErr, &forbiddenErr))
	//false --> contains ForbiddenError type
}

```

### Contributing
* Use PR for a small change (including test case and doc) 
* For a major change, open an issue for dicussion first.

### License

[MIT](https://github.com/finnomena/go-fnerror/blob/master/LICENSE)
