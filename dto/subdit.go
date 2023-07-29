package dto

import (
	"time"

	"github.com/fydhfzh/letter-notification/entity"
)

type CreateSubditRequest struct {
	Name string `json:"name" validate:"required"`
}

type CreateSubditResponse struct {
	Status    int       `json:"status"`
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetSubditByIDResponse struct {
	Status    int       `json:"status"`
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetAllSubditResponse struct {
	Status    int             `json:"status"`
	AllSubdit []entity.Subdit `json:"all_subdit"`
}

type DeleteSubditByIDResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type UpdateSubditByIDRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateSubditByIDResponse struct {
	Status    int       `json:"status"`
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
