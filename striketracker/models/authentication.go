package models

// AccessToken for controlling access to the API
type AccessToken struct {
	// Whether or not the token is active
	Active bool `json:"active"`
	// The name of the application which will use this token
	Application string `json:"application"`
	// 	Expiration date of the token
	Expiration string `json:"expiration"`
	// Unique ID for this AccessToken
	ID int `json:"id"`
	// The remote address of the client which created this token
	IP string `json:"ip"`
	// Whether or not the token can be used to refresh an access token
	Refresh bool `json:"refresh"`
	// 	The token used to authenticate with the API
	Token string `json:"token"`
}

// AccessTokenList ...
type AccessTokenList struct {
	List *AccessToken `json:"list"`
}

// APITokenRequest ...
type APITokenRequest struct {
	// The name of the application this token will be used for
	Application string `json:"application"`
	// The user's current password
	Password string `json:"password"`
}

// AuthToken ...
type AuthToken struct {
	AccessToken  string `json:"access_token"`
	Application  string `json:"application"`
	ExpiresIn    int    `json:"expires_in"`
	IP           string `json:"ip"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	UserAgent    string `json:"user_agent"`
}

// Authentication ...
type Authentication struct {
	Application string `json:"application"`
	IP          string `json:"ip"`
	Token       string `json:"token"`
}

// CreateTokenRequest is the payload for creating a new infinite expiration token
type CreateTokenRequest struct {
	AccountHash     string           `json:"account_hash"`
	UserID          string           `json:"user_id"`
	APITokenRequest *APITokenRequest `json:"api_token_request"`
}
