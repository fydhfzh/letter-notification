package user_letter_repository

import (
	"github.com/fydhfzh/letter-notification/entity"
	"github.com/fydhfzh/letter-notification/pkg/errs"
)

type UserLetterRepository interface {
	AssignUserLetter(letterID int, usersID []entity.User) ([]entity.UserLetter, errs.ErrMessage)
	ArchiveLetter(id int) errs.ErrMessage
	DeleteLetterByID(id int) errs.ErrMessage
}
