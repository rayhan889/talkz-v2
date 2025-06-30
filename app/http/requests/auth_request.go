package requests

type (
	RegisterRequest struct {
		Username string `json:"username" validate:"required,min=3,max=20"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6,max=100"`
	}
	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
	RefreshTokenRequest struct {
		Token string `json:"token" validate:"required"`
	}
)
