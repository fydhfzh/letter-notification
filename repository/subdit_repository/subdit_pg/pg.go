package subdit_pg

import (
	"errors"

	"github.com/fydhfzh/letter-notification/entity"
	"github.com/fydhfzh/letter-notification/pkg/errs"
	"github.com/fydhfzh/letter-notification/repository/subdit_repository"
	"gorm.io/gorm"
)

type subditRepository struct {
	db *gorm.DB
}

func NewSubditRepository(db *gorm.DB) subdit_repository.SubditRepository {
	return &subditRepository{
		db: db,
	}
}

func (s *subditRepository) CreateSubdit(subdit entity.Subdit) (*entity.Subdit, errs.ErrMessage) {
	result := s.db.Create(&subdit)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errs.NewBadRequestError("Subdit already exists")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}

	return &subdit, nil
}

func (s *subditRepository) GetSubditByID(id int) (*entity.Subdit, errs.ErrMessage) {
	var subdit entity.Subdit

	result := s.db.First(&subdit, id)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewBadRequestError("Subdit not found")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}

	return &subdit, nil
}

func (s *subditRepository) GetAllSubdit() ([]entity.Subdit, errs.ErrMessage) {
	var allSubdit []entity.Subdit

	result := s.db.Find(&allSubdit)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewBadRequestError("Subdit not found")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}

	if result.RowsAffected == 0 {
		return nil, errs.NewBadRequestError("Subdit not found")
	}

	return allSubdit, nil
}

func (s *subditRepository) DeleteSubditByID(id int) (string, errs.ErrMessage) {
	result := s.db.Delete(&entity.Subdit{}, id)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errs.NewBadRequestError("Subdit not found")
		}

		return "", errs.NewInternalServerError("Something went wrong")
	}

	return "Subdit has been deleted", nil
}

func (s *subditRepository) UpdateSubditByID(id int, subdit entity.Subdit) (*entity.Subdit, errs.ErrMessage) {
	result := s.db.Model(&entity.Subdit{}).Where("id = ?", id).Updates(subdit)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewBadRequestError("Subdit not found")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}

	_ = s.db.First(&subdit)

	return &subdit, nil
}
