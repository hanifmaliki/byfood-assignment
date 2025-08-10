package handlers

// ErrorResponse represents a standard error payload
// swagger:model ErrorResponse
type ErrorResponse struct {
	// Error message
	// example: invalid request payload
	Error string `json:"error"`
}

// MessageResponse represents a standard message payload
// swagger:model MessageResponse
type MessageResponse struct {
	// Informational message
	// example: operation completed successfully
	Message string `json:"message"`
}
