package service

import (
	"net/http"

	"github.com/fydhfzh/letter-notification/dto"
	"github.com/fydhfzh/letter-notification/entity"
	"github.com/fydhfzh/letter-notification/pkg/errs"
	"github.com/fydhfzh/letter-notification/pkg/mailer"
	"github.com/fydhfzh/letter-notification/repository/letter_repository"
	"github.com/fydhfzh/letter-notification/repository/user_letter_repository"
	"github.com/fydhfzh/letter-notification/repository/user_repository"
)

type letterService struct {
	letterRepo     letter_repository.LetterRepository
	userRepo       user_repository.UserRepository
	userLetterRepo user_letter_repository.UserLetterRepository
}

type LetterService interface {
	CreateLetter(letterPayload dto.CreateLetterRequest) (*dto.CreateLetterResponse, errs.ErrMessage)
	GetLetterByID(id int) (*dto.GetUserLetterByIDResponse, errs.ErrMessage)
	GetIncomingLettersByToSubditID(toSubditID int, userID int) (*dto.GetUserLettersByToSubditIDResponse, errs.ErrMessage)
	GetOutcomingLettersByToSubditID(toSubditID int, userID int) (*dto.GetUserLettersByToSubditIDResponse, errs.ErrMessage)
	GetArchivedLettersByToSubditID(toSubditID int, userID int) (*dto.GetUserLettersByToSubditIDResponse, errs.ErrMessage)
	ArchiveLetter(id int) (*dto.UserLetterArchiveResponse, errs.ErrMessage)
	DeleteLetterByID(id int) (*dto.UserLetterDeleteResponse, errs.ErrMessage)
}

func NewLetterService(letterRepo letter_repository.LetterRepository, userRepo user_repository.UserRepository, userLetterRepo user_letter_repository.UserLetterRepository) LetterService {
	return &letterService{
		letterRepo:     letterRepo,
		userRepo:       userRepo,
		userLetterRepo: userLetterRepo,
	}
}

func (l *letterService) CreateLetter(letterPayload dto.CreateLetterRequest) (*dto.CreateLetterResponse, errs.ErrMessage) {
	users, err := l.userRepo.GetUsersBySubditID(letterPayload.ToSubditID)

	if err != nil {
		return nil, err
	}

	letter := entity.Letter{
		Name:       letterPayload.Name,
		About:      letterPayload.About,
		Number:     letterPayload.Number,
		Datetime:   letterPayload.Datetime,
		Type:       letterPayload.Type,
		From:       letterPayload.From,
		ToSubditID: uint(letterPayload.ToSubditID),
	}

	createdLetter, err := l.letterRepo.CreateLetter(letter)

	if err != nil {
		return nil, err
	}

	_, err = l.userLetterRepo.AssignUserLetter(int(createdLetter.ID), users)

	if err != nil {
		return nil, err
	}

	scheduler := dto.SendLetterToMailScheduler{
		LetterID:   int(createdLetter.ID),
		About:      createdLetter.About,
		Datetime:   createdLetter.Datetime,
		Recipients: users,
	}

	mailer.SetSchedule(scheduler, l.letterRepo)

	response := dto.CreateLetterResponse{
		Status:     http.StatusCreated,
		ID:         int(createdLetter.ID),
		Name:       createdLetter.Name,
		About:      createdLetter.About,
		Number:     createdLetter.Number,
		Datetime:   createdLetter.Datetime,
		From:       createdLetter.From,
		ToSubditID: int(createdLetter.ToSubditID),
		CreatedAt:  createdLetter.CreatedAt,
		UpdatedAt:  createdLetter.UpdatedAt,
	}

	return &response, nil
}

func (l *letterService) GetLetterByID(id int) (*dto.GetUserLetterByIDResponse, errs.ErrMessage) {
	userLetter, err := l.letterRepo.GetLetterByID(id)

	if err != nil {
		return nil, err
	}

	response := dto.GetUserLetterByIDResponse{
		Status:     http.StatusOK,
		ID:         int(userLetter.ID),
		Name:       userLetter.Name,
		About:      userLetter.About,
		Number:     userLetter.Number,
		Datetime:   userLetter.Datetime,
		From:       userLetter.From,
		ToSubditID: int(userLetter.ToSubditID),
		CreatedAt:  userLetter.CreatedAt,
		UpdatedAt:  userLetter.UpdatedAt,
	}

	return &response, nil
}

