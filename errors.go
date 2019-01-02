package checkr

import "errors"

var (
	// ErrInvalidSignature ...
	ErrInvalidSignature = errors.New("Checkr Signature Mismatch Error")
	// ErrInvalidRequestInput ...
	ErrInvalidRequestInput = errors.New("Verify Input matches required parameters")
)
