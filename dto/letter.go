package dto

import (
	"time"

	"github.com/fydhfzh/letter-notification/entity"
)

type CreateLetterRequest struct {
	Name       string    `json:"name" validate:"required"`
	About      string    `json:"about"`
	Number     string    `json:"number"`
	Datetime   time.Time `json:"datetime"`
	From       string    `json:"from"`
	ToSubditID int       `json:"to_subdit_id"`
}

type CreateLetterResponse struct {
	Status     int       `json:"status"`
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	About      string    `json:"about"`
	Number     string    `json:"number"`
	Datetime   time.Time `json:"datetime"`
	From       string    `json:"from"`
	ToSubditID int       `json:"to_subdit_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type GetUserLetterByIDResponse struct {
	Status     int       `json:"status"`
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	About      string    `json:"about"`
	Number     string    `json:"number"`
	Datetime   time.Time `json:"datetime"`
	From       string    `json:"from"`
	ToSubditID int       `json:"to_subdit_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type GetUserLettersByToSubditIDResponse struct {
	Status  int                 `json:"status"`
	Letters []entity.UserLetter `json:"letters"`
}

type SendLetterToMailScheduler struct {
	LetterID   int           `json:"letter_id"`
	About      string        `json:"message"`
	Datetime   time.Time     `json:"datetime"`
	Recipients []entity.User `json:"recipients"`
}
