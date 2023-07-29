package subdit_repository

import (
	"github.com/fydhfzh/letter-notification/entity"
	"github.com/fydhfzh/letter-notification/pkg/errs"
)

type SubditRepository interface {
	CreateSubdit(subdit entity.Subdit) (*entity.Subdit, errs.ErrMessage)
	GetSubditByID(id int) (*entity.Subdit, errs.ErrMessage)
	GetAllSubdit() ([]entity.Subdit, errs.ErrMessage)
	DeleteSubditByID(id int) (string, errs.ErrMessage)
	UpdateSubditByID(id int, subdit entity.Subdit) (*entity.Subdit, errs.ErrMessage)
}
