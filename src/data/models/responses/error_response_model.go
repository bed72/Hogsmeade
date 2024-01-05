package responses

type (
	ErrorResponseModel struct {
		Message     string `json:"message"`
		Description string `json:"description"`
	}

	ValidatorResponseModel struct {
		Failure     bool
		FailedField string
		Tag         string
		Value       interface{}
	}
)
