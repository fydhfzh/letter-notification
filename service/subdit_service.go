package service

import (
	"net/http"

	"github.com/fydhfzh/letter-notification/dto"
	"github.com/fydhfzh/letter-notification/entity"
	"github.com/fydhfzh/letter-notification/pkg/errs"
	"github.com/fydhfzh/letter-notification/repository/subdit_repository"
)

type subditService struct {
	subditRepository subdit_repository.SubditRepository
}

type SubditService interface {
	CreateSubdit(subditPayload dto.CreateSubditRequest) (*dto.CreateSubditResponse, errs.ErrMessage)
	GetSubditByID(id int) (*dto.GetSubditByIDResponse, errs.ErrMessage)
	GetAllSubdit() (*dto.GetAllSubditResponse, errs.ErrMessage)
	DeleteSubditByID(id int) (*dto.DeleteSubditByIDResponse, errs.ErrMessage)
	UpdateSubditByID(id int, subditPayload dto.UpdateSubditByIDRequest) (*dto.UpdateSubditByIDResponse, errs.ErrMessage)
}

func NewSubditService(subditRepository subdit_repository.SubditRepository) SubditService {
	return &subditService{
		subditRepository: subditRepository,
	}
}

func (s *subditService) CreateSubdit(subditPayload dto.CreateSubditRequest) (*dto.CreateSubditResponse, errs.ErrMessage) {
	subdit := entity.Subdit{
		Name: subditPayload.Name,
	}

	createdSubdit, err := s.subditRepository.CreateSubdit(subdit)

	if err != nil {
		return nil, err
	}

	response := dto.CreateSubditResponse{
		Status:    http.StatusCreated,
		ID:        int(createdSubdit.ID),
		Name:      createdSubdit.Name,
		CreatedAt: createdSubdit.CreatedAt,
		UpdatedAt: createdSubdit.UpdatedAt,
	}

	return &response, nil
}

func (s *subditService) GetSubditByID(id int) (*dto.GetSubditByIDResponse, errs.ErrMessage) {
	subdit, err := s.subditRepository.GetSubditByID(id)

	if err != nil {
		return nil, err
	}

	response := dto.GetSubditByIDResponse{
		Status:    http.StatusOK,
		ID:        int(subdit.ID),
		Name:      subdit.Name,
		CreatedAt: subdit.CreatedAt,
		UpdatedAt: subdit.UpdatedAt,
	}

	return &response, nil
}

func (s *subditService) GetAllSubdit() (*dto.GetAllSubditResponse, errs.ErrMessage) {
	allSubdit, err := s.subditRepository.GetAllSubdit()

	if err != nil {
		return nil, err
	}

	var allSubditResponse []dto.SubditResponse

	for _, subdit := range allSubdit {
		subditResponse := dto.SubditResponse{
			ID:        int(subdit.ID),
			Name:      subdit.Name,
			CreatedAt: subdit.CreatedAt,
			UpdatedAt: subdit.UpdatedAt,
		}

		allSubditResponse = append(allSubditResponse, subditResponse)
	}

	response := dto.GetAllSubditResponse{
		Status:    http.StatusOK,
		AllSubdit: allSubditResponse,
	}

	return &response, nil
}

func (s *subditService) DeleteSubditByID(id int) (*dto.DeleteSubditByIDResponse, errs.ErrMessage) {
	message, err := s.subditRepository.DeleteSubditByID(id)

	if err != nil {
		return nil, err
	}

	response := dto.DeleteSubditByIDResponse{
		Status:  http.StatusOK,
		Message: message,
	}

	return &response, nil
}

func (s *subditService) UpdateSubditByID(id int, subditPayload dto.UpdateSubditByIDRequest) (*dto.UpdateSubditByIDResponse, errs.ErrMessage) {
	subdit := entity.Subdit{
		Name: subditPayload.Name,
	}

	updatedSubdit, err := s.subditRepository.UpdateSubditByID(id, subdit)

	if err != nil {
		return nil, err
	}

	response := dto.UpdateSubditByIDResponse{
		Status:    http.StatusOK,
		ID:        int(updatedSubdit.ID),
		Name:      updatedSubdit.Name,
		CreatedAt: updatedSubdit.CreatedAt,
		UpdatedAt: updatedSubdit.UpdatedAt,
	}

	return &response, nil
}
