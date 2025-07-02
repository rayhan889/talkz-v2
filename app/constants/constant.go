package constants

var (
	ErrorInvalidRequestBody         = "Invalid request body!"
	ErrorEmailAlreadyExists         = "Email already exists"
	JSONContentType                 = "application/json"
	InvalidEmailOrPassword          = "Invalid email or password"
	AccessTokenPrefix               = "Bearer "
	MissingAccessTokenError         = "Missing access token in request header"
	InvalidAccessTokenSigningMethod = "Invalid access token signing method"
	InvalidAccessToken              = "Invalid access token"
	RefreshTokenNotFound            = "Refresh token not found"
	RefreshTokenExpired             = "Refresh token expired"

	FeedCacheKey = "blogs:feed"
)
