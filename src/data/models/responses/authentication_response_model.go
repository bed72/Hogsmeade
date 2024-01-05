package responses

import "github.com/google/uuid"

type AuthenticationResponseModel struct {
	Id           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	ExpiresIn    int       `json:"expires_in"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
}
