package letter_pg

import (
	"errors"

	"github.com/fydhfzh/letter-notification/entity"
	"github.com/fydhfzh/letter-notification/pkg/errs"
	"github.com/fydhfzh/letter-notification/repository/letter_repository"
	"gorm.io/gorm"
)

type letterRepository struct {
	db *gorm.DB
}

func NewLetterRepository(db *gorm.DB) letter_repository.LetterRepository {
	return &letterRepository{
		db: db,
	}
}

func (l *letterRepository) CreateLetter(letter entity.Letter) (*entity.Letter, errs.ErrMessage) {
	result := l.db.Create(&letter)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errs.NewBadRequestError("Letter number already exists")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}

	return &letter, nil
}

// Query where to is_archived = false
func (l *letterRepository) GetLetterByID(id int) (*entity.UserLetter, errs.ErrMessage) {
	var letter entity.UserLetter

	result := l.db.Joins("Letter").Joins("User").Where("is_archived = ?", false).First(&letter, id)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewBadRequestError("Letter not found")
		}

		return nil, errs.NewBadRequestError("Something went wrong")
	}

	return &letter, nil
}

func (l *letterRepository) GetLettersByToSubditID(toSubditID int) ([]entity.UserLetter, errs.ErrMessage) {
	var letters []entity.UserLetter

	result := l.db.Joins("Letter").Joins("User").Where("is_archived = ?", false).Where("to_subdit_id = ?", toSubditID).Find(&letters)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewBadRequestError("Letter not found")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}

	if result.RowsAffected == 0 {
		return nil, errs.NewBadRequestError("Letter not found")
	}

	return letters, nil
}

func (l *letterRepository) SetIsNotifyTrue(id int) errs.ErrMessage {
	var letter entity.Letter

	result := l.db.Where("id = ?", id).First(&letter)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewBadRequestError("Letter not found")
		}

		return errs.NewInternalServerError("Something went wrong")
	}

	letter.IsNotified = true

	result = l.db.Save(&letter)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewBadRequestError("Letter not found")
		}

		return errs.NewInternalServerError("Something went wrong")
	}

	return nil
}
