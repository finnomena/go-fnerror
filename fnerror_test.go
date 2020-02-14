package fnerror_test

import (
	"errors"
	"testing"

	fne "github.com/finnomena/go-fnerror"
	"github.com/stretchr/testify/assert"
)

func getErrors() []fne.Apperror {
	var errs [13]fne.Apperror
	errs[0] = fne.NewBadRequestError("e0", nil)
	errs[1] = fne.NewForbiddenError("e1", errs[0])
	errs[2] = fne.NewMissingRequiredParamsError("e2", errs[1])
	errs[3] = fne.NewNotFoundError("e3", errs[2])
	errs[4] = fne.NewPayloadTooLargeError("e4", errs[3])
	errs[5] = fne.NewServiceUnavailableError("e5", errs[4])
	errs[6] = fne.NewTimoutError("e6", errs[5])
	errs[7] = fne.NewTooManyRequestError("e7", errs[6])
	errs[8] = fne.NewUnauthorizedError("e8", errs[7])
	errs[9] = fne.NewUnexpectedError("e9", errs[8])
	errs[10] = fne.NewValidationError("e10", errs[9])
	errs[11] = fne.NewUnprocessableEntityError("e11", errs[10])
	errs[12] = fne.NewWrongCredentialError("e12", errs[11])
	return errs[:]
}

func TestErrorsChain(t *testing.T) {
	errs := getErrors()

	var typeTests []bool
	for _, err := range errs {
		typeTests = append(typeTests, errors.As(errs[len(errs)-1], &err))
	}
	for _, test := range typeTests {
		assert.True(t, test)
	}

	var valueTests []bool
	for _, err := range errs {
		valueTests = append(valueTests, errors.Is(errs[len(errs)-1], err))
	}
	for _, test := range valueTests {
		assert.True(t, test)
	}
}

func TestNotZeroValue(t *testing.T) {
	errs := getErrors()
	for _, err := range errs {
		assert.NotEmpty(t, err.GetCode())
		assert.NotEmpty(t, err.GetHTTPCode())
		assert.NotEmpty(t, err.Error())
		assert.NotEmpty(t, err.GetTrace())
	}
}
