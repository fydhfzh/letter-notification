package handler

import (
	"strconv"

	"github.com/fydhfzh/letter-notification/dto"
	"github.com/fydhfzh/letter-notification/pkg/errs"
	"github.com/fydhfzh/letter-notification/pkg/helpers"
	"github.com/fydhfzh/letter-notification/service"
	"github.com/gin-gonic/gin"
)

type letterHandler struct {
	letterService service.LetterService
}

type LetterHandler interface {
	CreateLetter(c *gin.Context)
	GetLetterByID(c *gin.Context)
	GetLettersByToSubditID(c *gin.Context)
	ArchiveLetter(c *gin.Context)
	DeleteLetterByID(c *gin.Context)
}

func NewLetterHandler(letterService service.LetterService) LetterHandler {
	return &letterHandler{
		letterService: letterService,
	}
}

func (l *letterHandler) CreateLetter(c *gin.Context) {
	var letterPayload dto.CreateLetterRequest

	if err := c.ShouldBindJSON(&letterPayload); err != nil {
		bindErr := errs.NewUnprocessableEntityError(err.Error())

		c.JSON(bindErr.Status(), bindErr)
	}

	if err := helpers.ValidateStruct(letterPayload); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	response, err := l.letterService.CreateLetter(letterPayload)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.Status, response)
}

func (l *letterHandler) GetLetterByID(c *gin.Context) {
	id, convErr := strconv.Atoi(c.Param("letterID"))

	if convErr != nil {
		err := errs.NewBadRequestError("Invalid url parameter")

		c.JSON(err.Status(), err)
		return
	}

	response, err := l.letterService.GetLetterByID(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.Status, response)
}

func (l *letterHandler) GetLettersByToSubditID(c *gin.Context) {
	query, isOK := c.GetQuery("toSubditID")

	if !isOK {
		queryErr := errs.NewBadRequestError("Invalid query")

		c.JSON(queryErr.Status(), queryErr)
		return
	}

	toSubditID, convErr := strconv.Atoi(query)

	if convErr != nil {
		err := errs.NewBadRequestError("Invalid query")

		c.JSON(err.Status(), err)
		return
	}

	response, err := l.letterService.GetLettersByToSubditID(toSubditID)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.Status, response)
}

func (l *letterHandler) ArchiveLetter(c *gin.Context) {
	id, convErr := strconv.Atoi(c.Param("letterID"))

	if convErr != nil {
		err := errs.NewBadRequestError("Invalid letter id param")

		c.JSON(err.Status(), err)
		return
	}

	response, err := l.letterService.ArchiveLetter(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.Status, response)
}

func (l *letterHandler) DeleteLetterByID(c *gin.Context) {
	id, convErr := strconv.Atoi(c.Param("letterID"))

	if convErr != nil {
		err := errs.NewBadRequestError("Invalid letter id param")

		c.JSON(err.Status(), err)
		return
	}

	response, err := l.letterService.DeleteLetterByID(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.Status, response)
}
