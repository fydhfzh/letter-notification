package helpers

import (
	"os"
	"strconv"

	"github.com/fydhfzh/letter-notification/pkg/errs"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, errs.ErrMessage) {
	bytePass := []byte(password)
	salt, convErr := strconv.Atoi(os.Getenv("SALT"))

	if convErr != nil {
		err := errs.NewBadRequestError("Invalid salt")

		return "", err
	}

	hashed, err := bcrypt.GenerateFromPassword(bytePass, salt)

	if err != nil {
		if err == bcrypt.ErrHashTooShort {
			return "", errs.NewBadRequestError("Hast too short")
		}

		return "", errs.NewInternalServerError("Something went wrong")
	}

	return string(hashed), nil
}

func ComparePassword(hashed string, password string) errs.ErrMessage {
	byteHash := []byte(hashed)
	passwordHash := []byte(password)

	err := bcrypt.CompareHashAndPassword(byteHash, passwordHash)

	if err != nil {
		return errs.NewBadRequestError("Wrong email/password")
	}

	return nil
}
