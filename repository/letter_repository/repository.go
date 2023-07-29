package letter_repository

import (
	"github.com/fydhfzh/letter-notification/entity"
	"github.com/fydhfzh/letter-notification/pkg/errs"
)

type LetterRepository interface {
	CreateLetter(letter entity.Letter) (*entity.Letter, errs.ErrMessage)
	GetLetterByID(id int) (*entity.UserLetter, errs.ErrMessage)
	GetLettersByToSubditID(toSubditID int) ([]entity.UserLetter, errs.ErrMessage)
	SetIsNotifyTrue(id int) errs.ErrMessage
}
