package entities

type ErrorEntity struct {
	Error       *string `json:"error,omitempty"`
	Message     *string `json:"message,omitempty"`
	Description *string `json:"error_description,omitempty"`
}
