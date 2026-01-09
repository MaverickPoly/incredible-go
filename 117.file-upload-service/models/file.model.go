package models

import "gorm.io/gorm"

type File struct {
	gorm.Model

	UserId   uint   `json:"user_id" gorm:"not null"`
	Path     string `json:"path" gorm:"not null;unique"`
	Filename string `json:"filename" gorm:"not null"`
	Size     int64  `json:"size" gorm:"not null"` // in bytes
}