func (l *letterService) GetIncomingLettersByToSubditID(toSubditID int, userID int) (*dto.GetUserLettersByToSubditIDResponse, errs.ErrMessage) {
	userLetters, err := l.letterRepo.GetIncomingLetters(userID)

	if err != nil {
		return nil, err
	}

	var userLettersResponse []dto.LetterResponse

	for _, userLetter := range userLetters {
		userLetterResponse := dto.LetterResponse{
			ID:         int(userLetter.LetterID),
			Name:       userLetter.Letter.Name,
			About:      userLetter.Letter.About,
			Number:     userLetter.Letter.Number,
			Datetime:   userLetter.Letter.Datetime,
			From:       userLetter.Letter.From,
			Type:       userLetter.Letter.Type,
			ToSubditID: int(userLetter.Letter.ToSubditID),
			IsArchived: userLetter.IsArchived,
			CreatedAt:  userLetter.Letter.CreatedAt,
			UpdatedAt:  userLetter.Letter.UpdatedAt,
		}

		userLettersResponse = append(userLettersResponse, userLetterResponse)
	}

	response := dto.GetUserLettersByToSubditIDResponse{
		Status:  http.StatusOK,
		Letters: userLettersResponse,
	}

	return &response, nil
}

func (l *letterService) GetOutcomingLettersByToSubditID(toSubditID int, userID int) (*dto.GetUserLettersByToSubditIDResponse, errs.ErrMessage) {
	userLetters, err := l.letterRepo.GetOutcomingLetters(userID)

	if err != nil {
		return nil, err
	}

	var userLettersResponse []dto.LetterResponse

	for _, userLetter := range userLetters {
		userLetterResponse := dto.LetterResponse{
			ID:         int(userLetter.LetterID),
			Name:       userLetter.Letter.Name,
			About:      userLetter.Letter.About,
			Number:     userLetter.Letter.Number,
			Datetime:   userLetter.Letter.Datetime,
			From:       userLetter.Letter.From,
			Type:       userLetter.Letter.Type,
			ToSubditID: int(userLetter.Letter.ToSubditID),
			IsArchived: userLetter.IsArchived,
			CreatedAt:  userLetter.Letter.CreatedAt,
			UpdatedAt:  userLetter.Letter.UpdatedAt,
		}

		userLettersResponse = append(userLettersResponse, userLetterResponse)
	}

	response := dto.GetUserLettersByToSubditIDResponse{
		Status:  http.StatusOK,
		Letters: userLettersResponse,
	}

	return &response, nil
}

func (l *letterService) ArchiveLetter(id int) (*dto.UserLetterArchiveResponse, errs.ErrMessage) {
	_, err := l.letterRepo.GetLetterByID(id)

	if err != nil {
		return nil, err
	}

	err = l.userLetterRepo.ArchiveLetter(id)

	if err != nil {
		return nil, err
	}

	response := dto.UserLetterArchiveResponse{
		Status:  http.StatusOK,
		Message: "Letter has been archived successfully",
	}

	return &response, nil
}

func (l *letterService) DeleteLetterByID(id int) (*dto.UserLetterDeleteResponse, errs.ErrMessage) {
	_, err := l.letterRepo.GetLetterByID(id)

	if err != nil {
		return nil, err
	}

	err = l.userLetterRepo.DeleteLetterByID(id)

	if err != nil {
		return nil, err
	}

	response := dto.UserLetterDeleteResponse{
		Status:  http.StatusOK,
		Message: "Letter has been deleted successfully",
	}

	return &response, nil
}

func (l *letterService) GetArchivedLettersByToSubditID(toSubditID int, userID int) (*dto.GetUserLettersByToSubditIDResponse, errs.ErrMessage) {
	userLetters, err := l.letterRepo.GetArchivedLetters(userID)

	if err != nil {
		return nil, err
	}

	var userLettersResponse []dto.LetterResponse

	for _, userLetter := range userLetters {
		userLetterResponse := dto.LetterResponse{
			ID:         int(userLetter.LetterID),
			Name:       userLetter.Letter.Name,
			About:      userLetter.Letter.About,
			Number:     userLetter.Letter.Number,
			Datetime:   userLetter.Letter.Datetime,
			From:       userLetter.Letter.From,
			ToSubditID: int(userLetter.Letter.ToSubditID),
			IsArchived: userLetter.IsArchived,
			CreatedAt:  userLetter.Letter.CreatedAt,
			UpdatedAt:  userLetter.Letter.UpdatedAt,
		}

		userLettersResponse = append(userLettersResponse, userLetterResponse)
	}

	response := dto.GetUserLettersByToSubditIDResponse{
		Status:  http.StatusOK,
		Letters: userLettersResponse,
	}

	return &response, nil
}
