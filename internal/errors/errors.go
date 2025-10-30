package errors

import "errors"

var (
	ErrorNotYetImplemented = errors.New("this feature is not yet implemented")
	ErrorInvalidDefinition = errors.New("encountered one or more formatting errors")
)
