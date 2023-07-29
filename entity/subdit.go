package entity

import "gorm.io/gorm"

type Subdit struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}
