package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string `json:"name"`
	Email       string `json:"email" gorm:"unique"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role" gorm:"default:staff"`
	SubditID    uint   `json:"subdit_id"`
}
