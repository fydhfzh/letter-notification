package handler

import (
	"net/http"
	"strconv"

	"github.com/fydhfzh/letter-notification/dto"
	"github.com/fydhfzh/letter-notification/entity"
	"github.com/fydhfzh/letter-notification/pkg/errs"
	"github.com/fydhfzh/letter-notification/pkg/helpers"
	"github.com/fydhfzh/letter-notification/service"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

type UserHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
	ResetPassword(c *gin.Context)
	GetUserByID(c *gin.Context)
	GetUsersBySubditID(c *gin.Context)
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (u *userHandler) Register(c *gin.Context) {
	var userPayload dto.RegisterUserRequest

	if err := c.ShouldBindJSON(&userPayload); err != nil {
		bindErr := errs.NewUnprocessableEntityError(err.Error())

		c.JSON(bindErr.Status(), bindErr)
		return
	}

	if err := helpers.ValidateStruct(userPayload); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	response, err := u.userService.Register(userPayload)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.Status, response)
}

func (u *userHandler) GetUserByID(c *gin.Context) {
	id, convErr := strconv.Atoi(c.Param("userID"))

	if convErr != nil {
		err := errs.NewBadRequestError("Invalid user param")

		c.JSON(err.Status(), err)
		return
	}

	userData, isOK := c.MustGet("userData").(*entity.User)

	if !isOK {
		err := errs.NewUnauthenticatedError("You are not authenticated")

		c.JSON(err.Status(), err)
		return
	}

	if userData.ID != uint(id) {
		err := errs.NewUnauthorizedError("You are not authorized")

		c.JSON(err.Status(), err)
		return
	}

	response, err := u.userService.GetUserByID(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.Status, response)
}

func (u *userHandler) GetUsersBySubditID(c *gin.Context) {
	query, isOK := c.GetQuery("subditID")

	if !isOK {
		err := errs.NewBadRequestError("Invalid query")

		c.JSON(err.Status(), err)
		return
	}

	subditID, convErr := strconv.Atoi(query)

	if convErr != nil {
		err := errs.NewBadRequestError("Invalid query")

		c.JSON(err.Status(), err)
		return
	}

	response, err := u.userService.GetUsersBySubditID(subditID)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.Status, response)
}

func (u *userHandler) Login(c *gin.Context) {
	var userCredentials dto.LoginUserRequest

	if bindErr := c.ShouldBindJSON(&userCredentials); bindErr != nil {
		err := errs.NewUnprocessableEntityError("Invalid credentials")

		c.JSON(err.Status(), err)
		return
	}

	if err := helpers.ValidateStruct(userCredentials); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	response, err := u.userService.Login(userCredentials)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.Status, response)
}

func (u *userHandler) Logout(c *gin.Context) {
	c.Set("userData", nil)

	response := dto.LogoutResponse{
		Status:  http.StatusOK,
		Message: "You have successfully logged out",
	}

	c.JSON(response.Status, response)
}

func (u *userHandler) ResetPassword(c *gin.Context) {
	user, isOK := c.MustGet("userData").(*entity.User)

	if !isOK {
		unauthenticatedErr := errs.NewUnauthenticatedError("You are not authenticated")

		c.JSON(unauthenticatedErr.Status(), unauthenticatedErr)
		return
	}

	email := user.Email

	var userPayload dto.ResetPasswordRequest

	if bindErr := c.ShouldBindJSON(&userPayload); bindErr != nil {
		err := errs.NewBadRequestError(bindErr.Error())

		c.JSON(err.Status(), err)
		return
	}

	if err := helpers.ValidateStruct(userPayload); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	userPayload.Email = email

	response, err := u.userService.ResetPassword(userPayload)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.Status, response)
}
