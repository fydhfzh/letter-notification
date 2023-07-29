package service

import (
	"net/http"

	"github.com/fydhfzh/letter-notification/dto"
	"github.com/fydhfzh/letter-notification/entity"
	"github.com/fydhfzh/letter-notification/pkg/errs"
	"github.com/fydhfzh/letter-notification/pkg/helpers"
	"github.com/fydhfzh/letter-notification/repository/user_repository"
)

type userService struct {
	userRepo user_repository.UserRepository
}

type UserService interface {
	Register(userPayload dto.RegisterUserRequest) (*dto.RegisterUserResponse, errs.ErrMessage)
	Login(userCredentials dto.LoginUserRequest) (*dto.LoginUserResponse, errs.ErrMessage)
	GetUserByID(id int) (*dto.GetUserByIDResponse, errs.ErrMessage)
	GetUsersBySubditID(subditID int) (*dto.GetUsersBySubditIDResponse, errs.ErrMessage)
	ResetPassword(userPayload dto.ResetPasswordRequest) (*dto.ResetPasswordResponse, errs.ErrMessage)
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) Register(userPayload dto.RegisterUserRequest) (*dto.RegisterUserResponse, errs.ErrMessage) {
	user := entity.User{
		Name:        userPayload.Name,
		Email:       userPayload.Email,
		PhoneNumber: userPayload.PhoneNumber,
		SubditID:    uint(userPayload.SubditID),
	}

	hashedPassword, err := helpers.HashPassword(userPayload.Password)

	if err != nil {
		return nil, err
	}

	user.Password = hashedPassword

	createdUser, err := u.userRepo.Register(user)

	if err != nil {
		return nil, err
	}

	response := dto.RegisterUserResponse{
		Status:      http.StatusCreated,
		ID:          int(createdUser.ID),
		Name:        createdUser.Name,
		Email:       createdUser.Email,
		PhoneNumber: createdUser.PhoneNumber,
		Role:        createdUser.Role,
		SubditID:    int(createdUser.SubditID),
		CreatedAt:   createdUser.CreatedAt,
		UpdatedAt:   createdUser.UpdatedAt,
	}

	return &response, nil
}

func (u *userService) GetUserByID(id int) (*dto.GetUserByIDResponse, errs.ErrMessage) {
	user, err := u.userRepo.GetUserByID(id)

	if err != nil {
		return nil, err
	}

	response := dto.GetUserByIDResponse{
		Status:      http.StatusOK,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Role:        user.Role,
		SubditID:    int(user.SubditID),
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}

	return &response, nil
}

func (u *userService) GetUsersBySubditID(subditID int) (*dto.GetUsersBySubditIDResponse, errs.ErrMessage) {
	users, err := u.userRepo.GetUsersBySubditID(subditID)

	if err != nil {
		return nil, err
	}

	usersResponse := []dto.User{}

	for _, user := range users {
		userResponse := dto.User{
			Name:        user.Name,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			Role:        user.Role,
			SubditID:    int(user.SubditID),
			CreatedAt:   user.CreatedAt,
			UpdatedAt:   user.UpdatedAt,
		}

		usersResponse = append(usersResponse, userResponse)
	}

	response := dto.GetUsersBySubditIDResponse{
		Status:          http.StatusOK,
		UsersBySubditID: usersResponse,
	}

	return &response, nil
}

func (u *userService) Login(userCredentials dto.LoginUserRequest) (*dto.LoginUserResponse, errs.ErrMessage) {
	user, err := u.userRepo.GetUserByEmail(userCredentials.Email)

	if err != nil {
		return nil, err
	}

	// Compare password
	if err := helpers.ComparePassword(user.Password, userCredentials.Password); err != nil {
		return nil, err
	}

	token, err := helpers.GenerateToken(int(user.ID), user.Name, user.Email, user.Role)

	if err != nil {
		return nil, err
	}

	response := dto.LoginUserResponse{
		Status: http.StatusOK,
		Token:  token,
	}

	return &response, nil
}

func (u *userService) ResetPassword(userPayload dto.ResetPasswordRequest) (*dto.ResetPasswordResponse, errs.ErrMessage) {
	user, err := u.userRepo.GetUserByEmail(userPayload.Email)

	if err != nil {
		return nil, err
	}

	if err := helpers.ComparePassword(user.Password, userPayload.CurrentPassword); err != nil {
		return nil, err
	}

	newPassword, err := helpers.HashPassword(userPayload.NewPassword)

	if err != nil {
		return nil, err
	}

	user.Email = userPayload.Email
	user.Password = newPassword

	err = u.userRepo.Update(*user)

	if err != nil {
		return nil, err
	}

	response := dto.ResetPasswordResponse{
		Status:  http.StatusOK,
		Message: "Password has been changed successfully",
	}

	return &response, nil
}
