package dto

import "time"

type RegisterUserRequest struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=8"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	SubditID    int    `json:"subdit_id" validate:"required"`
}

type RegisterUserResponse struct {
	Status      int       `json:"status"`
	ID          int       `json:"user_id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Role        string    `json:"role"`
	SubditID    int       `json:"subdit_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginUserResponse struct {
	Status int    `json:"status"`
	Token  string `json:"token"`
}

type User struct {
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Role        string    `json:"role"`
	SubditID    int       `json:"subdit_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetUserByIDResponse struct {
	Status      int       `json:"status"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Role        string    `json:"role"`
	SubditID    int       `json:"subdit_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetUsersBySubditIDResponse struct {
	Status          int    `json:"status"`
	UsersBySubditID []User `json:"users"`
}

type LogoutResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ResetPasswordRequest struct {
	Email           string `json:"email"`
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required,min:8"`
}

type ResetPasswordResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
