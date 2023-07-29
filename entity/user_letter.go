package entity

import "gorm.io/gorm"

type UserLetter struct {
	gorm.Model
	UserID     uint   `json:"user_id" gorm:"primaryKey;autoIncrement:false"`
	LetterID   uint   `json:"letter_id" gorm:"primaryKey;autoIncrement:false"`
	IsArchived bool   `json:"is_archived" gorm:"default:false"`
	Letter     Letter `json:"letter"`
	User       User   `json:"user"`
}
