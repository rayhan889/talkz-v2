package responses

import "github.com/google/uuid"

type (
	RegisterReponse struct {
		ID        uuid.UUID `json:"id"`
		Username  string    `json:"username"`
		Email     string    `json:"email"`
		CreatedAt string    `json:"created_at"`
	}
	TokenResponse struct {
		AccessToken           string `json:"access_token"`
		AccessTokenExpiresIn  int64  `json:"access_token_expires_in"`
		RefreshToken          string `json:"refresh_token"`
		RefreshTokenExpiresIn int64  `json:"refresh_token_expires_in"`
	}
	LoggedUserResponse struct {
		ID       uuid.UUID `json:"id"`
		Username string    `json:"username"`
		Email    string    `json:"email"`
	}
)
