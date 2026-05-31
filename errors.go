package veego

import "errors"

var (
	// Returned by the controller if the API key passed into its initialization is
	// not a valid Govee API key
	ErrInvalidKey = errors.New("invalid Govee API key")

	// Returned if attempting to initialize an already-initialized Controller
	ErrInstanceAlreadyOpen = errors.New("instance already open")

	// Returned if attempting to use a Controller without it being initialized
	ErrInstanceNotOpen = errors.New("instance not active")
)
