package entity

import (
	"time"

	"gorm.io/gorm"
)

type Letter struct {
	gorm.Model
	Name       string    `json:"name"`
	About      string    `json:"about"`
	Number     string    `json:"number" gorm:"unique"`
	Datetime   time.Time `json:"datetime"`
	From       string    `json:"from"`
	Type       string    `json:"type"`
	ToSubditID uint      `json:"to_subdit_id"`
	IsNotified bool      `json:"is_notified" gorm:"default:false"`
}
