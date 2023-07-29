package helpers

import (
	"github.com/fydhfzh/letter-notification/pkg/errs"
	valid "github.com/go-playground/validator/v10"
)

var validator *valid.Validate = valid.New()

func ValidateStruct(payload interface{}) errs.ErrMessage {
	err := validator.Struct(payload)

	if err != nil {
		return errs.NewBadRequestError(err.Error())
	}

	return nil
}
