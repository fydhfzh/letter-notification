package user_pg

import (
	"errors"

	"github.com/fydhfzh/letter-notification/entity"
	"github.com/fydhfzh/letter-notification/pkg/errs"
	"github.com/fydhfzh/letter-notification/repository/user_repository"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user_repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Register(user entity.User) (*entity.User, errs.ErrMessage) {
	result := u.db.Create(&user)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errs.NewBadRequestError("Email already exists")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}

	return &user, nil
}

func (u *userRepository) GetUserByID(id int) (*entity.User, errs.ErrMessage) {
	var user entity.User

	result := u.db.First(&user, id)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewBadRequestError("User not found")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}

	return &user, nil
}

func (u *userRepository) GetUsersBySubditID(subditID int) ([]entity.User, errs.ErrMessage) {
	var users []entity.User

	result := u.db.Where("subdit_id = ?", subditID).Find(&users)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewBadRequestError("User not found")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}

	return users, nil
}

func (u *userRepository) GetUserByEmail(email string) (*entity.User, errs.ErrMessage) {
	var user entity.User

	result := u.db.Where("email = ?", email).First(&user)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewBadRequestError("Email not found")
		}

		return nil, errs.NewInternalServerError("Something went wrong")
	}

	return &user, nil
}

func (u *userRepository) Update(userUpdate entity.User) errs.ErrMessage {
	result := u.db.Model(&entity.User{}).Where("id = ?", userUpdate.ID).Update("password", userUpdate.Password)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewBadRequestError("User not found")
		}

		return errs.NewInternalServerError("Something went wrong")
	}

	return nil
}
