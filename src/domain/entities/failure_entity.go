package entities

type FailureEntity struct {
	Error       *string `json:"error,omitempty"`
	Message     *string `json:"msg,omitempty"`
	Description *string `json:"error_description,omitempty"`
}
