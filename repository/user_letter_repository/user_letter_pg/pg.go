package user_letter_pg

import (
	"errors"
	"fmt"

	"github.com/fydhfzh/letter-notification/entity"
	"github.com/fydhfzh/letter-notification/pkg/errs"
	"github.com/fydhfzh/letter-notification/repository/user_letter_repository"
	"gorm.io/gorm"
)

type userLetterRepository struct {
	db *gorm.DB
}

func NewUserLetterRepository(db *gorm.DB) user_letter_repository.UserLetterRepository {
	return &userLetterRepository{
		db: db,
	}
}

func (u *userLetterRepository) AssignUserLetter(letterID int, users []entity.User) ([]entity.UserLetter, errs.ErrMessage) {
	var userLetters []entity.UserLetter

	for _, user := range users {
		userLetter := entity.UserLetter{
			UserID:   uint(user.ID),
			LetterID: uint(letterID),
		}

		userLetters = append(userLetters, userLetter)
	}

	fmt.Println(userLetters)

	result := u.db.Create(userLetters)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrInvalidData) {
			return nil, errs.NewBadRequestError("Invalid param")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}

	return userLetters, nil
}

func (u *userLetterRepository) ArchiveLetter(id int) errs.ErrMessage {
	result := u.db.Model(&entity.UserLetter{}).Where("letter_id = ?", id).Update("is_archived", true)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewBadRequestError("User letter not found")
		}

		return errs.NewInternalServerError("Something went wrong")
	}

	return nil
}

func (u *userLetterRepository) DeleteLetterByID(id int) errs.ErrMessage {
	result := u.db.Where("letter_id = ?", id).Delete(&entity.UserLetter{})

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewBadRequestError("User letter not found")
		}

		return errs.NewInternalServerError("Something went wrong")
	}

	return nil
}
