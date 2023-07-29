package handler

import (
	"strconv"

	"github.com/fydhfzh/letter-notification/dto"
	"github.com/fydhfzh/letter-notification/pkg/errs"
	"github.com/fydhfzh/letter-notification/pkg/helpers"
	"github.com/fydhfzh/letter-notification/service"
	"github.com/gin-gonic/gin"
)

type subditHandler struct {
	subditService service.SubditService
}

type SubditHandler interface {
	CreateSubdit(c *gin.Context)
	GetSubditByID(c *gin.Context)
	GetAllSubdit(c *gin.Context)
	DeleteSubditByID(c *gin.Context)
	UpdateSubditByID(c *gin.Context)
}

func NewSubditHandler(subditService service.SubditService) SubditHandler {
	return &subditHandler{
		subditService: subditService,
	}
}

func (s *subditHandler) CreateSubdit(c *gin.Context) {
	var subditPayload dto.CreateSubditRequest

	if err := c.ShouldBindJSON(&subditPayload); err != nil {
		bindErr := errs.NewUnprocessableEntityError(err.Error())

		c.JSON(bindErr.Status(), bindErr)
		return
	}

	if err := helpers.ValidateStruct(subditPayload); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	response, err := s.subditService.CreateSubdit(subditPayload)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.Status, response)
}

func (s *subditHandler) GetSubditByID(c *gin.Context) {
	id, convErr := strconv.Atoi(c.Param("subditID"))

	if convErr != nil {
		err := errs.NewBadRequestError("Invalid subdit param")

		c.JSON(err.Status(), err)
		return
	}

	response, err := s.subditService.GetSubditByID(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.Status, response)
}

func (s *subditHandler) GetAllSubdit(c *gin.Context) {
	response, err := s.subditService.GetAllSubdit()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.Status, response)
}

func (s *subditHandler) DeleteSubditByID(c *gin.Context) {
	id, convErr := strconv.Atoi(c.Param("subditID"))

	if convErr != nil {
		err := errs.NewBadRequestError("Invalid subdit param")

		c.JSON(err.Status(), err)
		return
	}

	response, err := s.subditService.DeleteSubditByID(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.Status, response)
}

func (s *subditHandler) UpdateSubditByID(c *gin.Context) {
	id, convErr := strconv.Atoi(c.Param("subditID"))

	if convErr != nil {
		err := errs.NewBadRequestError("Invalid subdit param")

		c.JSON(err.Status(), err)
		return
	}

	var subditPayload dto.UpdateSubditByIDRequest

	if err := c.ShouldBindJSON(&subditPayload); err != nil {
		bindErr := errs.NewUnprocessableEntityError(err.Error())

		c.JSON(bindErr.Status(), bindErr)
		return
	}

	if err := helpers.ValidateStruct(subditPayload); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	response, err := s.subditService.UpdateSubditByID(id, subditPayload)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.Status, response)
}
