package dto

type UserLetterAssignLetterResponse struct {
	Status     int  `json:"status"`
	LetterID   int  `json:"letter_id"`
	UserID     int  `json:"user_id"`
	IsArchived bool `json:"is_archived"`
}

type UserLetterArchiveResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type UserLetterDeleteResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
