package models

import (
	"github.com/jinzhu/gorm"
)

type Image struct {
	gorm.Model
	Name string `gorm:"not null"`
	Data []byte `gorm:"type:blob;not null"`
}
