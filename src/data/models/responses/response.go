package responses

type Response struct {
	IsSuccess bool        `json:"is_success"`
	Data      interface{} `json:"data"`
}
