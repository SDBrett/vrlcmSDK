package types

type ErrorResponse struct {
	// The error message.
	// Required: true
	Message string `json:"message"`
}
