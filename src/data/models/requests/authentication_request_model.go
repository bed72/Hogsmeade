package models

type SignUpRequestModel struct {
	Email    string `json:"email" validate:"required,email,lte=128"`
	Password string `json:"password" validate:"required,lte=64"`
}
