package entities

import "github.com/google/uuid"

type AuthenticationEntity struct {
	AccessToken  string     `json:"access_token"`
	ExpiresIn    int        `json:"expires_in"`
	RefreshToken string     `json:"refresh_token"`
	User         UserEntity `json:"user"`
}

type UserEntity struct {
	Id    uuid.UUID `json:"id"`
	Email string    `json:"email"`
}
