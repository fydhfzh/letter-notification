package user_repository

import (
	"github.com/fydhfzh/letter-notification/entity"
	"github.com/fydhfzh/letter-notification/pkg/errs"
)

type UserRepository interface {
	Register(user entity.User) (*entity.User, errs.ErrMessage)
	GetUsersBySubditID(subditID int) ([]entity.User, errs.ErrMessage)
	GetUserByID(id int) (*entity.User, errs.ErrMessage)
	GetUserByEmail(email string) (*entity.User, errs.ErrMessage)
	Update(userUpdate entity.User) errs.ErrMessage
}
