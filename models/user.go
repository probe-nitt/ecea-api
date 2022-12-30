package models

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type ForgetPasswordRequest struct {
	Email string `json:"email"`
}

type ForgetPasswordResponse struct {
	Message string `json:"message"`
}

type ChangePasswordRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ChangePasswordResponse struct {
	Message string `json:"message"`
}
