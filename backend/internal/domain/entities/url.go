package entities

// URLRequest represents the input for URL processing
type URLRequest struct {
	URL       string `json:"url"`
	Operation string `json:"operation"`
}

// URLResponse represents the output for URL processing
type URLResponse struct {
	ProcessedURL string `json:"processed_url"`
}

// OperationType represents the type of URL processing operation
type OperationType string

const (
	OperationCanonical   OperationType = "canonical"
	OperationRedirection OperationType = "redirection"
	OperationAll         OperationType = "all"
)
