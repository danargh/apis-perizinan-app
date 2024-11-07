package auth

// Login
type AuthLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Register
type AuthRegisterRequest struct {
	Username string `json:"username"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type AuthRegisterResponse struct {
	Username string `json:"username"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
}
