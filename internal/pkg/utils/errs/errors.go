package errs

import "errors"

var (
	// ErrIdRequired indicates that an ID was not provided
	// but is required to satisfy the request.
	ErrIdRequired = errors.New("id is required")
	// ErrIdNotFound indicates that the provided ID was
	// not found on the remote resource.
	ErrIdNotFound = errors.New("id not found")
	// ErrUnimplemented is a catch-all used as a placeholder
	// in this example test stack.
	ErrUnimplemented = errors.New("unimplemented")
)
