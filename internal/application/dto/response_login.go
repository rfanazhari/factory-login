package dto

// LoginResponse represents the login response
type LoginResponse struct {
	Success     bool   `json:"success"`
	Message     string `json:"message"`
	AccessToken string `json:"access_token,omitempty"`
	UserID      string `json:"user_id,omitempty"`
	ExpiresAt   int64  `json:"expires_at,omitempty"`
}
