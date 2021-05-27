package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Name string `json:"name"`
	Tel  string `json:"tel" gorm:"uniqueIndex;index"`
}
