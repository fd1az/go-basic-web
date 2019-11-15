package models

import (
	"github.com/jinzhu/gorm"
	//_
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//User model
type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null; unique_index"`
}
