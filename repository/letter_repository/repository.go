package letter_repository

import (
	"github.com/fydhfzh/letter-notification/entity"
	"github.com/fydhfzh/letter-notification/pkg/errs"
)

type LetterRepository interface {
	CreateLetter(letter entity.Letter) (*entity.Letter, errs.ErrMessage)
	GetLetterByID(id int) (*entity.Letter, errs.ErrMessage)
	GetIncomingLettersByToSubditID(toSubditID int, userID int) ([]entity.UserLetter, errs.ErrMessage)
	GetOutcomingLettersByToSubditID(toSubditID int, userID int) ([]entity.UserLetter, errs.ErrMessage)
	GetArchivedLettersByToSubditID(toSubditID int, userID int) ([]entity.UserLetter, errs.ErrMessage)
	SetIsNotifyTrue(id int) errs.ErrMessage
}
