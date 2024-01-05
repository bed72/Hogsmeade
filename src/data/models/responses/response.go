package responses

type ResponseModel[T interface{}] struct {
	IsSuccess bool `json:"is_success"`
	Data      T    `json:"data"`
}
